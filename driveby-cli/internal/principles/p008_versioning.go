package principles

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

var semverRegex = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

// P008Versioning validates that the API has proper versioning strategy.
type P008Versioning struct{}

func (p *P008Versioning) ID() string { return "P008" }

func (p *P008Versioning) Check(_ context.Context, doc spec.APISpec, mode types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[7],
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	checks := make(map[string]bool)
	messages := make(map[string]string)

	info := doc.Info()

	// Check 1: API version is specified
	if info == nil || info.Version == "" {
		checks["API version is specified"] = false
		messages["API version is specified"] = "API version is not specified in the info section"
	} else {
		checks["API version is specified"] = true
	}

	// Check 2: Version follows semantic versioning
	if info != nil && info.Version != "" {
		if semverRegex.MatchString(info.Version) {
			checks["Version follows semantic versioning"] = true
		} else {
			checks["Version follows semantic versioning"] = false
			messages["Version follows semantic versioning"] = fmt.Sprintf("Version %q does not match semver format (expected MAJOR.MINOR.PATCH)", info.Version)
		}
	} else {
		checks["Version follows semantic versioning"] = false
		messages["Version follows semantic versioning"] = "Cannot validate semver format: version is missing"
	}

	// In minimal mode, only checks 1 and 2 apply
	if mode != types.ValidationModeMinimal {
		desc := ""
		if info != nil {
			desc = info.Description
		}
		descLower := strings.ToLower(desc)

		// Check 3: Versioning strategy is documented
		if strings.Contains(descLower, "version") {
			checks["Versioning strategy is documented"] = true
		} else {
			checks["Versioning strategy is documented"] = false
			messages["Versioning strategy is documented"] = "Info description does not mention versioning strategy"
		}

		// Check 4: Deprecation notices are present
		p.checkDeprecation(doc, checks, messages)

		// Check 5: Breaking changes are documented
		if strings.Contains(descLower, "breaking") || strings.Contains(descLower, "changelog") {
			checks["Breaking changes are documented"] = true
		} else {
			checks["Breaking changes are documented"] = false
			messages["Breaking changes are documented"] = "Info description does not reference breaking changes or a changelog"
		}

		// Check 6: Version compatibility is specified
		if strings.Contains(descLower, "compatibility") || strings.Contains(descLower, "backward") {
			checks["Version compatibility is specified"] = true
		} else {
			checks["Version compatibility is specified"] = false
			messages["Version compatibility is specified"] = "Info description does not mention version compatibility"
		}

		// Check 7: Migration guides are referenced
		if strings.Contains(descLower, "migration") || strings.Contains(descLower, "upgrade") {
			checks["Migration guides are referenced"] = true
		} else {
			checks["Migration guides are referenced"] = false
			messages["Migration guides are referenced"] = "Info description does not reference migration or upgrade guides"
		}
	}

	// Determine overall pass/fail
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
		var failed []string
		for check, passed := range checks {
			if !passed {
				failed = append(failed, fmt.Sprintf("%s: %s", check, messages[check]))
			}
		}
		result.Message = fmt.Sprintf("Versioning validation failed: %s", strings.Join(failed, "; "))
		result.SuggestedFix = "Update the info section with version details, deprecation notices, and migration references"
	} else {
		result.Message = "API versioning strategy is properly documented"
	}

	return result
}

// checkDeprecation scans operations for deprecated flags and verifies descriptions.
func (p *P008Versioning) checkDeprecation(doc spec.APISpec, checks map[string]bool, messages map[string]string) {
	paths := doc.Paths()
	var deprecatedOps []string
	var missingNotice []string

	for path, item := range paths {
		if item == nil || item.Operations == nil {
			continue
		}
		for method, op := range item.Operations {
			if op == nil || !op.Deprecated {
				continue
			}
			label := fmt.Sprintf("%s %s", method, path)
			deprecatedOps = append(deprecatedOps, label)
			if !strings.Contains(strings.ToLower(op.Description), "deprecat") {
				missingNotice = append(missingNotice, label)
			}
		}
	}

	if len(deprecatedOps) == 0 {
		// No deprecated operations — pass (nothing to flag)
		checks["Deprecation notices are present"] = true
		return
	}

	if len(missingNotice) > 0 {
		checks["Deprecation notices are present"] = false
		messages["Deprecation notices are present"] = fmt.Sprintf(
			"Deprecated operations missing deprecation details in description: %s",
			strings.Join(missingNotice, ", "))
	} else {
		checks["Deprecation notices are present"] = true
	}
}
