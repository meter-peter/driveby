package validation_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/meter-peter/driveby/internal/validation"
)

func TestAPIValidator(t *testing.T) {
	// Create a dummy OpenAPI spec file for testing
	specContent := `
openapi: 3.0.0
info:
  title: Test API
  version: 1.0.0
paths:
  /test:
    get:
      summary: A test endpoint
      responses:
        '200':
          description: Successful response
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                  status:
                    type: string
      parameters:
        - in: query
          name: queryParam
          schema:
            type: string
            example: testValue
    post:
      summary: A test post endpoint
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                value:
                  type: integer
              example: {"name": "test", "value": 123}
      responses:
        '201':
          description: Created
  /error:
    get:
      summary: An endpoint that returns an error
      responses:
        '500':
          description: Internal Server Error
`

	dummySpecFile := "test_openapi.yaml"
	if err := os.WriteFile(dummySpecFile, []byte(specContent), 0644); err != nil {
		t.Fatalf("Failed to create dummy spec file: %v", err)
	}
	defer os.Remove(dummySpecFile)

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/test":
			if r.Method == http.MethodGet {
				queryParam := r.URL.Query().Get("queryParam")
				if queryParam != "testValue" {
					http.Error(w, "Invalid query parameter", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message": "success", "status": "ok"}`))
			} else if r.Method == http.MethodPost {
				// In a real scenario, you would read and validate the request body
				w.WriteHeader(http.StatusCreated)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		case "/error":
			if r.Method == http.MethodGet {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Validator configuration
	config := validation.ValidatorConfig{
		BaseURL:     server.URL,
		SpecPath:    dummySpecFile,
		LogPath:     "./validation.log", // Specify a log file
		Environment: "test",
		Version:     "1.0",
		AutoFix:     false,
		Timeout:     5 * time.Second,
		PerformanceTarget: validation.PerformanceTargetConfig{
			MaxLatencyP95:  100 * time.Millisecond,
			MinSuccessRate: 99.0, // 99% success rate
		},
	}

	// Create a new APIValidator instance
	validator, err := validation.NewAPIValidator(config)
	if err != nil {
		t.Fatalf("Failed to create APIValidator: %v", err)
	}

	// Run the validation suite
	report, err := validator.Validate(context.Background())
	if err != nil {
		t.Fatalf("Validation failed: %v", err)
	}

	// Basic assertions on the report (you can add more detailed checks)
	if report == nil {
		t.Error("Validation report is nil")
	}

	t.Logf("Validation Report: %+v\n", report)
	// You can further inspect the report.Principles to check individual results

	// Example of checking a specific principle result (e.g., P006 Functional Testing)
	foundP006 := false
	for _, p := range report.Principles {
		if p.Principle.ID == "P006" {
			foundP006 = true
			if !p.Passed {
				t.Errorf("Principle P006 (Endpoint Functional Testing) failed: %s", p.Message)
			}
			// You can cast p.Details to []validation.EndpointValidation and inspect individual endpoint results
			break
		}
	}
	if !foundP006 {
		t.Error("Principle P006 not found in report")
	}

	// Example of checking a specific principle result (e.g., P007 Performance Compliance)
	foundP007 := false
	for _, p := range report.Principles {
		if p.Principle.ID == "P007" {
			foundP007 = true
			if !p.Passed {
				t.Errorf("Principle P007 (API Performance Compliance) failed: %s", p.Message)
			}
			// You can cast p.Details to *validation.PerformanceMetrics and inspect the metrics
			break
		}
	}
	if !foundP007 {
		t.Error("Principle P007 not found in report")
	}

	// You can also check the generated report files in the specified log path
}
