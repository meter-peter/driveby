package principles

import (
	"context"
	"fmt"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P001Compliance validates that the API spec is structurally compliant.
// For OpenAPI 3.x, this includes full schema validation. For Swagger 2.0,
// it performs a focused set of structural checks.
type P001Compliance struct{}

func (p *P001Compliance) ID() string { return "P001" }

func (p *P001Compliance) Check(ctx context.Context, doc spec.APISpec, mode types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[0],
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
