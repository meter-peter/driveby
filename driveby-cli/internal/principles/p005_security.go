package principles

import (
	"context"
	"fmt"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// P005Security validates that the API has comprehensive security definitions,
// authentication mechanisms, and consistent security requirements.
type P005Security struct{}

func (p *P005Security) ID() string { return "P005" }

func (p *P005Security) Check(_ context.Context, doc spec.APISpec, mode types.ValidationMode) types.PrincipleResult {
	result := types.PrincipleResult{
		Principle: types.CorePrinciples[4],
		Passed:    true,
		Details:   make(map[string]interface{}),
	}

	checks := make(map[string]bool)
	messages := make(map[string]string)

	comps := doc.Components()
	schemes := map[string]*spec.SecurityScheme{}
	if comps != nil && comps.SecuritySchemes != nil {
		schemes = comps.SecuritySchemes
	}

	// Check 1: Security schemes are defined
	if len(schemes) == 0 {
		checks["Security schemes are defined"] = false
		messages["Security schemes are defined"] = "No security schemes defined in components.securitySchemes"
	} else {
		checks["Security schemes are defined"] = true
	}

	// Check 2: Global security requirements are set
	globalSec := doc.Security()
	if globalSec == nil || len(globalSec) == 0 {
		checks["Global security requirements are set"] = false
		messages["Global security requirements are set"] = "No top-level security requirements defined"
	} else {
		checks["Global security requirements are set"] = true
	}

	// In minimal mode, only checks 1 and 2 are evaluated.
	if mode != types.ValidationModeMinimal {
		// Check 3: Operation-level security is defined
		var missingAuth []string
		for path, pathItem := range doc.Paths() {
			if pathItem == nil || pathItem.Operations == nil {
				continue
			}
			for method, op := range pathItem.Operations {
				if op == nil {
					continue
				}
				if op.Security == nil && len(globalSec) == 0 {
					missingAuth = append(missingAuth, fmt.Sprintf("%s %s", method, path))
				}
			}
		}
		if len(missingAuth) > 0 {
			checks["Operation-level security is defined"] = false
			messages["Operation-level security is defined"] = fmt.Sprintf("Endpoints without security: %s", strings.Join(missingAuth, ", "))
		} else {
			checks["Operation-level security is defined"] = true
		}

		// Check 4: OAuth2 scopes are documented
		hasOAuth := false
		oauthScopesOK := true
		for name, sc := range schemes {
			if sc.Type == "oauth2" || sc.Type == "openIdConnect" {
				hasOAuth = true
				// For OAuth2/OIDC schemes, scopes should appear in at least one security requirement.
				if !schemeReferencedWithScopes(name, globalSec, doc) {
					oauthScopesOK = false
					messages["OAuth2 scopes are documented"] = fmt.Sprintf("OAuth2/OIDC scheme %q is not referenced with scopes in any security requirement", name)
				}
			}
		}
		if !hasOAuth {
			// No OAuth2 schemes means this check is not applicable; pass by default.
			checks["OAuth2 scopes are documented"] = true
		} else {
			checks["OAuth2 scopes are documented"] = oauthScopesOK
		}

		// Check 5: API keys are properly described
		apiKeyIssues := []string{}
		for name, sc := range schemes {
			if sc.Type == "apiKey" && strings.TrimSpace(sc.Description) == "" {
				apiKeyIssues = append(apiKeyIssues, name)
			}
		}
		if len(apiKeyIssues) > 0 {
			checks["API keys are properly described"] = false
			messages["API keys are properly described"] = fmt.Sprintf("apiKey schemes missing description: %s", strings.Join(apiKeyIssues, ", "))
		} else {
			checks["API keys are properly described"] = true
		}

		// Check 6: Authentication headers are specified
		headerIssues := []string{}
		for name, sc := range schemes {
			if sc.Type == "apiKey" {
				if strings.TrimSpace(sc.Name) == "" || strings.TrimSpace(sc.In) == "" {
					headerIssues = append(headerIssues, name)
				}
			}
		}
		if len(headerIssues) > 0 {
			checks["Authentication headers are specified"] = false
			messages["Authentication headers are specified"] = fmt.Sprintf("apiKey schemes missing Name/In fields: %s", strings.Join(headerIssues, ", "))
		} else {
			checks["Authentication headers are specified"] = true
		}

		// Check 7: Security requirements are consistent
		inconsistent := collectInconsistentRefs(schemes, globalSec, doc)
		if len(inconsistent) > 0 {
			checks["Security requirements are consistent"] = false
			messages["Security requirements are consistent"] = fmt.Sprintf("Undefined security schemes referenced: %s", strings.Join(inconsistent, ", "))
		} else {
			checks["Security requirements are consistent"] = true
		}
	}

	// Aggregate results
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
		result.Message = fmt.Sprintf("Security validation failed: %s", strings.Join(failed, "; "))
		result.SuggestedFix = "Review security schemes, global security requirements, and per-operation security"
	} else {
		result.Message = "All security requirements are properly defined and consistent"
	}

	return result
}

// schemeReferencedWithScopes returns true if the given scheme name appears in
// any security requirement (global or operation-level) with at least one scope.
func schemeReferencedWithScopes(name string, globalSec []spec.SecurityRequirement, doc spec.APISpec) bool {
	for _, req := range globalSec {
		if scopes, ok := req[name]; ok && len(scopes) > 0 {
			return true
		}
	}
	for _, pi := range doc.Paths() {
		if pi == nil || pi.Operations == nil {
			continue
		}
		for _, op := range pi.Operations {
			if op == nil {
				continue
			}
			for _, req := range op.Security {
				if scopes, ok := req[name]; ok && len(scopes) > 0 {
					return true
				}
			}
		}
	}
	return false
}

// collectInconsistentRefs finds scheme names used in security requirements that
// are not defined in components.securitySchemes.
func collectInconsistentRefs(schemes map[string]*spec.SecurityScheme, globalSec []spec.SecurityRequirement, doc spec.APISpec) []string {
	seen := map[string]bool{}
	check := func(reqs []spec.SecurityRequirement) {
		for _, req := range reqs {
			for name := range req {
				if _, defined := schemes[name]; !defined {
					seen[name] = true
				}
			}
		}
	}
	check(globalSec)
	for _, pi := range doc.Paths() {
		if pi == nil || pi.Operations == nil {
			continue
		}
		for _, op := range pi.Operations {
			if op == nil {
				continue
			}
			check(op.Security)
		}
	}
	var out []string
	for name := range seen {
		out = append(out, name)
	}
	return out
}
