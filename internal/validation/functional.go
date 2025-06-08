package validation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/openapi"
)

// FunctionalTester handles functional testing of API endpoints
type FunctionalTester struct {
	config ValidatorConfig
	loader *openapi.Loader
	client *http.Client
}

// NewFunctionalTester creates a new functional tester instance
func NewFunctionalTester(config ValidatorConfig) *FunctionalTester {
	return &FunctionalTester{
		config: config,
		loader: openapi.NewLoader(),
		client: &http.Client{
			Timeout: 5 * time.Second,
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

	report := &ValidationReport{
		Version:      t.config.Version,
		Environment:  t.config.Environment,
		Timestamp:    time.Now(),
		Principles:   []PrincipleResult{principleResult},
		TotalChecks:  1,
		PassedChecks: 0,
		FailedChecks: 0,
	}
	if allSuccess {
		report.PassedChecks = 1
	} else {
		report.FailedChecks = 1
	}

	return report, nil
}

// validateEndpoints tests each endpoint in the OpenAPI spec
func (t *FunctionalTester) validateEndpoints(ctx context.Context, doc *openapi3.T) (*EndpointValidationResult, error) {
	result := &EndpointValidationResult{}

	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if operation.Deprecated {
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
				defer resp.Body.Close()
				validation.StatusCode = resp.StatusCode

				_, documented := operation.Responses.Map()[fmt.Sprintf("%d", resp.StatusCode)]
				if !documented {
					validation.Status = "warning"
					validation.Errors = []string{fmt.Sprintf("Status code %d is not documented in the OpenAPI spec", resp.StatusCode)}
				} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {
					validation.Status = "success"
				} else {
					validation.Status = "error"
					validation.Errors = []string{fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)}
				}
			}

			result.Endpoints = append(result.Endpoints, validation)
		}
	}

	return result, nil
}
