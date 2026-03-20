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
          functions.yaml               # 2 Function CRs (go-templating, auto-ready)
          provider-config.yaml         # ProviderConfig (InjectedIdentity)
          github-provider-config.yaml  # GitHub ProviderConfig (for branch protection)
          xrd-sdlc.yaml               # XSDLC XRD — single CR for full promotion pipeline
          composition-sdlc.yaml       # XSDLC composition — generates all resources
        argocd/         # ArgoCD AppProject
        secrets/        # Secret templates (api-auth, github-app, ghcr-creds, github-provider-token)
  examples/
    novelcore-perfect-api/  # Single-file XSDLC showcase
    gitops-promoter/    # GitOps Promoter CRD reference (ScmProvider, GitRepository, PromotionStrategy, CommitStatus)
  README.md
```

## Promotion Flow

Quality gates use a **multi-check model** — each gate defines an ordered list of checks.

### Repo Layout (Single-Repo Model)
The application repository holds both source code and Kubernetes manifests on `main`. Engineers own all manifest config (replicas, resources, env vars). The XSDLC only manages **promotion** and **quality gates** — it does not patch or rewrite manifests.

```
perfect-api/
  src/                    # developer's code
  Dockerfile              # developer's build
  manifests/
    deployment.yaml       # same manifest for all envs (image tag managed by CI)
    service.yaml
```

### Generated Workflow (1 — BYOCI showcase)

- `driveby-deploy.yml` — Manual trigger (`workflow_dispatch`): pick source branch, target environment, and image tag. The developer's own CI builds the image; this workflow showcases the BYOCI model by deploying manifests + image tag to `environment/<env>-next`. Promoter takes over from there.

### End-to-End Flow

```
Developer triggers "DriveBy Deploy" workflow (Actions → workflow_dispatch)
  Inputs: source_branch (e.g. main), environment (e.g. dev), image_tag (e.g. v1.2.3)
    → Checks out environment/<env>-next, cleans it
    → Copies manifests/ from source branch → flattens to branch root
    → Stamps image tag in deployment.yaml → pushes to env-next
    → Promoter auto-merges dev-next → dev (no gate) → ArgoCD syncs dev

  Gate 1 — Staging Gate (checks: validate-only + functional-test):
    → Promoter auto-PRs environment/staging-next → environment/staging
    → Webhook triggers DriveBy workflow against dev
    → Workflow runs checks sequentially → creates CommitStatus CRD (phase: success)
    → Promoter auto-merges → ArgoCD syncs staging

  Gate 2 — Prod Gate (checks: validate-only + load-test):
    → Promoter auto-PRs environment/prod-next → environment/prod
    → Webhook triggers DriveBy workflow against staging
    → Workflow runs checks sequentially → creates CommitStatus CRD (phase: success)
    → Promoter does NOT auto-merge (autoMerge: false — manual approval required)
```

## XSDLC XRD (v2.1.0)

### Required Fields (2)
| Field | Description |
|-------|-------------|
| `repository.{owner,name}` | GitHub repository owner and name (e.g., novelcore/perfect-api) |
| `environments` | Ordered list of promotion environments (min 2) |

### Per-Environment Fields
| Field | Default | Description |
|-------|---------|-------------|
| `name` | (required) | Environment name |
| `branch` | `environment/<name>` | Git branch |
| `autoMerge` | `true` | Auto-merge promotion PRs |
| `namespaceOverride` | `<appName>-<name>` | Override target namespace |
| `gate.checks` | (required if gate set) | Ordered list of checks (each becomes a DAG step) |
| `gate.commitStatusKey` | `<envName>-gate` | Override commit status key |
| `gate.sourceNamespaceOverride` | `<appName>-<prevEnv>` | Override source namespace |

### Check Types (gate.checks[].type)
| Type | CLI Command | Description |
|------|-------------|-------------|
| `validate-only` | `validate-only --validation-mode <mode>` | Static validation (P001-P009 based on mode) |
| `functional-test` | `function-only` | Functional API testing (P006) |
| `load-test` | `load-only` | k6 load testing with configurable thresholds |

### Per-Check Config
| Field | Applies To | Description |
|-------|-----------|-------------|
| `validationConfig.validationMode` | `validate-only` | Validation mode override |
| `loadTestConfig.concurrentUsers` | `load-test` | Number of concurrent users (default: 10) |
| `loadTestConfig.testDuration` | `load-test` | Duration of load test (default: 5m) |
| `loadTestConfig.maxLatencyP95` | `load-test` | Maximum P95 latency (default: 500ms) |
| `loadTestConfig.minSuccessRate` | `load-test` | Minimum success rate 0-1 (default: 0.99) |

### Optional Top-Level Fields
| Field | Default | Description |
|-------|---------|-------------|
| `manifestsPath` | `manifests` | Path within the repository where Kubernetes manifests live |
| `apiConfig.serviceName` | `metadata.name` | K8s service name |
| `apiConfig.port` | `8000` | API port |
| `apiConfig.openapiEndpoint` | `/openapi.json` | OpenAPI spec path |
| `validationDefaults.drivebyImage` | `ghcr.io/meter-peter/driveby:latest` | DriveBy image |
| `validationDefaults.validationMode` | `strict` | Default validation mode |

### Derived Fields (all automatic)
| Field | Derivation |
|-------|-----------|
| `appName` | `metadata.name` |
| `branch` | `environment/<envName>` |
| `sourceNamespace` | `<appName>-<previousEnvName>` |
| `targetNamespace` | `<appName>-<envName>` |
| `commitStatusKey` | `<envName>-gate` |
| All cluster config | `values.yaml` (baked at helm-template time) |

## Cluster State

### Namespaces
| Namespace | Purpose | Managed By |
|-----------|---------|------------|
| `driveby` | Workflow infrastructure (WorkflowTemplates, ServiceAccount, Sensors, Promoter resources) | Helm chart + Crossplane |
| `perfect-api-dev` | Dev environment for perfect-api (autoSync) | ArgoCD |
| `perfect-api-staging` | Staging environment for perfect-api (autoSync) | ArgoCD |
| `perfect-api-prod` | Production environment for perfect-api (manual sync) | ArgoCD |

### ArgoCD Resources (Crossplane-managed via XSDLC — always generated)
| Resource | Namespace | Details |
|----------|-----------|---------|
| AppProject `driveby` | argocd | Sources: `*`, Destinations: `*` namespace (Helm chart) |
| Application `perfect-api-dev` | argocd | Path: `.` (env branch root), autoSync + selfHeal + CreateNamespace (XSDLC) |
| Application `perfect-api-staging` | argocd | Path: `.`, autoSync + selfHeal (XSDLC) |
| Application `perfect-api-prod` | argocd | Path: `.`, manual sync (autoMerge: false) (XSDLC) |

### Argo Events (Crossplane-managed via XSDLC)
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| EventBus | `default` | driveby | 1-replica JetStream (NATS v2.10.10) |
| EventSource | `perfect-api-staging-gate-eventsource` | driveby | GitHub webhook on port 12000 |
| EventSource | `perfect-api-prod-gate-eventsource` | driveby | GitHub webhook on port 12000 |
| Sensor | `perfect-api-staging-gate-sensor` | driveby | Triggers workflow on PR to `environment/staging` |
| Sensor | `perfect-api-prod-gate-sensor` | driveby | Triggers workflow on PR to `environment/prod` |
| Ingress | `perfect-api-staging-gate-webhook-ingress` | driveby | TLS webhook endpoint |
| Ingress | `perfect-api-prod-gate-webhook-ingress` | driveby | TLS webhook endpoint |

### Secrets
| Secret | Namespace(s) | Keys |
|--------|-------------|------|
| `ghcr-creds` | perfect-api-dev, perfect-api-staging, perfect-api-prod, driveby | Docker registry auth for ghcr.io |
| `api-auth` | perfect-api-dev, perfect-api-staging, perfect-api-prod | `api-key`, `api-key-header` |
| `driveby-api-auth` | driveby | `api-key`, `api-key-header` |
| `github-app-credentials` | driveby | `githubAppID`, `githubInstallationID`, `githubAppPrivateKey` |

### GitOps Promoter Resources (Crossplane-managed via XSDLC)
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| ScmProvider | `perfect-api-github` | driveby | GitHub App auth for promoter SCM access |
| GitRepository | `perfect-api` | driveby | Points promoter to the application repository via ScmProvider |
| PromotionStrategy | `perfect-api-promotion` | driveby | Environment chain: dev → staging → prod, with commit status gates |
| ArgoCDCommitStatus | `perfect-api-argocd-health` | driveby | Aggregates ArgoCD app health into CommitStatus |
| ChangeTransferPolicy | (auto-created) | driveby | Auto-created by PromotionStrategy controller per environment |

### External Repos
| Repo | Purpose |
|------|---------|
| `novelcore/perfect-api` | Source code + Kubernetes manifests (`manifests/`); workflows auto-generated by XSDLC |

## Helm Chart (`helm/driveby/`)
Production Helm chart (v2.1.0) that installs the full DriveBy quality gate system:
- **Crossplane providers**: `provider-kubernetes` v0.14.1
- **Crossplane functions**: `function-go-templating`, `function-auto-ready`
- **XRD**: `xsdlcs.driveby.io` (v1alpha1) — single CR for full promotion pipeline
- **Composition**: XSDLC composition generates all resources (RBAC, WorkflowTemplates, EventBus, EventSource, Sensor, Ingress, Promoter, BranchProtection)
- **ProviderConfig**: `kubernetes-provider` with InjectedIdentity
- **GitHub ProviderConfig**: `github-provider` for `provider-upjet-github` (branch protection rules)

Install:
```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set crossplane.enabled=true
```

## WorkflowTemplate DAGs (Multi-Check Gates)

Gates use a dynamic **checks** array — each check becomes a sequential DAG step:

```
set-pending → health-check → check-0-<type> → check-1-<type> → ... → check-N-<type>
                                                                            ↓
                              report-success + comment-pr + update-commitstatus-success
```

### Example: staging gate (validate-only + functional-test)
```
set-pending → health-check → check-0-validate-only → check-1-functional-test
                                                             ↓
                              report-success + comment-pr + update-commitstatus-success
```

### Example: prod gate (validate-only + load-test)
```
set-pending → health-check → check-0-validate-only → check-1-load-test
                                                            ↓
                              report-success + comment-pr + update-commitstatus-success
```

- All pipelines validate against the **source environment** (previous env), not target
- Step templates (`driveby-validate`, `driveby-functional`, `driveby-loadtest`) accept input parameters from the DAG
- Shared `reports` PVC (64Mi) mounts at `/tmp/reports` across all DriveBy steps — `comment-pr` reads saved reports via `github-comment` instead of re-running tests
- **v2.0.0 breaking change**: `softwareRepository`, `DISPATCH_PAT`, and `GITOPS_PAT` removed — single-repo model requires only one GitHub App credential
- Environment-specific config (replicas, resources, env vars) is **engineer-owned** in the repository's `manifests/` directory

## Configurability Model
Compositions use a **two-tier** variable resolution pattern:
1. **values.yaml** (`defaults.*`) — chart-level defaults, baked at `helm template` time
2. **XRD spec fields** — per-CR overrides, resolved at composition runtime

EnvironmentConfigs have been removed. All cluster config lives in `values.yaml`.

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
