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
- Documentation
- Error Handling
- Schema
- Security
- Versioning


### Failed Tags
- documentation
- quality
- usability
- responses
- schema
- validation
- request
- errors
- versioning
- security
- authentication
- standards
- lifecycle
- authorization
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
- **Message:** Documentation quality issues found: All request/response bodies have examples: GET /transactions: 403 application/json response, GET /transactions: 422 application/json response, GET /transactions: 500 application/json response, GET /transactions: 401 application/json response, GET /transactions/{id}: 500 application/json response, GET /transactions/{id}: 401 application/json response, GET /transactions/{id}: 403 application/json response, GET /transactions/{id}: 422 application/json response, POST /transfers: request body, POST /transfers: 500 application/json response, POST /transfers: 202 application/json response, POST /transfers: 401 application/json response, POST /transfers: 403 application/json response, POST /transfers: 422 application/json response; All schemas have descriptions: JSONObject, CZLocalAccountIdentification, JSONPath, MerchantData, ResourceReference, PartyIdentification-2, SELocalAccountIdentification, Address-2, InvalidField, NameLocation, RestServiceError, CounterpartyV3, SGLocalAccountIdentification, USLocalAccountIdentification, NOLocalAccountIdentification, UKLocalAccountIdentification, BankAccountV3, AULocalAccountIdentification, UltimatePartyIdentification, PLLocalAccountIdentification, PaymentInstrument, TransferInfo, DKLocalAccountIdentification, NumberAndBicAccountIdentification, Links, CALocalAccountIdentification, AdditionalBankIdentification, HULocalAccountIdentification, Transaction, TransactionSearchResponse, CounterpartyInfoV3, Link, BRLocalAccountIdentification, Amount, Transfer, IbanAccountIdentification
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
        "GET /transactions: 403 application/json response",
        "GET /transactions: 422 application/json response",
        "GET /transactions: 500 application/json response",
        "GET /transactions: 401 application/json response",
        "GET /transactions/{id}: 500 application/json response",
        "GET /transactions/{id}: 401 application/json response",
        "GET /transactions/{id}: 403 application/json response",
        "GET /transactions/{id}: 422 application/json response",
        "POST /transfers: request body",
        "POST /transfers: 500 application/json response",
        "POST /transfers: 202 application/json response",
        "POST /transfers: 401 application/json response",
        "POST /transfers: 403 application/json response",
        "POST /transfers: 422 application/json response"
      ],
      "All schemas have descriptions": [
        "JSONObject",
        "CZLocalAccountIdentification",
        "JSONPath",
        "MerchantData",
        "ResourceReference",
        "PartyIdentification-2",
        "SELocalAccountIdentification",
        "Address-2",
        "InvalidField",
        "NameLocation",
        "RestServiceError",
        "CounterpartyV3",
        "SGLocalAccountIdentification",
        "USLocalAccountIdentification",
        "NOLocalAccountIdentification",
        "UKLocalAccountIdentification",
        "BankAccountV3",
        "AULocalAccountIdentification",
        "UltimatePartyIdentification",
        "PLLocalAccountIdentification",
        "PaymentInstrument",
        "TransferInfo",
        "DKLocalAccountIdentification",
        "NumberAndBicAccountIdentification",
        "Links",
        "CALocalAccountIdentification",
        "AdditionalBankIdentification",
        "HULocalAccountIdentification",
        "Transaction",
        "TransactionSearchResponse",
        "CounterpartyInfoV3",
        "Link",
        "BRLocalAccountIdentification",
        "Amount",
        "Transfer",
        "IbanAccountIdentification"
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
- **Message:** Request validation issues found: All string fields have length constraints: GET /transactions: parameter balancePlatform, GET /transactions: parameter paymentInstrumentId, GET /transactions: parameter accountHolderId, GET /transactions: parameter balanceAccountId, GET /transactions: parameter cursor, GET /transactions: parameter createdSince, GET /transactions: parameter createdUntil, GET /transactions/{id}: parameter id, POST /transfers.counterparty.transferInstrumentId: application/json schema, POST /transfers.counterparty.balanceAccountId: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.firstName: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.fullName: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.lastName: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.reference: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.type: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.postalCode: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.stateOrProvince: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.city: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.country: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.line1: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.address.line2: application/json schema, POST /transfers.counterparty.bankAccount.accountHolder.dateOfBirth: application/json schema, POST /transfers.paymentInstrumentId: application/json schema, POST /transfers.balanceAccountId: application/json schema, POST /transfers.description: application/json schema, POST /transfers.id: application/json schema, POST /transfers.category: application/json schema, POST /transfers.priority: application/json schema, POST /transfers.ultimateParty.type: application/json schema, POST /transfers.ultimateParty.address.postalCode: application/json schema, POST /transfers.ultimateParty.address.stateOrProvince: application/json schema, POST /transfers.ultimateParty.address.city: application/json schema, POST /transfers.ultimateParty.address.country: application/json schema, POST /transfers.ultimateParty.address.line1: application/json schema, POST /transfers.ultimateParty.address.line2: application/json schema, POST /transfers.ultimateParty.dateOfBirth: application/json schema, POST /transfers.ultimateParty.firstName: application/json schema, POST /transfers.ultimateParty.fullName: application/json schema, POST /transfers.ultimateParty.lastName: application/json schema, POST /transfers.ultimateParty.reference: application/json schema; All numeric fields have min/max values: GET /transactions: parameter limit, POST /transfers.amount.value: application/json schema; All schemas specify data types: POST /transfers.counterparty.bankAccount.accountIdentification: application/json schema
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
        "GET /transactions: parameter limit",
        "POST /transfers.amount.value: application/json schema"
      ],
      "All schemas specify data types": [
        "POST /transfers.counterparty.bankAccount.accountIdentification: application/json schema"
      ],
      "All string fields have length constraints": [
        "GET /transactions: parameter balancePlatform",
        "GET /transactions: parameter paymentInstrumentId",
        "GET /transactions: parameter accountHolderId",
        "GET /transactions: parameter balanceAccountId",
        "GET /transactions: parameter cursor",
        "GET /transactions: parameter createdSince",
        "GET /transactions: parameter createdUntil",
        "GET /transactions/{id}: parameter id",
        "POST /transfers.counterparty.transferInstrumentId: application/json schema",
        "POST /transfers.counterparty.balanceAccountId: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.firstName: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.fullName: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.lastName: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.reference: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.type: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.postalCode: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.stateOrProvince: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.city: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.country: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.line1: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.address.line2: application/json schema",
        "POST /transfers.counterparty.bankAccount.accountHolder.dateOfBirth: application/json schema",
        "POST /transfers.paymentInstrumentId: application/json schema",
        "POST /transfers.balanceAccountId: application/json schema",
        "POST /transfers.description: application/json schema",
        "POST /transfers.id: application/json schema",
        "POST /transfers.category: application/json schema",
        "POST /transfers.priority: application/json schema",
        "POST /transfers.ultimateParty.type: application/json schema",
        "POST /transfers.ultimateParty.address.postalCode: application/json schema",
        "POST /transfers.ultimateParty.address.stateOrProvince: application/json schema",
        "POST /transfers.ultimateParty.address.city: application/json schema",
        "POST /transfers.ultimateParty.address.country: application/json schema",
        "POST /transfers.ultimateParty.address.line1: application/json schema",
        "POST /transfers.ultimateParty.address.line2: application/json schema",
        "POST /transfers.ultimateParty.dateOfBirth: application/json schema",
        "POST /transfers.ultimateParty.firstName: application/json schema",
        "POST /transfers.ultimateParty.fullName: application/json schema",
        "POST /transfers.ultimateParty.lastName: application/json schema",
        "POST /transfers.ultimateParty.reference: application/json schema"
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
- **Message:** Error handling issues found: Error responses include error details schema: GET /transactions/{id}: 422 response, GET /transactions/{id}: 500 response, GET /transactions/{id}: 401 response, GET /transactions/{id}: 403 response, POST /transfers: 403 response, POST /transfers: 422 response, POST /transfers: 500 response, POST /transfers: 401 response, GET /transactions: 500 response, GET /transactions: 401 response, GET /transactions: 403 response, GET /transactions: 422 response
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
      "Error responses include error details schema": false
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "Error responses include error details schema": [
        "GET /transactions/{id}: 422 response",
        "GET /transactions/{id}: 500 response",
        "GET /transactions/{id}: 401 response",
        "GET /transactions/{id}: 403 response",
        "POST /transfers: 403 response",
        "POST /transfers: 422 response",
        "POST /transfers: 500 response",
        "POST /transfers: 401 response",
        "GET /transactions: 500 response",
        "GET /transactions: 401 response",
        "GET /transactions: 403 response",
        "GET /transactions: 422 response"
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
- **Message:** Versioning validation failed: Version follows semantic versioning: Version "3" does not match semver format (expected MAJOR.MINOR.PATCH); Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides
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
      "Version follows semantic versioning": "Version \"3\" does not match semver format (expected MAJOR.MINOR.PATCH)"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

