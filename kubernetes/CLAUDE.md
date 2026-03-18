# Kubernetes Directory — Deployment and GitOps Infrastructure

## Structure

```
kubernetes/
  helm/
    driveby/          # Helm chart for DriveBy deployment
      Chart.yaml
      values.yaml
      templates/
  examples/
    argo-workflows/   # Argo Workflow templates for validation pipelines
    argo-events/      # Argo Events triggers (webhook, git sensor)
    crossplane/       # Crossplane compositions for infrastructure provisioning
    gitops-promoter/  # GitOps promotion logic (environment promotion on validation pass)
  README.md
```

## Helm Chart (`helm/driveby/`)
Deploys the DriveBy CLI as a Kubernetes Job or CronJob for scheduled API validation. Configurable via `values.yaml`:
- Target API URL
- Validation mode (minimal/standard/strict)
- Report output destination
- GitHub integration credentials

## Example Resources (`examples/`)

| Directory | Purpose | Thesis Mapping |
|-----------|---------|---------------|
| `argo-workflows/` | Validation pipeline as Argo Workflow (spec fetch -> validate -> report -> promote) | Ch.5 end-to-end workflow |
| `argo-events/` | Event triggers: webhook on PR, git sensor on spec changes | Ch.5 event-driven architecture |
| `crossplane/` | Infrastructure-as-code compositions for API environments | Ch.5 Crossplane integration |
| `gitops-promoter/` | Promotion logic: advance environment on validation pass | Ch.5 feedback loop |

## Target Cluster
- **Cluster**: `private.novelcore.org`
- **ArgoCD**: Already installed -- do NOT reinstall
- **Strategy**: Use a separate namespace + AppProject for DriveBy resources
- Reference kubecore-operator patterns at `/home/meter-peter/development/novelcore/kubecore-operator/compositions/`

## Crossplane Workflow
When writing Crossplane resources:
1. Use `context7 resolve-library-id` first for provider docs
2. Use `context7 query-docs` for XRD schemas, Composition patterns
3. Only fall back to web search if context7 lacks the information

## Thesis Mapping
- **Chapter 5 (End-to-End Workflow)**: The full GitOps pipeline from event trigger through validation to environment promotion
- The Argo Workflows + Events + Crossplane + Promoter combination demonstrates DDT in a real GitOps context
