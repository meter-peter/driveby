# Documentation-Driven Testing: Axiomatic Foundation

## Preamble

Documentation-Driven Testing (DDT) is a methodology for automated API quality assurance that treats machine-readable API specifications as executable infrastructure. Rather than writing tests separately from documentation, DDT derives test cases, validation rules, and quality gates directly from the API specification itself.

This document formalizes the axiomatic foundation of DDT as presented in the thesis: *"Documentation-Driven Testing (DDT): A Paradigm for Automated API Quality Assurance in the GitOps Era."*

## The Three Axioms

### Axiom 1: Completeness

> **A fully documented API contains all necessary data to verify correctness.**

The Completeness Axiom asserts that a sufficiently detailed API specification is, by itself, a complete test oracle. If the specification describes every endpoint, every parameter, every schema, every error response, and every constraint, then no external test specification is needed — the documentation *is* the test specification.

This is a strong claim. Traditional testing methodologies assume that test cases are a separate artifact, written by testers who interpret requirements. DDT eliminates this interpretation step: the specification is both the requirement and the test, provided it is complete.

**Formal statement:** Let `S` be an API specification and `T(S)` be the set of test cases derivable from `S`. The Completeness Axiom holds if and only if `T(S)` covers all functional behaviors of the API implementation `I`. That is, for every observable behavior `b` of `I`, there exists a test `t` in `T(S)` that exercises `b`.

**Corollary:** An incomplete specification necessarily produces incomplete test coverage. The degree of specification completeness directly determines the degree of achievable automated test coverage.

**Principles derived from Completeness:**
- **P001 (OpenAPI Compliance):** The specification must be structurally parseable — a prerequisite for any derivation.
- **P002 (Documentation Quality):** Every API element must be described — the content that enables test derivation.
- **P003 (Error Handling):** Failure modes must be documented — negative testing requires error specifications.
- **P004 (Schema Definitions):** Data contracts must be precise — payload validation requires exact schemas.

### Axiom 2: Determinism

> **Documented examples imply reproducible behavior.**

The Determinism Axiom asserts that if the specification documents an example request-response pair, then the live API must produce the documented response when given the documented request, every time, without exception (given identical state). This is the bridge between static documentation and runtime verification.

Determinism is what makes contract testing possible. Without it, specifications are aspirational documents rather than executable contracts. DDT requires that specifications are not merely descriptive but *prescriptive* — they define what the API *must* do, not what it *might* do.

**Formal statement:** Let `(req, res)` be a request-response pair documented in specification `S`, and let `I(req)` be the response produced by implementation `I` when given request `req`. The Determinism Axiom holds if and only if `I(req) = res` for all documented pairs, given identical preconditions.

**Corollary:** Non-deterministic APIs (those whose responses depend on uncontrolled external state) require specification annotations that declare the non-deterministic dimensions, enabling test harnesses to account for variability.

**Principles derived from Determinism:**
- **P006 (Contract Testing):** The live API must match its specification — the empirical test of determinism.

### Axiom 3: Observability

> **Machine-readable descriptions enable automated analysis.**

The Observability Axiom asserts that API quality attributes — security posture, performance characteristics, versioning strategy — are only manageable if they are expressed in a machine-readable format. Human-readable prose in a wiki is not observable by automation; structured fields in an OpenAPI specification are.

This axiom extends DDT beyond functional correctness into operational quality. Security schemes, SLA thresholds, deprecation notices, and version metadata are all quality-relevant properties that can be automatically extracted, monitored, and enforced — but only if they are expressed in a format that tools can parse.

**Formal statement:** Let `Q` be a quality attribute of API implementation `I`. The Observability Axiom holds if and only if there exists a machine-readable representation `q` in specification `S` such that `Q` can be evaluated by an automated tool reading `q` without human interpretation.

**Corollary:** Quality attributes that lack machine-readable representation in the specification are invisible to DDT and must be addressed by other means. DDT's coverage is bounded by the expressiveness of the specification format.

**Principles derived from Observability:**
- **P005 (Security Standards):** Security schemes expressed as structured specification elements.
- **P007 (Performance Requirements):** Performance thresholds expressed as measurable, reportable metrics.
- **P008 (Versioning Strategy):** Version metadata expressed as structured fields enabling lifecycle automation.

## Axiom-to-Principle Mapping

```
  ┌───────────────────┐    ┌──────────────┐    ┌───────────────────┐
  │   COMPLETENESS    │    │ DETERMINISM  │    │  OBSERVABILITY    │
  │ "Spec describes   │    │ "Same input  │    │ "Results are      │
  │  the full API"    │    │  = same out" │    │  measurable"      │
  └─┬───┬───┬───┬────┘    └──────┬───────┘    └──┬───┬───┬───────┘
    │   │   │   │                │                │   │   │
    ▼   ▼   ▼   ▼                ▼                ▼   ▼   ▼
  P001 P002 P003 P004          P006             P005 P007 P008
  Comp Doc  Err  Schema        Func             Sec  Perf Vers
  ━━━━ ──── ──── ────          ━━━━             ━━━━ ──── ────
  CRIT WARN WARN WARN          CRIT             CRIT WARN WARN

  ━━━━ = Critical severity     ──── = Warning severity
```

| Axiom | Principles | Coverage Domain |
|-------|-----------|-----------------|
| Completeness | P001, P002, P003, P004 | Specification content and structure |
| Determinism | P006 | Runtime behavior verification |
| Observability | P005, P007, P008 | Machine-readable quality attributes |

## Completeness of the Axiomatic System

The three axioms form a complete system for API quality assurance:

1. **Completeness** ensures the specification has enough information (content coverage).
2. **Determinism** ensures the specification accurately reflects reality (behavioral correctness).
3. **Observability** ensures the specification is actionable by automation (machine readability).

Together, they address the three fundamental questions of API quality:
- *Is the documentation thorough?* (Completeness)
- *Does the API do what the documentation says?* (Determinism)
- *Can we automatically verify and enforce quality?* (Observability)

Any API quality concern that cannot be mapped to one of these three axioms falls outside the scope of DDT. This is a deliberate boundary: DDT does not claim to address business logic correctness, user experience, or domain-specific semantics — only those quality attributes that are derivable from the API specification.

## Novelty Claim

DDT is the **first methodology to treat API documentation as executable infrastructure**. Prior approaches fall into two categories:

1. **Test-then-document:** Tests are written first; documentation is generated or maintained separately. The documentation may drift from the tests.
2. **Document-then-test:** Documentation is written first; tests are written separately to verify the documentation. The tests may drift from the documentation.

DDT eliminates the drift problem by **collapsing documentation and test specification into a single artifact**. The OpenAPI specification is simultaneously the documentation, the test oracle, and the quality gate. Changes to the specification automatically change the validation criteria, and validation failures automatically indicate specification drift.

## Source References

- Thesis document: `docs/Thesis.md`
- Product vision: `docs/prd-vision-thesis.md`
- Principle specifications: `docs/principles/P001.md` through `docs/principles/P008.md`
- Implementation: `driveby-cli/internal/principles/` (principle checkers)
- Engine: `driveby-cli/internal/engine/` (orchestration)
