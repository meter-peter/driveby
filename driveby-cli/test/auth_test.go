package test

import (
	"encoding/base64"
	"net/http"
	"testing"

	dbtesting "github.com/meter-peter/driveby/driveby-cli/internal/testing"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

func TestAddAuthHeaders_Bearer(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	auth := &types.AuthConfig{Token: "my-jwt-token"}
	if err := dbtesting.AddAuthHeaders(req, auth); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := req.Header.Get("Authorization")
	expected := "Bearer my-jwt-token"
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestAddAuthHeaders_APIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	auth := &types.AuthConfig{APIKey: "test-api-key"}
	if err := dbtesting.AddAuthHeaders(req, auth); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := req.Header.Get("X-API-Key")
	if got != "test-api-key" {
		t.Errorf("expected 'test-api-key', got %q", got)
	}
}

func TestAddAuthHeaders_Basic(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	auth := &types.AuthConfig{Username: "user", Password: "pass"}
	if err := dbtesting.AddAuthHeaders(req, auth); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := req.Header.Get("Authorization")
	expected := "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass"))
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestAddAuthHeaders_Nil(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	if err := dbtesting.AddAuthHeaders(req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.Header.Get("Authorization"); got != "" {
		t.Errorf("expected no Authorization header, got %q", got)
	}
	if got := req.Header.Get("X-API-Key"); got != "" {
		t.Errorf("expected no X-API-Key header, got %q", got)
	}
}

func TestAddAuthHeaders_MultipleError(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	auth := &types.AuthConfig{Token: "tok", APIKey: "key"}
	err := dbtesting.AddAuthHeaders(req, auth)
	if err == nil {
		t.Fatal("expected error when multiple auth methods specified")
	}
}
