package validation

import (
	"time"
)

// ValidationReport represents the results of a validation run
type ValidationReport struct {
	Version      string
	Environment  string
	Timestamp    time.Time
	Principles   []PrincipleResult
	TotalChecks  int
	PassedChecks int
	FailedChecks int
	Summary      ValidationSummary
	AutoFixes    []AutoFixResult
	TestResults  *TestResults // Added test results to the main report
}

// TestResults contains results from functional and performance tests
type TestResults struct {
	Functional  *FunctionalTestResults
	Performance *PerformanceTestResults
	StartTime   time.Time
	EndTime     time.Time
	Status      TestStatus
}

// TestStatus represents the overall status of tests
type TestStatus string

const (
	TestStatusPassed     TestStatus = "passed"
	TestStatusFailed     TestStatus = "failed"
	TestStatusWarning    TestStatus = "warning"
	TestStatusSkipped    TestStatus = "skipped"
	TestStatusIncomplete TestStatus = "incomplete"
)

// FunctionalTestResults contains results from functional testing
type FunctionalTestResults struct {
	TotalEndpoints      int
	TestedEndpoints     int
	PassedEndpoints     int
	FailedEndpoints     int
	SkippedEndpoints    int
	EndpointResults     []EndpointTestResult
	AverageResponseTime time.Duration
	MaxResponseTime     time.Duration
	MinResponseTime     time.Duration
}

// EndpointTestResult represents a single endpoint test result
type EndpointTestResult struct {
	Method       string
	Path         string
	Status       TestStatus
	StatusCode   int
	ResponseTime time.Duration
	Errors       []string
	Warnings     []string
	TestCases    []TestCaseResult
}

// TestCaseResult represents a single test case result
type TestCaseResult struct {
	Name        string
	Status      TestStatus
	Description string
	Input       interface{}
	Expected    interface{}
	Actual      interface{}
	Error       string
}

// PerformanceTestResults contains results from performance testing
type PerformanceTestResults struct {
	TotalRequests     int64
	SuccessCount      int64
	ErrorCount        int64
	ErrorRate         float64
	LatencyP50        time.Duration
	LatencyP95        time.Duration
	LatencyP99        time.Duration
	RequestsPerSecond float64
	Duration          time.Duration
	Status            TestStatus
	FailedRequests    []FailedRequest
}

// FailedRequest represents a failed request during performance testing
type FailedRequest struct {
	Method     string
	Path       string
	StatusCode int
	Error      string
	Timestamp  time.Time
	Latency    time.Duration
}

// ValidationSummary provides a high-level summary of validation results
type ValidationSummary struct {
	CriticalIssues int
	Warnings       int
	Info           int
	Categories     []string
	FailedTags     []string
	TestSummary    *TestSummary
}

// TestSummary provides a high-level summary of test results
type TestSummary struct {
	FunctionalStatus  TestStatus
	PerformanceStatus TestStatus
	TotalTests        int
	PassedTests       int
	FailedTests       int
	Warnings          int
	SkippedTests      int
}

// PrincipleResult represents the result of validating a single principle
type PrincipleResult struct {
	Principle    Principle
	Passed       bool
	Message      string
	Details      interface{}
	Explanation  string
	SuggestedFix string
	TestImpact   *TestImpact // Added to show impact on testing
}

// TestImpact represents how a validation result impacts testing
type TestImpact struct {
	CanRunTests     bool
	AffectedTests   []string
	ImpactLevel     ImpactLevel
	Recommendations []string
}

// ImpactLevel represents the severity of test impact
type ImpactLevel string

const (
	ImpactLevelNone     ImpactLevel = "none"
	ImpactLevelLow      ImpactLevel = "low"
	ImpactLevelMedium   ImpactLevel = "medium"
	ImpactLevelHigh     ImpactLevel = "high"
	ImpactLevelCritical ImpactLevel = "critical"
)

// AutoFixResult represents the result of an automatic fix attempt
type AutoFixResult struct {
	PrincipleID string
	Timestamp   time.Time
	Location    string
	Success     bool
	Message     string
	Error       string
	Original    PrincipleResult
	Fixed       string
	TestImpact  *TestImpact // Added to show impact on testing
}

// ValidationMode defines the level of validation to perform
type ValidationMode string

const (
	ValidationModeStrict   ValidationMode = "strict"    // Comprehensive validation
	ValidationModeMinimal  ValidationMode = "minimal"   // Basic validation for test generation
	ValidationModeTestOnly ValidationMode = "test-only" // Skip validation, run tests only
	ValidationModeFlexible ValidationMode = "flexible"  // Allow tests even with some validation failures
)

// TestMode defines the type of testing to perform
type TestMode string

const (
	TestModeNone        TestMode = "none"        // No testing
	TestModeFunctional  TestMode = "functional"  // Only functional testing
	TestModePerformance TestMode = "performance" // Only performance testing
	TestModeAll         TestMode = "all"         // Both functional and performance testing
)

// ValidatorConfig holds configuration for the validator
type ValidatorConfig struct {
	BaseURL           string
	SpecPath          string
	Environment       string
	Version           string
	Timeout           time.Duration
	ValidationMode    ValidationMode
	Auth              *AuthConfig // Add back Auth field for token support
	PerformanceTarget *PerformanceTargetConfig
}

// PerformanceTargetConfig holds configuration for performance test targets
type PerformanceTargetConfig struct {
	MaxLatencyP95   time.Duration
	MinSuccessRate  float64
	ConcurrentUsers int
	Duration        time.Duration
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	Token        string
	TokenType    string
	TokenHeader  string
	Username     string
	Password     string
	APIKey       string
	APIKeyHeader string
}
