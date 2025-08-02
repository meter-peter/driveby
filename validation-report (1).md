# API Validation Report

Generated: 2025-06-08T20:32:21+03:00
Environment: development
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 1
- Failed Checks: 5
- Critical Issues: 2
- Warnings: 3
- Info: 0

### Categories
- Specification
- Documentation
- Error Handling
- Schema
- Security


### Failed Tags
- openapi
- compliance
- quality
- schema
- authentication
- usability
- responses
- standards
- security
- authorization
- specification
- documentation
- errors
- validation
- request


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: References are resolvable: invalid components: schema "Task": unsupported 'type' value "null"
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
      "OpenAPI version is 3.0.x or 3.1.0": true,
      "Paths are properly defined": true,
      "References are resolvable": false,
      "Required info fields (title, version) are present": true,
      "Valid HTTP methods used": true
    },
    "messages": {
      "References are resolvable": "invalid components: schema \"Task\": unsupported 'type' value \"null\""
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [warning]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All request/response bodies have examples: GET /products: 422 application/json response, POST /products: request body, POST /products: 201 application/json response, POST /products: 422 application/json response, DELETE /products/{product_id}: 422 application/json response, GET /products/{product_id}: 200 application/json response, GET /products/{product_id}: 422 application/json response, PUT /products/{product_id}: request body, PUT /products/{product_id}: 422 application/json response, PUT /products/{product_id}: 200 application/json response, POST /tasks: request body, POST /tasks: application/json request body, POST /tasks: 422 application/json response, POST /test/echo: request body, POST /test/echo: 422 application/json response; All enums have descriptions: ProductCategory: enum value electronics, ProductCategory: enum value clothing, ProductCategory: enum value food, ProductCategory: enum value books, ProductCategory: enum value other; All schemas have descriptions: ValidationError, HTTPValidationError
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
      "All enums have descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All enums have descriptions": [
        "ProductCategory: enum value electronics",
        "ProductCategory: enum value clothing",
        "ProductCategory: enum value food",
        "ProductCategory: enum value books",
        "ProductCategory: enum value other"
      ],
      "All request/response bodies have examples": [
        "GET /products: 422 application/json response",
        "POST /products: request body",
        "POST /products: 201 application/json response",
        "POST /products: 422 application/json response",
        "DELETE /products/{product_id}: 422 application/json response",
        "GET /products/{product_id}: 200 application/json response",
        "GET /products/{product_id}: 422 application/json response",
        "PUT /products/{product_id}: request body",
        "PUT /products/{product_id}: 422 application/json response",
        "PUT /products/{product_id}: 200 application/json response",
        "POST /tasks: request body",
        "POST /tasks: application/json request body",
        "POST /tasks: 422 application/json response",
        "POST /test/echo: request body",
        "POST /test/echo: 422 application/json response"
      ],
      "All schemas have descriptions": [
        "ValidationError",
        "HTTPValidationError"
      ]
    }
  }
```

**Suggested Fix:**
Add missing documentation including descriptions, examples, and operation details

---

### Schema

#### P004: Request Schema Definitions (Failed) [warning]

Ensures all API requests have comprehensive schema definitions with proper data types, validation rules, and constraints

- **Status:** Failed
- **Message:** Request validation issues found: All required fields are marked: POST /tasks: application/json.title, POST /products: application/json.name, POST /products: application/json.description, POST /products: application/json.price, POST /products: application/json.category, DELETE /products/{product_id}: parameter product_id, GET /products/{product_id}: parameter product_id, PUT /products/{product_id}: parameter product_id, PUT /products/{product_id}: application/json.name, PUT /products/{product_id}: application/json.description, PUT /products/{product_id}: application/json.price, PUT /products/{product_id}: application/json.category; All schemas specify data types: POST /tasks.description: application/json schema, GET /products: parameter category, GET /products: parameter min_price, GET /products: parameter max_price, GET /products: parameter in_stock; All string fields have length constraints: POST /products.category: application/json schema, POST /products.description: application/json schema, POST /products.tags[]: application/json schema, DELETE /products/{product_id}: parameter product_id, GET /products/{product_id}: parameter product_id, PUT /products/{product_id}: parameter product_id, PUT /products/{product_id}.category: application/json schema, PUT /products/{product_id}.description: application/json schema, PUT /products/{product_id}.tags[]: application/json schema; All numeric fields have min/max values: POST /products.price: application/json schema, PUT /products/{product_id}.price: application/json schema
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
      "All required fields are marked": false,
      "All schemas have appropriate constraints": true,
      "All schemas specify data types": false,
      "All string fields have length constraints": false
    },
    "messages": {},
    "missing_validation": {
      "All numeric fields have min/max values": [
        "POST /products.price: application/json schema",
        "PUT /products/{product_id}.price: application/json schema"
      ],
      "All required fields are marked": [
        "POST /tasks: application/json.title",
        "POST /products: application/json.name",
        "POST /products: application/json.description",
        "POST /products: application/json.price",
        "POST /products: application/json.category",
        "DELETE /products/{product_id}: parameter product_id",
        "GET /products/{product_id}: parameter product_id",
        "PUT /products/{product_id}: parameter product_id",
        "PUT /products/{product_id}: application/json.name",
        "PUT /products/{product_id}: application/json.description",
        "PUT /products/{product_id}: application/json.price",
        "PUT /products/{product_id}: application/json.category"
      ],
      "All schemas specify data types": [
        "POST /tasks.description: application/json schema",
        "GET /products: parameter category",
        "GET /products: parameter min_price",
        "GET /products: parameter max_price",
        "GET /products: parameter in_stock"
      ],
      "All string fields have length constraints": [
        "POST /products.category: application/json schema",
        "POST /products.description: application/json schema",
        "POST /products.tags[]: application/json schema",
        "DELETE /products/{product_id}: parameter product_id",
        "GET /products/{product_id}: parameter product_id",
        "PUT /products/{product_id}: parameter product_id",
        "PUT /products/{product_id}.category: application/json schema",
        "PUT /products/{product_id}.description: application/json schema",
        "PUT /products/{product_id}.tags[]: application/json schema"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive schema validation including data types, constraints, and required fields

---

### Error Handling

#### P003: Error Handling Standards (Failed) [warning]

Validates comprehensive error response documentation and consistent error handling patterns

- **Status:** Failed
- **Message:** Error handling issues found: All operations document 4xx error responses: GET /health, GET /test/health; All operations document 5xx error responses: GET /health, GET /products, POST /products, DELETE /products/{product_id}, GET /products/{product_id}, PUT /products/{product_id}, POST /tasks, POST /test/echo, GET /test/health; Error responses include error details schema: GET /products: 422 response, POST /products: 422 response, DELETE /products/{product_id}: 422 response, DELETE /products/{product_id}: 404 response, GET /products/{product_id}: 404 response, GET /products/{product_id}: 422 response, PUT /products/{product_id}: 404 response, PUT /products/{product_id}: 422 response, POST /tasks: 400 response, POST /tasks: 422 response, POST /test/echo: 422 response
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
      "Error responses include error details schema": false
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "All operations document 4xx error responses": [
        "GET /health",
        "GET /test/health"
      ],
      "All operations document 5xx error responses": [
        "GET /health",
        "GET /products",
        "POST /products",
        "DELETE /products/{product_id}",
        "GET /products/{product_id}",
        "PUT /products/{product_id}",
        "POST /tasks",
        "POST /test/echo",
        "GET /test/health"
      ],
      "Error responses include error details schema": [
        "GET /products: 422 response",
        "POST /products: 422 response",
        "DELETE /products/{product_id}: 422 response",
        "DELETE /products/{product_id}: 404 response",
        "GET /products/{product_id}: 404 response",
        "GET /products/{product_id}: 422 response",
        "PUT /products/{product_id}: 404 response",
        "PUT /products/{product_id}: 422 response",
        "POST /tasks: 400 response",
        "POST /tasks: 422 response",
        "POST /test/echo: 422 response"
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
- **Message:** No security schemes defined
- **Tags:** security, authentication, authorization

**Checks Performed:**
- Security schemes are defined
- Global security requirements are set
- Operation-level security is defined
- OAuth2 scopes are documented
- API keys are properly described
- Authentication headers are specified
- Security requirements are consistent

**Suggested Fix:**
Define security schemes in components.securitySchemes

---

### Versioning

#### P008: API Versioning Strategy (Passed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Passed
- **Message:** 
- **Tags:** versioning, compatibility, lifecycle

**Checks Performed:**
- API version is specified
- Version follows semantic versioning
- Versioning strategy is documented
- Deprecation notices are present
- Breaking changes are documented
- Version compatibility is specified
- Migration guides are referenced

---

