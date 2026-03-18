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
