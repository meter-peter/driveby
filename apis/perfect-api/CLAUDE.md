# Perfect API — FastAPI Reference Implementation

> **Moved to separate repo**: [`novelcore/perfect-api`](https://github.com/novelcore/perfect-api)
>
> **Naming convention (thesis vs source):** the thesis text refers to this API as **`non-critical-api`** (the rename rationale is in thesis Chapter 7 §7.2). Source identifiers in this directory, deployment manifests, and container images keep the historical name `perfect-api` for stability across past results, deployments, and clusters. The rename will be propagated to source in the v4.0 release. The artefact is the same.

## Overview
A FastAPI (Python) application designed as the "known-good" reference implementation for DriveBy validation. This API is intentionally well-documented and well-structured to serve as the positive control in thesis evaluation.

The source code and Dockerfile now live in their own repo. This directory retains a copy of the files for local development and `make validate`.

## Deployment

| Environment | URL | Managed By |
|-------------|-----|------------|
| Staging | `https://perfect-api-staging.private.novelcore.org` | ArgoCD (autoSync) |
| Prod | `https://perfect-api-prod.private.novelcore.org` | ArgoCD (autoSync, autoMerge: false) |

- **Container image**: `ghcr.io/novelcore/perfect-api:latest` (built by GitHub Actions on push to main)
- **GitOps manifests**: `novelcore/perfect-api-gitops` (two-repo model — dedicated gitops repo, auto-created by XSDLC)
- **ArgoCD AppProject**: `perfect-api` in `argocd` namespace
- **Database**: PostgreSQL 16 StatefulSet per environment (currently unused — API uses in-memory storage)

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/tasks` | Create a new task |
| GET | `/products` | List all products |
| POST | `/products` | Create a new product |
| GET | `/products/{id}` | Get product by ID |
| PUT | `/products/{id}` | Update product by ID |
| DELETE | `/products/{id}` | Delete product by ID |
| GET | `/test/health` | Health check endpoint (no auth) |
| POST | `/test/echo` | Echo endpoint for testing |

## Security
- **ApiKeyAuth**: API key via `X-API-Key` header
- **BearerAuth**: JWT Bearer token
- Global security applied to all endpoints (except `/test/health`)

## Expected Principle Results

| Principle | Expected | Notes |
|-----------|----------|-------|
| P001 OpenAPI Compliance | PASS | Valid OpenAPI 3.0.3 spec |
| P002 Documentation Quality | PASS | All operations have summaries, descriptions, examples |
| P003 Error Handling | PASS | 4xx/5xx responses defined on all endpoints |
| P004 Schema Definitions | PASS | Full request/response schemas with types |
| P005 Security Standards | PASS | Global security with ApiKeyAuth + BearerAuth |
| P006 Functional Testing | PASS | Endpoints respond correctly (requires live API) |
| P007 Performance Testing | PASS | Expected to meet latency thresholds |
| P008 Versioning Strategy | PASS | Version info in spec metadata |

## Thesis Mapping
- **Chapter 6 (Evaluation)**: Primary controlled evaluation target
- Used with `make validate` from the monorepo root via docker-compose
- DriveBy workflows in `driveby` namespace validate against `perfect-api.perfect-api-staging:8000`
