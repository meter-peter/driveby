package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/meter-peter/driveby/driveby-cli/internal/loader"
	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()

// Engine orchestrates validation, functional testing, and performance testing.
type Engine struct {
	config   types.ValidatorConfig
	loader   *loader.Loader
	registry *principles.Registry
}

// New creates a new Engine instance.
func New(config types.ValidatorConfig) (*Engine, error) {
	if err := types.ValidateConfig(config); err != nil {
		return nil, err
	}
	if config.ValidationMode == "" {
		config.ValidationMode = types.ValidationModeMinimal
	}
	return &Engine{
		config:   config,
		loader:   loader.NewLoader(),
		registry: principles.NewRegistry(),
	}, nil
}

// ValidateSpec runs spec validation (P001-P008) based on mode.
func (e *Engine) ValidateSpec(ctx context.Context) (*types.ValidationReport, error) {
	if err := e.loader.LoadFromFileOrURL(e.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load API spec: %w", err)
	}
	doc := e.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get API document")
	}

	report := &types.ValidationReport{
		Version:     e.config.Version,
		Environment: e.config.Environment,
		Timestamp:   time.Now(),
	}

	checkers := e.registry.ForMode(e.config.ValidationMode)
	for _, checker := range checkers {
		result := checker.Check(ctx, doc, e.config.ValidationMode)
		report.Principles = append(report.Principles, result)
		if result.Passed {
			report.PassedChecks++
		} else {
			report.FailedChecks++
		}
	}
	report.TotalChecks = len(checkers)
	updateSummary(report)

	return report, nil
}

// Validate runs spec validation wrapped through the logger (for backward compat with APIValidator).
func (e *Engine) Validate(ctx context.Context) (*types.ValidationReport, error) {
	report, err := e.ValidateSpec(ctx)
	if err != nil {
		return nil, err
	}

	logger, logErr := types.NewLogger("stdout")
	if logErr != nil {
		return nil, logErr
	}
	if err := logger.LogReport(report); err != nil {
		return nil, err
	}

	return report, nil
}

func updateSummary(report *types.ValidationReport) {
	summary := types.ValidationSummary{}
	categories := make(map[string]bool)
	failedTags := make(map[string]bool)

	for _, result := range report.Principles {
		if !result.Passed {
			switch result.Principle.Severity {
			case "critical":
				summary.CriticalIssues++
			case "warning":
				summary.Warnings++
			case "info":
				summary.Info++
			}
			categories[result.Principle.Category] = true
			for _, tag := range result.Principle.Tags {
				failedTags[tag] = true
			}
		}
	}

	for category := range categories {
		summary.Categories = append(summary.Categories, category)
	}
	for tag := range failedTags {
		summary.FailedTags = append(summary.FailedTags, tag)
	}
	report.Summary = summary
}
