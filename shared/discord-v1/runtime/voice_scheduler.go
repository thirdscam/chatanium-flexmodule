package runtime

import (
	"context"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
)

// QueueScheduler processes tasks from the queue and sends to VoiceConnection
type QueueScheduler struct {
	queue *VoiceTaskQueue
	vc    *discordgo.VoiceConnection
	log   hclog.Logger

	ctx    context.Context
	cancel context.CancelFunc

	// Statistics
	mu             sync.RWMutex
	tasksProcessed map[string]uint64 // module_id -> count
	bytesSent      map[string]uint64 // module_id -> bytes
}

// NewQueueScheduler creates a new queue scheduler
func NewQueueScheduler(queue *VoiceTaskQueue, vc *discordgo.VoiceConnection, ctx context.Context, log hclog.Logger) *QueueScheduler {
	ctx, cancel := context.WithCancel(ctx)

	return &QueueScheduler{
		queue:          queue,
		vc:             vc,
		log:            log,
		ctx:            ctx,
		cancel:         cancel,
		tasksProcessed: make(map[string]uint64),
		bytesSent:      make(map[string]uint64),
	}
}

// Start begins processing tasks from the queue
func (s *QueueScheduler) Start() {
	go s.processLoop()
	go s.cleanupLoop()
}

// Stop stops the scheduler
func (s *QueueScheduler) Stop() {
	s.cancel()
}

// processLoop continuously processes tasks from the queue
func (s *QueueScheduler) processLoop() {
	s.log.Info("Queue scheduler started")

	for {
		select {
		case <-s.ctx.Done():
			s.log.Info("Queue scheduler stopped")
			return

		default:
		}

		// Dequeue next task (blocking)
		task, err := s.queue.Dequeue()
		if err != nil {
			s.log.Debug("Queue dequeue error", "error", err)
			continue
		}

		// Check timeout
		if task.Timeout > 0 && time.Since(task.SubmitTime) > task.Timeout {
			s.log.Warn("Task timed out", "task_id", task.TaskID, "module_id", task.ModuleID)
			continue
		}

		// Send to VoiceConnection
		select {
		case s.vc.OpusSend <- task.OpusData:
			// Success
			s.mu.Lock()
			s.tasksProcessed[task.ModuleID]++
			s.bytesSent[task.ModuleID] += uint64(len(task.OpusData))
			s.mu.Unlock()

		case <-time.After(100 * time.Millisecond):
			s.log.Warn("Failed to send task (channel full)", "task_id", task.TaskID)

		case <-s.ctx.Done():
			return
		}

		// Maintain frame interval (20ms per Opus frame)
		time.Sleep(20 * time.Millisecond)
	}
}

// cleanupLoop periodically cleans expired tasks
func (s *QueueScheduler) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			expiredCount := s.queue.CleanExpired()
			if expiredCount > 0 {
				s.log.Debug("Cleaned expired tasks", "count", expiredCount)
			}

		case <-s.ctx.Done():
			return
		}
	}
}

// GetStats returns scheduler statistics
func (s *QueueScheduler) GetStats() (tasksProcessed, bytesSent map[string]uint64) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Copy maps
	tasksCopy := make(map[string]uint64)
	bytesCopy := make(map[string]uint64)

	for k, v := range s.tasksProcessed {
		tasksCopy[k] = v
	}
	for k, v := range s.bytesSent {
		bytesCopy[k] = v
	}

	return tasksCopy, bytesCopy
}
