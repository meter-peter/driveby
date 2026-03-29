package types

import (
	"fmt"
	"time"
)

// ValidationReport represents the results of a validation run
type ValidationReport struct {
	Status       string            `json:"status"`    // "passed" or "failed"
	ExitCode     int               `json:"exit_code"` // 0=pass, 1=validation_failed, 2=error
	Version      string            `json:"version"`
	Environment  string            `json:"environment"`
	Timestamp    time.Time         `json:"timestamp"`
	Principles   []PrincipleResult `json:"principles"`
	TotalChecks  int               `json:"total_checks"`
	PassedChecks int               `json:"passed_checks"`
	FailedChecks int               `json:"failed_checks"`
	Summary      ValidationSummary `json:"summary"`
	AutoFixes    []AutoFixResult   `json:"auto_fixes,omitempty"`
	TestResults  *TestResults      `json:"test_results,omitempty"`
}

// TestResults contains results from functional and performance tests
type TestResults struct {
	Functional  *FunctionalTestResults
	Performance *PerformanceTestResults
	StartTime   time.Time
	EndTime     time.Time
	Status      TestStatus
}

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

// Principle represents a validation principle
type Principle struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Severity    string   `json:"severity"`
	Tags        []string `json:"tags"`
	AutoFixable bool     `json:"auto_fixable"`
	Checks      []string `json:"checks,omitempty"`
}

// CorePrinciples defines the core validation principles
var CorePrinciples = []Principle{
	{
		ID:          "P001",
		Name:        "OpenAPI Specification Compliance",
		Description: "Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices",
		Category:    "Specification",
		Severity:    "critical",
		Tags:        []string{"openapi", "specification", "compliance"},
		AutoFixable: true,
		Checks: []string{
			"OpenAPI version is 3.0.x or 3.1.0",
			"Required info fields (title, version) are present",
			"Paths are properly defined",
			"Components are valid",
			"References are resolvable",
			"No duplicate operationIds",
			"Valid HTTP methods used",
		},
	},
	{
		ID:          "P002",
		Name:        "API Documentation Quality",
		Description: "Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines",
		Category:    "Documentation",
		Severity:    "critical",
		Tags:        []string{"documentation", "quality", "usability"},
		AutoFixable: true,
		Checks: []string{
			"All operations have clear summaries",
			"All operations have detailed descriptions",
			"All operations have unique operationIds",
			"All parameters have descriptions",
			"All request/response bodies have examples",
			"All schemas have descriptions",
			"All enums have descriptions",
			"API has a general description",
			"Contact information is provided",
			"License information is provided",
		},
	},
	{
		ID:          "P003",
		Name:        "Error Handling Standards",
		Description: "Validates comprehensive error response documentation and consistent error handling patterns",
		Category:    "Error Handling",
		Severity:    "critical",
		Tags:        []string{"errors", "responses", "standards"},
		AutoFixable: true,
		Checks: []string{
			"All operations document 4xx error responses",
			"All operations document 5xx error responses",
			"Error responses include error codes",
			"Error responses include error messages",
			"Error responses include error details schema",
			"Common error responses are defined in components",
			"Error responses follow consistent format",
		},
	},
	{
		ID:          "P004",
		Name:        "Request Schema Definitions",
		Description: "Ensures all API requests have comprehensive schema definitions with proper data types, validation rules, and constraints",
		Category:    "Schema",
		Severity:    "critical",
		Tags:        []string{"schema", "validation", "request"},
		AutoFixable: true,
		Checks: []string{
			"All path parameters have schemas",
			"All query parameters have schemas",
			"All header parameters have schemas",
			"All request bodies have content schemas",
			"All schemas specify data types",
			"All schemas have appropriate constraints",
			"All required fields are marked",
			"All enums have valid values",
			"All numeric fields have min/max values",
			"All string fields have length constraints",
		},
	},
	{
		ID:          "P005",
		Name:        "Security Standards",
		Description: "Validates comprehensive security requirements and authentication mechanisms",
		Category:    "Security",
		Severity:    "critical",
		Tags:        []string{"security", "authentication", "authorization"},
		AutoFixable: false,
		Checks: []string{
			"Security schemes are defined",
			"Global security requirements are set",
			"Operation-level security is defined",
			"OAuth2 scopes are documented",
			"API keys are properly described",
			"Authentication headers are specified",
			"Security requirements are consistent",
		},
	},
	{
		ID:          "P006",
		Name:        "API Contract Testing",
		Description: "Validates that the API implementation matches its specification through functional testing",
		Category:    "Testing",
		Severity:    "critical",
		Tags:        []string{"testing", "contract", "implementation"},
		AutoFixable: false,
		Checks: []string{
			"All endpoints are reachable",
			"Response status codes match documentation",
			"Response schemas match documentation",
			"Authentication works as documented",
			"Required parameters are enforced",
			"Request validation works as documented",
			"Error responses match documentation",
		},
	},
	{
		ID:          "P007",
		Name:        "Performance Requirements",
		Description: "Validates that the API meets performance targets and SLAs",
		Category:    "Performance",
		Severity:    "warning",
		Tags:        []string{"performance", "sla", "load-testing"},
		AutoFixable: false,
		Checks: []string{
			"Response time meets targets",
			"Success rate meets targets",
			"Error rate is within limits",
			"Latency percentiles are acceptable",
			"Throughput meets requirements",
			"Concurrent request handling",
			"Resource utilization is acceptable",
		},
	},
	{
		ID:          "P008",
		Name:        "API Versioning Strategy",
		Description: "Validates proper API versioning implementation and documentation",
		Category:    "Versioning",
		Severity:    "warning",
		Tags:        []string{"versioning", "compatibility", "lifecycle"},
		AutoFixable: true,
		Checks: []string{
			"API version is specified",
			"Version follows semantic versioning",
			"Versioning strategy is documented",
			"Deprecation notices are present",
			"Breaking changes are documented",
			"Version compatibility is specified",
			"Migration guides are referenced",
		},
	},
	{
		ID:          "P009",
		Name:        "Test Readiness",
		Description: "Validates that the specification provides enough detail for meaningful functional and load testing",
		Category:    "Testing",
		Severity:    "warning",
		Tags:        []string{"testing", "readiness", "examples", "schemas"},
		AutoFixable: false,
		Checks: []string{
			"Path parameters have examples or typed schemas",
			"Request bodies have examples or typed schema properties",
			"2xx responses have schemas defined",
			"2xx responses have examples defined",
		},
	},
}

// GateContext provides quality gate metadata for PR comments in XSDLC workflows.
// When present, comments include gate-specific headers, check details, and "How to Pass" guidance.
type GateContext struct {
	GateName       string `json:"gate_name"`       // e.g., "staging-gate"
	Environment    string `json:"environment"`      // e.g., "staging"
	AppName        string `json:"app_name"`         // e.g., "perfect-api"
	CheckTypes     string `json:"check_types"`      // e.g., "validate-only,functional-test"
	ValidationMode string `json:"validation_mode"`  // e.g., "strict"
	WorkflowURL    string `json:"workflow_url"`      // full URL to the specific Argo Workflow run
}

// ValidatorConfig holds configuration for the validator
type ValidatorConfig struct {
	BaseURL           string
	SpecPath          string
	Environment       string
	Version           string
	Timeout           time.Duration
	ValidationMode    ValidationMode
	Auth              *AuthConfig // Auth field for token support
	PerformanceTarget *PerformanceTargetConfig
	Retries           int // Number of retries for transient failures (default: 2)
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

// ValidateAuthConfig validates authentication configuration
func ValidateAuthConfig(auth *AuthConfig) error {
	if auth == nil {
		return nil // No auth is valid
	}

	// Count authentication methods
	authMethods := 0
	if auth.Token != "" {
		authMethods++
	}
	if auth.APIKey != "" {
		authMethods++
	}
	if auth.Username != "" {
		authMethods++
	}

	// Only one authentication method should be used
	if authMethods > 1 {
		return fmt.Errorf("only one authentication method can be specified (token, api-key, or username/password)")
	}

	// Validate specific auth methods
	if auth.Token != "" {
		if auth.TokenType == "" {
			auth.TokenType = "Bearer"
		}
		if auth.TokenHeader == "" {
			auth.TokenHeader = "Authorization"
		}
	} else if auth.APIKey != "" {
		if auth.APIKeyHeader == "" {
			auth.APIKeyHeader = "X-API-Key"
		}
	} else if auth.Username != "" {
		if auth.Password == "" {
			return fmt.Errorf("password is required when username is specified")
		}
	}

	return nil
}

// GetAuthMethod returns the authentication method being used
func GetAuthMethod(auth *AuthConfig) string {
	if auth == nil {
		return "none"
	}
	if auth.Token != "" {
		return "token"
	}
	if auth.APIKey != "" {
		return "api-key"
	}
	if auth.Username != "" {
		return "basic"
	}
	return "none"
}

// ValidateConfig validates the validator configuration
func ValidateConfig(config ValidatorConfig) error {
	if config.SpecPath == "" {
		return fmt.Errorf("spec path is required")
	}
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}
	if config.Timeout <= 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}
	if config.Auth != nil {
		authMethods := 0
		if config.Auth.Token != "" {
			authMethods++
		}
		if config.Auth.APIKey != "" {
			authMethods++
		}
		if config.Auth.Username != "" {
			authMethods++
		}
		if authMethods > 1 {
			return fmt.Errorf("only one authentication method can be specified")
		}
	}
	if config.PerformanceTarget != nil {
		if config.PerformanceTarget.Duration <= 0 {
			return fmt.Errorf("performance test duration must be greater than 0")
		}
		if config.PerformanceTarget.ConcurrentUsers <= 0 {
			return fmt.Errorf("concurrent users must be greater than 0")
		}
		if config.PerformanceTarget.MinSuccessRate < 0 || config.PerformanceTarget.MinSuccessRate > 1 {
			return fmt.Errorf("minimum success rate must be between 0 and 1")
		}
	}
	return nil
}
