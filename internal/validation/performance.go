package validation

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// RunPerformance runs load/performance tests (P002, P007)
func RunPerformance(ctx context.Context, config ValidatorConfig) (*ValidationReport, error) {
	loader := openapi.NewLoader()
	if err := loader.LoadFromFileOrURL(config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	var targets []vegeta.Target
	for path, pathItem := range doc.Paths.Map() {
		for method := range pathItem.Operations() {
			targets = append(targets, vegeta.Target{
				Method: method,
				URL:    fmt.Sprintf("%s%s", config.BaseURL, path),
			})
		}
	}

	validator := NewValidator()
	validator.SetBaseURL(config.BaseURL)
	perfResult, err := validator.RunPerformanceTests(targets, 50.0, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("performance testing failed: %w", err)
	}

	metTargets := true
	var failureReasons []string

	if perfResult.Performance.LatencyP95 > config.PerformanceTarget.MaxLatencyP95 && config.PerformanceTarget.MaxLatencyP95 != 0 {
		metTargets = false
		failureReasons = append(failureReasons, fmt.Sprintf("P95 latency (%s) exceeded target (%s)", perfResult.Performance.LatencyP95, config.PerformanceTarget.MaxLatencyP95))
	}
	if perfResult.Performance.ErrorRate > (1.0-config.PerformanceTarget.MinSuccessRate/100.0) && config.PerformanceTarget.MinSuccessRate != 0 {
		metTargets = false
		failureReasons = append(failureReasons, fmt.Sprintf("Error rate (%.2f%%) exceeded allowed (%.2f%% success rate target)", perfResult.Performance.ErrorRate*100, config.PerformanceTarget.MinSuccessRate))
	}

	principleResult := PrincipleResult{
		Principle: Principle{
			ID:   "P007",
			Name: "API Performance Compliance",
		},
		Passed:  metTargets,
		Details: perfResult.Performance,
	}
	if metTargets {
		principleResult.Message = "API performance meets configured targets."
	} else {
		principleResult.Message = fmt.Sprintf("API performance failed to meet targets: %s", strings.Join(failureReasons, ", "))
	}

	report := &ValidationReport{
		Version:     config.Version,
		Environment: config.Environment,
		Timestamp:   time.Now(),
		Principles:  []PrincipleResult{principleResult},
	}
	return report, nil
}
