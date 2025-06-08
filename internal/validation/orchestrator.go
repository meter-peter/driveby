// Package validation provides orchestration for OpenAPI validation phases.
// This file (orchestrator.go) contains the orchestrator logic for API validation, functional testing, and performance testing.
// Stateless validation functions are in stateless.go.

package validation

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	"github.com/sirupsen/logrus"
)

// APIValidator implements the validation logic
type APIValidator struct {
	config  ValidatorConfig
	logger  *Logger
	loader  *openapi.Loader
	client  *http.Client
	baseURL string
}

// NewAPIValidator creates a new validator instance
func NewAPIValidator(config ValidatorConfig) (*APIValidator, error) {
	logger, err := NewLogger(config.LogPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	// Set up HTTP client with authentication
	client := &http.Client{
		Timeout: config.Timeout,
		Transport: &authTransport{
			base: http.DefaultTransport,
			auth: config.Auth,
		},
	}

	// Determine the baseURL based on specPath
	baseURL := config.BaseURL
	logger.logger.Debugf("Initial specPath: %s, initial baseURL: %s", config.SpecPath, baseURL)
	if strings.HasPrefix(config.SpecPath, "http://") || strings.HasPrefix(config.SpecPath, "https://") {
		logger.logger.Debugf("specPath is a URL: %s", config.SpecPath)
		if strings.HasSuffix(strings.ToLower(config.SpecPath), ".json") || strings.HasSuffix(strings.ToLower(config.SpecPath), ".yaml") || strings.HasSuffix(strings.ToLower(config.SpecPath), ".yml") {
			logger.logger.Debug("specPath ends with .json/.yaml, extracting directory")
			lastSlash := strings.LastIndex(config.SpecPath, "/")
			if lastSlash != -1 {
				baseURL = config.SpecPath[:lastSlash]
				logger.logger.Debugf("Extracted baseURL from specPath: %s", baseURL)
			} else {
				baseURL = config.SpecPath
				logger.logger.Debugf("Could not find slash in specPath, using full specPath as baseURL: %s", baseURL)
			}
		} else {
			logger.logger.Debug("specPath is a URL but does not end with .json/.yaml, using provided config.BaseURL")
			baseURL = config.BaseURL
		}
	} else {
		logger.logger.Debug("specPath is not a URL, using provided config.BaseURL")
		baseURL = config.BaseURL
	}

	logger.logger.Debugf("Final baseURL determined: %s", baseURL)

	return &APIValidator{
		config:  config,
		logger:  logger,
		loader:  openapi.NewLoader(),
		client:  client,
		baseURL: baseURL,
	}, nil
}

// authTransport implements http.RoundTripper to add authentication headers
type authTransport struct {
	base http.RoundTripper
	auth AuthConfig
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.auth.Token != "" {
		headerName := t.auth.TokenHeader
		if headerName == "" {
			headerName = "Authorization"
		}

		tokenType := t.auth.TokenType
		if tokenType == "" {
			tokenType = "Bearer"
		}

		req.Header.Set(headerName, fmt.Sprintf("%s %s", tokenType, t.auth.Token))
	}

	return t.base.RoundTrip(req)
}

// Validate runs the complete validation suite
func (v *APIValidator) Validate(ctx context.Context) (*ValidationReport, error) {
	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// Load and validate OpenAPI spec
	if err := v.loader.LoadFromFileOrURL(v.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := v.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	// Run validation
	validator := NewOpenAPIValidator(v.config)
	validationReport, err := validator.ValidateSpec(ctx)
	if err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Copy validation results
	report.Principles = validationReport.Principles
	report.TotalChecks = validationReport.TotalChecks
	report.PassedChecks = validationReport.PassedChecks
	report.FailedChecks = validationReport.FailedChecks
	report.Summary = validationReport.Summary
	report.AutoFixes = validationReport.AutoFixes

	// Log the report
	if err := v.logger.LogReport(report); err != nil {
		return nil, fmt.Errorf("failed to log validation report: %w", err)
	}

	return report, nil
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[validation] Logger set to DEBUG (verbose) mode")
}

// ValidationReport represents a detailed report of validation results

// ValidationType represents the type of validation to run
type ValidationType string

const (
	ValidationTypeSpec        ValidationType = "spec"        // Only run OpenAPI spec validation (P001-P005, P008)
	ValidationTypeFunctional  ValidationType = "functional"  // Only run functional tests (P006)
	ValidationTypePerformance ValidationType = "performance" // Only run performance tests (P007)
)

// ValidationResult represents the result of a validation run
type ValidationResult struct {
	Type   ValidationType
	Report *ValidationReport
	Error  error
}

// Orchestrator handles running different types of validation
type Orchestrator struct {
	config ValidatorConfig
}

// NewOrchestrator creates a new orchestrator instance
func NewOrchestrator(config ValidatorConfig) *Orchestrator {
	return &Orchestrator{
		config: config,
	}
}

// RunValidation runs the specified type of validation
func (o *Orchestrator) RunValidation(ctx context.Context, validationType ValidationType) (*ValidationResult, error) {
	switch validationType {
	case ValidationTypeSpec:
		validator := NewOpenAPIValidator(o.config)
		report, err := validator.ValidateSpec(ctx)
		return &ValidationResult{
			Type:   ValidationTypeSpec,
			Report: report,
			Error:  err,
		}, nil

	case ValidationTypeFunctional:
		tester := NewFunctionalTester(o.config)
		report, err := tester.TestEndpoints(ctx)
		return &ValidationResult{
			Type:   ValidationTypeFunctional,
			Report: report,
			Error:  err,
		}, nil

	case ValidationTypePerformance:
		tester := NewPerformanceTester(o.config)
		report, err := tester.TestPerformance(ctx)
		return &ValidationResult{
			Type:   ValidationTypePerformance,
			Report: report,
			Error:  err,
		}, nil

	default:
		return nil, fmt.Errorf("unknown validation type: %s", validationType)
	}
}

// RunAllValidations runs all types of validation in sequence
func (o *Orchestrator) RunAllValidations(ctx context.Context) ([]ValidationResult, error) {
	results := make([]ValidationResult, 0, 3)

	// Run spec validation
	specResult, err := o.RunValidation(ctx, ValidationTypeSpec)
	if err != nil {
		return nil, fmt.Errorf("spec validation failed: %w", err)
	}
	results = append(results, *specResult)

	// Run functional tests
	funcResult, err := o.RunValidation(ctx, ValidationTypeFunctional)
	if err != nil {
		return nil, fmt.Errorf("functional testing failed: %w", err)
	}
	results = append(results, *funcResult)

	// Run performance tests
	perfResult, err := o.RunValidation(ctx, ValidationTypePerformance)
	if err != nil {
		return nil, fmt.Errorf("performance testing failed: %w", err)
	}
	results = append(results, *perfResult)

	return results, nil
}
