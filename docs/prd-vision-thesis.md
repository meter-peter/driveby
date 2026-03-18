Product Requirements Document: Pure Crossplane DriveBy Platform
1. Product Vision and DDT Methodology
The Pure Crossplane DriveBy Platform implements Documentation-Driven Testing (DDT), a methodology that transforms OpenAPI specifications into autonomous quality gates. By treating API documentation as executable infrastructure, the platform addresses the "GitOps Quality Assurance Gap" where manual testing fails to scale with declarative delivery.
1.1 The DDT Thesis
The platform operationalizes the "Documentation as Contracts" paradigm through three foundational axioms:
Completeness Axiom: A fully documented API contains all necessary data to verify correctness.
Determinism Axiom: Documented examples imply reproducible behavior.
Observability Axiom: Machine-readable descriptions enable automated analysis and enforcement.
1.2 Core Problem Statement
Discrepancies between API documentation and implementation account for 62% of software outages. Traditional QA requires manual synchronization, leading to documentation drift and high maintenance overhead. The DriveBy platform eliminates manual test authorship by using the OpenAPI spec as the single source of truth.
1.3 Key Performance Indicators (KPIs)
Implementation of the DDT methodology targets the following operational outcomes:
89% Reduction in manual QA effort through automated test generation.
40% Fewer production incidents traced to API documentation-implementation drift.
Zero-Overhead Maintenance: Quality gates update automatically as the OpenAPI spec evolves.
2. Platform Architecture Shift: From Kubebuilder to Pure Crossplane
The platform has transitioned from legacy imperative Go-based operators to a "Pure Crossplane" architecture. This model uses CompositeResourceDefinitions (XRDs) and Compositions to manage the SDLC through a tool-centric, declarative approach.
2.1 Core Composite Entities
XSDLC (Orchestrator): Manages the environment lifecycle and promotion logic. It leverages PromotionStrategy and ChangeTransferPolicy to coordinate state transitions across Dev, Stage, and Prod.
XQualityGate (Execution Engine): Encapsulates the DriveBy engine to enforce quality principles. It translates OpenAPI requirements into functional and performance tests within the cluster.
2.2 Technical "Wiring" and Logic
The relationship between these entities is managed via Crossplane EnvironmentConfigs and metadata-based selection:
Context Injection: KubeOrg (the top-level XRD) creates ProviderConfigs and org-wide EnvironmentConfigs (containing Org Name, AWS Account IDs, etc.).
Resource Selection: XQualityGate claims select these configs using labels (e.g., kubecore.io/kubeorg=<name>).
Ownership: Child Crossplane resources reference parent KubeCore resources via standard Kubernetes ownerReferences, ensuring cascading cleanup and state consistency.
2.3 Architectural Model Comparison
Feature
Legacy Kubebuilder Operator
Pure Crossplane (XRD-Driven)
Logic Type
Imperative (Go code/Controllers)
Declarative (YAML Compositions)
Sync Method
Manual reconciliation logic
Tool-centric convergence (Argo CD + Crossplane)
Extensibility
Recompilation/Release required
Pluggable Crossplane Functions
Infrastructure
Hardcoded controller logic
Managed via Crossplane Providers (AWS, GitHub)
3. Technical Requirements: The DriveBy CLI Engine
The DriveBy CLI is the execution core. It mandates explicit flag-based configuration to ensure portability and reproducible behavior across local, CI, and cluster environments.
3.1 Validation Principles (P001–P008)
The engine must enforce the following principles, executing specific checks derived from the validation-report.md:
P001: OpenAPI Specification Compliance (Critical)
Checks: Validates OpenAPI 3.0.x/3.1.0 versioning, presence of required info fields (title, version), resolvable references, and absence of duplicate operationIds.
P002: API Documentation Quality (Warning)
Checks: Verifies all operations have summaries/descriptions, parameters/schemas have descriptions, and all request/response bodies contain examples.
P003: Error Handling Standards (Warning)
Checks: Documents 4xx and 5xx responses for all operations; ensures error responses include standard error codes, messages, and detail schemas.
P004: Request Schema Definitions (Warning)
Checks: Validates data types for all schemas, length constraints for string fields, and min/max values for numeric fields.
P005: Security Standards (Critical)
Checks: Ensures security schemes are defined in components.securitySchemes, global/operation-level security requirements are set, and OAuth2 scopes are documented.
P006: Endpoint Functional Testing
Checks: Extracts examples from components.examples and validates them against live endpoints via combinatorial sampling.
P007: API Performance Compliance
Checks: Validates response times against SLOs defined in x-performance OpenAPI extensions.
P008: API Versioning Strategy (Warning)
Checks: Enforces semantic versioning (SemVer), documentation of breaking changes, and presence of deprecation notices.
3.2 Validation Modes
Mode
Purpose
Principals Covered
Requirement: Execution Speed
test-only
Pure functional/load testing
None (Skips all validation)
< 60s
minimal
Essential Dev validation
P001, P004 (Basic)
< 20s
strict
Production readiness
All (P001-P008)
< 5m
3.3 CLI Interface and Requirements
The CLI must support the following interface. If required flags are missing or conflicting auth methods are provided, the CLI must exit with Exit Code 3 (Invalid Arguments).
Mandatory Flags:
--openapi: Path or URL to OpenAPI specification.
--host: Target API host.
Configuration Defaults:
--port: Default 8080 (http) / 443 (https).
--protocol: Default http.
--timeout: Default 30 (seconds).
--validation-mode: Default minimal.
--report-dir: Default /tmp/driveby-reports.
GitHub and Security Flags:
--github-app-id: (Required for App Auth)
--github-installation-id: (Required for App Auth)
--github-private-key: (File path or PEM content)
--github-app-slug: (Required for App Auth)
--github-token: (Legacy PAT - Prohibited in Prod)
4. GitOps Promotion and Lifecycle Management
Promotion is managed as a series of Git operations governed by the PromotionStrategy and ChangeTransferPolicy CRDs.
4.1 Tool-Only GitOps Lifecycle
Source (GitHub): Developer pushes code or OpenAPI spec changes to a feature branch.
Build/Push: CI builds images and updates the GitOps repository manifests.
Promotion (Git Ops): ChangeTransferPolicy detects a commit in a "proposed" branch and opens a PR against the "live" branch.
Sync (Argo CD): Argo CD detects the commit/PR and syncs manifests (XRDs/XRs) to the cluster.
Reconcile (Crossplane): Crossplane materializes infrastructure and triggers the XQualityGate validation.
Feedback: Status flows back via Kubernetes conditions and GitHub PR comments.
4.2 Gating Mechanisms
CommitStatus resources act as the primary truth for promotion. Key gates include:
ArgoCD Health: Monitors application sync state.
WebRequest: Validates external API endpoints via configurable expressions.
Timed/Soak Time: Enforces "bake time" policies (e.g., 1h soak in Stage before Prod promotion).
5. Integration and Feedback Loop Requirements
5.1 Mandated GitHub App Authentication (Security P0)
The platform mandates GitHub App Authentication for all production environments. Legacy Personal Access Tokens (PATs) are prohibited due to their user-wide permission scopes. GitHub Apps provide granular, repository-specific permissions and a robust audit trail.
5.2 Feedback Channels
GitHub PR Commenting: Automated reports posted to PRs via --github-comment.
Crossplane Conditions: Resource state reflected in .status.conditions (Ready, Synced).
Argo CD Status: Real-time visibility of sync status and health (Healthy/Degraded).
5.3 Reporting and SARIF Support
The engine must generate reports in JSON, Markdown, and SARIF (Static Analysis Results Interchange Format). SARIF is required for native integration with GitHub Advanced Security and modern CI/CD security dashboards. Reports must include:
Summary Statistics (Passed/Failed/Critical).
Failed Tags (e.g., schema, compliance).
Actionable Suggested Fixes (e.g., "Add missing length constraints to string fields").
6. Functional and Performance Standards
6.1 Performance Thresholds (Defaults)
--max-latency-p95: 500ms
--min-success-rate: 0.99 (99%)
--concurrent-users: 10
--test-duration: 300s
6.2 Authentication Protocols
The engine supports Bearer Token, API Key, and Basic Auth. Rule: Only one method may be active at a time. Violation must trigger a validation error: "Only one authentication method can be used at a time."
6.3 Operational Resource Management
Secrets: All tokens/keys must be stored in K8s Secrets and referenced via flags; hardcoding is prohibited.
Pod Limits: Validation pods must have defined CPU/Memory limits to ensure they do not exhaust cluster resources during load tests.
7. Operational Verification and Troubleshooting
7.1 Verification Commands
KubeCore State: kubectl get kubeorgs.schema.kubecore.io
Infrastructure State: kubectl get xr
Argo CD Sync: kubectl get applications.argoproj.io -n argocd
CI/CD Pipeline: gh run view <run-id>
Promotion PRs: gh pr list
7.2 Common Failure Modes
Issue
Potential Cause
Recommended Solution
Authentication Failure
GitHub App Private Key mismatch
Verify secret mapping in ProviderConfig.
Schema Validation Error
Missing length constraints on strings
Update OpenAPI spec to include minLength/maxLength.
OpenAPI Incompatibility
Spec version is not 3.0.x or 3.1.0
Convert spec to a supported version.
Load Test Failure
API rate limiting
Increase --timeout or reduce --concurrent-users.
Promotion Blocked
TimedCommitStatus soak time not met
Wait for required duration or verify commitTime.