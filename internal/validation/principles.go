package validation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

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
		Severity:    "warning",
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
		Severity:    "warning",
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
		Severity:    "warning",
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
}

// Logger handles validation report logging
type Logger struct {
	logger *logrus.Logger
	path   string
}

// NewLogger creates a new validation logger
func NewLogger(logPath string) (*Logger, error) {
	if logPath == "" {
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stdout)
		logger.Infof("[validation/principles] Logger set to DEBUG (verbose) mode (stdout fallback)")
		return &Logger{
			logger: logger,
			path:   "",
		}, nil
	}
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.Infof("[validation/principles] Logger set to DEBUG (verbose) mode")

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger.SetOutput(file)

	return &Logger{
		logger: logger,
		path:   logPath,
	}, nil
}

// LogReport logs a validation report
func (l *Logger) LogReport(report *ValidationReport) error {
	report.Timestamp = time.Now()

	entry := l.logger.WithFields(logrus.Fields{
		"type":      "validation_report",
		"version":   report.Version,
		"env":       report.Environment,
		"timestamp": report.Timestamp,
	})

	// Log the full report as JSON
	reportJSON, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	entry.Info(string(reportJSON))
	return nil
}

// GetRecentReports retrieves recent validation reports
func (l *Logger) GetRecentReports(limit int) ([]ValidationReport, error) {
	// Implementation to read and parse recent reports from the log file
	// This would be implemented based on your specific needs
	return nil, nil
}
