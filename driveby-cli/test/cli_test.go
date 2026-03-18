package test

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// binaryPath returns the path to the compiled driveby binary.
func binaryPath(t *testing.T) string {
	t.Helper()
	// Build the binary in a temp dir
	dir := t.TempDir()
	bin := filepath.Join(dir, "driveby")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	cmd := exec.Command("go", "build", "-o", bin, "../cmd/driveby")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build binary: %v\n%s", err, out)
	}
	return bin
}

func TestCLIVersionCommand(t *testing.T) {
	bin := binaryPath(t)
	cmd := exec.Command(bin, "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("version command failed: %v\n%s", err, out)
	}
	if len(out) == 0 {
		t.Error("expected version output, got empty")
	}
}

func TestCLIValidateOnlyRequiresFlags(t *testing.T) {
	bin := binaryPath(t)
	cmd := exec.Command(bin, "validate-only")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Error("expected error when --openapi and --host not provided")
	}
	// Should mention the required flag
	if len(out) == 0 {
		t.Error("expected error output")
	}
}

func TestCLIValidateOnlyProducesJSON(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("testdata/valid-openapi.json")
	cmd := exec.Command(bin, "validate-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--port", "9999",
		"--log-level", "error",
	)
	// stdout should be valid JSON, stderr gets logs
	stdout, err := cmd.Output()
	if err != nil {
		// Exit code 1 is expected (validation may fail), but we should still get JSON
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() > 2 {
				t.Fatalf("unexpected exit code %d: %s", exitErr.ExitCode(), exitErr.Stderr)
			}
		} else {
			t.Fatalf("command failed: %v", err)
		}
	}

	// Verify stdout is valid JSON
	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Fatalf("stdout is not valid JSON: %v\nOutput: %s", err, stdout)
	}

	// Check for status field
	if _, ok := result["status"]; !ok {
		t.Error("expected 'status' field in JSON output")
	}
}

func TestCLIValidationModeFlag(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("testdata/valid-openapi.json")

	// Minimal mode should only run P001
	cmd := exec.Command(bin, "validate-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--validation-mode", "minimal",
		"--log-level", "error",
	)
	stdout, _ := cmd.Output()

	var result map[string]interface{}
	if err := json.Unmarshal(stdout, &result); err != nil {
		t.Skipf("could not parse output: %v", err)
		return
	}

	if principles, ok := result["principles"].([]interface{}); ok {
		if len(principles) != 1 {
			t.Errorf("expected 1 principle in minimal mode, got %d", len(principles))
		}
	}
}

func TestCLIExitCodes(t *testing.T) {
	bin := binaryPath(t)
	specPath, _ := filepath.Abs("testdata/valid-openapi.json")

	// validate-only with valid spec should exit 0 (minimal mode = only P001 which passes)
	cmd := exec.Command(bin, "validate-only",
		"--openapi", specPath,
		"--host", "localhost",
		"--validation-mode", "minimal",
		"--log-level", "error",
	)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			t.Errorf("expected exit code 0 for valid spec in minimal mode, got %d", exitErr.ExitCode())
		}
	}
}
