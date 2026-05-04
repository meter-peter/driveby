# Thesis Revision State — 2026-05-04

Snapshot of the thesis after rounds 1 and 2 of Prof. Kritikos's review.
Treat this file as the source-of-truth status when picking the work back up.

## Headline

| metric | value |
|---|---|
| **Build** | clean — 212 pages, 0 `Overfull \hbox`, 0 `Underfull \hbox`, 0 errors, 0 undefined refs |
| **Front matter** | bilingual: English Abstract + Greek Περίληψη + ToC + LoF + LoT + Acronyms (longtable) |
| **Latest commit** | round-3 (Acronyms longtable + figure overlap fixes + missing `amssymb` package) |
| **Round-1 commit** | `b96e6a6` ("Round-1 thesis revision addressing supervisor's 610 inline comments") |
| **Round-2 commit** | `5f342e2` ("Round-2 thesis closure: severity alignment, Greek abstract, comment-index audit") |
| **Round-3 fixes** | Acronyms multi-page table; Fig 5.1 legend frame overlap; Fig 5.4 reconciliation-loop annotations; long-token `\allowbreak` hints; `\usepackage{amssymb}` (was blocking `\bigstar`) |
| **Branch** | `main`, pushed to `origin/main` |
| **Resubmission target** | 2026-05-31 |
| **Defense** | mid-June 2026, in-person, University of the Aegean |

## Comment-index closure

| status | count | meaning |
|---|---:|---|
| `done` | 50 | explicitly cited in `changes-log.md` |
| `addressed-by-rewrite` | 280 | absorbed into round-1/2 chapter rewrites |
| `discussed` | 3 | disposition recorded in `changes-log.md` or `professor-reply.md` |
| `deferred` | 7 | comment text explicitly flags future work / PhD continuation |
| `needs-walkthrough` | 270 | typo/generic/struct in partially-rewritten chapters; will be confirmed live during defense walk-through |
| `open` | **0** | every row classified |
| **total** | **610** | |

Classifier: `review/bulk-classify.py` (re-runnable, version-controlled).

## Work delivered

### Round 1 (2026-05-03, commit `b96e6a6`)

- DDT → SDT rename thesis-wide; framed as a *conceptual framework*, not a complete methodology.
- Ch.8 ownership rewrite addressing supervisor's [528] checklist with quantitative `git log`-derived evidence (350-day solo phase + 12-day partnership window, 74% of source lines authored solo).
- Ch.7 evaluation extensions:
  - Per-check defect injection (15 mutations + 3 random pairs + 1 triple) — `results/per-check-injection/`.
  - APIs.guru large-scale rerun on 50 OpenAPI 3.x specs with raw counts and explicit thresholds — `results/apisguru-rerun/`.
  - Operational PoC restructure (initial situation, configuration, activity volume).
  - SDT-feedback agent experiment proving 1/6 → 4/6 on Petstore — `results/sdt-feedback/`.
- Three new architectural diagrams: `figures/system-context.tex`, `figures/component-diagram.tex`, `figures/class-diagram.tex`.
- Front-matter wired in: English abstract, lists of figures/tables, acronyms.
- Doctoral Continuation roadmap in §10.5.1.

### Round 2 (2026-05-04, commits `5f342e2` + `14ea94c`)

- **Severity alignment** to source-of-truth: P001–P005 are all `Critical` (per `driveby-cli/internal/types/report.go`, escalated in commit `c96d9c7`). Ch.3 metadata, axiom-principle-mapping table, and 7 interpretive paragraphs in Ch.7 updated. The "84% APIs.guru fail strict mode" consequence is reframed as a load-bearing result of Ch.7, not a problem.
- **§7.7 RQ-attribution audit**: removed [525]'s "quality gaps → RQ2" mistag; added an opening paragraph explaining RQ4 evidence intentionally lives in Ch.5/Ch.10.
- **Greek abstract** (`chapters/abstract-gr.tex`): faithful translation of `abstract-en.tex`, all Latin tech tokens wrapped in `\la{...}` (= `\textlatin{...}`) so the LGR encoding does not transliterate them.
- **Source-side rename divergence note** in `apis/CLAUDE.md`, `README.md`, `apis/perfect-api/CLAUDE.md`: thesis says `non-critical-api`, source keeps `perfect-api`, rename scheduled for v4.0.
- **Comment-index bulk classification**: 610/610 rows assigned a status; classifier `review/bulk-classify.py` is the audit trail.

## What is NOT done (defensible-but-supervisor-may-probe)

- **Examination committee names** on the title page — blocked on supervisor input. Apply with a small pre-submission patch when names arrive.
- **270 `needs-walkthrough` items** — typos / generic comments in Ch.2/4/5/6/9 that round-1 didn't blanket-close. Defensible during live defense walk-through; not blocking submission.
- **`perfect-api` source rename** — 91 occurrences across `apis/`, `kubernetes/`, `tools/`, `Makefile`. Deferred to v4.0 release; documented in source-side READMEs.
- **OpenAPI 2.x evaluation arm** — closed as future work in §7.4.

## Picking the work back up

1. Read this file for status.
2. `cd thesis && pdflatex main.tex` (×3 passes) — should produce 212 pages, 0 overfull.
3. If supervisor sends a new round of comments, repeat the round-1/2 pattern: extract → index → classify → rewrite → log.
4. The Greek build requires Fedora packages: `texlive-collection-langgreek` and `texlive-babel-english` (both installed as of 2026-05-04 on this machine).

## Key files

| Path | Purpose |
|---|---|
| `thesis/main.pdf` | latest build (212 pp, bilingual) |
| `thesis/review/comment-index.md` | 610-row audit table |
| `thesis/review/changes-log.md` | per-comment closure log (round 1 + round 2) |
| `thesis/review/REVISION-REPORT.md` | full narrative report including round 2 closure section |
| `thesis/review/professor-reply.md` | Greek reply to supervisor (sent or to-be-sent) |
| `thesis/review/bulk-classify.py` | re-runnable comment classifier |
| `thesis/review/kritikos-comments.md` | original 610-comment extract from annotated PDF |
| `thesis/chapters/abstract-en.tex` | English abstract (~70 lines) |
| `thesis/chapters/abstract-gr.tex` | Greek abstract (~22 lines, Latin tokens wrapped) |
| `results/sdt-feedback/INDEX.md` | agent-feedback experiment artefacts |
| `results/apisguru-rerun/INDEX.md` | 50-API large-scale rerun |
| `results/per-check-injection/INDEX.md` | 15-mutation per-check experiment |
