# API Validation Report

Generated: 2026-05-03T21:01:21+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 5
- Failed Checks: 1
- Critical Issues: 1
- Warnings: 0
- Info: 0

### Categories
- Documentation


### Failed Tags
- documentation
- quality
- usability


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Passed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Passed
- **Message:** OpenAPI specification is fully compliant with 3.0/3.1 standards
- **Tags:** openapi, specification, compliance

**Checks Performed:**
- OpenAPI version is 3.0.x or 3.1.0
- Required info fields (title, version) are present
- Paths are properly defined
- Components are valid
- References are resolvable
- No duplicate operationIds
- Valid HTTP methods used

**Details:**
```json
{
    "checks": {
      "Components are valid": true,
      "No duplicate operationIds": true,
      "Paths are properly defined": true,
      "Required info fields (title, version) are present": true,
      "Specification structure is valid": true,
      "Specification version is present": true,
      "Valid HTTP methods used": true
    },
    "messages": {}
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /test/echo: application/json request body, POST /test/echo: 401 application/json response, POST /test/echo: 422 application/json response, POST /test/echo: 500 application/json response, POST /test/echo: 200 application/json response, GET /test/health: 500 application/json response, GET /test/health: 200 application/json response, GET /test/health: 400 application/json response, GET /test/health: 404 application/json response, GET /legacy/products: 401 application/json response, GET /legacy/products: 500 application/json response, GET /legacy/products: 200 application/json response, GET /products: 500 application/json response, GET /products: 200 application/json response, GET /products: 400 application/json response, GET /products: 401 application/json response, GET /products: 422 application/json response, POST /products: application/json request body, POST /products: 201 application/json response, POST /products: 401 application/json response, POST /products: 422 application/json response, POST /products: 500 application/json response, DELETE /products/{product_id}: 401 application/json response, DELETE /products/{product_id}: 404 application/json response, DELETE /products/{product_id}: 422 application/json response, DELETE /products/{product_id}: 500 application/json response, GET /products/{product_id}: 200 application/json response, GET /products/{product_id}: 401 application/json response, GET /products/{product_id}: 404 application/json response, GET /products/{product_id}: 422 application/json response, GET /products/{product_id}: 500 application/json response, PUT /products/{product_id}: application/json request body, PUT /products/{product_id}: 404 application/json response, PUT /products/{product_id}: 422 application/json response, PUT /products/{product_id}: 500 application/json response, PUT /products/{product_id}: 200 application/json response, PUT /products/{product_id}: 401 application/json response, POST /tasks: application/json request body, POST /tasks: 400 application/json response, POST /tasks: 401 application/json response, POST /tasks: 422 application/json response, POST /tasks: 500 application/json response, POST /tasks: 201 application/json response
- **Tags:** documentation, quality, usability

**Checks Performed:**
- All operations have clear summaries
- All operations have detailed descriptions
- All operations have unique operationIds
- All parameters have descriptions
- All request/response bodies have examples
- All schemas have descriptions
- All enums have descriptions
- API has a general description
- Contact information is provided
- License information is provided

**Details:**
```json
{
    "checks": {
      "API has a general description": true,
      "All request/response bodies have examples": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All request/response bodies have examples": [
        "POST /test/echo: application/json request body",
        "POST /test/echo: 401 application/json response",
        "POST /test/echo: 422 application/json response",
        "POST /test/echo: 500 application/json response",
        "POST /test/echo: 200 application/json response",
        "GET /test/health: 500 application/json response",
        "GET /test/health: 200 application/json response",
        "GET /test/health: 400 application/json response",
        "GET /test/health: 404 application/json response",
        "GET /legacy/products: 401 application/json response",
        "GET /legacy/products: 500 application/json response",
        "GET /legacy/products: 200 application/json response",
        "GET /products: 500 application/json response",
        "GET /products: 200 application/json response",
        "GET /products: 400 application/json response",
        "GET /products: 401 application/json response",
        "GET /products: 422 application/json response",
        "POST /products: application/json request body",
        "POST /products: 201 application/json response",
        "POST /products: 401 application/json response",
        "POST /products: 422 application/json response",
        "POST /products: 500 application/json response",
        "DELETE /products/{product_id}: 401 application/json response",
        "DELETE /products/{product_id}: 404 application/json response",
        "DELETE /products/{product_id}: 422 application/json response",
        "DELETE /products/{product_id}: 500 application/json response",
        "GET /products/{product_id}: 200 application/json response",
        "GET /products/{product_id}: 401 application/json response",
        "GET /products/{product_id}: 404 application/json response",
        "GET /products/{product_id}: 422 application/json response",
        "GET /products/{product_id}: 500 application/json response",
        "PUT /products/{product_id}: application/json request body",
        "PUT /products/{product_id}: 404 application/json response",
        "PUT /products/{product_id}: 422 application/json response",
        "PUT /products/{product_id}: 500 application/json response",
        "PUT /products/{product_id}: 200 application/json response",
        "PUT /products/{product_id}: 401 application/json response",
        "POST /tasks: application/json request body",
        "POST /tasks: 400 application/json response",
        "POST /tasks: 401 application/json response",
        "POST /tasks: 422 application/json response",
        "POST /tasks: 500 application/json response",
        "POST /tasks: 201 application/json response"
      ]
    }
  }
```

**Suggested Fix:**
Add missing documentation including descriptions, examples, and operation details

---

### Schema

#### P004: Request Schema Definitions (Passed) [critical]

Ensures all API requests have comprehensive schema definitions with proper data types, validation rules, and constraints

- **Status:** Passed
- **Message:** All requests have comprehensive schema definitions with proper validation rules
- **Tags:** schema, validation, request

**Checks Performed:**
- All path parameters have schemas
- All query parameters have schemas
- All header parameters have schemas
- All request bodies have content schemas
- All schemas specify data types
- All schemas have appropriate constraints
- All required fields are marked
- All enums have valid values
- All numeric fields have min/max values
- All string fields have length constraints

**Details:**
```json
{
    "checks": {
      "All enums have valid values": true,
      "All header parameters have schemas": true,
      "All numeric fields have min/max values": true,
      "All path parameters have schemas": true,
      "All query parameters have schemas": true,
      "All request bodies have content schemas": true,
      "All required fields are marked": true,
      "All schemas have appropriate constraints": true,
      "All schemas specify data types": true,
      "All string fields have length constraints": true
    },
    "messages": {},
    "missing_validation": {}
  }
```

---

### Error Handling

#### P003: Error Handling Standards (Passed) [critical]

Validates comprehensive error response documentation and consistent error handling patterns

- **Status:** Passed
- **Message:** Error handling is well-documented and follows consistent patterns
- **Tags:** errors, responses, standards

**Checks Performed:**
- All operations document 4xx error responses
- All operations document 5xx error responses
- Error responses include error codes
- Error responses include error messages
- Error responses include error details schema
- Common error responses are defined in components
- Error responses follow consistent format

**Details:**
```json
{
    "checks": {
      "All operations document 4xx error responses": true,
      "All operations document 5xx error responses": true,
      "Common error responses are defined in components": true,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": true
    },
    "messages": {},
    "missing_errors": {}
  }
```

---

### Security

#### P005: Security Standards (Passed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Passed
- **Message:** All security requirements are properly defined and consistent
- **Tags:** security, authentication, authorization

**Checks Performed:**
- Security schemes are defined
- Global security requirements are set
- Operation-level security is defined
- OAuth2 scopes are documented
- API keys are properly described
- Authentication headers are specified
- Security requirements are consistent

**Details:**
```json
{
    "checks": {
      "API keys are properly described": true,
      "Authentication headers are specified": true,
      "Global security requirements are set": true,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": true,
      "Security requirements are consistent": true,
      "Security schemes are defined": true
    },
    "messages": {}
  }
```

---

### Versioning

#### P008: API Versioning Strategy (Passed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Passed
- **Message:** API versioning strategy is properly documented
- **Tags:** versioning, compatibility, lifecycle

**Checks Performed:**
- API version is specified
- Version follows semantic versioning
- Versioning strategy is documented
- Deprecation notices are present
- Breaking changes are documented
- Version compatibility is specified
- Migration guides are referenced

**Details:**
```json
{
    "checks": {
      "API version is specified": true,
      "Breaking changes are documented": true,
      "Deprecation notices are present": true,
      "Migration guides are referenced": true,
      "Version compatibility is specified": true,
      "Version follows semantic versioning": true,
      "Versioning strategy is documented": true
    },
    "messages": {}
  }
```

---

