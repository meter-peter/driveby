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

## Gate Results Summary (2026-03-30, v3.2.0)

| API | Staging Gate | Staging Result | Prod Gate | Prod Result |
|-----|-------------|---------------|-----------|-------------|
| perfect-api | validate-only (test-ready) + functional-test | **PASS** (5/5 + 8/8 endpoints) | load-test (auto-injected validate strict) | **PASS** (6/6 + load PASS) |
| slow-api | validate-only (strict) + functional-test | **PASS** (3/6 strict, zero critical; 8/8 endpoints with ~500ms p95) | validate-only (strict) + load-test | **FAIL** (P007 — P95 502ms vs 200ms target) |
| bad-docs-api | validate-only (strict) | **FAIL** (P002+P003+P004 critical) | N/A | NOT_REACHED |
| no-auth-api | validate-only (strict) | **FAIL** (P005+P002+P003+P004 critical) | N/A | NOT_REACHED |
| broken-api | validate-only (strict) + functional-test | **FAIL** at functional layer (P006 — 2/8 endpoints return undocumented status codes) | N/A | NOT_REACHED |

Notes:
- **v3.2.0**: P002, P003, P004 severity changed from warning to **critical** — validation now blocks incomplete specs from proceeding to functional/load testing
- **test-ready** mode checks P001, P002 (no contact/license), P003 (4xx + error schemas), P004, P009 — 5 principles
- **strict** mode checks P001, P002, P003, P004, P005, P008 — 6 principles
- **Layered defence**: only `perfect-api` (now `non-critical-api` in the thesis) completes the full pipeline. The other four are each blocked at exactly one of the three defence layers: `bad-docs-api`/`no-auth-api` at Layer 1 (static, staging), `broken-api` at Layer 2 (functional, staging), `slow-api` at Layer 3 (load-test, **prod** — not staging)
- Staging-results re-confirmed by the 2026-05-19 live rerun on `private.novelcore.org` (10/10 PRs decided; `slow-api` staging PR#2 **merged** after successful gate). See `results/poc-rerun-2026-05-19/SUMMARY.md`

## Local CLI Results (in `local/`)

| File | API | Test | Result |
|------|-----|------|--------|
| `perfect-api-validate-strict.json` | perfect-api | validate-only (strict) | passed (6/6) |
| `perfect-api-validate-test-ready.json` | perfect-api | validate-only (test-ready) | passed (5/5) |
| `perfect-api-functional.json` | perfect-api | function-only | passed (8/8 endpoints) |
| `perfect-api-loadtest.json` | perfect-api | load-only | passed (P95: ~7ms) |
| `bad-docs-api-validate-strict.json` | bad-docs-api | validate-only (strict) | **failed** (P002+P003+P004 critical) |
| `no-auth-api-validate-strict.json` | no-auth-api | validate-only (strict) | **failed** (P005+P002+P003+P004 critical) |
| `slow-api-validate-strict.json` | slow-api | validate-only (strict) | **failed** (P002+P003+P004 critical) |
| `slow-api-loadtest.json` | slow-api | load-only | **failed** (P95: ~505ms) |
| `broken-api-validate-strict.json` | broken-api | validate-only (strict) | **failed** (P002+P003+P004 critical) |
| `broken-api-functional.json` | broken-api | function-only | **failed** (2/8 endpoints) |

## Rules
- Every DriveBy test result gets saved here
- Use pretty-printed JSON (`python3 -m json.tool`)
- Cluster results include workflow JSONs, pod logs, PR details, and CommitStatus CRDs
- Results are additive — keep historical results
