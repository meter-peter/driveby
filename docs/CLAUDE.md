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
| `prd-01/` | Initial PRD iteration (historical) |

## Existing Documentation

The following files are complete and available:
- `principles/P001.md` through `P008.md` — per-principle documentation (all 8 principles documented)
- `ddt-axioms.md` — DDT methodology explanation (3 axioms: Completeness, Determinism, Observability)
- `architecture.md` — System design and dependency flow
- `gitops-pipeline.md` — Complete GitOps pipeline documentation (webhook → Argo Events → Workflow → PR feedback)
- `evaluation-methodology.md` — Chapter 6 evaluation protocol and metrics

## Thesis Feed Map
When writing thesis chapters, consult these docs:

| Thesis Chapter | Primary Docs |
|---------------|-------------|
| Ch.1 Introduction | `prd-vision-thesis.md`, `Thesis.md` |
| Ch.2 Related Work | `prd-vision-thesis.md` (positioning) |
| Ch.3 Methodology | `principles/P001-P008.md`, `ddt-axioms.md` |
| Ch.4 Architecture | `architecture.md`, `CLI_USAGE.md` |
| Ch.5 Workflow | `WORKFLOW.md`, `gitops-pipeline.md`, `MINIMAL_MODE_GUIDE.md` |
| Ch.6 Evaluation | `evaluation-methodology.md` |
| Ch.7 Conclusion | `Thesis.md` |

## Rules
1. **Consult docs/ first** before writing thesis content -- ensure alignment with documented decisions
2. **When modifying source code**, update the corresponding docs/ file if the change affects documented behavior
3. **When adding a new principle**, create `docs/principles/PXXX.md` alongside the source implementation
4. Keep documentation in sync with the codebase -- stale docs are worse than no docs
