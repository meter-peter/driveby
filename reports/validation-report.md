# API Validation Report

Generated: 2025-05-31T17:24:37+03:00

## Summary

- Total Checks: 8
- Passed Checks: 2
- Failed Checks: 6
- Critical Issues: 5
- Warnings: 1
- Info: 0

## Principle Results

### P001: OpenAPI Specification Compliance (Failed)

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: invalid components: schema "Task": unsupported 'type' value "null"

### P002: Response Time Performance (Failed)

- **Status:** Failed
- **Message:** Performance targets are not configured.

### P003: Error Response Documentation (Failed)

- **Status:** Failed
- **Message:** Endpoints missing error response documentation

#### Details

```json
[
    "GET /test/health",
    "GET /health"
  ]
```


### P004: Request Validation (Passed)

- **Status:** Passed
- **Message:** 

### P005: Authentication Requirements (Failed)

- **Status:** Failed
- **Message:** No security schemes defined

### P006: Endpoint Functional Testing (Failed)

- **Status:** Failed
- **Message:** Some endpoints failed functional tests (4/9 failed): 4 client errors (4xx)

### P007: API Performance Compliance (Failed)

- **Status:** Failed
- **Message:** Unknown principle ID: P007

### P008: API Versioning (Passed)

- **Status:** Passed
- **Message:** API version is specified as: 1.0.0

