# Docs Directory — Documentation Hub

## Purpose
Central documentation for the DriveBy project. Agents should consult this directory first before writing thesis content or modifying source code.

## Structure

| File/Directory | Purpose |
|---------------|---------|
| `prd-vision-thesis.md` | Product Requirements Document — project vision, goals, thesis alignment |
| `Thesis.md` | Thesis overview — chapter summaries, writing timeline |
| `CLI_USAGE.md` | CLI reference — all commands, flags, examples |
| `MINIMAL_MODE_GUIDE.md` | Guide to minimal validation mode — when and why to use it |
| `WORKFLOW.md` | End-to-end workflow documentation — local, remote URL, batch, and K8s deployment |
| `gitops-pipeline.md` | GitOps pipeline — PR webhook → Argo Events → Workflow → validation → PR feedback |
| `responsibility-matrix.md` | Responsibility matrix — what DriveBy owns vs. what clients own |
| `deployment-guide.md` | Full deployment guide — Helm install, secrets, XRDs, verification checklist |
| `quality-gate-sdlc.md` | Quality gate SDLC architecture — Crossplane XSDLC, flow diagrams, DDT mapping |
| `prd-01/` | Initial PRD iteration (historical) |

## Existing Documentation

The following files are complete and available:
- `principles/P001.md` through `P009.md` — per-principle documentation (all 9 principles documented)
- `ddt-axioms.md` — DDT methodology explanation (3 axioms: Completeness, Determinism, Observability)
- `architecture.md` — System design and dependency flow
- `gitops-pipeline.md` — Complete GitOps pipeline documentation (webhook → Argo Events → Workflow → PR feedback)
- `evaluation-methodology.md` — Chapter 6 evaluation protocol and metrics
- `deployment-guide.md` — Complete Helm chart installation and Crossplane quality gate setup
- `quality-gate-sdlc.md` — Quality gate SDLC architecture with Crossplane XRDs

## Thesis Feed Map
When writing thesis chapters, consult these docs:

| Thesis Chapter | Primary Docs |
|---------------|-------------|
| Ch.1 Introduction | `prd-vision-thesis.md`, `Thesis.md` |
| Ch.2 Related Work | `prd-vision-thesis.md` (positioning) |
| Ch.3 Methodology | `principles/P001-P009.md`, `ddt-axioms.md` |
| Ch.4 CLI Architecture | `architecture.md`, `CLI_USAGE.md` |
| Ch.5 Kubernetes Architecture | `deployment-guide.md`, `quality-gate-sdlc.md`, `responsibility-matrix.md` |
| Ch.6 GitOps Pipeline | `gitops-pipeline.md`, `evaluation-methodology.md` |
| Ch.7 Conclusion | `Thesis.md` |

## Diagram Coverage

| Doc File | Has Diagrams | Type |
|----------|-------------|------|
| `architecture.md` | Yes | 3 ASCII diagrams (component, data flow, K8s deployment) |
| `WORKFLOW.md` | Yes | 2 ASCII diagrams (simple flow + GitOps pipeline) |
| `gitops-pipeline.md` | Yes | 2 ASCII diagrams (pipeline + DAG) |
| `quality-gate-sdlc.md` | Yes | 1 ASCII diagram (gate flow) |
| `deployment-guide.md` | Yes | 3 ASCII diagrams (prerequisites, namespace layout, install order) |
| `ddt-axioms.md` | Yes | 1 ASCII diagram (axiom-to-principle mapping) |
| `evaluation-methodology.md` | Yes | 1 ASCII diagram (evaluation framework) |
| `CLI_USAGE.md` | No | Text-only reference (appropriate) |
| `MINIMAL_MODE_GUIDE.md` | No | Text-only guide (appropriate) |

## Rules
1. **Consult docs/ first** before writing thesis content -- ensure alignment with documented decisions
2. **When modifying source code**, update the corresponding docs/ file if the change affects documented behavior
3. **When adding a new principle**, create `docs/principles/PXXX.md` alongside the source implementation
4. Keep documentation in sync with the codebase -- stale docs are worse than no docs
