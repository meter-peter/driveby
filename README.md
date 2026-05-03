# DriveBy — Documentation-Driven Testing for GitOps

**Automated API quality gates powered by your OpenAPI spec.** DriveBy validates APIs against 9 documentation-driven principles and blocks promotions when they fail — no manual QA, no broken deployments.

```
Developer opens PR → Webhook → Argo Events → DriveBy validation → Commit status → Merge or block
```

## Why DriveBy?

Your OpenAPI spec already describes your API contract. DriveBy treats it as the single source of truth and validates your live API against it — automatically, on every PR.

- **Spec-first quality gates** — validates compliance, documentation, schemas, security, error handling, versioning
- **GitOps-native** — deploys as Crossplane custom resources, triggers via Argo Events webhooks
- **Declarative** — one YAML per API, one template per project. No scripts, no manual wiring
- **PR feedback** — commit status + PR comment with full validation report

## How It Works

### The Quality Gate Pipeline

```
┌──────────────────────────────────────────────────────────────────┐
│                        GitOps Repository                         │
│                                                                  │
│  Developer opens promotion PR (dev → staging)                    │
│       │                                                          │
│       ▼                                                          │
│  GitHub Webhook ──► Argo EventSource ──► Sensor                  │
│                                            │                     │
│                                            ▼                     │
│                                    Argo Workflow DAG              │
│                              ┌─────────────────────┐             │
│                              │  set-pending-status  │             │
│                              │         │            │             │
│                              │  wait-for-source     │             │
│                              │         │            │             │
│                              │  driveby validate    │ ◄── P001-P009
│                              │         │            │             │
│                              │  functional-test     │             │
│                              │        ╱ ╲           │             │
│                              │  report   comment-pr │             │
│                              └─────────────────────┘             │
│                                            │                     │
│       ┌────────────────────────────────────┘                     │
│       ▼                                                          │
│  Commit Status: success/failure                                  │
│       │                                                          │
│       ▼                                                          │
│  ✅ Merge allowed  ──► ArgoCD syncs staging                      │
│  ❌ PR blocked     ──► Fix and re-push                           │
└──────────────────────────────────────────────────────────────────┘
```

### Crossplane Quality Gates

DriveBy uses **a single Crossplane custom resource (XSDLC)** to manage the entire promotion pipeline declaratively:

> **Naming convention:** the thesis text refers to `perfect-api` as `non-critical-api` — see thesis Chapter 7 §7.2. Source identifiers (this repo, namespaces, container images, CRs) keep the historical `perfect-api` name for stability; the rename is scheduled for v4.0.

```yaml
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: perfect-api
  namespace: driveby
spec:
  gitopsRepository:
    owner: novelcore
    name: perfect-api-gitops
  environments:
    - name: dev
    - name: staging
      gate:
        checks:
          - type: validate-only
            validationConfig:
              validationMode: test-ready
          - type: functional-test
    - name: prod
      autoMerge: false
      gate:
        checks:
          - type: validate-only
          - type: load-test
            loadTestConfig:
              concurrentUsers: 50
              testDuration: "2m"
              maxLatencyP95: "200ms"
              minSuccessRate: 0.995
```

Only two required fields: `gitopsRepository` (owner + name) and `environments` (minimum 2). Each environment can optionally define a `gate` with an ordered `checks` array — each check becomes a DAG step in the validation workflow.

Creates: ServiceAccount, EventBus, RBAC, WorkflowTemplates, EventSources, Sensors, Ingresses, ScmProvider, GitRepository, PromotionStrategy, ArgoCDCommitStatus, BranchProtection, deploy workflow

**Add a new API?** Apply one more `XSDLC` — that's it.

## DDT Methodology

DriveBy implements **Documentation-Driven Testing (DDT)** — three axioms, nine validation principles:

| Axiom | Principle | What It Checks |
|-------|-----------|---------------|
| **Completeness** | P001: OpenAPI Compliance | Spec parses, valid structure |
| | P002: Documentation Quality | Descriptions, examples, completeness |
| | P003: Error Handling | Error responses documented (4xx/5xx) |
| | P004: Schema Definitions | Types, constraints, required fields |
| **Determinism** | P006: Functional Testing | Live endpoint responses match spec |
| | P009: Test Readiness | Spec ready for meaningful functional testing |
| **Observability** | P005: Security Standards | Auth schemes defined and applied |
| | P007: Performance Testing | P95 latency, success rate under load |
| | P008: Versioning Strategy | Version info present and consistent |

### Validation Modes

| Mode | Principles | Use Case |
|------|-----------|----------|
| `minimal` | P001 | Fast CI gate — basic spec compliance |
| `strict` | P001-P005, P008 | Production readiness — full static analysis |
| `test-ready` | P001, P004, P009 | Pre-flight check for meaningful testing |
| `test-only` | P006, P007 | Runtime testing only — functional + load |

## Installation

### Kubernetes (recommended)

```bash
# 1. Install the Helm chart
helm install driveby oci://ghcr.io/meter-peter/charts/driveby \
  --version 2.0.0 \
  --set crossplane.enabled=true

# 2. Apply an XSDLC CR for your API
kubectl apply -f kubernetes/examples/novelcore-perfect-api/xsdlc-perfect-api.yaml

# 3. Configure GitHub webhook on your gitops repo
#    URL: https://<app>-webhook.<domain>/<app>-<trigger>
```

See [docs/deployment-guide.md](docs/deployment-guide.md) for the full guide with prerequisites, secrets, and verification checklist.

### CLI Binary

```bash
# From GitHub Releases
curl -sL https://github.com/meter-peter/driveby/releases/latest/download/driveby_latest_linux_amd64.tar.gz | tar xz
sudo mv driveby /usr/local/bin/

# From source
git clone https://github.com/meter-peter/driveby.git
cd driveby/driveby-cli && go build -o driveby ./cmd/driveby
```

### Docker

```bash
docker run --rm ghcr.io/meter-peter/driveby:latest validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io --protocol https --port 443 \
  --validation-mode strict
```

## CLI Usage

```bash
# Validate API documentation (static analysis)
driveby validate-only --openapi openapi.json --host api.example.com --validation-mode strict

# Functional testing (live endpoint verification)
driveby function-only --openapi openapi.json --host api.example.com

# Load testing
driveby load-only --openapi openapi.json --host api.example.com \
  --concurrent-users 50 --test-duration 300

# Validate a public API by URL
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io --protocol https --port 443
```

### Authentication

```bash
# Bearer token
driveby validate-only --openapi spec.json --host api.example.com \
  --auth-token "your-token"

# API key
driveby validate-only --openapi spec.json --host api.example.com \
  --auth-api-key "your-key" --auth-api-key-header "X-API-Key"

# Basic auth
driveby validate-only --openapi spec.json --host api.example.com \
  --auth-username "user" --auth-password "pass"
```

### GitHub PR Integration

```bash
driveby validate-only --openapi spec.json --host api.example.com \
  --github-comment \
  --github-app-id "$APP_ID" \
  --github-installation-id "$INSTALL_ID" \
  --github-private-key "$PRIVATE_KEY" \
  --github-owner "your-org" --github-repo "your-repo" \
  --github-pr-number "$PR_NUMBER"
```

Full CLI reference: [docs/CLI_USAGE.md](docs/CLI_USAGE.md)

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster                         │
│                                                              │
│  ┌──────────────────────────────────────────────────┐       │
│  │              driveby namespace                    │       │
│  │                                                   │       │
│  │  Crossplane XRD ──► Composition                    │       │
│  │       │                    │                       │       │
│  │       ▼                    ▼                       │       │
│  │  XSDLC (single CR per API)                        │       │
│  │  (RBAC + Workflows + Events + Ingress + Promoter) │       │
│  │                             │                      │       │
│  │  Argo Events: EventBus ◄───┘                      │       │
│  │       │                                           │       │
│  │  Webhook ──► Sensor ──► Argo Workflow             │       │
│  │                              │                     │       │
│  │                         DriveBy CLI               │       │
│  │                         (container)               │       │
│  └──────────────────────────────────────────────────┘       │
│                                                              │
│  ┌────────────────┐  ┌────────────────┐  ┌───────────────┐  │
│  │ app-dev/       │  │ app-staging/   │  │ app-prod/     │  │
│  │ (autoSync)     │──│ (gate-blocked) │──│ (autoSync)    │  │
│  └────────────────┘  └────────────────┘  └───────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### Monorepo Structure

```
driveby-cli/     Go CLI — core validation engine (9 principles, spec abstraction)
kubernetes/      Helm chart (Crossplane XRDs + compositions, Argo Events/Workflows)
docs/            Deployment guide, quality gate SDLC, CLI reference
thesis/          LaTeX diploma thesis
tools/           Batch validation utilities
samples/         Example configs and reports
```

## Configuration Model

DriveBy uses a **two-tier configuration** so every value can be overridden without editing chart source:

| Tier | Scope | How to Set |
|------|-------|-----------|
| `values.yaml` | Chart-wide defaults | `--set defaults.X=Y` |
| XRD spec fields | Per-CR overrides in XSDLC | Set directly in XSDLC spec |

Resolution order (highest priority wins): **XRD spec > values.yaml**

## Documentation

| Document | Content |
|----------|---------|
| [Deployment Guide](docs/deployment-guide.md) | Full Kubernetes installation, secrets, verification |
| [Quality Gate SDLC](docs/quality-gate-sdlc.md) | Architecture, Crossplane XRDs, DDT mapping |
| [GitOps Pipeline](docs/gitops-pipeline.md) | Webhook → Events → Workflow → PR feedback |
| [CLI Usage](docs/CLI_USAGE.md) | All commands, flags, examples |
| [DDT Axioms](docs/ddt-axioms.md) | Methodology: Completeness, Determinism, Observability |
| [Principles P001-P009](docs/principles/) | Per-principle documentation |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, code standards, and how to add a new validation principle.

## License

MIT — see [LICENSE](LICENSE) for details.
