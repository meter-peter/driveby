The data is consistent with the live source (75 bib entries confirms the [NNN] numbers are stale; chapters match the 11 files). I have everything I need to write the deliverable.

# Kritikos Review — Action Plan

This plan consolidates **480 inline PDF comments** (triaged per-comment across 11 files) and **15 headline points** (independently re-verdicted) into a single working document. Every item was triaged against the **live `.tex` source**, not the annotated PDF — and the PDF was built from an older revision, so a number of items are already partly or fully resolved in the current source. Where an item is marked resolved, the action is *confirm-and-skip* (and, where useful, tell Kritikos exactly which file/line already satisfies him). The single biggest correctness risk is the multi-API table-vs-text contradiction (headline #13, Ch. 7); the two points where your email pushed back and should reconsider are #3 (single-artifact) and #9 (P009 redundancy).

---

## At a Glance

| File | Comments |
|------|---------:|
| 07-evaluation.tex | 97 |
| 03-methodology-ddt.tex | 61 |
| 02-related-work.tex | 53 |
| 04-cli-architecture.tex | 49 |
| 05-kubernetes-architecture.tex | 45 |
| 09-discussion.tex | 35 |
| 01-introduction.tex | 33 |
| 08-ai-assisted-development.tex | 33 |
| 10-conclusion.tex | 32 |
| 06-gitops-pipeline.tex | 27 |
| front-matter | 15 |
| **Total** | **480** |

| Severity | Count |
|----------|------:|
| Blocking | 79 |
| Major | 224 |
| Minor | 177 |

| Class | Count |
|-------|------:|
| rewrite | 229 |
| trivial | 81 |
| pushback | 52 |
| table-fix | 39 |
| structural | 29 |
| figure | 28 |
| reference-bug | 16 |
| cross-cutting | 6 |

**Mechanical (safely automatable) edits: 96.** These can be applied in one sweep without judgement (see the Mechanical Sweep section).

---

## The 15 Headline Points — Second Opinion

Legend: ⚠ = your email position is wrong/risky, or the verdict is *author-should-concede* with **high** risk. Address these first.

---

### ⚠ Point 1 — Agent-reconciliation evaluation is forgotten in several places
- **Kritikos asks:** Surface the new agent-feedback experiment (AI agent revises a spec from the DriveBy report alone) everywhere it belongs.
- **Your email said:** Framed it as fuzzy "general small comments," deferred to the inline PDF.
- **Source now:** The experiment IS live (abstract, Ch. 7 §`sec:sdt-feedback-experiment`, Ch. 6 control-loop, Ch. 8/9/10) — but genuinely missing from: Introduction contributions + RQ-map; related-work line 169 (the "correction role" claim never forward-refs the experiment that proves it); Ch. 7 figure caption still says "three-arm evaluation."
- **Verdict: author-should-concede.** The complaint is specific and true, not vague.
- **Do this:** (1) Add the experiment to the Intro contributions/RQ-map as RQ2/C5–C6 evidence. (2) Forward-ref from related-work line 169 to `sec:sdt-feedback-experiment`. (3) Fix the Ch. 7 methodology figure caption + "three-arm" wording to show it as a distinct arm.
- **Risk: high.**

### Point 2 — RQ-to-contribution correspondence
- **Kritikos asks:** Enrich/correct the contributions↔RQ mapping.
- **Your email said:** "If needed, it will be done."
- **Source now:** `sec:rq-contribution-mapping` + Table `tab:rq-contribution-mapping` already exist with ✓/~ marks, C1–C6 legend, per-row chapters. One real bug: RQ4 row marks only C3/C4, omitting C5 even though C5 = the IDP PoC that answers RQ4 in Ch. 9.
- **Verdict: already-resolved-in-source.**
- **Do this:** Reply "already done" (cite Sec. 1.5.1 / Table 2), plus add a ~ (supporting) **C5** mark to the RQ4 row.
- **Risk: low.**

### ⚠ Point 3 — "Single-artifact" criterion (Table 2.1) — YOU PUSHED BACK
- **Kritikos asks:** Fix the criterion — SDT relies on TWO artifacts (OpenAPI spec + XSDLC CR); non-functional thresholds live in the CR, so an unqualified "single-artifact ✓" is wrong.
- **Your email said:** Disagreed — the *process* relies on one artifact (the live spec); XSDLC is deployment, not methodology.
- **Source now:** The thesis already contradicts you: Ch. 6 line 288 says the load-test **thresholds** are declared in the XSDLC manifest (a second declarative artifact); Ch. 3 line 15 also says "SDT operates on TWO artefacts" (there: spec + implementation). So "two artefacts" is used with two different meanings, and Table 2.1 still says single-artifact — three statements that cannot all be defended in a viva.
- **Verdict: author-should-concede.** P007 (which #14 makes critical) cannot be derived from the spec alone; its thresholds come from the CR. The "process vs deployment" distinction does not save the unqualified ✓.
- **Do this:** Rename the criterion from "Single-Artifact" to **"Spec-Sole Authoring" / "Zero hand-authored test cases"**, add a footnote that non-functional thresholds are declared in the XSDLC CR, and reconcile the two meanings of "two artefacts" (Ch. 3 vs Ch. 6) into one consistent story.
- **Risk: high.**

### Point 4 — Research gaps not correlated with the DriveBy tool
- **Kritikos asks:** Tie each of the seven gaps (§2.6.3) to the concrete DriveBy/XSDLC feature that closes it.
- **Your email said:** "Will do an iteration on this" (concession).
- **Source now:** §2.6.3 phrases gaps in terms of the abstract paradigm; only Gap 6 names principles (P001–P004); the closing bridge maps gaps → *chapters*, not tool capabilities. No gap-to-DriveBy table exists.
- **Verdict: author-should-concede.** Mechanical, low-risk.
- **Do this:** Add a Gap→DriveBy table (Gap1→unified engine P006+P001–P005; Gap2→P002–P004; Gap3→axioms/registry; Gap4→XSDLC gate+validate-only; Gap5→spec-only validation modes; Gap6→P001–P004; Gap7→XSDLC CR).
- **Risk: low.**

### Point 5 — Stale [NNN] reference numbers + round-N narration
- **Kritikos asks:** Delete the fake bracketed comment-IDs and stop narrating how each comment was addressed; make the text self-contained.
- **Your email said:** Agreed — "I'll do a pass."
- **Source now:** Fully live. `main.tex` uses numeric biblatex (real cites = [1]–[75]); 40 raw [160]–[551] brackets remain (Ch. 3:6, Ch. 7:20, Ch. 8:8, Ch. 10:6) — all read as broken citations. Plus "round-1 review" ×15, "supervisor('s)" ×47, "earlier draft" ×11.
- **Verdict: author-should-concede.** Pure hygiene, no counter-argument.
- **Do this:** Delete all 40 brackets, collapse the meta-narration to self-contained prose, and add a grep/CI guard so they cannot return. Commit to *zero* references to prior review rounds.
- **Risk: low.**

### Point 6 — Fold accuracy/precision into the Completeness *definition*
- **Kritikos asks:** Since you subsume accuracy under Completeness, say so in the definition and map it to principles.
- **Your email said:** "I'll keep it in my head when I get there."
- **Source now:** ~80% done — line 28 justifies the fold; lines 39/41/43 route accuracy to P004 + P006. But the one-sentence formal definition (line 35, "must fully describe the API contract") never names precision/accuracy.
- **Verdict: author-partially-right** (presentation gap, not conceptual error).
- **Do this:** Add one clause to the line-35 definition naming precision/accuracy, and one sentence stating P004 = schema-level accuracy, P006 = spec-vs-runtime accuracy. Tell Kritikos it's near-resolved (cite line 28).
- **Risk: low.**

### Point 7 — Axiom-to-principle relationship (governing vs effective)
- **Kritikos asks:** Either all axioms govern all principles, or add an explicit Governing-vs-Effective distinction in the Table 3.1 mapping, with justification.
- **Your email said:** All axioms govern every principle; the per-principle mapping was already dropped.
- **Source now:** Line 17 + line 245 state exactly this ("does not partition the principles among the axioms"); determinism is justified by the side-effect-free parallel-application formalism (lines 88–105). The old one-to-one mapping figure is gone.
- **Verdict: already-resolved-in-source.**
- **Do this:** Reply citing lines 17/245 + the fp-formalism, **and** add one sentence naming Kritikos's governing/effective proposal and why you rejected it (an "effective subset" label re-introduces the partition the framework removed).
- **Risk: low.**

### ⚠ Point 8 — Mode-specific behaviour not fully specified per principle
- **Kritikos asks:** State each principle's mode behaviour completely; some modes covered for some principles, not all.
- **Your email said:** Soft defer — "I'll see what can be done."
- **Source now:** Line 111 *promises* every principle carries a bold "Mode-specific behaviour" or "Mode-uniform" note — but grep finds **zero** such notes in any P001–P009 entry. Worse, prose contradicts Table 3.2 (prose gives P003/P004/P005/P008 minimal-mode behaviour the table never lists).
- **Verdict: author-should-concede.** The promise was announced but never applied — indefensible if left.
- **Do this:** Append the promised note to all nine entries (Mode-uniform vs Mode-specific with per-mode checks), and reconcile every entry against Table 3.2. Treat line 111 as the checklist.
- **Risk: high.**

### ⚠ Point 9 — Is the P009 (Test-Readiness) PRINCIPLE redundant? — YOU PUSHED BACK
- **Kritikos asks:** P009 seems unnecessary — overlaps both the test-ready MODE and P002–P004.
- **Your email said:** Defended it as a reusable "building block"; argued changing it means new CLI classes for no benefit; asked Kritikos how serious the foul is.
- **Source now:** The principle-vs-mode formalism (lines 91–105) **does** rebut his point (a). But point (b) lands: `registry.go:48` shows P009 fires alongside P002/P003/P004 in test-ready mode, and P009's four checks measurably overlap them. The genuinely distinctive thing — P009's **disjunctive "example OR typed schema"** semantic, evaluated from the test-data generator's fallback perspective — exists in `p009_test_readiness.go` but is NOT raised in the thesis.
- **Verdict: author-partially-right.** Your email defense is weak: it argues implementation convenience and defends the *mode's* security-exclusion as if it justified the *principle*.
- **Do this:** Do NOT send the "building block / new classes" reply. Add 2–3 sentences after line 238 stating the one thing only P009 measures — the disjunctive precondition for non-degenerate test-data generation (example OR typed schema, per-operation, from the generator's fallback contract) — which P002 (examples, conjunctive) and P004 (types, conjunctive) do not express. Frame P009 as the only principle whose unit of analysis is the generated test fixture. Keep "warning" severity.
- **Risk: high.**

### Point 10 — Hive off the overall-system architecture into its own chapter (before the tool chapter)
- **Kritikos asks:** A standalone "System Architecture" chapter BEFORE the DriveBy tool chapter; tool chapter = internals only.
- **Your email said:** Agreed — "I'll add one more agent."
- **Source now:** No system chapter exists. The system-context figure sits *inside* the tool chapter (04:8–17); the cross-system evolution + three-layer model sit *after* it (05:9–91).
- **Verdict: author-should-concede.** This is the structural spine of all Ch. 4/5 comments.
- **Do this:** Create a new chapter "System Architecture" *before* Ch. 4; migrate into it the system-context diagram (from 04:8–17), the "From CLI Tool to Custom Resource" evolution + three-layer model (from 05:9–91), and the CLI→CR table. Narrow Ch. 4 to internal packages only. Update all `\ref`/`\label`. Do NOT just rename Ch. 5.
- **Risk: low.**

### Point 11 — Text-vs-figure inconsistencies in Chapter 4
- **Kritikos asks:** Fix the inconsistencies; make figures correct/professional.
- **Your email said:** "I'll do a pass."
- **Source now:** Still live. `class-diagram.tex` shows APISpec with `Version()/Title()/Operations()` (wrong; code has `Type()/RawVersion()/Info()/Security()`); shows PrincipleChecker with non-existent `Axiom()/Severity()/Modes()`; labels the swagger2 adapter "downgrades to 3.0.x via openapi2conv" while text+source say in-place normalization. The dependency-flow figure says "P001–P008"; the component diagram says "P001–P005, P008, P009".
- **Verdict: author-position-holds** — agreeing is right, *provided the fix is actually made*.
- **Do this:** Regenerate the two stale figures BEFORE replying: (1) APISpec node → `Type/RawVersion/Info/Paths/Components/Security/ValidateStructure(ctx)`; (2) PrincipleChecker node → only `ID()` + `Check(ctx,spec,mode)`; (3) swagger2 label → "converts/normalizes 2.0"; (4) unify principle set across both figures to "P001–P005, P008, P009". Then reply "done."
- **Risk: medium.**

### Point 12 — "Closed-loop negative feedback" → "without effective actuation"
- **Kritikos asks:** Rename — a failure only blocks promotion (preventive), it never improves the spec.
- **Your email said:** Grudging agreement, treated as a routing chore.
- **Source now:** Already conceded in substance: line 338 labels the verdict tier "(open-loop)" and says it "produces no actuation signal u that reduces e(t)… no spec re-write, no implementation patch." Only the section-opener (line 288) + chapter intro (line 4) still use the bare "closed-loop" phrase.
- **Verdict: already-resolved-in-source.** Note: the *platform* tier DOES self-heal (Crossplane/ArgoCD/Promoter), so scope the qualifier to the verdict tier only.
- **Do this:** One-line fix at 06:288 (and 06:4): the negative feedback on the API-contract loop is preventive, not corrective; forward-ref the open-loop analysis at line 338. Don't blanket-rename.
- **Risk: low.**

### ⚠ Point 13 — Multi-API tables contradict the single-injection claim (THE big one)
- **Kritikos asks:** Fix the wrong tables — you claim each defect API is one fault injection, yet tables show most failing MANY principles.
- **Your email said:** "I'll do a pass" (no concrete fix).
- **Source now:** Live contradiction. Lines 220/244/329 assert "violate exactly one principle… others unaffected", but `tab:defect-ground-truth` (342–347) and `tab:local-strict` (396–400) show bad-docs-api failing P002/P003/P004/P008 and no-auth-api failing P002/P003/P004/P005. Plus a wrong cell: staging-gate table (line 427) lists bad-docs-api blocking on "P005 Security" while ground-truth (343) + text (352/435) say it PASSES P005 and blocks on P002/P003/P004. Kritikos also notes slow-api/broken-api carry only RUNTIME defects — their static specs should be clean.
- **Verdict: author-should-concede.** Single biggest correctness item in the thesis.
- **Do this:** Real reconciliation, not a typo sweep. (1) Replace "exactly one principle / others unaffected" (lines 220/244/329) with the defensible framing already at 185/213: one defect *class* → one *distinguishing critical* principle, and documentation-class strips legitimately cascade across P002/P003/P004 (shared substrate). Keep the strict one-principle claim only for the single-API table. (2) Regenerate slow-api/broken-api results so they pass all static principles (blocked only at P006/P007 runtime). (3) Fix Table 7.11 bad-docs blocking cell → "P002/P003/P004 (critical)". (4) Add a sentence stating multi-API counts >1 are *expected* due to shared-substrate co-failure.
- **Risk: high.**

### Point 14 — P007 (Performance) must be Critical
- **Kritikos asks:** Reclassify P007 from Warning to Critical.
- **Your email said:** Full agreement — "Of course it will go in."
- **Source now:** Catalogue still prints P007 = Warning (03:200, 03:262; 07:369), yet Ch. 7 already shows P007 BLOCKING slow-api at staging+production (07:120, 435, 454, 464). The block fires via the load-only CLI command, not the static engine severity map.
- **Verdict: author-partially-right** — agree, but reconcile the self-contradiction.
- **Do this:** Flip P007 → Critical in Table 3.1 (03:262), the P007 subsection (03:200), and the Ch. 7 tables that still print "W" (07:369). Add a sentence to the severity-alignment note: a load-test threshold breach always blocks — there is no non-blocking reading. Keep Status "Planned (wrapper pending)" but note the load-only command already enforces it critically today.
- **Risk: low.**

### Point 15 — Split Table 8.2 effort analysis into code + report deliverables
- **Kritikos asks:** Give measured report-authorship effort, not just code.
- **Your email said:** Full agreement — "It will definitely go in."
- **Source now:** Not satisfied. Table 8.1 mixes Go+LaTeX+YAML+md in one bucket; Table 8.2's "Reviewing… code, manifests, and prose" merges code+report in one row. Report authorship exists only as prose (08:4, 284–289).
- **Verdict: author-should-concede.** Trap: a naive git-numstat split over `.tex` would *understate* report authorship (agent drafted, author edited).
- **Do this:** Add a deliverable dimension — split Table 8.1 into code/manifests vs thesis `.tex`, and add report-effort rows to Table 8.2 as **effort/role estimates** (outline→draft→edit→final-form per chapter), NOT raw line counts. Reuse the existing "author estimates, not wall-clock" caveat (08:65).
- **Risk: low.**

---

## Per-File Work Lists

Every file's complete item table is below (all 480 comments). Sorted within each file by severity (blocking → major → minor). `→pt` = related headline point (0 = none). `M` = mechanical (safe automatable edit).

### 07-evaluation.tex (97 comments)

*Chapter 7 (Evaluation) draws three dominant themes of supervisor concern. First and largest: pervasive table-vs-text inconsistency around the multi-API suite — Kritikos repeatedly objects that APIs derived from non-critical-api by a single fault injection nonetheless show MANY principle failures (P002/P003/P004 across the board) in Tables 7.8/7.10/7.11, which contradicts the single-change derivation claim and the no-auth-api "should fail only P005" expectation; this is his headline point #13 and the single biggest correctness risk in the chapter. Second: the round-N-review provenance scaffolding — non-existent bracketed reference numbers ([494],[509],[370],[466],[551],[473],[490],[550], etc.) and "this addresses my round-1 comment" prose must be stripped and the text made self-contained (headline #5). Third: the severity/criticality model is unsettled across the chapter — P007 (load) must be critical (headline #14), the 5/9-vs-5/7 static-critical ratio, P008/P009 inclusion in static validation, and "all W" table cells that contradict the post-revision critical assignment all need a single coherent pass. Smaller but real: RQ wording drift from Chapter 1, the missing agent-feedback/PoC cross-links to the RQ map, and several concluding paragraphs that Kritikos wants added.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 285 | 147 | blocking | rewrite |  | It seems that the content of the research questions has changed? I ma… | The chapter's RQ1 (line 15: 'Does the SDT conceptual framework provide automated quality… | 2 |
| 290 | 149 | blocking | reference-bug | ✓ | please be aware of the references ... | Line 36 contains non-existent bracketed reference numbers '~[494], [495], [510]'. Per hea… | 5 |
| 298 | 151 | blocking | reference-bug | ✓ | again, we have non-existing reference number. | Line 98: delete the non-existent reference '(per~[509], the supervisor's request for appr… | 5 |
| 300 | 151 | blocking | reference-bug | ✓ | again this is an unknown/non-existing reference number. | This comment attaches to the XSDLC-as-IDP paragraph (line 105-107) which contains 'Per~[3… | 5 |
| 301 | 152 | blocking | cross-cutting |  | ok but based on what was stated in Section 5.4.1, we expected 3*6=18… | Line 107 reports '~45-50 Kubernetes objects' and '44-47 spread' but Chapter 5 Section 5.4… | 13 |
| 307 | 155 | blocking | reference-bug | ✓ | again non-existing reference numbers ... | Line 176: the 'Naming note (round-1 review~[466], [476]).' lead-in contains non-existent… | 5 |
| 323 | 160 | blocking | reference-bug | ✓ | Maybe it is better in general not to indicate what was done on docume… | Line 352: remove the round-1-review provenance and reference number — 'Two observations f… | 5 |
| 324 | 160 | blocking | table-fix |  | Are you sure that the table's content is correct? It seems that many… | CORE headline-#13 issue. Table 7.8 (tab:defect-ground-truth, lines 343-346) shows bad-doc… | 13 |
| 328 | 161 | blocking | pushback |  | but the shared base specification does not have any critical failure… | Line 352 claims slow-api and broken-api 'inherit P002/P003/P004 critical failures from th… | 13 |
| 331 | 161 | blocking | table-fix |  | what do you mean with all W? Warning? But you have corrected this by… | Table 7.9 (tab:api-suite, line 367) bad-docs-api 'Expected Failure' cell reads 'P002, P00… | 14 |
| 333 | 162 | blocking | table-fix |  | Shouldn't no-auth-api fail only on P005? | Table 7.10 (tab:local-strict, line 398): no-auth-api row shows P002 \xmark, P003 \xmark,… | 13 |
| 334 | 162 | blocking | table-fix |  | These two APIs should have only runtime faults so they should pass th… | Table 7.10 (tab:local-strict, lines 399-400): slow-api and broken-api each show P002/P003… | 13 |
| 336 | 162 | blocking | table-fix |  | I do not understand this. You should have only P005. You do not get P… | Definitive statement of the headline-#13 problem for no-auth-api. Confirm in text (line 4… | 13 |
| 340 | 163 | blocking | table-fix |  | This is wrong. P002-P004 failed which are critical | Table 7.11 (tab:staging-gates, line 427): bad-docs-api 'Blocking Failure' column reads 'P… | 13 |
| 342 | 163 | blocking | rewrite | ✓ | This is totally wrong here + you talk again about revision while it i… | Line 435, final sentences: 'Note: this paragraph reflects the post-revision severity. In… | 5 |
| 344 | 164 | blocking | pushback |  | Question: if P007 is warning, then why do we have failure? In any cas… | This is headline #14: Kritikos insists P007 (Performance/Load) MUST be critical. The chap… | 14 |
| 347 | 166 | blocking | pushback |  | I disagree here for the last two APIs!!! | Sentence at line 508 (continuation): 'The four defect APIs (bad-docs-api, no-auth-api, sl… | 13 |
| 360 | 169 | blocking | figure |  | Based on what we see in Figure 7.2 there is no corroboration. The res… | Line 601 + Figure 7.2 (fig:pass-rates, figures/per-principle-pass-rates.tex): the text cl… | 13 |
| 367 | 172 | blocking | reference-bug | ✓ | Again, we have a non-existing reference number | Line 683: 'The experiment was added in direct response to the supervisor's round-1 review… | 5 |
| 376 | 176 | blocking | rewrite |  | correct, although it is not clear how P7 that is warning blocks the A… | Summary result 2 (line 758): 'slow-api at the runtime layer on P007'. Kritikos repeats th… | 14 |
| 284 | 147 | major | rewrite |  | You forget the SDT Feedback as Input to Agent-Driven Specification Im… | In the chapter intro (line 4, 'across three dimensions ... proof-of-concept ... controlle… | 1 |
| 286 | 148 | major | rewrite |  | RQ2 has two parts: automatic production of validation, testing and ga… | RQ2 paragraph (line 16) currently only cites Sections 7.2 and 7.5 and only addresses the… | 2 |
| 287 | 148 | major | rewrite |  | Is also large-scale evaluation relevant here? | In the RQ1 list of supporting experiments (line 15) the large-scale APIs.guru evaluation… | 2 |
| 288 | 148 | major | rewrite |  | Chapter 10 is a conclusions section that explains how the RQ is answe… | RQ4 paragraph (line 17) says it is 'Addressed through architectural analysis in Chapters… | 2 |
| 289 | 148 | major | figure |  | Where is PoC in this figure? In addition, you have effort comparison… | Figure 7.1 (fig:eval-methodology, figures/evaluation-methodology.tex) shows columns 'Cont… | 11 |
| 291 | 149 | major | structural |  | Ok but why do you have section 7.6 separate from 7.2? It can be argue… | Kritikos wants the CI/CD Integration Metrics section (currently \section sec:cicd-metrics… | 0 |
| 293 | 150 | major | rewrite |  | I believe that the first configuration step is to install the DriveBy… | In subsection 'Proof-of-Concept Configuration' (line 48, subsec:operational-poc-config) t… | 0 |
| 294 | 150 | major | table-fix |  | But in the production gate, we do not have the test-ready/validate-on… | Table 7.1 (tab:operational-poc-config, lines 54-70): the perfect-api production gate cell… | 13 |
| 295 | 150 | major | table-fix |  | See previous comment. We do not see this in table. | Same locus as idx 294: the prose at line 72 describes a strict check on the production ga… | 13 |
| 296 | 150 | major | table-fix |  | we do not see this on broken-api ... | Table 7.1 (tab:operational-poc-config) broken-api row (line 67) has '---' for the product… | 13 |
| 297 | 150 | major | table-fix |  | Please check table's content as it seems inconsistent with the text t… | Table 7.1 (tab:operational-poc-config) rows for bad-docs-api and no-auth-api (lines 64-65… | 13 |
| 299 | 151 | major | rewrite |  | if we have 44 workflow runs, how is it possible to have only 32 commi… | Line 98 reports 44 Workflow runs but 32 CommitStatus CRDs without explaining the discrepa… | 13 |
| 304 | 152 | major | rewrite |  | ok but the focus is on structural drift, not semantic. Semantic drift… | In the specification-contract enumerated item (line 119) the sentence 'Any drift between… | 0 |
| 305 | 153 | major | cross-cutting |  | This is at the very heart of your thesis. I would better position thi… | The scope-defining statement at line 123 ('This is not a claim that SDT replaces all test… | 0 |
| 306 | 154 | major | rewrite |  | ok but what does this mean in practice? Is there something that needs… | The health-check timeout edge case is marked with \xmark (fail) in Table 7.2 (tab:edge-ca… | 13 |
| 308 | 155 | major | rewrite |  | ok but does this mean that non-critical-api also passes dynamic valid… | In the naming-note paragraph (line 176) and Reference API Design (line 185), state explic… | 0 |
| 315 | 159 | major | pushback |  | yes, this is defensible. You actually removed something that would no… | In 'Interpretation of misses' (line 309), the text currently calls BOTH the p005-strip-op… | 0 |
| 317 | 159 | major | pushback |  | Well, I am not sure about the first issue. I believe that it is a pro… | Reinforces idx 315: the 'first issue' (p005-strip-op-security) should be reframed as an e… | 0 |
| 320 | 159 | major | rewrite |  | Missing a concluding paragraph here. Are you satisfied with the produ… | At the end of the Per-Check Extension subsection (after line 315, the Determinism check,… | 0 |
| 321 | 160 | major | structural |  | This table should be presented before Table 7.8. First, we see the in… | Reorder the two tables in subsection 'Evaluation API Suite Design': move Table 7.9 'Evalu… | 13 |
| 322 | 160 | major | table-fix |  | But you report them with these values and normal ones like 'runtime'… | Table 7.8 (tab:defect-ground-truth, lines 337-349): the P006/P007 column mixes the placeh… | 13 |
| 325 | 160 | major | rewrite |  | I would better explain how each API was derived from non-critical-api… | In subsection 'Evaluation API Suite Design' (line 325) add an explicit per-API derivation… | 13 |
| 326 | 161 | major | table-fix |  | /P008 — Please see also the table content and make necessary correcti… | Line 352 prose lists bad-docs-api blocking on 'P002/P003/P004' but Table 7.8 (line 343) a… | 13 |
| 327 | 161 | major | table-fix |  | but it seems that it also exhibits additional failures by observing T… | Same paragraph (line 352): the claim 'no-auth-api on P005' understates — Table 7.8 (line… | 13 |
| 332 | 161 | major | table-fix | ✓ | What does 'critical sweep' mean? It is not clear. Maybe remove it? | Table 7.9 (tab:api-suite, line 368) no-auth-api Expected Failure cell reads 'P005 (critic… | 13 |
| 335 | 162 | major | table-fix |  | It also fails P008, why? | Kritikos asks why bad-docs-api also fails P008. In Table 7.8 (tab:defect-ground-truth, li… | 13 |
| 337 | 162 | major | table-fix |  | See previous comments ... | Pointer back to idx 333-336. The paragraph at line 406 ('Under the post-revision severity… | 13 |
| 338 | 162 | major | rewrite |  | Shouldn't you explain first what was involved in each gate and which… | At the start of subsection 'Staging Gate Results' (line 409) and before Table 7.11 (tab:s… | 13 |
| 339 | 163 | major | table-fix |  | I am puzzled here as you need to have the same checks per gate. It se… | Table 7.11 (tab:staging-gates, lines 425-429): non-critical-api uses 'validate-only (test… | 13 |
| 341 | 163 | major | rewrite |  | But there was no blocking for this API for the current gate. As it pa… | In the staging-results prose (line 435), the description of broken-api being blocked at t… | 13 |
| 343 | 164 | major | rewrite |  | Ok but this delay should have been explained in 7.4.1 how it is appli… | The 500ms artificial delay of slow-api is discussed in Production Gate Results (line 464)… | 13 |
| 345 | 165 | major | rewrite |  | If these two principles are to be ignored (correct, as they are not y… | Subsection 'Severity Model in Practice' (line 500): the text uses a '5/9 static-Critical… | 14 |
| 346 | 165 | major | pushback |  | I do not agree with the position of this sentence. You talk only abou… | Bullet at line 508 ('As a result, only non-critical-api...passes all gates. The four defe… | 13 |
| 350 | 166 | major | rewrite |  | Do you also need to include P009 here? Isn't it also related to stati… | Line 524: 'Only static validation principles (P001--P005, P008) are evaluated'. Kritikos… | 8 |
| 351 | 167 | major | rewrite |  | Could indicate how many specs were included in the dataset in 7.5.1.… | In subsection 'Dataset and Protocol' (7.5.1, line 520) add the dataset size and OpenAPI-v… | 0 |
| 352 | 167 | major | rewrite |  | Ok but isn't this part of the original dataset? Shouldn't we indicate… | Line 550 presents the 50-spec 3.x sample, and Section 7.5.4 (subsec:swagger2-eval, line 6… | 0 |
| 353 | 167 | major | structural |  | Ok but why do you report them separately? This needs some explanation… | Kritikos wants global-then-local result presentation: first aggregate results across all… | 0 |
| 356 | 168 | major | reference-bug | ✓ | 'security' | Line 573 contains 'Per~[490], the P005 result deserves a separate reading'. The one-word… | 5 |
| 359 | 169 | major | rewrite |  | But if we have 46% for P001, even the specification form is not prope… | Line 599 ('The gap between P001 (medium pass rate, 46%) and P002--P004... confirms that e… | 0 |
| 361 | 170 | major | reference-bug | ✓ | round 6? Do you need to report such information in the formal thesis… | Line 619 (\paragraph{Method.}): 'The harvester ... was extended in round 6 with a 2.0 ver… | 5 |
| 362 | 171 | major | rewrite |  | But in the 3.x results, P001 was not passed universally!!!! | Line 644 (Swagger 2.0 score distribution): 'P001 again passes universally---each spec is… | 13 |
| 363 | 172 | major | table-fix |  | I would prefer to see a comparison table showcasing the improvement o… | This comment attaches near Table 7.17 (CI/CD integration metrics, tab:cicd-metrics, line… | 15 |
| 365 | 172 | major | table-fix |  | Are these all normal CI/CD integration metrics? They seem more like t… | Table 7.17 (tab:cicd-metrics, lines 655-674) lists footprint/technical metrics (binary si… | 0 |
| 368 | 173 | major | rewrite |  | Do you have the prompt? If yes, please document it somewhere (either… | In the agent-revision step (line 696), document the exact prompt/instructions given to th… | 0 |
| 371 | 174 | major | rewrite |  | Food for thought: (a) how can the agent know which operations are to… | In the Interpretation subsection (lines 737-745) add a discussion paragraph addressing Kr… | 1 |
| 373 | 175 | major | rewrite |  | Isn't this also covered by the PoC and its integration CI/CD metrics?… | Summary intro (line 752) lists six results tagged to RQs but never mentions the CI/CD int… | 0 |
| 374 | 175 | major | rewrite |  | Maybe this is also partially covered by PoC and Multi-API controlled… | In the Summary's RQ4 handling (line 752, 'RQ4 ... is intentionally absent from this summa… | 2 |
| 375 | 176 | major | rewrite |  | Please check if this rate/percentage is eventually correct | Summary result 1 (line 756): 'Twelve of fifteen single-check defects were detected (80%);… | 13 |
| 378 | 176 | major | rewrite |  | Forgot to say something about OpenAPI 2.x dataset and its evaluation… | Summary result 4 (line 762, 'Specification quality gap, quantified') reports only the 50-… | 0 |
| 380 | 177 | major | rewrite |  | Did not say something about CI/CD integration metrics | Summary result 6 (line 766, RQ3 / operational deployment) reports resource counts and edg… | 0 |
| 292 | 149 | minor | structural |  | Ok but is this part of the cluster? I guess not, so I would put it in… | In subsection 'Initial Cluster Situation' (line 39, label subsec:operational-initial-stat… | 0 |
| 302 | 152 | minor | trivial | ✓ | I would change paragraph here. | The XSDLC-as-IDP-feature paragraph (line 107) is a single very long paragraph running fro… | 0 |
| 303 | 152 | minor | rewrite |  | well, they are needed in one case (autoMerge: false). But the validat… | In the 'Quality Assurance Without Dedicated Testers' subsection (line 114), the claim tha… | 0 |
| 309 | 155 | minor | trivial | ✓ | current | OCR-fragment comment ('current') attaching to the naming-note paragraph (line 176), almos… | 5 |
| 310 | 155 | minor | rewrite |  | how would you characterise this API in terms of functionality? Produc… | In Reference API Design (line 185) add a one-line functional characterisation of the refe… | 0 |
| 311 | 157 | minor | trivial |  | 'produced' — to indicate that this was indeed the case — you actually… | Line 244: Kritikos endorses the determinism claim and suggests wording to make clear it w… | 0 |
| 312 | 158 | minor | rewrite |  | Maybe it is better to say that you have created three mutations cover… | This maps to the combination-robustness description. Line 287 / Table 7.6 (tab:combo-dete… | 0 |
| 313 | 158 | minor | pushback |  | In essence, this is good as an evaluation approach. But it is not com… | Kritikos concedes this is optional. Add a short limitations sentence after the per-check… | 0 |
| 314 | 158 | minor | rewrite |  | I would not count this as an issue as it actually leads to a parse er… | Table 7.5 (tab:per-check-defect-extension) row 'strip schema length/range constraints P00… | 0 |
| 316 | 159 | minor | rewrite |  | ok, so this indicates that this is truly a false negative, marking in… | Counterpart to idx 315: Kritikos agrees the p009-strip-param-examples miss IS a genuine f… | 0 |
| 318 | 159 | minor | pushback |  | I agree here. Again, this is a matter of how evaluation was implement… | Agreement comment reinforcing that the p005-strip-op-security case is an evaluation-imple… | 0 |
| 319 | 159 | minor | pushback |  | Ok, I agree that it relates to P001 theoretically but it is caught by… | Concerns the parse-error interpretation (line 313, 'Interpretation of the parse error').… | 0 |
| 329 | 161 | minor | trivial | ✓ | 'addresses' | OCR-fragment wording suggestion on line 352 ('address the contradictions flagged'). Likel… | 5 |
| 330 | 161 | minor | rewrite |  | ok but it is quite easy to create such an OpenAPI document. Much easi… | Relates to line 354 ('A critical-P001 example would require a deliberately malformed Open… | 13 |
| 348 | 166 | minor | trivial | ✓ | 'harvested' | Wording tweak at line 524: change 'The evaluation harvests OpenAPI specifications' to pas… | 0 |
| 349 | 166 | minor | trivial | ✓ | 'created' — includes .... | Wording tweak at line 524: Kritikos suggests 'created includes' — i.e. change 'The datase… | 0 |
| 354 | 167 | minor | rewrite |  | The mapping from qualitative values to quantitative constraints could… | Line 552 declares the bucket thresholds 'high >=80%, medium 30-79%, low <30%'. Kritikos w… | 0 |
| 355 | 167 | minor | trivial | ✓ | 'produced' — is the following: | OCR-fragment wording note at line 550/552. Likely a suggestion to introduce the results t… | 0 |
| 357 | 168 | minor | rewrite |  | and usage/documentation gap in the sense that it is not clear to the… | Line 573, P005 reading: the text calls the 37-failing-API pattern 'more often a test-orie… | 0 |
| 358 | 169 | minor | rewrite |  | So, this API (non-critical-api) has sufficient information for API te… | Line 597 references 'the controlled-evaluation result that DriveBy reports zero false pos… | 0 |
| 364 | 172 | minor | table-fix |  | In the context of given constraints, could supply also the actual ran… | Table 7.17 (tab:cicd-metrics, line 666): for constraint-style rows like 'Startup time <1s… | 0 |
| 366 | 172 | minor | table-fix |  | Can size-related metrics be considered as CI/CD integration metrics?… | Table 7.17 (tab:cicd-metrics, lines 664-665): 'CLI binary size ~15MB' and 'Docker image s… | 0 |
| 369 | 173 | minor | rewrite |  | Question: if it had in its possession the definition of the principle… | After the sole-feedback constraint (line 696) or in the Interpretation subsection (lines… | 1 |
| 370 | 174 | minor | trivial | ✓ | 'addressing' | OCR-fragment wording note on Table 7.18 (tab:sdt-feedback-delta) P003 row (line 720: 'Par… | 0 |
| 372 | 175 | minor | rewrite |  | ok, let's leave this as future work ... | Interpretation subsection, line 745 ('Broader claims ... require the multi-spec, multi-it… | 0 |
| 377 | 176 | minor | rewrite |  | what's 418? | Summary result 3 (line 760) references 'broken-api's 418' without explanation. HTTP 418 (… | 0 |
| 379 | 177 | minor | trivial | ✓ | 'maps to' .... | OCR-fragment wording note on Summary result 5 (line 764, RQ2 / agent-feedback experiment)… | 0 |

### 03-methodology-ddt.tex (61 comments)

*Kritikos's 61 comments on the methodology chapter cluster around four dominant themes. (1) The axiom-to-principle relationship: he repeatedly pushes back on any one-to-one "this principle is governed by axiom X" framing and wants every principle shown to relate to ALL three axioms, with an explicit Governing vs Impacted/Effective distinction in the mapping table plus justification (headline point #7). (2) Mode-specific behaviour discipline: he wants each principle's mode behaviour stated up front (right after Status), consistent terminology ("mode-specific" not "mode-aware"), the "promised" Mode-uniform/Mode-specific notes actually present, and test-ready behaviour enumerated as concrete checks rather than prose. (3) Stale provenance artefacts: non-existent bracketed reference numbers ([161],[172],[180],[192],[194]) and repetitive "round-N review" framing must be removed or converted to self-contained prose (headline point #5). (4) Severity/scope correctness: P007 performance must be Critical not Warning (headline #14), P009 may be redundant with the test-ready mode plus P001-P004, the test-only mode was neglected in the principle analysis, and the modes should map onto standard organisational gates (pre-flight/pre-merge/post-merge/release). Note the live .tex has been substantially restructured since the annotated PDF — per-principle "Axiom:" labels and the old "Mapping of SDT axioms" table (Table 3.1) are gone — so several comments are partially superseded but their underlying asks remain open.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 104 | 53 | blocking | reference-bug | ✓ | What do you mean here? In addition, you have 72 references and not [1… | Line 15 contains the stale provenance marker '\textit{What is being validated (round-1 re… | 5 |
| 108 | 53 | blocking | pushback |  | In the end, should they cover all axioms or at least one of them? It… | This is headline point #7 and the author already agrees (every axiom governs every princi… | 7 |
| 110 | 55 | blocking | pushback |  | To me, it seems that based on the way determinism is defined, it must… | Headline point #7. Live text at line 54 already de-emphasises P006 ('Determinism is there… | 7 |
| 111 | 55 | blocking | reference-bug | ✓ | Again here you have a non existing reference number ... | Line 56 contains '(round-1 review~[172])'. Delete the parenthetical so the heading reads… | 5 |
| 114 | 56 | blocking | pushback |  | So, the other principles do not cater for observability? This seems q… | Headline point #7. The Observability subsection (lines 65-75) must not imply observabilit… | 7 |
| 117 | 57 | blocking | reference-bug | ✓ | Again non-existing reference | Line 84 contains '\textit{Severity semantics (round-1 review~[180]).}'. Delete '(round-1… | 5 |
| 121 | 57 | blocking | structural |  | I disagree with the way the mode-specific behaviour is stressed. It i… | Headline point #8. The 'Mode coverage' note at line 111 promises a Mode-specific/Mode-uni… | 8 |
| 126 | 59 | blocking | reference-bug | ✓ | non-existing reference also here | Kritikos flags a non-existent reference in the P003 area. In the live .tex the stale numb… | 5 |
| 134 | 61 | blocking | pushback |  | Not sure whether this should be warning. This needs to be properly ju… | Headline point #14 — Kritikos INSISTS P007 must be Critical and the author has AGREED ('O… | 14 |
| 138 | 62 | blocking | pushback |  | The principle has a name that directly maps to the test-ready mode. T… | Headline point #9 — Kritikos questions whether P009 is redundant (overlaps test-ready MOD… | 9 |
| 142 | 63 | blocking | table-fix |  | In this case, it is better to enrich the table also with the impacted… | Headline point #7. Kritikos wants the catalogue/mapping table to carry axiom information… | 7 |
| 143 | 64 | blocking | table-fix |  | I propose to change the table content as follows: - provide first ID… | Headline point #7, explicit table spec. The live catalogue table (lines 252-266) already… | 7 |
| 146 | 65 | blocking | rewrite |  | I like the justification but it needs to be enhanced to clarify that… | Headline point #7 — the central ask. Kritikos explicitly endorses a Governing(overarching… | 7 |
| 149 | 65 | blocking | pushback |  | Sure but it is catastrophic to violate a non-functional requirement (… | Headline point #14. Reinforces that P007 (performance, a non-functional requirement) must… | 14 |
| 151 | 66 | blocking | table-fix |  | Correct, we see this in the first three modes where the number of pri… | Multi-part, high-priority. (1) Rationale for test-only mode: add a sentence in Section 3.… | 8 |
| 154 | 66 | blocking | pushback |  | I would remove P009 as it is actually implemented by P001 and partial… | Headline point #9 — strong redundancy challenge against P009 (overlaps P001 + partial P00… | 9 |
| 105 | 53 | major | rewrite |  | Again not clear what do you mean here. | The 'two artefacts must not be conflated' explanation at line 15 is dense and reads as an… | 5 |
| 106 | 53 | major | structural |  | I believe that this and previous sentences must be removed. | Kritikos wants the round-1-review explanatory block at line 15 (the 'two artefacts'/'earl… | 5 |
| 107 | 53 | major | rewrite |  | This is the third axiom that was not mentioned before in the paragrap… | At line 15 the trailing clause introduces observability ('if results are observable, they… | 0 |
| 109 | 54 | major | rewrite |  | Ok - we agree here. But in the definition of completeness in 3.2.1, y… | Headline point #6. The Axiom of Completeness definition at lines 35-43 does not make prec… | 6 |
| 112 | 55 | major | rewrite |  | Ok but functional testing is also about the actual content being retu… | At line 56 the text says P006 asserts 'status code, schema shape, declared header presenc… | 0 |
| 113 | 55 | major | rewrite |  | Exactly, this is also indicated in my previous comment. Thus, the spe… | Reinforce the link Kritikos draws: at line 56 the final sentence ('Specifications that de… | 6 |
| 115 | 56 | major | rewrite |  | Can't we consider completeness also as a property? If yes, then other… | Kritikos wants symmetry: if observability is a property of every principle's output, comp… | 7 |
| 123 | 58 | major | structural |  | Ok but is it mode-uniform? This is not clarified via a "promised" not… | Headline point #8. P001 ends at line 124 ('P001 runs in every validation mode...') with n… | 8 |
| 124 | 58 | major | rewrite |  | Correct - it also prohibits P006 & P007. Without examples, it is impo… | On the P002 severity-escalation note (line 133), Kritikos affirms and wants the consequen… | 0 |
| 125 | 59 | major | trivial | ✓ | mode-aware vs mode-specific. Maybe use the second term as it was intr… | Terminology consistency. Line 150 (P003), line 178 (P005), line 217 (P008) all say 'is mo… | 8 |
| 127 | 59 | major | pushback |  | Ok, I agree with this handling. But in essence, non-uniformity is not… | Two asks on P003 (lines 150, 154). (1) Confirm in the P003 test-ready note (Table footnot… | 8 |
| 128 | 59 | major | rewrite |  | Do not explain what is covered in test-ready mode as you indicate tha… | Headline point #8. P003 prose (line 150) describes minimal and strict modes but NOT test-… | 8 |
| 130 | 60 | major | rewrite |  | Again you do not indicate what holds in test-ready mode. | Headline point #8. P005 (line 178) describes only minimal and strict modes. Per Table 3.2… | 8 |
| 131 | 60 | major | rewrite |  | General question: if an API is totally public, does this principle ap… | Add a sentence to the P005 subsection (after line 180) addressing fully-public APIs: P005… | 0 |
| 132 | 61 | major | rewrite |  | Ok but you need to stress that you rely on the examples given to assu… | In the P006 subsection (line 193, 'P006 operationalizes the Axiom of Determinism: ... tes… | 0 |
| 133 | 61 | major | rewrite |  | See previous comment ... I totally agree but only the examples can be… | Add a short paragraph to the P006 subsection (after line 191) acknowledging two example-d… | 0 |
| 135 | 62 | major | rewrite |  | Issue & food for thought: in case we have non-determinism in the envi… | Add a paragraph to the P007 subsection (after line 204) addressing environmental non-dete… | 14 |
| 136 | 62 | major | rewrite |  | Did not justify why the affected axiom is observability ... Same for… | The annotated PDF showed P008 with an 'Axiom: Observability' label that the live .tex has… | 7 |
| 137 | 62 | major | rewrite |  | Do not clarify what happens with test-ready mode. | Headline point #8. P008 (line 217) describes minimal and strict modes only. Per Table 3.2… | 8 |
| 140 | 63 | major | pushback |  | I disagree here. As in functional testing you need to check the respo… | Kritikos disagrees that response examples are 'not required for meaningful validation' (l… | 0 |
| 141 | 63 | major | rewrite |  | Ok but this is risky in terms of non-determinism. | On P009 check (2) at line 233 ('request body builder prioritizes explicit examples, then… | 0 |
| 144 | 64 | major | structural |  | I believe that as this is a repetition, it does not need to be suppli… | Headline point #5 (handle comment once). The 'Severity alignment (round-2 revision)' bloc… | 5 |
| 145 | 65 | major | rewrite |  | Ok but don't they also have observability impact? | Headline point #7. This comment targets the old 'completeness principles P001-P005 enable… | 7 |
| 147 | 65 | major | rewrite |  | So, the overarching axiom is also determinism or it's just an impacte… | Headline point #7. Targets the removed 'P009 is a determinism check by nature' clause (AN… | 7 |
| 148 | 65 | major | rewrite |  | What about the remaining observability principles? Do they also affec… | Headline point #7. Targets the removed group-reasoning paragraph (ANCHOR-NOT-FOUND). If t… | 7 |
| 150 | 65 | major | pushback |  | What about non-functional testing? Isn't it also important? | Headline point #14. Same thrust as idx 149 — Kritikos wants non-functional testing (P007… | 14 |
| 156 | 67 | major | pushback |  | If I am not mistaken, usually a pre-merge gate also includes function… | On the Strict-mode paragraph (line 311, 'Strict mode evaluates P001 through P005 and P008… | 0 |
| 158 | 67 | major | pushback |  | Why not also include P005? It can enable functional testing to succee… | On the Test-ready-mode paragraph (line 315, currently 'evaluates P001, P002..., P003...,… | 0 |
| 159 | 67 | major | rewrite |  | Very good: maybe also indicate that this is normal practice as gates… | On the mode-progression paragraph (line 317, 'The four modes form a progression... minima… | 0 |
| 162 | 67 | major | rewrite |  | map to pre-merge or post-merge/release and justify this mapping. | On the mode-progression paragraph (line 317) where test-only is described for runtime ver… | 0 |
| 102 | 52 | minor | structural |  | I believe that this is a repetition so it could be safely removed. | The flagged opening sentence of Section 3.1 ('Chapter 2 identified three paradigms... SDT… | 5 |
| 103 | 53 | minor | trivial | ✓ | (consumer) | Inline word suggestion: Kritikos proposes 'consumer' where the text says 'contract-first… | 0 |
| 116 | 56 | minor | rewrite |  | Does the same hold for the other principles? | Same observability-for-all-principles concern as idx 114/115, attached to the longitudina… | 7 |
| 118 | 57 | minor | trivial |  | meaning | Inline 'meaning' margin note on the Severity-semantics paragraph (line 84), asking the au… | 0 |
| 119 | 57 | minor | trivial |  | meaning | Second inline 'meaning' note on the same Severity-semantics paragraph (line 84), likely o… | 0 |
| 120 | 57 | minor | trivial | ✓ | of the respective principle | Inline insertion suggestion: Kritikos wants 'of the respective principle' added for preci… | 0 |
| 122 | 57 | minor | trivial | ✓ | utilised | Inline word change in the P001 paragraph (lines 120-122): Kritikos suggests 'utilised' fo… | 0 |
| 129 | 60 | minor | trivial | ✓ | By global requirements you mean global authentication/security requir… | On P005 line 178, the phrase 'global requirements are set' is ambiguous. Change to 'globa… | 0 |
| 139 | 63 | minor | rewrite |  | Is this meaningful to say? If the severity is warning for this princi… | On P009 check (4) at line 235 ('2xx responses have examples defined. Advisory check (warn… | 0 |
| 152 | 66 | minor | trivial | ✓ | the test-only mode? | On the Table 3.2 footnote at line 303 ('In the current implementation, ForMode(test-only)… | 0 |
| 153 | 66 | minor | structural |  | I propose to remove the other note parts here as they should be suppl… | Headline point #5/#8. The Table 3.2 footnotes at lines 304-306 (dagger P004, ddagger P002… | 8 |
| 155 | 67 | minor | trivial | ✓ | a baseline? | On the Minimal-mode paragraph (line 309), Kritikos suggests describing minimal mode as 'a… | 0 |
| 157 | 67 | minor | trivial | ✓ | Ok, now it is clear why you have both strict mode and test-only mode. | Positive acknowledgement on the Test-only-mode paragraph (line 313) — no change required.… | 0 |
| 160 | 67 | minor | trivial | ✓ | pre-flight gate? | Inline suggestion on the mode-progression paragraph (line 317) where test-ready is descri… | 0 |
| 161 | 67 | minor | trivial | ✓ | maybe say pre-merge gate | Inline suggestion on the mode-progression paragraph (line 317) where strict is mapped to… | 0 |

### 02-related-work.tex (53 comments)

*Kritikos's 53 comments on the Related Work chapter cluster around three dominant themes. First, taxonomy and terminology rigor: he insists linting is not testing (it is static analysis / verification), that "contract testing" must be split into consumer-driven vs provider-driven sub-types (with Dredd reclassified as provider-driven), and that the chapter's paradigm should be called "Specification-Driven" to match the SDT name rather than "Documentation-driven." Second, a sustained challenge to the Table 2.1 "single-artifact" criterion (echoing headline points 3 and his repeated inline pushback at idx 86-89, 96): he argues every testing tool already consumes a single artifact (the OpenAPI spec), so the criterion as worded is wrong and either its definition or the per-tool evaluations must change. Third, a structural proposal to move the OpenAPI Ecosystem section to the front (as Section 2.1) and merge/clarify the overlapping Dredd section, plus a recurring demand to (a) correlate research gaps with the DriveBy tool and Chapter 4 (headline #4), (b) clarify whether validation/reconciliation targets the spec, the source code, or both, and (c) reference the agent-driven reconciliation evaluation in several places (headline #1).*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 78 | 43 | blocking | pushback |  | Ok but in principle you validate both the API documentation and the s… | Policy-as-Code subsection (lines 150-154). This is a recurring, defense-relevant ambiguit… | 3 |
| 82 | 44 | blocking | pushback |  | Correct. But the API source code should also be corrected or not? As… | Line 163, on the DriveBy reconciliation agent observing 'the specification's actual state… | 3 |
| 86 | 47 | blocking | pushback |  | OK but based on the definition of single-artifact, it seems that it s… | Table 2.1 legend / 'single-artifact' definition (line 209 and line 237). This is headline… | 3 |
| 87 | 47 | blocking | table-fix |  | I also do not understand why it is not a single-artifact approach, as… | Table 2.1 rows for Schemathesis and RESTler (lines 225-226) both mark Single-Artifact = \… | 3 |
| 88 | 47 | blocking | table-fix |  | See previous comment. All testing approaches take into consideration… | Continuation of idx 86/87 for the RESTler and Pact rows (lines 226-227). Pact's Single-Ar… | 3 |
| 93 | 48 | blocking | rewrite |  | Why don't you call this Specification-Driven? As you define your meth… | Paradigm list (line 252): the third paradigm is labelled '\item[Documentation-driven.]' b… | 0 |
| 96 | 49 | blocking | pushback |  | Still all are considered a single kind of source or artifact, right?… | Gap 5 'No single-artifact methodology' (line 273). Same core issue as idx 86-89. Kritikos… | 3 |
| 49 | 34 | major | rewrite |  | Linting is not a kind of testing. It is usually complementary to test… | On line 11, the sentence 'We organize these approaches into three categories based on the… | 0 |
| 50 | 34 | major | rewrite |  | It is better to indicate this as consumer-driven contract testing if… | Subsection 2.1.1 (line 16) currently treats 'contract testing' and 'consumer-driven contr… | 0 |
| 56 | 36 | major | rewrite |  | Yes, this is a kind of static analysis. Testing is about checking the… | This confirms idx 49. The 2.1.3 heading is already 'Specification Linting and Static Anal… | 0 |
| 60 | 37 | major | rewrite |  | Ok but if Chapter 3 does not explain the semantics of all these secti… | Line 56 ends 'The validation principles introduced in Chapter~\ref{ch:methodology} map di… | 0 |
| 63 | 37 | major | structural |  | I propose to move this section as Section 2.1, as it provides general… | Major reorganization request: move the entire 'The OpenAPI Ecosystem' section (lines 48-7… | 0 |
| 65 | 38 | major | structural |  | I do not understand why this section is needed as it greatly overlaps… | Subsection 2.2.3 'API Testing Tools: Dredd and Beyond' (lines 71-78) overlaps the contrac… | 0 |
| 67 | 39 | major | rewrite |  | Is there a proof for this? In the context of a reference? | Line 88: the claim 'These principles eliminate imperative drift' is asserted without supp… | 0 |
| 68 | 39 | major | rewrite |  | But these are specifications? They indicate the desired state of an A… | Kritikos wants the GitOps-principles paragraph (line 88) and the following paragraph (lin… | 1 |
| 69 | 39 | major | reference-bug | ✓ | missing reference here to Kong | Kong is now mentioned in the live text (line 92: 'API gateways such as Kong enforce speci… | 0 |
| 71 | 39 | major | rewrite |  | Ok but are these requirements somewhere documented and justified? How… | On the four OpenGitOps principles (line 88), Kritikos asks for justification linking them… | 1 |
| 75 | 41 | major | rewrite |  | Is there a reference to back this up? | Line 122: the GitOps Promoter description ('watching for commit status signals and mergin… | 0 |
| 76 | 42 | major | rewrite |  | Thus, while ..., SDT can now solve this ... | On the Terraform/IaC paragraph (lines 134-138), Kritikos wants an explicit 'while X is un… | 4 |
| 77 | 42 | major | rewrite |  | Would this be an API quality assessment or requirement? It depends on… | Crossplane subsection (lines 143-145), specifically the claim that the Kubernetes reconci… | 0 |
| 81 | 43 | major | rewrite |  | + XSDLC as it clarifies the environments, when promotion applies to t… | In the ontological-reconciliation tool list (line 159, 'Terraform, Crossplane, OPA, Argo~… | 3 |
| 83 | 45 | major | rewrite |  | I agree here. But you also cover the quality of the source code of th… | Platform-engineering subsection (lines 174-184), on the spec as a platform artifact reduc… | 0 |
| 85 | 46 | major | cross-cutting |  | Correct -- but also refer to the evaluation that led to a reconciliat… | Section 2.5 (line 199, the SDT-as-feedback-for-agents paragraph) must add a forward-refer… | 1 |
| 89 | 47 | major | table-fix |  | Why do we have partial support for this capability? The server stub i… | Table 2.1 Microcks row (line 229): Static = $\sim$ and GitOps-Native = $\sim$. Kritikos q… | 3 |
| 91 | 48 | major | rewrite | ✓ | Consumer Contract-First -> to make the distinction wrt the Provider C… | Paradigm list (line 250): rename the '\item[Contract-first.]' label to '\item[Consumer co… | 0 |
| 97 | 50 | major | rewrite |  | Of course, there is a tool constructed in the case of a PhD I supervi… | Gap 6 'No knowledge transfer assurance' (line 275): the absolute claim 'none assesses whe… | 0 |
| 99 | 50 | major | rewrite |  | In essence, apart from the 2nd sentence, you do not talk about gap de… | Line 280: Kritikos is correct that most sentences ('reflects the field's focus on...', 'f… | 0 |
| 100 | 50 | major | cross-cutting |  | You forgot to mention Chapter 4. Does it cover any gap? Maybe Gap 1? | Line 282: 'Chapter~\ref{ch:methodology} presents the SDT framework that addresses Gaps~1-… | 4 |
| 101 | 51 | major | cross-cutting |  | + reference to evaluation wrt the use of the agent to reconcile an ex… | Closing paragraph (line 284), the forward pointer to Chapter 8. Kritikos wants an explici… | 1 |
| 51 | 35 | minor | rewrite |  | Not clear what is meant by local testing. Is the API somehow mocked,… | On line 18, 'Consumers record interactions during local testing' is ambiguous. Clarify to… | 0 |
| 52 | 35 | minor | rewrite |  | Could also explain what are test stubs in a footnote. | Line 18 ends '...generating test stubs from contract definitions written in Groovy or YAM… | 0 |
| 53 | 35 | minor | trivial | ✓ | consumer-driven | Insertion marker: Kritikos wants 'consumer-driven' inserted before 'contracts'. Line 18 a… | 0 |
| 54 | 35 | minor | trivial | ✓ | of these properties | Wording insertion on the PBT sentence (line 25). Kritikos suggests 'generating random inp… | 0 |
| 55 | 35 | minor | structural |  | Could change paragraph here. | Kritikos suggests a paragraph break within the PBT discussion (line 25). The current para… | 0 |
| 57 | 36 | minor | rewrite |  | Is this production gated on the validation results? Could clarify in… | ANCHOR-NOT-FOUND for the exact gating sentence; the PDF context is the Spectral paragraph… | 0 |
| 58 | 36 | minor | cross-cutting | ✓ | (PR) -- could identify the acronym also in the acronym's table. | Kritikos flags an acronym (PR = pull request) that should be expanded on first use and ad… | 0 |
| 59 | 37 | minor | rewrite |  | Do you mean complex input parameters here or something else? | ANCHOR-NOT-FOUND: the PDF context is the OpenAPI-versions sentence (line 56) but Kritikos… | 0 |
| 61 | 37 | minor | structural |  | Code-First ... OpenAPI Production Workflows | Suggested heading rename for subsection 2.2.2 (line 62), currently 'Code-First and Design… | 0 |
| 62 | 37 | minor | rewrite |  | where annotations cover specific architectural components of RESTful… | Line 65, after 'frameworks generate the specification from annotated source code', Kritik… | 0 |
| 64 | 38 | minor | structural |  | Please change paragraph here. | Insert a paragraph break in subsection 2.2.2 (line 67) at the design-first transition. Cu… | 0 |
| 66 | 38 | minor | rewrite | ✓ | it does not assess the completeness or design quality of the specific… | This is suggested replacement prose for the Dredd-limitation sentence (line 78), which cu… | 0 |
| 70 | 39 | minor | trivial | ✓ | , such as Kong, ... | Duplicate of idx 69 -- Kritikos's inline insertion of ', such as Kong, ...'. The live tex… | 0 |
| 72 | 40 | minor | trivial | ✓ | container-based? | Line 99: 'Argo Workflows extends the Argo ecosystem with container-native workflow execut… | 0 |
| 73 | 40 | minor | trivial | ✓ | performance on these metrics | Line 108: 'High-performing teams achieve superior metrics through comprehensive automatio… | 0 |
| 74 | 41 | minor | trivial | ✓ | artifacts | Spelling/word insertion on the release-engineering paragraph (line 115). Kritikos flags '… | 0 |
| 79 | 43 | minor | rewrite |  | Maybe indicate that policy enforcement is analogous to quality gates?… | In the Policy-as-Code subsection (line 154, 'SDT applies the same pattern to validate API… | 0 |
| 80 | 43 | minor | rewrite |  | like code changes in SDT. Maybe this is stressed in next paragraph. | Ontological reconciliation subsection (line 161), on the 'correction mechanism' that driv… | 1 |
| 84 | 46 | minor | rewrite |  | Maybe it is more particularly suited for the latter two types of agen… | Section 2.5 intro (lines 191-199) defines three levels of AI assistance (code completion,… | 0 |
| 90 | 48 | minor | trivial | ✓ | and errors | Test-first paradigm description (line 248): 'It is labor-intensive, prone to incompletene… | 0 |
| 92 | 48 | minor | trivial | ✓ | consumer | Duplicate of idx 91 -- inline insertion of 'consumer' before 'Contract-first'. Resolved b… | 0 |
| 94 | 49 | minor | trivial | ✓ | Based on the analysis over all previous chapter sections, our ... | Line 262: 'Our survey identifies seven gaps in the current landscape that SDT and the XSD… | 0 |
| 95 | 49 | minor | trivial | ✓ | of both aspects? | Gap 1 (line 265): '...forcing teams to assemble ad-hoc toolchains with no guarantee of co… | 0 |
| 98 | 50 | minor | rewrite |  | Is this related to a specific gap? As you talk about gap dependencies… | The 'These gaps are not independent' paragraph (line 280). Kritikos notes the paragraph p… | 0 |

### 04-cli-architecture.tex (49 comments)

*Kritikos's dominant theme in Chapter 4 is a structural/conceptual one (headline #10): the chapter conflates "the DriveBy tool" with "the overall system," and he wants the system-level architecture (Kubernetes + Crossplane infrastructure with DriveBy and XSDLC on top, real external actors = API Developer/Platform Operator/GitHub) hived off into its own preceding chapter, leaving Chapter 4 to cover only DriveBy's internal architecture. The second dominant theme (headline #11) is pervasive text-vs-figure inconsistency: the dependency-flow figure says P001--P008 and omits report/util/github/logger, the component diagram says P001--P005/P008/P009, and the package table lists eleven packages and P009 — three different package/principle inventories that contradict each other and the prose. The remainder are clarification requests where the prose asserts behavior the reader cannot see in the listings/figures (registry, "combined modes," credential supply, content-level response validation, P007 criticality echoing headline #14).*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 164 | 69 | blocking | structural |  | "Tool" — As you mentioned in your email that the system includes Kube… | This is headline #10 and the dominant issue for this chapter. The chapter title is '\chap… | 10 |
| 169 | 70 | blocking | figure |  | If you attempt to cover the overall system and not the DriveBy tool,… | Redraw figures/system-context.tex (input on line 13). Kritikos's required model: an infra… | 11 |
| 177 | 72 | blocking | figure | ✓ | So, P009 was removed or not? As here we do not actually see it ... | In figures/dependency-flow.tex line 27 the principles node reads 'P001--P008 checkers, re… | 11 |
| 181 | 73 | blocking | figure | ✓ | Here now P009 appears ... In addition, in the previous image we have… | Direct contradiction between the two figures: dependency-flow.tex says 'P001--P008' (a co… | 11 |
| 170 | 70 | major | figure |  | Isn't this [GitOps Promoter] part of your system? As it enables the p… | In figures/system-context.tex, the GitOps Promoter is currently drawn as an EXTERNAL acto… | 10 |
| 171 | 70 | major | figure |  | In my view, this [OpenAPI Specification] is an artifact and not a com… | In figures/system-context.tex, the OpenAPI Specification is drawn with the same node styl… | 11 |
| 172 | 70 | major | figure |  | This [Validation Report (JSON, exit code)] is an artifact and not a c… | Same fix as idx=171: in figures/system-context.tex give the 'Validation Report (JSON, exi… | 11 |
| 173 | 70 | major | structural |  | DriveBy's internal architecture. From now on, you talk about the inte… | Kritikos marks the transition point (line 31, 'The DriveBy architecture is shaped by five… | 10 |
| 174 | 70 | major | rewrite |  | Why does this hold only for static validation? What about dynamic val… | Design goal 3 (line 38) reads 'Determinism by construction. Static validation must be fre… | 0 |
| 175 | 71 | major | rewrite |  | What about the other two goals? How do they relate to the SDT axioms? | Line 45 maps only three of the five design goals to axioms: 'Version agnosticism enables… | 0 |
| 176 | 71 | major | rewrite |  | You forgot to mention two packages here .... | Line 54 states the module 'is organized into eleven internal packages' and the verbatim c… | 11 |
| 178 | 72 | major | rewrite |  | What is registry? What does it include? | The word 'registry' appears in figures/dependency-flow.tex (line 27) and Table 4.1 (line… | 11 |
| 179 | 73 | major | figure |  | But now you introduce different layers wrt the previous figure. This… | The two figures use inconsistent layer vocabularies. dependency-flow.tex labels layers 'V… | 11 |
| 180 | 73 | major | figure |  | Can you improve the diagram? Some arrows are not fully represented (o… | Improve figures/component-diagram.tex. Verify all dependency arrows render fully (no clip… | 11 |
| 184 | 73 | major | figure |  | We do not see the last three packages in Figure 4.3. | Table 4.1 lists eleven packages including util (81 LOC), github (873 LOC), and logger (16… | 11 |
| 185 | 73 | major | figure |  | Report package is not mentioned in Figure 4.2. | The 'report' package appears in Table 4.1 (line 95) and in figures/component-diagram.tex… | 11 |
| 187 | 75 | major | rewrite |  | ok but if we have a critical failure, does it make sense to move to f… | Line 174 says TestImpact 'indicates whether functional or performance tests can proceed'… | 14 |
| 190 | 76 | major | rewrite |  | What is exactly covered by the Type? Does it relate to the kin-openap… | Line 204 mentions 'The \texttt{Type} and \texttt{RawVersion} methods allow callers to ins… | 0 |
| 191 | 76 | major | rewrite |  | ok but how does it return them? For instance, is the Schema type rela… | Line 213--235 introduce the normalized types (Table 4.2) including \texttt{Schema}, but t… | 0 |
| 192 | 78 | major | rewrite |  | ok but you need to argue if it is complete or not due to this restric… | Line 270 says Swagger 2.0 structural validation 'is implemented as a focused set of sanit… | 6 |
| 193 | 81 | major | rewrite |  | when you go from 2.X to 3.0, do you downgrade or upgrade? | This anchors to figures/class-diagram.tex (the class diagram input on line 349), where a… | 11 |
| 195 | 85 | major | rewrite |  | What is combined mode selections? I have the feeling you mean modes m… | Line 472 ends: 'they will be added to the registry and included in the test-only and comb… | 11 |
| 196 | 85 | major | rewrite |  | ok but in a previous figure we saw that the CLI could separately exec… | Line 484 claims the Engine 'is the single entry point for all validation operations', but… | 11 |
| 197 | 85 | major | rewrite |  | You say things that we do not see somewhere ... | Line 494 describes the New constructor's validation-at-construction behavior (requiring s… | 11 |
| 201 | 87 | major | rewrite |  | does it also comply with any kind of value restriction in general? | Functional-testing step 1 (line 555) generates path-parameter values from format/type def… | 0 |
| 202 | 87 | major | rewrite |  | where does the user supply credentials? It is not clear ... | Functional-testing step 3 (line 557) says auth headers are added 'when the user provides… | 0 |
| 203 | 87 | major | pushback |  | ok but it neglects the actual content which makes the functional test… | Functional-testing step 5 (line 559) validates that the status code is documented and the… | 0 |
| 205 | 88 | major | rewrite |  | which are supplied where? In the OpenAPI specification? Or in the qua… | Line 571 says results are 'compared against user-defined thresholds' for performance. Kri… | 3 |
| 163 | 69 | minor | rewrite |  | implementation tool/mechanism? To make it clear what this is about. | In the chapter-opening paragraph (line 4), the sentence 'This chapter describes how the D… | 10 |
| 165 | 69 | minor | trivial | ✓ | tool's system | Kritikos is marking another occurrence of 'system' that should read 'tool' (or be scoped… | 10 |
| 166 | 69 | minor | rewrite |  | via a context diagram | Kritikos suggests the sentence introducing Fig 4.1 should explicitly say the system is si… | 10 |
| 167 | 69 | minor | rewrite |  | takes a closer view on the boundary to ... "opening a boundary" looks… | Line 19 reads 'Figure~\ref{fig:component-diagram} then opens that boundary to expose the… | 11 |
| 168 | 70 | minor | trivial | ✓ | The system's context diagram. | Kritikos is labelling Fig 4.1 (line 15 caption) as 'the system's context diagram'. The ca… | 10 |
| 182 | 73 | minor | rewrite |  | Please explain in a footnote what is this registry. | Add a footnote to the 'principles' row of Table 4.1 (line 92, where 'registry' first appe… | 11 |
| 183 | 73 | minor | rewrite |  | not clear what is loaded, selected, evaluated and aggregated ... | Table 4.1 'engine' row (line 94) reads 'Orchestrator: load, select, evaluate, aggregate'… | 11 |
| 186 | 74 | minor | rewrite |  | You forgot to mention exit code 3. Maybe say: "while it also enables… | Line 138 explains only exit 1 vs exit 2 ('distinguish between validation failures (exit~1… | 0 |
| 188 | 75 | minor | trivial |  | In this respect, | OCR-fragment margin note: Kritikos is suggesting a connective phrase ('In this respect,')… | 0 |
| 189 | 76 | minor | rewrite |  | that was used in the tool implementation? | Line 190 first introduces kin-openapi: 'they never interact with the raw kin-openapi~\cit… | 0 |
| 194 | 85 | minor | trivial |  | The captions should be on the same page as their actual graphical ele… | Layout issue: Listing 4.8 (lst:formode, line 441) and similar listings/tables float so th… | 11 |
| 198 | 85 | minor | rewrite | ✓ | Shouldn't this be within a specific listing? | The Engine struct on lines 486--492 is in a bare \begin{verbatim} block with no caption/l… | 11 |
| 199 | 87 | minor | trivial |  | In this case. | OCR-fragment margin note suggesting a connective. Likely attaches to line 536 ('The only… | 0 |
| 200 | 87 | minor | rewrite |  | with respect to what? | Line 538 says the engine 'provides a \texttt{Validate} method that wraps \texttt{Validate… | 0 |
| 204 | 88 | minor | rewrite |  | like GET, HEAD, OPTIONS ones? | Line 571 (performance testing) says it 'constructs targets from non-destructive endpoints… | 0 |
| 206 | 88 | minor | rewrite |  | how to configure it and what does it usually include? | Line 571 mentions 'a configurable load profile' without specifying the configurable param… | 0 |
| 207 | 88 | minor | rewrite | ✓ | please provide an example in parenthesis as done for 1. | Report-generation item 1 (line 583) gives a filename example 'validation-report-20260319T… | 0 |
| 208 | 89 | minor | rewrite |  | correct. But it could also enable correcting source code, if the flaw… | This concerns the agent-feedback claim at line 597 ('an instruction set for autonomous AP… | 0 |
| 209 | 90 | minor | rewrite |  | what is Viper? | Line 631 references Viper without introduction: 'assemble a \texttt{ValidatorConfig} from… | 0 |
| 210 | 91 | minor | rewrite |  | what about github-status? | The CI/CD integration list at line 668 mentions only '--github-comment' under GitHub inte… | 11 |
| 211 | 92 | minor | trivial | ✓ | a | no action / OCR artefact — single stray character 'a' with no actionable content; verify… | 0 |

### 05-kubernetes-architecture.tex (45 comments)

*Kritikos's Chapter 5 feedback clusters around four dominant themes. First and largest: figure quality — nearly every TikZ diagram (k8s-system-context, provider-architecture, xrd-resource-hierarchy, declarative-lifecycle) has arrow-rendering defects (arrowheads without lines, unclear start/end, hidden arrows, dashed arrows ending inside rather than at rectangle boundaries) plus text/figure inconsistencies (figures show Role+RoleBinding while text says "RBAC"; figures omit the per-environment resources the text lists). Second: a structural request echoing his headline point #10 — pull the overall system architecture/context into its own chapter ahead of the CLI tool chapter. Third: correctness of the DAG step count (he counts 6, then 8 steps, against the text's "5-step/7-step" claims) and the worked-example DAG whose first two step names (set-pending, health-check) disagree with the named steps above. Fourth: many small clarifications and wording fixes (SDLC vs XSDLC acronym, "spec" vs "CR", default-value provenance, perfect-api naming, parameter-table ordering) plus a few trivial typos.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 252 | 121 | blocking | pushback |  | I have the impression that we have 6 steps in this case: 2 initial, 3… | Step-count correctness. Anchor at line 686: 'a gate with one \texttt{validate-only} check… | 13 |
| 253 | 121 | blocking | pushback |  | This should be 8. See previous comment. | Step-count correctness (companion to idx 252). Anchor at line 686: 'a gate with validate-… | 13 |
| 212 | 93 | major | structural |  | Almost as the architectural diagram (Figure 4.1) covered almost the w… | This is the chapter-level echo of headline point #10 (and the author has already agreed:… | 10 |
| 215 | 97 | major | figure |  | Correct but that figure requires some updating... | Refers to the k8s-system-context figure introduced at line 117 ('Figure~\ref{fig:k8s-syst… | 11 |
| 216 | 98 | major | figure |  | The figure is nice but still has issues, mostly w.r.t. the arrows aga… | Edit figures/k8s-system-context.tex (rendered as Figure~\ref{fig:k8s-system-context}, lin… | 11 |
| 217 | 98 | major | figure |  | What is perfect-api? Maybe remove the 'perfect' prefix to make the na… | The figure k8s-system-context.tex labels namespaces perfect-api-dev / perfect-api-staging… | 11 |
| 218 | 98 | major | figure |  | Is the promotion done directly from one environment to the other, or… | In figures/k8s-system-context.tex the 'promote' arrows between environment namespaces mus… | 11 |
| 222 | 102 | major | figure |  | Figure is almost ok. Again there is the issue with the arrows, this t… | Edit figures/provider-architecture.tex (Figure~\ref{fig:provider-architecture}, lines 183… | 11 |
| 223 | 103 | major | rewrite |  | SDLC or XSDLC? If the former holds, the SDLC acronym was never introd… | Anchor at line 204: '...the composition generates all resources needed for the full softw… | 0 |
| 230 | 105 | major | rewrite |  | Need to explain the semantics of the different parameters that can be… | Anchor at line 253 (same load-test bullet as idx 228). Kritikos wants each load-test para… | 0 |
| 231 | 106 | major | rewrite |  | Could explain why validate-only is executed twice in the two gates. I… | Anchor at the Listing lst:xsdlc-cr (lines 287-310), where both the staging gate (line 302… | 0 |
| 232 | 106 | major | table-fix |  | OK, but you also have environment resources that are not mentioned he… | Anchor at line 312: 'From this declaration, Crossplane generates the shared resources (Se… | 11 |
| 233 | 106 | major | rewrite | ✓ | In Figure 5.3 we see Role & RoleBinding while here you mention RBAC. | Text/figure inconsistency (headline #11). Anchor at line 312: 'the per-gate resources (RB… | 11 |
| 234 | 107 | major | figure |  | Again please try to fix the arrows to appear completely. | Edit figures/xrd-resource-hierarchy.tex (Figure~\ref{fig:xrd-resource-hierarchy}, lines 3… | 11 |
| 235 | 107 | major | figure |  | Do we also have environment resources that are created? We do not see… | Edit figures/xrd-resource-hierarchy.tex. The figure shows app-level shared and per-gate b… | 11 |
| 238 | 110 | major | rewrite |  | Do you mean CR? 'spec' gives the impression you change the meta-model… | Anchor at line 381: 'Modifying the \texttt{XSDLC} spec triggers a reconciliation delta.'… | 0 |
| 241 | 111 | major | figure |  | Some comments for improving the figure: (a) dashed arrows should not… | Edit figures/declarative-lifecycle.tex (Figure~\ref{fig:declarative-lifecycle}, lines 412… | 11 |
| 247 | 119 | major | rewrite | ✓ | The first two steps have different names than those declared above. | Inconsistency between the worked-example DAG and the named pipeline steps. Anchor at the… | 13 |
| 249 | 119 | major | pushback |  | I am puzzled slightly here, as I had the impression that in a previou… | Genuine contradiction to resolve. Anchor at line 659: 'Helm layer (evaluated at \texttt{h… | 0 |
| 250 | 120 | major | structural |  | As this was already explained in Section 5.6.2, you could make a simp… | De-duplicate the template-escaping explanation. Anchor at line 665: 'The composition solv… | 5 |
| 254 | 122 | major | rewrite |  | The correct or effective sequence? The chart should impose a specific… | Anchor at lines 711-718, the four-step install order ('Installation follows a four-step o… | 0 |
| 255 | 122 | major | rewrite |  | Can you provide a script or specific commands per phase? This would b… | Anchor at the install-order enumerate, lines 713-718 (item 1 'Pre-install: Crossplane pro… | 0 |
| 213 | 95 | minor | trivial |  | here | Vague one-word marker placed on the Stage-5 paragraph (line 48: 'The final stage---the \t… | 0 |
| 214 | 96 | minor | trivial | ✓ | , respectively | Wording insertion. Anchor at line 89: 'The same principle checkers that run in a develope… | 0 |
| 219 | 99 | minor | trivial | ✓ | actions | Wording. Anchor at line 140: 'In the act phase, the controller takes corrective action to… | 0 |
| 220 | 100 | minor | trivial | ✓ | Thus, | Connective insertion. Anchor at line 151, the sentence opening the self-healing logic: 'T… | 0 |
| 221 | 100 | minor | trivial | ✓ | concretely instantiated | Wording suggestion. Anchor at line 155: 'The three components of ontological state reconc… | 0 |
| 224 | 103 | minor | rewrite |  | The actual repository where the software project's main source code i… | Anchor at line 204: 'The software repository is never referenced---XSDLC manages only the… | 0 |
| 225 | 103 | minor | trivial | ✓ | (CR) | Acronym-definition request. Anchor at line 204: 'A single \texttt{XSDLC} custom resource… | 0 |
| 226 | 103 | minor | rewrite |  | Could explain what this is in a footnote. | Anchor at line 215 (list item 7): 'An \textbf{ArgoCD push secret} granting the Source Hyd… | 0 |
| 227 | 104 | minor | rewrite |  | Maybe say that it exploits the static validation facilities of the Dr… | Anchor at line 251: 'validate-only --- runs the DriveBy CLI in static validation mode, ev… | 8 |
| 228 | 105 | minor | rewrite |  | Could explain whether the default values are derived from common prac… | Anchor at line 253: 'load-test --- ... configurable parameters: \texttt{concurrentUsers}… | 0 |
| 229 | 105 | minor | rewrite |  | I presume this is an indicative distribution of checks to potential g… | Anchor at line 258: 'Different environment transitions compose different subsets. For the… | 0 |
| 236 | 108 | minor | rewrite |  | In the context of your system, right? This needs to be clarified. In… | Anchor at line 339: 'The Helm chart installs the \textit{platform layer}: the XRD definit… | 0 |
| 237 | 109 | minor | rewrite |  | composition? XSDLC composite? | Anchor at the caption of Figure~\ref{fig:crossplane-reconciliation-loop}, line 359: 'The… | 0 |
| 239 | 110 | minor | rewrite |  | + RepositoryFile based on 5.5.3 (last paragraph). | Anchor at line 381: 'Adding a new environment generates the new ArgoCD Application, branc… | 0 |
| 240 | 110 | minor | rewrite | ✓ | Across all gated environments? | Anchor at line 396 (status-fields list): '\texttt{gates[]} (array): per-gate status inclu… | 0 |
| 242 | 113 | minor | rewrite |  | Could explain why, although this is obvious. | Anchor at line 462: 'They use \texttt{deletionPolicy: Orphan} to preserve Git history whe… | 0 |
| 243 | 116 | minor | table-fix |  | I would propose ordering the parameters based on whether they are glo… | Reorder Table~\ref{tab:config-tiers} (the two-tier config table, lines 513-531). Kritikos… | 13 |
| 244 | 116 | minor | trivial | ✓ | support | Wording. Anchor at line 545: 'It is backed by NATS JetStream~\cite{nats-jetstream}, which… | 0 |
| 245 | 117 | minor | rewrite |  | What do you mean here? | Anchor in Table~\ref{tab:sensor-params} (lines 582-598), the rows where the 'Purpose' col… | 0 |
| 246 | 119 | minor | rewrite |  | Is authentication also required here? | Anchor at line 632: 'load-test $\rightarrow$ \texttt{driveby-loadtest}: Runs the DriveBy… | 0 |
| 248 | 119 | minor | rewrite |  | Could explain that -> means sequential execution and / means parallel… | Anchor at the verbatim DAG block lines 645-649 (and the analogous one at line 643 'produc… | 0 |
| 251 | 121 | minor | rewrite |  | Is this a default or an effective value taken from a specific configu… | Note: the JSON section label 'Helm Chart Packaging' is mis-attributed; the context text (… | 0 |
| 256 | 123 | minor | trivial | ✓ | a... a Go CLI, a..., and a... to a... (insert articles consistently) | Article/parallelism fix. Anchor at line 736: 'tracing the evolution from shell script thr… | 0 |

### 09-discussion.tex (35 comments)

*Kritikos's discussion-chapter comments cluster around three recurring demands. First, terminological precision: "methodology" should be "framework" throughout (RQ1 title and question wording), the RQ questions must match Chapter 1 Section 1.5 verbatim, and several claims overstate what was achieved (100% vs 80% detection, "fully" automated gates that actually need the XSDLC CR). Second, the central technical concern that runs through Threats-to-Validity and RQ answers: functional testing only checks structural validity of responses, not content/business-logic, and non-functional/performance testing is narrow — both create an upward bias in detection rates that the author must concede, plus missing statistical/conclusion validity and IDP-generalisability threats. Third, a conceptual prompt the author should engage with: whether non-functional (performance) constraints belong in the OpenAPI contract (via extension) rather than only in the XSDLC CR, which bears on the single-vs-dual-artifact claim (headline point #3).*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 415 | 195 | blocking | table-fix |  | 100% or 80%? | Line 18 claims '100% detection accuracy ... across all seven implemented principle checke… | 13 |
| 419 | 196 | blocking | table-fix |  | Are you sure about this? I had the impression that the first has othe… | Refers to line 24: 'bad-docs-api and no-auth-api, both lacking security definitions'. Kri… | 13 |
| 424 | 196 | blocking | pushback |  | I believe that quality gates are not produced fully solely from OpenA… | Line 33 answers RQ2 with 'Fully, for the implemented principles. All seven principle chec… | 3 |
| 431 | 199 | blocking | table-fix |  | Are you sure about this? This changed based on the new evaluation. | Line 74 lists the universal-fail set as 'P002, P003, P004, and P008' and line 76 names 'P… | 13 |
| 434 | 199 | blocking | pushback |  | Correct, but the question is whether it should also carry non-functio… | At the dual-specification paragraph (line 87) engage Kritikos's substantive proposal: tha… | 3 |
| 438 | 201 | blocking | rewrite |  | I would also add that functional testing relies on structural validat… | Add a new internal-validity threat (extend line 108 paragraph or add a sentence): P006 fu… | 13 |
| 441 | 201 | blocking | structural |  | You are lacking statistical validity. | The Threats section (lines 104-114) has Internal, External, Construct, and Agent-driven s… | 0 |
| 443 | 201 | blocking | rewrite |  | I am also concerned about the way functional testing is performed and… | Closely related to idx 438. In the threats section, add an explicit limitation that P006… | 13 |
| 445 | 201 | blocking | rewrite |  | Non-functional testing is also very restrained, covering only particu… | Add a threat (Internal or Construct validity) acknowledging that P007 non-functional test… | 14 |
| 414 | 195 | major | rewrite |  | I believe that again we have the issue that the question content is n… | Open Chapter 1 Section 1.5 (01-introduction.tex) and copy the RQ1 question verbatim. The… | 0 |
| 417 | 195 | major | rewrite |  | Maybe the name/title is misleading here as you are proposing a framew… | Rename the subsection heading on line 12 from 'RQ1: Methodology Soundness' to 'RQ1: Frame… | 0 |
| 420 | 196 | major | rewrite |  | They did not have any kind of warning or issue in terms of static val… | Refers to line 24 describing slow-api and broken-api passing staging. Clarify the mechani… | 13 |
| 421 | 196 | major | rewrite |  | You have to better clarify the exact layers and their exact hierarchy… | In the RQ1 layered-discrimination sentence (line 24, ending 'each validation layer catche… | 0 |
| 425 | 196 | major | pushback |  | In what forms could the business logic actually be described? Someone… | Line 35 claims SDT 'does not verify that the API's business logic produces correct result… | 0 |
| 426 | 197 | major | rewrite |  | Is it Chapter 8 or Chapter 7? I do not get the point about extending… | Line 37: 'The agent-driven evaluation (Chapter~\ref{ch:ai-development}) extended this fin… | 1 |
| 433 | 199 | major | rewrite |  | OK but mainly functional. Does not cover performance-oriented behavio… | Line 87: 'Where the OpenAPI document specifies API behaviour, the XSDLC document specifie… | 3 |
| 435 | 200 | major | rewrite |  | OK, but in one contribution we have OpenAPI specification validation;… | Line 91 frames the two contributions as 'complementary'. Kritikos asks what the SECOND (X… | 2 |
| 437 | 201 | major | pushback |  | In principle, someone could inspect existing OpenAPIs to collect thes… | Internal-validity paragraph (line 108) already concedes 'The defect set was hand-designed… | 13 |
| 439 | 201 | major | rewrite |  | A random selection strategy would have been better, which guarantees… | External-validity paragraph (line 110) admits 'The sample is alphabetically biased toward… | 0 |
| 440 | 201 | major | pushback |  | I am not sure about this: if you take into account a specific non-neg… | Line 110 argues the AWS-style over-representation 'likely STRENGTHENS rather than weakens… | 0 |
| 442 | 201 | major | rewrite |  | ...and did not cover other IDPs (i.e., IDPs deployed in other organis… | Line 110 ends 'The operational evaluation was conducted on a single IDP (Novelcore's Kube… | 0 |
| 448 | 203 | major | rewrite |  | This requires more detail, as a specific evaluation cannot be regarde… | Summary line 136 lists 'the single-API controlled evaluation, the public-API bias in the… | 0 |
| 416 | 195 | minor | rewrite |  | Could also say something about the extension of the controlled experi… | In the RQ1 paragraph (after line 18) add one sentence noting that the controlled experime… | 0 |
| 418 | 196 | minor | pushback |  | OK but the 'beyond theoretical scaffolding' was eventually covered on… | The paragraph at line 22 actually does address all three axioms (Completeness, Determinis… | 0 |
| 422 | 196 | minor | rewrite |  | With respect to what? Or do you mean that they are complementary to e… | Line 24 ends 'demonstrates that the nine-principle structure is not redundant but complem… | 0 |
| 423 | 196 | minor | trivial | ✓ | an OpenAPI specification | Wording nudge on line 31/33: the RQ2 question says 'derived automatically from OpenAPI sp… | 0 |
| 427 | 197 | minor | rewrite |  | OK but maybe this is fully addressed? In contrast to the previous one? | Refers to the RQ3 answer beginning line 46 ('demonstrated effective integration'). Kritik… | 0 |
| 428 | 197 | minor | rewrite |  | Could also refer here to the CI/CD metrics and the very good values a… | In the RQ3 answer (lines 46-48) add a sentence citing the concrete CI/CD-integration metr… | 0 |
| 429 | 198 | minor | rewrite |  | Here we talk about another control loop or the existing one? The focu… | Refers to the RQ4 answer (line 63, 'The XSDLC CRD functions as a machine-readable, struct… | 3 |
| 430 | 198 | minor | structural |  | You have 4 paragraphs here. Could map them to Section 9.2.1 'OpenAPI… | Wrap the four paragraphs at lines 74-80 (the APIs.guru pattern, the two implications, the… | 0 |
| 432 | 199 | minor | rewrite |  | Correct + seems that structural rules are not sufficient for full str… | On line 76 (Spectral 'enforce structural rules ... but do not assess whether the specific… | 0 |
| 436 | 200 | minor | rewrite |  | On the other hand, this means that... | Anchor is the contrast Kritikos wants drawn at line 89 ('An XSDLC resource, in contrast,… | 0 |
| 444 | 201 | minor | structural |  | Is this a kind of internal-validity issue? If yes, then why is it men… | Refers to the 'Agent-driven evaluation' subheading (line 114), which Kritikos thinks is i… | 1 |
| 446 | 202 | minor | rewrite |  | Plus reduce specification drift? As this is also covered, although no… | In the 'For API teams' implication (line 121), after 'enabling targeted remediation rathe… | 0 |
| 447 | 202 | minor | rewrite |  | And potentially incomplete results. | Same 'For API teams' paragraph (line 121) describing minimal mode (P001 only) as a starti… | 0 |

### 01-introduction.tex (33 comments)

*Kritikos's introduction comments cluster around three dominant themes. First and most consequential: XSDLC and the second user-supplied artifact (the CR expressing non-functional requirements) is under-represented in the opening framing, the contributions, and the RQ-to-contribution mapping table — he wants the quality-gate / XSDLC contribution and the reconciliation-agent evaluation surfaced consistently (his headline points #1, #2). Second: the RQ-contribution mapping table needs richer/corrected check-and-tilde marks (RQ2, RQ3, RQ4 rows). Third: a layer of small copy-edits (typos "an/a/analyses/details/limitations", missing commas), terminology precision (acronyms reused, CRD/CR expansions, "technical sense"), and one structural request to drop the "Technology Showcase" subsection from this position. No live stale bracketed reference numbers remain in this chapter (the author already converted them to \cite{}), so the reference-bug category does not apply here.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 16 | 22 | major | rewrite |  | Maybe say in parenthesis or footnote (based also on some extra user-s… | In the opening paragraph (line 4) or the SDT-introduction paragraph (line 6), add a brief… | 3 |
| 17 | 22 | major | rewrite |  | XSDLC is not mentioned at all as another contribution that focuses on… | The opening paragraph (line 4) and the SDT-overview paragraph (line 6) name only SDT/Driv… | 2 |
| 22 | 24 | major | rewrite |  | Maybe provide here a reference that validates this claim. | On line 43-47, the claim that GitOps/declarative operations have become dominant but API… | 0 |
| 23 | 25 | major | structural |  | I believe that this section is not needed in this place. We still exp… | The '\subsection{A Technology Showcase}' (line 58-61) interrupts the problem-to-contribut… | 0 |
| 24 | 26 | major | rewrite |  | Could indicate here the requirement for a methodology like SDT, such… | In the agentic-development passage (line 66-68, where 'SDT produces exactly this kind of… | 0 |
| 29 | 29 | major | pushback |  | Here with quality gate you mean just the conditions and not the infra… | RQ2 (line 133) currently asks whether gates 'can be derived automatically from OpenAPI sp… | 3 |
| 30 | 30 | major | rewrite |  | for what? Better to clearly indicate what is this system about. | Line 151 'This thesis presents a proof-of-concept system and makes six contributions.' is… | 0 |
| 31 | 30 | major | rewrite |  | (Custom Resource Definition) — Need to also explain in footnote what… | First use of 'CRD' is in RQ4 (line 143, 'the XSDLC CRD functions as...'). Expand the acro… | 0 |
| 35 | 31 | major | rewrite |  | Forgot also the last evaluation concerning the use of an agent to imp… | Contribution C5 (line 162) lists three experiments (controlled, large-scale, operational)… | 1 |
| 38 | 32 | major | table-fix |  | Maybe here we have ~ as the primary evaluation is on Chapter 7 and th… | In Table~\ref{tab:rq-contribution-mapping}, the RQ2 row (line 182) currently marks C6 wit… | 2 |
| 39 | 32 | major | table-fix |  | Maybe also ~ here as SDT touches somehow RQ3. | In the RQ3 row (line 183), Kritikos wants a tilde acknowledging that C1 (the SDT conceptu… | 2 |
| 40 | 32 | major | table-fix |  | I would also add a tick or tilde here as RQ4 is validated in C5. Chec… | In the RQ4 row (line 184), the C5 (empirical evaluation) cell is currently empty. The liv… | 2 |
| 45 | 33 | major | rewrite |  | Missing the last evaluation again here ... | Line 209 (Chapter 7 structure summary) lists controlled + large-scale + operational evalu… | 1 |
| 46 | 33 | major | rewrite |  | Also stress that the author's contribution degree in the thesis is de… | On line 211 (Chapter 8 summary), add that Chapter 8 also quantifies the AUTHOR's own cont… | 15 |
| 18 | 23 | minor | trivial | ✓ | I propose to remove "in Accelerate". | On line 20, delete the italic book title so it reads 'Forsgren, Humble, and Kim demonstra… | 0 |
| 19 | 23 | minor | trivial | ✓ | an (article correction) | Article fix near the documentation-drift definition (line 27). Kritikos marks an 'a' that… | 0 |
| 20 | 24 | minor | rewrite |  | So, what (transition / rhetorical-question concern) | Comment lands at the end of the Documentation-Drift subsection. The live text ends the su… | 0 |
| 21 | 24 | minor | trivial | ✓ | development / project teams? | On the GitOps paragraph (line 43) Kritikos queries an unqualified use of 'teams'. The liv… | 0 |
| 25 | 26 | minor | trivial | ✓ | a (article correction) | Single-letter article fix in the tribal-knowledge opening paragraph (line 80). Verify the… | 0 |
| 26 | 27 | minor | rewrite |  | through what? (externalization mechanism unclear) | On line 89, the sentence 'The central challenge of knowledge management is externalizatio… | 0 |
| 27 | 28 | minor | rewrite |  | or technical sense but ... (clarify 'aspirational' vs technical achie… | On line 104, 'This assumption is not aspirational---it is increasingly achievable.' Kriti… | 0 |
| 28 | 29 | minor | rewrite |  | Could also refer to EvoMaster here. | In the existing-tools enumeration (line 117, 'Spectral lints... Schemathesis generates...… | 4 |
| 32 | 31 | minor | trivial | ✓ | validation (DriveBy implements SDT validation, wording) | On line 156 (contribution C2), Kritikos annotates 'implements SDT for OpenAPI specificati… | 0 |
| 33 | 31 | minor | rewrite | ✓ | (Custom Resource) — expand CR on first use. | On line 158 (contribution C3), the phrase 'from a single CR' uses the acronym CR. Expand… | 0 |
| 34 | 31 | minor | trivial | ✓ | IDP -> could use the acronym as it has been already introduced. Can u… | On line 160 (C4) the phrase 'within an internal developer platform' is spelled out, but I… | 0 |
| 36 | 32 | minor | rewrite |  | Not clear what is Phase B. | On line 164 (C6), 'the 12-day Phase~B agent-partnership window' uses 'Phase B' without de… | 0 |
| 37 | 32 | minor | rewrite |  | You mean it could be directly consumable by AI agents? As the methodo… | On line 164, the clause 'demonstrates that SDT's structured output is directly consumable… | 0 |
| 41 | 33 | minor | trivial | ✓ | analyses (wording correction) | On line 203 (Chapter 4 summary) Kritikos marks the word 'describes' suggesting a stronger… | 0 |
| 42 | 33 | minor | trivial | ✓ | details (wording correction) | On line 207 (Chapter 6 summary) Kritikos suggests replacing the verb with 'details'. Chan… | 0 |
| 43 | 33 | minor | trivial | ✓ | a (article correction) | On line 209 (Chapter 7 summary) a single-letter article fix is marked. Likely 'controlled… | 0 |
| 44 | 33 | minor | trivial | ✓ | a (article correction) | Second single-letter mark on the same Chapter 7 sentence (line 209), e.g. 'a proof-of-con… | 0 |
| 47 | 33 | minor | trivial | ✓ | the current limitations (wording) | On line 215 (Chapter 10 summary), change 'limitations' to 'the current limitations' as Kr… | 0 |
| 48 | 33 | minor | trivial | ✓ | , (missing comma) | On line 215, add a missing comma in the future-work list. The phrase 'future work includi… | 0 |

### 08-ai-assisted-development.tex (33 comments)

*Chapter 8 is Kritikos's authorship-and-process audit, and his comments cluster around three demands: (1) prove authorship of the REPORT, not just the code — he wants Table 8.2 (effort allocation) split so the report deliverable has its own measured proofreading/writing/correction percentages, and he is uneasy that AI revised chapter content which the author then merely proofread (headline #15). (2) Strip stale round-N-review provenance — the chapter repeatedly says "the supervisor noted~([534/535/544/551/549/542/550/528])" and flags non-existent reference numbers; per headline #5 these must become self-contained prose. (3) A run of substantive process questions (IDE code-assist in "solo" Phase A, why lines were removed, how debugging was done, who implemented the fault injections, whether the derived OpenAPI spec is safe for production) plus a dense cluster of "follow-on vs follow-up" wording flags in Personal Lessons Learned.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 381 | 179 | blocking | pushback |  | Just one quick question at this point: did you use any code assistant… | In subsec:two-phases, the Phase A description ('During this phase the author worked alone… | 15 |
| 384 | 181 | blocking | table-fix |  | You do not provide any relevant percentage concerning thesis checking… | This is headline point #15 and the single most important comment in the chapter. Table ta… | 15 |
| 387 | 182 | blocking | pushback |  | So, the revisions were done by you? I am asking this as it seems that… | This is the academic-acceptability crux (echoed at idx 397 and idx 398). The sentence at… | 15 |
| 390 | 183 | blocking | reference-bug | ✓ | again this is a non-existing reference number | Per headline point #5: the chapter is littered with bracketed numbers that are stale prov… | 5 |
| 397 | 189 | blocking | pushback |  | ok but ideally I would expect to see you conduct the writing of the t… | Core defense risk (with idx 387, 398). The division-of-labour itemize (lines 286-290) cur… | 15 |
| 398 | 190 | blocking | structural |  | Ok, this is the most crucial point here. This makes it acceptable in… | Kritikos is pointing at the protocol step 4 (line 298): 'the author reads the draft, edit… | 15 |
| 400 | 192 | blocking | rewrite |  | ok but you do not explain what was the result and whether it is posit… | Line 340: 'The result of this experiment converts the theoretical connection of this sect… | 1 |
| 401 | 192 | blocking | rewrite |  | Sure, but in principle the reconciliation might not reflect reality e… | Strong technical objection to the line 349 example ('add maxLength and pattern properties… | 1 |
| 382 | 179 | major | pushback |  | + principle checking finalised and the DriveBy tool was fully produce… | This continues 381: Kritikos is asking whether the DriveBy tool / principle checking was… | 15 |
| 383 | 180 | major | rewrite |  | Why was there a need to remove lines? | Table tab:authorship-stats (lines 40-49) reports 'Lines removed 21,330 / 7,678'. Kritikos… | 15 |
| 386 | 182 | major | rewrite |  | Could you explain in the text somewhere how debugging was actually ca… | The table row 'Manual code changes, debugging, hot-fixes (author) 15%' (line 79) is asser… | 15 |
| 391 | 183 | major | rewrite |  | Although the supervisor..., the... You need to indicate what was sugg… | This is the 'make it self-contained' instruction applied to the 'Containment vs. repetiti… | 5 |
| 394 | 187 | major | rewrite |  | ok but there were multiple evaluations, not just one. Were all of the… | Line 252 only credits the author with designing 'The large-scale evaluation against APIs.… | 1 |
| 395 | 188 | major | rewrite |  | Why is there a need to validate reachability of the endpoints? Isn't… | Line 259 mentions the prober 'validates reachability by testing simple GET endpoints' but… | 0 |
| 396 | 189 | major | rewrite |  | ok but it is not clear who did the injections, i.e., the implementati… | The eval-coverage-limit paragraph (line 277) says 'the new run injects a defect at every… | 13 |
| 399 | 191 | major | rewrite |  | You need to be precise here. It should be one API based also on the c… | Line 340 says 'one or two real-world public APIs with low P002--P004 scores are taken as… | 1 |
| 402 | 192 | major | rewrite |  | Could refer to specific fix directions that were given in the context… | In subsec:validation-remediation (lines 347-351), the SuggestedFix discussion is abstract… | 1 |
| 404 | 192 | major | pushback |  | ...caters for conforming to this requirement by separating these two… | Lines 358-370 promise to separate 'lessons about the framework (generalise)' from 'lesson… | 5 |
| 405 | 193 | major | pushback |  | Correct. But couldn't the AI supply some nice ideas that could be imp… | Line 360 (and the chapter-wide claim, e.g. line 108 'No deviation in the project was driv… | 15 |
| 413 | 193 | major | rewrite |  | Did you develop other skills and knowledge that could be exploited in… | All five current lessons (lines 360-368) are about agentic development specifically. Krit… | 15 |
| 385 | 181 | minor | trivial | ✓ | the thesis report structure | On line 97 the parenthetical lists 'the chapter structure' as one of the owned items. Kri… | 15 |
| 388 | 182 | minor | trivial | ✓ | mapping to | Wording fix on line 97: Kritikos flags 'maps to a specific code element' style phrasing o… | 0 |
| 389 | 183 | minor | trivial |  | of what? | Line 146: 'the author's most reusable methodological contribution from the agentic-develo… | 0 |
| 392 | 186 | minor | trivial |  | flagging as stubs chapters that are under 20 lines? Was that what you… | Line 221 lists the check-docs.sh categories including 'thesis chapter status (flagging st… | 0 |
| 393 | 187 | minor | trivial | ✓ | writes | On line 245: 'The agent does not write to memory on its own initiative'. Kritikos's margi… | 0 |
| 403 | 192 | minor | trivial |  | to be... | OCR-fragmentary marginal note ('to be...') anchored to the Personal Lessons intro (line 3… | 5 |
| 406 | 193 | minor | trivial | ✓ | follow-on vs follow-up... | Terminology consistency: Kritikos repeatedly flags (idx 406-412) that the chapter uses 'f… | 0 |
| 407 | 193 | minor | trivial | ✓ | follow-up? | Duplicate of idx 406 on a second 'follow-on' occurrence in the Personal Lessons section.… | 0 |
| 408 | 193 | minor | trivial | ✓ | follow-up | Duplicate of idx 406 — another 'follow-on' to change to 'follow-up'. Corresponds to the '… | 0 |
| 409 | 193 | minor | trivial | ✓ | ? | Bare '?' marginal query in the Personal Lessons section, page 193. Most likely another 'f… | 0 |
| 410 | 193 | minor | trivial |  | sufficient | Single-word marginal 'sufficient' on page 193. Likely a suggested word replacement (inser… | 0 |
| 411 | 193 | minor | trivial | ✓ | follow-up | Duplicate 'follow-on' -> 'follow-up' flag, page 193. Corresponds to 'Three specific chang… | 0 |
| 412 | 193 | minor | trivial |  | follow-up project are planned | Kritikos is suggesting a rephrasing involving 'follow-up project are planned' — likely pr… | 0 |

### 10-conclusion.tex (32 comments)

*Kritikos's comments on the Conclusion cluster around three dominant themes: (1) the chapter forgets the newly-added evaluations — the single-API agent-feedback experiment and the reconciliation-agent study — in the RQ answers and future work (headline #1); (2) several overclaims need softening or precise grounding — the "100%" detection claim contradicts the actual large-scale numbers, the "proof" of RQ3 needs a section number, non-functional gating is wrongly attributed to the OpenAPI spec alone when it actually relies on the XSDLC CR too (headline #3), and P006/P007 "wrapped into strict mode" appears to collapse the static/dynamic layered-validation distinction; and (3) stale non-existent bracketed reference numbers ([160], [213], [484], [501], [502], [503]) and round-N provenance phrasing must be removed and made self-contained (headline #5). Secondary asks: introduce Limitations and Future Work with a paragraph tracing where each item came from, add security as a recurring dimension, and reconcile the P007-critical position (headline #14).*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 453 | 205 | blocking | rewrite |  | Where is this proof stated? Please indicate section number. | RQ3 answer (line 38) asserts the XSDLC generates ~45-50 objects 'in under 15 seconds' and… | 0 |
| 454 | 206 | blocking | pushback |  | No, it was less than that ... | RQ1 answer (line 30) claims '100% detection accuracy for injected defects with zero false… | 13 |
| 460 | 208 | blocking | pushback |  | Correct - although I foresee a specific issue: if all principles are… | The 'Principle coverage' limitation (line 73) says wrapping P006/P007 as PrincipleChecker… | 8 |
| 473 | 210 | blocking | pushback |  | This is what was indicated before in a previous comment. Do we desire… | The 'P006 and P007 as PrincipleCheckers' item (line 96) states this 'would enable a singl… | 8 |
| 476 | 210 | blocking | reference-bug | ✓ | Again, non-existing reference number. | The Doctoral Continuation subsection carries stale bracketed comment-IDs that are NOT bib… | 5 |
| 452 | 205 | major | rewrite |  | You forgot to mention the single-API experiment and the new evaluatio… | In the Summary of Contributions, the 'Empirical evaluation' bullet (line 19) describes a… | 1 |
| 455 | 206 | major | pushback |  | Again, non-functional testing does not rely on the OpenAPI specificat… | RQ2 answer (line 34) states 'All seven principle checkers derive their validation logic e… | 3 |
| 457 | 207 | major | rewrite | ✓ | potential ... as they have not been implemented yet. | RQ4 answer (line 42) claims 'a team onboarding controller can create XSDLC resources ...… | 0 |
| 459 | 208 | major | rewrite |  | Still this covers the documentation of versioning. It does not evalua… | The 'Keyword-based checks' limitation (line 75) describes P008 keyword matching as brittl… | 0 |
| 461 | 208 | major | structural |  | Maybe say in an introductory paragraph how these limitations were dis… | Add an introductory paragraph immediately after \section{Limitations} (before line 71) ex… | 0 |
| 462 | 209 | major | rewrite |  | Further, you only cover load testing and not other kinds of non-funct… | In the 'Load testing scope' limitation (line 79), add that load testing is the only non-f… | 0 |
| 464 | 209 | major | rewrite |  | Could also add lack of reconciliation to improve OpenAPI specificatio… | Add a Limitations bullet (around line 71-83) stating that the system does not currently r… | 1 |
| 465 | 209 | major | structural |  | Maybe indicate in a small paragraph where these directions come from… | Add an introductory paragraph after \section{Future Work} (before line 90) that explicitl… | 0 |
| 468 | 209 | major | rewrite |  | Partially ... in the sense that it improves the XSDLC CR but not the… | The adaptive-threshold item (line 90) claims it 'would close the feedback loop between ob… | 1 |
| 470 | 209 | major | rewrite |  | Should we also fix code or only the OpenAPI specification (and maybe… | In the 'Agent-driven remediation' item (line 94) the agent 'generate[s] specification fix… | 1 |
| 474 | 210 | major | rewrite |  | I agree, but there is a major issue: completeness. GraphQL introspect… | The 'Multi-format specification support' item (line 100) assumes GraphQL/gRPC/AsyncAPI ca… | 0 |
| 475 | 210 | major | reference-bug |  | Could provide a reference here that supports/backs up this claim. | The claim at line 94 that an agent 'can generate the appropriate maxLength and pattern co… | 0 |
| 477 | 211 | major | rewrite |  | and security | In the doctoral 'Coverage extension across testing types' item (line 115), Kritikos wants… | 0 |
| 478 | 211 | major | rewrite |  | More importantly, the contract would then cover both the ideal functi… | In the 'Ontological extension to performance metrics' item (line 119), add Kritikos's key… | 3 |
| 479 | 211 | major | pushback |  | Isn't that already supported? | The 'Conditional gate transitions' item (line 121) frames conditional gates ('promote to… | 0 |
| 449 | 204 | minor | trivial | ✓ | consumer/customer | Kritikos suggests the word 'consumer' should read 'customer' (or vice versa). The Summary… | 0 |
| 450 | 204 | minor | trivial | ✓ | OpenAPI | Kritikos flagged a spelling/capitalisation of 'OpenAPI' (likely written 'Open API' or 'op… | 0 |
| 451 | 204 | minor | rewrite |  | in suitable cases | Kritikos wants a hedge 'in suitable cases' inserted so an absolute claim is softened. The… | 0 |
| 456 | 206 | minor | rewrite |  | Could also refer to other CI/CD integration metrics. In addition, you… | RQ3 answer (line 38) ends on the proof-of-concept feasibility statement. Add a sentence c… | 0 |
| 458 | 207 | minor | rewrite |  | How does this correlate with security? Because we have authentication… | On the bullet at line 57 ('a security policy controller could enforce that all XSDLC reso… | 0 |
| 463 | 209 | minor | rewrite |  | Further, ... + multi-cluster evaluations ... | Extend the 'Evaluation scale' limitation (line 81) to also list multi-cluster evaluation… | 0 |
| 466 | 209 | minor | trivial | ✓ | We envision a ... | Kritikos suggests opening the 'Adaptive load testing thresholds' item (line 90) with a 'W… | 0 |
| 467 | 209 | minor | rewrite | ✓ | Maybe add 'constantly' or 'very often'? So that the improvement is ma… | In the adaptive-threshold item (line 90), the trigger 'When observed concurrent users exc… | 0 |
| 469 | 209 | minor | rewrite |  | Less critical than ...? I acknowledge that this can be distinguishabl… | In the 'Risk-based scoring model' item (line 92), the example 'a missing security scheme… | 0 |
| 471 | 209 | minor | rewrite |  | This was also proven in the evaluation ... | The 'Agent-driven remediation' item (line 94) presents agent remediation as purely specul… | 1 |
| 472 | 210 | minor | rewrite |  | OK but maybe this is accomplished in a multi-round setting in a compl… | In the 'Agent-driven remediation' item (line 94), the example agent fix for a single P004… | 1 |
| 480 | 211 | minor | rewrite | ✓ | single constraint or logical combinations of constraints (i.e., compo… | In the 'Conditional gate transitions' item (line 121), where it calls for 'a formal langu… | 0 |

### 06-gitops-pipeline.tex (27 comments)

*Kritikos's Chapter 6 comments cluster around three themes. First and dominant: the control-theory framing of the quality gate (Section 6.6) — he repeatedly insists the system is closed-loop WITHOUT effective actuation (blocking is preventive, not corrective; convergence relies on manual/agent reconciliation), and wants Table 6.1 and the matching figure re-mapped so the plant is the promotion/delivery system rather than the API. The live .tex has already been heavily rewritten with a two-tier (verdict/platform) model that addresses much of this, so most of these become verify-the-rewrite-covers-it items. Second: figure quality — cut vertical labels, reversed arrows, and invisible/too-thin arrows across the pipeline, branch-topology, human-in-the-loop, and control-loop TikZ figures. Third: cross-referencing and exposition gaps — a stale "Chapter 7?" cross-ref, requests for concrete change examples, clarifying the multi-environment promotion chain (x.next->x vs dev->staging->prod), naming the SDT gate "Static Validation Gate", and justifying the doc-check ordering in CI.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 266 | 133 | blocking | rewrite |  | Previously you indicate that we go from x.next to x, while here we ha… | Add a paragraph (best placed near line 186, the environment-count-agnostic PromotionStrat… | 0 |
| 272 | 139 | blocking | pushback |  | Maybe add 'with observation only' or 'without actuation', as there ar… | This is headline point #12. The live .tex line 288 already substantially answers it by sp… | 12 |
| 273 | 139 | blocking | table-fix |  | Ok but blocking is not a normal adaptation action. It just indicates… | Table 6.1 (tab:control-mapping, lines 297-316) currently maps Plant = 'API deployment (so… | 12 |
| 276 | 140 | blocking | figure |  | If you change the content of the previous table, the same should be d… | Cross-linked with idx 273 (Table 6.1 re-mapping). If Table 6.1's plant/setpoint/actuator… | 12 |
| 278 | 140 | blocking | pushback |  | Two issues: (a) if the goal is to conduct validation on the current c… | Anchored at the verdict-tier convergence paragraph (line 364) and the robustness paragrap… | 12 |
| 279 | 141 | blocking | pushback |  | ok but the system converges through two types of actions: (a) prevent… | Anchored at the 'This is not a metaphorical mapping' paragraph, now line 379. Kritikos's… | 12 |
| 259 | 126 | major | figure |  | Comments for this figure: (a) some vertical labels are cut and not re… | Fix figures/workflow-pipeline.tex (the figure at lines 23-30, fig:gitops-pipeline). Two c… | 11 |
| 260 | 127 | major | rewrite |  | what is domain here? Does it also map to the XSDLC spec? | On line 46, the URL pattern 'https://{app}-{gate}-webhook.{domain}/{app}-{gate}-pr-trigge… | 0 |
| 262 | 131 | major | rewrite |  | please provide an example of a change at this level | On line 133 (Manifest Delivery), after 'a commit to the main branch of the GitOps reposit… | 0 |
| 263 | 132 | major | rewrite |  | depending on the affected overlays, so which environment is affected… | In the caption of fig:single-repo-branch-flow (lines 146-153) and/or the lead-in on line… | 11 |
| 264 | 132 | major | figure |  | The arrows again here (and outgoing arrow for SDT gate) are not shown… | Fix figures/single-repo-branch-flow.tex (fig:single-repo-branch-flow, lines 146-153). Arr… | 11 |
| 265 | 132 | major | rewrite |  | Could also indicate that the gate naming is indicative. In addition,… | Two changes. (1) In figures/single-repo-branch-flow.tex (and the caption, line 151), rena… | 11 |
| 268 | 137 | major | figure |  | The arrows here could again become larger ... | Fix figures/human-in-the-loop.tex (fig:human-in-the-loop, lines 246-253). Increase arrow… | 11 |
| 269 | 137 | major | figure |  | See previous comment | Duplicate marker pointing at fig:human-in-the-loop (Env2 row, 'Promoter auto-PR / SDT + F… | 11 |
| 270 | 137 | major | figure |  | See previous comment | Duplicate marker pointing at fig:human-in-the-loop (EnvN row, 'Promoter auto-PR / SDT + L… | 11 |
| 274 | 140 | major | figure |  | This text here is difficult to be seen/observed. | In figures/quality-control-loop.tex (fig:quality-control-loop, lines 322-329), the annota… | 11 |
| 275 | 140 | major | figure |  | Here it seems that you have arrows that cannot be easily seen, if not… | In figures/quality-control-loop.tex (fig:quality-control-loop), the block-diagram arrows… | 11 |
| 277 | 140 | major | pushback |  | ok but it is out of scope of your system, so it is closed-loop with n… | Maps to headline #12. The live .tex line 338 already concedes this exactly: 'it produces… | 12 |
| 282 | 142 | major | rewrite |  | why is this last in the chain? This needs some justification, as norm… | On line 393 (Continuous Integration), the doc-completeness check is described as the last… | 0 |
| 283 | 145 | major | rewrite |  | So this should be done by every platform operator, right? I can also… | Anchored at the Verification subsection lead-in (line 451, 'Post-deployment verification… | 0 |
| 257 | 125 | minor | trivial |  | in Chapter 7? | ANCHOR-NOT-FOUND for the literal phrase, but this maps to the chapter intro. The live .te… | 0 |
| 258 | 126 | minor | trivial |  | source ... | Truncated marginal note attached to the sentence about the Source Hydrator rendering over… | 0 |
| 261 | 128 | minor | rewrite |  | and check type ... | Marginal note on the latency-analysis paragraph (line 64). Kritikos notes validation exec… | 0 |
| 267 | 136 | minor | rewrite |  | correct. When this mode would be used? For manual testing done by the… | On line 241 (automation-spectrum item 4, 'Fully manual merge: no quality gate, manual mer… | 0 |
| 271 | 138 | minor | trivial | ✓ | (PR) | On line 272, the first use of 'pull request' in this section — Kritikos suggests introduc… | 0 |
| 280 | 141 | minor | pushback |  | Totally agree here. One small issue though: do you cover observabilit… | Anchored at the footnote on line 377, which already explicitly addresses this: 'the use o… | 0 |
| 281 | 141 | minor | trivial | ✓ | as we will see in the following sub-sections | Suggested wording insertion. At the end of the CI/CD intro paragraph (line 386), after li… | 0 |

### front-matter (15 comments)

*Kritikos's front-matter comments cluster on three themes. First, structural framing of both abstracts: he twice asks for a short orienting paragraph that states the starting point and the overall contribution package (the SDT/DriveBy/XSDLC layering) before diving into layer-specific paragraphs, and he prefers the Greek abstract's framing over the English one. Second, a string of precision/consistency queries on the evaluation summary numbers — the "three arms" claim that actually lists four experiments, the 80% mutation detection rate, the 3/3 multi-principle combinations, and OpenAPI version coverage. Third, two housekeeping items: add the two examination-committee members' names to the title page once known, and strip all pre-Introduction matter from the Table of Contents.*

| idx | pg | sev | class | M | Kritikos | Proposed action | →pt |
|----:|---:|-----|-------|:-:|----------|-----------------|:---:|
| 7 | 4 | blocking | rewrite |  | are these three or four? | Real inconsistency. abstract-en.tex line 15 states 'evaluated empirically across three ar… | 0 |
| 13 | 6 | blocking | rewrite |  | 3 h' 4? (three or four?) | Greek-side counterpart of idx 7 — the same three-vs-four inconsistency. abstract-gr.tex l… | 0 |
| 1 | 1 | major | rewrite |  | Please provide the names of the two examination committee members whe… | In main.tex the English title page (lines 57-69) and Greek title page (lines 86-100) curr… | 0 |
| 3 | 3 | major | structural |  | based on the following contributions. Maybe it is better to explain i… | In abstract-en.tex, insert a short orienting paragraph between the problem-statement para… | 2 |
| 4 | 4 | major | rewrite |  | So, is this OpenAPI version also supported? If yes, then it needs to… | The DriveBy paragraph (abstract-en.tex line 11) opens 'a Go command-line validator that l… | 0 |
| 5 | 4 | major | rewrite |  | not so clear what this means ... | Anchored to the XSDLC paragraph (abstract-en.tex line 13). The most likely unclear phrase… | 0 |
| 6 | 4 | major | table-fix |  | all 3 of 3 ... combinations detected. | Anchored to the controlled-evaluation sentence (abstract-en.tex line 15): '15 check-level… | 13 |
| 8 | 4 | major | table-fix |  | mutations with a 80% detection rate? | Anchored to '15 check-level mutations, 80% detection rate' (abstract-en.tex line 15). Kri… | 13 |
| 10 | 5 | major | structural |  | This is a better statement rather than in the English version that ta… | Same structural request as idx 3, applied to the Greek abstract. Kritikos prefers the Gre… | 2 |
| 11 | 6 | major | rewrite |  | tou OpenAPI (i.e. 'of OpenAPI' — version/coverage query on the OpenAP… | Greek-side equivalent of idx 4. In abstract-gr.tex line 20 the DriveBy paragraph opens 'φ… | 0 |
| 12 | 6 | major | rewrite |  | provlimatwn poy aforoyn parabiash pollaplwn arxwn (problems concernin… | Greek-side counterpart of idx 6. Kritikos wants the multi-principle (multiple-violation)… | 13 |
| 15 | 8 | major | structural | ✓ | I propose to remove everything before Introduction from the ToC. | The front matter currently force-adds itself to the ToC via \addcontentsline calls: Ackno… | 0 |
| 2 | 3 | minor | trivial |  | the API | Marginal wording insertion. In abstract-en.tex line 7, the opening sentence reads 'contra… | 0 |
| 9 | 5 | minor | trivial |  | to API? (i.e. '... of the API?') | Greek-side equivalent of idx 2. In abstract-gr.tex line 16, the contract-testing clause '… | 0 |
| 14 | 7 | minor | trivial |  | odhgoymeni apo? (driven by? — query on the phrase 'καθοδηγούμενη/οδηγ… | Anchored near the strict-mode results / keywords region (abstract-gr.tex; the '84%' sente… | 0 |

## Mechanical Sweep (do these in one pass)

Safe, judgement-free edits. Grouped by file.

**96 mechanical items total.**

**07-evaluation.tex**

- [ ] `07-evaluation.tex`:#290 (p149) — Line 36 contains non-existent bracketed reference numbers '~[494], [495], [510]'. Per headline #5, delete the bracketed numbers and the rou…
- [ ] `07-evaluation.tex`:#298 (p151) — Line 98: delete the non-existent reference '(per~[509], the supervisor's request for approximate per-repo commit counts)'. Rewrite as self-…
- [ ] `07-evaluation.tex`:#300 (p151) — This comment attaches to the XSDLC-as-IDP paragraph (line 105-107) which contains 'Per~[370], the count decomposes...' at line 107. Delete…
- [ ] `07-evaluation.tex`:#302 (p152) — The XSDLC-as-IDP-feature paragraph (line 107) is a single very long paragraph running from 'Within this IDP...' through '...reconciled auto…
- [ ] `07-evaluation.tex`:#307 (p155) — Line 176: the 'Naming note (round-1 review~[466], [476]).' lead-in contains non-existent reference numbers. Delete '[466], [476]' and recas…
- [ ] `07-evaluation.tex`:#309 (p155) — OCR-fragment comment ('current') attaching to the naming-note paragraph (line 176), almost certainly a wording suggestion to replace 'in th…
- [ ] `07-evaluation.tex`:#323 (p160) — Line 352: remove the round-1-review provenance and reference number — 'Two observations follow from the measured matrix and address the con…
- [ ] `07-evaluation.tex`:#329 (p161) — OCR-fragment wording suggestion on line 352 ('address the contradictions flagged'). Likely Kritikos suggesting 'addresses' singular or rewo…
- [ ] `07-evaluation.tex`:#332 (p161) — Table 7.9 (tab:api-suite, line 368) no-auth-api Expected Failure cell reads 'P005 (critical) + critical sweep'. The term 'critical sweep' i…
- [ ] `07-evaluation.tex`:#342 (p163) — Line 435, final sentences: 'Note: this paragraph reflects the post-revision severity. In the legacy severity assignment (P002--P004 warning…
- [ ] `07-evaluation.tex`:#348 (p166) — Wording tweak at line 524: change 'The evaluation harvests OpenAPI specifications' to past tense 'The evaluation harvested OpenAPI specific…
- [ ] `07-evaluation.tex`:#349 (p166) — Wording tweak at line 524: Kritikos suggests 'created includes' — i.e. change 'The dataset includes specifications from major API providers…
- [ ] `07-evaluation.tex`:#355 (p167) — OCR-fragment wording note at line 550/552. Likely a suggestion to introduce the results table with 'The produced pass rates are the followi…
- [ ] `07-evaluation.tex`:#356 (p168) — Line 573 contains 'Per~[490], the P005 result deserves a separate reading'. The one-word comment 'security' marks the P005/security sentenc…
- [ ] `07-evaluation.tex`:#361 (p170) — Line 619 (\paragraph{Method.}): 'The harvester ... was extended in round 6 with a 2.0 version selector'. Per headline #5, remove the develo…
- [ ] `07-evaluation.tex`:#367 (p172) — Line 683: 'The experiment was added in direct response to the supervisor's round-1 review~[550], which noted that "what I am missing is a p…
- [ ] `07-evaluation.tex`:#370 (p174) — OCR-fragment wording note on Table 7.18 (tab:sdt-feedback-delta) P003 row (line 720: 'Partial: 4xx/5xx wired, but components.responses reus…
- [ ] `07-evaluation.tex`:#379 (p177) — OCR-fragment wording note on Summary result 5 (line 764, RQ2 / agent-feedback experiment). Likely suggesting 'this experiment maps to RQ2'…

**03-methodology-ddt.tex**

- [ ] `03-methodology-ddt.tex`:#103 (p53) — Inline word suggestion: Kritikos proposes 'consumer' where the text says 'contract-first requires consumers to publish expectations'. The l…
- [ ] `03-methodology-ddt.tex`:#104 (p53) — Line 15 contains the stale provenance marker '\textit{What is being validated (round-1 review~[161])}:'. Delete the parenthetical '(round-1…
- [ ] `03-methodology-ddt.tex`:#111 (p55) — Line 56 contains '(round-1 review~[172])'. Delete the parenthetical so the heading reads '\textit{Context-dependent APIs and the determinis…
- [ ] `03-methodology-ddt.tex`:#117 (p57) — Line 84 contains '\textit{Severity semantics (round-1 review~[180]).}'. Delete '(round-1 review~[180])' so the lead-in reads '\textit{Sever…
- [ ] `03-methodology-ddt.tex`:#120 (p57) — Inline insertion suggestion: Kritikos wants 'of the respective principle' added for precision, likely on 'a failure of this principle indic…
- [ ] `03-methodology-ddt.tex`:#122 (p57) — Inline word change in the P001 paragraph (lines 120-122): Kritikos suggests 'utilised' for a word like 'consumed' or 'used'. The live P001…
- [ ] `03-methodology-ddt.tex`:#125 (p59) — Terminology consistency. Line 150 (P003), line 178 (P005), line 217 (P008) all say 'is mode-aware', but line 111 introduced 'Mode-specific…
- [ ] `03-methodology-ddt.tex`:#126 (p59) — Kritikos flags a non-existent reference in the P003 area. In the live .tex the stale number is at line 154: '\textit{Uniformity vs complete…
- [ ] `03-methodology-ddt.tex`:#129 (p60) — On P005 line 178, the phrase 'global requirements are set' is ambiguous. Change to 'global security requirements are set' (the live text ac…
- [ ] `03-methodology-ddt.tex`:#152 (p66) — On the Table 3.2 footnote at line 303 ('In the current implementation, ForMode(test-only) returns no static checkers; the mode instead trig…
- [ ] `03-methodology-ddt.tex`:#155 (p67) — On the Minimal-mode paragraph (line 309), Kritikos suggests describing minimal mode as 'a baseline'. Add the word: reword 'It serves as a f…
- [ ] `03-methodology-ddt.tex`:#157 (p67) — Positive acknowledgement on the Test-only-mode paragraph (line 313) — no change required. Kritikos confirms the strict-vs-test-only distinc…
- [ ] `03-methodology-ddt.tex`:#160 (p67) — Inline suggestion on the mode-progression paragraph (line 317) where test-ready is described as a pre-flight check — Kritikos endorses call…
- [ ] `03-methodology-ddt.tex`:#161 (p67) — Inline suggestion on the mode-progression paragraph (line 317) where strict is mapped to 'release pipelines' — Kritikos suggests labelling…

**02-related-work.tex**

- [ ] `02-related-work.tex`:#53 (p35) — Insertion marker: Kritikos wants 'consumer-driven' inserted before 'contracts'. Line 18 already reads 'the most widely adopted implementati…
- [ ] `02-related-work.tex`:#54 (p35) — Wording insertion on the PBT sentence (line 25). Kritikos suggests 'generating random inputs to search for violations of these properties'…
- [ ] `02-related-work.tex`:#58 (p36) — Kritikos flags an acronym (PR = pull request) that should be expanded on first use and added to the acronyms list. 'pull request' first app…
- [ ] `02-related-work.tex`:#66 (p38) — This is suggested replacement prose for the Dredd-limitation sentence (line 78), which currently reads '...it does not assess whether the s…
- [ ] `02-related-work.tex`:#69 (p39) — Kong is now mentioned in the live text (line 92: 'API gateways such as Kong enforce specification-derived policies at the edge') but it has…
- [ ] `02-related-work.tex`:#70 (p39) — Duplicate of idx 69 -- Kritikos's inline insertion of ', such as Kong, ...'. The live text already reads 'API gateways such as Kong' (line…
- [ ] `02-related-work.tex`:#72 (p40) — Line 99: 'Argo Workflows extends the Argo ecosystem with container-native workflow execution.' Kritikos questions 'container-native'; consi…
- [ ] `02-related-work.tex`:#73 (p40) — Line 108: 'High-performing teams achieve superior metrics through comprehensive automation'. Kritikos suggests 'superior performance on the…
- [ ] `02-related-work.tex`:#74 (p41) — Spelling/word insertion on the release-engineering paragraph (line 115). Kritikos flags 'artifacts'. The live text uses 'artifacts' consist…
- [ ] `02-related-work.tex`:#90 (p48) — Test-first paradigm description (line 248): 'It is labor-intensive, prone to incompleteness, and disconnected from the specification.' Krit…
- [ ] `02-related-work.tex`:#91 (p48) — Paradigm list (line 250): rename the '\item[Contract-first.]' label to '\item[Consumer contract-first.]' to distinguish it from provider-co…
- [ ] `02-related-work.tex`:#92 (p48) — Duplicate of idx 91 -- inline insertion of 'consumer' before 'Contract-first'. Resolved by the relabel to 'Consumer contract-first' in idx…
- [ ] `02-related-work.tex`:#94 (p49) — Line 262: 'Our survey identifies seven gaps in the current landscape that SDT and the XSDLC pipeline address:'. Kritikos suggests opening w…
- [ ] `02-related-work.tex`:#95 (p49) — Gap 1 (line 265): '...forcing teams to assemble ad-hoc toolchains with no guarantee of coverage.' Kritikos suggests 'no guarantee of covera…

**04-cli-architecture.tex**

- [ ] `04-cli-architecture.tex`:#165 (p69) — Kritikos is marking another occurrence of 'system' that should read 'tool' (or be scoped to the tool). On line 4 'through the data flow thr…
- [ ] `04-cli-architecture.tex`:#168 (p70) — Kritikos is labelling Fig 4.1 (line 15 caption) as 'the system's context diagram'. The caption already says 'System context'; ensure the in…
- [ ] `04-cli-architecture.tex`:#177 (p72) — In figures/dependency-flow.tex line 27 the principles node reads 'P001--P008 checkers, registry' — P009 is absent, yet the component diagra…
- [ ] `04-cli-architecture.tex`:#181 (p73) — Direct contradiction between the two figures: dependency-flow.tex says 'P001--P008' (a contiguous range implying P006/P007 included), while…
- [ ] `04-cli-architecture.tex`:#198 (p85) — The Engine struct on lines 486--492 is in a bare \begin{verbatim} block with no caption/label, unlike every other code excerpt in the chapt…
- [ ] `04-cli-architecture.tex`:#207 (p88) — Report-generation item 1 (line 583) gives a filename example 'validation-report-20260319T120000Z.json'. Item 3 (line 585), the timestamped…
- [ ] `04-cli-architecture.tex`:#211 (p92) — no action / OCR artefact — single stray character 'a' with no actionable content; verify in PDF that no real margin note is attached near l…

**05-kubernetes-architecture.tex**

- [ ] `05-kubernetes-architecture.tex`:#214 (p96) — Wording insertion. Anchor at line 89: 'The same principle checkers that run in a developer's terminal run inside an Argo Workflow container…
- [ ] `05-kubernetes-architecture.tex`:#219 (p99) — Wording. Anchor at line 140: 'In the act phase, the controller takes corrective action to drive the actual state...'. Kritikos suggests 'ac…
- [ ] `05-kubernetes-architecture.tex`:#220 (p100) — Connective insertion. Anchor at line 151, the sentence opening the self-healing logic: 'The system \textit{converges} toward the desired st…
- [ ] `05-kubernetes-architecture.tex`:#221 (p100) — Wording suggestion. Anchor at line 155: 'The three components of ontological state reconciliation are instantiated concretely in the Kubern…
- [ ] `05-kubernetes-architecture.tex`:#225 (p103) — Acronym-definition request. Anchor at line 204: 'A single \texttt{XSDLC} custom resource generates three categories of resources.' This is…
- [ ] `05-kubernetes-architecture.tex`:#233 (p106) — Text/figure inconsistency (headline #11). Anchor at line 312: 'the per-gate resources (RBAC, WorkflowTemplate, EventSource, Sensor, Service…
- [ ] `05-kubernetes-architecture.tex`:#240 (p110) — Anchor at line 396 (status-fields list): '\texttt{gates[]} (array): per-gate status including environment name, gate type, commit status ke…
- [ ] `05-kubernetes-architecture.tex`:#244 (p116) — Wording. Anchor at line 545: 'It is backed by NATS JetStream~\cite{nats-jetstream}, which provides persistent, at-least-once delivery seman…
- [ ] `05-kubernetes-architecture.tex`:#247 (p119) — Inconsistency between the worked-example DAG and the named pipeline steps. Anchor at the verbatim block lines 645-649: 'set-pending -> heal…
- [ ] `05-kubernetes-architecture.tex`:#256 (p123) — Article/parallelism fix. Anchor at line 736: 'tracing the evolution from shell script through Go CLI, container image, and workflow step to…

**09-discussion.tex**

- [ ] `09-discussion.tex`:#423 (p196) — Wording nudge on line 31/33: the RQ2 question says 'derived automatically from OpenAPI specifications' (plural) and line 33 says 'any valid…

**01-introduction.tex**

- [ ] `01-introduction.tex`:#18 (p23) — On line 20, delete the italic book title so it reads 'Forsgren, Humble, and Kim demonstrate that high-performing teams...'. Change 'demonst…
- [ ] `01-introduction.tex`:#19 (p23) — Article fix near the documentation-drift definition (line 27). Kritikos marks an 'a' that should be 'an'. The candidate is 'becomes a histo…
- [ ] `01-introduction.tex`:#21 (p24) — On the GitOps paragraph (line 43) Kritikos queries an unqualified use of 'teams'. The live line 43 does not contain a bare 'teams' — the wo…
- [ ] `01-introduction.tex`:#25 (p26) — Single-letter article fix in the tribal-knowledge opening paragraph (line 80). Verify the exact word in the PDF mark; likely an 'a'/'an' or…
- [ ] `01-introduction.tex`:#32 (p31) — On line 156 (contribution C2), Kritikos annotates 'implements SDT for OpenAPI specifications' suggesting 'implements SDT validation for Ope…
- [ ] `01-introduction.tex`:#33 (p31) — On line 158 (contribution C3), the phrase 'from a single CR' uses the acronym CR. Expand on first bare use: 'from a single Custom Resource…
- [ ] `01-introduction.tex`:#34 (p31) — On line 160 (C4) the phrase 'within an internal developer platform' is spelled out, but IDP was already defined at line 143 (RQ4 footnote).…
- [ ] `01-introduction.tex`:#41 (p33) — On line 203 (Chapter 4 summary) Kritikos marks the word 'describes' suggesting a stronger verb like 'analyses'. Change 'Chapter~\ref{ch:arc…
- [ ] `01-introduction.tex`:#42 (p33) — On line 207 (Chapter 6 summary) Kritikos suggests replacing the verb with 'details'. Change 'Chapter~\ref{ch:workflow} describes the GitOps…
- [ ] `01-introduction.tex`:#43 (p33) — On line 209 (Chapter 7 summary) a single-letter article fix is marked. Likely 'controlled validation against a reference API' is fine; veri…
- [ ] `01-introduction.tex`:#44 (p33) — Second single-letter mark on the same Chapter 7 sentence (line 209), e.g. 'a proof-of-concept operational evaluation'. Verify exact locatio…
- [ ] `01-introduction.tex`:#47 (p33) — On line 215 (Chapter 10 summary), change 'limitations' to 'the current limitations' as Kritikos suggests, to signal these are present-state…
- [ ] `01-introduction.tex`:#48 (p33) — On line 215, add a missing comma in the future-work list. The phrase 'future work including adaptive load testing thresholds and agent-driv…

**08-ai-assisted-development.tex**

- [ ] `08-ai-assisted-development.tex`:#385 (p181) — On line 97 the parenthetical lists 'the chapter structure' as one of the owned items. Kritikos suggests sharpening 'chapter structure' to '…
- [ ] `08-ai-assisted-development.tex`:#388 (p182) — Wording fix on line 97: Kritikos flags 'maps to a specific code element' style phrasing or 'every claim that maps to...'. The suggestion 'm…
- [ ] `08-ai-assisted-development.tex`:#390 (p183) — Per headline point #5: the chapter is littered with bracketed numbers that are stale provenance from a previous review round and resolve to…
- [ ] `08-ai-assisted-development.tex`:#393 (p187) — On line 245: 'The agent does not write to memory on its own initiative'. Kritikos's marginal 'writes' is a verb/agreement note. Verify the…
- [ ] `08-ai-assisted-development.tex`:#406 (p193) — Terminology consistency: Kritikos repeatedly flags (idx 406-412) that the chapter uses 'follow-on project' where he expects 'follow-up proj…
- [ ] `08-ai-assisted-development.tex`:#407 (p193) — Duplicate of idx 406 on a second 'follow-on' occurrence in the Personal Lessons section. Covered by the global 'follow-on' -> 'follow-up' r…
- [ ] `08-ai-assisted-development.tex`:#408 (p193) — Duplicate of idx 406 — another 'follow-on' to change to 'follow-up'. Corresponds to the 'Code review remained the bottleneck' lesson (line…
- [ ] `08-ai-assisted-development.tex`:#409 (p193) — Bare '?' marginal query in the Personal Lessons section, page 193. Most likely another 'follow-on/follow-up' flag or a query on the 'follow…
- [ ] `08-ai-assisted-development.tex`:#411 (p193) — Duplicate 'follow-on' -> 'follow-up' flag, page 193. Corresponds to 'Three specific changes for a follow-on project' (line 368). Covered by…

**10-conclusion.tex**

- [ ] `10-conclusion.tex`:#449 (p204) — Kritikos suggests the word 'consumer' should read 'customer' (or vice versa). The Summary of Contributions opening paragraph (line 11) has…
- [ ] `10-conclusion.tex`:#450 (p204) — Kritikos flagged a spelling/capitalisation of 'OpenAPI' (likely written 'Open API' or 'openapi' in the old text). The rewritten line 11 doe…
- [ ] `10-conclusion.tex`:#457 (p207) — RQ4 answer (line 42) claims 'a team onboarding controller can create XSDLC resources ... and an observability controller can adjust load te…
- [ ] `10-conclusion.tex`:#466 (p209) — Kritikos suggests opening the 'Adaptive load testing thresholds' item (line 90) with a 'We envision a...' framing rather than a bare defini…
- [ ] `10-conclusion.tex`:#467 (p209) — In the adaptive-threshold item (line 90), the trigger 'When observed concurrent users exceed the configured threshold by a margin, the cont…
- [ ] `10-conclusion.tex`:#476 (p210) — The Doctoral Continuation subsection carries stale bracketed comment-IDs that are NOT bibliography references: line 113 'Kritikos's review~…
- [ ] `10-conclusion.tex`:#480 (p211) — In the 'Conditional gate transitions' item (line 121), where it calls for 'a formal language for gate conditions', add Kritikos's refinemen…

**06-gitops-pipeline.tex**

- [ ] `06-gitops-pipeline.tex`:#271 (p138) — On line 272, the first use of 'pull request' in this section — Kritikos suggests introducing the acronym (PR) at first use since the rest o…
- [ ] `06-gitops-pipeline.tex`:#281 (p141) — Suggested wording insertion. At the end of the CI/CD intro paragraph (line 386), after listing the three pipeline design principles, add a…

**front-matter**

- [ ] `front-matter`:#15 (p8) — The front matter currently force-adds itself to the ToC via \addcontentsline calls: Acknowledgements (acknowledgements.tex line 5), Abstrac…

### Reference-bug cleanup (headline point #5 — spans many files)

**16 stale `[NNN]` / round-N-review provenance items.** Delete the bracket or convert the provenance note to self-contained prose.

- [ ] `07-evaluation.tex`:#290 (p149) — Line 36 contains non-existent bracketed reference numbers '~[494], [495], [510]'. Per headline #5, delete the bracketed numbers a…
- [ ] `07-evaluation.tex`:#298 (p151) — Line 98: delete the non-existent reference '(per~[509], the supervisor's request for approximate per-repo commit counts)'. Rewrit…
- [ ] `07-evaluation.tex`:#300 (p151) — This comment attaches to the XSDLC-as-IDP paragraph (line 105-107) which contains 'Per~[370], the count decomposes...' at line 10…
- [ ] `07-evaluation.tex`:#307 (p155) — Line 176: the 'Naming note (round-1 review~[466], [476]).' lead-in contains non-existent reference numbers. Delete '[466], [476]'…
- [ ] `07-evaluation.tex`:#323 (p160) — Line 352: remove the round-1-review provenance and reference number — 'Two observations follow from the measured matrix and addre…
- [ ] `07-evaluation.tex`:#356 (p168) — Line 573 contains 'Per~[490], the P005 result deserves a separate reading'. The one-word comment 'security' marks the P005/securi…
- [ ] `07-evaluation.tex`:#361 (p170) — Line 619 (\paragraph{Method.}): 'The harvester ... was extended in round 6 with a 2.0 version selector'. Per headline #5, remove…
- [ ] `07-evaluation.tex`:#367 (p172) — Line 683: 'The experiment was added in direct response to the supervisor's round-1 review~[550], which noted that "what I am miss…
- [ ] `03-methodology-ddt.tex`:#104 (p53) — Line 15 contains the stale provenance marker '\textit{What is being validated (round-1 review~[161])}:'. Delete the parenthetical…
- [ ] `03-methodology-ddt.tex`:#111 (p55) — Line 56 contains '(round-1 review~[172])'. Delete the parenthetical so the heading reads '\textit{Context-dependent APIs and the…
- [ ] `03-methodology-ddt.tex`:#117 (p57) — Line 84 contains '\textit{Severity semantics (round-1 review~[180]).}'. Delete '(round-1 review~[180])' so the lead-in reads '\te…
- [ ] `03-methodology-ddt.tex`:#126 (p59) — Kritikos flags a non-existent reference in the P003 area. In the live .tex the stale number is at line 154: '\textit{Uniformity v…
- [ ] `02-related-work.tex`:#69 (p39) — Kong is now mentioned in the live text (line 92: 'API gateways such as Kong enforce specification-derived policies at the edge')…
- [ ] `08-ai-assisted-development.tex`:#390 (p183) — Per headline point #5: the chapter is littered with bracketed numbers that are stale provenance from a previous review round and…
- [ ] `10-conclusion.tex`:#475 (p210) — The claim at line 94 that an agent 'can generate the appropriate maxLength and pattern constraints without human guidance' needs…
- [ ] `10-conclusion.tex`:#476 (p210) — The Doctoral Continuation subsection carries stale bracketed comment-IDs that are NOT bibliography references: line 113 'Kritikos…

## Recommended Order of Attack

Rough total: ~79 blocking + 224 major items, but a large fraction are already resolved-in-source (confirm-and-skip) or mechanical. Realistic critical-path work is the data regeneration (#13), the chapter split (#10), the severity reconciliation (#14), and the two pushback decisions (#3, #9).

1. **First — the data-level correctness fix (headline #13, blocking, ~half a day).** Decide the single defect-derivation story (one defect *class* → one distinguishing critical principle; documentation-class strips legitimately cascade), then regenerate the multi-API JSON results so slow-api/broken-api pass static validation and no-auth fails only P005. Everything in Ch. 7 §7.4–7.5, Figure 7.2, and the summary depends on this; doing it first prevents re-editing the same tables twice. (idx 324, 328, 333, 334, 336, 337, 340, 342, 347, 360, 375 + prose.)

2. **Second — the structural chapter split (headline #10, blocking).** Create the new "System Architecture" chapter before Ch. 4, migrate the system-context figure + three-layer model + CLI→CR evolution into it, narrow Ch. 4 to tool internals, and fix all `\ref`/`\label`. Do this before the Ch. 4/5 figure pass so figures land in the right chapter. (idx 164, 169, 170, 173 + Ch. 5 split items.)

3. **Third — the severity reconciliation (headline #14, blocking, mostly mechanical once decided).** Flip P007 → Critical everywhere (03:200/262, 07:369, "(all W)" labels), update gate rationales, add the threshold-breach-always-blocks sentence. (idx 134, 149, 150, 331, 344, 376.)

4. **Batch — the Mechanical Sweep (96 items, ~half a day, scriptable).** Run the reference-bug cleanup first (40 stale brackets + provenance narration, with the grep guard), then the per-file mechanical edits. This is the safest, highest-count win and clears the "broken citations" impression for any reader. (headline #5.)

5. **Batch — the figure pass (headline #11, Ch. 4/5/6).** Regenerate the stale `class-diagram.tex` (APISpec/PrincipleChecker/swagger2 labels), unify principle inventories across both Ch. 4 figures, fix the arrow-rendering defects across Ch. 5/6 TikZ diagrams, and give artifacts a distinct glyph. Do all figures together so layer vocabularies/colours stay consistent.

6. **Author decisions required — the pushback points (do not delegate).**
   - **#3 single-artifact (HIGH risk):** concede — rename the Table 2.1 criterion, footnote the XSDLC-CR thresholds, reconcile the two meanings of "two artefacts". (idx 78, 82, 86–89, 96, 81.)
   - **#9 P009 redundancy (HIGH risk):** do NOT send the "building block" reply; add the disjunctive-precondition justification. (idx 138, 154.)
   - **#7 governing-vs-effective:** decide flat vs split model and apply consistently (catalogue table columns). (idx 142, 143, 146.)
   - **#8 mode-specific notes (HIGH risk):** append the promised note to all nine principles; reconcile prose vs Table 3.2. (idx 121, 123, 128, 130, 137.)

7. **Then — the cross-cutting threading (headline #1, #2).** Add the agent-feedback experiment to Intro contributions/RQ-map, related-work line 169, Ch. 7 figure caption, Ch. 10 RQ answers; add ~C5 to the RQ4 row; sync RQ wording verbatim across Ch. 1/7/9. (idx 284, 285, 100, 101, 85 + intro/conclusion items.)

8. **Then — the report-authorship evidence (headline #15).** Split Table 8.1/8.2 with effort-based (not line-count) report estimates; reuse the existing caveat.

9. **Last — minor polish.** The remaining `trivial`/`minor` rewrites and the verify-the-rewrite-covers-it items (headline #12 is mostly already resolved — just the one-line label fix at 06:288). Front-matter abstract framing + ToC strip. Committee names when known.

> Confirm-and-skip as you go: headline #2, #7, #12 are largely resolved-in-source — verify the cited lines, reply with the exact file/line that satisfies Kritikos, and apply only the one small correction each carries (RQ4-row C5 mark; the governing/effective acknowledgement sentence; the 06:288 label). This converts soft email deferrals into "already done plus one fix," the strongest possible response.