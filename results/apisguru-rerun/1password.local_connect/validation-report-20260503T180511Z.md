# API Validation Report

Generated: 2026-05-03T21:05:11+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 0
- Failed Checks: 6
- Critical Issues: 5
- Warnings: 1
- Info: 0

### Categories
- Specification
- Documentation
- Error Handling
- Schema
- Security
- Versioning


### Failed Tags
- versioning
- lifecycle
- documentation
- usability
- compatibility
- authentication
- authorization
- openapi
- quality
- responses
- schema
- request
- security
- specification
- compliance
- errors
- standards
- validation


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid paths: invalid path /vaults/{vaultUuid}/items/{itemUuid}: invalid operation PATCH: example PatchItemAttr: Error at "/0/value": value must be an object
Schema:
  {
    "type": "object"
  }

Value:
  true

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
      "Specification structure is valid": false,
      "Specification version is present": true,
      "Valid HTTP methods used": true
    },
    "messages": {
      "Specification structure is valid": "invalid paths: invalid path /vaults/{vaultUuid}/items/{itemUuid}: invalid operation PATCH: example PatchItemAttr: Error at \"/0/value\": value must be an object\nSchema:\n  {\n    \"type\": \"object\"\n  }\n\nValue:\n  true\n"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All operations have detailed descriptions: GET /heartbeat, GET /vaults, GET /vaults/{vaultUuid}/items, POST /vaults/{vaultUuid}/items, GET /vaults/{vaultUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}/files, GET /health, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content, GET /activity, DELETE /vaults/{vaultUuid}/items/{itemUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}, PUT /vaults/{vaultUuid}/items/{itemUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}; All request/response bodies have examples: GET /heartbeat: 200 text/plain response, GET /metrics: 200 text/plain response, GET /vaults: 200 application/json response, GET /vaults/{vaultUuid}/items: 200 application/json response, POST /vaults/{vaultUuid}/items: request body, POST /vaults/{vaultUuid}/items: application/json request body, POST /vaults/{vaultUuid}/items: 200 application/json response, GET /vaults/{vaultUuid}: 200 application/json response, GET /vaults/{vaultUuid}/items/{itemUuid}/files: 200 application/json response, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: 200 application/octet-stream response, GET /activity: 200 application/json response, GET /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response, PATCH /vaults/{vaultUuid}/items/{itemUuid}: request body, PATCH /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response, PUT /vaults/{vaultUuid}/items/{itemUuid}: request body, PUT /vaults/{vaultUuid}/items/{itemUuid}: application/json request body, PUT /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: 200 application/json response; All schemas have descriptions: Patch, Field, File, Vault, FullItem, ErrorResponse, Item
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
      "All operations have detailed descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": false
    },
    "messages": {
      "License information is provided": "License information is missing"
    },
    "missing_docs": {
      "All operations have detailed descriptions": [
        "GET /heartbeat",
        "GET /vaults",
        "GET /vaults/{vaultUuid}/items",
        "POST /vaults/{vaultUuid}/items",
        "GET /vaults/{vaultUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files",
        "GET /health",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content",
        "GET /activity",
        "DELETE /vaults/{vaultUuid}/items/{itemUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}"
      ],
      "All request/response bodies have examples": [
        "GET /heartbeat: 200 text/plain response",
        "GET /metrics: 200 text/plain response",
        "GET /vaults: 200 application/json response",
        "GET /vaults/{vaultUuid}/items: 200 application/json response",
        "POST /vaults/{vaultUuid}/items: request body",
        "POST /vaults/{vaultUuid}/items: application/json request body",
        "POST /vaults/{vaultUuid}/items: 200 application/json response",
        "GET /vaults/{vaultUuid}: 200 application/json response",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files: 200 application/json response",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: 200 application/octet-stream response",
        "GET /activity: 200 application/json response",
        "GET /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response",
        "PATCH /vaults/{vaultUuid}/items/{itemUuid}: request body",
        "PATCH /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}: request body",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}: application/json request body",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}: 200 application/json response",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: 200 application/json response"
      ],
      "All schemas have descriptions": [
        "Patch",
        "Field",
        "File",
        "Vault",
        "FullItem",
        "ErrorResponse",
        "Item"
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
- **Message:** Request validation issues found: All schemas specify data types: POST /vaults/{vaultUuid}/items: application/json schema, PUT /vaults/{vaultUuid}/items/{itemUuid}: application/json schema; All string fields have length constraints: GET /vaults: parameter filter, GET /vaults/{vaultUuid}/items: parameter filter, /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter vaultUuid, /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter itemUuid, /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter fileUuid, PATCH /vaults/{vaultUuid}/items/{itemUuid}[].op: application/json schema, PATCH /vaults/{vaultUuid}/items/{itemUuid}[].path: application/json schema, GET /vaults/{vaultUuid}/items/{itemUuid}/files: parameter vaultUuid, GET /vaults/{vaultUuid}/items/{itemUuid}/files: parameter itemUuid, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter vaultUuid, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter itemUuid, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter fileUuid; All numeric fields have min/max values: GET /activity: parameter limit, GET /activity: parameter offset
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
        "GET /activity: parameter limit",
        "GET /activity: parameter offset"
      ],
      "All schemas specify data types": [
        "POST /vaults/{vaultUuid}/items: application/json schema",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}: application/json schema"
      ],
      "All string fields have length constraints": [
        "GET /vaults: parameter filter",
        "GET /vaults/{vaultUuid}/items: parameter filter",
        "/vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter vaultUuid",
        "/vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter itemUuid",
        "/vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content: parameter fileUuid",
        "PATCH /vaults/{vaultUuid}/items/{itemUuid}[].op: application/json schema",
        "PATCH /vaults/{vaultUuid}/items/{itemUuid}[].path: application/json schema",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files: parameter vaultUuid",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files: parameter itemUuid",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter vaultUuid",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter itemUuid",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}: parameter fileUuid"
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
- **Message:** Error handling issues found: All operations document 4xx error responses: GET /heartbeat, GET /metrics, GET /health; All operations document 5xx error responses: GET /heartbeat, GET /metrics, GET /vaults, GET /vaults/{vaultUuid}/items, POST /vaults/{vaultUuid}/items, GET /vaults/{vaultUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content, GET /activity, PUT /vaults/{vaultUuid}/items/{itemUuid}, DELETE /vaults/{vaultUuid}/items/{itemUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}, PATCH /vaults/{vaultUuid}/items/{itemUuid}, GET /vaults/{vaultUuid}/items/{itemUuid}/files, GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}, GET /health
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
      "All operations document 4xx error responses": false,
      "All operations document 5xx error responses": false,
      "Common error responses are defined in components": false,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": true
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "All operations document 4xx error responses": [
        "GET /heartbeat",
        "GET /metrics",
        "GET /health"
      ],
      "All operations document 5xx error responses": [
        "GET /heartbeat",
        "GET /metrics",
        "GET /vaults",
        "GET /vaults/{vaultUuid}/items",
        "POST /vaults/{vaultUuid}/items",
        "GET /vaults/{vaultUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}/content",
        "GET /activity",
        "PUT /vaults/{vaultUuid}/items/{itemUuid}",
        "DELETE /vaults/{vaultUuid}/items/{itemUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}",
        "PATCH /vaults/{vaultUuid}/items/{itemUuid}",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files",
        "GET /vaults/{vaultUuid}/items/{itemUuid}/files/{fileUuid}",
        "GET /health"
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
- **Message:** Security validation failed: Global security requirements are set: No top-level security requirements defined; Operation-level security is defined: Endpoints without security: GET /heartbeat, GET /health, GET /metrics
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
      "Security schemes are defined": true
    },
    "messages": {
      "Global security requirements are set": "No top-level security requirements defined",
      "Operation-level security is defined": "Endpoints without security: GET /heartbeat, GET /health, GET /metrics"
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
- **Message:** Versioning validation failed: Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides
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
      "Version follows semantic versioning": true,
      "Versioning strategy is documented": false
    },
    "messages": {
      "Breaking changes are documented": "Info description does not reference breaking changes or a changelog",
      "Migration guides are referenced": "Info description does not reference migration or upgrade guides",
      "Version compatibility is specified": "Info description does not mention version compatibility",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

