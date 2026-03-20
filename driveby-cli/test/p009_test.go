package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP009PassesWithExamples(t *testing.T) {
	l := loadTestDoc(t, "testdata/test-ready-api.json")
	checker := &principles.P009TestReadiness{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	if !result.Passed {
		t.Errorf("expected P009 to pass for spec with examples, got: %s", result.Message)
	}
}

func TestP009PassesWithTypedSchemasNoExamples(t *testing.T) {
	l := loadTestDoc(t, "testdata/test-ready-typed-no-examples-api.json")
	checker := &principles.P009TestReadiness{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	if !result.Passed {
		t.Errorf("expected P009 to pass for spec with typed schemas (no examples), got: %s", result.Message)
	}
}

func TestP009FailsWithNoSchemasNoExamples(t *testing.T) {
	l := loadTestDoc(t, "testdata/not-test-ready-api.json")
	checker := &principles.P009TestReadiness{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	if result.Passed {
		t.Error("expected P009 to fail for spec with no schemas and no examples")
	}
}

func TestP009FailsMissingResponseSchemas(t *testing.T) {
	// valid-openapi.json has a 201 response with no content/schema on POST /items
	l := loadTestDoc(t, "testdata/valid-openapi.json")
	checker := &principles.P009TestReadiness{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	if result.Passed {
		t.Error("expected P009 to fail for spec missing response schemas on 201")
	}
}

func TestP009PassesPerfectAPI(t *testing.T) {
	l := loadTestDoc(t, "../../apis/perfect-api/openapi.json")
	checker := &principles.P009TestReadiness{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	if !result.Passed {
		t.Errorf("expected P009 to pass for perfect-api, got: %s", result.Message)
	}
}
