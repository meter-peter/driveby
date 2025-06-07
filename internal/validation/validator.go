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
)

// PerformanceTargetConfig holds configuration for performance test targets
type PerformanceTargetConfig struct {
	MaxLatencyP95  time.Duration
	MinSuccessRate float64
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	Token       string
	TokenType   string
	TokenHeader string
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
	Auth              AuthConfig
}

// APIValidator implements the validation logic
type APIValidator struct {
	config    ValidatorConfig
	logger    *Logger
	loader    *openapi.Loader
	client    *http.Client
	validator *Validator
	baseURL   string
}

// NewAPIValidator creates a new validator instance
func NewAPIValidator(config ValidatorConfig) (*APIValidator, error) {
	logger, err := NewLogger(config.LogPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	validationValidator := NewValidator()
	validationValidator.SetBaseURL(config.BaseURL)

	// Set up HTTP client with authentication
	client := &http.Client{
		Timeout: config.Timeout,
		Transport: &authTransport{
			base: http.DefaultTransport,
			auth: config.Auth,
		},
	}

	// Determine the baseURL based on specPath. If specPath is a URL ending in a file extension (.json/.yaml), extract the directory path. Otherwise, use config.BaseURL.
	baseURL := config.BaseURL // Default to provided BaseURL
	log.Debugf("Initial specPath: %s, initial baseURL: %s", config.SpecPath, baseURL)
	if strings.HasPrefix(config.SpecPath, "http://") || strings.HasPrefix(config.SpecPath, "https://") {
		log.Debugf("specPath is a URL: %s", config.SpecPath)
		// Check if the URL ends with a file extension indicating the spec file
		if strings.HasSuffix(strings.ToLower(config.SpecPath), ".json") || strings.HasSuffix(strings.ToLower(config.SpecPath), ".yaml") || strings.HasSuffix(strings.ToLower(config.SpecPath), ".yml") {
			log.Debug("specPath ends with .json/.yaml, extracting directory")
			// Extract the base URL by removing the filename part
			lastSlash := strings.LastIndex(config.SpecPath, "/")
			if lastSlash != -1 {
				baseURL = config.SpecPath[:lastSlash]
				log.Debugf("Extracted baseURL from specPath: %s", baseURL)
			} else {
				// Should not happen for a valid URL, but handle defensively
				baseURL = config.SpecPath
				log.Debugf("Could not find slash in specPath, using full specPath as baseURL: %s", baseURL)
			}
		} else {
			log.Debug("specPath is a URL but does not end with .json/.yaml, using provided config.BaseURL")
			// Use the provided config.BaseURL if the specPath is a base URL itself
			baseURL = config.BaseURL
		}
	} else {
		log.Debug("specPath is not a URL, using provided config.BaseURL")
		// If specPath is a local file path, use the provided config.BaseURL for API calls
		baseURL = config.BaseURL
	}

	// The internal validator instance uses the calculated baseURL for making requests
	validationValidator.SetBaseURL(baseURL)

	log.Debugf("Final baseURL determined: %s", baseURL)

	return &APIValidator{
		config:    config,
		logger:    logger,
		loader:    openapi.NewLoader(),
		client:    client,
		validator: validationValidator,
		baseURL:   baseURL,
	}, nil
}

// authTransport implements http.RoundTripper to add authentication headers
type authTransport struct {
	base http.RoundTripper
	auth AuthConfig
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.auth.Token != "" {
		headerName := t.auth.TokenHeader
		if headerName == "" {
			headerName = "Authorization"
		}

		tokenType := t.auth.TokenType
		if tokenType == "" {
			tokenType = "Bearer"
		}

		req.Header.Set(headerName, fmt.Sprintf("%s %s", tokenType, t.auth.Token))
	}

	return t.base.RoundTrip(req)
}

// Validate runs the complete validation suite
func (v *APIValidator) Validate(ctx context.Context) (*ValidationReport, error) {
	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// Load and validate OpenAPI spec
	if err := v.loader.LoadFromFileOrURL(v.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := v.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	// Process principles sequentially to handle dependencies
	var principlesResults []PrincipleResult
	var p006Result PrincipleResult // Store P006 result for P007 dependency

	for _, principle := range CorePrinciples {

		// Handle P007 dependency on P006 before validating P007
		// If P006 indicates widespread functional failures (e.g., auth or network errors), skip P007

		if principle.ID == "P007" {
			skipPerformance := false
			// Check P006 result if available and failed
			if p006Result.Principle.ID != "" && !p006Result.Passed {
				if detailsMap, ok := p006Result.Details.(map[string]interface{}); ok {
					totalEndpoints := 0
					if totalVal, totalOk := detailsMap["total_endpoints"].(int); totalOk {
						totalEndpoints = totalVal
					}

					if totalEndpoints > 0 {
						authFailed, authOk := detailsMap["auth_failed_count"].(int)
						otherFailed, otherOk := detailsMap["other_failed_count"].(int)

						if (authOk && authFailed > totalEndpoints/2) || (otherOk && otherFailed > totalEndpoints/2) {
							skipPerformance = true

							// Create a skipped result for P007
							result := PrincipleResult{
								Principle:   principle,
								Passed:      false, // Mark as failed or skipped in report
								Message:     "Performance test skipped due to widespread functional test failures (likely auth/network issues).",
								Explanation: "Functional tests (P006) indicated widespread failures, making performance metrics unreliable. Address functional issues first.",
								Details:     nil, // No performance details if skipped
							}
							principlesResults = append(principlesResults, result)
							log.Debug("P007 skipped due to P006 widespread failure")
						}
					}
				}
			}

			// If not skipped, validate P007 as usual
			if !skipPerformance {
				result := v.validatePrinciple(ctx, principle, doc)
				principlesResults = append(principlesResults, result)

				// Update passed/failed counts based on P007 result
				if result.Passed {
					report.PassedChecks++
				} else {
					report.FailedChecks++
				}
			}

		} else { // For all other principles (including P006)
			result := v.validatePrinciple(ctx, principle, doc)
			principlesResults = append(principlesResults, result)

			// Store P006 result for P007 dependency check
			if principle.ID == "P006" {
				p006Result = result
			}

			// Update passed/failed counts for non-P007 principles
			if result.Passed {
				report.PassedChecks++
			} else {
				report.FailedChecks++

				// Attempt auto-fix if enabled and principle is auto-fixable (only for non-P007)
				if v.config.AutoFix && principle.AutoFixable {
					fixResult := v.attemptAutoFix(ctx, principle, result, doc)
					report.AutoFixes = append(report.AutoFixes, fixResult)

					if fixResult.Success {
						// Revalidate after fix
						revalidatedResult := v.validatePrinciple(ctx, principle, doc)
						// Replace the original result with the revalidated one
						for i, r := range principlesResults {
							if r.Principle.ID == principle.ID {
								// Adjust counts before replacing
								if principlesResults[i].Passed && !revalidatedResult.Passed {
									report.PassedChecks--
									report.FailedChecks++
								} else if !principlesResults[i].Passed && revalidatedResult.Passed {
									report.PassedChecks++
									report.FailedChecks--
								}
								principlesResults[i] = revalidatedResult
								log.Debugf("Revalidated principle %s after autofix. Passed: %v", principle.ID, revalidatedResult.Passed)
								break
							}
						}
					}
				}
			}
		}
	}

	report.Principles = principlesResults // Assign the collected results
	report.TotalChecks = len(CorePrinciples)

	// Ensure total checks is correct even if principles were skipped/modified
	report.TotalChecks = len(report.Principles) + len(report.AutoFixes) // Adjust if auto-fixes add principles or counts? Revisit this.

	// Recalculate passed/failed checks based on final principles list
	report.PassedChecks = 0
	report.FailedChecks = 0
	for _, res := range report.Principles {
		if res.Passed {
			report.PassedChecks++
		} else {
			report.FailedChecks++
		}
	}

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
		log.Debug("Running P006: Endpoint Functional Testing")
		endpointResult, err := v.validator.ValidateEndpoints(ctx, doc, v.baseURL)
		if err != nil {
			result.Passed = false
			result.Message = fmt.Sprintf("Endpoint functional testing failed to run: %v", err)
			result.Details = map[string]interface{}{"error": err.Error()}
			log.WithError(err).Error("P006 failed to run")
		} else {
			result.Details = endpointResult.Endpoints

			// Count different failure statuses
			total := len(endpointResult.Endpoints)
			passedCount := 0
			authFailedCount := 0
			clientErrorCount := 0
			serverErrorCount := 0
			failedCount := 0 // Network errors, timeouts, etc.
			undocumentedCount := 0

			for _, epVal := range endpointResult.Endpoints {
				switch epVal.Status {
				case "success":
					passedCount++
				case "auth_failed":
					authFailedCount++
				case "client_error":
					clientErrorCount++
				case "server_error":
					serverErrorCount++
				case "undocumented":
					undocumentedCount++
				case "failed": // Catch all for other failures like network errors
					failedCount++
				}
			}

			// Determine overall status and message
			if passedCount == total && total > 0 {
				result.Passed = true
				result.Message = fmt.Sprintf("All %d documented endpoints are reachable and returned documented status codes.", total)
				log.Debug("P006 passed: All endpoints successful")
			} else if total == 0 {
				result.Passed = true // Or handle as 'info' or 'skipped'
				result.Message = "No endpoints found in OpenAPI spec for functional testing."
				log.Debug("P006 passed: No endpoints to test")
			} else {
				result.Passed = false
				messageParts := []string{}
				if authFailedCount > 0 {
					messageParts = append(messageParts, fmt.Sprintf("%d authentication failures (401/403)", authFailedCount))
				}
				if clientErrorCount > 0 {
					messageParts = append(messageParts, fmt.Sprintf("%d client errors (4xx)", clientErrorCount))
				}
				if serverErrorCount > 0 {
					messageParts = append(messageParts, fmt.Sprintf("%d server errors (5xx)", serverErrorCount))
				}
				if undocumentedCount > 0 {
					messageParts = append(messageParts, fmt.Sprintf("%d undocumented status codes", undocumentedCount))
				}
				if failedCount > 0 {
					messageParts = append(messageParts, fmt.Sprintf("%d network/other failures", failedCount))
				}

				result.Message = fmt.Sprintf("Some endpoints failed functional tests (%d/%d failed): %s", total-passedCount, total, strings.Join(messageParts, ", "))
				result.Details = map[string]interface{}{
					"total_endpoints":    total,
					"passed_count":       passedCount,
					"auth_failed_count":  authFailedCount,
					"client_error_count": clientErrorCount,
					"server_error_count": serverErrorCount,
					"undocumented_count": undocumentedCount,
					"other_failed_count": failedCount,
					"all_results":        endpointResult.Endpoints,
				}
				log.Debugf("P006 failed: %s", result.Message)
			}
		}
	case "P008": // API Versioning
		log.Debug("Running P008: API Versioning")
		if doc.Info == nil || doc.Info.Version == "" {
			result.Passed = false
			result.Message = "API version is not specified in the OpenAPI document info section."
			result.SuggestedFix = "Add or update the 'version' field in the 'info' section of your OpenAPI specification."
			log.Debug("P008 failed: Version not specified")
		} else {
			result.Message = fmt.Sprintf("API version is specified as: %s", doc.Info.Version)
			log.Debugf("P008 passed: Version specified as %s", doc.Info.Version)
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

// ValidationReport represents a detailed report of validation results
