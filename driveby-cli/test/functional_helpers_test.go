package test

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	dbtesting "github.com/meter-peter/driveby/driveby-cli/internal/testing"
)

func TestSubstitutePathParams_UUID(t *testing.T) {
	op := &spec.Operation{
		Parameters: []*spec.Parameter{
			{
				Name:    "product_id",
				In:      "path",
				Example: "abc-123-def",
			},
		},
	}
	result := dbtesting.SubstitutePathParams("/products/{product_id}", op)
	if result != "/products/abc-123-def" {
		t.Errorf("expected '/products/abc-123-def', got %q", result)
	}
}

func TestSubstitutePathParams_Integer(t *testing.T) {
	op := &spec.Operation{
		Parameters: []*spec.Parameter{
			{
				Name:   "id",
				In:     "path",
				Schema: &spec.Schema{Type: "integer"},
			},
		},
	}
	result := dbtesting.SubstitutePathParams("/items/{id}", op)
	if result != "/items/1" {
		t.Errorf("expected '/items/1', got %q", result)
	}
}

func TestSubstitutePathParams_NoParams(t *testing.T) {
	op := &spec.Operation{}
	result := dbtesting.SubstitutePathParams("/health", op)
	if result != "/health" {
		t.Errorf("expected '/health', got %q", result)
	}
}

func TestSubstitutePathParams_Multiple(t *testing.T) {
	op := &spec.Operation{
		Parameters: []*spec.Parameter{
			{
				Name:    "org_id",
				In:      "path",
				Example: "org-1",
			},
			{
				Name:   "user_id",
				In:     "path",
				Schema: &spec.Schema{Type: "integer"},
			},
		},
	}
	result := dbtesting.SubstitutePathParams("/orgs/{org_id}/users/{user_id}", op)
	if result != "/orgs/org-1/users/1" {
		t.Errorf("expected '/orgs/org-1/users/1', got %q", result)
	}
}

func TestBuildRequestBody_WithExample(t *testing.T) {
	example := map[string]interface{}{"name": "test", "value": float64(42)}
	op := &spec.Operation{
		RequestBody: &spec.RequestBody{
			Content: map[string]*spec.MediaType{
				"application/json": {
					Example: example,
				},
			},
		},
	}
	reader, err := dbtesting.BuildRequestBody(op)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reader == nil {
		t.Fatal("expected non-nil reader")
	}
	data, _ := io.ReadAll(reader)
	var got map[string]interface{}
	json.Unmarshal(data, &got)
	if got["name"] != "test" {
		t.Errorf("expected name 'test', got %v", got["name"])
	}
}

func TestBuildRequestBody_NilBody(t *testing.T) {
	op := &spec.Operation{} // GET — no request body
	reader, err := dbtesting.BuildRequestBody(op)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reader != nil {
		t.Error("expected nil reader for GET operation")
	}
}

func TestBuildRequestBody_FromSchema(t *testing.T) {
	op := &spec.Operation{
		RequestBody: &spec.RequestBody{
			Content: map[string]*spec.MediaType{
				"application/json": {
					Schema: &spec.Schema{
						Type: "object",
						Properties: map[string]*spec.Schema{
							"name":  {Type: "string"},
							"count": {Type: "integer"},
						},
						Required: []string{"name"},
					},
				},
			},
		},
	}
	reader, err := dbtesting.BuildRequestBody(op)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reader == nil {
		t.Fatal("expected non-nil reader")
	}
	data, _ := io.ReadAll(reader)
	var got map[string]interface{}
	json.Unmarshal(data, &got)
	if _, ok := got["name"]; !ok {
		t.Error("expected 'name' field in generated body")
	}
	if _, ok := got["count"]; !ok {
		t.Error("expected 'count' field in generated body")
	}
}

func TestGenerateExampleFromSchema_Object(t *testing.T) {
	schema := &spec.Schema{
		Type: "object",
		Properties: map[string]*spec.Schema{
			"id":    {Type: "string", Format: "uuid"},
			"name":  {Type: "string"},
			"count": {Type: "integer"},
			"active": {Type: "boolean"},
			"tags": {
				Type:  "array",
				Items: &spec.Schema{Type: "string"},
			},
		},
	}
	result := dbtesting.GenerateExampleFromSchema(schema)
	obj, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map, got %T", result)
	}
	if obj["id"] != "f7cfc49d-824b-4728-a4c4-45e5901e3d42" {
		t.Errorf("expected UUID for id, got %v", obj["id"])
	}
	if obj["name"] != "example" {
		t.Errorf("expected 'example' for name, got %v", obj["name"])
	}
	if obj["count"] != 1 {
		t.Errorf("expected 1 for count, got %v", obj["count"])
	}
	if obj["active"] != true {
		t.Errorf("expected true for active, got %v", obj["active"])
	}
	tags, ok := obj["tags"].([]interface{})
	if !ok || len(tags) == 0 {
		t.Error("expected non-empty array for tags")
	}
}

func TestValidateResponseBody_Correct(t *testing.T) {
	op := &spec.Operation{
		Responses: map[string]*spec.Response{
			"200": {
				Content: map[string]*spec.MediaType{
					"application/json": {
						Schema: &spec.Schema{
							Type:     "object",
							Required: []string{"id", "name"},
							Properties: map[string]*spec.Schema{
								"id":   {Type: "string"},
								"name": {Type: "string"},
							},
						},
					},
				},
			},
		},
	}
	body := []byte(`{"id": "123", "name": "test"}`)
	warnings := dbtesting.ValidateResponseBody(body, op, 200)
	if len(warnings) > 0 {
		t.Errorf("expected no warnings, got %v", warnings)
	}
}

func TestValidateResponseBody_WrongType(t *testing.T) {
	op := &spec.Operation{
		Responses: map[string]*spec.Response{
			"200": {
				Content: map[string]*spec.MediaType{
					"application/json": {
						Schema: &spec.Schema{
							Type: "object",
						},
					},
				},
			},
		},
	}
	body := []byte(`[1, 2, 3]`) // array, but schema says object
	warnings := dbtesting.ValidateResponseBody(body, op, 200)
	if len(warnings) == 0 {
		t.Error("expected warning for type mismatch")
	}
}

func TestValidateResponseBody_MissingRequired(t *testing.T) {
	op := &spec.Operation{
		Responses: map[string]*spec.Response{
			"200": {
				Content: map[string]*spec.MediaType{
					"application/json": {
						Schema: &spec.Schema{
							Type:     "object",
							Required: []string{"id", "name"},
							Properties: map[string]*spec.Schema{
								"id":   {Type: "string"},
								"name": {Type: "string"},
							},
						},
					},
				},
			},
		},
	}
	body := []byte(`{"id": "123"}`) // missing "name"
	warnings := dbtesting.ValidateResponseBody(body, op, 200)
	if len(warnings) == 0 {
		t.Error("expected warning for missing required field")
	}
	found := false
	for _, w := range warnings {
		if w == "Missing required field in response: name" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected 'Missing required field in response: name', got %v", warnings)
	}
}
