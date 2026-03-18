# Perfect API — FastAPI Reference Implementation

## Overview
A FastAPI (Python) application designed as the "known-good" reference implementation for DriveBy validation. This API is intentionally well-documented and well-structured to serve as the positive control in thesis evaluation.

## Files

| File | Purpose |
|------|---------|
| `perfect-api.py` | FastAPI application source |
| `openapi.json` | Exported OpenAPI 3.1.0 specification |
| `requirements.txt` | Python dependencies |
| `Dockerfile` | Container build definition |

## Endpoints

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

## Security
- **ApiKeyAuth**: API key via `X-API-Key` header
- **BearerAuth**: JWT Bearer token
- Global security applied to all endpoints

## Expected Principle Results

| Principle | Expected | Notes |
|-----------|----------|-------|
| P001 OpenAPI Compliance | PASS | Valid OpenAPI 3.1.0 spec |
| P002 Documentation Quality | PASS | All operations have summaries, descriptions, examples |
| P003 Error Handling | PASS | 4xx/5xx responses defined on all endpoints |
| P004 Schema Definitions | PASS | Full request/response schemas with types |
| P005 Security Standards | PASS | Global security with ApiKeyAuth + BearerAuth |
| P006 Functional Testing | PASS | Endpoints respond correctly (requires live API) |
| P007 Performance Testing | PASS | Expected to meet latency thresholds |
| P008 Versioning Strategy | PASS | Version info in spec metadata |

## Docker
```bash
docker build -t perfect-api .
docker run -p 8000:8000 perfect-api
```
The OpenAPI spec is served at `http://localhost:8000/openapi.json`.

## Thesis Mapping
- **Chapter 6 (Evaluation)**: Primary controlled evaluation target
- Used with `make validate` from the monorepo root via docker-compose
