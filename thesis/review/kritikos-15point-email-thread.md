# Kritikos 15-Point Summary + Author's Email Replies

This is the executive-summary email Kritikos sent (2026-06-09) alongside the 480 inline
PDF comments, plus the author's (Petros) point-by-point reply. The inline comments are
the detailed work list; these 15 are the "must be final" headline issues.

The author's replies are paraphrased/positioned below. The workflow should RE-EVALUATE
each point independently on technical merit and flag where the author's email position
may be weak before he locks it in.

---

## Kritikos's framing
"Let us consider these the last issues so the change/feedback cycle does not run forever.
Summary of the most important issues below, but look carefully at the inline comments too."

---

### (1) Reconciliation agent evaluation forgotten in several places
**Kritikos:** The text forgets, in various relevant places, the new evaluation that was
performed concerning the use of a reconciliation agent.
**Author's reply:** "I need further clarification — I think this is general small comments
that I'll look at. You clearly have stronger things in mind, so I'll check the inline comments."

### (2) Enrich / change the mapping of contributions to research questions
**Kritikos:** There is a need to enrich/change the correspondence between contributions and
research questions.
**Author's reply:** "If needed, it will be done."

### (3) Table 2.1 last criterion + single-artifact claim
**Kritikos:** A more correct definition of the last criterion in Table 2.1 is needed, and its
evaluation in that table may change. Note: your approach does NOT consider only one artifact
but TWO — the OpenAPI specification AND the XSDLC CR — to produce tests and gates. The problem
is that the non-functional requirements are expressed in the CR and not in the OpenAPI spec.
If those requirements could be moved into the OpenAPI spec, then your approach would indeed be
single-artifact. So the question is whether the criterion should change to be clearer and correct.
**Author's reply (PUSHBACK):** "I disagree. We rely on a single artifact for the PROCESS — the
live OpenAPI spec published at runtime. What you describe sounds analogous to: if we dropped XSDLC
entirely and used only the CLI, THEN it would be single-artifact? In my mind, deployment should
not be conflated with the process, which IS based on a single artifact. I'll review my wording in
the text and come back on where confusion may exist, if any."

### (4) No correlation of research gaps with DriveBy tool in Section 2.6.3
**Kritikos:** The research gaps have not been correlated with the DriveBy tool in Section 2.6.3.
**Author's reply:** "Will do an iteration on this."

### (5) Non-existent reference numbers + repetitive in-text comment-handling
**Kritikos:** Non-existent bibliographic reference numbers are continually provided in the text
(these likely refer to comments I made on the previous version, but since the corresponding
references don't exist in the text, they should be removed). Also, there is repetition in how
comments are addressed within the text. Good solution: refer to my comments and how they were
addressed ONCE only. A better alternative: make the text self-contained, not referring to previous
versions at all — this removes any reference to old comments and reduces text volume. Otherwise the
text risks becoming too large.
**Author's reply:** "I'll do a pass, thanks for the pinpointing."

### (6) If precision is part of completeness, make it explicit in the definition
**Kritikos:** Since you state that precision (accuracy) falls within completeness, this should be
made evident in the definition of the latter and mapped to relevant principles.
**Author's reply:** "I'll do an iteration on this having read the document once and keeping it in
my head when I reach that point."

### (7) Do all axioms govern all principles? Governing vs effective distinction
**Kritikos:** Maybe in the end all axioms concern all principles and not only some? Or it can be
considered that each principle corresponds to at most 1-2 axioms. I agree a separation could be
made between whether an axiom is GOVERNING or EFFECTIVE, shown in Table 3.1. I propose you apply
this, but there should be relevant justification in the text.
**Author's reply:** "In a previous mail I mentioned the document changed about this — yes, all
axioms govern every principle and therefore every check. (This was why I stopped, to avoid drift
between documents.)"

### (8) Principle analysis doesn't always fully specify mode-specific behavior
**Kritikos:** During the analysis of the principles, their specific behavior with respect to mode
is not always fully specified (possibly only some modes are covered but not always all relevant ones).
**Author's reply:** "Hmm okay, I'll see what can be done by reading the document once and having it
in my head when I get to that point."

### (9) test-ready principle seems unnecessary (overlaps mode + P002-P004)
**Kritikos:** The test-ready PRINCIPLE seems unnecessary: (a) it overlaps with the test-ready MODE,
and (b) it overlaps with other related principles (P002-P004), giving the impression that the latter
could simply change their implementation to apply only specific checkers in test-ready mode.
**Author's reply (DEFENSE):** "As the CLI was designed, it was convenient for it to be a principle
because it could be used as a building block, and test-ready can for example be used to test whether
something is test-worthy, i.e. whether functional will 'blow up.' test-ready does not contain security
for example. I can change this design in the CLI, but it would need new classes, and at the point the
thesis is at, I don't think it would benefit functionality or any defense. It's a design assumption that
serves the purpose that a mode is a subset of principles which work as building blocks. Please stress the
importance if you consider it that much of a foul."

### (10) Move overall system architecture to its own chapter, before the DriveBy tool chapter
**Kritikos:** I would propose the architecture of the whole system appear in a separate chapter,
before the DriveBy tool one. Then, correctly, the (internal) architecture of the tool can appear
in the chapter dedicated to it.
**Author's reply:** "I'll add one more agent." (i.e. will action this)

### (11) Many inconsistencies between text and figures in Chapter 4; figures need to be more professional
**Kritikos:** There are many inconsistencies between the text content and the figures in Chapter 4.
Also, in general, many figures need changes to become more correct and professional.
**Author's reply:** "I'll do a pass, thanks for the pinpointing."

### (12) Rename "closed-loop negative feedback" to "...without effective actuation"
**Kritikos:** Unless relevant changes are made to the interpretation and representation of the loop,
I would propose renaming "closed-loop negative feedback" to "closed-loop negative feedback without
effective actuation," with the logic that no actions are taken to improve the OpenAPI specs but the
promotion is simply blocked (this is a preventive action that is not explicitly related to improving
the specs).
**Author's reply:** "Since you insist, I'll see how it can be scheduled/routed." (will action)

### (13) Table content + corresponding text wrong in many places; multi-API fault-injection mismatch
**Kritikos:** In many places the content of the provided tables is wrong, and the corresponding text
has errors. Indicatively: while you state that in the multi-API controlled experiment all problematic
APIs are produced from the non-critical-api with a SINGLE change (one fault injection), the corresponding
tables show that many problematic APIs present NOT just one but MANY principle violations — which is wrong.
**Author's reply:** "I'll do a pass." (acknowledged)

### (14) P007 (Performance) must be critical
**Kritikos:** I insist that principle P007 should be critical. Otherwise it makes no sense to block a
promotion when the non-functional requirements that were set are not satisfied.
**Author's reply (AGREES):** "Of course it will go in, yes I agree!"

### (15) Authorship time analysis (Table 8.2) should separate the two deliverables (code vs report)
**Kritikos:** I would propose the time analysis in the authorship section (Table 8.2) separate the two
main deliverables — the code and the report. So we want relevant estimations for the report too. This
must be done because significant contribution from you is expected in both. While the contribution to
the code is demonstrated, this is not done for the report.
**Author's reply (AGREES):** "It will definitely go in, yes I agree!"

---

## Author's closing
"Await my message with the new document and my decisions until then. Thank you again for the diligent feedback."
