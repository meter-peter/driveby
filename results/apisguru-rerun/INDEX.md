# APIs.guru Large-Scale Rerun — Results Index

Experiment described in `thesis/chapters/07-evaluation.tex` § 7.4 (`sec:wild-eval`).
Date: 2026-05-03. Harvest script: `/tmp/fetch_apisguru.py` (corrected for actual `list.json` schema).
Batch script: `/tmp/batch_validate.py` (parallel-safe). Validation mode: strict.

## Dataset

| Metric | Value |
|---|---|
| Total APIs harvested | 50 |
| OpenAPI 3.0.x | 26 |
| OpenAPI 3.1.0 | 24 |
| Successful validation runs | 50/50 (100%) |
| Failed/timeout runs | 0 |

## Per-principle pass rates (raw counts + percentages)

| Principle | Pass | Fail | Pass rate | Bucket |
|---|---:|---:|---:|---|
| P001 OpenAPI Compliance | 23 | 27 | 46.0% | low |
| P002 Documentation | 0 | 50 | 0.0% | low |
| P003 Error Handling | 0 | 50 | 0.0% | low |
| P004 Schema | 6 | 44 | 12.0% | low |
| P005 Security | 13 | 37 | 26.0% | low |
| P008 Versioning | 0 | 50 | 0.0% | low |

**Bucket thresholds (declared up-front):** high ≥ 80%, medium 30–79%, low < 30%.

## Score distribution (out of 6 evaluated principles in strict mode)

| Score | Count | % |
|---:|---:|---:|
| 0/6 | 16 | 32.0% |
| 1/6 | 26 | 52.0% |
| 2/6 | 8  | 16.0% |
| 3/6 | 0  | 0.0% |
| 4/6 | 0  | 0.0% |
| 5/6 | 0  | 0.0% |
| 6/6 | 0  | 0.0% |

**Headline:** 84% of public APIs score 1/6 or worse in strict mode. Zero APIs pass even half the principles.

## Files

- `apis-3x.csv` — input dataset (name, url, host, openapi_version)
- `summary.json` — per-API outcomes + aggregate per-principle counts
- `<api-name>/validation-report-latest.json` — per-API DriveBy report
