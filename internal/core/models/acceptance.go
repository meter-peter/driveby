package models

import (
	"fmt"
	"time"
)

// AcceptanceTest represents an acceptance test configuration
type AcceptanceTest struct {
	TestBase
	BaseURL           string                 `json:"base_url"`
	Timeout           time.Duration          `json:"timeout"`
	Headers           map[string]string      `json:"headers"`
	GlobalVariables   map[string]interface{} `json:"global_variables"`
	TestCases         []TestCase             `json:"test_cases"`
	Result            *AcceptanceResult      `json:"result,omitempty"`
	GitHubIssueRequest *GitHubIssueRequest   `json:"github_issue_request,omitempty"`
}

// TestCase represents a single test case in an acceptance test
type TestCase struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Path        string                 `json:"path"`
	Method      string                 `json:"method"`
	Headers     map[string]string      `json:"headers,omitempty"`
	QueryParams map[string]string      `json:"query_params,omitempty"`
	Body        interface{}            `json:"body,omitempty"`
	Variables   map[string]interface{} `json:"variables,omitempty"`
	Assertions  []Assertion            `json:"assertions"`
	DependsOn   []string               `json:"depends_on,omitempty"`
	Weight      int                    `json:"weight"`
	Timeout     *time.Duration         `json:"timeout,omitempty"`
}

// Assertion represents a validation check for a test case
type Assertion struct {
	Type    string      `json:"type"`    // "status", "json", "header", "time"
	Target  string      `json:"target"`  // JSON path, header name, etc.
	Value   interface{} `json:"value"`   // Expected value
	Command string      `json:"command"` // Comparison: "eq", "neq", "contains", "gt", "lt"
}

// NewAcceptanceTest creates a new acceptance test
func NewAcceptanceTest(name, description, baseURL string) *AcceptanceTest {
	base := NewTestBase(TestTypeAcceptance, name, description)
	return &AcceptanceTest{
		TestBase:        base,
		BaseURL:         baseURL,
		Timeout:         30 * time.Second,
		Headers:         map[string]string{},
		GlobalVariables: map[string]interface{}{},
		TestCases:       []TestCase{},
	}
}

// AcceptanceResult represents the results of an acceptance test
type AcceptanceResult struct {
	BaseTestResult
	TestCaseResults []TestCaseResult `json:"test_case_results"`
	PassedTests     int              `json:"passed_tests"`
	FailedTests     int              `json:"failed_tests"`
	SkippedTests    int              `json:"skipped_tests"`
	TotalTests      int              `json:"total_tests"`
	PassRate        float64          `json:"pass_rate"`
	ReportURL       string           `json:"report_url,omitempty"`
	ReportPath      string           `json:"report_path,omitempty"`
}

// TestCaseResult represents the result of a single test case
type TestCaseResult struct {
	Name           string                  `json:"name"`
	Status         string                  `json:"status"` // "passed", "failed", "skipped", "error"
	Duration       time.Duration           `json:"duration"`
	StatusCode     int                     `json:"status_code,omitempty"`
	ResponseBody   string                  `json:"response_body,omitempty"`
	ResponseHeaders map[string]string      `json:"response_headers,omitempty"`
	Error          string                  `json:"error,omitempty"`
	AssertionResults []AssertionResult     `json:"assertion_results,omitempty"`
	Variables      map[string]interface{}  `json:"variables,omitempty"`
}

// AssertionResult represents the result of a single assertion
type AssertionResult struct {
	Type    string      `json:"type"`
	Target  string      `json:"target"`
	Command string      `json:"command"`
	Expected interface{} `json:"expected"`
	Actual   interface{} `json:"actual"`
	Success  bool        `json:"success"`
	Error    string      `json:"error,omitempty"`
}

// NewAcceptanceResult creates a new acceptance test result
func NewAcceptanceResult(testID string) *AcceptanceResult {
	now := time.Now()
	return &AcceptanceResult{
		BaseTestResult: BaseTestResult{
			TestID:    testID,
			Status:    TestStatusRunning,
			StartTime: now,
		},
	}
}

// IsSuccessful returns true if the acceptance test passed
func (r *AcceptanceResult) IsSuccessful() bool {
	return r.Status == TestStatusCompleted && r.PassRate >= 100.0
}

// GetSummary returns a summary of the acceptance test result
func (r *AcceptanceResult) GetSummary() string {
	if r.Status != TestStatusCompleted {
		return r.BaseTestResult.GetSummary()
	}

	if r.IsSuccessful() {
		return fmt.Sprintf("Acceptance test passed with %.2f%% success rate (%d/%d test cases)", 
			r.PassRate, r.PassedTests, r.TotalTests)
	}

	return fmt.Sprintf("Acceptance test failed with %.2f%% success rate (%d/%d test cases)", 
		r.PassRate, r.PassedTests, r.TotalTests)
}

// AcceptanceTestRequest represents a request to run an acceptance test
type AcceptanceTestRequest struct {
	Name             string                 `json:"name" binding:"required"`
	Description      string                 `json:"description"`
	BaseURL          string                 `json:"base_url" binding:"required,url"`
	Timeout          *int                   `json:"timeout,omitempty"` // In seconds
	Headers          map[string]string      `json:"headers,omitempty"`
	GlobalVariables  map[string]interface{} `json:"global_variables,omitempty"`
	TestCases        []TestCase             `json:"test_cases" binding:"required,min=1"`
	CreateGitHubIssue bool                  `json:"create_github_issue"`
	GitHubRepo        *GitHubIssueRequest   `json:"github_repo,omitempty"`
	Tags              []string              `json:"tags"`
}

// AcceptanceTestResponse represents the response to an acceptance test request
type AcceptanceTestResponse struct {
	TestID    string             `json:"test_id"`
	Status    TestStatus         `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	Result    *AcceptanceResult  `json:"result,omitempty"`
}