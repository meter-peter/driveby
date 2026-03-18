package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP004SchemaStrictPassesWithConstraints(t *testing.T) {
	l := loadTestDoc(t, "testdata/strict-schema-api.json")
	checker := &principles.P004Schema{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if !result.Passed {
		t.Errorf("expected P004 to pass for strict-schema API, got: %s", result.Message)
	}
}

func TestP004SchemaMinimalPassesBasicSpec(t *testing.T) {
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P004Schema{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	if !result.Passed {
		t.Errorf("expected P004 minimal to pass for valid spec, got: %s", result.Message)
	}
}

func TestP004SchemaStrictFailsWithoutConstraints(t *testing.T) {
	// valid-openapi.json has schemas but no constraints (no minLength, no min/max)
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P004Schema{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if result.Passed {
		t.Error("expected P004 strict to fail for API without schema constraints")
	}
}

func TestP004SchemaAllOfComposition(t *testing.T) {
	l := loadTestDoc(t, "testdata/allof-composition-api.json")
	checker := &principles.P004Schema{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	if !result.Passed {
		t.Errorf("expected P004 minimal to pass for allOf composition API, got: %s", result.Message)
	}
}
