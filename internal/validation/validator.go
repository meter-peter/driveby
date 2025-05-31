package validation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/openapi"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// PerformanceTargetConfig holds configuration for performance test targets
type PerformanceTargetConfig struct {
	MaxLatencyP95  time.Duration
	MinSuccessRate float64
}

// ValidatorConfig holds configuration for the validator
type ValidatorConfig struct {
	BaseURL           string
	SpecPath          string
	LogPath           string
	Environment       string
	Version           string
	AutoFix           bool
	Timeout           time.Duration
	PerformanceTarget PerformanceTargetConfig
}

// APIValidator implements the validation logic
type APIValidator struct {
	config    ValidatorConfig
	logger    *Logger
	loader    *openapi.Loader
	client    *http.Client
	validator *Validator
}

// NewAPIValidator creates a new validator instance
func NewAPIValidator(config ValidatorConfig) (*APIValidator, error) {
	logger, err := NewLogger(config.LogPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	validationValidator := NewValidator()
	validationValidator.SetBaseURL(config.BaseURL)

	return &APIValidator{
		config: config,
		logger: logger,
		loader: openapi.NewLoader(),
		client: &http.Client{
			Timeout: config.Timeout,
		},
		validator: validationValidator,
	}, nil
}

// Validate runs the complete validation suite
func (v *APIValidator) Validate(ctx context.Context) (*ValidationReport, error) {
	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// Load and validate OpenAPI spec
	if err := v.loader.LoadFromFile(v.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := v.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	// Validate against each principle
	for _, principle := range CorePrinciples {
		result := v.validatePrinciple(ctx, principle, doc)
		report.Principles = append(report.Principles, result)

		if result.Passed {
			report.PassedChecks++
		} else {
			report.FailedChecks++

			// Attempt auto-fix if enabled and principle is auto-fixable
			if v.config.AutoFix && principle.AutoFixable {
				fixResult := v.attemptAutoFix(ctx, principle, result, doc)
				report.AutoFixes = append(report.AutoFixes, fixResult)

				if fixResult.Success {
					// Revalidate after fix
					result = v.validatePrinciple(ctx, principle, doc)
					if result.Passed {
						report.PassedChecks++
						report.FailedChecks--
					}
				}
			}
		}
	}

	report.TotalChecks = len(CorePrinciples)
	v.updateSummary(report)

	// Log the report
	if err := v.logger.LogReport(report); err != nil {
		return nil, fmt.Errorf("failed to log validation report: %w", err)
	}

	return report, nil
}

// validatePrinciple checks a single principle
func (v *APIValidator) validatePrinciple(ctx context.Context, principle Principle, doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: principle,
		Passed:    true,
	}

	switch principle.ID {
	case "P001": // OpenAPI Specification Compliance
		result = ValidateOpenAPICompliance(doc)
	case "P002": // Response Time Performance (Documentation check)
		if v.config.PerformanceTarget.MaxLatencyP95 == 0 || v.config.PerformanceTarget.MinSuccessRate == 0 {
			result.Passed = false
			result.Message = "Performance targets are not configured."
			result.SuggestedFix = "Configure performance_target.max_latency_p95 and performance_target.min_success_rate in the config."
		}
	case "P003": // Error Response Documentation
		result = ValidateErrorDocumentation(doc)
	case "P004": // Request Validation (Documentation check)
		result = ValidateRequestValidation(doc)
	case "P005": // Authentication Requirements (Documentation check)
		result = ValidateAuthentication(doc)
	case "P006": // Endpoint Functional Testing
		endpointResult, err := v.validator.ValidateEndpoints(ctx, doc, v.config.BaseURL)
		if err != nil {
			result.Passed = false
			result.Message = fmt.Sprintf("Endpoint functional testing failed: %v", err)
			result.Details = map[string]interface{}{"error": err.Error()}
		} else {
			result.Details = endpointResult.Endpoints

			allSuccess := true
			var failedEndpoints []string
			for _, epVal := range endpointResult.Endpoints {
				if epVal.Status != "success" {
					allSuccess = false
					failedEndpoints = append(failedEndpoints, fmt.Sprintf("%s %s (Status: %s, Code: %d)", epVal.Method, epVal.Path, epVal.Status, epVal.StatusCode))
				}
			}

			if allSuccess {
				result.Passed = true
				result.Message = "All documented endpoints are reachable and return documented status codes."
			} else {
				result.Passed = false
				result.Message = fmt.Sprintf("Some endpoints failed functional tests. Failed: %d/%d", len(failedEndpoints), len(endpointResult.Endpoints))
				result.Details = map[string]interface{}{"failed_endpoints": failedEndpoints, "all_results": endpointResult.Endpoints}
			}
		}
	case "P007": // API Performance Compliance (Execution and Validation)
		var targets []vegeta.Target
		for path, pathItem := range doc.Paths.Map() {
			for method := range pathItem.Operations() {
				targets = append(targets, vegeta.Target{
					Method: method,
					URL:    fmt.Sprintf("%s%s", v.config.BaseURL, path),
				})
			}
		}

		if len(targets) == 0 {
			result.Passed = true
			result.Message = "No targets found in OpenAPI spec for performance testing."
			break
		}

		rate := 50.0
		duration := 10 * time.Second

		perfResult, err := v.validator.RunPerformanceTests(targets, rate, duration)
		if err != nil {
			result.Passed = false
			result.Message = fmt.Sprintf("Performance testing failed: %v", err)
			result.Details = map[string]interface{}{"error": err.Error()}
		} else {
			result.Details = perfResult.Performance

			metTargets := true
			var failureReasons []string

			if perfResult.Performance.LatencyP95 > v.config.PerformanceTarget.MaxLatencyP95 && v.config.PerformanceTarget.MaxLatencyP95 != 0 {
				metTargets = false
				failureReasons = append(failureReasons, fmt.Sprintf("P95 latency (%s) exceeded target (%s)", perfResult.Performance.LatencyP95, v.config.PerformanceTarget.MaxLatencyP95))
			}
			if perfResult.Performance.ErrorRate > (1.0-v.config.PerformanceTarget.MinSuccessRate/100.0) && v.config.PerformanceTarget.MinSuccessRate != 0 {
				metTargets = false
				failureReasons = append(failureReasons, fmt.Sprintf("Error rate (%.2f%%) exceeded allowed (%.2f%% success rate target)", perfResult.Performance.ErrorRate*100, v.config.PerformanceTarget.MinSuccessRate))
			}

			if metTargets {
				result.Passed = true
				result.Message = "API performance meets configured targets."
			} else {
				result.Passed = false
				result.Message = fmt.Sprintf("API performance failed to meet targets: %s", strings.Join(failureReasons, ", "))
			}
		}
	default:
		result.Passed = false
		result.Message = fmt.Sprintf("Unknown principle ID: %s", principle.ID)
	}

	return result
}

// ValidateOpenAPICompliance checks OpenAPI specification compliance
func ValidateOpenAPICompliance(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[0], // P001
		Passed:    true,
	}

	// Validate spec structure
	if err := doc.Validate(context.Background()); err != nil {
		result.Passed = false
		result.Message = fmt.Sprintf("OpenAPI spec validation failed: %v", err)
		return result
	}

	// Check for required fields
	if doc.Info == nil || doc.Info.Title == "" || doc.Info.Version == "" {
		result.Passed = false
		result.Message = "Missing required OpenAPI info fields (title, version)"
		return result
	}

	// Check paths
	if doc.Paths == nil || len(doc.Paths.Map()) == 0 {
		result.Passed = false
		result.Message = "No paths defined in OpenAPI spec"
		return result
	}

	return result
}

// ValidateErrorDocumentation checks error response documentation
func ValidateErrorDocumentation(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[2], // P003
		Passed:    true,
	}

	var missingErrors []string
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if operation.Responses == nil {
				missingErrors = append(missingErrors, fmt.Sprintf("%s %s", method, path))
				continue
			}

			hasErrorResponse := false
			for code := range operation.Responses.Map() {
				if code >= "400" && code < "600" {
					hasErrorResponse = true
					break
				}
			}

			if !hasErrorResponse {
				missingErrors = append(missingErrors, fmt.Sprintf("%s %s", method, path))
			}
		}
	}

	if len(missingErrors) > 0 {
		result.Passed = false
		result.Message = "Endpoints missing error response documentation"
		result.Details = missingErrors
		result.SuggestedFix = "Add error responses (4xx, 5xx) to all endpoints"
	}

	return result
}

// ValidateRequestValidation checks request parameter validation
func ValidateRequestValidation(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[3], // P004
		Passed:    true,
	}

	var missingValidation []string
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			// Check path parameters
			for _, param := range operation.Parameters {
				if param.Value == nil || param.Value.Schema == nil {
					missingValidation = append(missingValidation,
						fmt.Sprintf("%s %s: parameter %s", method, path, param.Value.Name))
				}
			}

			// Check request body
			if operation.RequestBody != nil && operation.RequestBody.Value != nil {
				if operation.RequestBody.Value.Content == nil {
					missingValidation = append(missingValidation,
						fmt.Sprintf("%s %s: request body", method, path))
				}
			}
		}
	}

	if len(missingValidation) > 0 {
		result.Passed = false
		result.Message = "Endpoints missing request validation"
		result.Details = missingValidation
		result.SuggestedFix = "Add schema validation for all parameters and request bodies"
	}

	return result
}

// ValidateAuthentication checks authentication requirements
func ValidateAuthentication(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[4], // P005
		Passed:    true,
	}

	if doc.Components == nil || doc.Components.SecuritySchemes == nil {
		result.Passed = false
		result.Message = "No security schemes defined"
		result.SuggestedFix = "Define security schemes in components.securitySchemes"
		return result
	}

	var missingAuth []string
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if operation.Security == nil && doc.Security == nil {
				missingAuth = append(missingAuth, fmt.Sprintf("%s %s", method, path))
			}
		}
	}

	if len(missingAuth) > 0 {
		result.Passed = false
		result.Message = "Endpoints missing authentication requirements"
		result.Details = missingAuth
		result.SuggestedFix = "Add security requirements to endpoints or global security"
	}

	return result
}

// attemptAutoFix tries to automatically fix a validation issue
func (v *APIValidator) attemptAutoFix(ctx context.Context, principle Principle, result PrincipleResult, doc *openapi3.T) AutoFixResult {
	fixResult := AutoFixResult{
		PrincipleID: principle.ID,
		Timestamp:   time.Now(),
		Location:    v.config.SpecPath,
	}

	switch principle.ID {
	case "P001": // OpenAPI Specification Compliance
		if strings.Contains(result.Message, "Missing required OpenAPI info fields") {
			if doc.Info == nil {
				doc.Info = &openapi3.Info{}
			}
			if doc.Info.Title == "" {
				doc.Info.Title = "API"
			}
			if doc.Info.Version == "" {
				doc.Info.Version = "1.0.0"
			}
			fixResult.Success = true
			fixResult.Message = "Added missing OpenAPI info fields"
		}

	case "P003": // Error Response Documentation
		if result.Details != nil {
			if missingErrors, ok := result.Details.([]string); ok {
				for _, endpoint := range missingErrors {
					parts := strings.Split(endpoint, " ")
					if len(parts) != 2 {
						continue
					}
					method, path := parts[0], parts[1]

					pathItem := doc.Paths.Find(path)
					if pathItem != nil {
						if operation, ok := pathItem.Operations()[method]; ok {
							if operation.Responses == nil {
								operation.Responses = openapi3.NewResponses()
							}

							// Add common error responses
							badRequest := "Bad Request"
							unauthorized := "Unauthorized"
							internalError := "Internal Server Error"

							operation.Responses.Set("400", &openapi3.ResponseRef{
								Value: &openapi3.Response{
									Description: &badRequest,
								},
							})
							operation.Responses.Set("401", &openapi3.ResponseRef{
								Value: &openapi3.Response{
									Description: &unauthorized,
								},
							})
							operation.Responses.Set("500", &openapi3.ResponseRef{
								Value: &openapi3.Response{
									Description: &internalError,
								},
							})
						}
					}
				}
				fixResult.Success = true
				fixResult.Message = "Added common error responses to endpoints"
			}
		}

	case "P004": // Request Validation
		if result.Details != nil {
			if missingValidation, ok := result.Details.([]string); ok {
				for _, endpoint := range missingValidation {
					parts := strings.Split(endpoint, ": ")
					if len(parts) != 2 {
						continue
					}

					endpointInfo := strings.Split(parts[0], " ")
					if len(endpointInfo) != 2 {
						continue
					}

					method, path := endpointInfo[0], endpointInfo[1]
					paramName := parts[1]

					pathItem := doc.Paths.Find(path)
					if pathItem != nil {
						if operation, ok := pathItem.Operations()[method]; ok {
							if paramName == "request body" {
								if operation.RequestBody == nil {
									operation.RequestBody = &openapi3.RequestBodyRef{
										Value: &openapi3.RequestBody{
											Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
												Type: "object",
											}),
										},
									}
								}
							} else {
								// Add schema for parameter
								for i, param := range operation.Parameters {
									if param.Value != nil && param.Value.Name == paramName {
										param.Value.Schema = &openapi3.SchemaRef{
											Value: &openapi3.Schema{
												Type: "string",
											},
										}
										operation.Parameters[i] = param
										break
									}
								}
							}
						}
					}
				}
				fixResult.Success = true
				fixResult.Message = "Added basic validation schemas to parameters and request bodies"
			}
		}
	}

	if fixResult.Success {
		// Save the modified spec
		specJSON, err := json.MarshalIndent(doc, "", "  ")
		if err != nil {
			fixResult.Success = false
			fixResult.Error = fmt.Sprintf("Failed to marshal modified spec: %v", err)
			return fixResult
		}

		if err := ioutil.WriteFile(v.config.SpecPath, specJSON, 0644); err != nil {
			fixResult.Success = false
			fixResult.Error = fmt.Sprintf("Failed to save modified spec: %v", err)
			return fixResult
		}

		fixResult.Original = result
		fixResult.Fixed = "Spec file updated with fixes"
	}

	return fixResult
}

// updateSummary updates the validation summary
func (v *APIValidator) updateSummary(report *ValidationReport) {
	summary := ValidationSummary{}
	categories := make(map[string]bool)
	failedTags := make(map[string]bool)

	for _, result := range report.Principles {
		// Count issues by severity
		if !result.Passed {
			switch result.Principle.Severity {
			case "critical":
				summary.CriticalIssues++
			case "warning":
				summary.Warnings++
			case "info":
				summary.Info++
			}

			// Track categories and tags
			categories[result.Principle.Category] = true
			for _, tag := range result.Principle.Tags {
				failedTags[tag] = true
			}
		}
	}

	// Convert maps to slices
	for category := range categories {
		summary.Categories = append(summary.Categories, category)
	}
	for tag := range failedTags {
		summary.FailedTags = append(summary.FailedTags, tag)
	}

	report.Summary = summary
}
