package runtime

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/thirdscam/chatanium-flexmodule/proto"
	pb "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// VoiceHelper implements the VoiceStream gRPC service
type VoiceHelper struct {
	pb.UnimplementedVoiceStreamServer

	dgSession *discordgo.Session
	log       hclog.Logger

	// Guild별 음성 세션 관리
	mu            sync.RWMutex
	voiceSessions map[string]*VoiceSession // guild_id -> session
}

// VoiceSession represents a voice session for a guild
type VoiceSession struct {
	GuildID   string
	ChannelID string

	vc          *discordgo.VoiceConnection
	queue       *VoiceTaskQueue
	scheduler   *QueueScheduler
	broadcaster *ReceiveBroadcaster

	// 모듈 구독 관리
	mu          sync.RWMutex
	subscribers map[string]*ModuleSubscription // module_id -> subscription

	ctx    context.Context
	cancel context.CancelFunc
}

// ModuleSubscription represents a module's subscription to a voice session
type ModuleSubscription struct {
	ModuleID     string
	ConnectionID string
	SubmitChan   chan *VoiceTask
	RecvChan     chan *pb.VoicePacket
	Stream       pb.VoiceStream_VoiceStreamServer

	CreatedAt    time.Time
	LastActivity time.Time

	// Statistics
	mu             sync.RWMutex
	tasksSubmitted uint64
	bytesReceived  uint64
}

// NewVoiceHelper creates a new voice helper
func NewVoiceHelper(session *discordgo.Session, log hclog.Logger) *VoiceHelper {
	return &VoiceHelper{
		dgSession:     session,
		log:           log.Named("voice-helper"),
		voiceSessions: make(map[string]*VoiceSession),
	}
}

// VoiceJoin implements VoiceStream.VoiceJoin
func (h *VoiceHelper) VoiceJoin(ctx context.Context, req *pb.VoiceJoinRequest) (*pb.VoiceJoinResponse, error) {
	moduleID, err := getModuleIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "module ID not found")
	}

	h.log.Info("Voice join request", "module_id", moduleID, "guild_id", req.GuildId, "channel_id", req.ChannelId)

	h.mu.Lock()
	defer h.mu.Unlock()

	// Get or create voice session
	session, exists := h.voiceSessions[req.GuildId]
	if !exists || session.ChannelID != req.ChannelId {
		// Create new session
		if exists {
			// Moving to different channel, close existing session
			h.log.Info("Moving voice session to different channel", "guild_id", req.GuildId, "old_channel", session.ChannelID, "new_channel", req.ChannelId)
			session.cancel()
		}

		session, err = h.createVoiceSession(req.GuildId, req.ChannelId, req.Mute, req.Deaf)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create voice session: %v", err)
		}

		h.voiceSessions[req.GuildId] = session
	}

	// Create module subscription
	bufferSize := req.BufferSize
	if bufferSize <= 0 {
		bufferSize = 100
	}

	subscription := &ModuleSubscription{
		ModuleID:     moduleID,
		ConnectionID: generateConnectionID(),
		SubmitChan:   make(chan *VoiceTask, 100),
		RecvChan:     make(chan *pb.VoicePacket, bufferSize),
		CreatedAt:    time.Now(),
		LastActivity: time.Now(),
	}

	// Register subscription
	session.mu.Lock()
	session.subscribers[moduleID] = subscription
	session.mu.Unlock()

	// Subscribe to broadcaster
	session.broadcaster.Subscribe(moduleID, subscription.RecvChan, nil)

	// Start task submission goroutine
	go h.submitTaskLoop(session, subscription)

	h.log.Info("Module joined voice session", "module_id", moduleID, "connection_id", subscription.ConnectionID)

	return &pb.VoiceJoinResponse{
		ConnectionId:      subscription.ConnectionID,
		GuildId:           req.GuildId,
		ChannelId:         req.ChannelId,
		Ready:             session.vc.Ready,
		QueueLength:       int32(session.queue.Length()),
		ActiveSubscribers: int32(len(session.subscribers)),
	}, nil
}

// VoiceStream implements VoiceStream.VoiceStream
func (h *VoiceHelper) VoiceStream(stream pb.VoiceStream_VoiceStreamServer) error {
	// Get module ID
	moduleID, err := getModuleIDFromContext(stream.Context())
	if err != nil {
		return status.Error(codes.Unauthenticated, "module ID not found")
	}

	// Get first packet to obtain connection_id
	firstPacket, err := stream.Recv()
	if err != nil {
		return err
	}
	connectionID := firstPacket.ConnectionId

	h.log.Info("Voice stream started", "module_id", moduleID, "connection_id", connectionID)

	// Find subscription
	subscription, session, err := h.findSubscription(connectionID, moduleID)
	if err != nil {
		return status.Errorf(codes.NotFound, "subscription not found: %v", err)
	}

	// Store stream
	subscription.Stream = stream

	// Error channel
	errChan := make(chan error, 2)

	// Send goroutine (module -> queue)
	go func() {
		// Process first packet
		if len(firstPacket.OpusData) > 0 {
			task := &VoiceTask{
				TaskID:     generateTaskID(),
				ModuleID:   moduleID,
				OpusData:   firstPacket.OpusData,
				Priority:   int(firstPacket.Priority),
				SubmitTime: time.Now(),
				Timeout:    time.Duration(firstPacket.TimeoutMs) * time.Millisecond,
			}

			select {
			case subscription.SubmitChan <- task:
				subscription.mu.Lock()
				subscription.tasksSubmitted++
				subscription.mu.Unlock()
			case <-stream.Context().Done():
				errChan <- stream.Context().Err()
				return
			}
		}

		// Process subsequent packets
		for {
			packet, err := stream.Recv()
			if err != nil {
				errChan <- err
				return
			}

			task := &VoiceTask{
				TaskID:     generateTaskID(),
				ModuleID:   moduleID,
				OpusData:   packet.OpusData,
				Priority:   int(packet.Priority),
				SubmitTime: time.Now(),
				Timeout:    time.Duration(packet.TimeoutMs) * time.Millisecond,
			}

			select {
			case subscription.SubmitChan <- task:
				subscription.LastActivity = time.Now()
				subscription.mu.Lock()
				subscription.tasksSubmitted++
				subscription.mu.Unlock()

			case <-stream.Context().Done():
				errChan <- stream.Context().Err()
				return
			}
		}
	}()

	// Receive goroutine (broadcaster -> module)
	go func() {
		for {
			select {
			case packet, ok := <-subscription.RecvChan:
				if !ok {
					errChan <- nil
					return
				}

				packet.ConnectionId = connectionID

				err := stream.Send(packet)
				if err != nil {
					errChan <- err
					return
				}

				subscription.LastActivity = time.Now()
				subscription.mu.Lock()
				subscription.bytesReceived += uint64(len(packet.OpusData))
				subscription.mu.Unlock()

			case <-stream.Context().Done():
				errChan <- stream.Context().Err()
				return
			}
		}
	}()

	// Wait for error
	err = <-errChan

	// Cleanup
	h.cleanupSubscription(subscription, session)

	h.log.Info("Voice stream ended", "module_id", moduleID, "connection_id", connectionID)

	return err
}

// VoiceLeave implements VoiceStream.VoiceLeave
func (h *VoiceHelper) VoiceLeave(ctx context.Context, req *pb.VoiceLeaveRequest) (*proto.Empty, error) {
	moduleID, err := getModuleIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "module ID not found")
	}

	h.log.Info("Voice leave request", "module_id", moduleID, "connection_id", req.ConnectionId)

	subscription, session, err := h.findSubscription(req.ConnectionId, moduleID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "subscription not found: %v", err)
	}

	h.cleanupSubscription(subscription, session)

	// Check if session should be closed
	session.mu.RLock()
	subscriberCount := len(session.subscribers)
	session.mu.RUnlock()

	if subscriberCount == 0 {
		h.log.Info("Last subscriber left, closing voice session", "guild_id", session.GuildID)
		h.mu.Lock()
		delete(h.voiceSessions, session.GuildID)
		h.mu.Unlock()

		session.cancel()
		session.vc.Disconnect()
	}

	return &proto.Empty{}, nil
}

// VoiceSpeaking implements VoiceStream.VoiceSpeaking
func (h *VoiceHelper) VoiceSpeaking(ctx context.Context, req *pb.VoiceSpeakingRequest) (*proto.Empty, error) {
	moduleID, err := getModuleIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "module ID not found")
	}

	_, session, err := h.findSubscription(req.ConnectionId, moduleID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "subscription not found: %v", err)
	}

	err = session.vc.Speaking(req.Speaking)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set speaking state: %v", err)
	}

	h.log.Debug("Speaking state changed", "module_id", moduleID, "speaking", req.Speaking)

	return &proto.Empty{}, nil
}

// GetQueueStatus implements VoiceStream.GetQueueStatus
func (h *VoiceHelper) GetQueueStatus(ctx context.Context, req *pb.QueueStatusRequest) (*pb.QueueStatusResponse, error) {
	h.mu.RLock()
	session, exists := h.voiceSessions[req.GuildId]
	h.mu.RUnlock()

	if !exists {
		return nil, status.Error(codes.NotFound, "voice session not found")
	}

	// Get queue stats
	queueLength := session.queue.Length()

	// Get subscriber count
	session.mu.RLock()
	subscriberCount := len(session.subscribers)
	session.mu.RUnlock()

	// Get module stats
	tasksProcessed, bytesSent := session.scheduler.GetStats()

	moduleStats := make([]*pb.ModuleStats, 0)
	session.mu.RLock()
	for moduleID, sub := range session.subscribers {
		sub.mu.RLock()
		stats := &pb.ModuleStats{
			ModuleId:        moduleID,
			TasksSubmitted:  sub.tasksSubmitted,
			TasksProcessed:  tasksProcessed[moduleID],
			BytesSent:       bytesSent[moduleID],
			BytesReceived:   sub.bytesReceived,
		}
		sub.mu.RUnlock()
		moduleStats = append(moduleStats, stats)
	}
	session.mu.RUnlock()

	return &pb.QueueStatusResponse{
		QueueLength:       int32(queueLength),
		ActiveSubscribers: int32(subscriberCount),
		ModuleStats:       moduleStats,
	}, nil
}

// Helper functions

func (h *VoiceHelper) createVoiceSession(guildID, channelID string, mute, deaf bool) (*VoiceSession, error) {
	// Join voice channel
	vc, err := h.dgSession.ChannelVoiceJoin(guildID, channelID, mute, deaf)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create queue system
	queue := NewVoiceTaskQueue()
	scheduler := NewQueueScheduler(queue, vc, ctx, h.log.Named("scheduler"))
	broadcaster := NewReceiveBroadcaster(vc, ctx, h.log.Named("broadcaster"))

	session := &VoiceSession{
		GuildID:     guildID,
		ChannelID:   channelID,
		vc:          vc,
		queue:       queue,
		scheduler:   scheduler,
		broadcaster: broadcaster,
		subscribers: make(map[string]*ModuleSubscription),
		ctx:         ctx,
		cancel:      cancel,
	}

	// Start queue and broadcaster
	scheduler.Start()
	broadcaster.Start()

	h.log.Info("Voice session created", "guild_id", guildID, "channel_id", channelID)

	return session, nil
}

func (h *VoiceHelper) submitTaskLoop(session *VoiceSession, subscription *ModuleSubscription) {
	for {
		select {
		case task := <-subscription.SubmitChan:
			err := session.queue.Enqueue(task)
			if err != nil {
				h.log.Error("Failed to enqueue task", "error", err, "module_id", subscription.ModuleID)
			}

		case <-session.ctx.Done():
			return
		}
	}
}

func (h *VoiceHelper) findSubscription(connectionID, moduleID string) (*ModuleSubscription, *VoiceSession, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, session := range h.voiceSessions {
		session.mu.RLock()
		subscription, exists := session.subscribers[moduleID]
		session.mu.RUnlock()

		if exists && subscription.ConnectionID == connectionID {
			return subscription, session, nil
		}
	}

	return nil, nil, fmt.Errorf("subscription not found")
}

func (h *VoiceHelper) cleanupSubscription(subscription *ModuleSubscription, session *VoiceSession) {
	// Remove from session
	session.mu.Lock()
	delete(session.subscribers, subscription.ModuleID)
	session.mu.Unlock()

	// Unsubscribe from broadcaster
	session.broadcaster.Unsubscribe(subscription.ModuleID)

	// Close channels
	close(subscription.SubmitChan)
	close(subscription.RecvChan)

	h.log.Info("Subscription cleaned up", "module_id", subscription.ModuleID, "connection_id", subscription.ConnectionID)
}

func getModuleIDFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("no metadata in context")
	}

	moduleIDs := md.Get("module-id")
	if len(moduleIDs) == 0 {
		return "", fmt.Errorf("module-id not found in metadata")
	}

	return moduleIDs[0], nil
}

func generateConnectionID() string {
	return "voice_" + uuid.New().String()
}

func generateTaskID() string {
	return "task_" + uuid.New().String()
}
