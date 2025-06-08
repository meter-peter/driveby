package validation

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// PerformanceTester handles performance testing of API endpoints
type PerformanceTester struct {
	config ValidatorConfig
	loader *openapi.Loader
}

// NewPerformanceTester creates a new performance tester instance
func NewPerformanceTester(config ValidatorConfig) *PerformanceTester {
	return &PerformanceTester{
		config: config,
		loader: openapi.NewLoader(),
	}
}

// TestPerformance runs performance tests against all endpoints
func (t *PerformanceTester) TestPerformance(ctx context.Context) (*ValidationReport, error) {
	// Load OpenAPI spec
	if err := t.loader.LoadFromFileOrURL(t.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := t.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	// Create targets for all endpoints
	var targets []vegeta.Target
	for path, pathItem := range doc.Paths.Map() {
		for method := range pathItem.Operations() {
			targets = append(targets, vegeta.Target{
				Method: method,
				URL:    fmt.Sprintf("%s%s", t.config.BaseURL, path),
			})
		}
	}

	// Run performance tests
	perfResult, err := t.runPerformanceTests(targets)
	if err != nil {
		return nil, fmt.Errorf("performance testing failed: %w", err)
	}

	// Check against performance targets
	metTargets := true
	var failureReasons []string

	if perfResult.Performance.LatencyP95 > t.config.PerformanceTarget.MaxLatencyP95 && t.config.PerformanceTarget.MaxLatencyP95 != 0 {
		metTargets = false
		failureReasons = append(failureReasons, fmt.Sprintf("P95 latency (%s) exceeded target (%s)", perfResult.Performance.LatencyP95, t.config.PerformanceTarget.MaxLatencyP95))
	}
	if perfResult.Performance.ErrorRate > (1.0-t.config.PerformanceTarget.MinSuccessRate/100.0) && t.config.PerformanceTarget.MinSuccessRate != 0 {
		metTargets = false
		failureReasons = append(failureReasons, fmt.Sprintf("Error rate (%.2f%%) exceeded allowed (%.2f%% success rate target)", perfResult.Performance.ErrorRate*100, t.config.PerformanceTarget.MinSuccessRate))
	}

	// Create report
	principleResult := PrincipleResult{
		Principle: CorePrinciples[6], // P007: API Performance Compliance
		Passed:    metTargets,
		Details:   perfResult.Performance,
	}
	if metTargets {
		principleResult.Message = "API performance meets configured targets."
	} else {
		principleResult.Message = fmt.Sprintf("API performance failed to meet targets: %s", strings.Join(failureReasons, ", "))
	}

	report := &ValidationReport{
		Version:      t.config.Version,
		Environment:  t.config.Environment,
		Timestamp:    time.Now(),
		Principles:   []PrincipleResult{principleResult},
		TotalChecks:  1,
		PassedChecks: 0,
		FailedChecks: 0,
	}
	if metTargets {
		report.PassedChecks = 1
	} else {
		report.FailedChecks = 1
	}

	return report, nil
}

// runPerformanceTests executes a load test against the specified targets
func (t *PerformanceTester) runPerformanceTests(targets []vegeta.Target) (*PerformanceTestResult, error) {
	attacker := vegeta.NewAttacker()
	targeter := vegeta.NewStaticTargeter(targets...)
	pacer := vegeta.Rate{Freq: 50, Per: time.Second} // 50 requests per second
	duration := 10 * time.Second
	metrics := &vegeta.Metrics{}

	for res := range attacker.Attack(targeter, pacer, duration, "Performance Test") {
		metrics.Add(res)
	}
	metrics.Close()

	perfMetrics := &PerformanceMetrics{
		StartTime:      time.Now().Add(-duration),
		EndTime:        time.Now(),
		TotalRequests:  metrics.Requests,
		SuccessCount:   uint64(float64(metrics.Requests) * metrics.Success),
		ErrorCount:     uint64(float64(metrics.Requests) * (1 - metrics.Success)),
		ErrorRate:      1 - metrics.Success,
		LatencyP50:     metrics.Latencies.P50,
		LatencyP95:     metrics.Latencies.P95,
		LatencyP99:     metrics.Latencies.P99,
		RequestsPerSec: float64(metrics.Requests) / duration.Seconds(),
	}

	return &PerformanceTestResult{Performance: perfMetrics}, nil
}
