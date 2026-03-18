package test

import (
	"testing"

	"github.com/meter-peter/driveby/driveby-cli/internal/loader"
	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
)

func TestLoadValidOpenAPI(t *testing.T) {
	l := loader.NewLoader()
	err := l.LoadFromFile("testdata/valid-openapi.json")
	if err != nil {
		t.Fatalf("expected no error loading valid openapi, got: %v", err)
	}
	doc := l.GetDocument()
	if doc == nil {
		t.Fatal("expected document to be non-nil")
	}
	if doc.Type() != spec.SpecTypeOpenAPI3 {
		t.Errorf("expected type openapi3, got %s", doc.Type())
	}
	if doc.RawVersion() != "3.0.3" {
		t.Errorf("expected version 3.0.3, got %s", doc.RawVersion())
	}
	info := doc.Info()
	if info == nil {
		t.Fatal("expected info to be non-nil")
	}
	if info.Title != "Test API" {
		t.Errorf("expected title 'Test API', got '%s'", info.Title)
	}
	paths := doc.Paths()
	if len(paths) != 2 {
		t.Errorf("expected 2 paths, got %d", len(paths))
	}
}

func TestLoadSwagger2(t *testing.T) {
	l := loader.NewLoader()
	err := l.LoadFromFile("testdata/swagger2-sample.json")
	if err != nil {
		t.Fatalf("expected no error loading swagger 2.0, got: %v", err)
	}
	doc := l.GetDocument()
	if doc == nil {
		t.Fatal("expected document to be non-nil")
	}
	if doc.Type() != spec.SpecTypeSwagger2 {
		t.Errorf("expected type swagger2, got %s", doc.Type())
	}
	if doc.RawVersion() != "2.0" {
		t.Errorf("expected version 2.0, got %s", doc.RawVersion())
	}
}

func TestLoadNullableOpenAPI31(t *testing.T) {
	l := loader.NewLoader()
	err := l.LoadFromFile("testdata/nullable-openapi31.json")
	if err != nil {
		t.Fatalf("expected no error loading nullable OpenAPI 3.1.0 spec, got: %v", err)
	}
	doc := l.GetDocument()
	if doc == nil {
		t.Fatal("expected document to be non-nil")
	}
	// Should load despite anyOf with null type
	paths := doc.Paths()
	if len(paths) == 0 {
		t.Fatal("expected at least one path")
	}
	if _, ok := paths["/items"]; !ok {
		t.Error("expected /items path")
	}

	// Check that schemas loaded correctly
	comps := doc.Components()
	if comps == nil {
		t.Fatal("expected components to be non-nil")
	}
	item, ok := comps.Schemas["Item"]
	if !ok {
		t.Fatal("expected Item schema in components")
	}
	desc, ok := item.Properties["description"]
	if !ok {
		t.Fatal("expected description property in Item schema")
	}
	if desc.Type != "string" {
		t.Errorf("expected description type 'string', got %q", desc.Type)
	}
	if !desc.Nullable {
		t.Error("expected description to be nullable after preprocessing")
	}
}

func TestLoadInvalidOpenAPI(t *testing.T) {
	l := loader.NewLoader()
	err := l.LoadFromFile("testdata/invalid-openapi.json")
	if err != nil {
		t.Fatalf("expected no error loading (invalid is still parseable), got: %v", err)
	}
	doc := l.GetDocument()
	if doc == nil {
		t.Fatal("expected document to be non-nil")
	}
	// info should be mostly empty
	info := doc.Info()
	if info != nil && info.Title != "" {
		t.Errorf("expected empty title, got '%s'", info.Title)
	}
}
