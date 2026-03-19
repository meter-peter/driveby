# Kubernetes Directory — Deployment and GitOps Infrastructure

## Structure

```
kubernetes/
  helm/
    driveby/            # Helm chart — installs full quality gate system
      Chart.yaml
      values.yaml
      .helmignore
      templates/
        crossplane/     # Crossplane resources (providers, functions, XRDs, compositions)
          provider-kubernetes.yaml     # Provider CR
          functions.yaml               # 4 Function CRs
          provider-config.yaml         # ProviderConfig (InjectedIdentity)
          xrd-template.yaml            # XQualityGateTemplate XRD
          xrd-instance.yaml            # XQualityGate XRD
          composition-template.yaml    # Template composition (RBAC + WorkflowTemplate)
          composition-instance.yaml    # Instance composition (EventBus + EventSource + Sensor + Ingress)
        argo-events/    # Argo Events (non-Crossplane fallback, gated behind !crossplane.enabled)
        argocd/         # ArgoCD AppProject
        secrets/        # Secret templates (api-auth, github-pat)
  examples/
    argo-workflows/     # Argo Workflow templates for validation pipelines
    argo-events/        # Argo Events triggers (webhook, git sensor)
    crossplane/         # Crossplane XQualityGate examples
      template-example.yaml      # XQualityGateTemplate CR example
      instance-example.yaml      # XQualityGate CR example
      environment-configs.yaml   # Template + Instance EnvironmentConfigs
    gitops-promoter/    # GitOps promotion logic (environment promotion on validation pass)
  README.md
```

> **Note**: The `manifests/` directory (raw YAML duplicates) was removed in v0.3.0 — all resources are now managed by the Helm chart + Crossplane compositions.

## Promotion Flow

Quality gates trigger on **promotion PRs in the gitops repo** (`perfect-api-gitops`), validate against the **source environment** (dev), and gate promotion to the **target environment** (staging):

```
Developer merges to perfect-api (source) → CI builds image → updates gitops/overlays/dev
    → ArgoCD auto-syncs dev
    → Developer opens promotion PR on gitops repo (dev → staging)
    → Webhook triggers quality gate → validates API in dev
    → Pass → commit status "success" on gitops PR → merge allowed → ArgoCD syncs staging
    → Fail → commit status "failure" → PR blocked
```

## Cluster State

### Namespaces
| Namespace | Purpose | Managed By |
|-----------|---------|------------|
| `driveby` | Workflow infrastructure (WorkflowTemplates, ServiceAccount, Sensors) | Helm chart + Crossplane |
| `perfect-api-dev` | Dev environment for perfect-api (autoSync) | ArgoCD |
| `perfect-api-staging` | Staging environment for perfect-api (autoSync) | ArgoCD |
| `perfect-api-prod` | Production environment for perfect-api (manual sync) | ArgoCD |

### ArgoCD Resources
| Resource | Namespace | Details |
|----------|-----------|---------|
| AppProject `perfect-api` | argocd | Sources: `meter-peter/perfect-api-gitops`, Destinations: dev + staging + prod |
| Application `perfect-api-dev` | argocd | Path: `overlays/dev`, autoSync + selfHeal + CreateNamespace |
| Application `perfect-api-staging` | argocd | Path: `overlays/staging`, autoSync + selfHeal |
| Application `perfect-api-prod` | argocd | Path: `overlays/prod`, manual sync |

### Argo Events (Crossplane-managed)
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| EventBus | `default` | driveby | 1-replica JetStream (NATS v2.10.10) |
| EventSource | `perfect-api-staging-promotion-eventsource` | driveby | GitHub webhook on port 12000, endpoint `/perfect-api-staging-promotion-pr-validation` |
| Sensor | `perfect-api-staging-promotion-sensor` | driveby | Triggers `driveby-staging-promotion` WorkflowTemplate on PR open/reopen/sync |
| Ingress | `perfect-api-staging-promotion-webhook-ingress` | driveby | `perfect-api-staging-promotion-webhook.private.novelcore.org` → EventSource svc |

GitHub webhook (ID: 601590418) is configured on `meter-peter/perfect-api-gitops` to POST `pull_request` events to `https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation`.

### Secrets
| Secret | Namespace(s) | Keys |
|--------|-------------|------|
| `ghcr-creds` | perfect-api-dev, perfect-api-staging, perfect-api-prod, driveby | Docker registry auth for ghcr.io |
| `api-auth` | perfect-api-dev, perfect-api-staging, perfect-api-prod | `api-key`, `api-key-header` |
| `driveby-api-auth` | driveby | `api-key`, `api-key-header` |
| `github-pat` | driveby | `token` (GitHub PAT for commit status + PR comments) |

### External Repos
| Repo | Purpose |
|------|---------|
| `meter-peter/perfect-api` | App code (FastAPI), CI builds `ghcr.io/meter-peter/perfect-api:latest` |
| `meter-peter/perfect-api-gitops` | Kustomize base + overlays (dev/staging/prod), synced by ArgoCD. Webhook triggers quality gate on promotion PRs. |

## Helm Chart (`helm/driveby/`)
Production Helm chart (v0.3.0) that installs the full DriveBy quality gate system:
- **Crossplane providers**: `provider-kubernetes` v0.14.1
- **Crossplane functions**: `function-go-templating`, `function-auto-ready`, `function-environment-configs` (sequencer removed — not needed for flat resource sets)
- **XRDs**: `xqualitygatetemplates.driveby.io`, `xqualitygates.driveby.io` (both v1alpha1)
- **Compositions**: Template (generates RBAC + WorkflowTemplate), Instance (generates EventBus + EventSource + Sensor + Ingress)
- **ProviderConfig**: `kubernetes-provider` with InjectedIdentity
- Raw Argo templates gated behind `not .Values.crossplane.enabled` (legacy fallback)

Install:
```bash
# Pre-install provider-kubernetes (CRDs must exist before helm install)
kubectl apply -f <provider-kubernetes-cr>
kubectl wait --for=condition=Healthy provider/provider-kubernetes --timeout=120s
# Then helm install
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set crossplane.enabled=true \
  --set github.pat=$(gh auth token)
```

## WorkflowTemplate DAG (Promotion Pipeline)

```
set-pending-status → wait-for-source-ready → validate-source → functional-test-source
                                                                      ↓
                                                            report-success + comment-pr
```

- Validates against the **source environment** (dev), not staging
- `wait-for-source-ready`: health-check loop (max 2min) against `http://{service-name}.{source-namespace}:{port}{openapi-endpoint}`
- Commit status descriptions: "Validating dev environment..." / "Dev validation passed, safe to promote" / "Dev validation failed, promotion blocked"

## Configurability Model
Compositions use a three-tier variable resolution pattern:
1. **values.yaml** (`defaults.*`) — chart-level defaults, baked at `helm template` time
2. **EnvironmentConfig** — cluster-level overrides, resolved at Crossplane composition runtime via `.environment.*`
3. **XRD spec fields** — per-CR overrides, resolved at composition runtime via `$xr.spec.*`

Resolution order in go-templates: `$xr.spec.X | default ($env.Y | default "<helm-baked-default>")`

## Target Cluster
- **Cluster**: `private.novelcore.org` (via `access.kubecore.eu`)
- **ArgoCD**: Already installed — do NOT reinstall
- **Strategy**: `driveby` namespace for workflows, `perfect-api-{dev,staging,prod}` for app environments

## Crossplane Workflow
When writing Crossplane resources:
1. Use `context7 resolve-library-id` first for provider docs
2. Use `context7 query-docs` for XRD schemas, Composition patterns
3. Only fall back to web search if context7 lacks the information

### Template Escaping
Argo Workflow parameter references (`{{workflow.parameters.xxx}}`) inside Crossplane go-templates need triple escaping:
- Helm layer: `{{` `` `{{ "{{" }}workflow.parameters.xxx{{ "}}" }}` `` `}}`
- Renders in Crossplane go-template as: `{{ "{{" }}workflow.parameters.xxx{{ "}}" }}`
- Crossplane renders as literal: `{{workflow.parameters.xxx}}` for Argo

## Thesis Mapping
- **Chapter 4 (Architecture)**: Namespace separation, ArgoCD AppProject isolation
- **Chapter 5 (End-to-End Workflow)**: The full GitOps pipeline from event trigger through validation to environment promotion
- The Argo Workflows + Events + Crossplane combination demonstrates DDT in a real GitOps context
