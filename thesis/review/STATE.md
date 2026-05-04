# Thesis Revision State — 2026-05-04

Snapshot of the thesis after rounds 1 and 2 of Prof. Kritikos's review.
Treat this file as the source-of-truth status when picking the work back up.

## Headline

| metric | value |
|---|---|
| **Build** | clean — 213 pages, 0 `Overfull \hbox`, 0 `Underfull \hbox`, 0 errors, 0 undefined refs |
| **Front matter** | bilingual: English Abstract + Greek Περίληψη + ToC + LoF + LoT + Acronyms (longtable) |
| **Latest commit** | round-5 (must-priority closures: Fig 3.1 removed [228], Ch.7 evaluation arms reordered [455]) |
| **Round-1 commit** | `b96e6a6` ("Round-1 thesis revision addressing supervisor's 610 inline comments") |
| **Round-2 commit** | `5f342e2` ("Round-2 thesis closure: severity alignment, Greek abstract, comment-index audit") |
| **Round-3 commit** | `56efe41` ("Round-3 thesis layout closure: zero overfull, zero errors, 212 pp") |
| **Round-4 commit** | `f850ed6` ("Round-4 thesis closure: full needs-walkthrough sweep") |
| **Round-3 fixes** | Acronyms multi-page table; Fig 5.1 legend frame overlap; Fig 5.4 reconciliation-loop annotations; long-token `\allowbreak` hints; `\usepackage{amssymb}` (was blocking `\bigstar`) |
| **Round-4 fixes** | Ch.2 prose adds (EvoMaster, OpenAPI 3.2, Smardas+Kritikos 2025, Optic/APIClarity/42Crunch named, code-first rewording, Table 2.1 symbol legend); Ch.4 vegeta "attack" disclaimed; comment-index 270 needs-walkthrough rows hand-classified per-row |
| **Round-5 fixes** | [228]: Fig 3.1 (axiom-principle-mapping figure) removed — Table 3.x is the single source for the mapping. [455]: Ch.7 evaluation arms physically reordered to PoC → controlled → multi-API → large-scale; chapter intro, RQ-attribution paragraph, and Fig 7.1 caption rewritten to match the new progression |
| **Branch** | `main`, pushed to `origin/main` |
| **Resubmission target** | 2026-05-31 |
| **Defense** | mid-June 2026, in-person, University of the Aegean |

## Comment-index closure (post round-4 hand audit)

| status | count | meaning |
|---|---:|---|
| `done` | 64 | explicitly cited in `changes-log.md` (round-1 closure), hand-edited in round 4, or directly addressed in round 5 |
| `addressed-by-rewrite` | 516 | absorbed into round-1/2/4 chapter rewrites; per-row resolution note in `[r4: ...]` trailer |
| `discussed` | 23 | design-debate comment whose disposition is recorded per row and in `changes-log.md` / `professor-reply.md` |
| `deferred` | 7 | comment text explicitly flags future work / PhD continuation |
| `needs-walkthrough` | **0** | round-4 sweep classified every row case-by-case |
| `open` | **0** | every row classified |
| **total** | **610** | |

Round-1 classifier: `review/bulk-classify.py` (re-runnable, version-controlled).
Round-4 hand audit: 270 previously-`needs-walkthrough` rows were walked one-by-one against the current chapter prose; each carries a `[r4: ...]` trailer naming the section/subsection that addresses the supervisor's concern. Any individual comment can therefore be answered live at defense without searching.

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
- **`perfect-api` source rename** — 91 occurrences across `apis/`, `kubernetes/`, `tools/`, `Makefile`. Deferred to v4.0 release; documented in source-side READMEs.
- **OpenAPI 2.x evaluation arm** — closed as future work in §7.4.

## Picking the work back up

1. Read this file for status.
2. `cd thesis && pdflatex main.tex` (×3 passes) — should produce 213 pages, 0 overfull.
3. If supervisor sends a new round of comments, repeat the round-1/2 pattern: extract → index → classify → rewrite → log.
4. The Greek build requires Fedora packages: `texlive-collection-langgreek` and `texlive-babel-english` (both installed as of 2026-05-04 on this machine).

## Key files

| Path | Purpose |
|---|---|
| `thesis/main.pdf` | latest build (213 pp, bilingual) |
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
