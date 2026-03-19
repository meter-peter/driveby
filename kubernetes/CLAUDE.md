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
Production Helm chart that installs the full DriveBy quality gate system:
- **Crossplane providers**: `provider-kubernetes` v0.14.1
- **Crossplane functions**: `function-go-templating`, `function-auto-ready`, `function-sequencer`, `function-environment-configs`
- **XRDs**: `xqualitygatetemplates.driveby.io`, `xqualitygates.driveby.io` (both namespace-scoped, v1alpha1)
- **Compositions**: Template (generates RBAC + WorkflowTemplate), Instance (generates EventBus + EventSource + Sensor + Ingress)
- **ProviderConfig**: `kubernetes-provider` with InjectedIdentity
- Plus existing: Argo Workflows templates, Argo Events, ArgoCD resources, RBAC, secrets

Install: `helm install driveby ./kubernetes/helm/driveby/ --set crossplane.enabled=true`
See `docs/deployment-guide.md` for full installation guide.

## Example Resources (`examples/`)

| Directory | Purpose | Thesis Mapping |
|-----------|---------|---------------|
| `argo-workflows/` | Validation pipeline as Argo Workflow (spec fetch -> validate -> report -> promote) | Ch.5 end-to-end workflow |
| `argo-events/` | Event triggers: webhook on PR, git sensor on spec changes | Ch.5 event-driven architecture |
| `crossplane/` | XQualityGateTemplate + XQualityGate examples, EnvironmentConfigs | Ch.5 Crossplane XSDLC |
| `gitops-promoter/` | Promotion logic: advance environment on validation pass | Ch.5 feedback loop |

## Configurability Model
Compositions use a three-tier variable resolution pattern:
1. **values.yaml** (`defaults.*`) — chart-level defaults, baked at `helm template` time
2. **EnvironmentConfig** — cluster-level overrides, resolved at Crossplane composition runtime via `.environment.*`
3. **XRD spec fields** — per-CR overrides, resolved at composition runtime via `$xr.spec.*`

Resolution order in go-templates: `$xr.spec.X | default ($env.Y | default "<helm-baked-default>")`

All previously hardcoded values (ingress class, cert-manager issuer, secret names, container images, webhook ports, GitHub API URL, EventBus config) are now configurable through this model. See `docs/deployment-guide.md` for the full table of configurable knobs.

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
