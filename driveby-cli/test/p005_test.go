package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP005SecurityWithSchemes(t *testing.T) {
	l := loadTestDoc(t, "testdata/security-api.json")
	checker := &principles.P005Security{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if !result.Passed {
		t.Errorf("expected P005 to pass for security-api.json, got: %s", result.Message)
	}
}

func TestP005SecurityNoSchemes(t *testing.T) {
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P005Security{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if result.Passed {
		t.Error("expected P005 to fail when no security schemes defined")
	}
}

func TestP005SecurityMissingGlobalSecurity(t *testing.T) {
	l := loadTestDoc(t, "testdata/missing-global-security-api.json")
	checker := &principles.P005Security{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	// Has security schemes but no global security, and GET /items has no operation-level security
	if result.Passed {
		t.Error("expected P005 to fail when global security is missing and some operations lack security")
	}

	// Verify the specific check that failed
	details, ok := result.Details.(map[string]interface{})
	if !ok {
		t.Fatal("expected details to be map[string]interface{}")
	}
	checks, ok := details["checks"].(map[string]bool)
	if !ok {
		t.Fatal("expected checks to be map[string]bool")
	}
	if checks["Global security requirements are set"] {
		t.Error("expected 'Global security requirements are set' to be false")
	}
	if checks["Security schemes are defined"] != true {
		t.Error("expected 'Security schemes are defined' to be true")
	}
}

func TestP005SecurityMinimalMode(t *testing.T) {
	l := loadTestDoc(t, "testdata/missing-global-security-api.json")
	checker := &principles.P005Security{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	// Minimal mode only checks: schemes defined (pass) + global security (fail)
	if result.Passed {
		t.Error("expected P005 minimal to fail when global security is missing")
	}

	details, ok := result.Details.(map[string]interface{})
	if !ok {
		t.Fatal("expected details to be map")
	}
	checks, ok := details["checks"].(map[string]bool)
	if !ok {
		t.Fatal("expected checks to be map[string]bool")
	}
	// Should NOT have strict-only checks
	if _, exists := checks["OAuth2 scopes are documented"]; exists {
		t.Error("minimal mode should not evaluate OAuth2 scopes check")
	}
}

func TestP005SecurityConsistentRefs(t *testing.T) {
	l := loadTestDoc(t, "testdata/security-api.json")
	checker := &principles.P005Security{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	// security-api.json references bearerAuth and apiKeyAuth, both defined in securitySchemes
	if !result.Passed {
		t.Errorf("expected P005 to pass for consistent security refs, got: %s", result.Message)
	}
}
