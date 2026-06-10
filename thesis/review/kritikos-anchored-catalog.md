# Kritikos Comment Catalog — all 480 inline comments, anchored & explained

Generated from `kritikos-anchored-annotations.json` (rect-precise anchors, identity-verified idx) 
+ per-chapter interpretation against the live source (2026-06-10).

**Status:** 🔴 open 287 · ✅ already-addressed 163 · 🟡 needs-author-decision 19 · ❓ unclear 11 · triage corrections: 76


## 🟡 Needs your decision

- **#313** (p158, 07-evaluation.tex) `In essence, this is good as an evaluation approach.But it is not complete. In my view, complete wo`
  - on: "between two and three mutations per implemented principle, each targeting a" → 07-evaluation.tex:255
  - The per-check protocol is good but not complete — full coverage would mutate all checks across all principles, repeated with randomised violation realisations. He explicitly does not impose this ('my comment is addressed at a good level... up to you'); at most add a limitation/future-work note.
- **#371** (p174, 07-evaluation.tex) `Question / Food for thought:(a) how the agent can now which operations are to be protected and whic`
  - on: "P005 Security (row of Table 7.18)" → 07-evaluation.tex:739-747 (interpretation)
  - Food-for-thought: how can the agent know which operations should be protected or which errors each operation really produces without reading the service source code? Report-only revision risks a hallucinated spec diverging from the implementation, and functional testing only catches that if coverage exercises every error/auth path. He asks the author to consider source-aware revision (or argue P006 suffices) — explicitly a question, not a directive.
- **#140** (p63, 03-methodology-ddt.tex) `I disagree here. As in functional testing you need to check the response content if it is the expect`
  - on: "Response examples improve test reporting but are not required for meaningful validation" → 03-methodology-ddt.tex:235
  - He disagrees: functional testing should check that response content is the expected one, so response examples ARE required for meaningful validation. The author must take one consistent position (this clashes with the schema-shape-only assertion stance of line 56 and with idx 112): either defend schema-level assertion scope or soften/remove this claim. Unchanged in the live text.
- **#153** (p66, 03-methodology-ddt.tex) `I propose to remove the other note parts here as they should be supplied in the analysis of each pri`
  - on: "P004 in test-ready mode checks schema existence and type specifications only" → 03-methodology-ddt.tex:308
  - Proposal (explicitly left to the author): remove the per-principle detail footnotes under Table 3.2 since that detail should live in each principle's analysis; he notes this repetition is not harmful, so retaining is acceptable. The detail now exists in both places (per-principle mode notes added today; footnotes at lines 308-310 retained) — the author must decide whether to trim the footnotes.
- **#63** (p37, 02-related-work.tex) `I propose to move this section as Section 2.1. As it provides general information about OpenAPI and `
  - on: "2.2 The OpenAPI Ecosystem" → 02-related-work.tex:49
  - Proposes moving the OpenAPI Ecosystem section to be Section 2.1 (general OpenAPI background first), with API Testing following as 2.2, where contract testing covers both provider-driven and consumer-driven sub-types. Major restructure with cross-reference renumbering cost.
- **#65** (p38, 02-related-work.tex) `I do not understand why this section is needed as it greatly overlaps with 2.1. Maybe clarify its `
  - on: "2.2.3 API Testing Tools: Dredd and Beyond" → 02-related-work.tex:71
  - Section 2.2.3 overlaps heavily with 2.1: either clarify its distinct scope or merge it into 2.1. He notes Dredd is provider-driven contract testing, so if 2.1.1 is extended to cover both contract-testing sub-types, Dredd could move there.
- **#203** (p87, 04-cli-architecture.tex) `ok but it neglects the actual content which makes the functional testing incomplete.`
  - on: "the response body conforms to the declared schema (required fields, type checks)" → 04-cli-architecture.tex:548
  - He argues that ignoring the actual response CONTENT/values makes functional testing incomplete. The author must either add an explicit scope statement (contract conformance, not semantic correctness — deliberate boundary) or concede the limitation.
- **#217** (p98, 05-kubernetes-architecture.tex) `What is perfect-api?Maybe remove perfect prefix to make the naming independent of a specific API?`
  - on: "driveby ⋆ (namespace labels perfect-api-dev/-staging/-prod in Figure 5.1)" → figures/k8s-system-context.tex:85-87 (footnote at 05-kubernetes-architecture.tex:200)
  - Asks what perfect-api is and suggests dropping the 'perfect' prefix in the figure so the naming is independent of a specific API. The chapter text now carries a footnote (line 200) explaining the historical name, but the figure labels (k8s-system-context.tex:85-87) still read perfect-api-*; the author's recorded decision is that source identifiers keep the name. Author must decide whether to genericize the figure labels or extend the footnote to the figure.
- **#434** (p199, 09-discussion.tex) `correct but the question is whether it should also carry non-functional constraints like performance`
  - on: "the XSDLC document specifies gating policy" → 09-discussion.tex:87
  - Substantive design pushback: he believes non-functional constraints (e.g. performance thresholds) are part of the API's non-functional requirements and should arguably live IN the OpenAPI specification/contract, potentially via an x- extension, rather than only in the XSDLC CR. 'Please think about it' — the author must either defend the current split or flag an OpenAPI performance extension as future work. This decision also determines the single-vs-dual-artifact framing.
- **#444** (p201, 09-discussion.tex) `is this kind of internal validity issue? If yes, then why is it mentioned separately?`
  - on: "the interpretation of results may be influenced by" → 09-discussion.tex:114
  - He asks whether the agent-confirmation-bias threat is itself a kind of internal-validity issue, and if so why it gets its own separate category. The author must either fold the 'Agent-driven evaluation' paragraph into Internal validity or keep it separate and justify why agent-mediated interpretation warrants its own category.
- **#29** (p29, 01-introduction.tex) `here with quality gate you mean just the conditions and not the infrastructure. But in essence, the `
  - on: "quality gates be derived automatically from OpenAPI specifications (RQ2)" → chapters/01-introduction.tex:133
  - Substantive scoping note on RQ2: 'quality gate' here means the conditions, not the infrastructure, and the gates also depend on the XSDLC specification (user-supplied non-functional constraints). So the OpenAPI spec ALONG WITH user requirements drives gate production — RQ2's 'from OpenAPI specifications [alone]' should be qualified. Pairs with idx 16. The author has flagged this as the headline-point-3 pushback topic, so the resolution (qualify RQ2 vs. add a footnote vs. argue the scope) is the author's call.
- **#40** (p32, 01-introduction.tex) `I would also add a tick or tilde here as RQ4 is validated in C5. Check which from the two to add ...`
  - on: "✓ (RQ4 row of the RQ-contribution mapping table)" → chapters/01-introduction.tex:184
  - He would add a tick or tilde in the RQ4 row's C5 cell because RQ4 is validated in the empirical evaluation (C5) — and explicitly leaves the choice between \cmark and ~ to the author ('Check which from the two to add'). Live row 'RQ4 & & & \cmark & \cmark & &' still has an empty C5 cell.
- **#405** (p193, 08-ai-assisted-development.tex) `correct. But the AI couldn't supply some nice ideas that could be implanted in the original design?`
  - on: "the up-front design work is the thing that makes later acceleration possible." → 08-ai-assisted-development.tex:360
  - He accepts the claim ('correct') but asks whether the AI couldn't have supplied some good ideas worth implanting in the original design, and suggests mentioning that. The author must decide whether to concede that the agent contributed any design-level ideas (weakening the 'no AI influence on design' claim at line 108) or hold the line with an explicit acknowledgement of tactical-only suggestions.
- **#460** (p208, 10-conclusion.tex) `correct - although I foresee a specific issue: if all principles are applicable in strict or other m`
  - on: "Full principle coverage would strengthen the SDT framework's claim of comprehensive" → 10-conclusion.tex:73
  - He agrees with the limitation but warns: if P006/P007 become PrincipleCheckers that run in strict (or all) modes, dynamic validation gets conflated with static validation, contradicting the layered-validation design. The text should clarify that dynamic validation is a separate mode, not mixed into static modes. Line 73 (and the paired future-work item at line 96) still do not address the static/dynamic mode separation; the resolution is a framework design choice.
- **#470** (p209, 10-conclusion.tex) `Should we also fix code or only the OpenAPI specification (and maybe XSDLC CR)? I am asking this as `
  - on: "generate specification fixes or code patches, and re-push to trigger re-validation" → 10-conclusion.tex:94
  - Should agents fix code too, or only the OpenAPI spec (and maybe the XSDLC CR)? He recommends restricting remediation scope to the specifications so there is a stable ground truth; letting the agent edit either code or spec is problematic when it is uncertain which is correct, unless properly constrained. Line 94 still says 'specification fixes or code patches'; the remediation-scope policy is the author's design decision.
- **#473** (p210, 10-conclusion.tex) `This what was actually indicated before in a previous comment.Do we desire this if the goal is to h`
  - on: "validate-only --validation-mode strict invocation to produce a complete quality assessment" → 10-conclusion.tex:96
  - Strong objection (pairs with idx 460): strict mode was understood as full STATIC validation; saying a single strict invocation covers all nine principles means strict also runs dynamic tests, destroying the static/dynamic distinction. He proposes alternatives (full validation plus aspect-specific static and dynamic validations) and asks whether the framework/DriveBy would need adjustments. Line 96 still contains the contested sentence; resolving the mode design is the author's call.
- **#273** (p139, 06-gitops-pipeline.tex) `Ok but blocking is not a normal adaptation action. It just indicates that the live API does not sati`
  - on: "Actuator   CommitStatus + Promoter   Merges or blocks promotion PR (Table 6.1 row)" → 06-gitops-pipeline.tex:297-316 (Plant row at 308)
  - Blocking is not a corrective adaptation action on the API. His IMPORTANT instruction: the mapping is only valid if the plant is the delivery system, so Table 6.1 must be modified (e.g., Plant = pipeline & promotion workflow, Setpoint = allowed-transitions policy; or Plant = promotion-PR/release-decision system, Setpoint = predicate over artifacts and state). The live source took a different route: Table 6.1 keeps Plant = 'API deployment (source env)' but is now explicitly framed as the verdict-tier view only, with a new platform-tier section and fig:controller-stack carrying the delivery-system plant. The author must decide whether to defend this two-tier framing or actually remap/split Table 6.1 as instructed.
- **#276** (p140, 06-gitops-pipeline.tex) `If you change the content of the previous table, the same should be done for the current figure.`
  - on: "Argo Workflow DAG (controller block in Figure 6.5, fig:quality-control-loop)" → figures/quality-control-loop.tex + 06-gitops-pipeline.tex:327 (caption) and 342 (fig:controller-stack)
  - Conditional follow-up to idx 273: if Table 6.1's plant/setpoint mapping changes, Figure 6.5 must be updated identically — table and figure must stay in lockstep. Live source already adds fig:controller-stack for the platform tier and recaptions Figure 6.5 as the verdict-tier view; whether Figure 6.5 itself changes follows directly from the idx-273 decision.
- **#1** (p1, front-matter (main.tex / abstract / acronyms)) `Please provide the names of the two examination committee members when they become available.`
  - on: "Supervisor: Kyriakos Kritikos" → main.tex:65 (English title page) and main.tex:96 (Greek title page)
  - Kritikos asks that the names of the two examination committee members be added to the title page(s) once the department assigns them. Both the English and Greek title pages currently list only the supervisor.

## ❓ Unclear — candidate readings

- **#370** (p174, 07-evaluation.tex) `addressing`
  - on: "✗ (P002/P003 rows of Table 7.18)" → 07-evaluation.tex:720-722
  - Single word 'addressing' on the delta table — most plausibly a wording suggestion for a Notes cell (e.g., 'examples added' / '4xx/5xx wired' → 'addressing ...'), but the intent cannot be decoded with confidence from one word on a table cell. Candidates: reword the P002 or P003 Notes cell, or caption wording.
- **#103** (p53, 03-methodology-ddt.tex) `(consumer)`
  - on: "contract-first requires consumers to publish expectations" → 03-methodology-ddt.tex:13
  - One-word note '(consumer)' with the icon at the start of 'contract-first'. Candidate readings: (a) add a '(consumer)' / 'consumer-driven' gloss to 'contract-first' to identify it as the consumer-driven contract paradigm; (b) a singular/plural nit on 'consumers'. The intent cannot be pinned down from the icon alone.
- **#118** (p57, 03-methodology-ddt.tex) `meaning`
  - on: "Critical severity means: a failure of this principle indicates that the specification" → 03-methodology-ddt.tex:84
  - One-word note 'meaning' with the icon on 'Critical severity means:'. Candidate readings: (a) a wording suggestion on the 'means:' construction (e.g. 'Critical severity means that a failure...' or 'meaning:'), or (b) a neutral margin gloss marking this sentence as the definition. The implied edit cannot be determined with confidence.
- **#119** (p57, 03-methodology-ddt.tex) `meaning`
  - on: "Warning severity means: a failure indicates" → 03-methodology-ddt.tex:84
  - Identical one-word note 'meaning' on the parallel Warning-severity definition sentence. Same candidate readings as idx 118 — most likely a wording note on the 'X severity means:' construction.
- **#211** (p92, 04-cli-architecture.tex) `a`
  - on: "full potential is realized when embedded within declarative infrastructure" → 04-cli-architecture.tex:668
  - Single-character comment 'a'. Two candidate readings: (1) insert the article — 'embedded within A declarative infrastructure' (the icon sits right at 'declarative infrastructure', making this the likelier reading); (2) a stray/accidental annotation with no content. Verify against the PDF before editing.
- **#213** (p95, 05-kubernetes-architecture.tex) `here`
  - on: "and Crossplane's reconciliation loop provisions everything." → 04-system-architecture.tex:107
  - One-word marker 'here' on the final word of the Stage-5 paragraph. Candidate readings: (a) flags the operator(sing.)/'they declare'(pl.) agreement slip in that sentence, (b) marks a spot where something (a reference or detail) should be inserted. Intent cannot be decoded with confidence.
- **#427** (p197, 09-discussion.tex) `ok but maybe this is fully addressed?In contrast to the previous one?`
  - on: "the composition complexity tension: the Crossplane composition" → 09-discussion.tex:52
  - 'OK but maybe this is fully addressed? In contrast to the previous one?' Most likely reading: he asks whether this second tension (composition complexity) is fully addressed/mitigated, unlike the first tension (webhook reliability) which was only partially mitigated — and wants the text to say so. Alternative reading: he is musing that the issue may already be handled elsewhere in the thesis. The paragraph currently ends by calling the complexity 'inherent', with no statement of whether/how it is addressed.
- **#19** (p23, 01-introduction.tex) `an`
  - on: "implementation. Drift is not a one-time event" → chapters/01-introduction.tex:27
  - One-word comment 'an' with the icon at the end of 'its actual implementation. Drift is not a one-time event'. Most plausible reading: he wants 'an one-time event' instead of 'a one-time event' (a common letter-based article rule; actually incorrect English since 'one' starts with a /w/ sound). Alternative: an article tweak somewhere else in that sentence, but 'an API's specification' earlier is already correct.
- **#388** (p182, 08-ai-assisted-development.tex) `mapping to`
  - on: "both author-driven decisions made before Phase B:" → 08-ai-assisted-development.tex:101
  - One-phrase comment 'mapping to' whose intent cannot be pinned with confidence. Candidate readings: (a) a wording suggestion for the nearby 'the author's role would map most cleanly to architect-and-tech-lead' (line 99), rephrasing 'would map most cleanly to' as 'mapping to'; (b) a suggested insertion into 'both author-driven decisions made before Phase B, mapping to ...'. Needs the author to check the PDF placement.
- **#403** (p192, 08-ai-assisted-development.tex) `to be ...`
  - on: "for the latter, distinct from the former. This section provides them." → 08-ai-assisted-development.tex:358
  - Fragmentary comment 'to be ...' on the Personal Lessons intro sentence. Candidate readings: a started-but-unfinished wording suggestion (e.g. '... are to be provided' or smoothing 'distinct from the former'), or the beginning of a longer remark he abandoned. The sentence was rewritten today (supervisor bracket removed; now 'This section provides the latter, distinct from the former'), which may have mooted it, but the original intent cannot be decoded.
- **#449** (p204, 10-conclusion.tex) `consumer/customer`
  - on: "contract-first and test-first approaches by deriving all validation rules from" → 10-conclusion.tex:11
  - Comment 'consumer/customer' sits on 'contract-first and test-first approaches'. Most plausible reading: Kritikos is pointing at the established term 'consumer-driven (customer-driven) contract testing' and wants it named/acknowledged when the framework distinguishes itself from contract-first approaches. Alternative reading: a consumer-vs-customer terminology fix, but neither word appears in this chapter.

---

# Full catalog (by chapter, in comment order)


## front-matter (main.tex / abstract / acronyms) — 15 comments (1 open)

### 🟡 #1 (p1) — needs-author-decision
**Kritikos:** `Please provide the names of the two examination committee members when they become available.`

**Anchored on:** "Supervisor: Kyriakos Kritikos"  → **main.tex:65 (English title page) and main.tex:96 (Greek title page)**

**Meaning:** Kritikos asks that the names of the two examination committee members be added to the title page(s) once the department assigns them. Both the English and Greek title pages currently list only the supervisor.

### ✅ #2 (p3) — already-addressed
**Kritikos:** `the API`

**Anchored on:** "verify that a server matches its specification"  → **abstract-en.tex:6**

**Meaning:** One-word substitution: replace 'a server' with 'the API' so the antecedent of 'its specification' is unambiguous in the English abstract's opening sentence.

### ✅ #3 (p3) — already-addressed
**Kritikos:** `based on the following contributions+ Maybe it is better to explain in a small paragraph what is the starting point and what is the overall contribution package/architecture in terms of specific layers. Then layer-specific paragraphs can follow.`

**Anchored on:** "This thesis proposes a different starting point."  → **abstract-en.tex:8 (new orienting paragraph; opening sentence at abstract-en.tex:6)**

**Meaning:** Kritikos wants the sentence reworked to something like 'based on the following contributions', and more substantially asks for a short orienting paragraph that states the starting point and the overall contribution package/architecture in layers, before the layer-specific paragraphs. The live abstract now has exactly such a paragraph ('The proposal is developed as a three-layer contribution package: SDT... DriveBy... XSDLC...') and the opening sentence was changed to 'a different approach to this problem'.

### ✅ #4 (p4) — already-addressed
**Kritikos:** `So, is this OpenAPI version also supported? If yes, then it needs to be clarified in the beginning of the paragraph.`

**Anchored on:** "3.1.0, or a converted Swagger 2.0 document"  → **abstract-en.tex:12**

**Meaning:** The DriveBy paragraph originally opened claiming OpenAPI 3.0/3.1 support but only revealed Swagger 2.0 (converted) support at the end. Kritikos asks: is Swagger 2.0 also supported? If so, state it at the beginning of the paragraph. The live paragraph now opens 'accepts OpenAPI 3.0.x and 3.1.0 specifications natively---and Swagger 2.0 documents through a transparent conversion adapter'.

### ✅ #5 (p4) — already-addressed
**Kritikos:** `not so clear what this means ...`

**Anchored on:** "verifies single-defect detection (15 check-level mutations,"  → **abstract-en.tex:16**

**Meaning:** 'not so clear what this means ...' — the phrase 'single-defect detection (15 check-level mutations, 80% detection rate...)' is opaque to an abstract reader: what is a check-level mutation, what was detected? The live text rewrites it as 'verifies defect detection at the level of individual checks: 12 of 15 injected single-check mutations are detected...', which is self-explanatory.

**⚠ Triage correction:** Earlier triage anchored this to 'ArgoCD applications with source-hydration overlays' in the XSDLC paragraph; the icon actually sits on the controlled-evaluation sentence at 'single-defect detection (15 check-level mutations' — it is the mutation phrasing, not the hydration jargon, that he found unclear.

### ✅ #6 (p4) — already-addressed
**Kritikos:** `all 3 of 3 ... combinations detected.`

**Anchored on:** "verifies single-defect detection (15 check-level mutations, 80% detection rate; 3 of 3 random multi-principle combinations all detected)"  → **abstract-en.tex:16**

**Meaning:** Kritikos quotes a suggested rewording — 'all 3 of 3 ... combinations detected.' — i.e. move 'all' before the count instead of the awkward trailing '...combinations all detected'. The live text now reads 'all 3 evaluable random multi-principle combinations are fully detected, every targeted principle flagged simultaneously', which adopts the 'all ... detected' word order.

**⚠ Triage correction:** Earlier triage read this as a figure-verification request ('confirm the 3-of-3 count against Ch.7'); the comment is actually a wording suggestion — Kritikos is dictating the rephrased clause, not querying the number.

### ✅ #7 (p4) — already-addressed
**Kritikos:** `are these three or four?`

**Anchored on:** "The framework, the CLI, and the composition are evaluated empirically across three arms."  → **abstract-en.tex:16**

**Meaning:** 'are these three or four?' — the abstract said 'three arms' but then described four experiments (controlled, large-scale, operational PoC, agent-feedback). The live text now says 'across four arms' and the agent-feedback experiment is the explicit fourth.

### 🔴 #8 (p4) — open
**Kritikos:** `mutations with a 80% detection rate?`

**Anchored on:** "A controlled evaluation on a hand-crafted reference API (non-critical-api) ... 15 check-level mutations, 80% detection rate"  → **abstract-en.tex:16 (and the parallel error at abstract-gr.tex:25)**

**Meaning:** 'mutations with a 80% detection rate?' — Kritikos questions the 80% figure: in a controlled single-defect experiment, why were not all mutations detected, and does the number hold up? Today's rewrite explains it ('12 of 15 injected single-check mutations are detected... misses analysed as documented framework limitations'), BUT it introduced an arithmetic error: the live English abstract says 'the two misses' and the Greek says 'τις δύο αστοχίες', while 12 of 15 implies THREE misses — Ch.7 (07-evaluation.tex:758) confirms 'The three undetected single-check defects'. Both abstracts need 'two' changed to 'three'.

### ✅ #9 (p5) — already-addressed
**Kritikos:** `to API?`

**Anchored on:** "επαληθεύουν ότι ένας διακομιστής συμμορφώνεται με τις προδιαγραφές του"  → **abstract-gr.tex:15**

**Meaning:** 'to API?' = 'το API?' — Greek-side twin of idx 2: replace 'ένας διακομιστής' (a server) with 'το API' so the referent is explicit. The live Greek abstract now reads 'επαληθεύουν ότι ένα API συμμορφώνεται με τις προδιαγραφές του'.

### ✅ #10 (p5) — already-addressed
**Kritikos:** `This is a better statement rather than in English version that talks about a starting point. In any case, the proposed approach must be explained in a small paragraph before delving into the details of its main contributions`

**Anchored on:** "Η παρούσα διατριβή προτείνει μια διαφορετική προσέγγιση στο ζήτημα αυτό."  → **abstract-gr.tex:17 (new orienting paragraph)**

**Meaning:** Kritikos prefers the Greek 'different approach' framing over the English 'starting point' phrasing, and repeats his idx 3 request: a small paragraph explaining the proposed approach before the per-contribution details. Live: the English opening was aligned to 'a different approach to this problem' (abstract-en.tex:6) and both abstracts gained the three-layer package paragraph (abstract-gr.tex:17, abstract-en.tex:8).

### ✅ #11 (p6) — already-addressed
**Kritikos:** `tou OpenAPI`

**Anchored on:** "παρέχει ένα επίπεδο αφαίρεσης (abstraction) για τις λεπτομέρειες που αφορούν συγκεκριμένες εκδόσεις"  → **abstract-gr.tex:21**

**Meaning:** 'tou OpenAPI' = 'του OpenAPI' — insert the words 'του OpenAPI' so the phrase reads 'συγκεκριμένες εκδόσεις του OpenAPI' (version-specific details OF OPENAPI). The live text now contains exactly 'συγκεκριμένες εκδόσεις του OpenAPI'.

**⚠ Triage correction:** Earlier triage anchored this to the paragraph opening 'φορτώνει προδιαγραφές OpenAPI 3.0/3.1' and read it as the Greek counterpart of the idx 4 version-coverage query; the icon actually sits on the APISpec-adapter sentence and the comment is a simple two-word insertion ('εκδόσεις' + 'του OpenAPI'), unrelated to Swagger 2.0 support.

### ✅ #12 (p6) — already-addressed
**Kritikos:** `provlimatwn poy aforoyn parabiash pollaplwn arxwn`

**Anchored on:** "εντοπίστηκαν επιτυχώς 3 στους 3 τυχαίους συνδυασμούς πολλαπλών αρχών"  → **abstract-gr.tex:25**

**Meaning:** Greeklish dictation of replacement wording: 'προβλημάτων που αφορούν παραβίαση πολλαπλών αρχών' (problems concerning violation of multiple principles) — i.e. the combinations are combinations of PROBLEMS violating multiple principles, not 'combinations of principles'. The live text adopts his wording verbatim: 'οι 3 αξιολογήσιμοι τυχαίοι συνδυασμοί προβλημάτων που αφορούν παραβίαση πολλαπλών αρχών'.

### ✅ #13 (p6) — already-addressed
**Kritikos:** `3 h' 4?`

**Anchored on:** "Το πλαίσιο, η διεπαφή γραμμής εντολών (CLI) και η σύνθεση XSDLC αξιολογούνται εμπειρικά σε τρεις άξονες"  → **abstract-gr.tex:25**

**Meaning:** '3 ή 4?' (three or four?) — Greek twin of idx 7: the abstract claimed three evaluation axes but enumerated four. The live text now says 'σε τέσσερις άξονες' with Πρώτον/Δεύτερον/Τρίτον/Τέταρτον.

### ✅ #14 (p7) — already-addressed
**Kritikos:** `odhgoymeni apo?`

**Anchored on:** "υποδομή ανατροφοδότησης για τη βελτίωση προδιαγραφών με γνώμονα τους πράκτορες (agent-driven specification improvement)"  → **abstract-gr.tex:25 (keywords at abstract-gr.tex:29)**

**Meaning:** 'odhgoymeni apo?' = 'οδηγούμενη από?' — he queries the Greek rendering of 'agent-driven': the body used 'με γνώμονα τους πράκτορες' while the keywords used 'καθοδηγούμενη από πράκτορες'; he proposes a 'driven by' ('οδηγούμενη/καθοδηγούμενη από') form. The live body now reads 'που καθοδηγείται από πράκτορες', consistent with the keyword 'καθοδηγούμενη από πράκτορες'.

**⚠ Triage correction:** Earlier triage placed the anchor near the keywords line; the icon actually sits on the body sentence about feedback infrastructure, where 'με γνώμονα τους πράκτορες' stood at annotation time — the interpretation (standardise the 'driven-by' translation) was nonetheless correct.

### ✅ #15 (p8) — already-addressed
**Kritikos:** `I propose to remove everything before Introduction from ToC`

**Anchored on:** "Περίληψη (entry in the Table of Contents, between 'Abstract 1' and 'List of Figures 7')"  → **main.tex:109 (\tableofcontents; the former \addcontentsline lines in acknowledgements.tex, abstract-en.tex, abstract-gr.tex, lists-and-acronyms.tex are GONE)**

**Meaning:** Remove everything before the Introduction (Acknowledgements, Abstract, Περίληψη, List of Figures, List of Tables, Acronyms) from the Table of Contents so the ToC starts at Chapter 1. Verified against the live source: no \addcontentsline remains anywhere in chapters/ or main.tex, all front matter uses \chapter* (unnumbered, not auto-listed), and \listoffigures/\listoftables do not self-register — the ToC now begins at the Introduction.


## 01-introduction.tex — 33 comments (22 open)

### 🔴 #16 (p22) — open
**Kritikos:** `Maybe say in parenthesis or footnote (based also on some extra user-specific information given in the form an XSDLC that expresses his/her requirements over ...)`

**Anchored on:** "single source of truth from which validation rules, test cases, and quality gates are derived automatically"  → **chapters/01-introduction.tex:6**

**Meaning:** Kritikos suggests adding a parenthesis or footnote qualifying that gate derivation is not from the OpenAPI spec alone: it is 'based also on some extra user-specific information given in the form of an XSDLC that expresses his/her requirements' (i.e., non-functional thresholds come from the XSDLC CR). Implies adding a footnote on the 'derived automatically' claim acknowledging the second input.

### 🔴 #17 (p22) — open
**Kritikos:** `XSDLC is not mentioned at all as another contribution that focuses on realising quality gates`

**Anchored on:** "demonstrating its effectiveness across both controlled and large-scale settings."  → **chapters/01-introduction.tex:6**

**Meaning:** The thesis-overview paragraph names only SDT and DriveBy; Kritikos points out XSDLC is never mentioned here even though it is a separate contribution that realises the quality gates. Implies adding XSDLC to the opening enumeration of what the thesis presents.

### ✅ #18 (p23) — already-addressed
**Kritikos:** `I propose to remove "in Accelerate"`

**Anchored on:** "Forsgren, Humble, and Kim demonstrate"  → **chapters/01-introduction.tex:20**

**Meaning:** He proposes deleting the book-title phrase 'in Accelerate' so the sentence reads 'Forsgren, Humble, and Kim demonstrate that...'; the \cite{accelerate} already identifies the source.

### ❓ #19 (p23) — unclear
**Kritikos:** `an`

**Anchored on:** "implementation. Drift is not a one-time event"  → **chapters/01-introduction.tex:27**

**Meaning:** One-word comment 'an' with the icon at the end of 'its actual implementation. Drift is not a one-time event'. Most plausible reading: he wants 'an one-time event' instead of 'a one-time event' (a common letter-based article rule; actually incorrect English since 'one' starts with a /w/ sound). Alternative: an article tweak somewhere else in that sentence, but 'an API's specification' earlier is already correct.

**⚠ Triage correction:** Triage anchored this to 'becomes a historical artifact' a sentence later; the icon actually sits on 'implementation. Drift is not a one-time event', so the candidate is the article in 'a one-time event' (which is already correct English).

### 🔴 #20 (p24) — open
**Kritikos:** `So, what`

**Anchored on:** "What if the documentation itself could be the testing infrastructure?"  → **chapters/01-introduction.tex:33**

**Meaning:** Comment 'So, what' is literal replacement text: he wants the rhetorical question to begin 'So, what if the documentation itself could be the testing infrastructure?' — adding 'So,' to tie the question to the preceding argument.

**⚠ Triage correction:** Triage read this as a vague 'transition reads abruptly / consider softening the rhetorical question' concern; the simpler and correct reading is that the comment supplies the literal prefix 'So,' for the existing question.

### ✅ #21 (p24) — already-addressed
**Kritikos:** `development / project teams?`

**Anchored on:** "remains imperative: teams write tests by hand, run them"  → **chapters/01-introduction.tex:47**

**Meaning:** He queries the unqualified word 'teams' and offers 'development / project teams?' — i.e., specify which teams, e.g. 'development teams write tests by hand'.

**⚠ Triage correction:** Triage hedged that line 43 had no bare 'teams' and the anchor was uncertain; the icon is squarely on 'teams write tests by hand' and today's source already reads 'development teams write tests by hand'.

### 🔴 #22 (p24) — open
**Kritikos:** `Maybe provide here a reference that validates this claim`

**Anchored on:** "remains imperative: teams write tests by hand, run them in ad-hoc pipelines"  → **chapters/01-introduction.tex:47**

**Meaning:** He asks for a reference validating the claim that API validation remains imperative/manual (teams hand-writing tests in ad-hoc pipelines). Implies adding a citation on this sentence — e.g. reusing \cite{postman-state-of-api} or another survey of testing practice.

**⚠ Triage correction:** Triage attached the citation request to 'This model has proven remarkably effective for infrastructure and application delivery' (the GitOps claim, line 45); the icon actually sits on the imperative-testing claim in the next sentence — it is the manual-testing assertion that needs the reference.

### 🔴 #23 (p25) — open
**Kritikos:** `I believe that this section is not needed in this place.We still expect to see the current situation and its main problems/gaps and then your contribution and how it addresses these problems/gaps.`

**Anchored on:** "1.2.3 A Technology Showcase"  → **chapters/01-introduction.tex:58-61**

**Meaning:** He believes the 'A Technology Showcase' subsection does not belong here: the introduction should flow from current situation and its problems/gaps to the contribution that addresses them. Implies moving the subsection out of the problem narrative (e.g., into an architecture chapter or after the contributions).

### 🔴 #24 (p26) — open
**Kritikos:** `Could indicate here the requirement for a methodology like SDT. Such that afterwards, it is possible to state that this requirement is fulfilled by your work.`

**Anchored on:** "Validation results are structured, machine-readable quality signals tied directly to the specification. A CI"  → **chapters/01-introduction.tex:66-68**

**Meaning:** He suggests stating here an explicit REQUIREMENT for a methodology like SDT (agents need structured, spec-grounded signals), so that later the thesis can claim this requirement is fulfilled by the work. Implies adding a requirement-framing sentence before 'SDT produces exactly this kind of output'.

### 🔴 #25 (p26) — open
**Kritikos:** `a`

**Anchored on:** "the API contract is captured in structured, machine-readable form"  → **chapters/01-introduction.tex:84**

**Meaning:** One-word comment 'a' with the icon at 'is captured in structured, machine-readable form': insert the article so it reads 'captured in a structured, machine-readable form'.

**⚠ Triage correction:** Triage anchored this to 'Every software organization accumulates two kinds of knowledge' (line 80) and suspected an OCR artefact; the icon actually sits on line 84's 'captured in structured, machine-readable form' — insert 'a' before 'structured'.

### 🔴 #26 (p27) — open
**Kritikos:** `through what?`

**Anchored on:** "in a specification whose completeness is continuously validated"  → **chapters/01-introduction.tex:97**

**Meaning:** 'through what?' — he asks by what mechanism the specification's completeness is 'continuously validated'. Implies naming the mechanism, e.g. '...continuously validated through SDT/DriveBy checks on every pipeline run'.

**⚠ Triage correction:** Triage anchored this to the knowledge-externalization sentence on line 89 ('converting tacit knowledge into explicit form'); the icon actually sits on line 97's 'continuously validated' — the missing mechanism is how the validation happens, not how externalization happens.

### 🔴 #27 (p28) — open
**Kritikos:** `or technical sense but ...`

**Anchored on:** "valid in the schema sense but insufficient in the documentation sense"  → **chapters/01-introduction.tex:106**

**Meaning:** Comment 'or technical sense but ...' is literal insertion text: he wants 'valid in the schema (or technical) sense but insufficient in the documentation sense' — broadening 'schema sense' to 'schema or technical sense'.

**⚠ Triage correction:** Triage attached this to 'This assumption is not aspirational---it is increasingly achievable' (line 104) and invented an aspirational-vs-technical contrast; the icon sits on line 106's 'valid in the schema sense' and the comment is simply the words to insert: 'or technical sense but'.

### 🔴 #28 (p29) — open
**Kritikos:** `Could also refer to Evo Master here`

**Anchored on:** "Spectral [65] lints specifi-"  → **chapters/01-introduction.tex:117**

**Meaning:** He suggests also citing EvoMaster in the existing-tools enumeration (Spectral / Schemathesis / Pact) of the problem statement. The bib entry already exists (references.bib:560) and EvoMaster is discussed in Chapter 2, but it is not named here in the introduction's fragment-tools list.

### 🟡 #29 (p29) — needs-author-decision
**Kritikos:** `here with quality gate you mean just the conditions and not the infrastructure. But in essence, the quality gates might also depends on the XSDLC specification, especially in terms of non-functional constraints to be respected. Thus, OpenAPI specification along with user requirements drive the production/incarnation of the quality gates`

**Anchored on:** "quality gates be derived automatically from OpenAPI specifications (RQ2)"  → **chapters/01-introduction.tex:133**

**Meaning:** Substantive scoping note on RQ2: 'quality gate' here means the conditions, not the infrastructure, and the gates also depend on the XSDLC specification (user-supplied non-functional constraints). So the OpenAPI spec ALONG WITH user requirements drives gate production — RQ2's 'from OpenAPI specifications [alone]' should be qualified. Pairs with idx 16. The author has flagged this as the headline-point-3 pushback topic, so the resolution (qualify RQ2 vs. add a footnote vs. argue the scope) is the author's call.

### 🔴 #30 (p30) — open
**Kritikos:** `for what? Better to clearly indicate what is this system about.`

**Anchored on:** "This thesis presents a proof-of-concept system and makes six contributions."  → **chapters/01-introduction.tex:151**

**Meaning:** 'for what? Better to clearly indicate what is this system about.' — the sentence names a 'proof-of-concept system' without saying what it does. Implies expanding to identify the system's purpose (specification-driven API quality assurance via DriveBy + XSDLC).

### 🔴 #31 (p30) — open
**Kritikos:** `(Custom Resource Definition)Need to also explain in footnote what is CRD`

**Anchored on:** "within an internal developer platform? ... the XSDLC CRD functions as"  → **chapters/01-introduction.tex:141-143**

**Meaning:** '(Custom Resource Definition)' plus 'Need to also explain in footnote what is CRD' — expand the acronym CRD at first use (RQ4) and add a footnote explaining what a Custom Resource Definition is. The live RQ4 has an IDP footnote but still uses bare 'CRD' (line 143) without expansion or footnote.

### 🔴 #32 (p31) — open
**Kritikos:** `validation`

**Anchored on:** "structured reports suitable for CI/CD integration"  → **chapters/01-introduction.tex:156**

**Meaning:** One-word comment 'validation' with the icon on 'structured reports suitable for CI/CD integration': insert the word so it reads 'structured validation reports suitable for CI/CD integration' (matching the RQ2 wording at line 135).

**⚠ Triage correction:** Triage read it as 'implements SDT validation for OpenAPI specifications' (which the live text now already says); the icon actually sits on the 'structured reports' phrase at the end of C2, so the implied edit is 'structured validation reports' — still missing in the live text.

### ✅ #33 (p31) — already-addressed
**Kritikos:** `(Custom Resource)`

**Anchored on:** "from a single CR (~25 lines of YAML → 45–50 Kuber-"  → **chapters/01-introduction.tex:158**

**Meaning:** '(Custom Resource)' — expand the acronym CR at this first bare use in contribution C3.

### ✅ #34 (p31) — already-addressed
**Kritikos:** `IDP -> could use the acronym as it has been already introduced.Can use acronyms across the whole report when properly intoduced -> this could also slightly compress the current thesis report length.`

**Anchored on:** "composable building block within an internal developer plat-form"  → **chapters/01-introduction.tex:160**

**Meaning:** Use the already-introduced acronym IDP instead of spelling out 'internal developer platform' in C4; more generally, use defined acronyms consistently across the whole report to compress its length. The anchored instance now reads 'within an IDP', but the general acronym-consistency pass is a document-wide ask.

### 🔴 #35 (p31) — open
**Kritikos:** `Forgot also the last evaluation concerning the use of an agent to improve an OpenAPI spec based on its validation report.`

**Anchored on:** "5. Empirical Evaluation. We evaluate SDT through three complementary experiments"  → **chapters/01-introduction.tex:162**

**Meaning:** Contribution C5 forgets the last evaluation: the agent-driven experiment in which an AI agent improves an OpenAPI specification based on its DriveBy validation report. Implies adding it as a fourth experiment arm in C5 (live text still lists only three: controlled, large-scale, operational).

### 🔴 #36 (p32) — open
**Kritikos:** `Not clear what is Phase B`

**Anchored on:** "the 12-day Phase B agent-partnership window"  → **chapters/01-introduction.tex:164**

**Meaning:** 'Not clear what is Phase B' — the term Phase B is used without definition in C6. Implies an inline gloss or footnote defining Phase B (the agent-partnership development phase detailed in Chapter 8), and verifying Chapter 8 defines the phases.

### 🔴 #37 (p32) — open
**Kritikos:** `you mean it could be directly consumable by AI agents? As the methodology does not rely on the implemented SDT framework. But it follows similar principles: like the (deterministic) production of machine-readable reports.`

**Anchored on:** "demonstrates that SDT's structured output is directly consumable by AI agents"  → **chapters/01-introduction.tex:164**

**Meaning:** He questions the claim's framing: 'you mean it COULD be directly consumable by AI agents?' — the CLAUDE.md methodology does not rely on the implemented SDT framework, it merely follows similar principles (deterministic production of machine-readable reports). Implies softening/clarifying the clause to say the output is consumable (or could be consumed) by agents, decoupling C6's methodology from the SDT implementation.

### 🔴 #38 (p32) — open
**Kritikos:** `Maybe here we have ~ as the primary evaluation is on Chapter 7 and thus C5. A weak validation is supplied in C6 in my opinion that deserves a ~.`

**Anchored on:** "✓ (RQ2 row, C6 cell of the RQ-contribution mapping table)"  → **chapters/01-introduction.tex:182**

**Meaning:** In the mapping table he argues the RQ2/C6 check mark should be a tilde: the primary evaluation of RQ2 is Chapter 7 (C5); C6 supplies only weak validation and deserves '~'. Live row 'RQ2 & ~ & \cmark & & & \cmark & \cmark' still has \cmark in C6.

### 🔴 #39 (p32) — open
**Kritikos:** `Maybe also ~ here as SDT touches somehow RQ3.`

**Anchored on:** "✓ (RQ3 row of the RQ-contribution mapping table)"  → **chapters/01-introduction.tex:183**

**Meaning:** He suggests also adding a '~' in the RQ3 row for C1, since SDT (the conceptual framework) 'touches somehow' RQ3. Live row 'RQ3 & & ~ & \cmark & & \cmark &' still has an empty C1 cell.

### 🟡 #40 (p32) — needs-author-decision
**Kritikos:** `I would also add a tick or tilde here as RQ4 is validated in C5. Check which from the two to add ...`

**Anchored on:** "✓ (RQ4 row of the RQ-contribution mapping table)"  → **chapters/01-introduction.tex:184**

**Meaning:** He would add a tick or tilde in the RQ4 row's C5 cell because RQ4 is validated in the empirical evaluation (C5) — and explicitly leaves the choice between \cmark and ~ to the author ('Check which from the two to add'). Live row 'RQ4 & & & \cmark & \cmark & &' still has an empty C5 cell.

### ✅ #41 (p33) — already-addressed
**Kritikos:** `analyses`

**Anchored on:** "Chapter 4 describes the DriveBy CLI architecture"  → **chapters/01-introduction.tex:205**

**Meaning:** One-word comment 'analyses': replace the verb 'describes' with 'analyses' in the Chapter-4 (CLI architecture) structure summary, for verb variety/precision.

### ✅ #42 (p33) — already-addressed
**Kritikos:** `details`

**Anchored on:** "Chapter 6 describes the GitOps pipeline"  → **chapters/01-introduction.tex:209**

**Meaning:** One-word comment 'details': replace 'describes' with 'details' in the Chapter-6 (GitOps pipeline) structure summary.

### 🔴 #43 (p33) — open
**Kritikos:** `a`

**Anchored on:** "presents the empirical evaluation: controlled validation against a reference API"  → **chapters/01-introduction.tex:211**

**Meaning:** One-word comment 'a' near the start of the Chapter-7 summary list: insert the article before the first list item for parallelism with 'a proof-of-concept operational evaluation' — i.e., 'a controlled validation against a reference API'. Pairs with idx 44.

**⚠ Triage correction:** Triage suspected an OCR artefact on an already-correct article; read together with idx 44, the pair of 'a' marks asks for parallel articles: 'a controlled validation ... a large-scale validation ... and a proof-of-concept ...'.

### 🔴 #44 (p33) — open
**Kritikos:** `a`

**Anchored on:** "(non-critical-api), large-scale validation against public APIs"  → **chapters/01-introduction.tex:211**

**Meaning:** Second 'a' mark in the same sentence, icon right before 'large-scale validation': insert the article — 'a large-scale validation against public APIs from APIs.guru' — completing the parallel a/a/a structure with idx 43.

**⚠ Triage correction:** Triage guessed it sat on 'a proof-of-concept operational evaluation' as duplicate noise; the icon position shows it is on 'large-scale validation', which is missing its article.

### 🔴 #45 (p33) — open
**Kritikos:** `Missing the last evaluation again here ...`

**Anchored on:** "declarative load testing contracts, and 15-second pipeline provisioning."  → **chapters/01-introduction.tex:211**

**Meaning:** 'Missing the last evaluation again here' — the Chapter-7 structure summary, like contribution C5 (idx 35), omits the agent-driven specification-remediation evaluation. Implies appending it to the Chapter-7 item.

### 🔴 #46 (p33) — open
**Kritikos:** `Also stress that the authors contribution degree in the thesis is also detailed. The authorship as you mention it (in the title of this section)`

**Anchored on:** "SDT as feedback infrastructure for autonomous agent loops."  → **chapters/01-introduction.tex:213**

**Meaning:** He asks the Chapter-8 summary to also stress that the author's degree of contribution to the thesis is detailed there ('the authorship, as you mention it in the title of this section') — i.e., add that Chapter 8 quantifies the author's own contribution to code and report.

### ✅ #47 (p33) — already-addressed
**Kritikos:** `the current limitations`

**Anchored on:** "for operators and agents, limitations, and"  → **chapters/01-introduction.tex:217**

**Meaning:** Comment 'the current limitations' is literal replacement text: change 'limitations' to 'the current limitations' in the Chapter-10 summary.

### ✅ #48 (p33) — already-addressed
**Kritikos:** `,`

**Anchored on:** "directions for future work including adaptive load testing thresholds"  → **chapters/01-introduction.tex:217**

**Meaning:** Comment ',' — insert a comma before 'including': 'directions for future work, including adaptive load testing thresholds...'.


## 02-related-work.tex — 53 comments (32 open)

### 🔴 #49 (p34) — open
**Kritikos:** `Linting is not a kind of testing. It is usually complementary to testing. Both testing and linting are verification techniques/approaches/methods. Possible (weak) remedy: "we organise these testing and verification approaches into ...:"`

**Anchored on:** "contract testing, property-based and fuzz testing, and specification linting."  → **02-related-work.tex:11**

**Meaning:** Linting is not a kind of testing — it is a complementary verification technique. He suggests rewording the framing sentence to "we organise these testing and verification approaches into ..." so linting is not subsumed under testing.

### 🔴 #50 (p34) — open
**Kritikos:** `Please note that it is better to indicate this as consumer-driven contract testing if the consumer defines the contract based on which the testing is conducted.Contract testing is more general and consumer-driven contract testing a specific sub-type.`

**Anchored on:** "instead of the provider defining what is correct and the consumer adapting"  → **02-related-work.tex:16**

**Meaning:** The subsection should present contract testing as the general category with two sub-types — provider-driven and consumer-driven — and label the Robinson/Pact material explicitly as consumer-driven contract testing (the specific sub-type).

### 🔴 #51 (p35) — open
**Kritikos:** `not clear what is meant by local testing. Is the API somehow mocked, covering its expected behaviour?`

**Anchored on:** "Consumers record interactions during local testing"  → **02-related-work.tex:18**

**Meaning:** "Local testing" is unclear: he asks whether the provider API is mocked/stubbed during the consumer's local test runs. Clarify how the interactions are recorded (i.e., against a mock provider).

### 🔴 #52 (p35) — open
**Kritikos:** `Could also explain what are test stubs in footnote`

**Anchored on:** "generating test stubs from contract definitions"  → **02-related-work.tex:18**

**Meaning:** Add a footnote defining "test stub" (a lightweight stand-in returning canned responses) since the term is used without explanation.

### 🔴 #53 (p35) — open
**Kritikos:** `consumer-driven`

**Anchored on:** "A contract can pass while the API's documentation remains incomplete"  → **02-related-work.tex:20**

**Meaning:** One-word comment "consumer-driven" with the icon right before "A contract can pass": insert the qualifier so the limitation sentence reads "A consumer-driven contract can pass while ..." (consistent with his consumer/provider sub-type split).

**⚠ Triage correction:** Earlier triage anchored this to line 18's already-qualified "consumer-driven contracts" and marked it satisfied; the icon actually sits on line 20's "A contract can pass while", which still lacks the qualifier — open, not resolved.

### ✅ #54 (p35) — already-addressed
**Kritikos:** `of these properties`

**Anchored on:** "then generating random inputs to search for violations"  → **02-related-work.tex:25**

**Meaning:** Insert "of these properties" so the object of "violations" is explicit: "to search for violations of these properties".

### 🔴 #55 (p35) — open
**Kritikos:** `Could change paragraph here`

**Anchored on:** "crashes and security violations. EvoMaster [7] pursues"  → **02-related-work.tex:29**

**Meaning:** "Could change paragraph here": start a new paragraph at the EvoMaster sentence, splitting it off from the RESTler discussion.

**⚠ Triage correction:** Earlier triage placed the break in the QuickCheck/PBT paragraph (before "In the API-testing context..."); the icon actually sits on the RESTler paragraph at "...security violations. EvoMaster pursues" — the requested break is before the EvoMaster sentence on line 29.

### 🔴 #56 (p36) — open
**Kritikos:** `Yes, this is a kind of static analysis. Testing is about checking the actual behaviour of an API, which is different.`

**Anchored on:** "2.1.3 Specification Linting and Static Analysis"  → **02-related-work.tex:33**

**Meaning:** He affirms that linting is static analysis, not testing (testing checks actual API behaviour). This reinforces idx 49: the chapter must stop classifying linting as one of the "testing" approaches.

### 🔴 #57 (p36) — open
**Kritikos:** `is this production gated on the validation results? Could clarify in the text.`

**Anchored on:** "Redocly [56] combines linting with documentation generation. Its CLI"  → **02-related-work.tex:42**

**Meaning:** He asks whether production/promotion is actually gated on these linters' validation results, and wants the text to clarify that (e.g., that Redocly/Spectral run as advisory CI checks, not promotion gates).

**⚠ Triage correction:** Earlier triage reported ANCHOR-NOT-FOUND and attributed the comment to the Spectral paragraph; the icon actually sits on the Redocly CLI sentence (line 42).

### ✅ #58 (p36) — already-addressed
**Kritikos:** `(PR) + could identify the acronym also in the acronym's table`

**Anchored on:** "integrating this analysis into pull request workflows"  → **02-related-work.tex:44**

**Meaning:** Expand the acronym on first use — "pull request (PR)" — and add PR to the acronyms table.

### 🔴 #59 (p37) — open
**Kritikos:** `do you mean complex input parameters here or something else?`

**Anchored on:** "components (reusable schemas, parameters,"  → **02-related-work.tex:56**

**Meaning:** He asks what "parameters" means here — complex input parameters or something else? Clarify that components/parameters are reusable request-parameter definitions (path, query, header, cookie).

**⚠ Triage correction:** Earlier triage reported ANCHOR-NOT-FOUND and guessed the OpenAPI 3.1/JSON Schema sentence; the icon actually sits on "components (reusable schemas, parameters, ...)" — the question is about the meaning of 'parameters' in the components object.

### 🔴 #60 (p37) — open
**Kritikos:** `Ok but if Chapter 3 does not explain the semantics of all these sections, these sections needs to be analysed here.`

**Anchored on:** "structured as a tree rooted at openapi, info, and servers"  → **02-related-work.tex:56**

**Meaning:** If Chapter 3 does not explain the semantics of these top-level sections (openapi, info, servers, paths, components), they must be analysed here. Verified: 03-methodology-ddt.tex never mentions \texttt{paths}/\texttt{components}/\texttt{info}/\texttt{servers}, so the expansion belongs in this chapter.

### 🔴 #61 (p37) — open
**Kritikos:** `Code-First ... OpenAPI Production Workflows`

**Anchored on:** "2.2.2 Code-First and Design-First Workflows"  → **02-related-work.tex:62**

**Meaning:** Suggested reframing of the subsection title as OpenAPI *production* workflows, e.g. "OpenAPI Production Workflows: Code-First and Design-First" — these are ways specifications are produced.

### 🔴 #62 (p37) — open
**Kritikos:** `where annotations cover specific architectural components of RESTful APIs`

**Anchored on:** "frameworks generate the specification from annotated source code"  → **02-related-work.tex:65**

**Meaning:** Append his suggested clause: "...where annotations cover specific architectural components of RESTful APIs" (i.e., say what the annotations describe — endpoints, models, parameters, status codes).

### 🟡 #63 (p37) — needs-author-decision
**Kritikos:** `I propose to move this section as Section 2.1. As it provides general information about OpenAPI and how API specifications conforming to it are generated.Then, Section 2.2 will focus on how API Testing is conducted. There you could cover both forms of contract testing: provider-driven and consumer-driven contract testing along with property-based and fuzz testing ...`

**Anchored on:** "2.2 The OpenAPI Ecosystem"  → **02-related-work.tex:49**

**Meaning:** Proposes moving the OpenAPI Ecosystem section to be Section 2.1 (general OpenAPI background first), with API Testing following as 2.2, where contract testing covers both provider-driven and consumer-driven sub-types. Major restructure with cross-reference renumbering cost.

### 🔴 #64 (p38) — open
**Kritikos:** `Please change paragraph here`

**Anchored on:** "manually updated. Recent research prototypes attempt to bridge the two---Smardas"  → **02-related-work.tex:67**

**Meaning:** "Please change paragraph here": start a new paragraph before "Recent research prototypes attempt to bridge the two..." (the Smardas & Kritikos sentence).

**⚠ Triage correction:** Earlier triage proposed the break before "In design-first workflows..." (start of line 67); the icon sits later, at "updated. Recent research prototypes" — the requested break is before the Smardas/bridging sentence.

### 🟡 #65 (p38) — needs-author-decision
**Kritikos:** `I do not understand why this section is needed as it greatly overlaps with 2.1. Maybe clarify its scope and attempt to differentiate it wrt Section 2.1. Otherwise, it is better to merge it with Section 2.1Note: Dredd seems to be a provider-contract based testing tool. So, it maps to contract testing but one of its two types: provider-driven contract testing. While 2.1.1 covers consumer-driven contract testing.So, if 2.1.1 is extended to cover both types of contract testing, Dredd could be moved there.`

**Anchored on:** "2.2.3 API Testing Tools: Dredd and Beyond"  → **02-related-work.tex:71**

**Meaning:** Section 2.2.3 overlaps heavily with 2.1: either clarify its distinct scope or merge it into 2.1. He notes Dredd is provider-driven contract testing, so if 2.1.1 is extended to cover both contract-testing sub-types, Dredd could move there.

### ✅ #66 (p38) — already-addressed
**Kritikos:** `it does not assess the completeness or design quality of the specification itself, such as the presence of documented error responses, operation descriptions, security scheme definitions, or adherence to versioning and governance convention`

**Anchored on:** "GET /users returns a 200 response matching the declared schema, but it does not"  → **02-related-work.tex:78**

**Meaning:** He supplies replacement prose: "...it does not assess the completeness or design quality of the specification itself, such as the presence of documented error responses, operation descriptions, security scheme definitions, or adherence to versioning and governance conventions." The live text now contains this wording verbatim.

### 🔴 #67 (p39) — open
**Kritikos:** `Is there a proof for this? In the context of a reference?`

**Anchored on:** "its reconciliation model has not been systematically applied to API quality assurance"  → **02-related-work.tex:92**

**Meaning:** He asks for proof or a supporting reference for the claim that GitOps reconciliation has not been systematically applied to API QA. The adjacent-tools enumeration is suggestive but no citation backs the negative claim itself.

**⚠ Triage correction:** Earlier triage anchored this to "These principles eliminate imperative drift" (line 88); the icon actually sits on the "has not been systematically applied to API quality assurance" claim in line 92.

### 🔴 #68 (p39) — open
**Kritikos:** `But these are specifications? They indicate the desired state of an API. So, an API must converge towards the declared state in its specification -- provided that the API specification is not drifted (wrt the implementation). Thus, being the single source of truth.`

**Anchored on:** "Deployments converge toward declared state; API specifications, in current practice, do not"  → **02-related-work.tex:92**

**Meaning:** He challenges the framing: API specifications ARE declared desired state, so it is the API that should converge toward its specification (provided the spec has not drifted from the implementation), making the spec the single source of truth. Rework the sentence to make the spec-as-declared-state analogy explicit.

**⚠ Triage correction:** Earlier triage anchored to the OpenGitOps principle (4) sentence on line 88; the icon actually sits on line 92's "Deployments converge...; API specifications ... do not". Interpretation was otherwise correct.

### ✅ #69 (p39) — already-addressed
**Kritikos:** `missing reference here to Kong`

**Anchored on:** "API gateways such as Kong enforce specification-derived policies at the edge"  → **02-related-work.tex:92**

**Meaning:** Kong is named without a reference — add a citation. Live text now has \cite{kong} and references.bib:591 contains the @misc{kong} entry.

### ✅ #70 (p39) — already-addressed
**Kritikos:** `, such as Kong, ...`

**Anchored on:** "and API gateways such as"  → **02-related-work.tex:92**

**Meaning:** Inline insertion ", such as Kong, ..." — name Kong as the example gateway. The live text reads "API gateways such as Kong~\cite{kong}".

### 🔴 #71 (p39) — open
**Kritikos:** `Ok but are these requirements somewhere documented and justified? How they derive from the need to converge APIs into a desired, declared state?`

**Anchored on:** "a declarative principle catalogue, machine-readable per-principle output, and a reconciliation-style promotion gate"  → **02-related-work.tex:92**

**Meaning:** He asks whether these three requirements (the triple used to define the gap) are documented and justified anywhere — how do they derive from the need to converge APIs toward a declared state? Add a justification or forward-reference deriving the triple from the convergence argument.

**⚠ Triage correction:** Earlier triage anchored this to the four OpenGitOps principles (line 88); the icon actually sits on the requirements triple in line 92 — the question is about justifying SDT's three gap-defining requirements, not the OpenGitOps principles.

### ✅ #72 (p40) — already-addressed
**Kritikos:** `container-based?`

**Anchored on:** "directed acyclic graphs (DAGs) of container steps"  → **02-related-work.tex:99**

**Meaning:** "container-based?" — he questions the "container-native"/"container steps" terminology. The live text keeps "container-native" but adds the gloss "(each step runs in its own container)", resolving the ambiguity.

### ✅ #73 (p40) — already-addressed
**Kritikos:** `performance on these metrics`

**Anchored on:** "High-performing teams achieve superior metrics"  → **02-related-work.tex:108**

**Meaning:** Precision fix: "achieve superior performance on these metrics" instead of "achieve superior metrics". Live text now reads exactly that.

### 🔴 #74 (p41) — open
**Kritikos:** `artifacts`

**Anchored on:** "from commit to validated, multi-platform, registry-hosted artifact with zero manual intervention"  → **02-related-work.tex:117**

**Meaning:** One-word comment "artifacts": pluralize — "registry-hosted artifacts". The live sentence still uses the singular.

**⚠ Triage correction:** Earlier triage anchored this to line 115's "produces artifacts for both amd64 and arm64" and called it already satisfied; the icon actually sits on line 117's singular "registry-hosted artifact", which is still unpluralized — open.

### 🔴 #75 (p41) — open
**Kritikos:** `Is there a reference to back this up?`

**Anchored on:** "promotion decisions should be automated, evidence-based, and reversible"  → **02-related-work.tex:124**

**Meaning:** He asks for a reference backing the claim that promotion decisions should be automated, evidence-based, and reversible (the progressive-delivery "key insight"). Add a citation or attribute the claim.

**⚠ Triage correction:** Earlier triage anchored this to the GitOps Promoter pull-based-deployment sentence (line 122); the icon actually sits on line 124's "automated, evidence-based, and reversible" claim.

### 🔴 #76 (p42) — open
**Kritikos:** `Thus, while ..., SDT can now solve this ...`

**Anchored on:** "The declarative pattern that eliminated infrastructure drift has not yet been systematically applied"  → **02-related-work.tex:138**

**Meaning:** Add a closing turn: "Thus, while [this remains unsolved for APIs], SDT can now solve this..." — i.e., end the IaC analogy by explicitly stating SDT fills the gap.

### 🔴 #77 (p42) — open
**Kritikos:** `Would this be a API quality assessment or requirement? It depends on including maps to the handled resources or to the reconciliation model (i.e., what drives the reconciliation). This is not clear ...`

**Anchored on:** "including, in principle, API quality assessments. Second, it provides the"  → **02-related-work.tex:145**

**Meaning:** Unclear whether "API quality assessments" is meant as a governed resource or a requirement: does the Crossplane relevance lie in the resources it manages or in the reconciliation model itself? Clarify that SDT borrows the reconciliation mechanism and treats API quality as a new class of governed resource.

### 🔴 #78 (p43) — open
**Kritikos:** `Ok but in principle you validate both the API documentation and the source code. So, validation results cover both. In the current context, the resource is the API specification, the source code or both?`

**Anchored on:** "SDT applies the same pattern to validate API documentation"  → **02-related-work.tex:154**

**Meaning:** He presses that SDT validates both the API documentation AND the running API/source code (runtime principles), and asks which is "the resource" here: spec, code, or both. Line 163 (edited today) now states the governed resource is dual, but the anchored sentence at line 154 still says "validate API documentation" without qualification — a small consistency edit remains.

### 🔴 #79 (p43) — open
**Kritikos:** `Maybe indicate that policy enforcement is analogous to quality gates? A quality gate assesses "policies" and rejects the promotion. Same is done for policies, right?`

**Anchored on:** "reject resources that violate policy"  → **02-related-work.tex:154**

**Meaning:** Suggests stating explicitly that policy enforcement is analogous to quality gates: a quality gate assesses "policies" and rejects promotion, just as OPA/Kyverno reject non-conforming resources at admission. One sentence drawing the gate-rejection ↔ admission-rejection parallel.

### ✅ #80 (p43) — already-addressed
**Kritikos:** `like code changes in SDT. Maybe this is stressed in next paragraph.`

**Anchored on:** "rejecting non-conforming changes (OPA admission control)"  → **02-related-work.tex:161**

**Meaning:** He notes the SDT analogue of the correction mechanism is code changes, and wonders if the next paragraph stresses this. The next paragraph (line 163, edited today) now states remediation is dual: spec enrichment for static findings, source-code fixes for runtime findings.

### ✅ #81 (p43) — already-addressed
**Kritikos:** `+ XSDLC as it clarifies the environments, when promotion applies to them and also imposes non-functional constraints on the runtime API.`

**Anchored on:** "the OpenAPI specification---a structured declaration of what entities exist"  → **02-related-work.tex:165**

**Meaning:** Add XSDLC to the ontological-artifact discussion, since the XSDLC CR declares the environments, when promotion applies, and the non-functional constraints on the runtime API. Line 165 (added today) does exactly this.

### ✅ #82 (p44) — already-addressed
**Kritikos:** `correct. But the API source code should be also corrected or not? As some principles concern the actual/running API and not its specification.`

**Anchored on:** "specification quality convergence driven by AI agents"  → **02-related-work.tex:163**

**Meaning:** "Correct, but should the API source code also be corrected?" — some principles concern the running API, not the spec. The dual-remediation sentence added today at line 163 answers this explicitly (runtime findings are corrected by fixing the implementation's source code or capacity).

### 🔴 #83 (p45) — open
**Kritikos:** `I agree here. But you also cover the quality of the source code of the API. Is this relevant to the current discussion? I presume yes as if there is a gap between the specification and the actual API, this means that the API will not be consumed properly. This will create frustration to the consumers.`

**Anchored on:** "targeted remediation rather than blanket documentation mandates"  → **02-related-work.tex:184**

**Meaning:** He agrees, but asks whether source-code quality (the spec-vs-implementation gap) is relevant here — he presumes yes, because spec/impl drift frustrates consumers. Add a sentence in the platform-engineering subsection noting the runtime principles guard the spec/impl gap, so the self-service guarantee holds only when the running API conforms.

### 🔴 #84 (p46) — open
**Kritikos:** `maybe it is more particularly suited for the latter two types of agents?`

**Anchored on:** "principled delegation of development tasks"  → **02-related-work.tex:201**

**Meaning:** Qualify that SDT's structured output is particularly suited to the latter two AI-assistance levels (autonomous coding agents and agentic workflows), not code completion. The live text still says "naturally positioned for this ecosystem" without the "latter two" qualification.

### 🔴 #85 (p46) — open
**Kritikos:** `Correct - but also refer to the evaluation that led to a reconciliation of a specific API based on the validation feedback given by your work to an autonomous SE agent.`

**Anchored on:** "drawing on the experience of using AI-assisted development throughout the construction of"  → **02-related-work.tex:201**

**Meaning:** Also reference the empirical evaluation in which an autonomous agent reconciled a specific API based on SDT's validation feedback. That experiment exists (07-evaluation.tex, sec:sdt-feedback-experiment, Petstore 1/6→4/6), but the chapter-2 paragraph still lacks the cross-reference.

### ✅ #86 (p47) — already-addressed
**Kritikos:** `ok but based on the definition of single-artifact, it seems that it satisfies it.Unless you mean that the tool relies on test definitions that are specific to the OpenAPI spec so they can cover any OpenAPI specification? But isn't this what is also offered by DriveBy tool/CLI? Please better make this distinction as there is currently a specific overlap ...`

**Anchored on:** "since it consumes only the specification by construction"  → **02-related-work.tex:239**

**Meaning:** He disputes the "single-artifact" column: by its own definition, linters (and other spec-consuming tools) satisfy it, so the criterion overlaps with what DriveBy offers and must be redefined. Resolved today: the criterion was renamed "Zero Test Authorship" (authorship, not input cardinality), the legend rewritten, and the anchored sentence replaced.

### ✅ #87 (p47) — already-addressed
**Kritikos:** `I also do not understand why it is not a single-artifact approach. As it takes into account only the OpenAPI specification to test the running code.Maybe the definition of single-artifact is problematic and has to be improved such that all of the evaluations of this criterion are correct.`

**Anchored on:** "Schemathesis ... single-artifact ✗ (table cell)"  → **02-related-work.tex:227**

**Meaning:** He does not understand why Schemathesis/RESTler are marked non-single-artifact when they consume only the OpenAPI spec; the definition must be fixed so all marks are correct. Resolved today: under the renamed Zero-Test-Authorship criterion Schemathesis and RESTler are now \cmark, with justification in the prose at line 242.

### ✅ #88 (p47) — already-addressed
**Kritikos:** `See previous comment. All testing approaches take into consideration a single-artifact, the OpenAPI specification. So, there satisfy this criterion.`

**Anchored on:** "RESTler (table row)"  → **02-related-work.tex:228**

**Meaning:** Same point continued: all generation-based testing approaches consume the single OpenAPI artifact, so they satisfy the criterion as worded. Resolved by the criterion rename and corrected cell values (RESTler now \cmark; Pact kept \xmark and explicitly called the canonical counter-example in the prose).

### ✅ #89 (p47) — already-addressed
**Kritikos:** `Why do we have partial support for this capability? The server stub is generated if only the OpenAPI specification is syntactically valid?But isn't that the case of all testing tools? If they cannot parse the OpenAPI specification, they would not be able to test it.`

**Anchored on:** "Microcks Mock+Contract ∼ ✓ ... (table row)"  → **02-related-work.tex:231**

**Meaning:** Why partial (∼) support? If the partial "static" mark only reflects that stub generation needs a syntactically valid spec, that is true of every testing tool and is not static analysis. Resolved today: Microcks's Static cell changed from ∼ to \xmark in the live table.

### ✅ #90 (p48) — already-addressed
**Kritikos:** `and errors`

**Anchored on:** "prone to incompleteness, and disconnected from the specification"  → **02-related-work.tex:250**

**Meaning:** Insert "and errors": hand-written tests are also error-prone. Live text now reads "prone to incompleteness and errors, and disconnected from the specification".

### ✅ #91 (p48) — already-addressed
**Kritikos:** `Consumer Contract-First -> to make the distinction wrt the Provider Contract-First approaches, i.e., the documentation-driven ones.`

**Anchored on:** "Contract-first. Consumer expectations define the quality contract"  → **02-related-work.tex:252**

**Meaning:** Rename the paradigm label to "Consumer contract-first" to distinguish it from provider contract-first (documentation-driven) approaches. Live item label is now "\item[Consumer contract-first.]".

### 🔴 #92 (p48) — open
**Kritikos:** `consumer`

**Anchored on:** "implement this paradigm. The contract is the primary artifact"  → **02-related-work.tex:252**

**Meaning:** One-word comment "consumer" with the icon on "The contract": insert the qualifier so the body reads "The consumer contract is the primary artifact". The item label was renamed, but this in-body insertion was not made.

**⚠ Triage correction:** Earlier triage treated this as a duplicate of idx 91, resolved by the label rename; the icon actually sits on the item body's "The contract is the primary artifact", which still lacks "consumer" — a separate (trivial) open edit.

### 🔴 #93 (p48) — open
**Kritikos:** `Why don't you call this Specification-Driven?As you define your methodology Specification-Driven Testing (SDT).`

**Anchored on:** "Documentation-driven. The API specification is the primary artifact"  → **02-related-work.tex:254**

**Meaning:** Why not call this paradigm "Specification-driven", since the methodology is named Specification-Driven Testing (SDT)? Rename the item label; note line 257 already says "the specification-driven paradigm", so the label is internally inconsistent too.

### ✅ #94 (p49) — already-addressed
**Kritikos:** `Based on the analysis over all previous chapter sections, our ...`

**Anchored on:** "Our survey identifies seven gaps in the current landscape"  → **02-related-work.tex:264**

**Meaning:** Open with a connective tying back to the preceding analysis: "Based on the analysis across the preceding sections, our survey identifies...". Live text now reads exactly that.

### ✅ #95 (p49) — already-addressed
**Kritikos:** `of both aspects?`

**Anchored on:** "ad-hoc toolchains with no guarantee of coverage"  → **02-related-work.tex:267**

**Meaning:** Add "of both aspects" — i.e., coverage across both the static and runtime dimensions. Live text now reads "no guarantee of coverage across both static and runtime aspects".

### ✅ #96 (p49) — already-addressed
**Kritikos:** `still all are considered a single kind of source or artifact, right? Or do you believe that as we deal with plural (e.g., test cases), we have multiple artifacts to consider?This is not so clear in my opinion as in some cases we have a single artifact (a consumer contract is a set of interactions that could be imprinted in a single document, the custom rule configurations are rule-sets that constitute the normative document based on which API validation can be performed, much as in the case of your principle-based tests).Again, this is matter of what is defined as the criterion, the single-artifact one and why all other approaches fail to satisfy it.`

**Anchored on:** "Every existing tool requires inputs beyond the API specification: hand-written test cases, consumer contracts"  → **02-related-work.tex:275**

**Meaning:** His strongest single-artifact objection: contracts and rule-sets are themselves single normative documents, so the criterion as defined does not discriminate; either fix the definition or the evaluations. Resolved today: Gap 5 was rewritten as "No specification-derived full quality verdict", explicitly conceding that single-artifact operation is not novel and re-basing the gap on zero hand-authored quality artifacts plus full three-dimension verdicts.

### 🔴 #97 (p50) — open
**Kritikos:** `of course, there is a tool constructed in the case of a PhD I supervised which evaluates the maturity level of an API and supplies quality issues. But let's leave it internal between us.`

**Anchored on:** "but none assesses whether the specification is sufficient for knowledge transfer"  → **02-related-work.tex:277**

**Meaning:** He privately notes a tool from a PhD he supervised that evaluates API maturity and reports quality issues — contradicting the absolute "none assesses". Soften the claim (e.g., "no widely adopted tool assesses") without citing the internal tool. Live Gap 6 still says "none assesses".

### 🔴 #98 (p50) — open
**Kritikos:** `is this related to a specific gap? As you talk about gap dependencies here.`

**Anchored on:** "focus on testing from specifications rather than testing the specification. The absence of"  → **02-related-work.tex:282**

**Meaning:** He asks whether this sentence relates to a specific gap, since the paragraph claims to describe dependencies between gaps — each clause should name which gap depends on which.

### 🔴 #99 (p50) — open
**Kritikos:** `In essence, apart from 2nd sentence, you do not talk about gap dependencies but actually you provide justification or evidence concerning these gaps.`

**Anchored on:** "The absence of GitOps integration (Gap 4) follows from the imperative design of existing tools"  → **02-related-work.tex:282**

**Meaning:** Apart from the second sentence, the paragraph does not actually describe gap dependencies — it provides justification/evidence for each gap. Reframe the paragraph (e.g., "these gaps share common root causes") or restate genuine inter-gap dependencies.

### 🔴 #100 (p50) — open
**Kritikos:** `You forgot to mention Chapter 4. Does it cover any gap? Maybe Gap 1?`

**Anchored on:** "Chapter 3 presents the SDT framework that addresses Gaps 1--6, and"  → **02-related-work.tex:284**

**Meaning:** The gap-to-chapter mapping omits the CLI/tool chapter (DriveBy implementation) — does it cover a gap, maybe Gap 1 (unified static+runtime in one tool)? Credit that chapter in the mapping sentence. Live line 284 still maps only ch:methodology and ch:kubernetes.

### 🔴 #101 (p51) — open
**Kritikos:** `+ reference to evaluation wrt the use of the agent to reconcile an existing API specification`

**Anchored on:** "Chapter 8 returns to this thread in detail, examining how the SDT validation report served as feedback infrastructure"  → **02-related-work.tex:286**

**Meaning:** Add an explicit reference to the evaluation experiment in which an autonomous agent reconciled an existing API specification using SDT feedback (the Petstore sole-feedback experiment, sec:sdt-feedback-experiment in 07-evaluation.tex). The closing paragraph still points only to the AI-development chapter, not the empirical experiment.


## 03-methodology-ddt.tex — 61 comments (19 open)

### 🔴 #102 (p52) — open
**Kritikos:** `I believe that this is a repetition so it could be safely removed.`

**Anchored on:** "the nine principles that operationalize the axioms today, and the validation-mode discipline"  → **03-methodology-ddt.tex:6**

**Meaning:** Kritikos says the chapter-intro closing sentence ('What this chapter presents is the axiomatic foundation, the nine principles..., and the validation-mode discipline...') repeats the enumeration already given in the first paragraph and can be safely deleted.

**⚠ Triage correction:** Triage anchored this to the Section 3.1 lead-in ('SDT formalizes the third paradigm', line 13); the icon actually sits on the final sentence of the chapter-intro paragraph (line 6), which duplicates line 4's axioms/principles/modes enumeration.

### ❓ #103 (p53) — unclear
**Kritikos:** `(consumer)`

**Anchored on:** "contract-first requires consumers to publish expectations"  → **03-methodology-ddt.tex:13**

**Meaning:** One-word note '(consumer)' with the icon at the start of 'contract-first'. Candidate readings: (a) add a '(consumer)' / 'consumer-driven' gloss to 'contract-first' to identify it as the consumer-driven contract paradigm; (b) a singular/plural nit on 'consumers'. The intent cannot be pinned down from the icon alone.

**⚠ Triage correction:** Triage concluded 'likely already satisfied' because the text reads 'consumers'; but the icon sits on 'contract-first', so the more plausible ask is a '(consumer-driven)' clarifying gloss on the paradigm name, not a plural fix.

### ✅ #104 (p53) — already-addressed
**Kritikos:** `What do you mean here? In addition, you have 72 references and not [161]. So, there is a problem with the references in the document. I observe this problem to exist also in other report parts.So, please fix it.`

**Anchored on:** "What is being validated (round-1 review [161])"  → **03-methodology-ddt.tex:15**

**Meaning:** Two asks: he does not understand the '(round-1 review [161])' provenance notation, and [161] is a non-existent reference (the bibliography has only 72 entries) — a document-wide problem he asks to be fixed.

### ✅ #105 (p53) — already-addressed
**Kritikos:** `Again not clear what do you mean here.`

**Anchored on:** "The Axiom of Completeness governs the first reading; the Axiom of Determinism governs the second"  → **03-methodology-ddt.tex:15**

**Meaning:** He finds the 'first reading / second reading' axiom-assignment sentence unclear and wants it clarified or removed. In the live text this sentence is gone, replaced by 'The three axioms apply uniformly to both readings: ...'.

**⚠ Triage correction:** Triage anchored to the 'SDT operates on two artefacts' opener; the icon actually sits on the 'Axiom of Completeness governs the first reading' sentence — same passage, but the specific confusing sentence was the target (and is now removed).

### ✅ #106 (p53) — already-addressed
**Kritikos:** `I believe that this and previous sentences must be removed.`

**Anchored on:** "The earlier draft used the word “complete” loosely across both readings"  → **03-methodology-ddt.tex:15**

**Meaning:** He wants the round-1 meta-commentary ('The earlier draft used the word complete loosely...') and the preceding reviewer-reply sentence removed. The live text deleted them and kept only a neutral scoping sentence ('Throughout this thesis, the term refers specifically to the specification document...').

### ✅ #107 (p53) — already-addressed
**Kritikos:** `This is the third axiom that was not mentioned before in the paragraph!`

**Anchored on:** "regardless of when or where validation runs. If results are observable"  → **03-methodology-ddt.tex:15**

**Meaning:** Observability (the third axiom) appears in the if-clause enumeration without the paragraph having introduced the three axioms first. The live text now prefaces the enumeration with 'The three axioms apply uniformly to both readings:', announcing all three before the if-clauses; the names follow immediately at line 17.

### ✅ #108 (p53) — already-addressed
**Kritikos:** `in the end, they should cover all axioms or at least one of them? It is important to clarify!I have the impression that every principle should relate to all axioms. If not, this has to be explained.`

**Anchored on:** "each grounded in one or more axioms, that specify what the framework checks"  → **03-methodology-ddt.tex:17**

**Meaning:** He asks whether each principle covers all axioms or only some, and says this must be clarified — his impression is that every principle should relate to all three axioms. The live text replaced 'grounded in one or more axioms' with 'Every principle is bound to satisfy all three axioms simultaneously' plus the governing-vs-impacted assignment pointing to Table 3.1.

### 🔴 #109 (p54) — open
**Kritikos:** `Ok - we agree here. But in the definition of completeness in 3.2.1, you do not mention accuracy as part of completeness and you do not provide principles that cover it.`

**Anchored on:** "a specification cannot be fully “complete” in the SDT sense if its content is inaccurate"  → **03-methodology-ddt.tex:37**

**Meaning:** He agrees with folding accuracy into Completeness, but points out that the Completeness definition in Section 3.2.1 never mentions accuracy and no principle is named as covering it. The live 3.2.1 (lines 35-43) still does not state that completeness subsumes accuracy or which principles measure it (line 43 mentions P006 measuring runtime completeness but never uses the word accuracy).

### ✅ #110 (p55) — already-addressed
**Kritikos:** `to me, it seems that based on the way determinism is defined, it must be respected by all principles. So, I do not understand the sole focus on P006.`

**Anchored on:** "P006 (Functional Testing) operationalizes determinism by defining runtime tests"  → **03-methodology-ddt.tex:54**

**Meaning:** Given how Determinism is defined, all principles must respect it, so he does not understand why only P006 is singled out in the axiom subsection. The live Determinism subsection now says determinism 'is not a property of one particular principle but a structural property of how the principles compose' (line 54); P006's line 193 keeps 'operationalizes the Axiom of Determinism', which is now consistent with the governing-axiom model of line 17/Table 3.1.

### ✅ #111 (p55) — already-addressed
**Kritikos:** `again here you have a non existing reference number ...`

**Anchored on:** "Context-dependent APIs and the determinism axiom (round-1 review [172])"  → **03-methodology-ddt.tex:56**

**Meaning:** Another non-existent reference number [172] in a '(round-1 review ...)' marker; he wants it removed. The live heading at line 56 reads plain 'Context-dependent APIs and the determinism axiom.' with no marker.

### 🔴 #112 (p55) — open
**Kritikos:** `ok but functional testing is also about the actual content being returned. Thus, everything should match (status code, response schema/structure, response content). In essence, the functional testing of context-dependent APIs will require on examples that should be context-aware. So, I agree that determinism affects validation results. But it puts pressure on examples (or test case derivation techniques) which should be context-aware in this case.`

**Anchored on:** "not the literal response body. A multi-tenant API"  → **03-methodology-ddt.tex:56**

**Meaning:** He argues functional testing should also verify actual response content (status code AND schema AND content), and that for context-dependent APIs this requires context-aware examples/test-derivation — determinism shifts the burden onto the examples. The live line 56 still asserts only spec-declared properties (not literal bodies) and does not acknowledge the content-assertion question or the pressure on examples to be context-aware.

### 🔴 #113 (p55) — open
**Kritikos:** `Exactly, this is also indicated in my previous comment. Thus, the specification must be complete also in this case - thus, context-aware including context-dependent behaviour`

**Anchored on:** "Specifications that declare context-dependent behaviour (e.g. via x-context extensions"  → **03-methodology-ddt.tex:56**

**Meaning:** He agrees, and wants the point tied back to Completeness: for context-dependent APIs the specification must be complete in the context-aware sense too, i.e. declaring the context-dependent behaviour is part of completeness. The live sentence calls such specs 'first-class SDT inputs' but never connects this to the Completeness axiom.

### ✅ #114 (p56) — already-addressed
**Kritikos:** `So, the other principles do not cater for observability? This seems quite restrictive and contradictory as the whole report should be structured, machine-readable and actionable and not parts of it that map to specific principles.`

**Anchored on:** "Principles P005 (Security Standards), P007 (Performance Testing), and P008"  → **03-methodology-ddt.tex:69**

**Meaning:** He objects that naming only P005/P007/P008 as 'operationalizing observability' implies the other principles do not produce structured, machine-readable output — contradictory, since the whole report must be. The live Observability subsection deleted that sentence and now states 'Observability is a property of every principle's output' (line 69).

### ✅ #115 (p56) — already-addressed
**Kritikos:** `can't we consider completeness also as a property? If yes, then other principles also apply here.`

**Anchored on:** "operationalize observability by validating properties—security posture, performance characteristics"  → **03-methodology-ddt.tex:69**

**Meaning:** Symmetric point: if observability is about validating 'properties', completeness is also a property, so the other principles would equally qualify — i.e. the per-subset axiom partition is untenable. Resolved by the flat all-axioms model: the offending sentence is gone and lines 17, 69, and 247 state every principle satisfies all three axioms.

### ✅ #116 (p56) — already-addressed
**Kritikos:** `does the same hold for the other principles?`

**Anchored on:** "governed by Determinism (its enabling axiom) but observable by impact"  → **03-methodology-ddt.tex:69**

**Meaning:** On the old claim that P006 is 'governed by Determinism but observable by impact' he asks whether the same nature-vs-impact treatment holds for the other principles. The live text replaced the ad-hoc remark with the systematic Governing/Impacted columns in Table 3.1 covering all nine principles, and line 69 treats P006 output identically to static findings.

### ✅ #117 (p57) — already-addressed
**Kritikos:** `Again non-existing reference`

**Anchored on:** "Severity semantics (round-1 review [180])"  → **03-methodology-ddt.tex:84**

**Meaning:** Another non-existent reference [180] in a round-1-review marker; remove it. The live line 84 reads plain 'Severity semantics.'

### ❓ #118 (p57) — unclear
**Kritikos:** `meaning`

**Anchored on:** "Critical severity means: a failure of this principle indicates that the specification"  → **03-methodology-ddt.tex:84**

**Meaning:** One-word note 'meaning' with the icon on 'Critical severity means:'. Candidate readings: (a) a wording suggestion on the 'means:' construction (e.g. 'Critical severity means that a failure...' or 'meaning:'), or (b) a neutral margin gloss marking this sentence as the definition. The implied edit cannot be determined with confidence.

**⚠ Triage correction:** Triage guessed the note asked for a gloss on a circled term like 'security gating'; the icon actually sits on 'Critical severity means:', so the comment concerns the 'means'/definition wording itself.

### ❓ #119 (p57) — unclear
**Kritikos:** `meaning`

**Anchored on:** "Warning severity means: a failure indicates"  → **03-methodology-ddt.tex:84**

**Meaning:** Identical one-word note 'meaning' on the parallel Warning-severity definition sentence. Same candidate readings as idx 118 — most likely a wording note on the 'X severity means:' construction.

**⚠ Triage correction:** Triage guessed a clarifying gloss on 'degrades the specification's usefulness'; the icon sits on 'Warning severity means:', mirroring idx 118.

### 🔴 #120 (p57) — open
**Kritikos:** `of the respective principle`

**Anchored on:** "a failure indicates a quality gap that degrades the specification’s usefulness"  → **03-methodology-ddt.tex:84**

**Meaning:** Insertion suggestion: in the Warning definition, write 'a failure of the respective principle indicates a quality gap...'. The live line 84 applied 'of the respective principle' to the Critical sentence but the Warning sentence still reads bare 'a failure indicates a quality gap'.

**⚠ Triage correction:** Triage anchored the insertion to the Critical sentence ('a failure of this principle'); the icon sits on the Warning sentence, which is the one still missing the phrase — today's edit fixed only the Critical sentence.

### ✅ #121 (p57) — already-addressed
**Kritikos:** `I disagree here with the way the mode-specific behaviour is stressed. It is better to indicate that behaviour in the very beginning of the principle presentation. For instance, after the value of the Status property. In addition, while you indicate that mode-related behaviour notes are supplied, I do not see them in the first two principles!`

**Anchored on:** "Principles whose check set differs across modes carry a Mode-specific behaviour note at the end of their entry"  → **03-methodology-ddt.tex:111**

**Meaning:** He disagrees with placing mode-specific behaviour at the end of each principle entry — it should appear at the very beginning (right after the Status property) — and he notes the promised notes are missing from the first two principles (P001, P002). The live text now promises and delivers a 'Mode behaviour' note immediately after the Severity/Status line for every principle, including P001 (line 118) and P002 (line 131).

### 🔴 #122 (p57) — open
**Kritikos:** `utilised`

**Anchored on:** "and (7) all HTTP methods are standard"  → **03-methodology-ddt.tex:122**

**Meaning:** One-word insertion 'utilised': P001 check (7) should read 'all HTTP methods utilised are standard (GET, POST, ...)' — i.e. it is the methods actually used in the specification that must be standard. The live line 122 still reads 'all HTTP methods are standard'.

**⚠ Triage correction:** Triage anchored 'utilised' to 'cannot be reliably consumed by code generators' in the P001 intro; the icon actually sits on check (7) in the check list, implying 'all HTTP methods utilised are standard'.

### ✅ #123 (p58) — already-addressed
**Kritikos:** `Ok but is it mode-uniform? This is not clarified via a "promised" note ...`

**Anchored on:** "P001 runs in every validation mode. A structurally invalid specification cannot be meaningfully assessed"  → **03-methodology-ddt.tex:118**

**Meaning:** He asks whether P001 is mode-uniform, noting the promised mode-behaviour note is absent from the P001 entry. The live P001 header (line 118) now carries 'Mode behaviour: Mode-uniform—P001 runs in every static validation mode... and applies the identical seven-check set in each.'

### 🔴 #124 (p58) — open
**Kritikos:** `correct - it also prohibits P006 & P007. Without examples, it is impossible to execute them without incurring nondeterminism (as there will be the need to e.g. utilise random sampling in the test logic).`

**Anchored on:** "systematically prevent downstream agents and code generators from producing usable artefacts"  → **03-methodology-ddt.tex:133**

**Meaning:** He affirms the P002 escalation and adds a stronger argument: missing examples also prohibit P006 and P007, because without examples runtime tests cannot execute without incurring non-determinism (e.g. random sampling in test logic). The live escalation note (line 133) and the gate rationale (line 277) still do not state this P006/P007 non-determinism cascade explicitly.

### ✅ #125 (p59) — already-addressed
**Kritikos:** `mode-aware vs mode-specific. Maybe use the second term as it was introduced in the introduction of Section 3.3 and is expected to be used.`

**Anchored on:** "P003 is mode-aware. In minimal mode"  → **03-methodology-ddt.tex:146**

**Meaning:** Terminology consistency: the section intro introduced 'Mode-specific behaviour', so use 'mode-specific' instead of 'mode-aware' throughout. The live text contains no occurrence of 'mode-aware'; all principles now use the 'Mode behaviour: Mode-specific/Mode-uniform' vocabulary (e.g. line 146).

### ✅ #126 (p59) — already-addressed
**Kritikos:** `non-existing reference also here`

**Anchored on:** "Uniformity vs completeness (round-1 review [194]). Check (7) raises"  → **03-methodology-ddt.tex:154**

**Meaning:** Another non-existent reference [194] in a round-1-review marker; remove it. Live line 154 reads plain 'Uniformity vs completeness.'

### ✅ #127 (p59) — already-addressed
**Kritikos:** `Ok, I agree here with this handling. But in essence, non-uniformity is not a blocking issue. As such, potentially check (7) might not be needed in the context of the test-ready mode.Currently, we have the impression that strict and test-ready mode have the same application of all checks. Another thing: is (7) the only check related to uniformity or there are other relevant checks to be mentioned?`

**Anchored on:** "the heterogeneity is a completeness gap at the meta-level"  → **03-methodology-ddt.tex:146**

**Meaning:** He accepts folding uniformity into completeness but argues non-uniformity is not blocking, so check (7) might be dropped from test-ready mode; he also says the text gives the impression strict and test-ready run the same checks, and asks whether (7) is the only uniformity-related check. The live P003 mode note (line 146) now states test-ready runs exactly checks (1) and (5) and skips format consistency, dispelling the same-checks impression and excluding (7) from test-ready as he proposed; only the minor 'is (7) the only uniformity check' question remains implicitly answered.

### ✅ #128 (p59) — already-addressed
**Kritikos:** `Do not explain what is covered in test-ready mode as you indicate that we have mode-specific behaviour in this principle. Thus, please indicate which checks are applied in that mode.`

**Anchored on:** "(1) all operations document at least one 4xx error response"  → **03-methodology-ddt.tex:146**

**Meaning:** Since P003 is declared mode-specific, the entry must state exactly which checks run in test-ready mode rather than leaving it implied. The live P003 mode note (line 146) enumerates the test-ready subset precisely: check (1) (4xx documentation) and check (5) (structured error-schema fields).

### ✅ #129 (p60) — already-addressed
**Kritikos:** `by global requirements you mean global authentication/security requirements?`

**Anchored on:** "verifies security schemes exist and global requirements are set"  → **03-methodology-ddt.tex:178**

**Meaning:** Disambiguation question: does 'global requirements' mean global authentication/security requirements? The ambiguous minimal-mode summary sentence no longer exists (P005 now runs only in strict mode), and the surviving check (2) wording reads 'global security requirements are set at the top level'.

### ✅ #130 (p60) — already-addressed
**Kritikos:** `Again you do not indicate what holds in test-ready mode.`

**Anchored on:** "In strict mode, P005 performs all seven checks: (1) at least one security scheme is defined"  → **03-methodology-ddt.tex:174**

**Meaning:** P005's entry described minimal and strict but not test-ready mode; he wants the test-ready behaviour stated. The live P005 header (line 174) now states P005 is evaluated only in strict mode and explicitly 'not evaluated in minimal or test-ready mode', with a design rationale for the test-ready exclusion.

### 🔴 #131 (p60) — open
**Kritikos:** `general question: if an API is totally public, does this principle apply? Can we avoid it?`

**Anchored on:** "P005: Security Standards"  → **03-methodology-ddt.tex:176**

**Meaning:** General question: if an API is totally public, does P005 apply, and can it be avoided? The live P005 subsection (lines 170-180) still contains no discussion of intentionally public APIs (e.g. that declaring 'no auth required' explicitly is itself the P005-compliant posture, versus omission).

### 🔴 #132 (p61) — open
**Kritikos:** `ok but you need to stress that you rely on the examples given to assure determinism ...`

**Anchored on:** "It derives test cases from the specification—HTTP requests with appropriate parameters, headers, and bodies"  → **03-methodology-ddt.tex:193**

**Meaning:** He agrees but wants the text to stress that P006's determinism relies on the examples given in the specification — example data is what guarantees deterministic test derivation. The live P006 (lines 189-193) says inputs are 'derived deterministically from the specification' but never makes the dependence on declared examples (and P009's preconditions) explicit.

### 🔴 #133 (p61) — open
**Kritikos:** `See previous comment ... I totally agree here but only the examples can be used in this case to perform the testing. Food for thought: if examples map to real data, then potentially sensitive values are reported. Further, how can we deal with authentication? If we rely on example data, isn't that very risky?`

**Anchored on:** "functional tests must produce the same pass/fail results"  → **03-methodology-ddt.tex:191**

**Meaning:** Follow-on food for thought: only spec examples can drive deterministic functional testing, which raises two risks — (a) if examples map to real data, sensitive values may leak into reports; (b) authentication cannot safely come from example data, so how are credentials handled? The live P006 mentions 'authentication probing' exists in the implementation (line 191) but addresses neither the sensitive-example risk nor how credentials are supplied.

### ✅ #134 (p61) — already-addressed
**Kritikos:** `Not sure whether this should be warning. This needs to be properly justified. In my view, performance testing covers non-functional requirements, which are as important as the functional ones.`

**Anchored on:** "Axiom: Observability. Severity: Warning. Status: Defined theoretically"  → **03-methodology-ddt.tex:200**

**Meaning:** He disputes P007 being Warning — performance covers non-functional requirements as important as functional ones — and demands proper justification or escalation. The live line 200 now reads 'Severity: Critical.' and Table 3.1 (line 265) lists P007 as Critical, matching the agreed headline-#14 change applied today.

### 🔴 #135 (p62) — open
**Kritikos:** `Issue & food for thought: in case we have non-determinism in the environment itself, it is possible that the measured values vary across executions. Thus, it can be possible that one time the checks pass and in another time, they don't. How to deal with this? Do we consider statistically reliable values that cover the non-determinism in context/environment? Thus, only specific metrics whose values are statistically reliable?In addition, the API performance also depends on the underlying resources so it is the job of the user to properly define them, right?`

**Anchored on:** "performance metrics are inherently measurable quantities that produce structured, actionable output"  → **03-methodology-ddt.tex:204**

**Meaning:** Issue: environmental non-determinism means measured performance values vary across runs, so checks may flip between pass and fail — should P007 use only statistically reliable aggregates, and isn't resource provisioning the user's responsibility? The live P007 subsection (lines 202-206) contains no discussion of statistical robustness of thresholds or the user's responsibility for representative resources.

### ✅ #136 (p62) — already-addressed
**Kritikos:** `Did not justify why the affected axiom is observability ...Same for severity.`

**Anchored on:** "P008: API Versioning Strategy — Axiom: Observability. Severity: Warning. Status: Implemented."  → **03-methodology-ddt.tex:213**

**Meaning:** The P008 header gave no justification for the Observability axiom assignment nor for the Warning severity. In the live text the governing-axiom assignment is justified in the Section 3.5 group reasoning (line 247: observability-governed group surfaces version-lifecycle signals) and the Warning severity in the gate rationale (line 277: versioning degrades gracefully); the justification lives outside the P008 entry itself but exists.

### ✅ #137 (p62) — already-addressed
**Kritikos:** `Do not clarify what happens with test-ready mode.`

**Anchored on:** "In strict mode, P008 performs all seven checks: (1) the API version is specified"  → **03-methodology-ddt.tex:213**

**Meaning:** P008's entry did not say what happens in test-ready mode. The live P008 header (line 213) now states it is evaluated only in strict mode and 'not evaluated in minimal or test-ready mode', with a parenthetical rationale.

### ✅ #138 (p62) — already-addressed
**Kritikos:** `The principle has a name that directly maps to the test-ready mode. This gives the impression that you confuse mode with principle here. In essence, it could be argued that test-readiness could just rely on the test-ready mode as being applied to all previous relevant principles. It is really not clear why do you have a separate principle to cover all test-ready mode cases.`

**Anchored on:** "P009: Test Readiness"  → **03-methodology-ddt.tex:240**

**Meaning:** Structural challenge: P009's name maps directly onto the test-ready mode, suggesting mode and principle are confused — test-readiness could just be the test-ready mode applied over the other principles, so why a separate principle? The live P009 now ends with a dedicated 'Why a principle and not a mode' paragraph (line 240) giving a three-part defense (category distinction, disjunctive preconditions no other principle expresses, composability).

### 🔴 #139 (p63) — open
**Kritikos:** `Is this meaningful to say? If the severity is warning for this principle?`

**Anchored on:** "2xx responses have examples defined. Advisory check (warning, not failure)."  → **03-methodology-ddt.tex:235**

**Meaning:** He questions whether labelling check (4) 'warning, not failure' is meaningful given the whole P009 principle is already Warning severity — the intra-principle 'advisory' status versus the principle-level severity class needs disambiguation. The live line 235 still uses the same 'Advisory check (warning, not failure)' wording unchanged.

### 🟡 #140 (p63) — needs-author-decision
**Kritikos:** `I disagree here. As in functional testing you need to check the response content if it is the expected one.`

**Anchored on:** "Response examples improve test reporting but are not required for meaningful validation"  → **03-methodology-ddt.tex:235**

**Meaning:** He disagrees: functional testing should check that response content is the expected one, so response examples ARE required for meaningful validation. The author must take one consistent position (this clashes with the schema-shape-only assertion stance of line 56 and with idx 112): either defend schema-level assertion scope or soften/remove this claim. Unchanged in the live text.

### 🔴 #141 (p63) — open
**Kritikos:** `ok but this is risky in terms of non-determinism`

**Anchored on:** "then falls back to schema-based generation. If neither is available"  → **03-methodology-ddt.tex:233**

**Meaning:** He flags that the schema-based fallback generation in P009 check (2) is risky in terms of non-determinism (generated payloads may vary). The live line 233 is unchanged; no note that the fallback must be deterministic/seeded or that flagging missing examples is precisely how P009 avoids the non-deterministic path.

### ✅ #142 (p63) — already-addressed
**Kritikos:** `In this case, it is better to enrich the table also with the impacted axioms.I prefer this rather than just a paragraph that explains this impact per principle.`

**Anchored on:** "Several principles, however, also have downstream impact on a different axiom"  → **03-methodology-ddt.tex:257**

**Meaning:** Rather than a prose paragraph explaining per-principle axiom impact, he wants the table itself enriched with the impacted axioms. The live Table 3.1 (lines 249-271) now carries both a 'Governing Axiom' and an 'Impacted Axioms' column for every principle.

### ✅ #143 (p64) — already-addressed
**Kritikos:** `I propose to change the table content as follows:- provide first ID and Principle- order the items/rows based on ID- rename Axiom to Governing Axiom- add Impacted Axioms`

**Anchored on:** "Table 3.1: Mapping of SDT axioms to validation principles"  → **03-methodology-ddt.tex:257**

**Meaning:** Concrete table restructure: put ID and Principle first, order rows by ID, rename 'Axiom' to 'Governing Axiom', and add an 'Impacted Axioms' column. The live tab:principle-catalogue implements all four asks (header: ID, Principle, Governing Axiom, Impacted Axioms, Severity, Status; rows ordered P001-P009).

### 🔴 #144 (p64) — open
**Kritikos:** `I believe that as this is a repetition, it does not need to be also supplied here. I would just include the additional info that is supplied here in the previous part of the report that first mentions this change.`

**Anchored on:** "specifications missing those properties systematically prevented downstream agent and code-generator consumption"  → **03-methodology-ddt.tex:273**

**Meaning:** The 'Severity alignment' block repeats the P002 escalation story already told earlier; he wants the duplicate removed and any additional information (84% APIs.guru figure, commit reference) folded into the first mention. The live text still carries both blocks: the escalation note at line 133 (P002 entry) and the near-duplicate 'Severity alignment.' paragraph at line 273.

### ✅ #145 (p65) — already-addressed
**Kritikos:** `ok but don't they also have observability impact?`

**Anchored on:** "but they enable determinism by impact: a structurally invalid specification (P001)"  → **03-methodology-ddt.tex:247**

**Meaning:** On the old claim that completeness principles P001-P005 'enable determinism by impact', he asks: don't they also have observability impact? The old paragraph is gone; the live Table 3.1 answers it directly — P001-P004 rows list impacted axioms 'D, O' (observability included) and line 247 states every principle contributes to all three axioms.

### ✅ #146 (p65) — already-addressed
**Kritikos:** `I like the justification but it needs to be enhanced in order to clarify that while all principles have an overarching axiom, they also impact the other two axioms. And then you reason over groups of principles (those that have different axiom as governing/overarching).`

**Anchored on:** "the deterministic generation of valid runtime tests. Likewise, P009 is a determinism check by nature"  → **03-methodology-ddt.tex:247**

**Meaning:** He likes the cross-impact justification but wants it enhanced: state that every principle has an overarching (governing) axiom while also impacting the other two, then reason over the groups of principles sharing a governing axiom. The live Section 3.5 opening paragraph (line 247) implements exactly this model: governing vs impacted assignment plus group-by-group reasoning (completeness-, observability-, and determinism-governed groups).

**⚠ Triage correction:** Triage claimed the live text 'DENIES any governing axiom' and flagged a divergence from Kritikos's proposal; that is outdated — today's revision adopts precisely the governing+impacted two-sided model he asked for (line 17 and line 247).

### ✅ #147 (p65) — already-addressed
**Kritikos:** `So, the overarching axiom is also determinism or its just an impacted one?`

**Anchored on:** "The static principles (P001–P005, P008) are deterministic by construction"  → **03-methodology-ddt.tex:247**

**Meaning:** On 'static principles are deterministic by construction' he asks whether determinism is then their overarching axiom or just an impacted one. The live model resolves this: Table 3.1 marks Determinism as an impacted axiom (not governing) for the static principles, and line 247 explains governing = primary motivation vs impacted = also satisfied.

### ✅ #148 (p65) — already-addressed
**Kritikos:** `What about the remaining observability principles. Do they also affect determinism?`

**Anchored on:** "Determinism becomes a non-trivial requirement when validation involves runtime interaction"  → **03-methodology-ddt.tex:263**

**Meaning:** He asks whether the observability-governed principles (P005, P007, P008) also affect determinism. The live Table 3.1 rows for P005/P007/P008 list impacted axioms 'C, D' — determinism explicitly included for each.

### ✅ #149 (p65) — already-addressed
**Kritikos:** `Sure but it is catastrophic to violate a non-functional requirement (e.g., safety, performance).`

**Anchored on:** "P007 and P008 are classified as warnings because they represent quality dimensions that degrade gracefully"  → **03-methodology-ddt.tex:277**

**Meaning:** He objects that violating a non-functional requirement (safety, performance) is catastrophic, so P007 cannot be a gracefully-degrading warning. The live gate-rationale sentence (line 277) now reads 'Only versioning (P008) and test readiness advisory checks (P009) remain at warning severity' — P007 has been removed from the warning class and is Critical throughout.

### 🔴 #150 (p65) — open
**Kritikos:** `what about non-functional testing? Isn't it also important?`

**Anchored on:** "validation gates functional testing. The specification completeness principles (P001–P005)"  → **03-methodology-ddt.tex:277**

**Meaning:** He asks: what about non-functional testing — isn't it also important enough to be part of what validation gates? Although P007 is now Critical, the core design-principle sentence at line 277 still frames gating exclusively around functional testing ('validation gates functional testing') and the paragraph never mentions performance/non-functional testing.

### 🔴 #151 (p66) — open
**Kritikos:** `Correct, we can see this in the first three modes where the number of principles gets increased with the strictness of the mode. However, we cannot observe this wrt the last mode. What is really the actual rationale for that mode?In addition, please note that this mode was totally neglected from the previous analysis of the principles!!!Final thing: maybe it is not just a matter of number of principles but also the amount of checks and cost of the checks within the principles. Thus, the modes indicate eventually cost intensity. In this sense, I would recommend the following to update Table 3.2:(a) use different symbols to indicate full principle and partial principle coverage. I propose to use ~ for partial principle coverage (i.e., a subset of checks is applied)(b) use normal tick with small circle for test-only mode. This indicates that the principles surely apply in this model but with the small circle it is indicated that the principles are not yet fully implemented.`

**Anchored on:** "Modes enable incremental adoption: a team can begin with minimal validation and progressively enable stricter checks"  → **03-methodology-ddt.tex:294**

**Meaning:** Multi-part: (a) the strictness progression breaks at test-only — explain that mode's actual rationale; (b) test-only was neglected in the per-principle analysis; (c) modes really encode cost intensity (number AND cost of checks); (d) concrete Table 3.2 proposal: use '~' for partial principle coverage and a tick-with-small-circle for test-only. Parts (a) and (b) are addressed (test-only paragraph at line 317; P006/P007 mode notes name test-only), but Table 3.2 still uses cmark+dagger footnotes for partial coverage and a bare circle for test-only, and the cost-intensity framing is absent.

### ✅ #152 (p66) — already-addressed
**Kritikos:** `the test-only mode?`

**Anchored on:** "ForMode(test-only) returns no static checkers; the mode instead triggers the runtime testing pipeline"  → **03-methodology-ddt.tex:307**

**Meaning:** Short query 'the test-only mode?' — he is asking to confirm/clarify that this footnote and its circle symbol describe the test-only mode. The live footnote (line 307) explicitly names ForMode(test-only) and ties the circle marker to the Test-only row; the association is stated.

### 🟡 #153 (p66) — needs-author-decision
**Kritikos:** `I propose to remove the other note parts here as they should be supplied in the analysis of each principle. However, this is a proposal and you can make the final decision of whether to retain or not the current token. The repetition here is not so harmful than wrt the main report content.`

**Anchored on:** "P004 in test-ready mode checks schema existence and type specifications only"  → **03-methodology-ddt.tex:308**

**Meaning:** Proposal (explicitly left to the author): remove the per-principle detail footnotes under Table 3.2 since that detail should live in each principle's analysis; he notes this repetition is not harmful, so retaining is acceptable. The detail now exists in both places (per-principle mode notes added today; footnotes at lines 308-310 retained) — the author must decide whether to trim the footnotes.

### ✅ #154 (p66) — already-addressed
**Kritikos:** `I would remove P009 as it is actually implemented by P001 and partial support of P002-004. Thus, there is no need to have a separate principle here.`

**Anchored on:** "Mode / P001 ... P009 (Table 3.2 header row)"  → **03-methodology-ddt.tex:240**

**Meaning:** He proposes removing P009 entirely, arguing it is implemented by P001 plus partial P002-P004, so a separate principle is unnecessary. The author chose to keep P009 and added an explicit three-part rebuttal ('Why a principle and not a mode', line 240) showing P009's disjunctive preconditions are not a subset of P001-P004 — the challenge is answered in the text, though Kritikos may still push back.

### ✅ #155 (p67) — already-addressed
**Kritikos:** `a baseline?`

**Anchored on:** "A specification that passes minimal mode is structurally valid and safe for code generation"  → **03-methodology-ddt.tex:313**

**Meaning:** Suggestion 'a baseline?' — describe minimal mode as a baseline. The live line 313 now reads 'It serves as a baseline: a fast quality gate suitable for pull request pipelines...'.

### 🔴 #156 (p67) — open
**Kritikos:** `If I am not mistaken, usually a pre-merge gate also includes functional tests (unit, integration, system) and might include also security and performance tests.Thus, why not including all principles apart from P009?`

**Anchored on:** "appropriate for release pipelines, pre-merge gates, and periodic audits"  → **03-methodology-ddt.tex:315**

**Meaning:** He notes a pre-merge gate usually also includes functional tests (unit/integration/system) and possibly security and performance tests, so why doesn't strict mode include all principles except P009? The implied edit: justify that strict is the static-only gate (no live API required) and that a full pre-merge gate combines strict with the dynamic test-only stage. The live strict-mode paragraph (line 315) is unchanged and offers no such justification.

### ✅ #157 (p67) — already-addressed
**Kritikos:** `ok, now it is clear why you have both strict mode and test-only mode.`

**Anchored on:** "decoupling static and runtime validation into separate pipeline stages"  → **03-methodology-ddt.tex:317**

**Meaning:** Affirmation: 'ok, now it is clear why you have both strict mode and test-only mode.' No change requested; the test-only paragraph (line 317) remains intact.

### ✅ #158 (p67) — already-addressed
**Kritikos:** `Why not also include P005? It can enable functional testing to succeed when many parts of an API are secured/access controlled. Thus, in my view, some checks of P005 should have been included in the test-ready mode.`

**Anchored on:** "Test-ready mode evaluates P001, P002 (test-critical documentation...), P003..., P004 (schema types without constraints)"  → **03-methodology-ddt.tex:174**

**Meaning:** He argues some P005 (security) checks should be in test-ready mode, since functional tests against access-controlled APIs need the security declarations to succeed. The live P005 entry (line 174) now defends the exclusion explicitly: test-readiness concerns test-data derivability while authentication is exercised at the runtime layer — a documented design decision answering the question.

### 🔴 #159 (p67) — open
**Kritikos:** `Very good: maybe also indicate that this is normal practice as gates follow a cost-based scheme going from lower (e.g., initial CI gate) towards higher cost (release/production) gate. However, there is one major issue to resolve: strict mode maps to a static gate that could be mapped to a pre-merge gate. While dynamic validation maps to pre-merge or post-merge or release gate depending on its content. So, maybe you need to provide a better content for the paragraph indicating the role that each mode can play in terms of standard gates that are applied in an organisation, including pre-flight, pre-merge, post-merge and release`

**Anchored on:** "begin with minimal mode for initial CI gates, add test-ready as a pre-flight check before functional testing pipelines"  → **03-methodology-ddt.tex:321**

**Meaning:** He endorses the progression but wants it framed as the standard cost-increasing gate scheme and each mode mapped onto standard organisational gates (pre-flight, pre-merge, post-merge, release), noting strict maps to a static pre-merge gate while dynamic validation maps to pre-merge/post-merge/release depending on content. The live line 321 adopted 'pre-flight gate' and 'static pre-merge gate' but lacks the cost-based framing, the post-merge/release mapping for dynamic validation, and still confusingly says strict is a 'static pre-merge gate for release pipelines'.

### ✅ #160 (p67) — already-addressed
**Kritikos:** `pre-flight gate?`

**Anchored on:** "add test-ready as a pre-flight check before functional testing pipelines"  → **03-methodology-ddt.tex:321**

**Meaning:** Term suggestion 'pre-flight gate?' — call test-ready's role a pre-flight gate rather than a pre-flight check. The live line 321 now reads 'add test-ready as a pre-flight gate before functional testing pipelines'.

### ✅ #161 (p67) — already-addressed
**Kritikos:** `maybe say pre-merge gate`

**Anchored on:** "adopt strict for release pipelines"  → **03-methodology-ddt.tex:321**

**Meaning:** Term suggestion: label strict's role a 'pre-merge gate'. The live line 321 now reads 'adopt strict as a static pre-merge gate for release pipelines' — the term was added, though the combined phrase 'pre-merge gate for release pipelines' is awkward and worth smoothing during the idx 159 rewrite.

### 🔴 #162 (p67) — open
**Kritikos:** `map to pre-merge or post-merge/release and justify this mapping.`

**Anchored on:** "layer in test-only for runtime verification once P006 and P007 are available"  → **03-methodology-ddt.tex:321**

**Meaning:** He wants test-only (dynamic validation) explicitly mapped to a pre-merge or post-merge/release gate, with justification for the mapping (dynamic tests need a running instance, so placement depends on environment availability and cost). The live line 321 still describes test-only only as 'runtime verification' with no gate mapping or justification.


## 04-cli-architecture.tex — 49 comments (31 open)

### ✅ #163 (p69) — already-addressed
**Kritikos:** `implementation tool/mechanism?To make it clear what is this about.`

**Anchored on:** "DriveBy as a quality gate is presented in Chapter 5, and the GitOps pipelines"  → **04-cli-architecture.tex:6**

**Meaning:** Kritikos asks whether DriveBy here denotes the implementation tool/mechanism, and wants the chapter-scoping sentence to make explicit what 'DriveBy' is about (tool vs system).

**⚠ Triage correction:** Triage anchored this to the line-4 sentence 'This chapter describes how the DriveBy command-line tool implements those'; the icon actually sits on the Chapter-5 scoping sentence ('The Kubernetes infrastructure that hosts DriveBy as a quality gate...'). Same fix family, wrong anchor. Today's rewrite (line 4: 'descends into that single component') now frames DriveBy explicitly as the tool/component.

### ✅ #164 (p69) — already-addressed
**Kritikos:** `"Tool"As you mentioned in your email that the system includes Kubernetes and Crossplane. So, this is a component/tool in the overall system architecture.However, after checking the first two pages in this chapter, I am confused as you consider the tool as the system itself, which is also wrong. In my opinion, this is a serious mental problem. I believe that the overall system is the one that supports the SDT framework and the quality gate-based pipelines. Thus, I propose to have separate chapters:(a) one that reflects the overall system architecture(b) others that dive into each separate main/major component of the architecture, that is DriveBy and XSDLC, starting with the DriveBy chapter.Thus, my proposition is:Chapter 4: Overall System ArchitectureChapter 5: The DriveBy Tool...`

**Anchored on:** "The DriveBy System"  → **04-cli-architecture.tex:1**

**Meaning:** The headline structural objection: the chapter title calls the tool 'the system'. He proposes a separate preceding chapter for the overall system architecture (Kubernetes+Crossplane substrate, DriveBy and XSDLC as components) and a dedicated DriveBy-tool chapter after it.

### ✅ #165 (p69) — already-addressed
**Kritikos:** `tool's system`

**Anchored on:** "Before describing the internal structure of the system"  → **04-cli-architecture.tex:4 (successor: 'data flow through the tool'; original sentence GONE)**

**Meaning:** Another tool-vs-system flag: 'the system' here actually means the tool ('tool's system' = it is the tool's internals, not the system's). Implies replacing 'the system' with 'the tool'.

**⚠ Triage correction:** Triage anchored it to 'the data flow through the system: from specification loading'; the icon sits on 'Before describing the internal structure of the system, Figure 4.1 situates'. That sentence was removed in today's restructure; line 4 now says 'through the tool'.

### ✅ #166 (p69) — already-addressed
**Kritikos:** `via a context diagram`

**Anchored on:** "Figure 4.1 situates DriveBy within its environment"  → **04-cli-architecture.tex:8 (now: 'Where the system context diagram (Figure ...) drew the boundary'); figure itself moved to 04-system-architecture.tex:17-24**

**Meaning:** He wants the text to say explicitly that DriveBy is situated 'via a context diagram' — name the diagram type when introducing the figure.

### ✅ #167 (p69) — already-addressed
**Kritikos:** `takes a closer view on the boundary to ..."opening a boundary" looks like extending the boundary ...`

**Anchored on:** "Figure 4.3 then opens that boundary to expose the major internal Go"  → **04-cli-architecture.tex:8**

**Meaning:** 'Opening a boundary' reads as extending/enlarging the boundary; he suggests 'takes a closer view on/inside the boundary'. Implies rewording the verb phrase.

**⚠ Triage correction:** Today's rewrite changed the object: line 8 now reads 'Figure ... opens the DriveBy binary itself to expose its major internal Go packages' — the boundary-extension misreading is gone, though the verb 'opens' survives; harmless, but 'looks inside' would match his suggestion even more closely.

### ✅ #168 (p70) — already-addressed
**Kritikos:** `The system's content diagram.`

**Anchored on:** "Figure 4.1: System context. The dashed boundary encloses the system under study"  → **04-system-architecture.tex:22 (figure caption, moved out of ch.4)**

**Meaning:** Margin label: this is 'the system's context diagram' (his 'content' is a typo for 'context') — he is naming the figure type, consistent with idx 166; ensure caption/text call it the context diagram.

### ✅ #169 (p70) — already-addressed
**Kritikos:** `In my opinion, if you attempt to cover the overall system and not the DriveBy tool, then the system should be organised differently and its context diagram should be different with a more expanded boundary. In particular, in my opinion, Kubernetes and CrossPlane define the system's infrastructure while DriveBy & XSDLC lay on top of that infrastructure. The real external actors are the API Developer, the Platform Operator and external repositories (like Github). Thus, I propose to modify the diagram to reflect this.`

**Anchored on:** "ArgoCD (node in the Figure 4.1 context diagram)"  → **figures/system-context.tex:51-77 (redrawn; input at 04-system-architecture.tex:20)**

**Meaning:** Blocking figure request: if the system is the overall system (not just the tool), the boundary must expand — Kubernetes+Crossplane as infrastructure substrate, DriveBy & XSDLC on top, and external actors reduced to API Developer, Platform Operator, and external repositories (GitHub).

### ✅ #170 (p70) — already-addressed
**Kritikos:** `isn't this part of your system? As it enables the promotion ...That is why I indicate that there is a confusion about what constitutes your system.`

**Anchored on:** "GitOps Promoter (node outside the boundary in Figure 4.1)"  → **figures/system-context.tex:54; justification at 04-system-architecture.tex:26**

**Meaning:** He argues the GitOps Promoter is part of the system, since it performs the promotion the gates control — placing it outside proves the system-boundary confusion. Implies moving it inside the boundary.

### ✅ #171 (p70) — already-addressed
**Kritikos:** `In my view, this is an artifact and not a component. So, I would remove it or use a different symbol to denote it.`

**Anchored on:** "OpenAPI Specification (node in Figure 4.1)"  → **figures/system-context.tex:74 (violet 'Target API & its spec' legend category, lines 95-96)**

**Meaning:** The OpenAPI Specification is a data artifact, not a component; remove it from the diagram or denote it with a different symbol than component boxes.

**⚠ Triage correction:** Largely addressed: the redrawn figure separates the spec from components via a distinct colour category and legend, but it still uses the same rectangular node shape — if Kritikos insists on a different SYMBOL (e.g., document glyph), a residual tweak remains.

### ✅ #172 (p70) — already-addressed
**Kritikos:** `This is an artifact and not a component. So, it can be removed or represented in a different way that normal components.`

**Anchored on:** "Validation Report (JSON, exit code) (node in Figure 4.1)"  → **figures/system-context.tex:77 (gray 'Output' legend category, lines 103-104)**

**Meaning:** Same as 171 for the output: the validation report is an artifact, not a component; remove or render with a distinct artifact notation.

**⚠ Triage correction:** Same caveat as idx 171: distinguished by colour/legend ('Output'), not by node shape.

### ✅ #173 (p70) — already-addressed
**Kritikos:** `DriveBy's internal architecture.From now on, you talk about the internal architecture of a specific component wrt the overall system architecture ...`

**Anchored on:** "The DriveBy architecture is shaped by five design goals derived from the SDT"  → **04-cli-architecture.tex:20**

**Meaning:** He marks this as the point where the text switches to the INTERNAL architecture of one component ('DriveBy's internal architecture') — i.e., the natural seam for the chapter split he proposes in idx 164.

### 🔴 #174 (p70) — open
**Kritikos:** `Why does this hold only for static validation? What about dynamic validation. Shouldn't it also be deterministic?`

**Anchored on:** "Determinism by construction. Static validation must be free of side effects, randomness, and environment-dependent behavior"  → **04-cli-architecture.tex:27**

**Meaning:** Why is determinism scoped only to static validation? Dynamic (runtime) validation should also be deterministic. Implies extending design goal 3 to cover deterministic test-case derivation in runtime testing (already claimed later at line 553).

### 🔴 #175 (p71) — open
**Kritikos:** `What about the other two goals? How do they relate to the SDT axioms?`

**Anchored on:** "These goals map directly to the SDT axioms. Version agnosticism enables Completeness"  → **04-cli-architecture.tex:34**

**Meaning:** Only three of the five design goals are mapped to axioms; he asks how the other two (Extensibility, Separation of concerns) relate to the SDT axioms. Implies adding a sentence mapping or explicitly excluding them.

### 🔴 #176 (p71) — open
**Kritikos:** `You forgot to mention two packages here ....`

**Anchored on:** "The types package defines what validation produces. The spec package defines what validation consumes."  → **04-cli-architecture.tex:93**

**Meaning:** 'You forgot to mention two packages here': the separation-of-concerns summary paragraph names types, spec, principles, engine, cli but omits two packages from the seven-package dependency chain — loader and testing. Implies adding their roles to the paragraph.

**⚠ Triage correction:** Triage misanchored this to the 'eleven internal packages' chain (line 43) and read it as a four-package gap (report/util/github/logger). The icon sits on the closing summary paragraph, where exactly TWO chain packages (loader, testing) are unmentioned — matching Kritikos's count.

### ✅ #177 (p72) — already-addressed
**Kritikos:** `So, P009 was removed or not? As here we do not actually see it ...`

**Anchored on:** "principles — P001–P008 checkers, registry (node in Figure 4.2)"  → **figures/dependency-flow.tex:27 (now 'P001--P005, P008, P009 checkers, registry')**

**Meaning:** The dependency-flow figure said 'P001–P008', implying P009 was removed — contradicting the table and the component diagram. Implies relabelling the node with the actual registered set.

### 🔴 #178 (p72) — open
**Kritikos:** `What is registry? What does it include?`

**Anchored on:** "principles — P001–P008 checkers, registry (the word 'registry' in Figure 4.2)"  → **04-cli-architecture.tex:81 (table row) and figures/dependency-flow.tex:27**

**Meaning:** 'What is registry? What does it include?' — the term appears in the figure and package table long before it is defined in Section 4.5.3. Implies a gloss/footnote at first mention.

### 🔴 #179 (p73) — open
**Kritikos:** `But now you introduce different layers wrt the previous figure. This creates confusion to the reader.`

**Anchored on:** "Colour denotes architectural layer (entry point, orchestration, principles, spec abstraction, foundation)"  → **04-cli-architecture.tex:63 (caption) vs figures/dependency-flow.tex:41-46 (still 'Data/Validation/Interface layer')**

**Meaning:** The component-diagram caption introduces a five-layer taxonomy that differs from Figure 4.2's layers (Data/Validation/Interface), confusing the reader. Implies unifying layer names/colours across both figures.

### ✅ #180 (p73) — already-addressed
**Kritikos:** `Can you improve the diagram?Some arrows are not fully represented (only their end). In addition, isn't the engine calling principles and testing? If yes, then we do not see an arrow from engine to testing.`

**Anchored on:** "spec — APISpec interface, adapters (Figure 4.3, component diagram)"  → **figures/component-diagram.tex (engine→testing edge exists at line 47; header notes a round-1 redraw for clipping); input still wrapped in \resizebox at 04-cli-architecture.tex:62**

**Meaning:** Improve the diagram: some arrows render only partially (clipped), and the expected engine→testing arrow appears missing. Implies a rendering/layout pass on the component diagram.

**⚠ Triage correction:** The figure header says it was redrawn after round-1 review specifically for clipping, and the engine→testing edge is in the source — but the chapter still applies \resizebox{\textwidth} (contradicting the figure's 'no resize' note), so a visual check of the compiled PDF is still warranted.

### ✅ #181 (p73) — already-addressed
**Kritikos:** `Here now P009 appears ...In addition, in previous image, we have P001-P008 principles while here we see a subset of them ...`

**Anchored on:** "principles P001–P005, P008, P009 checkers / spec APISpec interface, adapters (Figure 4.3)"  → **figures/dependency-flow.tex:27 and figures/component-diagram.tex:32 (both now 'P001--P005, P008, P009'); table at 04-cli-architecture.tex:81 matches**

**Meaning:** Direct figure-vs-figure contradiction: Fig 4.2 said P001–P008 while Fig 4.3 shows P001–P005, P008, P009. Implies making both figures (and the table) state the identical registered set.

### 🔴 #182 (p73) — open
**Kritikos:** `Please explain in a footnote what is this registry`

**Anchored on:** "Principle checkers (P001–P008), registry (Table 4.1 row)"  → **04-cli-architecture.tex:81**

**Meaning:** Explicit request: explain in a footnote what the registry is. Implies adding a footnote to the principles row of the package table (or at first textual mention).

### 🔴 #183 (p73) — open
**Kritikos:** `not clear what is loaded, selected, evaluated and aggregated ...`

**Anchored on:** "Orchestrator: load, select, evaluate, aggregate (Table 4.1 engine row)"  → **04-cli-architecture.tex:83**

**Meaning:** The four verbs are opaque: WHAT is loaded, selected, evaluated, aggregated? Implies expanding the cell or footnoting/cross-referencing the four-step pipeline of Section 4.6.2.

### 🔴 #184 (p73) — open
**Kritikos:** `We do not see the last three packages in Figure 4.3`

**Anchored on:** "util — Preprocessing utilities (Table 4.1 rows util/github/logger)"  → **figures/component-diagram.tex (still only 8 nodes) / 04-cli-architecture.tex:86-88**

**Meaning:** The last three table packages (util, github, logger) do not appear in Figure 4.3. Implies adding them to the component diagram or stating they are auxiliary/cross-cutting packages intentionally omitted.

### 🔴 #185 (p73) — open
**Kritikos:** `Report package is not mentioned in Figure 4.2.`

**Anchored on:** "engine (node in Figure 4.3) — 'Report package is not mentioned in Figure 4.2'"  → **figures/dependency-flow.tex (still no report node) / 04-cli-architecture.tex:47,53-58**

**Meaning:** The report package appears in Fig 4.3 and the table but is absent from Fig 4.2 (dependency flow). Implies adding report to Fig 4.2 or explaining the figures' differing scope.

### 🔴 #186 (p74) — open
**Kritikos:** `you forgot to mention exit code 3. Maybe say that: "while it also enables to distinguish between normal validation failures and tool usage failures (user mistakes)".`

**Anchored on:** "distinguish between validation failures (exit 1, actionable) and execution errors (exit 2, ...)"  → **04-cli-architecture.tex:127**

**Meaning:** The prose omits exit code 3 although the listing defines it. He even supplies wording: also say it distinguishes normal validation failures from tool-usage failures (user mistakes).

### 🔴 #187 (p75) — open
**Kritikos:** `ok but if we have a critical failure, does it make sense to move to functional or performance testing?`

**Anchored on:** "when a critical principle fails, the impact assessment indicates whether functional or performance tests can proceed"  → **04-cli-architecture.tex:163**

**Meaning:** If a CRITICAL principle fails, does it make sense to proceed to functional/performance testing at all? Implies clarifying the gating semantics: are downstream tests skipped/blocked on critical failure, and what exactly TestImpact records.

### 🔴 #188 (p75) — open
**Kritikos:** `In this respect,`

**Anchored on:** "self-documenting. A consumer of the JSON report can understand what was"  → **04-cli-architecture.tex:165**

**Meaning:** A connective suggestion: begin the sentence 'In this respect, a consumer of the JSON report can understand...'. Pure flow improvement.

**⚠ Triage correction:** Triage attached the connective to the following sentence ('The Principle struct itself carries metadata...'); the icon sits at 'self-documenting. A consumer...', so the connective belongs before 'A consumer of the JSON report'.

### 🔴 #189 (p76) — open
**Kritikos:** `that was used in the tool implementation?`

**Anchored on:** "they never interact with the raw kin-openapi [31] types (*openapi3.T,"  → **04-cli-architecture.tex:179**

**Meaning:** 'that was used in the tool implementation?' — make explicit at first mention that kin-openapi is the third-party parser library the tool implementation is built on.

### 🔴 #190 (p76) — open
**Kritikos:** `What is exactly covered by the Type? Does it relate to the kin-openapi types? Not clear.`

**Anchored on:** "and a structural validation hook (ValidateStructure). The Type and RawVersion methods"  → **04-cli-architecture.tex:193**

**Meaning:** 'What exactly is covered by Type? Does it relate to the kin-openapi types?' — disambiguate that Type() returns DriveBy's own SpecType discriminator (OpenAPI 3.x vs Swagger 2.0), independent of kin-openapi's types.

### 🔴 #191 (p76) — open
**Kritikos:** `ok but how does it return them? For instance, the Schema type is related to any field within the APISpec interface?`

**Anchored on:** "The APISpec interface returns normalized types that abstract away version-specific representation differences"  → **04-cli-architecture.tex:202**

**Meaning:** 'How does it return them? Is the Schema type related to any field within the APISpec interface?' — connect the normalized types (esp. Schema) to the interface methods that reach them (Components().Schemas, Paths() → Operation parameters/bodies).

### 🔴 #192 (p78) — open
**Kritikos:** `ok but you need to argue if it is complete or not due to this restriction.`

**Anchored on:** "is implemented as a focused set of sanity checks rather than full schema validation"  → **04-cli-architecture.tex:259**

**Meaning:** He accepts the restriction but wants an explicit completeness argument: state what the Swagger 2.0 sanity checks do and do not cover, and conclude whether validation is complete or knowingly partial for that path.

### 🔴 #193 (p81) — open
**Kritikos:** `when you go from 2.X to 3.0, do you downgrade or upgrade?`

**Anchored on:** "downgrades to 3.0.x via openapi2conv (Figure 4.4 class diagram, swagger2Spec node)"  → **figures/class-diagram.tex:76 (input at 04-cli-architecture.tex:338)**

**Meaning:** 'When you go from 2.x to 3.0, do you downgrade or upgrade?' — calling the 2.0→3.0 conversion a 'downgrade' is backwards; the label should say upgrades/converts.

### 🔴 #194 (p85) — open
**Kritikos:** `The captions should be in the same page as their actual graphical elements (figures, listings, tables, etc)`

**Anchored on:** "Listing 4.8: Mode-based checker selection (principles/registry.go)."  → **04-cli-architecture.tex:430 (lst:formode) and chapter-wide float/listing placement**

**Meaning:** Typesetting rule: captions must land on the same page as their figure/listing/table body. Implies a layout pass (keep listings with captions, \needspace, float placement) across the chapter.

**⚠ Triage correction:** Cannot be verified from source alone — requires a fresh PDF build after today's restructure; page breaks have certainly moved, so check before treating as fixed.

### 🔴 #195 (p85) — open
**Kritikos:** `what is combined mode selections? I have the feeling that you mean modes mapping to multiple principles?Or you mean cases where multiple modes are given? But the latter is not supported by the Registry as we can see in Listing 4.8.`

**Anchored on:** "they will be added to the registry and included in the test-only and combined mode selections"  → **04-cli-architecture.tex:461**

**Meaning:** 'Combined mode selections' is undefined; Listing 4.8's switch takes a single mode, so combining modes is unsupported. Implies deleting or precisely defining the phrase.

### 🔴 #196 (p85) — open
**Kritikos:** `ok but in a previous figure, we saw that the CLI could separately execute the functional and performance testing functionality which maps to validation operation.So, is what you say here consistent?`

**Anchored on:** "It is the single entry point for all validation operations."  → **04-cli-architecture.tex:473**

**Meaning:** Inconsistency with earlier figures/table: the CLI can run functional and load testing directly (function-only, load-only), bypassing the engine. Implies scoping the claim (e.g., 'all static/principle-based validation') or explaining the bypass.

### 🔴 #197 (p85) — open
**Kritikos:** `You say things that we do not see somewhere ...`

**Anchored on:** "The New constructor validates the configuration (requiring a spec path,"  → **04-cli-architecture.tex:483**

**Meaning:** 'You say things that we do not see somewhere' — the prose asserts New-constructor behaviour with no listing or code anchor. Implies adding a short listing of New or at least a source-file citation.

### ✅ #198 (p85) — already-addressed
**Kritikos:** `Shouldn't this be within a specific listing?`

**Anchored on:** "type Engine struct { config types.ValidatorConfig ..."  → **04-cli-architecture.tex:475-481 (now lstlisting lst:engine-struct with caption)**

**Meaning:** 'Shouldn't this be within a specific listing?' — the Engine struct sat in a bare verbatim block; convert to a captioned, labelled lstlisting like the other excerpts.

### 🔴 #199 (p87) — open
**Kritikos:** `In this case.`

**Anchored on:** "backward compatibility. The CLI commands use [Validate]"  → **04-cli-architecture.tex:527**

**Meaning:** Connective suggestion at the sentence boundary: open the next sentence with 'In this case, the CLI commands use Validate...'. Minor flow edit; intent low-stakes but plausible.

### 🔴 #200 (p87) — open
**Kritikos:** `with respect to what?`

**Anchored on:** "with logging output for backward compatibility"  → **04-cli-architecture.tex:527**

**Meaning:** 'With respect to what?' — backward compatibility with which prior interface/consumer? Implies naming the legacy caller or dropping the phrase if none exists.

### 🔴 #201 (p87) — open
**Kritikos:** `does it also comply with any kind of value restriction in general?`

**Anchored on:** "or type-based defaults (e.g., UUID format generates a valid UUID,"  → **04-cli-architecture.tex:544**

**Meaning:** Do generated parameter values also comply with general value restrictions (enum, min/max, pattern, length)? Implies stating whether the generator honours schema constraints or acknowledging the limitation.

### 🔴 #202 (p87) — open
**Kritikos:** `where does the user supply credentials? It is not clear ...`

**Anchored on:** "Adds authentication headers when the user provides credentials"  → **04-cli-architecture.tex:546**

**Meaning:** 'Where does the user supply credentials?' — the mechanism is unstated here. Implies cross-referencing the root-command auth flags (--auth-token etc.) documented in Section 4.9.1 (live line 622).

### 🟡 #203 (p87) — needs-author-decision
**Kritikos:** `ok but it neglects the actual content which makes the functional testing incomplete.`

**Anchored on:** "the response body conforms to the declared schema (required fields, type checks)"  → **04-cli-architecture.tex:548**

**Meaning:** He argues that ignoring the actual response CONTENT/values makes functional testing incomplete. The author must either add an explicit scope statement (contract conformance, not semantic correctness — deliberate boundary) or concede the limitation.

### 🔴 #204 (p88) — open
**Kritikos:** `like GET, HEAD, OPTIONS ones?`

**Anchored on:** "from non-destructive endpoints, runs a configurable load profile"  → **04-cli-architecture.tex:560**

**Meaning:** 'Like GET, HEAD, OPTIONS ones?' — define 'non-destructive endpoints' by naming the safe/idempotent methods in a parenthetical.

### 🔴 #205 (p88) — open
**Kritikos:** `which are supplied where? In the OpenAPI specification? Or in the quality gates derived from XSDLC?`

**Anchored on:** "Results are compared against user-defined thresholds, producing binary pass/fail outcomes"  → **04-cli-architecture.tex:560**

**Meaning:** Where are the thresholds supplied — in the OpenAPI specification or in the XSDLC quality gates? Implies stating the source explicitly (CLI flags / XSDLC loadTestConfig, NOT the spec); ties to the single-artifact debate.

**⚠ Triage correction:** Note: the single-artifact concession was made elsewhere today (Ch.2/abstract), but this sentence is unchanged — the local clarification is still missing.

### 🔴 #206 (p88) — open
**Kritikos:** `how to configure it and what does it usually include?`

**Anchored on:** "runs a configurable load profile—vegeta's"  → **04-cli-architecture.tex:560**

**Meaning:** 'How to configure it and what does it usually include?' — enumerate the load-profile knobs (rate, duration, concurrency, p95/success thresholds) and the configuration mechanism. The sentence already names rate/duration/concurrency; the HOW (flags / XSDLC loadTestConfig) is still missing.

### ✅ #207 (p88) — already-addressed
**Kritikos:** `please provide an example in parenthesis as done for 1.`

**Anchored on:** "A timestamped Markdown file with the same content rendered as"  → **04-cli-architecture.tex:574 (now includes 'e.g., validation-report-20260319T120000Z.md')**

**Meaning:** Provide a filename example in parentheses for item 3, matching item 1's style.

### 🔴 #208 (p89) — open
**Kritikos:** `correct. But it could also enable to correct source code, if the flaw concerns it and not the API specification itself (as it could be the case in performance testing).`

**Anchored on:** "it is an instruction set for autonomous API specification improvement"  → **04-cli-architecture.tex:586**

**Meaning:** Correct, but the remediation target is not only the specification: behavioural failures (e.g., performance) require fixing the SOURCE CODE. Implies broadening the claim to spec-or-implementation remediation.

### 🔴 #209 (p90) — open
**Kritikos:** `what is Viper?`

**Anchored on:** "from CLI flags (via Viper for environment variable binding)"  → **04-cli-architecture.tex:620**

**Meaning:** 'What is Viper?' — gloss and cite the library at first use, as done for Cobra.

### 🔴 #210 (p91) — open
**Kritikos:** `what about github-status?`

**Anchored on:** "GitHub integration—the –github-comment flag enables DriveBy"  → **04-cli-architecture.tex:657**

**Meaning:** 'What about github-status?' — the integration summary mentions only --github-comment, but the command table also defines github-status (commit status on a SHA). Implies adding it to the summary.

### ❓ #211 (p92) — unclear
**Kritikos:** `a`

**Anchored on:** "full potential is realized when embedded within declarative infrastructure"  → **04-cli-architecture.tex:668**

**Meaning:** Single-character comment 'a'. Two candidate readings: (1) insert the article — 'embedded within A declarative infrastructure' (the icon sits right at 'declarative infrastructure', making this the likelier reading); (2) a stray/accidental annotation with no content. Verify against the PDF before editing.

**⚠ Triage correction:** Triage dismissed it as a pure OCR artifact; given the icon position directly on 'within declarative infrastructure', the insert-article-'a' reading is at least as plausible and should be checked rather than dropped.


## 05-kubernetes-architecture.tex — 45 comments (32 open)

### ✅ #212 (p93) — already-addressed
**Kritikos:** `almost as the architectural diagram (Figure 4.1) covered almost the whole system.In any case, I propose to create a new Chapter 4 to cover the Overall Systems and its context and components.`

**Anchored on:** "Chapter 4 described the DriveBy CLI as a standalone validation binary."  → **05-kubernetes-architecture.tex:4**

**Meaning:** Kritikos notes the old 'From CLI Tool to Custom Resource' section overlapped with Figure 4.1 and proposes creating a new Chapter 4 covering the overall system, its context, and its components.

### ❓ #213 (p95) — unclear
**Kritikos:** `here`

**Anchored on:** "and Crossplane's reconciliation loop provisions everything."  → **04-system-architecture.tex:107**

**Meaning:** One-word marker 'here' on the final word of the Stage-5 paragraph. Candidate readings: (a) flags the operator(sing.)/'they declare'(pl.) agreement slip in that sentence, (b) marks a spot where something (a reference or detail) should be inserted. Intent cannot be decoded with confidence.

**⚠ Triage correction:** Triage anchored it to 'The operator does not configure individual components; they declare' and asserted the agreement-slip reading; the icon actually sits on the sentence-final word 'everything.', so that reading is only one candidate, not established. The paragraph also moved to the new system-architecture chapter.

### ✅ #214 (p96) — already-addressed
**Kritikos:** `, respectively`

**Anchored on:** "validation pipelines (Sections 5.8 and 5.9)."  → **GONE**

**Meaning:** Add ', respectively' so 'Argo Events and Workflows ... (Sections 5.8 and 5.9, respectively)' pairs each technology with its section. The entire chapter-roadmap sentence was deleted in today's restructure, so the fix is moot.

**⚠ Triage correction:** Triage anchored it to 'how the binary is invoked and where the results are consumed' (Stage-5 enabler paragraph); the icon actually sits on '(Sections 5.8 and 5.9)' in the chapter-roadmap sentence.

### ✅ #215 (p97) — already-addressed
**Kritikos:** `Correct but that figure requires some updating ...`

**Anchored on:** "it is not a context diagram in the strict SSADM sense (the true context diagram,"  → **05-kubernetes-architecture.tex:32**

**Meaning:** He accepts the prose distinction but says Figure 4.1 (the true SSADM context diagram, fig:system-context) itself needs updating. figures/system-context.tex was substantially rewritten today (expanded system boundary, substrate layer) for the new system-architecture chapter.

### 🔴 #216 (p98) — open
**Kritikos:** `The figure is nice but still has issues, mostly wrt the arrows again.Sometimes they have only their ending and not the respective line.In other cases, it is not so clear from where the arrow starts and ends. Finally, potentially some arrows are hidden.`

**Anchored on:** "perfect-api-prod (label inside Figure 5.1)"  → **figures/k8s-system-context.tex:82-110 (input at 05-kubernetes-architecture.tex:37)**

**Meaning:** Figure 5.1 (k8s-system-context) has arrow defects: some arrows show only the arrowhead without the line, start/end points are unclear, and some arrows may be hidden behind nodes. Fix the TikZ draw commands. The figure file is unmodified since the commented PDF.

### 🟡 #217 (p98) — needs-author-decision
**Kritikos:** `What is perfect-api?Maybe remove perfect prefix to make the naming independent of a specific API?`

**Anchored on:** "driveby ⋆ (namespace labels perfect-api-dev/-staging/-prod in Figure 5.1)"  → **figures/k8s-system-context.tex:85-87 (footnote at 05-kubernetes-architecture.tex:200)**

**Meaning:** Asks what perfect-api is and suggests dropping the 'perfect' prefix in the figure so the naming is independent of a specific API. The chapter text now carries a footnote (line 200) explaining the historical name, but the figure labels (k8s-system-context.tex:85-87) still read perfect-api-*; the author's recorded decision is that source identifiers keep the name. Author must decide whether to genericize the figure labels or extend the footnote to the figure.

### 🔴 #218 (p98) — open
**Kritikos:** `The promotion is done directly from one environment to the other or does it pass from the GitOps Promoter?In addition, it is dashed here while it is not dashed in the next case.`

**Anchored on:** "promote / promote / validate (arrows in Figure 5.1)"  → **figures/k8s-system-context.tex:104-105**

**Meaning:** Asks whether promotion goes directly between environments or through the GitOps Promoter (the text says Promoter-mediated, but the figure draws direct dev→staging→prod arrows), and notes the promote arrow is dashed here but not dashed in the comparable figure elsewhere. Figure still has direct dashed arrows (lines 104-105) not routed via the Promoter node.

### ✅ #219 (p99) — already-addressed
**Kritikos:** `actions`

**Anchored on:** "the controller takes corrective action to"  → **05-kubernetes-architecture.tex:55**

**Meaning:** Make it plural: 'takes corrective actions'. Today's edit changed line 55 to 'takes corrective actions'.

### 🔴 #220 (p100) — open
**Kritikos:** `Thus,`

**Anchored on:** "creating deleted resources or reverting drifted configurations is the correction"  → **05-kubernetes-architecture.tex:70**

**Meaning:** Insert 'Thus,' as a concluding connective. The icon sits inside the ontological-reconciliation paragraph (the 'The act of re-creating ... is the correction mechanism' / 'The three components ... are concretely instantiated' sentences), so the connective most plausibly belongs before the concluding 'The three components...' sentence of that paragraph.

**⚠ Triage correction:** Triage anchored it to 'This eventual consistency model enables self-healing' one paragraph earlier, and today's edit added 'Thus,' there (line 68). The icon actually sits in the ontological-reconciliation paragraph (line 70), so the applied 'Thus,' may be at the wrong sentence; author should verify and possibly move it.

### ✅ #221 (p100) — already-addressed
**Kritikos:** `concretely instantiated`

**Anchored on:** "are instantiated concretely in the Kubernetes controller model."  → **05-kubernetes-architecture.tex:70**

**Meaning:** Reorder to 'are concretely instantiated'. Live line 70 now reads 'are concretely instantiated in the Kubernetes controller model.'

### 🔴 #222 (p102) — open
**Kritikos:** `Figure is almost ok. Again, there is the issue with the arrows, this time internally the two provider rectangles.`

**Anchored on:** "In-Cluster Resources (box inside Figure 5.2)"  → **figures/provider-architecture.tex (input at 05-kubernetes-architecture.tex:101)**

**Meaning:** Figure 5.2 (provider-architecture) is almost OK, but the arrows inside the two provider rectangles (provider-kubernetes and provider-upjet-github boxes) are incomplete/unclear and must be fixed. Figure file unmodified since the commented PDF.

### 🔴 #223 (p103) — open
**Kritikos:** `SDLC or XSDLC? If the former holds, SDLC acronym was never introduced before.`

**Anchored on:** "one custom resource defines the entire SDLC pipeline for an application"  → **05-kubernetes-architecture.tex:112**

**Meaning:** Asks: SDLC or XSDLC? If bare 'SDLC' is intended, the acronym was never introduced. Live text still uses bare 'SDLC pipeline' (lines 112, 200, 655) and only XSDLC is expanded (acronym list, abstract); 'SDLC' as such is never defined in prose. Add the expansion 'software development lifecycle (SDLC)' at first use or rephrase.

### 🔴 #224 (p103) — open
**Kritikos:** `actual ... where the software project's main source code is situated is never referenced ...`

**Anchored on:** "The software repository is never referenced"  → **05-kubernetes-architecture.tex:119**

**Meaning:** He highlights that the repository holding the application's actual main source code is never referenced by XSDLC; he wants the text to make explicit what 'software repository' means (the source-code repo) and that its exclusion is deliberate. Sentence is unchanged in the live source.

### ✅ #225 (p103) — already-addressed
**Kritikos:** `(CR)`

**Anchored on:** "A single XSDLC custom resource generates three categories of resources."  → **05-kubernetes-architecture.tex:121**

**Meaning:** Introduce the acronym: 'custom resource (CR)', since 'CR' is used shortly after ('once per CR'). The acronym is now expanded at first use in Chapter 1 ('a single Custom Resource (CR)', 01-introduction.tex:158), satisfying the define-on-first-use convention.

### 🔴 #226 (p103) — open
**Kritikos:** `Could explain what is this in a footnote`

**Anchored on:** "An ArgoCD push secret granting the Source Hydrator write access"  → **05-kubernetes-architecture.tex:130**

**Meaning:** Add a footnote explaining what an ArgoCD push secret / the Source Hydrator is and why write access to the gitops repo is needed. No footnote exists; the Source Hydrator is only explained much later (Section on ArgoCD Application Generation, lines 390-398). Add a footnote or a forward reference.

### 🔴 #227 (p104) — open
**Kritikos:** `maybe say that it exploits the static validation facilities of the DriveBy CLI focusing on the specification-level principles, which are fully applied or not depending on the validation mode (strict means full application, minimal means partial application)`

**Anchored on:** "runs the DriveBy CLI in static validation mode,"  → **05-kubernetes-architecture.tex:166**

**Meaning:** Suggests phrasing that validate-only exploits the DriveBy CLI's static validation facilities for the specification-level principles, which are fully or partially applied depending on validation mode (strict = full application, minimal = partial). The live bullet still has the terser original wording he commented on.

### 🔴 #228 (p105) — open
**Kritikos:** `could explain whether the default values are derived from common practice (or whatever other reason actually exists)`

**Anchored on:** "concurrentUsers (default: 10)"  → **05-kubernetes-architecture.tex:168**

**Meaning:** Explain where the load-test default values come from (common practice, experimentation, or another rationale). No justification has been added.

### 🔴 #229 (p105) — open
**Kritikos:** `I presume that this is an indicative distribution of checks to potential gates, right?The actual distribution depends on the user and his/her requirements`

**Anchored on:** "Different environment transitions compose different subsets."  → **05-kubernetes-architecture.tex:173**

**Meaning:** He presumes the staging/production gate examples are an indicative distribution of checks to gates, with the actual distribution chosen by the user per requirements; he wants this stated explicitly. No 'illustrative/example' qualifier has been added.

### 🔴 #230 (p105) — open
**Kritikos:** `Need to explain the semantics of the different parameters that can be configured, although they might be obvious to you.`

**Anchored on:** "concurrentUsers (default: 10), testDuration (default: 5m), maxLatencyP95"  → **05-kubernetes-architecture.tex:168**

**Meaning:** Explain the semantics of each configurable load-test parameter (what concurrentUsers, testDuration, maxLatencyP95, minSuccessRate actually mean), not just their defaults. Not done.

### 🔴 #231 (p106) — open
**Kritikos:** `could explain why validate-only is executed twice in the two gates. Is this normal practice?`

**Anchored on:** "- type: validate-only (in Listing 5.1, appearing in both gates)"  → **05-kubernetes-architecture.tex:217 (and 223)**

**Meaning:** Asks why validate-only is executed twice (once per gate) and whether re-running static validation at every gate is normal practice. No explanation added near the listing or the fail-fast paragraph.

### 🔴 #232 (p106) — open
**Kritikos:** `ok but you also have environment resources that are not mentioned here!`

**Anchored on:** "and the per-gate resources (...) for each gated environment."  → **05-kubernetes-architecture.tex:227**

**Meaning:** The 'From this declaration, Crossplane generates...' sentence lists shared and per-gate resources but omits the per-environment resources (PromotionStrategy, Branches, overlay RepositoryFile, ghcr-creds, ArgoCD Application) enumerated earlier in 5.4.1. Still omitted in the live sentence.

### ✅ #233 (p106) — already-addressed
**Kritikos:** `In Figure 5.3 we see Role & RoleBinding while here you mention RBAC`

**Anchored on:** "shared resources (ServiceAccount, EventBus, ScmProvider, GitRepository, PromotionStrategy,"  → **05-kubernetes-architecture.tex:227**

**Meaning:** Inconsistency: the text said 'RBAC' while Figure 5.3 shows 'Role' and 'RoleBinding'. Today's edits replaced 'RBAC' with 'Role, RoleBinding' at line 227 (and in the figure caption and the install-order list), aligning text and figure.

### 🔴 #234 (p107) — open
**Kritikos:** `Again please try to fix the arrows to appear completely.`

**Anchored on:** "XSDLC (single CR per application) / App-level (shared) / Per-gate (one set each) (Figure 5.3)"  → **figures/xrd-resource-hierarchy.tex (input at 05-kubernetes-architecture.tex:234)**

**Meaning:** Fix the arrows in the xrd-resource-hierarchy figure so they render completely (full lines, visible endpoints). Figure file unmodified since the commented PDF.

### 🔴 #235 (p107) — open
**Kritikos:** `Do we also have environment resources that are created? We do not see them in this figure ...`

**Anchored on:** "XSDLC (root node of Figure 5.3)"  → **figures/xrd-resource-hierarchy.tex:44-68 (caption at 05-kubernetes-architecture.tex:236)**

**Meaning:** Asks whether per-environment resources are also created, since they are not visible in Figure 5.3. The caption (line 236) lists a per-environment group, but the figure itself still has only 'App-level (shared)' and 'Per-gate' groups (PromotionStrategy sits in the app-level column), so figure and caption disagree. Add a per-environment group to the figure.

### 🔴 #236 (p108) — open
**Kritikos:** `in the context of your system, right? This needs to be clarified. In addition, is there a URL for this chart? It could be supplied here unless it was already provided earlier.`

**Anchored on:** "The Helm chart installs the platform layer:"  → **05-kubernetes-architecture.tex:254**

**Meaning:** Clarify that 'the Helm chart' is DriveBy's own chart (i.e., in the context of this system), and supply the chart's URL here unless given earlier. The URL (oci://ghcr.io/meter-peter/charts/driveby) appears only later (line 644), not earlier; the sentence is unchanged.

### 🔴 #237 (p109) — open
**Kritikos:** `composition?XSDLC composite?`

**Anchored on:** "to mark the composite as ready. (caption of Figure 5.4)"  → **05-kubernetes-architecture.tex:274**

**Meaning:** Terminology question: 'composition? XSDLC composite?' — he wants the caption to say precisely which object is marked ready (the XSDLC composite resource, not the composition). Caption unchanged ('the composite').

### 🔴 #238 (p110) — open
**Kritikos:** `do you mean CR?As spec gives the impression that you change the meta-model and not a specific model / CR definition.`

**Anchored on:** "Modifying the XSDLC spec triggers a reconciliation delta."  → **05-kubernetes-architecture.tex:296**

**Meaning:** Asks whether 'spec' means the CR: as written, 'modifying the XSDLC spec' could read as changing the meta-model (the XRD) rather than a specific CR's spec field. Rephrase to e.g. 'Modifying an XSDLC CR's spec'. Unchanged.

### 🔴 #239 (p110) — open
**Kritikos:** `+ RepositoryFile based on 5.5.3 (last paragraph)`

**Anchored on:** "deletionPolicy: Orphan (ServiceAccount, EventBus, Branches) remain."  → **05-kubernetes-architecture.tex:298**

**Meaning:** The Orphan list in the Deletion paragraph should also include RepositoryFile, because Section 5.5.3's last paragraph states that both Branch and RepositoryFile resources use deletionPolicy: Orphan. Live line 298 still lists only (ServiceAccount, EventBus, Branches).

**⚠ Triage correction:** Triage anchored it to 'Adding a new environment generates the new ArgoCD Application, branches...' in the Updates paragraph; the icon actually sits on the Deletion paragraph's Orphan list, and the requested change is adding RepositoryFile to that list.

### ✅ #240 (p110) — already-addressed
**Kritikos:** `across all gated-environments?`

**Anchored on:** "gates[] (array): per-gate status including environment name, gate"  → **05-kubernetes-architecture.tex:311**

**Meaning:** Asks whether the gates[] array spans all gated environments. Today's edit reworded it to 'one entry per gated environment, reporting environment name, ...'.

### 🔴 #241 (p111) — open
**Kritikos:** `Some comments for improving the figure:(a) the dashed arrows should not end inside the rectangles but at their boundary. (b) the first dashed arrow should not overlap with the label below the respective rectangle(c) e.g. Deleted Sensor -> re-created to indicate that this is an example.`

**Anchored on:** "Day N Update / Delete (Figure 5.5)"  → **figures/declarative-lifecycle.tex (input at 05-kubernetes-architecture.tex:330)**

**Meaning:** Figure improvements: (a) dashed arrows should end at rectangle boundaries, not inside; (b) the first dashed arrow should not overlap the label below its rectangle; (c) prefix 'e.g.' to 'Deleted Sensor → re-created' to mark it as an example. Figure file unmodified since the commented PDF.

### 🔴 #242 (p113) — open
**Kritikos:** `could explain why although this is obvious`

**Anchored on:** "uses deletionPolicy: Orphan to preserve the repository when the XSDLC CR is deleted"  → **05-kubernetes-architecture.tex:379**

**Meaning:** Explain why the gitops repository is orphaned rather than deleted, even though it seems obvious (e.g., deleting it would destroy manifest history and audit trail for environments that may still be running). No rationale added.

**⚠ Triage correction:** Triage anchored it to the Branches paragraph ('Orphan to preserve Git history', line 377); the icon actually sits on the Repository paragraph ('preserve the repository when the XSDLC CR is deleted', line 379) — adjacent but distinct; the rationale request targets the repository's Orphan policy.

### 🔴 #243 (p116) — open
**Kritikos:** `I would propose to order the parameters based on whether they are global or resource-specific. In each group, you could apply alphabetical ordering`

**Anchored on:** "Table 5.2: Representative configuration parameters and their two-tier resolution"  → **05-kubernetes-architecture.tex:437-443**

**Meaning:** Reorder the table rows: group parameters into global (no per-resource override) vs resource-specific, and order alphabetically within each group. Table rows are still in the original mixed order.

### ✅ #244 (p116) — already-addressed
**Kritikos:** `support`

**Anchored on:** "The JetStream backing ensures that webhook events are not lost if the"  → **05-kubernetes-architecture.tex:460-462**

**Meaning:** One-word comment 'support'. The applied reading (today): change 'which provides persistent, at-least-once delivery semantics' to 'which supports ...' in the preceding sentence (line 460). Alternative reading suggested by the icon position: replace 'backing' with 'support' ('The JetStream support ensures...') in the anchored sentence (line 462). One plausible reading has been applied.

**⚠ Triage correction:** Triage anchored it to 'which provides persistent, at-least-once delivery semantics'; the icon actually sits at the start of the next sentence ('The JetStream backing ensures...'), so the intended target is ambiguous between the two adjacent sentences — the applied 'provides→supports' edit is plausible but not certain.

### 🔴 #245 (p117) — open
**Kritikos:** `What do you mean here?`

**Anchored on:** "github-owner / body.repository.owner.login / GitHub API calls (Table 5.3 row)"  → **05-kubernetes-architecture.tex:504-505**

**Meaning:** 'What do you mean here?' — the Purpose entry 'GitHub API calls' for github-owner/github-repo is too vague; spell out what these parameters are used for (e.g., addressing the repo when posting commit statuses and PR comments). Table unchanged.

### 🔴 #246 (p119) — open
**Kritikos:** `Is Authentication also required here?`

**Anchored on:** "load-test → driveby-loadtest: Runs the DriveBy CLI in load-only mode with vegeta-based load testing"  → **05-kubernetes-architecture.tex:547**

**Meaning:** Asks whether authentication credentials are also required/injected for the load-test step, as stated for functional-test in the bullet above. The load-test bullet still says nothing about authentication; the author must state the factual answer.

### ✅ #247 (p119) — already-addressed
**Kritikos:** `The first two steps have different names than those declared above.`

**Anchored on:** "set-pending → health-check → check-0-validate-only (example DAG)"  → **05-kubernetes-architecture.tex:561**

**Meaning:** The first two steps in the example DAG had different names from those declared in the Pipeline Steps subsection (set-pending-status, wait-for-source-ready). Today's edit fixed the verbatim block to use the declared names.

### 🔴 #248 (p119) — open
**Kritikos:** `could explain that -> mean sequential execution and / parallel. Just to be as typical as possible.`

**Anchored on:** "check-1-functional-test → report-success / comment-pr /"  → **05-kubernetes-architecture.tex:558-564**

**Meaning:** Add a legend for the DAG notation: '→' means sequential execution and '/' means parallel execution, to be as conventional as possible. No legend has been added near the verbatim block.

### 🔴 #249 (p119) — open
**Kritikos:** `I am puzzled slightly here as I had the impression that in a previous table, a default port could be modified in a CR definition. Please check this and make corrections if needed.`

**Anchored on:** "These are baked into the composition at chart installation and cannot change per CR."  → **05-kubernetes-architecture.tex:574 (vs table row at 05-kubernetes-architecture.tex:440)**

**Meaning:** Apparent contradiction: this sentence says default ports cannot change per CR, but Table 5.2 lists the webhook port as overridable per-resource (Tier 2). He asks to check and correct. Both passages are unchanged; the contradiction remains.

### 🔴 #250 (p120) — open
**Kritikos:** `As this was already explained in Section 5.6.2, could make a simple reference to that section indicating that the mechanism is explained there.This can save some space and avoid the repetition.`

**Anchored on:** "The composition solves this with triple escaping"  → **05-kubernetes-architecture.tex:580**

**Meaning:** The escaping mechanism was already explained in the Template Escaping subsection (5.6.2); replace the repeated explanation with a short reference to that section to save space. The full triple-escaping explanation remains at line 580, duplicating lines 360-368 (which call it a 'double-escaping convention' — also an internal naming inconsistency).

### 🔴 #251 (p121) — open
**Kritikos:** `is this a default or an effective value taken from a specific configuration?`

**Anchored on:** "(concurrent-users, test-duration, max-latency-p95, min-success-rate), each with a default"  → **05-kubernetes-architecture.tex:599**

**Meaning:** Asks whether the cited values (10 users, 5 minutes, 500 ms P95, 99% success rate) are defaults or effective values taken from a specific configuration; clarify their provenance (and consistency with the defaults listed in 5.4.2). Sentence unchanged.

**⚠ Triage correction:** Triage anchored it to 'a gate with one validate-only check produces a 5-step DAG' and read it as a step-count question; the icon actually sits on the per-check parameter-defaults sentence ('test-duration, max-latency-p95, min-success-rate, each with a default').

### 🔴 #252 (p121) — open
**Kritikos:** `I have the impression that we have 6 steps in this case: 2 initial, 3 last and the validation one.`

**Anchored on:** "a gate with one validate-only check produces a 5-step DAG"  → **05-kubernetes-architecture.tex:601**

**Meaning:** He counts 6 steps for the single-check gate (2 preamble + 1 check + 3 postamble), not 5. The live text still says '5-step DAG'; the counts need recounting against the actual WorkflowTemplate using one consistent counting convention.

### 🔴 #253 (p121) — open
**Kritikos:** `This should be 8. See previous comment.`

**Anchored on:** "produces a 7-step DAG with three different step template definitions"  → **05-kubernetes-architecture.tex:601**

**Meaning:** Companion to the previous comment: by his accounting the three-check gate yields 8 steps (2 preamble + 3 checks + 3 postamble), not 7. Still says '7-step DAG'.

### 🔴 #254 (p122) — open
**Kritikos:** `The correct or effective sequence?I am asking this as the chart should impose a specific order, right? If not, then what is the solution to apply by an operator? To apply the chart in waves/phases?In addition, I presume that step 4 is out of scope in terms of system installation. The chart handles the system installation while applications/environments are installed via CRs afterwards. I would thus make this distinction clear to the reader.`

**Anchored on:** "the XRD must be installed before the composition that implements it. The recommended sequence is:"  → **05-kubernetes-architecture.tex:626-635**

**Meaning:** Asks whether the four-step order is enforced by the chart or must be applied manually in waves/phases, and notes step 4 (applying the XSDLC CR) is outside system installation — the chart installs the system, applications/environments come later via CRs — and wants this distinction made explicit. Line 635 partially notes step 4 is Crossplane-handled, but the chart-vs-manual-ordering question and the explicit scope distinction are still unaddressed.

### 🔴 #255 (p122) — open
**Kritikos:** `Can you provide a script or specific commands per phase? This would be ideal in the context of your work as you can facilitate its exploitation by potential adopters.`

**Anchored on:** "Pre-install: Crossplane provider and functions (wait for CRDs)"  → **05-kubernetes-architecture.tex:629-633**

**Meaning:** Provide a script or the specific commands for each installation phase, to facilitate adoption by potential users. No commands/listing added.

### ✅ #256 (p123) — already-addressed
**Kritikos:** `a ... a GO CLI, a ..., and a ... to a ...`

**Anchored on:** "tracing the evolution from shell script through Go CLI, container image, and"  → **05-kubernetes-architecture.tex:4 (summary successor at 05-kubernetes-architecture.tex:651)**

**Meaning:** Insert consistent indefinite articles: 'from a shell script through a Go CLI, a container image, and a workflow step to a Crossplane custom resource'. Today's edits rewrote the summary sentence (line 651) and the surviving intro sentence (line 4) now carries the full article series exactly as requested.


## 06-gitops-pipeline.tex — 27 comments (18 open)

### 🔴 #257 (p125) — open
**Kritikos:** `in Chapter 7?`

**Anchored on:** "The answer is demonstrated through a proof-of-concept deployment"  → **06-gitops-pipeline.tex:8**

**Meaning:** On the RQ3 sentence claiming 'the answer is demonstrated through a proof-of-concept deployment', Kritikos asks 'in Chapter 7?' — i.e., isn't the demonstration/answer actually delivered in the evaluation chapter? He wants an explicit cross-reference to the evaluation chapter (Chapter~\ref{ch:evaluation}) added to this sentence.

**⚠ Triage correction:** Triage anchored this to the chapter-intro 'Chapter~\ref{ch:kubernetes} described...' sentence and read it as a cross-ref-number sanity check. Wrong anchor: the icon sits on the RQ3 'answer is demonstrated through' sentence; he wants a forward reference to the evaluation chapter, not a label audit.

### 🔴 #258 (p126) — open
**Kritikos:** `source ...`

**Anchored on:** "validation against the development environment---the environment where the new code is already running"  → **06-gitops-pipeline.tex:17**

**Meaning:** One-word note 'source ...' on the phrase 'validation against the development environment'. He is pointing back to the chapter's own slogan 'validate the source environment, not the target': he wants this sentence to say 'source environment' (or 'source (development) environment') for terminological consistency.

**⚠ Triage correction:** Triage anchored this to the Source Hydrator sentence and dismissed it as an OCR-truncated note. Wrong: the icon sits on 'validation against the development environment'; 'source ...' means substitute/add 'source environment' terminology there.

### ✅ #259 (p126) — already-addressed
**Kritikos:** `Comments for this figure:(a) Some vertical labels are "cut" and not represented fully(b) two arrows have wrong direction:(1) EventSource -> Traefik Ingress(2) ArgoWorkflow -> Sensor`

**Anchored on:** "BYOCI (vertical swim-lane label in Figure 6.1, fig:gitops-pipeline)"  → **figures/workflow-pipeline.tex:46-83 (input via 06-gitops-pipeline.tex:26)**

**Meaning:** Two figure defects in fig:gitops-pipeline: (a) vertical lane labels (BYOCI, Hydrator, GitHub, Ingress, Argo Events, Argo Workflows) are clipped; (b) two arrows render in the wrong direction — should be Traefik Ingress -> EventSource and Sensor -> Argo Workflow.

**⚠ Triage correction:** Earlier triage reading was right; status update only: commit 214ec71 ('close item [408]') moved the vertical swimlane labels inside the lanes, and the current source already draws traefik->es (line 75) and sensor->wf (line 83) in the correct direction. Needs one visual confirmation on the next build, but the source is fixed.

### 🔴 #260 (p127) — open
**Kritikos:** `what is domain here? Does it also map to the XSDLC spec?`

**Anchored on:** "https://{app}-{gate}-webhook.{domain}/{app}-{gate}-pr-trigger"  → **06-gitops-pipeline.tex:46**

**Meaning:** He asks what {domain} is in the URL template and whether it also maps to an XSDLC spec field like {app} and {gate} do. The text says the URL derives from 'two identity fields' but uses three template variables. Fix: state that {domain} is the platform ingress domain inherited from values.yaml (e.g., private.novelcore.org), not a per-app XSDLC field.

### 🔴 #261 (p128) — open
**Kritikos:** `and check type ...`

**Anchored on:** "depending on specification size and validation mode). In practice, a successful validation cycle"  → **06-gitops-pipeline.tex:64**

**Meaning:** Note 'and check type ...': validation execution time also depends on the configured check type (static vs functional vs load test), not only spec size and validation mode — the very next sentence's 10-12s vs 18-28s split proves it. Add 'and check type' to the factor list.

### 🔴 #262 (p131) — open
**Kritikos:** `please provide an example of a change at this level`

**Anchored on:** "cross-environment changes or a specific overlay (dry/overlays/<env>/) for environment-specific customization"  → **06-gitops-pipeline.tex:135**

**Meaning:** 'Please provide an example of a change at this level': he wants a concrete worked example of a dry-manifest edit (e.g., bumping an image tag in dry/overlays/staging/kustomization.yaml) so the base-vs-overlay integration contract is illustrated rather than only asserted.

### 🔴 #263 (p132) — open
**Kritikos:** `depending on the affected overlays.So, which environment is affected can actually vary. Maybe it is better to indicate here that Figure 6.2 covers a cross-environment updating scenario?`

**Anchored on:** "Hydrator renders per-environment overlays and writes hydrated output to -next branches (Figure 6.2 caption)"  → **06-gitops-pipeline.tex:151 (caption; lead-in at 144)**

**Meaning:** Which environments are affected depends on which overlays a commit touches, so the fan-out to all three -next branches shown in Figure 6.2 is not universal. He suggests stating in the caption/lead-in that the figure depicts a cross-environment (base-change) scenario, while single-overlay commits re-hydrate only the affected environment.

### 🔴 #264 (p132) — open
**Kritikos:** `The arrows again here (and outgoing arrow for SDT gate) are not shown completely.`

**Anchored on:** "SDT Gate (node in Figure 6.2, fig:single-repo-branch-flow)"  → **figures/single-repo-branch-flow.tex:77-87 (input via 06-gitops-pipeline.tex:149)**

**Meaning:** In the rendered Figure 6.2 some arrows — specifically the outgoing arrow from the SDT Gate node — are not shown completely (clipped). The TikZ source does draw gate1->staging (figure line 86), so this is a rendering/bounding-box problem to fix and verify visually.

### 🔴 #265 (p132) — open
**Kritikos:** `Could also indicate that the gate naming is indicative.In addition, maybe name the first Static Validation Gate. This could make more sense. As SDT gives the impression that all kinds of validation in your (SDT) work are applied.`

**Anchored on:** "quality gates control promotion; ArgoCD auto-syncs from active branches (Figure 6.2 caption)"  → **figures/single-repo-branch-flow.tex:74 + 06-gitops-pipeline.tex:151**

**Meaning:** Two suggestions: (1) note that the gate names in the figure are indicative/illustrative, not fixed; (2) rename the 'SDT Gate' node to 'Static Validation Gate', because 'SDT Gate' wrongly implies the full SDT principle set runs there when it is a validate-only gate. Same label should be checked in fig:multi-env-promotion and fig:human-in-the-loop.

### 🔴 #266 (p133) — open
**Kritikos:** `Previously you indicate that we go from x.next to x. While here we have x->y->z. This creates the natural question: how do you support the latter transition chain?In addition, if one change affects one environment, e.g., staging, does this mean that we start from staging and we can end up to production?Please clarify in the text.`

**Anchored on:** "Manual Approval (label in Figure 6.3, fig:multi-env-promotion)"  → **06-gitops-pipeline.tex:162-186 (fig:multi-env-promotion and PromotionStrategy discussion)**

**Meaning:** The chapter first describes per-environment x-next -> x promotion, but Figure 6.3 shows a linear dev -> staging -> prod chain. He asks how the x->y->z transition chain is actually supported, and whether a change affecting only staging can end up in production — and explicitly asks for clarification in the text. The reconciliation (each env hydrates independently from main; no cross-env propagation) exists at line 113 but not at Section 6.4 where the apparent contradiction arises.

### 🔴 #267 (p136) — open
**Kritikos:** `correct. When this mode would be used? For manual testing done by the developer/QA engineer?`

**Anchored on:** "it offers no advantage over the gated model. ArgoCD still auto-syncs once the branch is updated"  → **06-gitops-pipeline.tex:241**

**Meaning:** He agrees ('correct') but asks when the fully-manual-merge mode would actually be used — e.g., for manual testing by a developer/QA engineer? Add one sentence naming a plausible use case instead of leaving the mode purely vestigial.

### 🔴 #268 (p137) — open
**Kritikos:** `The arrows here could again become larger ...`

**Anchored on:** "Env2 (row label in Figure 6.4, fig:human-in-the-loop)"  → **figures/human-in-the-loop.tex:83-107 (input via 06-gitops-pipeline.tex:249)**

**Meaning:** The arrows in Figure 6.4 should be made larger/thicker — after the \textwidth resizebox they are hard to see. The source still uses thin '->,thick,black!40' arrows unchanged since the reviewed build.

### 🔴 #269 (p137) — open
**Kritikos:** `See previous comment`

**Anchored on:** "Auto-merge (cell in Figure 6.4, fig:human-in-the-loop)"  → **figures/human-in-the-loop.tex:83-107 (input via 06-gitops-pipeline.tex:249)**

**Meaning:** 'See previous comment' — duplicate marker for idx 268: enlarge/thicken the arrows in fig:human-in-the-loop. Resolve together with 268.

### 🔴 #270 (p137) — open
**Kritikos:** `See previous comment`

**Anchored on:** "Auto-sync (cell in Figure 6.4, fig:human-in-the-loop)"  → **figures/human-in-the-loop.tex:83-107 (input via 06-gitops-pipeline.tex:249)**

**Meaning:** 'See previous comment' — second duplicate marker for idx 268: same arrow-size fix in fig:human-in-the-loop. No separate action.

### ✅ #271 (p138) — already-addressed
**Kritikos:** `(PR)`

**Anchored on:** "Every environment transition is a GitHub pull request"  → **06-gitops-pipeline.tex:272**

**Meaning:** '(PR)' — introduce the acronym at this first prose use: 'GitHub pull request (PR)'. The live source now reads exactly that (uncommitted edit), so the change is in.

### ✅ #272 (p139) — already-addressed
**Kritikos:** `Maybe add "with observation only" or "without actuation". As there are not update actions to bring the system to the desired state ...It depends also on what is meant by "desired state" here.`

**Anchored on:** "forms a closed-loop negative feedback controller"  → **06-gitops-pipeline.tex:288 (with 338)**

**Meaning:** He suggests qualifying the closed-loop claim with 'with observation only' / 'without actuation', since the gate does not act to bring the system to the desired state; he also notes it depends on what 'desired state' means. The live rewrite answers this: the same opening paragraph now introduces the two-tier decomposition ('a verdict tier that is open-loop and a platform tier that performs continuous closed-loop repair'), and line 338 states explicitly that the verdict tier 'produces no actuation signal u that reduces e(t)'. Optionally the very first sentence could still carry the qualifier, but the substance is addressed.

### 🟡 #273 (p139) — needs-author-decision
**Kritikos:** `Ok but blocking is not a normal adaptation action. It just indicates that the live API does not satisfy the contract. It does not act to bring the live API close to the contract. IMPORTANT: if the plant is the delivery system, then your assumptions and mappings are correct. But you need to modify Table 6.1 to reflect this (e.g., plant: pipeline & promotion workflow, setpoint: allows transitions policy).Another alternative: plant: Promotion PR / release decision system and setpoint: A predicate over artifacts and state.`

**Anchored on:** "Actuator   CommitStatus + Promoter   Merges or blocks promotion PR (Table 6.1 row)"  → **06-gitops-pipeline.tex:297-316 (Plant row at 308)**

**Meaning:** Blocking is not a corrective adaptation action on the API. His IMPORTANT instruction: the mapping is only valid if the plant is the delivery system, so Table 6.1 must be modified (e.g., Plant = pipeline & promotion workflow, Setpoint = allowed-transitions policy; or Plant = promotion-PR/release-decision system, Setpoint = predicate over artifacts and state). The live source took a different route: Table 6.1 keeps Plant = 'API deployment (source env)' but is now explicitly framed as the verdict-tier view only, with a new platform-tier section and fig:controller-stack carrying the delivery-system plant. The author must decide whether to defend this two-tier framing or actually remap/split Table 6.1 as instructed.

### 🔴 #274 (p140) — open
**Kritikos:** `This text here is difficult to be seen/observed`

**Anchored on:** "(SDT validator) (annotation in Figure 6.5, fig:quality-control-loop)"  → **figures/quality-control-loop.tex:64 (input via 06-gitops-pipeline.tex:325)**

**Meaning:** The '(SDT validator)' label and surrounding feedback-path text in Figure 6.5 are too small to read ('difficult to be seen'). The source still renders it \tiny italic; enlarge or relocate it. Figure unchanged since the reviewed build.

### 🔴 #275 (p140) — open
**Kritikos:** `Here it seems that you have arrows that cannot be easily seen, if not at all.`

**Anchored on:** "API Deployment (source env) (plant block in Figure 6.5, fig:quality-control-loop)"  → **figures/quality-control-loop.tex (input via 06-gitops-pipeline.tex:325)**

**Meaning:** The block-diagram arrows in Figure 6.5 (setpoint -> summing junction -> controller -> actuator -> plant, plus the feedback path) are barely visible or invisible in the rendered PDF. Thicken arrows/arrowheads; figure source unchanged since the reviewed build.

### 🟡 #276 (p140) — needs-author-decision
**Kritikos:** `If you change the content of the previous table, the same should be done for the current figure.`

**Anchored on:** "Argo Workflow DAG (controller block in Figure 6.5, fig:quality-control-loop)"  → **figures/quality-control-loop.tex + 06-gitops-pipeline.tex:327 (caption) and 342 (fig:controller-stack)**

**Meaning:** Conditional follow-up to idx 273: if Table 6.1's plant/setpoint mapping changes, Figure 6.5 must be updated identically — table and figure must stay in lockstep. Live source already adds fig:controller-stack for the platform tier and recaptions Figure 6.5 as the verdict-tier view; whether Figure 6.5 itself changes follows directly from the idx-273 decision.

### ✅ #277 (p140) — already-addressed
**Kritikos:** `ok but it is out of scope of your system, so it is closed-loop with no actuation.`

**Anchored on:** "convergence is driven by the developer's corrective action, gated"  → **06-gitops-pipeline.tex:364 (with 338)**

**Meaning:** 'OK but it is out of scope of your system, so it is closed-loop with no actuation' — the corrective action lies outside the system boundary. The live text now concedes this exactly: line 338 ('no actuation signal u... corrective action is performed externally') and line 364 ('convergence is driven by external corrective action... the verdict tier converges over PRs').

### 🔴 #278 (p140) — open
**Kritikos:** `Ok, but I have two issues here:(a) if the goal is to conduct the validation on the current change, then this goal is not achieved. Shouldn't the system retry to execute the workflow in case of such a failure?(b) if the system waits the next branch change, this could take a long time, thus causing the develop to create a new commit in the GitOps repository. This is manual re-triggering ...`

**Anchored on:** "posts a "failure" status. On the next branch change, the Promoter triggers re-evaluation automatically"  → **06-gitops-pipeline.tex:366**

**Meaning:** Two objections: (a) if a workflow fails transiently, the validation of the CURRENT change never completed — shouldn't the system retry the workflow itself? (b) waiting for the 'next branch change' can take long and forces the developer to push a new commit, which IS manual re-triggering. Line 340 now mentions Argo Workflows' step retry policy (partially answering (a)), but the annotated sentence survives verbatim at line 366 and objection (b) — no automatic same-change re-evaluation — is still unanswered; the transient-infra-retry vs re-commit-after-real-failure distinction is not drawn.

### ✅ #279 (p141) — already-addressed
**Kritikos:** `ok but the system converges through two types of actions: (a) preventive blocking promotion and (b) manual corrective. So, it relies eventually on manual reconciliation actions so as to converge.Isn't this problematic in the context of closed-loop control theory?`

**Anchored on:** "why the system converges, why it self-heals, and why the axiomatic foundation is necessary"  → **06-gitops-pipeline.tex:379**

**Meaning:** Theoretical objection: convergence relies on two action types — preventive blocking plus manual corrective action — so the system ultimately depends on manual reconciliation; isn't that problematic for a closed-loop control claim? The live rewrite answers it head-on: line 379 now reads 'why the platform tier converges autonomously and self-heals, why the verdict tier requires external loop closure today' and 'makes explicit the open-loop boundary', plus the footnote conceding the framing is illustrative, not formally rigorous.

### ✅ #280 (p141) — already-addressed
**Kritikos:** `Totally agree here. One small issue though: do you cover observability in the sense of state-reconstructibility?`

**Anchored on:** "the use of the latter term in the SDT axiom set is deliberately coincidental"  → **06-gitops-pipeline.tex:377 (footnote)**

**Meaning:** 'Totally agree... but do you cover observability in the sense of state-reconstructibility?' The footnote he anchored on already answers: SDT's Observability axiom means structured machine-readable output, NOT state-reconstructibility, and the system does not claim classical control-theoretic observability. Optionally surface this distinction in the body text, but the question is answered where he asked it.

### ✅ #281 (p141) — already-addressed
**Kritikos:** `as we will see in the following sub-sections`

**Anchored on:** "Three GitHub Actions workflows implement these principles"  → **06-gitops-pipeline.tex:386**

**Meaning:** Suggested wording: append 'as we will see in the following sub-sections' to connect the three pipeline principles to subsections 6.7.1-6.7.3. The live source now ends the sentence with ', as detailed in the following subsections' (uncommitted edit) — done.

### 🔴 #282 (p142) — open
**Kritikos:** `why is this last in the chain? This needs some justification as normally this check is fast so in principle it should have been moved earlier in the workflow.`

**Anchored on:** "and a documentation check enforces meta-validation---that documentation artifacts remain synchronized"  → **06-gitops-pipeline.tex:393**

**Meaning:** Why is the documentation check LAST in the CI chain? It is normally fast, so fail-fast ordering would put it earlier. He wants a justification: either defend the ordering (it checks docs against the final validated tree) or state the real rationale/parallelism from ci.yml. No justification has been added.

### 🔴 #283 (p145) — open
**Kritikos:** `So, this should be done by every platform operator, right?I can also see that there are specific steps where some relate to the system installation, others to the CR application and finally others to PR commits.Maybe this is clearly indicated at each bullet or initially in the introduction paragraph (you can decide).`

**Anchored on:** "Post-deployment verification follows a layer-by-layer strategy (Section 6.8.2 Verification)"  → **06-gitops-pipeline.tex:451-461**

**Meaning:** He asks to confirm this verification is the platform operator's job, and observes the four steps belong to different lifecycle phases (system installation vs CR application vs PR commits). He wants each bullet — or the intro paragraph — to indicate the phase and responsible actor (author's choice of placement). Unchanged in the live source.


## 07-evaluation.tex — 97 comments (58 open)

### 🔴 #284 (p147) — open
**Kritikos:** `You forget the SDT Feedback as Input to Agent-Driven Specification Improvement. This issue was mentioned also in previous chapters.`

**Anchored on:** "The chapter concludes with CI/CD integration metrics"  → **07-evaluation.tex:4**

**Meaning:** The chapter intro frames the evaluation as three arms (PoC, controlled, large-scale) and omits the fourth experiment, the SDT-feedback/agent-driven specification-improvement study (Section 7.7) — Kritikos notes this experiment is forgotten here as in previous chapters. Add it to the framing.

### 🔴 #285 (p147) — open
**Kritikos:** `It seems that the content of the research questions has changed?I made a check wrt RQ1 and this is really the case. Please fix this.`

**Anchored on:** "RQ1 (Framework soundness): Does the SDT conceptual framework provide"  → **07-evaluation.tex:13-18**

**Meaning:** The chapter's RQ texts have drifted from the canonical wording in Chapter 1 (he verified RQ1 differs: intro asks about feasibility of QA 'without dedicated testers', chapter asks about deterministic, false-positive-free detection). Restore verbatim RQ wording from 01-introduction.tex:129-141 for all four RQs.

### 🔴 #286 (p148) — open
**Kritikos:** `See previous comment. RQ2 has two parts: automatic production of validation, testing and gates & actionable observability ...Further, in terms of the second, only the experiment in Section 7.5 is relevant or also other experiments?`

**Anchored on:** "RQ2 (Automation extent): To what extent can validation and testing be"  → **07-evaluation.tex:16**

**Meaning:** RQ2 has two parts — automatic production of validation/testing/gates AND actionable observability. The chapter's RQ2 entry only covers the first half (citing 7.2 and 7.5); he asks whether the actionability half is evidenced only by Section 7.7 or by other experiments too, and wants that stated explicitly.

### 🔴 #287 (p148) — open
**Kritikos:** `Is also large-scale evaluation relevant here?`

**Anchored on:** "the multi-API controlled evaluation"  → **07-evaluation.tex:15**

**Meaning:** Asks whether the large-scale APIs.guru evaluation (7.5) also belongs in RQ1's evidence list — the Summary tags the 50-API result to RQ1 but the RQ1 entry here omits Section 7.5. Add it for consistency.

### 🔴 #288 (p148) — open
**Kritikos:** `10 is a conclusions section that explains how the RQ is answered. Here you should focus solely on evaluation. So, the evaluation evidence that supports answering the research question.`

**Anchored on:** "Chapters 5 and 10 and operationally"  → **07-evaluation.tex:17**

**Meaning:** Chapter 10 is a conclusions chapter explaining how RQs are answered; the evaluation chapter's RQ4 entry should cite only evaluation evidence, not point forward to the conclusion. Drop the ch:conclusion pointer and keep the operational evidence (Section 7.2).

### 🔴 #289 (p148) — open
**Kritikos:** `Where is PoC in this figure?In addition, you have effort comparison in figure. Does this correspond to any arm of the evaluation design?`

**Anchored on:** "(non-critical-api)"  → **figures/evaluation-methodology.tex:38-59 (via 07-evaluation.tex:25)**

**Meaning:** Figure 7.1 is inconsistent with the chapter: the PoC/operational arm (presented first in the text) is missing from the figure, and the figure shows an 'Effort Comparison (SDT vs Manual)' arm that corresponds to no section. Add the PoC arm and remove or back the effort-comparison arm.

### ✅ #290 (p149) — already-addressed
**Kritikos:** `please be aware of the references ...`

**Anchored on:** "tions [494], [495], [510], which noted that earlier drafts mixed information"  → **07-evaluation.tex:36**

**Meaning:** The bracketed numbers [494], [495], [510] are not real bibliography references (they were supervisor-comment IDs); remove them and the round-1-review provenance. The live line 36 now describes the section structure self-containedly with no bracketed numbers.

### 🔴 #291 (p149) — open
**Kritikos:** `Ok but why do you have section 7.6 separate from 7.2?It can be argued that it is a very important part of the PoC, so it needs to be moved as last section of 7.2 (i.e., 7.2.8).`

**Anchored on:** "Section 7.6 concludes"  → **07-evaluation.tex:651 (pointer at :36)**

**Meaning:** Asks why the CI/CD Integration Metrics section is separate from the PoC section 7.2 — he argues it is an important PoC output and should be moved to become the last subsection of 7.2 (e.g., 7.2.8). The metrics section is still a standalone \section (sec:cicd-metrics).

### 🔴 #292 (p149) — open
**Kritikos:** `Ok but is this part of the cluster?I guess not so I would put it in a separate paragraph.I would also change the section title as Initial Situation in this case.`

**Anchored on:** "Five GitHub repositories under the novelcore"  → **07-evaluation.tex:39,43**

**Meaning:** The five GitHub repositories are not part of the cluster, so this off-cluster state should be split into its own paragraph; he also suggests retitling the subsection from 'Initial Cluster Situation' to 'Initial Situation'. Neither change is made in the live text.

### 🔴 #293 (p150) — open
**Kritikos:** `I believe that the first configuration step is to install the DriveBy system. Maybe this could be alternatively made a separate subsection (7.2 DriveBy Installation).In any case, you need to mention what else is (indeed) added to the cluster.`

**Anchored on:** "7.2.2 Proof-of-Concept Configuration"  → **07-evaluation.tex:48-52**

**Meaning:** The first configuration step should be installing the DriveBy system itself — possibly as its own subsection ('DriveBy Installation') — and the text must say what else was actually added to the cluster before the five XSDLC CRs. Line 45 lists the thesis-built additions, but 7.2.2 still jumps straight to applying CRs with no installation step.

### 🔴 #294 (p150) — open
**Kritikos:** `but in the production gate, we do not have the test-ready/validate-only check so this does not make sense.Unless you forgot to add this in the first row of Table 7.1`

**Anchored on:** "the production gate uses strict. test-ready runs"  → **07-evaluation.tex:72 (vs :63)**

**Meaning:** The prose says the perfect-api production gate uses strict validation, but Table 7.1's production-gate cell for perfect-api lists only the load-test check — either the table forgot the validate-only(strict) check or the prose is wrong. Still inconsistent in the live source.

### 🔴 #295 (p150) — open
**Kritikos:** `see previous comment. We do not see this in table.`

**Anchored on:** "non-critical-api clears both, so the choice has no operational"  → **07-evaluation.tex:63,72**

**Meaning:** Follow-up to 294: 'we do not see this in table' — the strict production-gate check that non-critical-api supposedly clears is not visible in Table 7.1. Same table/prose reconciliation still pending.

### 🔴 #296 (p150) — open
**Kritikos:** `we do not see this on broken-api ...`

**Anchored on:** "broken-api) carry a production gate with autoMerge:false"  → **07-evaluation.tex:67 (vs :72)**

**Meaning:** The prose claims broken-api carries a production gate with autoMerge:false, but Table 7.1's broken-api row shows '---' for the production gate. Verify broken-api's real gate config and make the table and prose agree. Still inconsistent.

### 🔴 #297 (p150) — open
**Kritikos:** `Please check table's content as it seems inconsistent with the text that follows it.`

**Anchored on:** "bad-docs-api (row of Table 7.1)"  → **07-evaluation.tex:54-72**

**Meaning:** General request to check Table 7.1's content against the paragraph that follows it (the 'two design choices' text about strict production checks and autoMerge:false on three apps) — the same inconsistencies as 294-296. Still open.

**⚠ Triage correction:** Triage pointed the reconciliation at the staging-gate results table (P005 vs P002/P003/P004) much later in the chapter; the comment actually targets Table 7.1 versus the paragraph immediately following it (strict production check and autoMerge claims), i.e. the same locus as idx 294-296.

### ✅ #298 (p151) — already-addressed
**Kritikos:** `again, we have non-existing reference number.`

**Anchored on:** "CRDs (per [509], the supervisor's request for approx-"  → **07-evaluation.tex:98**

**Meaning:** '[509]' is another non-existent reference number; delete it and the supervisor-request provenance. Live line 98 now reads 'The 32 CommitStatus CRDs decompose by phase as follows' with no bracketed number.

### 🔴 #299 (p151) — open
**Kritikos:** `if we have 44 workflow runs, how is it possible to have only 32 commit status CRDs? Please comment on this in the current paragraph.`

**Anchored on:** "The 44 Workflow runs include both staging and"  → **07-evaluation.tex:98**

**Meaning:** Asks how 44 workflow runs can coexist with only 32 CommitStatus CRDs, and wants this explained in the same paragraph (e.g., retries and multiple runs map to one CommitStatus per commit/environment). The retry/staging+production sentence exists but the 44-vs-32 arithmetic is still not explicitly reconciled.

### ✅ #300 (p151) — already-addressed
**Kritikos:** `again this is an unknown/non-existing reference number.`

**Anchored on:** "Per [370], the count decomposes"  → **07-evaluation.tex:107**

**Meaning:** '[370]' is a non-existent reference number; remove it. Live line 107 now reads 'For the perfect-api CR observed on the cluster, the count decomposes as:' — self-contained.

### 🔴 #301 (p152) — open
**Kritikos:** `ok but based on what was stated in Section 5.4.1, we expected 3 * 6 = 18 environment-specific resources, 8 app resources and 8 CR resources. This makes 34 objects and not 44 or 47. So, maybe there is a need to update Section 5.4.1 to be more precise?`

**Anchored on:** "The 44–47 spread"  → **07-evaluation.tex:107**

**Meaning:** His arithmetic from Section 5.4.1 (3x6 env resources + 8 app + 8 CR = 34 objects) does not match the 44-47/45-50 counts reported here; he suggests Section 5.4.1 needs to be made more precise. Chapter 5 now says '~44 managed resources' while this chapter says 'approximately 45--50' and a '44--47 spread' — the cross-chapter decomposition is still not reconciled to one authoritative count.

### ✅ #302 (p152) — already-addressed
**Kritikos:** `I would change paragraph here.`

**Anchored on:** "branch and protection objects once or per environment. Without XSDLC,"  → **07-evaluation.tex:109**

**Meaning:** Requests a paragraph break before 'Without XSDLC, a platform operator would need...'. The live source now starts a new paragraph at line 109 exactly there.

### 🔴 #303 (p152) — open
**Kritikos:** `well, they are needed in one case (autoMerge: false). But the validation result is clear. The user must decide in this case.`

**Anchored on:** "human authoring of test cases, or human interpretation of results"  → **07-evaluation.tex:116**

**Meaning:** The zero-human claim overstates: with autoMerge:false a human must still decide to merge (though the validation verdict itself is clear). Add a qualifier acknowledging the human approval step in that configuration. Live line 116 is still unqualified.

### 🔴 #304 (p152) — open
**Kritikos:** `ok but the focus is on structural drift, not semantic. Semantic drift would require also doing full functional checking (checking also the real content and not just the structure of responses).`

**Anchored on:** "Any drift between the"  → **07-evaluation.tex:121**

**Meaning:** The static contract detects structural drift only; semantic drift (correctness of actual response content) would require full functional checking. Qualify the 'any drift is detected' claim accordingly. Live line 121 still says 'Any drift ... is detected before promotion' without the structural/semantic distinction.

### 🔴 #305 (p153) — open
**Kritikos:** `This is at the very heart of your thesis. I would better position this at the beginning of the thesis to be crystal clear what is the overall work scope.`

**Anchored on:** "SDT replaces the promotion-gate testing that would"  → **07-evaluation.tex:125 (cross-cutting to 01-introduction.tex)**

**Meaning:** The 'SDT does not replace all testing; it replaces promotion-gate testing (BYOCI boundary)' statement is the very heart of the thesis scope and should be positioned at the beginning of the thesis (Chapter 1), not only here. No such scope statement exists in 01-introduction.tex yet.

### 🔴 #306 (p154) — open
**Kritikos:** `ok but what does this mean in practice?Is there something that needs improvement?I am asking this as you indicate that this is a positive finding. However, at the table is marked with X.`

**Anchored on:** "means the 5-minute timeout path was only observed when the staging gate"  → **07-evaluation.tex:160,169**

**Meaning:** The health-check-timeout case is described as a positive resilience finding in prose but marked ✗ (fail) in Table 7.2 — a contradiction. Clarify what it means in practice and whether anything needs improvement (e.g., mark the row 'not reproducible' instead of fail). Live table row 160 still shows \xmark while line 169 calls it positive.

### ✅ #307 (p155) — already-addressed
**Kritikos:** `again non-existing reference numbers ...`

**Anchored on:** "Naming note (round-1 review [466], [476])."  → **07-evaluation.tex:178**

**Meaning:** '[466], [476]' are non-existent reference numbers; drop them and the round-1 provenance. Live line 178 now reads 'Naming note. The reference API ... is renamed to non-critical-api because the previous name was misleading' — self-contained.

### 🔴 #308 (p155) — open
**Kritikos:** `ok but does this mean that non-critical-api also passes dynamic validation tests?If this is the case, then it needs to be communicated earlier in this paragraph.This must be, however, true as the main rationale is to use a robust/reliable API as a base for producing problematic API versions through error/fault injection.`

**Anchored on:** "runtime layer in slow-api/broken-api). Source-code identifiers"  → **07-evaluation.tex:178,187**

**Meaning:** Asks whether non-critical-api also passes the dynamic/runtime tests (P006/P007) — it must, since it is the robust base for fault injection — and wants this stated earlier in this paragraph, not only implied. The naming note and Reference API Design still mention only static principles P001-P005/P008/P009; runtime cleanliness is only implicit (line 495).

### ✅ #309 (p155) — already-addressed
**Kritikos:** `current`

**Anchored on:** "The controlled evaluation uses this reference API as a positive control"  → **07-evaluation.tex:178**

**Meaning:** One-word comment 'current' — a wording fix for the preceding clause 'in this revised text it is renamed': use 'current text' or drop the revision framing. The live text deleted 'in this revised text' entirely ('it is renamed to non-critical-api because...'), so the issue is resolved.

### 🔴 #310 (p155) — open
**Kritikos:** `how would you characterise this API in terms of functionality?Product Management API?`

**Anchored on:** "The non-critical-api is a Python FastAPI application"  → **07-evaluation.tex:187**

**Meaning:** Asks for a functional characterisation of the reference API — e.g., 'a Product Management API'. The live design paragraph lists the endpoints (/products CRUD etc.) but still gives no one-line functional characterisation.

### 🔴 #311 (p157) — open
**Kritikos:** `producedTo indicate that this was indeed the case - you actually attempted this to prove determinism ...`

**Anchored on:** "against the same defective specification produces identical results"  → **07-evaluation.tex:246**

**Meaning:** One-word 'produced': switch to past tense to indicate the determinism check was actually performed (you really ran it multiple times), not just a property statement. Live line 246 still says present-tense 'produces identical results'.

### 🔴 #312 (p158) — open
**Kritikos:** `Maybe it is better to say that you have created three mutations covering two principles and one mutation covering three principles.This is what I understand from this sentence. I leave it to you to decide how to modify the sentence to better reflect reality.`

**Anchored on:** "generates three random pairs and one random"  → **07-evaluation.tex:255**

**Meaning:** Suggests clearer phrasing: 'three mutation-combinations covering two principles each and one combination covering three principles' instead of 'three random pairs and one random triple'. He leaves exact wording to the author. Live line 255 keeps the original phrasing.

### 🟡 #313 (p158) — needs-author-decision
**Kritikos:** `In essence, this is good as an evaluation approach.But it is not complete. In my view, complete would mean to cover all checks across all principles. And to conduct this coverage multiple times to introduce some randomness in the way a specific violation (wrt a check)  is realised.However, I do not impose this as my comment is addressed at a good level. It is up to you to decide whether you can fully address it.`

**Anchored on:** "between two and three mutations per implemented principle, each targeting a"  → **07-evaluation.tex:255**

**Meaning:** The per-check protocol is good but not complete — full coverage would mutate all checks across all principles, repeated with randomised violation realisations. He explicitly does not impose this ('my comment is addressed at a good level... up to you'); at most add a limitation/future-work note.

### ✅ #314 (p158) — already-addressed
**Kritikos:** `I would not count this as an issue as it actually leads to a parse error so somehow it is covered, right?So, the actual critical issues are the missed detections. As these are justified below, I am OK!However, I would prefer to avoid the load problem in an ideal evaluation.`

**Anchored on:** "strip schema length/range constraints"  → **07-evaluation.tex:275,315**

**Meaning:** The loader-rejected P004 mutation should not count as an issue — the parse error means it is effectively covered; the real critical items are the two missed detections, which the text justifies, so he is OK. He would merely prefer an ideal evaluation to avoid generating unparseable mutations. The existing 'Interpretation of the parse error' paragraph (line 315) covers his acceptance.

### 🔴 #315 (p159) — open
**Kritikos:** `yes, this is defensible. You actually removed sth that would not cause any issue. So, it is a faulty issue from the very beginning. Thus, correctly maps to a true negative that is indeed covered by your tool!This is a problem caused by the way the faults are injected (evaluation problem). It does not have to do sth with your implemented work ...`

**Anchored on:** "dant when a global security requirement is already in place—a defensible"  → **07-evaluation.tex:311**

**Meaning:** The p005-strip-op-security miss is a true negative caused by how the fault was injected (removing per-op security that is redundant under a global requirement causes no real defect) — an evaluation problem, not a framework limitation. The live text still classifies it as one of two 'genuine framework limitations' (line 311); the reclassification has not been made.

### ✅ #316 (p159) — already-addressed
**Kritikos:** `ok, so this indicates that this is truly a false negative, marking incompleteness in the way P009 principle was implemented.`

**Anchored on:** "examples rather than parameter-level examples; the targeted defect bypasses"  → **07-evaluation.tex:311**

**Meaning:** Confirms that the p009-strip-param-examples miss IS a genuine false negative revealing incompleteness in P009's implementation (schema-level vs parameter-level examples). The live text already states exactly this; no change implied beyond keeping the classification.

### 🔴 #317 (p159) — open
**Kritikos:** `Well, I am not sure about the first issue. I believe that it is a problem of the way the evaluation has been performed. See previous^2 comment.`

**Anchored on:** "the check by definition. Both misses are genuine framework limitations rather"  → **07-evaluation.tex:311**

**Meaning:** Disputes the 'both misses are genuine framework limitations' claim for the first (p005) miss — per his earlier comment it is an artefact of how the evaluation was performed. Live line 311 still says 'Both misses are genuine framework limitations'; needs the asymmetric reclassification (p005 = evaluation artefact, p009 = framework limitation).

### ✅ #318 (p159) — already-addressed
**Kritikos:** `I agree here. Again, this is a matter of how evaluation was implemented, not a limitation of your framework.`

**Anchored on:** "talk (a real overlap between checks), not a false positive"  → **07-evaluation.tex:313**

**Meaning:** Agreement: the P002→P008 collateral is a matter of how the evaluation was implemented, not a framework limitation. The live cross-talk explanation already matches his reading; no edit required.

### ✅ #319 (p159) — already-addressed
**Kritikos:** `Ok, I agree that it relates to P001 theoretically but it is caught by the tool before P001 could be executed. Again, as indicated in a previous comment of mine, this is an evaluation error, not a framework issue ...`

**Anchored on:** "any principle checker ran. This is a P001-level (compliance) catch occurring"  → **07-evaluation.tex:315**

**Meaning:** Agreement: the parse-error case relates to P001 theoretically but is caught by the loader before P001 runs — an evaluation error, not a framework issue. The live interpretation paragraph already frames it as a loader-level catch, not a P004 miss, matching his view.

### 🔴 #320 (p159) — open
**Kritikos:** `Missing a concluding paragraph here.Are you satisfied with the produced evaluation results. Do they showcase sth important for your framework? Do they also signify limitations that must be covered by future work?`

**Anchored on:** "confirming the Axiom of Determinism at the experimental scale"  → **07-evaluation.tex:317**

**Meaning:** Section 7.3.4 lacks a concluding paragraph: state whether the results are satisfactory, what they showcase about the framework, and which limitations they signify for future work. The live section still ends abruptly at the determinism check before the next \section.

### 🔴 #321 (p160) — open
**Kritikos:** `This table should be presented before Table 7.8. First, we see the intent and then the outcome ...`

**Anchored on:** "Table 7.9; the empirically measured strict-mode outcome is given in Table 7.8"  → **07-evaluation.tex:331-376**

**Meaning:** Present the intent table (api-suite, Table 7.9) BEFORE the measured-outcome table (defect-ground-truth, Table 7.8): first intent, then outcome. In the live source the outcome table (line 333) still precedes the intent table (line 358).

### 🔴 #322 (p160) — open
**Kritikos:** `But you report them with these values and normal ones like runtime and runtime fail ....`

**Anchored on:** "are reported as "–"."  → **07-evaluation.tex:335,344-348**

**Meaning:** The caption says P006/P007 are 'reported as "--"' but the column actually mixes '--' with values like 'runtime' and 'runtime fail (P007/P006)'. Make caption and cells consistent. Live caption (line 335) and cells (lines 344-348) still mismatch.

### ✅ #323 (p160) — already-addressed
**Kritikos:** `Maybe it is better in general not to indicate what was done on document revisions.The main goal is to present the current version of the evaluation, not how it was evolved. This will also reduce the text in the report (as it now becomes verbose in some placed and the verbosity is further enhanced through the addition of evolution statements).`

**Anchored on:** "these, P002, P003, and P004 are critical under the post-revision severity"  → **07-evaluation.tex:354**

**Meaning:** Stop narrating document revisions ('post-revision severity', 'round-1 review [473]'); present only the current evaluation, reducing verbosity. The live paragraph (line 354) now describes the measured matrix with no revision narration or fake references.

### ✅ #324 (p160) — already-addressed
**Kritikos:** `Are you sure that the table's content is correct?It seems that many APIs have common principle failures while the goal was that they fail in just one principle (as they are just derived from the non-critical-api that has not principle failures).`

**Anchored on:** "bad-docs-api (Table 7.8 rows)"  → **07-evaluation.tex:344-348**

**Meaning:** Core data-integrity challenge: the table showed many APIs sharing P002-P004 failures although each API was supposed to fail in exactly one principle (being derived from a clean base). The defect APIs were rebuilt from the clean base and the results regenerated: live Table 7.8 now shows no-auth-api failing only P005 and slow/broken-api passing all static principles; bad-docs-api's P002-P004/P008 cascade is explained as a legitimate documentation-class strip.

### 🔴 #325 (p160) — open
**Kritikos:** `I would better explain how each API was derived from non-critical-api and present Table 7.9  first.`

**Anchored on:** "7.4.1 Evaluation API Suite Design"  → **07-evaluation.tex:331**

**Meaning:** Wants 7.4.1 to first explain how each API was derived from non-critical-api and to present Table 7.9 (intent) first. The derivation classes are now described (lines 178/187/331) but the per-API derivation detail in 7.4.1 is still thin and the table order is still outcome-first.

### ✅ #326 (p161) — already-addressed
**Kritikos:** `/P008Please see also the table content and make necessary corrections in the main report text where appropriate.`

**Anchored on:** "bad-docs-api blocks on P002/P003/P004"  → **07-evaluation.tex:354**

**Meaning:** '/P008' — the prose omitted P008 from bad-docs-api's failure list relative to the table. The live prose (line 354) now includes 'versioning documentation (P008)' in the documented cascade, matching Table 7.8.

### ✅ #327 (p161) — already-addressed
**Kritikos:** `but it seems that it also exhibits additional failures by observing Table's 7.8 content`

**Anchored on:** "no-auth-api on P005, and the runtime defects (slow-api, broken-api)"  → **07-evaluation.tex:354,346**

**Meaning:** Observed that no-auth-api exhibited additional failures in Table 7.8 beyond P005. Resolved by the data regeneration: no-auth-api now fails exactly one principle (P005) in both the table and the prose ('surfaces as exactly one violation').

### ✅ #328 (p161) — already-addressed
**Kritikos:** `but the shared base specification does not have any critical failure ...`

**Anchored on:** "herit P002/P003/P004 critical failures from their shared base specification."  → **07-evaluation.tex:354**

**Meaning:** Logical impossibility: the shared base (non-critical-api) has no critical failures, so slow/broken-api cannot 'inherit' P002-P004 failures from it. Resolved: the live text says slow/broken-api 'pass all six static principles: their specifications are byte-identical to the reference'.

### ✅ #329 (p161) — already-addressed
**Kritikos:** `addresses`

**Anchored on:** "outcomes also corrects supervisor comment [467]"  → **07-evaluation.tex:356**

**Meaning:** One-word 'addresses' — replace 'corrects' with 'addresses' (and implicitly the fake [467] reference is part of the same problem). The whole sentence referencing supervisor comment [467] was deleted; the successor passage (line 356) is self-contained.

**⚠ Triage correction:** Triage anchored the one-word comment to the 'address the contradictions flagged' clause; the icon actually sits on 'outcomes also corrects supervisor comment [467]' and proposes 'corrects'→'addresses'. Both readings are mooted because the sentence was removed.

### 🔴 #330 (p161) — open
**Kritikos:** `ok but it is quite easy to create such an OpenAPI document. Much easier that violating the other principles.`

**Anchored on:** "critical-P001 example would require a deliberately malformed OpenAPI docu-"  → **07-evaluation.tex:356**

**Meaning:** Counters the implied difficulty: a malformed OpenAPI document is much easier to create than violations of the other principles, so deferring the critical-P001 example as future work reads oddly — either add one or own the choice. Live line 356 still says it 'would require a deliberately malformed OpenAPI document ... recorded as a candidate addition'.

### ✅ #331 (p161) — already-addressed
**Kritikos:** `what do you mean with all W? Warning? But you have corrected this by indicating that now the P002-P004 principles are critical ..`

**Anchored on:** "Stripped documentation (bad-docs-api expected-failure cell)"  → **07-evaluation.tex:369**

**Meaning:** Questions '(all W)': with P002-P004 corrected to critical, the expected-failure label 'P002, P003, P004, P008 (all W)' is wrong. Live Table 7.9 cell now reads 'P002, P003, P004 (C); P008 (W)'.

### ✅ #332 (p161) — already-addressed
**Kritikos:** `What does critical sweep mean? It is not clear. Maybe remove it?`

**Anchored on:** "P005 (critical) + critical sweep"  → **07-evaluation.tex:370**

**Meaning:** 'critical sweep' is undefined and confusing — remove or define it. The live no-auth-api cell now reads just 'P005 (critical)'.

### ✅ #333 (p162) — already-addressed
**Kritikos:** `Shouldn't no-auth-api fail only on P005?`

**Anchored on:** "no-auth-api row: ✗ on P002–P005"  → **07-evaluation.tex:400**

**Meaning:** no-auth-api should fail only on P005 given the single-injection design. Resolved at the data level: live Table 7.10 row shows no-auth-api passing P001-P004 and P008, failing only P005 (5/6, 1 critical failure).

### ✅ #334 (p162) — already-addressed
**Kritikos:** `These two APIs should have only runtime faults so they should pass the checks ...`

**Anchored on:** "slow-api row: ✗ on P002–P004"  → **07-evaluation.tex:401-402**

**Meaning:** slow-api and broken-api carry only runtime faults, so they should pass all static checks. Resolved: live Table 7.10 shows both at a clean 6/6 with 0 critical static failures; they fail only at the P007/P006 runtime layer.

### ✅ #335 (p162) — already-addressed
**Kritikos:** `It also fails P008, why?`

**Anchored on:** "critical principles (P002/P003/P004 stripped from the documentation, error"  → **07-evaluation.tex:354,345,399**

**Meaning:** Asks why bad-docs-api also fails P008. The live text explains the documentation strip removes the versioning-strategy evidence P008 checks (info.description keywords), and Tables 7.8 and 7.10 now agree (P008 = F/✗ for bad-docs-api).

### ✅ #336 (p162) — already-addressed
**Kritikos:** `I do not understand this. You should have only P005. You do not get P002-P004 if you strip only authentication information. In addition, if no-auth-api was based on bad-docs-api wouldn't have also P008?In any case, it is not based on bad-docs-api but non-critical-api, right?`

**Anchored on:** "inherited from the same documentation strip plus P005 from the security strip)"  → **07-evaluation.tex:408**

**Meaning:** no-auth-api is derived from non-critical-api (not from bad-docs-api), so it cannot 'inherit' P002-P004 failures and should fail only P005; he also notes the inheritance story would imply a P008 failure that isn't there. Resolved: the inheritance claim is gone and live prose says no-auth-api fails 'a single critical principle (P005, the security strip), with all of its other principles passing'.

### ✅ #337 (p162) — already-addressed
**Kritikos:** `See previous comments ...`

**Anchored on:** "and slow-api/broken-api on three each (P002/P003/P004 inherited from"  → **07-evaluation.tex:408**

**Meaning:** 'See previous comments' — same erroneous-propagation paragraph. Resolved: the paragraph was rewritten (line 408) so the static gate blocks exactly bad-docs-api and no-auth-api, and slow/broken-api score clean 6/6 statically and are caught at runtime.

### 🔴 #338 (p162) — open
**Kritikos:** `Shouldn't you explain first what was involved in each gate and which gates were involved?`

**Anchored on:** "Each API was deployed through the XSDLC pipeline, triggering staging gate"  → **07-evaluation.tex:415**

**Meaning:** Before showing staging-gate outcomes, first explain what checks each gate comprised and which gates each API traversed (intent before outcome again). The live 7.4.3 still jumps straight from one deployment sentence to the results table with no gate-composition explanation or link back to Table 7.1.

### 🔴 #339 (p163) — open
**Kritikos:** `I am puzzled here as you need to have the same checks per gate. While it seems that you apply validate-only with test-ready for non-critical-api and validate-only with strict the APIs. Is this really correct?`

**Anchored on:** "non-critical-api: validate-only (test-ready) ... ✓PASS"  → **07-evaluation.tex:427**

**Meaning:** Puzzled that the staging gate applies validate-only(test-ready) for non-critical-api but validate-only(strict) for the other APIs — gates should apply the same checks per gate for an apples-to-apples comparison; is this configuration really correct? The live table still shows the non-uniform modes with no local justification.

### ✅ #340 (p163) — already-addressed
**Kritikos:** `This is wrong. P002-P004 failed which are critical`

**Anchored on:** "bad-docs-api ... ✗FAIL P005 Security (critical)"  → **07-evaluation.tex:429**

**Meaning:** The bad-docs-api blocking-failure cell wrongly said P005; its actual critical failures are P002-P004. Resolved: live Table 7.11 cell reads 'P002/P003/P004 Documentation (critical)'.

### 🔴 #341 (p163) — open
**Kritikos:** `But there was no blocking for this API for the current gate. As it passed all checks ...`

**Anchored on:** "slow-api was blocked at"  → **07-evaluation.tex:437 (vs :428)**

**Meaning:** In the staging-results prose, slow-api is described as 'blocked at the runtime layer on P007 ... exceeding the 200 ms staging threshold', but per the table slow-api PASSED its staging gate (it has no load-test check at staging; the P007 block happened at the production gate). The live line 437 still carries this contradiction — slow-api 'blocked ... staging threshold' vs Table 7.11 row PASS.

**⚠ Triage correction:** Triage anchored this to broken-api and read it as blurring static-vs-runtime blocking; the icon actually sits on 'slow-api was blocked at' and the point is that slow-api passed all staging checks — its P007 failure belongs to the production gate, so the staging subsection must not say it was blocked.

### ✅ #342 (p163) — already-addressed
**Kritikos:** `This is totally wrong here + you talk again about revision while it is not needed`

**Anchored on:** "the legacy severity assignment (P002–P004 warning) used in earlier drafts"  → **07-evaluation.tex:437 (passage removed)**

**Meaning:** The legacy-severity/post-revision note is both wrong and unnecessary — drop the revision narration and state only current behaviour. Resolved: the entire 'Note: this paragraph reflects the post-revision severity...' passage was deleted from the live staging-results prose.

### 🔴 #343 (p164) — open
**Kritikos:** `Ok but this delay should have been explain in 7.4.1 how it is applied ...`

**Anchored on:** "The 500 ms artificial delay, invisible to single-request functional testing"  → **07-evaluation.tex:371 (suite design) / :466**

**Meaning:** The 500 ms delay mechanism should be explained back in 7.4.1 (how the delay is applied, e.g., a sleep in every handler), not first surface in the production-gate discussion. The live suite-design subsection still only says 'Performance degradation (500 ms delay)' with no mechanism.

### ✅ #344 (p164) — already-addressed
**Kritikos:** `Question: if P007 is warning, then why do we have failure?In any case, violating non-functional requirements isn't critical?`

**Anchored on:** "delay, invisible to single-request functional testing, manifested clearly under"  → **07-evaluation.tex:371,502**

**Meaning:** If P007 is a warning, why does it cause a blocking failure? And shouldn't violating a declared non-functional requirement be critical? Resolved by the chapter-wide P007→critical flip: Table 7.9 now reads 'P007 (load test, C)' and the severity model states a P007 breach always blocks ('no advisory mode for a violated SLA').

### 🔴 #345 (p165) — open
**Kritikos:** `If these two principles are to be ignored, sth which is correct as they are not yet implemented, then the right value here is 5/7 and not 5/9. As only 7 principles also affect the static validation and not 9. The other two concern runtime aspects.`

**Anchored on:** "but are not counted in the 5/9 static-Critical"  → **07-evaluation.tex:504**

**Meaning:** If P006/P007 are runtime principles excluded from static validation, the static-critical ratio should be 5/7 (seven static principles), not 5/9. Partially fixed (line 502 now counts seven critical of nine overall), but line 504 still says 'This 5/9 critical ratio', which is now stale and still not the 5/7 he asked for.

### ✅ #346 (p165) — already-addressed
**Kritikos:** `I do not agree with the position of this sentence. As you talk only about static validation.So, we should not care whether non-critical-api passes all gates. We care that APIs with no documentation-oriented defects indeed pass Layer 1 and thus static validation. This accounts for non-critical-api, slow-api and broken-api.`

**Anchored on:** "plete, hand-crafted specification—passes all gates. The four defect APIs"  → **07-evaluation.tex:510**

**Meaning:** Misplaced sentence: in a static-validation discussion the point is that APIs without documentation defects (non-critical, slow, broken) pass Layer 1, not that only non-critical-api passes all gates. Resolved: the live bullet now distinguishes the two specification-defect APIs (blocked statically) from the two behavioural-defect APIs that 'pass static validation, and are blocked at the runtime stage instead'.

### ✅ #347 (p166) — already-addressed
**Kritikos:** `I disagree here for the last two APIs!!!`

**Anchored on:** "(bad-docs-api, no-auth-api, slow-api, broken-api) are all blocked"  → **07-evaluation.tex:510**

**Meaning:** Strong disagreement that all four defect APIs are blocked at the validation stage — slow-api and broken-api have complete specs and must pass static validation. Resolved by the same rewrite: only bad-docs-api and no-auth-api are now said to be blocked at the static stage.

### ✅ #348 (p166) — already-addressed
**Kritikos:** `harvested`

**Anchored on:** "The evaluation harvests OpenAPI specifications from the APIs.guru registry"  → **07-evaluation.tex:526**

**Meaning:** One-word 'harvested': use past tense for performed evaluation steps. Live line 526 now reads 'The evaluation harvested OpenAPI specifications...'.

### ✅ #349 (p166) — already-addressed
**Kritikos:** `created includes ....`

**Anchored on:** "The dataset includes specifi-"  → **07-evaluation.tex:526**

**Meaning:** 'created includes ....': change to 'The created dataset includes...'. Live line 526 now reads 'The created dataset includes specifications...'.

### 🔴 #350 (p166) — open
**Kritikos:** `Do you also need to include P009 here? Isn't it also related to static validation?`

**Anchored on:** "tion principles (P001–P005, P008) are evaluated; P006 (Functional Testing)"  → **07-evaluation.tex:526**

**Meaning:** Asks whether P009 — also a static principle — should be included in the large-scale static evaluation, or at least why it is excluded. The live protocol still evaluates P001-P005/P008 only and explains the exclusion of P006/P007 but says nothing about P009.

### 🔴 #351 (p167) — open
**Kritikos:** `Could indicate how many specs were included in the dataset in 7.5.1. The same for the distribution of the specs wrt OpenAPI version.`

**Anchored on:** "against 50 OpenAPI 3.x specifications sampled from APIs.guru: 26"  → **07-evaluation.tex:526 (7.5.1) vs :552**

**Meaning:** State the dataset size and its OpenAPI-version distribution up front in 7.5.1 (Dataset and Protocol), not first in the results subsection. The counts (50 specs; 26 at 3.0.x, 24 at 3.1.0) still appear only at line 552, not in 7.5.1.

### 🔴 #352 (p167) — open
**Kritikos:** `Ok but isn't this part of the original dataset? Thus, shouldn't we indicate that the dataset has 70 OpenAPI specs, where 50 conform to 3.x.x and 20 to Swagger/OpenAPI 2.x?`

**Anchored on:** "validated successfully; no timeouts or harness errors occurred. The Ope-"  → **07-evaluation.tex:552,619**

**Meaning:** The Swagger 2.0 arm is part of the same harvest, so frame one dataset of 70 specs (50 conforming to 3.x, 20 to 2.x) rather than presenting the 2.x specs as a wholly separate dataset. The live text still introduces 50 specs and defers 'the OpenAPI 2.x portion' to a separate section without a 70-spec framing.

### 🔴 #353 (p167) — open
**Kritikos:** `Ok but why do you report them separately? This needs some explanation.Maybe you could first present the results at a global level and then go to local level with local results per major OpenAPI version ...`

**Anchored on:** "nAPI 2.x portion of the registry is reported separately in Section 7.5.4"  → **07-evaluation.tex:552,615-648**

**Meaning:** Why report 2.x separately? Either justify the split or present global results first (across all 70 specs) and then per-version local results. No global-then-local restructuring or justification has been added.

### 🔴 #354 (p167) — open
**Kritikos:** `The mapping from qualitative values to quantitative constraints could be justified.`

**Anchored on:** "high ≥80%, medium 30–79%, low < 30%"  → **07-evaluation.tex:554**

**Meaning:** Justify the mapping from qualitative buckets to the quantitative thresholds (why 80/30?). Live line 554 declares the thresholds with no justification.

### 🔴 #355 (p167) — open
**Kritikos:** `produced is the following:`

**Anchored on:** "The headline pattern: three of the six implemented principles"  → **07-evaluation.tex:575**

**Meaning:** Wording fragment 'produced is the following:' — he wants the results introduced as produced output, e.g., 'The headline pattern produced is the following: ...'. The preceding sentence gained 'The produced per-principle pass rates...' (line 554), but the anchored sentence itself still reads 'The headline pattern: three of the six...'.

### 🔴 #356 (p168) — open
**Kritikos:** `security`

**Anchored on:** "typically declare schemes but do not require them"  → **07-evaluation.tex:575**

**Meaning:** One-word 'security' at the icon: make the clause parallel — 'the 37 failing APIs typically declare security schemes but do not require them'. The live sentence still says just 'declare schemes'.

**⚠ Triage correction:** Triage read this as a request to delete the fake reference '[490]' earlier in the sentence (which was deleted anyway); the one-word comment is more plausibly an insertion of the word 'security' before 'schemes' at the icon position.

### 🔴 #357 (p168) — open
**Kritikos:** `and usage/documentation gap in the sense that it is not clear to the user/consumer which operations are protected and which are not.`

**Anchored on:** "test-orientation gap (the spec author did not document which operations are"  → **07-evaluation.tex:575**

**Meaning:** Add that this is also a usage/documentation gap: an API consumer cannot tell from the spec which operations are protected and which are not. The live sentence still frames it only as a test-orientation gap vs security gap.

### 🔴 #358 (p169) — open
**Kritikos:** `So, this API (no-critical-api) has sufficient information for API testing and usage?`

**Anchored on:** "the controlled-evaluation result that DriveBy reports zero false positives on"  → **07-evaluation.tex:599**

**Meaning:** Asks whether non-critical-api therefore carries sufficient information for both API testing and API usage — wants an explicit affirmation tying the zero-false-positive reference back to completeness/usability. No such clause has been added.

### 🔴 #359 (p169) — open
**Kritikos:** `But if we have 46% for P001, even the specification form is not properly validated by the current tooling! This is even worse result than the one indicated here. It signifies that current tooling is insufficient even for structural validation!Right?`

**Anchored on:** "tural compliance) but not substance (documentation completeness, schema"  → **07-evaluation.tex:601**

**Meaning:** With P001 at 46%, even the specification's FORM is not properly validated by current tooling for the majority of public APIs — a more severe finding than the form-vs-substance framing conveys; state this stronger reading. The live form-vs-substance paragraph (line 601) is unchanged, though line 575 partially notes structural compliance 'is not universal in the wild'.

### ✅ #360 (p169) — already-addressed
**Kritikos:** `Based on what we see in Figure 7.2, there is no corroboration.However, the results in this Figure are wrong!The same of course goes for the current paragraph that comments them!`

**Anchored on:** "The five-API controlled evaluation (Section 7.4) corroborates this pattern"  → **07-evaluation.tex:603 + figures/per-principle-pass-rates.tex**

**Meaning:** The 'corroborates' claim and Figure 7.2's values were wrong against the (then-broken) tables. Resolved: data regenerated — the figure now shows P001 100% and all other principles 80% matching Table 7.10, and the paragraph was rewritten as a 'contrasting baseline' (80-100% engineered vs 0-46% public) instead of corroboration.

### ✅ #361 (p170) — already-addressed
**Kritikos:** `round 6?Do you need to report such information in the formal thesis report?`

**Anchored on:** "tended in round 6 with a 2.0 version selector that accepts specifications"  → **07-evaluation.tex:621**

**Meaning:** 'round 6?' — development-round provenance does not belong in a formal thesis. Resolved: live line 621 reads 'The harvester ... supports a 2.0 version selector' with no round mention.

### 🔴 #362 (p171) — open
**Kritikos:** `But in the 3.x results, P001 was not passed universally!!!!`

**Anchored on:** "P001 again passes universally—each spec is structurally valid 2.0"  → **07-evaluation.tex:646**

**Meaning:** 'again' is false: in the 3.x arm P001 was NOT passed universally (46%). Remove 'again' and explicitly contrast/explain the 2.0 (100%) vs 3.x (46%) P001 inversion. Live line 646 still says 'P001 again passes universally'.

### 🔴 #363 (p172) — open
**Kritikos:** `I would prefer to see a comparison table showcasing the improvement obtained through the proposed automated approach.`

**Anchored on:** "Table 7.17: CI/CD integration metrics for the DriveBy framework"  → **07-evaluation.tex:657-678**

**Meaning:** Wants a comparison table showing the improvement of the automated SDT approach over a manual approach (effort/time per activity). No SDT-vs-manual comparison table exists; only one prose sentence ('compares favorably against the hours of manual configuration').

### 🔴 #364 (p172) — open
**Kritikos:** `In the context of given constraints, could supply also the actual ranges that support/satisfy them.`

**Anchored on:** "<1 s (Startup time)"  → **07-evaluation.tex:668**

**Meaning:** For threshold-style entries, also supply the actual measured ranges that satisfy the constraints (e.g., observed startup times), not only the bound. Table 7.17 values are unchanged.

### 🔴 #365 (p172) — open
**Kritikos:** `Are these all normal CI/CD integration metrics?They seem more as technical integration metrics.Please see a list of standard metrics that could be adopted: Deployment FrequencyLead Time for ChangesChange Failure RateMean Time to Recovery (MTTR)Build Success RatePipeline DurationTest Pass RateDeployment Success RateDefect Escape RateSecurity Gate Pass Rate`

**Anchored on:** "Table 7.17: CI/CD integration metrics for the DriveBy framework"  → **07-evaluation.tex:651-676**

**Meaning:** These look like technical-integration metrics, not standard CI/CD metrics; he lists standard ones (Deployment Frequency, Lead Time for Changes, Change Failure Rate, MTTR, Build Success Rate, Pipeline Duration, Test Pass Rate, Deployment Success Rate, Defect Escape Rate, Security Gate Pass Rate) to adopt or to motivate renaming the section. Unchanged in live source.

### 🔴 #366 (p172) — open
**Kritikos:** `Can size-related metrics be considered as CI/CD integration metrics? Could they also be measured for a manual-based approach?`

**Anchored on:** "CLI binary size"  → **07-evaluation.tex:666-667**

**Meaning:** Questions whether size metrics qualify as CI/CD integration metrics and whether they have a manual-approach counterpart for comparison — regroup or justify. Unchanged.

### ✅ #367 (p172) — already-addressed
**Kritikos:** `Again, we have a non-existing reference number`

**Anchored on:** "added in direct response to the supervisor's round-1 review [550], which noted"  → **07-evaluation.tex:685**

**Meaning:** '[550]' is a non-existent reference; drop it and the round-1 provenance. Resolved: live line 685 motivates the experiment self-containedly ('tests whether the feedback from SDT validation indeed leads to concrete specification corrections') with no bracketed number.

### 🔴 #368 (p173) — open
**Kritikos:** `Do you have the prompt? If yes, please document it somewhere (either in Appendix here or in DriveBy code)`

**Anchored on:** "(Claude Code, Sonnet 4.6) is instructed to revise the specification with"  → **07-evaluation.tex:698**

**Meaning:** If the agent prompt exists, document it (thesis appendix or the DriveBy repo) for reproducibility. No prompt is documented in the thesis or visible under results/sdt-feedback/ (only INDEX.md, petstore-baseline, petstore-revised).

### 🔴 #369 (p173) — open
**Kritikos:** `Question: if it had in its possession the definition of the principles, maybe it could lead to better reconciliation actions done by the agent employed? This might enable to reduce the number of rounds required to produce a proper OpenAPI spec that passes all static validation.Please think about it. It does not require to be done, of course, now, even if decided to do it - could be left as part of future work ...`

**Anchored on:** "the agent may not access the web, the DriveBy CLI source, the thesis"  → **07-evaluation.tex:698,743**

**Meaning:** Suggestion: giving the agent the principle DEFINITIONS (not only the validation report) might yield better reconciliation in fewer rounds; he explicitly allows deferring this to future work. No such future-work note has been added to the design or interpretation subsections.

### ❓ #370 (p174) — unclear
**Kritikos:** `addressing`

**Anchored on:** "✗ (P002/P003 rows of Table 7.18)"  → **07-evaluation.tex:720-722**

**Meaning:** Single word 'addressing' on the delta table — most plausibly a wording suggestion for a Notes cell (e.g., 'examples added' / '4xx/5xx wired' → 'addressing ...'), but the intent cannot be decoded with confidence from one word on a table cell. Candidates: reword the P002 or P003 Notes cell, or caption wording.

### 🟡 #371 (p174) — needs-author-decision
**Kritikos:** `Question / Food for thought:(a) how the agent can now which operations are to be protected and which not?(b) how the agent can now which errors are to be produced by each operation in the OpenAPI spec?This introduces the following issue: should the agent perform the changes on the OpenAPI spec irrespective of the content of the service's source code?This can create a hallucinated OpenAPI spec, which does not correspond to the actual API implementation. So, maybe the agent should also take into account the API's source code?Or you believe that this is not needed as the functional testing can reveal discrepancies?In my view, this is not always possible and depends on the test coverage: does functional testing includes test cases that cover all possible error situations for each specific API operation?If yes, then source code inspection might not be needed.The same holds for security: if an operation does not require authenatication and another requires it, would this be unveiled through functional testing? I am not so sure about this ...`

**Anchored on:** "P005 Security (row of Table 7.18)"  → **07-evaluation.tex:739-747 (interpretation)**

**Meaning:** Food-for-thought: how can the agent know which operations should be protected or which errors each operation really produces without reading the service source code? Report-only revision risks a hallucinated spec diverging from the implementation, and functional testing only catches that if coverage exercises every error/auth path. He asks the author to consider source-aware revision (or argue P006 suffices) — explicitly a question, not a directive.

### ✅ #372 (p175) — already-addressed
**Kritikos:** `ok, let's leave this as future work ...`

**Anchored on:** "The DriveBy report from the post-revision run is,"  → **07-evaluation.tex:743,747**

**Meaning:** 'ok, let's leave this as future work' — he accepts that the iterative validate-remediate loop extension stays future work. The live interpretation already frames the multi-iteration study as future work recorded in the conclusion chapter.

### 🔴 #373 (p175) — open
**Kritikos:** `Isn't this also covered by the PoC and its integration CI/CD metrics? You did not mention the latter ...`

**Anchored on:** "tive pipeline provisioning) is covered by the operational PoC"  → **07-evaluation.tex:754,768**

**Meaning:** RQ3 is also covered by the PoC's CI/CD integration metrics, which the summary never mentions — add them as RQ3 evidence. The live summary's RQ3 result (line 768) still omits the CI/CD metrics.

### 🔴 #374 (p175) — open
**Kritikos:** `Maybe this is also partially covered by PoC and Multi-API controlled evaluations? In the second case, we have definitely the definition of quality gated spec which is respected by the proposed system? In fact, this is not just in terms of constructing the quality-gated workflow but also properly executing and delivering well-expected results based on the way the evaluation was designed (e.g., gate failures in expected cases/APIs)?`

**Anchored on:** "RQ4 (XSDLC as Kubernetes-native ontological specification)"  → **07-evaluation.tex:754**

**Meaning:** Disputes excluding RQ4 from the evaluation summary: the PoC and multi-API evaluation partially evidence RQ4 — the declared quality-gated specification was respected, executed, and delivered the expected gate failures. The live summary still says RQ4 'is intentionally absent from this summary'.

### 🔴 #375 (p176) — open
**Kritikos:** `Please check if this rate/percentage is eventually correct`

**Anchored on:** "single-check defects were detected (80%); three of three combinations were"  → **07-evaluation.tex:758**

**Meaning:** Check whether 80% (12/15) is eventually correct: the body reports two genuine misses plus one loader-rejected mutation (argued to be effectively caught), so 'three undetected ... all reflect known framework limitations' is doubly inconsistent — the count story and the limitation attribution (the p005 miss being an evaluation artefact per idx 315). The live summary sentence is unchanged.

### ✅ #376 (p176) — already-addressed
**Kritikos:** `correct although it is not clear how P7 that is warning blocks the API at the production stage.`

**Anchored on:** "runtime layer on P007. Each API was blocked by the principle corresponding"  → **07-evaluation.tex:760 (severity at :502)**

**Meaning:** Correct, but unclear how P007 — then labelled warning — can block at the production stage. Resolved by the P007→critical flip: the severity model now declares P006/P007 critical and a P007 breach always blocking.

### 🔴 #377 (p176) — open
**Kritikos:** `what's 418?`

**Anchored on:** "(broken-api's 418, slow-api's"  → **07-evaluation.tex:762**

**Meaning:** 'what's 418?' — the summary uses the bare number; spell out that broken-api returns HTTP 418 instead of the documented 200 (or gloss it), since a summary reader won't recall the earlier explanation. Live line 762 still says "broken-api's 418" unexplained.

### 🔴 #378 (p176) — open
**Kritikos:** `Forgot to say sth about OpenAPI 2.x dataset and its evaluation ....`

**Anchored on:** "no API in the sample passed more than 2/6 principles. The"  → **07-evaluation.tex:764**

**Meaning:** The summary's quality-gap result omits the OpenAPI/Swagger 2.x dataset and its evaluation (Section 7.5.4: 80% scored 1/6, pattern reproduces) — add it. Live summary result 4 (line 764) still reports only the 50-spec 3.x arm.

### ✅ #379 (p177) — already-addressed
**Kritikos:** `maps to ....`

**Anchored on:** "RQ2 explicitly asks whether the validation output is "actionable, machine-"  → **07-evaluation.tex:766**

**Meaning:** 'maps to ....' — wording: say the result 'maps to' RQ2. Resolved: live line 766 now reads 'The result maps to RQ2 directly'.

### 🔴 #380 (p177) — open
**Kritikos:** `Did not say sth about CI/CD integration metrics`

**Anchored on:** "operational edge cases passed; the sole failure (the health-check timeout) was"  → **07-evaluation.tex:768**

**Meaning:** The RQ3 summary result says nothing about the CI/CD integration metrics (startup, validation, provisioning, gate latency) — add them as quantitative integration evidence (same gap as idx 373). Live line 768 unchanged.


## 08-ai-assisted-development.tex — 33 comments (24 open)

### 🔴 #381 (p179) — open
**Kritikos:** `Just one quick question at this point: did you use any code assistant in the IDE that accelerated somehow the development (e.g., through the use of specific prompts or the use of auto-completion features)?Or this was pure manual work?`

**Anchored on:** "they are entirely the author's work."  → **08-ai-assisted-development.tex:22**

**Meaning:** Kritikos asks whether the 350-day 'solo' Phase A really involved no AI at all: did the author use an IDE code assistant (auto-completion, Copilot-style prompts) that accelerated development, or was it pure manual work? Implies adding an explicit sentence to the Phase A description stating whether any IDE-level AI assistance was used.

### 🔴 #382 (p179) — open
**Kritikos:** `+ principle checking finalised and the DriveBy tool was fully produced?`

**Anchored on:** "the GitOps pipeline was wired together, OSS-polish cleanup"  → **08-ai-assisted-development.tex:24**

**Meaning:** The '+' marks a suggested addition to the Phase B activity list: '+ principle checking finalised and the DriveBy tool was fully produced?'. He is asking whether principle checking was finalised and DriveBy fully completed during Phase B, and wants the Phase B list to say so explicitly, making the Phase A/B completion boundary unambiguous.

**⚠ Triage correction:** Earlier triage anchored this to the Phase A sentence 'the first principle checkers were written'; the icon actually sits in the Phase B activity list (line 24), where Kritikos proposes adding 'principle checking finalised and the DriveBy tool fully produced' as Phase B items.

### 🔴 #383 (p180) — open
**Kritikos:** `Why there was a need to remove lines?`

**Anchored on:** "Lines removed"  → **08-ai-assisted-development.tex:46**

**Meaning:** Looking at Table 8.1's 'Lines removed 21,330 / 7,678' row, he asks why ~29k lines were deleted across the project. Implies adding a sentence of prose explaining the removals (v3.0.0 restructure, samples/ removal, single-repo to two-repo redesign) so the deletions read as design iteration, not unexplained churn.

### 🔴 #384 (p181) — open
**Kritikos:** `You do not provide any relevant percentage concerning thesis checking, proofreading and correction. Maybe this corresponds to some percentages in the above table but this requires to be well covered. As you need to prove authorship not just with respect to the code but also to the report. Both are the two main deliverables of this thesis!!!`

**Anchored on:** "every thesis chapter was read end-to-end and edited before being pushed"  → **08-ai-assisted-development.tex:90**

**Meaning:** He observes that no effort percentage is given for thesis checking, proofreading, and correction — the effort table (tab:effort-allocation, lines 67-84) only quantifies code/productisation activity. Since both code AND report are the thesis deliverables, he demands an explicit, measured report-side effort breakdown to prove authorship of the report deliverable too.

### ✅ #385 (p181) — already-addressed
**Kritikos:** `the thesis report structure`

**Anchored on:** "the evaluation design, the chapter structure), the quantitative line-level evidence"  → **08-ai-assisted-development.tex:97**

**Meaning:** Suggested word replacement: change 'the chapter structure' to 'the thesis report structure' in the conceptual-ownership parenthetical, making explicit that the report deliverable's structure is author-owned.

### 🔴 #386 (p181) — open
**Kritikos:** `Could explain in the text somewhere how debugging was actually carried out.Did you produce test code? Or was it manually conducted based on your experience?`

**Anchored on:** "(au-thor) — wrapped row label in the Phase-B effort table"  → **08-ai-assisted-development.tex:79**

**Meaning:** Prompted by the table row 'Manual code changes, debugging, hot-fixes (author) 15%', he asks the text to explain HOW debugging was actually carried out: did the author write test code, or was it manual debugging based on experience? Implies adding 2-3 sentences describing the debugging method (Go test suite vs manual reproduction vs JSON inspection).

### 🔴 #387 (p182) — open
**Kritikos:** `So, the revisions were done by you?I am asking this as it seems that the chapter content was also revised by the AI and was then proofread by you ...`

**Anchored on:** "initial drafts of several thesis sections that were then reviewed and edited by the author"  → **08-ai-assisted-development.tex:97**

**Meaning:** He asks pointedly: were the revisions done by you? The current wording reads as if the AI revised chapter content and the author merely proofread, which he finds academically problematic. Implies rewriting to state unambiguously that the author authored the final content of every chapter, with AI limited to initial drafts.

### ❓ #388 (p182) — unclear
**Kritikos:** `mapping to`

**Anchored on:** "both author-driven decisions made before Phase B:"  → **08-ai-assisted-development.tex:101**

**Meaning:** One-phrase comment 'mapping to' whose intent cannot be pinned with confidence. Candidate readings: (a) a wording suggestion for the nearby 'the author's role would map most cleanly to architect-and-tech-lead' (line 99), rephrasing 'would map most cleanly to' as 'mapping to'; (b) a suggested insertion into 'both author-driven decisions made before Phase B, mapping to ...'. Needs the author to check the PDF placement.

**⚠ Triage correction:** Earlier triage anchored this to the 'reviewer-of-record ... claims overall authorship' sentence (line 97); the icon actually sits on 'both author-driven decisions made before Phase B:' (line 101), so the triage's target phrase was wrong and the intent remains ambiguous.

### 🔴 #389 (p183) — open
**Kritikos:** `of what?`

**Anchored on:** "distributed across the monorepo."  → **08-ai-assisted-development.tex:146**

**Meaning:** 'of what?' sits on 'across the monorepo' — he asks which monorepo / the monorepo of what. Implies specifying it: 'distributed across the DriveBy project monorepo' (first mention in this chapter should name the repository).

**⚠ Triage correction:** Earlier triage read 'of what?' as querying 'most reusable methodological contribution ... of what?'; the icon actually sits on 'across the monorepo', so the question is which/whose monorepo, not what the contribution belongs to.

### ✅ #390 (p183) — already-addressed
**Kritikos:** `again this is a non-existing reference number`

**Anchored on:** "The supervisor noted ([534]) that repetition"  → **08-ai-assisted-development.tex:188**

**Meaning:** '[534]' is a stale review-round reference number that resolves to nothing in the bibliography; he wants all such '(The supervisor noted ([NNN]))' constructions removed and rewritten as self-contained prose. The live text now reads 'Repetition is in tension with the layered-loading model' with no bracket — this and all other [5NN] brackets are gone from the chapter (grep confirms zero remaining).

### 🔴 #391 (p183) — open
**Kritikos:** `Although the supervisor ..., the ...You need to indicate what was suggested and why you did not follow this based on an actual observation made during the project running ...`

**Anchored on:** "The author's choice was to repeat"  → **08-ai-assisted-development.tex:188**

**Meaning:** He wants the Containment-vs-repetition paragraph restructured as 'Although the supervisor suggested X, the author did Y': state concretely what was suggested AND justify not following it with an actual observation made while running the project, not just an abstract principle. The live text now states the suggestion plainly (supervisor bracket removed) but still justifies the choice only by a hypothetical truncation argument, not an observed incident.

### 🔴 #392 (p186) — open
**Kritikos:** `flagging as stubs chapters that are under 20 lines?Was that what you wanted to say here?`

**Anchored on:** "thesis chapter status (flagging stubs under 20 lines)"  → **08-ai-assisted-development.tex:221**

**Meaning:** He finds 'flagging stubs under 20 lines' ambiguous and asks whether the intended meaning is 'flagging as stubs any chapter under 20 lines'. Implies rewording the parenthetical so the subject (chapter files) and the criterion (under 20 lines) are explicit, verified against tools/check-docs.sh.

### 🔴 #393 (p187) — open
**Kritikos:** `writes`

**Anchored on:** "the author authors every entry, the agent obeys"  → **08-ai-assisted-development.tex:245**

**Meaning:** Suggested word replacement: change the awkward 'the author authors every entry' to 'the author writes every entry'. The live text still says 'the author authors'.

**⚠ Triage correction:** Earlier triage read 'writes' as a verb-agreement note on 'does not write to memory' and called it likely already resolved; the icon actually spans 'the author authors every entry' — it is a replacement suggestion (authors -> writes) and is still unapplied.

### 🔴 #394 (p187) — open
**Kritikos:** `ok but there were multiple evaluations, not just one.Were all of them designed by you then?`

**Anchored on:** "The large-scale evaluation against APIs.guru was designed by the author, tooled"  → **08-ai-assisted-development.tex:252**

**Meaning:** He notes there were multiple evaluations (controlled, large-scale, operational, agent-feedback), not just the APIs.guru one, and asks whether ALL of them were designed by the author. Implies generalising the sentence to claim authorship of every evaluation arm. The new Demarcation subsection (line 125) does credit the author with the full three-arm design plus the agent-feedback experiment, but the flagged sentence itself still names only the large-scale study.

### 🔴 #395 (p188) — open
**Kritikos:** `Why is there a need to validate reachability of the endpoints?Isn't the OpenAPI specs that interest you more?Or you wanted to produce evaluation results that reflect the actual reality concerning non outdated APIs?`

**Anchored on:** "The prober validates reachability by testing simple GET endpoints."  → **08-ai-assisted-development.tex:259**

**Meaning:** He asks why reachability validation is needed at all when the static OpenAPI spec is the object of interest — and offers the answer himself: was it to ensure results reflect live, non-outdated APIs? Implies adding one justifying sentence (reachability filters out abandoned specs and is a precondition for the functional/performance arms).

### 🔴 #396 (p189) — open
**Kritikos:** `ok but it is not clear who did the injections, i.e., the implementation and re-implementation of that evaluation (controlled evaluation of non-critical API).`

**Anchored on:** "and addresses it as part of the controlled-evaluation rerun in Chapter 7"  → **08-ai-assisted-development.tex:277**

**Meaning:** He says it is unclear WHO did the defect injections — i.e., who implemented and re-implemented the controlled evaluation of the non-critical API. Implies stating explicitly whether the author wrote the injection harness or the agent did under author specification. The paragraph was rewritten today (supervisor bracket removed, rerun now described as the evaluation itself) but still does not attribute the injection implementation.

### 🔴 #397 (p189) — open
**Kritikos:** `ok but ideally I would expect to see you conduct the writing of the thesis report. It is ok if the agent produces part of the code. However, concerning the thesis report, it is usually assumed that the thesis is mainly written by the author/student with the AI being used to get some ideas and do some text polishing. Nothing more than that. In the current case, it could also be acceptable if only the initial draft of the thesis was written based on the actual code situation and then you took over completely to produce the full content of each report chapter.`

**Anchored on:** "the thesis was written by"  → **08-ai-assisted-development.tex:284**

**Meaning:** His core academic-acceptability position: the thesis report should be mainly written by the student, with AI only for ideas and polishing — OR, acceptably, the AI writes only an initial code-grounded draft after which the author takes over completely to produce the full content of each chapter. He wants the section's framing to commit to that acceptable model. The live opening (line 284, supervisor bracket removed) still hedges with 'partially correct but not complete'.

### 🔴 #398 (p190) — open
**Kritikos:** `Ok, this is the most crucial point here. This makes it acceptable in academic terms. I would urge to clarify this from the very beginning of this section. See also a previous comment of mine!`

**Anchored on:** "Substantive sections were fully rewritten by the author after the draft."  → **08-ai-assisted-development.tex:298**

**Meaning:** He says this sentence is THE crucial point that makes the workflow academically acceptable, and urges that it be stated at the very beginning of the section instead of being buried as item 4 of the protocol list. The live text still has it at the end of the enumerate (line 298), not in the section opening (line 284).

### 🔴 #399 (p191) — open
**Kritikos:** `you need to be precise here. It should be one API based also on content of Section 7.7`

**Anchored on:** "one or two real-world public APIs with low P002–P004 scores"  → **08-ai-assisted-development.tex:340**

**Meaning:** He demands precision consistent with Section 7.7: the experiment used exactly ONE API. Chapter 7's sec:sdt-feedback-experiment confirms a single subject (Swagger Petstore). The live Chapter 8 sentence still says 'one or two real-world public APIs', so it must be changed to the one named API.

### 🔴 #400 (p192) — open
**Kritikos:** `ok but you do not explain what was the result and whether it is positive/affirmative or not ...`

**Anchored on:** "The result of this experiment converts the theoretical connection of this section into an empirical one."  → **08-ai-assisted-development.tex:340**

**Meaning:** He notes the text never states what the result actually WAS or whether it was positive. The experiment has been run (Petstore 1/6 baseline improved to 4/6 in a single iteration, per Chapter 7), so the sentence should report the affirmative outcome and magnitude, not just promise that a result exists.

### 🔴 #401 (p192) — open
**Kritikos:** `Sure but in principle the reconciliation might not reflect the reality even in this case: maxLength and pattern shouldn't be confined by the implementation itself?Otherwise, the reconciliation will tend to create a new version of the spec that could succeed in passing the principles but does not reflect reality and that latter issue could not be unveiled by functional testing ...This is really a bad scenario as you are producing an OpenAPI spec that can lead to a wrong usage of the respective API by its clients. This means that the derived OpenAPI spec should not be included in a production environment!This is an issue that of course needs to be deal with in future work ...`

**Anchored on:** "add maxLength and pattern properties to the field's schema."  → **08-ai-assisted-development.tex:349**

**Meaning:** Substantive technical objection: an agent inventing maxLength/pattern values not grounded in the actual implementation can produce a spec that passes the principles but misrepresents reality — undetectable even by functional testing — and such a derived spec would mislead API clients and must NOT go to production. He wants an explicit limitation/future-work caveat acknowledging this. No such caveat exists in the live subsection.

### 🔴 #402 (p192) — open
**Kritikos:** `Could refer to specific fix directions that were given in the context of Section 7.7. And could stress that here ...`

**Anchored on:** "the remediation"  → **08-ai-assisted-development.tex:349**

**Meaning:** He suggests grounding the abstract SuggestedFix discussion in the specific fix directions actually produced and acted on in the Section 7.7 experiment, and stressing here that those were real executed remediations. No cross-reference to 7.7's concrete fixes exists in the live subsection.

### ❓ #403 (p192) — unclear
**Kritikos:** `to be ...`

**Anchored on:** "for the latter, distinct from the former. This section provides them."  → **08-ai-assisted-development.tex:358**

**Meaning:** Fragmentary comment 'to be ...' on the Personal Lessons intro sentence. Candidate readings: a started-but-unfinished wording suggestion (e.g. '... are to be provided' or smoothing 'distinct from the former'), or the beginning of a longer remark he abandoned. The sentence was rewritten today (supervisor bracket removed; now 'This section provides the latter, distinct from the former'), which may have mooted it, but the original intent cannot be decoded.

### 🔴 #404 (p192) — open
**Kritikos:** `caters for conforming to this requirement by separating these two forms of lessons learned ...Do you really cover both of them here?Or only the personal lessons learned experience?`

**Anchored on:** "This section provides them."  → **08-ai-assisted-development.tex:358**

**Meaning:** He asks whether the chapter genuinely covers BOTH forms of lessons learned — generalisable framework lessons AND personal lessons — or only the personal ones. The live section explicitly delivers only the personal lessons; implies either adding the framework-level lessons or explicitly pointing to where they are treated (e.g. the Discussion chapter).

### 🟡 #405 (p193) — needs-author-decision
**Kritikos:** `correct. But the AI couldn't supply some nice ideas that could be implanted in the original design?Maybe this could be also mentioned ...`

**Anchored on:** "the up-front design work is the thing that makes later acceleration possible."  → **08-ai-assisted-development.tex:360**

**Meaning:** He accepts the claim ('correct') but asks whether the AI couldn't have supplied some good ideas worth implanting in the original design, and suggests mentioning that. The author must decide whether to concede that the agent contributed any design-level ideas (weakening the 'no AI influence on design' claim at line 108) or hold the line with an explicit acknowledgement of tactical-only suggestions.

### ✅ #406 (p193) — already-addressed
**Kritikos:** `follow-on vs follow-up ...`

**Anchored on:** "In a follow-on project,"  → **08-ai-assisted-development.tex:360**

**Meaning:** Terminology flag: 'follow-on' should be 'follow-up'. The live chapter has been globally swept — every former 'follow-on' (lines 4, 340, 360, 362, 364, 366, 368) now reads 'follow-up'; grep confirms zero 'follow-on' remaining.

### ✅ #407 (p193) — already-addressed
**Kritikos:** `follow-up?`

**Anchored on:** "a chat session dies with the session; ... In a follow-on project"  → **08-ai-assisted-development.tex:362**

**Meaning:** Another 'follow-up?' flag on the 'follow-on project' in the Persisting-corrections lesson. The live text now reads 'In a follow-up project' — covered by the global swap.

### ✅ #408 (p193) — already-addressed
**Kritikos:** `follow-up`

**Anchored on:** "follow-on project budget time explicitly for verification"  → **08-ai-assisted-development.tex:364**

**Meaning:** Another 'follow-up' replacement flag, on the Code-review-bottleneck lesson. The live text now reads 'in a follow-up project budget time' — addressed by the global swap.

### 🔴 #409 (p193) — open
**Kritikos:** `?`

**Anchored on:** "even if that means a quieter sprint."  → **08-ai-assisted-development.tex:364**

**Meaning:** A bare '?' on the phrase 'a quieter sprint' — he does not understand or objects to the informal expression. Implies rephrasing in plainer academic language (e.g. 'even if that reduces the volume of output per iteration'). The phrase is unchanged in the live text.

**⚠ Triage correction:** Earlier triage guessed this was another follow-on/follow-up flag resolved by the terminology sweep; the icon actually sits on 'that means a quieter sprint', so it is a clarity query on that phrase and is still open.

### 🔴 #410 (p193) — open
**Kritikos:** `sufficient`

**Anchored on:** "budget time explicitly for verification before each session"  → **08-ai-assisted-development.tex:364**

**Meaning:** Single-word insertion suggestion: 'sufficient' — i.e. 'budget sufficient time explicitly for verification before each session'. The icon sits exactly on 'project budget time explicitly...', making the insertion point clear. Unapplied in the live text.

**⚠ Triage correction:** Earlier triage could not pin the target and speculated about 'enough -> sufficient' elsewhere; the icon position pins it to 'budget time', so the reading is 'budget sufficient time'.

### ✅ #411 (p193) — already-addressed
**Kritikos:** `follow-up`

**Anchored on:** "In follow-on projects the author would adopt this pattern by default"  → **08-ai-assisted-development.tex:366**

**Meaning:** Another 'follow-up' terminology flag, on the thesis-as-artifact lesson. The live text now reads 'In follow-up projects' — addressed by the global swap.

### 🔴 #412 (p193) — open
**Kritikos:** `follow-up project are planned`

**Anchored on:** "a follow-on project: (i) record exact wall-clock time per activity"  → **08-ai-assisted-development.tex:368**

**Meaning:** Suggested rewording: 'follow-up project are planned' — i.e. beyond the follow-on->follow-up swap (done), he proposes phrasing like 'Three specific changes for a follow-up project are planned:', stating the changes as planned intent rather than hypothetical. The term swap is in the live text but the 'are planned' phrasing is not.

### 🔴 #413 (p193) — open
**Kritikos:** `Did you develop other skills and knowledge that could be exploited in the future?Other best practices that you could follow, irrespective of agentic development?`

**Anchored on:** "about how to run the next one."  → **08-ai-assisted-development.tex:370**

**Meaning:** He asks whether the author developed OTHER transferable skills, knowledge, or best practices exploitable in the future irrespective of agentic development — all five current lessons are AI-tooling-specific. Implies adding at least one AI-independent lesson (e.g. up-front architecture discipline, documentation as a CI-gated artifact, version-control authorship traceability). Not present in the live text.


## 09-discussion.tex — 35 comments (29 open)

### 🔴 #414 (p195) — open
**Kritikos:** `I believe that again we have the issue that the question content is not the one original supplied in Chapter 1, Section 1.5. Please make it identical ...`

**Anchored on:** "Can a methodology grounded in specification completeness, behavioral determinism, and operational observability..."  → **09-discussion.tex:16**

**Meaning:** The RQ1 question text quoted here differs from the official wording in Chapter 1 Section 1.5; he wants the two made character-identical. The divergence is now even larger: Chapter 1 (01-introduction.tex:129) was reworded today-ish to 'Can a conceptual framework ... operate without dedicated testers in the promotion pipeline?' while this chapter still says 'Can a methodology ... comparable to manual approaches?'. Same identity check applies to RQ2-RQ4 prompts.

### 🔴 #415 (p195) — open
**Kritikos:** `100% or 80%`

**Anchored on:** "demonstrated 100% detection accuracy for injected defects"  → **09-discussion.tex:18**

**Meaning:** '100% or 80%?' — he questions whether the 100% headline is honest, since elsewhere the evaluation reports misses (the per-check extension detected 13/15 single-check defects; line 108 of this chapter concedes two undetected defects). Implied change: scope the 100% to the 7-defect principle-level arm and state the per-check figure honestly, so the chapter does not contradict itself.

### 🔴 #416 (p195) — open
**Kritikos:** `Could also say sth about the extension of the controlled experiment to cover all checkers in all principles and combination of checkings (across different principles).`

**Anchored on:** "confirming the independence of the principle checkers."  → **09-discussion.tex:18**

**Meaning:** Suggests the RQ1 answer also mention that the controlled experiment was extended to cover all checkers in all principles and combinations of checks across different principles. The extension exists (Ch.7 per-check extension, 3 pairs + 1 triple) but the RQ1 paragraph never cites it; add a sentence referencing subsec:per-check-extension.

### 🔴 #417 (p195) — open
**Kritikos:** `Maybe the name/title is misleading here as you are proposing a framework and not a whole methodology. So, a better title would have been Framework Soundness. Similar fixes need to be done within the question content.`

**Anchored on:** "RQ1: Methodology Soundness"  → **09-discussion.tex:12**

**Meaning:** The thesis proposes a framework, not a whole methodology, so the subsection title should be 'RQ1: Framework Soundness', with matching fixes inside the question text. Chapter 1 already concedes 'conceptual framework rather than a complete methodology', so this heading and line 16 are now inconsistent with the conceded terminology.

### 🔴 #418 (p196) — open
**Kritikos:** `Ok but the beyond theoretical scaffolding was eventually covered only for the first axiom?Or all of them? Did not mention something about the other two.`

**Anchored on:** "Determinism was necessary because the quality gate operates in an automated loop"  → **09-discussion.tex:22**

**Meaning:** He asks whether 'more than theoretical scaffolding' was actually demonstrated for all three axioms or only the first: Completeness is backed by evaluation evidence ('the evaluation revealed...'), but Determinism and Observability get only design-rationale arguments ('was necessary because...'), not empirical demonstration. Implied change: cite the empirical evidence for Determinism (repeated identical runs, Ch.7 line 246) and Observability (agent JSON consumption, Ch.8/agent-feedback experiment).

**⚠ Triage correction:** Earlier triage claimed the paragraph 'actually does address all three axioms' and treated the comment as stale; the paragraph is unchanged from the PDF and Kritikos's real point is that only Completeness has evaluation evidence — the other two are asserted, not shown.

### ✅ #419 (p196) — already-addressed
**Kritikos:** `are you sure about this? I had the impression that the first has other issues than the "security" one while the second had only "security" issues ...`

**Anchored on:** "both lacking security definitions) were blocked at the staging gate"  → **09-discussion.tex:24**

**Meaning:** He doubts that bad-docs-api and no-auth-api both fail on security: bad-docs-api fails on documentation issues, only no-auth-api is the security failure. Today's rewrite of this paragraph now attributes each API correctly: bad-docs-api blocked on P002–P004 (documentation-class), no-auth-api on P005 (security).

### ✅ #420 (p196) — already-addressed
**Kritikos:** `they did not have any kind of warning or issue in terms of static validation. That is why they passed the staging gate.`

**Anchored on:** "APIs with only warning-severity failures (slow-api) passed staging but were caught"  → **09-discussion.tex:24**

**Meaning:** He corrects the mechanism: slow-api (and broken-api) passed staging because they had no static-validation issues at all, not because their failures were merely 'warning-severity'. Today's rewrite now says 'The APIs whose specifications are clean but whose behaviour is defective passed static validation and were caught at the runtime layer', which is exactly his point.

### ✅ #421 (p196) — already-addressed
**Kritikos:** `You have to better clarify the exact layers and their exact hierarchy. We have static validation, then functional testing and finally non-functional/performance testing.`

**Anchored on:** "The broken-api, which passed all static checks but diverged from its specification at runtime"  → **09-discussion.tex:24**

**Meaning:** Asks for an explicit statement of the validation layers and their hierarchy: static validation, then functional testing, finally non-functional/performance testing. Today's rewrite names the layers and their order (staging-gate static validation for spec defects; runtime layer with functional testing P006 and load testing P007), satisfying the request, though an explicit 'three layers in order' sentence could make it even more direct.

### 🔴 #422 (p196) — open
**Kritikos:** `with respect to what? Or do you mean that they are complementary to each other?`

**Anchored on:** "the nine-principle structure is not redundant but complementary."  → **09-discussion.tex:24**

**Meaning:** 'Complementary with respect to what?' — the word 'complementary' has no object. Implied change: write 'complementary to each other' or 'mutually complementary, each principle covering a failure class the others cannot detect'. Today's rewrite strengthened the preceding clause ('catches exactly the failure class it is designed for') but the dangling 'complementary' survives verbatim.

### 🔴 #423 (p196) — open
**Kritikos:** `an OpenAPI specification`

**Anchored on:** "derive their validation logic entirely from the specification"  → **09-discussion.tex:33**

**Meaning:** One-phrase substitution: the icon sits on 'entirely from the specification' and the comment supplies the replacement — 'derive their validation logic entirely from an OpenAPI specification'. He wants the word OpenAPI made explicit at this spot.

**⚠ Triage correction:** Earlier triage anchored this to the RQ2 question line ('from OpenAPI specifications', line 31) and that line was duly edited today, but the icon actually sits on line 33's 'entirely from the specification', which is still unchanged.

### 🔴 #424 (p196) — open
**Kritikos:** `I believe that quality gates are not produced fully solely from OpenAPI specifications. They also require additional knowledge that is supplied in the CR (XSLDC). Please make the respective correction in the main report text.`

**Anchored on:** "To what extent can validation rules, test cases, and quality gates be derived automatically from OpenAPI specifications"  → **09-discussion.tex:33**

**Meaning:** Pushback on the 'Fully' answer: quality gates are NOT derived solely from the OpenAPI specification — they also need knowledge supplied in the XSDLC custom resource (thresholds, environment chain, gate composition). He explicitly asks for the correction in the main report text: scope 'Fully' to validation rules and test cases, and state that gate policy/parameters come from the XSDLC CR. The single-artifact concession was applied elsewhere (Ch.2 Table) but this RQ2 answer still claims 'Fully, for the implemented principles.'

### 🔴 #425 (p196) — open
**Kritikos:** `In what forms the business logic could be actually described? I am asking this as someone could argue that the textual description of each operation along with the examples and parameter descriptions is enough to convey the business logic.In this sense, is there really an omission here?And what should be supplied in addition?`

**Anchored on:** "cannot test business logic that the specification does not describe."  → **09-discussion.tex:35**

**Meaning:** He challenges the claim: in what forms could business logic be described at all? One could argue operation descriptions, examples, and parameter descriptions already convey business logic — so is there really an omission, and what extra input would be needed? Implied change: distinguish spec-expressible business logic (descriptions, examples, constraints that P006 could in principle assert against) from logic requiring out-of-band oracles (stateful workflows, cross-call invariants), and say precisely what would still need manual authorship.

### 🔴 #426 (p197) — open
**Kritikos:** `Is it Chapter 8 or Chapter 7?In addition, I do not get the point about extending the finding. And which one?The RQ talks about extracting validation rules, tests and gates from OpenAPI specifications.In this sense, how the MultiAPI controlled evaluation supports this question's answering?Does it show the sufficiency of the principle checkers and the tests produced from the OpenAPI specifications?Does it show their completeness? Their accuracy?`

**Anchored on:** "The agent-driven evaluation (Chapter 8) extended this finding."  → **09-discussion.tex:37**

**Meaning:** Three issues: (a) verify the chapter number (7 vs 8 — with the new system-architecture chapter, ch:ai-development now resolves to Chapter 9, so the \ref is self-correcting but should be checked in the build); (b) 'extended this finding' is vague — which finding?; (c) explain how the agent evaluation actually supports RQ2 — does it show sufficiency, completeness, or accuracy of the spec-derived checkers/tests? Implied change: state explicitly that it evidences sufficiency of the structured output for autonomous consumption, and that it does not establish completeness or accuracy.

### ❓ #427 (p197) — unclear
**Kritikos:** `ok but maybe this is fully addressed?In contrast to the previous one?`

**Anchored on:** "the composition complexity tension: the Crossplane composition"  → **09-discussion.tex:52**

**Meaning:** 'OK but maybe this is fully addressed? In contrast to the previous one?' Most likely reading: he asks whether this second tension (composition complexity) is fully addressed/mitigated, unlike the first tension (webhook reliability) which was only partially mitigated — and wants the text to say so. Alternative reading: he is musing that the issue may already be handled elsewhere in the thesis. The paragraph currently ends by calling the complexity 'inherent', with no statement of whether/how it is addressed.

**⚠ Triage correction:** Earlier triage anchored this to the RQ3 verdict ('demonstrated effective integration', line 46) and read it as 'RQ3 is more fully answered than RQ2'; the icon actually sits on the composition-complexity tension paragraph, so the comment is about whether that tension is fully addressed in contrast to the webhook tension.

### 🔴 #428 (p197) — open
**Kritikos:** `could also refer here to the CI/CD metrics and the very good values achieved, much better than manual processes ...`

**Anchored on:** "the potential to replace hours of manual ClickOps. The system was installed"  → **09-discussion.tex:46**

**Meaning:** Suggests citing the CI/CD metrics here (gate latency, provisioning times from Chapter 7's CI/CD Integration Metrics) and the very good values achieved, contrasting them with manual processes, to substantiate the ClickOps-replacement claim with numbers rather than 'potential'.

### 🔴 #429 (p198) — open
**Kritikos:** `Here we talk about another control loop or the existing one?As the focus until now was on the API specification. While here it seems to extend towards XSDLC (resources).`

**Anchored on:** "not yet exercised by autonomous consumers. No production controller currently"  → **09-discussion.tex:67**

**Meaning:** He asks whether programmatic creation/modification of XSDLC resources is the same control loop discussed so far or a new one: the thesis focus until here was the API specification, and this passage extends the scope to XSDLC resources themselves. Implied change: add a bridging clarification that this is the same GitOps reconciliation model applied to a second, parallel specification layer (the pipeline contract), not a different loop.

**⚠ Triage correction:** Earlier triage anchored it to the RQ4 opening sentence (line 63); the icon sits on the limitation paragraph (line 67) about autonomous consumers, so the clarification belongs there, though the interpretation (clarify the API-spec to XSDLC scope shift) was essentially right.

### 🔴 #430 (p198) — open
**Kritikos:** `You have 4 paragraphs here. Could map them to Section 9.2.1 OpenAPI Specifications Quality Gap`

**Anchored on:** "The large-scale evaluation revealed a consistent pattern: structural compliance"  → **09-discussion.tex:74**

**Meaning:** Structural request: the four paragraphs sitting bare under Section 9.2 should be wrapped in their own subsection, e.g. '9.2.1 OpenAPI Specifications Quality Gap', mirroring the existing 9.2.x subsection 'XSDLC as Specification of the Specification'. The live source still has the four paragraphs (lines 74-80) orphaned directly under the section heading.

### 🔴 #431 (p199) — open
**Kritikos:** `Are you sure about this? This changed based on the new evaluation.`

**Anchored on:** "fail P002 (Documentation Quality), P003 (Error Handling), P004 (Schema Definitions), and P008 (Versioning Strategy)"  → **09-discussion.tex:74**

**Meaning:** 'Are you sure about this? This changed based on the new evaluation.' — the universal-fail principle list is stale relative to the regenerated evaluation, which reports P002/P003/P008 universal fail and P004/P005 NEAR-universal fail (this chapter's own line 110 carries the new pattern). Implied change: update line 74 to match the round-1-revision numbers (P004 demoted to near-universal, P005 added).

### 🔴 #432 (p199) — open
**Kritikos:** `correct + seems that structural rules are not sufficient for full structural validation`

**Anchored on:** "do not assess whether the specification provides the information a consumer needs"  → **09-discussion.tex:76**

**Meaning:** 'Correct +' — he agrees, and adds a reinforcing observation to include: structural linting rules are not even sufficient for full STRUCTURAL validation (the form/substance gap is wider than documentation alone). Implied change: append a sentence to the Spectral comparison making this point.

### 🔴 #433 (p199) — open
**Kritikos:** `ok but mainly functional. Does not cover performance-oriented behaviour!`

**Anchored on:** "Where the OpenAPI document specifies API behaviour"  → **09-discussion.tex:87**

**Meaning:** Qualify 'API behaviour': the OpenAPI document specifies mainly FUNCTIONAL behaviour and does not cover performance-oriented behaviour. Implied change: '...specifies the API's functional behaviour...' plus an acknowledgment that the OpenAPI contract is silent on performance (which is why thresholds live in XSDLC).

### 🟡 #434 (p199) — needs-author-decision
**Kritikos:** `correct but the question is whether it should also carry non-functional constraints like performance ones.In my opinion, these constraints are part of non-functional requirements that should be part of the OpenAPI specification/contract (potentially through an extension). Please think about it ...`

**Anchored on:** "the XSDLC document specifies gating policy"  → **09-discussion.tex:87**

**Meaning:** Substantive design pushback: he believes non-functional constraints (e.g. performance thresholds) are part of the API's non-functional requirements and should arguably live IN the OpenAPI specification/contract, potentially via an x- extension, rather than only in the XSDLC CR. 'Please think about it' — the author must either defend the current split or flag an OpenAPI performance extension as future work. This decision also determines the single-vs-dual-artifact framing.

### 🔴 #435 (p200) — open
**Kritikos:** `ok but at one contribution we have OpenAPI specification validation. At the second, we have just a declarative specification or something additional? Or the extras will be part of future work?`

**Anchored on:** "a pipeline that is itself a declarative specification---rather than"  → **09-discussion.tex:91**

**Meaning:** He asks what the second (XSDLC) contribution adds beyond being 'just a declarative specification' — is there something additional now, or are the extras future work? Implied change: make the present-vs-future boundary explicit (delivered today: discoverability, inspectability, composability, audit trail; deferred: autonomous controllers managing XSDLC). Today's edit only changed the framing phrase ('a distinct contribution'); the present/future boundary is still implicit.

### 🔴 #436 (p200) — open
**Kritikos:** `On the other hand, this means that`

**Anchored on:** "But it means that XSDLC's"  → **09-discussion.tex:101**

**Meaning:** A wording substitution: the comment text 'On the other hand, this means that' is the replacement for the phrase the icon sits on — change 'But it means that XSDLC's quality assurance is scoped...' to 'On the other hand, this means that XSDLC's quality assurance is scoped...'.

**⚠ Triage correction:** Earlier triage anchored this to line 89's inspectability sentence and read it as a request to complete a trailing contrast; the icon actually sits on 'But it means that XSDLC's' in the BYOCI trade-off paragraph (line 101), and the comment is simply the suggested replacement wording.

### 🔴 #437 (p201) — open
**Kritikos:** `In principle, someone could inspect existing OpenAPIs in order to collect these defects and then utilise them in an evaluation ....Selecting specific defects could create a bias towards the success of already implemented checkers.`

**Anchored on:** "The defect set was hand-designed and may still understate the long tail of naturally occurring defects"  → **09-discussion.tex:108**

**Meaning:** He strengthens the internal-validity concession: hand-selecting defects biases the evaluation toward the success of already-implemented checkers; an unbiased design would mine real defects from existing public OpenAPI specs and evaluate against those. Implied change: name this selection bias explicitly and note real-defect mining as the mitigation. The live sentence concedes hand-design but not the bias-toward-checker-success point.

### 🔴 #438 (p201) — open
**Kritikos:** `I would also add the issue that functional testing relies on structural validation over responses - thus output structural validity is applied.While true functional testing would require also checking the output content. This creates a bias towards successful functional testing even in cases where it should have failed. Thus, this creates a bias towards higher detection rates.`

**Anchored on:** "Internal validity. The controlled evaluation uses a single reference API"  → **09-discussion.tex:108**

**Meaning:** Add a missing internal-validity threat: P006 functional testing only applies STRUCTURAL validation over responses (schema conformance), not content checking; true functional testing would also verify output content. A content-wrong but shape-correct response passes, biasing detection rates upward. This concession is absent from the live Threats section.

### 🔴 #439 (p201) — open
**Kritikos:** `a random selection strategy would have been better.Which guarantees a balanced representation of both big and smaller API providers.`

**Anchored on:** "toward providers whose names sort early (Adyen, Amazon Web Services, Amadeus)"  → **09-discussion.tex:110**

**Meaning:** A random selection strategy over the APIs.guru population would have been better than alphabetical, guaranteeing balanced representation of big and small API providers. Implied change: state random sampling as the correct remedy / future-rerun strategy in the external-validity paragraph.

### 🔴 #440 (p201) — open
**Kritikos:** `I am not sure about this: if you take into consideration a specific non-negligent percentage of surely poor OpenAPI specifications, this creates a bias towards poor specifications and the specific relevant evaluation results that have been produced.`

**Anchored on:** "this likely strengthens rather than weakens the thesis result"  → **09-discussion.tex:110**

**Meaning:** He disputes the one-sided 'strengthens' claim: including a non-negligible share of known-poor specifications biases the results toward poor-spec findings and may inflate the apparent quality gap. Implied change: present it as a two-sided threat (the over-representation may overstate the gap's prevalence) instead of only a strengthening.

### 🔴 #441 (p201) — open
**Kritikos:** `You are lacking statistical validity ...`

**Anchored on:** "Threats to Validity"  → **09-discussion.tex:104**

**Meaning:** 'You are lacking statistical validity' — the Threats section has Internal/External/Construct/Agent-driven categories but no statistical (conclusion) validity paragraph. Implied change: add one acknowledging purely descriptive pass/fail counts, no confidence intervals or significance testing, small n (15 check-level defects), so claims are demonstrative rather than statistically generalizable.

### 🔴 #442 (p201) — open
**Kritikos:** `and did not cover other IDPs (i.e., IDPs deployed in other organisations)`

**Anchored on:** "multi-cluster operational dynamics are not exercised."  → **09-discussion.tex:110**

**Meaning:** Extend the single-IDP concession: the operational evaluation also did not cover other IDPs deployed in other organisations, so cross-platform portability of XSDLC is untested. Implied change: append that clause to the external-validity paragraph.

### 🔴 #443 (p201) — open
**Kritikos:** `I am also concerned about the way functional testing is performed and whether it can really supply complete oversight of specification drift. The current implementation covers structural discrepancies but not content discrepancies.`

**Anchored on:** "Construct validity. SDT's principle-based scoring treats each principle"  → **09-discussion.tex:112**

**Meaning:** Companion to 438, framed as construct validity: he doubts functional testing can supply complete oversight of specification DRIFT, because the current implementation detects structural discrepancies (response shape vs schema) but not content discrepancies. Implied change: concede that drift detection is structural-only and a content-aware oracle would be needed for complete drift oversight.

### 🟡 #444 (p201) — needs-author-decision
**Kritikos:** `is this kind of internal validity issue? If yes, then why is it mentioned separately?`

**Anchored on:** "the interpretation of results may be influenced by"  → **09-discussion.tex:114**

**Meaning:** He asks whether the agent-confirmation-bias threat is itself a kind of internal-validity issue, and if so why it gets its own separate category. The author must either fold the 'Agent-driven evaluation' paragraph into Internal validity or keep it separate and justify why agent-mediated interpretation warrants its own category.

### 🔴 #445 (p201) — open
**Kritikos:** `Non-functional testing is also very restrained, covering only particular testing kinds and limited constraints/requirements. Again, this creates a bias towards high detection rate for specific performance constraints but not others. Thus, the real evaluation over different non-functional requirements would have delivered less favourable results.`

**Anchored on:** "Construct validity. SDT's principle-based scoring treats each principle"  → **09-discussion.tex:112**

**Meaning:** Add a further threat: non-functional testing (P007) is very restrained — it covers only particular test kinds and limited constraints (latency p95, success rate under k6 load), not the broader NFR space (throughput, resources, soak/spike, concurrency). This biases detection rates upward for the tested constraints; a broader NFR evaluation would deliver less favourable results. Absent from the live Threats section.

### 🔴 #446 (p202) — open
**Kritikos:** `plus reduce specification drift? As this is also covered, although not fully for now.`

**Anchored on:** "improve specification quality without requiring immediate perfection."  → **09-discussion.tex:121**

**Meaning:** Suggests adding that SDT adoption also reduces specification drift ('plus reduce specification drift? As this is also covered, although not fully for now') — i.e. mention drift reduction as a benefit for API teams, with the caveat that drift coverage is currently partial (structural only).

### 🔴 #447 (p202) — open
**Kritikos:** `and potentially incomplete results`

**Anchored on:** "the capability of the model alone. A capable model without structured context produces inconsistent results"  → **09-discussion.tex:127**

**Meaning:** Append his addition: 'and potentially incomplete results' — i.e. a capable model without structured context produces inconsistent AND potentially incomplete results.

**⚠ Triage correction:** Earlier triage anchored this to the 'For API teams' minimal-mode paragraph (line 121) and read it as 'minimal mode gives incomplete quality coverage'; the icon actually sits in the AI-assisted-development paragraph (line 127), and the comment is a two-word insertion into 'inconsistent results'.

### 🔴 #448 (p203) — open
**Kritikos:** `this requires more detail as a specific evaluation cannot be regarded as a threat to validity.`

**Anchored on:** "Threats to validity include the single-API controlled evaluation"  → **09-discussion.tex:136**

**Meaning:** He objects that 'a specific evaluation cannot be regarded as a threat to validity' — the summary lists evaluation design choices as if they were threats. Implied change: reword each item to name the threat the design introduces (limited generalizability from a single reference API, sampling bias in the public-API arm, lost severity granularity from binary scoring), and sync the list with whatever new threats (statistical validity, content drift, narrow NFR coverage) get added to Section 9.3.


## 10-conclusion.tex — 32 comments (21 open)

### ❓ #449 (p204) — unclear
**Kritikos:** `consumer/customer`

**Anchored on:** "contract-first and test-first approaches by deriving all validation rules from"  → **10-conclusion.tex:11**

**Meaning:** Comment 'consumer/customer' sits on 'contract-first and test-first approaches'. Most plausible reading: Kritikos is pointing at the established term 'consumer-driven (customer-driven) contract testing' and wants it named/acknowledged when the framework distinguishes itself from contract-first approaches. Alternative reading: a consumer-vs-customer terminology fix, but neither word appears in this chapter.

**⚠ Triage correction:** Triage reported ANCHOR-NOT-FOUND and guessed a consumer→customer word swap somewhere; the icon actually sits on 'contract-first and test-first approaches' (line 11), suggesting he means consumer/customer-driven contract testing terminology.

### 🔴 #450 (p204) — open
**Kritikos:** `OpenAPI`

**Anchored on:** "the specification itself, requiring zero manual test authorship"  → **10-conclusion.tex:11**

**Meaning:** One-word comment 'OpenAPI' on 'the specification itself': he wants the word OpenAPI inserted, i.e. 'deriving all validation rules from the OpenAPI specification itself' — make explicit which specification is meant. The live line 11 still says only 'the specification itself'.

**⚠ Triage correction:** Triage read it as an OpenAPI capitalisation/spelling fix anchored to 'OpenAPI~3.x and Swagger~2.0'; the icon actually sits on 'the specification itself' — he wants the word OpenAPI added there.

### 🔴 #451 (p204) — open
**Kritikos:** `in suitable cases`

**Anchored on:** "with mode-aware behavior; and the engine orchestrates evaluation"  → **10-conclusion.tex:13**

**Meaning:** Insert the hedge 'in suitable cases' at this point in the DriveBy contribution paragraph — i.e. make explicit that the principle checks/engine evaluation are applied only in the suitable cases (modes), e.g. '...implement the static validation principles with mode-aware behavior, applying checks in suitable cases; and the engine orchestrates...'. Line 13 is unchanged, so the insertion is still missing.

**⚠ Triage correction:** Triage anchored this to 'deriving all validation rules...zero manual test authorship' (line 11) as a claim-softening hedge; the icon actually sits on the mode-aware-behavior/engine sentence (line 13) — the hedge belongs to the mode-aware evaluation claim.

### 🔴 #452 (p205) — open
**Kritikos:** `You forgot to mention single-API experiment & the new evaluation concerning the use of an reconciliation agent`

**Anchored on:** "large-scale study over public OpenAPI specifications harvested from"  → **10-conclusion.tex:19**

**Meaning:** The 'Empirical evaluation' contribution bullet describes only a three-arm evaluation; Kritikos says it forgets the single-API (single-spec) agent-feedback experiment and the new reconciliation-agent evaluation. Add these arms explicitly, cross-referencing Section sec:sdt-feedback-experiment. Live line 19 still lists only three arms.

### 🔴 #453 (p205) — open
**Kritikos:** `where is this proof stated? Please indicate section number.`

**Anchored on:** "consumable by AI agents for automated evaluation analysis, validating the"  → **10-conclusion.tex:21**

**Meaning:** He asks where the proof that 'SDT's structured output proved to be directly consumable by AI agents' is stated, and wants a section number, not just '(Chapter 8)'. Live line 21 still cites only Chapter~\ref{ch:ai-development}; the proof actually lives in Section~\ref{sec:sdt-feedback-experiment} (07-evaluation.tex:682) — add that section reference.

**⚠ Triage correction:** Triage anchored this to the '45–50 objects in under 15 seconds' claim in the RQ3 answer (line 38); the icon actually sits on the AI-agent consumability claim in the agent-driven framework contribution (line 21).

### 🔴 #454 (p206) — open
**Kritikos:** `no, it was less than that ...`

**Anchored on:** "demonstrated 100% detection accuracy"  → **10-conclusion.tex:30**

**Meaning:** 'No, it was less than that' — he disputes the 100% detection claim in the RQ1 answer. The regenerated evaluation now reports the per-check extension at 12/15 (80%) with two missed checks (07-evaluation.tex:283), though the seven principle-level injected defects are all detected. Line 30 still claims unqualified '100% detection accuracy with zero false positives' and must be corrected/qualified to match the per-check figures.

### 🔴 #455 (p206) — open
**Kritikos:** `Again, non-functional testing does not rely on OpenAPI specification along but also the respective XSDLC CR.Maybe also refer to the evaluations that signify that the right gates are produced, which enable transition only when their conditions are satisfied.`

**Anchored on:** "no test data is manually curated. The test-ready"  → **10-conclusion.tex:34**

**Meaning:** Recurring point: non-functional testing does not derive from the OpenAPI specification alone — performance thresholds come from the XSDLC CR. The RQ2 answer should qualify that 'entirely from the OpenAPI specification' holds for static principles only, and should also cite the evaluations showing the right gates are produced and block promotion until conditions are satisfied. Line 34 is unqualified (the single-artifact concession was made in 02-related-work.tex:275 but not threaded into this RQ2 answer).

### 🔴 #456 (p206) — open
**Kritikos:** `Could also refer to other CI/CD integration metrics. In addition, you could mention the scalability of the system as it could handle 5 API CRs without any problem as well as appropriately addressed specific injected issues ...`

**Anchored on:** "generates approximately 45–50 managed Kubernetes objects in under 15 seconds"  → **10-conclusion.tex:38**

**Meaning:** Strengthen the RQ3 answer with more evidence: cite other CI/CD integration metrics from Chapter 7, mention scalability (the system handled 5 API CRs without problems) and that the gates correctly addressed the specific injected defects. Line 38 still lacks these additions.

### ✅ #457 (p207) — already-addressed
**Kritikos:** `potential ...As they have not been implemented yet.`

**Anchored on:** "inspected for its declared pipeline configuration, and manageable by"  → **10-conclusion.tex:42**

**Meaning:** Mark the onboarding/observability controllers in the RQ4 answer as potential, since they have not been implemented. The live line 42 now says 'a team onboarding controller could create... an observability controller could adjust... these controllers are not yet implemented (see Section~\ref{sec:future-work})' — exactly the requested hedge.

### 🔴 #458 (p207) — open
**Kritikos:** `how this correlates with security? Because we have authentication covered in this way?Maybe security could be better covered in the near future to cater for more advanced security testing scenarios?`

**Anchored on:** "a security policy controller could enforce that all XSDLC"  → **10-conclusion.tex:57**

**Meaning:** He asks how the 'security policy controller' bullet correlates with security — is security covered only via authentication-scheme presence (P005)? He suggests stating that, and adding near-future work on more advanced security-testing scenarios. The bullet (line 57) is unchanged; security testing is only mentioned as a doctoral coverage-extension item (line 115), so the clarification and dedicated future-work mention are still missing.

### 🔴 #459 (p208) — open
**Kritikos:** `Still this covers the documentation of versioning.It does not evaluate how well versioning is supported. The latter requires understanding the semantics of the versioning strategy applied, if it exists.`

**Anchored on:** "rely on keyword matching in the info.description field"  → **10-conclusion.tex:75**

**Meaning:** Deeper limitation than brittleness: P008 only checks that versioning is documented, not how well versioning is actually supported — the latter requires understanding the semantics of the applied versioning strategy. Add this semantic-adequacy limitation to the 'Keyword-based checks' paragraph. Line 75 is unchanged.

### 🟡 #460 (p208) — needs-author-decision
**Kritikos:** `correct - although I foresee a specific issue: if all principles are applicable in strict or other modes, then this means that these two implemented principles will be always applied. This contracts the layered validation approach that you have designated. Unless my expectation is not correct - dynamic validation is another mode that should not be perplexed with static validation ...`

**Anchored on:** "Full principle coverage would strengthen the SDT framework's claim of comprehensive"  → **10-conclusion.tex:73**

**Meaning:** He agrees with the limitation but warns: if P006/P007 become PrincipleCheckers that run in strict (or all) modes, dynamic validation gets conflated with static validation, contradicting the layered-validation design. The text should clarify that dynamic validation is a separate mode, not mixed into static modes. Line 73 (and the paired future-work item at line 96) still do not address the static/dynamic mode separation; the resolution is a framework design choice.

### 🔴 #461 (p208) — open
**Kritikos:** `Maybe say in an introductory paragraph how these limitations were discovered (e.g., thesis evaluation feedback, the conducted evaluations, etc.)`

**Anchored on:** "10.4 Limitations (section heading)"  → **10-conclusion.tex:67**

**Meaning:** Add an introductory paragraph to the Limitations section explaining how the limitations were discovered (evaluation results, thesis review feedback, implementation scope decisions). The section still opens directly with the first bullet — no provenance paragraph exists.

### 🔴 #462 (p209) — open
**Kritikos:** `Further, you only cover load testing and not other kinds of non-functional testing.In addition, functional testing covers output structural compliance. So, no actual output content compliance.`

**Anchored on:** "not an adaptive contract."  → **10-conclusion.tex:79**

**Meaning:** Two additions to the limitations: (1) only load testing is covered, no other kinds of non-functional testing; (2) functional testing (P006) checks output structural compliance only, not actual output content compliance. Neither point appears in the live Load-testing-scope (line 79) or Principle-coverage (line 73) paragraphs.

### ✅ #463 (p209) — already-addressed
**Kritikos:** `Further, ... + multi-cluster evaluations ...`

**Anchored on:** "The operational evaluation is limited to a single IDP"  → **10-conclusion.tex:81**

**Meaning:** Add that multi-cluster evaluation is also missing ('Further, ... + multi-cluster evaluations'). The live file now has a dedicated 'Single delivery cluster' limitation (line 83) stating multi-cluster promotion is not exercised — this resolves the comment.

### ✅ #464 (p209) — already-addressed
**Kritikos:** `Could also add lack of reconciliation to improve OpenAPI specification quality`

**Anchored on:** "(end of Limitations, before 10.5 Future Work)"  → **10-conclusion.tex:71**

**Meaning:** Add as a limitation the lack of a reconciliation loop that improves OpenAPI specification quality automatically. The new 'Verdict tier is open-loop' limitation (line 71) states exactly this: no automated spec re-write or autonomous remediation, the corrective action is external, and closing the loop is recorded as future work.

### 🔴 #465 (p209) — open
**Kritikos:** `Maybe indicate in a small paragraph where these directions come from the limitations that were stated in Section 10.4.`

**Anchored on:** "10.5 Future Work (section heading)"  → **10-conclusion.tex:86**

**Meaning:** Add a small introductory paragraph to Future Work indicating that the directions derive from the limitations stated in Section 10.4 (mapping each direction to its motivating limitation). The section still opens directly with the first item — no such paragraph exists.

### ✅ #466 (p209) — already-addressed
**Kritikos:** `We envision a ...`

**Anchored on:** "Adaptive load testing thresholds. A Kubernetes controller that watches..."  → **10-conclusion.tex:90**

**Meaning:** Reword the opener to 'We envision a ...' so the future-work item reads as a vision rather than a bare statement. Live line 90 now begins 'We envision a Kubernetes controller that watches production observability metrics' — done.

### ✅ #467 (p209) — already-addressed
**Kritikos:** `maybe add "constantly" or "very often"? So that the improvement is made only when the pattern repeats ...`

**Anchored on:** "adjusts loadTestConfig fields based on actual traffic patterns. When observed"  → **10-conclusion.tex:90**

**Meaning:** Add 'constantly'/'very often' so the controller only tightens the gate when the traffic pattern repeats, not on a one-off spike. Live line 90 now reads 'When observed concurrent users consistently exceed the configured threshold' — done.

### 🔴 #468 (p209) — open
**Kritikos:** `partially ... -> in the sense that it improves XSLDC CR but not the OpenAPI specification. While the latter might also have to be improved.`

**Anchored on:** "This would close the feedback loop between observability and quality gating"  → **10-conclusion.tex:90**

**Meaning:** Qualify with 'partially': the adaptive-threshold controller only improves the XSDLC CR (gate thresholds), not the OpenAPI specification, which might also need improvement. Live line 90 still claims it 'would close the feedback loop' without the qualifier.

### 🔴 #469 (p209) — open
**Kritikos:** `less critical than ...?I acknowledge that this can be distinguishable but the risk should be also different.`

**Anchored on:** "a partial documentation gap is distinguishable from a complete one"  → **10-conclusion.tex:92**

**Meaning:** 'Less critical than ...?' — distinguishability is not enough; the assigned risk/weight should also differ. Reword so a partial documentation gap carries explicitly lower operational risk (weight) than a complete one. Line 92 unchanged.

### 🟡 #470 (p209) — needs-author-decision
**Kritikos:** `Should we also fix code or only the OpenAPI specification (and maybe XSDLC CR)? I am asking this as it is much better to reduce the scope over the specifications by considering that it could be argued that you need to have a stable ground truth. If it is not certain what is the correct, the code or the specifcation, this could be very problematic, if not properly constrained somehow.`

**Anchored on:** "generate specification fixes or code patches, and re-push to trigger re-validation"  → **10-conclusion.tex:94**

**Meaning:** Should agents fix code too, or only the OpenAPI spec (and maybe the XSDLC CR)? He recommends restricting remediation scope to the specifications so there is a stable ground truth; letting the agent edit either code or spec is problematic when it is uncertain which is correct, unless properly constrained. Line 94 still says 'specification fixes or code patches'; the remediation-scope policy is the author's design decision.

### 🔴 #471 (p209) — open
**Kritikos:** `This was also proven in the evaluation ...`

**Anchored on:** "is precisely the interface that AI agents need to operate autonomously"  → **10-conclusion.tex:94**

**Meaning:** 'This was also proven in the evaluation' — ground the agent-driven-remediation future work in the agent-feedback experiment (Section~\ref{sec:sdt-feedback-experiment}) instead of presenting it as purely speculative. The experiment exists and is cited in the open-loop limitation (line 71) and the doctoral item (line 125), but the agent-driven-remediation paragraph at line 94 still does not reference it.

### 🔴 #472 (p210) — open
**Kritikos:** `Ok but maybe this is accomplished in a multi-round setting in a complete manner? As we need to perfection the current situation, not just slightly improve it (e.g., reduce principle failures but not mitigate them all).`

**Anchored on:** "can generate the appropriate maxLength and pattern constraints without human guidance"  → **10-conclusion.tex:94**

**Meaning:** The one-shot single-failure example is too modest: remediation should be a multi-round process that converges to complete mitigation of all principle failures, not a slight improvement. Line 94 still gives only the one-shot example; the multi-iteration study at line 125 partially covers this for doctoral work but the future-work item itself should state the iterative-convergence goal (or cross-reference line 125).

### 🟡 #473 (p210) — needs-author-decision
**Kritikos:** `This what was actually indicated before in a previous comment.Do we desire this if the goal is to have layered validation?As had the impression that strict validate-only covers full static validation and not dynamic one.So, with this sentence, you actually indicate that it covers both!!!!So, there will be no way to distinguish between them?Or maybe we could have a full validation plus aspect-specific validations (static & dynamic ones)?Would that require further improvements/adjustments in the framework and DriveBy tool?`

**Anchored on:** "validate-only --validation-mode strict invocation to produce a complete quality assessment"  → **10-conclusion.tex:96**

**Meaning:** Strong objection (pairs with idx 460): strict mode was understood as full STATIC validation; saying a single strict invocation covers all nine principles means strict also runs dynamic tests, destroying the static/dynamic distinction. He proposes alternatives (full validation plus aspect-specific static and dynamic validations) and asks whether the framework/DriveBy would need adjustments. Line 96 still contains the contested sentence; resolving the mode design is the author's call.

### 🔴 #474 (p210) — open
**Kritikos:** `I can agree with you but there is a major issue here: completeness. GraphQL introspection schemas and gRPC protobuf service definitions cover partially the information aspects covered by OpenAPI. Thus, there will be a need to extend these specifications to cover all the necessary information aspects (which are obviously required for testing/quality assurance purposes).`

**Anchored on:** "adapters for GraphQL introspection schemas, gRPC protobuf service definitions"  → **10-conclusion.tex:100**

**Meaning:** Completeness caveat: GraphQL introspection and gRPC protobuf only partially cover the information aspects OpenAPI covers, so multi-format support requires extending/supplementing those specifications to carry the information the principles need for testing/QA. Line 100 still presents the adapters as a simple interface extension without this caveat.

### 🔴 #475 (p210) — open
**Kritikos:** `Could provide a reference here that supports/backs up this claim.`

**Anchored on:** "cross-cluster ArgoCD ApplicationSets and multi-cluster Crossplane providers"  → **10-conclusion.tex:102**

**Meaning:** Provide a reference backing the claim that cross-cluster ArgoCD ApplicationSets and multi-cluster Crossplane providers support multi-cluster delivery. Line 102 still has no citation.

**⚠ Triage correction:** Triage anchored this to 'without human guidance' in the agent-remediation item and proposed an LLM/spec-generation citation; the icon actually sits on 'multi-cluster Crossplane providers' in the Enterprise multi-cluster patterns item — the missing reference is for the ApplicationSets/multi-cluster-providers claim.

### ✅ #476 (p210) — already-addressed
**Kritikos:** `Again, non-existing reference number.`

**Anchored on:** "software-testing methodology in the sense of Kritikos's review [160]"  → **10-conclusion.tex:113**

**Meaning:** '[160]' was a stale, non-existent reference number (a leftover comment-ID, not a bibliography entry) and had to be removed. The live line 113 has been rewritten to '...to a complete software-testing methodology. This requires three additions...' and grep confirms no bare bracketed numbers ([160], [501], [484], [503], [213]) remain anywhere in the chapter.

### 🔴 #477 (p211) — open
**Kritikos:** `and security`

**Anchored on:** "first-class concepts for performance characteristics (expected latency,"  → **10-conclusion.tex:119**

**Meaning:** 'and security' — the ontological/metamodel extension should cover not only performance characteristics but also security (first-class security-testing concepts, beyond the basic security schemes OpenAPI already declares; consistent with his idx 458 push for advanced security testing). Line 119 still lists only performance characteristics.

**⚠ Triage correction:** Triage anchored this to 'completeness and minimality theorems' in the formal-methodology item (line 113); the icon actually sits on 'first-class concepts for performance characteristics' in the Ontological-extension item (line 119) — security should be added to that extension.

### 🔴 #478 (p211) — open
**Kritikos:** `more importantly, the contract would then cover both the ideal functional and non-functional behaviour of an API!So, XSDLC would not have to cover the non-functional behaviour and would be restricted to specifying mainly gate policies.`

**Anchored on:** "would close the loop between contract specification and performance gating"  → **10-conclusion.tex:119**

**Meaning:** Add the key consequence: with a performance-annotation vocabulary, the OpenAPI contract would cover BOTH ideal functional and non-functional behaviour of the API, so the XSDLC CR would no longer carry non-functional behaviour and would be restricted to specifying mainly gate policies. Line 119 still ends without this consequence.

### 🔴 #479 (p211) — open
**Kritikos:** `isn't that already supported?`

**Anchored on:** "an evaluator that integrates with cluster observability, and an audit trail that records why each promotion"  → **10-conclusion.tex:121**

**Meaning:** 'Isn't that already supported?' — he questions whether conditional gating / promotion audit trail is already provided by the current system. The text should clarify the delta: current gates evaluate point-in-time load-test results declared in the XSDLC CR, whereas the proposed conditional gates evaluate constraints over live, time-windowed production observability — which is not supported today. Line 121 does not yet make this distinction.

### ✅ #480 (p211) — already-addressed
**Kritikos:** `single constraint or logical combinations of constraints (i.e., composite constraints)`

**Anchored on:** "conditional gate that lets a build through subject to constraints (e.g. "promote"  → **10-conclusion.tex:121**

**Meaning:** Gate conditions should support not only single constraints but logical combinations (composite constraints). Live line 121 now reads 'a formal language for gate conditions supporting single constraints and their logical combinations (composite constraints)' — done.
