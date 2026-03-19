# Quality Gates in the Software Development Lifecycle

How DriveBy implements documentation-driven quality gates using Crossplane and Argo Events.

## Quality Gate Concept

A **quality gate** is an automated checkpoint in the SDLC that validates software against defined criteria before allowing progression. In DriveBy, quality gates enforce DDT principles against API specifications at deployment boundaries.

```
Developer → PR → Webhook → EventSource → Sensor → Workflow → Commit Status → Merge/Block
```

## End-to-End Flow

### 1. Pull Request Created

A developer opens a PR against the API repository (e.g., `meter-peter/perfect-api`).

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

### 5. Validation Workflow (7-Step DAG)

```
set-pending-status
        │
    db-sync
        │
  spec-sync-check
        │
  validate-staging     ← DriveBy validate-only (DDT principles P001-P008)
        │
functional-test-staging ← DriveBy function-only (P006)
       ╱ ╲
report-success  comment-pr
```

### 6. Commit Status as Gate Signal

The workflow reports a GitHub commit status:
- **Context**: `driveby/staging-promotion`
- **State**: `pending` → `success` or `failure`
- **Description**: Human-readable validation result

### 7. GitOps Promoter Integration

The commit status serves as a gate signal for the GitOps Promoter's `PromotionStrategy`:

```yaml
apiVersion: promoter.argoproj.io/v1alpha1
kind: PromotionStrategy
spec:
  environments:
    - branch: env/staging
      activeCommitStatuses:
        - key: driveby/staging-promotion
    - branch: env/prod
      # Only promote to prod after staging gate passes
```

## Crossplane XQualityGate: Declarative Approach

Instead of manually deploying EventBus + EventSource + Sensor + Ingress + WorkflowTemplate + RBAC, DriveBy uses two Crossplane XRDs:

### XQualityGateTemplate (Shared Infrastructure)

Declares **what validation looks like** for a project:
- WorkflowTemplate with the validation DAG
- ServiceAccount + RBAC for workflow execution

One template serves multiple API instances.

### XQualityGate (Per-API Instance)

Declares **which API to validate** and **how to trigger it**:
- EventBus for event transport
- EventSource for GitHub webhooks
- Sensor for event filtering and workflow triggering
- Ingress for webhook endpoint exposure

Each API gets its own instance, all sharing the same template.

### Resource Flow

```
XQualityGateTemplate (1 per project)
├── ServiceAccount + imagePullSecrets
├── Role (workflows, secrets, pods)
├── RoleBinding
└── WorkflowTemplate (7-step DAG)

XQualityGate (1 per API)
├── EventBus (JetStream)
├── EventSource (GitHub webhook)
├── Sensor (event → workflow trigger)
├── Service (webhook endpoint)
└── Ingress (TLS termination)
```

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
| `test-only` | P006 | Functional testing only |

Configure per-template in `XQualityGateTemplate.spec.validationConfig.validationMode`.
