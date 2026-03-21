# Novelcore Perfect-API — XSDLC v3.0.0 Two-Repo Pipeline

Real-world showcase: one XSDLC CR turns `novelcore/perfect-api` into a fully automated quality-gated delivery pipeline with a dedicated gitops repo.

## How It Works

XSDLC owns **delivery**, not CI. The developer updates dry manifests on the `main` branch of the auto-created gitops repo (`dry/base/` or `dry/overlays/<env>/`). The ArgoCD Source Hydrator renders each environment's Kustomize overlay into hydrated manifests on `-next` branches. The Promoter then drives the promotion pipeline through quality gates.

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
    │
    ├──► Hydrator renders dry/overlays/dev → environment/dev-next:manifests/ + hydrator.metadata
    │    └── Promoter auto-merges dev-next → dev (no gate) → ArgoCD syncs dev
    │
    ├──► Hydrator renders dry/overlays/staging → environment/staging-next:manifests/ + hydrator.metadata
    │    └── Promoter PRs staging-next → staging → Gate 1 fires
    │    └── validate-only (test-ready) + functional-test against dev
    │    └── Auto-merge on success → ArgoCD syncs staging
    │
    └──► Hydrator renders dry/overlays/prod → environment/prod-next:manifests/ + hydrator.metadata
         └── Promoter PRs prod-next → prod → Gate 2 fires
         └── validate-only (strict) + load-test against staging
         └── Manual approval required (autoMerge: false)
```

## Repository Layout

```
perfect-api/           (software repo — untouched by XSDLC)
  src/                    # developer's code
  Dockerfile              # developer's build

perfect-api-gitops/    (gitops repo — auto-created by XSDLC)
  main branch:
    dry/
      base/
        kustomization.yaml    # references deployment.yaml + service.yaml
        deployment.yaml        # boilerplate deployment
        service.yaml           # boilerplate service
      overlays/
        dev/
          kustomization.yaml  # resources: [../../base], namespace: perfect-api-dev
        staging/
          kustomization.yaml  # resources: [../../base], namespace: perfect-api-staging
        prod/
          kustomization.yaml  # resources: [../../base], namespace: perfect-api-prod

  environment/<env>-next branches:    (written by ArgoCD hydrator)
    manifests/
      deployment.yaml
      service.yaml
    hydrator.metadata                 # {"drySha": "abc123..."}

  environment/<env> branches:         (merged by Promoter from <env>-next)
    manifests/ + hydrator.metadata
```

Dry manifests on `main` are the single source of truth. The ArgoCD hydrator renders per-environment overlays into hydrated output on `-next` branches. Gates control **when** content promotes, not **what** gets deployed.

## The Single CR

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
    name: perfect-api-gitops

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
            validationConfig:
              validationMode: strict
          - type: load-test
            loadTestConfig:
              concurrentUsers: 50
              testDuration: "2m"
              maxLatencyP95: "200ms"
              minSuccessRate: 0.995
```

**~40 lines** — generates the entire delivery pipeline.

## What Gets Created

From this single CR, the XSDLC composition generates:

**App-level (shared):**
- GitOps Repository `perfect-api-gitops` (auto-created via `provider-upjet-github`)
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `perfect-api-github` (GitHub App auth)
- GitRepository `perfect-api-gitops` (promoter repo reference — points to gitops repo)
- PromotionStrategy `perfect-api-promotion` (3-env chain with commit status gates)
- ArgoCDCommitStatus `perfect-api-argocd-health` (when `argocdHealthCheck: true`)

**Per-environment (x3):**
- ArgoCD Application (`perfect-api-dev`, `perfect-api-staging`, `perfect-api-prod`) — all autoSync; dev uses `sourceHydrator` (drySource=`main:dry/overlays/dev`, hydrateTo=`environment/dev-next`), staging and prod use regular `source` pointing to `manifests/`
- Git branches in gitops repo (`environment/dev` + `environment/dev-next`, etc.)

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding (RBAC for workflow execution)
- WorkflowTemplate (multi-check DAG)
- EventSource (GitHub webhook receiver — watches gitops repo)
- Sensor (PR event filter → workflow trigger)
- Service + Ingress (TLS webhook endpoint)
- BranchProtection (required status checks — on gitops repo)

**Total: ~33 resources** for a 3-env, 2-gate setup.

## Everything Is Derived

| Value | Source |
|-------|--------|
| `appName` | `metadata.name` |
| `gitopsRepoName` | `gitopsRepository.name` or `<repository.name>-gitops` |
| `branch` | `environment/<envName>` |
| `sourceNamespace` | `<appName>-<previousEnvName>` |
| `targetNamespace` | `<appName>-<envName>` |
| `commitStatusKey` | `<envName>-gate` |
| Cluster config | `values.yaml` (baked at helm-template time) |

## Prerequisites

1. **DriveBy Helm chart v3.0.0** installed with `crossplane.enabled=true` and `githubProvider.enabled=true`
2. **Secrets** in `driveby` namespace: `github-app-credentials`, `driveby-api-auth`, `ghcr-creds`
3. **GitHub App** with webhook + commit status + repo create permissions on `novelcore` org
4. **Dockerfile** in the software repo for building the application image

That's it. GitOps repo, branches, ArgoCD apps, promoter, quality gates, branch protection — all created by the CR.

## Apply

```bash
kubectl apply -f xsdlc-perfect-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# GitOps repo created
gh repo view novelcore/perfect-api-gitops

# Branches in gitops repo
gh api repos/novelcore/perfect-api-gitops/branches --jq '.[].name'

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby

# Event infrastructure
kubectl get eventbus,eventsources,sensors -n driveby
kubectl get ingress -n driveby

# Promoter resources
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy -n driveby

# ArgoCD apps
kubectl get applications -n argocd | grep perfect-api
```

## Uninstall

```bash
kubectl delete xsdlc perfect-api -n driveby
```
