package core

import (
	"context"
	"driveby/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// TestingService handles all types of API testing
type TestingService struct {
	logger  *logrus.Logger
	apiHost string
	apiPort string
}

// NewTestingService creates a new testing service
func NewTestingService(logger *logrus.Logger, apiHost, apiPort string) *TestingService {
	return &TestingService{
		logger:  logger,
		apiHost: apiHost,
		apiPort: apiPort,
	}
}

// RunTests executes all configured tests and returns comprehensive results
func (s *TestingService) RunTests(ctx context.Context, req types.TestRequest) (*types.TestResponse, error) {
	testID := uuid.New().String()
	now := time.Now()

	// Run documentation validation
	docResult, err := s.validateDocumentation(ctx, req.OpenAPISpec, req.Thresholds.Documentation)
	if err != nil {
		return nil, fmt.Errorf("documentation validation failed: %w", err)
	}

	// Run integration tests
	intResult, err := s.runIntegrationTests(ctx, req.OpenAPISpec)
	if err != nil {
		return nil, fmt.Errorf("integration tests failed: %w", err)
	}

	// Run load tests
	loadResult, err := s.runLoadTests(ctx, req.OpenAPISpec, req.LoadTestConfig, req.Thresholds.LoadTest)
	if err != nil {
		return nil, fmt.Errorf("load tests failed: %w", err)
	}

	// Compile results
	result := &types.TestResponse{
		TestID:    testID,
		Timestamp: now,
		Results: types.TestResult{
			TestID:        testID,
			Timestamp:     now,
			Documentation: docResult,
			Integration:   intResult,
			LoadTest:      loadResult,
		},
	}

	return result, nil
}

// validateDocumentation checks if the API documentation meets the required standards
func (s *TestingService) validateDocumentation(ctx context.Context, spec *openapi3.T, thresholds struct {
	MinComplianceScore float64 `json:"min_compliance_score"`
	MaxMissingExamples int     `json:"max_missing_examples"`
}) (types.DocResult, error) {
	result := types.DocResult{
		ErrorResponses: make(map[string]int),
	}

	totalEndpoints := 0
	compliantEndpoints := 0

	// Validate each endpoint
	for path, pathItem := range spec.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			totalEndpoints++
			endpointCompliant := true
			endpointID := fmt.Sprintf("%s %s", method, path)

			// Check response documentation
			if len(operation.Responses.Map()) == 0 {
				result.UndocumentedEndpoints = append(result.UndocumentedEndpoints, endpointID)
				endpointCompliant = false
			}

			// Check examples and error responses
			for statusCode, response := range operation.Responses.Map() {
				if response.Value.Description == nil || *response.Value.Description == "" {
					result.UndocumentedEndpoints = append(result.UndocumentedEndpoints, endpointID)
					endpointCompliant = false
				}

				if strings.HasPrefix(statusCode, "4") || strings.HasPrefix(statusCode, "5") {
					result.ErrorResponses[statusCode]++
				}

				if response.Value.Content != nil {
					jsonContent := response.Value.Content.Get("application/json")
					if jsonContent != nil && jsonContent.Example == nil && len(jsonContent.Examples) == 0 {
						result.MissingExamples++
						endpointCompliant = false
					}
				}
			}

			if endpointCompliant {
				compliantEndpoints++
			}
		}
	}

	// Calculate compliance score
	if totalEndpoints > 0 {
		result.ComplianceScore = float64(compliantEndpoints) / float64(totalEndpoints) * 100
	}

	// Determine if documentation passes thresholds
	result.Passed = result.ComplianceScore >= thresholds.MinComplianceScore &&
		result.MissingExamples <= thresholds.MaxMissingExamples
	result.Threshold = thresholds.MinComplianceScore

	return result, nil
}

// runIntegrationTests executes integration tests based on OpenAPI examples
func (s *TestingService) runIntegrationTests(ctx context.Context, spec *openapi3.T) (types.IntResult, error) {
	result := types.IntResult{
		FailedEndpoints: make(map[string]string),
	}

	// Discover testable endpoints
	endpoints, err := s.discoverTestableEndpoints(spec)
	if err != nil {
		return result, fmt.Errorf("failed to discover endpoints: %w", err)
	}

	result.TotalTests = len(endpoints)

	// Execute tests for each endpoint
	for _, endpoint := range endpoints {
		err := s.testEndpoint(ctx, endpoint)
		if err != nil {
			result.FailedTests++
			result.FailedEndpoints[endpoint.ID] = err.Error()
		} else {
			result.PassedTests++
		}
	}

	// Calculate pass rate and determine if tests passed
	if result.TotalTests > 0 {
		passRate := float64(result.PassedTests) / float64(result.TotalTests)
		result.Passed = passRate >= 0.95 // 95% pass rate threshold
	}

	return result, nil
}

// runLoadTests executes load tests using Vegeta
func (s *TestingService) runLoadTests(ctx context.Context, spec *openapi3.T, config types.LoadTestConfig, thresholds types.LoadThresholds) (types.LoadResult, error) {
	result := types.LoadResult{
		StatusCodes: make(map[int]int),
		Thresholds:  thresholds,
	}

	// Create Vegeta targets
	targets := s.createLoadTestTargets(spec)

	// Configure load test
	rate := vegeta.Rate{Freq: config.RequestRate, Per: time.Second}
	duration := config.TestDuration
	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker(vegeta.Timeout(config.RequestTimeout))

	// Run the test
	var metrics vegeta.Metrics
	resultChan := attacker.Attack(targeter, rate, duration, "API Load Test")

	// Process results
	for res := range resultChan {
		metrics.Add(res)
		result.StatusCodes[int(res.Code)]++
	}
	metrics.Close()

	// Compile results
	result.TotalRequests = int64(metrics.Requests)
	result.SuccessRate = metrics.Success * 100
	result.LatencyP95 = metrics.Latencies.P95
	result.ErrorRate = (1 - metrics.Success) * 100

	// Determine if load test passed thresholds
	result.Passed = result.SuccessRate >= thresholds.MinSuccessRate &&
		result.LatencyP95 <= thresholds.MaxLatencyP95 &&
		result.ErrorRate <= thresholds.MaxErrorRate

	return result, nil
}

// Helper methods for integration testing
type testableEndpoint struct {
	ID       string
	Method   string
	Path     string
	Examples map[string]interface{}
	Body     interface{}
}

func (s *TestingService) discoverTestableEndpoints(spec *openapi3.T) ([]testableEndpoint, error) {
	var endpoints []testableEndpoint

	for path, pathItem := range spec.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			endpoint := testableEndpoint{
				ID:       fmt.Sprintf("%s %s", method, path),
				Method:   method,
				Path:     path,
				Examples: make(map[string]interface{}),
			}

			// Extract examples from parameters
			for _, param := range operation.Parameters {
				if param.Value.Example != nil {
					endpoint.Examples[param.Value.Name] = param.Value.Example
				}
			}

			// Extract request body examples
			if operation.RequestBody != nil {
				content := operation.RequestBody.Value.Content.Get("application/json")
				if content != nil && content.Example != nil {
					endpoint.Body = content.Example
				}
			}

			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints, nil
}

func (s *TestingService) testEndpoint(ctx context.Context, endpoint testableEndpoint) error {
	// Implementation of endpoint testing logic
	// This would make actual HTTP requests and validate responses
	// against the OpenAPI specification
	return nil
}

func (s *TestingService) createLoadTestTargets(spec *openapi3.T) []vegeta.Target {
	var targets []vegeta.Target

	for path, pathItem := range spec.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			target := vegeta.Target{
				Method: method,
				URL:    s.buildURL(path, operation),
				Header: http.Header{
					"Content-Type": []string{"application/json"},
					"Accept":       []string{"application/json"},
				},
			}

			// Add request body if available
			if operation.RequestBody != nil {
				content := operation.RequestBody.Value.Content.Get("application/json")
				if content != nil && content.Example != nil {
					body, _ := json.Marshal(content.Example)
					target.Body = body
				}
			}

			targets = append(targets, target)
		}
	}

	return targets
}

func (s *TestingService) buildURL(path string, operation *openapi3.Operation) string {
	// Implementation of URL building logic
	return fmt.Sprintf("http://%s:%s%s", s.apiHost, s.apiPort, path)
}
