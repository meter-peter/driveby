# Novelcore Broken-API — XSDLC Failure Case: Implementation Bugs

Negative test case: an API with a perfect OpenAPI specification but a buggy implementation. Passes static validation (the spec is flawless) but fails functional testing at the staging gate. Demonstrates the gap between documentation correctness and runtime correctness.

## The Defect

The `broken-api` has a **perfect OpenAPI specification** (loaded from disk, hand-crafted to be complete), but the implementation has three bugs:

1. **`GET /products`** returns `418 I'm a Teapot` instead of `200 OK`
2. **`POST /products`** returns `200 OK` instead of `201 Created`
3. **`GET /products/{id}`** returns a response body that does not match the declared schema

The spec itself passes every static validation principle. The bugs only surface when DriveBy sends actual requests and compares responses against the specification.

## How It Works

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
    |
    +---> Hydrator renders dry/overlays/dev -> environment/dev-next
    |     +-- Promoter auto-merges dev-next -> dev (no gate) -> ArgoCD syncs dev
    |
    +---> Hydrator renders dry/overlays/staging -> environment/staging-next
    |     +-- Promoter PRs staging-next -> staging -> Gate fires
    |     +-- validate-only (strict) -> PASS (spec is perfect)
    |     +-- functional-test -> FAILS (P006)
    |     +-- PR blocked, DriveBy report posted as comment
    |
    +---> Hydrator renders dry/overlays/prod -> environment/prod-next
          +-- Promoter PRs prod-next -> prod -> Gate NEVER REACHED
          +-- (staging gate failure blocks the pipeline)
```

## The Single CR

```yaml
apiVersion: driveby.io/v1alpha1
kind: XSDLC
metadata:
  name: broken-api
  namespace: driveby
spec:
  repository:
    owner: novelcore
    name: broken-api
  gitopsRepository:
    name: broken-api-gitops

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
```

**~30 lines** — generates the entire delivery pipeline.

## What Gets Created

From this single CR, the XSDLC composition generates:

**App-level (shared):**
- GitOps Repository `broken-api-gitops` (auto-created via `provider-upjet-github`)
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `broken-api-github` (GitHub App auth)
- GitRepository `broken-api-gitops` (promoter repo reference)
- PromotionStrategy `broken-api-promotion` (3-env chain with commit status gates)

**Per-environment (x3):**
- ArgoCD Application (`broken-api-dev`, `broken-api-staging`, `broken-api-prod`) — all autoSync with sourceHydrator

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding, WorkflowTemplate, EventSource, Sensor, Service + Ingress, BranchProtection

## Verification Matrix

| Gate | Check | Result | Reason |
|------|-------|--------|--------|
| staging | `validate-only` (strict) | PASS | Spec is perfect — loaded from disk, all principles pass |
| staging | `functional-test` | **FAIL** | P006 (Functional Testing) fails — wrong status codes (418, 200 vs 201) and schema mismatches |
| prod | `validate-only` | NEVER REACHED | Staging gate failure blocks promotion |
| prod | `load-test` | NEVER REACHED | Staging gate failure blocks promotion |

## Key Insight

This case demonstrates the critical distinction between **specification correctness** and **implementation correctness**. Static validation (P001-P005, P008-P009) can only verify the specification itself. It cannot detect when the implementation diverges from the specification.

This is why DDT defines the Determinism axiom: the same specification plus the same API must produce the same validation results. When the API's runtime behavior contradicts its specification, functional testing (P006) catches the divergence. The three specific bugs illustrate common implementation errors:

- **Wrong status code** (418 on GET) — catastrophic, the endpoint is non-functional
- **Wrong success code** (200 vs 201 on POST) — subtle, clients relying on 201 semantics break
- **Schema mismatch** (wrong body on GET by ID) — data contract violation

All three are invisible to spec-only validation but immediately caught by functional testing.

## Apply

```bash
kubectl apply -f xsdlc-broken-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# GitOps repo created
gh repo view novelcore/broken-api-gitops

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby | grep broken-api

# Watch the staging gate fail on functional-test
kubectl get workflows -n driveby | grep broken-api-staging
```

## Uninstall

```bash
kubectl delete xsdlc broken-api -n driveby
```
