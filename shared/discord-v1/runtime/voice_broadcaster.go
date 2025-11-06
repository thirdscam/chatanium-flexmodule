package runtime

import (
	"context"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	pb "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// VoiceReceiver represents a module subscribed to receive voice data
type VoiceReceiver struct {
	ModuleID string
	RecvChan chan *pb.VoicePacket
	Filter   func(*discordgo.Packet) bool // Optional filter
}

// ReceiveBroadcaster broadcasts received voice packets to all subscribers
type ReceiveBroadcaster struct {
	vc  *discordgo.VoiceConnection
	log hclog.Logger

	mu        sync.RWMutex
	receivers map[string]*VoiceReceiver // module_id -> receiver

	ctx    context.Context
	cancel context.CancelFunc

	// Statistics
	packetsReceived uint64
	packetsDropped  map[string]uint64 // module_id -> dropped count
}

// NewReceiveBroadcaster creates a new receive broadcaster
func NewReceiveBroadcaster(vc *discordgo.VoiceConnection, ctx context.Context, log hclog.Logger) *ReceiveBroadcaster {
	ctx, cancel := context.WithCancel(ctx)

	return &ReceiveBroadcaster{
		vc:              vc,
		log:             log,
		receivers:       make(map[string]*VoiceReceiver),
		ctx:             ctx,
		cancel:          cancel,
		packetsDropped:  make(map[string]uint64),
	}
}

// Start begins broadcasting received voice packets
func (b *ReceiveBroadcaster) Start() {
	go b.broadcastLoop()
}

// Stop stops the broadcaster
func (b *ReceiveBroadcaster) Stop() {
	b.cancel()
}

// broadcastLoop continuously broadcasts received packets
func (b *ReceiveBroadcaster) broadcastLoop() {
	b.log.Info("Receive broadcaster started")

	for {
		select {
		case <-b.ctx.Done():
			b.log.Info("Receive broadcaster stopped")
			return

		case packet, ok := <-b.vc.OpusRecv:
			if !ok {
				b.log.Warn("OpusRecv channel closed")
				return
			}

			b.packetsReceived++

			// Convert to protobuf packet
			pbPacket := &pb.VoicePacket{
				OpusData:  packet.Opus,
				Ssrc:      packet.SSRC,
				Sequence:  uint32(packet.Sequence),
				Timestamp: packet.Timestamp,
				Type:      packet.Type,
			}

			// Broadcast to all subscribers
			b.mu.RLock()
			for moduleID, receiver := range b.receivers {
				// Apply filter if exists
				if receiver.Filter != nil && !receiver.Filter(packet) {
					continue
				}

				// Non-blocking send
				select {
				case receiver.RecvChan <- pbPacket:
					// Success
				default:
					// Channel full, drop packet
					b.mu.RUnlock()
					b.mu.Lock()
					b.packetsDropped[moduleID]++
					b.mu.Unlock()
					b.mu.RLock()

					b.log.Debug("Dropped packet for module (channel full)", "module_id", moduleID)
				}
			}
			b.mu.RUnlock()
		}
	}
}

// Subscribe adds a new receiver
func (b *ReceiveBroadcaster) Subscribe(moduleID string, recvChan chan *pb.VoicePacket, filter func(*discordgo.Packet) bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.receivers[moduleID] = &VoiceReceiver{
		ModuleID: moduleID,
		RecvChan: recvChan,
		Filter:   filter,
	}

	b.log.Info("Module subscribed to voice receive", "module_id", moduleID)
}

// Unsubscribe removes a receiver
func (b *ReceiveBroadcaster) Unsubscribe(moduleID string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.receivers, moduleID)
	delete(b.packetsDropped, moduleID)

	b.log.Info("Module unsubscribed from voice receive", "module_id", moduleID)
}

// GetSubscriberCount returns the number of active subscribers
func (b *ReceiveBroadcaster) GetSubscriberCount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.receivers)
}

// GetStats returns broadcaster statistics
func (b *ReceiveBroadcaster) GetStats() (packetsReceived uint64, packetsDropped map[string]uint64) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	// Copy map
	droppedCopy := make(map[string]uint64)
	for k, v := range b.packetsDropped {
		droppedCopy[k] = v
	}

	return b.packetsReceived, droppedCopy
}
