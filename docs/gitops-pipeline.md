# GitOps Pipeline: PR to Validated Deployment

This document describes the end-to-end automated pipeline that triggers when a pull request is opened on the `meter-peter/perfect-api` repository. It covers every component in the chain: GitHub webhook, Argo Events, Argo Workflows, DriveBy validation, and feedback to the PR.

## Pipeline Overview

```
 Developer opens PR on perfect-api
         |
         v
 GitHub sends webhook (pull_request event)
         |
         v
 Traefik Ingress (driveby-webhook.private.novelcore.org)
         |
         v
 Argo Events EventSource (driveby-github, port 12000)
         |
         v
 JetStream EventBus (3-replica NATS cluster)
         |
         v
 Argo Events Sensor (driveby-sensor)
   filters: opened | reopened | synchronize
         |
         v
 Argo Workflow submitted (driveby-staging-promotion)
   parameters: app-name, pr-number, head-sha, github-owner, github-repo
         |
         v
 DAG Pipeline:
   set-pending-status
         |
     db-sync (pg-dump-restore prod -> staging)
         |
     spec-sync-check (compare staging vs prod OpenAPI)
         |
     validate-staging (driveby validate-only --validation-mode strict)
         |
     functional-test (driveby function-only against staging)
         |
     report-success (set commit status = success)
      + comment-pr (post validation report as PR comment)
```

## Components

### 1. GitHub Webhook

A repository webhook on `meter-peter/perfect-api` sends `pull_request` events to the cluster.

- **URL:** `https://driveby-webhook.private.novelcore.org/github/driveby`
- **Content type:** `application/json`
- **Events:** `pull_request`
- **Active:** yes

The webhook is created manually via `gh api` (not auto-managed by Argo Events). To recreate:

```bash
gh api repos/meter-peter/perfect-api/hooks --method POST \
  -f name=web \
  -f 'config[url]=https://driveby-webhook.private.novelcore.org/github/driveby' \
  -f 'config[content_type]=json' \
  -f 'config[insecure_ssl]=0' \
  -F 'events[]=pull_request' \
  -F 'active=true'
```

### 2. Ingress + Service

Traefik routes external HTTPS traffic to the EventSource pod.

| Resource | Name | Details |
|----------|------|---------|
| Service | `driveby-github-eventsource-svc` | ClusterIP, port 12000, selects EventSource pods |
| Ingress | `driveby-webhook` | TLS via cert-manager (`sys-prod-issuer-dns`), host `driveby-webhook.private.novelcore.org` |

Manifest: `kubernetes/manifests/webhook-ingress.yaml`

### 3. EventBus

JetStream-based NATS cluster providing event transport between EventSource and Sensor.

- **Name:** `default` (required name for Argo Events auto-discovery)
- **Replicas:** 3 (NATS requires clustered mode for JetStream streams with replicas > 1)
- **Version:** NATS 2.10.10

Manifest: `kubernetes/manifests/eventbus.yaml`

**Important:** The EventBus must be applied and healthy before EventSource or Sensor resources. If NATS pods aren't clustered, EventSources will fail with `replicas > 1 not supported in non-clustered mode`.

### 4. EventSource

Listens for GitHub webhook payloads on HTTP.

- **Name:** `driveby-github`
- **Endpoint:** `/github/driveby` on port 12000
- **Events:** `pull_request`, `push`
- **API Token:** reads from secret `github-pat` (key: `token`) for webhook signature verification

Manifest: `kubernetes/manifests/eventsource.yaml`

When a webhook arrives, the EventSource publishes an event to the JetStream bus on subject `default.driveby-github.driveby-webhook`.

### 5. Sensor

Subscribes to EventSource events, filters for PR actions, and submits Argo Workflows.

- **Name:** `driveby-sensor`
- **Dependency:** `driveby-github` EventSource, event name `driveby-webhook`
- **Filter:** `body.action` in `[opened, reopened, synchronize]`
- **Trigger:** submits a Workflow referencing `driveby-staging-promotion` WorkflowTemplate

Parameter extraction from webhook payload:

| Parameter | JSON path | Example |
|-----------|-----------|---------|
| `app-name` | hardcoded | `perfect-api` |
| `pr-number` | `body.number` | `1` |
| `head-sha` | `body.pull_request.head.sha` | `d6eca53...` |
| `github-owner` | `body.repository.owner.login` | `meter-peter` |
| `github-repo` | `body.repository.name` | `perfect-api` |

Manifest: `kubernetes/manifests/sensor.yaml`

### 6. Staging Promotion Workflow

The `driveby-staging-promotion` WorkflowTemplate is a DAG pipeline with 7 steps:

```
set-pending-status ──> db-sync ──> spec-sync-check ──> validate-staging ──> functional-test ──> report-success
                                                                                               └──> comment-pr
```

| Step | Template | Purpose |
|------|----------|---------|
| set-pending-status | `github-commit-status` | Sets GitHub commit status to "pending" |
| db-sync | `pg-dump-restore` | Syncs production DB to staging |
| spec-sync-check | `spec-comparison` | Compares staging OpenAPI spec against prod |
| validate-staging | `driveby-validate` | Runs `driveby validate-only --validation-mode strict` against staging |
| functional-test | `driveby-functional` | Runs `driveby function-only` against staging |
| report-success | `github-commit-status` | Sets commit status to "success" if all checks pass |
| comment-pr | `github-pr-comment` | Posts a DriveBy validation report as a PR comment |

**Exit handler:** On failure, sets commit status to "failure" with an error description.

Manifest: `kubernetes/manifests/staging-promotion.yaml`

### 7. DriveBy Container Image

The workflow steps use `ghcr.io/meter-peter/driveby:latest` as the container image.

- Built from `driveby-cli/Dockerfile`
- Contains the compiled `driveby` binary
- Pulled using `ghcr-creds` imagePullSecret in the `driveby` namespace

## Secrets

| Secret | Namespace | Purpose |
|--------|-----------|---------|
| `github-pat` | driveby | GitHub PAT for commit statuses, PR comments, and webhook verification |
| `driveby-api-auth` | driveby | API key and header name for authenticating to the target API |
| `ghcr-creds` | driveby | Docker registry credentials for pulling images from ghcr.io |

## Namespace Layout

```
driveby/                          # Workflow infrastructure
  EventBus, EventSource, Sensor
  WorkflowTemplates
  ServiceAccount + RBAC
  Secrets

perfect-api-staging/              # Staging environment (ArgoCD autoSync)
  Deployment, Service, ConfigMap
  Ingress, Secrets

perfect-api-prod/                 # Production environment (ArgoCD manual sync)
  Deployment, Service, ConfigMap
  Ingress, Secrets
```

## Troubleshooting

### Webhook not received
1. Check webhook deliveries: `gh api repos/meter-peter/perfect-api/hooks/<id>/deliveries`
2. Verify Ingress resolves: `curl -sk https://driveby-webhook.private.novelcore.org/`
3. Check EventSource pod is running: `kubectl get pods -n driveby -l eventsource-name=driveby-github`
4. Check EventSource logs for incoming requests: `kubectl logs -n driveby -l eventsource-name=driveby-github`

### EventSource/Sensor crash-looping
1. Check EventBus is healthy: `kubectl get eventbus -n driveby`
2. Verify all 3 NATS pods are Running: `kubectl get pods -n driveby | grep eventbus`
3. Check for stream creation errors in EventSource logs (common: `replicas > 1 not supported`)
4. If NATS cluster won't form: delete EventBus, wait for pods to terminate, recreate

### Workflow not triggered
1. Verify Sensor is subscribed: `kubectl logs -n driveby -l sensor-name=driveby-sensor` (look for "Subscribing to subject")
2. Check the PR action is one of: `opened`, `reopened`, `synchronize`
3. Verify the EventSource received the webhook (check logs for the HTTP request)

### Workflow fails
1. Check workflow status: `kubectl get workflows -n driveby`
2. Get detailed node status: `kubectl get workflow <name> -n driveby -o yaml`
3. Check individual step logs: `kubectl logs -n driveby <pod-name> -c main`
4. Common failures:
   - `github-commit-status`: PAT token expired or lacks `repo:status` scope
   - `driveby-validate`: API not reachable from staging namespace
   - `pg-dump-restore`: Database connectivity issues

## Thesis Mapping

This pipeline directly implements the concepts from:
- **Chapter 4 (Architecture):** Namespace separation, ArgoCD application model, event-driven trigger architecture
- **Chapter 5 (End-to-End Workflow):** The complete GitOps feedback loop from developer PR to validated deployment
- **Chapter 6 (Evaluation):** The controlled evaluation (perfect-api) runs through this exact pipeline
