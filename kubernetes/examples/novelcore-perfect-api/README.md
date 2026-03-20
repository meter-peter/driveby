# Novelcore Perfect-API — XSDLC v2.1.0 Single-Repo Pipeline

Real-world showcase: one XSDLC CR turns `novelcore/perfect-api` into a fully automated quality-gated delivery pipeline.

## How It Works

XSDLC owns **delivery**, not CI. This is the **BYOCI (Bring Your Own CI) model**: the developer's own CI builds the container image, then the `driveby-deploy.yml` workflow showcases what happens when that image gets deployed to an environment and goes through quality gates.

```
Developer triggers "DriveBy Deploy" (Actions → workflow_dispatch)
  Inputs: source_branch, environment, image_tag
    │
    ▼
driveby-deploy.yml: copies manifests/ → stamps image tag → pushes to environment/<env>-next
    │
    ├──► Promoter auto-merges dev-next → dev (no gate) → ArgoCD syncs dev
    │
    ├──► Promoter PRs staging-next → staging → Gate 1 fires
    │    └── validate-only (test-ready) + functional-test against dev
    │    └── Auto-merge on success → ArgoCD syncs staging
    │
    └──► Promoter PRs prod-next → prod → Gate 2 fires
         └── validate-only (strict) + load-test against staging
         └── Manual approval required (autoMerge: false)
```

## Repository Layout

```
perfect-api/
  src/                    # developer's code
  Dockerfile              # developer's build
  manifests/
    deployment.yaml       # K8s manifests — same for all envs
    service.yaml
```

Same manifests for all environments. Gates control **when** code promotes, not **what** gets deployed. Environment-specific config (replicas, resources) is set per-environment in the ArgoCD Application or via Kustomize in the env branches.

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
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `perfect-api-github` (GitHub App auth)
- GitRepository `perfect-api` (promoter repo reference)
- PromotionStrategy `perfect-api-promotion` (3-env chain with commit status gates)
- ArgoCDCommitStatus `perfect-api-argocd-health`

**Per-environment (x3):**
- ArgoCD Application (`perfect-api-dev`, `perfect-api-staging`, `perfect-api-prod`)
- Git branches (`environment/dev` + `environment/dev-next`, etc.)

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding (RBAC for workflow execution)
- WorkflowTemplate (multi-check DAG)
- EventSource (GitHub webhook receiver)
- Sensor (PR event filter → workflow trigger)
- Service + Ingress (TLS webhook endpoint)
- BranchProtection (required status checks)

**Repo workflow (x1):**
- `driveby-deploy.yml` — manual trigger: pick source branch, target environment, and image tag. Deploys manifests to chosen env-next.

**Total: ~34 resources** for a 3-env, 2-gate setup.

## Everything Is Derived

| Value | Source |
|-------|--------|
| `appName` | `metadata.name` |
| `branch` | `environment/<envName>` |
| `sourceNamespace` | `<appName>-<previousEnvName>` |
| `targetNamespace` | `<appName>-<envName>` |
| `commitStatusKey` | `<envName>-gate` |
| Cluster config | `values.yaml` (baked at helm-template time) |

## Prerequisites

1. **DriveBy Helm chart v2.1.0** installed with `crossplane.enabled=true` and `githubProvider.enabled=true`
2. **Secrets** in `driveby` namespace: `github-app-credentials`, `driveby-api-auth`, `ghcr-creds`
3. **GitHub App** with webhook + commit status permissions on `novelcore/perfect-api`
4. **Dockerfile** in the repo for building the application image

That's it. Branches, workflows, ArgoCD apps, promoter, quality gates, branch protection — all created by the CR.

## Apply

```bash
kubectl apply -f xsdlc-perfect-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby

# Event infrastructure
kubectl get eventbus,eventsources,sensors -n driveby
kubectl get ingress -n driveby

# Promoter resources
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy -n driveby

# ArgoCD apps
kubectl get applications -n argocd | grep perfect-api

# Repo workflows
gh api repos/novelcore/perfect-api/contents/.github/workflows --jq '.[].name'
# Expected: driveby-deploy.yml
```

## Uninstall

```bash
kubectl delete xsdlc perfect-api -n driveby
```
