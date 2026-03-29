# DriveBy Deployment Guide

Complete guide to installing DriveBy quality gates on a Kubernetes cluster (XSDLC v3.0.0 — two-repo GitOps model).

## Prerequisites

```
┌─────────────────────────────────────────────────────┐
│                  Kubernetes Cluster                  │
│                                                     │
│  ┌──────────────┐  ┌──────────────┐  ┌───────────┐ │
│  │  Crossplane   │  │   ArgoCD     │  │  Traefik  │ │
│  │  (1.14+)      │  │   (2.8+)     │  │  (2.10+)  │ │
│  └──────────────┘  └──────────────┘  └───────────┘ │
│  ┌──────────────┐  ┌──────────────┐  ┌───────────┐ │
│  │ Argo Workflows│  │ Argo Events  │  │cert-manager│ │
│  │   (3.5+)      │  │   (1.9+)     │  │  (1.13+)  │ │
│  └──────────────┘  └──────────────┘  └───────────┘ │
└─────────────────────────────────────────────────────┘
```

| Component | Minimum Version | Notes |
|-----------|----------------|-------|
| Kubernetes | 1.27+ | |
| Crossplane | 1.14+ | Must be installed in `crossplane-system` namespace |
| ArgoCD | 2.8+ | Pre-installed; chart does NOT install ArgoCD |
| Argo Workflows | 3.5+ | Controller running |
| Argo Events | 1.9+ | Controller running |
| Traefik | 2.10+ | Ingress controller (`ingressClassName: traefik-system`) |
| cert-manager | 1.13+ | With a configured ClusterIssuer |
| GitOps Promoter | 0.1+ | Required; install controller with ≥512Mi memory for 5+ apps |

### GitHub App Requirements

The GitHub App used by both the Promoter and Argo Events must have access to **all gitops repositories** created by XSDLC. Two options:

1. **Recommended**: Set the App's repository access to **"All repositories"** in the GitHub org settings. This ensures any new gitops repo auto-created by XSDLC is automatically accessible.
2. **Manual**: After each XSDLC CR is applied, add the newly created gitops repo to the App's installation (GitHub → Org Settings → Applications → Configure).

If the App doesn't have access to a gitops repo, the Promoter will fail with `repository not found` errors and no promotion PRs will be created.

### GitOps Promoter Memory

The Promoter controller needs sufficient memory for the number of GitRepositories it manages. Default 128Mi is insufficient for 5+ apps:

| Apps | Recommended Memory (manager container) |
|------|----------------------------------------|
| 1-3 | 256Mi |
| 4-10 | 1Gi |
| 10+ | 2Gi |

**Important**: The `manager` container is typically at index `[0]` in the deployment spec. Ensure you patch the correct container — the `kube-rbac-proxy` sidecar needs only 128Mi.

## Repository Structure

XSDLC v3.0.0 uses a **two-repo model** — the software repo and the GitOps repo are separate concerns:

### Software Repository (developer-owned, untouched by XSDLC)

The application repository (e.g., `novelcore/perfect-api`) contains only the source code and build artifacts. XSDLC does not create, modify, or generate any files in this repo.

```
perfect-api/                # Software repo — developer-owned
  src/                      # application source code
  Dockerfile                # container build
  .github/workflows/        # developer's own CI (build, test, push image)
  manifests/                # Kubernetes manifests (optional — developers may keep them here)
    deployment.yaml
    service.yaml
```

### GitOps Repository (auto-created by XSDLC)

The GitOps repository (e.g., `novelcore/perfect-api-gitops`) is a dedicated repo that holds **dry manifests** on the `main` branch using a Kustomize overlay structure. XSDLC auto-creates this repo via `provider-upjet-github` when `githubProvider.enabled`.

```
perfect-api-gitops/         # GitOps repo — auto-created
  main branch:
    dry/
      base/                 # Shared base manifests
        deployment.yaml
        service.yaml
        kustomization.yaml
      overlays/
        dev/                # Per-environment Kustomize overlays
          kustomization.yaml
        staging/
          kustomization.yaml
        prod/
          kustomization.yaml
  environment branches (written by ArgoCD Source Hydrator — not by developers):
    environment/dev-next:   manifests/ + hydrator.metadata
    environment/dev:        manifests/ + hydrator.metadata
    environment/staging-next: ...
    ...
```

- The `gitopsRepository` field is required — specify `owner` (GitHub org) and `name` (gitops repo name).
- Dry manifests live on `main` in `dry/base/` and `dry/overlays/<env>/`. Each overlay references `../../base` and can customize replicas, env vars, images, etc.
- ArgoCD Source Hydrator renders each overlay and writes the hydrated output to `environment/<env>-next` branches automatically. Developers do **not** push to `-next` branches directly.

### Branch Structure (Auto-Created in GitOps Repo)

When `githubProvider.enabled`, the XSDLC auto-creates **6 environment branches** in the **GitOps repository** (e.g., `perfect-api-gitops`):

```
environment/dev-next      ← Hydrator writes rendered manifests here (from dry/overlays/dev on main)
environment/dev           ← Promoter auto-PRs from dev-next; ArgoCD tracks this
environment/staging-next  ← Hydrator writes rendered manifests here (from dry/overlays/staging on main)
environment/staging       ← Promoter auto-PRs from staging-next; Gate 1 blocks this
environment/prod-next     ← Hydrator writes rendered manifests here (from dry/overlays/prod on main)
environment/prod          ← Promoter auto-PRs from prod-next; Gate 2 blocks this
```

Each `-next` branch is written by the ArgoCD Source Hydrator, not by developers. The hydrator renders the per-environment Kustomize overlay from `main` and commits the result to the corresponding `-next` branch. Each environment hydrates independently — there is no linear propagation between environments.

### ArgoCD Applications (Auto-Created)

ArgoCD Applications are always generated by the XSDLC composition for every environment. Each application uses **`sourceHydrator`** to render its per-environment Kustomize overlay from the `main` branch of the GitOps repository. The hydrator reads from `dry/overlays/<env>/` on `main`, renders the Kustomize output, and writes the hydrated manifests to the `environment/<env>-next` branch. All applications get `syncPolicy.automated` with `selfHeal: true` and `prune: true` regardless of the `autoMerge` setting — `autoMerge` only controls whether the Promoter auto-merges promotion PRs.

### Complete Lifecycle (Two-Repo Model)

The developer's own CI builds the container image; the developer updates dry manifests on the `main` branch of the GitOps repo (e.g., updating the image tag in an overlay). XSDLC does not generate any CI workflow — developers are free to update dry manifests however they want (CI pipeline, script, manual push, etc.).

```
Developer updates dry manifests on main (e.g., image tag in dry/overlays/<env>/kustomization.yaml)
    → ArgoCD Source Hydrator renders overlay → writes to environment/<env>-next branch
    → Promoter creates promotion PR (env-next → env)
    → Quality gates validate → auto-merge (or manual approval)
    → ArgoCD syncs environment (always autoSync)
```

### How to Deploy

Deployment starts by updating the dry manifests on the `main` branch of the GitOps repo. The ArgoCD Source Hydrator renders the overlay and writes the result to the `-next` branch; the Promoter handles everything from there.

**Option A**: Update from a CI pipeline:

```bash
# Clone the gitops repo, update dry manifests on main, push
git clone https://github.com/novelcore/perfect-api-gitops.git
cd perfect-api-gitops
# Update image tag in the dev overlay (or base for all environments)
cd dry/overlays/dev
cat > kustomization.yaml <<EOF
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - ../../base
images:
  - name: ghcr.io/novelcore/perfect-api
    newTag: v1.2.3
EOF
git add . && git commit -m "deploy v1.2.3 to dev" && git push origin main
```

**Option B**: Update manually:

```bash
# Same flow — edit dry manifests on main
git clone https://github.com/novelcore/perfect-api-gitops.git
cd perfect-api-gitops
# Edit base or overlay manifests as needed
vim dry/overlays/dev/kustomization.yaml
git add . && git commit -m "deploy v1.2.3 to dev" && git push origin main
```

The ArgoCD Source Hydrator detects the change on `main`, renders the Kustomize overlay, and writes the hydrated manifests to the `-next` branch. The Promoter then creates a promotion PR, and the quality gates take over.

## Namespace Layout

```
Kubernetes Cluster
├── driveby/                     ← DriveBy control plane
│   ├── Crossplane XRD + Composition
│   ├── Argo Events (EventBus, EventSource, Sensor)
│   ├── Argo Workflows (WorkflowTemplates, RBAC)
│   ├── GitOps Promoter (ScmProvider, GitRepository, PromotionStrategy)
│   └── Secrets (driveby-api-auth, github-app-credentials)
├── perfect-api-dev/             ← Dev environment (autoSync)
│   └── API deployment + service
├── perfect-api-staging/         ← Staging environment (autoSync)
│   └── API deployment + service
└── perfect-api-prod/            ← Production environment (autoSync — autoMerge: false)
    └── API deployment + service
```

**All ArgoCD Applications use autoSync** — every application gets `syncPolicy.automated` with `selfHeal: true` and `prune: true`, regardless of the `autoMerge` setting. The `autoMerge` field only controls whether the Promoter automatically merges promotion PRs:
- `autoMerge: true` (default) → Promoter auto-merges the PR from `env-next` → `env` when gates pass
- `autoMerge: false` → Promoter creates the PR but requires manual approval before merging

Once a promotion PR is merged (automatically or manually), ArgoCD always syncs the environment immediately.

## Promotion Flow

Quality gates use a **multi-check model**: each gate defines an ordered list of checks. Developers update dry manifests on `main`; the ArgoCD Source Hydrator renders per-environment overlays and writes to `-next` branches; the Promoter handles promotion through the environment chain.

```
Developer updates dry manifests on main (dry/base/ or dry/overlays/<env>/)
    → Hydrator renders dry/overlays/dev/ → writes to environment/dev-next
    → Promoter auto-PRs environment/dev-next → environment/dev → ArgoCD syncs dev

Gate 1 (staging-gate):
    → Hydrator renders dry/overlays/staging/ → writes to environment/staging-next
    → Promoter auto-PRs environment/staging-next → environment/staging
    → Webhook triggers DriveBy workflow (validate-only + functional-test) against dev
    → Workflow creates CommitStatus CRD (staging-gate, phase: success)
    → Promoter auto-merges → ArgoCD syncs staging

Gate 2 (prod-gate):
    → Hydrator renders dry/overlays/prod/ → writes to environment/prod-next
    → Promoter auto-PRs environment/prod-next → environment/prod
    → Webhook triggers DriveBy workflow (validate-only + load-test) against staging
    → Workflow creates CommitStatus CRD (prod-gate, phase: success)
    → Promoter does NOT auto-merge (autoMerge: false — manual approval required)
```

Each environment hydrates independently from `main` — there is no linear propagation between environments. All environments can receive updates simultaneously when dry manifests on `main` change.

## Install Order

```
1. Pre-install providers ──► provider-kubernetes must be healthy before Helm
       │
2. Secrets           ──► GitHub App, API auth, registry creds
       │
3. Helm Install      ──► Providers, Functions, XRD, Composition
       │
4. Apply XSDLC      ──► Single CR generates gitops repo + all gate resources
       │
5. GitHub Webhooks   ──► Auto-created per gate (one per Ingress URL)
       │
6. GitOps Promoter   ──► Install controller + GitHub App secret
```

## 1. Pre-install Provider

The Helm chart includes the Provider CR, but the ProviderConfig CRD won't exist until the provider is installed and running:

```bash
kubectl apply -f - <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-kubernetes
spec:
  package: xpkg.upbound.io/crossplane-contrib/provider-kubernetes:v0.14.1
EOF

kubectl wait --for=condition=Healthy provider/provider-kubernetes --timeout=120s
```

## 2. Install the Helm Chart

**Important**: The `driveby` namespace must exist before install. If reinstalling after `helm uninstall`, the namespace may already exist — do NOT use `--create-namespace` in that case (Helm will error). If installing for the first time, create the namespace first:

```bash
# First-time install only — create namespace if it doesn't exist
kubectl create namespace driveby --dry-run=client -o yaml | kubectl apply -f -
```

Full install command with all secrets baked in (recommended — avoids manual secret creation):

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --namespace driveby \
  --set crossplane.enabled=true \
  --set githubProvider.enabled=true \
  --set githubProvider.token="$(gh auth token)" \
  --set githubProvider.owner=novelcore \
  --set secrets.ghcr.enabled=true \
  --set secrets.ghcr.username=meter-peter \
  --set secrets.ghcr.token="$(gh auth token)" \
  --set-file gitopsPromoter.githubApp.privateKey=/path/to/driveby-promoter.private-key.pem
```

This creates all required secrets automatically:
- `github-provider-token` — for `provider-upjet-github` (branch protection, repo creation)
- `github-app-credentials` — for Promoter SCM access, webhook creation, commit status
- `ghcr-creds` — for pulling DriveBy container image from ghcr.io

**Minimal install** (without GitHub branch protection or Promoter PEM baked in):

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --namespace driveby \
  --set crossplane.enabled=true
```

This creates:
- **Provider**: `provider-kubernetes` (v0.14.1) — adopted if pre-installed
- **Functions**: `function-go-templating`, `function-auto-ready`
- **ProviderConfig**: `kubernetes-provider` (InjectedIdentity)
- **XRD**: `xsdlcs.driveby.io`
- **Composition**: XSDLC composition

Wait for functions to become healthy:

```bash
kubectl get providers -w
kubectl get functions -w
```

## 3. Create Secrets

### GitHub App Credentials (for all GitHub operations)

Required for commit status, PR comments, and GitOps Promoter SCM access:

```bash
kubectl create secret generic github-app-credentials \
  --namespace driveby \
  --from-literal=githubAppID=YOUR_APP_ID \
  --from-literal=githubInstallationID=YOUR_INSTALLATION_ID \
  --from-file=githubAppPrivateKey=/path/to/private-key.pem
```

Or via Helm values:

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set gitopsPromoter.githubApp.appID=12345 \
  --set gitopsPromoter.githubApp.installationID=67890 \
  --set-file gitopsPromoter.githubApp.privateKey=/path/to/private-key.pem
```

### ArgoCD Push Secret (for Source Hydrator)

**Auto-managed by XSDLC** — the composition automatically creates a `repository-write` secret in the `argocd` namespace for each gitops repo. No manual secret creation needed.

After applying an XSDLC CR, verify the push secret was created:

```bash
kubectl get secrets -n argocd -l argocd.argoproj.io/secret-type=repository-write
# Expected: one secret per XSDLC (e.g., perfect-api-push-secret)
```

**Important**: After the first XSDLC is applied, restart the ArgoCD commit-server pod so it picks up the new credentials:

```bash
kubectl delete pod -n argocd -l app.kubernetes.io/name=argocd-commit-server
```

### Container Registry Credentials (per-namespace)

**Auto-managed by XSDLC** — the composition automatically copies the `ghcr-creds` secret from the `driveby` namespace into each environment namespace (e.g., `perfect-api-dev`, `perfect-api-staging`). No manual per-namespace secret creation needed.

The source `ghcr-creds` secret must exist in the `driveby` namespace (created via Helm values or manually — see Container Registry section below).

### API Auth (for DriveBy validation)

```bash
kubectl create secret generic driveby-api-auth \
  --namespace driveby \
  --from-literal=api-key=YOUR_API_KEY \
  --from-literal=api-key-header=X-API-Key
```

### Container Registry (for private images)

**Option A**: Via Helm values (recommended — creates the secret automatically):

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set secrets.ghcr.enabled=true \
  --set secrets.ghcr.username=YOUR_USERNAME \
  --set secrets.ghcr.token=YOUR_TOKEN
```

**Option B**: Manually:

```bash
kubectl create secret docker-registry ghcr-creds \
  --namespace driveby \
  --docker-server=ghcr.io \
  --docker-username=YOUR_USERNAME \
  --docker-password=YOUR_TOKEN
```

### GitHub Provider Token (for branch protection)

**Option A**: Via Helm values (recommended — creates the secret automatically):

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set githubProvider.token=YOUR_GITHUB_PAT \
  --set githubProvider.owner=YOUR_GITHUB_ORG
```

**Option B**: Manually:

```bash
kubectl create secret generic github-provider-token \
  --namespace driveby \
  --from-literal=credentials='{"token":"YOUR_GITHUB_PAT","owner":"YOUR_GITHUB_ORG"}'
```

## 4. Apply XSDLC

A single XSDLC CR defines the entire promotion pipeline:

```bash
kubectl apply -f kubernetes/examples/novelcore-perfect-api/xsdlc-perfect-api.yaml
```

This single CR generates all resources:
- **GitOps repo**: Dedicated repository via `provider-upjet-github` (e.g., `perfect-api-gitops`)
- **App-level**: ServiceAccount, EventBus, ScmProvider, GitRepository, PromotionStrategy, ArgoCDCommitStatus
- **Per-environment**: ArgoCD Application (using sourceHydrator), environment branches (env + env-next), ghcr-creds secret
- **Per-gate**: Role, RoleBinding, WorkflowTemplate, EventSource, Sensor, Service, Ingress, BranchProtection
- **ArgoCD**: Push secret (repository-write) for the gitops repo

Verify:

```bash
kubectl get xsdlcs -n driveby                          # XSDLC status (ENVIRONMENTS=3, GATES=2)
kubectl get workflowtemplates -n driveby               # 2 WorkflowTemplates per API
kubectl get eventbus,eventsources,sensors -n driveby   # Event infrastructure
kubectl get ingress -n driveby                         # 2 webhook Ingresses per API
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy -n driveby
```

**Important: Restart ArgoCD commit-server** after the first XSDLC is applied, so it picks up the new push secrets:

```bash
kubectl delete pod -n argocd -l app.kubernetes.io/name=argocd-commit-server
```

**Important: Hard-refresh all ArgoCD apps** to trigger the Source Hydrator immediately (otherwise it may take up to 3 minutes per app):

```bash
for app in $(kubectl get applications -n argocd -o name | grep -E 'perfect-api|bad-docs|no-auth|slow-api|broken'); do
  kubectl patch $app -n argocd --type=merge -p '{"metadata":{"annotations":{"argocd.argoproj.io/refresh":"hard"}}}'
done
```

Check that the GitOps repository was created:

```bash
gh repo view novelcore/perfect-api-gitops --json name,defaultBranchRef
# Expected: repository exists

# Verify environment branches were created in the gitops repo
gh api repos/novelcore/perfect-api-gitops/branches --jq '.[].name'
# Expected: environment/dev, environment/dev-next, environment/staging, environment/staging-next, environment/prod, environment/prod-next
```

## 5. Verify GitHub Webhook (auto-created)

Argo Events will automatically create the **repository webhook** when the XSDLC is reconciled, using the GitHub App credentials and the generated EventSource Ingress URL.

Expected webhook URL format:
- **URL**: `https://<app>-<env>-<type>-webhook.<baseDomain>/<app>-<env>-<type>-<trigger>`
- **Content type**: `application/json`
- **Events**: `pull_request`

## 6. Verification Checklist

```bash
# Crossplane resources healthy
kubectl get providers         # provider-kubernetes: Healthy
kubectl get functions         # All functions: Healthy

# XRD registered
kubectl get xrd               # xsdlcs.driveby.io

# XSDLC resources created
kubectl get xsdlcs -n driveby                 # Ready: True
kubectl get workflowtemplates -n driveby      # 2 templates
kubectl get role,rolebinding -n driveby       # RBAC for both gates
kubectl get eventbus -n driveby               # default: Running
kubectl get eventsources -n driveby           # 2 EventSources: Running
kubectl get sensors -n driveby                # 2 Sensors: Running
kubectl get ingress -n driveby                # 2 webhook Ingresses with TLS

# Promoter resources
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy -n driveby
kubectl get changetransferpolicies -n driveby      # Auto-created by PromotionStrategy
kubectl get argocdcommitstatuses -n driveby
```

## Configuration Model

DriveBy uses a **two-tier configuration model**:

| Tier | Resolved at | Scope | How to set |
|------|-------------|-------|------------|
| `values.yaml` | `helm template` time | Chart-wide defaults | `--set defaults.X=Y` or custom values file |
| XRD spec fields | Crossplane composition runtime | Per-CR overrides | Set directly in `XSDLC` spec |

**Resolution order** (highest priority wins): XRD spec → values.yaml baked default.

### Configurable Knobs

| Knob | values.yaml path | XRD field |
|------|------------------|-----------|
| Ingress class | `defaults.ingress.className` | (values.yaml only) |
| cert-manager issuer | `defaults.ingress.clusterIssuer` | (values.yaml only) |
| Base domain | `defaults.ingress.baseDomain` | (values.yaml only) |
| Webhook port | `defaults.webhookPort` | (values.yaml only) |
| Workflows UI URL | `defaults.argoWorkflowsUrl` | (values.yaml only) |
| API port | — | `apiConfig.port` |
| OpenAPI endpoint | — | `apiConfig.openapiEndpoint` |
| Service name | — | `apiConfig.serviceName` |
| DriveBy image | — | `validationDefaults.drivebyImage` |
| Validation mode | — | `validationDefaults.validationMode` or per-gate `validationConfig.validationMode` |
| Load test params | — | Per-gate `loadTestConfig.*` |

## Customization

### Validation Mode

Control strictness via per-check `validationConfig.validationMode` or top-level `validationDefaults.validationMode`:
- `minimal` — Basic OpenAPI compliance only (P001)
- `strict` — All static validation principles (P001-P005, P008)
- `test-ready` — Validates spec is ready for testing (P001, P002, P003, P004, P009)
- `test-only` — Functional testing only (P006)

### Multiple APIs

Adding a new API is a single XSDLC CR:

```yaml
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: another-api
  namespace: driveby
spec:
  gitopsRepository:
    owner: meter-peter
    name: another-api-gitops
  environments:
    - name: dev
    - name: staging
      gate:
        checks:
          - type: validate-only
            validationConfig:
              validationMode: strict
          - type: functional-test
    - name: prod
      autoMerge: false
      gate:
        checks:
          - type: validate-only
            validationConfig:
              validationMode: strict
          - type: load-test
```

No manual branch creation, workflow wiring, or EnvironmentConfigs needed. The GitOps repository, ArgoCD Applications, and environment branches are all auto-generated.

## Clean Reinstall

When upgrading the composition (e.g., source hydrator changes), a full clean reinstall ensures no stale resources remain:

```bash
# 1. Delete XSDLC (cascade-deletes all generated resources)
kubectl delete xsdlc perfect-api -n driveby

# 2. Delete the gitops repo (so XSDLC creates fresh with new structure)
gh repo delete novelcore/perfect-api-gitops --yes

# 3. Uninstall Helm chart
helm uninstall driveby -n driveby

# 4. Reinstall (namespace persists — do NOT use --create-namespace)
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --namespace driveby \
  --set crossplane.enabled=true \
  --set githubProvider.enabled=true \
  --set githubProvider.token="$(gh auth token)" \
  --set githubProvider.owner=novelcore \
  --set secrets.ghcr.enabled=true \
  --set secrets.ghcr.username=meter-peter \
  --set secrets.ghcr.token="$(gh auth token)" \
  --set-file gitopsPromoter.githubApp.privateKey=/path/to/driveby-promoter.private-key.pem

# 5. Wait for providers + functions
kubectl get providers -w
kubectl get functions -w

# 6. Re-apply XSDLC
kubectl apply -f kubernetes/examples/novelcore-perfect-api/xsdlc-perfect-api.yaml
```

## Uninstall

### 1. Delete XSDLC

This cascade-deletes all resources (event infrastructure, RBAC, WorkflowTemplates, promoter resources):

```bash
kubectl delete xsdlc perfect-api -n driveby
```

### 2. Uninstall the Helm Chart

This removes the XRD, Composition, Functions, ProviderConfig, and chart RBAC:

```bash
helm uninstall driveby -n driveby
```

### 3. Clean Up Namespace

```bash
kubectl delete namespace driveby
```

### 4. Verify GitHub Webhooks Removed

If the EventSource was configured with `deleteHookOnFinish: true`, webhooks are automatically removed. Verify:

```bash
gh api repos/novelcore/perfect-api-gitops/hooks | jq '.[].config.url'
```

### 5. Verify No Orphaned Resources

```bash
kubectl get crd | grep driveby.io
kubectl get managed | grep driveby
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy --all-namespaces | grep driveby
```

## Troubleshooting

### Webhook not received
1. Check webhook deliveries: `gh api repos/novelcore/perfect-api-gitops/hooks/<id>/deliveries`
2. Verify Ingress resolves: `dig perfect-api-staging-validation-webhook.private.novelcore.org`
3. Check EventSource pod is running: `kubectl get pods -n driveby -l eventsource-name=<name>`
4. Check EventSource logs: `kubectl logs -n driveby -l eventsource-name=<name>`

### EventSource/Sensor crash-looping
1. Check EventBus is healthy: `kubectl get eventbus -n driveby`
2. Check for stream creation errors
3. Verify NATS pods are Running: `kubectl get pods -n driveby | grep eventbus`

### Workflow not triggered
1. Verify Sensor is subscribed: `kubectl logs -n driveby -l sensor-name=<name>`
2. Check the PR action is one of: `opened`, `reopened`, `synchronize`
3. Verify the EventSource received the webhook

### Workflow fails
1. Check workflow status: `kubectl get workflows -n driveby`
2. Get step logs: `kubectl logs -n driveby <pod-name> -c main`
3. Common failures:
   - `github-commit-status`: GitHub App credentials invalid or missing
   - `driveby-validate`: API not reachable from source namespace
   - `health-check-source`: Source environment not ready (5 min for first gates, 10 min for downstream gates)

## Teardown

### Remove individual XSDLCs

```bash
kubectl delete xsdlc <name> -n driveby
```

This deletes all managed resources including the gitops repository on GitHub (`deletionPolicy: Delete`).

### Full uninstall

```bash
# 1. Delete all XSDLCs first (cascade deletes managed resources + gitops repos)
kubectl delete xsdlc --all -n driveby

# 2. Wait for Crossplane cleanup
kubectl get object --no-headers | wc -l  # should be 0

# 3. If namespace gets stuck in Terminating (Promoter finalizers):
kubectl get pullrequests.promoter.argoproj.io -n driveby --no-headers | \
  awk '{print $1}' | xargs -I{} kubectl patch pullrequest.promoter.argoproj.io {} \
  -n driveby --type=merge -p '{"metadata":{"finalizers":null}}'

# 4. Helm uninstall
helm uninstall driveby -n driveby
```

**Known issue**: The GitOps Promoter's PullRequest CRDs have finalizers that can block namespace deletion. If the namespace gets stuck in `Terminating`, remove the finalizers manually (step 3 above).

### Reinstall after teardown

When reinstalling after a full teardown, the `driveby` namespace needs Helm ownership labels:

```bash
kubectl create namespace driveby --dry-run=client -o yaml | kubectl apply -f -
kubectl label namespace driveby app.kubernetes.io/managed-by=Helm --overwrite
kubectl annotate namespace driveby meta.helm.sh/release-name=driveby meta.helm.sh/release-namespace=driveby --overwrite
```

Then run the full install command from Section 2.
