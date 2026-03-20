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
2.1 Core Composite Entity
XSDLC (Pipeline Orchestrator): A single Crossplane XRD (`driveby.io/v1alpha1`) that defines the full delivery pipeline — environments, quality gates, promoter resources, ArgoCD applications, GitHub workflows, and branch protection. One CR (~35 lines YAML) generates ~34 Kubernetes resources. The environment chain is fully dynamic (minimum 2 environments, no maximum). XSDLC follows a Bring-Your-Own-CI (BYOCI) model: it is a delivery pipeline, not a CI system.
2.2 Technical "Wiring" and Logic
The relationship between XSDLC and the generated resources is managed through Crossplane compositions. The XSDLC CR declares the desired pipeline state; the composition generates all required Kubernetes resources (RBAC, workflows, event routing, promotion strategy).
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
3.1 Validation Principles (P001–P009)
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
P009: Test Readiness (Warning)
Checks: Validates that the specification provides sufficient testable data — typed schemas with constraints, request/response examples, and documented parameters — to enable meaningful functional testing (P006).
3.2 Validation Modes
Mode
Purpose
Principals Covered
Requirement: Execution Speed
test-only
Pure functional/load testing
P006, P007 (Skips all static validation)
< 60s
minimal
Essential Dev validation
P001
< 20s
strict
Production readiness
P001-P005, P008
< 5m
test-ready
Pre-flight for functional testing
P001, P004, P009
< 30s
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
4.1 Bring-Your-Own-CI (BYOCI) GitOps Lifecycle
XSDLC is a delivery pipeline, not a CI system. The developer's own CI (GitHub Actions, GitLab CI, Jenkins, etc.) handles building, testing, and updating manifests. XSDLC owns only promotion and quality gates.
Source (GitHub): Developer pushes code to main. Their CI builds the container image and updates the image tag in manifests/deployment.yaml.
Deployment: The developer triggers the auto-generated driveby-deploy.yml workflow (manual dispatch) to deploy manifests from the chosen source branch to a specific environment/*-next branch.
Promotion (GitOps Promoter): ChangeTransferPolicy detects a commit in a "-next" branch and opens a PR against the corresponding environment branch.
Quality Gate: The PR triggers a webhook → Argo Events → Argo Workflow → DriveBy validation against the source environment. The workflow creates a CommitStatus CRD.
Sync (Argo CD): On successful promotion (auto-merge or manual approval), Argo CD detects the commit on the environment branch and syncs manifests to the cluster.
Feedback: Status flows back via GitHub commit status, PR comments, and CommitStatus CRDs.
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


DOCUMENTATION-DRIVEN TESTING (DDT) & THE ZERO-TOUCH GITOPS PLATFORM
The Comprehensive Architectural and Technical Master Document
This document serves as the definitive guide to your thesis and platform architecture: Documentation-Driven Testing (DDT): A Paradigm for Automated API Quality Assurance in the GitOps Era
. It details the theoretical foundations, the rejection of custom platform operators in favor of pure declarative GitOps, the mechanics of GitOps Promoter, and the rigorous API validation enforced by the DriveBy CLI framework.

--------------------------------------------------------------------------------
PART 1: The Problem and the DDT Paradigm
1.1 The GitOps Quality Assurance Gap
Contemporary software delivery relies heavily on Internal Developer Platforms (IDPs) and GitOps workflows, where infrastructure is managed declaratively through code
. However, while infrastructure deployment has been automated, API reliability remains a critical vulnerability
. Currently, 62% of system outages are traced directly to discrepancies between API documentation and the live implementation—a phenomenon known as "documentation drift"
. Traditional testing frameworks exacerbate this issue because they require manual test authorship, creating massive maintenance overhead and remaining disconnected from the declarative GitOps pipeline
.
1.2 Documentation-Driven Testing (DDT)
This thesis introduces Documentation-Driven Testing (DDT), a novel methodology that transforms OpenAPI specifications into autonomous, executable quality gates
. DDT unifies static specification linting, example-driven functional testing, and synthetic load monitoring into a single Kubernetes-native framework
.
The methodology rests on three core theoretical axioms:
The Completeness Axiom: A fully documented API contains all the necessary data to verify its correctness
.
The Determinism Axiom: Documented examples strictly imply reproducible, verifiable behavior
.
The Observability Axiom: Machine-readable descriptions (OpenAPI 3.0.x/3.1.0) enable automated, dynamic analysis
.
By implementing DDT within a GitOps pipeline, organizations can achieve an 89% reduction in QA effort and 40% fewer production incidents compared to traditional testing methods
.

--------------------------------------------------------------------------------
PART 2: The Tool-Centric GitOps Architecture
A core design principle of this platform is the rejection of custom, monolithic Kubernetes operators (like KubeOrg, KubeProject, or KubeApp)
. Instead, the architecture is entirely "tool-centric," relying on a combination of best-in-class open-source tools to achieve a zero-touch pipeline driven purely by Git operations
.
2.1 The Four Pillars of the Platform
GitHub (The Source of Truth): All desired state, including application manifests, Crossplane resources, and Argo CD configurations, lives entirely in Git repositories
. Promotion between environments occurs strictly via Git operations, such as merging branches from dev to stage to prod
.
Crossplane (The Infrastructure Engine): Instead of custom Go operators, Crossplane provisions cloud and GitHub resources via standard CompositeResourceDefinitions (XRDs) and Compositions
. Crossplane connects to external APIs using a ProviderConfig (e.g., provider-github) to autonomously create repositories, manage teams, and configure webhooks
.
Argo CD (Continuous Delivery): Argo CD acts as the continuous delivery controller, watching specific Git branches or folders and syncing those manifests directly into the Kubernetes clusters
. Argo CD guarantees that the live cluster state converges to what Git declares, regardless of whether the manifests are basic workloads or complex Crossplane XRs
.
GitOps Promoter (The Promotion Engine): Rather than using imperative scripts, GitOps Promoter automates the flow of code across environments by autonomously opening Pull Requests and gating them based on automated checks
.

--------------------------------------------------------------------------------
PART 3: Inside GitOps Promoter
GitOps Promoter is the engine that moves changes between environments (which are simply Argo CD applications) while enforcing strict quality checks
.
3.1 Core Promotion CRDs
PromotionStrategy: This is the user's primary interface. It configures the sequence of live environment branches (e.g., environment/dev, environment/test, environment/prod) and defines the absolute gates—known as proposedCommitStatuses—that must pass before a promotion can occur
.
ChangeTransferPolicy: Generated automatically by the PromotionStrategy, this CRD represents a pair of environments (e.g., the proposed dev-next branch and the live dev branch)
. When a new commit appears, it opens a Pull Request to merge the code to the next environment
.
PullRequest: A thin wrapper around the Source Control Management (SCM) Pull Request API, used by the ChangeTransferPolicy to manage the actual Git merge
.
CommitStatus: The primary source of truth for promotion gates. It represents a specific check (like driveby-validation) and its phase (pending, success, or failure). If a CommitStatus is marked as a failure, the ChangeTransferPolicy will strictly refuse to merge the PR
.
3.2 Authentication & Status Aggregation
ScmProvider & ClusterScmProvider: These CRDs allow GitOps Promoter to securely authenticate to GitHub (or GitLab, Bitbucket, etc.) via Kubernetes Secrets to manage PRs and statuses
. ClusterScmProvider is cluster-scoped, allowing any repository to reference a central set of credentials (like a GitHub App)
.
ArgoCDCommitStatus: This controller monitors the health of Argo CD Applications and aggregates them into a CommitStatus
. It ensures that code is not promoted unless the underlying Argo CD Application is completely healthy and synced
.
TimedCommitStatus & WebRequestCommitStatus: Advanced gating mechanisms that allow promotions to be blocked based on time (e.g., requiring a 1-hour "bake time" in an environment) or external HTTP webhooks
.
To prevent orphaned resources in GitHub, GitOps Promoter utilizes Kubernetes finalizers to enforce a strict deletion order: Pull Requests are closed in the SCM before the underlying GitRepository or ScmProvider secrets can be deleted from the cluster
.

--------------------------------------------------------------------------------
PART 4: The DriveBy CLI Framework
If GitOps Promoter builds the roads, DriveBy is the mandatory tollbooth. DriveBy is a modern API validation framework that implements the DDT methodology to eliminate documentation drift
.
4.1 CLI-First Design & Configuration
DriveBy is configured entirely through explicit command-line flags, actively deprecating legacy environment variables to ensure executions are scriptable, self-documenting, and portable
.
Core CLI Flags:
--openapi: The path or URL to the OpenAPI specification (Required)
.
--host: The hostname of the live API to test (Required)
.
--validation-mode: Defines the rigor of the test (test-only, minimal, or strict)
.
Authentication flags: Supports --auth-token (Bearer), --auth-api-key, and --auth-username/--auth-password (Basic Auth). DriveBy strictly validates that only one authentication method is used at a time
.
4.2 GitHub App Integration
To securely post validation feedback directly to developers, DriveBy integrates with GitHub. While legacy Personal Access Tokens (--github-token) are supported, GitHub App Authentication is highly recommended for granular, repository-specific security
. When executing in a workflow, DriveBy uses the --github-app-id, --github-installation-id, and --github-private-key flags (along with --github-owner, --github-repo, and --github-pr-number) to autonomously post comprehensive Markdown validation reports directly to the Pull Request
.
4.3 The Four Execution Modes
To balance the need for rigorous production checks against the need for rapid feedback, DriveBy offers four validation modes
:
Test-Only Mode (--validation-mode=test-only):
Behavior: Skips all static OpenAPI validation entirely to run pure functional (P006) and performance (P007) tests
.
Performance: Executes in ~30-60 seconds, utilizing minimal CPU and memory
. Ideal for pipelines where the spec is already trusted
.
Minimal Mode (Default, --validation-mode=minimal):
Behavior: Focuses on essential structural validation. It runs basic OpenAPI specification compliance (P001) but skips deep schema validation, functional, and performance testing
.
Performance: Fastest static execution at ~10-20 seconds with low CPU/memory usage, perfect for development environments
.
Strict Mode (--validation-mode=strict):
Behavior: The production quality gate. It enforces all static validation principles (P001-P005, P008), enforcing schema constraints, documentation completeness, and security standards
.
Performance: Executes in ~2-5 minutes, designed for production-readiness checks
.
Test-Ready Mode (--validation-mode=test-ready):
Behavior: Pre-flight check that validates the specification provides sufficient data for meaningful functional testing. Runs P001, P004, and P009 (Test Readiness)
.
Performance: Executes in ~10-30 seconds. Use as a gate before functional-test checks to catch under-documented specs early
.
4.4 The Validation Principles (P001 - P009)
DriveBy enforces nine validation principles, each mapped to a DDT axiom
:
P001 (Specification Compliance — Completeness): Ensures the spec strictly follows OpenAPI 3.0.x/3.1.0 standards, checking that paths are defined, HTTP methods are valid, and components are resolvable
.
P002 (Documentation Quality — Completeness): Mandates that all operations have clear summaries, and all request/response bodies have concrete examples
.
P003 (Error Handling Standards — Completeness): Requires explicit documentation of 4xx and 5xx error responses, enforcing consistent error details schemas
.
P004 (Request Schema Definitions — Completeness): Validates rigorous data constraints. String fields must have length constraints, numeric fields require min/max values, and required fields must be explicitly marked
.
P005 (Security Standards — Observability): Ensures global and operation-level security schemes (e.g., OAuth2, API Keys) are properly defined
.
P006 (Functional Testing — Determinism): Automatically extracts examples from the OpenAPI spec to verify that the live API behaves exactly as documented
.
P007 (Performance Compliance — Observability): Executes load tests against the live API, driven by CLI flags like --max-latency-p95 (default 500ms), --min-success-rate (default 0.99), and --concurrent-users
.
P008 (Versioning Strategy — Observability): Ensures the API declares semantic versions, deprecation notices, and breaking changes
.
P009 (Test Readiness — Determinism): Validates that the specification provides sufficient testable data — typed schemas, examples, and documented parameters — to enable meaningful functional testing (P006)
.
4.5 Execution and Exit Codes
DriveBy's deterministic design outputs strict exit codes for CI/CD interpretation
:
0: Success (All validation checks passed)
.
1: Tests ran but failed validation (e.g., API deviated from the spec)
.
2: Error executing tests (System/Network failure)
.
3: Invalid command line arguments
.

--------------------------------------------------------------------------------
PART 5: The End-to-End "Zero-Touch" Workflow (BYOCI Model)
When these tools are combined, they create an autonomous, closed-loop GitOps lifecycle where documentation dictates reality. XSDLC follows a strict Bring-Your-Own-CI (BYOCI) model: it is a delivery pipeline, not a CI system. The developer's own CI handles building and updating manifests. XSDLC owns only promotion, quality gates, and environment sync.

Here is the step-by-step flow of a code change:
The Commit: A developer pushes a code change to main. Their own CI pipeline (GitHub Actions, GitLab CI, Jenkins, etc.) builds the container image and updates the image tag in manifests/deployment.yaml on the main branch
.
Manual Deploy: The developer triggers the auto-generated driveby-deploy.yml workflow (workflow_dispatch), selecting the source branch, target environment, and image tag. The workflow copies the manifests directory to the target environment/*-next branch and stamps the image tag
.
Dev Promotion (No Gate): GitOps Promoter auto-PRs environment/dev-next → environment/dev. With no gate on dev, the PR is auto-merged. Argo CD detects the commit on environment/dev, marks the Application as OutOfSync, and syncs the deployment to the Dev cluster
.
The Promotion Trigger (GitOps Promoter): GitOps Promoter observes the new commit in dev-next and autonomously opens a Pull Request from environment/staging-next to environment/staging
.
The Quality Gate (DriveBy via Argo Workflows):
The creation of the PR triggers a GitHub webhook to the cluster-side EventSource
.
The Sensor filters the event and submits an Argo Workflow based on the gate's checks array. Each check becomes a sequential DAG step
.
For example, a staging gate with validate-only + functional-test runs: set-pending → health-check → check-0-validate-only → check-1-functional-test → report-success
.
DriveBy validates against the source environment (dev), not the target (staging)
.
Feedback & Enforcement:
If the developer forgot to document a 500 Internal Server Error schema, DriveBy fails (Exit Code 1)
.
The Argo Workflow posts a detailed Markdown validation report as a PR comment and sets the GitHub commit status to failure
.
The workflow creates a CommitStatus CRD with phase: failure
.
Because the CommitStatus failed, GitOps Promoter strictly blocks the merge to staging.
Resolution: The developer adds the missing error schema to the OpenAPI document and pushes the fix to main. They re-trigger the driveby-deploy workflow targeting staging, the deploy updates staging-next, Promoter updates the PR, DriveBy re-runs, exits with 0 (Success), the CommitStatus turns green, and GitOps Promoter autonomously merges the PR to staging
. Argo CD immediately syncs the staging cluster
.
Production (Manual Approval): The same flow repeats for prod, but with autoMerge: false — the CommitStatus turns green but a human must approve the merge
.
The environment chain is fully dynamic — developers can define any number of environments (minimum 2) with any combination of gates. A 4-environment setup (dev → qa → staging → prod) works identically, with each gated environment getting its own webhook, workflow, and branch protection
.
