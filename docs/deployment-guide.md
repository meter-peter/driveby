# DriveBy Deployment Guide

Complete guide to installing DriveBy quality gates on a Kubernetes cluster.

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
| GitOps Promoter | 0.1+ | Required; install controller for automated promotion |

## Namespace Layout

```
Kubernetes Cluster
├── driveby/                     ← DriveBy control plane
│   ├── Crossplane XRDs + Compositions
│   ├── Argo Events (EventBus, EventSource, Sensor)
│   ├── Argo Workflows (WorkflowTemplates, RBAC)
│   ├── GitOps Promoter (ScmProvider, GitRepository, PromotionStrategy)
│   └── Secrets (driveby-api-auth, github-app-credentials)
├── perfect-api-dev/             ← Dev environment (autoSync)
│   └── API deployment + service
├── perfect-api-staging/         ← Staging environment (autoSync)
│   └── API deployment + service
└── perfect-api-prod/            ← Production environment (manual sync)
    └── API deployment + service
```

## Promotion Flow

Quality gates trigger on **promotion PRs in the gitops repo**, validate against the **source environment** (dev), and gate promotion to the **target environment** (staging):

```
Push to main (DRY branch) → Hydrator builds environment/*-next branches
    → Promoter auto-PRs environment/dev-next → environment/dev → ArgoCD syncs dev
    → Promoter auto-PRs environment/staging-next → environment/staging
    → Webhook triggers DriveBy validation against dev
    → Workflow creates CommitStatus CRD (phase: success)
    → Promoter auto-merges → ArgoCD syncs staging
    → Promoter auto-PRs environment/prod-next → environment/prod (autoMerge: false — manual approval)
```

## Install Order

```
1. Pre-install providers ──► provider-kubernetes must be healthy before Helm
       │
2. Secrets           ──► GitHub App, API auth, registry creds
       │
3. Helm Install      ──► Providers, Functions, XRDs, Compositions
       │
4. EnvironmentConfigs ──► Cluster-specific settings
       │
5. XQualityGateTemplate ──► RBAC + WorkflowTemplate
       │
6. XQualityGate      ──► EventBus + EventSource + Sensor + Ingress
       │
7. GitHub Webhook    ──► Point gitops repo to Ingress URL
       │
8. GitOps Promoter ──► Install controller + GitHub App secret
```

## 1. Pre-install Provider

The Helm chart includes the Provider CR, but the ProviderConfig CRD won't exist until the provider is installed and running. Pre-install to avoid CRD chicken-and-egg:

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

```bash
helm upgrade --install driveby ./kubernetes/helm/driveby/ \
  --set crossplane.enabled=true
```

This creates:
- **Provider**: `provider-kubernetes` (v0.14.1) — adopted if pre-installed
- **Functions**: `function-go-templating`, `function-auto-ready`, `function-environment-configs`
- **ProviderConfig**: `kubernetes-provider` (InjectedIdentity)
- **XRDs**: `xqualitygatetemplates.driveby.io`, `xqualitygates.driveby.io`
- **Compositions**: Template and instance compositions
- **RBAC**: ServiceAccount, Role, RoleBinding

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

### API Auth (for DriveBy validation)

```bash
kubectl create secret generic driveby-api-auth \
  --namespace driveby \
  --from-literal=api-key=YOUR_API_KEY \
  --from-literal=api-key-header=X-API-Key
```

### Container Registry (for private images)

```bash
kubectl create secret docker-registry ghcr-creds \
  --namespace driveby \
  --docker-server=ghcr.io \
  --docker-username=YOUR_USERNAME \
  --docker-password=YOUR_TOKEN
```

## 4. Apply EnvironmentConfigs

EnvironmentConfigs provide cluster-specific settings to the compositions:

```bash
kubectl apply -f kubernetes/examples/crossplane/environment-configs.yaml
```

## 5. Create XQualityGateTemplate

The template defines the validation workflow (RBAC + WorkflowTemplate):

```bash
kubectl apply -f kubernetes/examples/crossplane/template-example.yaml
```

Verify:

```bash
kubectl get xqualitygatetemplates -n driveby
kubectl get workflowtemplates -n driveby
kubectl get serviceaccounts -n driveby
```

## 6. Create XQualityGate

The instance creates the event pipeline (EventBus + EventSource + Sensor + Ingress):

```bash
kubectl apply -f kubernetes/examples/crossplane/instance-example.yaml
```

Verify:

```bash
kubectl get xqualitygates -n driveby
kubectl get eventbus -n driveby
kubectl get eventsources -n driveby
kubectl get sensors -n driveby
kubectl get ingress -n driveby
```

## 6b. Verify Promoter Resources

```bash
kubectl get scmproviders -n driveby                # ScmProvider created
kubectl get gitrepositories.promoter.argoproj.io -n driveby  # GitRepository created
kubectl get promotionstrategy -n driveby           # PromotionStrategy created
kubectl get argocdcommitstatuses -n driveby        # ArgoCDCommitStatus created
kubectl get changetransferpolicies -n driveby      # Auto-created by PromotionStrategy
```

## 7. Verify GitHub Webhook (auto-created)

Argo Events will automatically create the **gitops repository webhook** when your `XQualityGate` is reconciled, using the GitHub App credentials (`github-app-credentials`) and the generated EventSource Ingress URL.

Prerequisite: your GitHub App must have permission to manage webhooks (typically `Administration: Read & write` or at least the `Webhooks` permission) on the target repository/org.

Expected webhook URL format:
- **URL**: `https://<app>-<gate>-webhook.<baseDomain>/<app>-<gate>-<trigger>`
- **Content type**: `application/json`
- **Events**: Select the events matching your trigger config (e.g., `pull_request`)

For the default example:
- **Repo**: `novelcore/perfect-api-gitops`
- **URL**: `https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation`
- **Events**: Pull requests

If you need to troubleshoot or manually recreate it:
```bash
gh api repos/novelcore/perfect-api-gitops/hooks \
  --method POST \
  -f name=web -F active=true \
  -f "config[url]=https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation" \
  -f "config[content_type]=json" \
  -f "events[]=pull_request"
```

## 8. Verification Checklist

```bash
# Crossplane resources healthy
kubectl get providers         # provider-kubernetes: Healthy
kubectl get functions         # All functions: Healthy

# XRDs registered
kubectl get xrd               # xqualitygatetemplates.driveby.io, xqualitygates.driveby.io

# Template resources created
kubectl get xqualitygatetemplates -n driveby   # Ready: True
kubectl get workflowtemplates -n driveby       # driveby-staging-promotion exists
kubectl get role,rolebinding -n driveby        # driveby-staging-promotion exists

# Instance resources created
kubectl get xqualitygates -n driveby           # Ready: True
kubectl get eventbus -n driveby                # default: Running
kubectl get eventsources -n driveby            # EventSource: Running
kubectl get sensors -n driveby                 # Sensor: Running
kubectl get ingress -n driveby                 # Webhook ingress with TLS

# Promoter resources
kubectl get scmproviders,gitrepositories.promoter,promotionstrategy -n driveby
kubectl get changetransferpolicies -n driveby      # Auto-created by PromotionStrategy controller
```

## Configuration Model

DriveBy uses a **three-tier configuration model** so that every hardcoded value can be overridden without editing chart source:

| Tier | Resolved at | Scope | How to set |
|------|-------------|-------|------------|
| `values.yaml` | `helm template` time | Chart-wide defaults | `--set defaults.X=Y` or custom values file |
| EnvironmentConfig | Crossplane composition runtime | Per-cluster overrides | `kubectl apply -f environment-configs.yaml` |
| XRD spec fields | Crossplane composition runtime | Per-CR overrides | Set directly in `XQualityGateTemplate` or `XQualityGate` spec |

**Resolution order** (highest priority wins): XRD spec → EnvironmentConfig → values.yaml baked default.

### Configurable Knobs

| Knob | values.yaml path | EnvironmentConfig key | XRD field |
|------|------------------|-----------------------|-----------|
| Image pull secret | `defaults.secrets.imagePullSecret` | `imagePullSecret` | `secretsConfig.imagePullSecret` |
| API auth secret | `defaults.secrets.apiAuth.name` | `apiAuthSecretName` | `secretsConfig.apiAuthSecretName` |
| API key header | `defaults.secrets.apiKeyHeader` | `apiKeyHeader` | `secretsConfig.apiKeyHeader` |
| GitHub API URL | `defaults.github.apiUrl` | `githubApiUrl` | `githubConfig.apiUrl` |
| Workflows UI URL | `defaults.argoWorkflowsUrl` | `workflowsUrl` | `workflowsUrl` |
| curl image | `defaults.images.curl` | `curlImage` | `images.curl` |
| postgres image | `defaults.images.postgres` | `postgresImage` | `images.postgres` |
| Ingress class | `defaults.ingress.className` | `ingressClassName` | `ingressConfig.className` |
| cert-manager issuer | `defaults.ingress.clusterIssuer` | `clusterIssuer` | `ingressConfig.clusterIssuer` |
| Extra ingress annotations | `defaults.ingress.annotations` | — | `ingressConfig.annotations` |
| Webhook port | `defaults.webhookPort` | `webhookPort` | `webhookPort` |
| JetStream version | `defaults.eventBus.jetstream.version` | `jetstreamVersion` | `eventBusConfig.jetstreamVersion` |
| EventBus replicas | `defaults.eventBus.jetstream.replicas` | `eventBusReplicas` | `eventBusConfig.replicas` |
| GitHub App secret | `gitopsPromoter.githubApp.secretName` | `githubAppSecretName` | `promoterConfig.githubApp.secretName` |
| GitHub App ID | `gitopsPromoter.githubApp.appID` | `githubAppID` | `promoterConfig.githubApp.appID` |
| GitHub App Installation ID | `gitopsPromoter.githubApp.installationID` | `githubInstallationID` | `promoterConfig.githubApp.installationID` |

## Customization

### Validation Mode

Control strictness via the template's `validationConfig.validationMode`:
- `minimal` — Basic OpenAPI compliance only (P001)
- `strict` — All static validation principles (P001-P005, P008)
- `test-only` — Functional testing only (P006)

### Multiple APIs

Create additional `XQualityGate` instances for each API:

```yaml
apiVersion: driveby.io/v1alpha1
kind: XQualityGate
metadata:
  name: another-api-staging-gate
  namespace: driveby
spec:
  appName: another-api
  projectName: driveby
  gateName: staging-promotion
  templateRef:
    name: driveby-staging-promotion
  repositoryConfig:
    owner: meter-peter
    name: another-api-gitops
    fullName: meter-peter/another-api-gitops
  apiConfig:
    serviceName: another-api
    sourceEnvironment:
      namespace: another-api-dev
    targetEnvironment:
      namespace: another-api-staging
  # ... rest of config
```

## Troubleshooting

### Webhook not received
1. Check webhook deliveries: `gh api repos/novelcore/perfect-api-gitops/hooks/<id>/deliveries`
2. Verify Ingress resolves: `dig perfect-api-staging-promotion-webhook.private.novelcore.org`
3. Check EventSource pod is running: `kubectl get pods -n driveby -l eventsource-name=perfect-api-staging-promotion-eventsource`
4. Check EventSource logs: `kubectl logs -n driveby -l eventsource-name=perfect-api-staging-promotion-eventsource`

### EventSource/Sensor crash-looping
1. Check EventBus is healthy: `kubectl get eventbus -n driveby`
2. Check for stream creation errors (common: `replicas > 1 not supported` — set `streamConfig.replicas` to match bus replicas)
3. Verify NATS pods are Running: `kubectl get pods -n driveby | grep eventbus`

### Workflow not triggered
1. Verify Sensor is subscribed: `kubectl logs -n driveby -l sensor-name=perfect-api-staging-promotion-sensor`
2. Check the PR action is one of: `opened`, `reopened`, `synchronize`
3. Verify the EventSource received the webhook

### Workflow fails
1. Check workflow status: `kubectl get workflows -n driveby`
2. Get step logs: `kubectl logs -n driveby <pod-name> -c main`
3. Common failures:
   - `github-commit-status`: GitHub App credentials invalid or missing
   - `driveby-validate`: API not reachable from source namespace
   - `health-check-source`: Source environment not ready within 2 minutes
