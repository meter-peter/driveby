package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP002DocumentationWellDocumented(t *testing.T) {
	l := loadTestDoc(t, "testdata/well-documented-api.json")
	checker := &principles.P002Documentation{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if !result.Passed {
		t.Errorf("expected P002 to pass for well-documented API, got: %s", result.Message)
	}
}

func TestP002DocumentationMissingDocs(t *testing.T) {
	// valid-openapi.json lacks examples and schema descriptions
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P002Documentation{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if result.Passed {
		t.Error("expected P002 to fail for API missing examples/schema descriptions")
	}
}

func TestP002DocumentationMinimalSpec(t *testing.T) {
	l := loadTestDoc(t, "testdata/invalid-openapi.json")
	checker := &principles.P002Documentation{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if result.Passed {
		t.Error("expected P002 to fail for API with empty info section")
	}
}
