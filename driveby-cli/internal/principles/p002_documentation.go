package principles

import (
	"context"
	"fmt"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P002Documentation validates the quality and completeness of API documentation.
type P002Documentation struct{}

func (p *P002Documentation) ID() string { return "P002" }

func (p *P002Documentation) Check(_ context.Context, doc spec.APISpec, _ types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[1],
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
						if schema.Description == "" || !strings.Contains(schema.Description, strEnum) {
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
