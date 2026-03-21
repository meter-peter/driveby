# Novelcore No-Auth-API — XSDLC Failure Case: Missing Security

Negative test case: an API with no security definitions that fails static validation at the staging gate. Demonstrates how DDT catches security gaps (P005) before any runtime testing is attempted.

## The Defect

The `no-auth-api` OpenAPI specification has:
- **No `securitySchemes`** defined in the components section
- **No `security`** requirements on any endpoint

The API functions correctly and returns valid responses, but it exposes all endpoints without any authentication or authorization.

## How It Works

```
Developer updates dry manifests on main branch (dry/base/ or dry/overlays/<env>/)
    |
    +---> Hydrator renders dry/overlays/dev -> environment/dev-next
    |     +-- Promoter auto-merges dev-next -> dev (no gate) -> ArgoCD syncs dev
    |
    +---> Hydrator renders dry/overlays/staging -> environment/staging-next
    |     +-- Promoter PRs staging-next -> staging -> Gate fires
    |     +-- validate-only (strict) -> FAILS (P005)
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
  name: no-auth-api
  namespace: driveby
spec:
  repository:
    owner: novelcore
    name: no-auth-api
  gitopsRepository:
    name: no-auth-api-gitops

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
          - type: functional-test
```

**~30 lines** — generates the entire delivery pipeline.

## What Gets Created

From this single CR, the XSDLC composition generates:

**App-level (shared):**
- GitOps Repository `no-auth-api-gitops` (auto-created via `provider-upjet-github`)
- ServiceAccount `driveby` + imagePullSecrets
- EventBus `default` (JetStream/NATS)
- ScmProvider `no-auth-api-github` (GitHub App auth)
- GitRepository `no-auth-api-gitops` (promoter repo reference)
- PromotionStrategy `no-auth-api-promotion` (3-env chain with commit status gates)

**Per-environment (x3):**
- ArgoCD Application (`no-auth-api-dev`, `no-auth-api-staging`, `no-auth-api-prod`) — all autoSync with sourceHydrator

**Per-gate (x2 for staging + prod):**
- Role + RoleBinding, WorkflowTemplate, EventSource, Sensor, Service + Ingress, BranchProtection

## Verification Matrix

| Gate | Check | Result | Reason |
|------|-------|--------|--------|
| staging | `validate-only` (strict) | FAIL | P005 (Security Standards) fails — no security schemes or security requirements defined |
| prod | `validate-only` | NEVER REACHED | Staging gate failure blocks promotion |
| prod | `functional-test` | NEVER REACHED | Staging gate failure blocks promotion |

## Key Insight

Security is caught by **static validation alone** — no runtime testing is needed to detect the absence of authentication. DDT's Observability axiom (P005) ensures that APIs without security definitions cannot promote past the first gate. This is a fast-feedback loop: the developer sees the P005 failure in the DriveBy report on the staging PR within seconds, not after a lengthy functional test run.

## Apply

```bash
kubectl apply -f xsdlc-no-auth-api.yaml
```

## Verify

```bash
# XSDLC status
kubectl get xsdlcs -n driveby

# GitOps repo created
gh repo view novelcore/no-auth-api-gitops

# WorkflowTemplates (2: staging-gate + prod-gate)
kubectl get workflowtemplates -n driveby | grep no-auth-api

# Watch the staging gate fail
kubectl get workflows -n driveby | grep no-auth-api-staging
```

## Uninstall

```bash
kubectl delete xsdlc no-auth-api -n driveby
```
