package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/models"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// QueueService defines the interface for a task queue
type QueueService interface {
	// Enqueue adds a task to the queue
	Enqueue(ctx context.Context, taskType string, payload interface{}) (string, error)
	
	// Dequeue gets a task from the queue
	Dequeue(ctx context.Context, taskTypes []string) (*models.QueueTask, error)
	
	// Complete marks a task as completed
	Complete(ctx context.Context, taskID string) error
	
	// Fail marks a task as failed with an error message
	Fail(ctx context.Context, taskID string, err error) error
	
	// Retry requeues a task with a backoff delay
	Retry(ctx context.Context, task *models.QueueTask) error
	
	// RegisterHandler registers a handler for a task type
	RegisterHandler(taskType string, handler TaskHandler)
	
	// StartWorkers starts the worker goroutines
	StartWorkers(ctx context.Context, workerCount int) error
	
	// Close closes the queue client connection
	Close() error
}

// TaskHandler is a function that processes a task
type TaskHandler func(ctx context.Context, task *models.QueueTask) error

// RedisQueue implements QueueService using Redis
type RedisQueue struct {
	client           *redis.Client
	logger           *logrus.Logger
	handlers         map[string]TaskHandler
	pendingQueueName string
	processingPrefix string
	completedPrefix  string
	failedPrefix     string
}

// NewRedisQueue creates a new Redis queue client
func NewRedisQueue(ctx context.Context, config config.RedisConfig) (*RedisQueue, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	// Verify connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisQueue{
		client:           client,
		logger:           logrus.New(),
		handlers:         make(map[string]TaskHandler),
		pendingQueueName: "driveby:queue:pending",
		processingPrefix: "driveby:queue:processing:",
		completedPrefix:  "driveby:queue:completed:",
		failedPrefix:     "driveby:queue:failed:",
	}, nil
}

// SetLogger sets the logger for the queue
func (q *RedisQueue) SetLogger(logger *logrus.Logger) {
	q.logger = logger
}

// Enqueue adds a task to the queue
func (q *RedisQueue) Enqueue(ctx context.Context, taskType string, payload interface{}) (string, error) {
	task := models.NewQueueTask(taskType, payload)
	
	// Serialize task
	taskData, err := json.Marshal(task)
	if err != nil {
		return "", fmt.Errorf("failed to serialize task: %w", err)
	}

	// Add to pending queue
	err = q.client.LPush(ctx, q.pendingQueueName, taskData).Err()
	if err != nil {
		return "", fmt.Errorf("failed to enqueue task: %w", err)
	}

	q.logger.WithFields(logrus.Fields{
		"task_id":   task.ID,
		"task_type": task.Type,
	}).Info("Task enqueued")

	return task.ID, nil
}

// Dequeue gets a task from the queue
func (q *RedisQueue) Dequeue(ctx context.Context, taskTypes []string) (*models.QueueTask, error) {
	// Get a task from the pending queue
	result, err := q.client.BRPop(ctx, 1*time.Second, q.pendingQueueName).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // No tasks available
		}
		return nil, fmt.Errorf("failed to dequeue task: %w", err)
	}

	// Unmarshal task
	var task models.QueueTask
	err = json.Unmarshal([]byte(result[1]), &task)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize task: %w", err)
	}

	// Check if task type is in the list of accepted types
	if len(taskTypes) > 0 {
		typeMatch := false
		for _, t := range taskTypes {
			if task.Type == t {
				typeMatch = true
				break
			}
		}
		if !typeMatch {
			// Put the task back into the queue
			err = q.client.LPush(ctx, q.pendingQueueName, result[1]).Err()
			if err != nil {
				q.logger.WithError(err).Error("Failed to put task back into queue")
			}
			return nil, nil // No tasks of the requested type
		}
	}

	// Move to processing queue
	err = q.client.Set(ctx, q.processingPrefix+task.ID, result[1], 24*time.Hour).Err()
	if err != nil {
		q.logger.WithError(err).Error("Failed to mark task as processing")
	}

	q.logger.WithFields(logrus.Fields{
		"task_id":   task.ID,
		"task_type": task.Type,
	}).Info("Task dequeued")

	return &task, nil
}

// Complete marks a task as completed
func (q *RedisQueue) Complete(ctx context.Context, taskID string) error {
	// Get task from processing queue
	taskData, err := q.client.Get(ctx, q.processingPrefix+taskID).Result()
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Move to completed queue
	err = q.client.Set(ctx, q.completedPrefix+taskID, taskData, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to mark task as completed: %w", err)
	}

	// Remove from processing queue
	err = q.client.Del(ctx, q.processingPrefix+taskID).Err()
	if err != nil {
		q.logger.WithError(err).Error("Failed to remove task from processing queue")
	}

	q.logger.WithField("task_id", taskID).Info("Task completed")
	return nil
}

// Fail marks a task as failed with an error message
func (q *RedisQueue) Fail(ctx context.Context, taskID string, taskErr error) error {
	// Get task from processing queue
	taskDataStr, err := q.client.Get(ctx, q.processingPrefix+taskID).Result()
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	var task models.QueueTask
	err = json.Unmarshal([]byte(taskDataStr), &task)
	if err != nil {
		return fmt.Errorf("failed to deserialize task: %w", err)
	}

	// Add error information
	task.Attempts++
	
	// Serialize updated task
	taskData, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to serialize task: %w", err)
	}

	// Move to failed queue
	err = q.client.Set(ctx, q.failedPrefix+taskID, taskData, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to mark task as failed: %w", err)
	}

	// Remove from processing queue
	err = q.client.Del(ctx, q.processingPrefix+taskID).Err()
	if err != nil {
		q.logger.WithError(err).Error("Failed to remove task from processing queue")
	}

	q.logger.WithFields(logrus.Fields{
		"task_id": taskID,
		"error":   taskErr.Error(),
	}).Error("Task failed")
	return nil
}

// Retry requeues a task with a backoff delay
func (q *RedisQueue) Retry(ctx context.Context, task *models.QueueTask) error {
	task.Attempts++
	
	// Serialize task
	taskData, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to serialize task: %w", err)
	}

	// Calculate backoff (exponential backoff with jitter)
	backoff := time.Duration(1<<task.Attempts) * time.Second
	if backoff > 1*time.Hour {
		backoff = 1 * time.Hour
	}

	// Add to pending queue after backoff
	err = q.client.LPush(ctx, q.pendingQueueName, taskData).Err()
	if err != nil {
		return fmt.Errorf("failed to requeue task: %w", err)
	}

	// Remove from processing queue
	err = q.client.Del(ctx, q.processingPrefix+task.ID).Err()
	if err != nil {
		q.logger.WithError(err).Error("Failed to remove task from processing queue")
	}

	q.logger.WithFields(logrus.Fields{
		"task_id":  task.ID,
		"attempts": task.Attempts,
		"backoff":  backoff.String(),
	}).Info("Task requeued for retry")
	return nil
}

// RegisterHandler registers a handler for a task type
func (q *RedisQueue) RegisterHandler(taskType string, handler TaskHandler) {
	q.handlers[taskType] = handler
	q.logger.WithField("task_type", taskType).Info("Registered task handler")
}

// StartWorkers starts the worker goroutines
func (q *RedisQueue) StartWorkers(ctx context.Context, workerCount int) error {
	if len(q.handlers) == 0 {
		return fmt.Errorf("no task handlers registered")
	}

	// Get task types from handlers
	taskTypes := make([]string, 0, len(q.handlers))
	for t := range q.handlers {
		taskTypes = append(taskTypes, t)
	}

	q.logger.WithFields(logrus.Fields{
		"worker_count": workerCount,
		"task_types":   taskTypes,
	}).Info("Starting queue workers")

	// Start workers
	for i := 0; i < workerCount; i++ {
		workerID := i
		go func() {
			q.runWorker(ctx, workerID, taskTypes)
		}()
	}

	return nil
}

// runWorker runs a worker goroutine
func (q *RedisQueue) runWorker(ctx context.Context, workerID int, taskTypes []string) {
	logger := q.logger.WithField("worker_id", workerID)
	logger.Info("Worker started")

	for {
		select {
		case <-ctx.Done():
			logger.Info("Worker stopped")
			return
		default:
			task, err := q.Dequeue(ctx, taskTypes)
			if err != nil {
				logger.WithError(err).Error("Failed to dequeue task")
				time.Sleep(1 * time.Second)
				continue
			}

			if task == nil {
				// No tasks available, wait a bit
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// Get handler for task type
			handler, ok := q.handlers[task.Type]
			if !ok {
				logger.WithField("task_type", task.Type).Error("No handler registered for task type")
				_ = q.Fail(ctx, task.ID, fmt.Errorf("no handler for task type %s", task.Type))
				continue
			}

			// Process task
			logger.WithFields(logrus.Fields{
				"task_id":   task.ID,
				"task_type": task.Type,
			}).Info("Processing task")

			err = handler(ctx, task)
			if err != nil {
				logger.WithError(err).Error("Failed to process task")
				_ = q.Fail(ctx, task.ID, err)
				continue
			}

			// Mark task as completed
			err = q.Complete(ctx, task.ID)
			if err != nil {
				logger.WithError(err).Error("Failed to complete task")
			}
		}
	}
}

// Close closes the queue client connection
func (q *RedisQueue) Close() error {
	return q.client.Close()
}