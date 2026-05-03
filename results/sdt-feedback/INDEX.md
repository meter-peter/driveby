# SDT Feedback Experiment — Results Index

Experiment described in `thesis/chapters/07-evaluation.tex` § 7.6 (`sec:sdt-feedback-experiment`).
Date: 2026-05-03. DriveBy commit: `main` HEAD as of run.

## Files

| File | Purpose |
|------|---------|
| `petstore-baseline/openapi.json` | Baseline Swagger Petstore spec (fetched from `petstore3.swagger.io`) |
| `petstore-baseline/validation-report-latest.json` | DriveBy strict-mode report on baseline (the agent's sole feedback signal) |
| `petstore-baseline/validation-report-latest.md` | Markdown rendering of the same |
| `petstore-revised/openapi.json` | Revised spec produced by the agent |
| `petstore-revised/agent-changes.md` | Agent's audit trail: every field added/changed, grouped by principle |
| `petstore-revised/validation-report-latest.json` | DriveBy strict-mode report on the revised spec |

## Headline result

| Metric | Baseline | Revised | $\Delta$ |
|---|---|---|---|
| Strict-mode score | 1/6 | 4/6 | +3 |
| Critical principles passing | 1/5 | 3/5 | +2 |
| Warning principles passing | 0/1 | 1/1 | +1 |

## Per-principle outcome

| Principle | Baseline | Revised | Notes |
|---|---|---|---|
| P001 Compliance | PASS | FAIL | Regressed: stale XML example on Pet body |
| P002 Documentation | FAIL | PASS | Operation/schema descriptions, examples added |
| P003 Errors | FAIL | FAIL | Partial: 4xx/5xx wired, but no `components.responses` reuse |
| P004 Schema | FAIL | PASS | Length/format/min/max constraints added |
| P005 Security | FAIL | PASS | Top-level + per-op security wired |
| P008 Versioning | FAIL | PASS | Semver + breaking-changes + migration narrative |
