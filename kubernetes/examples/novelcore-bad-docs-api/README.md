# Novelcore Bad-Docs-API — XSDLC Failure Case: Documentation Gaps

Negative test case: an API with incomplete documentation that fails static validation at the staging gate. Demonstrates how DDT catches documentation quality issues (P002 + P003) before any runtime testing is attempted.

## The Defect

The `bad-docs-api` OpenAPI specification has:
- **Missing descriptions** on endpoints and parameters (fails P002 — Documentation Quality)
- **Missing error responses** for standard HTTP error codes (fails P003 — Error Handling)

The API itself works correctly at runtime, but its specification is incomplete.

## How It Works

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
    |
    +---> Hydrator renders dry/overlays/dev -> environment/dev-next
    |     +-- Promoter auto-merges dev-next -> dev (no gate) -> ArgoCD syncs dev
    |
    +---> Hydrator renders dry/overlays/staging -> environment/staging-next
    |     +-- Promoter PRs staging-next -> staging -> Gate fires
    |     +-- validate-only (strict) -> FAILS (P002 + P003)
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
  name: bad-docs-api
  namespace: driveby
spec:
  repository:
    owner: novelcore
    name: bad-docs-api
  gitopsRepository:
    name: bad-docs-api-gitops

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
- GitOps Repository `bad-docs-api-gitops` (auto-created via `provider-upjet-github`)
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `bad-docs-api-github` (GitHub App auth)
- GitRepository `bad-docs-api-gitops` (promoter repo reference)
- PromotionStrategy `bad-docs-api-promotion` (3-env chain with commit status gates)

**Per-environment (x3):**
- ArgoCD Application (`bad-docs-api-dev`, `bad-docs-api-staging`, `bad-docs-api-prod`) — all autoSync with sourceHydrator

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding, WorkflowTemplate, EventSource, Sensor, Service + Ingress, BranchProtection

## Verification Matrix

| Gate | Check | Result | Reason |
|------|-------|--------|--------|
| staging | `validate-only` (strict) | FAIL | P002 (Documentation Quality) + P003 (Error Handling) fail — missing descriptions and error responses |
| prod | `validate-only` | NEVER REACHED | Staging gate failure blocks promotion |
| prod | `load-test` | NEVER REACHED | Staging gate failure blocks promotion |

## Key Insight

Static validation catches documentation gaps **before any runtime testing is attempted**. The API may work perfectly at runtime, but without complete documentation, DDT's Completeness axiom (P002, P003) prevents promotion. This enforces the DDT principle that the specification is the contract — if the contract is incomplete, the API is not ready for promotion regardless of runtime behavior.

## Apply

```bash
kubectl apply -f xsdlc-bad-docs-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# GitOps repo created
gh repo view novelcore/bad-docs-api-gitops

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby | grep bad-docs-api

# Watch the staging gate fail
kubectl get workflows -n driveby | grep bad-docs-api-staging
```

## Uninstall

```bash
kubectl delete xsdlc bad-docs-api -n driveby
```
