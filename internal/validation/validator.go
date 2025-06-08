package validation

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/openapi"
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

	// Load OpenAPI spec
	if err := v.loader.LoadFromFileOrURL(v.config.SpecPath); err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := v.loader.GetDocument()
	if doc == nil {
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	report := &ValidationReport{
		Version:     v.config.Version,
		Environment: v.config.Environment,
		Timestamp:   time.Now(),
	}

	// In minimal mode, only run basic validation principles (P001, P004)
	// Skip functional and performance testing
	var validationPrinciples []Principle
	if v.config.ValidationMode == ValidationModeMinimal {
		validationPrinciples = []Principle{
			CorePrinciples[0], // P001: OpenAPI Specification Compliance
			CorePrinciples[3], // P004: Request Validation (basic schema checks only)
		}
		log.Debug("Running in minimal mode - skipping functional and performance testing")
	} else {
		// Strict mode - run all validation principles
		validationPrinciples = []Principle{
			CorePrinciples[0], // P001: OpenAPI Specification Compliance
			CorePrinciples[1], // P002: API Documentation Completeness
			CorePrinciples[2], // P003: Error Response Documentation
			CorePrinciples[3], // P004: Request Validation
			CorePrinciples[4], // P005: Authentication Requirements
			CorePrinciples[7], // P008: API Versioning
		}
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
func (v *OpenAPIValidator) validatePrinciple(ctx context.Context, principle Principle, doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: principle,
		Passed:    true,
	}

	switch principle.ID {
	case "P001": // OpenAPI Specification Compliance
		result = v.validateOpenAPICompliance(doc)
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

// validateOpenAPICompliance validates that the OpenAPI spec is compliant with the OpenAPI 3.0/3.1 schema
func (v *OpenAPIValidator) validateOpenAPICompliance(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[0], // P001
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// Track individual check results
	checks := make(map[string]bool)
	messages := make(map[string]string)

	// Check OpenAPI version
	if doc.OpenAPI == "" {
		checks["OpenAPI version is 3.0.x or 3.1.0"] = false
		messages["OpenAPI version is 3.0.x or 3.1.0"] = "OpenAPI version is not specified"
	} else if !strings.HasPrefix(doc.OpenAPI, "3.0.") && doc.OpenAPI != "3.1.0" {
		checks["OpenAPI version is 3.0.x or 3.1.0"] = false
		messages["OpenAPI version is 3.0.x or 3.1.0"] = fmt.Sprintf("OpenAPI version %s is not 3.0.x or 3.1.0", doc.OpenAPI)
	} else {
		checks["OpenAPI version is 3.0.x or 3.1.0"] = true
	}

	// Check required info fields
	if doc.Info == nil {
		checks["Required info fields (title, version) are present"] = false
		messages["Required info fields (title, version) are present"] = "Info section is missing"
	} else {
		missingFields := []string{}
		if doc.Info.Title == "" {
			missingFields = append(missingFields, "title")
		}
		if doc.Info.Version == "" {
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
	if doc.Paths == nil || len(doc.Paths.Map()) == 0 {
		checks["Paths are properly defined"] = false
		messages["Paths are properly defined"] = "No paths defined in the API"
	} else {
		checks["Paths are properly defined"] = true
	}

	// Check components
	if doc.Components == nil {
		checks["Components are valid"] = false
		messages["Components are valid"] = "Components section is missing"
	} else {
		checks["Components are valid"] = true
	}

	// Check references
	refErrors := []string{}
	if err := doc.Validate(context.Background()); err != nil {
		checks["References are resolvable"] = false
		refErrors = append(refErrors, err.Error())
	} else {
		checks["References are resolvable"] = true
	}
	if len(refErrors) > 0 {
		messages["References are resolvable"] = strings.Join(refErrors, "; ")
	}

	// Check for duplicate operationIds
	operationIDs := make(map[string][]string)
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if operation.OperationID != "" {
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
	for path, pathItem := range doc.Paths.Map() {
		for method := range pathItem.Operations() {
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

	// Check references and components
	if doc.Components != nil {
		// Check for null type in schemas
		for name, schema := range doc.Components.Schemas {
			if schema.Value != nil {
				// Allow null type in OpenAPI 3.1.0
				if schema.Value.Type == "null" && doc.OpenAPI == "3.1.0" {
					continue
				}
				// For OpenAPI 3.0.x, null type should be represented as ["null", "type"]
				if schema.Value.Type == "null" && strings.HasPrefix(doc.OpenAPI, "3.0.") {
					checks["References are resolvable"] = false
					messages["References are resolvable"] = fmt.Sprintf("invalid components: schema %q: 'null' type should be represented as [\"null\", \"type\"] in OpenAPI 3.0.x", name)
					continue
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
func (v *OpenAPIValidator) validateDocumentationQuality(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[1], // P002
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingDocs := make(map[string][]string)

	// Check API-level documentation
	if doc.Info == nil {
		checks["API has a general description"] = false
		messages["API has a general description"] = "Info section is missing"
	} else if doc.Info.Description == "" {
		checks["API has a general description"] = false
		messages["API has a general description"] = "API description is missing"
	} else {
		checks["API has a general description"] = true
	}

	// Check contact information
	if doc.Info == nil || doc.Info.Contact == nil {
		checks["Contact information is provided"] = false
		messages["Contact information is provided"] = "Contact information is missing"
	} else {
		hasContact := doc.Info.Contact.Name != "" || doc.Info.Contact.Email != "" || doc.Info.Contact.URL != ""
		checks["Contact information is provided"] = hasContact
		if !hasContact {
			messages["Contact information is provided"] = "Contact information is empty"
		}
	}

	// Check license information
	if doc.Info == nil || doc.Info.License == nil {
		checks["License information is provided"] = false
		messages["License information is provided"] = "License information is missing"
	} else if doc.Info.License.Name == "" {
		checks["License information is provided"] = false
		messages["License information is provided"] = "License name is missing"
	} else {
		checks["License information is provided"] = true
	}

	// Check operation documentation
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
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
				if param.Value == nil {
					continue
				}
				if param.Value.Description == "" {
					missingDocs["All parameters have descriptions"] = append(missingDocs["All parameters have descriptions"],
						fmt.Sprintf("%s: parameter %s", opKey, param.Value.Name))
					checks["All parameters have descriptions"] = false
				}
			}

			// Check request body documentation
			if operation.RequestBody != nil && operation.RequestBody.Value != nil {
				if operation.RequestBody.Value.Description == "" {
					missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
						fmt.Sprintf("%s: request body", opKey))
					checks["All request/response bodies have examples"] = false
				}
				// Check for examples in content
				for contentType, content := range operation.RequestBody.Value.Content {
					if content.Example == nil && len(content.Examples) == 0 {
						missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
							fmt.Sprintf("%s: %s request body", opKey, contentType))
						checks["All request/response bodies have examples"] = false
					}
				}
			}

			// Check response documentation
			for status, response := range operation.Responses.Map() {
				if response.Value == nil {
					continue
				}
				if response.Value.Description == nil || *response.Value.Description == "" {
					missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
						fmt.Sprintf("%s: %s response", opKey, status))
					checks["All request/response bodies have examples"] = false
				}
				// Check for examples in content
				for contentType, content := range response.Value.Content {
					if content.Example == nil && len(content.Examples) == 0 {
						missingDocs["All request/response bodies have examples"] = append(missingDocs["All request/response bodies have examples"],
							fmt.Sprintf("%s: %s %s response", opKey, status, contentType))
						checks["All request/response bodies have examples"] = false
					}
				}
			}
		}
	}

	// Check schema documentation
	if doc.Components != nil && doc.Components.Schemas != nil {
		for name, schema := range doc.Components.Schemas {
			if schema.Value == nil {
				continue
			}
			if schema.Value.Description == "" {
				missingDocs["All schemas have descriptions"] = append(missingDocs["All schemas have descriptions"], name)
				checks["All schemas have descriptions"] = false
			}
			// Check enum descriptions
			if len(schema.Value.Enum) > 0 {
				for _, enum := range schema.Value.Enum {
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
func (v *OpenAPIValidator) validateErrorHandling(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[2], // P003
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check documentation for present error codes
	if v.config.ValidationMode == ValidationModeMinimal {
		for path, pathItem := range doc.Paths.Map() {
			for method, operation := range pathItem.Operations() {
				opKey := fmt.Sprintf("%s %s", method, path)

				// Only check documentation for present error responses
				for code, response := range operation.Responses.Map() {
					if code >= "400" && code < "600" && response.Value != nil {
						if response.Value.Description == nil || *response.Value.Description == "" {
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
	if doc.Components != nil && doc.Components.Responses != nil {
		commonCodes := []string{"400", "401", "403", "404", "500"}
		for _, code := range commonCodes {
			if _, exists := doc.Components.Responses[code]; exists {
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
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check for 4xx errors
			has4xx := false
			for code := range operation.Responses.Map() {
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
			for code := range operation.Responses.Map() {
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
			for code, response := range operation.Responses.Map() {
				if code >= "400" && code < "600" {
					if response.Value == nil {
						continue
					}

					// Check error code documentation
					if response.Value.Description == nil || *response.Value.Description == "" {
						missingErrors["Error responses include error codes"] = append(
							missingErrors["Error responses include error codes"],
							fmt.Sprintf("%s: %s response", opKey, code))
						checks["Error responses include error codes"] = false
					}

					// Check error message schema
					hasErrorSchema := false
					if response.Value.Content != nil {
						for _, content := range response.Value.Content {
							if content.Schema != nil && content.Schema.Value != nil {
								// Look for common error message fields
								schema := content.Schema.Value
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
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			opKey := fmt.Sprintf("%s %s", method, path)
			for code, response := range operation.Responses.Map() {
				if code >= "400" && code < "600" && response.Value != nil && response.Value.Content != nil {
					for contentType, content := range response.Value.Content {
						if content.Schema != nil && content.Schema.Value != nil {
							schema := content.Schema.Value
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
func (v *OpenAPIValidator) validateRequestSchema(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[3], // P004
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check for basic schema existence
	if v.config.ValidationMode == ValidationModeMinimal {
		for path, pathItem := range doc.Paths.Map() {
			for method, operation := range pathItem.Operations() {
				opKey := fmt.Sprintf("%s %s", method, path)

				// Check if request body has schema
				if operation.RequestBody != nil && operation.RequestBody.Value != nil {
					if operation.RequestBody.Value.Content == nil {
						result.Passed = false
						result.Message = fmt.Sprintf("Request body missing content schema: %s", opKey)
						return result
					}
					for contentType, content := range operation.RequestBody.Value.Content {
						if content.Schema == nil {
							result.Passed = false
							result.Message = fmt.Sprintf("Request body missing schema for %s: %s", contentType, opKey)
							return result
						}
					}
				}

				// Check if parameters have schemas
				for _, param := range operation.Parameters {
					if param.Value != nil && param.Value.Schema == nil {
						result.Passed = false
						result.Message = fmt.Sprintf("Parameter missing schema: %s %s", param.Value.Name, opKey)
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

	for path, pathItem := range doc.Paths.Map() {
		// Check path-level parameters
		for _, param := range pathItem.Parameters {
			if param.Value == nil {
				continue
			}
			paramKey := fmt.Sprintf("%s: parameter %s", path, param.Value.Name)

			// Check schema existence
			if param.Value.Schema == nil {
				missingValidation["All path parameters have schemas"] = append(
					missingValidation["All path parameters have schemas"], paramKey)
				checks["All path parameters have schemas"] = false
				continue
			}

			// Check schema type
			if param.Value.Schema.Value.Type == "" {
				missingValidation["All schemas specify data types"] = append(
					missingValidation["All schemas specify data types"], paramKey)
				checks["All schemas specify data types"] = false
			}

			// Check constraints
			schema := param.Value.Schema.Value
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
		for method, operation := range pathItem.Operations() {
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check operation parameters
			for _, param := range operation.Parameters {
				if param.Value == nil {
					continue
				}
				paramKey := fmt.Sprintf("%s: parameter %s", opKey, param.Value.Name)

				// Check schema existence
				if param.Value.Schema == nil {
					switch param.Value.In {
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
				if param.Value.Schema.Value.Type == "" {
					missingValidation["All schemas specify data types"] = append(
						missingValidation["All schemas specify data types"], paramKey)
					checks["All schemas specify data types"] = false
				}

				// Check constraints
				schema := param.Value.Schema.Value
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
				if param.Value.Required {
					required := false
					for _, r := range param.Value.Schema.Value.Required {
						if r == param.Value.Name {
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
			if operation.RequestBody != nil && operation.RequestBody.Value != nil {
				if operation.RequestBody.Value.Content == nil {
					missingValidation["All request bodies have content schemas"] = append(
						missingValidation["All request bodies have content schemas"],
						fmt.Sprintf("%s: request body", opKey))
					checks["All request bodies have content schemas"] = false
				} else {
					for contentType, content := range operation.RequestBody.Value.Content {
						if content.Schema == nil {
							missingValidation["All request bodies have content schemas"] = append(
								missingValidation["All request bodies have content schemas"],
								fmt.Sprintf("%s: %s request body", opKey, contentType))
							checks["All request bodies have content schemas"] = false
							continue
						}

						// Validate schema recursively
						v.validateSchemaConstraints(content.Schema.Value, opKey, contentType, checks, missingValidation)
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

// validateSchemaConstraints recursively validates schema constraints
func (v *OpenAPIValidator) validateSchemaConstraints(schema *openapi3.Schema, context, contentType string, checks map[string]bool, missingValidation map[string][]string) {
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
				if prop, exists := schema.Properties[required]; exists && prop.Value != nil {
					found := false
					for _, r := range prop.Value.Required {
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
			if prop.Value != nil {
				v.validateSchemaConstraints(prop.Value, fmt.Sprintf("%s.%s", context, name), contentType, checks, missingValidation)
			}
		}
	}

	// Check array items
	if schema.Type == "array" && schema.Items != nil && schema.Items.Value != nil {
		v.validateSchemaConstraints(schema.Items.Value, fmt.Sprintf("%s[]", context), contentType, checks, missingValidation)
	}
}

// validateAuthentication validates that all operations have proper authentication requirements
func (v *OpenAPIValidator) validateAuthentication(doc *openapi3.T) PrincipleResult {
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

// validateVersioning validates that the API has proper versioning
func (v *OpenAPIValidator) validateVersioning(doc *openapi3.T) PrincipleResult {
	result := PrincipleResult{
		Principle: CorePrinciples[7], // P008
		Passed:    true,
	}

	if doc.Info == nil || doc.Info.Version == "" {
		result.Passed = false
		result.Message = "API version is not specified in the OpenAPI document info section"
		result.SuggestedFix = "Add or update the 'version' field in the 'info' section"
		return result
	}

	// Check if version is in semantic versioning format
	version := doc.Info.Version
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
