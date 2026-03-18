//go:build integration

package test

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIntegrationValidateOnly(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("../apis/perfect-api/openapi.json")

	cmd := exec.Command(bin, "validate-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--port", "8000",
		"--validation-mode", "strict",
		"--log-level", "error",
	)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code: %d", exitErr.ExitCode())
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v", err)
	}
	if _, ok := result["status"]; !ok {
		t.Error("expected 'status' field in output")
	}
}

func TestIntegrationFunctionOnly(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("../apis/perfect-api/openapi.json")

	cmd := exec.Command(bin, "function-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--port", "8000",
		"--log-level", "error",
	)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code: %d", exitErr.ExitCode())
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v", err)
	}
}

func TestIntegrationTestOnly(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("../apis/perfect-api/openapi.json")

	cmd := exec.Command(bin, "test-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--port", "8000",
		"--test-duration", "2",
		"--concurrent-users", "2",
		"--log-level", "error",
	)
	stdout, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code: %d", exitErr.ExitCode())
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v", err)
	}
}
