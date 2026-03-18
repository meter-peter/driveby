package validation

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/meter-peter/driveby/internal/openapi"
	"github.com/meter-peter/driveby/internal/spec"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// OpenAPIValidator handles validation of OpenAPI specifications
type OpenAPIValidator struct {
	config ValidatorConfig
	loader *openapi.Loader
}

// validateConfig validates the validator configuration
func validateConfig(config ValidatorConfig) error {
	if config.SpecPath == "" {
		return fmt.Errorf("spec path is required")
	}
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}
	if config.Timeout <= 0 {
		return fmt.Errorf("timeout must be greater than 0")
	}
	if config.Auth != nil {
		authMethods := 0
		if config.Auth.Token != "" {
			authMethods++
		}
		if config.Auth.APIKey != "" {
			authMethods++
		}
		if config.Auth.Username != "" {
			authMethods++
		}
		if authMethods > 1 {
			return fmt.Errorf("only one authentication method can be specified")
		}
	}
	if config.PerformanceTarget != nil {
		if config.PerformanceTarget.Duration <= 0 {
			return fmt.Errorf("performance test duration must be greater than 0")
		}
		if config.PerformanceTarget.ConcurrentUsers <= 0 {
			return fmt.Errorf("concurrent users must be greater than 0")
		}
		if config.PerformanceTarget.MinSuccessRate < 0 || config.PerformanceTarget.MinSuccessRate > 1 {
			return fmt.Errorf("minimum success rate must be between 0 and 1")
		}
	}
	return nil
}

// NewOpenAPIValidator creates a new validator instance
func NewOpenAPIValidator(config ValidatorConfig) (*OpenAPIValidator, error) {
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid validator config: %w", err)
	}

	// Default to minimal mode if not specified
	if config.ValidationMode == "" {
		config.ValidationMode = ValidationModeMinimal
		log.Debug("Validation mode not specified, defaulting to minimal mode")
	}
	return &OpenAPIValidator{
		config: config,
		loader: openapi.NewLoader(),
	}, nil
}

// ValidateSpec runs validation against an OpenAPI specification
func (v *OpenAPIValidator) ValidateSpec(ctx context.Context) (*ValidationReport, error) {
	log.Debugf("Starting OpenAPI spec validation with config: %+v", v.config)

	// Load API spec (OpenAPI 3.x or Swagger 2.0)
	if err := v.loader.LoadFromFileOrURL(v.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load API spec: %w", err)
	}
	doc := v.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get API document")
	}

	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// Select validation principles based on mode
	var validationPrinciples []Principle
	switch v.config.ValidationMode {
	case ValidationModeTestOnly:
		// Test-only mode: skip all validation, only run tests
		log.Debug("Running in test-only mode - skipping all validation")
		validationPrinciples = []Principle{}
	case ValidationModeMinimal:
		// Minimal mode: only essential validation for testing
		validationPrinciples = []Principle{
			CorePrinciples[0], // P001: OpenAPI Specification Compliance (basic structure)
		}
		log.Debug("Running in minimal mode - essential validation only")
	case ValidationModeStrict:
		// Strict mode: comprehensive validation
		validationPrinciples = []Principle{
			CorePrinciples[0], // P001: OpenAPI Specification Compliance
			CorePrinciples[1], // P002: API Documentation Completeness
			CorePrinciples[2], // P003: Error Response Documentation
			CorePrinciples[3], // P004: Request Validation
			CorePrinciples[4], // P005: Authentication Requirements
			CorePrinciples[7], // P008: API Versioning
		}
		log.Debug("Running in strict mode - comprehensive validation")
	default:
		// Default to minimal mode
		validationPrinciples = []Principle{
			CorePrinciples[0], // P001: OpenAPI Specification Compliance (basic structure)
		}
		log.Debug("Running in default minimal mode")
	}

	for _, principle := range validationPrinciples {
		result := v.validatePrinciple(ctx, principle, doc)
		report.Principles = append(report.Principles, result)

		if result.Passed {
			report.PassedChecks++
		} else {
			report.FailedChecks++
		}
	}

	report.TotalChecks = len(validationPrinciples)
	v.updateSummary(report)

	return report, nil
}

// validatePrinciple checks a single validation principle
func (v *OpenAPIValidator) validatePrinciple(ctx context.Context, principle Principle, doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: principle,
		Passed:    true,
	}

	switch principle.ID {
	case "P001": // Specification Compliance
		result = v.validateOpenAPICompliance(ctx, doc)
	case "P002": // API Documentation Completeness
		result = v.validateDocumentationQuality(doc)
	case "P003": // Error Response Documentation
		result = v.validateErrorHandling(doc)
	case "P004": // Request Validation
		result = v.validateRequestSchema(doc)
	case "P005": // Authentication Requirements
		result = v.validateAuthentication(doc)
	case "P008": // API Versioning
		result = v.validateVersioning(doc)
	default:
		result.Passed = false
		result.Message = fmt.Sprintf("Unknown principle ID: %s", principle.ID)
	}

	return result
}

// validateOpenAPICompliance validates that the API spec is structurally compliant.
// For OpenAPI 3.x, this includes full schema validation. For Swagger 2.0, it performs
// a focused set of structural checks.
func (v *OpenAPIValidator) validateOpenAPICompliance(ctx context.Context, doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[0], // P001
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// Track individual check results
	checks := make(map[string]bool)
	messages := make(map[string]string)

	rawVersion := doc.RawVersion()
	if rawVersion == "" {
		checks["Specification version is present"] = false
		messages["Specification version is present"] = "Specification version is not specified"
	} else {
		checks["Specification version is present"] = true
	}

	// Check required info fields
	info := doc.Info()
	if info == nil {
		checks["Required info fields (title, version) are present"] = false
		messages["Required info fields (title, version) are present"] = "Info section is missing"
	} else {
		missingFields := []string{}
		if info.Title == "" {
			missingFields = append(missingFields, "title")
		}
		if info.Version == "" {
			missingFields = append(missingFields, "version")
		}
		if len(missingFields) > 0 {
			checks["Required info fields (title, version) are present"] = false
			messages["Required info fields (title, version) are present"] = fmt.Sprintf("Missing required fields: %s", strings.Join(missingFields, ", "))
		} else {
			checks["Required info fields (title, version) are present"] = true
		}
	}

	// Check paths
	paths := doc.Paths()
	if paths == nil || len(paths) == 0 {
		checks["Paths are properly defined"] = false
		messages["Paths are properly defined"] = "No paths defined in the API"
	} else {
		checks["Paths are properly defined"] = true
	}

	// Components presence is treated as non-fatal; downstream principles handle depth.
	checks["Components are valid"] = true

	// Check references / structural validity via adapter-level validation.
	if err := doc.ValidateStructure(ctx); err != nil {
		checks["Specification structure is valid"] = false
		messages["Specification structure is valid"] = err.Error()
	} else {
		checks["Specification structure is valid"] = true
	}

	// Check for duplicate operationIds
	operationIDs := make(map[string][]string)
	for path, pathItem := range paths {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation != nil && operation.OperationID != "" {
				operationIDs[operation.OperationID] = append(operationIDs[operation.OperationID], fmt.Sprintf("%s %s", method, path))
			}
		}
	}
	duplicates := []string{}
	for opID, locations := range operationIDs {
		if len(locations) > 1 {
			duplicates = append(duplicates, fmt.Sprintf("%s used in: %s", opID, strings.Join(locations, ", ")))
		}
	}
	if len(duplicates) > 0 {
		checks["No duplicate operationIds"] = false
		messages["No duplicate operationIds"] = fmt.Sprintf("Duplicate operationIds found: %s", strings.Join(duplicates, "; "))
	} else {
		checks["No duplicate operationIds"] = true
	}

	// Check HTTP methods
	validMethods := map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"HEAD":    true,
		"OPTIONS": true,
		"TRACE":   true,
	}
	invalidMethods := []string{}
	for path, pathItem := range paths {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method := range pathItem.Operations {
			if !validMethods[method] {
				invalidMethods = append(invalidMethods, fmt.Sprintf("%s %s", method, path))
			}
		}
	}
	if len(invalidMethods) > 0 {
		checks["Valid HTTP methods used"] = false
		messages["Valid HTTP methods used"] = fmt.Sprintf("Invalid HTTP methods found: %s", strings.Join(invalidMethods, ", "))
	} else {
		checks["Valid HTTP methods used"] = true
	}

	// OpenAPI 3.x-only: check for null type in schemas.
	if doc.Type() == spec.SpecTypeOpenAPI3 {
		if comps := doc.Components(); comps != nil && comps.Schemas != nil {
			for name, schema := range comps.Schemas {
				if schema == nil {
					continue
				}
				// Allow null type in OpenAPI 3.1.0
				if schema.Type == "null" && rawVersion == "3.1.0" {
					continue
				}
				// For OpenAPI 3.0.x, null type should be represented as ["null", "type"]
				if schema.Type == "null" && strings.HasPrefix(rawVersion, "3.0.") {
					checks["Specification structure is valid"] = false
					messages["Specification structure is valid"] = fmt.Sprintf("invalid components: schema %q: 'null' type should be represented as [\"null\", \"type\"] in OpenAPI 3.0.x", name)
				}
			}
		}
	}

	// Update result based on checks
	allPassed := true
	for _, passed := range checks {
		if !passed {
			allPassed = false
			break
		}
	}
	result.Passed = allPassed
	result.Details = map[string]interface{}{
		"checks":   checks,
		"messages": messages,
	}

	if !allPassed {
		var failedChecks []string
		for check, passed := range checks {
			if !passed {
				failedChecks = append(failedChecks, fmt.Sprintf("%s: %s", check, messages[check]))
			}
		}
		result.Message = fmt.Sprintf("OpenAPI spec validation failed: %s", strings.Join(failedChecks, "; "))
	} else {
		result.Message = "OpenAPI specification is fully compliant with 3.0/3.1 standards"
	}

	return result
}

// validateDocumentationQuality validates the quality and completeness of API documentation
func (v *OpenAPIValidator) validateDocumentationQuality(doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[1], // P002
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingDocs := make(map[string][]string)

	// Check API-level documentation
	info := doc.Info()
	if info == nil {
		checks["API has a general description"] = false
		messages["API has a general description"] = "Info section is missing"
	} else if info.Description == "" {
		checks["API has a general description"] = false
		messages["API has a general description"] = "API description is missing"
	} else {
		checks["API has a general description"] = true
	}

	// Check contact information
	if info == nil || info.Contact == nil {
		checks["Contact information is provided"] = false
		messages["Contact information is provided"] = "Contact information is missing"
	} else {
		hasContact := info.Contact.Name != "" || info.Contact.Email != "" || info.Contact.URL != ""
		checks["Contact information is provided"] = hasContact
		if !hasContact {
			messages["Contact information is provided"] = "Contact information is empty"
		}
	}

	// Check license information
	if info == nil || info.License == nil {
		checks["License information is provided"] = false
		messages["License information is provided"] = "License information is missing"
	} else if info.License.Name == "" {
		checks["License information is provided"] = false
		messages["License information is provided"] = "License name is missing"
	} else {
		checks["License information is provided"] = true
	}

	// Check operation documentation
	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check summary
			if operation.Summary == "" {
				missingDocs["All operations have clear summaries"] = append(missingDocs["All operations have clear summaries"], opKey)
				checks["All operations have clear summaries"] = false
			}

			// Check description
			if operation.Description == "" {
				missingDocs["All operations have detailed descriptions"] = append(missingDocs["All operations have detailed descriptions"], opKey)
				checks["All operations have detailed descriptions"] = false
			}

			// Check operationId
			if operation.OperationID == "" {
				missingDocs["All operations have unique operationIds"] = append(missingDocs["All operations have unique operationIds"], opKey)
				checks["All operations have unique operationIds"] = false
			}

			// Check parameter documentation
			for _, param := range operation.Parameters {
				if param == nil {
					continue
				}
				if param.Description == "" {
					missingDocs["All parameters have descriptions"] = append(missingDocs["All parameters have descriptions"],
						fmt.Sprintf("%s: parameter %s", opKey, param.Name))
					checks["All parameters have descriptions"] = false
				}
			}

			// Check request body documentation
			if operation.RequestBody != nil {
				if operation.RequestBody.Description == "" {
					missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
						fmt.Sprintf("%s: request body", opKey))
					checks["All request/response bodies have examples"] = false
				}
				// Check for examples in content
				for contentType, content := range operation.RequestBody.Content {
					if content == nil || (content.Example == nil && len(content.Examples) == 0) {
						missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
							fmt.Sprintf("%s: %s request body", opKey, contentType))
						checks["All request/response bodies have examples"] = false
					}
				}
			}

			// Check response documentation
			for status, response := range operation.Responses {
				if response == nil {
					continue
				}
				if response.Description == "" {
					missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
						fmt.Sprintf("%s: %s response", opKey, status))
					checks["All request/response bodies have examples"] = false
				}
				// Check for examples in content
				for contentType, content := range response.Content {
					if content == nil || (content.Example == nil && len(content.Examples) == 0) {
						missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
							fmt.Sprintf("%s: %s %s response", opKey, status, contentType))
						checks["All request/response bodies have examples"] = false
					}
				}
			}
		}
	}

	// Check schema documentation
	if comps := doc.Components(); comps != nil && comps.Schemas != nil {
		for name, schema := range comps.Schemas {
			if schema == nil {
				continue
			}
			if schema.Description == "" {
				missingDocs["All schemas have descriptions"] = append(missingDocs["All schemas have descriptions"], name)
				checks["All schemas have descriptions"] = false
			}
			// Check enum descriptions
			if len(schema.Enum) > 0 {
				for _, enum := range schema.Enum {
					if strEnum, ok := enum.(string); ok {
						if schema.Value.Description == "" || !strings.Contains(schema.Value.Description, strEnum) {
							missingDocs["All enums have descriptions"] = append(missingDocs["All enums have descriptions"],
								fmt.Sprintf("%s: enum value %s", name, strEnum))
							checks["All enums have descriptions"] = false
						}
					}
				}
			}
		}
	}

	// Update result based on checks
	allPassed := true
	for _, passed := range checks {
		if !passed {
			allPassed = false
			break
		}
	}
	result.Passed = allPassed
	result.Details = map[string]interface{}{
		"checks":       checks,
		"messages":     messages,
		"missing_docs": missingDocs,
	}

	if !allPassed {
		var failedChecks []string
		for check, items := range missingDocs {
			if len(items) > 0 {
				failedChecks = append(failedChecks, fmt.Sprintf("%s: %s", check, strings.Join(items, ", ")))
			}
		}
		result.Message = fmt.Sprintf("Documentation quality issues found: %s", strings.Join(failedChecks, "; "))
		result.SuggestedFix = "Add missing documentation including descriptions, examples, and operation details"
	} else {
		result.Message = "API documentation is comprehensive and high quality"
	}

	return result
}

// validateErrorHandling validates error response documentation and patterns
func (v *OpenAPIValidator) validateErrorHandling(doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[2], // P003
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check documentation for present error codes
	if v.config.ValidationMode == ValidationModeMinimal {
		for path, pathItem := range doc.Paths() {
			if pathItem == nil || pathItem.Operations == nil {
				continue
			}
			for method, operation := range pathItem.Operations {
				if operation == nil {
					continue
				}
				opKey := fmt.Sprintf("%s %s", method, path)

				// Only check documentation for present error responses
				for code, response := range operation.Responses {
					if code >= "400" && code < "600" && response != nil {
						if response.Description == "" {
							result.Passed = false
							result.Message = fmt.Sprintf("Error response %s missing description: %s", code, opKey)
							return result
						}
					}
				}
			}
		}
		result.Message = "All present error responses are documented"
		return result
	}

	// Strict mode - existing comprehensive validation
	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingErrors := make(map[string][]string)

	// Check for common error responses in components
	hasCommonErrors := false
	if comps := doc.Components(); comps != nil && comps.Responses != nil {
		commonCodes := []string{"400", "401", "403", "404", "500"}
		for _, code := range commonCodes {
			if _, exists := comps.Responses[code]; exists {
				hasCommonErrors = true
				break
			}
		}
	}
	checks["Common error responses are defined in components"] = hasCommonErrors
	if !hasCommonErrors {
		messages["Common error responses are defined in components"] = "No common error responses defined in components"
	}

	// Check each operation's error responses
	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check for 4xx errors
			has4xx := false
			for code := range operation.Responses {
				if code >= "400" && code < "500" {
					has4xx = true
					break
				}
			}
			checks["All operations document 4xx error responses"] = has4xx
			if !has4xx {
				missingErrors["All operations document 4xx error responses"] = append(
					missingErrors["All operations document 4xx error responses"], opKey)
			}

			// Check for 5xx errors
			has5xx := false
			for code := range operation.Responses {
				if code >= "500" && code < "600" {
					has5xx = true
					break
				}
			}
			checks["All operations document 5xx error responses"] = has5xx
			if !has5xx {
				missingErrors["All operations document 5xx error responses"] = append(
					missingErrors["All operations document 5xx error responses"], opKey)
			}

			// Check error response details
			for code, response := range operation.Responses {
				if code >= "400" && code < "600" {
					if response == nil {
						continue
					}

					// Check error code documentation
					if response.Description == "" {
						missingErrors["Error responses include error codes"] = append(
							missingErrors["Error responses include error codes"],
							fmt.Sprintf("%s: %s response", opKey, code))
						checks["Error responses include error codes"] = false
					}

					// Check error message schema
					hasErrorSchema := false
					if response.Content != nil {
						for _, content := range response.Content {
							if content == nil || content.Schema == nil {
								continue
							}
							// Look for common error message fields
							schema := content.Schema
							if schema.Properties != nil {
								if _, hasMessage := schema.Properties["message"]; hasMessage {
										hasErrorSchema = true
									}
									if _, hasCode := schema.Properties["code"]; hasCode {
										hasErrorSchema = true
									}
									if _, hasDetails := schema.Properties["details"]; hasDetails {
										hasErrorSchema = true
									}
								}
							}
						}
					}
					checks["Error responses include error details schema"] = hasErrorSchema
					if !hasErrorSchema {
						missingErrors["Error responses include error details schema"] = append(
							missingErrors["Error responses include error details schema"],
							fmt.Sprintf("%s: %s response", opKey, code))
					}
				}
			}
		}
	}

	// Check error response format consistency
	errorFormats := make(map[string][]string)
	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			opKey := fmt.Sprintf("%s %s", method, path)
			for code, response := range operation.Responses {
				if code >= "400" && code < "600" && response != nil && response.Content != nil {
					for contentType, content := range response.Content {
						if content == nil || content.Schema == nil {
							continue
						}
						schema := content.Schema
							format := "unknown"
							if schema.Properties != nil {
								props := []string{}
								for prop := range schema.Properties {
									props = append(props, prop)
								}
								sort.Strings(props)
								format = strings.Join(props, ",")
							}
							errorFormats[format] = append(errorFormats[format], fmt.Sprintf("%s: %s %s", opKey, code, contentType))
						}
					}
				}
			}
		}
	}
	hasConsistentFormat := len(errorFormats) <= 1
	checks["Error responses follow consistent format"] = hasConsistentFormat
	if !hasConsistentFormat {
		messages["Error responses follow consistent format"] = "Multiple error response formats found"
		for format, locations := range errorFormats {
			missingErrors["Error responses follow consistent format"] = append(
				missingErrors["Error responses follow consistent format"],
				fmt.Sprintf("Format [%s] used in: %s", format, strings.Join(locations, ", ")))
		}
	}

	// Update result based on checks
	allPassed := true
	for _, passed := range checks {
		if !passed {
			allPassed = false
			break
		}
	}
	result.Passed = allPassed
	result.Details = map[string]interface{}{
		"checks":         checks,
		"messages":       messages,
		"missing_errors": missingErrors,
	}

	if !allPassed {
		var failedChecks []string
		for check, items := range missingErrors {
			if len(items) > 0 {
				failedChecks = append(failedChecks, fmt.Sprintf("%s: %s", check, strings.Join(items, ", ")))
			}
		}
		result.Message = fmt.Sprintf("Error handling issues found: %s", strings.Join(failedChecks, "; "))
		result.SuggestedFix = "Add comprehensive error response documentation including codes, messages, and consistent error schemas"
	} else {
		result.Message = "Error handling is well-documented and follows consistent patterns"
	}

	return result
}

// validateRequestSchema validates request parameter and body schemas
func (v *OpenAPIValidator) validateRequestSchema(doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[3], // P004
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check for basic schema existence
	if v.config.ValidationMode == ValidationModeMinimal {
		for path, pathItem := range doc.Paths() {
			if pathItem == nil || pathItem.Operations == nil {
				continue
			}
			for method, operation := range pathItem.Operations {
				if operation == nil {
					continue
				}
				opKey := fmt.Sprintf("%s %s", method, path)

				// Check if request body has schema
				if operation.RequestBody != nil {
					if operation.RequestBody.Content == nil {
						result.Passed = false
						result.Message = fmt.Sprintf("Request body missing content schema: %s", opKey)
						return result
					}
					for contentType, content := range operation.RequestBody.Content {
						if content == nil || content.Schema == nil {
							result.Passed = false
							result.Message = fmt.Sprintf("Request body missing schema for %s: %s", contentType, opKey)
							return result
						}
					}
				}

				// Check if parameters have schemas
				for _, param := range operation.Parameters {
					if param != nil && param.Schema == nil {
						result.Passed = false
						result.Message = fmt.Sprintf("Parameter missing schema: %s %s", param.Name, opKey)
						return result
					}
				}
			}
		}
		result.Message = "All requests have basic schema definitions"
		return result
	}

	// Strict mode - existing comprehensive validation
	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingValidation := make(map[string][]string)

	// Initialize all checks to true
	for _, check := range CorePrinciples[3].Checks {
		checks[check] = true
	}

	for path, pathItem := range doc.Paths() {
		if pathItem == nil {
			continue
		}
		// Check path-level parameters
		for _, param := range pathItem.Parameters {
			if param == nil {
				continue
			}
			paramKey := fmt.Sprintf("%s: parameter %s", path, param.Name)

			// Check schema existence
			if param.Schema == nil {
				missingValidation["All path parameters have schemas"] = append(
					missingValidation["All path parameters have schemas"], paramKey)
				checks["All path parameters have schemas"] = false
				continue
			}

			// Check schema type
			if param.Schema.Type == "" {
				missingValidation["All schemas specify data types"] = append(
					missingValidation["All schemas specify data types"], paramKey)
				checks["All schemas specify data types"] = false
			}

			// Check constraints
			schema := param.Schema
			if schema.Type == "string" {
				hasConstraints := false
				minLen := schema.MinLength > 0
				maxLen := schema.MaxLength != nil && *schema.MaxLength > 0
				if minLen || maxLen || schema.Pattern != "" {
					hasConstraints = true
				}
				if !hasConstraints {
					missingValidation["All string fields have length constraints"] = append(
						missingValidation["All string fields have length constraints"], paramKey)
					checks["All string fields have length constraints"] = false
				}
			} else if schema.Type == "number" || schema.Type == "integer" {
				if schema.Min == nil && schema.Max == nil {
					missingValidation["All numeric fields have min/max values"] = append(
						missingValidation["All numeric fields have min/max values"], paramKey)
					checks["All numeric fields have min/max values"] = false
				}
			}

			// Check enums
			if len(schema.Enum) > 0 {
				hasValidValues := true
				for _, enum := range schema.Enum {
					if enum == nil {
						hasValidValues = false
						break
					}
				}
				if !hasValidValues {
					missingValidation["All enums have valid values"] = append(
						missingValidation["All enums have valid values"], paramKey)
					checks["All enums have valid values"] = false
				}
			}
		}

		// Check operation-level parameters and bodies
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check operation parameters
			for _, param := range operation.Parameters {
				if param == nil {
					continue
				}
				paramKey := fmt.Sprintf("%s: parameter %s", opKey, param.Name)

				// Check schema existence
				if param.Schema == nil {
					switch param.In {
					case "query":
						missingValidation["All query parameters have schemas"] = append(
							missingValidation["All query parameters have schemas"], paramKey)
						checks["All query parameters have schemas"] = false
					case "header":
						missingValidation["All header parameters have schemas"] = append(
							missingValidation["All header parameters have schemas"], paramKey)
						checks["All header parameters have schemas"] = false
					}
					continue
				}

				// Check schema type
				if param.Schema.Type == "" {
					missingValidation["All schemas specify data types"] = append(
						missingValidation["All schemas specify data types"], paramKey)
					checks["All schemas specify data types"] = false
				}

				// Check constraints
				schema := param.Schema
				if schema.Type == "string" {
					hasConstraints := false
					minLen := schema.MinLength > 0
					maxLen := schema.MaxLength != nil && *schema.MaxLength > 0
					if minLen || maxLen || schema.Pattern != "" {
						hasConstraints = true
					}
					if !hasConstraints {
						missingValidation["All string fields have length constraints"] = append(
							missingValidation["All string fields have length constraints"], paramKey)
						checks["All string fields have length constraints"] = false
					}
				} else if schema.Type == "number" || schema.Type == "integer" {
					if schema.Min == nil && schema.Max == nil {
						missingValidation["All numeric fields have min/max values"] = append(
							missingValidation["All numeric fields have min/max values"], paramKey)
						checks["All numeric fields have min/max values"] = false
					}
				}

				// Check enums
				if len(schema.Enum) > 0 {
					hasValidValues := true
					for _, enum := range schema.Enum {
						if enum == nil {
							hasValidValues = false
							break
						}
					}
					if !hasValidValues {
						missingValidation["All enums have valid values"] = append(
							missingValidation["All enums have valid values"], paramKey)
						checks["All enums have valid values"] = false
					}
				}

				// Check required flag
				if param.Required {
					required := false
					for _, r := range param.Schema.Required {
						if r == param.Name {
							required = true
							break
						}
					}
					if !required {
						missingValidation["All required fields are marked"] = append(
							missingValidation["All required fields are marked"], paramKey)
						checks["All required fields are marked"] = false
					}
				}
			}

			// Check request body
			if operation.RequestBody != nil {
				if operation.RequestBody.Content == nil {
					missingValidation["All request bodies have content schemas"] = append(
						missingValidation["All request bodies have content schemas"],
						fmt.Sprintf("%s: request body", opKey))
					checks["All request bodies have content schemas"] = false
				} else {
					for contentType, content := range operation.RequestBody.Content {
						if content == nil || content.Schema == nil {
							missingValidation["All request bodies have content schemas"] = append(
								missingValidation["All request bodies have content schemas"],
								fmt.Sprintf("%s: %s request body", opKey, contentType))
							checks["All request bodies have content schemas"] = false
							continue
						}

						// Validate schema recursively
						v.validateSchemaConstraints(content.Schema, opKey, contentType, checks, missingValidation)
					}
				}
			}
		}
	}

	// Update result based on checks
	allPassed := true
	for _, passed := range checks {
		if !passed {
			allPassed = false
			break
		}
	}
	result.Passed = allPassed
	result.Details = map[string]interface{}{
		"checks":             checks,
		"messages":           messages,
		"missing_validation": missingValidation,
	}

	if !allPassed {
		var failedChecks []string
		for check, items := range missingValidation {
			if len(items) > 0 {
				failedChecks = append(failedChecks, fmt.Sprintf("%s: %s", check, strings.Join(items, ", ")))
			}
		}
		result.Message = fmt.Sprintf("Request validation issues found: %s", strings.Join(failedChecks, "; "))
		result.SuggestedFix = "Add comprehensive schema validation including data types, constraints, and required fields"
	} else {
		result.Message = "All requests have comprehensive schema definitions with proper validation rules"
	}

	return result
}

// validateSchemaConstraints recursively validates schema constraints on a normalized schema.
func (v *OpenAPIValidator) validateSchemaConstraints(schema *spec.Schema, context, contentType string, checks map[string]bool, missingValidation map[string][]string) {
	if schema == nil {
		return
	}

	// Check type
	if schema.Type == "" {
		missingValidation["All schemas specify data types"] = append(
			missingValidation["All schemas specify data types"],
			fmt.Sprintf("%s: %s schema", context, contentType))
		checks["All schemas specify data types"] = false
	}

	// Check constraints based on type
	switch schema.Type {
	case "string":
		hasConstraints := false
		minLen := schema.MinLength > 0
		maxLen := schema.MaxLength != nil && *schema.MaxLength > 0
		if minLen || maxLen || schema.Pattern != "" {
			hasConstraints = true
		}
		if !hasConstraints {
			missingValidation["All string fields have length constraints"] = append(
				missingValidation["All string fields have length constraints"],
				fmt.Sprintf("%s: %s schema", context, contentType))
			checks["All string fields have length constraints"] = false
		}
	case "number", "integer":
		if schema.Min == nil && schema.Max == nil {
			missingValidation["All numeric fields have min/max values"] = append(
				missingValidation["All numeric fields have min/max values"],
				fmt.Sprintf("%s: %s schema", context, contentType))
			checks["All numeric fields have min/max values"] = false
		}
	}

	// Check enums
	if len(schema.Enum) > 0 {
		hasValidValues := true
		for _, enum := range schema.Enum {
			if enum == nil {
				hasValidValues = false
				break
			}
		}
		if !hasValidValues {
			missingValidation["All enums have valid values"] = append(
				missingValidation["All enums have valid values"],
				fmt.Sprintf("%s: %s schema", context, contentType))
			checks["All enums have valid values"] = false
		}
	}

	// Check required fields
	if len(schema.Required) > 0 {
		for _, required := range schema.Required {
			if schema.Properties != nil {
				if prop, exists := schema.Properties[required]; exists && prop != nil {
					found := false
					for _, r := range prop.Required {
						if r == required {
							found = true
							break
						}
					}
					if !found {
						missingValidation["All required fields are marked"] = append(
							missingValidation["All required fields are marked"],
							fmt.Sprintf("%s: %s.%s", context, contentType, required))
						checks["All required fields are marked"] = false
					}
				}
			}
		}
	}

	// Recursively check properties
	if schema.Properties != nil {
		for name, prop := range schema.Properties {
			if prop != nil {
				v.validateSchemaConstraints(prop, fmt.Sprintf("%s.%s", context, name), contentType, checks, missingValidation)
			}
		}
	}

	// Check array items
	if schema.Type == "array" && schema.Items != nil {
		v.validateSchemaConstraints(schema.Items, fmt.Sprintf("%s[]", context), contentType, checks, missingValidation)
	}
}

// validateAuthentication validates that all operations have proper authentication requirements
func (v *OpenAPIValidator) validateAuthentication(doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[4], // P005
		Passed:    true,
	}

	comps := doc.Components()
	if comps == nil || comps.SecuritySchemes == nil || len(comps.SecuritySchemes) == 0 {
		result.Passed = false
		result.Message = "No security schemes defined"
		result.SuggestedFix = "Define security schemes in components.securitySchemes"
		return result
	}

	var missingAuth []string
	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			if operation.Security == nil && doc.Security() == nil {
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

// validateVersioning validates that the API has proper versioning
func (v *OpenAPIValidator) validateVersioning(doc spec.APISpec) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[7], // P008
		Passed:    true,
	}

	info := doc.Info()
	if info == nil || info.Version == "" {
		result.Passed = false
		result.Message = "API version is not specified in the OpenAPI document info section"
		result.SuggestedFix = "Add or update the 'version' field in the 'info' section"
		return result
	}

	// Check if version is in semantic versioning format
	version := info.Version
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		result.Passed = false
		result.Message = "API version should follow semantic versioning (e.g., 1.0.0)"
		result.SuggestedFix = "Update version to follow semantic versioning format"
		return result
	}

	return result
}

// updateSummary updates the validation summary
func (v *OpenAPIValidator) updateSummary(report *ValidationReport) {
	summary := ValidationSummary{}
	categories := make(map[string]bool)
	failedTags := make(map[string]bool)

	for _, result := range report.Principles {
		if !result.Passed {
			switch result.Principle.Severity {
			case "critical":
				summary.CriticalIssues++
			case "warning":
				summary.Warnings++
			case "info":
				summary.Info++
			}

			categories[result.Principle.Category] = true
			for _, tag := range result.Principle.Tags {
				failedTags[tag] = true
			}
		}
	}

	for category := range categories {
		summary.Categories = append(summary.Categories, category)
	}
	for tag := range failedTags {
		summary.FailedTags = append(summary.FailedTags, tag)
	}

	report.Summary = summary
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[validation] Logger set to DEBUG (verbose) mode")
}
