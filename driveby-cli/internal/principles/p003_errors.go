package principles

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P003Errors validates error response documentation and patterns.
type P003Errors struct{}

func (p *P003Errors) ID() string { return "P003" }

func (p *P003Errors) Check(_ context.Context, doc spec.APISpec, mode types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[2],
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check documentation for present error codes
	if mode == types.ValidationModeMinimal {
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

	// Test-ready mode: check what functional tests need — 4xx responses exist
	// and error responses have schemas so test assertions can validate error shapes.
	// Skip: common components, 5xx on every op, format consistency.
	isTestReady := mode == types.ValidationModeTestReady

	// Strict mode - comprehensive validation
	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingErrors := make(map[string][]string)

	// Check for common error responses in components (strict only)
	if !isTestReady {
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

			// Check for 5xx errors (strict only — functional tests focus on 4xx)
			if !isTestReady {
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

	// Check error response format consistency (strict only)
	if !isTestReady {
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
