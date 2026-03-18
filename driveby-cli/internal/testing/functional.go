package testing

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/meter-peter/driveby/driveby-cli/internal/loader"
	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// FunctionalTester handles functional testing of API endpoints
type FunctionalTester struct {
	config  types.ValidatorConfig
	loader  *loader.Loader
	client  *http.Client
	retries int
}

// NewFunctionalTester creates a new functional tester instance
func NewFunctionalTester(config types.ValidatorConfig) *FunctionalTester {
	if config.Timeout == 0 {
		config.Timeout = 5 * time.Second
	}
	retries := 2
	if config.Retries > 0 {
		retries = config.Retries
	}
	return &FunctionalTester{
		config:  config,
		loader:  loader.NewLoader(),
		client:  &http.Client{Timeout: config.Timeout},
		retries: retries,
	}
}

// TestEndpoints runs functional tests against all endpoints
func (t *FunctionalTester) TestEndpoints(ctx context.Context) (*types.ValidationReport, error) {
	// Load OpenAPI spec
	if err := t.loader.LoadFromFileOrURL(t.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := t.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	// Test all endpoints
	endpointResult, err := t.validateEndpoints(ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("endpoint functional testing failed: %w", err)
	}

	// Analyze results
	allSuccess := true
	var failedEndpoints []string
	for _, epVal := range endpointResult.Endpoints {
		if epVal.Status != "success" {
			allSuccess = false
			failedEndpoints = append(failedEndpoints, fmt.Sprintf("%s %s (Status: %s, Code: %d)", epVal.Method, epVal.Path, epVal.Status, epVal.StatusCode))
		}
	}

	// Create report
	principleResult := types.PrincipleResult{
		Principle: types.CorePrinciples[5], // P006: Endpoint Functional Testing
		Passed:    allSuccess,
		Details:   endpointResult.Endpoints,
	}
	if allSuccess {
		principleResult.Message = "All documented endpoints are reachable and return documented status codes."
	} else {
		principleResult.Message = fmt.Sprintf("Some endpoints failed functional tests. Failed: %d/%d", len(failedEndpoints), len(endpointResult.Endpoints))
		principleResult.Details = map[string]interface{}{"failed_endpoints": failedEndpoints, "all_results": endpointResult.Endpoints}
	}

	// Convert EndpointValidation to EndpointTestResult
	var endpointResults []types.EndpointTestResult
	for _, epVal := range endpointResult.Endpoints {
		status := types.TestStatusPassed
		if epVal.Status != "success" {
			status = types.TestStatusFailed
		}

		endpointResults = append(endpointResults, types.EndpointTestResult{
			Method:       epVal.Method,
			Path:         epVal.Path,
			Status:       status,
			StatusCode:   epVal.StatusCode,
			ResponseTime: epVal.ResponseTime,
			Errors:       epVal.Errors,
		})
	}

	// Create functional test results
	functionalResults := &types.FunctionalTestResults{
		TotalEndpoints:      len(endpointResult.Endpoints),
		TestedEndpoints:     len(endpointResult.Endpoints),
		PassedEndpoints:     countSuccessfulEndpoints(endpointResult.Endpoints),
		FailedEndpoints:     countFailedEndpoints(endpointResult.Endpoints),
		SkippedEndpoints:    0,
		EndpointResults:     endpointResults,
		AverageResponseTime: calculateAverageResponseTime(endpointResult.Endpoints),
		MaxResponseTime:     calculateMaxResponseTime(endpointResult.Endpoints),
		MinResponseTime:     calculateMinResponseTime(endpointResult.Endpoints),
	}

	report := &types.ValidationReport{
		Version:     t.config.Version,
		Environment: t.config.Environment,
		Timestamp:   time.Now(),
		Principles:  []types.PrincipleResult{principleResult},
		TotalChecks: 1,
		PassedChecks: 0,
		FailedChecks: 0,
		TestResults: &types.TestResults{
			Functional: functionalResults,
			StartTime:  time.Now(),
			EndTime:    time.Now(),
			Status:     types.TestStatusPassed,
		},
	}
	if allSuccess {
		report.PassedChecks = 1
		report.TestResults.Status = types.TestStatusPassed
	} else {
		report.FailedChecks = 1
		report.TestResults.Status = types.TestStatusFailed
	}

	return report, nil
}

// Helper functions for calculating test statistics
func countSuccessfulEndpoints(endpoints []types.EndpointValidation) int {
	count := 0
	for _, ep := range endpoints {
		if ep.Status == "success" {
			count++
		}
	}
	return count
}

func countFailedEndpoints(endpoints []types.EndpointValidation) int {
	count := 0
	for _, ep := range endpoints {
		if ep.Status != "success" {
			count++
		}
	}
	return count
}

func calculateAverageResponseTime(endpoints []types.EndpointValidation) time.Duration {
	if len(endpoints) == 0 {
		return 0
	}

	total := time.Duration(0)
	for _, ep := range endpoints {
		total += ep.ResponseTime
	}
	return total / time.Duration(len(endpoints))
}

func calculateMaxResponseTime(endpoints []types.EndpointValidation) time.Duration {
	if len(endpoints) == 0 {
		return 0
	}

	max := endpoints[0].ResponseTime
	for _, ep := range endpoints {
		if ep.ResponseTime > max {
			max = ep.ResponseTime
		}
	}
	return max
}

func calculateMinResponseTime(endpoints []types.EndpointValidation) time.Duration {
	if len(endpoints) == 0 {
		return 0
	}

	min := endpoints[0].ResponseTime
	for _, ep := range endpoints {
		if ep.ResponseTime < min {
			min = ep.ResponseTime
		}
	}
	return min
}

// isRetryable returns true for transient HTTP errors worth retrying.
func isRetryable(resp *http.Response) bool {
	return resp.StatusCode == 502 || resp.StatusCode == 503 || resp.StatusCode == 504
}

// doWithRetry performs the HTTP request with exponential backoff on transient failures.
func (t *FunctionalTester) doWithRetry(req *http.Request) (*http.Response, error) {
	var lastErr error
	backoff := 500 * time.Millisecond

	for attempt := 0; attempt <= t.retries; attempt++ {
		if attempt > 0 {
			time.Sleep(backoff)
			backoff *= 2
		}
		resp, err := t.client.Do(req)
		if err != nil {
			lastErr = err
			continue // connection error — retry
		}
		if isRetryable(resp) {
			io.ReadAll(resp.Body)
			resp.Body.Close()
			lastErr = fmt.Errorf("transient HTTP %d", resp.StatusCode)
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("after %d retries: %w", t.retries, lastErr)
}

// validateEndpoints tests each endpoint in the API spec.
func (t *FunctionalTester) validateEndpoints(ctx context.Context, doc spec.APISpec) (*types.EndpointValidationResult, error) {
	result := &types.EndpointValidationResult{}

	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil || operation.Deprecated {
				continue
			}

			url := fmt.Sprintf("%s%s", t.config.BaseURL, path)
			req, err := http.NewRequestWithContext(ctx, method, url, nil)
			if err != nil {
				result.Endpoints = append(result.Endpoints, types.EndpointValidation{
					Method: method,
					Path:   path,
					Status: "error",
					Errors: []string{fmt.Sprintf("Failed to create request: %v", err)},
				})
				continue
			}

			// Add authentication if configured
			if t.config.Auth != nil {
				if err := AddAuthHeaders(req, t.config.Auth); err != nil {
					result.Endpoints = append(result.Endpoints, types.EndpointValidation{
						Method: method,
						Path:   path,
						Status: "error",
						Errors: []string{fmt.Sprintf("Failed to add authentication: %v", err)},
					})
					continue
				}
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")

			startTime := time.Now()
			resp, err := t.doWithRetry(req)
			responseTime := time.Since(startTime)

			validation := types.EndpointValidation{
				Method:       method,
				Path:         path,
				ResponseTime: responseTime,
			}

			if err != nil {
				validation.Status = "error"
				validation.Errors = []string{fmt.Sprintf("Request failed: %v", err)}
			} else {
				// Ensure response body is closed
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					validation.Status = "error"
					validation.Errors = []string{fmt.Sprintf("Failed to read response body: %v", err)}
				} else {
					validation.StatusCode = resp.StatusCode
					validation.ResponseBody = body

					// Check if status code is documented
					if operation.Responses != nil {
						if _, documented := operation.Responses[fmt.Sprintf("%d", resp.StatusCode)]; documented {
							// If documented, it's a success regardless of status code
							validation.Status = "success"
						} else {
							validation.Status = "warning"
							validation.Errors = []string{fmt.Sprintf("Status code %d is not documented in the OpenAPI spec", resp.StatusCode)}
						}
					} else {
						validation.Status = "warning"
						validation.Errors = []string{fmt.Sprintf("Status code %d is not documented in the OpenAPI spec", resp.StatusCode)}
					}
				}
			}

			result.Endpoints = append(result.Endpoints, validation)
		}
	}

	return result, nil
}
