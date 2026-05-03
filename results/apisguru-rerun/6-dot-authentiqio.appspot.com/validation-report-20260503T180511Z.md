# API Validation Report

Generated: 2026-05-03T21:05:11+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 1
- Failed Checks: 5
- Critical Issues: 4
- Warnings: 1
- Info: 0

### Categories
- Documentation
- Error Handling
- Schema
- Security
- Versioning


### Failed Tags
- documentation
- responses
- schema
- security
- authorization
- errors
- validation
- authentication
- compatibility
- standards
- request
- versioning
- lifecycle
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
- **Message:** Documentation quality issues found: All operations have clear summaries: POST /key, DELETE /key, GET /key/{PK}, HEAD /key/{PK}, POST /key/{PK}, PUT /key/{PK}, DELETE /key/{PK}, POST /login, POST /scope, HEAD /scope/{job}, POST /scope/{job}, PUT /scope/{job}, DELETE /scope/{job}, GET /scope/{job}; All request/response bodies have examples: POST /key: application/jwt request body, POST /key: 201 application/json response, POST /key: 409 application/json response, POST /key: default */* response, DELETE /key: default */* response, DELETE /key: 200 application/json response, DELETE /key: 401 application/json response, DELETE /key: 404 application/json response, DELETE /key: 409 application/json response, GET /key/{PK}: 200 application/json response, GET /key/{PK}: 404 application/json response, GET /key/{PK}: 410 application/json response, GET /key/{PK}: default */* response, HEAD /key/{PK}: 404 */* response, HEAD /key/{PK}: 410 */* response, HEAD /key/{PK}: default */* response, POST /key/{PK}: application/jwt request body, POST /key/{PK}: 404 application/json response, POST /key/{PK}: default */* response, POST /key/{PK}: 200 application/json response, PUT /key/{PK}: application/jwt request body, PUT /key/{PK}: 404 application/json response, PUT /key/{PK}: 409 application/json response, PUT /key/{PK}: default */* response, PUT /key/{PK}: 200 application/json response, DELETE /key/{PK}: default */* response, DELETE /key/{PK}: 200 application/json response, DELETE /key/{PK}: 401 application/json response, DELETE /key/{PK}: 404 application/json response, POST /login: application/jwt request body, POST /login: 200 application/json response, POST /login: 401 application/json response, POST /login: default */* response, POST /scope: application/jwt request body, POST /scope: 201 application/json response, POST /scope: 429 application/json response, POST /scope: default */* response, HEAD /scope/{job}: 404 application/json response, HEAD /scope/{job}: default */* response, POST /scope/{job}: 405 application/json response, POST /scope/{job}: default */* response, POST /scope/{job}: 202 application/json response, POST /scope/{job}: 401 application/json response, POST /scope/{job}: 404 application/json response, PUT /scope/{job}: 200 application/jwt response, PUT /scope/{job}: 404 application/jwt response, PUT /scope/{job}: 409 application/jwt response, PUT /scope/{job}: default */* response, DELETE /scope/{job}: 404 application/json response, DELETE /scope/{job}: default */* response, DELETE /scope/{job}: 200 application/json response, GET /scope/{job}: 200 application/json response, GET /scope/{job}: 200 application/jwt response, GET /scope/{job}: 404 application/json response, GET /scope/{job}: 404 application/jwt response, GET /scope/{job}: default */* response; All operations have unique operationIds: HEAD /key/{PK}; All schemas have descriptions: Error
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
      "All operations have clear summaries": false,
      "All operations have unique operationIds": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All operations have clear summaries": [
        "POST /key",
        "DELETE /key",
        "GET /key/{PK}",
        "HEAD /key/{PK}",
        "POST /key/{PK}",
        "PUT /key/{PK}",
        "DELETE /key/{PK}",
        "POST /login",
        "POST /scope",
        "HEAD /scope/{job}",
        "POST /scope/{job}",
        "PUT /scope/{job}",
        "DELETE /scope/{job}",
        "GET /scope/{job}"
      ],
      "All operations have unique operationIds": [
        "HEAD /key/{PK}"
      ],
      "All request/response bodies have examples": [
        "POST /key: application/jwt request body",
        "POST /key: 201 application/json response",
        "POST /key: 409 application/json response",
        "POST /key: default */* response",
        "DELETE /key: default */* response",
        "DELETE /key: 200 application/json response",
        "DELETE /key: 401 application/json response",
        "DELETE /key: 404 application/json response",
        "DELETE /key: 409 application/json response",
        "GET /key/{PK}: 200 application/json response",
        "GET /key/{PK}: 404 application/json response",
        "GET /key/{PK}: 410 application/json response",
        "GET /key/{PK}: default */* response",
        "HEAD /key/{PK}: 404 */* response",
        "HEAD /key/{PK}: 410 */* response",
        "HEAD /key/{PK}: default */* response",
        "POST /key/{PK}: application/jwt request body",
        "POST /key/{PK}: 404 application/json response",
        "POST /key/{PK}: default */* response",
        "POST /key/{PK}: 200 application/json response",
        "PUT /key/{PK}: application/jwt request body",
        "PUT /key/{PK}: 404 application/json response",
        "PUT /key/{PK}: 409 application/json response",
        "PUT /key/{PK}: default */* response",
        "PUT /key/{PK}: 200 application/json response",
        "DELETE /key/{PK}: default */* response",
        "DELETE /key/{PK}: 200 application/json response",
        "DELETE /key/{PK}: 401 application/json response",
        "DELETE /key/{PK}: 404 application/json response",
        "POST /login: application/jwt request body",
        "POST /login: 200 application/json response",
        "POST /login: 401 application/json response",
        "POST /login: default */* response",
        "POST /scope: application/jwt request body",
        "POST /scope: 201 application/json response",
        "POST /scope: 429 application/json response",
        "POST /scope: default */* response",
        "HEAD /scope/{job}: 404 application/json response",
        "HEAD /scope/{job}: default */* response",
        "POST /scope/{job}: 405 application/json response",
        "POST /scope/{job}: default */* response",
        "POST /scope/{job}: 202 application/json response",
        "POST /scope/{job}: 401 application/json response",
        "POST /scope/{job}: 404 application/json response",
        "PUT /scope/{job}: 200 application/jwt response",
        "PUT /scope/{job}: 404 application/jwt response",
        "PUT /scope/{job}: 409 application/jwt response",
        "PUT /scope/{job}: default */* response",
        "DELETE /scope/{job}: 404 application/json response",
        "DELETE /scope/{job}: default */* response",
        "DELETE /scope/{job}: 200 application/json response",
        "GET /scope/{job}: 200 application/json response",
        "GET /scope/{job}: 200 application/jwt response",
        "GET /scope/{job}: 404 application/json response",
        "GET /scope/{job}: 404 application/jwt response",
        "GET /scope/{job}: default */* response"
      ],
      "All schemas have descriptions": [
        "Error"
      ]
    }
  }
```

**Suggested Fix:**
Add missing documentation including descriptions, examples, and operation details

---

### Schema

#### P004: Request Schema Definitions (Failed) [critical]

Ensures all API requests have comprehensive schema definitions with proper data types, validation rules, and constraints

- **Status:** Failed
- **Message:** Request validation issues found: All numeric fields have min/max values: POST /scope: parameter test, POST /login.exp: application/jwt schema, POST /login.iat: application/jwt schema, POST /login.nbf: application/jwt schema; All schemas specify data types: POST /scope: application/jwt schema, POST /key: application/jwt schema, POST /key/{PK}: application/jwt schema, PUT /key/{PK}: application/jwt schema, POST /login: application/jwt schema; All string fields have length constraints: POST /scope.type: application/jwt schema, POST /scope.email: application/jwt schema, POST /scope.phone: application/jwt schema, POST /scope.scope: application/jwt schema, POST /scope.sub: application/jwt schema, PUT /scope/{job}: parameter job, DELETE /scope/{job}: parameter job, GET /scope/{job}: parameter job, HEAD /scope/{job}: parameter job, POST /scope/{job}: parameter job, POST /key.devtoken: application/jwt schema, POST /key.sub: application/jwt schema, DELETE /key: parameter email, DELETE /key: parameter phone, DELETE /key: parameter code, DELETE /key/{PK}: parameter PK, DELETE /key/{PK}: parameter secret, GET /key/{PK}: parameter PK, HEAD /key/{PK}: parameter PK, POST /key/{PK}: parameter PK, POST /key/{PK}.devtoken: application/jwt schema, POST /key/{PK}.sub: application/jwt schema, PUT /key/{PK}: parameter PK, PUT /key/{PK}.devtoken: application/jwt schema, PUT /key/{PK}.sub: application/jwt schema, POST /login: parameter callback, POST /login.sub: application/jwt schema, POST /login.aud: application/jwt schema, POST /login.iss: application/jwt schema
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
      "All numeric fields have min/max values": false,
      "All path parameters have schemas": true,
      "All query parameters have schemas": true,
      "All request bodies have content schemas": true,
      "All required fields are marked": true,
      "All schemas have appropriate constraints": true,
      "All schemas specify data types": false,
      "All string fields have length constraints": false
    },
    "messages": {},
    "missing_validation": {
      "All numeric fields have min/max values": [
        "POST /scope: parameter test",
        "POST /login.exp: application/jwt schema",
        "POST /login.iat: application/jwt schema",
        "POST /login.nbf: application/jwt schema"
      ],
      "All schemas specify data types": [
        "POST /scope: application/jwt schema",
        "POST /key: application/jwt schema",
        "POST /key/{PK}: application/jwt schema",
        "PUT /key/{PK}: application/jwt schema",
        "POST /login: application/jwt schema"
      ],
      "All string fields have length constraints": [
        "POST /scope.type: application/jwt schema",
        "POST /scope.email: application/jwt schema",
        "POST /scope.phone: application/jwt schema",
        "POST /scope.scope: application/jwt schema",
        "POST /scope.sub: application/jwt schema",
        "PUT /scope/{job}: parameter job",
        "DELETE /scope/{job}: parameter job",
        "GET /scope/{job}: parameter job",
        "HEAD /scope/{job}: parameter job",
        "POST /scope/{job}: parameter job",
        "POST /key.devtoken: application/jwt schema",
        "POST /key.sub: application/jwt schema",
        "DELETE /key: parameter email",
        "DELETE /key: parameter phone",
        "DELETE /key: parameter code",
        "DELETE /key/{PK}: parameter PK",
        "DELETE /key/{PK}: parameter secret",
        "GET /key/{PK}: parameter PK",
        "HEAD /key/{PK}: parameter PK",
        "POST /key/{PK}: parameter PK",
        "POST /key/{PK}.devtoken: application/jwt schema",
        "POST /key/{PK}.sub: application/jwt schema",
        "PUT /key/{PK}: parameter PK",
        "PUT /key/{PK}.devtoken: application/jwt schema",
        "PUT /key/{PK}.sub: application/jwt schema",
        "POST /login: parameter callback",
        "POST /login.sub: application/jwt schema",
        "POST /login.aud: application/jwt schema",
        "POST /login.iss: application/jwt schema"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive schema validation including data types, constraints, and required fields

---

### Error Handling

#### P003: Error Handling Standards (Failed) [critical]

Validates comprehensive error response documentation and consistent error handling patterns

- **Status:** Failed
- **Message:** Error handling issues found: All operations document 5xx error responses: POST /login, POST /scope, DELETE /scope/{job}, GET /scope/{job}, HEAD /scope/{job}, POST /scope/{job}, PUT /scope/{job}, DELETE /key, POST /key, PUT /key/{PK}, DELETE /key/{PK}, GET /key/{PK}, HEAD /key/{PK}, POST /key/{PK}; Error responses include error details schema: POST /login: 401 response, POST /scope: 429 response, DELETE /scope/{job}: 404 response, GET /scope/{job}: 404 response, HEAD /scope/{job}: 404 response, POST /scope/{job}: 401 response, POST /scope/{job}: 404 response, POST /scope/{job}: 405 response, PUT /scope/{job}: 404 response, PUT /scope/{job}: 409 response, DELETE /key: 401 response, DELETE /key: 404 response, DELETE /key: 409 response, POST /key: 409 response, PUT /key/{PK}: 404 response, PUT /key/{PK}: 409 response, DELETE /key/{PK}: 401 response, DELETE /key/{PK}: 404 response, GET /key/{PK}: 410 response, GET /key/{PK}: 404 response, HEAD /key/{PK}: 404 response, HEAD /key/{PK}: 410 response, POST /key/{PK}: 404 response
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
      "All operations document 5xx error responses": false,
      "Common error responses are defined in components": false,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": false
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "All operations document 5xx error responses": [
        "POST /login",
        "POST /scope",
        "DELETE /scope/{job}",
        "GET /scope/{job}",
        "HEAD /scope/{job}",
        "POST /scope/{job}",
        "PUT /scope/{job}",
        "DELETE /key",
        "POST /key",
        "PUT /key/{PK}",
        "DELETE /key/{PK}",
        "GET /key/{PK}",
        "HEAD /key/{PK}",
        "POST /key/{PK}"
      ],
      "Error responses include error details schema": [
        "POST /login: 401 response",
        "POST /scope: 429 response",
        "DELETE /scope/{job}: 404 response",
        "GET /scope/{job}: 404 response",
        "HEAD /scope/{job}: 404 response",
        "POST /scope/{job}: 401 response",
        "POST /scope/{job}: 404 response",
        "POST /scope/{job}: 405 response",
        "PUT /scope/{job}: 404 response",
        "PUT /scope/{job}: 409 response",
        "DELETE /key: 401 response",
        "DELETE /key: 404 response",
        "DELETE /key: 409 response",
        "POST /key: 409 response",
        "PUT /key/{PK}: 404 response",
        "PUT /key/{PK}: 409 response",
        "DELETE /key/{PK}: 401 response",
        "DELETE /key/{PK}: 404 response",
        "GET /key/{PK}: 410 response",
        "GET /key/{PK}: 404 response",
        "HEAD /key/{PK}: 404 response",
        "HEAD /key/{PK}: 410 response",
        "POST /key/{PK}: 404 response"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive error response documentation including codes, messages, and consistent error schemas

---

### Security

#### P005: Security Standards (Failed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Failed
- **Message:** Security validation failed: Security schemes are defined: No security schemes defined in components.securitySchemes; Global security requirements are set: No top-level security requirements defined; Operation-level security is defined: Endpoints without security: POST /key, DELETE /key, DELETE /key/{PK}, GET /key/{PK}, HEAD /key/{PK}, POST /key/{PK}, PUT /key/{PK}, POST /login, POST /scope, DELETE /scope/{job}, GET /scope/{job}, HEAD /scope/{job}, POST /scope/{job}, PUT /scope/{job}
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
      "Global security requirements are set": false,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": false,
      "Security requirements are consistent": true,
      "Security schemes are defined": false
    },
    "messages": {
      "Global security requirements are set": "No top-level security requirements defined",
      "Operation-level security is defined": "Endpoints without security: POST /key, DELETE /key, DELETE /key/{PK}, GET /key/{PK}, HEAD /key/{PK}, POST /key/{PK}, PUT /key/{PK}, POST /login, POST /scope, DELETE /scope/{job}, GET /scope/{job}, HEAD /scope/{job}, POST /scope/{job}, PUT /scope/{job}",
      "Security schemes are defined": "No security schemes defined in components.securitySchemes"
    }
  }
```

**Suggested Fix:**
Review security schemes, global security requirements, and per-operation security

---

### Versioning

#### P008: API Versioning Strategy (Failed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Failed
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "6" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy
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
      "Breaking changes are documented": false,
      "Deprecation notices are present": true,
      "Migration guides are referenced": false,
      "Version compatibility is specified": false,
      "Version follows semantic versioning": false,
      "Versioning strategy is documented": false
    },
    "messages": {
      "Breaking changes are documented": "Info description does not reference breaking changes or a changelog",
      "Migration guides are referenced": "Info description does not reference migration or upgrade guides",
      "Version compatibility is specified": "Info description does not mention version compatibility",
      "Version follows semantic versioning": "Version \"6\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

