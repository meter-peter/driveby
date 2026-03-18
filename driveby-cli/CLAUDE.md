# DriveBy CLI — Core Validation Engine

## Module
`github.com/meter-peter/driveby/driveby-cli`

## Dependency Flow
```
types <- spec <- loader <- principles <- testing <- engine <- cli
```
No package may import a package to its left. `types` has zero internal dependencies.

## Package Map

| Package | Path | Purpose |
|---------|------|---------|
| main | `cmd/driveby/main.go` | Entry point, calls `cli.Execute()` |
| cli | `internal/cli/` | Cobra commands: root, validate, functional, loadtest, combined, github, version. Config shared via `config.go`. |
| engine | `internal/engine/engine.go` | Orchestrator: loads spec, selects principles by mode, runs checkers, aggregates results |
| principles | `internal/principles/` | 8 DDT principle checkers (P001-P008), registry, interface definition |
| spec | `internal/spec/` | `APISpec` interface + adapters for OpenAPI 3.x and Swagger 2.x |
| loader | `internal/loader/` | Loads OpenAPI specs from local file or remote URL, auto-detects OpenAPI 3.x vs Swagger 2.0 |
| testing | `internal/testing/` | Runtime testing: `auth.go` (auth probing), `functional.go` (endpoint testing), `performance.go` (load testing) |
| types | `internal/types/` | Standalone types: `ValidationMode`, `PrincipleResult`, `Report`, error types, metrics, logger config |
| report | `internal/report/` | Report generation: `generator.go` (orchestrator), `markdown.go` (Markdown renderer) |
| github | `internal/github/client.go` | GitHub API client for PR comment posting |
| logger | `internal/logger/logger.go` | Structured logging setup |
| util | `internal/util/util.go` | Shared helper functions |

## Build & Test
```bash
go build ./cmd/driveby           # Build binary
go test -count=1 ./test/...      # Run tests (tests live in test/ directory)
go vet ./...                     # Lint
```

## Test Directory
Tests live in `test/`, not alongside source files:
- `cli_test.go` — CLI command tests
- `integration_test.go` — Full pipeline integration tests
- `loader_test.go` — Spec loading tests
- `principles_test.go` — Principle checker unit tests
- `p002_test.go`, `p003_test.go`, `p004_test.go` — Per-principle tests
- `auth_test.go` — Authentication validation tests
- `functional_helpers_test.go` — Functional testing helper tests
- `util_test.go` — Utility function tests
- `testdata/` — Test fixtures (sample OpenAPI specs including `auth-required-api.json`, `nullable-openapi31.json`)

## Validation Modes
- **minimal** — P001 (compliance) only; fast CI gate
- **strict** — P001-P005, P008; comprehensive static analysis
- **test-only** — Skips static validation, runs runtime tests only (P006/P007 when implemented)
- **flexible** — P001 only (same as minimal); allows tests even with some validation failures

## Thesis Mapping
- **Chapter 3 (Methodology)**: Principles P001-P008 map directly to `internal/principles/`
- **Chapter 4 (Architecture)**: Engine orchestration, spec abstraction, dependency flow
- **Chapter 5 (Workflow)**: CLI commands, validation modes, CI integration

## Key Rules
1. `spec.APISpec` abstraction is mandatory -- no raw `*openapi3.T` in principle checkers
2. Each principle checker implements `PrincipleChecker` interface from `checker.go`
3. No `os.Exit()` in command handlers -- return errors from `RunE`
4. Types in `internal/types/` must remain standalone with zero internal imports
5. CLI config shared via `config.go` -- no flag duplication across commands
