# Quality Gates in the Software Development Lifecycle

How DriveBy implements documentation-driven quality gates using Crossplane and Argo Events.

## Quality Gate Concept

A **quality gate** is an automated checkpoint in the SDLC that validates software against defined criteria before allowing progression. In DriveBy, quality gates enforce DDT principles against API specifications at deployment boundaries.

```
Developer -> PR -> Webhook -> EventSource -> Sensor -> Workflow -> Commit Status -> CommitStatus CRD -> Promoter Auto-Merge -> ArgoCD Sync
```

## End-to-End Flow

### 1. Pull Request Created

A developer opens a PR against the API repository (e.g., `novelcore/perfect-api`).

### 2. GitHub Webhook

GitHub sends a `pull_request` event to the webhook endpoint exposed by the Argo Events EventSource via Traefik Ingress.

### 3. EventSource Processing

The Argo Events EventSource receives the webhook payload and publishes it to the JetStream EventBus.

### 4. Sensor Triggers Workflow

The Sensor filters events (e.g., `action: opened|reopened|synchronize`) and triggers an Argo Workflow, mapping event payload fields to workflow parameters:

| Event Field | Workflow Parameter |
|---|---|
| `body.number` | `pr-number` |
| `body.pull_request.head.sha` | `head-sha` |
| `body.repository.owner.login` | `github-owner` |
| `body.repository.name` | `github-repo` |

### 5. Validation Workflow (DAG)

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

### 6. Commit Status as Gate Signal

The workflow reports a GitHub commit status:
- **Context**: `driveby/<env>-<gate-type>` (e.g., `driveby/staging-validation`)
- **State**: `pending` -> `success` or `failure`
- **Description**: Human-readable validation result

### 7. GitOps Promoter Integration

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
# GitRepository -- repo reference
apiVersion: promoter.argoproj.io/v1alpha1
kind: GitRepository
metadata:
  name: perfect-api
spec:
  owner: novelcore
  name: perfect-api
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
    name: perfect-api
  activeCommitStatuses:
    - key: staging-validation
    - key: prod-loadtest
  proposedCommitStatuses:
    - key: staging-validation
    - key: prod-loadtest
  environments:
    - branch: environment/dev
      autoMerge: true
    - branch: environment/staging
      autoMerge: true
    - branch: environment/prod
      autoMerge: false
```

The workflow also creates CommitStatus CRDs via the `update-commitstatus` step template:

```yaml
apiVersion: promoter.argoproj.io/v1alpha1
kind: CommitStatus
metadata:
  name: staging-validation-<head-sha>
  labels:
    promoter.argoproj.io/commit-status: staging-validation
spec:
  gitRepositoryRef:
    name: perfect-api
  sha: <head-sha>
  name: driveby/staging-validation
  phase: success  # or pending/failure
  description: "Dev validation passed, safe to promote"
```

## Crossplane XSDLC: Declarative Approach

Instead of manually deploying EventBus + EventSource + Sensor + Ingress + WorkflowTemplate + RBAC + Promoter resources, DriveBy uses a single Crossplane XRD: **XSDLC** (`driveby.io/v1alpha1`).

One XSDLC CR (~30 lines of YAML) defines the entire promotion pipeline for an API, including all environments, quality gates, and promoter integration. Only two fields are required: `repository` and `environments`.

### Resource Flow

```
XSDLC (1 per API, ~35 lines of YAML)
+-- App-level (shared)
|   +-- ServiceAccount + imagePullSecrets
|   +-- EventBus (JetStream)
|   +-- ScmProvider (GitHub App auth)
|   +-- GitRepository (gitops repo ref)
|   +-- PromotionStrategy (env chain + gates)
|   +-- ArgoCDCommitStatus (ArgoCD health)
+-- Per-environment (always generated)
|   +-- ArgoCD Application (path: ".", syncs from env branch root)
|   +-- Git branches (env + env-next)
+-- Per-gate (one set per gated environment)
|   +-- Role + RoleBinding
|   +-- WorkflowTemplate (dynamic checks DAG)
|   +-- EventSource (GitHub webhook)
|   +-- Sensor (event -> workflow trigger)
|   +-- Service (webhook endpoint)
|   +-- Ingress (TLS termination)
|   +-- BranchProtection
+-- Repo workflow (1)
    +-- driveby-deploy.yml (manual trigger: deploy manifests + image tag to env-next)
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
    defaultBranch: main
    manifestsPath: manifests

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

### Manifest-Agnostic Design

The XSDLC is **manifest-agnostic** — it does NOT manage environment-specific configuration (replicas, resources, env vars). Engineers own their Kubernetes manifests under the `manifests/` directory in the application repo (configured via `repository.manifestsPath`). The XSDLC only manages **quality gates** and **promotion flow**. There are no per-environment directories; ArgoCD applications sync from the relevant environment branch, all pointing at the same `manifests/` path.

### Two-Tier Configuration

All cluster-level configuration (ingress domains, JetStream settings, image registries, secrets) comes from `values.yaml`, baked into the composition at Helm template time. Per-API overrides are set directly in the XSDLC spec. There are no EnvironmentConfigs.

| Tier | Source | Examples |
|---|---|---|
| Cluster defaults | `values.yaml` (`defaults.*`) | `baseDomain`, `clusterIssuer`, `webhookPort`, `imagePullSecret` |
| Per-API overrides | XSDLC spec fields | `apiConfig.port`, `validationDefaults.validationMode`, per-gate config |

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
