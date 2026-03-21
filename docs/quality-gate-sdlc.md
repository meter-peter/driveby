# Quality Gates in the Software Development Lifecycle

How DriveBy implements documentation-driven quality gates using Crossplane and Argo Events in a two-repo GitOps model.

## Quality Gate Concept

A **quality gate** is an automated checkpoint in the SDLC that validates software against defined criteria before allowing progression. In DriveBy, quality gates enforce DDT principles against API specifications at deployment boundaries.

```
Dry manifests on main -> Hydrator renders overlay -> writes to -next branch -> Promoter PR (gitops repo) -> Webhook -> EventSource -> Sensor -> Workflow -> Commit Status -> CommitStatus CRD -> Promoter Auto-Merge -> ArgoCD Sync
```

## End-to-End Flow

### 1. Hydrator Renders Overlay

When dry manifests change on the `main` branch of the gitops repo, the ArgoCD Source Hydrator renders the per-environment Kustomize overlay (e.g., `dry/overlays/dev/`) and writes the hydrated manifests to the corresponding `environment/<env>-next` branch.

### 2. Pull Request Created

The GitOps Promoter detects the new commits on the `-next` branch and opens a PR in the dedicated gitops repository (e.g., `novelcore/perfect-api-gitops`) to promote changes from `environment/<env>-next` to `environment/<env>`. The software repository (e.g., `novelcore/perfect-api`) is never touched by the promotion pipeline.

### 3. GitHub Webhook

GitHub sends a `pull_request` event from the gitops repository to the webhook endpoint exposed by the Argo Events EventSource via Traefik Ingress.

### 4. EventSource Processing

The Argo Events EventSource receives the webhook payload and publishes it to the JetStream EventBus.

### 5. Sensor Triggers Workflow

The Sensor filters events (e.g., `action: opened|reopened|synchronize`) and triggers an Argo Workflow, mapping event payload fields to workflow parameters:

| Event Field | Workflow Parameter |
|---|---|
| `body.number` | `pr-number` |
| `body.pull_request.head.sha` | `head-sha` |
| `body.repository.owner.login` | `github-owner` |
| `body.repository.name` | `github-repo` |

### 6. Validation Workflow (DAG)

Each gate defines an ordered list of **checks** that become sequential DAG steps. The workflow dynamically chains checks based on the gate's `checks` array:

```
set-pending -> health-check -> check-0-<type> -> check-1-<type> -> ... -> check-N-<type>
                                                                                |
                                                   report-success + comment-pr + update-commitstatus-success
```

**Example: staging gate** (validate-only + functional-test):
```
set-pending -> health-check -> check-0-validate-only -> check-1-functional-test
                                                               |
                                   report-success + comment-pr + update-commitstatus-success
```

**Example: prod gate** (validate-only + load-test):
```
set-pending -> health-check -> check-0-validate-only -> check-1-load-test
                                                              |
                                   report-success + comment-pr + update-commitstatus-success
```

### 7. Commit Status as Gate Signal

The workflow reports a GitHub commit status:
- **Context**: `driveby/<env>-gate` (e.g., `driveby/staging-gate`)
- **State**: `pending` -> `success` or `failure`
- **Description**: Human-readable validation result

### 8. GitOps Promoter Integration

The commit status and CommitStatus CRD serve as gate signals for the GitOps Promoter. The XSDLC composition generates all promoter resources:

```yaml
# ScmProvider -- GitHub App auth
apiVersion: promoter.argoproj.io/v1alpha1
kind: ScmProvider
metadata:
  name: perfect-api-github
spec:
  github:
    domain: github.com
  secretRef:
    name: github-app-credentials
  isApp: true
---
# GitRepository -- points to the dedicated GITOPS repo (not the software repo)
apiVersion: promoter.argoproj.io/v1alpha1
kind: GitRepository
metadata:
  name: perfect-api-gitops
spec:
  owner: novelcore
  name: perfect-api-gitops
  scmProviderRef:
    name: perfect-api-github
---
# PromotionStrategy -- environment chain with commit status gates
apiVersion: promoter.argoproj.io/v1alpha1
kind: PromotionStrategy
metadata:
  name: perfect-api-promotion
spec:
  gitRepositoryRef:
    name: perfect-api-gitops
  environments:
    - branch: environment/dev
      autoMerge: true
    - branch: environment/staging
      autoMerge: true
      activeCommitStatuses:
        - key: staging-gate
      proposedCommitStatuses:
        - key: staging-gate
    - branch: environment/prod
      autoMerge: false
      activeCommitStatuses:
        - key: prod-gate
      proposedCommitStatuses:
        - key: prod-gate
```

The workflow also creates CommitStatus CRDs via the `update-commitstatus` step template:

```yaml
apiVersion: promoter.argoproj.io/v1alpha1
kind: CommitStatus
metadata:
  name: staging-gate-<head-sha>
  labels:
    promoter.argoproj.io/commit-status: staging-gate
spec:
  gitRepositoryRef:
    name: perfect-api-gitops
  sha: <head-sha>
  name: driveby/staging-gate
  phase: success  # or pending/failure
  description: "Dev validation passed, safe to promote"
```

## Crossplane XSDLC: Declarative Approach

Instead of manually deploying EventBus + EventSource + Sensor + Ingress + WorkflowTemplate + RBAC + Promoter resources, DriveBy uses a single Crossplane XRD: **XSDLC** (`driveby.io/v1alpha1`).

One XSDLC CR (~30 lines of YAML) defines the entire promotion pipeline for an API, including a dedicated gitops repository, all environments, quality gates, and promoter integration. Only two fields are required: `repository` and `environments`. The XSDLC auto-creates a separate gitops repository (default: `<repo-name>-gitops`) where all environment branches, PRs, and webhooks live. The software repository is never modified.

### Resource Flow

```
XSDLC (1 per API, ~35 lines of YAML)
+-- App-level (shared)
|   +-- ServiceAccount + imagePullSecrets
|   +-- EventBus (JetStream)
|   +-- ScmProvider (GitHub App auth)
|   +-- GitRepository (dedicated gitops repo ref)
|   +-- PromotionStrategy (env chain + gates)
|   +-- ArgoCDCommitStatus (ArgoCD health)
+-- GitOps repo (auto-created: <repo-name>-gitops)
|   +-- Per-environment branches (env + env-next)
|   +-- PRs and webhooks originate here
+-- Per-environment (always generated)
|   +-- ArgoCD Application (sourceHydrator: renders dry/overlays/<env>/ from main → writes to env-next, always autoSync)
+-- Per-gate (one set per gated environment)
    +-- Role + RoleBinding
    +-- WorkflowTemplate (dynamic checks DAG)
    +-- EventSource (GitHub webhook from gitops repo)
    +-- Sensor (event -> workflow trigger)
    +-- Service (webhook endpoint)
    +-- Ingress (TLS termination)
    +-- BranchProtection
```

### XSDLC Example

```yaml
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: perfect-api
  namespace: driveby
spec:
  repository:
    owner: novelcore
    name: perfect-api
  gitopsRepository:
    name: perfect-api-gitops   # auto-created; defaults to <repository.name>-gitops

  apiConfig:
    port: 8000
    openapiEndpoint: /openapi.json

  validationDefaults:
    validationMode: strict

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

### Check Types

Each gate defines an ordered `checks` array. Each check becomes a sequential DAG step in the workflow:

| Check Type | CLI Command | Description |
|---|---|---|
| `validate-only` | `validate-only --validation-mode <mode>` | Static validation (P001-P009 based on mode) |
| `functional-test` | `function-only` | Functional API testing (P006) |
| `load-test` | `load-only` | k6 load testing with configurable thresholds |

### Manifest-Agnostic Design (Two-Repo Model with Source Hydrator)

The XSDLC is **manifest-agnostic** — it does NOT manage environment-specific configuration (replicas, resources, env vars). In the v3.0.0 two-repo model, dry manifests live on the `main` branch of a dedicated gitops repository (auto-created by the XSDLC, default name: `<repo-name>-gitops`) using a Kustomize overlay structure: `dry/base/` for shared manifests and `dry/overlays/<env>/` for per-environment customization. The software repository is never modified by the promotion pipeline.

Engineers update dry manifests on `main` — the ArgoCD Source Hydrator renders each per-environment overlay and writes the hydrated output to the corresponding `environment/<env>-next` branch. Each environment hydrates independently; there is no linear propagation between environments. The XSDLC only manages **quality gates** and **promotion flow** — all branches, PRs, and webhooks originate in the gitops repo while the software repo remains solely for application source code.

### Two-Tier Configuration

All cluster-level configuration (ingress domains, JetStream settings, image registries, secrets) comes from `values.yaml`, baked into the composition at Helm template time. Per-API overrides are set directly in the XSDLC spec. There are no EnvironmentConfigs.

| Tier | Source | Examples |
|---|---|---|
| Cluster defaults | `values.yaml` (`defaults.*`) | `baseDomain`, `clusterIssuer`, `webhookPort`, `imagePullSecret` |
| Per-API overrides | XSDLC spec fields | `apiConfig.port`, `validationDefaults.validationMode`, `gitopsRepository.name`, per-gate config |

## Mapping to DDT Axioms

| DDT Axiom | Quality Gate Implementation |
|---|---|
| **Completeness** | Validates P001-P004: OpenAPI compliance, documentation quality, error handling, schema definitions |
| **Determinism** | Same spec + same API = same validation result; automated, no human judgment |
| **Observability** | Commit status on every PR; PR comment with full validation report; workflow logs |

## Gate Strictness Configuration

The `validationMode` field controls which DDT principles are evaluated:

| Mode | Principles | Use Case |
|---|---|---|
| `minimal` | P001 only | Quick feedback on OpenAPI compliance |
| `strict` | P001-P005, P008 | Full static analysis |
| `test-ready` | P001, P004, P009 | Validates spec is ready for testing |
| `test-only` | P006 | Functional testing only |

Configure via `validationDefaults.validationMode` (applies to all validation gates) or per-gate `gate.validationConfig.validationMode` in the XSDLC spec.
