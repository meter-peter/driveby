# Bad Docs API — Deliberately Undocumented FastAPI Implementation

## Overview
A FastAPI application derived from `perfect-api/perfect-api.py` that deliberately strips all documentation metadata while preserving functional correctness, security, schemas, and versioning. Designed as a negative control for DriveBy P002 (Documentation Quality) and P003 (Error Handling) validation.

## What Was Removed (vs perfect-api)
- `description`, `contact`, `license_info` from `FastAPI()` constructor
- `summary`, `description`, `response_description` from ALL route decorators
- `responses={}` error response dicts (400/401/422/500) from ALL routes
- `description` and `example` kwargs from ALL Pydantic `Field()` calls
- `ErrorResponse` model class and `CommonErrorResponses` helper class (not needed without error responses on routes)

## What Was Kept
- All endpoint business logic (identical to perfect-api)
- Security schemes: `ApiKeyAuth` (X-API-Key header) + `BearerAuth` (JWT Bearer)
- Global security requirement on all authenticated endpoints
- `verify_auth` dependency with API key and Bearer token support
- Pydantic validation constraints (`min_length`, `max_length`, `gt`, `ge`, `pattern`, etc.)
- Auto-generated Pydantic schemas in OpenAPI components (P004)
- Versioning strategy text injected into `info.description` (P008)
- `ErrorResponse` schema manually added to `components.schemas` in `custom_openapi()`
- `patch_exclusive_min_max` function for OpenAPI 3.0 compatibility
- Startup event with example products and tasks
- SemVer version `1.0.0`, OpenAPI version `3.0.3`

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/tasks` | Create a new task |
| GET | `/products` | List all products |
| POST | `/products` | Create a new product |
| GET | `/products/{product_id}` | Get product by ID |
| PUT | `/products/{product_id}` | Update product by ID |
| DELETE | `/products/{product_id}` | Delete product by ID |
| GET | `/legacy/products` | Legacy product listing (deprecated) |
| GET | `/test/health` | Health check (no auth) |
| POST | `/test/echo` | Echo endpoint for testing |

## Expected Principle Results

| Principle | Expected | Why |
|-----------|----------|-----|
| P001 OpenAPI Compliance | PASS | Valid OpenAPI 3.0.3 spec, well-formed |
| P002 Documentation Quality | FAIL | No summaries, descriptions, or examples on operations or fields |
| P003 Error Handling | FAIL | No error responses (4xx/5xx) defined on any operation |
| P004 Schema Definitions | PASS | Pydantic auto-generates typed schemas with validation constraints |
| P005 Security Standards | PASS | Global security with ApiKeyAuth + BearerAuth |
| P006 Functional Testing | N/A | Requires live API |
| P007 Performance Testing | N/A | Requires live API |
| P008 Versioning Strategy | PASS | Versioning text in info.description, SemVer version |
| P009 Test Readiness | FAIL | Depends on P002+P003 passing |

## Files

| File | Purpose |
|------|---------|
| `bad_docs_api.py` | FastAPI application source |
| `openapi.json` | Placeholder (API generates spec at runtime via `/openapi.json`) |
| `Dockerfile` | Container build — Python 3.11-slim with FastAPI stack |
| `requirements.txt` | Python dependencies (pinned versions) |

## Thesis Mapping
- **Chapter 6 (Evaluation)**: Negative control for documentation quality validation
- Demonstrates that DriveBy correctly identifies missing documentation and error handling
- Paired with `perfect-api/` (positive control) for comparative evaluation
