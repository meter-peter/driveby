# Swagger 2.0 Evaluation — Round-6 (2026-05-04)

Evaluation arm added in round 6 of the supervisor revision cycle in
response to comment [465] (Kritikos): *"you seem to focus mainly on
3.X versions. Thus, it would be nice to attempt to also cover 2.X
versions to showcase that your work is fully operational also for
them."*

## Method

1. `tools/harvest-openapi-apisguru.py results/swagger2-evaluation/apis-2.0-raw.csv 2.0`
   — extended harvester accepts `swagger:"2.0"` specs; scans every
   version listed for each API entry (because APIs.guru aggressively
   prefers 3.x as the "preferred" version when an upgraded 3.x spec
   exists). A per-provider cap (2 specs per provider) ensures the
   sample contains 20 distinct providers rather than ~30 sub-APIs of
   a single vendor.
2. DriveBy CLI run in `validate-only --validation-mode strict` against
   each of the 20 harvested URLs. Reports written under
   `results/swagger2-evaluation/reports/`.
3. Aggregated to `results/swagger2-evaluation/per-principle-summary.csv`.

## Results — per-principle pass rate (n = 20)

| Principle | Severity | Pass | Rate |
|---|---|---:|---:|
| P001 — OpenAPI Compliance | Critical | 20/20 | 100% |
| P002 — Documentation Quality | Critical | 0/20 | 0% |
| P003 — Error Handling | Critical | 0/20 | 0% |
| P004 — Schema Definitions | Critical | 1/20 | 5% |
| P005 — Security Standards | Critical | 3/20 | 15% |
| P008 — Versioning Strategy | Warning | 0/20 | 0% |

## Score distribution

| Score (out of 6) | APIs |
|---:|---:|
| 1/6 | 16 |
| 2/6 | 4 |
| 3/6+ | 0 |

## Operational findings

- **DriveBy's Swagger 2.0 adapter is fully operational on real-world
  inputs.** All 20 harvested specs were parsed and validated
  successfully; no parse errors, no adapter crashes, no
  spec-conversion failures. The same `validate-only` command line that
  drives the 3.x large-scale arm (Section 7.5) drives the 2.0 arm with
  no flag changes.
- **The systemic specification-quality gap observed in 3.x specs is
  reproduced in 2.0 specs.** P002/P003/P008 fail universally; P004 and
  P005 fail on 95% / 85% of specs. The 1/6 mode score for 3.x APIs in
  Section 7.5.3 (84% of APIs scored 1/6 or worse) is matched by an
  80% rate at 1/6 for 2.0 specs.
- **The naming conventions of Swagger 2.0 do not produce systematically
  weaker outcomes than OpenAPI 3.x.** P001 passes universally (each
  spec is structurally valid 2.0), and the 1/6–2/6 score band
  observed in 2.0 mirrors the 3.x band — there is no evidence that
  DriveBy under-detects in 2.0 because of any version-specific
  parsing limitation.

## Mapping to supervisor comment [465]

This evaluation closes [465]:
- DriveBy's source-code adapter for Swagger 2.0
  (`driveby-cli/internal/spec/swagger2.go`, ~150 LOC) handles the
  spec-version split at load time
  (`driveby-cli/internal/loader/loader.go:163`) and routes to the
  same `PrincipleChecker` registry used for 3.x.
- The harvester now accepts a `2.0` selector and successfully
  produces 20 distinct-provider samples.
- The framework runs end-to-end on the produced sample with the same
  per-principle behaviour reported for 3.x.
- The 2.0 arm is reported as Section §7.5.4 of Chapter 7.

## Reproducibility

- `apis-2.0-raw.csv` — input dataset (20 rows).
- `reports/<api>/validation-report-latest.json` — full DriveBy report
  per API.
- `per-principle-summary.csv` — aggregated table used in §7.5.4.
- The harvester and batch script are version-controlled
  (`tools/harvest-openapi-apisguru.py`,
  `tools/run-openapi-batch.sh`).
- The aggregator is `/tmp/agg_swagger2.py` (recreated from the
  thesis run; can be lifted into `tools/` if useful long-term).
