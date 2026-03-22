# Verification Matrix -- XSDLC Gate Results (2026-03-22)

Complete pass/fail matrix for all 5 evaluation APIs across both staging and production gates.
Data source: Argo Workflow logs captured from `private.novelcore.org` cluster, namespace `driveby`.

## Gate Configuration Per API

| API | Staging Gate Checks | Staging Mode | Prod Gate Checks | Prod Mode |
|-----|-------------------|-------------|-----------------|-----------|
| perfect-api | validate-only, functional-test | test-ready | load-test (auto-injected validate-only) | strict |
| slow-api | validate-only, functional-test | strict | load-test (auto-injected validate-only) | strict |
| bad-docs-api | validate-only | strict | N/A (not reached) | strict |
| no-auth-api | validate-only | strict | N/A (not reached) | strict |
| broken-api | validate-only, functional-test | strict | N/A (not reached) | strict |

## Full Verification Matrix

### Staging Gate

| API | P001 Compliance (critical) | P002 Docs (warning) | P003 Errors (warning) | P004 Schema (warning) | P005 Security (critical) | P006 Functional (critical) | P008 Versioning (warning) | P009 Test Readiness (warning) | Gate Result |
|-----|---------------------------|--------------------|--------------------|---------------------|------------------------|--------------------------|-------------------------|------------------------------|------------|
| perfect-api | PASS | N/A (test-ready) | N/A (test-ready) | PASS | N/A (test-ready) | PASS (8/8 endpoints) | N/A (test-ready) | PASS | **PASS** |
| slow-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | PASS | PASS (8/8 endpoints) | PASS | N/A (strict) | **PASS** |
| bad-docs-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | **FAIL (critical)** | skipped | FAIL (warning) | N/A (strict) | **FAIL** |
| no-auth-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | **FAIL (critical)** | N/A | PASS | N/A (strict) | **FAIL** |
| broken-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | PASS | **FAIL (1/5 failed)** | FAIL (warning) | N/A (strict) | **FAIL** |

### Production Gate

| API | P001 Compliance (critical) | P002 Docs (warning) | P003 Errors (warning) | P004 Schema (warning) | P005 Security (critical) | P007 Performance (warning) | P008 Versioning (warning) | Gate Result |
|-----|---------------------------|--------------------|--------------------|---------------------|------------------------|--------------------------|-------------------------|------------|
| perfect-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | PASS | PASS (P95: 1.8ms, 150 req, 0.7% err) | PASS | **PASS** (manual merge) |
| slow-api | PASS | FAIL (warning) | FAIL (warning) | FAIL (warning) | PASS | **FAIL** (P95: 502ms > 200ms target) | PASS | **FAIL** |
| bad-docs-api | -- | -- | -- | -- | -- | -- | -- | **NOT_REACHED** |
| no-auth-api | -- | -- | -- | -- | -- | -- | -- | **NOT_REACHED** |
| broken-api | -- | -- | -- | -- | -- | -- | -- | **NOT_REACHED** |

Notes:
- "N/A (test-ready)" means the principle is not evaluated in test-ready validation mode (only P001, P004, P009)
- "N/A (strict)" means the principle is not part of the strict validation mode set
- "--" means the gate was never reached because the staging gate failed

## Full Gate Chain: Staging to Production

```
                           STAGING GATE                              PROD GATE
                    +-----------------------+                +-----------------------+
                    | 1. health-check       |                | 1. health-check       |
                    |    (source: dev ns)   |                |    (source: staging ns)|
                    | 2. validate-only      |   on success   | 2. validate-only      |
  PR on staging --> |    (strict/test-ready)|  ----------->  |    (auto-injected)    |
  -next branch      | 3. functional-test    | Promoter       | 3. load-test          |
                    |    (if configured)    | creates PR     |    (if configured)    |
                    | 4. exit-handler       | on prod-next   | 4. exit-handler       |
                    +-----------------------+                +-----------------------+
                           |                                        |
                        CommitStatus                             CommitStatus
                        + PR comment                             + PR comment
```

## Detailed Principle Results Per API

### perfect-api

**Staging gate** (validate-only mode: test-ready + functional-test):
- P001 OpenAPI Compliance: PASS -- fully compliant with 3.0/3.1 standards
- P004 Request Schema Definitions: PASS -- all schemas have type specifications
- P009 Test Readiness: PASS -- all examples, schemas, and response definitions present
- P006 Functional Testing: PASS -- 8/8 endpoints passed (POST /products 201, GET /products 200, GET /products/{id} 404, PUT /products/{id} 404, DELETE /products/{id} 404, POST /tasks 201, POST /test/echo 200, GET /test/health 200)
- Validation score: 3/3 passed, functional: 8/8 endpoints

**Prod gate** (auto-injected validate-only mode: strict + load-test):
- P001 OpenAPI Compliance: PASS
- P002 Documentation Quality: FAIL (warning) -- missing request body examples, enum descriptions, contact/license info
- P003 Error Handling: FAIL (warning) -- missing 5xx docs, missing error detail schemas, no common components
- P004 Request Schema Definitions: FAIL (warning) -- some schemas missing explicit data types
- P005 Security Standards: PASS -- all security schemes properly defined
- P008 Versioning Strategy: PASS -- semantic versioning, deprecation notices present
- P007 Performance: PASS -- P50: 1.6ms, P95: 1.8ms, P99: 5.2ms, 150 requests, 10 req/s, 1 error (0.7%)
- Validation score: 3/6 (0 critical, 3 warnings), load-test: PASS
- PR left open (autoMerge: false), manually merged

### slow-api

**Staging gate** (validate-only mode: strict + functional-test):
- P001 OpenAPI Compliance: PASS
- P002 Documentation Quality: FAIL (warning) -- same issues as perfect-api (shared spec template)
- P003 Error Handling: FAIL (warning) -- same issues as perfect-api
- P004 Request Schema Definitions: FAIL (warning) -- missing explicit data types for some schemas
- P005 Security Standards: PASS
- P008 Versioning Strategy: PASS
- P006 Functional Testing: PASS -- 8/8 endpoints (all ~501ms response time due to artificial delay)
- Validation score: 3/6 (0 critical, 3 warnings), functional: 8/8 endpoints

**Prod gate** (auto-injected validate-only mode: strict + load-test):
- P001-P005, P008: Same as staging (PASS on critical, FAIL on warnings)
- P007 Performance: **FAIL** -- P50: 1.9ms, P95: 501.8ms (target: 200ms), P99: 502.1ms, 150 requests, 1 error (0.7%)
- Validation score: 3/6 (0 critical, 3 warnings), load-test: **FAIL**
- Failure reason: The 500ms artificial delay causes P95 latency (502ms) to exceed the 200ms target

### bad-docs-api

**Staging gate** (validate-only mode: strict):
- P001 OpenAPI Compliance: PASS
- P002 Documentation Quality: FAIL (warning) -- missing operation descriptions, parameter descriptions, request/response examples, schema descriptions, contact/license info
- P003 Error Handling: FAIL (warning) -- missing 4xx/5xx docs, missing error detail schemas
- P004 Request Schema Definitions: FAIL (warning) -- missing string length constraints, numeric min/max values
- P005 Security Standards: **FAIL (critical)** -- no security schemes defined, no global security, no per-operation security
- P008 Versioning Strategy: FAIL (warning) -- no migration guides, no versioning strategy, no breaking changes docs, no compatibility info
- Validation score: 1/6 (1 critical, 4 warnings)
- Functional-test: **skipped** (validate-only failed, so subsequent checks were not executed)

**Prod gate**: NOT_REACHED -- staging gate CommitStatus was "failure", Promoter never promoted to prod

### no-auth-api

**Staging gate** (validate-only mode: strict):
- P001 OpenAPI Compliance: PASS
- P002 Documentation Quality: FAIL (warning) -- missing request body examples, enum descriptions, contact/license info
- P003 Error Handling: FAIL (warning) -- missing 5xx docs for some endpoints, missing error detail schemas
- P004 Request Schema Definitions: FAIL (warning) -- missing explicit data types for some schemas
- P005 Security Standards: **FAIL (critical)** -- no security schemes defined, no global security, 9 endpoints without security
- P008 Versioning Strategy: PASS
- Validation score: 2/6 (1 critical, 3 warnings)
- Functional-test: N/A (not configured for this gate, only validate-only)

**Prod gate**: NOT_REACHED -- staging gate CommitStatus was "failure", health-check timed out (60 attempts, 5 minutes, HTTP 000000 - no staging deployment exists because staging gate failed)

### broken-api

**Staging gate** (validate-only mode: strict + functional-test):
- P001 OpenAPI Compliance: PASS
- P002 Documentation Quality: FAIL (warning) -- missing request/response examples, parameter/schema descriptions, contact/license info
- P003 Error Handling: FAIL (warning) -- missing 5xx docs, inconsistent error formats (two different schemas), missing 4xx for health endpoint
- P004 Request Schema Definitions: FAIL (warning) -- missing string length constraints, numeric min/max
- P005 Security Standards: PASS -- security schemes properly defined
- P008 Versioning Strategy: FAIL (warning) -- missing migration guides, version compatibility
- P006 Functional Testing: **FAIL** -- 4/5 passed, 1/5 failed
  - GET /test/health: PASS (200, 0.9ms)
  - POST /widgets: PASS (500 returned, but response schema mismatch -- missing `error` and `code` fields)
  - GET /widgets: **FAIL** (500, undocumented status code -- "Internal Server Error")
  - DELETE /widgets/{widget_id}: PASS (204, 504ms)
  - GET /widgets/{widget_id}: PASS (200, 15s response time)
- Validation score: 2/6 (0 critical, 4 warnings), functional: **1/5 failed** (P006 critical failure)

**Prod gate**: NOT_REACHED -- health-check timed out at staging namespace (60 attempts, 5 minutes, HTTP 000000)

## Summary Statistics

| Metric | Value |
|--------|-------|
| Total APIs tested | 5 |
| Staging gates passed | 2 (perfect-api, slow-api) |
| Staging gates failed | 3 (bad-docs-api, no-auth-api, broken-api) |
| Prod gates reached | 2 (perfect-api, slow-api) |
| Prod gates passed | 1 (perfect-api) |
| Prod gates failed | 1 (slow-api -- load-test) |
| Full pipeline passed | 1/5 (perfect-api) |
| Blocked by P005 Security (critical) | 2 (bad-docs-api, no-auth-api) |
| Blocked by P006 Functional (critical) | 1 (broken-api) |
| Blocked by P007 Performance | 1 (slow-api) |
| Warning-only failures that did NOT block | P002, P003, P004, P008 across all APIs |
