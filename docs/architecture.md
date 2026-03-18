# DriveBy System Architecture

## Overview

DriveBy is a CLI-based API validation framework that implements Documentation-Driven Testing (DDT). It reads an OpenAPI specification, runs it through a configurable set of principle checkers, and produces a structured validation report. The system is designed for CI/CD integration, Kubernetes-native deployment, and GitOps workflows.

## Component Diagram

```
+-------------------+
|    CLI (cobra)    |  User-facing commands: validate, report, version
+-------------------+
         |
         v
+-------------------+
|   Engine          |  Orchestrates validation: loads spec, selects principles, runs checks
+-------------------+
         |
    +----+----+
    |         |
    v         v
+--------+ +----------+
|Registry| |  Config  |  Registry: maps principle IDs to checkers; Config: CLI flags + defaults
+--------+ +----------+
    |
    v
+-------------------+
|  PrincipleChecker |  Interface: Check(spec APISpec, config ValidatorConfig) PrincipleResult
+-------------------+
  |   |   |   |   |
  v   v   v   v   v
P001 P002 P003 ... P008   (each in its own file)
         |
         v
+-------------------+
|    spec.APISpec   |  Abstraction over kin-openapi types — mandatory interface
+-------------------+
         |
         v
+-------------------+
|   loader          |  Loads OpenAPI from file, URL, or stdin; returns APISpec
+-------------------+
         |
         v
+-------------------+
|   types           |  Standalone type definitions — no dependencies on other internal packages
+-------------------+
```

## Data Flow

The complete data flow from user invocation to output:

```
CLI flags (--spec, --mode, --base-url, --output)
    |
    v
ValidatorConfig (parsed from CLI flags + defaults)
    |
    v
Engine.Validate(config)
    |
    +---> loader.Load(config.SpecPath) ---> spec.APISpec
    |
    +---> Registry.GetPrinciples(config.Mode) ---> []PrincipleChecker
    |
    +---> for each checker:
    |         checker.Check(apiSpec, config) ---> PrincipleResult
    |
    +---> Aggregate []PrincipleResult ---> ValidationReport
    |
    v
ValidationReport ---> JSON stdout (or file, depending on --output flag)
```

### Key Types in the Flow

1. **ValidatorConfig** — Aggregated configuration from CLI flags: spec path, validation mode, base URL, output format, thresholds.
2. **spec.APISpec** — Abstract interface wrapping the parsed OpenAPI document. Exposes methods like `GetPaths()`, `GetSchemas()`, `GetSecuritySchemes()` without leaking kin-openapi internals.
3. **PrincipleResult** — Result of a single principle check: principle ID, pass/fail, severity, list of findings, execution time.
4. **ValidationReport** — Aggregated results: all principle results, overall pass/fail, summary statistics, metadata (spec path, timestamp, mode).

## Package Dependency Graph

```
types           (standalone, no internal imports)
  ^
  |
spec            (imports types)
  ^
  |
loader          (imports spec, types)
  ^
  |
principles      (imports spec, types)
  ^
  |
testing         (imports spec, types — functional.go, performance.go)
  ^
  |
engine          (imports principles, testing, loader, spec, types)
  ^
  |
cli             (imports engine, types)
```

**Dependency rule:** Each package may only import packages below it in this graph. Circular dependencies are forbidden and will be caught by `go vet`.

## Spec Abstraction: Design Rationale

The `spec.APISpec` interface is the most critical architectural decision in DriveBy. It serves three purposes:

### 1. Isolation from Third-Party Types
The `kin-openapi` library exposes `*openapi3.T` and related types. If principle checkers used these types directly, every kin-openapi version bump could break every principle checker. The `APISpec` interface insulates checkers from library internals.

### 2. Adapter Pattern for Multiple Spec Formats
The `APISpec` interface enables support for multiple specification formats (OpenAPI 3.0, OpenAPI 3.1, potentially AsyncAPI) through different adapter implementations. Each adapter translates format-specific types into the common `APISpec` interface.

### 3. Testability
Principle checkers can be tested with mock `APISpec` implementations, without needing to parse actual OpenAPI files. This enables unit tests that are fast, deterministic, and focused on checker logic.

**Invariant:** No file in `internal/principles/` or `internal/testing/` may import `github.com/getkin/kin-openapi`. All access to parsed specification data must go through `spec.APISpec`.

## Extension Points

### Adding a New Principle

1. Create `driveby-cli/internal/principles/p00N_name.go`
2. Implement the `PrincipleChecker` interface:
   ```go
   type PrincipleChecker interface {
       ID() string
       Name() string
       Description() string
       Severity() Severity
       Check(spec spec.APISpec, config ValidatorConfig) PrincipleResult
   }
   ```
3. Register the checker in `driveby-cli/internal/principles/registry.go`:
   ```go
   func init() {
       Register(&P00NNameChecker{})
   }
   ```
4. Add test file: `driveby-cli/test/p00N_test.go`
5. Add test data: `driveby-cli/testdata/p00N-test-api.json`
6. Document: `docs/principles/P00N.md`

### Adding a New Spec Format

1. Create `driveby-cli/internal/spec/format_adapter.go`
2. Implement the `APISpec` interface for the new format
3. Update `driveby-cli/internal/loader/loader.go` to detect and use the new adapter
4. Add test data in the new format

## Key Design Decisions

### 1. No `os.Exit()` in Command Handlers
All Cobra command handlers use `RunE` (not `Run`) and return errors instead of calling `os.Exit()`. This ensures:
- Testability: handlers can be invoked in tests without killing the test process
- Composability: errors propagate cleanly to the CLI framework
- Consistency: Cobra handles exit codes based on returned errors

### 2. Types Package is Standalone
The `internal/types/` package has zero internal imports. It defines:
- `Severity` (Critical, Warning, Info)
- `ValidationMode` (Minimal, Strict, TestOnly)
- `PrincipleResult`, `Finding`, `ValidationReport`
- `ValidatorConfig`

This ensures types can be imported by any package without creating circular dependencies.

### 3. Config Shared via `config.go`
CLI configuration (flag defaults, environment variable mappings, config file parsing) lives in a single `config.go` file within the `cli` package. No configuration logic is duplicated across commands.

### 4. JSON as Primary Output Format
The validation report is emitted as JSON to stdout by default. This enables:
- Piping to `jq` for ad-hoc analysis
- Parsing by CI/CD systems for gate decisions
- Storage in artifact repositories for historical comparison
- Consumption by the thesis evaluation scripts

### 5. Mode-Aware Principle Checking
Each principle checker receives the `ValidatorConfig` including the validation mode. Checkers adjust their strictness:
- **Minimal:** Quick checks, no network calls, lenient thresholds
- **Strict:** All checks, including optional ones, strict thresholds
- **Test-only:** Only P006 (contract) and P007 (performance), requires live API

## Deployment Architecture

### Local Development
```bash
cd driveby-cli && go build ./cmd/driveby
./driveby validate --spec path/to/openapi.json --mode strict
```

### Docker Compose (Integration Testing)
```bash
make up        # Starts perfect-api + any dependencies
make validate  # Runs driveby against perfect-api
```

### Kubernetes (Production)
- Helm chart in `kubernetes/helm/driveby/`
- Raw YAML examples in `kubernetes/raw/`
- Designed for CronJob-based periodic validation or CI/CD job integration
- ArgoCD-compatible: GitOps deployment via Application manifest

## Source References

- CLI entry point: `driveby-cli/cmd/driveby/main.go`
- Engine: `driveby-cli/internal/engine/`
- Principles: `driveby-cli/internal/principles/`
- Spec abstraction: `driveby-cli/internal/spec/`
- Loader: `driveby-cli/internal/loader/`
- Types: `driveby-cli/internal/types/`
- Testing (functional/perf): `driveby-cli/internal/testing/`
