package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TestStatus represents the status of a test
type TestStatus string

const (
	// TestStatusPending indicates the test is pending
	TestStatusPending TestStatus = "pending"
	// TestStatusRunning indicates the test is running
	TestStatusRunning TestStatus = "running"
	// TestStatusCompleted indicates the test completed successfully
	TestStatusCompleted TestStatus = "completed"
	// TestStatusFailed indicates the test failed
	TestStatusFailed TestStatus = "failed"
	// TestStatusCancelled indicates the test was cancelled
	TestStatusCancelled TestStatus = "cancelled"
)

// TestType represents the type of a test
type TestType string

const (
	// TestTypeValidation indicates a documentation validation test
	TestTypeValidation TestType = "validation"
	// TestTypeLoadTest indicates a load test
	TestTypeLoadTest TestType = "load_test"
	// TestTypeAcceptance indicates an acceptance test
	TestTypeAcceptance TestType = "acceptance"
)

// TestBase contains common fields for all test types
type TestBase struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      TestStatus `json:"status"`
	Type        TestType   `json:"type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	UserID      string     `json:"user_id"`
	Tags        []string   `json:"tags"`
}

// NewTestBase creates a new TestBase with default values
func NewTestBase(testType TestType, name, description string) TestBase {
	now := time.Now()
	return TestBase{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      TestStatusPending,
		Type:        testType,
		CreatedAt:   now,
		UpdatedAt:   now,
		Tags:        []string{},
	}
}

// GitHubIssueRequest represents a request to create a GitHub issue
type GitHubIssueRequest struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Labels     []string `json:"labels"`
}

// GitHubIssueResponse represents the response from creating a GitHub issue
type GitHubIssueResponse struct {
	IssueNumber int    `json:"issue_number"`
	IssueURL    string `json:"issue_url"`
}

// TestResult is an interface that all test result types must implement
type TestResult interface {
	GetStatus() TestStatus
	IsSuccessful() bool
	GetSummary() string
	GetTestID() string
}

// BaseTestResult contains common fields for all test result types
type BaseTestResult struct {
	TestID      string     `json:"test_id"`
	Status      TestStatus `json:"status"`
	StartTime   time.Time  `json:"start_time"`
	EndTime     time.Time  `json:"end_time"`
	Duration    string     `json:"duration"`
	ErrorDetail string     `json:"error_detail,omitempty"`
}

// GetStatus returns the status of the test
func (r *BaseTestResult) GetStatus() TestStatus {
	return r.Status
}

// IsSuccessful returns true if the test was successful
func (r *BaseTestResult) IsSuccessful() bool {
	return r.Status == TestStatusCompleted
}

// GetTestID returns the ID of the test
func (r *BaseTestResult) GetTestID() string {
	return r.TestID
}

// GetSummary returns a summary of the test result
func (r *BaseTestResult) GetSummary() string {
	if r.IsSuccessful() {
		return "Test completed successfully"
	}
	return "Test failed: " + r.ErrorDetail
}

// QueueTask represents a task that can be queued
type QueueTask struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	CreatedAt time.Time   `json:"created_at"`
	Attempts  int         `json:"attempts"`
}

// NewQueueTask creates a new QueueTask
func NewQueueTask(taskType string, payload interface{}) QueueTask {
	return QueueTask{
		ID:        uuid.New().String(),
		Type:      taskType,
		Payload:   payload,
		CreatedAt: time.Now(),
		Attempts:  0,
	}
}

// UnmarshalPayload unmarshals the payload of a QueueTask into the provided interface
func (task *QueueTask) UnmarshalPayload(v interface{}) error {
	payloadBytes, err := json.Marshal(task.Payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}
	
	if err := json.Unmarshal(payloadBytes, v); err != nil {
		return fmt.Errorf("failed to unmarshal task payload: %w", err)
	}
	
	return nil
}