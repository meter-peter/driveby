# Perfect API — FastAPI Reference Implementation

> **Moved to separate repo**: [`novelcore/perfect-api`](https://github.com/novelcore/perfect-api)

## Overview
A FastAPI (Python) application designed as the "known-good" reference implementation for DriveBy validation. This API is intentionally well-documented and well-structured to serve as the positive control in thesis evaluation.

The source code and Dockerfile now live in their own repo. This directory retains a copy of the files for local development and `make validate`.

## Deployment

| Environment | URL | Managed By |
|-------------|-----|------------|
| Staging | `https://perfect-api-staging.private.novelcore.org` | ArgoCD (autoSync) |
| Prod | `https://perfect-api-prod.private.novelcore.org` | ArgoCD (manual sync) |

- **Container image**: `ghcr.io/novelcore/perfect-api:latest` (built by GitHub Actions on push to main)
- **GitOps manifests**: `manifests/` directory in `novelcore/perfect-api` (single-repo model — source and manifests co-located)
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
