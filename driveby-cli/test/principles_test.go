package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/loader"
	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func loadTestDoc(t *testing.T, path string) *loader.Loader {
	t.Helper()
	l := loader.NewLoader()
	if err := l.LoadFromFile(path); err != nil {
		t.Fatalf("failed to load %s: %v", path, err)
	}
	return l
}

func TestP001ComplianceValidSpec(t *testing.T) {
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P001Compliance{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	if !result.Passed {
		t.Errorf("expected P001 to pass for valid spec, got: %s", result.Message)
	}
}

func TestP001ComplianceInvalidSpec(t *testing.T) {
	l := loadTestDoc(t, "testdata/invalid-openapi.json")
	checker := &principles.P001Compliance{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	if result.Passed {
		t.Error("expected P001 to fail for invalid spec (missing info fields, empty paths)")
	}
}

func TestRegistryForMode(t *testing.T) {
	registry := principles.NewRegistry()

	minimal := registry.ForMode(types.ValidationModeMinimal)
	if len(minimal) != 1 {
		t.Errorf("expected 1 checker for minimal mode, got %d", len(minimal))
	}

	strict := registry.ForMode(types.ValidationModeStrict)
	if len(strict) != 6 {
		t.Errorf("expected 6 checkers for strict mode, got %d", len(strict))
	}

	testReady := registry.ForMode(types.ValidationModeTestReady)
	if len(testReady) != 5 {
		t.Errorf("expected 5 checkers for test-ready mode, got %d", len(testReady))
	}

	testOnly := registry.ForMode(types.ValidationModeTestOnly)
	if len(testOnly) != 0 {
		t.Errorf("expected 0 checkers for test-only mode, got %d", len(testOnly))
	}
}

// P005 and P008 tests moved to p005_test.go and p008_test.go
