# Results Directory — DriveBy Validation Test Results

## Purpose
Stores all DriveBy validation results: local CLI tests, cluster gate captures, edge case tests, and thesis summary documents. These serve as the ground truth for thesis evaluation (Chapter 6).

## Directory Structure

```
results/
├── CLAUDE.md                       ← This file
├── local/                          ← Local CLI test results (9 JSON files)
├── cluster/                        ← Cluster gate captures
│   ├── driveby-namespace-snapshot.txt
│   ├── all-workflows.json
│   ├── all-commitstatuses.json
│   ├── all-argocd-apps.json
│   ├── perfect-api/               ← Infrastructure + gate results
│   │   ├── xsdlc-status.json
│   │   ├── argocd-apps.json
│   │   ├── promoter-resources.json
│   │   ├── events-infra.json
│   │   ├── workflow-templates.json
│   │   ├── ingresses.json
│   │   ├── staging-gate/          ← PASS (validate + functional)
│   │   └── prod-gate/             ← PASS (validate + load-test, manual merge)
│   ├── bad-docs-api/              ← PASS staging (warnings only), PASS prod
│   ├── no-auth-api/               ← FAIL staging (P005 critical)
│   ├── slow-api/                  ← PASS staging, FAIL prod (load-test P95>200ms)
│   └── broken-api/                ← FAIL staging (P006 functional)
├── edge-cases/                     ← Edge case test results (7 tests)
│   ├── webhook-redeliver.json
│   ├── workflow-rerun.json
│   ├── health-check-timeout.json
│   ├── manual-approval.json
│   ├── branch-protection.json
│   ├── commitstatus-lifecycle.json
│   └── argocd-sync-after-merge.json
└── summary/                        ← Thesis summary documents
    ├── verification-matrix.md
    ├── gate-behavior.md
    └── thesis-insights.md
```

## Gate Results Summary (2026-03-29, v3.1.0)

| API | Staging Gate | Staging Result | Prod Gate | Prod Result |
|-----|-------------|---------------|-----------|-------------|
| perfect-api | validate-only (test-ready) + functional-test | **PASS** | load-test (auto-injected validate) | **PASS** (manual merge) |
| slow-api | validate-only (strict) + functional-test | **PASS** | validate-only + load-test (10u/15s/200ms) | **FAIL** (P95>200ms) |
| bad-docs-api | validate-only (strict) | **PASS** (0 critical, 4 warnings) | validate-only + load-test | **PASS** |
| no-auth-api | validate-only (strict) | **FAIL** (P005 critical) | N/A | NOT_REACHED |
| broken-api | validate-only (strict) + functional-test | **FAIL** (P006 functional) | N/A | NOT_REACHED |

Notes:
- **test-ready** mode checks P001, P002 (no contact/license), P003 (4xx + error schemas), P004, P009 — 5 principles
- **strict** mode checks P001, P002, P003, P004, P005, P008 — 6 principles
- bad-docs-api staging now PASSES (v3.1.0 confirms P005 passes — security schemes are present)

## Local CLI Results (in `local/`)

| File | API | Test | Result |
|------|-----|------|--------|
| `perfect-api-validate-strict.json` | perfect-api | validate-only (strict) | passed (6/6) |
| `perfect-api-validate-test-ready.json` | perfect-api | validate-only (test-ready) | passed (5/5) |
| `perfect-api-functional.json` | perfect-api | function-only | passed (8/8 endpoints) |
| `perfect-api-loadtest.json` | perfect-api | load-only | passed (P95: ~7ms) |
| `bad-docs-api-validate-strict.json` | bad-docs-api | validate-only (strict) | passed (2/6, warnings only) |
| `no-auth-api-validate-strict.json` | no-auth-api | validate-only (strict) | **failed** (P005 critical) |
| `slow-api-validate-strict.json` | slow-api | validate-only (strict) | passed (3/6) |
| `slow-api-loadtest.json` | slow-api | load-only | **failed** (P95: ~505ms) |
| `broken-api-validate-strict.json` | broken-api | validate-only (strict) | passed (3/6) |
| `broken-api-functional.json` | broken-api | function-only | **failed** (2/8 endpoints) |

## Rules
- Every DriveBy test result gets saved here
- Use pretty-printed JSON (`python3 -m json.tool`)
- Cluster results include workflow JSONs, pod logs, PR details, and CommitStatus CRDs
- Results are additive — keep historical results
