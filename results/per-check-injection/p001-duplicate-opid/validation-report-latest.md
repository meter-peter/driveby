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
- Specification


### Failed Tags
- openapi
- specification
- compliance


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid paths: operations "GET /products" and "POST /tasks" have the same operation id "create_task_tasks_post"; No duplicate operationIds: Duplicate operationIds found: create_task_tasks_post used in: POST /tasks, GET /products
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
      "No duplicate operationIds": false,
      "Paths are properly defined": true,
      "Required info fields (title, version) are present": true,
      "Specification structure is valid": false,
      "Specification version is present": true,
      "Valid HTTP methods used": true
    },
    "messages": {
      "No duplicate operationIds": "Duplicate operationIds found: create_task_tasks_post used in: POST /tasks, GET /products",
      "Specification structure is valid": "invalid paths: operations \"GET /products\" and \"POST /tasks\" have the same operation id \"create_task_tasks_post\""
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Passed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Passed
- **Message:** API documentation is comprehensive and high quality
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
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {}
  }
```

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

