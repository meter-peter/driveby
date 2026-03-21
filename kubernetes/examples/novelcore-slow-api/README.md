# Novelcore Slow-API — XSDLC Failure Case: Performance Under Load

**The most interesting case for the thesis.** A well-documented, functionally correct API that only fails under concurrent load at the prod gate. Demonstrates a defect class that is invisible to both static validation and single-request functional testing — only DDT's layered validation approach catches it.

## The Defect

The `slow-api` has:
- A **complete, well-documented OpenAPI specification** (passes all static checks)
- **Correct functional behavior** (every endpoint returns the right status codes and schemas)
- A **hidden performance bottleneck** — response times degrade sharply under concurrent load (e.g., unoptimized database queries, missing connection pooling, synchronous blocking)

Single requests complete in ~50ms. Under 10 concurrent users, P95 latency exceeds 200ms.

## How It Works

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
    |
    +---> Hydrator renders dry/overlays/dev -> environment/dev-next
    |     +-- Promoter auto-merges dev-next -> dev (no gate) -> ArgoCD syncs dev
    |
    +---> Hydrator renders dry/overlays/staging -> environment/staging-next
    |     +-- Promoter PRs staging-next -> staging -> Gate fires
    |     +-- validate-only (strict) -> PASS (spec is complete)
    |     +-- functional-test -> PASS (single-request behavior is correct)
    |     +-- Auto-merge on success -> ArgoCD syncs staging
    |
    +---> Hydrator renders dry/overlays/prod -> environment/prod-next
          +-- Promoter PRs prod-next -> prod -> Gate fires
          +-- validate-only -> PASS
          +-- load-test (10 users, 15s, P95 < 200ms) -> FAILS
          +-- PR blocked, DriveBy report posted as comment
```

## The Single CR

```yaml
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: slow-api
  namespace: driveby
spec:
  repository:
    owner: novelcore
    name: slow-api
  gitopsRepository:
    name: slow-api-gitops

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
              validationMode: strict
          - type: functional-test

    - name: prod
      autoMerge: false
      gate:
        checks:
          - type: validate-only
          - type: load-test
            loadTestConfig:
              concurrentUsers: 10
              testDuration: "15s"
              maxLatencyP95: "200ms"
              minSuccessRate: 0.99
```

**~35 lines** — generates the entire delivery pipeline.

## What Gets Created

From this single CR, the XSDLC composition generates:

**App-level (shared):**
- GitOps Repository `slow-api-gitops` (auto-created via `provider-upjet-github`)
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `slow-api-github` (GitHub App auth)
- GitRepository `slow-api-gitops` (promoter repo reference)
- PromotionStrategy `slow-api-promotion` (3-env chain with commit status gates)

**Per-environment (x3):**
- ArgoCD Application (`slow-api-dev`, `slow-api-staging`, `slow-api-prod`) — all autoSync with sourceHydrator

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding, WorkflowTemplate, EventSource, Sensor, Service + Ingress, BranchProtection

## Verification Matrix

| Gate | Check | Result | Reason |
|------|-------|--------|--------|
| staging | `validate-only` (strict) | PASS | Spec is complete — all principles (P001-P005, P008-P009) pass |
| staging | `functional-test` | PASS | Single-request behavior is correct — all endpoints return expected status codes and schemas |
| prod | `validate-only` | PASS | Same spec, still valid |
| prod | `load-test` | **FAIL** | P95 latency exceeds 200ms under 10 concurrent users; performance degrades under load |

## Why This Matters for the Thesis

This case proves the necessity of DDT's **layered validation approach**:

1. **Static validation alone is insufficient** — the specification is perfect, all documentation principles pass. A framework relying only on spec validation would promote this API to production.

2. **Functional testing alone is insufficient** — every endpoint behaves correctly when tested with single sequential requests. Traditional API testing would give this API a clean bill of health.

3. **Only load testing reveals the defect** — the performance bottleneck manifests exclusively under concurrent load. This is the class of defect that DDT's multi-check gate model is designed to catch.

The progression from `validate-only` (static) to `functional-test` (single-request) to `load-test` (concurrent) demonstrates that each validation layer catches a distinct class of defects. Removing any layer creates a blind spot.

## Apply

```bash
kubectl apply -f xsdlc-slow-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# GitOps repo created
gh repo view novelcore/slow-api-gitops

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby | grep slow-api

# Watch staging gate pass
kubectl get workflows -n driveby | grep slow-api-staging

# Watch prod gate fail on load-test
kubectl get workflows -n driveby | grep slow-api-prod
```

## Uninstall

```bash
kubectl delete xsdlc slow-api -n driveby
```
