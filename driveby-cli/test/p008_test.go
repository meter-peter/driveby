package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP008VersioningValidSemVer(t *testing.T) {
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P008Versioning{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	// valid-openapi.json has version "1.0.0" which is valid semver
	if !result.Passed {
		t.Errorf("expected P008 to pass for valid semver, got: %s", result.Message)
	}
}

func TestP008VersioningInvalidFormat(t *testing.T) {
	l := loadTestDoc(t, "testdata/invalid-version-api.json")
	checker := &principles.P008Versioning{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	// invalid-version-api.json has version "v1" which is NOT valid semver
	if result.Passed {
		t.Error("expected P008 to fail for non-semver version 'v1'")
	}
}

func TestP008VersioningDeprecationNotices(t *testing.T) {
	l := loadTestDoc(t, "testdata/deprecated-endpoints-api.json")
	checker := &principles.P008Versioning{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	// deprecated-endpoints-api.json has deprecated ops with proper descriptions
	// and version info mentions versioning, breaking, compatibility, migration
	if !result.Passed {
		t.Errorf("expected P008 strict to pass for well-documented versioned API, got: %s", result.Message)
	}
}

func TestP008VersioningMinimalMode(t *testing.T) {
	l := loadTestDoc(t, "testdata/deprecated-endpoints-api.json")
	checker := &principles.P008Versioning{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	// Minimal mode only checks version specified + semver format
	if !result.Passed {
		t.Errorf("expected P008 minimal to pass, got: %s", result.Message)
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
	if _, exists := checks["Breaking changes are documented"]; exists {
		t.Error("minimal mode should not evaluate breaking changes check")
	}
}

func TestP008VersioningStrictMissingDocs(t *testing.T) {
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P008Versioning{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	// valid-openapi.json has version 1.0.0 but description "A test API for validation"
	// — no mention of versioning strategy, breaking changes, compatibility, or migration
	if result.Passed {
		t.Error("expected P008 strict to fail for API without versioning documentation in description")
	}
}
