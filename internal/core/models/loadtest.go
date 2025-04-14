package models

import (
	"fmt"
	"time"
)

// LoadTest represents a load test configuration
type LoadTest struct {
	TestBase
	TargetURL       string            `json:"target_url"`
	RequestRate     int               `json:"request_rate"`     // Requests per second
	Duration        time.Duration     `json:"duration"`         // Test duration
	Timeout         time.Duration     `json:"timeout"`          // Request timeout
	Method          string            `json:"method"`           // HTTP method
	Headers         map[string]string `json:"headers"`          // HTTP headers
	Body            string            `json:"body,omitempty"`   // Request body
	SuccessThreshold float64          `json:"success_threshold"` // Minimum success rate to pass
	Endpoints       []LoadTestEndpoint `json:"endpoints,omitempty"` // Multiple endpoints to test
	Result          *LoadTestResult   `json:"result,omitempty"`
	GitHubIssueRequest *GitHubIssueRequest `json:"github_issue_request,omitempty"`
}

// LoadTestEndpoint represents a single endpoint in a load test
type LoadTestEndpoint struct {
	Path    string            `json:"path"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
	Weight  int               `json:"weight"` // Relative frequency of this endpoint in the test
}

// NewLoadTest creates a new load test
func NewLoadTest(name, description, targetURL string, requestRate int, duration time.Duration) *LoadTest {
	base := NewTestBase(TestTypeLoadTest, name, description)
	return &LoadTest{
		TestBase:         base,
		TargetURL:        targetURL,
		RequestRate:      requestRate,
		Duration:         duration,
		Timeout:          5 * time.Second,
		Method:           "GET",
		Headers:          map[string]string{},
		SuccessThreshold: 95.0,
		Endpoints:        []LoadTestEndpoint{},
	}
}

// LoadTestResult represents the result of a load test
type LoadTestResult struct {
	BaseTestResult
	Throughput         float64                 `json:"throughput"`           // Actual throughput in requests/sec
	SuccessRate        float64                 `json:"success_rate"`         // Percentage of successful requests
	TotalRequests      int                     `json:"total_requests"`       // Total number of requests made
	FailedRequests     int                     `json:"failed_requests"`      // Number of failed requests
	Latencies          LatencyMetrics          `json:"latencies"`            // Latency statistics
	StatusCodeCounts   map[string]int          `json:"status_code_counts"`   // Count of each status code
	ErrorCounts        map[string]int          `json:"error_counts"`         // Count of each error type
	EndpointPerformance []EndpointPerformance  `json:"endpoint_performance"` // Performance by endpoint
	ReportURL          string                  `json:"report_url,omitempty"` // URL to detailed report
	ReportPath         string                  `json:"report_path,omitempty"` // Path to detailed report
}

// LatencyMetrics represents latency statistics
type LatencyMetrics struct {
	Min    time.Duration `json:"min"`
	Mean   time.Duration `json:"mean"`
	P50    time.Duration `json:"p50"` // 50th percentile (median)
	P90    time.Duration `json:"p90"` // 90th percentile
	P95    time.Duration `json:"p95"` // 95th percentile
	P99    time.Duration `json:"p99"` // 99th percentile
	Max    time.Duration `json:"max"`
}

// EndpointPerformance represents performance metrics for a specific endpoint
type EndpointPerformance struct {
	Path        string        `json:"path"`
	Method      string        `json:"method"`
	SuccessRate float64       `json:"success_rate"`
	Latency     LatencyMetrics `json:"latency"`
	Requests    int           `json:"requests"`
}

// NewLoadTestResult creates a new load test result
func NewLoadTestResult(testID string) *LoadTestResult {
	now := time.Now()
	return &LoadTestResult{
		BaseTestResult: BaseTestResult{
			TestID:    testID,
			Status:    TestStatusRunning,
			StartTime: now,
		},
		StatusCodeCounts: make(map[string]int),
		ErrorCounts:      make(map[string]int),
	}
}

// IsSuccessful returns true if the load test passed
func (r *LoadTestResult) IsSuccessful() bool {
	return r.Status == TestStatusCompleted && r.SuccessRate >= 95.0
}

// GetSummary returns a summary of the load test result
func (r *LoadTestResult) GetSummary() string {
	if r.Status != TestStatusCompleted {
		return r.BaseTestResult.GetSummary()
	}

	if r.IsSuccessful() {
		return fmt.Sprintf("Load test passed with %.2f%% success rate. Throughput: %.2f req/s, Median latency: %s", 
			r.SuccessRate, r.Throughput, r.Latencies.P50)
	}

	return fmt.Sprintf("Load test failed with %.2f%% success rate (threshold: %.2f%%)", 
		r.SuccessRate, 95.0)
}

// LoadTestRequest represents a request to run a load test
type LoadTestRequest struct {
	Name             string            `json:"name" binding:"required"`
	Description      string            `json:"description"`
	TargetURL        string            `json:"target_url" binding:"required,url"`
	RequestRate      int               `json:"request_rate" binding:"required,min=1"`
	Duration         int               `json:"duration" binding:"required,min=1"` // Duration in seconds
	Timeout          *int              `json:"timeout,omitempty"`                 // Timeout in seconds
	Method           string            `json:"method,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
	Body             string            `json:"body,omitempty"`
	SuccessThreshold *float64          `json:"success_threshold,omitempty"`
	Endpoints        []LoadTestEndpoint `json:"endpoints,omitempty"`
	CreateGitHubIssue bool             `json:"create_github_issue"`
	GitHubRepo       *GitHubIssueRequest `json:"github_repo,omitempty"`
	Tags             []string          `json:"tags"`
}

// LoadTestResponse represents the response to a load test request
type LoadTestResponse struct {
	TestID    string         `json:"test_id"`
	Status    TestStatus     `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	Result    *LoadTestResult `json:"result,omitempty"`
}