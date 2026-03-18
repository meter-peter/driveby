package testing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
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

	// Analyze results — skipped endpoints are neither success nor failure
	allSuccess := true
	var failedEndpoints []string
	for _, epVal := range endpointResult.Endpoints {
		if epVal.Status != "success" && epVal.Status != "skipped" {
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
		switch epVal.Status {
		case "success":
			status = types.TestStatusPassed
		case "skipped":
			status = types.TestStatusSkipped
		default:
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
		SkippedEndpoints:    countSkippedEndpoints(endpointResult.Endpoints),
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
		if ep.Status != "success" && ep.Status != "skipped" {
			count++
		}
	}
	return count
}

func countSkippedEndpoints(endpoints []types.EndpointValidation) int {
	count := 0
	for _, ep := range endpoints {
		if ep.Status == "skipped" {
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

// pathParamRegex matches {paramName} placeholders in URL paths.
var pathParamRegex = regexp.MustCompile(`\{([^}]+)\}`)

// SubstitutePathParams replaces {paramName} placeholders with concrete values
// from the operation's parameter examples or generated defaults.
func SubstitutePathParams(path string, operation *spec.Operation) string {
	return pathParamRegex.ReplaceAllStringFunc(path, func(match string) string {
		paramName := match[1 : len(match)-1] // strip { and }

		// Find matching path parameter
		for _, p := range operation.Parameters {
			if p.In == "path" && p.Name == paramName {
				// Priority 1: parameter-level example
				if p.Example != nil {
					return fmt.Sprintf("%v", p.Example)
				}
				// Priority 2: generate from schema type
				if p.Schema != nil {
					return generatePathValue(p.Schema)
				}
			}
		}

		// Fallback
		return "test-value"
	})
}

// generatePathValue creates a sensible default value for a path parameter based on schema type.
func generatePathValue(s *spec.Schema) string {
	switch s.Type {
	case "string":
		if s.Format == "uuid" {
			return "f7cfc49d-824b-4728-a4c4-45e5901e3d42"
		}
		return "test-value"
	case "integer":
		return "1"
	case "number":
		return "1"
	default:
		return "test-value"
	}
}

// BuildRequestBody constructs a JSON request body from the operation's request body definition.
// Returns nil if the operation has no request body.
func BuildRequestBody(operation *spec.Operation) (io.Reader, error) {
	if operation.RequestBody == nil {
		return nil, nil
	}

	jsonContent, ok := operation.RequestBody.Content["application/json"]
	if !ok {
		return nil, nil
	}

	var bodyData interface{}

	// Priority 1: MediaType-level example
	if jsonContent.Example != nil {
		bodyData = jsonContent.Example
	}

	// Priority 2: First of MediaType-level examples
	if bodyData == nil && len(jsonContent.Examples) > 0 {
		for _, ex := range jsonContent.Examples {
			bodyData = ex
			break
		}
	}

	// Priority 3: Generate from schema
	if bodyData == nil && jsonContent.Schema != nil {
		bodyData = GenerateExampleFromSchema(jsonContent.Schema)
	}

	if bodyData == nil {
		return nil, nil
	}

	data, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	return bytes.NewReader(data), nil
}

// GenerateExampleFromSchema creates example data from a spec.Schema definition.
func GenerateExampleFromSchema(schema *spec.Schema) interface{} {
	if schema == nil {
		return nil
	}

	// If allOf, flatten first
	if len(schema.AllOf) > 0 {
		schema = spec.FlattenAllOf(schema)
	}

	switch schema.Type {
	case "string":
		return generateExampleString(schema.Format)
	case "integer":
		return 1
	case "number":
		return 1.0
	case "boolean":
		return true
	case "array":
		if schema.Items != nil {
			return []interface{}{GenerateExampleFromSchema(schema.Items)}
		}
		return []interface{}{}
	case "object", "":
		obj := make(map[string]interface{})
		for name, prop := range schema.Properties {
			obj[name] = GenerateExampleFromSchema(prop)
		}
		return obj
	default:
		return nil
	}
}

// generateExampleString creates a sensible default string based on format.
func generateExampleString(format string) string {
	switch format {
	case "uuid":
		return "f7cfc49d-824b-4728-a4c4-45e5901e3d42"
	case "date-time":
		return "2023-01-15T14:30:00Z"
	case "date":
		return "2023-01-15"
	case "email":
		return "test@example.com"
	case "uri", "url":
		return "https://example.com"
	default:
		return "example"
	}
}

// ValidateResponseBody performs lightweight schema validation on the response body.
// It returns a list of warning strings (empty means valid).
func ValidateResponseBody(body []byte, operation *spec.Operation, statusCode int) []string {
	if len(body) == 0 || operation.Responses == nil {
		return nil
	}

	statusStr := fmt.Sprintf("%d", statusCode)
	resp, ok := operation.Responses[statusStr]
	if !ok {
		return nil
	}

	jsonContent, ok := resp.Content["application/json"]
	if !ok || jsonContent.Schema == nil {
		return nil
	}

	schema := jsonContent.Schema
	if len(schema.AllOf) > 0 {
		schema = spec.FlattenAllOf(schema)
	}

	var parsed interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return []string{fmt.Sprintf("Response body is not valid JSON: %v", err)}
	}

	var warnings []string

	// Type check
	switch schema.Type {
	case "object":
		obj, ok := parsed.(map[string]interface{})
		if !ok {
			warnings = append(warnings, fmt.Sprintf("Expected object response, got %T", parsed))
			return warnings
		}
		// Check required fields
		for _, req := range schema.Required {
			if _, exists := obj[req]; !exists {
				warnings = append(warnings, fmt.Sprintf("Missing required field in response: %s", req))
			}
		}
	case "array":
		if _, ok := parsed.([]interface{}); !ok {
			warnings = append(warnings, fmt.Sprintf("Expected array response, got %T", parsed))
		}
	}

	return warnings
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

			// Substitute path parameters with concrete values
			resolvedPath := SubstitutePathParams(path, operation)
			url := fmt.Sprintf("%s%s", t.config.BaseURL, resolvedPath)

			// Build request body for POST/PUT/PATCH
			body, err := BuildRequestBody(operation)
			if err != nil {
				result.Endpoints = append(result.Endpoints, types.EndpointValidation{
					Method: method,
					Path:   path,
					Status: "error",
					Errors: []string{fmt.Sprintf("Failed to build request body: %v", err)},
				})
				continue
			}

			req, err := http.NewRequestWithContext(ctx, method, url, body)
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
				respBody, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					validation.Status = "error"
					validation.Errors = []string{fmt.Sprintf("Failed to read response body: %v", err)}
				} else {
					validation.StatusCode = resp.StatusCode
					validation.ResponseBody = respBody

					// Auth-awareness: if we got 401 without auth config, mark as skipped
					if resp.StatusCode == 401 && t.config.Auth == nil {
						validation.Status = "skipped"
						validation.Errors = []string{"Endpoint requires authentication (got 401 without auth config)"}
					} else if operation.Responses != nil {
						// Check if status code is documented
						if _, documented := operation.Responses[fmt.Sprintf("%d", resp.StatusCode)]; documented {
							validation.Status = "success"

							// Response schema validation for successful responses
							warnings := ValidateResponseBody(respBody, operation, resp.StatusCode)
							if len(warnings) > 0 {
								validation.Errors = append(validation.Errors, warnings...)
							}
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
