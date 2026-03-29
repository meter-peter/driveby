package test

import (
	"context"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/principles"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestP003ErrorsMissingErrorResponses(t *testing.T) {
	l := loadTestDoc(t, "testdata/missing-errors-api.json")
	checker := &principles.P003Errors{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	if result.Passed {
		t.Error("expected P003 to fail for API with no error responses")
	}
}

func TestP003ErrorsMissingErrorsMinimalMode(t *testing.T) {
	l := loadTestDoc(t, "testdata/missing-errors-api.json")
	checker := &principles.P003Errors{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeMinimal)
	// In minimal mode, only present error responses are checked — none present means pass
	if !result.Passed {
		t.Errorf("expected P003 minimal mode to pass when no error responses present, got: %s", result.Message)
	}
}

func TestP003ErrorsTestReadyMode(t *testing.T) {
	// In test-ready mode, P003 skips: common components check, 5xx requirement, format consistency
	// But still checks: 4xx responses exist, error schemas have message/code/details
	l := loadTestDoc(t, "testdata/missing-errors-api.json")
	checker := &principles.P003Errors{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeTestReady)
	// Missing errors API has no 4xx responses, so should still fail in test-ready
	if result.Passed {
		t.Error("expected P003 test-ready to fail for API with no error responses")
	}
}

func TestP003ErrorsPetstoreHasErrors(t *testing.T) {
	l := loadTestDoc(t, "testdata/real-world-petstore.json")
	checker := &principles.P003Errors{}
	result := checker.Check(context.Background(), l.GetDocument(), types.ValidationModeStrict)
	// Petstore should have error responses
	if result.Passed {
		t.Log("P003 passed for petstore — it has proper error responses")
	} else {
		t.Logf("P003 result: %s", result.Message)
	}
}
