# Defense Presentation — Slide Deck Outline

**Defense:** Friday 19 June 2026, 16:30, University of the Aegean, Samos
**Duration:** 30 minutes (~20 slides → ~90s each)
**Closing visual:** Closed control loop diagram (Kritikos [605])

## Design principles

- **One idea per slide.** No bullets longer than 4 lines.
- **Diagrams > prose.** Every slide has a primary visual; text supports the visual, not the reverse.
- **Three pillars structure**: procedures → primitive ops → closed control loop.
- **Reuse existing TikZ figures from the thesis** wherever possible — they're already done, supervisor-approved, and stylistically consistent.
- **Greek or English?** Recommendation: **English slides, Greek delivery** (most University of the Aegean defenses in English-titled MSc/diploma theses do this; the supervisor reads English with no friction; slides match the thesis).

---

## Act 1 — The Problem (slides 1–5, ~7 minutes)

### Slide 1 — Title

```
Specification-Driven Testing (SDT):
A Paradigm for Automated API Quality
Assurance in the GitOps Era

Πέτρος Ευάγγελος Τριανταφύλλης
Supervisor: Prof. Kyriakos Kritikos
University of the Aegean — Dept. of Information &
Communication Systems Engineering
19 June 2026
```

No visual. Just clean title.

### Slide 2 — The procedure today (manual / fragmented)

**Visual:** a horizontal flow with broken arrows / question marks:

```
[Developer writes OpenAPI] → [Spec linter (Spectral)] → ?
                                                       ↘
                       [Contract test (Dredd)] → [Server]
                                                       ↗
                       [Load test (k6)] → ?
                                                       ↘
                                      [Manual ClickOps]
                                                       ↓
                                    [Deploy → discover problem]
```

**Caption (≤2 lines):** Each tool answers one question; nothing answers the meta-question — *is the specification itself good enough?*

### Slide 3 — The three failure modes

**Visual:** a 3-cell strip, one icon per cell:

| ❌ | ❌ | ❌ |
|---|---|---|
| **Spec is incomplete** (linters say it parses) | **Implementation drifts** (tests run too late) | **Pipeline is manual** (ClickOps everywhere) |

### Slide 4 — Research questions (RQ1–RQ4)

Plain text, one per line, no visual:

- **RQ1** — Can a framework demonstrate automated API quality assurance without dedicated testers in the promotion path?
- **RQ2** — To what extent can validation rules, test cases, and quality gates be derived automatically from OpenAPI specifications?
- **RQ3** — How effectively does this integrate into cloud-native GitOps pipelines?
- **RQ4** — Can quality-gated delivery pipelines be expressed as Kubernetes-native *ontological specifications*?

### Slide 5 — Three contributions (the deck's roadmap)

**Visual:** three stacked horizontal bands, each with its own colour:

```
┌──────────────────────────────────────────────────┐
│ 1. SDT framework — 3 axioms, 9 principles        │   ← Act 1.5
│                  (a conceptual framework)        │
├──────────────────────────────────────────────────┤
│ 2. DriveBy CLI + XSDLC composition               │   ← Act 2
│                  (the primitive ops)             │
├──────────────────────────────────────────────────┤
│ 3. Closed-loop quality control                   │   ← Act 3
│                  (the vision)                    │
└──────────────────────────────────────────────────┘
```

---

## Act 2 — Primitive ops (slides 6–14, ~13 minutes)

### Slide 6 — The three axioms (the foundation)

**Visual:** a clean triangle, one axiom per vertex:

```
                Completeness
                     △
                   /   \
                  /     \
                 /       \
                /         \
       Determinism ─── Observability
```

**Caption:** Three properties any specification-driven validation must satisfy. The whole framework derives from these.

### Slide 7 — The nine principles (the primitive operations)

**Visual:** the existing `axiom-principle-mapping` table from Ch.3 (or a simplified version):

| ID | Principle | Axiom | Severity |
|---|---|---|---|
| P001 | OpenAPI Compliance | Completeness | Critical |
| P002 | Documentation Quality | Completeness | Critical |
| P003 | Error Handling | Completeness | Critical |
| P004 | Schema Definitions | Completeness | Critical |
| P005 | Security Standards | Observability | Critical |
| P006 | Functional Testing | Determinism | Critical (runtime) |
| P007 | Performance Testing | Observability | Warning |
| P008 | Versioning Strategy | Observability | Warning |
| P009 | Test Readiness | Determinism | Warning |

**Caption:** Each principle is an atomic check derivable from the OpenAPI document. P001–P005 are static; P006–P007 are runtime; P008–P009 are advisory.

### Slide 8 — Validation modes (composition primitives)

**Visual:** a 2×2 quadrant or a simple table:

```
                      Static       Runtime
              ┌─────────────────────────────┐
   Light      │  minimal       test-ready  │
              │  (P001)        (P001–P004+P009) │
              ├─────────────────────────────┤
   Heavy      │  strict        test-only   │
              │  (P001–P005,   (P006+P007)  │
              │   P008)                     │
              └─────────────────────────────┘
```

**Caption:** Modes compose principles into use-case-shaped check sets. `test-only` is **orthogonal** to the static modes — they're not incremental (this answers Kritikos [244]/[245] proactively).

### Slide 9 — DriveBy: the validator (single artifact)

**Visual:** a simple flow diagram:

```
[OpenAPI spec] → [DriveBy CLI] → [JSON report]
                       │
                       ├── load (auto-detect 3.x or 2.0)
                       ├── normalize (APISpec adapter)
                       ├── run principle checkers
                       └── emit (Critical / Warning / Pass)
```

**Caption:** One Go binary. Reads any OpenAPI version. Produces structured JSON. The JSON Schema for that output is now published — `driveby-cli/schemas/report.schema.json`.

### Slide 10 — XSDLC: from CLI tool to platform feature

**Visual:** the existing `xrd-resource-hierarchy` figure (Fig 5.3 from Ch.5) — shows ~45 K8s objects fanning out from one CR.

**Caption:** A single Kubernetes Custom Resource (~35 lines of YAML) declaratively provisions an entire quality-gated GitOps delivery pipeline. The CR *is* the specification.

### Slide 11 — The two-repo BYOCI model

**Visual:** reuse the now-fixed `fig:single-repo-branch-flow` (Fig 6.2) — software repo on one side, GitOps repo on the other, BYOCI boundary line between.

**Caption:** Developer owns CI (build/test/image). XSDLC owns delivery (promotion/gates/sync). The boundary is enforced by the GitOps repo separation.

### Slide 12 — The gate procedure (end-to-end)

**Visual:** reuse `fig:gitops-pipeline` (Fig 6.1) — the swimlane diagram with the labels-inside-lanes you just fixed.

**Caption:** 6 lanes, 1 webhook, 1 Argo Workflow, 1 commit status → 1 promotion decision. End to end in <15s for a passing run.

### Slide 13 — Multi-environment promotion (the procedural primitive)

**Visual:** reuse `fig:multi-env-promotion` (Fig 6.3) — Development → SDT Gate → Staging → Manual Approval → Production.

**Caption:** Same gate template instantiated per environment. Different modes per stage (test-ready in staging, strict + load in production).

### Slide 14 — What's actually being measured

**Visual:** a stacked-bar of the per-principle check counts (mini-version of `fig:per-principle-pass-rates`):

```
P001: ████████ 7 checks
P002: ██████████ 10 checks
P003: ███████ 7 checks
P004: ██████████ 10 checks
P005: ███████ 7 checks
P006: ███████ 7 checks (runtime)
P008: ███████ 7 checks
P009: ████ 4 checks
```

**Caption:** 59 atomic checks across 8 implemented principles. Each one derivable from the specification.

---

## Act 3 — Evidence + the closed loop (slides 15–20, ~10 minutes)

### Slide 15 — Evaluation: three arms

**Visual:** reuse `fig:evaluation-methodology` (Fig 7.1) — three columns:

```
[PoC operational]    [Controlled]         [Large-scale]
  5 APIs              non-critical-api +    50 OpenAPI 3.x
  Real Novelcore      4 defect variants     + 20 Swagger 2.0
  cluster             15 mutations         from APIs.guru
  44 workflows        12/15 detected
```

**Caption:** Proof-of-concept proves it runs. Controlled proves it detects what you inject. Large-scale proves it generalizes.

### Slide 16 — The 84% finding

**Visual:** a big stark chart or table — the headline empirical result:

```
APIs.guru large-scale (50 OpenAPI 3.x + 20 Swagger 2.0):

  84% of public APIs score 1/6 or worse in strict mode
  No API in the sample passes more than 2/6 principles

  P001 (Compliance):     46% pass
  P002 (Documentation):   0% pass
  P003 (Errors):          0% pass
  P004 (Schemas):        12% pass
  P005 (Security):       26% pass
  P008 (Versioning):      0% pass
```

**Caption:** This is the *systemic* specification-quality gap. Not anecdotal. Not edge cases.

### Slide 17 — Agent-feedback experiment (the small punch)

**Visual:** before-after of Swagger Petstore:

```
Baseline:           After 1 agent iteration:

   P001: ✓             P001: ✓
   P002: ✗             P002: ✓
   P003: ✗             P003: ✓
   P004: ✗             P004: ✓ ← agent added typed schemas
   P005: ✗             P005: ✗ (still missing security)
   P008: ✗             P008: ✗

   1/6                 4/6 (+300%)

   The agent received only the DriveBy report as input.
   No prompts about what to fix. Just the structured output.
```

**Caption:** SDT is feedback infrastructure for agent-driven correction. The report is machine-readable enough that an AI coding agent can act on it autonomously.

### Slide 18 — From procedure to closed loop (the transition)

**Visual:** a side-by-side, showing the conceptual lift:

```
  Procedure (Ch.6 §6.1–§6.4)        Closed loop (Ch.6 §6.5)

  Developer pushes code             Specification = setpoint
       ↓                                  ↓
  Pipeline runs                     DriveBy = sensor
       ↓                                  ↓
  Gate decides                      Workflow = controller
       ↓                                  ↓
  Promotion happens                 Promoter = actuator
                                          ↓
                                    Deployment state → feedback
                                          ↑
                                    [back to setpoint]
```

**Caption:** The pipeline isn't just a procedure — it's a control loop. The specification is the setpoint, and the gate enforces convergence toward it.

### Slide 19 — The closed control loop (Kritikos's vision, [605])

**Visual:** reuse `fig:quality-control-loop` (Fig 6.5) — the full setpoint/sensor/controller/actuator/plant diagram in colour. This is the centerpiece visual.

**Caption (the line you say out loud):**
> "Comment [605] in the supervisor's review framed this as the main vision: a closed control loop with multiple levels, potentially operating across clusters, that self-adjusts as the specification evolves. The thesis delivers the inner loop; the outer loops — adaptive thresholds, agent-driven remediation, multi-cluster orchestration — are the natural research continuation."

This is **the slide you build everything toward**. It does triple duty:
1. Acknowledges Kritikos's contribution explicitly
2. Names the future-work directions concretely
3. Lands on a vision, not a result

### Slide 20 — Contributions, limitations, future work, thanks

Three columns:

| Contributions | Limitations | Future work |
|---|---|---|
| 1. SDT framework | P006/P007 wrappers pending | Adaptive thresholds |
| 2. DriveBy CLI | Single-cluster PoC | Risk-based scoring |
| 3. XSDLC composition | English/EN OpenAPI only | Agent-driven remediation |
| 4. Empirical evidence | | Multi-format (GraphQL, gRPC) |
| 5. Agent-feedback proof | | Closed multi-loop control |
| 6. Agent-driven dev method | | |

**Centered bottom line:**
> Thank you. Questions?

---

## Production checklist

| step | tool | effort |
|------|------|--------|
| 1. Pick template (LaTeX Beamer or PowerPoint/Keynote?) | — | 0 |
| 2. Extract the 6 figures that are reused (Fig 5.3, 6.1, 6.2, 6.3, 6.5, 7.1) into standalone PDF crops | `pdfcrop` or screenshot | 30 min |
| 3. Build slides 1–5 (Act 1) | template | 1 hour |
| 4. Build slides 6–14 (Act 2 — primitives) | template | 2 hours |
| 5. Build slides 15–20 (Act 3 — evidence + loop) | template | 1.5 hours |
| 6. Practice run 1 — timing check | clock | 40 min |
| 7. Practice run 2 — polish transitions | — | 40 min |
| 8. Practice run 3 — full delivery | — | 40 min |

**Total: ~7 hours of slide work + ~2 hours of rehearsal.**

## Open questions (decide before slides start)

1. **Format**: LaTeX Beamer (matches thesis style, ugly transitions) or PowerPoint/Keynote (prettier, two-format mismatch)?
2. **Live demo**: include a 60-second `driveby validate-only` run on Petstore, live, or just screenshots? (Live = risky but memorable; screenshots = safe.)
3. **Backup slides** (after slide 20, hidden): one slide per `discussed` rebuttal in case Kritikos probes? I'd say yes — 5–10 hidden slides cost nothing and protect against surprise questions.

## What I recommend you do next

1. Pick **LaTeX Beamer vs PowerPoint** (item 1 above) — I can scaffold either.
2. Pick **live demo vs screenshots** (item 2) — both work, just commit.
3. Then I write the actual deck.
