# Thesis Writing Agent

## Thesis
**Title**: Documentation-Driven Testing (DDT): A Paradigm for Automated API Quality Assurance in the GitOps Era

## Files
```
thesis/
  main.tex            # Root LaTeX document, includes all chapters
  references.bib      # BibTeX references — cite all claims
  Makefile             # Build: make pdf, make clean
  figures/             # All figures referenced via \includegraphics
  chapters/
    01-introduction.tex
    02-related-work.tex
    03-methodology-ddt.tex
    04-cli-architecture.tex
    05-kubernetes-architecture.tex
    06-gitops-pipeline.tex
    07-conclusion.tex
```

## Chapter-by-Chapter Specification

### Chapter 1 — Introduction (`01-introduction.tex`)
**Expected sections**: Problem Statement, Motivation, Research Questions, Contributions, Thesis Structure
**Source files to read**: `docs/prd-vision-thesis.md`, `docs/Thesis.md`
**Code files to read**: None (conceptual chapter)
**Key content**: Define the problem (API quality assurance lacks automation tied to documentation), state 2-3 research questions, list contributions (DDT paradigm, DriveBy tool, evaluation)

### Chapter 2 — Related Work (`02-related-work.tex`)
**Expected sections**: API Testing (contract testing, fuzz testing), OpenAPI Ecosystem, GitOps and Continuous Delivery, Crossplane and Infrastructure-as-Code, Quality Assurance Frameworks
**Source files to read**: `docs/prd-vision-thesis.md` (positioning against existing tools)
**Code files to read**: None
**Key content**: Position DDT against Schemathesis, Dredd, Spectral, Optic. Distinguish documentation-driven from contract-driven testing. Cover GitOps fundamentals (ArgoCD, Flux).

### Chapter 3 — Methodology: DDT (`03-methodology-ddt.tex`)
**Expected sections**: DDT Paradigm Overview, The Three Axioms (Completeness, Determinism, Observability), The Nine Principles (P001-P009), Axiom-to-Principle Mapping, Validation Modes
**Source files to read**: `docs/principles/P001-P009.md`, `docs/ddt-axioms.md`
**Code files to read**:
- `driveby-cli/internal/principles/checker.go` — PrincipleChecker interface
- `driveby-cli/internal/principles/p001_compliance.go` through `p008_versioning.go` — each principle's checks
- `driveby-cli/internal/principles/registry.go` — mode-based selection
- `driveby-cli/internal/types/modes.go` — ValidationMode definitions
**Key content**:
- Define 3 axioms: Completeness (spec fully describes API), Determinism (same input = same output), Observability (validation results are measurable)
- Define 9 principles with formal identifiers P001-P009
- Axiom mapping: Completeness -> P001-P004, Determinism -> P006, Observability -> P005/P007/P008
- Validation modes: minimal, strict, test-only, test-ready

### Chapter 4 — CLI Architecture (`04-cli-architecture.tex`)
**Expected sections**: Architecture Overview, Dependency Flow, Spec Abstraction Layer, Engine Design, Principle Checker Pattern, Report Generation, CLI Design
**Source files to read**: `docs/CLI_USAGE.md`, `docs/architecture.md`
**Code files to read**:
- `driveby-cli/internal/spec/spec.go` — APISpec interface
- `driveby-cli/internal/spec/openapi3.go` — OpenAPI 3.x adapter
- `driveby-cli/internal/spec/swagger2.go` — Swagger 2.x adapter
- `driveby-cli/internal/engine/engine.go` — orchestration logic
- `driveby-cli/internal/loader/loader.go` — spec loading
- `driveby-cli/internal/cli/root.go`, `validate.go`, `config.go` — CLI structure
- `driveby-cli/internal/report/generator.go`, `markdown.go` — report generation
- `driveby-cli/internal/types/` — all type definitions
**Key content**:
- Dependency flow diagram: types <- spec <- loader <- principles <- testing <- engine <- cli
- APISpec adapter pattern (version-agnostic validation)
- Engine orchestration: load spec -> select principles by mode -> run checkers -> aggregate -> report

### Chapter 5 — Kubernetes Architecture (`05-kubernetes-architecture.tex`)
**Expected sections**: Crossplane XRD Design, XSDLC Composition Architecture, Declarative Resource Generation, Two-Tier Configuration, Single-Repo Branch Model, BYOCI Delivery Pipeline
**Source files to read**: `docs/deployment-guide.md`, `docs/quality-gate-sdlc.md`, `docs/responsibility-matrix.md`
**Code files to read**:
- `kubernetes/helm/driveby/templates/crossplane/composition-sdlc.yaml` — XSDLC composition
- `kubernetes/helm/driveby/templates/crossplane/xrd-sdlc.yaml` — XRD schema
- `kubernetes/examples/novelcore-perfect-api/` — reference XSDLC CR
- `kubernetes/examples/gitops-promoter/` — promoter CRD examples
- `kubernetes/helm/driveby/values.yaml` — two-tier defaults
- `kubernetes/helm/driveby/Chart.yaml` — chart metadata
**Key content**:
- XSDLC as single-CR declarative pipeline (~35 YAML lines → ~34 resources)
- BYOCI model: XSDLC is a delivery pipeline, not a CI system
- Two-tier config: values.yaml defaults → XRD spec overrides (no EnvironmentConfigs)
- Dynamic environment chain (any N environments, minimum 2)
- Auto-generated GitHub workflow (driveby-deploy — manual trigger deploy pipeline)
- Target cluster: private.novelcore.org

### Chapter 6 — GitOps Pipeline (`06-gitops-pipeline.tex`)
**Expected sections**: Single-Repo Architecture, BYOCI Model, Promotion Flow, Quality Gate Implementation, Event-Driven Validation, Production Deployment, Evaluation
**Source files to read**: `docs/gitops-pipeline.md`, `docs/quality-gate-sdlc.md`, `docs/evaluation-methodology.md`
**Code files to read**:
- `kubernetes/helm/driveby/templates/crossplane/composition-sdlc.yaml` — composition
- `kubernetes/examples/novelcore-perfect-api/` — reference deployment
- `apis/perfect-api/perfect-api.py` — reference API implementation
- `apis/perfect-api/openapi.json` — reference spec
- `tools/harvest-openapi-apisguru.py` — dataset harvesting
- `tools/run-openapi-batch.sh` — batch validation
**Key content**:
- End-to-end flow: developer triggers deploy → manifests to env-next → promote → gate → merge → sync
- BYOCI: developer's CI builds; XSDLC delivers
- Production deployment on private.novelcore.org
- Controlled evaluation: perfect-api against all principles
- Large-scale evaluation: APIs.guru dataset

### Chapter 7 — Conclusion (`07-conclusion.tex`)
**Expected sections**: Summary of Contributions, Answers to Research Questions, Limitations, Future Work
**Source files to read**: `docs/Thesis.md`
**Code files to read**: None
**Key content**: Summarize DDT paradigm contribution, DriveBy tool contribution, XSDLC delivery pipeline contribution, evaluation findings. Future work: broader language support, enterprise adoption, additional check types.

## Complete Feed Map

| Thesis Section | Source Code | Documentation |
|---------------|-------------|---------------|
| DDT Axioms | `internal/types/modes.go` | `docs/ddt-axioms.md` |
| P001 Compliance | `internal/principles/p001_compliance.go` | `docs/principles/P001.md` |
| P002 Documentation | `internal/principles/p002_documentation.go` | `docs/principles/P002.md` |
| P003 Errors | `internal/principles/p003_errors.go` | `docs/principles/P003.md` |
| P004 Schema | `internal/principles/p004_schema.go` | `docs/principles/P004.md` |
| P005 Security | `internal/principles/p005_security.go` | `docs/principles/P005.md` |
| P006 Functional | `internal/testing/functional.go` | `docs/principles/P006.md` |
| P007 Performance | `internal/testing/performance.go` | `docs/principles/P007.md` |
| P008 Versioning | `internal/principles/p008_versioning.go` | `docs/principles/P008.md` |
| Spec Abstraction | `internal/spec/spec.go`, `openapi3.go`, `swagger2.go` | `docs/architecture.md` |
| Engine Design | `internal/engine/engine.go` | `docs/architecture.md` |
| CLI Design | `internal/cli/root.go`, `validate.go`, `config.go` | `docs/CLI_USAGE.md` |
| Validation Modes | `internal/types/modes.go`, `internal/principles/registry.go` | `docs/MINIMAL_MODE_GUIDE.md` |
| XSDLC Architecture | `kubernetes/helm/driveby/templates/crossplane/` | `docs/quality-gate-sdlc.md`, `docs/deployment-guide.md` |
| GitOps Pipeline | `kubernetes/examples/novelcore-perfect-api/` | `docs/gitops-pipeline.md` |
| Evaluation Data | `tools/`, `apis/perfect-api/` | `docs/evaluation-methodology.md` |
| CI/CD Integration | `.github/workflows/ci.yml` | `docs/WORKFLOW.md` |

## LaTeX Conventions
- Use `\section{}`, `\subsection{}`, `\subsubsection{}` — no deeper nesting
- Citations: `\cite{key}` with entries in `references.bib`
- Cross-references: `\label{sec:name}` and `\ref{sec:name}`
- Code listings: `\begin{lstlisting}[language=Go]` for source code
- Figures: placed in `thesis/figures/`, referenced with `\includegraphics[width=\textwidth]{figures/name.pdf}`
- Tables: use `\begin{table}[htbp]` with `\caption{}` and `\label{tab:name}`
- Acronyms: define on first use, e.g., "Documentation-Driven Testing (DDT)"

## Citation Rules
- Every factual claim must have a `\cite{}` reference
- Self-citations to the DriveBy repository are acceptable for implementation details
- Use BibTeX entry types: `@article`, `@inproceedings`, `@misc` (for URLs/tools), `@book`
- All references go in `references.bib` — no inline URLs without BibTeX entries

## Figure Requirements
- [DONE] Architecture diagram: dependency flow (Ch.4) — `figures/dependency-flow.tex`, `fig:dependency-flow`
- [DONE] Axiom-to-principle mapping (Ch.3) — `figures/axiom-principle-mapping.tex`, `fig:axiom-mapping`
- [DONE] Pipeline diagram: GitOps event-driven workflow (Ch.5) — `figures/workflow-pipeline.tex`, `fig:gitops-pipeline`
- [DONE] Staging promotion DAG (Ch.5) — `figures/staging-promotion-dag.tex`, `fig:staging-dag`
- [DONE] Evaluation framework (Ch.6) — `figures/evaluation-methodology.tex`, `fig:eval-methodology`
- [REMOVED] Provider architecture (Ch.5) — `figures/provider-architecture.tex` — figure cut, text kept inline
- [DONE] Crossplane reconciliation loop (Ch.5) — `figures/crossplane-reconciliation-loop.tex`, `fig:crossplane-reconciliation-loop`
- [DONE] Declarative lifecycle (Ch.5) — `figures/declarative-lifecycle.tex`, `fig:declarative-lifecycle`
- [DONE] Single-repo branch flow (Ch.6) — `figures/single-repo-branch-flow.tex`, `fig:single-repo-branch-flow`
- [DONE] Human-in-the-loop (Ch.6) — `figures/human-in-the-loop.tex`, `fig:human-in-the-loop`
- [DONE] Quality control loop (Ch.6) — `figures/quality-control-loop.tex`, `fig:quality-control-loop`
- Bar chart: per-principle pass rates across APIs.guru dataset (Ch.6) — TODO (requires evaluation data)
- Table: perfect-api validation results (Ch.6) — TODO (requires evaluation data)
- All figures are TikZ vector diagrams (inline LaTeX, no external image files needed)

## Writing Rules
- Academic tone, third person, present tense for methodology
- Past tense for evaluation results ("The evaluation showed...")
- Cite all claims via `references.bib`
- Map code to chapters: P001-P009 -> Chapter 3, architecture -> Chapter 4
- Pull technical details from `driveby-cli/` source for accuracy
- Reference PRD docs in `docs/` for vision alignment
- Figures go in `thesis/figures/`, referenced with `\includegraphics`
- When describing a principle, include: formal definition, implementation details (from source), example check output
