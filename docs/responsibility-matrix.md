# DriveBy Responsibility Matrix

Ownership map for every resource in the DriveBy quality gate system. Use this to understand what DriveBy automates, what clients must provide, and where the boundary lies.

## Lifecycle Overview

```
┌──────────────────────────────────────────────────────────────────────────┐
│                        CLIENT RESPONSIBILITY                             │
│                                                                          │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌───────────────┐  │
│  │  Crossplane  │  │   ArgoCD    │  │Argo Workflows│  │  Argo Events  │  │
│  │  (1.14+)     │  │   (2.8+)    │  │   (3.5+)     │  │   (1.9+)      │  │
│  └─────────────┘  └─────────────┘  └─────────────┘  └───────────────┘  │
│  ┌─────────────┐  ┌─────────────┐  ┌───────────────────────────────┐   │
│  │  Traefik     │  │cert-manager │  │  GitOps Promoter (0.1+)      │   │
│  │  (2.10+)     │  │  (1.13+)    │  │  + GitHub App credentials    │   │
│  └─────────────┘  └─────────────┘  └───────────────────────────────┘   │
│                                                                          │
│  Secrets: github-app-credentials, ghcr-creds*, github-provider-token*   │
│  (* = optionally Helm-managed via values.yaml)                           │
│                                                                          │
│  Per-API: Repo with manifests/ directory                                 │
└──────────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────────────┐
│                     DRIVEBY HELM CHART CREATES                           │
│                                                                          │
│  Namespace ─ SA ─ Role ─ RoleBinding ─ Secrets (api-auth, github-app)   │
│  AppProject ─ Provider ─ ProviderConfig (×2) ─ Functions (×2)           │
│  XRD ─ Composition ─ GitHub Provider (optional)                         │
│  ghcr-creds* ─ github-provider-token* (* if enabled in values.yaml)     │
└──────────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌──────────────────────────────────────────────────────────────────────────┐
│                     XSDLC CR CREATES (per app)                           │
│                                                                          │
│  App-level (~6):                                                         │
│    ServiceAccount ─ EventBus ─ ScmProvider ─ GitRepository               │
│    PromotionStrategy ─ ArgoCDCommitStatus                                │
│                                                                          │
│  Per-environment (~3 each, ×3 envs = ~9):                                │
│    ArgoCD Application ─ Branch (env) ─ Branch (env-next)                 │
│                                                                          │
│  Per-gate (~9 each, ×2 gates = ~18):                                     │
│    Role ─ RoleBinding ─ WorkflowTemplate ─ EventSource                   │
│    Sensor ─ Service ─ Ingress ─ CommitStatus ─ BranchProtection          │
│                                                                          │
│  Repo workflow (1): driveby-deploy                                      │
│                                                                          │
│  Total: ~34 resources for a 3-env, 2-gate setup                         │
└──────────────────────────────────────────────────────────────────────────┘
```

## DriveBy Helm Chart Creates

Resources created by `helm upgrade --install driveby ./kubernetes/helm/driveby/`:

| # | Resource | Kind | Name | Namespace | Condition |
|---|----------|------|------|-----------|-----------|
| 1 | Namespace | `Namespace` | `driveby` | — | Always |
| 2 | Service Account | `ServiceAccount` | `driveby-sa` | `driveby` | `rbac.create` |
| 3 | Role | `Role` | `driveby-role` | `driveby` | `rbac.create` |
| 4 | RoleBinding | `RoleBinding` | `driveby-rolebinding` | `driveby` | `rbac.create` |
| 5 | API Auth Secret | `Secret` | `driveby-api-auth` | `driveby` | Always |
| 6 | GitHub App Secret | `Secret` | `github-app-credentials` | `driveby` | `gitopsPromoter.githubApp.appID` set |
| 7 | ArgoCD AppProject | `AppProject` | `driveby` | `argocd` | `argocd.enabled` |
| 8 | Kubernetes Provider | `Provider` | `provider-kubernetes` | — | `crossplane.enabled` |
| 9 | K8s ProviderConfig | `ProviderConfig` | `kubernetes-provider` | — | `crossplane.enabled` |
| 10 | GitHub Provider | `Provider` | `provider-upjet-github` | — | `githubProvider.enabled` |
| 11 | GitHub ProviderConfig | `ProviderConfig` | `github-provider` | — | `githubProvider.enabled` |
| 12 | Go-Templating Function | `Function` | `function-go-templating` | — | `crossplane.enabled` |
| 13 | Auto-Ready Function | `Function` | `function-auto-ready` | — | `crossplane.enabled` |
| 14 | XSDLC XRD | `CompositeResourceDefinition` | `xsdlcs.driveby.io` | — | `crossplane.enabled` |
| 15 | XSDLC Composition | `Composition` | `xsdlc-composition` | — | `crossplane.enabled` |
| 16 | GHCR Credentials | `Secret` | `ghcr-creds` | `driveby` | `secrets.ghcr.enabled` |
| 17 | GitHub Provider Token | `Secret` | `github-provider-token` | `driveby` | `githubProvider.token` set |

Items 16–17 are optional secrets added to reduce manual `kubectl create secret` steps.

## XSDLC CR Creates (Per App)

Resources generated by a single `XSDLC` CR via the Crossplane Composition:

### App-Level Resources (~6)

| # | Resource | Kind | Name Pattern | Namespace |
|---|----------|------|-------------|-----------|
| 1 | Service Account | `ServiceAccount` | `<app>-workflows` | `driveby` |
| 2 | Event Bus | `EventBus` | `default` | `driveby` |
| 3 | SCM Provider | `ScmProvider` | `<app>-github` | `driveby` |
| 4 | Git Repository | `GitRepository` | `<app>-gitops` | `driveby` |
| 5 | Promotion Strategy | `PromotionStrategy` | `<app>-promotion` | `driveby` |
| 6 | ArgoCD Commit Status | `ArgoCDCommitStatus` | `<app>-argocd-health` | `driveby` |

### Per-Environment Resources (always generated)

For each environment:

| # | Resource | Kind | Name Pattern | Namespace |
|---|----------|------|-------------|-----------|
| 1 | ArgoCD Application | `Application` | `<app>-<env>` | `argocd` |
| 2 | Environment Branch | `Branch` | `environment/<env>` | — |
| 3 | Environment Next Branch | `Branch` | `environment/<env>-next` | — |

### Per-Gate Resources (~9 each)

For each environment with a `gate` defined:

| # | Resource | Kind | Name Pattern | Namespace |
|---|----------|------|-------------|-----------|
| 1 | Role | `Role` | `<app>-<env>-gate` | `driveby` |
| 2 | Role Binding | `RoleBinding` | `<app>-<env>-gate` | `driveby` |
| 3 | Workflow Template | `WorkflowTemplate` | `<app>-<env>-gate` | `driveby` |
| 4 | Event Source | `EventSource` | `<app>-<env>-gate-eventsource` | `driveby` |
| 5 | Sensor | `Sensor` | `<app>-<env>-gate-sensor` | `driveby` |
| 6 | Service | `Service` | `<app>-<env>-gate-eventsource-svc` | `driveby` |
| 7 | Ingress | `Ingress` | `<app>-<env>-gate-webhook-ingress` | `driveby` |
| 8 | Commit Status | `CommitStatus` (Promoter CRD) | `<env>-gate` | `driveby` |
| 9 | Branch Protection | `BranchProtection` | `<app>-<env>-protection` | `driveby` |

Each gate's WorkflowTemplate contains a dynamic DAG with `check-<idx>-<type>` steps (e.g., `check-0-validate-only`, `check-1-functional-test`).

**Example**: For `perfect-api` with 3 environments (dev, staging, prod) and 2 gates (staging-gate, prod-gate):
- 6 app-level + (3 × 3 per-env) + (9 × 2 gates) + 1 repo workflow = **~34 resources**

## Client Prerequisites — Cluster Infrastructure

Components the client must install **before** the Helm chart:

| # | Component | Min Version | Purpose | Install Notes |
|---|-----------|-------------|---------|---------------|
| 1 | Kubernetes | 1.27+ | Cluster | Any distribution |
| 2 | Crossplane | 1.14+ | XRD/Composition engine | `crossplane-system` namespace |
| 3 | ArgoCD | 2.8+ | GitOps deployment | Pre-installed; chart does NOT install it |
| 4 | Argo Workflows | 3.5+ | Validation/loadtest pipelines | Controller running |
| 5 | Argo Events | 1.9+ | Webhook → workflow triggers | Controller running |
| 6 | Traefik | 2.10+ | Ingress controller | `ingressClassName: traefik-system` |
| 7 | cert-manager | 1.13+ | TLS certificates | With a configured ClusterIssuer |
| 8 | GitOps Promoter | 0.1+ | Automated promotion | Controller + CRDs installed |

**Why these can't be automated**: Each is a complex system with its own lifecycle, RBAC, CRDs, and storage requirements. Organizations typically have opinionated installations of these components. DriveBy integrates with them rather than installing them.

## Client Prerequisites — Secrets

| # | Secret | Namespace | Keys | Helm-Managed? |
|---|--------|-----------|------|---------------|
| 1 | `github-app-credentials` | `driveby` | `githubAppID`, `githubInstallationID`, `githubAppPrivateKey` | Yes — via `gitopsPromoter.githubApp.*` values |
| 2 | `driveby-api-auth` | `driveby` | `api-key`, `api-key-header` | Yes — via `apiAuth.*` values |
| 3 | `ghcr-creds` | `driveby` + each env namespace | `.dockerconfigjson` | Partially — Helm creates in `driveby` ns if `secrets.ghcr.enabled=true`; env namespaces need separate creation |
| 4 | `github-provider-token` | `driveby` | `credentials` (JSON: token + owner) | Yes — via `githubProvider.token` and `githubProvider.owner` values |
| 5 | `ghcr-creds` | per-env (`<app>-dev`, etc.) | `.dockerconfigjson` | No — must be created manually or via ArgoCD |
| 6 | `api-auth` | per-env (`<app>-dev`, etc.) | `api-key`, `api-key-header` | No — must be created manually or via ArgoCD |

**Why per-env secrets can't be Helm-managed**: The Helm chart doesn't know which applications or environments will exist — those are defined later by XSDLC CRs. Per-environment secrets should be managed via ArgoCD SealedSecrets, ExternalSecrets, or Kustomize overlays in the gitops repo.

## Client Prerequisites — Per-API Setup

For each API onboarded to DriveBy quality gates:

| # | Task | Details | One-Time? |
|---|------|---------|-----------|
| 1 | Have a repo with source code + `manifests/` directory | The repo contains both application source and a `manifests/` directory with K8s manifests (deployment.yaml, service.yaml, etc.) — same manifests for all environments; gates control when, not what | Yes |
| 2 | Create K8s secret with GitHub PAT | `kubectl create secret generic driveby-gitops-pat -n driveby --from-literal=token=ghp_xxx` — used for branch operations and workflow provisioning | Yes |
| 3 | Apply XSDLC CR | Single ~35-line YAML — generates branches, workflows, ArgoCD apps, promoter, quality gates, branch protection | Yes |

**That's it.** Branches, workflows, ArgoCD apps, promoter, quality gates, branch protection — all automated. Environment config (replicas, resources, env vars) is managed by the engineer in the `manifests/` directory.

## What Cannot Be Automated (and Why)

| Item | Reason |
|------|--------|
| Cluster infrastructure (8 components) | Each is a complex system with its own RBAC, storage, and lifecycle. Organizations have opinionated installations. |
| GitHub App creation | Requires manual GitHub UI interaction, organization approval, and private key download. |
| Repository `manifests/` directory | Engineers must create the `manifests/` directory with their manifests (deployment.yaml, service.yaml, etc.) — this is app-specific configuration. Same manifests are used for all environments; the developer's CI updates the image tag before pushing to `main`. |
| Per-environment secrets | Unknown at Helm time — app namespaces are created later by XSDLC + ArgoCD. Best managed via SealedSecrets or ExternalSecrets in the manifests/ directory. |
| GitHub PAT secret | Must be manually created in-cluster (`driveby-gitops-pat`) — used for branch operations and workflow provisioning. |

## Setup Checklist

Copy-paste ready checklist for onboarding a new cluster and API.

### One-Time Cluster Setup

```bash
# 1. Verify prerequisites are running
kubectl get pods -n crossplane-system        # Crossplane
kubectl get pods -n argocd                   # ArgoCD
kubectl get pods -n argo                     # Argo Workflows
kubectl get pods -n argo-events              # Argo Events
kubectl get pods -n traefik-system           # Traefik (or your ingress namespace)
kubectl get pods -n cert-manager             # cert-manager
kubectl get pods -n promoter-system          # GitOps Promoter

# 2. Pre-install Crossplane provider (CRDs must exist before Helm ProviderConfig)
kubectl apply -f - <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-kubernetes
spec:
  package: xpkg.upbound.io/crossplane-contrib/provider-kubernetes:v0.14.1
EOF
kubectl wait --for=condition=Healthy provider/provider-kubernetes --timeout=120s

# 3. Install Helm chart
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set crossplane.enabled=true \
  --set gitopsPromoter.githubApp.appID=YOUR_APP_ID \
  --set gitopsPromoter.githubApp.installationID=YOUR_INSTALLATION_ID \
  --set-file gitopsPromoter.githubApp.privateKey=/path/to/key.pem \
  --set secrets.ghcr.enabled=true \
  --set secrets.ghcr.username=YOUR_GITHUB_USER \
  --set secrets.ghcr.token=YOUR_GHCR_TOKEN \
  --set githubProvider.token=YOUR_GITHUB_PAT \
  --set githubProvider.owner=YOUR_GITHUB_ORG

# 4. Verify Helm resources
kubectl get providers,functions               # All Healthy
kubectl get xrd                               # xsdlcs.driveby.io
kubectl get secrets -n driveby                # All secrets created
```

### Per-API Onboarding

```bash
# 1. Ensure your repo has a manifests/ directory with K8s manifests
# Structure: manifests/deployment.yaml, manifests/service.yaml, etc.
# Same manifests for all environments — gates control when, not what

# 2. Create K8s secret with GitHub PAT (one-time, shared across APIs)
kubectl create secret generic driveby-gitops-pat \
  --namespace driveby \
  --from-literal=token=ghp_xxxxxxxxxxxx

# 3. Apply XSDLC CR — everything else is automated
kubectl apply -f - <<EOF
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: your-api
  namespace: driveby
spec:
  repository:
    owner: your-org
    name: your-api
  manifestsPath: manifests
  environments:
    - name: dev
    - name: staging
      gate:
        checks:
          - type: validate-only
          - type: functional-test
    - name: prod
      autoMerge: false
      gate:
        checks:
          - type: validate-only
            validationConfig:
              validationMode: strict
          - type: load-test
EOF

# 4. Verify
kubectl get xsdlcs -n driveby
kubectl get workflowtemplates,eventsources,sensors,ingress -n driveby
kubectl get applications -n argocd | grep your-api   # ArgoCD apps auto-created
gh api repos/your-org/your-api/contents/.github/workflows --jq '.[].name'
# Expected: driveby-deploy.yml
```
