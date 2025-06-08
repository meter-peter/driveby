package validation

import "time"

// EndpointValidation represents the result of validating a single endpoint
type EndpointValidation struct {
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Status       string        `json:"status"`
	StatusCode   int           `json:"status_code"`
	ResponseTime time.Duration `json:"response_time"`
	Errors       []string      `json:"errors,omitempty"`
}

// EndpointValidationResult holds the results of endpoint validation
type EndpointValidationResult struct {
	Endpoints []EndpointValidation `json:"endpoints"`
}

// PerformanceMetrics holds metrics from a performance test run
type PerformanceMetrics struct {
	StartTime      time.Time     `json:"start_time"`
	EndTime        time.Time     `json:"end_time"`
	TotalRequests  uint64        `json:"total_requests"`
	SuccessCount   uint64        `json:"success_count"`
	ErrorCount     uint64        `json:"error_count"`
	ErrorRate      float64       `json:"error_rate"`
	LatencyP50     time.Duration `json:"latency_p50"`
	LatencyP95     time.Duration `json:"latency_p95"`
	LatencyP99     time.Duration `json:"latency_p99"`
	RequestsPerSec float64       `json:"requests_per_sec"`
}

// PerformanceTestResult holds the result of a performance test run
type PerformanceTestResult struct {
	Performance *PerformanceMetrics `json:"performance"`
}
