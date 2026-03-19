# DriveBy — Documentation-Driven Testing (DDT)

## Project Vision
DriveBy is a thesis-ready API validation framework implementing Documentation-Driven Testing (DDT) — a methodology for automated API quality assurance in the GitOps era. This is the diploma thesis project: "Documentation-Driven Testing (DDT): A Paradigm for Automated API Quality Assurance in the GitOps Era".

## DDT Quick Context

### Three Axioms
1. **Completeness** — The API specification must fully describe the API contract
2. **Determinism** — Same specification + same API = same validation results
3. **Observability** — Validation produces measurable, actionable results

### Eight Principles
| ID | Principle | Axiom | Status |
|----|-----------|-------|--------|
| P001 | OpenAPI Compliance | Completeness | Complete |
| P002 | Documentation Quality | Completeness | Complete |
| P003 | Error Handling | Completeness | Complete |
| P004 | Schema Definitions | Completeness | Complete |
| P005 | Security Standards | Observability | Complete |
| P006 | Functional Testing | Determinism | Not yet implemented (PrincipleChecker wrapper) |
| P007 | Performance Testing | Observability | Not yet implemented (PrincipleChecker wrapper) |
| P008 | Versioning Strategy | Observability | Complete |

## Monorepo Structure

| Directory | Purpose | Status | CLAUDE.md |
|-----------|---------|--------|-----------|
| `driveby-cli/` | Go CLI tool — the core validation engine | Active development | [`driveby-cli/CLAUDE.md`](driveby-cli/CLAUDE.md) |
| `apis/` | Sample APIs for testing (perfect-api) — source now in [`novelcore/perfect-api`](https://github.com/novelcore/perfect-api) | 1 API complete | [`apis/CLAUDE.md`](apis/CLAUDE.md) |
| `kubernetes/` | Helm chart (Crossplane XRDs + compositions), examples | Deployed | [`kubernetes/CLAUDE.md`](kubernetes/CLAUDE.md) |
| `samples/` | Example configs, reports, demo workflows | Configs complete | [`samples/CLAUDE.md`](samples/CLAUDE.md) |
| `thesis/` | LaTeX thesis document (7 chapters) | In progress | [`thesis/CLAUDE.md`](thesis/CLAUDE.md) |
| `tools/` | Python/bash utilities for batch testing | Complete | [`tools/CLAUDE.md`](tools/CLAUDE.md) |
| `docs/` | Documentation hub | Partial | [`docs/CLAUDE.md`](docs/CLAUDE.md) |
| `.github/` | CI/CD workflows | Active | [`.github/CLAUDE.md`](.github/CLAUDE.md) |

### Key Sub-Directory CLAUDE.md Files
- [`driveby-cli/internal/principles/CLAUDE.md`](driveby-cli/internal/principles/CLAUDE.md) — Principle checker implementations
- [`driveby-cli/internal/spec/CLAUDE.md`](driveby-cli/internal/spec/CLAUDE.md) — APISpec abstraction layer
- [`apis/perfect-api/CLAUDE.md`](apis/perfect-api/CLAUDE.md) — Reference API implementation

## Architecture
- Module: `github.com/meter-peter/driveby/driveby-cli`
- Dependency flow: `types` <- `spec` <- `loader` <- `principles` <- `testing` <- `engine` <- `cli`
- The `spec.APISpec` abstraction is **mandatory** — no raw `*openapi3.T` in validators
- 8 validation principles (P001-P008), each in its own file implementing `PrincipleChecker`
- Engine orchestrates principle checking based on validation mode (minimal/strict/test-only)

## Build & Test
```bash
cd driveby-cli && go build ./cmd/driveby    # Build
cd driveby-cli && go test ./test/...         # Test
cd driveby-cli && go vet ./...               # Lint
make up                                       # Start docker-compose
make validate                                 # Run validation against perfect-api
```

## Public API Validation
DriveBy supports validating remote API specs by URL (no local files needed):
```bash
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io --protocol https --port 443 \
  --validation-mode strict
```
Reference result: Swagger Petstore v3 scores 1/6 in strict mode (only P001 passes).
See `docs/WORKFLOW.md` for full workflow, `tools/` for batch validation of public APIs.

## Code Quality Rules
1. `spec.APISpec` abstraction is mandatory — no raw `*openapi3.T` in principle checkers
2. Each principle checker implements `PrincipleChecker` interface
3. No `os.Exit()` in command handlers — return errors from `RunE`
4. Types live in `internal/types/` — standalone, no circular deps
5. CLI config shared via `config.go` — no duplication

## Maintenance Rules
1. **When modifying source files**, update the corresponding CLAUDE.md in that directory
2. **When modifying source files**, update the corresponding docs/ file if documented behavior changes
3. **When adding a new principle**, create: source file, test file, docs/principles/PXXX.md, and update principles/CLAUDE.md
4. **When adding a new directory**, create a CLAUDE.md in it and add it to this file's monorepo table

## Crossplane Workflow
When writing Crossplane resources:
1. Use `context7 resolve-library-id` first for provider docs
2. Use `context7 query-docs` for XRD schemas, Composition patterns
3. Reference kubecore-operator at `/home/meter-peter/development/novelcore/kubecore-operator/compositions/`
4. Only fall back to web search if context7 lacks the information

## Crossplane Quality Gates (XSDLC)
- **XQualityGateTemplate** (`driveby.io/v1alpha1`) — declares validation workflow (RBAC + WorkflowTemplate + CommitStatus steps)
- **XQualityGate** (`driveby.io/v1alpha1`) — declares per-API event pipeline (EventBus + EventSource + Sensor + Ingress) + promoter resources (ScmProvider, GitRepository, PromotionStrategy, ArgoCDCommitStatus)
- Helm chart (v0.4.0) installs `provider-kubernetes` + Crossplane functions + XRDs + compositions
- Quality gates trigger on **promotion PRs in the gitops repo**, validate against the **source env** (dev), and gate promotion to **target env** (staging)
- **GitOps Promoter integration**: auto-provisions ScmProvider, GitRepository, PromotionStrategy, and CommitStatus CRD workflow steps for fully automated environment promotion
- See `docs/deployment-guide.md` for installation, `docs/quality-gate-sdlc.md` for architecture

## Releasing

1. Ensure CI is green on `main`
2. Tag: `git tag -a v0.4.0 -m "DriveBy v0.4.0: automated promotion pipeline"`
3. Push: `git push origin v0.4.0`
4. The `release.yml` workflow produces:
   - GitHub Release with 6 platform binaries + SHA256 checksums (via GoReleaser)
   - Docker images: `ghcr.io/meter-peter/driveby:0.4.0`, `:0.4`, `:latest`
   - Helm chart: `oci://ghcr.io/meter-peter/charts/driveby:0.4.0`

## Target Cluster
- Cluster: `private.novelcore.org`
- ArgoCD already installed — do NOT reinstall, use separate namespace + AppProject
