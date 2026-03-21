# Results Directory — DriveBy Validation Test Results

## Purpose
Stores JSON output from every DriveBy validation run. These results serve as the ground truth for thesis evaluation (Chapter 6) and provide a historical record of gate behavior across the 5-API test suite.

## Naming Convention
```
<api-name>-<test-type>.json
```

Where `<test-type>` is one of:
- `validate-strict` — `validate-only --validation-mode strict`
- `validate-test-ready` — `validate-only --validation-mode test-ready`
- `validate-minimal` — `validate-only --validation-mode minimal`
- `functional` — `function-only`
- `loadtest` — `load-only`

## Current Results (2026-03-21)

| File | API | Test | Result |
|------|-----|------|--------|
| `perfect-api-validate-strict.json` | perfect-api | validate-only (strict) | passed (3/6) |
| `perfect-api-functional.json` | perfect-api | function-only | passed (8/8 endpoints) |
| `perfect-api-loadtest.json` | perfect-api | load-only (10 users, 15s, 200ms P95) | passed (P95: ~7ms) |
| `bad-docs-api-validate-strict.json` | bad-docs-api | validate-only (strict) | passed (2/6, warnings only) |
| `no-auth-api-validate-strict.json` | no-auth-api | validate-only (strict) | **failed** (P005 critical) |
| `slow-api-validate-strict.json` | slow-api | validate-only (strict) | passed (3/6) |
| `slow-api-loadtest.json` | slow-api | load-only (10 users, 15s, 200ms P95) | **failed** (P95: ~505ms) |
| `broken-api-validate-strict.json` | broken-api | validate-only (strict) | passed (3/6) |
| `broken-api-functional.json` | broken-api | function-only | **failed** (2/8 endpoints) |

## Rules
- Every DriveBy test result gets saved here
- Use pretty-printed JSON (`python3 -m json.tool`)
- File names must follow the naming convention above
- Results are additive — keep historical results, don't overwrite unless re-running the same exact test
