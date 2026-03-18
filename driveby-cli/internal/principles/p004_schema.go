package principles

import (
	"context"
	"fmt"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P004Schema validates request parameter and body schemas.
type P004Schema struct{}

func (p *P004Schema) ID() string { return "P004" }

func (p *P004Schema) Check(_ context.Context, doc spec.APISpec, mode types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[3],
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	// In minimal mode, only check for basic schema existence
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

	// Strict mode - comprehensive validation
	checks := make(map[string]bool)
	messages := make(map[string]string)
	missingValidation := make(map[string][]string)

	// Initialize all checks to true
	for _, check := range types.CorePrinciples[3].Checks {
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
						p.validateSchemaConstraints(content.Schema, opKey, contentType, checks, missingValidation)
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
func (p *P004Schema) validateSchemaConstraints(schema *spec.Schema, ctx, contentType string, checks map[string]bool, missingValidation map[string][]string) {
	if schema == nil {
		return
	}

	// Check type
	if schema.Type == "" {
		missingValidation["All schemas specify data types"] = append(
			missingValidation["All schemas specify data types"],
			fmt.Sprintf("%s: %s schema", ctx, contentType))
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
				fmt.Sprintf("%s: %s schema", ctx, contentType))
			checks["All string fields have length constraints"] = false
		}
	case "number", "integer":
		if schema.Min == nil && schema.Max == nil {
			missingValidation["All numeric fields have min/max values"] = append(
				missingValidation["All numeric fields have min/max values"],
				fmt.Sprintf("%s: %s schema", ctx, contentType))
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
				fmt.Sprintf("%s: %s schema", ctx, contentType))
			checks["All enums have valid values"] = false
		}
	}

	// Check required fields: verify that fields listed in schema.Required exist in schema.Properties
	if len(schema.Required) > 0 && schema.Properties != nil {
		for _, required := range schema.Required {
			if _, exists := schema.Properties[required]; !exists {
				missingValidation["All required fields are marked"] = append(
					missingValidation["All required fields are marked"],
					fmt.Sprintf("%s: %s.%s (listed as required but not in properties)", ctx, contentType, required))
				checks["All required fields are marked"] = false
			}
		}
	}

	// Recursively check properties
	if schema.Properties != nil {
		for name, prop := range schema.Properties {
			if prop != nil {
				p.validateSchemaConstraints(prop, fmt.Sprintf("%s.%s", ctx, name), contentType, checks, missingValidation)
			}
		}
	}

	// Check array items
	if schema.Type == "array" && schema.Items != nil {
		p.validateSchemaConstraints(schema.Items, fmt.Sprintf("%s[]", ctx), contentType, checks, missingValidation)
	}
}
