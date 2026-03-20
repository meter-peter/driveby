package principles

import (
	"context"
	"fmt"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P009TestReadiness validates that the API specification provides enough
// detail (examples, typed schemas, response schemas) for meaningful
// functional and load testing.
type P009TestReadiness struct{}

func (p *P009TestReadiness) ID() string { return "P009" }

func (p *P009TestReadiness) Check(_ context.Context, doc spec.APISpec, _ types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[8], // P009 is index 8
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	checks := map[string]bool{
		"Path parameters have examples or typed schemas":      true,
		"Request bodies have examples or typed schema properties": true,
		"2xx responses have schemas defined":                  true,
		"2xx responses have examples defined":                 true,
	}
	issues := make(map[string][]string)
	var warnings []string

	for path, pathItem := range doc.Paths() {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		for method, operation := range pathItem.Operations {
			if operation == nil {
				continue
			}
			opKey := fmt.Sprintf("%s %s", method, path)

			// Check path parameters have example or typed schema
			for _, param := range operation.Parameters {
				if param == nil || param.In != "path" {
					continue
				}
				hasExample := param.Example != nil
				hasTypedSchema := param.Schema != nil && param.Schema.Type != ""
				if !hasExample && !hasTypedSchema {
					issues["Path parameters have examples or typed schemas"] = append(
						issues["Path parameters have examples or typed schemas"],
						fmt.Sprintf("%s: param %s", opKey, param.Name))
					checks["Path parameters have examples or typed schemas"] = false
				}
			}

			// Also check path-level parameters
			for _, param := range pathItem.Parameters {
				if param == nil || param.In != "path" {
					continue
				}
				hasExample := param.Example != nil
				hasTypedSchema := param.Schema != nil && param.Schema.Type != ""
				if !hasExample && !hasTypedSchema {
					issues["Path parameters have examples or typed schemas"] = append(
						issues["Path parameters have examples or typed schemas"],
						fmt.Sprintf("%s: param %s (path-level)", opKey, param.Name))
					checks["Path parameters have examples or typed schemas"] = false
				}
			}

			// Check request bodies have examples or typed schema properties
			if operation.RequestBody != nil && operation.RequestBody.Content != nil {
				for contentType, media := range operation.RequestBody.Content {
					if media == nil {
						continue
					}
					hasExample := media.Example != nil
					hasExamples := len(media.Examples) > 0
					hasTypedProps := p.schemaHasTypedProperties(media.Schema)
					if !hasExample && !hasExamples && !hasTypedProps {
						issues["Request bodies have examples or typed schema properties"] = append(
							issues["Request bodies have examples or typed schema properties"],
							fmt.Sprintf("%s: %s body", opKey, contentType))
						checks["Request bodies have examples or typed schema properties"] = false
					}
				}
			}

			// Check 2xx responses have schemas and examples
			if operation.Responses != nil {
				for statusCode, resp := range operation.Responses {
					if resp == nil || !strings.HasPrefix(statusCode, "2") {
						continue
					}
					// 204 No Content legitimately has no body
					if statusCode == "204" {
						continue
					}
					respKey := fmt.Sprintf("%s: %s response", opKey, statusCode)

					if resp.Content == nil || len(resp.Content) == 0 {
						// No content at all — schema missing
						issues["2xx responses have schemas defined"] = append(
							issues["2xx responses have schemas defined"], respKey)
						checks["2xx responses have schemas defined"] = false
						continue
					}

					for _, media := range resp.Content {
						if media == nil || media.Schema == nil {
							issues["2xx responses have schemas defined"] = append(
								issues["2xx responses have schemas defined"], respKey)
							checks["2xx responses have schemas defined"] = false
						}
						// Response examples are advisory (warning, not failure)
						if media != nil && media.Example == nil && len(media.Examples) == 0 {
							warnings = append(warnings,
								fmt.Sprintf("%s: no response example", respKey))
							checks["2xx responses have examples defined"] = false
						}
					}
				}
			}
		}
	}

	// Response examples are advisory — don't fail the check
	allPassed := checks["Path parameters have examples or typed schemas"] &&
		checks["Request bodies have examples or typed schema properties"] &&
		checks["2xx responses have schemas defined"]

	result.Passed = allPassed
	result.Details = map[string]interface{}{
		"checks":   checks,
		"issues":   issues,
		"warnings": warnings,
	}

	if !allPassed {
		var failedChecks []string
		for check, items := range issues {
			if len(items) > 0 {
				failedChecks = append(failedChecks, fmt.Sprintf("%s: %s", check, strings.Join(items, ", ")))
			}
		}
		result.Message = fmt.Sprintf("Test readiness issues found: %s", strings.Join(failedChecks, "; "))
		result.SuggestedFix = "Add examples to path parameters and request bodies, and schemas to 2xx responses"
	} else {
		result.Message = "Specification is ready for meaningful functional and load testing"
		if len(warnings) > 0 {
			result.Message += fmt.Sprintf(" (advisory: %d response(s) missing examples)", len(warnings))
		}
	}

	return result
}

// schemaHasTypedProperties returns true if the schema has at least one
// property with a non-empty type, meaning GenerateExampleFromSchema can
// produce typed values.
func (p *P009TestReadiness) schemaHasTypedProperties(schema *spec.Schema) bool {
	if schema == nil {
		return false
	}
	// A typed schema itself is enough (e.g., type: string)
	if schema.Type != "" && schema.Type != "object" {
		return true
	}
	// For objects, check that properties exist with types
	if schema.Properties != nil {
		for _, prop := range schema.Properties {
			if prop != nil && prop.Type != "" {
				return true
			}
		}
	}
	// Check allOf
	for _, sub := range schema.AllOf {
		if p.schemaHasTypedProperties(sub) {
			return true
		}
	}
	return false
}
