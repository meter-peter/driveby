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
    novelcore-perfect-api/  # Single-file XSDLC showcase (two-repo model)
    novelcore-bad-docs-api/ # XSDLC: fails P002+P003 at staging validate-only
    novelcore-no-auth-api/  # XSDLC: fails P005 at staging validate-only
    novelcore-slow-api/     # XSDLC: passes staging, fails prod load-test
    novelcore-broken-api/   # XSDLC: passes validate-only, fails functional-test
    gitops-promoter/    # GitOps Promoter CRD reference
  README.md
```

## Promotion Flow

Quality gates use a **multi-check model** — each gate defines an ordered list of checks.

### Repo Layout (Two-Repo Model with Source Hydrator)
XSDLC v3.0.0 uses a **two-repo model** with ArgoCD Source Hydrator: the software repo is untouched by XSDLC, and a separate gitops repo (auto-created by XSDLC) holds dry manifests on `main` and hydrated output on per-environment branches. Engineers own all manifest config (replicas, resources, env vars). The XSDLC only manages **promotion** and **quality gates** — it does not patch or rewrite manifests.

```
perfect-api/           (software repo — untouched by XSDLC)
  src/
  Dockerfile

perfect-api-gitops/    (gitops repo — auto-created by XSDLC)
  main branch:
    dry/
      base/
        kustomization.yaml       # references deployment.yaml + service.yaml
        deployment.yaml           # boilerplate deployment
        service.yaml              # boilerplate service
      overlays/
        dev/
          kustomization.yaml     # resources: [../../base], namespace: perfect-api-dev
        staging/
          kustomization.yaml     # resources: [../../base], namespace: perfect-api-staging
        prod/
          kustomization.yaml     # resources: [../../base], namespace: perfect-api-prod

  environment/dev-next:           # written by ArgoCD hydrator (from dry/overlays/dev)
    manifests/
      deployment.yaml
      service.yaml
    hydrator.metadata             # {"drySha": "abc123..."}

  environment/dev:                # merged by Promoter from dev-next
    manifests/ + hydrator.metadata

  environment/staging-next:       # written by ArgoCD hydrator (from dry/overlays/staging)
    manifests/ + hydrator.metadata

  environment/staging:            # merged by Promoter from staging-next (after gate passes)
    manifests/ + hydrator.metadata

  environment/prod-next:          # written by ArgoCD hydrator (from dry/overlays/prod)
    manifests/ + hydrator.metadata

  environment/prod:               # merged by Promoter from prod-next (after gate passes)
    manifests/ + hydrator.metadata
```

### Source Hydrator Flow

ALL environments use ArgoCD Source Hydrator. Each ArgoCD Application points its `drySource` to `main:dry/overlays/<env>` and `hydrateTo` to `environment/<env>-next`. The hydrator independently renders each overlay and writes the hydrated output (with `hydrator.metadata`) to that environment's `-next` branch. There is no linear propagation of hydrated content between environments — each env hydrates independently from `main`. The Promoter only handles merging `-next` into the active branch (with optional quality gates).

### No Generated Workflows

XSDLC v3.0.0 generates **no workflows** in the software repo. Developers update dry manifests on `main` in the gitops repo (either `dry/base/` for all-env changes or `dry/overlays/<env>/` for per-env changes). The hydrator independently renders each overlay to its corresponding `-next` branch, and the Promoter drives the promotion pipeline from there.

### End-to-End Flow

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
  → ArgoCD hydrator renders each overlay → writes to environment/<env>-next:manifests/ + hydrator.metadata
  → For ungated envs (dev): Promoter auto-merges dev-next → dev → ArgoCD syncs from manifests/
  → For gated envs (staging): Promoter PRs staging-next → staging → Webhook fires → DriveBy validates → auto-merge on success
  → For gated envs (prod): Promoter PRs prod-next → prod → Webhook fires → DriveBy validates → autoMerge: false → manual approval
```

## XSDLC XRD (v3.0.0)

### Required Fields (2)
| Field | Description |
|-------|-------------|
| `gitopsRepository.{owner,name}` | GitHub org/user and gitops repo name (e.g., novelcore/perfect-api-gitops) |
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
| `apiConfig.image` | `ghcr.io/<owner>/<appName>:latest` | Container image for the API |
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
| `perfect-api-prod` | Production environment for perfect-api (autoSync — `autoMerge: false` only controls Promoter PR merge) | ArgoCD |

### ArgoCD Resources (Crossplane-managed via XSDLC — always generated)
| Resource | Namespace | Details |
|----------|-----------|---------|
| AppProject `driveby` | argocd | Sources: `*`, Destinations: `*` namespace (Helm chart) |
| Application `perfect-api-dev` | argocd | sourceHydrator: drySource=main:dry/overlays/dev, syncSource=env/dev:manifests/, hydrateTo=env/dev-next; autoSync + selfHeal (XSDLC) |
| Application `perfect-api-staging` | argocd | sourceHydrator: drySource=main:dry/overlays/staging, syncSource=env/staging:manifests/, hydrateTo=env/staging-next; autoSync + selfHeal (XSDLC) |
| Application `perfect-api-prod` | argocd | sourceHydrator: drySource=main:dry/overlays/prod, syncSource=env/prod:manifests/, hydrateTo=env/prod-next; autoSync + selfHeal (XSDLC) — `autoMerge: false` only controls Promoter PR merge, not ArgoCD sync |

### Argo Events (Crossplane-managed via XSDLC)
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| EventBus | `default` | driveby | 1-replica JetStream (NATS v2.10.10) |
| EventSource | `perfect-api-staging-gate-eventsource` | driveby | GitHub webhook on port 12000 |
| EventSource | `perfect-api-prod-gate-eventsource` | driveby | GitHub webhook on port 12000 |
| Sensor | `perfect-api-staging-gate-sensor` | driveby | Triggers workflow on PR to `environment/staging` in gitops repo |
| Sensor | `perfect-api-prod-gate-sensor` | driveby | Triggers workflow on PR to `environment/prod` in gitops repo |
| Ingress | `perfect-api-staging-gate-webhook-ingress` | driveby | TLS webhook endpoint |
| Ingress | `perfect-api-prod-gate-webhook-ingress` | driveby | TLS webhook endpoint |

### Secrets
| Secret | Namespace(s) | Keys | Managed By |
|--------|-------------|------|------------|
| `ghcr-creds` | driveby | Docker registry auth for ghcr.io | Helm chart (`secrets.ghcr.enabled`) |
| `ghcr-creds` | `<app>-{dev,staging,prod}` | Docker registry auth (copied from driveby ns) | **XSDLC composition** (auto) |
| `<app>-push-secret` | argocd | ArgoCD repository-write for gitops repo | **XSDLC composition** (auto) |
| `driveby-api-auth` | driveby | `api-key`, `api-key-header` | Helm chart |
| `github-app-credentials` | driveby | `githubAppID`, `githubInstallationID`, `githubAppPrivateKey` | Helm chart |

### GitOps Promoter Resources (Crossplane-managed via XSDLC)
| Resource | Name | Namespace | Details |
|----------|------|-----------|---------|
| ScmProvider | `perfect-api-github` | driveby | GitHub App auth for promoter SCM access |
| GitRepository | `perfect-api` | driveby | Points promoter to the gitops repository via ScmProvider |
| PromotionStrategy | `perfect-api-promotion` | driveby | Environment chain: dev → staging → prod, with commit status gates |
| ArgoCDCommitStatus | `perfect-api-argocd-health` | driveby | Aggregates ArgoCD app health into CommitStatus |
| ChangeTransferPolicy | (auto-created) | driveby | Auto-created by PromotionStrategy controller per environment |

### External Repos
| Repo | Purpose |
|------|---------|
| `novelcore/perfect-api` | Software repo — source code + Dockerfile (untouched by XSDLC) |
| `novelcore/perfect-api-gitops` | GitOps repo — auto-created by XSDLC, holds manifests on per-environment branches |

## Helm Chart (`helm/driveby/`)
Production Helm chart (v3.0.0) that installs the full DriveBy quality gate system:
- **Crossplane providers**: `provider-kubernetes` v0.14.1
- **Crossplane functions**: `function-go-templating`, `function-auto-ready`
- **XRD**: `xsdlcs.driveby.io` (v1alpha1) — single CR for full promotion pipeline
- **Composition**: XSDLC composition generates all resources (RBAC, WorkflowTemplates, EventBus, EventSource, Sensor, Ingress, Promoter, BranchProtection, ArgoCD push secrets, per-namespace ghcr-creds)
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

Exit handler (on failure):
  report-failure-github + update-commitstatus-failure + comment-pr-failure (parallel)
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

### Check Ordering & Auto-Injection
- **Auto-injection**: If a gate has runtime checks (`functional-test`/`load-test`) but no `validate-only`, a `validate-only` check is auto-injected using `validationDefaults.validationMode`
- **Sorting**: `validate-only` checks always run before runtime checks regardless of the order in `gate.checks[]`
- **Failure reporting**: Exit handler posts the DriveBy report as a PR comment on failure (`comment-pr-failure`), so developers see which principles failed — not just a generic "checks failed" status

### Notes
- All pipelines validate against the **source environment** (previous env), not target
- Step templates (`driveby-validate`, `driveby-functional`, `driveby-loadtest`) accept input parameters from the DAG
- Shared `reports` PVC (64Mi) mounts at `/tmp/reports` across all DriveBy steps — `comment-pr` reads saved reports via `github-comment` instead of re-running tests
- Environment-specific config (replicas, resources, env vars) is **engineer-owned** in the gitops repo's per-environment overlays (`dry/overlays/<env>/` on `main`)

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
