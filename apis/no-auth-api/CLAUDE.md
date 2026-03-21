# No-Auth API — Deliberately Insecure FastAPI Variant

## Overview
A FastAPI application derived from `perfect-api` that deliberately omits all authentication and security definitions. This API is the **negative control** for DriveBy P005 (Security Standards) evaluation: it should pass all documentation/schema principles but fail security validation.

## How It Differs from Perfect API
- **Removed**: `APIKeyHeader`, `HTTPBearer`, `Security` imports from `fastapi.security`
- **Removed**: `api_key_header`, `bearer_scheme` declarations
- **Removed**: `verify_auth` dependency function
- **Removed**: `Depends(verify_auth)` / `_auth=Depends(verify_auth)` from all route handlers
- **Removed**: `securitySchemes` from custom OpenAPI schema
- **Removed**: Global `security` requirement from custom OpenAPI schema
- **Removed**: Per-operation security scheme remap loop
- **Removed**: All `401 Unauthorized` entries from route `responses={}` dicts
- **Kept**: All documentation (summaries, descriptions, examples, response_description)
- **Kept**: All Pydantic model Field() descriptions and examples
- **Kept**: All business logic, error handling, and other response codes
- **Kept**: ErrorResponse schema injection and versioning strategy text in custom_openapi()
- **Kept**: `patch_exclusive_min_max` for OpenAPI 3.0 compatibility

## Endpoints
Same as perfect-api:

| Method | Path | Description |
|--------|------|-------------|
| POST | `/tasks` | Create a new task |
| GET | `/products` | List all products |
| POST | `/products` | Create a new product |
| GET | `/products/{id}` | Get product by ID |
| PUT | `/products/{id}` | Update product by ID |
| DELETE | `/products/{id}` | Delete product by ID |
| GET | `/test/health` | Health check endpoint |
| POST | `/test/echo` | Echo endpoint for testing |
| GET | `/legacy/products` | Legacy product listing (deprecated) |

## Expected DriveBy Principle Results

| Principle | Expected | Notes |
|-----------|----------|-------|
| P001 OpenAPI Compliance | PASS | Valid OpenAPI 3.0.3 spec |
| P002 Documentation Quality | PASS | All operations have summaries, descriptions, examples |
| P003 Error Handling | PASS | 4xx/5xx responses defined on all endpoints |
| P004 Schema Definitions | PASS | Full request/response schemas with types |
| P005 Security Standards | FAIL | No securitySchemes, no global security, no per-operation security |
| P008 Versioning Strategy | PASS | Version info in spec metadata |
| P009 Test Readiness | PASS | Should pass (depends on P001-P004 + P008) |

## Thesis Mapping
- **Chapter 6 (Evaluation)**: Negative control for P005 validation
- Demonstrates that DriveBy correctly identifies missing security definitions
- Paired with perfect-api (positive control) to show principle isolation

## Running Locally
```bash
docker build -t no-auth-api .
docker run -p 8000:8000 no-auth-api
# OpenAPI spec at http://localhost:8000/openapi.json
```
