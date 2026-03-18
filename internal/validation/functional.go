package validation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"encoding/base64"

	"github.com/meter-peter/driveby/internal/openapi"
	"github.com/meter-peter/driveby/internal/spec"
)

// FunctionalTester handles functional testing of API endpoints
type FunctionalTester struct {
	config ValidatorConfig
	loader *openapi.Loader
	client *http.Client
}

// NewFunctionalTester creates a new functional tester instance
func NewFunctionalTester(config ValidatorConfig) *FunctionalTester {
	if config.Timeout == 0 {
		config.Timeout = 5 * time.Second // Default timeout if not specified
	}
	return &FunctionalTester{
		config: config,
		loader: openapi.NewLoader(),
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// TestEndpoints runs functional tests against all endpoints
func (t *FunctionalTester) TestEndpoints(ctx context.Context) (*ValidationReport, error) {
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
	principleResult := PrincipleResult{
		Principle: CorePrinciples[5], // P006: Endpoint Functional Testing
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
	var endpointResults []EndpointTestResult
	for _, epVal := range endpointResult.Endpoints {
		status := TestStatusPassed
		if epVal.Status != "success" {
			status = TestStatusFailed
		}
		
		endpointResults = append(endpointResults, EndpointTestResult{
			Method:       epVal.Method,
			Path:         epVal.Path,
			Status:       status,
			StatusCode:   epVal.StatusCode,
			ResponseTime: epVal.ResponseTime,
			Errors:       epVal.Errors,
		})
	}

	// Create functional test results
	functionalResults := &FunctionalTestResults{
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

	report := &ValidationReport{
		Version:      t.config.Version,
		Environment:  t.config.Environment,
		Timestamp:    time.Now(),
		Principles:   []PrincipleResult{principleResult},
		TotalChecks:  1,
		PassedChecks: 0,
		FailedChecks: 0,
		TestResults: &TestResults{
			Functional: functionalResults,
			StartTime:  time.Now(),
			EndTime:    time.Now(),
			Status:     TestStatusPassed,
		},
	}
	if allSuccess {
		report.PassedChecks = 1
		report.TestResults.Status = TestStatusPassed
	} else {
		report.FailedChecks = 1
		report.TestResults.Status = TestStatusFailed
	}

	return report, nil
}

// Helper functions for calculating test statistics
func countSuccessfulEndpoints(endpoints []EndpointValidation) int {
	count := 0
	for _, ep := range endpoints {
		if ep.Status == "success" {
			count++
		}
	}
	return count
}

func countFailedEndpoints(endpoints []EndpointValidation) int {
	count := 0
	for _, ep := range endpoints {
		if ep.Status != "success" {
			count++
		}
	}
	return count
}

func calculateAverageResponseTime(endpoints []EndpointValidation) time.Duration {
	if len(endpoints) == 0 {
		return 0
	}
	
	total := time.Duration(0)
	for _, ep := range endpoints {
		total += ep.ResponseTime
	}
	return total / time.Duration(len(endpoints))
}

func calculateMaxResponseTime(endpoints []EndpointValidation) time.Duration {
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

func calculateMinResponseTime(endpoints []EndpointValidation) time.Duration {
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

// validateEndpoints tests each endpoint in the API spec.
func (t *FunctionalTester) validateEndpoints(ctx context.Context, doc spec.APISpec) (*EndpointValidationResult, error) {
	result := &EndpointValidationResult{}

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
				result.Endpoints = append(result.Endpoints, EndpointValidation{
					Method: method,
					Path:   path,
					Status: "error",
					Errors: []string{fmt.Sprintf("Failed to create request: %v", err)},
				})
				continue
			}

			// Add authentication if configured
			if t.config.Auth != nil {
				if err := t.addAuthHeaders(req); err != nil {
					result.Endpoints = append(result.Endpoints, EndpointValidation{
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
			resp, err := t.client.Do(req)
			responseTime := time.Since(startTime)

			validation := EndpointValidation{
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

// addAuthHeaders adds authentication headers to the request based on the configured auth method
func (t *FunctionalTester) addAuthHeaders(req *http.Request) error {
	if t.config.Auth == nil {
		return nil
	}

	// Only one authentication method should be used
	authMethods := 0
	if t.config.Auth.Token != "" {
		authMethods++
	}
	if t.config.Auth.APIKey != "" {
		authMethods++
	}
	if t.config.Auth.Username != "" {
		authMethods++
	}
	if authMethods > 1 {
		return fmt.Errorf("only one authentication method can be specified")
	}

	// Add the appropriate auth header
	if t.config.Auth.Token != "" {
		headerName := t.config.Auth.TokenHeader
		if headerName == "" {
			headerName = "Authorization"
		}
		tokenType := t.config.Auth.TokenType
		if tokenType == "" {
			tokenType = "Bearer"
		}
		req.Header.Set(headerName, fmt.Sprintf("%s %s", tokenType, t.config.Auth.Token))
	} else if t.config.Auth.APIKey != "" {
		headerName := t.config.Auth.APIKeyHeader
		if headerName == "" {
			headerName = "X-API-Key"
		}
		req.Header.Set(headerName, t.config.Auth.APIKey)
	} else if t.config.Auth.Username != "" {
		// Basic auth
		auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.config.Auth.Username, t.config.Auth.Password)))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
	}

	return nil
}
