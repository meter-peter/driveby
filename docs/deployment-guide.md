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

## Namespace Layout

```
Kubernetes Cluster
├── driveby/                     ← DriveBy control plane
│   ├── Crossplane XRDs + Compositions
│   ├── Argo Events (EventBus, EventSource, Sensor)
│   ├── Argo Workflows (WorkflowTemplates, RBAC)
│   └── Secrets (github-pat, driveby-api-auth)
├── perfect-api-staging/         ← Staging environment (per-API)
│   └── API deployment + service
└── perfect-api-prod/            ← Production environment (per-API)
    └── API deployment + service
```

## Install Order

```
1. Secrets           ──► GitHub PAT, API auth, registry creds
       │
2. Helm Install      ──► Providers, Functions, XRDs, Compositions
       │
3. EnvironmentConfigs ──► Cluster-specific settings
       │
4. XQualityGateTemplate ──► RBAC + WorkflowTemplate
       │
5. XQualityGate      ──► EventBus + EventSource + Sensor + Ingress
       │
6. GitHub Webhook    ──► Point repo to Ingress URL
```

## 1. Install the Helm Chart

The chart installs Crossplane providers, functions, XRDs, and compositions.

### From OCI Registry (recommended)

```bash
helm install driveby oci://ghcr.io/meter-peter/charts/driveby \
  --version 0.3.0 \
  --set github.pat=ghp_YOUR_TOKEN \
  --set crossplane.enabled=true
```

### From Local Source

```bash
helm install driveby ./kubernetes/helm/driveby/ \
  --set github.pat=ghp_YOUR_TOKEN \
  --set crossplane.enabled=true
```

This creates:
- **Provider**: `provider-kubernetes` (v0.14.1)
- **Functions**: `function-go-templating`, `function-auto-ready`, `function-sequencer`, `function-environment-configs`
- **ProviderConfig**: `kubernetes-provider` (InjectedIdentity)
- **XRDs**: `xqualitygatetemplates.driveby.io`, `xqualitygates.driveby.io`
- **Compositions**: Template and instance compositions
- **Argo Workflows**: WorkflowTemplates, RBAC
- **Argo Events**: EventSource, Sensor

Wait for providers and functions to become healthy:

```bash
kubectl get providers -w
kubectl get functions -w
```

## 2. Create Secrets

### GitHub PAT (for commit status + PR comments)

If not set via `--set github.pat`:

```bash
kubectl create secret generic github-pat \
  --namespace driveby \
  --from-literal=token=ghp_YOUR_TOKEN
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

## 3. Apply EnvironmentConfigs

EnvironmentConfigs provide cluster-specific settings to the compositions:

```bash
kubectl apply -f kubernetes/examples/crossplane/environment-configs.yaml
```

## 4. Create XQualityGateTemplate

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

## 5. Create XQualityGate

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

## 6. Configure GitHub Webhook

Point your repository's webhook to the Ingress URL:

- **URL**: `https://<app>-<gate>-webhook.<baseDomain>/<app>-<gate>-<trigger>`
- **Content type**: `application/json`
- **Events**: Select the events matching your trigger config (e.g., `pull_request`)

For the default example:
- URL: `https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation`
- Events: Pull requests

## 7. Verification Checklist

```bash
# Crossplane resources healthy
kubectl get providers         # provider-kubernetes: Healthy
kubectl get functions         # All 4 functions: Healthy

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
| GitHub PAT secret name | `defaults.secrets.githubPat.name` | `githubPatSecretName` | `secretsConfig.githubPatSecretName` |
| GitHub PAT secret key | `defaults.secrets.githubPat.key` | `githubPatSecretKey` | `secretsConfig.githubPatSecretKey` |
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

### Example: nginx Ingress (no cert-manager)

```bash
helm install driveby ./kubernetes/helm/driveby/ \
  --set github.pat=ghp_xxx \
  --set defaults.ingress.className=nginx \
  --set defaults.ingress.clusterIssuer=""
```

### Example: Traefik with cert-manager

```bash
helm install driveby ./kubernetes/helm/driveby/ \
  --set github.pat=ghp_xxx \
  --set defaults.ingress.className=traefik-system \
  --set defaults.ingress.clusterIssuer=letsencrypt-prod \
  --set defaults.ingress.annotations."traefik\.ingress\.kubernetes\.io/router\.entrypoints"=websecure \
  --set defaults.ingress.annotations."traefik\.ingress\.kubernetes\.io/router\.tls"=true
```

### Example: Istio (VirtualService via annotations)

```bash
helm install driveby ./kubernetes/helm/driveby/ \
  --set github.pat=ghp_xxx \
  --set defaults.ingress.className=istio \
  --set defaults.ingress.clusterIssuer=""
```

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
  # ... rest of config
```

### Multiple Gates

Create different gate types with separate templates:

```yaml
# Security-focused gate
apiVersion: driveby.io/v1alpha1
kind: XQualityGateTemplate
metadata:
  name: driveby-security-scan
spec:
  projectName: driveby
  gateName: security-scan
  validationConfig:
    validationMode: minimal  # Focus on P005 security
```
