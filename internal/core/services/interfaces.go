package services

import (
	"context"

	"driveby/internal/core/models"
)

// ValidationService defines operations for documentation validation
type ValidationService interface {
	// ValidateOpenAPI performs validation on an OpenAPI specification
	ValidateOpenAPI(ctx context.Context, test *models.ValidationTest) (*models.ValidationResult, error)

	// GetValidationTest retrieves a validation test by ID
	GetValidationTest(ctx context.Context, testID string) (*models.ValidationTest, error)

	// ListValidationTests retrieves all validation tests
	ListValidationTests(ctx context.Context) ([]*models.ValidationTest, error)

	// QueueValidationTest queues a validation test for asynchronous processing
	QueueValidationTest(ctx context.Context, test *models.ValidationTest) error

	// GenerateReport creates a validation report for a completed test
	GenerateReport(ctx context.Context, testID string) (string, error)
}

// LoadTestService defines operations for load testing
type LoadTestService interface {
	// RunLoadTest performs a load test against a target
	RunLoadTest(ctx context.Context, test *models.LoadTest) (*models.LoadTestResult, error)

	// GetLoadTest retrieves a load test by ID
	GetLoadTest(ctx context.Context, testID string) (*models.LoadTest, error)

	// ListLoadTests retrieves all load tests
	ListLoadTests(ctx context.Context) ([]*models.LoadTest, error)

	// QueueLoadTest queues a load test for asynchronous processing
	QueueLoadTest(ctx context.Context, test *models.LoadTest) error

	// GenerateReport creates a load test report for a completed test
	GenerateReport(ctx context.Context, testID string) (string, error)
}

// AcceptanceTestService defines operations for acceptance testing
type AcceptanceTestService interface {
	// RunAcceptanceTest performs an acceptance test against a target
	RunAcceptanceTest(ctx context.Context, test *models.AcceptanceTest) (*models.AcceptanceResult, error)

	// GetAcceptanceTest retrieves an acceptance test by ID
	GetAcceptanceTest(ctx context.Context, testID string) (*models.AcceptanceTest, error)

	// ListAcceptanceTests retrieves all acceptance tests
	ListAcceptanceTests(ctx context.Context) ([]*models.AcceptanceTest, error)

	// QueueAcceptanceTest queues an acceptance test for asynchronous processing
	QueueAcceptanceTest(ctx context.Context, test *models.AcceptanceTest) error

	// GenerateReport creates an acceptance test report for a completed test
	GenerateReport(ctx context.Context, testID string) (string, error)
}

// GitHubService defines operations for GitHub integration
type GitHubService interface {
	// CreateIssue creates a GitHub issue
	CreateIssue(ctx context.Context, request *models.GitHubIssueRequest) (*models.GitHubIssueResponse, error)

	// GetIssue retrieves a GitHub issue by number
	GetIssue(ctx context.Context, owner, repo string, issueNumber int) (*models.GitHubIssueResponse, error)
}

// StorageService defines operations for storage
type StorageService interface {
	// SaveTest saves a test to storage
	SaveTest(ctx context.Context, testType models.TestType, testID string, data interface{}) error

	// GetTest retrieves a test from storage
	GetTest(ctx context.Context, testType models.TestType, testID string, result interface{}) error

	// ListTests retrieves all tests of a specific type from storage
	ListTests(ctx context.Context, testType models.TestType) ([]string, error)

	// SaveReport saves a test report to storage
	SaveReport(ctx context.Context, testType models.TestType, testID string, reportContent string) (string, error)

	// GetReport retrieves a test report from storage
	GetReport(ctx context.Context, reportPath string) (string, error)
}

// TestWorker defines operations for a test worker
type TestWorker interface {
	// Start starts the worker
	Start(ctx context.Context) error

	// Stop stops the worker
	Stop(ctx context.Context) error
}
