# API Validation Report

Generated: 2026-05-03T20:52:55+03:00
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
- request
- authentication
- authorization
- versioning
- compatibility
- errors
- security
- lifecycle
- documentation
- usability
- responses
- quality
- standards
- schema
- validation


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: GET /store/order/{orderId}: 200 application/json response, GET /store/order/{orderId}: 200 application/xml response, GET /user/login: 200 application/json response, GET /user/login: 200 application/xml response, POST /store/order: request body, POST /store/order: application/x-www-form-urlencoded request body, POST /store/order: application/xml request body, POST /store/order: application/json request body, POST /store/order: 200 application/json response, POST /user/createWithList: request body, POST /user/createWithList: application/json request body, POST /user/createWithList: 200 application/xml response, POST /user/createWithList: 200 application/json response, GET /pet/findByStatus: 200 application/json response, GET /pet/findByStatus: 200 application/xml response, GET /pet/{petId}: 200 application/json response, GET /pet/{petId}: 200 application/xml response, POST /pet/{petId}: 200 application/json response, POST /pet/{petId}: 200 application/xml response, POST /pet/{petId}/uploadImage: request body, POST /pet/{petId}/uploadImage: application/octet-stream request body, POST /pet/{petId}/uploadImage: 200 application/json response, POST /user: application/x-www-form-urlencoded request body, POST /user: application/xml request body, POST /user: application/json request body, POST /user: 200 application/json response, POST /user: 200 application/xml response, GET /user/{username}: 200 application/xml response, GET /user/{username}: 200 application/json response, PUT /user/{username}: application/json request body, PUT /user/{username}: application/x-www-form-urlencoded request body, PUT /user/{username}: application/xml request body, POST /pet: application/json request body, POST /pet: application/x-www-form-urlencoded request body, POST /pet: application/xml request body, POST /pet: 200 application/xml response, POST /pet: 200 application/json response, PUT /pet: application/x-www-form-urlencoded request body, PUT /pet: application/xml request body, PUT /pet: application/json request body, PUT /pet: 200 application/json response, PUT /pet: 200 application/xml response, GET /pet/findByTags: 200 application/json response, GET /pet/findByTags: 200 application/xml response, GET /store/inventory: 200 application/json response; All parameters have descriptions: DELETE /pet/{petId}: parameter api_key; All schemas have descriptions: Category, Order, Pet, Tag, User, ApiResponse
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
      "All parameters have descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All parameters have descriptions": [
        "DELETE /pet/{petId}: parameter api_key"
      ],
      "All request/response bodies have examples": [
        "GET /store/order/{orderId}: 200 application/json response",
        "GET /store/order/{orderId}: 200 application/xml response",
        "GET /user/login: 200 application/json response",
        "GET /user/login: 200 application/xml response",
        "POST /store/order: request body",
        "POST /store/order: application/x-www-form-urlencoded request body",
        "POST /store/order: application/xml request body",
        "POST /store/order: application/json request body",
        "POST /store/order: 200 application/json response",
        "POST /user/createWithList: request body",
        "POST /user/createWithList: application/json request body",
        "POST /user/createWithList: 200 application/xml response",
        "POST /user/createWithList: 200 application/json response",
        "GET /pet/findByStatus: 200 application/json response",
        "GET /pet/findByStatus: 200 application/xml response",
        "GET /pet/{petId}: 200 application/json response",
        "GET /pet/{petId}: 200 application/xml response",
        "POST /pet/{petId}: 200 application/json response",
        "POST /pet/{petId}: 200 application/xml response",
        "POST /pet/{petId}/uploadImage: request body",
        "POST /pet/{petId}/uploadImage: application/octet-stream request body",
        "POST /pet/{petId}/uploadImage: 200 application/json response",
        "POST /user: application/x-www-form-urlencoded request body",
        "POST /user: application/xml request body",
        "POST /user: application/json request body",
        "POST /user: 200 application/json response",
        "POST /user: 200 application/xml response",
        "GET /user/{username}: 200 application/xml response",
        "GET /user/{username}: 200 application/json response",
        "PUT /user/{username}: application/json request body",
        "PUT /user/{username}: application/x-www-form-urlencoded request body",
        "PUT /user/{username}: application/xml request body",
        "POST /pet: application/json request body",
        "POST /pet: application/x-www-form-urlencoded request body",
        "POST /pet: application/xml request body",
        "POST /pet: 200 application/xml response",
        "POST /pet: 200 application/json response",
        "PUT /pet: application/x-www-form-urlencoded request body",
        "PUT /pet: application/xml request body",
        "PUT /pet: application/json request body",
        "PUT /pet: 200 application/json response",
        "PUT /pet: 200 application/xml response",
        "GET /pet/findByTags: 200 application/json response",
        "GET /pet/findByTags: 200 application/xml response",
        "GET /store/inventory: 200 application/json response"
      ],
      "All schemas have descriptions": [
        "Category",
        "Order",
        "Pet",
        "Tag",
        "User",
        "ApiResponse"
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
- **Message:** Request validation issues found: All string fields have length constraints: PUT /pet.status: application/x-www-form-urlencoded schema, PUT /pet.tags[].name: application/x-www-form-urlencoded schema, PUT /pet.category.name: application/x-www-form-urlencoded schema, PUT /pet.name: application/x-www-form-urlencoded schema, PUT /pet.photoUrls[]: application/x-www-form-urlencoded schema, PUT /pet.name: application/xml schema, PUT /pet.photoUrls[]: application/xml schema, PUT /pet.status: application/xml schema, PUT /pet.tags[].name: application/xml schema, PUT /pet.category.name: application/xml schema, PUT /pet.tags[].name: application/json schema, PUT /pet.category.name: application/json schema, PUT /pet.name: application/json schema, PUT /pet.photoUrls[]: application/json schema, PUT /pet.status: application/json schema, POST /pet.status: application/json schema, POST /pet.tags[].name: application/json schema, POST /pet.category.name: application/json schema, POST /pet.name: application/json schema, POST /pet.photoUrls[]: application/json schema, POST /pet.name: application/x-www-form-urlencoded schema, POST /pet.photoUrls[]: application/x-www-form-urlencoded schema, POST /pet.status: application/x-www-form-urlencoded schema, POST /pet.tags[].name: application/x-www-form-urlencoded schema, POST /pet.category.name: application/x-www-form-urlencoded schema, POST /pet.category.name: application/xml schema, POST /pet.name: application/xml schema, POST /pet.photoUrls[]: application/xml schema, POST /pet.status: application/xml schema, POST /pet.tags[].name: application/xml schema, POST /pet/{petId}: parameter name, POST /pet/{petId}: parameter status, DELETE /pet/{petId}: parameter api_key, GET /user/login: parameter username, GET /user/login: parameter password, POST /store/order.shipDate: application/xml schema, POST /store/order.status: application/xml schema, POST /store/order.status: application/json schema, POST /store/order.shipDate: application/json schema, POST /store/order.shipDate: application/x-www-form-urlencoded schema, POST /store/order.status: application/x-www-form-urlencoded schema, POST /user/createWithList[].phone: application/json schema, POST /user/createWithList[].username: application/json schema, POST /user/createWithList[].email: application/json schema, POST /user/createWithList[].firstName: application/json schema, POST /user/createWithList[].lastName: application/json schema, POST /user/createWithList[].password: application/json schema, GET /user/{username}: parameter username, PUT /user/{username}: parameter username, PUT /user/{username}.username: application/x-www-form-urlencoded schema, PUT /user/{username}.email: application/x-www-form-urlencoded schema, PUT /user/{username}.firstName: application/x-www-form-urlencoded schema, PUT /user/{username}.lastName: application/x-www-form-urlencoded schema, PUT /user/{username}.password: application/x-www-form-urlencoded schema, PUT /user/{username}.phone: application/x-www-form-urlencoded schema, PUT /user/{username}.username: application/xml schema, PUT /user/{username}.email: application/xml schema, PUT /user/{username}.firstName: application/xml schema, PUT /user/{username}.lastName: application/xml schema, PUT /user/{username}.password: application/xml schema, PUT /user/{username}.phone: application/xml schema, PUT /user/{username}.lastName: application/json schema, PUT /user/{username}.password: application/json schema, PUT /user/{username}.phone: application/json schema, PUT /user/{username}.username: application/json schema, PUT /user/{username}.email: application/json schema, PUT /user/{username}.firstName: application/json schema, DELETE /user/{username}: parameter username, GET /pet/findByStatus: parameter status, POST /pet/{petId}/uploadImage: parameter additionalMetadata, POST /pet/{petId}/uploadImage: application/octet-stream schema, POST /user.password: application/json schema, POST /user.phone: application/json schema, POST /user.username: application/json schema, POST /user.email: application/json schema, POST /user.firstName: application/json schema, POST /user.lastName: application/json schema, POST /user.username: application/x-www-form-urlencoded schema, POST /user.email: application/x-www-form-urlencoded schema, POST /user.firstName: application/x-www-form-urlencoded schema, POST /user.lastName: application/x-www-form-urlencoded schema, POST /user.password: application/x-www-form-urlencoded schema, POST /user.phone: application/x-www-form-urlencoded schema, POST /user.email: application/xml schema, POST /user.firstName: application/xml schema, POST /user.lastName: application/xml schema, POST /user.password: application/xml schema, POST /user.phone: application/xml schema, POST /user.username: application/xml schema; All numeric fields have min/max values: PUT /pet.tags[].id: application/x-www-form-urlencoded schema, PUT /pet.category.id: application/x-www-form-urlencoded schema, PUT /pet.id: application/x-www-form-urlencoded schema, PUT /pet.id: application/xml schema, PUT /pet.tags[].id: application/xml schema, PUT /pet.category.id: application/xml schema, PUT /pet.tags[].id: application/json schema, PUT /pet.category.id: application/json schema, PUT /pet.id: application/json schema, POST /pet.tags[].id: application/json schema, POST /pet.category.id: application/json schema, POST /pet.id: application/json schema, POST /pet.id: application/x-www-form-urlencoded schema, POST /pet.tags[].id: application/x-www-form-urlencoded schema, POST /pet.category.id: application/x-www-form-urlencoded schema, POST /pet.category.id: application/xml schema, POST /pet.id: application/xml schema, POST /pet.tags[].id: application/xml schema, POST /pet/{petId}: parameter petId, DELETE /pet/{petId}: parameter petId, GET /pet/{petId}: parameter petId, DELETE /store/order/{orderId}: parameter orderId, GET /store/order/{orderId}: parameter orderId, POST /store/order.id: application/xml schema, POST /store/order.petId: application/xml schema, POST /store/order.quantity: application/xml schema, POST /store/order.id: application/json schema, POST /store/order.petId: application/json schema, POST /store/order.quantity: application/json schema, POST /store/order.id: application/x-www-form-urlencoded schema, POST /store/order.petId: application/x-www-form-urlencoded schema, POST /store/order.quantity: application/x-www-form-urlencoded schema, POST /user/createWithList[].userStatus: application/json schema, POST /user/createWithList[].id: application/json schema, PUT /user/{username}.id: application/x-www-form-urlencoded schema, PUT /user/{username}.userStatus: application/x-www-form-urlencoded schema, PUT /user/{username}.userStatus: application/xml schema, PUT /user/{username}.id: application/xml schema, PUT /user/{username}.userStatus: application/json schema, PUT /user/{username}.id: application/json schema, POST /pet/{petId}/uploadImage: parameter petId, POST /user.userStatus: application/json schema, POST /user.id: application/json schema, POST /user.id: application/x-www-form-urlencoded schema, POST /user.userStatus: application/x-www-form-urlencoded schema, POST /user.id: application/xml schema, POST /user.userStatus: application/xml schema
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
      "All schemas specify data types": true,
      "All string fields have length constraints": false
    },
    "messages": {},
    "missing_validation": {
      "All numeric fields have min/max values": [
        "PUT /pet.tags[].id: application/x-www-form-urlencoded schema",
        "PUT /pet.category.id: application/x-www-form-urlencoded schema",
        "PUT /pet.id: application/x-www-form-urlencoded schema",
        "PUT /pet.id: application/xml schema",
        "PUT /pet.tags[].id: application/xml schema",
        "PUT /pet.category.id: application/xml schema",
        "PUT /pet.tags[].id: application/json schema",
        "PUT /pet.category.id: application/json schema",
        "PUT /pet.id: application/json schema",
        "POST /pet.tags[].id: application/json schema",
        "POST /pet.category.id: application/json schema",
        "POST /pet.id: application/json schema",
        "POST /pet.id: application/x-www-form-urlencoded schema",
        "POST /pet.tags[].id: application/x-www-form-urlencoded schema",
        "POST /pet.category.id: application/x-www-form-urlencoded schema",
        "POST /pet.category.id: application/xml schema",
        "POST /pet.id: application/xml schema",
        "POST /pet.tags[].id: application/xml schema",
        "POST /pet/{petId}: parameter petId",
        "DELETE /pet/{petId}: parameter petId",
        "GET /pet/{petId}: parameter petId",
        "DELETE /store/order/{orderId}: parameter orderId",
        "GET /store/order/{orderId}: parameter orderId",
        "POST /store/order.id: application/xml schema",
        "POST /store/order.petId: application/xml schema",
        "POST /store/order.quantity: application/xml schema",
        "POST /store/order.id: application/json schema",
        "POST /store/order.petId: application/json schema",
        "POST /store/order.quantity: application/json schema",
        "POST /store/order.id: application/x-www-form-urlencoded schema",
        "POST /store/order.petId: application/x-www-form-urlencoded schema",
        "POST /store/order.quantity: application/x-www-form-urlencoded schema",
        "POST /user/createWithList[].userStatus: application/json schema",
        "POST /user/createWithList[].id: application/json schema",
        "PUT /user/{username}.id: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.userStatus: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.userStatus: application/xml schema",
        "PUT /user/{username}.id: application/xml schema",
        "PUT /user/{username}.userStatus: application/json schema",
        "PUT /user/{username}.id: application/json schema",
        "POST /pet/{petId}/uploadImage: parameter petId",
        "POST /user.userStatus: application/json schema",
        "POST /user.id: application/json schema",
        "POST /user.id: application/x-www-form-urlencoded schema",
        "POST /user.userStatus: application/x-www-form-urlencoded schema",
        "POST /user.id: application/xml schema",
        "POST /user.userStatus: application/xml schema"
      ],
      "All string fields have length constraints": [
        "PUT /pet.status: application/x-www-form-urlencoded schema",
        "PUT /pet.tags[].name: application/x-www-form-urlencoded schema",
        "PUT /pet.category.name: application/x-www-form-urlencoded schema",
        "PUT /pet.name: application/x-www-form-urlencoded schema",
        "PUT /pet.photoUrls[]: application/x-www-form-urlencoded schema",
        "PUT /pet.name: application/xml schema",
        "PUT /pet.photoUrls[]: application/xml schema",
        "PUT /pet.status: application/xml schema",
        "PUT /pet.tags[].name: application/xml schema",
        "PUT /pet.category.name: application/xml schema",
        "PUT /pet.tags[].name: application/json schema",
        "PUT /pet.category.name: application/json schema",
        "PUT /pet.name: application/json schema",
        "PUT /pet.photoUrls[]: application/json schema",
        "PUT /pet.status: application/json schema",
        "POST /pet.status: application/json schema",
        "POST /pet.tags[].name: application/json schema",
        "POST /pet.category.name: application/json schema",
        "POST /pet.name: application/json schema",
        "POST /pet.photoUrls[]: application/json schema",
        "POST /pet.name: application/x-www-form-urlencoded schema",
        "POST /pet.photoUrls[]: application/x-www-form-urlencoded schema",
        "POST /pet.status: application/x-www-form-urlencoded schema",
        "POST /pet.tags[].name: application/x-www-form-urlencoded schema",
        "POST /pet.category.name: application/x-www-form-urlencoded schema",
        "POST /pet.category.name: application/xml schema",
        "POST /pet.name: application/xml schema",
        "POST /pet.photoUrls[]: application/xml schema",
        "POST /pet.status: application/xml schema",
        "POST /pet.tags[].name: application/xml schema",
        "POST /pet/{petId}: parameter name",
        "POST /pet/{petId}: parameter status",
        "DELETE /pet/{petId}: parameter api_key",
        "GET /user/login: parameter username",
        "GET /user/login: parameter password",
        "POST /store/order.shipDate: application/xml schema",
        "POST /store/order.status: application/xml schema",
        "POST /store/order.status: application/json schema",
        "POST /store/order.shipDate: application/json schema",
        "POST /store/order.shipDate: application/x-www-form-urlencoded schema",
        "POST /store/order.status: application/x-www-form-urlencoded schema",
        "POST /user/createWithList[].phone: application/json schema",
        "POST /user/createWithList[].username: application/json schema",
        "POST /user/createWithList[].email: application/json schema",
        "POST /user/createWithList[].firstName: application/json schema",
        "POST /user/createWithList[].lastName: application/json schema",
        "POST /user/createWithList[].password: application/json schema",
        "GET /user/{username}: parameter username",
        "PUT /user/{username}: parameter username",
        "PUT /user/{username}.username: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.email: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.firstName: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.lastName: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.password: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.phone: application/x-www-form-urlencoded schema",
        "PUT /user/{username}.username: application/xml schema",
        "PUT /user/{username}.email: application/xml schema",
        "PUT /user/{username}.firstName: application/xml schema",
        "PUT /user/{username}.lastName: application/xml schema",
        "PUT /user/{username}.password: application/xml schema",
        "PUT /user/{username}.phone: application/xml schema",
        "PUT /user/{username}.lastName: application/json schema",
        "PUT /user/{username}.password: application/json schema",
        "PUT /user/{username}.phone: application/json schema",
        "PUT /user/{username}.username: application/json schema",
        "PUT /user/{username}.email: application/json schema",
        "PUT /user/{username}.firstName: application/json schema",
        "DELETE /user/{username}: parameter username",
        "GET /pet/findByStatus: parameter status",
        "POST /pet/{petId}/uploadImage: parameter additionalMetadata",
        "POST /pet/{petId}/uploadImage: application/octet-stream schema",
        "POST /user.password: application/json schema",
        "POST /user.phone: application/json schema",
        "POST /user.username: application/json schema",
        "POST /user.email: application/json schema",
        "POST /user.firstName: application/json schema",
        "POST /user.lastName: application/json schema",
        "POST /user.username: application/x-www-form-urlencoded schema",
        "POST /user.email: application/x-www-form-urlencoded schema",
        "POST /user.firstName: application/x-www-form-urlencoded schema",
        "POST /user.lastName: application/x-www-form-urlencoded schema",
        "POST /user.password: application/x-www-form-urlencoded schema",
        "POST /user.phone: application/x-www-form-urlencoded schema",
        "POST /user.email: application/xml schema",
        "POST /user.firstName: application/xml schema",
        "POST /user.lastName: application/xml schema",
        "POST /user.password: application/xml schema",
        "POST /user.phone: application/xml schema",
        "POST /user.username: application/xml schema"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: POST /pet/{petId}/uploadImage, POST /store/order, POST /user, POST /pet, PUT /pet, GET /pet/findByStatus, DELETE /pet/{petId}, GET /pet/{petId}, POST /pet/{petId}, GET /user/login, POST /user/createWithList, GET /user/{username}, PUT /user/{username}, DELETE /user/{username}, GET /pet/findByTags, GET /store/inventory, GET /store/order/{orderId}, DELETE /store/order/{orderId}, GET /user/logout; Error responses include error details schema: POST /pet/{petId}/uploadImage: 400 response, POST /pet/{petId}/uploadImage: 404 response, POST /store/order: 400 response, POST /store/order: 422 response, POST /pet: 400 response, POST /pet: 422 response, PUT /pet: 404 response, PUT /pet: 422 response, PUT /pet: 400 response, GET /pet/findByStatus: 400 response, DELETE /pet/{petId}: 400 response, GET /pet/{petId}: 404 response, GET /pet/{petId}: 400 response, POST /pet/{petId}: 400 response, GET /user/login: 400 response, GET /user/{username}: 400 response, GET /user/{username}: 404 response, PUT /user/{username}: 400 response, PUT /user/{username}: 404 response, DELETE /user/{username}: 404 response, DELETE /user/{username}: 400 response, GET /pet/findByTags: 400 response, GET /store/order/{orderId}: 400 response, GET /store/order/{orderId}: 404 response, DELETE /store/order/{orderId}: 400 response, DELETE /store/order/{orderId}: 404 response; All operations document 4xx error responses: POST /user, POST /user/createWithList, GET /store/inventory, GET /user/logout
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
        "POST /user",
        "POST /user/createWithList",
        "GET /store/inventory",
        "GET /user/logout"
      ],
      "All operations document 5xx error responses": [
        "POST /pet/{petId}/uploadImage",
        "POST /store/order",
        "POST /user",
        "POST /pet",
        "PUT /pet",
        "GET /pet/findByStatus",
        "DELETE /pet/{petId}",
        "GET /pet/{petId}",
        "POST /pet/{petId}",
        "GET /user/login",
        "POST /user/createWithList",
        "GET /user/{username}",
        "PUT /user/{username}",
        "DELETE /user/{username}",
        "GET /pet/findByTags",
        "GET /store/inventory",
        "GET /store/order/{orderId}",
        "DELETE /store/order/{orderId}",
        "GET /user/logout"
      ],
      "Error responses include error details schema": [
        "POST /pet/{petId}/uploadImage: 400 response",
        "POST /pet/{petId}/uploadImage: 404 response",
        "POST /store/order: 400 response",
        "POST /store/order: 422 response",
        "POST /pet: 400 response",
        "POST /pet: 422 response",
        "PUT /pet: 404 response",
        "PUT /pet: 422 response",
        "PUT /pet: 400 response",
        "GET /pet/findByStatus: 400 response",
        "DELETE /pet/{petId}: 400 response",
        "GET /pet/{petId}: 404 response",
        "GET /pet/{petId}: 400 response",
        "POST /pet/{petId}: 400 response",
        "GET /user/login: 400 response",
        "GET /user/{username}: 400 response",
        "GET /user/{username}: 404 response",
        "PUT /user/{username}: 400 response",
        "PUT /user/{username}: 404 response",
        "DELETE /user/{username}: 404 response",
        "DELETE /user/{username}: 400 response",
        "GET /pet/findByTags: 400 response",
        "GET /store/order/{orderId}: 400 response",
        "GET /store/order/{orderId}: 404 response",
        "DELETE /store/order/{orderId}: 400 response",
        "DELETE /store/order/{orderId}: 404 response"
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
- **Message:** Security validation failed: Global security requirements are set: No top-level security requirements defined; Operation-level security is defined: Endpoints without security: PUT /user/{username}, DELETE /user/{username}, GET /user/{username}, DELETE /store/order/{orderId}, GET /store/order/{orderId}, GET /user/login, GET /user/logout, POST /user, POST /user/createWithList, POST /store/order; API keys are properly described: apiKey schemes missing description: api_key
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
      "API keys are properly described": false,
      "Authentication headers are specified": true,
      "Global security requirements are set": false,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": false,
      "Security requirements are consistent": true,
      "Security schemes are defined": true
    },
    "messages": {
      "API keys are properly described": "apiKey schemes missing description: api_key",
      "Global security requirements are set": "No top-level security requirements defined",
      "Operation-level security is defined": "Endpoints without security: PUT /user/{username}, DELETE /user/{username}, GET /user/{username}, DELETE /store/order/{orderId}, GET /store/order/{orderId}, GET /user/login, GET /user/logout, POST /user, POST /user/createWithList, POST /store/order"
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

