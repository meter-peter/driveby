# API Validation Report

Generated: 2026-05-03T21:05:14+03:00
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
- Documentation
- Error Handling
- Schema
- Security
- Versioning
- Specification


### Failed Tags
- openapi
- quality
- errors
- specification
- usability
- schema
- validation
- authentication
- compliance
- request
- security
- versioning
- lifecycle
- documentation
- responses
- standards
- authorization
- compatibility


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid paths: invalid path /accountHolders/{id}: invalid operation PATCH: example requestAccountHolderCapability: Error at "/id": property "id" is missing
Schema:
  {
    "properties": {
      "balancePlatform": {
        "description": "The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.",
        "type": "string"
      },
      "capabilities": {
        "additionalProperties": {
          "$ref": "#/components/schemas/AccountHolderCapability"
        },
        "description": "Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.",
        "type": "object"
      },
      "contactDetails": {
        "$ref": "#/components/schemas/ContactDetails"
      },
      "description": {
        "description": "Your description for the account holder, maximum 300 characters.",
        "maxLength": 300,
        "type": "string"
      },
      "id": {
        "description": "The unique identifier of the account holder.",
        "readOnly": true,
        "type": "string"
      },
      "legalEntityId": {
        "description": "The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.",
        "type": "string"
      },
      "primaryBalanceAccount": {
        "description": "The ID of the account holder's primary balance account. By default, this is set to the first balance account that you create for the account holder. To assign a different balance account, send a PATCH request.",
        "type": "string"
      },
      "reference": {
        "description": "Your reference for the account holder, maximum 150 characters.",
        "maxLength": 150,
        "type": "string"
      },
      "status": {
        "description": "The status of the account holder.\n\nPossible values: \n\n * **active**: The account holder is active. This is the default status when creating an account holder. \n\n * **inactive (Deprecated)**: The account holder is temporarily inactive due to missing KYC details. You can set the account back to active by providing the missing KYC details. \n\n * **suspended**: The account holder is permanently deactivated by Adyen. This action cannot be undone. \n\n* **closed**: The account holder is permanently deactivated by you. This action cannot be undone.",
        "enum": [
          "active",
          "closed",
          "inactive",
          "suspended"
        ],
        "type": "string"
      },
      "timeZone": {
        "description": "The [time zone](https://www.iana.org/time-zones) of the account holder. For example, **Europe/Amsterdam**.\nDefaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).",
        "type": "string"
      },
      "verificationDeadlines": {
        "description": "List of verification deadlines and the capabilities that will be disallowed if verification errors are not resolved.",
        "items": {
          "$ref": "#/components/schemas/VerificationDeadline"
        },
        "readOnly": true,
        "type": "array",
        "x-addedInVersion": "2"
      }
    },
    "required": [
      "legalEntityId",
      "id"
    ],
    "type": "object"
  }

Value:
  {
    "capabilities": {
      "receivePayments": {
        "requested": true
      }
    },
    "description": "Liable account holder used for international payments and payouts",
    "legalEntityId": "LE322JV223222D5GG42KN6869",
    "reference": "S.Eller-001"
  }

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
      "Specification structure is valid": "invalid paths: invalid path /accountHolders/{id}: invalid operation PATCH: example requestAccountHolderCapability: Error at \"/id\": property \"id\" is missing\nSchema:\n  {\n    \"properties\": {\n      \"balancePlatform\": {\n        \"description\": \"The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.\",\n        \"type\": \"string\"\n      },\n      \"capabilities\": {\n        \"additionalProperties\": {\n          \"$ref\": \"#/components/schemas/AccountHolderCapability\"\n        },\n        \"description\": \"Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.\",\n        \"type\": \"object\"\n      },\n      \"contactDetails\": {\n        \"$ref\": \"#/components/schemas/ContactDetails\"\n      },\n      \"description\": {\n        \"description\": \"Your description for the account holder, maximum 300 characters.\",\n        \"maxLength\": 300,\n        \"type\": \"string\"\n      },\n      \"id\": {\n        \"description\": \"The unique identifier of the account holder.\",\n        \"readOnly\": true,\n        \"type\": \"string\"\n      },\n      \"legalEntityId\": {\n        \"description\": \"The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.\",\n        \"type\": \"string\"\n      },\n      \"primaryBalanceAccount\": {\n        \"description\": \"The ID of the account holder's primary balance account. By default, this is set to the first balance account that you create for the account holder. To assign a different balance account, send a PATCH request.\",\n        \"type\": \"string\"\n      },\n      \"reference\": {\n        \"description\": \"Your reference for the account holder, maximum 150 characters.\",\n        \"maxLength\": 150,\n        \"type\": \"string\"\n      },\n      \"status\": {\n        \"description\": \"The status of the account holder.\\n\\nPossible values: \\n\\n * **active**: The account holder is active. This is the default status when creating an account holder. \\n\\n * **inactive (Deprecated)**: The account holder is temporarily inactive due to missing KYC details. You can set the account back to active by providing the missing KYC details. \\n\\n * **suspended**: The account holder is permanently deactivated by Adyen. This action cannot be undone. \\n\\n* **closed**: The account holder is permanently deactivated by you. This action cannot be undone.\",\n        \"enum\": [\n          \"active\",\n          \"closed\",\n          \"inactive\",\n          \"suspended\"\n        ],\n        \"type\": \"string\"\n      },\n      \"timeZone\": {\n        \"description\": \"The [time zone](https://www.iana.org/time-zones) of the account holder. For example, **Europe/Amsterdam**.\\nDefaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).\",\n        \"type\": \"string\"\n      },\n      \"verificationDeadlines\": {\n        \"description\": \"List of verification deadlines and the capabilities that will be disallowed if verification errors are not resolved.\",\n        \"items\": {\n          \"$ref\": \"#/components/schemas/VerificationDeadline\"\n        },\n        \"readOnly\": true,\n        \"type\": \"array\",\n        \"x-addedInVersion\": \"2\"\n      }\n    },\n    \"required\": [\n      \"legalEntityId\",\n      \"id\"\n    ],\n    \"type\": \"object\"\n  }\n\nValue:\n  {\n    \"capabilities\": {\n      \"receivePayments\": {\n        \"requested\": true\n      }\n    },\n    \"description\": \"Liable account holder used for international payments and payouts\",\n    \"legalEntityId\": \"LE322JV223222D5GG42KN6869\",\n    \"reference\": \"S.Eller-001\"\n  }\n"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All request/response bodies have examples: PATCH /paymentInstruments/{id}: request body, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: request body, GET /paymentInstruments/{id}/reveal: 200 application/json response, PATCH /transactionRules/{transactionRuleId}: request body, POST /balanceAccounts: request body, POST /balanceAccounts: 200 application/json response, PATCH /balanceAccounts/{id}: request body, PATCH /balanceAccounts/{id}: application/json request body, PATCH /balanceAccounts/{id}: 200 application/json response, POST /paymentInstruments: request body, POST /transactionRules: request body, POST /transactionRules: 200 application/json response, GET /grantOffers/{grantOfferId}: 200 application/json response, POST /paymentInstrumentGroups: request body, POST /paymentInstrumentGroups: 200 application/json response, POST /balanceAccounts/{balanceAccountId}/sweeps: request body, GET /grantAccounts/{id}: 200 application/json response, GET /grantOffers: 200 application/json response, POST /accountHolders: request body, POST /accountHolders: 200 application/json response, POST /validateBankAccountIdentification: request body, POST /validateBankAccountIdentification: 200 application/json response, PATCH /accountHolders/{id}: request body; All schemas have descriptions: PaymentInstrumentGroupInfo, Repayment, AdditionalBankIdentification, Amount, RestServiceError, UKLocalAccountIdentification, BalanceAccountUpdateRequest, VerificationDeadline, Card, GrantOffers, JSONObject, TimeOfDay, CountriesRestriction, InvalidField, BrandVariantsRestriction, JSONPath, IbanAccountIdentification, PaginatedPaymentInstrumentsResponse, BalanceSweepConfigurationsResponse, ThresholdRepayment, BalancePlatform, PLLocalAccountIdentification, NumberAndBicAccountIdentification, TransactionRuleResponse, Void, UpdatePaymentInstrument, CZLocalAccountIdentification, MerchantAcquirerPair, PaymentInstrumentRevealInfo, CardConfiguration, MatchingTransactionsRestriction, PaymentInstrumentGroup, GrantLimit, DayOfWeekRestriction, CardInfo, SweepSchedule, BankAccountIdentificationValidationRequest, PaginatedAccountHoldersResponse, SweepConfigurationV2, PaymentInstrument, DeliveryContact, DifferentCurrenciesRestriction, Name, BalanceAccountInfo, BalanceAccount, SELocalAccountIdentification, TransactionRuleRestrictions, PaymentInstrumentInfo, MerchantNamesRestriction, ContactDetails, Address-2, NOLocalAccountIdentification, PaginatedBalanceAccountsResponse, Duration, TimeOfDayRestriction, TransactionRule, EntryModesRestriction, TransactionRuleInterval, AccountHolderInfo, MerchantsRestriction, CapitalGrantAccount, Fee, TotalAmountRestriction, Expiry, TransactionRuleEntityKey, LocalDateTime, MccsRestriction, CronSweepSchedule, Balance, CALocalAccountIdentification, PaymentInstrumentUpdateRequest, SweepCounterparty, USLocalAccountIdentification, AccountHolderCapability, PhoneNumber, ProcessingTypesRestriction, RepaymentTerm, ActiveNetworkTokensRestriction, TransactionRulesResponse, BulkAddress, GrantOffer, void, StringMatch, TransactionRuleInfo, BalanceAccountBase, AccountHolder, Authentication, Phone, HULocalAccountIdentification, InternationalTransactionRestriction, SGLocalAccountIdentification, CapitalBalance, CapabilityProblem, AULocalAccountIdentification, AccountSupportingEntityCapability, Address
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
        "PATCH /paymentInstruments/{id}: request body",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: request body",
        "GET /paymentInstruments/{id}/reveal: 200 application/json response",
        "PATCH /transactionRules/{transactionRuleId}: request body",
        "POST /balanceAccounts: request body",
        "POST /balanceAccounts: 200 application/json response",
        "PATCH /balanceAccounts/{id}: request body",
        "PATCH /balanceAccounts/{id}: application/json request body",
        "PATCH /balanceAccounts/{id}: 200 application/json response",
        "POST /paymentInstruments: request body",
        "POST /transactionRules: request body",
        "POST /transactionRules: 200 application/json response",
        "GET /grantOffers/{grantOfferId}: 200 application/json response",
        "POST /paymentInstrumentGroups: request body",
        "POST /paymentInstrumentGroups: 200 application/json response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: request body",
        "GET /grantAccounts/{id}: 200 application/json response",
        "GET /grantOffers: 200 application/json response",
        "POST /accountHolders: request body",
        "POST /accountHolders: 200 application/json response",
        "POST /validateBankAccountIdentification: request body",
        "POST /validateBankAccountIdentification: 200 application/json response",
        "PATCH /accountHolders/{id}: request body"
      ],
      "All schemas have descriptions": [
        "PaymentInstrumentGroupInfo",
        "Repayment",
        "AdditionalBankIdentification",
        "Amount",
        "RestServiceError",
        "UKLocalAccountIdentification",
        "BalanceAccountUpdateRequest",
        "VerificationDeadline",
        "Card",
        "GrantOffers",
        "JSONObject",
        "TimeOfDay",
        "CountriesRestriction",
        "InvalidField",
        "BrandVariantsRestriction",
        "JSONPath",
        "IbanAccountIdentification",
        "PaginatedPaymentInstrumentsResponse",
        "BalanceSweepConfigurationsResponse",
        "ThresholdRepayment",
        "BalancePlatform",
        "PLLocalAccountIdentification",
        "NumberAndBicAccountIdentification",
        "TransactionRuleResponse",
        "Void",
        "UpdatePaymentInstrument",
        "CZLocalAccountIdentification",
        "MerchantAcquirerPair",
        "PaymentInstrumentRevealInfo",
        "CardConfiguration",
        "MatchingTransactionsRestriction",
        "PaymentInstrumentGroup",
        "GrantLimit",
        "DayOfWeekRestriction",
        "CardInfo",
        "SweepSchedule",
        "BankAccountIdentificationValidationRequest",
        "PaginatedAccountHoldersResponse",
        "SweepConfigurationV2",
        "PaymentInstrument",
        "DeliveryContact",
        "DifferentCurrenciesRestriction",
        "Name",
        "BalanceAccountInfo",
        "BalanceAccount",
        "SELocalAccountIdentification",
        "TransactionRuleRestrictions",
        "PaymentInstrumentInfo",
        "MerchantNamesRestriction",
        "ContactDetails",
        "Address-2",
        "NOLocalAccountIdentification",
        "PaginatedBalanceAccountsResponse",
        "Duration",
        "TimeOfDayRestriction",
        "TransactionRule",
        "EntryModesRestriction",
        "TransactionRuleInterval",
        "AccountHolderInfo",
        "MerchantsRestriction",
        "CapitalGrantAccount",
        "Fee",
        "TotalAmountRestriction",
        "Expiry",
        "TransactionRuleEntityKey",
        "LocalDateTime",
        "MccsRestriction",
        "CronSweepSchedule",
        "Balance",
        "CALocalAccountIdentification",
        "PaymentInstrumentUpdateRequest",
        "SweepCounterparty",
        "USLocalAccountIdentification",
        "AccountHolderCapability",
        "PhoneNumber",
        "ProcessingTypesRestriction",
        "RepaymentTerm",
        "ActiveNetworkTokensRestriction",
        "TransactionRulesResponse",
        "BulkAddress",
        "GrantOffer",
        "void",
        "StringMatch",
        "TransactionRuleInfo",
        "BalanceAccountBase",
        "AccountHolder",
        "Authentication",
        "Phone",
        "HULocalAccountIdentification",
        "InternationalTransactionRestriction",
        "SGLocalAccountIdentification",
        "CapitalBalance",
        "CapabilityProblem",
        "AULocalAccountIdentification",
        "AccountSupportingEntityCapability",
        "Address"
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
- **Message:** Request validation issues found: All string fields have length constraints: GET /balancePlatforms/{id}/accountHolders: parameter id, GET /paymentInstruments/{id}/transactionRules: parameter id, GET /grantOffers/{grantOfferId}: parameter grantOfferId, GET /paymentInstrumentGroups/{id}/transactionRules: parameter id, PATCH /transactionRules/{transactionRuleId}: parameter transactionRuleId, PATCH /transactionRules/{transactionRuleId}.startDate: application/json schema, PATCH /transactionRules/{transactionRuleId}.aggregationLevel: application/json schema, PATCH /transactionRules/{transactionRuleId}.requestType: application/json schema, PATCH /transactionRules/{transactionRuleId}.type: application/json schema, PATCH /transactionRules/{transactionRuleId}.endDate: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.duration.unit: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.timeOfDay: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.timeZone: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.type: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.dayOfWeek: application/json schema, PATCH /transactionRules/{transactionRuleId}.outcomeType: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.totalAmount.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.dayOfWeek.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.dayOfWeek.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.differentCurrencies.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.value[].acquirerId: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.value[].merchantId: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.activeNetworkTokens.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.countries.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.countries.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.value[].operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.value[].value: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.processingTypes.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.processingTypes.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.internationalTransaction.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.matchingTransactions.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.value.endTime: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.value.startTime: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.brandVariants.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.brandVariants.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.entryModes.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.entryModes.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.mccs.operation: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.mccs.value[]: application/json schema, PATCH /transactionRules/{transactionRuleId}.status: application/json schema, PATCH /transactionRules/{transactionRuleId}.entityKey.entityType: application/json schema, PATCH /transactionRules/{transactionRuleId}.entityKey.entityReference: application/json schema, DELETE /transactionRules/{transactionRuleId}: parameter transactionRuleId, GET /transactionRules/{transactionRuleId}: parameter transactionRuleId, GET /accountHolders/{id}: parameter id, PATCH /accountHolders/{id}: parameter id, PATCH /accountHolders/{id}.id: application/json schema, PATCH /accountHolders/{id}.status: application/json schema, PATCH /accountHolders/{id}.verificationDeadlines[].capabilities[]: application/json schema, PATCH /accountHolders/{id}.verificationDeadlines[].expiresAt: application/json schema, PATCH /accountHolders/{id}.legalEntityId: application/json schema, PATCH /accountHolders/{id}.primaryBalanceAccount: application/json schema, PATCH /accountHolders/{id}.timeZone: application/json schema, PATCH /accountHolders/{id}.balancePlatform: application/json schema, PATCH /accountHolders/{id}.contactDetails.webAddress: application/json schema, PATCH /accountHolders/{id}.contactDetails.address.country: application/json schema, PATCH /accountHolders/{id}.contactDetails.address.postalCode: application/json schema, PATCH /accountHolders/{id}.contactDetails.address.stateOrProvince: application/json schema, PATCH /accountHolders/{id}.contactDetails.email: application/json schema, PATCH /accountHolders/{id}.contactDetails.phone.number: application/json schema, PATCH /accountHolders/{id}.contactDetails.phone.type: application/json schema, GET /balanceAccounts/{balanceAccountId}/sweeps: parameter balanceAccountId, POST /balanceAccounts/{balanceAccountId}/sweeps: parameter balanceAccountId, POST /balanceAccounts/{balanceAccountId}/sweeps.id: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.reason: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.type: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.category: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.priorities[]: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.status: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.merchantAccount: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.transferInstrumentId: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.balanceAccountId: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.currency: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.description: application/json schema, GET /paymentInstruments/{id}: parameter id, PATCH /paymentInstruments/{id}: parameter id, PATCH /paymentInstruments/{id}.statusComment: application/json schema, PATCH /paymentInstruments/{id}.statusReason: application/json schema, PATCH /paymentInstruments/{id}.balanceAccountId: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.envelope: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.insert: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.language: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.logoImageId: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.houseNumberOrName: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.postalCode: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.street: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.city: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.stateOrProvince: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.company: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.country: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.email: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.mobile: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.carrierImageId: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.currency: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.pinMailer: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.shipmentMethod: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.activation: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.cardImageId: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.carrier: application/json schema, PATCH /paymentInstruments/{id}.card.configuration.configurationProfileId: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.line2: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.line3: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.postalCode: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.stateOrProvince: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.city: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.country: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.address.line1: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.email: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.fullPhoneNumber: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.name.firstName: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.name.lastName: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneCountryCode: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneNumber: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneType: application/json schema, PATCH /paymentInstruments/{id}.card.deliveryContact.webAddress: application/json schema, PATCH /paymentInstruments/{id}.card.formFactor: application/json schema, PATCH /paymentInstruments/{id}.card.authentication.phone.number: application/json schema, PATCH /paymentInstruments/{id}.card.authentication.phone.type: application/json schema, PATCH /paymentInstruments/{id}.card.authentication.email: application/json schema, PATCH /paymentInstruments/{id}.card.brand: application/json schema, PATCH /paymentInstruments/{id}.card.brandVariant: application/json schema, PATCH /paymentInstruments/{id}.status: application/json schema, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.description: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.id: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.status: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.balanceAccountId: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.merchantAccount: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.transferInstrumentId: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.currency: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.category: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.priorities[]: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.reason: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.type: application/json schema, GET /paymentInstrumentGroups/{id}: parameter id, GET /balanceAccounts/{id}/paymentInstruments: parameter id, GET /balancePlatforms/{id}: parameter id, POST /paymentInstruments.issuingCountryCode: application/json schema, POST /paymentInstruments.paymentInstrumentGroupId: application/json schema, POST /paymentInstruments.statusReason: application/json schema, POST /paymentInstruments.card.brand: application/json schema, POST /paymentInstruments.card.brandVariant: application/json schema, POST /paymentInstruments.card.configuration.envelope: application/json schema, POST /paymentInstruments.card.configuration.insert: application/json schema, POST /paymentInstruments.card.configuration.language: application/json schema, POST /paymentInstruments.card.configuration.carrierImageId: application/json schema, POST /paymentInstruments.card.configuration.pinMailer: application/json schema, POST /paymentInstruments.card.configuration.configurationProfileId: application/json schema, POST /paymentInstruments.card.configuration.logoImageId: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.stateOrProvince: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.company: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.email: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.mobile: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.postalCode: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.country: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.houseNumberOrName: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.street: application/json schema, POST /paymentInstruments.card.configuration.bulkAddress.city: application/json schema, POST /paymentInstruments.card.configuration.currency: application/json schema, POST /paymentInstruments.card.configuration.shipmentMethod: application/json schema, POST /paymentInstruments.card.configuration.activation: application/json schema, POST /paymentInstruments.card.configuration.cardImageId: application/json schema, POST /paymentInstruments.card.configuration.carrier: application/json schema, POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneCountryCode: application/json schema, POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneNumber: application/json schema, POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneType: application/json schema, POST /paymentInstruments.card.deliveryContact.webAddress: application/json schema, POST /paymentInstruments.card.deliveryContact.address.country: application/json schema, POST /paymentInstruments.card.deliveryContact.address.line1: application/json schema, POST /paymentInstruments.card.deliveryContact.address.line2: application/json schema, POST /paymentInstruments.card.deliveryContact.address.line3: application/json schema, POST /paymentInstruments.card.deliveryContact.address.postalCode: application/json schema, POST /paymentInstruments.card.deliveryContact.address.stateOrProvince: application/json schema, POST /paymentInstruments.card.deliveryContact.address.city: application/json schema, POST /paymentInstruments.card.deliveryContact.email: application/json schema, POST /paymentInstruments.card.deliveryContact.fullPhoneNumber: application/json schema, POST /paymentInstruments.card.deliveryContact.name.firstName: application/json schema, POST /paymentInstruments.card.deliveryContact.name.lastName: application/json schema, POST /paymentInstruments.card.formFactor: application/json schema, POST /paymentInstruments.card.authentication.email: application/json schema, POST /paymentInstruments.card.authentication.phone.number: application/json schema, POST /paymentInstruments.card.authentication.phone.type: application/json schema, POST /paymentInstruments.balanceAccountId: application/json schema, POST /paymentInstruments.status: application/json schema, POST /paymentInstruments.type: application/json schema, POST /balanceAccounts.timeZone: application/json schema, POST /balanceAccounts.accountHolderId: application/json schema, POST /balanceAccounts.defaultCurrencyCode: application/json schema, GET /grantOffers: parameter accountHolderId, POST /paymentInstrumentGroups.txVariant: application/json schema, POST /paymentInstrumentGroups.balancePlatform: application/json schema, POST /accountHolders.timeZone: application/json schema, POST /accountHolders.balancePlatform: application/json schema, POST /accountHolders.contactDetails.phone.type: application/json schema, POST /accountHolders.contactDetails.phone.number: application/json schema, POST /accountHolders.contactDetails.webAddress: application/json schema, POST /accountHolders.contactDetails.address.stateOrProvince: application/json schema, POST /accountHolders.contactDetails.address.country: application/json schema, POST /accountHolders.contactDetails.address.postalCode: application/json schema, POST /accountHolders.contactDetails.email: application/json schema, POST /accountHolders.legalEntityId: application/json schema, GET /balanceAccounts/{id}: parameter id, PATCH /balanceAccounts/{id}: parameter id, PATCH /balanceAccounts/{id}.status: application/json schema, PATCH /balanceAccounts/{id}.timeZone: application/json schema, PATCH /balanceAccounts/{id}.accountHolderId: application/json schema, PATCH /balanceAccounts/{id}.defaultCurrencyCode: application/json schema, GET /paymentInstruments/{id}/reveal: parameter id, POST /transactionRules.requestType: application/json schema, POST /transactionRules.endDate: application/json schema, POST /transactionRules.entityKey.entityReference: application/json schema, POST /transactionRules.entityKey.entityType: application/json schema, POST /transactionRules.interval.dayOfWeek: application/json schema, POST /transactionRules.interval.duration.unit: application/json schema, POST /transactionRules.interval.timeOfDay: application/json schema, POST /transactionRules.interval.timeZone: application/json schema, POST /transactionRules.interval.type: application/json schema, POST /transactionRules.outcomeType: application/json schema, POST /transactionRules.startDate: application/json schema, POST /transactionRules.aggregationLevel: application/json schema, POST /transactionRules.type: application/json schema, POST /transactionRules.ruleRestrictions.differentCurrencies.operation: application/json schema, POST /transactionRules.ruleRestrictions.matchingTransactions.operation: application/json schema, POST /transactionRules.ruleRestrictions.merchants.operation: application/json schema, POST /transactionRules.ruleRestrictions.merchants.value[].acquirerId: application/json schema, POST /transactionRules.ruleRestrictions.merchants.value[].merchantId: application/json schema, POST /transactionRules.ruleRestrictions.timeOfDay.operation: application/json schema, POST /transactionRules.ruleRestrictions.timeOfDay.value.startTime: application/json schema, POST /transactionRules.ruleRestrictions.timeOfDay.value.endTime: application/json schema, POST /transactionRules.ruleRestrictions.mccs.value[]: application/json schema, POST /transactionRules.ruleRestrictions.mccs.operation: application/json schema, POST /transactionRules.ruleRestrictions.merchantNames.value[].operation: application/json schema, POST /transactionRules.ruleRestrictions.merchantNames.value[].value: application/json schema, POST /transactionRules.ruleRestrictions.merchantNames.operation: application/json schema, POST /transactionRules.ruleRestrictions.processingTypes.operation: application/json schema, POST /transactionRules.ruleRestrictions.processingTypes.value[]: application/json schema, POST /transactionRules.ruleRestrictions.internationalTransaction.operation: application/json schema, POST /transactionRules.ruleRestrictions.activeNetworkTokens.operation: application/json schema, POST /transactionRules.ruleRestrictions.brandVariants.value[]: application/json schema, POST /transactionRules.ruleRestrictions.brandVariants.operation: application/json schema, POST /transactionRules.ruleRestrictions.countries.value[]: application/json schema, POST /transactionRules.ruleRestrictions.countries.operation: application/json schema, POST /transactionRules.ruleRestrictions.entryModes.operation: application/json schema, POST /transactionRules.ruleRestrictions.entryModes.value[]: application/json schema, POST /transactionRules.ruleRestrictions.totalAmount.operation: application/json schema, POST /transactionRules.ruleRestrictions.dayOfWeek.operation: application/json schema, POST /transactionRules.ruleRestrictions.dayOfWeek.value[]: application/json schema, POST /transactionRules.status: application/json schema, GET /grantAccounts/{id}: parameter id, GET /accountHolders/{id}/balanceAccounts: parameter id; All numeric fields have min/max values: GET /balancePlatforms/{id}/accountHolders: parameter offset, GET /balancePlatforms/{id}/accountHolders: parameter limit, PATCH /transactionRules/{transactionRuleId}.score: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.duration.value: application/json schema, PATCH /transactionRules/{transactionRuleId}.interval.dayOfMonth: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.totalAmount.value.value: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.activeNetworkTokens.value: application/json schema, PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.matchingTransactions.value: application/json schema, GET /balanceAccounts/{balanceAccountId}/sweeps: parameter offset, GET /balanceAccounts/{balanceAccountId}/sweeps: parameter limit, POST /balanceAccounts/{balanceAccountId}/sweeps.targetAmount.value: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.triggerAmount.value: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.sweepAmount.value: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.triggerAmount.value: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.sweepAmount.value: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.targetAmount.value: application/json schema, GET /balanceAccounts/{id}/paymentInstruments: parameter offset, GET /balanceAccounts/{id}/paymentInstruments: parameter limit, POST /transactionRules.interval.duration.value: application/json schema, POST /transactionRules.interval.dayOfMonth: application/json schema, POST /transactionRules.ruleRestrictions.matchingTransactions.value: application/json schema, POST /transactionRules.ruleRestrictions.activeNetworkTokens.value: application/json schema, POST /transactionRules.ruleRestrictions.totalAmount.value.value: application/json schema, POST /transactionRules.score: application/json schema, GET /accountHolders/{id}/balanceAccounts: parameter offset, GET /accountHolders/{id}/balanceAccounts: parameter limit; All schemas specify data types: POST /validateBankAccountIdentification.accountIdentification: application/json schema, POST /balanceAccounts/{balanceAccountId}/sweeps.schedule: application/json schema, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.schedule: application/json schema
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
        "GET /balancePlatforms/{id}/accountHolders: parameter offset",
        "GET /balancePlatforms/{id}/accountHolders: parameter limit",
        "PATCH /transactionRules/{transactionRuleId}.score: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.duration.value: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.dayOfMonth: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.totalAmount.value.value: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.activeNetworkTokens.value: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.matchingTransactions.value: application/json schema",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: parameter offset",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: parameter limit",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.targetAmount.value: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.triggerAmount.value: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.sweepAmount.value: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.triggerAmount.value: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.sweepAmount.value: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.targetAmount.value: application/json schema",
        "GET /balanceAccounts/{id}/paymentInstruments: parameter offset",
        "GET /balanceAccounts/{id}/paymentInstruments: parameter limit",
        "POST /transactionRules.interval.duration.value: application/json schema",
        "POST /transactionRules.interval.dayOfMonth: application/json schema",
        "POST /transactionRules.ruleRestrictions.matchingTransactions.value: application/json schema",
        "POST /transactionRules.ruleRestrictions.activeNetworkTokens.value: application/json schema",
        "POST /transactionRules.ruleRestrictions.totalAmount.value.value: application/json schema",
        "POST /transactionRules.score: application/json schema",
        "GET /accountHolders/{id}/balanceAccounts: parameter offset",
        "GET /accountHolders/{id}/balanceAccounts: parameter limit"
      ],
      "All schemas specify data types": [
        "POST /validateBankAccountIdentification.accountIdentification: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.schedule: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.schedule: application/json schema"
      ],
      "All string fields have length constraints": [
        "GET /balancePlatforms/{id}/accountHolders: parameter id",
        "GET /paymentInstruments/{id}/transactionRules: parameter id",
        "GET /grantOffers/{grantOfferId}: parameter grantOfferId",
        "GET /paymentInstrumentGroups/{id}/transactionRules: parameter id",
        "PATCH /transactionRules/{transactionRuleId}: parameter transactionRuleId",
        "PATCH /transactionRules/{transactionRuleId}.startDate: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.aggregationLevel: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.requestType: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.type: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.endDate: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.duration.unit: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.timeOfDay: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.timeZone: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.type: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.interval.dayOfWeek: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.outcomeType: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.totalAmount.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.dayOfWeek.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.dayOfWeek.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.differentCurrencies.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.value[].acquirerId: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchants.value[].merchantId: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.activeNetworkTokens.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.countries.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.countries.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.value[].operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.value[].value: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.merchantNames.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.processingTypes.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.processingTypes.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.internationalTransaction.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.matchingTransactions.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.value.endTime: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.timeOfDay.value.startTime: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.brandVariants.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.brandVariants.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.entryModes.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.entryModes.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.mccs.operation: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.ruleRestrictions.mccs.value[]: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.status: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.entityKey.entityType: application/json schema",
        "PATCH /transactionRules/{transactionRuleId}.entityKey.entityReference: application/json schema",
        "DELETE /transactionRules/{transactionRuleId}: parameter transactionRuleId",
        "GET /transactionRules/{transactionRuleId}: parameter transactionRuleId",
        "GET /accountHolders/{id}: parameter id",
        "PATCH /accountHolders/{id}: parameter id",
        "PATCH /accountHolders/{id}.id: application/json schema",
        "PATCH /accountHolders/{id}.status: application/json schema",
        "PATCH /accountHolders/{id}.verificationDeadlines[].capabilities[]: application/json schema",
        "PATCH /accountHolders/{id}.verificationDeadlines[].expiresAt: application/json schema",
        "PATCH /accountHolders/{id}.legalEntityId: application/json schema",
        "PATCH /accountHolders/{id}.primaryBalanceAccount: application/json schema",
        "PATCH /accountHolders/{id}.timeZone: application/json schema",
        "PATCH /accountHolders/{id}.balancePlatform: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.webAddress: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.address.country: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.address.postalCode: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.address.stateOrProvince: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.email: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.phone.number: application/json schema",
        "PATCH /accountHolders/{id}.contactDetails.phone.type: application/json schema",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: parameter balanceAccountId",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: parameter balanceAccountId",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.id: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.reason: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.type: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.category: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.priorities[]: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.status: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.merchantAccount: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.transferInstrumentId: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.counterparty.balanceAccountId: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.currency: application/json schema",
        "POST /balanceAccounts/{balanceAccountId}/sweeps.description: application/json schema",
        "GET /paymentInstruments/{id}: parameter id",
        "PATCH /paymentInstruments/{id}: parameter id",
        "PATCH /paymentInstruments/{id}.statusComment: application/json schema",
        "PATCH /paymentInstruments/{id}.statusReason: application/json schema",
        "PATCH /paymentInstruments/{id}.balanceAccountId: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.envelope: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.insert: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.language: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.logoImageId: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.houseNumberOrName: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.postalCode: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.street: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.city: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.stateOrProvince: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.company: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.country: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.email: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.bulkAddress.mobile: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.carrierImageId: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.currency: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.pinMailer: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.shipmentMethod: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.activation: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.cardImageId: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.carrier: application/json schema",
        "PATCH /paymentInstruments/{id}.card.configuration.configurationProfileId: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.line2: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.line3: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.postalCode: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.stateOrProvince: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.city: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.country: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.address.line1: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.email: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.fullPhoneNumber: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.name.firstName: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.name.lastName: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneCountryCode: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneNumber: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.phoneNumber.phoneType: application/json schema",
        "PATCH /paymentInstruments/{id}.card.deliveryContact.webAddress: application/json schema",
        "PATCH /paymentInstruments/{id}.card.formFactor: application/json schema",
        "PATCH /paymentInstruments/{id}.card.authentication.phone.number: application/json schema",
        "PATCH /paymentInstruments/{id}.card.authentication.phone.type: application/json schema",
        "PATCH /paymentInstruments/{id}.card.authentication.email: application/json schema",
        "PATCH /paymentInstruments/{id}.card.brand: application/json schema",
        "PATCH /paymentInstruments/{id}.card.brandVariant: application/json schema",
        "PATCH /paymentInstruments/{id}.status: application/json schema",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter balanceAccountId",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: parameter sweepId",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.description: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.id: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.status: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.balanceAccountId: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.merchantAccount: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.counterparty.transferInstrumentId: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.currency: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.category: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.priorities[]: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.reason: application/json schema",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}.type: application/json schema",
        "GET /paymentInstrumentGroups/{id}: parameter id",
        "GET /balanceAccounts/{id}/paymentInstruments: parameter id",
        "GET /balancePlatforms/{id}: parameter id",
        "POST /paymentInstruments.issuingCountryCode: application/json schema",
        "POST /paymentInstruments.paymentInstrumentGroupId: application/json schema",
        "POST /paymentInstruments.statusReason: application/json schema",
        "POST /paymentInstruments.card.brand: application/json schema",
        "POST /paymentInstruments.card.brandVariant: application/json schema",
        "POST /paymentInstruments.card.configuration.envelope: application/json schema",
        "POST /paymentInstruments.card.configuration.insert: application/json schema",
        "POST /paymentInstruments.card.configuration.language: application/json schema",
        "POST /paymentInstruments.card.configuration.carrierImageId: application/json schema",
        "POST /paymentInstruments.card.configuration.pinMailer: application/json schema",
        "POST /paymentInstruments.card.configuration.configurationProfileId: application/json schema",
        "POST /paymentInstruments.card.configuration.logoImageId: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.stateOrProvince: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.company: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.email: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.mobile: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.postalCode: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.country: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.houseNumberOrName: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.street: application/json schema",
        "POST /paymentInstruments.card.configuration.bulkAddress.city: application/json schema",
        "POST /paymentInstruments.card.configuration.currency: application/json schema",
        "POST /paymentInstruments.card.configuration.shipmentMethod: application/json schema",
        "POST /paymentInstruments.card.configuration.activation: application/json schema",
        "POST /paymentInstruments.card.configuration.cardImageId: application/json schema",
        "POST /paymentInstruments.card.configuration.carrier: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneCountryCode: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneNumber: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.phoneNumber.phoneType: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.webAddress: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.country: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.line1: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.line2: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.line3: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.postalCode: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.stateOrProvince: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.address.city: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.email: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.fullPhoneNumber: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.name.firstName: application/json schema",
        "POST /paymentInstruments.card.deliveryContact.name.lastName: application/json schema",
        "POST /paymentInstruments.card.formFactor: application/json schema",
        "POST /paymentInstruments.card.authentication.email: application/json schema",
        "POST /paymentInstruments.card.authentication.phone.number: application/json schema",
        "POST /paymentInstruments.card.authentication.phone.type: application/json schema",
        "POST /paymentInstruments.balanceAccountId: application/json schema",
        "POST /paymentInstruments.status: application/json schema",
        "POST /paymentInstruments.type: application/json schema",
        "POST /balanceAccounts.timeZone: application/json schema",
        "POST /balanceAccounts.accountHolderId: application/json schema",
        "POST /balanceAccounts.defaultCurrencyCode: application/json schema",
        "GET /grantOffers: parameter accountHolderId",
        "POST /paymentInstrumentGroups.txVariant: application/json schema",
        "POST /paymentInstrumentGroups.balancePlatform: application/json schema",
        "POST /accountHolders.timeZone: application/json schema",
        "POST /accountHolders.balancePlatform: application/json schema",
        "POST /accountHolders.contactDetails.phone.type: application/json schema",
        "POST /accountHolders.contactDetails.phone.number: application/json schema",
        "POST /accountHolders.contactDetails.webAddress: application/json schema",
        "POST /accountHolders.contactDetails.address.stateOrProvince: application/json schema",
        "POST /accountHolders.contactDetails.address.country: application/json schema",
        "POST /accountHolders.contactDetails.address.postalCode: application/json schema",
        "POST /accountHolders.contactDetails.email: application/json schema",
        "POST /accountHolders.legalEntityId: application/json schema",
        "GET /balanceAccounts/{id}: parameter id",
        "PATCH /balanceAccounts/{id}: parameter id",
        "PATCH /balanceAccounts/{id}.status: application/json schema",
        "PATCH /balanceAccounts/{id}.timeZone: application/json schema",
        "PATCH /balanceAccounts/{id}.accountHolderId: application/json schema",
        "PATCH /balanceAccounts/{id}.defaultCurrencyCode: application/json schema",
        "GET /paymentInstruments/{id}/reveal: parameter id",
        "POST /transactionRules.requestType: application/json schema",
        "POST /transactionRules.endDate: application/json schema",
        "POST /transactionRules.entityKey.entityReference: application/json schema",
        "POST /transactionRules.entityKey.entityType: application/json schema",
        "POST /transactionRules.interval.dayOfWeek: application/json schema",
        "POST /transactionRules.interval.duration.unit: application/json schema",
        "POST /transactionRules.interval.timeOfDay: application/json schema",
        "POST /transactionRules.interval.timeZone: application/json schema",
        "POST /transactionRules.interval.type: application/json schema",
        "POST /transactionRules.outcomeType: application/json schema",
        "POST /transactionRules.startDate: application/json schema",
        "POST /transactionRules.aggregationLevel: application/json schema",
        "POST /transactionRules.type: application/json schema",
        "POST /transactionRules.ruleRestrictions.differentCurrencies.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.matchingTransactions.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchants.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchants.value[].acquirerId: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchants.value[].merchantId: application/json schema",
        "POST /transactionRules.ruleRestrictions.timeOfDay.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.timeOfDay.value.startTime: application/json schema",
        "POST /transactionRules.ruleRestrictions.timeOfDay.value.endTime: application/json schema",
        "POST /transactionRules.ruleRestrictions.mccs.value[]: application/json schema",
        "POST /transactionRules.ruleRestrictions.mccs.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchantNames.value[].operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchantNames.value[].value: application/json schema",
        "POST /transactionRules.ruleRestrictions.merchantNames.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.processingTypes.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.processingTypes.value[]: application/json schema",
        "POST /transactionRules.ruleRestrictions.internationalTransaction.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.activeNetworkTokens.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.brandVariants.value[]: application/json schema",
        "POST /transactionRules.ruleRestrictions.brandVariants.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.countries.value[]: application/json schema",
        "POST /transactionRules.ruleRestrictions.countries.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.entryModes.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.entryModes.value[]: application/json schema",
        "POST /transactionRules.ruleRestrictions.totalAmount.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.dayOfWeek.operation: application/json schema",
        "POST /transactionRules.ruleRestrictions.dayOfWeek.value[]: application/json schema",
        "POST /transactionRules.status: application/json schema",
        "GET /grantAccounts/{id}: parameter id",
        "GET /accountHolders/{id}/balanceAccounts: parameter id"
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
- **Message:** Error handling issues found: Error responses include error details schema: GET /balancePlatforms/{id}: 400 response, GET /balancePlatforms/{id}: 401 response, GET /balancePlatforms/{id}: 403 response, GET /balancePlatforms/{id}: 422 response, GET /balancePlatforms/{id}: 500 response, POST /paymentInstruments: 500 response, POST /paymentInstruments: 400 response, POST /paymentInstruments: 401 response, POST /paymentInstruments: 403 response, POST /paymentInstruments: 422 response, GET /paymentInstruments/{id}: 403 response, GET /paymentInstruments/{id}: 422 response, GET /paymentInstruments/{id}: 500 response, GET /paymentInstruments/{id}: 400 response, GET /paymentInstruments/{id}: 401 response, PATCH /paymentInstruments/{id}: 401 response, PATCH /paymentInstruments/{id}: 403 response, PATCH /paymentInstruments/{id}: 422 response, PATCH /paymentInstruments/{id}: 500 response, PATCH /paymentInstruments/{id}: 400 response, GET /accountHolders/{id}/balanceAccounts: 403 response, GET /accountHolders/{id}/balanceAccounts: 422 response, GET /accountHolders/{id}/balanceAccounts: 500 response, GET /accountHolders/{id}/balanceAccounts: 400 response, GET /accountHolders/{id}/balanceAccounts: 401 response, GET /accountHolders/{id}: 500 response, GET /accountHolders/{id}: 400 response, GET /accountHolders/{id}: 401 response, GET /accountHolders/{id}: 403 response, GET /accountHolders/{id}: 422 response, PATCH /accountHolders/{id}: 400 response, PATCH /accountHolders/{id}: 401 response, PATCH /accountHolders/{id}: 403 response, PATCH /accountHolders/{id}: 422 response, PATCH /accountHolders/{id}: 500 response, GET /grantAccounts/{id}: 400 response, GET /grantAccounts/{id}: 401 response, GET /grantAccounts/{id}: 403 response, GET /grantAccounts/{id}: 422 response, GET /grantAccounts/{id}: 500 response, GET /grantOffers/{grantOfferId}: 400 response, GET /grantOffers/{grantOfferId}: 401 response, GET /grantOffers/{grantOfferId}: 403 response, GET /grantOffers/{grantOfferId}: 422 response, GET /grantOffers/{grantOfferId}: 500 response, POST /paymentInstrumentGroups: 400 response, POST /paymentInstrumentGroups: 401 response, POST /paymentInstrumentGroups: 403 response, POST /paymentInstrumentGroups: 422 response, POST /paymentInstrumentGroups: 500 response, GET /paymentInstrumentGroups/{id}: 400 response, GET /paymentInstrumentGroups/{id}: 401 response, GET /paymentInstrumentGroups/{id}: 403 response, GET /paymentInstrumentGroups/{id}: 422 response, GET /paymentInstrumentGroups/{id}: 500 response, POST /accountHolders: 500 response, POST /accountHolders: 400 response, POST /accountHolders: 401 response, POST /accountHolders: 403 response, POST /accountHolders: 422 response, GET /balanceAccounts/{id}: 401 response, GET /balanceAccounts/{id}: 403 response, GET /balanceAccounts/{id}: 422 response, GET /balanceAccounts/{id}: 500 response, GET /balanceAccounts/{id}: 400 response, PATCH /balanceAccounts/{id}: 400 response, PATCH /balanceAccounts/{id}: 401 response, PATCH /balanceAccounts/{id}: 403 response, PATCH /balanceAccounts/{id}: 422 response, PATCH /balanceAccounts/{id}: 500 response, GET /balanceAccounts/{id}/paymentInstruments: 500 response, GET /balanceAccounts/{id}/paymentInstruments: 400 response, GET /balanceAccounts/{id}/paymentInstruments: 401 response, GET /balanceAccounts/{id}/paymentInstruments: 403 response, GET /balanceAccounts/{id}/paymentInstruments: 422 response, GET /balancePlatforms/{id}/accountHolders: 400 response, GET /balancePlatforms/{id}/accountHolders: 401 response, GET /balancePlatforms/{id}/accountHolders: 403 response, GET /balancePlatforms/{id}/accountHolders: 422 response, GET /balancePlatforms/{id}/accountHolders: 500 response, GET /paymentInstruments/{id}/reveal: 500 response, GET /paymentInstruments/{id}/reveal: 400 response, GET /paymentInstruments/{id}/reveal: 401 response, GET /paymentInstruments/{id}/reveal: 403 response, GET /paymentInstruments/{id}/reveal: 422 response, GET /paymentInstruments/{id}/transactionRules: 422 response, GET /paymentInstruments/{id}/transactionRules: 500 response, GET /paymentInstruments/{id}/transactionRules: 400 response, GET /paymentInstruments/{id}/transactionRules: 401 response, GET /paymentInstruments/{id}/transactionRules: 403 response, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response, DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response, GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response, PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response, POST /transactionRules: 400 response, POST /transactionRules: 401 response, POST /transactionRules: 403 response, POST /transactionRules: 422 response, POST /transactionRules: 500 response, DELETE /transactionRules/{transactionRuleId}: 401 response, DELETE /transactionRules/{transactionRuleId}: 403 response, DELETE /transactionRules/{transactionRuleId}: 422 response, DELETE /transactionRules/{transactionRuleId}: 500 response, DELETE /transactionRules/{transactionRuleId}: 400 response, GET /transactionRules/{transactionRuleId}: 400 response, GET /transactionRules/{transactionRuleId}: 401 response, GET /transactionRules/{transactionRuleId}: 403 response, GET /transactionRules/{transactionRuleId}: 422 response, GET /transactionRules/{transactionRuleId}: 500 response, PATCH /transactionRules/{transactionRuleId}: 400 response, PATCH /transactionRules/{transactionRuleId}: 401 response, PATCH /transactionRules/{transactionRuleId}: 403 response, PATCH /transactionRules/{transactionRuleId}: 422 response, PATCH /transactionRules/{transactionRuleId}: 500 response, POST /validateBankAccountIdentification: 403 response, POST /validateBankAccountIdentification: 422 response, POST /validateBankAccountIdentification: 500 response, POST /validateBankAccountIdentification: 401 response, POST /balanceAccounts: 422 response, POST /balanceAccounts: 500 response, POST /balanceAccounts: 400 response, POST /balanceAccounts: 401 response, POST /balanceAccounts: 403 response, GET /balanceAccounts/{balanceAccountId}/sweeps: 422 response, GET /balanceAccounts/{balanceAccountId}/sweeps: 500 response, GET /balanceAccounts/{balanceAccountId}/sweeps: 400 response, GET /balanceAccounts/{balanceAccountId}/sweeps: 401 response, GET /balanceAccounts/{balanceAccountId}/sweeps: 403 response, POST /balanceAccounts/{balanceAccountId}/sweeps: 400 response, POST /balanceAccounts/{balanceAccountId}/sweeps: 401 response, POST /balanceAccounts/{balanceAccountId}/sweeps: 403 response, POST /balanceAccounts/{balanceAccountId}/sweeps: 422 response, POST /balanceAccounts/{balanceAccountId}/sweeps: 500 response, GET /grantOffers: 401 response, GET /grantOffers: 403 response, GET /grantOffers: 422 response, GET /grantOffers: 500 response, GET /grantOffers: 400 response, GET /paymentInstrumentGroups/{id}/transactionRules: 400 response, GET /paymentInstrumentGroups/{id}/transactionRules: 401 response, GET /paymentInstrumentGroups/{id}/transactionRules: 403 response, GET /paymentInstrumentGroups/{id}/transactionRules: 422 response, GET /paymentInstrumentGroups/{id}/transactionRules: 500 response
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
        "GET /balancePlatforms/{id}: 400 response",
        "GET /balancePlatforms/{id}: 401 response",
        "GET /balancePlatforms/{id}: 403 response",
        "GET /balancePlatforms/{id}: 422 response",
        "GET /balancePlatforms/{id}: 500 response",
        "POST /paymentInstruments: 500 response",
        "POST /paymentInstruments: 400 response",
        "POST /paymentInstruments: 401 response",
        "POST /paymentInstruments: 403 response",
        "POST /paymentInstruments: 422 response",
        "GET /paymentInstruments/{id}: 403 response",
        "GET /paymentInstruments/{id}: 422 response",
        "GET /paymentInstruments/{id}: 500 response",
        "GET /paymentInstruments/{id}: 400 response",
        "GET /paymentInstruments/{id}: 401 response",
        "PATCH /paymentInstruments/{id}: 401 response",
        "PATCH /paymentInstruments/{id}: 403 response",
        "PATCH /paymentInstruments/{id}: 422 response",
        "PATCH /paymentInstruments/{id}: 500 response",
        "PATCH /paymentInstruments/{id}: 400 response",
        "GET /accountHolders/{id}/balanceAccounts: 403 response",
        "GET /accountHolders/{id}/balanceAccounts: 422 response",
        "GET /accountHolders/{id}/balanceAccounts: 500 response",
        "GET /accountHolders/{id}/balanceAccounts: 400 response",
        "GET /accountHolders/{id}/balanceAccounts: 401 response",
        "GET /accountHolders/{id}: 500 response",
        "GET /accountHolders/{id}: 400 response",
        "GET /accountHolders/{id}: 401 response",
        "GET /accountHolders/{id}: 403 response",
        "GET /accountHolders/{id}: 422 response",
        "PATCH /accountHolders/{id}: 400 response",
        "PATCH /accountHolders/{id}: 401 response",
        "PATCH /accountHolders/{id}: 403 response",
        "PATCH /accountHolders/{id}: 422 response",
        "PATCH /accountHolders/{id}: 500 response",
        "GET /grantAccounts/{id}: 400 response",
        "GET /grantAccounts/{id}: 401 response",
        "GET /grantAccounts/{id}: 403 response",
        "GET /grantAccounts/{id}: 422 response",
        "GET /grantAccounts/{id}: 500 response",
        "GET /grantOffers/{grantOfferId}: 400 response",
        "GET /grantOffers/{grantOfferId}: 401 response",
        "GET /grantOffers/{grantOfferId}: 403 response",
        "GET /grantOffers/{grantOfferId}: 422 response",
        "GET /grantOffers/{grantOfferId}: 500 response",
        "POST /paymentInstrumentGroups: 400 response",
        "POST /paymentInstrumentGroups: 401 response",
        "POST /paymentInstrumentGroups: 403 response",
        "POST /paymentInstrumentGroups: 422 response",
        "POST /paymentInstrumentGroups: 500 response",
        "GET /paymentInstrumentGroups/{id}: 400 response",
        "GET /paymentInstrumentGroups/{id}: 401 response",
        "GET /paymentInstrumentGroups/{id}: 403 response",
        "GET /paymentInstrumentGroups/{id}: 422 response",
        "GET /paymentInstrumentGroups/{id}: 500 response",
        "POST /accountHolders: 500 response",
        "POST /accountHolders: 400 response",
        "POST /accountHolders: 401 response",
        "POST /accountHolders: 403 response",
        "POST /accountHolders: 422 response",
        "GET /balanceAccounts/{id}: 401 response",
        "GET /balanceAccounts/{id}: 403 response",
        "GET /balanceAccounts/{id}: 422 response",
        "GET /balanceAccounts/{id}: 500 response",
        "GET /balanceAccounts/{id}: 400 response",
        "PATCH /balanceAccounts/{id}: 400 response",
        "PATCH /balanceAccounts/{id}: 401 response",
        "PATCH /balanceAccounts/{id}: 403 response",
        "PATCH /balanceAccounts/{id}: 422 response",
        "PATCH /balanceAccounts/{id}: 500 response",
        "GET /balanceAccounts/{id}/paymentInstruments: 500 response",
        "GET /balanceAccounts/{id}/paymentInstruments: 400 response",
        "GET /balanceAccounts/{id}/paymentInstruments: 401 response",
        "GET /balanceAccounts/{id}/paymentInstruments: 403 response",
        "GET /balanceAccounts/{id}/paymentInstruments: 422 response",
        "GET /balancePlatforms/{id}/accountHolders: 400 response",
        "GET /balancePlatforms/{id}/accountHolders: 401 response",
        "GET /balancePlatforms/{id}/accountHolders: 403 response",
        "GET /balancePlatforms/{id}/accountHolders: 422 response",
        "GET /balancePlatforms/{id}/accountHolders: 500 response",
        "GET /paymentInstruments/{id}/reveal: 500 response",
        "GET /paymentInstruments/{id}/reveal: 400 response",
        "GET /paymentInstruments/{id}/reveal: 401 response",
        "GET /paymentInstruments/{id}/reveal: 403 response",
        "GET /paymentInstruments/{id}/reveal: 422 response",
        "GET /paymentInstruments/{id}/transactionRules: 422 response",
        "GET /paymentInstruments/{id}/transactionRules: 500 response",
        "GET /paymentInstruments/{id}/transactionRules: 400 response",
        "GET /paymentInstruments/{id}/transactionRules: 401 response",
        "GET /paymentInstruments/{id}/transactionRules: 403 response",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response",
        "DELETE /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 403 response",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 422 response",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 500 response",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 400 response",
        "PATCH /balanceAccounts/{balanceAccountId}/sweeps/{sweepId}: 401 response",
        "POST /transactionRules: 400 response",
        "POST /transactionRules: 401 response",
        "POST /transactionRules: 403 response",
        "POST /transactionRules: 422 response",
        "POST /transactionRules: 500 response",
        "DELETE /transactionRules/{transactionRuleId}: 401 response",
        "DELETE /transactionRules/{transactionRuleId}: 403 response",
        "DELETE /transactionRules/{transactionRuleId}: 422 response",
        "DELETE /transactionRules/{transactionRuleId}: 500 response",
        "DELETE /transactionRules/{transactionRuleId}: 400 response",
        "GET /transactionRules/{transactionRuleId}: 400 response",
        "GET /transactionRules/{transactionRuleId}: 401 response",
        "GET /transactionRules/{transactionRuleId}: 403 response",
        "GET /transactionRules/{transactionRuleId}: 422 response",
        "GET /transactionRules/{transactionRuleId}: 500 response",
        "PATCH /transactionRules/{transactionRuleId}: 400 response",
        "PATCH /transactionRules/{transactionRuleId}: 401 response",
        "PATCH /transactionRules/{transactionRuleId}: 403 response",
        "PATCH /transactionRules/{transactionRuleId}: 422 response",
        "PATCH /transactionRules/{transactionRuleId}: 500 response",
        "POST /validateBankAccountIdentification: 403 response",
        "POST /validateBankAccountIdentification: 422 response",
        "POST /validateBankAccountIdentification: 500 response",
        "POST /validateBankAccountIdentification: 401 response",
        "POST /balanceAccounts: 422 response",
        "POST /balanceAccounts: 500 response",
        "POST /balanceAccounts: 400 response",
        "POST /balanceAccounts: 401 response",
        "POST /balanceAccounts: 403 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: 422 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: 500 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: 400 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: 401 response",
        "GET /balanceAccounts/{balanceAccountId}/sweeps: 403 response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: 400 response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: 401 response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: 403 response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: 422 response",
        "POST /balanceAccounts/{balanceAccountId}/sweeps: 500 response",
        "GET /grantOffers: 401 response",
        "GET /grantOffers: 403 response",
        "GET /grantOffers: 422 response",
        "GET /grantOffers: 500 response",
        "GET /grantOffers: 400 response",
        "GET /paymentInstrumentGroups/{id}/transactionRules: 400 response",
        "GET /paymentInstrumentGroups/{id}/transactionRules: 401 response",
        "GET /paymentInstrumentGroups/{id}/transactionRules: 403 response",
        "GET /paymentInstrumentGroups/{id}/transactionRules: 422 response",
        "GET /paymentInstrumentGroups/{id}/transactionRules: 500 response"
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
- **Message:** Versioning validation failed: Version follows semantic versioning: Version "2" does not match semver format (expected MAJOR.MINOR.PATCH); Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides
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
      "Version follows semantic versioning": "Version \"2\" does not match semver format (expected MAJOR.MINOR.PATCH)"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

