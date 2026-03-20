# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.8.0] - 2026-03-20

### Added
- P009 Test Readiness principle: validates examples, typed schemas, and response schemas
- `test-ready` validation mode: fast pre-flight check (P001 + P004 + P009) for meaningful testing
- P004 intermediate branch for `test-ready` mode: schema existence + type checks (no constraints)
- `github-comment` CLI command: reads pre-saved report files and posts a combined rich PR comment (no re-running tests)
- Rich PR comments: validation results + functional test tables + performance metrics in a single comment
- Shared `reports` PVC in WorkflowTemplate DAGs — all DriveBy steps write to and read from the same volume
- Fallback P007 metrics extraction: `CreateValidationComment` extracts `PerformanceMetrics` from `PrincipleResult.Details` when `TestResults.Performance` is nil

### Changed
- Replaced unused `flexible` validation mode with `test-ready`
- Updated XRD enum to include `test-ready` as valid validation mode
- Helm chart version 0.8.0

### Removed
- `ValidationModeFlexible` constant (was declared but never used)

## [0.7.0] - 2026-03-19

### Added
- XSDLC single-CR model: one custom resource defines the full promotion pipeline
- Two gate types: `validation` (DriveBy validate+functional) and `loadtest` (k6)
- `loadTestConfig` fields: concurrentUsers, testDuration, maxLatencyP95, minSuccessRate
- Simplified UX: ~25 lines of YAML per API

### Changed
- Replaced two-XRD model (XQualityGateTemplate + XQualityGate) with single XSDLC CR
- Replaced three-tier configuration with two-tier (values.yaml + XRD spec)
- Removed EnvironmentConfig dependency
- Reduced Crossplane functions from 4 to 2 (go-templating, auto-ready)
- Helm chart version 0.7.0

### Removed
- XQualityGateTemplate XRD
- XQualityGate XRD
- EnvironmentConfig resources
- function-environment-configs
- function-sequencer

## [0.3.0] - 2026-03-19

### Added
- Crossplane XRDs: `XQualityGateTemplate` and `XQualityGate` (v1alpha1)
- Crossplane compositions for declarative quality gate management
- Helm chart with Crossplane providers, functions, XRDs, and compositions
- Three-tier configuration model (values.yaml, EnvironmentConfig, XRD spec)
- Release automation: GoReleaser, GitHub Actions release workflow
- OCI Helm chart publishing to `ghcr.io/meter-peter/charts/driveby`
- `CONTRIBUTING.md` with development setup and PR guidelines

### Changed
- Helm chart appVersion aligned to CLI version (0.3.0)
- Dockerfile Go version fixed to match go.mod (1.21)
- CI version injection uses `git describe` instead of raw commit SHA
- README principle names corrected (P002, P004, P005)

### Removed
- `kubernetes/manifests/` directory (superseded by Helm chart + Crossplane)
- Dead Helm Argo Workflow templates (gated behind non-existent `.Values.apps`)
- Stale Crossplane examples (superseded by current XRD/Composition architecture)

## [0.2.0] - 2026-02-15

### Added
- GitOps infrastructure: ArgoCD AppProject, staging/production Applications
- Argo Events webhook integration (EventBus, EventSource, Sensor)
- Argo Workflows validation pipeline (staging-promotion DAG)
- Traefik Ingress for webhook endpoint
- GitHub PR commenting with commit status updates
- GitHub App authentication support (recommended over PAT)

## [0.1.0] - 2026-01-20

### Added
- Core validation engine with 8 DDT principles (P001-P008)
- `spec.APISpec` abstraction for OpenAPI 3.x and Swagger 2.0
- CLI commands: `validate-only`, `function-only`, `load-only`, `test-only`
- Validation modes: minimal, strict, test-only, test-ready
- Remote spec validation by URL
- JSON and Markdown report generation
- Docker image with multi-stage build
- Batch validation tooling for public APIs
- Authentication support: Bearer token, API key, Basic auth
