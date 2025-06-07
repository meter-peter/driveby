# API Validation Report

Generated: 2025-06-07T16:50:07+03:00

## Summary

- Total Checks: 8
- Passed Checks: 3
- Failed Checks: 5
- Critical Issues: 4
- Warnings: 1
- Info: 0

## Principle Results

### P001: OpenAPI Specification Compliance (Passed)

- **Status:** Passed
- **Message:** 

### P002: Response Time Performance (Failed)

- **Status:** Failed
- **Message:** Performance targets are not configured.

### P003: Error Response Documentation (Failed)

- **Status:** Failed
- **Message:** Endpoints missing error response documentation

#### Details

```json
[
    "GET /admin/initialized"
  ]
```


### P004: Request Validation (Passed)

- **Status:** Passed
- **Message:** 

### P005: Authentication Requirements (Failed)

- **Status:** Failed
- **Message:** Endpoints missing authentication requirements

#### Details

```json
[
    "POST /admin/register",
    "POST /auth/login",
    "POST /auth/logout",
    "GET /admin/initialized",
    "POST /auth/register"
  ]
```


### P006: Endpoint Functional Testing (Failed)

- **Status:** Failed
- **Message:** Some endpoints failed functional tests (48/50 failed): 46 authentication failures (401/403), 2 client errors (4xx)

### P007: API Performance Compliance (Failed)

- **Status:** Failed
- **Message:** Performance test skipped due to widespread functional test failures (likely auth/network issues).

### P008: API Versioning (Passed)

- **Status:** Passed
- **Message:** API version is specified as: 1.0.0

