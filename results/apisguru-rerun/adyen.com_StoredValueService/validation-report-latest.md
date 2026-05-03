# API Validation Report

Generated: 2026-05-03T21:05:19+03:00
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
- Versioning
- Documentation
- Error Handling
- Schema
- Security


### Failed Tags
- schema
- usability
- errors
- standards
- request
- documentation
- security
- authorization
- versioning
- compatibility
- quality
- responses
- validation
- authentication
- lifecycle


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /load: request body, POST /load: application/json request body, POST /load: 500 application/json response, POST /load: 200 application/json response, POST /load: 400 application/json response, POST /load: 401 application/json response, POST /load: 403 application/json response, POST /mergeBalance: request body, POST /mergeBalance: application/json request body, POST /mergeBalance: 401 application/json response, POST /mergeBalance: 403 application/json response, POST /mergeBalance: 500 application/json response, POST /mergeBalance: 200 application/json response, POST /mergeBalance: 400 application/json response, POST /voidTransaction: request body, POST /voidTransaction: application/json request body, POST /voidTransaction: 400 application/json response, POST /voidTransaction: 401 application/json response, POST /voidTransaction: 403 application/json response, POST /voidTransaction: 500 application/json response, POST /voidTransaction: 200 application/json response, POST /changeStatus: request body, POST /changeStatus: application/json request body, POST /changeStatus: 500 application/json response, POST /changeStatus: 200 application/json response, POST /changeStatus: 400 application/json response, POST /changeStatus: 401 application/json response, POST /changeStatus: 403 application/json response, POST /checkBalance: request body, POST /checkBalance: application/json request body, POST /checkBalance: 500 application/json response, POST /checkBalance: 200 application/json response, POST /checkBalance: 400 application/json response, POST /checkBalance: 401 application/json response, POST /checkBalance: 403 application/json response, POST /issue: request body, POST /issue: application/json request body, POST /issue: 403 application/json response, POST /issue: 500 application/json response, POST /issue: 200 application/json response, POST /issue: 400 application/json response, POST /issue: 401 application/json response; All schemas have descriptions: StoredValueBalanceMergeRequest, StoredValueStatusChangeResponse, StoredValueVoidResponse, StoredValueLoadRequest, StoredValueVoidRequest, ServiceError, StoredValueStatusChangeRequest, StoredValueBalanceCheckResponse, StoredValueBalanceMergeResponse, StoredValueIssueRequest, StoredValueIssueResponse, StoredValueLoadResponse, Amount, StoredValueBalanceCheckRequest
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
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": false
    },
    "messages": {
      "License information is provided": "License information is missing"
    },
    "missing_docs": {
      "All request/response bodies have examples": [
        "POST /load: request body",
        "POST /load: application/json request body",
        "POST /load: 500 application/json response",
        "POST /load: 200 application/json response",
        "POST /load: 400 application/json response",
        "POST /load: 401 application/json response",
        "POST /load: 403 application/json response",
        "POST /mergeBalance: request body",
        "POST /mergeBalance: application/json request body",
        "POST /mergeBalance: 401 application/json response",
        "POST /mergeBalance: 403 application/json response",
        "POST /mergeBalance: 500 application/json response",
        "POST /mergeBalance: 200 application/json response",
        "POST /mergeBalance: 400 application/json response",
        "POST /voidTransaction: request body",
        "POST /voidTransaction: application/json request body",
        "POST /voidTransaction: 400 application/json response",
        "POST /voidTransaction: 401 application/json response",
        "POST /voidTransaction: 403 application/json response",
        "POST /voidTransaction: 500 application/json response",
        "POST /voidTransaction: 200 application/json response",
        "POST /changeStatus: request body",
        "POST /changeStatus: application/json request body",
        "POST /changeStatus: 500 application/json response",
        "POST /changeStatus: 200 application/json response",
        "POST /changeStatus: 400 application/json response",
        "POST /changeStatus: 401 application/json response",
        "POST /changeStatus: 403 application/json response",
        "POST /checkBalance: request body",
        "POST /checkBalance: application/json request body",
        "POST /checkBalance: 500 application/json response",
        "POST /checkBalance: 200 application/json response",
        "POST /checkBalance: 400 application/json response",
        "POST /checkBalance: 401 application/json response",
        "POST /checkBalance: 403 application/json response",
        "POST /issue: request body",
        "POST /issue: application/json request body",
        "POST /issue: 403 application/json response",
        "POST /issue: 500 application/json response",
        "POST /issue: 200 application/json response",
        "POST /issue: 400 application/json response",
        "POST /issue: 401 application/json response"
      ],
      "All schemas have descriptions": [
        "StoredValueBalanceMergeRequest",
        "StoredValueStatusChangeResponse",
        "StoredValueVoidResponse",
        "StoredValueLoadRequest",
        "StoredValueVoidRequest",
        "ServiceError",
        "StoredValueStatusChangeRequest",
        "StoredValueBalanceCheckResponse",
        "StoredValueBalanceMergeResponse",
        "StoredValueIssueRequest",
        "StoredValueIssueResponse",
        "StoredValueLoadResponse",
        "Amount",
        "StoredValueBalanceCheckRequest"
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
- **Message:** Request validation issues found: All string fields have length constraints: POST /changeStatus.recurringDetailReference: application/json schema, POST /changeStatus.merchantAccount: application/json schema, POST /changeStatus.status: application/json schema, POST /changeStatus.reference: application/json schema, POST /changeStatus.shopperInteraction: application/json schema, POST /changeStatus.shopperReference: application/json schema, POST /checkBalance.recurringDetailReference: application/json schema, POST /checkBalance.reference: application/json schema, POST /checkBalance.shopperInteraction: application/json schema, POST /checkBalance.shopperReference: application/json schema, POST /checkBalance.merchantAccount: application/json schema, POST /issue.merchantAccount: application/json schema, POST /issue.recurringDetailReference: application/json schema, POST /issue.reference: application/json schema, POST /issue.shopperInteraction: application/json schema, POST /issue.shopperReference: application/json schema, POST /load.merchantAccount: application/json schema, POST /load.recurringDetailReference: application/json schema, POST /load.reference: application/json schema, POST /load.shopperReference: application/json schema, POST /load.loadType: application/json schema, POST /load.shopperInteraction: application/json schema, POST /mergeBalance.recurringDetailReference: application/json schema, POST /mergeBalance.shopperInteraction: application/json schema, POST /mergeBalance.reference: application/json schema, POST /mergeBalance.merchantAccount: application/json schema, POST /mergeBalance.shopperReference: application/json schema, POST /voidTransaction.merchantAccount: application/json schema, POST /voidTransaction.originalReference: application/json schema, POST /voidTransaction.reference: application/json schema, POST /voidTransaction.tenderReference: application/json schema, POST /voidTransaction.uniqueTerminalId: application/json schema; All numeric fields have min/max values: POST /changeStatus.amount.value: application/json schema, POST /checkBalance.amount.value: application/json schema, POST /issue.amount.value: application/json schema, POST /load.amount.value: application/json schema, POST /mergeBalance.amount.value: application/json schema
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
        "POST /changeStatus.amount.value: application/json schema",
        "POST /checkBalance.amount.value: application/json schema",
        "POST /issue.amount.value: application/json schema",
        "POST /load.amount.value: application/json schema",
        "POST /mergeBalance.amount.value: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /changeStatus.recurringDetailReference: application/json schema",
        "POST /changeStatus.merchantAccount: application/json schema",
        "POST /changeStatus.status: application/json schema",
        "POST /changeStatus.reference: application/json schema",
        "POST /changeStatus.shopperInteraction: application/json schema",
        "POST /changeStatus.shopperReference: application/json schema",
        "POST /checkBalance.recurringDetailReference: application/json schema",
        "POST /checkBalance.reference: application/json schema",
        "POST /checkBalance.shopperInteraction: application/json schema",
        "POST /checkBalance.shopperReference: application/json schema",
        "POST /checkBalance.merchantAccount: application/json schema",
        "POST /issue.merchantAccount: application/json schema",
        "POST /issue.recurringDetailReference: application/json schema",
        "POST /issue.reference: application/json schema",
        "POST /issue.shopperInteraction: application/json schema",
        "POST /issue.shopperReference: application/json schema",
        "POST /load.merchantAccount: application/json schema",
        "POST /load.recurringDetailReference: application/json schema",
        "POST /load.reference: application/json schema",
        "POST /load.shopperReference: application/json schema",
        "POST /load.loadType: application/json schema",
        "POST /load.shopperInteraction: application/json schema",
        "POST /mergeBalance.recurringDetailReference: application/json schema",
        "POST /mergeBalance.shopperInteraction: application/json schema",
        "POST /mergeBalance.reference: application/json schema",
        "POST /mergeBalance.merchantAccount: application/json schema",
        "POST /mergeBalance.shopperReference: application/json schema",
        "POST /voidTransaction.merchantAccount: application/json schema",
        "POST /voidTransaction.originalReference: application/json schema",
        "POST /voidTransaction.reference: application/json schema",
        "POST /voidTransaction.tenderReference: application/json schema",
        "POST /voidTransaction.uniqueTerminalId: application/json schema"
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
- **Message:** Error handling issues found: 
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
      "Common error responses are defined in components": false,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": true
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {}
  }
```

**Suggested Fix:**
Add comprehensive error response documentation including codes, messages, and consistent error schemas

---

### Security

#### P005: Security Standards (Failed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Failed
- **Message:** Security validation failed: API keys are properly described: apiKey schemes missing description: ApiKeyAuth; Global security requirements are set: No top-level security requirements defined
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
      "Operation-level security is defined": true,
      "Security requirements are consistent": true,
      "Security schemes are defined": true
    },
    "messages": {
      "API keys are properly described": "apiKey schemes missing description: ApiKeyAuth",
      "Global security requirements are set": "No top-level security requirements defined"
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
- **Message:** Versioning validation failed: Version follows semantic versioning: Version "46" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides
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
      "Version follows semantic versioning": "Version \"46\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

