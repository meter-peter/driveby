# GitOps Pipeline: Promotion PR to Validated Deployment

This document describes the end-to-end automated pipeline that triggers when a promotion pull request is opened on the **gitops repository**. It covers every component in the chain: GitHub webhook, Argo Events, Argo Workflows, DriveBy validation, and feedback to the PR.

## Two-Repo Model

XSDLC v3.0.0 uses a **two-repo model** with **ArgoCD Source Hydrator** that cleanly separates concerns:

| Repo | Purpose | Example |
|------|---------|---------|
| **Software repo** | Source code, Dockerfiles, CI — where the API lives and is built | `novelcore/perfect-api` |
| **GitOps repo** | Dry manifests on `main` (Kustomize base + per-env overlays), auto-created by XSDLC | `novelcore/perfect-api-gitops` |

The software repo is owned by the developer and contains everything needed to build and test the application. The gitops repo is the deployment source-of-truth: dry manifests live on `main` in `dry/base/` and `dry/overlays/<env>/`. The ArgoCD Source Hydrator renders each overlay and writes the hydrated output to the corresponding `environment/<env>-next` branch. Each environment has a pair of branches (`environment/<env>` and `environment/<env>-next`), and ArgoCD syncs from the environment branches. XSDLC auto-creates the gitops repo and its branch structure but does **not** generate any CI workflows — developers update dry manifests on `main` however they choose (manually, CI scripts, automation).

## Pre-Pipeline: Dry Manifests Updated, Hydrator Writes to -next Branches

Before the promotion pipeline triggers, the developer updates dry manifests on the `main` branch of the gitops repo. The ArgoCD Source Hydrator then renders the per-environment Kustomize overlays and writes hydrated manifests to the corresponding `environment/<env>-next` branches. XSDLC does **not** generate any CI workflow — developers choose their own mechanism for updating dry manifests:

- **Manual push:** `git push` changes to dry manifests on `main`
- **CI script:** A step in the software repo's CI pipeline updates image tags or overlay configuration on `main`
- **Automation tool:** Renovate, or any tool that commits to the gitops repo's `main` branch

```
 Developer updates dry manifests on main (dry/base/ or dry/overlays/<env>/)
   (via manual push, CI script, or any automation)
         |
         v
 ArgoCD Source Hydrator renders the per-env Kustomize overlay
   → writes hydrated manifests to environment/<env>-next branch (manifests/ + hydrator.metadata)
         |
         v
 Promoter detects new commits on <env>-next, opens promotion PR (<env>-next → <env>)
         |
         v
 Webhook fires on the gitops repo → triggers the quality gate pipeline below
```

**Required setup:**
- GitOps repo with dry manifests on `main` and environment branches (auto-created by XSDLC)
- ArgoCD Source Hydrator configured per-environment (auto-configured by XSDLC ArgoCD Applications)
- Developer's own mechanism for updating dry manifests on `main`

## Pipeline Overview

```
 Hydrator renders dry/overlays/<env>/ from main → writes to environment/<env>-next
         |
         v
 Promoter opens promotion PR on gitops repo (environment/dev-next → environment/dev)
         |
         v
 GitHub sends webhook (pull_request event) from the gitops repo
         |
         v
 Traefik Ingress (perfect-api-staging-promotion-webhook.private.novelcore.org)
         |
         v
 Argo Events EventSource (perfect-api-staging-promotion-eventsource, port 12000)
         |
         v
 JetStream EventBus (1-replica NATS cluster)
         |
         v
 Argo Events Sensor (perfect-api-staging-promotion-sensor)
   filters: opened | reopened | synchronize
         |
         v
 Argo Workflow submitted (driveby-staging-promotion)
   parameters: app-name, pr-number, head-sha, github-owner, github-repo,
               source-namespace, target-namespace, service-name, openapi-endpoint
         |
         v
 DAG Pipeline:
   set-pending-status ("Validating dev environment...")
    + update-commitstatus-pending (CommitStatus CRD: phase=pending)
         |
     wait-for-source-ready (health-check dev, max 2min)
         |
     validate-source (driveby validate-only --validation-mode strict against dev)
         |
     functional-test-source (driveby function-only against dev)
         |
     report-success (commit status = "Dev validation passed, safe to promote")
      + comment-pr (post validation report as PR comment)
      + update-commitstatus-success (CommitStatus CRD: phase=success)
         |
     GitOps Promoter sees CommitStatus success → auto-merges promotion PR
         |
     ArgoCD syncs staging environment
```

## Components

### 1. GitHub Webhook

A repository webhook on the **gitops repo** (e.g., `novelcore/perfect-api-gitops`) sends `pull_request` events to the cluster. The webhook targets the gitops repo because promotion PRs — opened by the Promoter when new commits land on `-next` branches — happen there.

- **URL:** `https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation`
- **Content type:** `application/json`
- **Events:** `pull_request`
- **Active:** yes

This webhook is auto-created by Argo Events when the corresponding `XSDLC` is reconciled.
To manually recreate or troubleshoot:

```bash
gh api repos/novelcore/perfect-api-gitops/hooks --method POST \
  -f name=web -F active=true \
  -f 'config[url]=https://perfect-api-staging-promotion-webhook.private.novelcore.org/perfect-api-staging-promotion-pr-validation' \
  -f 'config[content_type]=json' \
  -f 'events[]=pull_request'
```

### 2. Ingress + Service (Crossplane-managed)

Traefik routes external HTTPS traffic to the EventSource pod.

| Resource | Name | Details |
|----------|------|---------|
| Service | `perfect-api-staging-promotion-eventsource-svc` | ClusterIP, port 12000, selects EventSource pods |
| Ingress | `perfect-api-staging-promotion-webhook-ingress` | TLS via cert-manager (`sys-prod-issuer-dns`), `traefik-system` class |

Both are generated by the `XSDLC` Crossplane composition.

### 3. EventBus

JetStream-based NATS server providing event transport between EventSource and Sensor.

- **Name:** `default` (required name for Argo Events auto-discovery)
- **Replicas:** 1 (single node; `streamConfig.replicas` also set to 1 to match)
- **Version:** NATS 2.10.10

Generated by the `XSDLC` Crossplane composition.

### 4. EventSource

Listens for GitHub webhook payloads on HTTP.

- **Name:** `perfect-api-staging-promotion-eventsource`
- **Endpoint:** `/perfect-api-staging-promotion-pr-validation` on port 12000
- **Events:** `pull_request`
- **Webhook auth:** GitHub App credentials (`github-app-credentials`) used by Argo Events to provision the webhook

Generated by the `XSDLC` Crossplane composition.

### 5. Sensor

Subscribes to EventSource events, filters for PR actions, and submits Argo Workflows.

- **Name:** `perfect-api-staging-promotion-sensor`
- **Dependency:** EventSource event name `pr-validation`
- **Filter:** `body.action` in `[opened, reopened, synchronize]`
- **Trigger:** submits a Workflow referencing `driveby-staging-promotion` WorkflowTemplate

Parameter extraction from webhook payload:

| Parameter | Source | Example |
|-----------|--------|---------|
| `app-name` | static (from XSDLC spec) | `perfect-api` |
| `pr-number` | `body.number` | `1` |
| `head-sha` | `body.pull_request.head.sha` | `d6eca53...` |
| `github-owner` | `body.repository.owner.login` | `meter-peter` |
| `github-repo` | `body.repository.name` | `perfect-api-gitops` |
| `source-namespace` | static | `perfect-api-dev` |
| `target-namespace` | static | `perfect-api-staging` |
| `service-name` | static | `perfect-api` |
| `openapi-endpoint` | static | `/openapi.json` |
| `app-port` | static | `8000` |

Generated by the `XSDLC` Crossplane composition.

### 6. Promotion Pipeline Workflow

The `driveby-staging-promotion` WorkflowTemplate is a DAG pipeline with 5 steps:

```
set-pending-status ──────────────────────┐
update-commitstatus-pending ─────────────┤ (parallel)
                                         ├──> wait-for-source-ready ──> validate-source ──> functional-test-source ──> report-success
                                         │                                                                             ├──> comment-pr
                                         │                                                                             └──> update-commitstatus-success
```

| Step | Template | Purpose |
|------|----------|---------|
| set-pending-status | `github-commit-status` | Sets GitHub commit status to "pending" ("Validating dev environment...") |
| wait-for-source-ready | `health-check-source` | Health-check loop against source env (max 2min, 5s interval) |
| validate-source | `driveby-validate` | Runs `driveby validate-only --validation-mode strict` against source env (dev) |
| functional-test-source | `driveby-functional` | Runs `driveby function-only` against source env (dev) |
| report-success | `github-commit-status` | Sets commit status to "success" ("Dev validation passed, safe to promote") |
| comment-pr | `github-pr-comment` | Posts a DriveBy validation report as a PR comment |
| update-commitstatus-pending | `update-commitstatus` | Creates CommitStatus CRD with phase=pending |
| update-commitstatus-success | `update-commitstatus` | Creates CommitStatus CRD with phase=success |

**Exit handler:** On failure, sets commit status to "failure" and creates CommitStatus CRD with phase=failure.

All validation targets the **source environment** (dev), not the target (staging). The source env must be deployed and healthy before promotion.

Generated by the `XSDLC` Crossplane composition.

### 7. DriveBy Container Image

The workflow steps use `ghcr.io/meter-peter/driveby:latest` as the container image.

- Built from `driveby-cli/Dockerfile`
- Contains the compiled `driveby` binary
- Pulled using `ghcr-creds` imagePullSecret in the `driveby` namespace

## Crossplane Resource Model

All Argo Events and Workflow resources are generated by a single Crossplane composition:

| XRD | Generates | Purpose |
|-----|-----------|---------|
| `XSDLC` | ServiceAccount, EventBus, Role, RoleBinding, WorkflowTemplate, EventSource, Sensor, Service, Ingress, ScmProvider, GitRepository, PromotionStrategy, ArgoCDCommitStatus | All quality gate resources from a single CR |

This means adding a quality gate for a new API is a single `XSDLC` CR — no manual YAML.

### GitOps Promoter Loop

The XSDLC composition generates promoter resources that close the automation loop:

| Resource | Purpose |
|----------|---------|
| ScmProvider | GitHub App auth for promoter SCM operations |
| GitRepository | Points promoter to the **gitops repo** (e.g., `novelcore/perfect-api-gitops`) |
| PromotionStrategy | Defines environment chain (dev → staging → prod) with commit status gates |
| ArgoCDCommitStatus | Aggregates ArgoCD app health into a CommitStatus |

The Argo Workflow's `update-commitstatus` step creates CommitStatus CRDs that the promoter watches. When the CommitStatus phase transitions to `success`, the promoter auto-merges the pending promotion PR (if `autoMerge: true`).

## Secrets

| Secret | Namespace | Purpose |
|--------|-----------|---------|
| `github-app-credentials` | driveby | GitHub App credentials for all GitHub operations: commit status, PR comments, promoter SCM access (appID, installationID, privateKey) |
| `driveby-api-auth` | driveby | API key and header name for authenticating to the target API |
| `ghcr-creds` | driveby | Docker registry credentials for pulling images from ghcr.io |

## Thesis Mapping

This pipeline directly implements the concepts from:
- **Chapter 4 (Architecture):** Namespace separation, ArgoCD application model, event-driven trigger architecture
- **Chapter 5 (End-to-End Workflow):** The complete GitOps feedback loop from promotion PR to validated deployment
- **Chapter 6 (Evaluation):** The controlled evaluation (perfect-api) runs through this exact pipeline
