package runtime

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// VoiceTask represents a voice task in the queue
type VoiceTask struct {
	TaskID     string
	ModuleID   string
	OpusData   []byte
	Priority   int
	SubmitTime time.Time
	Timeout    time.Duration
}

// VoiceTaskQueue manages voice tasks with priority support
type VoiceTaskQueue struct {
	mu     sync.RWMutex
	tasks  []*VoiceTask
	cond   *sync.Cond
	closed bool

	// Statistics
	tasksEnqueued uint64
	tasksDequeued uint64
	tasksExpired  uint64
}

// NewVoiceTaskQueue creates a new voice task queue
func NewVoiceTaskQueue() *VoiceTaskQueue {
	q := &VoiceTaskQueue{
		tasks: make([]*VoiceTask, 0),
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Enqueue adds a task to the queue
func (q *VoiceTaskQueue) Enqueue(task *VoiceTask) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return fmt.Errorf("queue closed")
	}

	q.tasks = append(q.tasks, task)
	q.tasksEnqueued++

	// Sort by priority (higher priority first)
	sort.Slice(q.tasks, func(i, j int) bool {
		if q.tasks[i].Priority != q.tasks[j].Priority {
			return q.tasks[i].Priority > q.tasks[j].Priority
		}
		return q.tasks[i].SubmitTime.Before(q.tasks[j].SubmitTime)
	})

	q.cond.Signal()
	return nil
}

// Dequeue removes and returns the next task (blocking)
func (q *VoiceTaskQueue) Dequeue() (*VoiceTask, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.tasks) == 0 && !q.closed {
		q.cond.Wait()
	}

	if q.closed && len(q.tasks) == 0 {
		return nil, fmt.Errorf("queue closed")
	}

	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	q.tasksDequeued++

	return task, nil
}

// Length returns the current queue length
func (q *VoiceTaskQueue) Length() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return len(q.tasks)
}

// Close closes the queue
func (q *VoiceTaskQueue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.closed {
		q.closed = true
		q.cond.Broadcast()
	}
}

// CleanExpired removes expired tasks
func (q *VoiceTaskQueue) CleanExpired() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	now := time.Now()
	validTasks := make([]*VoiceTask, 0, len(q.tasks))
	expiredCount := 0

	for _, task := range q.tasks {
		if task.Timeout > 0 && now.Sub(task.SubmitTime) > task.Timeout {
			expiredCount++
			q.tasksExpired++
		} else {
			validTasks = append(validTasks, task)
		}
	}

	q.tasks = validTasks
	return expiredCount
}

// Stats returns queue statistics
func (q *VoiceTaskQueue) Stats() (enqueued, dequeued, expired uint64) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.tasksEnqueued, q.tasksDequeued, q.tasksExpired
}
