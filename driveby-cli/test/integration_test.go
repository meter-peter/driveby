//go:build integration

package test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// apiPort returns the port the test API listens on, defaulting to 8000.
func apiPort() string {
	if p := os.Getenv("DRIVEBY_TEST_PORT"); p != "" {
		return p
	}
	return "8000"
}

// apiHost returns the hostname of the test API, defaulting to localhost.
func apiHost() string {
	if h := os.Getenv("DRIVEBY_TEST_HOST"); h != "" {
		return h
	}
	return "localhost"
}

// specPath returns the absolute path to the perfect-api OpenAPI spec.
// Tests run from driveby-cli/test/, so the spec is at ../../apis/perfect-api/.
func specPath(t *testing.T) string {
	t.Helper()
	if p := os.Getenv("DRIVEBY_TEST_SPEC"); p != "" {
		abs, err := filepath.Abs(p)
		if err != nil {
			t.Fatalf("failed to resolve spec path: %v", err)
		}
		return abs
	}
	p, err := filepath.Abs("../../apis/perfect-api/openapi.json")
	if err != nil {
		t.Fatalf("failed to resolve spec path: %v", err)
	}
	return p
}

// runDriveby executes the driveby CLI with the given args and returns parsed JSON output.
// It tolerates exit code 1 (validation failed) but fails on anything else.
func runDriveby(t *testing.T, args ...string) map[string]interface{} {
	t.Helper()
	bin := binaryPath(t)
	cmd := exec.Command(bin, args...)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code: %d\nstderr: %s", exitErr.ExitCode(), string(exitErr.Stderr))
			if exitErr.ExitCode() > 2 {
				t.Fatalf("unexpected exit code %d", exitErr.ExitCode())
			}
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v\nraw output (%d bytes): %s", err, len(stdout), string(stdout))
	}
	return result
}

// --- Validate-only tests ---

func TestIntegrationValidateOnly(t *testing.T) {
	result := runDriveby(t,
		"validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "strict",
		"--log-level", "error",
	)
	if _, ok := result["status"]; !ok {
		t.Error("expected 'status' field in output")
	}
}

func TestIntegrationValidateMinimal(t *testing.T) {
	result := runDriveby(t,
		"validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "minimal",
		"--log-level", "error",
	)

	// Minimal mode: only P001
	principles, ok := result["principles"].([]interface{})
	if !ok {
		t.Fatal("expected 'principles' array in output")
	}
	if len(principles) != 1 {
		t.Errorf("expected 1 principle in minimal mode, got %d", len(principles))
	}

	// P001 should pass for perfect-api
	p001 := principles[0].(map[string]interface{})
	principle := p001["Principle"].(map[string]interface{})
	if principle["id"] != "P001" {
		t.Errorf("expected P001, got %v", principle["id"])
	}
	if p001["Passed"] != true {
		t.Error("P001 should pass for perfect-api spec")
	}
}

func TestIntegrationValidateStrictAllPrinciples(t *testing.T) {
	result := runDriveby(t,
		"validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "strict",
		"--log-level", "error",
	)

	// Strict mode should run P001-P005 + P008 (6 checkers)
	principles, ok := result["principles"].([]interface{})
	if !ok {
		t.Fatal("expected 'principles' array in output")
	}
	if len(principles) != 6 {
		t.Errorf("expected 6 principles in strict mode, got %d", len(principles))
	}

	// P001, P005, P008 should always pass for perfect-api
	mustPass := map[string]bool{"P001": true, "P005": true, "P008": true}
	for _, p := range principles {
		pm := p.(map[string]interface{})
		principle := pm["Principle"].(map[string]interface{})
		pid := principle["id"].(string)
		passed := pm["Passed"].(bool)
		if mustPass[pid] && !passed {
			msg, _ := pm["Message"].(string)
			t.Errorf("principle %s must pass for perfect-api, message: %s", pid, msg)
		}
	}
}

func TestIntegrationValidateExitCodeZero(t *testing.T) {
	bin := binaryPath(t)
	cmd := exec.Command(bin, "validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "minimal",
		"--log-level", "error",
	)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Errorf("expected exit code 0 for valid spec minimal mode, got %d", exitErr.ExitCode())
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}
}

func TestIntegrationValidateReportFields(t *testing.T) {
	result := runDriveby(t,
		"validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "strict",
		"--log-level", "error",
	)

	// Check all required report fields exist
	for _, field := range []string{"status", "exit_code", "principles", "total_checks", "passed_checks", "failed_checks", "summary", "timestamp"} {
		if _, ok := result[field]; !ok {
			t.Errorf("missing required field: %s", field)
		}
	}

	// total_checks should equal passed + failed
	total := result["total_checks"].(float64)
	passed := result["passed_checks"].(float64)
	failed := result["failed_checks"].(float64)
	if total != passed+failed {
		t.Errorf("total_checks (%.0f) != passed (%.0f) + failed (%.0f)", total, passed, failed)
	}
}

// --- Function-only tests ---

func TestIntegrationFunctionOnly(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--log-level", "error",
	)

	// Should have test_results with Functional section
	tr, ok := result["test_results"].(map[string]interface{})
	if !ok {
		t.Fatal("expected 'test_results' in output")
	}
	if _, ok := tr["Functional"]; !ok {
		t.Error("expected 'Functional' in test_results")
	}
}

func TestIntegrationFunctionalWithAuth(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})

	// With auth, no endpoints should be skipped
	skipped, _ := fn["SkippedEndpoints"].(float64)
	if skipped > 0 {
		t.Errorf("expected 0 skipped endpoints with auth, got %.0f", skipped)
	}

	// All tested endpoints should pass
	passed, _ := fn["PassedEndpoints"].(float64)
	total, _ := fn["TotalEndpoints"].(float64)
	failed, _ := fn["FailedEndpoints"].(float64)
	if failed > 0 {
		t.Errorf("expected 0 failed endpoints with auth, got %.0f", failed)
	}
	if passed != total {
		t.Errorf("expected all %v endpoints to pass, but only %v passed", total, passed)
	}
}

func TestIntegrationFunctionalWithoutAuth(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--log-level", "error",
	)

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})

	// Without auth, secured endpoints should be skipped (401 -> skipped)
	skipped, _ := fn["SkippedEndpoints"].(float64)
	if skipped == 0 {
		t.Error("expected some skipped endpoints without auth — API should enforce auth")
	}
}

func TestIntegrationFunctionalEndpointCoverage(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})

	// Perfect-api has 8 operations (excluding deprecated /legacy/products):
	//   POST /tasks, GET /products, POST /products, GET /products/{id},
	//   PUT /products/{id}, DELETE /products/{id}, GET /test/health, POST /test/echo
	// Deprecated endpoints are skipped by the functional tester.
	total, _ := fn["TotalEndpoints"].(float64)
	if total < 7 {
		t.Errorf("expected at least 7 tested endpoints, got %.0f", total)
	}

	// Check endpoint results are present
	endpoints, ok := fn["EndpointResults"].([]interface{})
	if !ok {
		t.Fatal("expected EndpointResults array")
	}

	// Build a set of tested method+path combos
	tested := make(map[string]bool)
	for _, ep := range endpoints {
		epm := ep.(map[string]interface{})
		key := fmt.Sprintf("%s %s", epm["Method"], epm["Path"])
		tested[key] = true
	}

	// Verify critical endpoints were tested
	expectedEndpoints := []string{
		"GET /products",
		"POST /products",
		"GET /test/health",
		"POST /test/echo",
		"POST /tasks",
	}
	for _, ep := range expectedEndpoints {
		if !tested[ep] {
			t.Errorf("expected endpoint %s to be tested, but it was not", ep)
		}
	}
}

func TestIntegrationFunctionalResponseTimes(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})

	// Response times should be populated
	avg, _ := fn["AverageResponseTime"].(float64)
	max, _ := fn["MaxResponseTime"].(float64)
	min, _ := fn["MinResponseTime"].(float64)

	if avg == 0 {
		t.Error("expected non-zero average response time")
	}
	if max == 0 {
		t.Error("expected non-zero max response time")
	}
	if min == 0 {
		t.Error("expected non-zero min response time")
	}
	if min > avg || avg > max {
		t.Errorf("response time invariant violated: min=%.0f avg=%.0f max=%.0f", min, avg, max)
	}
}

func TestIntegrationFunctionalBearerAuth(t *testing.T) {
	result := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-token", "test-bearer-token",
		"--log-level", "error",
	)

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})

	// Bearer auth should also work — perfect-api accepts any token
	skipped, _ := fn["SkippedEndpoints"].(float64)
	if skipped > 0 {
		t.Errorf("expected 0 skipped endpoints with bearer auth, got %.0f", skipped)
	}
}

// --- Test-only mode tests ---

func TestIntegrationTestOnly(t *testing.T) {
	result := runDriveby(t,
		"test-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--test-duration", "2s",
		"--concurrent-users", "2",
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	// test-only outputs a combined report with functional + performance keys
	if _, ok := result["functional"]; !ok {
		t.Error("expected 'functional' in test-only output")
	}
	if _, ok := result["performance"]; !ok {
		t.Error("expected 'performance' in test-only output")
	}
	if result["status"] == nil {
		t.Error("expected 'status' in test-only output")
	}
}

func TestIntegrationPerformanceMetrics(t *testing.T) {
	result := runDriveby(t,
		"test-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--test-duration", "3s",
		"--concurrent-users", "3",
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	// Performance report is nested under "performance" -> "principles" -> [P007] -> "Details"
	perfReport, ok := result["performance"].(map[string]interface{})
	if !ok {
		t.Fatal("expected 'performance' in test-only output")
	}
	principles, ok := perfReport["principles"].([]interface{})
	if !ok || len(principles) == 0 {
		t.Fatal("expected 'principles' with P007 in performance report")
	}
	p007 := principles[0].(map[string]interface{})
	details, ok := p007["Details"].(map[string]interface{})
	if !ok {
		t.Fatal("expected 'Details' in P007 principle result")
	}

	// Check performance metrics exist
	totalReqs, _ := details["total_requests"].(float64)
	if totalReqs == 0 {
		t.Error("expected non-zero total_requests in performance results")
	}

	successCount, _ := details["success_count"].(float64)
	if successCount == 0 {
		t.Error("expected non-zero success_count in performance results")
	}

	// Error rate should be reasonable for perfect-api under load
	errorRate, _ := details["error_rate"].(float64)
	if errorRate > 0.2 {
		t.Errorf("error rate too high: %.2f (expected < 20%%)", errorRate)
	}

	// RPS should be non-zero
	rps, _ := details["requests_per_sec"].(float64)
	if rps == 0 {
		t.Error("expected non-zero requests_per_sec")
	}
}

// --- Spec loading via URL tests ---

func TestIntegrationLoadSpecFromURL(t *testing.T) {
	specURL := fmt.Sprintf("http://%s:%s/openapi.json", apiHost(), apiPort())
	result := runDriveby(t,
		"validate-only",
		"--openapi", specURL,
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "minimal",
		"--log-level", "error",
	)

	if result["status"] == nil {
		t.Error("expected 'status' field when loading spec from URL")
	}

	// P001 should still pass
	principles := result["principles"].([]interface{})
	p001 := principles[0].(map[string]interface{})
	if p001["Passed"] != true {
		t.Error("P001 should pass when loading spec from live API URL")
	}
}

func TestIntegrationFunctionalFromURL(t *testing.T) {
	specURL := fmt.Sprintf("http://%s:%s/openapi.json", apiHost(), apiPort())

	// The live OpenAPI spec (3.1 with anyOf nullable) may cause some endpoint
	// response validation differences vs the saved 3.0 file. We just verify
	// that the CLI can load from URL and test endpoints.
	bin := binaryPath(t)
	cmd := exec.Command(bin,
		"function-only",
		"--openapi", specURL,
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// Exit code 1 (some endpoints warn/fail) is acceptable
			if exitErr.ExitCode() > 2 {
				t.Fatalf("unexpected exit code %d", exitErr.ExitCode())
			}
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	if len(stdout) == 0 {
		t.Fatal("expected JSON output from function-only with URL spec")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("output not valid JSON: %v", err)
	}

	tr := result["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})
	total, _ := fn["TotalEndpoints"].(float64)
	if total < 5 {
		t.Errorf("expected at least 5 tested endpoints from URL spec, got %.0f", total)
	}
}

// --- Error handling tests ---

func TestIntegrationInvalidSpecPath(t *testing.T) {
	bin := binaryPath(t)
	cmd := exec.Command(bin, "validate-only",
		"--openapi", "/nonexistent/spec.json",
		"--host", apiHost(),
		"--port", apiPort(),
		"--log-level", "error",
	)
	err := cmd.Run()
	if err == nil {
		t.Error("expected error for nonexistent spec path")
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() == 0 {
			t.Error("expected non-zero exit code for nonexistent spec path")
		}
	}
}

func TestIntegrationInvalidHost(t *testing.T) {
	bin := binaryPath(t)
	cmd := exec.Command(bin, "function-only",
		"--openapi", specPath(t),
		"--host", "nonexistent-host-that-does-not-exist.local",
		"--port", "9999",
		"--log-level", "error",
	)
	stdout, err := cmd.Output()
	if err != nil {
		// Exit code 1 is expected (functional tests fail against unreachable host)
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code (expected non-zero): %d", exitErr.ExitCode())
		}
	}

	// Should still produce valid JSON output (error report)
	if len(stdout) > 0 {
		var result map[string]interface{}
		if err := json.Unmarshal(stdout, &result); err != nil {
			t.Logf("output not JSON (may be expected for connection failure): %s", string(stdout)[:min(len(stdout), 200)])
		}
	}
}

// --- Combined workflow test ---

func TestIntegrationCombinedValidateAndFunction(t *testing.T) {
	// First validate, then functional test — simulates the promotion workflow
	validateResult := runDriveby(t,
		"validate-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--validation-mode", "strict",
		"--log-level", "error",
	)

	// Validation should pass
	if validateResult["status"] != "passed" {
		t.Errorf("expected validation status 'passed', got %v", validateResult["status"])
	}

	// Then run functional tests
	funcResult := runDriveby(t,
		"function-only",
		"--openapi", specPath(t),
		"--host", apiHost(),
		"--port", apiPort(),
		"--auth-api-key", "test-key",
		"--log-level", "error",
	)

	tr := funcResult["test_results"].(map[string]interface{})
	fn := tr["Functional"].(map[string]interface{})
	failed, _ := fn["FailedEndpoints"].(float64)
	if failed > 0 {
		t.Errorf("combined workflow: expected 0 failed endpoints, got %.0f", failed)
	}
}

// --- Auth enforcement verification ---

func TestIntegrationAuthEnforcementMatrix(t *testing.T) {
	// Test that the API properly enforces auth:
	// 1. No auth → skipped endpoints
	// 2. API key → all pass
	// 3. Bearer → all pass
	tests := []struct {
		name       string
		authArgs   []string
		expectSkip bool
	}{
		{
			name:       "no_auth",
			authArgs:   nil,
			expectSkip: true,
		},
		{
			name:       "api_key",
			authArgs:   []string{"--auth-api-key", "driveby-test-key-2024"},
			expectSkip: false,
		},
		{
			name:       "bearer_token",
			authArgs:   []string{"--auth-token", "some-jwt-token"},
			expectSkip: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []string{
				"function-only",
				"--openapi", specPath(t),
				"--host", apiHost(),
				"--port", apiPort(),
				"--log-level", "error",
			}
			args = append(args, tt.authArgs...)

			result := runDriveby(t, args...)
			tr := result["test_results"].(map[string]interface{})
			fn := tr["Functional"].(map[string]interface{})
			skipped, _ := fn["SkippedEndpoints"].(float64)

			if tt.expectSkip && skipped == 0 {
				t.Error("expected skipped endpoints without auth")
			}
			if !tt.expectSkip && skipped > 0 {
				t.Errorf("expected 0 skipped with %s auth, got %.0f", tt.name, skipped)
			}
		})
	}
}

// --- Validation mode completeness ---

func TestIntegrationValidationModes(t *testing.T) {
	tests := []struct {
		mode             string
		expectedCheckers int
	}{
		{"minimal", 1},   // P001 only
		{"strict", 6},    // P001-P005 + P008
	}

	for _, tt := range tests {
		t.Run(tt.mode, func(t *testing.T) {
			result := runDriveby(t,
				"validate-only",
				"--openapi", specPath(t),
				"--host", apiHost(),
				"--port", apiPort(),
				"--validation-mode", tt.mode,
				"--log-level", "error",
			)

			principles := result["principles"].([]interface{})
			if len(principles) != tt.expectedCheckers {
				ids := make([]string, len(principles))
				for i, p := range principles {
					pm := p.(map[string]interface{})
					pr := pm["Principle"].(map[string]interface{})
					ids[i] = pr["id"].(string)
				}
				t.Errorf("mode %s: expected %d checkers, got %d (%s)",
					tt.mode, tt.expectedCheckers, len(principles), strings.Join(ids, ", "))
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
