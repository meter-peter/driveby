# Per-Check Defect Injection — Results Index

Experiment described in `thesis/chapters/07-evaluation.tex` § 7.2.4 (`subsec:per-check-extension`).
Date: 2026-05-03. Harness: `tools/per-check-defect-injection.py`. Random seed: 42.

## Headline

| Metric | Value |
|---|---|
| Single-check mutations | 15 |
| Single-check detection rate | 12/15 (80%) |
| Combination tests (pairs + triple) | 4 |
| Combination detection rate | 3/3 evaluated, 1 skipped |

## Misses (real framework limitations)

- `p005-strip-op-security` — P005 treats per-op security as redundant when global is set. Recorded as a candidate strengthening in Ch.10.
- `p009-strip-param-examples` — P009 examines schema-level rather than parameter-level examples. Same.

## Collaterals

- `p002-strip-descriptions` also flags P008 because the versioning-strategy check looks for keywords in `info.description` — principled cross-talk, not a false positive.

## Parse failures

- `p004-strip-constraints` produces a spec the loader rejects (a P001-level catch at the I/O boundary). Demonstrates that structural invalidity short-circuits semantic validation.
