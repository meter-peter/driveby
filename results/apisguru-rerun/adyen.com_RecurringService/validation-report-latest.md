# API Validation Report

Generated: 2026-05-03T21:05:18+03:00
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
- Error Handling
- Schema
- Security
- Versioning
- Documentation


### Failed Tags
- responses
- security
- authorization
- lifecycle
- documentation
- versioning
- usability
- errors
- standards
- validation
- authentication
- quality
- schema
- request
- compatibility


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /listRecurringDetails: request body, POST /listRecurringDetails: 403 application/json response, POST /listRecurringDetails: 422 application/json response, POST /listRecurringDetails: 500 application/json response, POST /listRecurringDetails: 200 application/json response, POST /listRecurringDetails: 400 application/json response, POST /listRecurringDetails: 401 application/json response, POST /notifyShopper: request body, POST /notifyShopper: 422 application/json response, POST /notifyShopper: 500 application/json response, POST /notifyShopper: 400 application/json response, POST /notifyShopper: 401 application/json response, POST /notifyShopper: 403 application/json response, POST /scheduleAccountUpdater: request body, POST /scheduleAccountUpdater: 400 application/json response, POST /scheduleAccountUpdater: 401 application/json response, POST /scheduleAccountUpdater: 403 application/json response, POST /scheduleAccountUpdater: 422 application/json response, POST /scheduleAccountUpdater: 500 application/json response, POST /scheduleAccountUpdater: 200 application/json response, POST /createPermit: request body, POST /createPermit: application/json request body, POST /createPermit: 403 application/json response, POST /createPermit: 422 application/json response, POST /createPermit: 500 application/json response, POST /createPermit: 200 application/json response, POST /createPermit: 400 application/json response, POST /createPermit: 401 application/json response, POST /disable: request body, POST /disable: 500 application/json response, POST /disable: 200 application/json response, POST /disable: 400 application/json response, POST /disable: 401 application/json response, POST /disable: 403 application/json response, POST /disable: 422 application/json response, POST /disablePermit: request body, POST /disablePermit: application/json request body, POST /disablePermit: 401 application/json response, POST /disablePermit: 403 application/json response, POST /disablePermit: 422 application/json response, POST /disablePermit: 500 application/json response, POST /disablePermit: 200 application/json response, POST /disablePermit: 400 application/json response; All schemas have descriptions: ScheduleAccountUpdaterResult, ServiceError, DisablePermitRequest, DisableRequest, Amount, ScheduleAccountUpdaterRequest, DisablePermitResult, Address, RecurringDetail, BankAccount, NotifyShopperResult, Name, PermitResult, Recurring, RecurringDetailsRequest, RecurringDetailWrapper, TokenDetails, CreatePermitRequest, Permit, RecurringDetailsResult, DisableResult, Card, CreatePermitResult, NotifyShopperRequest, PermitRestriction
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
        "POST /listRecurringDetails: request body",
        "POST /listRecurringDetails: 403 application/json response",
        "POST /listRecurringDetails: 422 application/json response",
        "POST /listRecurringDetails: 500 application/json response",
        "POST /listRecurringDetails: 200 application/json response",
        "POST /listRecurringDetails: 400 application/json response",
        "POST /listRecurringDetails: 401 application/json response",
        "POST /notifyShopper: request body",
        "POST /notifyShopper: 422 application/json response",
        "POST /notifyShopper: 500 application/json response",
        "POST /notifyShopper: 400 application/json response",
        "POST /notifyShopper: 401 application/json response",
        "POST /notifyShopper: 403 application/json response",
        "POST /scheduleAccountUpdater: request body",
        "POST /scheduleAccountUpdater: 400 application/json response",
        "POST /scheduleAccountUpdater: 401 application/json response",
        "POST /scheduleAccountUpdater: 403 application/json response",
        "POST /scheduleAccountUpdater: 422 application/json response",
        "POST /scheduleAccountUpdater: 500 application/json response",
        "POST /scheduleAccountUpdater: 200 application/json response",
        "POST /createPermit: request body",
        "POST /createPermit: application/json request body",
        "POST /createPermit: 403 application/json response",
        "POST /createPermit: 422 application/json response",
        "POST /createPermit: 500 application/json response",
        "POST /createPermit: 200 application/json response",
        "POST /createPermit: 400 application/json response",
        "POST /createPermit: 401 application/json response",
        "POST /disable: request body",
        "POST /disable: 500 application/json response",
        "POST /disable: 200 application/json response",
        "POST /disable: 400 application/json response",
        "POST /disable: 401 application/json response",
        "POST /disable: 403 application/json response",
        "POST /disable: 422 application/json response",
        "POST /disablePermit: request body",
        "POST /disablePermit: application/json request body",
        "POST /disablePermit: 401 application/json response",
        "POST /disablePermit: 403 application/json response",
        "POST /disablePermit: 422 application/json response",
        "POST /disablePermit: 500 application/json response",
        "POST /disablePermit: 200 application/json response",
        "POST /disablePermit: 400 application/json response"
      ],
      "All schemas have descriptions": [
        "ScheduleAccountUpdaterResult",
        "ServiceError",
        "DisablePermitRequest",
        "DisableRequest",
        "Amount",
        "ScheduleAccountUpdaterRequest",
        "DisablePermitResult",
        "Address",
        "RecurringDetail",
        "BankAccount",
        "NotifyShopperResult",
        "Name",
        "PermitResult",
        "Recurring",
        "RecurringDetailsRequest",
        "RecurringDetailWrapper",
        "TokenDetails",
        "CreatePermitRequest",
        "Permit",
        "RecurringDetailsResult",
        "DisableResult",
        "Card",
        "CreatePermitResult",
        "NotifyShopperRequest",
        "PermitRestriction"
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
- **Message:** Request validation issues found: All numeric fields have min/max values: POST /notifyShopper.amount.value: application/json schema, POST /createPermit.permits[].restriction.maxAmount.value: application/json schema, POST /createPermit.permits[].restriction.singleTransactionLimit.value: application/json schema; All string fields have length constraints: POST /disable.contract: application/json schema, POST /disable.merchantAccount: application/json schema, POST /disable.recurringDetailReference: application/json schema, POST /disable.shopperReference: application/json schema, POST /disablePermit.merchantAccount: application/json schema, POST /disablePermit.token: application/json schema, POST /listRecurringDetails.merchantAccount: application/json schema, POST /listRecurringDetails.recurring.recurringDetailName: application/json schema, POST /listRecurringDetails.recurring.recurringExpiry: application/json schema, POST /listRecurringDetails.recurring.recurringFrequency: application/json schema, POST /listRecurringDetails.recurring.tokenService: application/json schema, POST /listRecurringDetails.recurring.contract: application/json schema, POST /listRecurringDetails.shopperReference: application/json schema, POST /notifyShopper.billingDate: application/json schema, POST /notifyShopper.displayedReference: application/json schema, POST /notifyShopper.merchantAccount: application/json schema, POST /notifyShopper.shopperReference: application/json schema, POST /notifyShopper.billingSequenceNumber: application/json schema, POST /notifyShopper.reference: application/json schema, POST /notifyShopper.recurringDetailReference: application/json schema, POST /notifyShopper.storedPaymentMethodId: application/json schema, POST /scheduleAccountUpdater.merchantAccount: application/json schema, POST /scheduleAccountUpdater.reference: application/json schema, POST /scheduleAccountUpdater.selectedRecurringDetailReference: application/json schema, POST /scheduleAccountUpdater.shopperReference: application/json schema, POST /createPermit.permits[].resultKey: application/json schema, POST /createPermit.permits[].validTillDate: application/json schema, POST /createPermit.permits[].partnerId: application/json schema, POST /createPermit.permits[].profileReference: application/json schema, POST /createPermit.recurringDetailReference: application/json schema, POST /createPermit.shopperReference: application/json schema, POST /createPermit.merchantAccount: application/json schema
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
        "POST /notifyShopper.amount.value: application/json schema",
        "POST /createPermit.permits[].restriction.maxAmount.value: application/json schema",
        "POST /createPermit.permits[].restriction.singleTransactionLimit.value: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /disable.contract: application/json schema",
        "POST /disable.merchantAccount: application/json schema",
        "POST /disable.recurringDetailReference: application/json schema",
        "POST /disable.shopperReference: application/json schema",
        "POST /disablePermit.merchantAccount: application/json schema",
        "POST /disablePermit.token: application/json schema",
        "POST /listRecurringDetails.merchantAccount: application/json schema",
        "POST /listRecurringDetails.recurring.recurringDetailName: application/json schema",
        "POST /listRecurringDetails.recurring.recurringExpiry: application/json schema",
        "POST /listRecurringDetails.recurring.recurringFrequency: application/json schema",
        "POST /listRecurringDetails.recurring.tokenService: application/json schema",
        "POST /listRecurringDetails.recurring.contract: application/json schema",
        "POST /listRecurringDetails.shopperReference: application/json schema",
        "POST /notifyShopper.billingDate: application/json schema",
        "POST /notifyShopper.displayedReference: application/json schema",
        "POST /notifyShopper.merchantAccount: application/json schema",
        "POST /notifyShopper.shopperReference: application/json schema",
        "POST /notifyShopper.billingSequenceNumber: application/json schema",
        "POST /notifyShopper.reference: application/json schema",
        "POST /notifyShopper.recurringDetailReference: application/json schema",
        "POST /notifyShopper.storedPaymentMethodId: application/json schema",
        "POST /scheduleAccountUpdater.merchantAccount: application/json schema",
        "POST /scheduleAccountUpdater.reference: application/json schema",
        "POST /scheduleAccountUpdater.selectedRecurringDetailReference: application/json schema",
        "POST /scheduleAccountUpdater.shopperReference: application/json schema",
        "POST /createPermit.permits[].resultKey: application/json schema",
        "POST /createPermit.permits[].validTillDate: application/json schema",
        "POST /createPermit.permits[].partnerId: application/json schema",
        "POST /createPermit.permits[].profileReference: application/json schema",
        "POST /createPermit.recurringDetailReference: application/json schema",
        "POST /createPermit.shopperReference: application/json schema",
        "POST /createPermit.merchantAccount: application/json schema"
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
- **Message:** Security validation failed: Global security requirements are set: No top-level security requirements defined; API keys are properly described: apiKey schemes missing description: ApiKeyAuth
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
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "68" does not match semver format (expected MAJOR.MINOR.PATCH)
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
      "Versioning strategy is documented": true
    },
    "messages": {
      "Breaking changes are documented": "Info description does not reference breaking changes or a changelog",
      "Migration guides are referenced": "Info description does not reference migration or upgrade guides",
      "Version compatibility is specified": "Info description does not mention version compatibility",
      "Version follows semantic versioning": "Version \"68\" does not match semver format (expected MAJOR.MINOR.PATCH)"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

