export const meta = {
  name: 'kritikos-triage',
  description: 'Triage all 480 Kritikos PDF comments into a prioritized action plan + second-opinion on his 15 headline points',
  phases: [
    { title: 'Triage', detail: 'one agent per chapter .tex: classify every comment + propose concrete action' },
    { title: 'SecondOpinion', detail: 'independently re-evaluate each of Kritikos 15 summary points vs live source + author email reply' },
    { title: 'Synthesize', detail: 'merge into thesis/review/kritikos-action-plan.md' },
  ],
}

const REVIEW = '/home/meter-peter/development/driveby/thesis/review'
const CHAPTERS = '/home/meter-peter/development/driveby/thesis/chapters'

// The 11 target files and how many comments each carries (from the tagged JSON).
const FILES = [
  { tex: '07-evaluation.tex',               n: 97, chapter: 'Evaluation' },
  { tex: '03-methodology-ddt.tex',          n: 61, chapter: 'A Conceptual Framework: Specification-Driven Testing' },
  { tex: '02-related-work.tex',             n: 53, chapter: 'Related Work' },
  { tex: '04-cli-architecture.tex',         n: 49, chapter: 'The DriveBy System' },
  { tex: '05-kubernetes-architecture.tex',  n: 45, chapter: 'Kubernetes Architecture' },
  { tex: '09-discussion.tex',               n: 35, chapter: 'Discussion' },
  { tex: '01-introduction.tex',             n: 33, chapter: 'Introduction' },
  { tex: '08-ai-assisted-development.tex',  n: 33, chapter: 'AI-Assisted Development' },
  { tex: '10-conclusion.tex',               n: 32, chapter: 'Conclusion' },
  { tex: '06-gitops-pipeline.tex',          n: 27, chapter: 'GitOps Pipeline and CI/CD' },
  { tex: 'front-matter',                    n: 15, chapter: 'front-matter (main.tex / abstract / acronyms)' },
]

const TRIAGE_SCHEMA = {
  type: 'object',
  required: ['tex', 'totalComments', 'items', 'fileSummary'],
  properties: {
    tex: { type: 'string' },
    totalComments: { type: 'number', description: 'how many comments you actually triaged for this file' },
    fileSummary: { type: 'string', description: '2-3 sentence overview of the dominant themes Kritikos raises in this chapter' },
    items: {
      type: 'array',
      items: {
        type: 'object',
        required: ['idx', 'page', 'section', 'kritikosComment', 'classification', 'severity', 'proposedAction', 'effort'],
        properties: {
          idx: { type: 'number', description: 'the comment idx from the JSON' },
          page: { type: 'number' },
          section: { type: 'string', description: 'section title from the comment record' },
          anchorQuote: { type: 'string', description: 'the exact ~8-word phrase in the live .tex this comment attaches to (quote it so the author can find it). If anchor no longer found in source, say "ANCHOR-NOT-FOUND".' },
          kritikosComment: { type: 'string', description: 'Kritikos comment, lightly cleaned of OCR artefacts' },
          classification: {
            type: 'string',
            enum: ['trivial', 'rewrite', 'pushback', 'table-fix', 'figure', 'structural', 'reference-bug', 'cross-cutting'],
            description: 'trivial=typo/wording/one-liner; rewrite=substantive prose change; pushback=author may want to defend, not concede; table-fix=table data/values wrong; figure=figure inconsistency or quality; structural=move/reorganize content; reference-bug=non-existent citation number [NNN]; cross-cutting=touches multiple chapters',
          },
          severity: { type: 'string', enum: ['blocking', 'major', 'minor'], description: 'blocking=defense/correctness risk if unaddressed; major=clearly expected; minor=nice-to-have polish' },
          relatedHeadlinePoint: { type: 'number', description: 'which of Kritikos 15 summary points (1-15) this maps to, or 0 if none' },
          proposedAction: { type: 'string', description: 'concrete, specific action the author should take. Reference the actual line/text in the live .tex. If mechanical, give the exact replacement.' },
          mechanical: { type: 'boolean', description: 'true if this is a safe automatable edit (typo, ref number, wording) needing no judgement; false if it needs author thought' },
          effort: { type: 'string', enum: ['1min', '5min', '15min', '30min+'] },
        },
      },
    },
  },
}

phase('Triage')
const triaged = await parallel(FILES.map(f => () =>
  agent(
    `You are triaging supervisor (Prof. Kritikos) PDF comments on a master's thesis chapter, for the author to act on.

CONTEXT YOU MUST READ FIRST:
1. ${REVIEW}/kritikos-inline-annotations.json — JSON array of ALL 480 comments. Each record has: idx, page, type, author, comment (Kritikos's text), context (surrounding thesis text from the PDF), chapter, section, tex. FILTER to records where "tex" == "${f.tex}". Those are YOUR comments (~${f.n} of them).
2. The LIVE chapter source. For a real chapter file, read: ${CHAPTERS}/${f.tex}. For "front-matter", the relevant files are ${CHAPTERS}/abstract-en.tex, ${CHAPTERS}/abstract-gr.tex, ${CHAPTERS}/lists-and-acronyms.tex and /home/meter-peter/development/driveby/thesis/main.tex — read whichever the comment's page (1-21, front matter) points to.
3. ${REVIEW}/kritikos-15point-email-thread.md — Kritikos's 15 headline points + the author's email replies. Use this ONLY to set the "relatedHeadlinePoint" field and to recognize where the author has already taken a position.

CRITICAL FACTS:
- The PDF Kritikos annotated was built from an OLDER source than the live .tex. The author has since edited principles and other content. So the "context" field in the JSON (PDF text) may NOT match the live .tex. DO NOT assume the comment is resolved — treat EVERY comment as still open and produce an action for it. But when you quote the anchor, quote it from the LIVE .tex; if you cannot find the anchored text in the live source at all, set anchorQuote to "ANCHOR-NOT-FOUND" and note in proposedAction that the author should verify whether newer edits already cover it.
- Do NOT edit any files. This is triage only.
- Be concrete. "Clarify this sentence" is useless. Say WHICH sentence (quote it) and HOW to clarify it. For reference-bug comments (Kritikos flags non-existent citation numbers like [161], [176], [192], [194]), the action is almost always: remove the stale bracketed number / convert the round-N-review provenance note into self-contained prose (per his headline point #5). For table-fix comments, say which table and what's numerically wrong.

For EACH of your ~${f.n} comments, produce one item in the schema. Triage ALL of them — do not sample or truncate. If a comment is vague OCR noise with no actionable content, still emit it classified as trivial with proposedAction "no action / OCR artefact — verify in PDF".

Return the structured object. totalComments must equal the number of items you emit.`,
    { label: `triage:${f.tex}`, phase: 'Triage', schema: TRIAGE_SCHEMA }
  )
))

const triagedOk = triaged.filter(Boolean)
const allItems = triagedOk.flatMap(t => t.items || [])
log(`Triage done: ${triagedOk.length}/${FILES.length} files, ${allItems.length} comments classified`)

// ---- Phase 2: independent second opinion on the 15 headline points ----
const POINTS = Array.from({ length: 15 }, (_, i) => i + 1)

const OPINION_SCHEMA = {
  type: 'object',
  required: ['point', 'kritikosAsk', 'authorEmailPosition', 'independentVerdict', 'recommendation', 'liveSourceEvidence'],
  properties: {
    point: { type: 'number' },
    kritikosAsk: { type: 'string', description: 'one-line restatement of what Kritikos demands' },
    authorEmailPosition: { type: 'string', description: 'one-line restatement of what the author replied in the email' },
    liveSourceEvidence: { type: 'string', description: 'what the LIVE thesis source actually says about this now — quote file:line or note "not yet in source". Does the source already do what Kritikos asks, partially, or not at all?' },
    independentVerdict: {
      type: 'string',
      enum: ['author-position-holds', 'author-should-concede', 'author-partially-right', 'already-resolved-in-source', 'needs-author-decision'],
    },
    reasoning: { type: 'string', description: 'why you reached that verdict, on technical merit — 2-4 sentences' },
    recommendation: { type: 'string', description: 'the single concrete thing the author should do before locking his reply' },
    risk: { type: 'string', enum: ['high', 'medium', 'low'], description: 'risk to the defense if the author keeps his current email position' },
  },
}

phase('SecondOpinion')
const opinions = await parallel(POINTS.map(p => () =>
  agent(
    `You are a thesis examiner giving the AUTHOR a second opinion on ONE of his supervisor's 15 headline review points, BEFORE he locks in his email reply.

READ:
1. ${REVIEW}/kritikos-15point-email-thread.md — find point (${p}). Read Kritikos's demand AND the author's email reply for that exact point.
2. The relevant LIVE thesis source to check what the document ACTUALLY says now. Point you should chase into source:
   - Point 3 (single-artifact / Table 2.1): ${CHAPTERS}/02-related-work.tex — find Table 2.1 and its last criterion.
   - Point 6 (precision within completeness) and 7 (axioms govern all principles) and 8 (mode-specific behavior) and 9 (test-ready principle): ${CHAPTERS}/03-methodology-ddt.tex.
   - Point 10 (architecture as own chapter) and 11 (Ch.4 text/figure inconsistency): ${CHAPTERS}/04-cli-architecture.tex and ${CHAPTERS}/05-kubernetes-architecture.tex.
   - Point 12 (closed-loop negative feedback naming): grep "closed-loop" across ${CHAPTERS}/06-gitops-pipeline.tex and 09-discussion.tex.
   - Point 13 (table content wrong / multi-API one-fault claim): ${CHAPTERS}/07-evaluation.tex — find the multi-API controlled experiment tables.
   - Point 14 (P007 must be critical): ${CHAPTERS}/03-methodology-ddt.tex (severity) + ${CHAPTERS}/07-evaluation.tex.
   - Point 15 (Table 8.2 split code vs report): ${CHAPTERS}/08-ai-assisted-development.tex.
   - Points 1,2,4,5: chase the relevant chapter; use grep/Read to verify the current state.

YOUR JOB: independently judge, on technical merit, whether the AUTHOR'S EMAIL POSITION is sound. The author has chosen "re-evaluate independently" — so where his email pushes back (points 3, 9) or agrees (14, 15) or defers, tell him honestly whether that position will survive a defense, or whether he should change it. Quote what the live source already says. If the source already satisfies Kritikos, say so (verdict already-resolved-in-source) — that's the strongest reply he can give.

Be a critical friend, not a yes-man. If the author is wrong (e.g. on the single-artifact claim, Kritikos has a real point that two artifacts carry the requirements), say so plainly and give him the defensible reframing.

Return the structured object for point ${p}.`,
    { label: `opinion:P${p}`, phase: 'SecondOpinion', schema: OPINION_SCHEMA }
  )
))

const opinionsOk = opinions.filter(Boolean).sort((a, b) => a.point - b.point)
log(`Second opinions done: ${opinionsOk.length}/15 points`)

// ---- Phase 3: synthesize the action plan ----
phase('Synthesize')

// Pre-compute rollups in plain JS so the synthesizer has clean numbers.
const byFile = {}
for (const t of triagedOk) byFile[t.tex] = { summary: t.fileSummary, count: (t.items || []).length }
const bySeverity = { blocking: 0, major: 0, minor: 0 }
const byClass = {}
let mechanicalCount = 0
for (const it of allItems) {
  if (bySeverity[it.severity] !== undefined) bySeverity[it.severity]++
  byClass[it.classification] = (byClass[it.classification] || 0) + 1
  if (it.mechanical) mechanicalCount++
}

const rollup = { totalTriaged: allItems.length, byFile, bySeverity, byClass, mechanicalCount }

const planMd = await agent(
  `You are writing the master action plan that the thesis author will work from to address all of Prof. Kritikos's review. Write it as a single, well-structured Markdown document. This is the deliverable.

You are given two JSON datasets inline:

=== TRIAGE (per-comment classification, ${allItems.length} items across ${triagedOk.length} files) ===
${JSON.stringify({ rollup, files: triagedOk.map(t => ({ tex: t.tex, fileSummary: t.fileSummary, items: t.items })) }).slice(0, 240000)}

=== SECOND OPINIONS (independent verdicts on the 15 headline points) ===
${JSON.stringify(opinionsOk)}

Write the document with these sections, in this order:

# Kritikos Review — Action Plan
(one short paragraph: 480 inline comments + 15 headline points, triaged against the LIVE source; note the PDF was built from older source so some items may already be resolved and the author should confirm-and-skip those.)

## At a Glance
A compact table: per-file comment counts, and the severity rollup (blocking/major/minor), and the count of mechanical (safely automatable) edits. Use the rollup numbers given.

## The 15 Headline Points — Second Opinion
For EACH of the 15 points, a short block: **Point N — <one-line>**, then: Kritikos asks / Your email said / What the source says now / **Verdict** (in bold) / **Do this** (the recommendation) / Risk. Make the verdict scannable. Put the points where the author's email position is WRONG or RISKY (verdict author-should-concede or risk high) at visual prominence — flag them with a ⚠ marker. Be explicit and honest about point 3 (single-artifact) and point 9 (test-ready principle) since the author pushed back there.

## Per-File Work Lists
For each of the 11 files, in descending comment-count order: a "### <filename> (<N> comments)" heading, the file summary, then a Markdown table of its items with columns: idx | page | severity | class | Kritikos's point (abbreviated) | Proposed action (abbreviated to ~15 words) | →#15pt. Sort each file's table by severity (blocking, then major, then minor). Keep proposed-action cells tight but specific.

## Mechanical Sweep (do these in one pass)
A flat checklist of every item with mechanical=true, grouped by file, formatted as "- [ ] {file}:{idx} (p{page}) — {action}". These are the safe edits the author (or a follow-up apply-workflow) can do without judgement. Call out the reference-bug cleanup (stale [NNN] numbers + round-N-review provenance notes, headline point #5) as its own subsection since it spans many files.

## Recommended Order of Attack
A numbered list: what to do first (blocking + structural like the architecture-chapter split point 10), what to batch (mechanical sweep), what needs a decision from the author (the pushback points), what's last (minor polish). Give a rough total-effort sense.

Rules: be concrete, quote anchors where helpful, never invent comments not in the data, and keep it skimmable — the author will work top to bottom. Output ONLY the Markdown document, no preamble.`,
  { label: 'synthesize:action-plan', phase: 'Synthesize' }
)

return { planMd, rollup, opinionCount: opinionsOk.length, triagedFiles: triagedOk.length, totalComments: allItems.length }
