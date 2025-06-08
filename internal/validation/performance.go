package validation

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// PerformanceTester handles performance testing of API endpoints
type PerformanceTester struct {
	config  ValidatorConfig
	loader  *openapi.Loader
	metrics *vegeta.Metrics
	mu      sync.Mutex // Protect metrics access
}

// NewPerformanceTester creates a new performance tester instance
func NewPerformanceTester(config ValidatorConfig) (*PerformanceTester, error) {
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid validator config: %w", err)
	}
	return &PerformanceTester{
		config:  config,
		loader:  openapi.NewLoader(),
		metrics: &vegeta.Metrics{},
	}, nil
}

// cleanup releases any resources held by the tester
func (t *PerformanceTester) cleanup() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.metrics != nil {
		t.metrics.Close()
		t.metrics = nil
	}
	if t.loader != nil {
		t.loader = nil
	}
}

// TestPerformance runs performance tests against all endpoints
func (t *PerformanceTester) TestPerformance(ctx context.Context) (*ValidationReport, error) {
	defer t.cleanup()

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
			// Skip endpoints that are not suitable for load testing
			if method == "DELETE" || method == "PATCH" {
				continue
			}
			targets = append(targets, vegeta.Target{
				Method: method,
				URL:    fmt.Sprintf("%s%s", t.config.BaseURL, path),
			})
		}
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("no suitable endpoints found for load testing")
	}

	// Configure the attack
	rate := vegeta.Rate{
		Freq: t.config.PerformanceTarget.ConcurrentUsers,
		Per:  time.Second,
	}
	duration := t.config.PerformanceTarget.Duration
	if duration == 0 {
		duration = 5 * time.Minute // Default duration
	}

	attacker := vegeta.NewAttacker()
	targeter := vegeta.NewStaticTargeter(targets...)

	// Run the attack with context cancellation
	done := make(chan struct{})
	go func() {
		defer close(done)
		for res := range attacker.Attack(targeter, rate, duration, "DriveBy Load Test") {
			t.mu.Lock()
			t.metrics.Add(res)
			t.mu.Unlock()
		}
	}()

	// Wait for either context cancellation or attack completion
	select {
	case <-ctx.Done():
		attacker.Stop()
		return nil, ctx.Err()
	case <-done:
		// Attack completed normally
	}

	t.mu.Lock()
	t.metrics.Close()
	metrics := t.metrics
	t.metrics = nil // Prevent double close
	t.mu.Unlock()

	// Create performance report
	report := &ValidationReport{
		Version:     t.config.Version,
		Environment: t.config.Environment,
		Timestamp:   time.Now(),
		Principles: []PrincipleResult{
			{
				Principle: CorePrinciples[6], // P007: API Performance Compliance
				Passed:    true,
				Details: &PerformanceMetrics{
					StartTime:      time.Now().Add(-duration),
					EndTime:        time.Now(),
					TotalRequests:  metrics.Requests,
					SuccessCount:   metrics.Requests - uint64(len(metrics.Errors)),
					ErrorCount:     uint64(len(metrics.Errors)),
					ErrorRate:      float64(len(metrics.Errors)) / float64(metrics.Requests),
					LatencyP50:     metrics.Latencies.P50,
					LatencyP95:     metrics.Latencies.P95,
					LatencyP99:     metrics.Latencies.P99,
					RequestsPerSec: metrics.Rate,
				},
			},
		},
	}

	// Check against performance targets
	var failedChecks []string
	if t.config.PerformanceTarget.MaxLatencyP95 > 0 && metrics.Latencies.P95 > t.config.PerformanceTarget.MaxLatencyP95 {
		failedChecks = append(failedChecks, fmt.Sprintf("P95 latency (%s) exceeded target (%s)",
			metrics.Latencies.P95, t.config.PerformanceTarget.MaxLatencyP95))
	}

	successRate := 1.0 - (float64(len(metrics.Errors)) / float64(metrics.Requests))
	if t.config.PerformanceTarget.MinSuccessRate > 0 && successRate < t.config.PerformanceTarget.MinSuccessRate {
		failedChecks = append(failedChecks, fmt.Sprintf("Success rate (%.2f%%) below target (%.2f%%)",
			successRate*100, t.config.PerformanceTarget.MinSuccessRate*100))
	}

	if len(failedChecks) > 0 {
		report.Principles[0].Passed = false
		report.Principles[0].Message = strings.Join(failedChecks, "; ")
	} else {
		report.Principles[0].Message = "All performance targets met"
	}

	report.TotalChecks = 1
	if report.Principles[0].Passed {
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
