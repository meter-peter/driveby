package services

import (
	"context"
	"fmt"

	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/models"
	"github.com/example/driveby/internal/queue"
	"github.com/sirupsen/logrus"
)

// ServiceManager manages all the services and provides a unified API
type ServiceManager struct {
	config       *config.Config
	queueService queue.QueueService
	logger       *logrus.Logger

	// Service instances
	validationService  ValidationService
	loadTestService    LoadTestService
	acceptanceService  AcceptanceTestService
	githubService      GitHubService
	storageService     StorageService

	// Task handlers
	handlers map[string]queue.TaskHandler
}

// TaskTypes for queue tasks
const (
	TaskTypeValidation  = "validation_test"
	TaskTypeLoadTest    = "load_test"
	TaskTypeAcceptance  = "acceptance_test"
	TaskTypeGitHubIssue = "github_issue"
)

// NewManager creates a new service manager
func NewManager(cfg *config.Config, queueService queue.QueueService, logger *logrus.Logger) *ServiceManager {
	manager := &ServiceManager{
		config:       cfg,
		queueService: queueService,
		logger:       logger,
		handlers:     make(map[string]queue.TaskHandler),
	}

	// Initialize services
	manager.initServices()

	// Register task handlers
	manager.registerTaskHandlers()

	return manager
}

// initServices initializes all the services
func (m *ServiceManager) initServices() {
	// Initialize storage service
	storageService, err := NewMinioStorageService(m.config, m.logger)
	if err != nil {
		m.logger.WithError(err).Error("Failed to initialize Minio storage service")
	} else {
		m.storageService = storageService
		m.logger.Info("Minio storage service initialized")
	}

	// Initialize GitHub service
	m.githubService = NewGitHubService(m.config, m.logger)
	m.logger.Info("GitHub service initialized")

	// Initialize validation service
	m.validationService = NewValidationService(
		m.config,
		m.logger,
		m.storageService,
		m.githubService,
		m.queueService,
	)
	m.logger.Info("Validation service initialized")

	// Initialize load test service (placeholder for now)
	m.loadTestService = nil

	// Initialize acceptance test service (placeholder for now)
	m.acceptanceService = nil

	m.logger.Info("Services initialized")
}

// registerTaskHandlers registers task handlers for the queue
func (m *ServiceManager) registerTaskHandlers() {
	// Register validation test handler
	m.queueService.RegisterHandler(TaskTypeValidation, m.handleValidationTask)

	// Register load test handler
	m.queueService.RegisterHandler(TaskTypeLoadTest, m.handleLoadTestTask)

	// Register acceptance test handler
	m.queueService.RegisterHandler(TaskTypeAcceptance, m.handleAcceptanceTask)

	// Register GitHub issue handler
	m.queueService.RegisterHandler(TaskTypeGitHubIssue, m.handleGitHubIssueTask)
}

// StartWorkers starts the queue workers
func (m *ServiceManager) StartWorkers(ctx context.Context) error {
	if !m.config.Features.EnableWorkers {
		m.logger.Info("Workers disabled by configuration")
		return nil
	}

	// Default worker count is numCPU, but can be configured
	workerCount := 5 // This should come from config

	return m.queueService.StartWorkers(ctx, workerCount)
}

// handleValidationTask handles a validation test task
func (m *ServiceManager) handleValidationTask(ctx context.Context, task *models.QueueTask) error {
	m.logger.WithField("task_id", task.ID).Info("Processing validation task")

	// If validation service not initialized, return error
	if m.validationService == nil {
		return fmt.Errorf("validation service not initialized")
	}

	// Unmarshal payload to validation test
	var test models.ValidationTest
	if err := task.UnmarshalPayload(&test); err != nil {
		return fmt.Errorf("failed to unmarshal validation test: %w", err)
	}

	// Run validation
	result, err := m.validationService.ValidateOpenAPI(ctx, &test)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Generate report if validation completed
	if result.Status == models.TestStatusCompleted {
		reportPath, err := m.validationService.GenerateReport(ctx, test.ID)
		if err != nil {
			m.logger.WithError(err).Error("Failed to generate validation report")
		} else {
			result.ReportPath = reportPath
		}
	}

	// Create GitHub issue if requested and validation failed
	if test.Result.Status == models.TestStatusCompleted && !test.Result.IsSuccessful() && test.GitHubIssueRequest != nil {
		m.logger.WithField("test_id", test.ID).Info("Creating GitHub issue for failed validation")
		
		// Queue GitHub issue creation
		if m.config.Features.EnableGitHub {
			_, err := m.queueService.Enqueue(ctx, TaskTypeGitHubIssue, 
				&models.GitHubIssueRequest{
					Owner:      test.GitHubIssueRequest.Owner,
					Repository: test.GitHubIssueRequest.Repository,
					Title:      fmt.Sprintf("API Documentation Validation Failed: %s", test.Name),
					Body:       m.generateValidationIssueBody(&test),
					Labels:     []string{"validation", "documentation", "api"},
				})
			if err != nil {
				m.logger.WithError(err).Error("Failed to queue GitHub issue creation")
			}
		}
	}

	return nil
}

// handleLoadTestTask handles a load test task
func (m *ServiceManager) handleLoadTestTask(ctx context.Context, task *models.QueueTask) error {
	m.logger.WithField("task_id", task.ID).Info("Processing load test task")

	// If load test service not initialized, return error
	if m.loadTestService == nil {
		return fmt.Errorf("load test service not initialized")
	}

	// Unmarshal payload to load test
	var test models.LoadTest
	if err := task.UnmarshalPayload(&test); err != nil {
		return fmt.Errorf("failed to unmarshal load test: %w", err)
	}

	// Run load test
	result, err := m.loadTestService.RunLoadTest(ctx, &test)
	if err != nil {
		return fmt.Errorf("load test failed: %w", err)
	}

	// Generate report if load test completed
	if result.Status == models.TestStatusCompleted {
		reportPath, err := m.loadTestService.GenerateReport(ctx, test.ID)
		if err != nil {
			m.logger.WithError(err).Error("Failed to generate load test report")
		} else {
			result.ReportPath = reportPath
		}
	}

	// Create GitHub issue if requested and load test failed
	if test.Result.Status == models.TestStatusCompleted && !test.Result.IsSuccessful() && test.GitHubIssueRequest != nil {
		m.logger.WithField("test_id", test.ID).Info("Creating GitHub issue for failed load test")
		
		// Queue GitHub issue creation
		if m.config.Features.EnableGitHub {
			_, err := m.queueService.Enqueue(ctx, TaskTypeGitHubIssue, 
				&models.GitHubIssueRequest{
					Owner:      test.GitHubIssueRequest.Owner,
					Repository: test.GitHubIssueRequest.Repository,
					Title:      fmt.Sprintf("Load Test Failed: %s", test.Name),
					Body:       m.generateLoadTestIssueBody(&test),
					Labels:     []string{"load-test", "performance", "api"},
				})
			if err != nil {
				m.logger.WithError(err).Error("Failed to queue GitHub issue creation")
			}
		}
	}

	return nil
}

// handleAcceptanceTask handles an acceptance test task
func (m *ServiceManager) handleAcceptanceTask(ctx context.Context, task *models.QueueTask) error {
	m.logger.WithField("task_id", task.ID).Info("Processing acceptance test task")

	// If acceptance service not initialized, return error
	if m.acceptanceService == nil {
		return fmt.Errorf("acceptance test service not initialized")
	}

	// Unmarshal payload to acceptance test
	var test models.AcceptanceTest
	if err := task.UnmarshalPayload(&test); err != nil {
		return fmt.Errorf("failed to unmarshal acceptance test: %w", err)
	}

	// Run acceptance test
	result, err := m.acceptanceService.RunAcceptanceTest(ctx, &test)
	if err != nil {
		return fmt.Errorf("acceptance test failed: %w", err)
	}

	// Generate report if acceptance test completed
	if result.Status == models.TestStatusCompleted {
		reportPath, err := m.acceptanceService.GenerateReport(ctx, test.ID)
		if err != nil {
			m.logger.WithError(err).Error("Failed to generate acceptance test report")
		} else {
			result.ReportPath = reportPath
		}
	}

	// Create GitHub issue if requested and acceptance test failed
	if test.Result.Status == models.TestStatusCompleted && !test.Result.IsSuccessful() && test.GitHubIssueRequest != nil {
		m.logger.WithField("test_id", test.ID).Info("Creating GitHub issue for failed acceptance test")
		
		// Queue GitHub issue creation
		if m.config.Features.EnableGitHub {
			_, err := m.queueService.Enqueue(ctx, TaskTypeGitHubIssue, 
				&models.GitHubIssueRequest{
					Owner:      test.GitHubIssueRequest.Owner,
					Repository: test.GitHubIssueRequest.Repository,
					Title:      fmt.Sprintf("Acceptance Test Failed: %s", test.Name),
					Body:       m.generateAcceptanceTestIssueBody(&test),
					Labels:     []string{"acceptance-test", "testing", "api"},
				})
			if err != nil {
				m.logger.WithError(err).Error("Failed to queue GitHub issue creation")
			}
		}
	}

	return nil
}

// handleGitHubIssueTask handles a GitHub issue creation task
func (m *ServiceManager) handleGitHubIssueTask(ctx context.Context, task *models.QueueTask) error {
	m.logger.WithField("task_id", task.ID).Info("Processing GitHub issue task")

	// If GitHub service not initialized, return error
	if m.githubService == nil {
		return fmt.Errorf("GitHub service not initialized")
	}

	// Unmarshal payload to GitHub issue request
	var request models.GitHubIssueRequest
	if err := task.UnmarshalPayload(&request); err != nil {
		return fmt.Errorf("failed to unmarshal GitHub issue request: %w", err)
	}

	// Create GitHub issue
	response, err := m.githubService.CreateIssue(ctx, &request)
	if err != nil {
		return fmt.Errorf("failed to create GitHub issue: %w", err)
	}

	m.logger.WithFields(logrus.Fields{
		"issue_number": response.IssueNumber,
		"issue_url":    response.IssueURL,
	}).Info("GitHub issue created")

	return nil
}

// Helper methods for generating GitHub issue bodies

// generateValidationIssueBody generates a GitHub issue body for a failed validation test
func (m *ServiceManager) generateValidationIssueBody(test *models.ValidationTest) string {
	result := test.Result
	if result == nil {
		return "No validation results available."
	}

	body := fmt.Sprintf(`## API Documentation Validation Failed

**Test:** %s  
**Compliance Score:** %.2f%% (Threshold: %.2f%%)  
**Missing Examples:** %d  
**Error Responses:** %d

### Undocumented Endpoints

`,
		test.Name,
		result.ComplianceScore,
		test.ComplianceThreshold,
		result.MissingExamples,
		len(result.ErrorResponses),
	)

	// Add undocumented endpoints
	if len(result.UndocumentedEndpoints) > 0 {
		for _, endpoint := range result.UndocumentedEndpoints {
			body += fmt.Sprintf("- `%s`\n", endpoint)
		}
	} else {
		body += "No undocumented endpoints.\n"
	}

	// Add validation errors
	body += "\n### Validation Errors\n\n"
	if len(result.ValidationErrors) > 0 {
		for _, err := range result.ValidationErrors {
			body += fmt.Sprintf("- **%s**: %s\n", err.EndpointID, err.Message)
		}
	} else {
		body += "No validation errors.\n"
	}

	// Add report link if available
	if result.ReportURL != "" {
		body += fmt.Sprintf("\n### [View Full Report](%s)\n", result.ReportURL)
	}

	return body
}

// generateLoadTestIssueBody generates a GitHub issue body for a failed load test
func (m *ServiceManager) generateLoadTestIssueBody(test *models.LoadTest) string {
	result := test.Result
	if result == nil {
		return "No load test results available."
	}

	body := fmt.Sprintf(`## Load Test Failed

**Test:** %s  
**Target URL:** %s  
**Request Rate:** %d req/sec  
**Duration:** %s  
**Success Rate:** %.2f%% (Threshold: %.2f%%)  
**Total Requests:** %d  
**Failed Requests:** %d

### Latency Statistics

- **Min:** %s
- **Mean:** %s
- **P50 (Median):** %s
- **P95:** %s
- **P99:** %s
- **Max:** %s

`,
		test.Name,
		test.TargetURL,
		test.RequestRate,
		test.Duration.String(),
		result.SuccessRate,
		test.SuccessThreshold,
		result.TotalRequests,
		result.FailedRequests,
		result.Latencies.Min.String(),
		result.Latencies.Mean.String(),
		result.Latencies.P50.String(),
		result.Latencies.P95.String(),
		result.Latencies.P99.String(),
		result.Latencies.Max.String(),
	)

	// Add status code distribution
	body += "### Status Code Distribution\n\n"
	if len(result.StatusCodeCounts) > 0 {
		for code, count := range result.StatusCodeCounts {
			percentage := float64(count) / float64(result.TotalRequests) * 100
			body += fmt.Sprintf("- **%s**: %d (%.2f%%)\n", code, count, percentage)
		}
	} else {
		body += "No status code information available.\n"
	}

	// Add error distribution
	body += "\n### Error Distribution\n\n"
	if len(result.ErrorCounts) > 0 {
		for err, count := range result.ErrorCounts {
			percentage := float64(count) / float64(result.FailedRequests) * 100
			body += fmt.Sprintf("- **%s**: %d (%.2f%%)\n", err, count, percentage)
		}
	} else {
		body += "No error information available.\n"
	}

	// Add report link if available
	if result.ReportURL != "" {
		body += fmt.Sprintf("\n### [View Full Report](%s)\n", result.ReportURL)
	}

	return body
}

// generateAcceptanceTestIssueBody generates a GitHub issue body for a failed acceptance test
func (m *ServiceManager) generateAcceptanceTestIssueBody(test *models.AcceptanceTest) string {
	result := test.Result
	if result == nil {
		return "No acceptance test results available."
	}

	body := fmt.Sprintf(`## Acceptance Test Failed

**Test:** %s  
**Base URL:** %s  
**Pass Rate:** %.2f%%  
**Total Test Cases:** %d  
**Passed:** %d  
**Failed:** %d  
**Skipped:** %d

### Failed Test Cases

`,
		test.Name,
		test.BaseURL,
		result.PassRate,
		result.TotalTests,
		result.PassedTests,
		result.FailedTests,
		result.SkippedTests,
	)

	// Add failed test cases
	for _, testCase := range result.TestCaseResults {
		if testCase.Status == "failed" {
			body += fmt.Sprintf("#### %s\n\n", testCase.Name)
			body += fmt.Sprintf("- **Status Code:** %d\n", testCase.StatusCode)
			body += fmt.Sprintf("- **Error:** %s\n", testCase.Error)
			
			// Add failed assertions
			body += "\n**Failed Assertions:**\n\n"
			for _, assertion := range testCase.AssertionResults {
				if !assertion.Success {
					body += fmt.Sprintf("- **%s %s %s**\n", assertion.Type, assertion.Target, assertion.Command)
					body += fmt.Sprintf("  - Expected: `%v`\n", assertion.Expected)
					body += fmt.Sprintf("  - Actual: `%v`\n", assertion.Actual)
					if assertion.Error != "" {
						body += fmt.Sprintf("  - Error: %s\n", assertion.Error)
					}
				}
			}
			body += "\n"
		}
	}

	// Add report link if available
	if result.ReportURL != "" {
		body += fmt.Sprintf("\n### [View Full Report](%s)\n", result.ReportURL)
	}

	return body
}

// GetValidationService returns the validation service
func (m *ServiceManager) GetValidationService() ValidationService {
	return m.validationService
}

// GetLoadTestService returns the load test service
func (m *ServiceManager) GetLoadTestService() LoadTestService {
	return m.loadTestService
}

// GetAcceptanceService returns the acceptance test service
func (m *ServiceManager) GetAcceptanceService() AcceptanceTestService {
	return m.acceptanceService
}

// GetGitHubService returns the GitHub service
func (m *ServiceManager) GetGitHubService() GitHubService {
	return m.githubService
}

// GetStorageService returns the storage service
func (m *ServiceManager) GetStorageService() StorageService {
	return m.storageService
}