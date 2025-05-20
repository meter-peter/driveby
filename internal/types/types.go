package types

import (
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResult represents the overall result of an API test
type TestResult struct {
	TestID        string     `json:"test_id"`
	Timestamp     time.Time  `json:"timestamp"`
	Documentation DocResult  `json:"documentation"`
	Integration   IntResult  `json:"integration"`
	LoadTest      LoadResult `json:"load_test"`
}

// DocResult represents documentation validation results
type DocResult struct {
	ComplianceScore       float64        `json:"compliance_score"`
	MissingExamples       int            `json:"missing_examples"`
	UndocumentedEndpoints []string       `json:"undocumented_endpoints"`
	ErrorResponses        map[string]int `json:"error_responses"`
	Passed                bool           `json:"passed"`
	Threshold             float64        `json:"threshold"`
}

// IntResult represents integration test results
type IntResult struct {
	TotalTests      int               `json:"total_tests"`
	PassedTests     int               `json:"passed_tests"`
	FailedTests     int               `json:"failed_tests"`
	FailedEndpoints map[string]string `json:"failed_endpoints"` // endpoint -> error message
	Passed          bool              `json:"passed"`
}

// LoadResult represents load test results
type LoadResult struct {
	TotalRequests int64          `json:"total_requests"`
	SuccessRate   float64        `json:"success_rate"`
	LatencyP95    time.Duration  `json:"latency_p95"`
	ErrorRate     float64        `json:"error_rate"`
	StatusCodes   map[int]int    `json:"status_codes"`
	Passed        bool           `json:"passed"`
	Thresholds    LoadThresholds `json:"thresholds"`
}

// LoadThresholds defines the thresholds for load test criteria
type LoadThresholds struct {
	MinSuccessRate float64       `json:"min_success_rate"`
	MaxLatencyP95  time.Duration `json:"max_latency_p95"`
	MaxErrorRate   float64       `json:"max_error_rate"`
}

// TestRequest represents a request to run tests
type TestRequest struct {
	OpenAPISpec    *openapi3.T    `json:"openapi_spec"`
	LoadTestConfig LoadTestConfig `json:"load_test_config"`
	Thresholds     TestThresholds `json:"thresholds"`
}

// LoadTestConfig defines the configuration for load testing
type LoadTestConfig struct {
	RequestRate    int           `json:"request_rate"` // requests per second
	TestDuration   time.Duration `json:"test_duration"`
	RequestTimeout time.Duration `json:"request_timeout"`
}

// TestThresholds defines all thresholds for different test types
type TestThresholds struct {
	Documentation struct {
		MinComplianceScore float64 `json:"min_compliance_score"`
		MaxMissingExamples int     `json:"max_missing_examples"`
	} `json:"documentation"`
	Integration struct {
		MinPassRate float64 `json:"min_pass_rate"`
	} `json:"integration"`
	LoadTest LoadThresholds `json:"load_test"`
}

// TestResponse represents the response from running tests
type TestResponse struct {
	TestID    string     `json:"test_id"`
	Timestamp time.Time  `json:"timestamp"`
	Results   TestResult `json:"results"`
	Metadata  struct {
		ServiceName    string `json:"service_name"`
		ServiceVersion string `json:"service_version"`
		Environment    string `json:"environment"`
	} `json:"metadata"`
}
