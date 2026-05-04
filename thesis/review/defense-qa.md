# Defense Q&A — Rehearsed Answers for the Risky Items

Last updated: 2026-05-04 (round-7 closure).

This file holds rehearsed answers for the comments where the supervisor's
position diverges from the thesis or where the thesis pushes back. Use it
as a memory aid before defense; do not read it verbatim.

Each entry: comment ID, supervisor concern, our position, the answer, and
the strongest counter-argument we should be ready for.

---

## 1. [65] / [556] — "Why is Chapter 9 not merged into Chapter 7?"

**Supervisor:** Ch.9 (Discussion) discusses results that Ch.7 (Evaluation)
produced. Why are these separate? Could include threats to validity too.

**Our position:** Kept them separate.

**Answer.** Ch.7 reports evidence; Ch.9 interprets it against the research
questions and surfaces threats to validity. Merging would conflate "what
we measured" with "what it means" and would force the reader to absorb
new evidence and the interpretive frame in the same pass. Ch.7 is now
~30 pages; Ch.9 is ~6 pages of focused interpretation. Splitting them
keeps Ch.7 readable as a record of what was tested and Ch.9 readable as a
distilled response to RQ1–RQ4. Ch.8 (AI-assisted development) sits
between the two because it is itself an evaluation arm of a different
kind, and its placement before Ch.9 is what enables Ch.9 to draw on
both Ch.7's empirical results and Ch.8's process evidence.

**Strongest counter we may face.** "Six pages is short enough to merge."
Concession answer: "I see the argument; if you would prefer them merged
for the bound copy I can fold the §9.5 threats-to-validity block into
Ch.7 as a closing section. The interpretive narrative would still need
its own home, which I would keep as a short Ch.9. The current structure
follows the convention used in Hevner & March's design-science papers
that I cite in §1.6."

---

## 2. [557] — "'methodology' is too strong a word"

**Supervisor:** SDT is a *framework*, not yet a *methodology*. Calling it a
methodology overclaims.

**Our position:** Agreed in round 1; renamed DDT→SDT and
re-positioned SDT as a *conceptual framework, a strong theoretical core
for a future formal methodology* (Ch.1 §1.4 first contribution; Ch.10
§10.5 future-work item 1). Round-7 softened the remaining methodology→
framework wording in Ch.1, Ch.7, Ch.8.

**Answer.** "I agree — that is exactly the framing I adopted. The thesis
explicitly distinguishes three things: the *axioms* (the theoretical
foundation), the *framework* (the axioms plus the principle catalogue
plus the validation modes), and a future *methodology* (the framework
plus a formal test-derivation procedure, execution and failure
semantics, and lifecycle governance). What this thesis delivers is the
framework. Section 10.5 lays out the formal-methodology programme as
the principal future-work direction."

**Counter we may face.** "Then why does Chapter 3 still call itself
'Methodology: SDT'?" — Answer: "Section 3.1 explains that the chapter
title is methodological in the research-process sense (it presents the
framework's structure and rationale), not in the formal
software-testing-methodology sense. I can rename it to 'The SDT
Framework' if that is clearer."

---

## 3. [447] — "Control-theory mapping mixes formal properties with informal language"

**Supervisor:** Stability, controllability, and observability are
specific control-theory properties. Mapping setpoint/process-variable to
the SDT loop is loose.

**Our position:** Agreed in round 4. §6.5 reframed the control-theory
discussion as an *illustrative analogy*, not a formal claim. A footnote
acknowledges the analogy is not a stability-theoretic argument.

**Answer.** "You are right that I am not making a formal control-theory
claim. The analogy is purely illustrative — I use the
setpoint/process-variable/feedback-loop vocabulary because the GitOps
community itself uses this language to motivate Argo CD's reconciliation
loop, and the SDT loop fits that informal pattern. I do not claim
Lyapunov stability or formal observability properties. Section 6.5 now
says this in a footnote. If the analogy is more confusing than helpful,
I can drop it and replace it with the simpler 'reconciliation loop'
framing used elsewhere in the thesis."

---

## 4. [297] — "Composite/stateful API testing is missing"

**Supervisor:** P006 (functional testing) covers single interactions.
Real APIs need multi-step sequences and prerequisite state. Calling this
'preliminary' would be honest.

**Our position:** Acknowledged. Future work in Ch.10 §10.5 explicitly
names stateful and multi-step test generation as the natural P006
extension (this is also where RESTler's stateful-fuzzing and EvoMaster's
sequence-evolution work, both surveyed in §2.1.2, would integrate).

**Answer.** "Correct. P006 in this thesis runs single-shot, schema-derived
requests and verifies the response against the declared schema. It does
not handle producer-consumer dependencies between endpoints, stateful
preconditions like database fixtures, or multi-step session flows. This
is a deliberate scope: P006 was the first runtime principle to land and
the thesis focuses on getting the static principles, the GitOps
integration, and the methodology framing right. Stateful and
sequence-based testing is the most consequential P006 extension and it
is named as such in Section 10.5. The architecture supports it — the
APISpec abstraction exposes the producer-consumer dependency
information through the operation tree — so the work is bounded
implementation, not new theory."

**Counter.** "Then why claim P006 detects functional defects?" —
Answer: "It detects the subset of functional defects that surface in a
single request/response pair: schema mismatches, undocumented response
codes, type errors. That is a real subset; the multi-API table in
Section 7.4 shows P006 catching the broken-api defect that the static
principles miss. The thesis does not claim P006 is a complete
functional-testing solution."

---

## 5. [580] — "Risk-based scoring would be more valuable than severity-weighted"

**Supervisor:** Severity is binary; a risk model captures more nuance.

**Our position:** Severity-weighted is the implemented model;
risk-based is named as future work in Ch.10 §10.5.

**Answer.** "The severity model in the thesis is the simplest decision
that maps cleanly to a quality gate: critical → block promotion,
warning → record but pass. A risk-based score — combining severity,
likelihood of exploitation, blast radius, and fix cost — would be more
informative but it would also require empirical calibration data the
thesis does not collect. Section 10.5 names the risk-based extension as
a future-work direction with a specific hook: the per-principle JSON
output already carries the structured information a risk-scoring
function would need (severity, fail count, suggested-fix metadata),
so adding the risk score is a downstream report-generator change, not
a framework change."

---

## 6. [246] — "Chapter 4 should be titled 'The DriveBy System' not 'CLI Architecture'"

**Supervisor:** "CLI architecture" sounds like a chapter about CLIs in
general.

**Our position:** Kept "DriveBy CLI Architecture" because Ch.5 covers
the wider DriveBy system (Kubernetes operator, XSDLC); Ch.4 is
specifically the CLI binary architecture.

**Answer.** "Chapter 4 is specifically about the binary — the dependency
flow across eleven internal Go packages, the APISpec abstraction layer,
the principle-checker pattern, the engine. Chapter 5 then takes the
CLI as a building block and shows how it composes into the wider
DriveBy system: the XSDLC custom resource, the Crossplane composition,
the GitOps quality gate. The two chapters split this way because the
CLI has a coherent internal architecture worth chapter-length treatment,
and the system has a different coherent architecture also worth
chapter-length treatment. If you would prefer 'The DriveBy CLI' as the
chapter title rather than 'DriveBy CLI Architecture', that is a clean
edit."

---

## 7. [140] — "'Static Analysis' should be 'Syntactic Validation'"

**Supervisor:** Static analysis covers more than syntactic validation.

**Our position:** Kept "Static Analysis" — it is the field-standard term
for what Spectral, Vacuum, and Redocly do.

**Answer.** "Static analysis is the standard term in the API-tooling
literature — Spectral, Vacuum, and Redocly all describe themselves as
static analysers. Calling that category 'syntactic validation' would
narrow the term: Spectral's rules cover documentation completeness,
naming conventions, and security-scheme presence — not just syntax.
The §2.6 comparison table contrasts static-vs-runtime analysis, which
is the meaningful axis. I take the underlying point that 'static
analysis' is broader than the current tools' actual coverage; I will
say so explicitly in the §2.6 caption if you would prefer."

---

## 8. [115] — "'Reconciliation' or 'adaptation'?"

**Supervisor:** "Adaptation" better captures convergence to a desired state.

**Our position:** Kept "reconciliation" — the GitOps community
(CNCF OpenGitOps, Argo CD docs, Crossplane docs) uses this term.

**Answer.** "'Reconciliation' is the term-of-art in the GitOps and
Kubernetes community — Argo CD reconciles, Crossplane controllers
reconcile, the OpenGitOps principles document calls the loop a
reconciliation loop. Using that term is what makes the §2.4.4
'ontological state reconciliation' framing legible to a reader from
that community. 'Adaptation' is the better metaphor in cybernetics
and control theory; 'reconciliation' is the better term in
practitioner-facing GitOps literature. I chose the practitioner term
because the audience for the thesis includes engineers who will
recognise it instantly."

---

## 9. P006 / P007 PrincipleChecker wrappers are "planned"

**Supervisor risk:** May ask why the thesis ships with 6 of 9 principles
as full PrincipleCheckers and the other 3 (P006, P007, plus P009 in some
modes) as separate test runners.

**Our position:** Architectural decision — the runtime-test runners
predated the PrincipleChecker abstraction and the wrapper integration
is queued as mechanical work, not new design.

**Answer.** "The static principles (P001–P005, P008) are pure functions
over the parsed specification — given a spec, return a result. The
runtime principles (P006 functional, P007 performance) are I/O-bound:
they need an HTTP client, a configurable target host, and timeout
handling. The current architecture exposes them as separate test
runners that produce the same per-principle JSON output the
PrincipleChecker registry produces, so downstream consumers (the report
generator, the GitOps gate, the agent feedback loop) cannot tell the
difference. The wrapper that registers P006/P007 in the PrincipleChecker
registry is queued as a mechanical refactor — it would replace three
ad-hoc CLI subcommands (`function-only`, `load-only`, `test-only`)
with a registry lookup. Not in the thesis because it is not part of the
methodological contribution."

**Counter.** "Then are the runtime principles really part of SDT?" —
Answer: "Yes — the framework defines nine principles grounded in three
axioms. P006 is grounded in Determinism, P007 in Observability. The
implementation status is recorded in Table 3.x with explicit Implemented
vs Planned labels. The framework is complete; the implementation is
seven full PrincipleCheckers plus two test runners that produce the
same output shape."

---

## 10. The operational PoC is single-cluster, single-team

**Supervisor risk:** May ask how many real developers used the system.

**Our position:** Honest — the PoC is one cluster, one platform team
operator (the author), five evaluation APIs.

**Answer.** "The PoC is small by industrial standards: one cluster,
one platform-team operator, five APIs, hundreds of gate executions
captured over the evaluation window. It is not a multi-tenant
deployment with many development teams, and the thesis does not claim
otherwise — Section 7.2.1 describes exactly the IDP baseline that
existed before the PoC and Section 7.2.2 describes exactly what the
PoC added. The PoC's role is to demonstrate that XSDLC actually
runs end-to-end on a real cluster — the alternative would be a
synthetic minikube demo, which would not show the platform-inheritance
property that Sections 5.2 and 7.2 emphasise. Multi-cluster and
multi-team scaling is named as future work in Section 10.5."

---

## Spot-check defense moves

If the supervisor asks "where did you address comment X" and X is in
the `addressed-by-rewrite` pile, the answer pattern is:

1. Open `thesis/review/comment-index.md`, find the row, read the
   `[r4: ...]` trailer aloud — it names the section that addresses it.
2. Open the named section in the PDF, point to the relevant paragraph.

If the section does not in fact address the comment, the honest answer
is: "I marked this as absorbed by the round-1 rewrite of §X.Y. Let me
reread the comment and the current section together — if the resolution
is wrong I would rather correct it now than defend a stale claim."

The audit trail in `[r4: ...]` lets this answer be given without
shuffling through pages.
