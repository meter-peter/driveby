# Kubernetes Directory — Deployment and GitOps Infrastructure

## Structure

```
kubernetes/
  manifests/            # Raw YAML for workflow infra (applied directly to cluster)
    rbac.yaml           # ServiceAccount + Role + RoleBinding for driveby namespace
    workflow-templates.yaml  # WorkflowTemplates: driveby-validate, driveby-full-pipeline
    staging-promotion.yaml   # WorkflowTemplate: driveby-staging-promotion (DAG pipeline)
    eventbus.yaml       # JetStream EventBus (3-replica NATS cluster)
    eventsource.yaml    # GitHub webhook EventSource for perfect-api PRs
    sensor.yaml         # Sensor: PR events → driveby-staging-promotion workflow
    webhook-ingress.yaml # Service + Traefik Ingress for webhook endpoint
  helm/
    driveby/            # Helm chart (reference templates, not actively deployed)
      Chart.yaml
      values.yaml
      templates/
  examples/
    argo-workflows/     # Argo Workflow templates for validation pipelines
    argo-events/        # Argo Events triggers (webhook, git sensor)
    crossplane/         # Crossplane compositions for infrastructure provisioning
    gitops-promoter/    # GitOps promotion logic (environment promotion on validation pass)
  README.md
```

## Raw Manifests (`manifests/`)
Production-deployed workflow infrastructure in the `driveby` namespace:
- **rbac.yaml** — ServiceAccount `driveby`, Role (workflows, secrets, pods access), RoleBinding
- **workflow-templates.yaml** — Reusable `driveby-validate` and `driveby-full-pipeline` templates
- **staging-promotion.yaml** — Full DAG pipeline: set-pending → db-sync → spec-check → validate → functional-test → report-success → comment-pr
- **eventbus.yaml** — JetStream EventBus (`default`), 3-replica NATS cluster for event transport
- **eventsource.yaml** — GitHub webhook EventSource (`driveby-github`), listens on port 12000 at `/github/driveby` for `pull_request` and `push` events on `meter-peter/perfect-api`
- **sensor.yaml** — Sensor (`driveby-sensor`), filters PR `opened`/`reopened`/`synchronize` actions and submits `driveby-staging-promotion` workflow with PR number, head SHA, owner, repo as parameters
- **webhook-ingress.yaml** — ClusterIP Service + Traefik Ingress exposing the EventSource at `https://driveby-webhook.private.novelcore.org`

Apply with: `kubectl apply -f kubernetes/manifests/`
Apply order: eventbus.yaml first (NATS cluster must be ready before EventSource/Sensor can connect).

## Cluster State

### Namespaces
| Namespace | Purpose | Managed By |
|-----------|---------|------------|
| `driveby` | Workflow infrastructure (WorkflowTemplates, ServiceAccount, Sensors) | Manual kubectl apply |
| `perfect-api-staging` | Staging environment for perfect-api (autoSync) | ArgoCD |
| `perfect-api-prod` | Production environment for perfect-api (manual sync) | ArgoCD |

### ArgoCD Resources
| Resource | Namespace | Details |
|----------|-----------|---------|
| AppProject `perfect-api` | argocd | Sources: `meter-peter/perfect-api-gitops`, Destinations: staging + prod |
| Application `perfect-api-staging` | argocd | Path: `overlays/staging`, autoSync + selfHeal |
| Application `perfect-api-prod` | argocd | Path: `overlays/prod`, manual sync |

### Argo Events
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| EventBus | `default` | driveby | 3-replica JetStream (NATS v2.10.10) |
| EventSource | `driveby-github` | driveby | GitHub webhook on port 12000, endpoint `/github/driveby` |
| Sensor | `driveby-sensor` | driveby | Triggers `driveby-staging-promotion` on PR open/reopen/sync |
| Ingress | `driveby-webhook` | driveby | `driveby-webhook.private.novelcore.org` → EventSource svc |

GitHub webhook (ID: 601488783) is configured on `meter-peter/perfect-api` to POST `pull_request` events to `https://driveby-webhook.private.novelcore.org/github/driveby`.

### Secrets
| Secret | Namespace(s) | Keys |
|--------|-------------|------|
| `ghcr-creds` | perfect-api-staging, perfect-api-prod | Docker registry auth for ghcr.io |
| `api-auth` | perfect-api-staging, perfect-api-prod | `api-key`, `api-key-header` |
| `driveby-api-auth` | driveby | `api-key`, `api-key-header` |
| `github-pat` | driveby | `token` (GitHub PAT for commit status + PR comments) |
| `ghcr-creds` | driveby | Docker registry auth for ghcr.io |

### External Repos
| Repo | Purpose |
|------|---------|
| `meter-peter/perfect-api` | App code (FastAPI), CI builds `ghcr.io/meter-peter/perfect-api:latest` |
| `meter-peter/perfect-api-gitops` | Kustomize base + overlays (staging/prod), synced by ArgoCD |

## Helm Chart (`helm/driveby/`)
Reference Helm chart with templated versions of the workflow infrastructure. Not actively deployed — the raw manifests in `manifests/` are the source of truth for the cluster.

## Example Resources (`examples/`)

| Directory | Purpose | Thesis Mapping |
|-----------|---------|---------------|
| `argo-workflows/` | Validation pipeline as Argo Workflow (spec fetch -> validate -> report -> promote) | Ch.5 end-to-end workflow |
| `argo-events/` | Event triggers: webhook on PR, git sensor on spec changes | Ch.5 event-driven architecture |
| `crossplane/` | Infrastructure-as-code compositions for API environments | Ch.5 Crossplane integration |
| `gitops-promoter/` | Promotion logic: advance environment on validation pass | Ch.5 feedback loop |

## Target Cluster
- **Cluster**: `private.novelcore.org` (via `access.kubecore.eu`)
- **ArgoCD**: Already installed — do NOT reinstall
- **Strategy**: `driveby` namespace for workflows, `perfect-api-{staging,prod}` for app environments
- Reference kubecore-operator patterns at `/home/meter-peter/development/novelcore/kubecore-operator/compositions/`

## Crossplane Workflow
When writing Crossplane resources:
1. Use `context7 resolve-library-id` first for provider docs
2. Use `context7 query-docs` for XRD schemas, Composition patterns
3. Only fall back to web search if context7 lacks the information

## Thesis Mapping
- **Chapter 4 (Architecture)**: Namespace separation, ArgoCD AppProject isolation
- **Chapter 5 (End-to-End Workflow)**: The full GitOps pipeline from event trigger through validation to environment promotion
- The Argo Workflows + Events + Crossplane + Promoter combination demonstrates DDT in a real GitOps context
