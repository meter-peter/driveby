# Changelog for the reply to Kritikos (running log — NOT the email itself)

Tracks every change made in response to the 2026-06-09 feedback round (480 inline PDF
comments + 15-point email summary), with the author's decisions. Source of truth for
drafting the reply email when the new version is sent.

Format: headline point → decision taken → what changed in the document/code/data.

---

## Headline #5 — Stale reference numbers & review-round narration — DONE

**Decision:** Adopted Kritikos's preferred alternative: the text is now fully
**self-contained** — it reads as a clean first version with zero references to
previous versions, review rounds, or comment numbers (rather than keeping a one-time
summary of comment handling).

**What changed:**
- Mechanical sweep (77 applied edits across all 10 chapters + front matter): deleted
  all stale bracketed numbers (`~[161]`, `[466], [476]`, `[494]`, …) and round-1/round-6
  provenance narration; sentences smoothed to read as v1 prose.
- Second pass (~35 additional entangled sites in Ch. 1, 3, 7, 8, 9 + acronym table):
  removed "per the supervisor's review", "earlier draft", "in this revised text",
  "post-revision severity assignment", "round-2 revision", "Closing comment [465]"
  paragraph heading, supervisor-question framing throughout Ch. 8.
- Sections formerly motivated by "the supervisor asked for X" are now motivated from
  the domain (e.g. Swagger 2.0 arm: motivated by 2.0's continued production share).
- DDT→SDT rename footnote kept (practical necessity: tool/repo keep the old name) but
  no longer attributes the rename to supervisor review.
- Verified by grep: zero non-`\cite` bracketed numbers, zero review/round/draft
  narration anywhere in chapters/. Acknowledgements (thanking the supervisor) intact.
- Also stripped Abstract/List of Figures/List of Tables/Acronyms from the ToC
  (front-matter comment).

## Headline #13 — Multi-API tables contradicted the one-defect claim — DONE

**Decision:** Reword + regenerate. The claim is now "each API targets exactly one
*defect class*" (documentation-class strips legitimately cascade across P002–P004 +
P008); AND the underlying data was regenerated so it actually shows clean per-layer
discrimination.

**Root cause found:** the defect APIs' specs had already been fixed in the repo
(slow/broken specs are md5-identical to the reference; no-auth = reference minus
security only), but `results/local/*.json` dated 2026-03-30 predated that fix — the
old reports captured specs inherited from a dirty shared base.

**What changed (data):**
- Re-ran the full local arm live (docker compose, DriveBy CLI built from source):
  strict validation ×5, functional ×2 (with auth), load ×2 (10 users / 15 s /
  P95 ≤ 200 ms), test-ready ×1. All 10 reports re-saved (pretty-printed) in
  `results/local/`.
- New measured matrix: non-critical 6/6; bad-docs fails P002/P003/P004 (+P008 W),
  **passes P005** (fixes the wrong table cell flagged in the review); no-auth fails
  **only P005**; slow & broken **pass all six static principles** and fail only at the
  runtime layer (P007: P95 505 ms > 200 ms; P006: 418 vs documented 200, 200 vs 201).

**What changed (text):**
- Tables updated to measured data: `tab:defect-ground-truth` (incl. severity header
  W→C for P002–P004), `tab:local-strict` (passed counts 5/6, 6/6, critical-failure
  counts), `tab:api-suite` (bad-docs expected failure "P002, P003, P004 (C); P008 (W)"),
  `tab:staging-gates` (bad-docs blocking cell corrected from "P005 Security" to
  "P002/P003/P004 Documentation"), `tab:prod-gates` (slow-api 505 ms).
- Figure `per-principle-pass-rates.tex` regenerated (was 0% for P002–P004; now 80%
  for every principle except P001 at 100%) + caption.
- Prose rewritten at: suite-design observations, local-results paragraph, staging
  narrative, production narrative (load numbers: P50 4.2/P95 6.8/P99 9.1 ms),
  severity-model bullet, large-scale-comparison paragraph, RQ1 summary, Ch. 9
  RQ1 discussion paragraph (also fixed factual error "bad-docs … lacking security
  definitions").
- New framing: static gate blocks exactly the specification-defect APIs; runtime layer
  catches exactly the behavioural ones — layered defence with clean discrimination.

**Caveat to mention in the email (or fix before sending):** the staging/production gate
tables describe the cluster PoC; the PASS/FAIL verdicts are identical under the
corrected specs, but the cluster gates were not re-triggered. Optional: re-run on
private.novelcore.org for full rigor.

## Decisions locked (work queued, not yet done)

- **#10 chapter split — DONE.** New Chapter 4 "System Architecture"
  (`chapters/04-system-architecture.tex`, `ch:system-architecture`) before the CLI
  chapter. Contains: (a) explicit system-vs-substrate definition — the system is
  DriveBy CLI + XSDLC composition + generated gate infrastructure + **GitOps Promoter
  (moved inside the boundary, per comment #170)**; Kubernetes + Crossplane are the
  infrastructure substrate, not system components (per #169); (b) the system-context
  figure **redrawn** with expanded boundary, substrate band, and external actors
  reduced to developer, operator, GitHub repositories, target API/spec, report;
  (c) the three-layer model + progression table (migrated from the K8s chapter);
  (d) the five-stage CLI→CR evolution (migrated from the K8s chapter); (e) roadmap.
  The CLI chapter was retitled "The DriveBy CLI: Internal Architecture", its scope
  narrowed to the binary, with a signpost referencing the system chapter (per #173).
  K8s chapter intro/summary updated; Ch.1 thesis-structure roadmap gained the new
  chapter entry; Ch.6 references the three-layer model. Chapter count 10 → 11.
- **#14 P007 severity — DONE.** P007 flipped to Critical **in the source of truth
  first** (`driveby-cli/internal/types/report.go`; build + vet + tests green), then
  the load-test reports were regenerated with the new binary so the saved JSON in
  `results/local/` carries `severity: critical`, then the thesis: P007 subsection
  severity line, Table 3.x principle catalogue (Warning→Critical), Table 7.9 slow-api
  cell (W→C), and the Ch. 7 severity-model paragraph now states "seven of nine carry
  Critical" with the new sentence: a declared performance threshold is a contract —
  a P007 breach always blocks promotion, no advisory mode for a violated SLA.
  `docs/principles/P007.md` updated to match. Also fixed an incorrect naming-note
  parenthetical (said "P001 in bad-docs-api"; now states the defect classes).
- **#3 single-artifact — CONCEDED & DONE.** Kritikos is right on both counts: (a) most
  runtime testers consume only the OpenAPI spec, so "single-artifact" as worded did
  not discriminate; (b) SDT's own performance thresholds live in the XSDLC CR, a
  second artifact. Changes: Table 2.1 criterion renamed **"Zero Test Authorship"**
  (authorship, not input cardinality, is the discriminator) and redefined in the
  taxonomy prose; cells corrected honestly — Schemathesis/RESTler/Dredd now ✓ (they
  derive tests from the spec automatically), Pact stays ✗ (the canonical counter-
  example: hand-authored consumer contracts), Microcks Static ~→✗ (per comment #89);
  SDT's cell carries a †-footnote stating that test *content* is spec-derived while
  P007 thresholds are declared in the XSDLC CR (deployment policy). Post-table prose
  and research Gap 5 rewritten: the gap is no longer "single-artifact" but "no
  specification-derived FULL quality verdict (spec quality + functional + performance)
  with zero test authorship". Ch. 3 gained a scoping note: spec-sole derivation holds
  for test content, not operational policy. Also per comments #78/#82: the
  ontological-reconciliation section now states the governed resource is dual (static
  principles govern the document; runtime principles govern the implementation) and
  remediation is correspondingly dual (enrich the spec vs fix the source code); per
  #81, XSDLC added to the declared-ontology family (Terraform/Crossplane/OPA/ArgoCD).
- **#9 P009/test-ready — DONE (kept design, strengthened justification).** No CLI
  redesign. New "Why a principle and not a mode" paragraph in the P009 subsection:
  (1) modes select principles, principles carry verdicts — test-ready = {P001–P004
  reduced, P009}, and without P009 the mode's defining question would have no
  reportable verdict; (2) the disjunctive-precondition argument — P009 asks "example
  OR typed schema" (the weakest condition for typed test data), strictly weaker than
  P002's conjunctive documentation checks; a spec with typed schemas but no examples
  fails P002 yet passes P009, exactly when functional testing is still meaningful —
  so P009 is NOT expressible as "P001 + partial P002–P004"; (3) composability.
  This supersedes the "building block" email reply with a sharper version of it.
- **#7 axiom/principle mapping — DONE (hybrid per Kritikos's own proposal in #146).**
  Kritikos's inline comments endorse a Governing-Axiom + Impacted-Axioms model — which
  reconciles with the author's "all axioms govern every principle" position rather
  than contradicting it. Implemented exactly his table spec (#143): the principle
  catalogue now has Governing Axiom and Impacted Axioms columns (P001–P004
  Completeness; P005/P007/P008 Observability; P006/P009 Determinism; impacted = the
  other two, always). Prose reworked: governing groups rather than partitions; group-
  level reasoning added (completeness / observability / determinism groups). Chapter
  intro's "no partition" sentence reconciled to the two-sided claim.
- **#8 mode-specific notes — DONE (all nine principles).** Every principle now carries
  a **Mode behaviour** note immediately after its Severity/Status line (where Kritikos
  asked for it, comment #121): P001 mode-uniform; P002/P003/P004 mode-specific with
  the exact test-ready check subsets enumerated (#128); P005 and P008 explicitly "not
  evaluated in test-ready/minimal" with one-line rationale (#130, #137); P006/P007
  runtime principles (test-only mode / gate check types); P009 mode-uniform.
  Also fixed a real prose-vs-table inconsistency the audit surfaced: the P003/P005/
  P008 bodies described "minimal mode" behaviour that cannot occur (the registry's
  ForMode(minimal) returns only P001) — that dead text is removed and the prose now
  matches the registry source of truth.
- **#1 agent-reconciliation evaluation threading, #2 RQ-contribution map, #4 research
  gaps ↔ DriveBy in §2.6.3, #11 figure pass, #15 report-authorship evidence** → in the
  bulk apply pass (run `wf_fac7a8c7`), see below.

## Anchored catalog + author interview (2026-06-10)

All 480 comments were re-anchored rect-precisely (identity-verified against the
canonical comment ids after an ordering bug was caught and fixed — 110 ids had been
shuffled between extractions) and explained against the live source:
`kritikos-anchored-catalog.md` / `kritikos-anchored-explained.json`.
Verdicts: 287 open, 163 already-addressed, 19 needed author decisions, 11 unclear.

**Author decisions (interview, 2026-06-10):**
1. **Modes (#460/#473):** strict stays STATIC-only permanently; P006/P007 wrappers
   target test-only / a future composed "full validation" (strict + test-only),
   never strict. The "single strict invocation = complete quality assessment"
   sentence is removed. (Adopts Kritikos's separation.)
2. **NFR thresholds in spec (#434/#29):** acknowledged as future work — an OpenAPI
   vendor extension (x-performance) carrying thresholds would make the approach
   single-artifact in the strict sense; cross-referenced in Ch.2 (Table 2.1 †),
   Ch.9, Ch.10; RQ2 wording qualified ("specification together with declared
   operational policy").
3. **P006 scope (#140/#203/#371):** explicit contract-conformance scope statement
   (status codes, schema shape, types); semantic content correctness is a recorded
   limitation; content assertions derived from response examples = future work;
   P009 check-4 wording reconciled. Source-aware agent revision noted as future work.
4. **Control loop (#273/#276):** remap as instructed — Table 6.1 Plant = the
   delivery/promotion system, Setpoint = promotion-policy predicate; Figure 6.5
   updated in lockstep; two-tier material kept as supporting decomposition.
5. **Ch.2 restructure (#63/#65):** full — OpenAPI Ecosystem becomes §2.1; contract
   testing split into provider-driven (Dredd folded in) / consumer-driven (Pact).
6. **CI/CD metrics (#291):** moved into the operational-PoC section as its final
   subsection.
7. **Cluster figure naming (#217):** namespace labels genericized to <app>-dev/
   -staging/-prod; caption notes the PoC instance used the historical name.
8. **Committee names (#1):** not yet announced — title pages keep supervisor only.
9. **RQ4×C5 (#40):** tilde (partial empirical validation via the PoC).
10. **Agent remediation scope (#470):** spec-first (stable ground truth, per
    Kritikos); extension to code patches is conditional on a reliable procedure/
    workflow for deciding which artifact is authoritative.
11. **AI design ideas (#405):** demarcation sharpened, not conceded — strict
    architect-executor division; agent suggestions were implementation-level
    tactics under author direction, never design input.
12. **Validity (#444):** agent-confirmation-bias folded into Internal Validity.
13. **Per-check completeness (#313):** limitation note added (full check-coverage
    mutation with randomised realisations = future evaluation work).
14. **Table 3.2 footnotes (#153):** kept (Kritikos accepted the repetition).

Also fixed during the audit: comment #256 (missing articles in the CLI→CR evolution
enumeration) — articles restored in the Ch.5 intro.

**Page-count note for the email:** the document grew (229→235 before the bulk pass)
because the removals were phrase-level (review narration) while the additions were
page-level demanded content (new System Architecture chapter, mode notes, P009
justification, axiom columns). The repetition-removal items land in the bulk pass.

## Mechanical sweep (the 96 pre-vetted small edits) — DONE

77 applied, 16 were already resolved in the current source, 3 skipped as ambiguous
(OCR-fragment comments at 07:722, 03:120 "consumed vs utilised", 08: grammatical
fragment). Includes: terminology unification ("mode-aware" → "mode-specific";
"critical sweep" removed), missing Kong citation added to references.bib, wrong
principle inventory fixed in `figures/dependency-flow.tex` (P001–P008 → P001–P005,
P008, P009).

## Bulk apply pass (run `wf_fac7a8c7`, 2026-06-10) — DONE

312 catalog items applied across all chapters (49 skips audited — all benign or
re-routed). Highlights beyond the decisions above:
- **RQ wording synced verbatim** from Ch.1 into Ch.7's RQ list; evidence pointers
  corrected (7.5 added to RQ1; RQ2's two halves mapped explicitly; conclusion
  pointer dropped from RQ4). Four-arm evaluation framing threaded through
  Ch.1/7/8/9/10 and both abstracts (was "three arms" + an unexplained fourth).
- **Table 7.1 reconciled against the actual XSDLC CRs** in kubernetes/examples/:
  all five apps carry production gates with autoMerge:false; broken-api row fixed;
  slow-api load parameters corrected; false "production gate uses strict" prose
  removed. 44-vs-32 Workflow/CommitStatus arithmetic explained. Resource count
  unified (~44, up to 47 by counting convention) across Ch.5/Ch.7.
- **Evaluation-methodology figure** rebuilt: PoC arm added, unbacked "Effort
  Comparison" arm removed, agent-feedback arm added.
- **Ch.2 restructured** per the proposal (OpenAPI Ecosystem → 2.1; provider- vs
  consumer-driven contract testing; Dredd folded in). CI/CD metrics moved into the
  PoC section. Control loop remapped (Table 6.1 + Figure 6.5 in lockstep: Plant =
  delivery/promotion system). BYOCI scope sentence added to Ch.1 (#305).
- **Six orphaned figure comments fixed** (#222 provider arrows, #234/#235
  xrd-hierarchy arrows + per-environment group, #241 declarative-lifecycle arrow
  endpoints/labels, #264 clipped SDT-gate arrow, #268 thicker arrows in
  human-in-the-loop). Cluster figure namespaces genericized (#217).
- **Honesty fixes:** zero-human claim qualified (autoMerge:false requires a human
  merge); "any drift" qualified by P006 coverage; health-check edge case reported
  as not-reproducible rather than failed; p005 miss reclassified as evaluation
  artefact vs p009 as genuine limitation; per-check combination protocol now
  reports 3/3 evaluable + 1 pair unevaluable due to an upstream kin-openapi panic.

## Condensation pass (run `wf_9eead6b3`, 2026-06-10) — DONE

Kritikos warned the text risks growing; the apply pass added ~11,300 words. A trim
pass removed 3,212 words of redundancy from today's additions (every number,
decision, and demanded justification kept). Net: 235 → 258 → **252 pages**.
For the email: the residual growth over the reviewed version is the new System
Architecture chapter plus the explicitly requested justifications; the review-round
narration that was removed was phrase-level, not page-level.

## Deep-trim pass (run `wf_8f91e497`, 2026-06-10) — DONE

Per the author's instruction ("he asked for deletions and expects less"), a
whole-document deletion pass removed **12,553 words**: cross-chapter
re-explanations replaced by one sentence + \ref to each concept's canonical home,
chapter-end summaries collapsed to ≤3 sentences, table-reciting prose deleted,
roadmap paragraphs cut to one sentence, oversized listings elided. Everything
answering a supervisor comment, every number, table, figure, and label preserved
(orphaned labels relocated, never dropped). Plus: redundant intro footnote removed
(healed a broken page), ToC limited to section depth (−3 pages of front matter).

**Final size: 217 pages — one page FEWER than the 218-page version reviewed,**
despite containing the new System Architecture chapter (+5) he ordered and all the
demanded justifications. Per-chapter vs reviewed: Intro −1, Related −1, Methodology
+1, SysArch +5 (new), CLI 0, Kubernetes −5, GitOps 0, Evaluation −1, AI 0,
Discussion +1, Conclusion +1.

**Email line:** «Το κείμενο είναι πλέον κατά μία σελίδα μικρότερο από την έκδοση που
σχολιάσατε, παρότι ενσωματώνει το νέο κεφάλαιο Αρχιτεκτονικής Συστήματος που
ζητήσατε και όλες τις ζητούμενες τεκμηριώσεις — η συρρίκνωση προήλθε από την
αφαίρεση των επαναλήψεων που επισημάνατε.»

## Build status

Compiles clean (pdflatex + biber): zero errors, zero undefined references or
citations, zero multiply-defined labels. **217 pages.** All work uncommitted, on `main`.
