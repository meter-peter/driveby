# Gate Behavior Analysis -- XSDLC Quality Gates (2026-03-22)

Documented observations from running 5 APIs through the XSDLC quality gate pipeline on `private.novelcore.org`.

## 1. Severity-Based Blocking

The gate decision is based exclusively on principle severity, not on pass/fail count.

**Critical severity (blocks the gate):**
- P001 OpenAPI Specification Compliance
- P005 Security Standards
- P006 API Contract Testing (functional-test)
- P007 Performance Requirements (load-test -- via exit code)

**Warning severity (does NOT block the gate):**
- P002 API Documentation Quality
- P003 Error Handling Standards
- P004 Request Schema Definitions
- P008 API Versioning Strategy
- P009 Test Readiness

**Evidence from results:**
- slow-api passed staging with 3/6 principles failing (P002, P003, P004 all warnings) because P001, P005, and P008 passed and no critical principle failed
- perfect-api passed prod with 3/6 principles failing (same warning-severity failures) -- only critical principles matter for gate decisions
- bad-docs-api failed staging because P005 (critical) failed, even though P001 (critical) passed
- broken-api passed validate-only (0 critical failures, 4 warnings) but failed on functional-test (P006 critical)

**Key insight:** An API can fail the majority of its principles and still pass the gate, as long as no critical-severity principle fails. This is by design -- warnings produce feedback (PR comments) but do not block promotion.

## 2. Multi-Check Ordering and Short-Circuiting

Gate checks execute as a DAG within an Argo Workflow. Observed behavior:

1. **Health check always runs first** -- `health-check-source` verifies the source environment pod is reachable before any DriveBy checks execute
2. **validate-only runs before functional-test/load-test** -- the composition auto-sorts checks so static validation precedes runtime testing
3. **Short-circuit on validate failure** -- when validate-only exits with code 1 (critical failure), subsequent checks (functional-test, load-test) are skipped entirely

**Evidence:**
- bad-docs-api staging: validate-only failed (P005 critical) at 18:02:33. No `driveby-functional` pod was ever created.
- broken-api staging: validate-only passed (exit code 0, only warnings) at 18:02:44, then functional-test ran at 18:02:54 and failed (P006) at 18:03:10
- perfect-api staging: validate-only passed at 17:58:53, functional-test ran at 17:59:04 and completed at 17:59:05

## 3. Auto-Injection of validate-only

When a gate defines `functional-test` or `load-test` but no explicit `validate-only`, the composition auto-injects a `validate-only` check using `validationDefaults` from the XSDLC spec.

**Evidence:**
- perfect-api prod gate was defined with only `load-test` in the XSDLC spec, but the workflow logs show `driveby-validate` running before `driveby-loadtest`
- slow-api prod gate: same behavior -- `driveby-validate` at 18:02:52, then `driveby-loadtest` at 18:03:02
- The auto-injected validate-only uses the `validationDefaults.validationMode` (strict) from the XSDLC spec

## 4. Exit Handler Feedback

Every workflow has an exit handler that runs regardless of success or failure. The exit handler performs:

1. **CommitStatus update** -- updates the GitHub CommitStatus from `pending` to `success` or `failure`
2. **PR comment** -- posts a detailed comment on the Promoter PR with principle-by-principle results

**Evidence from pod listing:**
- Every workflow run shows these pods:
  - `update-commitstatus` (runs twice: once at start to set `pending`, once in exit handler for final status)
  - `github-commit-status` (in exit handler)
  - `github-pr-comment` (in exit handler)
  - `exit-handler-github` (orchestrates the above)

**On failure (bad-docs-api staging):**
- `driveby-validate` pod: Error status
- `exit-handler-github` pod: Completed (posted failure CommitStatus + failure PR comment)

**On success (perfect-api staging):**
- `driveby-validate` pod: Completed
- `driveby-functional` pod: Completed
- `exit-handler-github` pod: Completed (posted success CommitStatus + success PR comment)

## 5. Source Environment Validation

Each gate validates against the **source** environment (the environment being promoted FROM), not the target:

| Gate | Source Environment | Target Environment |
|------|-------------------|-------------------|
| staging gate | dev namespace (`<api>-dev`) | staging branch |
| prod gate | staging namespace (`<api>-staging`) | prod branch |

**Evidence:**
- perfect-api staging gate loaded spec from: `http://perfect-api.perfect-api-dev:8000/openapi.json`
- perfect-api prod gate loaded spec from: `http://perfect-api.perfect-api-staging:8000/openapi.json`
- slow-api staging gate: `http://slow-api.slow-api-dev:8000/openapi.json`
- slow-api prod gate: `http://slow-api.slow-api-staging:8000/openapi.json`

**Health check implications for NOT_REACHED gates:**
- bad-docs-api prod gate health-checked `http://bad-docs-api.bad-docs-api-staging:8000/openapi.json` -- returned HTTP 000000 for 60 attempts (5 minutes) because no staging deployment existed (staging gate had failed, so nothing was promoted to staging)
- Same behavior for no-auth-api and broken-api prod gates

## 6. CommitStatus Lifecycle

Each gate workflow manages CommitStatus through a defined lifecycle:

```
Promoter creates PR --> Webhook fires --> Sensor triggers workflow
    |
    v
update-commitstatus (pending) --> health-check --> driveby checks --> exit-handler
    |                                                                      |
    v                                                                      v
CommitStatus: pending                                      CommitStatus: success/failure
                                                           + PR comment with details
```

**Two `update-commitstatus` pods per workflow:**
1. First runs at workflow start -- sets CommitStatus to `pending`
2. Second runs in exit handler -- sets CommitStatus to `success` or `failure`

## 7. autoMerge vs autoSync

These are two independent controls:

| Setting | Controls | Scope |
|---------|----------|-------|
| `autoMerge` (XSDLC spec) | Whether Promoter automatically merges the PR after all CommitStatuses pass | Per-environment, Promoter behavior |
| `autoSync` (ArgoCD) | Whether ArgoCD automatically syncs when the branch is updated | Always on for all apps, with selfHeal |

**Evidence:**
- perfect-api prod: `autoMerge: false` -- the Promoter PR passed all gates but stayed open, requiring manual merge. ArgoCD was configured with autoSync + selfHeal.
- All other environments: `autoMerge: true` -- successful gates trigger automatic PR merge by Promoter
- Result: perfect-api prod gate PASSED but PR was not auto-merged (result.txt: "PR stays open (autoMerge: false)")

## 8. Health Check Behavior

The `health-check-source` step runs before any DriveBy checks:

- **Target:** OpenAPI endpoint of the source environment API
- **Retry policy:** 60 attempts, 5 seconds apart (5-minute timeout)
- **Success condition:** HTTP 200 from the OpenAPI endpoint
- **Failure mode:** After 60 failed attempts, exits with code 1, triggering the exit handler

**Evidence of health check timeout (3 APIs):**
- bad-docs-api prod: 60/60 attempts failed (HTTP 000000) -- no staging deployment
- no-auth-api prod: 60/60 attempts failed (HTTP 000000) -- no staging deployment
- broken-api prod: 60/60 attempts failed (HTTP 000000) -- no staging deployment

**Evidence of health check success:**
- perfect-api staging: health check passed quickly (no retry logs visible, went straight to validate)
- slow-api staging: same -- immediate success

## 9. Workflow Timing

Observed durations from workflow logs:

| API | Gate | Health Check | Validate | Functional/Load | Total (approx) |
|-----|------|-------------|----------|-----------------|----------------|
| perfect-api | staging | <2s | ~2s | ~2s (functional) | ~12s |
| perfect-api | prod | <2s | ~3s | ~18s (load-test, 15s duration) | ~28s |
| slow-api | staging | <2s | ~2s | ~6s (functional, 500ms delay per endpoint) | ~18s |
| slow-api | prod | <2s | ~2s | ~17s (load-test, 15s duration) | ~28s |
| bad-docs-api | staging | <2s | ~2s (failed) | skipped | ~10s |
| no-auth-api | staging | <2s | ~2s (failed) | N/A | ~10s |
| broken-api | staging | <2s | ~3s | ~17s (functional, 15s timeout on one endpoint) | ~28s |
| bad-docs-api | prod | 5 min (timeout) | N/A | N/A | ~5 min |
| no-auth-api | prod | 5 min (timeout) | N/A | N/A | ~5 min |
| broken-api | prod | 5 min (timeout) | N/A | N/A | ~5 min |

Typical successful workflow: 10-30 seconds. Failed health checks: 5 minutes (full timeout).

## 10. Load Test Bug Discovery

During testing, a bug was identified in the load-test exit code handling:

- slow-api prod load-test output: `"exit_code":0` in the JSON output even though the test FAILED (`"status":"failed"`)
- The CLI printed `Error: load testing failed: performance targets not met` and the process exited with code 1 (from cobra)
- This means the JSON `exit_code` field was incorrect (reported 0) but the actual process exit code was correct (1)
- The workflow correctly detected the failure via the process exit code, not the JSON field

This demonstrates DDT's self-testing capability: the framework's own test infrastructure exposed a reporting inconsistency in the CLI.
