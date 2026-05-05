# Thesis Revision State — 2026-05-04

Snapshot of the thesis after rounds 1 and 2 of Prof. Kritikos's review.
Treat this file as the source-of-truth status when picking the work back up.

## Headline

| metric | value |
|---|---|
| **Build** | clean — 215 pages, 0 `Overfull \hbox`, 0 `Underfull \hbox`, 0 errors, 0 undefined refs |
| **Front matter** | bilingual: English Abstract + Greek Περίληψη + ToC + LoF + LoT + Acronyms (longtable) |
| **Latest commit** | round-7 (defensive polish: typo grep-pass, methodology→framework softening, defense Q&A doc) |
| **Latest revision pass** | round-9 (consistency polish: 4 surgical fixes addressing inter-chapter drift surfaced by a parallel-agent audit; uncommitted) |
| **Round-8 status** | uncommitted; trailer-sharpening for must-priority Ch.3/Ch.7 |
| **Round-1 commit** | `b96e6a6` ("Round-1 thesis revision addressing supervisor's 610 inline comments") |
| **Round-2 commit** | `5f342e2` ("Round-2 thesis closure: severity alignment, Greek abstract, comment-index audit") |
| **Round-3 commit** | `56efe41` ("Round-3 thesis layout closure: zero overfull, zero errors, 212 pp") |
| **Round-4 commit** | `f850ed6` ("Round-4 thesis closure: full needs-walkthrough sweep") |
| **Round-3 fixes** | Acronyms multi-page table; Fig 5.1 legend frame overlap; Fig 5.4 reconciliation-loop annotations; long-token `\allowbreak` hints; `\usepackage{amssymb}` (was blocking `\bigstar`) |
| **Round-4 fixes** | Ch.2 prose adds (EvoMaster, OpenAPI 3.2, Smardas+Kritikos 2025, Optic/APIClarity/42Crunch named, code-first rewording, Table 2.1 symbol legend); Ch.4 vegeta "attack" disclaimed; comment-index 270 needs-walkthrough rows hand-classified per-row |
| **Round-5 fixes** | [228]: Fig 3.1 (axiom-principle-mapping figure) removed — Table 3.x is the single source for the mapping. [455]: Ch.7 evaluation arms physically reordered to PoC → controlled → multi-API → large-scale; chapter intro, RQ-attribution paragraph, and Fig 7.1 caption rewritten to match the new progression |
| **Round-6 fixes** | [465]: Swagger 2.0 evaluation arm added as Ch.7 §7.5.4 (`subsec:swagger2-eval`). Harvester extended with `2.0` selector + per-provider cap; 20 distinct-provider Swagger 2.0 specs harvested from APIs.guru and validated in strict mode; results table reports per-principle pass rates; the systemic specification-quality gap observed in 3.x is reproduced in 2.0 (80% of specs score 1/6 vs 84% for 3.x). Ch.9 §9.5 external-validity paragraph updated to cite the 70-API combined population. All three `must`-priority `discussed` items now closed |
| **Round-7 fixes** | Defensive polish before submission. (a) Typo grep-pass: 10 marginal-typo rows verified against current text and reclassified `done` (PVC expanded on first use; remaining 9 confirmed already absorbed). (b) [557] follow-through: "methodology"→"framework"/"approach" in 5 prominent places (Ch.1 ×3, Ch.7 ×1, Ch.8 ×1) where the formal-methodology connotation overclaimed. Disclaimer phrasing in Ch.1 §1.4 and Ch.10 §10.5 left intact (those uses are deliberately the framework-vs-methodology contrast). (c) New `review/defense-qa.md`: rehearsed answers for the 10 highest-risk discussed/scope items (Ch.9-merge, methodology, control-theory analogy, composite testing, risk scoring, Ch.4 title, static-vs-syntactic, reconciliation-vs-adaptation, P006/P007 wrapper status, single-cluster PoC) |
| **Round-8 fixes** | Trailer-sharpening pass triggered by an audit that found ~70% of `addressed-by-rewrite` trailers were generic boilerplate (`[r2: ... in heavily-rewritten ch3-methodology]` / `[r2: evaluation extensions delivered in round 1 ...]`) that could not be defended live. Sharpened **all 75 must-priority `addressed-by-rewrite` trailers in Ch.3 and Ch.7** to cite specific section labels and line numbers, naming the table or paragraph that resolves each concern. Mis-trailed Ch.8 cluster [520]–[527] re-pointed to the actual answer locations (Ch.3 §3.1 for "what does specification mean", Ch.7 §7.5/§7.7 for defect-classification and RQ-correlation, Ch.10 §10.3 for RQ4-ontological-spec). [481] (4/9 critical-principles) reclassified `addressed-by-rewrite` → `discussed` since post-revision severity escalation deliberately moves the count to 5/9 — knowingly-deviant point. Zero `[r2: ...]` boilerplate remains on must-priority rows for Ch.3 and Ch.7 |
| **Round-9 fixes** | Consistency polish triggered by a parallel-agent audit across all 10 chapters. Four surgical edits resolve inter-chapter drift surfaced by the audit: (a) Ch.3 §3.5 round-2 severity footnote at line 242 fixed — was "P001--P006 critical", corrected to "P001--P005 critical (P006 implementation pending)" to match the P006 status table at line 171. (b) Ch.7 §7.6 line 486 reworded — earlier wording read as "5+2=7 critical" on a careless read; new wording explicitly separates the 5 static-Critical principles from the 2 runtime-Critical-by-impact principles, matching the 5/9 ratio used elsewhere. (c) Ch.8 §8.1 line 97 — precise/rounded reconciliation now visible: "approximately 74%---73.7% added and 73.5% removed---authored solo", closing the [547] precise-vs-estimate probe target. (d) Ch.5 §5.7 added a footnote on first `perfect-api` mention (line 283) explicitly disclaiming the source-rename divergence with a forward-pointer to Ch.7 §7.3; removes the supervisor's confusion when reading Ch.5 before Ch.7. Build verified clean: 215 pp, 0 overfull, 0 errors, 0 undefined refs. |
| **Branch** | `main`, pushed to `origin/main` (rounds 8 + 9 uncommitted on main) |
| **Resubmission target** | 2026-05-31 |
| **Defense** | mid-June 2026, in-person, University of the Aegean |

## Comment-index closure (post round-4 hand audit)

| status | count | meaning |
|---|---:|---|
| `done` | 75 | explicitly cited in `changes-log.md` (round-1 closure), hand-edited in round 4, addressed in round 5, addressed in round 6 (Swagger 2.0 arm), or verified in round 7 grep-pass (10 typo rows) |
| `addressed-by-rewrite` | 505 | absorbed into round-1/2/4 chapter rewrites; per-row resolution note in `[r4: ...]` or `[r7: ...]` trailer (round-8 sharpened all 75 must-priority Ch.3/Ch.7 trailers to cite specific §X.Y or line ranges, plus the 8-row Ch.8 mis-trail cluster [520]–[527] re-pointed to their real answer locations) |
| `discussed` | 23 | design-debate comment whose disposition is recorded per row and in `changes-log.md` / `professor-reply.md` / `defense-qa.md` (round-8 reclassified [481] — the 4/9-vs-5/9 critical-principle count is a deliberate post-revision deviation, not a rewrite-absorption) |
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

## Picking the work back up

1. Read this file for status.
2. `cd thesis && pdflatex main.tex` (×3 passes) — should produce 215 pages, 0 overfull.
3. If supervisor sends a new round of comments, repeat the round-1/2 pattern: extract → index → classify → rewrite → log.
4. The Greek build requires Fedora packages: `texlive-collection-langgreek` and `texlive-babel-english` (both installed as of 2026-05-04 on this machine).

## Key files

| Path | Purpose |
|---|---|
| `thesis/main.pdf` | latest build (215 pp, bilingual) |
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
