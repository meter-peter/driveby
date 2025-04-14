package models

import (
	"fmt"
	"time"
)

// ValidationTest represents a documentation validation test
type ValidationTest struct {
	TestBase
	OpenAPIURL           string             `json:"openapi_url"`
	ComplianceThreshold  float64            `json:"compliance_threshold"`
	FailOnValidation     bool               `json:"fail_on_validation"`
	ValidationParameters map[string]string  `json:"validation_parameters"`
	Result               *ValidationResult  `json:"result,omitempty"`
	GitHubIssueRequest   *GitHubIssueRequest `json:"github_issue_request,omitempty"`
}

// NewValidationTest creates a new validation test
func NewValidationTest(name, description, openAPIURL string, complianceThreshold float64) *ValidationTest {
	base := NewTestBase(TestTypeValidation, name, description)
	return &ValidationTest{
		TestBase:            base,
		OpenAPIURL:          openAPIURL,
		ComplianceThreshold: complianceThreshold,
		FailOnValidation:    true,
		ValidationParameters: map[string]string{},
	}
}

// ValidationResult represents the result of a validation test
type ValidationResult struct {
	BaseTestResult
	ComplianceScore       float64            `json:"compliance_score"`
	MissingExamples       int                `json:"missing_examples"`
	UndocumentedEndpoints []string           `json:"undocumented_endpoints"`
	ErrorResponses        map[string]int     `json:"error_responses"`
	ValidationErrors      []ValidationError  `json:"validation_errors"`
	ReportURL             string             `json:"report_url,omitempty"`
	ReportPath            string             `json:"report_path,omitempty"`
}

// ValidationError represents an error found during validation
type ValidationError struct {
	EndpointID string `json:"endpoint_id"`
	Message    string `json:"message"`
	Severity   string `json:"severity"` // "error", "warning", "info"
}

// NewValidationResult creates a new validation result
func NewValidationResult(testID string) *ValidationResult {
	now := time.Now()
	return &ValidationResult{
		BaseTestResult: BaseTestResult{
			TestID:    testID,
			Status:    TestStatusRunning,
			StartTime: now,
		},
		ErrorResponses: make(map[string]int),
	}
}

// IsSuccessful returns true if the validation passed
func (r *ValidationResult) IsSuccessful() bool {
	return r.Status == TestStatusCompleted && r.ComplianceScore >= 95.0
}

// GetSummary returns a summary of the validation result
func (r *ValidationResult) GetSummary() string {
	if r.Status != TestStatusCompleted {
		return r.BaseTestResult.GetSummary()
	}

	if r.IsSuccessful() {
		return "Validation passed with compliance score: " + 
			fmt.Sprintf("%.2f%%", r.ComplianceScore)
	}

	return "Validation failed with compliance score: " + 
		fmt.Sprintf("%.2f%% (threshold: %.2f%%)", r.ComplianceScore, 95.0)
}

// ValidationRequest represents a request to run a validation test
type ValidationRequest struct {
	Name               string            `json:"name" binding:"required"`
	Description        string            `json:"description"`
	OpenAPIURL         string            `json:"openapi_url" binding:"required,url"`
	ComplianceThreshold *float64         `json:"compliance_threshold"`
	FailOnValidation   *bool             `json:"fail_on_validation"`
	CreateGitHubIssue  bool              `json:"create_github_issue"`
	GitHubRepo         *GitHubIssueRequest `json:"github_repo,omitempty"`
	Tags               []string          `json:"tags"`
}

// ValidationResponse represents the response to a validation request
type ValidationResponse struct {
	TestID      string      `json:"test_id"`
	Status      TestStatus  `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	Result      *ValidationResult `json:"result,omitempty"`
}

// DocumentationReport holds metrics related to API documentation quality
type DocumentationReport struct {
	ComplianceScore       float64        `json:"compliance_score"`
	MissingExamples       int            `json:"missing_examples"`
	UndocumentedEndpoints []string       `json:"undocumented_endpoints"`
	ErrorResponses        map[string]int `json:"error_responses"` // count per status code
}