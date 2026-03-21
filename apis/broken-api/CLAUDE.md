# Broken API — Deliberately Flawed Implementation

## Overview
A FastAPI application derived from `perfect-api` that serves the **exact same OpenAPI spec** but deliberately violates it in three endpoints. This API is designed to pass all static DDT principles (P001-P005, P008) while failing P006 (Functional Testing), demonstrating that spec quality alone is insufficient without runtime validation.

The key mechanism: `custom_openapi()` loads `openapi.json` from disk (a copy of perfect-api's spec), so the served spec is perfect. But the implementation diverges in three places.

## Implementation Bugs

| Endpoint | Spec Says | Implementation Does | Failure Type |
|----------|-----------|---------------------|-------------|
| `GET /products` | 200 with `List[Product]` | 418 with `{"error": "I'm a teapot"}` | Wrong status code (undocumented 418) |
| `POST /products` | 201 with `Product` | 200 with `Product` | Wrong status code (200 vs 201) |
| `GET /products/{id}` | 200 with `Product` | 200 with `{"result": "not found"}` | Wrong response schema |

## Working Endpoints
All other endpoints behave correctly and match the spec:
- `POST /tasks` — creates tasks, returns 201
- `PUT /products/{id}` — updates products, returns 200
- `DELETE /products/{id}` — deletes products, returns 204
- `GET /test/health` — health check, returns 200
- `POST /test/echo` — echo, returns 200
- `GET /legacy/products` — legacy listing, returns 200

## Expected Principle Results

| Principle | Expected | Notes |
|-----------|----------|-------|
| P001 OpenAPI Compliance | PASS | Spec is loaded from perfect-api's openapi.json |
| P002 Documentation Quality | PASS | All operations have summaries, descriptions, examples |
| P003 Error Handling | PASS | 4xx/5xx responses defined on all endpoints |
| P004 Schema Definitions | PASS | Full request/response schemas with types |
| P005 Security Standards | PASS | Global security with ApiKeyAuth + BearerAuth |
| P006 Functional Testing | FAIL | Three endpoints return wrong status codes or schemas |
| P007 Performance Testing | N/A | Not yet implemented |
| P008 Versioning Strategy | PASS | Version info in spec metadata |
| P009 Test Readiness | PASS | Spec has examples and testable endpoints |

## Thesis Mapping
- **Chapter 6 (Evaluation)**: Negative control — demonstrates that a "perfect" spec with a buggy implementation is caught by P006
- Validates the DDT thesis: static analysis (P001-P005) is necessary but not sufficient; runtime validation (P006) catches spec-implementation drift

## Running Locally
```bash
docker build -t broken-api .
docker run -p 8000:8000 broken-api
```
