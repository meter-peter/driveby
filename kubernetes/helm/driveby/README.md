# DriveBy Helm Chart

## Installation

```bash
helm install driveby . -n driveby --create-namespace \
  --set github.owner=your-org \
  --set github.repo=your-repo
```

## Configuration

See `values.yaml` for all configurable options.

## Components

- **Argo WorkflowTemplates**: Validation, functional testing, performance testing, full pipeline
- **Argo Events**: GitHub webhook EventSource + Sensor
- **ArgoCD**: Isolated AppProject (uses existing ArgoCD installation)
- **RBAC**: ServiceAccount, Role, RoleBinding for DriveBy workflows
- **Crossplane** (optional): XRD + Composition for QualityGate resources
- **GitOps Promoter** (optional): Promotion strategies with commit status gates
