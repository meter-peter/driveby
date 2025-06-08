// Package validation provides orchestration for OpenAPI validation phases.
// This file (orchestrator.go) contains the orchestrator logic for API validation, functional testing, and performance testing.
// Stateless validation functions are in stateless.go.

package validation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	"github.com/sirupsen/logrus"
)

// APIValidator implements the validation logic
type APIValidator struct {
	config    ValidatorConfig
	logger    *Logger
	loader    *openapi.Loader
	client    *http.Client
	baseURL   string
	validator *OpenAPIValidator
}

// NewAPIValidator creates a new validator instance
func NewAPIValidator(config ValidatorConfig) (*APIValidator, error) {
	logger, err := NewLogger("stdout") // Force stdout for Kubernetes environment
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	// Set up HTTP client
	client := &http.Client{
		Timeout: config.Timeout,
	}

	validator, err := NewOpenAPIValidator(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAPI validator: %w", err)
	}

	return &APIValidator{
		config:    config,
		logger:    logger,
		loader:    openapi.NewLoader(),
		client:    client,
		baseURL:   config.BaseURL,
		validator: validator,
	}, nil
}

// Validate runs the complete validation suite
func (v *APIValidator) Validate(ctx context.Context) (*ValidationReport, error) {
	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// Use the OpenAPI validator to validate the spec
	validationReport, err := v.validator.ValidateSpec(ctx)
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
		validator, err := NewOpenAPIValidator(o.config)
		if err != nil {
			return nil, fmt.Errorf("failed to create validator: %w", err)
		}
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
		tester, err := NewPerformanceTester(o.config)
		if err != nil {
			return nil, fmt.Errorf("failed to create performance tester: %w", err)
		}
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
	if specResult.Error != nil {
		return nil, fmt.Errorf("spec validation error: %w", specResult.Error)
	}
	results = append(results, *specResult)

	// Run functional tests
	funcResult, err := o.RunValidation(ctx, ValidationTypeFunctional)
	if err != nil {
		return nil, fmt.Errorf("functional testing failed: %w", err)
	}
	if funcResult.Error != nil {
		return nil, fmt.Errorf("functional testing error: %w", funcResult.Error)
	}
	results = append(results, *funcResult)

	// Run performance tests
	perfResult, err := o.RunValidation(ctx, ValidationTypePerformance)
	if err != nil {
		return nil, fmt.Errorf("performance testing failed: %w", err)
	}
	if perfResult.Error != nil {
		return nil, fmt.Errorf("performance testing error: %w", perfResult.Error)
	}
	results = append(results, *perfResult)

	return results, nil
}
