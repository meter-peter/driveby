# API Validation Report

Generated: 2026-05-03T21:05:18+03:00
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
- Error Handling
- Schema
- Security
- Versioning
- Specification
- Documentation


### Failed Tags
- schema
- request
- documentation
- errors
- responses
- authentication
- lifecycle
- specification
- quality
- standards
- validation
- security
- authorization
- compatibility
- openapi
- compliance
- usability
- versioning


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "PayoutResponse": extra sibling fields: [examples]
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
      "Specification structure is valid": "invalid components: schema \"PayoutResponse\": extra sibling fields: [examples]"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /storeDetail: request body, POST /storeDetail: 500 application/json response, POST /storeDetail: 401 application/json response, POST /storeDetail: 403 application/json response, POST /storeDetail: 422 application/json response, POST /storeDetailAndSubmitThirdParty: request body, POST /storeDetailAndSubmitThirdParty: 500 application/json response, POST /storeDetailAndSubmitThirdParty: 200 application/json response, POST /storeDetailAndSubmitThirdParty: 401 application/json response, POST /storeDetailAndSubmitThirdParty: 403 application/json response, POST /storeDetailAndSubmitThirdParty: 422 application/json response, POST /submitThirdParty: request body, POST /submitThirdParty: 200 application/json response, POST /submitThirdParty: 401 application/json response, POST /submitThirdParty: 403 application/json response, POST /submitThirdParty: 422 application/json response, POST /submitThirdParty: 500 application/json response, POST /confirmThirdParty: request body, POST /confirmThirdParty: 500 application/json response, POST /confirmThirdParty: 401 application/json response, POST /confirmThirdParty: 403 application/json response, POST /confirmThirdParty: 422 application/json response, POST /declineThirdParty: request body, POST /declineThirdParty: 401 application/json response, POST /declineThirdParty: 403 application/json response, POST /declineThirdParty: 422 application/json response, POST /declineThirdParty: 500 application/json response, POST /payout: request body, POST /payout: 200 application/json response, POST /payout: 401 application/json response, POST /payout: 403 application/json response, POST /payout: 422 application/json response, POST /payout: 500 application/json response; All schemas have descriptions: FundSource, StoreDetailResponse, ModifyResponse, ResponseAdditionalDataInstallments, ResponseAdditionalData3DSecure, BankAccount, Address, ResponseAdditionalDataNetworkTokens, ResponseAdditionalDataOpi, Name, FraudCheckResultWrapper, StoreDetailRequest, Recurring, ResponseAdditionalDataSepa, Amount, FraudCheckResult, StoreDetailAndSubmitResponse, Card, PayoutResponse, StoreDetailAndSubmitRequest, SubmitRequest, ResponseAdditionalDataCommon, SubmitResponse, FraudResult, ModifyRequest, ResponseAdditionalDataCard, PayoutRequest, ResponseAdditionalDataBillingAddress, ServiceError
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
        "POST /storeDetail: request body",
        "POST /storeDetail: 500 application/json response",
        "POST /storeDetail: 401 application/json response",
        "POST /storeDetail: 403 application/json response",
        "POST /storeDetail: 422 application/json response",
        "POST /storeDetailAndSubmitThirdParty: request body",
        "POST /storeDetailAndSubmitThirdParty: 500 application/json response",
        "POST /storeDetailAndSubmitThirdParty: 200 application/json response",
        "POST /storeDetailAndSubmitThirdParty: 401 application/json response",
        "POST /storeDetailAndSubmitThirdParty: 403 application/json response",
        "POST /storeDetailAndSubmitThirdParty: 422 application/json response",
        "POST /submitThirdParty: request body",
        "POST /submitThirdParty: 200 application/json response",
        "POST /submitThirdParty: 401 application/json response",
        "POST /submitThirdParty: 403 application/json response",
        "POST /submitThirdParty: 422 application/json response",
        "POST /submitThirdParty: 500 application/json response",
        "POST /confirmThirdParty: request body",
        "POST /confirmThirdParty: 500 application/json response",
        "POST /confirmThirdParty: 401 application/json response",
        "POST /confirmThirdParty: 403 application/json response",
        "POST /confirmThirdParty: 422 application/json response",
        "POST /declineThirdParty: request body",
        "POST /declineThirdParty: 401 application/json response",
        "POST /declineThirdParty: 403 application/json response",
        "POST /declineThirdParty: 422 application/json response",
        "POST /declineThirdParty: 500 application/json response",
        "POST /payout: request body",
        "POST /payout: 200 application/json response",
        "POST /payout: 401 application/json response",
        "POST /payout: 403 application/json response",
        "POST /payout: 422 application/json response",
        "POST /payout: 500 application/json response"
      ],
      "All schemas have descriptions": [
        "FundSource",
        "StoreDetailResponse",
        "ModifyResponse",
        "ResponseAdditionalDataInstallments",
        "ResponseAdditionalData3DSecure",
        "BankAccount",
        "Address",
        "ResponseAdditionalDataNetworkTokens",
        "ResponseAdditionalDataOpi",
        "Name",
        "FraudCheckResultWrapper",
        "StoreDetailRequest",
        "Recurring",
        "ResponseAdditionalDataSepa",
        "Amount",
        "FraudCheckResult",
        "StoreDetailAndSubmitResponse",
        "Card",
        "PayoutResponse",
        "StoreDetailAndSubmitRequest",
        "SubmitRequest",
        "ResponseAdditionalDataCommon",
        "SubmitResponse",
        "FraudResult",
        "ModifyRequest",
        "ResponseAdditionalDataCard",
        "PayoutRequest",
        "ResponseAdditionalDataBillingAddress",
        "ServiceError"
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
- **Message:** Request validation issues found: All string fields have length constraints: POST /confirmThirdParty.originalReference: application/json schema, POST /confirmThirdParty.merchantAccount: application/json schema, POST /declineThirdParty.originalReference: application/json schema, POST /declineThirdParty.merchantAccount: application/json schema, POST /payout.shopperEmail: application/json schema, POST /payout.billingAddress.country: application/json schema, POST /payout.billingAddress.postalCode: application/json schema, POST /payout.billingAddress.stateOrProvince: application/json schema, POST /payout.selectedRecurringDetailReference: application/json schema, POST /payout.shopperReference: application/json schema, POST /payout.telephoneNumber: application/json schema, POST /payout.fundSource.billingAddress.country: application/json schema, POST /payout.fundSource.billingAddress.postalCode: application/json schema, POST /payout.fundSource.billingAddress.stateOrProvince: application/json schema, POST /payout.fundSource.shopperEmail: application/json schema, POST /payout.fundSource.shopperName.firstName: application/json schema, POST /payout.fundSource.shopperName.lastName: application/json schema, POST /payout.fundSource.telephoneNumber: application/json schema, POST /payout.recurring.contract: application/json schema, POST /payout.recurring.recurringDetailName: application/json schema, POST /payout.recurring.recurringExpiry: application/json schema, POST /payout.recurring.recurringFrequency: application/json schema, POST /payout.recurring.tokenService: application/json schema, POST /payout.shopperName.firstName: application/json schema, POST /payout.shopperName.lastName: application/json schema, POST /payout.merchantAccount: application/json schema, POST /payout.shopperInteraction: application/json schema, POST /payout.reference: application/json schema, POST /storeDetail.selectedBrand: application/json schema, POST /storeDetail.bank.ownerName: application/json schema, POST /storeDetail.bank.bankLocationId: application/json schema, POST /storeDetail.bank.iban: application/json schema, POST /storeDetail.bank.taxId: application/json schema, POST /storeDetail.bank.bankAccountNumber: application/json schema, POST /storeDetail.bank.bic: application/json schema, POST /storeDetail.bank.bankCity: application/json schema, POST /storeDetail.bank.bankName: application/json schema, POST /storeDetail.bank.countryCode: application/json schema, POST /storeDetail.shopperEmail: application/json schema, POST /storeDetail.socialSecurityNumber: application/json schema, POST /storeDetail.dateOfBirth: application/json schema, POST /storeDetail.merchantAccount: application/json schema, POST /storeDetail.recurring.contract: application/json schema, POST /storeDetail.recurring.recurringDetailName: application/json schema, POST /storeDetail.recurring.recurringExpiry: application/json schema, POST /storeDetail.recurring.recurringFrequency: application/json schema, POST /storeDetail.recurring.tokenService: application/json schema, POST /storeDetail.shopperName.firstName: application/json schema, POST /storeDetail.shopperName.lastName: application/json schema, POST /storeDetail.entityType: application/json schema, POST /storeDetail.telephoneNumber: application/json schema, POST /storeDetail.billingAddress.country: application/json schema, POST /storeDetail.billingAddress.postalCode: application/json schema, POST /storeDetail.billingAddress.stateOrProvince: application/json schema, POST /storeDetail.shopperReference: application/json schema, POST /storeDetailAndSubmitThirdParty.selectedBrand: application/json schema, POST /storeDetailAndSubmitThirdParty.reference: application/json schema, POST /storeDetailAndSubmitThirdParty.dateOfBirth: application/json schema, POST /storeDetailAndSubmitThirdParty.socialSecurityNumber: application/json schema, POST /storeDetailAndSubmitThirdParty.telephoneNumber: application/json schema, POST /storeDetailAndSubmitThirdParty.recurring.contract: application/json schema, POST /storeDetailAndSubmitThirdParty.recurring.recurringDetailName: application/json schema, POST /storeDetailAndSubmitThirdParty.recurring.recurringExpiry: application/json schema, POST /storeDetailAndSubmitThirdParty.recurring.recurringFrequency: application/json schema, POST /storeDetailAndSubmitThirdParty.recurring.tokenService: application/json schema, POST /storeDetailAndSubmitThirdParty.shopperStatement: application/json schema, POST /storeDetailAndSubmitThirdParty.billingAddress.country: application/json schema, POST /storeDetailAndSubmitThirdParty.billingAddress.postalCode: application/json schema, POST /storeDetailAndSubmitThirdParty.billingAddress.stateOrProvince: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.taxId: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.bankCity: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.bankName: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.bankAccountNumber: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.bic: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.countryCode: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.ownerName: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.bankLocationId: application/json schema, POST /storeDetailAndSubmitThirdParty.bank.iban: application/json schema, POST /storeDetailAndSubmitThirdParty.merchantAccount: application/json schema, POST /storeDetailAndSubmitThirdParty.shopperName.firstName: application/json schema, POST /storeDetailAndSubmitThirdParty.shopperName.lastName: application/json schema, POST /storeDetailAndSubmitThirdParty.shopperEmail: application/json schema, POST /storeDetailAndSubmitThirdParty.shopperReference: application/json schema, POST /storeDetailAndSubmitThirdParty.entityType: application/json schema, POST /submitThirdParty.shopperReference: application/json schema, POST /submitThirdParty.selectedRecurringDetailReference: application/json schema, POST /submitThirdParty.entityType: application/json schema, POST /submitThirdParty.nationality: application/json schema, POST /submitThirdParty.shopperStatement: application/json schema, POST /submitThirdParty.reference: application/json schema, POST /submitThirdParty.shopperEmail: application/json schema, POST /submitThirdParty.recurring.contract: application/json schema, POST /submitThirdParty.recurring.recurringDetailName: application/json schema, POST /submitThirdParty.recurring.recurringExpiry: application/json schema, POST /submitThirdParty.recurring.recurringFrequency: application/json schema, POST /submitThirdParty.recurring.tokenService: application/json schema, POST /submitThirdParty.shopperName.firstName: application/json schema, POST /submitThirdParty.shopperName.lastName: application/json schema, POST /submitThirdParty.socialSecurityNumber: application/json schema, POST /submitThirdParty.dateOfBirth: application/json schema, POST /submitThirdParty.merchantAccount: application/json schema; All numeric fields have min/max values: POST /payout.fraudOffset: application/json schema, POST /payout.amount.value: application/json schema, POST /storeDetail.fraudOffset: application/json schema, POST /storeDetailAndSubmitThirdParty.amount.value: application/json schema, POST /storeDetailAndSubmitThirdParty.fraudOffset: application/json schema, POST /submitThirdParty.amount.value: application/json schema, POST /submitThirdParty.fraudOffset: application/json schema
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
        "POST /payout.fraudOffset: application/json schema",
        "POST /payout.amount.value: application/json schema",
        "POST /storeDetail.fraudOffset: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.amount.value: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.fraudOffset: application/json schema",
        "POST /submitThirdParty.amount.value: application/json schema",
        "POST /submitThirdParty.fraudOffset: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /confirmThirdParty.originalReference: application/json schema",
        "POST /confirmThirdParty.merchantAccount: application/json schema",
        "POST /declineThirdParty.originalReference: application/json schema",
        "POST /declineThirdParty.merchantAccount: application/json schema",
        "POST /payout.shopperEmail: application/json schema",
        "POST /payout.billingAddress.country: application/json schema",
        "POST /payout.billingAddress.postalCode: application/json schema",
        "POST /payout.billingAddress.stateOrProvince: application/json schema",
        "POST /payout.selectedRecurringDetailReference: application/json schema",
        "POST /payout.shopperReference: application/json schema",
        "POST /payout.telephoneNumber: application/json schema",
        "POST /payout.fundSource.billingAddress.country: application/json schema",
        "POST /payout.fundSource.billingAddress.postalCode: application/json schema",
        "POST /payout.fundSource.billingAddress.stateOrProvince: application/json schema",
        "POST /payout.fundSource.shopperEmail: application/json schema",
        "POST /payout.fundSource.shopperName.firstName: application/json schema",
        "POST /payout.fundSource.shopperName.lastName: application/json schema",
        "POST /payout.fundSource.telephoneNumber: application/json schema",
        "POST /payout.recurring.contract: application/json schema",
        "POST /payout.recurring.recurringDetailName: application/json schema",
        "POST /payout.recurring.recurringExpiry: application/json schema",
        "POST /payout.recurring.recurringFrequency: application/json schema",
        "POST /payout.recurring.tokenService: application/json schema",
        "POST /payout.shopperName.firstName: application/json schema",
        "POST /payout.shopperName.lastName: application/json schema",
        "POST /payout.merchantAccount: application/json schema",
        "POST /payout.shopperInteraction: application/json schema",
        "POST /payout.reference: application/json schema",
        "POST /storeDetail.selectedBrand: application/json schema",
        "POST /storeDetail.bank.ownerName: application/json schema",
        "POST /storeDetail.bank.bankLocationId: application/json schema",
        "POST /storeDetail.bank.iban: application/json schema",
        "POST /storeDetail.bank.taxId: application/json schema",
        "POST /storeDetail.bank.bankAccountNumber: application/json schema",
        "POST /storeDetail.bank.bic: application/json schema",
        "POST /storeDetail.bank.bankCity: application/json schema",
        "POST /storeDetail.bank.bankName: application/json schema",
        "POST /storeDetail.bank.countryCode: application/json schema",
        "POST /storeDetail.shopperEmail: application/json schema",
        "POST /storeDetail.socialSecurityNumber: application/json schema",
        "POST /storeDetail.dateOfBirth: application/json schema",
        "POST /storeDetail.merchantAccount: application/json schema",
        "POST /storeDetail.recurring.contract: application/json schema",
        "POST /storeDetail.recurring.recurringDetailName: application/json schema",
        "POST /storeDetail.recurring.recurringExpiry: application/json schema",
        "POST /storeDetail.recurring.recurringFrequency: application/json schema",
        "POST /storeDetail.recurring.tokenService: application/json schema",
        "POST /storeDetail.shopperName.firstName: application/json schema",
        "POST /storeDetail.shopperName.lastName: application/json schema",
        "POST /storeDetail.entityType: application/json schema",
        "POST /storeDetail.telephoneNumber: application/json schema",
        "POST /storeDetail.billingAddress.country: application/json schema",
        "POST /storeDetail.billingAddress.postalCode: application/json schema",
        "POST /storeDetail.billingAddress.stateOrProvince: application/json schema",
        "POST /storeDetail.shopperReference: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.selectedBrand: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.reference: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.dateOfBirth: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.socialSecurityNumber: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.telephoneNumber: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.recurring.contract: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.recurring.recurringDetailName: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.recurring.recurringExpiry: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.recurring.recurringFrequency: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.recurring.tokenService: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.shopperStatement: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.billingAddress.country: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.billingAddress.postalCode: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.billingAddress.stateOrProvince: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.taxId: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.bankCity: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.bankName: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.bankAccountNumber: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.bic: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.countryCode: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.ownerName: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.bankLocationId: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.bank.iban: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.merchantAccount: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.shopperName.firstName: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.shopperName.lastName: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.shopperEmail: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.shopperReference: application/json schema",
        "POST /storeDetailAndSubmitThirdParty.entityType: application/json schema",
        "POST /submitThirdParty.shopperReference: application/json schema",
        "POST /submitThirdParty.selectedRecurringDetailReference: application/json schema",
        "POST /submitThirdParty.entityType: application/json schema",
        "POST /submitThirdParty.nationality: application/json schema",
        "POST /submitThirdParty.shopperStatement: application/json schema",
        "POST /submitThirdParty.reference: application/json schema",
        "POST /submitThirdParty.shopperEmail: application/json schema",
        "POST /submitThirdParty.recurring.contract: application/json schema",
        "POST /submitThirdParty.recurring.recurringDetailName: application/json schema",
        "POST /submitThirdParty.recurring.recurringExpiry: application/json schema",
        "POST /submitThirdParty.recurring.recurringFrequency: application/json schema",
        "POST /submitThirdParty.recurring.tokenService: application/json schema",
        "POST /submitThirdParty.shopperName.firstName: application/json schema",
        "POST /submitThirdParty.shopperName.lastName: application/json schema",
        "POST /submitThirdParty.socialSecurityNumber: application/json schema",
        "POST /submitThirdParty.dateOfBirth: application/json schema",
        "POST /submitThirdParty.merchantAccount: application/json schema"
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
- **Message:** Versioning validation failed: Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "68" does not match semver format (expected MAJOR.MINOR.PATCH); Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility
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

