# Principles Package — DDT Principle Checkers

## Interface
```go
type PrincipleChecker interface {
    ID() string
    Check(ctx context.Context, apiSpec spec.APISpec, mode types.ValidationMode) types.PrincipleResult
}
```
Defined in `checker.go`. Every principle checker must implement this interface. All checkers receive `spec.APISpec` (never raw kin-openapi types).

## File Status

| File | Principle | Description | Status | LOC |
|------|-----------|-------------|--------|-----|
| `p001_compliance.go` | P001 | OpenAPI Compliance — validates spec structure, required fields, valid HTTP methods | Complete | ~180 |
| `p002_documentation.go` | P002 | Documentation Quality — checks descriptions, summaries, examples on operations/schemas | Complete | ~197 |
| `p003_errors.go` | P003 | Error Handling — verifies error response definitions (4xx/5xx), consistent error schemas | Complete | ~234 |
| `p004_schema.go` | P004 | Schema Definitions — validates request/response schemas, type coverage, allOf flattening | Complete | ~353 |
| `p005_security.go` | P005 | Security Standards — checks security schemes, global security, HTTPS enforcement | Complete | ~221 |
| `p008_versioning.go` | P008 | Versioning Strategy — validates API versioning in paths/headers, deprecation markers | Complete | ~163 |
| — | P006 | Functional Testing — runtime endpoint testing (auth, responses, contracts) | **Not yet implemented** | — |
| — | P007 | Performance Testing — load testing, latency thresholds, rate limiting | **Not yet implemented** | — |

P006 and P007 require live API access and are partially covered by `internal/testing/` (functional.go, performance.go, auth.go) but do not yet have `PrincipleChecker` wrappers.

## Supporting Files

| File | Purpose |
|------|---------|
| `checker.go` | `PrincipleChecker` interface definition (14 LOC) |
| `registry.go` | Registry pattern: `GetCheckers(mode)` returns checkers filtered by `ValidationMode` (~58 LOC) |

## DDT Axiom Assignments

| Axiom | Principles | Rationale |
|-------|-----------|-----------|
| **Completeness** | P001, P002, P003, P004 | Spec must fully describe the API contract |
| **Determinism** | P006 | Same spec + same API = same test results |
| **Observability** | P005, P007, P008 | Security, performance, and versioning are observable properties |

## Thesis Mapping
- **Chapter 3 (Methodology)**: Each principle maps to a section in Ch.3. The axiom-to-principle mapping is a core thesis contribution.
- When writing about a principle, read the corresponding `pXXX_*.go` file for the exact checks performed.

## Adding a New Principle
1. Create `pXXX_name.go` implementing `PrincipleChecker`
2. Register in `registry.go` under the appropriate modes
3. Add test in `../../test/pXXX_test.go`
4. Document in `docs/principles/PXXX.md`
5. Update this CLAUDE.md
