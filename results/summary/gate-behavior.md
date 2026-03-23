# Gate Behavior Analysis -- XSDLC Quality Gates (Clean Reinstall, 2026-03-22)

Documented observations from running 5 APIs through the XSDLC quality gate pipeline on `private.novelcore.org`. All infrastructure was provisioned from scratch via XSDLC Crossplane composition.

## 1. Severity-Based Blocking

The gate decision is based exclusively on principle severity, not on pass/fail count.

**Critical severity (blocks the gate):**
- P001 OpenAPI Specification Compliance
- P005 Security Standards
- P006 API Contract Testing (functional-test check type)
- P007 Performance Requirements (load-test check type -- via process exit code)

**Warning severity (does NOT block the gate):**
- P002 API Documentation Quality
- P003 Error Handling Standards
- P004 Request Schema Definitions
- P008 API Versioning Strategy
- P009 Test Readiness

**Evidence from results:**
- slow-api passed staging with 3/6 principles failing (P002, P003, P004 -- all warnings) because P001 and P005 passed and P006 functional-test passed (8/8 endpoints)
- perfect-api passed prod with 3/6 principles failing (same warning-severity failures) -- only critical principles matter for gate decisions
- bad-docs-api failed staging because P005 (critical) failed, even though P001 (also critical) passed
- broken-api passed validate-only (0 critical failures, 4 warnings) but failed on the subsequent functional-test step (P006 critical)
- no-auth-api failed staging with 2/6 passing -- P005 (critical) was the sole gate blocker

**Key insight:** An API can fail the majority of its principles and still pass the gate, as long as no critical-severity principle fails. This is by design -- warnings produce feedback (PR comments with "How to Pass This Gate" guidance) but do not block promotion.

## 2. Multi-Check Ordering and Short-Circuiting

Gate checks execute as a DAG within an Argo Workflow. Each check is a DAG step with dependency ordering.

**Observed execution order:**
1. `set-pending-status` + `update-commitstatus-pending` -- set CommitStatus to pending
2. `wait-for-source-ready` (health-check) -- verify source environment is reachable
3. `check-0-validate-only` -- static spec validation
4. `check-1-functional-test` or `check-1-load-test` -- runtime testing (depends on check-0 success)
5. Exit handler: `update-commitstatus-{success|failure}` + `report-{success|failure}-github` + `comment-pr-{success|failure}`

**Short-circuit behavior:**
- When validate-only exits with code 1 (critical failure), subsequent checks are skipped entirely because of DAG dependencies
- When health-check fails, no DriveBy checks execute at all

**Evidence:**
- bad-docs-api staging: validate-only failed (P005 critical, exit code 1) at 18:02:33. No `driveby-functional` pod was created. Logs show: "No driveby-validate/functional/loadtest pods -- pipeline stopped at health-check" for the prod gate.
- broken-api staging: validate-only passed (exit code 0, only warnings) at 18:02:44, then functional-test ran at 18:02:54 and failed (P006) at 18:03:10
- perfect-api staging: validate-only passed at 17:58:53, functional-test ran at 17:59:04 and completed at 17:59:05

## 3. Auto-Injection of validate-only

When a gate defines `functional-test` or `load-test` but no explicit `validate-only`, the composition auto-injects a `validate-only` check as the first DAG step using `validationDefaults` from the XSDLC spec.

**Evidence:**
- perfect-api prod gate was defined with only `load-test` in the XSDLC spec, but the workflow logs show `driveby-validate` running before `driveby-loadtest`:
  - `driveby-validate` at 17:54:02 (loaded spec from `perfect-api-staging:8000`)
  - `driveby-loadtest` at 17:54:12 (started 10 req/s attack)
- slow-api prod gate: same behavior -- `driveby-validate` at 18:02:50, then `driveby-loadtest` at 18:03:00
- The auto-injected validate-only uses `validationDefaults.validationMode` (strict) from the XSDLC spec
- This ensures no runtime tests run against a spec that fails critical static checks

## 4. Exit Handler Feedback Loop

Every workflow has an exit handler that runs regardless of success or failure. The exit handler performs two actions:

1. **CommitStatus update** -- updates the Promoter CommitStatus CRD from `pending` to `success` or `failure`
2. **PR comment** -- posts a detailed DriveBy report as a comment on the Promoter PR

**PR comment improvements verified in clean install:**
- Gate-specific headers with environment icon: "Staging Gate -- bad-docs-api" (red circle for failure, green for pass)
- Correct validation mode displayed (strict/test-ready, not minimal)
- Gate Details table showing environment, check types, validation mode, and clickable workflow link
- "How to Pass This Gate" section on ALL failed gates listing critical blockers and warnings
- Workflow URL points to `workflow.private.novelcore.org`
- Red icon for load-test failures rendered via `--overall-status` flag
- XSDLC status reports showing ENVIRONMENTS=3, GATES=2

**Exit handler pods per workflow:**
- `set-pending-status` / `update-commitstatus-pending` (at workflow start)
- `update-commitstatus-{success|failure}` (in exit handler)
- `report-{success|failure}-github` (in exit handler)
- `comment-pr-{success|failure}` (in exit handler)

## 5. Source Environment Validation

Each gate validates against the **source** environment (the environment being promoted FROM), not the target:

| Gate | Source Environment | Target Environment |
|------|-------------------|-------------------|
| staging gate | dev namespace (`<api>-dev`) | staging branch (gitops repo) |
| prod gate | staging namespace (`<api>-staging`) | prod branch (gitops repo) |

**Evidence from OpenAPI spec loading URLs:**
- perfect-api staging gate: `http://perfect-api.perfect-api-dev:8000/openapi.json`
- perfect-api prod gate: `http://perfect-api.perfect-api-staging:8000/openapi.json`
- slow-api staging gate: `http://slow-api.slow-api-dev:8000/openapi.json`
- slow-api prod gate: `http://slow-api.slow-api-staging:8000/openapi.json`
- bad-docs-api staging gate: `http://bad-docs-api.bad-docs-api-dev:8000/openapi.json`

**Health check implications for NOT_REACHED gates:**
- bad-docs-api prod gate health-checked `http://bad-docs-api.bad-docs-api-staging:8000/openapi.json` -- returned HTTP 000000 for 60 attempts because no staging deployment existed (staging gate had failed, Promoter never merged staging PR, ArgoCD had nothing to sync)
- Same behavior for no-auth-api and broken-api prod gates
- This is the correct cascade: staging failure prevents staging deployment, which causes prod health check to time out, which prevents prod gate from running

## 6. CommitStatus Lifecycle

Each gate workflow manages a CommitStatus CRD through a defined lifecycle:

```
Promoter creates PR --> Webhook fires --> Sensor triggers workflow
    |
    v
set-pending-status + update-commitstatus (pending)
    |
    v
health-check-source (wait-for-source-ready)
    |
    v
driveby-validate --> driveby-functional/loadtest (if applicable)
    |
    v
exit-handler: update-commitstatus (success/failure) + PR comment
```

**CommitStatus CRD naming:** `{gate-name}-{sha-prefix}` (e.g., `staging-gate-0a7a0a3a15ba`)

**Labels for Promoter matching:**
- `promoter.argoproj.io/commit-status: staging-gate` -- required label for Promoter to find the CRD
- `driveby.io/app: perfect-api` -- for filtering by API
- `driveby.io/gate: staging-gate` -- for filtering by gate

**CRD accumulation:** CommitStatus CRDs persist per SHA. The cluster accumulated 32 CRDs across 5 APIs over two test runs (initial deploy + clean reinstall). perfect-api alone has 15 CRDs showing its iterative development history: pending -> failure -> success lifecycle as gate definitions evolved.

## 7. autoMerge vs autoSync

These are two independent controls operating at different layers:

| Setting | Controls | Scope | Layer |
|---------|----------|-------|-------|
| `autoMerge` (XSDLC spec) | Whether Promoter auto-merges the PR after all CommitStatuses pass | Per-environment | Git (Promoter) |
| `autoSync` (ArgoCD) | Whether ArgoCD auto-syncs when the target branch is updated | Always on for all apps, with selfHeal | Kubernetes (ArgoCD) |

**Evidence:**
- perfect-api prod: `autoMerge: false` -- the Promoter PR (#11) passed all gates (CommitStatus: success) but remained open (state: open, mergeable: true, merged: false). Required manual merge.
- All other environments: `autoMerge: true` -- successful staging gates triggered automatic PR merge by Promoter, then ArgoCD auto-synced the staging deployment within seconds.
- After staging PRs were merged (perfect-api, slow-api), ArgoCD synced both applications:
  - perfect-api-staging: Synced/Healthy, finished at 18:03:27Z
  - slow-api-staging: Synced/Healthy, finished at 18:00:28Z

## 8. Health Check Behavior

The `health-check-source` (wait-for-source-ready) step runs before any DriveBy checks:

- **Target:** OpenAPI endpoint of the source environment API (`/openapi.json`)
- **Retry policy:** 60 attempts, 5 seconds apart (5-minute total timeout)
- **Success condition:** HTTP 200 from the OpenAPI endpoint
- **Failure mode:** After 60 failed attempts, exits with code 1, triggering the exit handler with failure CommitStatus

**Evidence of health check timeout (3 APIs):**
- bad-docs-api prod: 60/60 attempts failed (HTTP 000000), started at 17:55:09, timed out at 18:00:11 (5m02s)
- no-auth-api prod: 60/60 attempts failed (HTTP 000000), started at 17:55:25, timed out at 18:00:27 (5m02s)
- broken-api prod: 60/60 attempts failed (HTTP 000000), started at 17:55:54, timed out at 18:00:55 (5m01s)

**Evidence of health check success:**
- perfect-api staging: no retry logs visible -- health check passed on first attempt, went straight to validate
- slow-api staging: same -- immediate success

**ArgoCD selfHeal interaction (edge case):**
- When a deployment was manually scaled to 0 replicas for testing, ArgoCD's selfHeal (automated.selfHeal=true) restored the deployment to 1 replica before the health check could detect the outage. This means health check timeouts can only occur when there is genuinely no deployment (as in the staging-failure cascade), not from transient outages when selfHeal is enabled.

## 9. Workflow Timing

Observed durations from workflow logs:

| API | Gate | Health Check | Validate | Functional/Load | Total (approx) |
|-----|------|-------------|----------|-----------------|----------------|
| perfect-api | staging | <2s | ~2s | ~2s (functional, 8 endpoints) | ~12s |
| perfect-api | prod | <2s | ~3s | ~18s (load-test: 15s duration) | ~28s |
| slow-api | staging | <2s | ~2s | ~6s (functional, 500ms delay/endpoint) | ~18s |
| slow-api | prod | <2s | ~2s | ~17s (load-test: 15s duration) | ~28s |
| bad-docs-api | staging | <2s | ~2s (failed) | skipped | ~10s |
| no-auth-api | staging | <2s | ~2s (failed) | N/A (not configured) | ~10s |
| broken-api | staging | <2s | ~3s | ~17s (functional: 15s from one slow endpoint) | ~28s |
| bad-docs-api | prod | 5m02s (timeout) | N/A | N/A | ~5 min |
| no-auth-api | prod | 5m02s (timeout) | N/A | N/A | ~5 min |
| broken-api | prod | 5m01s (timeout) | N/A | N/A | ~5 min |

Typical successful workflow: 10-30 seconds. Failed health checks: ~5 minutes (full timeout).

## 10. Load Test Exit Code Bug

During testing, an inconsistency was identified in the load-test JSON output:

- slow-api prod load-test output: `"exit_code":0` in the JSON even though the test FAILED (`"status":"failed"`)
- The CLI printed `Error: load testing failed: performance targets not met` and the **process** exited with code 1 (from cobra)
- The JSON `exit_code` field was incorrect (reported 0) but the actual process exit code was correct (1)
- The Argo Workflow correctly detected the failure via the process exit code, not the JSON field

This demonstrates DDT's self-testing capability: the framework's own test infrastructure exposed a reporting inconsistency in the CLI. The P95 message was accurate: "P95 latency (501.828154ms) exceeded target (200ms)".

## 11. Branch Protection Integration

XSDLC auto-provisions GitHub BranchProtection rules for every gated environment:

**Staging protection (all 5 APIs):**
- Required status check: `staging-gate` (strict mode)
- Enforce admins: true
- No force pushes, no deletions

**Prod protection (all 5 APIs):**
- Required status check: `prod-gate` (strict mode)
- Required PR reviews: 1 approving review, dismiss stale reviews
- Enforce admins: true
- No force pushes, no deletions

**Additional:** `*-next` branch protection rules prevent direct pushes to hydrated branches.

Total BranchProtection CRDs: 15 (3 per API: staging, prod, next-branches). All 15 Synced=True, Ready=True.

## 12. Webhook and Workflow Retriggering

**Webhook redelivery:** Redelivering a GitHub webhook for a pull_request event successfully triggers event processing. Tested on bad-docs-api-gitops: redelivered webhook (id: 3810071934675190000) received HTTP 200 from the EventSource ingress.

**PR close/reopen retriggering:** Closing and reopening a Promoter PR retriggers the quality gate workflow via the EventSource/Sensor pipeline. Tested on no-auth-api-gitops: workflow count increased from 5 to 6 after PR reopen. New workflow `no-auth-api-staging-gate-pipeline-hf6xm` completed (Failed, as expected for no-auth-api).

## 13. Composition-Managed Infrastructure

Resources generated per XSDLC CR in the clean install:

| Category | Resources | Count |
|----------|-----------|-------|
| Git | GitOps Repository (provider-upjet-github), base manifests, overlay scaffolds | 1 repo + files |
| Secrets | Push secrets (per gitops repo, for Source Hydrator), ghcr-creds (per namespace) | 4+ |
| ArgoCD | Applications with sourceHydrator config (dev, staging, prod) | 3 |
| Promotion | PromotionStrategy, ArgoCDCommitStatus (per gate), ScmProvider, GitRepository | 5 |
| Events | EventBus, EventSources (per gate), Sensors (per gate) | 5 |
| Workflows | WorkflowTemplates (per gate), Ingresses (per gate webhook) | 4 |
| RBAC | ServiceAccount, Role, RoleBinding | 3 |
| Branch Protection | BranchProtection (per gated env + next branches) | 3 |

**Estimated total: 28+ Kubernetes resources per API from ~35 lines of XSDLC YAML input.**

New in this version:
- Auto-created push secrets per gitops repo (for Source Hydrator write access to the gitops repo)
- Auto-copied ghcr-creds to per-environment namespaces
- Base files managed with `overwriteOnCreate` (composition-managed, not developer-managed)
- Overlay includes namespace field (triggers re-hydration when base changes)
- Extended health-check timeout (10 min max) for non-first gates
- Repository `deletionPolicy: Delete` (removing the XSDLC CR deletes the gitops repo)
- Load-test exit code fix (returns non-zero on failure)
