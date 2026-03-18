# DriveBy Kubernetes Manifests

## Prerequisites
- Kubernetes cluster (target: private.novelcore.org)
- ArgoCD (already installed)
- Argo Workflows
- Argo Events
- Crossplane with provider-kubernetes

## Structure
- `helm/driveby/` — Helm chart for full installation
- `examples/` — Raw YAML examples for reference

## Quick Start
```bash
helm install driveby helm/driveby/ -n driveby --create-namespace
```
