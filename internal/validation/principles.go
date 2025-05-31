package validation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

// Principle represents a validation principle that the API must adhere to
type Principle struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Severity    string   `json:"severity"` // "critical", "warning", "info"
	AutoFixable bool     `json:"auto_fixable"`
	Tags        []string `json:"tags"`
}

// ValidationReport represents a detailed report of validation results
type ValidationReport struct {
	Timestamp    time.Time         `json:"timestamp"`
	Version      string            `json:"version"`
	Environment  string            `json:"environment"`
	TotalChecks  int               `json:"total_checks"`
	PassedChecks int               `json:"passed_checks"`
	FailedChecks int               `json:"failed_checks"`
	Principles   []PrincipleResult `json:"principles"`
	AutoFixes    []AutoFixResult   `json:"auto_fixes"`
	Summary      ValidationSummary `json:"summary"`
}

// PrincipleResult represents the result of checking a single principle
type PrincipleResult struct {
	Principle    Principle `json:"principle"`
	Passed       bool      `json:"passed"`
	Message      string    `json:"message"`
	Location     string    `json:"location,omitempty"`
	Details      any       `json:"details,omitempty"`
	SuggestedFix string    `json:"suggested_fix,omitempty"`
}

// AutoFixResult represents the result of an automatic fix attempt
type AutoFixResult struct {
	PrincipleID string    `json:"principle_id"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	Location    string    `json:"location"`
	Original    any       `json:"original,omitempty"`
	Fixed       any       `json:"fixed,omitempty"`
	Error       string    `json:"error,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

// ValidationSummary provides a high-level overview of the validation results
type ValidationSummary struct {
	CriticalIssues int      `json:"critical_issues"`
	Warnings       int      `json:"warnings"`
	Info           int      `json:"info"`
	Categories     []string `json:"categories"`
	FailedTags     []string `json:"failed_tags"`
}

// Define our core principles
var CorePrinciples = []Principle{
	{
		ID:          "P001",
		Name:        "OpenAPI Specification Compliance",
		Description: "API must fully comply with OpenAPI 3.0 specification",
		Category:    "Documentation",
		Severity:    "critical",
		AutoFixable: true,
		Tags:        []string{"openapi", "spec", "documentation"},
	},
	{
		ID:          "P002",
		Name:        "Response Time Performance",
		Description: "API endpoints must respond within acceptable time limits",
		Category:    "Performance",
		Severity:    "critical",
		AutoFixable: false,
		Tags:        []string{"performance", "latency"},
	},
	{
		ID:          "P003",
		Name:        "Error Response Documentation",
		Description: "All endpoints must document possible error responses",
		Category:    "Documentation",
		Severity:    "warning",
		AutoFixable: true,
		Tags:        []string{"errors", "documentation"},
	},
	{
		ID:          "P004",
		Name:        "Request Validation",
		Description: "All endpoints must validate request parameters",
		Category:    "Security",
		Severity:    "critical",
		AutoFixable: true,
		Tags:        []string{"validation", "security"},
	},
	{
		ID:          "P005",
		Name:        "Authentication Requirements",
		Description: "Endpoints must specify authentication requirements",
		Category:    "Security",
		Severity:    "critical",
		AutoFixable: false,
		Tags:        []string{"security", "auth"},
	},
	{
		ID:          "P006",
		Name:        "Endpoint Functional Testing",
		Description: "API endpoints should be reachable and return documented responses",
		Category:    "Testing",
		Severity:    "critical",
		AutoFixable: false,
		Tags:        []string{"testing", "functional", "endpoints"},
	},
	{
		ID:          "P007",
		Name:        "API Performance Compliance",
		Description: "API performance metrics must meet defined targets",
		Category:    "Performance",
		Severity:    "critical",
		AutoFixable: false,
		Tags:        []string{"testing", "performance", "sla"},
	},
}

// Logger handles validation report logging
type Logger struct {
	logger *logrus.Logger
	path   string
}

// NewLogger creates a new validation logger
func NewLogger(logPath string) (*Logger, error) {
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
