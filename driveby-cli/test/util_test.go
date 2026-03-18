package test

import (
	"encoding/json"
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/util"
)

func parseJSON(t *testing.T, s string) map[string]interface{} {
	t.Helper()
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}
	return m
}

func TestPreprocessNullableAnyOf_StringNull(t *testing.T) {
	m := parseJSON(t, `{
		"properties": {
			"name": {
				"anyOf": [{"type": "string"}, {"type": "null"}],
				"title": "Name"
			}
		}
	}`)
	util.PreprocessNullableAnyOf(m)

	props := m["properties"].(map[string]interface{})
	name := props["name"].(map[string]interface{})

	if _, hasAnyOf := name["anyOf"]; hasAnyOf {
		t.Error("anyOf should have been removed")
	}
	if name["type"] != "string" {
		t.Errorf("expected type 'string', got %v", name["type"])
	}
	if name["nullable"] != true {
		t.Error("expected nullable: true")
	}
	if name["title"] != "Name" {
		t.Error("title should be preserved")
	}
}

func TestPreprocessNullableAnyOf_NumberNull(t *testing.T) {
	m := parseJSON(t, `{
		"properties": {
			"count": {
				"anyOf": [{"type": "integer", "minimum": 0}, {"type": "null"}],
				"title": "Count"
			}
		}
	}`)
	util.PreprocessNullableAnyOf(m)

	props := m["properties"].(map[string]interface{})
	count := props["count"].(map[string]interface{})

	if _, hasAnyOf := count["anyOf"]; hasAnyOf {
		t.Error("anyOf should have been removed")
	}
	if count["type"] != "integer" {
		t.Errorf("expected type 'integer', got %v", count["type"])
	}
	if count["nullable"] != true {
		t.Error("expected nullable: true")
	}
	if count["minimum"] != float64(0) {
		t.Errorf("expected minimum 0, got %v", count["minimum"])
	}
}

func TestPreprocessNullableAnyOf_ThreeMembers(t *testing.T) {
	m := parseJSON(t, `{
		"field": {
			"anyOf": [{"type": "string"}, {"type": "integer"}, {"type": "null"}]
		}
	}`)
	util.PreprocessNullableAnyOf(m)

	field := m["field"].(map[string]interface{})
	if _, hasAnyOf := field["anyOf"]; !hasAnyOf {
		t.Error("3-member anyOf should NOT be converted")
	}
}

func TestPreprocessNullableAnyOf_Nested(t *testing.T) {
	m := parseJSON(t, `{
		"components": {
			"schemas": {
				"Item": {
					"properties": {
						"desc": {
							"anyOf": [{"type": "string"}, {"type": "null"}]
						}
					}
				}
			}
		}
	}`)
	util.PreprocessNullableAnyOf(m)

	desc := m["components"].(map[string]interface{})["schemas"].(map[string]interface{})["Item"].(map[string]interface{})["properties"].(map[string]interface{})["desc"].(map[string]interface{})
	if desc["type"] != "string" {
		t.Errorf("expected nested type 'string', got %v", desc["type"])
	}
	if desc["nullable"] != true {
		t.Error("expected nested nullable: true")
	}
}

func TestPreprocessNullableAnyOf_NoAnyOf(t *testing.T) {
	m := parseJSON(t, `{
		"properties": {
			"name": {"type": "string"},
			"count": {"type": "integer"}
		}
	}`)
	original := `{"properties":{"name":{"type":"string"},"count":{"type":"integer"}}}`
	util.PreprocessNullableAnyOf(m)

	result, _ := json.Marshal(m)
	// Just verify it doesn't crash and types are preserved
	var check map[string]interface{}
	json.Unmarshal([]byte(original), &check)

	props := m["properties"].(map[string]interface{})
	if props["name"].(map[string]interface{})["type"] != "string" {
		t.Error("no-op case: name type should still be string")
	}
	if props["count"].(map[string]interface{})["type"] != "integer" {
		t.Error("no-op case: count type should still be integer")
	}
	_ = result
}
