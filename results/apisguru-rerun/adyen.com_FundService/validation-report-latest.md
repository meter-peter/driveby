# API Validation Report

Generated: 2026-05-03T21:05:16+03:00
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
- Schema
- Security
- Versioning
- Documentation
- Error Handling


### Failed Tags
- compatibility
- responses
- standards
- validation
- lifecycle
- versioning
- quality
- schema
- security
- authentication
- authorization
- documentation
- usability
- errors
- request


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /accountHolderBalance: request body, POST /accountHolderBalance: 401 application/json response, POST /accountHolderBalance: 403 application/json response, POST /accountHolderBalance: 422 application/json response, POST /accountHolderBalance: 500 application/json response, POST /accountHolderBalance: 200 application/json response, POST /accountHolderBalance: 202 application/json response, POST /accountHolderBalance: 400 application/json response, POST /accountHolderTransactionList: request body, POST /accountHolderTransactionList: 422 application/json response, POST /accountHolderTransactionList: 500 application/json response, POST /accountHolderTransactionList: 200 application/json response, POST /accountHolderTransactionList: 202 application/json response, POST /accountHolderTransactionList: 400 application/json response, POST /accountHolderTransactionList: 401 application/json response, POST /accountHolderTransactionList: 403 application/json response, POST /debitAccountHolder: request body, POST /debitAccountHolder: 422 application/json response, POST /debitAccountHolder: 500 application/json response, POST /debitAccountHolder: 202 application/json response, POST /debitAccountHolder: 400 application/json response, POST /debitAccountHolder: 401 application/json response, POST /debitAccountHolder: 403 application/json response, POST /payoutAccountHolder: request body, POST /payoutAccountHolder: 400 application/json response, POST /payoutAccountHolder: 401 application/json response, POST /payoutAccountHolder: 403 application/json response, POST /payoutAccountHolder: 422 application/json response, POST /payoutAccountHolder: 500 application/json response, POST /payoutAccountHolder: 200 application/json response, POST /payoutAccountHolder: 202 application/json response, POST /refundFundsTransfer: request body, POST /refundFundsTransfer: 200 application/json response, POST /refundFundsTransfer: 202 application/json response, POST /refundFundsTransfer: 400 application/json response, POST /refundFundsTransfer: 401 application/json response, POST /refundFundsTransfer: 403 application/json response, POST /refundFundsTransfer: 422 application/json response, POST /refundFundsTransfer: 500 application/json response, POST /refundNotPaidOutTransfers: request body, POST /refundNotPaidOutTransfers: 403 application/json response, POST /refundNotPaidOutTransfers: 422 application/json response, POST /refundNotPaidOutTransfers: 500 application/json response, POST /refundNotPaidOutTransfers: 200 application/json response, POST /refundNotPaidOutTransfers: 202 application/json response, POST /refundNotPaidOutTransfers: 400 application/json response, POST /refundNotPaidOutTransfers: 401 application/json response, POST /setupBeneficiary: request body, POST /setupBeneficiary: 400 application/json response, POST /setupBeneficiary: 401 application/json response, POST /setupBeneficiary: 403 application/json response, POST /setupBeneficiary: 422 application/json response, POST /setupBeneficiary: 500 application/json response, POST /setupBeneficiary: 200 application/json response, POST /setupBeneficiary: 202 application/json response, POST /transferFunds: request body, POST /transferFunds: 403 application/json response, POST /transferFunds: 422 application/json response, POST /transferFunds: 500 application/json response, POST /transferFunds: 200 application/json response, POST /transferFunds: 202 application/json response, POST /transferFunds: 400 application/json response, POST /transferFunds: 401 application/json response; All schemas have descriptions: ErrorFieldType, BankAccountDetail, Split, TransferFundsRequest, PayoutAccountHolderResponse, RefundFundsTransferResponse, AccountDetailBalance, AccountHolderBalanceResponse, RefundFundsTransferRequest, AccountHolderBalanceRequest, AccountHolderTransactionListResponse, RefundNotPaidOutTransfersRequest, AccountHolderTransactionListRequest, SplitAmount, TransactionListForAccount, ServiceError, Amount, DebitAccountHolderResponse, DetailBalance, FieldType, Transaction, AccountTransactionList, DebitAccountHolderRequest, SetupBeneficiaryRequest, PayoutAccountHolderRequest, SetupBeneficiaryResponse, RefundNotPaidOutTransfersResponse, TransferFundsResponse
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
        "POST /accountHolderBalance: request body",
        "POST /accountHolderBalance: 401 application/json response",
        "POST /accountHolderBalance: 403 application/json response",
        "POST /accountHolderBalance: 422 application/json response",
        "POST /accountHolderBalance: 500 application/json response",
        "POST /accountHolderBalance: 200 application/json response",
        "POST /accountHolderBalance: 202 application/json response",
        "POST /accountHolderBalance: 400 application/json response",
        "POST /accountHolderTransactionList: request body",
        "POST /accountHolderTransactionList: 422 application/json response",
        "POST /accountHolderTransactionList: 500 application/json response",
        "POST /accountHolderTransactionList: 200 application/json response",
        "POST /accountHolderTransactionList: 202 application/json response",
        "POST /accountHolderTransactionList: 400 application/json response",
        "POST /accountHolderTransactionList: 401 application/json response",
        "POST /accountHolderTransactionList: 403 application/json response",
        "POST /debitAccountHolder: request body",
        "POST /debitAccountHolder: 422 application/json response",
        "POST /debitAccountHolder: 500 application/json response",
        "POST /debitAccountHolder: 202 application/json response",
        "POST /debitAccountHolder: 400 application/json response",
        "POST /debitAccountHolder: 401 application/json response",
        "POST /debitAccountHolder: 403 application/json response",
        "POST /payoutAccountHolder: request body",
        "POST /payoutAccountHolder: 400 application/json response",
        "POST /payoutAccountHolder: 401 application/json response",
        "POST /payoutAccountHolder: 403 application/json response",
        "POST /payoutAccountHolder: 422 application/json response",
        "POST /payoutAccountHolder: 500 application/json response",
        "POST /payoutAccountHolder: 200 application/json response",
        "POST /payoutAccountHolder: 202 application/json response",
        "POST /refundFundsTransfer: request body",
        "POST /refundFundsTransfer: 200 application/json response",
        "POST /refundFundsTransfer: 202 application/json response",
        "POST /refundFundsTransfer: 400 application/json response",
        "POST /refundFundsTransfer: 401 application/json response",
        "POST /refundFundsTransfer: 403 application/json response",
        "POST /refundFundsTransfer: 422 application/json response",
        "POST /refundFundsTransfer: 500 application/json response",
        "POST /refundNotPaidOutTransfers: request body",
        "POST /refundNotPaidOutTransfers: 403 application/json response",
        "POST /refundNotPaidOutTransfers: 422 application/json response",
        "POST /refundNotPaidOutTransfers: 500 application/json response",
        "POST /refundNotPaidOutTransfers: 200 application/json response",
        "POST /refundNotPaidOutTransfers: 202 application/json response",
        "POST /refundNotPaidOutTransfers: 400 application/json response",
        "POST /refundNotPaidOutTransfers: 401 application/json response",
        "POST /setupBeneficiary: request body",
        "POST /setupBeneficiary: 400 application/json response",
        "POST /setupBeneficiary: 401 application/json response",
        "POST /setupBeneficiary: 403 application/json response",
        "POST /setupBeneficiary: 422 application/json response",
        "POST /setupBeneficiary: 500 application/json response",
        "POST /setupBeneficiary: 200 application/json response",
        "POST /setupBeneficiary: 202 application/json response",
        "POST /transferFunds: request body",
        "POST /transferFunds: 403 application/json response",
        "POST /transferFunds: 422 application/json response",
        "POST /transferFunds: 500 application/json response",
        "POST /transferFunds: 200 application/json response",
        "POST /transferFunds: 202 application/json response",
        "POST /transferFunds: 400 application/json response",
        "POST /transferFunds: 401 application/json response"
      ],
      "All schemas have descriptions": [
        "ErrorFieldType",
        "BankAccountDetail",
        "Split",
        "TransferFundsRequest",
        "PayoutAccountHolderResponse",
        "RefundFundsTransferResponse",
        "AccountDetailBalance",
        "AccountHolderBalanceResponse",
        "RefundFundsTransferRequest",
        "AccountHolderBalanceRequest",
        "AccountHolderTransactionListResponse",
        "RefundNotPaidOutTransfersRequest",
        "AccountHolderTransactionListRequest",
        "SplitAmount",
        "TransactionListForAccount",
        "ServiceError",
        "Amount",
        "DebitAccountHolderResponse",
        "DetailBalance",
        "FieldType",
        "Transaction",
        "AccountTransactionList",
        "DebitAccountHolderRequest",
        "SetupBeneficiaryRequest",
        "PayoutAccountHolderRequest",
        "SetupBeneficiaryResponse",
        "RefundNotPaidOutTransfersResponse",
        "TransferFundsResponse"
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
- **Message:** Request validation issues found: All string fields have length constraints: POST /refundNotPaidOutTransfers.accountHolderCode: application/json schema, POST /refundNotPaidOutTransfers.accountCode: application/json schema, POST /setupBeneficiary.sourceAccountCode: application/json schema, POST /setupBeneficiary.destinationAccountCode: application/json schema, POST /setupBeneficiary.merchantReference: application/json schema, POST /transferFunds.destinationAccountCode: application/json schema, POST /transferFunds.merchantReference: application/json schema, POST /transferFunds.sourceAccountCode: application/json schema, POST /transferFunds.transferCode: application/json schema, POST /accountHolderBalance.accountHolderCode: application/json schema, POST /accountHolderTransactionList.accountHolderCode: application/json schema, POST /accountHolderTransactionList.transactionListsPerAccount[].accountCode: application/json schema, POST /accountHolderTransactionList.transactionStatuses[]: application/json schema, POST /debitAccountHolder.merchantAccount: application/json schema, POST /debitAccountHolder.splits[].account: application/json schema, POST /debitAccountHolder.splits[].description: application/json schema, POST /debitAccountHolder.splits[].reference: application/json schema, POST /debitAccountHolder.splits[].type: application/json schema, POST /debitAccountHolder.accountHolderCode: application/json schema, POST /debitAccountHolder.bankAccountUUID: application/json schema, POST /payoutAccountHolder.bankAccountUUID: application/json schema, POST /payoutAccountHolder.merchantReference: application/json schema, POST /payoutAccountHolder.payoutMethodCode: application/json schema, POST /payoutAccountHolder.payoutSpeed: application/json schema, POST /payoutAccountHolder.accountCode: application/json schema, POST /payoutAccountHolder.accountHolderCode: application/json schema, POST /refundFundsTransfer.merchantReference: application/json schema, POST /refundFundsTransfer.originalReference: application/json schema; All numeric fields have min/max values: POST /transferFunds.amount.value: application/json schema, POST /accountHolderTransactionList.transactionListsPerAccount[].page: application/json schema, POST /debitAccountHolder.splits[].amount.value: application/json schema, POST /debitAccountHolder.amount.value: application/json schema, POST /payoutAccountHolder.amount.value: application/json schema, POST /refundFundsTransfer.amount.value: application/json schema
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
        "POST /transferFunds.amount.value: application/json schema",
        "POST /accountHolderTransactionList.transactionListsPerAccount[].page: application/json schema",
        "POST /debitAccountHolder.splits[].amount.value: application/json schema",
        "POST /debitAccountHolder.amount.value: application/json schema",
        "POST /payoutAccountHolder.amount.value: application/json schema",
        "POST /refundFundsTransfer.amount.value: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /refundNotPaidOutTransfers.accountHolderCode: application/json schema",
        "POST /refundNotPaidOutTransfers.accountCode: application/json schema",
        "POST /setupBeneficiary.sourceAccountCode: application/json schema",
        "POST /setupBeneficiary.destinationAccountCode: application/json schema",
        "POST /setupBeneficiary.merchantReference: application/json schema",
        "POST /transferFunds.destinationAccountCode: application/json schema",
        "POST /transferFunds.merchantReference: application/json schema",
        "POST /transferFunds.sourceAccountCode: application/json schema",
        "POST /transferFunds.transferCode: application/json schema",
        "POST /accountHolderBalance.accountHolderCode: application/json schema",
        "POST /accountHolderTransactionList.accountHolderCode: application/json schema",
        "POST /accountHolderTransactionList.transactionListsPerAccount[].accountCode: application/json schema",
        "POST /accountHolderTransactionList.transactionStatuses[]: application/json schema",
        "POST /debitAccountHolder.merchantAccount: application/json schema",
        "POST /debitAccountHolder.splits[].account: application/json schema",
        "POST /debitAccountHolder.splits[].description: application/json schema",
        "POST /debitAccountHolder.splits[].reference: application/json schema",
        "POST /debitAccountHolder.splits[].type: application/json schema",
        "POST /debitAccountHolder.accountHolderCode: application/json schema",
        "POST /debitAccountHolder.bankAccountUUID: application/json schema",
        "POST /payoutAccountHolder.bankAccountUUID: application/json schema",
        "POST /payoutAccountHolder.merchantReference: application/json schema",
        "POST /payoutAccountHolder.payoutMethodCode: application/json schema",
        "POST /payoutAccountHolder.payoutSpeed: application/json schema",
        "POST /payoutAccountHolder.accountCode: application/json schema",
        "POST /payoutAccountHolder.accountHolderCode: application/json schema",
        "POST /refundFundsTransfer.merchantReference: application/json schema",
        "POST /refundFundsTransfer.originalReference: application/json schema"
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
- **Message:** Versioning validation failed: Version follows semantic versioning: Version "6" does not match semver format (expected MAJOR.MINOR.PATCH); Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides
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
      "Version follows semantic versioning": "Version \"6\" does not match semver format (expected MAJOR.MINOR.PATCH)"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

