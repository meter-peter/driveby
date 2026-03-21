# DriveBy — Documentation-Driven Testing (DDT)

## Project Vision
DriveBy is a thesis-ready API validation framework implementing Documentation-Driven Testing (DDT) — a methodology for automated API quality assurance in the GitOps era. This is the diploma thesis project: "Documentation-Driven Testing (DDT): A Paradigm for Automated API Quality Assurance in the GitOps Era".

## DDT Quick Context

### Three Axioms
1. **Completeness** — The API specification must fully describe the API contract
2. **Determinism** — Same specification + same API = same validation results
3. **Observability** — Validation produces measurable, actionable results

### Nine Principles
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
| P009 | Test Readiness | Determinism | Complete |

## Monorepo Structure

| Directory | Purpose | Status | CLAUDE.md |
|-----------|---------|--------|-----------|
| `driveby-cli/` | Go CLI tool — the core validation engine | Active development | [`driveby-cli/CLAUDE.md`](driveby-cli/CLAUDE.md) |
| `apis/` | Sample APIs for testing (5 APIs: perfect, bad-docs, no-auth, slow, broken) | 5 APIs complete | [`apis/CLAUDE.md`](apis/CLAUDE.md) |
| `kubernetes/` | Helm chart (Crossplane XRDs + compositions), examples | Deployed | [`kubernetes/CLAUDE.md`](kubernetes/CLAUDE.md) |
| `samples/` | Example configs, reports | Removed (v3.0.0) | — |
| `thesis/` | LaTeX thesis document (10 chapters) | In progress | [`thesis/CLAUDE.md`](thesis/CLAUDE.md) |
| `tools/` | Python/bash utilities for batch testing | Complete | [`tools/CLAUDE.md`](tools/CLAUDE.md) |
| `results/` | DriveBy validation test results (JSON) | Active | [`results/CLAUDE.md`](results/CLAUDE.md) |
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
- 9 validation principles (P001-P009), each in its own file implementing `PrincipleChecker`
- Engine orchestrates principle checking based on validation mode (minimal/strict/test-ready/test-only)

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
- **XSDLC** (`driveby.io/v1alpha1`) — fully turnkey GitOps delivery pipeline. One CR provisions a dedicated gitops repo, branches, ArgoCD apps, promoter, and quality gates
- A single XSDLC CR generates: GitOps Repository, ServiceAccount, EventBus, RBAC, WorkflowTemplates, EventSources, Sensors, Ingresses, ScmProvider, GitRepository, PromotionStrategy, ArgoCDCommitStatus, BranchProtection rules
- **Two-repo model with Source Hydrator** (v3.0.0): XSDLC auto-creates a dedicated gitops repo per app (via `provider-upjet-github` Repository). The software repo is untouched. The gitops repo's `main` branch holds dry manifests in `dry/base/` + `dry/overlays/<env>/` (Kustomize). ArgoCD Source Hydrator renders each overlay and writes hydrated output (with `hydrator.metadata`) to `environment/<env>-next` branches. The Promoter creates PRs from `-next` to active branches; gates fire on those PRs.
- **Per-environment overlays**: Each environment has its own Kustomize overlay (`dry/overlays/<env>/`), enabling per-env customization (replicas, env vars, images). Each ArgoCD Application uses `sourceHydrator` pointing to its overlay — no linear propagation between environments.
- **BYOCI boundary**: Developers update dry manifests on `main` in the gitops repo. The hydrator + Promoter handle the rest. XSDLC does NOT generate any CI/CD workflows.
- **Multi-check gates**: Each gate defines an ordered `checks` array — each check becomes a DAG step
- **Check types**: `validate-only` (static validation), `functional-test` (P006), `load-test` (k6 load testing)
- **`loadTestConfig`**: configures load testing parameters — `concurrentUsers`, `testDuration`, `maxLatencyP95`, `minSuccessRate`
- Helm chart (v3.0.0) installs `provider-kubernetes` + Crossplane functions + XRD + composition
- **ArgoCD Application generation**: Always generated for every environment using `sourceHydrator` (drySource → overlay, syncSource → env branch, hydrateTo → env-next branch). All apps get `autoSync` with selfHeal. `autoMerge: false` only controls Promoter PR merge behavior, not ArgoCD sync.
- **Simplified UX**: ~35 lines of YAML per app. Only `repository` and `environments` required. All cluster config from `values.yaml`. No EnvironmentConfigs.
- **Two-tier config**: `values.yaml` defaults → XRD spec overrides (EnvironmentConfig layer removed)
- **`gitopsRepository.name`**: Optional override for the gitops repo name (default: `<repository.name>-gitops`)
- **Branch protection**: When `githubProvider.enabled`, XSDLC auto-provisions GitHub BranchProtection rules per gated environment + `*-next` branch protection (on the gitops repo)
- See `docs/deployment-guide.md` for installation, `docs/quality-gate-sdlc.md` for architecture

## Releasing

1. Ensure CI is green on `main`
2. Tag: `git tag -a v3.0.0 -m "DriveBy v3.0.0: two-repo GitOps delivery pipeline"`
3. Push: `git push origin v3.0.0`
4. The `release.yml` workflow produces:
   - GitHub Release with 6 platform binaries + SHA256 checksums (via GoReleaser)
   - Docker images: `ghcr.io/meter-peter/driveby:3.0.0`, `:3.0`, `:latest`
   - Helm chart: `oci://ghcr.io/meter-peter/charts/driveby:3.0.0`

## Target Cluster
- Cluster: `private.novelcore.org`
- ArgoCD already installed — do NOT reinstall, use separate namespace + AppProject
