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
- Schema
- Security
- Versioning
- Specification
- Documentation
- Error Handling


### Failed Tags
- versioning
- openapi
- documentation
- authentication
- lifecycle
- compliance
- quality
- responses
- schema
- request
- usability
- standards
- validation
- security
- compatibility
- specification
- errors
- authorization


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "AdjustAuthorisationRequest": extra sibling fields: [contentEncoding]
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
      "Specification structure is valid": "invalid components: schema \"AdjustAuthorisationRequest\": extra sibling fields: [contentEncoding]"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All schemas have descriptions: ThreeDSecureData, AdditionalDataWallets, ThreeDS2Result, FraudResult, AdditionalDataTemporaryServices, PaymentResult, ResponseAdditionalDataInstallments, SDKEphemPubKey, PlatformChargebackLogic, AdditionalDataOpi, CancelOrRefundRequest, CaptureRequest, AdditionalDataAirline, ExternalPlatform, FraudCheckResult, MerchantDevice, Recurring, CancelRequest, AdditionalDataLevel23, SplitAmount, Card, Installments, SubMerchant, AdditionalData3DSecure, Name, ApplicationInfo, CommonField, ForexQuote, ResponseAdditionalDataCommon, ResponseAdditionalDataSepa, AdditionalDataCarRental, FundSource, ResponseAdditionalDataNetworkTokens, Split, MerchantRiskIndicator, ResponseAdditionalDataOpi, AdjustAuthorisationRequest, ThreeDS1Result, AccountInfo, AdditionalDataRisk, AdditionalDataRiskStandalone, ResponseAdditionalDataCard, ResponseAdditionalDataBillingAddress, VoidPendingRefundRequest, PaymentRequest3ds2, ThreeDS2ResultResponse, PaymentRequest3d, AdditionalDataModifications, AdditionalDataSubMerchant, ThreeDS2RequestData, TechnicalCancelRequest, AuthenticationResultResponse, FraudCheckResultWrapper, BrowserInfo, AdditionalDataRetry, DonationRequest, PaymentRequest, AdditionalDataOpenInvoice, AdditionalDataCommon, FundDestination, ServiceError, AdditionalDataRatepay, Amount, ShopperInteractionDevice, Address, AuthenticationResultRequest, Phone, ModificationResult, DeviceRenderOptions, ThreeDSRequestorAuthenticationInfo, AdditionalDataLodging, ThreeDS2ResultRequest, Mandate, AcctInfo, ResponseAdditionalData3DSecure, BankAccount, ThreeDSRequestorPriorAuthenticationInfo, RefundRequest; All request/response bodies have examples: POST /cancelOrRefund: request body, POST /cancelOrRefund: 403 application/json response, POST /cancelOrRefund: 422 application/json response, POST /capture: request body, POST /capture: 403 application/json response, POST /capture: 422 application/json response, POST /refund: request body, POST /refund: 403 application/json response, POST /refund: 422 application/json response, POST /retrieve3ds2Result: request body, POST /retrieve3ds2Result: 200 application/json response, POST /retrieve3ds2Result: 403 application/json response, POST /retrieve3ds2Result: 422 application/json response, POST /technicalCancel: request body, POST /technicalCancel: 403 application/json response, POST /technicalCancel: 422 application/json response, POST /voidPendingRefund: request body, POST /voidPendingRefund: 403 application/json response, POST /voidPendingRefund: 422 application/json response, POST /authorise: request body, POST /authorise: 403 application/json response, POST /authorise: 422 application/json response, POST /authorise: 200 application/json response, POST /authorise3d: request body, POST /authorise3d: 200 application/json response, POST /authorise3d: 403 application/json response, POST /authorise3d: 422 application/json response, POST /donate: request body, POST /donate: 403 application/json response, POST /donate: 422 application/json response, POST /donate: 200 application/json response, POST /getAuthenticationResult: request body, POST /getAuthenticationResult: 200 application/json response, POST /getAuthenticationResult: 403 application/json response, POST /getAuthenticationResult: 422 application/json response, POST /adjustAuthorisation: request body, POST /adjustAuthorisation: 403 application/json response, POST /adjustAuthorisation: 422 application/json response, POST /authorise3ds2: request body, POST /authorise3ds2: 403 application/json response, POST /authorise3ds2: 422 application/json response, POST /authorise3ds2: 200 application/json response, POST /cancel: request body, POST /cancel: 403 application/json response, POST /cancel: 422 application/json response, POST /cancel: 200 application/json response
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
        "POST /cancelOrRefund: request body",
        "POST /cancelOrRefund: 403 application/json response",
        "POST /cancelOrRefund: 422 application/json response",
        "POST /capture: request body",
        "POST /capture: 403 application/json response",
        "POST /capture: 422 application/json response",
        "POST /refund: request body",
        "POST /refund: 403 application/json response",
        "POST /refund: 422 application/json response",
        "POST /retrieve3ds2Result: request body",
        "POST /retrieve3ds2Result: 200 application/json response",
        "POST /retrieve3ds2Result: 403 application/json response",
        "POST /retrieve3ds2Result: 422 application/json response",
        "POST /technicalCancel: request body",
        "POST /technicalCancel: 403 application/json response",
        "POST /technicalCancel: 422 application/json response",
        "POST /voidPendingRefund: request body",
        "POST /voidPendingRefund: 403 application/json response",
        "POST /voidPendingRefund: 422 application/json response",
        "POST /authorise: request body",
        "POST /authorise: 403 application/json response",
        "POST /authorise: 422 application/json response",
        "POST /authorise: 200 application/json response",
        "POST /authorise3d: request body",
        "POST /authorise3d: 200 application/json response",
        "POST /authorise3d: 403 application/json response",
        "POST /authorise3d: 422 application/json response",
        "POST /donate: request body",
        "POST /donate: 403 application/json response",
        "POST /donate: 422 application/json response",
        "POST /donate: 200 application/json response",
        "POST /getAuthenticationResult: request body",
        "POST /getAuthenticationResult: 200 application/json response",
        "POST /getAuthenticationResult: 403 application/json response",
        "POST /getAuthenticationResult: 422 application/json response",
        "POST /adjustAuthorisation: request body",
        "POST /adjustAuthorisation: 403 application/json response",
        "POST /adjustAuthorisation: 422 application/json response",
        "POST /authorise3ds2: request body",
        "POST /authorise3ds2: 403 application/json response",
        "POST /authorise3ds2: 422 application/json response",
        "POST /authorise3ds2: 200 application/json response",
        "POST /cancel: request body",
        "POST /cancel: 403 application/json response",
        "POST /cancel: 422 application/json response",
        "POST /cancel: 200 application/json response"
      ],
      "All schemas have descriptions": [
        "ThreeDSecureData",
        "AdditionalDataWallets",
        "ThreeDS2Result",
        "FraudResult",
        "AdditionalDataTemporaryServices",
        "PaymentResult",
        "ResponseAdditionalDataInstallments",
        "SDKEphemPubKey",
        "PlatformChargebackLogic",
        "AdditionalDataOpi",
        "CancelOrRefundRequest",
        "CaptureRequest",
        "AdditionalDataAirline",
        "ExternalPlatform",
        "FraudCheckResult",
        "MerchantDevice",
        "Recurring",
        "CancelRequest",
        "AdditionalDataLevel23",
        "SplitAmount",
        "Card",
        "Installments",
        "SubMerchant",
        "AdditionalData3DSecure",
        "Name",
        "ApplicationInfo",
        "CommonField",
        "ForexQuote",
        "ResponseAdditionalDataCommon",
        "ResponseAdditionalDataSepa",
        "AdditionalDataCarRental",
        "FundSource",
        "ResponseAdditionalDataNetworkTokens",
        "Split",
        "MerchantRiskIndicator",
        "ResponseAdditionalDataOpi",
        "AdjustAuthorisationRequest",
        "ThreeDS1Result",
        "AccountInfo",
        "AdditionalDataRisk",
        "AdditionalDataRiskStandalone",
        "ResponseAdditionalDataCard",
        "ResponseAdditionalDataBillingAddress",
        "VoidPendingRefundRequest",
        "PaymentRequest3ds2",
        "ThreeDS2ResultResponse",
        "PaymentRequest3d",
        "AdditionalDataModifications",
        "AdditionalDataSubMerchant",
        "ThreeDS2RequestData",
        "TechnicalCancelRequest",
        "AuthenticationResultResponse",
        "FraudCheckResultWrapper",
        "BrowserInfo",
        "AdditionalDataRetry",
        "DonationRequest",
        "PaymentRequest",
        "AdditionalDataOpenInvoice",
        "AdditionalDataCommon",
        "FundDestination",
        "ServiceError",
        "AdditionalDataRatepay",
        "Amount",
        "ShopperInteractionDevice",
        "Address",
        "AuthenticationResultRequest",
        "Phone",
        "ModificationResult",
        "DeviceRenderOptions",
        "ThreeDSRequestorAuthenticationInfo",
        "AdditionalDataLodging",
        "ThreeDS2ResultRequest",
        "Mandate",
        "AcctInfo",
        "ResponseAdditionalData3DSecure",
        "BankAccount",
        "ThreeDSRequestorPriorAuthenticationInfo",
        "RefundRequest"
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
- **Message:** Request validation issues found: All string fields have length constraints: POST /authorise3ds2.installments.plan: application/json schema, POST /authorise3ds2.selectedBrand: application/json schema, POST /authorise3ds2.billingAddress.country: application/json schema, POST /authorise3ds2.billingAddress.postalCode: application/json schema, POST /authorise3ds2.billingAddress.stateOrProvince: application/json schema, POST /authorise3ds2.merchantAccount: application/json schema, POST /authorise3ds2.shopperEmail: application/json schema, POST /authorise3ds2.accountInfo.paymentAccountIndicator: application/json schema, POST /authorise3ds2.accountInfo.accountChangeDate: application/json schema, POST /authorise3ds2.accountInfo.passwordChangeIndicator: application/json schema, POST /authorise3ds2.accountInfo.workPhone: application/json schema, POST /authorise3ds2.accountInfo.passwordChangeDate: application/json schema, POST /authorise3ds2.accountInfo.mobilePhone: application/json schema, POST /authorise3ds2.accountInfo.paymentAccountAge: application/json schema, POST /authorise3ds2.accountInfo.accountType: application/json schema, POST /authorise3ds2.accountInfo.accountCreationDate: application/json schema, POST /authorise3ds2.accountInfo.deliveryAddressUsageDate: application/json schema, POST /authorise3ds2.accountInfo.deliveryAddressUsageIndicator: application/json schema, POST /authorise3ds2.accountInfo.accountAgeIndicator: application/json schema, POST /authorise3ds2.accountInfo.accountChangeIndicator: application/json schema, POST /authorise3ds2.accountInfo.homePhone: application/json schema, POST /authorise3ds2.mcc: application/json schema, POST /authorise3ds2.threeDS2Result.transStatus: application/json schema, POST /authorise3ds2.threeDS2Result.authenticationValue: application/json schema, POST /authorise3ds2.threeDS2Result.messageVersion: application/json schema, POST /authorise3ds2.threeDS2Result.threeDSServerTransID: application/json schema, POST /authorise3ds2.threeDS2Result.challengeCancel: application/json schema, POST /authorise3ds2.threeDS2Result.challengeIndicator: application/json schema, POST /authorise3ds2.threeDS2Result.transStatusReason: application/json schema, POST /authorise3ds2.threeDS2Result.cavvAlgorithm: application/json schema, POST /authorise3ds2.threeDS2Result.dsTransID: application/json schema, POST /authorise3ds2.threeDS2Result.exemptionIndicator: application/json schema, POST /authorise3ds2.threeDS2Result.whiteListStatus: application/json schema, POST /authorise3ds2.threeDS2Result.eci: application/json schema, POST /authorise3ds2.threeDS2Result.riskScore: application/json schema, POST /authorise3ds2.threeDS2Result.timestamp: application/json schema, POST /authorise3ds2.merchantRiskIndicator.giftCardCurr: application/json schema, POST /authorise3ds2.merchantRiskIndicator.preOrderPurchaseInd: application/json schema, POST /authorise3ds2.merchantRiskIndicator.reorderItemsInd: application/json schema, POST /authorise3ds2.merchantRiskIndicator.deliveryAddressIndicator: application/json schema, POST /authorise3ds2.merchantRiskIndicator.preOrderDate: application/json schema, POST /authorise3ds2.merchantRiskIndicator.deliveryEmail: application/json schema, POST /authorise3ds2.merchantRiskIndicator.shipIndicator: application/json schema, POST /authorise3ds2.merchantRiskIndicator.deliveryTimeframe: application/json schema, POST /authorise3ds2.applicationInfo.adyenLibrary.version: application/json schema, POST /authorise3ds2.applicationInfo.adyenLibrary.name: application/json schema, POST /authorise3ds2.applicationInfo.adyenPaymentSource.name: application/json schema, POST /authorise3ds2.applicationInfo.adyenPaymentSource.version: application/json schema, POST /authorise3ds2.applicationInfo.externalPlatform.version: application/json schema, POST /authorise3ds2.applicationInfo.externalPlatform.integrator: application/json schema, POST /authorise3ds2.applicationInfo.externalPlatform.name: application/json schema, POST /authorise3ds2.applicationInfo.merchantApplication.version: application/json schema, POST /authorise3ds2.applicationInfo.merchantApplication.name: application/json schema, POST /authorise3ds2.applicationInfo.merchantDevice.os: application/json schema, POST /authorise3ds2.applicationInfo.merchantDevice.osVersion: application/json schema, POST /authorise3ds2.applicationInfo.merchantDevice.reference: application/json schema, POST /authorise3ds2.applicationInfo.shopperInteractionDevice.os: application/json schema, POST /authorise3ds2.applicationInfo.shopperInteractionDevice.osVersion: application/json schema, POST /authorise3ds2.applicationInfo.shopperInteractionDevice.locale: application/json schema, POST /authorise3ds2.telephoneNumber: application/json schema, POST /authorise3ds2.socialSecurityNumber: application/json schema, POST /authorise3ds2.reference: application/json schema, POST /authorise3ds2.dccQuote.type: application/json schema, POST /authorise3ds2.dccQuote.accountType: application/json schema, POST /authorise3ds2.dccQuote.reference: application/json schema, POST /authorise3ds2.dccQuote.signature: application/json schema, POST /authorise3ds2.dccQuote.source: application/json schema, POST /authorise3ds2.dccQuote.validTill: application/json schema, POST /authorise3ds2.dccQuote.account: application/json schema, POST /authorise3ds2.deliveryAddress.country: application/json schema, POST /authorise3ds2.deliveryAddress.postalCode: application/json schema, POST /authorise3ds2.deliveryAddress.stateOrProvince: application/json schema, POST /authorise3ds2.splits[].reference: application/json schema, POST /authorise3ds2.splits[].type: application/json schema, POST /authorise3ds2.splits[].account: application/json schema, POST /authorise3ds2.splits[].description: application/json schema, POST /authorise3ds2.shopperReference: application/json schema, POST /authorise3ds2.deliveryDate: application/json schema, POST /authorise3ds2.dateOfBirth: application/json schema, POST /authorise3ds2.shopperInteraction: application/json schema, POST /authorise3ds2.orderReference: application/json schema, POST /authorise3ds2.sessionId: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkReferenceNumber: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccChange: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema, POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccString: application/json schema, POST /authorise3ds2.threeDS2RequestData.transactionType: application/json schema, POST /authorise3ds2.threeDS2RequestData.acquirerBIN: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema, POST /authorise3ds2.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkAppID: application/json schema, POST /authorise3ds2.threeDS2RequestData.messageVersion: application/json schema, POST /authorise3ds2.threeDS2RequestData.mcc: application/json schema, POST /authorise3ds2.threeDS2RequestData.whiteListStatus: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorID: application/json schema, POST /authorise3ds2.threeDS2RequestData.challengeIndicator: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkVersion: application/json schema, POST /authorise3ds2.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema, POST /authorise3ds2.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkEncData: application/json schema, POST /authorise3ds2.threeDS2RequestData.recurringExpiry: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.x: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.y: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema, POST /authorise3ds2.threeDS2RequestData.acquirerMerchantID: application/json schema, POST /authorise3ds2.threeDS2RequestData.deviceChannel: application/json schema, POST /authorise3ds2.threeDS2RequestData.notificationURL: application/json schema, POST /authorise3ds2.threeDS2RequestData.merchantName: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorURL: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSCompInd: application/json schema, POST /authorise3ds2.threeDS2RequestData.threeDSRequestorName: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkTransID: application/json schema, POST /authorise3ds2.selectedRecurringDetailReference: application/json schema, POST /authorise3ds2.browserInfo.userAgent: application/json schema, POST /authorise3ds2.browserInfo.language: application/json schema, POST /authorise3ds2.browserInfo.acceptHeader: application/json schema, POST /authorise3ds2.merchantOrderReference: application/json schema, POST /authorise3ds2.recurringProcessingModel: application/json schema, POST /authorise3ds2.shopperName.lastName: application/json schema, POST /authorise3ds2.shopperName.firstName: application/json schema, POST /authorise3ds2.shopperStatement: application/json schema, POST /authorise3ds2.recurring.recurringDetailName: application/json schema, POST /authorise3ds2.recurring.recurringExpiry: application/json schema, POST /authorise3ds2.recurring.recurringFrequency: application/json schema, POST /authorise3ds2.recurring.tokenService: application/json schema, POST /authorise3ds2.recurring.contract: application/json schema, POST /authorise3ds2.threeDS2Token: application/json schema, POST /authorise3ds2.shopperLocale: application/json schema, POST /authorise3ds2.shopperIP: application/json schema, POST /voidPendingRefund.originalMerchantReference: application/json schema, POST /voidPendingRefund.originalReference: application/json schema, POST /voidPendingRefund.platformChargebackLogic.behavior: application/json schema, POST /voidPendingRefund.platformChargebackLogic.costAllocationAccount: application/json schema, POST /voidPendingRefund.platformChargebackLogic.targetAccount: application/json schema, POST /voidPendingRefund.reference: application/json schema, POST /voidPendingRefund.merchantAccount: application/json schema, POST /voidPendingRefund.splits[].type: application/json schema, POST /voidPendingRefund.splits[].account: application/json schema, POST /voidPendingRefund.splits[].description: application/json schema, POST /voidPendingRefund.splits[].reference: application/json schema, POST /voidPendingRefund.tenderReference: application/json schema, POST /voidPendingRefund.uniqueTerminalId: application/json schema, POST /voidPendingRefund.mpiData.cavv: application/json schema, POST /voidPendingRefund.mpiData.cavvAlgorithm: application/json schema, POST /voidPendingRefund.mpiData.directoryResponse: application/json schema, POST /voidPendingRefund.mpiData.xid: application/json schema, POST /voidPendingRefund.mpiData.challengeCancel: application/json schema, POST /voidPendingRefund.mpiData.eci: application/json schema, POST /voidPendingRefund.mpiData.transStatusReason: application/json schema, POST /voidPendingRefund.mpiData.authenticationResponse: application/json schema, POST /voidPendingRefund.mpiData.riskScore: application/json schema, POST /voidPendingRefund.mpiData.dsTransID: application/json schema, POST /voidPendingRefund.mpiData.threeDSVersion: application/json schema, POST /voidPendingRefund.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /authorise.bankAccount.bankName: application/json schema, POST /authorise.bankAccount.bic: application/json schema, POST /authorise.bankAccount.ownerName: application/json schema, POST /authorise.bankAccount.countryCode: application/json schema, POST /authorise.bankAccount.taxId: application/json schema, POST /authorise.bankAccount.bankAccountNumber: application/json schema, POST /authorise.bankAccount.bankCity: application/json schema, POST /authorise.bankAccount.bankLocationId: application/json schema, POST /authorise.bankAccount.iban: application/json schema, POST /authorise.dateOfBirth: application/json schema, POST /authorise.billingAddress.postalCode: application/json schema, POST /authorise.billingAddress.stateOrProvince: application/json schema, POST /authorise.billingAddress.country: application/json schema, POST /authorise.applicationInfo.adyenLibrary.name: application/json schema, POST /authorise.applicationInfo.adyenLibrary.version: application/json schema, POST /authorise.applicationInfo.adyenPaymentSource.name: application/json schema, POST /authorise.applicationInfo.adyenPaymentSource.version: application/json schema, POST /authorise.applicationInfo.externalPlatform.integrator: application/json schema, POST /authorise.applicationInfo.externalPlatform.name: application/json schema, POST /authorise.applicationInfo.externalPlatform.version: application/json schema, POST /authorise.applicationInfo.merchantApplication.name: application/json schema, POST /authorise.applicationInfo.merchantApplication.version: application/json schema, POST /authorise.applicationInfo.merchantDevice.os: application/json schema, POST /authorise.applicationInfo.merchantDevice.osVersion: application/json schema, POST /authorise.applicationInfo.merchantDevice.reference: application/json schema, POST /authorise.applicationInfo.shopperInteractionDevice.locale: application/json schema, POST /authorise.applicationInfo.shopperInteractionDevice.os: application/json schema, POST /authorise.applicationInfo.shopperInteractionDevice.osVersion: application/json schema, POST /authorise.merchantOrderReference: application/json schema, POST /authorise.orderReference: application/json schema, POST /authorise.fundSource.billingAddress.country: application/json schema, POST /authorise.fundSource.billingAddress.postalCode: application/json schema, POST /authorise.fundSource.billingAddress.stateOrProvince: application/json schema, POST /authorise.fundSource.shopperEmail: application/json schema, POST /authorise.fundSource.shopperName.firstName: application/json schema, POST /authorise.fundSource.shopperName.lastName: application/json schema, POST /authorise.fundSource.telephoneNumber: application/json schema, POST /authorise.platformChargebackLogic.costAllocationAccount: application/json schema, POST /authorise.platformChargebackLogic.targetAccount: application/json schema, POST /authorise.platformChargebackLogic.behavior: application/json schema, POST /authorise.deliveryDate: application/json schema, POST /authorise.merchantRiskIndicator.deliveryEmail: application/json schema, POST /authorise.merchantRiskIndicator.preOrderPurchaseInd: application/json schema, POST /authorise.merchantRiskIndicator.reorderItemsInd: application/json schema, POST /authorise.merchantRiskIndicator.shipIndicator: application/json schema, POST /authorise.merchantRiskIndicator.deliveryAddressIndicator: application/json schema, POST /authorise.merchantRiskIndicator.deliveryTimeframe: application/json schema, POST /authorise.merchantRiskIndicator.giftCardCurr: application/json schema, POST /authorise.merchantRiskIndicator.preOrderDate: application/json schema, POST /authorise.mpiData.threeDSVersion: application/json schema, POST /authorise.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /authorise.mpiData.transStatusReason: application/json schema, POST /authorise.mpiData.cavv: application/json schema, POST /authorise.mpiData.cavvAlgorithm: application/json schema, POST /authorise.mpiData.xid: application/json schema, POST /authorise.mpiData.eci: application/json schema, POST /authorise.mpiData.authenticationResponse: application/json schema, POST /authorise.mpiData.directoryResponse: application/json schema, POST /authorise.mpiData.riskScore: application/json schema, POST /authorise.mpiData.challengeCancel: application/json schema, POST /authorise.mpiData.dsTransID: application/json schema, POST /authorise.sessionId: application/json schema, POST /authorise.splits[].reference: application/json schema, POST /authorise.splits[].type: application/json schema, POST /authorise.splits[].account: application/json schema, POST /authorise.splits[].description: application/json schema, POST /authorise.shopperName.firstName: application/json schema, POST /authorise.shopperName.lastName: application/json schema, POST /authorise.shopperStatement: application/json schema, POST /authorise.fundDestination.billingAddress.stateOrProvince: application/json schema, POST /authorise.fundDestination.billingAddress.country: application/json schema, POST /authorise.fundDestination.billingAddress.postalCode: application/json schema, POST /authorise.fundDestination.shopperName.lastName: application/json schema, POST /authorise.fundDestination.shopperName.firstName: application/json schema, POST /authorise.fundDestination.shopperReference: application/json schema, POST /authorise.fundDestination.subMerchant.city: application/json schema, POST /authorise.fundDestination.subMerchant.country: application/json schema, POST /authorise.fundDestination.subMerchant.mcc: application/json schema, POST /authorise.fundDestination.subMerchant.name: application/json schema, POST /authorise.fundDestination.subMerchant.taxId: application/json schema, POST /authorise.fundDestination.telephoneNumber: application/json schema, POST /authorise.fundDestination.shopperEmail: application/json schema, POST /authorise.fundDestination.selectedRecurringDetailReference: application/json schema, POST /authorise.installments.plan: application/json schema, POST /authorise.selectedBrand: application/json schema, POST /authorise.reference: application/json schema, POST /authorise.merchantAccount: application/json schema, POST /authorise.fundingSource: application/json schema, POST /authorise.mandate.startsAt: application/json schema, POST /authorise.mandate.amount: application/json schema, POST /authorise.mandate.amountRule: application/json schema, POST /authorise.mandate.billingAttemptsRule: application/json schema, POST /authorise.mandate.billingDay: application/json schema, POST /authorise.mandate.endsAt: application/json schema, POST /authorise.mandate.frequency: application/json schema, POST /authorise.mandate.remarks: application/json schema, POST /authorise.shopperInteraction: application/json schema, POST /authorise.threeDS2RequestData.recurringExpiry: application/json schema, POST /authorise.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema, POST /authorise.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema, POST /authorise.threeDS2RequestData.sdkEphemPubKey.x: application/json schema, POST /authorise.threeDS2RequestData.sdkEphemPubKey.y: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema, POST /authorise.threeDS2RequestData.acquirerMerchantID: application/json schema, POST /authorise.threeDS2RequestData.deviceChannel: application/json schema, POST /authorise.threeDS2RequestData.mcc: application/json schema, POST /authorise.threeDS2RequestData.notificationURL: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorName: application/json schema, POST /authorise.threeDS2RequestData.merchantName: application/json schema, POST /authorise.threeDS2RequestData.threeDSCompInd: application/json schema, POST /authorise.threeDS2RequestData.challengeIndicator: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorURL: application/json schema, POST /authorise.threeDS2RequestData.sdkAppID: application/json schema, POST /authorise.threeDS2RequestData.messageVersion: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema, POST /authorise.threeDS2RequestData.transactionType: application/json schema, POST /authorise.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema, POST /authorise.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema, POST /authorise.threeDS2RequestData.sdkVersion: application/json schema, POST /authorise.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema, POST /authorise.threeDS2RequestData.sdkEncData: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.chAccString: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.chAccChange: application/json schema, POST /authorise.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema, POST /authorise.threeDS2RequestData.sdkTransID: application/json schema, POST /authorise.threeDS2RequestData.whiteListStatus: application/json schema, POST /authorise.threeDS2RequestData.sdkReferenceNumber: application/json schema, POST /authorise.threeDS2RequestData.acquirerBIN: application/json schema, POST /authorise.threeDS2RequestData.threeDSRequestorID: application/json schema, POST /authorise.entityType: application/json schema, POST /authorise.accountInfo.accountCreationDate: application/json schema, POST /authorise.accountInfo.paymentAccountIndicator: application/json schema, POST /authorise.accountInfo.accountChangeDate: application/json schema, POST /authorise.accountInfo.accountChangeIndicator: application/json schema, POST /authorise.accountInfo.accountType: application/json schema, POST /authorise.accountInfo.deliveryAddressUsageDate: application/json schema, POST /authorise.accountInfo.workPhone: application/json schema, POST /authorise.accountInfo.mobilePhone: application/json schema, POST /authorise.accountInfo.paymentAccountAge: application/json schema, POST /authorise.accountInfo.accountAgeIndicator: application/json schema, POST /authorise.accountInfo.passwordChangeIndicator: application/json schema, POST /authorise.accountInfo.deliveryAddressUsageIndicator: application/json schema, POST /authorise.accountInfo.homePhone: application/json schema, POST /authorise.accountInfo.passwordChangeDate: application/json schema, POST /authorise.shopperReference: application/json schema, POST /authorise.recurring.recurringFrequency: application/json schema, POST /authorise.recurring.tokenService: application/json schema, POST /authorise.recurring.contract: application/json schema, POST /authorise.recurring.recurringDetailName: application/json schema, POST /authorise.recurring.recurringExpiry: application/json schema, POST /authorise.shopperEmail: application/json schema, POST /authorise.telephoneNumber: application/json schema, POST /authorise.browserInfo.userAgent: application/json schema, POST /authorise.browserInfo.language: application/json schema, POST /authorise.browserInfo.acceptHeader: application/json schema, POST /authorise.mcc: application/json schema, POST /authorise.recurringProcessingModel: application/json schema, POST /authorise.deliveryAddress.country: application/json schema, POST /authorise.deliveryAddress.postalCode: application/json schema, POST /authorise.deliveryAddress.stateOrProvince: application/json schema, POST /authorise.dccQuote.validTill: application/json schema, POST /authorise.dccQuote.type: application/json schema, POST /authorise.dccQuote.signature: application/json schema, POST /authorise.dccQuote.source: application/json schema, POST /authorise.dccQuote.account: application/json schema, POST /authorise.dccQuote.accountType: application/json schema, POST /authorise.dccQuote.reference: application/json schema, POST /authorise.shopperLocale: application/json schema, POST /authorise.selectedRecurringDetailReference: application/json schema, POST /authorise.shopperIP: application/json schema, POST /authorise.socialSecurityNumber: application/json schema, POST /capture.platformChargebackLogic.targetAccount: application/json schema, POST /capture.platformChargebackLogic.behavior: application/json schema, POST /capture.platformChargebackLogic.costAllocationAccount: application/json schema, POST /capture.reference: application/json schema, POST /capture.originalMerchantReference: application/json schema, POST /capture.tenderReference: application/json schema, POST /capture.uniqueTerminalId: application/json schema, POST /capture.originalReference: application/json schema, POST /capture.splits[].type: application/json schema, POST /capture.splits[].account: application/json schema, POST /capture.splits[].description: application/json schema, POST /capture.splits[].reference: application/json schema, POST /capture.merchantAccount: application/json schema, POST /capture.mpiData.eci: application/json schema, POST /capture.mpiData.threeDSVersion: application/json schema, POST /capture.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /capture.mpiData.cavv: application/json schema, POST /capture.mpiData.riskScore: application/json schema, POST /capture.mpiData.challengeCancel: application/json schema, POST /capture.mpiData.dsTransID: application/json schema, POST /capture.mpiData.transStatusReason: application/json schema, POST /capture.mpiData.authenticationResponse: application/json schema, POST /capture.mpiData.cavvAlgorithm: application/json schema, POST /capture.mpiData.directoryResponse: application/json schema, POST /capture.mpiData.xid: application/json schema, POST /getAuthenticationResult.pspReference: application/json schema, POST /getAuthenticationResult.merchantAccount: application/json schema, POST /retrieve3ds2Result.merchantAccount: application/json schema, POST /retrieve3ds2Result.pspReference: application/json schema, POST /technicalCancel.merchantAccount: application/json schema, POST /technicalCancel.mpiData.cavvAlgorithm: application/json schema, POST /technicalCancel.mpiData.riskScore: application/json schema, POST /technicalCancel.mpiData.eci: application/json schema, POST /technicalCancel.mpiData.threeDSVersion: application/json schema, POST /technicalCancel.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /technicalCancel.mpiData.cavv: application/json schema, POST /technicalCancel.mpiData.directoryResponse: application/json schema, POST /technicalCancel.mpiData.xid: application/json schema, POST /technicalCancel.mpiData.challengeCancel: application/json schema, POST /technicalCancel.mpiData.dsTransID: application/json schema, POST /technicalCancel.mpiData.transStatusReason: application/json schema, POST /technicalCancel.mpiData.authenticationResponse: application/json schema, POST /technicalCancel.originalMerchantReference: application/json schema, POST /technicalCancel.platformChargebackLogic.targetAccount: application/json schema, POST /technicalCancel.platformChargebackLogic.behavior: application/json schema, POST /technicalCancel.platformChargebackLogic.costAllocationAccount: application/json schema, POST /technicalCancel.tenderReference: application/json schema, POST /technicalCancel.uniqueTerminalId: application/json schema, POST /technicalCancel.reference: application/json schema, POST /technicalCancel.splits[].account: application/json schema, POST /technicalCancel.splits[].description: application/json schema, POST /technicalCancel.splits[].reference: application/json schema, POST /technicalCancel.splits[].type: application/json schema, POST /authorise3d.sessionId: application/json schema, POST /authorise3d.installments.plan: application/json schema, POST /authorise3d.applicationInfo.adyenLibrary.version: application/json schema, POST /authorise3d.applicationInfo.adyenLibrary.name: application/json schema, POST /authorise3d.applicationInfo.adyenPaymentSource.name: application/json schema, POST /authorise3d.applicationInfo.adyenPaymentSource.version: application/json schema, POST /authorise3d.applicationInfo.externalPlatform.name: application/json schema, POST /authorise3d.applicationInfo.externalPlatform.version: application/json schema, POST /authorise3d.applicationInfo.externalPlatform.integrator: application/json schema, POST /authorise3d.applicationInfo.merchantApplication.name: application/json schema, POST /authorise3d.applicationInfo.merchantApplication.version: application/json schema, POST /authorise3d.applicationInfo.merchantDevice.os: application/json schema, POST /authorise3d.applicationInfo.merchantDevice.osVersion: application/json schema, POST /authorise3d.applicationInfo.merchantDevice.reference: application/json schema, POST /authorise3d.applicationInfo.shopperInteractionDevice.locale: application/json schema, POST /authorise3d.applicationInfo.shopperInteractionDevice.os: application/json schema, POST /authorise3d.applicationInfo.shopperInteractionDevice.osVersion: application/json schema, POST /authorise3d.shopperStatement: application/json schema, POST /authorise3d.browserInfo.acceptHeader: application/json schema, POST /authorise3d.browserInfo.userAgent: application/json schema, POST /authorise3d.browserInfo.language: application/json schema, POST /authorise3d.merchantAccount: application/json schema, POST /authorise3d.telephoneNumber: application/json schema, POST /authorise3d.billingAddress.country: application/json schema, POST /authorise3d.billingAddress.postalCode: application/json schema, POST /authorise3d.billingAddress.stateOrProvince: application/json schema, POST /authorise3d.mcc: application/json schema, POST /authorise3d.shopperReference: application/json schema, POST /authorise3d.accountInfo.paymentAccountIndicator: application/json schema, POST /authorise3d.accountInfo.accountAgeIndicator: application/json schema, POST /authorise3d.accountInfo.deliveryAddressUsageDate: application/json schema, POST /authorise3d.accountInfo.passwordChangeIndicator: application/json schema, POST /authorise3d.accountInfo.homePhone: application/json schema, POST /authorise3d.accountInfo.mobilePhone: application/json schema, POST /authorise3d.accountInfo.accountChangeIndicator: application/json schema, POST /authorise3d.accountInfo.deliveryAddressUsageIndicator: application/json schema, POST /authorise3d.accountInfo.passwordChangeDate: application/json schema, POST /authorise3d.accountInfo.accountChangeDate: application/json schema, POST /authorise3d.accountInfo.accountType: application/json schema, POST /authorise3d.accountInfo.accountCreationDate: application/json schema, POST /authorise3d.accountInfo.workPhone: application/json schema, POST /authorise3d.accountInfo.paymentAccountAge: application/json schema, POST /authorise3d.recurring.recurringDetailName: application/json schema, POST /authorise3d.recurring.recurringExpiry: application/json schema, POST /authorise3d.recurring.recurringFrequency: application/json schema, POST /authorise3d.recurring.tokenService: application/json schema, POST /authorise3d.recurring.contract: application/json schema, POST /authorise3d.reference: application/json schema, POST /authorise3d.shopperInteraction: application/json schema, POST /authorise3d.splits[].type: application/json schema, POST /authorise3d.splits[].account: application/json schema, POST /authorise3d.splits[].description: application/json schema, POST /authorise3d.splits[].reference: application/json schema, POST /authorise3d.shopperEmail: application/json schema, POST /authorise3d.shopperLocale: application/json schema, POST /authorise3d.socialSecurityNumber: application/json schema, POST /authorise3d.merchantOrderReference: application/json schema, POST /authorise3d.selectedBrand: application/json schema, POST /authorise3d.md: application/json schema, POST /authorise3d.recurringProcessingModel: application/json schema, POST /authorise3d.deliveryAddress.country: application/json schema, POST /authorise3d.deliveryAddress.postalCode: application/json schema, POST /authorise3d.deliveryAddress.stateOrProvince: application/json schema, POST /authorise3d.orderReference: application/json schema, POST /authorise3d.shopperIP: application/json schema, POST /authorise3d.dateOfBirth: application/json schema, POST /authorise3d.deliveryDate: application/json schema, POST /authorise3d.shopperName.firstName: application/json schema, POST /authorise3d.shopperName.lastName: application/json schema, POST /authorise3d.selectedRecurringDetailReference: application/json schema, POST /authorise3d.dccQuote.reference: application/json schema, POST /authorise3d.dccQuote.signature: application/json schema, POST /authorise3d.dccQuote.source: application/json schema, POST /authorise3d.dccQuote.validTill: application/json schema, POST /authorise3d.dccQuote.type: application/json schema, POST /authorise3d.dccQuote.accountType: application/json schema, POST /authorise3d.dccQuote.account: application/json schema, POST /authorise3d.merchantRiskIndicator.shipIndicator: application/json schema, POST /authorise3d.merchantRiskIndicator.deliveryAddressIndicator: application/json schema, POST /authorise3d.merchantRiskIndicator.deliveryTimeframe: application/json schema, POST /authorise3d.merchantRiskIndicator.giftCardCurr: application/json schema, POST /authorise3d.merchantRiskIndicator.preOrderDate: application/json schema, POST /authorise3d.merchantRiskIndicator.reorderItemsInd: application/json schema, POST /authorise3d.merchantRiskIndicator.deliveryEmail: application/json schema, POST /authorise3d.merchantRiskIndicator.preOrderPurchaseInd: application/json schema, POST /authorise3d.threeDS2RequestData.notificationURL: application/json schema, POST /authorise3d.threeDS2RequestData.whiteListStatus: application/json schema, POST /authorise3d.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema, POST /authorise3d.threeDS2RequestData.challengeIndicator: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorID: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSCompInd: application/json schema, POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema, POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.x: application/json schema, POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.y: application/json schema, POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorURL: application/json schema, POST /authorise3d.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema, POST /authorise3d.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema, POST /authorise3d.threeDS2RequestData.sdkVersion: application/json schema, POST /authorise3d.threeDS2RequestData.messageVersion: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema, POST /authorise3d.threeDS2RequestData.merchantName: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema, POST /authorise3d.threeDS2RequestData.acquirerMerchantID: application/json schema, POST /authorise3d.threeDS2RequestData.deviceChannel: application/json schema, POST /authorise3d.threeDS2RequestData.recurringExpiry: application/json schema, POST /authorise3d.threeDS2RequestData.sdkAppID: application/json schema, POST /authorise3d.threeDS2RequestData.sdkTransID: application/json schema, POST /authorise3d.threeDS2RequestData.sdkReferenceNumber: application/json schema, POST /authorise3d.threeDS2RequestData.sdkEncData: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema, POST /authorise3d.threeDS2RequestData.threeDSRequestorName: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.chAccString: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.chAccChange: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema, POST /authorise3d.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema, POST /authorise3d.threeDS2RequestData.mcc: application/json schema, POST /authorise3d.threeDS2RequestData.transactionType: application/json schema, POST /authorise3d.threeDS2RequestData.acquirerBIN: application/json schema, POST /authorise3d.paResponse: application/json schema, POST /cancel.originalReference: application/json schema, POST /cancel.tenderReference: application/json schema, POST /cancel.mpiData.dsTransID: application/json schema, POST /cancel.mpiData.eci: application/json schema, POST /cancel.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /cancel.mpiData.authenticationResponse: application/json schema, POST /cancel.mpiData.cavv: application/json schema, POST /cancel.mpiData.cavvAlgorithm: application/json schema, POST /cancel.mpiData.riskScore: application/json schema, POST /cancel.mpiData.threeDSVersion: application/json schema, POST /cancel.mpiData.transStatusReason: application/json schema, POST /cancel.mpiData.directoryResponse: application/json schema, POST /cancel.mpiData.xid: application/json schema, POST /cancel.mpiData.challengeCancel: application/json schema, POST /cancel.reference: application/json schema, POST /cancel.uniqueTerminalId: application/json schema, POST /cancel.originalMerchantReference: application/json schema, POST /cancel.platformChargebackLogic.behavior: application/json schema, POST /cancel.platformChargebackLogic.costAllocationAccount: application/json schema, POST /cancel.platformChargebackLogic.targetAccount: application/json schema, POST /cancel.splits[].description: application/json schema, POST /cancel.splits[].reference: application/json schema, POST /cancel.splits[].type: application/json schema, POST /cancel.splits[].account: application/json schema, POST /cancel.merchantAccount: application/json schema, POST /cancelOrRefund.uniqueTerminalId: application/json schema, POST /cancelOrRefund.originalReference: application/json schema, POST /cancelOrRefund.platformChargebackLogic.behavior: application/json schema, POST /cancelOrRefund.platformChargebackLogic.costAllocationAccount: application/json schema, POST /cancelOrRefund.platformChargebackLogic.targetAccount: application/json schema, POST /cancelOrRefund.tenderReference: application/json schema, POST /cancelOrRefund.mpiData.xid: application/json schema, POST /cancelOrRefund.mpiData.dsTransID: application/json schema, POST /cancelOrRefund.mpiData.eci: application/json schema, POST /cancelOrRefund.mpiData.authenticationResponse: application/json schema, POST /cancelOrRefund.mpiData.cavv: application/json schema, POST /cancelOrRefund.mpiData.directoryResponse: application/json schema, POST /cancelOrRefund.mpiData.riskScore: application/json schema, POST /cancelOrRefund.mpiData.challengeCancel: application/json schema, POST /cancelOrRefund.mpiData.threeDSVersion: application/json schema, POST /cancelOrRefund.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /cancelOrRefund.mpiData.transStatusReason: application/json schema, POST /cancelOrRefund.mpiData.cavvAlgorithm: application/json schema, POST /cancelOrRefund.originalMerchantReference: application/json schema, POST /cancelOrRefund.reference: application/json schema, POST /cancelOrRefund.merchantAccount: application/json schema, POST /donate.merchantAccount: application/json schema, POST /donate.originalReference: application/json schema, POST /donate.platformChargebackLogic.targetAccount: application/json schema, POST /donate.platformChargebackLogic.behavior: application/json schema, POST /donate.platformChargebackLogic.costAllocationAccount: application/json schema, POST /donate.reference: application/json schema, POST /donate.donationAccount: application/json schema, POST /refund.platformChargebackLogic.targetAccount: application/json schema, POST /refund.platformChargebackLogic.behavior: application/json schema, POST /refund.platformChargebackLogic.costAllocationAccount: application/json schema, POST /refund.uniqueTerminalId: application/json schema, POST /refund.mpiData.challengeCancel: application/json schema, POST /refund.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /refund.mpiData.transStatusReason: application/json schema, POST /refund.mpiData.authenticationResponse: application/json schema, POST /refund.mpiData.cavv: application/json schema, POST /refund.mpiData.directoryResponse: application/json schema, POST /refund.mpiData.xid: application/json schema, POST /refund.mpiData.dsTransID: application/json schema, POST /refund.mpiData.eci: application/json schema, POST /refund.mpiData.threeDSVersion: application/json schema, POST /refund.mpiData.cavvAlgorithm: application/json schema, POST /refund.mpiData.riskScore: application/json schema, POST /refund.originalReference: application/json schema, POST /refund.splits[].description: application/json schema, POST /refund.splits[].reference: application/json schema, POST /refund.splits[].type: application/json schema, POST /refund.splits[].account: application/json schema, POST /refund.reference: application/json schema, POST /refund.originalMerchantReference: application/json schema, POST /refund.tenderReference: application/json schema, POST /refund.merchantAccount: application/json schema, POST /adjustAuthorisation.platformChargebackLogic.targetAccount: application/json schema, POST /adjustAuthorisation.platformChargebackLogic.behavior: application/json schema, POST /adjustAuthorisation.platformChargebackLogic.costAllocationAccount: application/json schema, POST /adjustAuthorisation.splits[].description: application/json schema, POST /adjustAuthorisation.splits[].reference: application/json schema, POST /adjustAuthorisation.splits[].type: application/json schema, POST /adjustAuthorisation.splits[].account: application/json schema, POST /adjustAuthorisation.tenderReference: application/json schema, POST /adjustAuthorisation.merchantAccount: application/json schema, POST /adjustAuthorisation.mpiData.threeDSVersion: application/json schema, POST /adjustAuthorisation.mpiData.tokenAuthenticationVerificationValue: application/json schema, POST /adjustAuthorisation.mpiData.authenticationResponse: application/json schema, POST /adjustAuthorisation.mpiData.cavv: application/json schema, POST /adjustAuthorisation.mpiData.directoryResponse: application/json schema, POST /adjustAuthorisation.mpiData.riskScore: application/json schema, POST /adjustAuthorisation.mpiData.xid: application/json schema, POST /adjustAuthorisation.mpiData.dsTransID: application/json schema, POST /adjustAuthorisation.mpiData.eci: application/json schema, POST /adjustAuthorisation.mpiData.transStatusReason: application/json schema, POST /adjustAuthorisation.mpiData.cavvAlgorithm: application/json schema, POST /adjustAuthorisation.mpiData.challengeCancel: application/json schema, POST /adjustAuthorisation.reference: application/json schema, POST /adjustAuthorisation.uniqueTerminalId: application/json schema, POST /adjustAuthorisation.originalMerchantReference: application/json schema, POST /adjustAuthorisation.originalReference: application/json schema; All numeric fields have min/max values: POST /authorise3ds2.installments.value: application/json schema, POST /authorise3ds2.fraudOffset: application/json schema, POST /authorise3ds2.additionalAmount.value: application/json schema, POST /authorise3ds2.accountInfo.purchasesLast6Months: application/json schema, POST /authorise3ds2.accountInfo.pastTransactionsYear: application/json schema, POST /authorise3ds2.accountInfo.pastTransactionsDay: application/json schema, POST /authorise3ds2.accountInfo.addCardAttemptsDay: application/json schema, POST /authorise3ds2.merchantRiskIndicator.giftCardAmount.value: application/json schema, POST /authorise3ds2.merchantRiskIndicator.giftCardCount: application/json schema, POST /authorise3ds2.dccQuote.interbank.value: application/json schema, POST /authorise3ds2.dccQuote.basePoints: application/json schema, POST /authorise3ds2.dccQuote.baseAmount.value: application/json schema, POST /authorise3ds2.dccQuote.sell.value: application/json schema, POST /authorise3ds2.dccQuote.buy.value: application/json schema, POST /authorise3ds2.splits[].amount.value: application/json schema, POST /authorise3ds2.amount.value: application/json schema, POST /authorise3ds2.threeDS2RequestData.sdkMaxTimeout: application/json schema, POST /authorise3ds2.browserInfo.screenWidth: application/json schema, POST /authorise3ds2.browserInfo.screenHeight: application/json schema, POST /authorise3ds2.browserInfo.colorDepth: application/json schema, POST /authorise3ds2.browserInfo.timeZoneOffset: application/json schema, POST /authorise3ds2.captureDelayHours: application/json schema, POST /voidPendingRefund.splits[].amount.value: application/json schema, POST /voidPendingRefund.modificationAmount.value: application/json schema, POST /authorise.amount.value: application/json schema, POST /authorise.fraudOffset: application/json schema, POST /authorise.merchantRiskIndicator.giftCardCount: application/json schema, POST /authorise.merchantRiskIndicator.giftCardAmount.value: application/json schema, POST /authorise.splits[].amount.value: application/json schema, POST /authorise.captureDelayHours: application/json schema, POST /authorise.installments.value: application/json schema, POST /authorise.threeDS2RequestData.sdkMaxTimeout: application/json schema, POST /authorise.accountInfo.pastTransactionsYear: application/json schema, POST /authorise.accountInfo.pastTransactionsDay: application/json schema, POST /authorise.accountInfo.addCardAttemptsDay: application/json schema, POST /authorise.accountInfo.purchasesLast6Months: application/json schema, POST /authorise.browserInfo.timeZoneOffset: application/json schema, POST /authorise.browserInfo.screenHeight: application/json schema, POST /authorise.browserInfo.screenWidth: application/json schema, POST /authorise.browserInfo.colorDepth: application/json schema, POST /authorise.additionalAmount.value: application/json schema, POST /authorise.dccQuote.sell.value: application/json schema, POST /authorise.dccQuote.basePoints: application/json schema, POST /authorise.dccQuote.buy.value: application/json schema, POST /authorise.dccQuote.interbank.value: application/json schema, POST /authorise.dccQuote.baseAmount.value: application/json schema, POST /capture.modificationAmount.value: application/json schema, POST /capture.splits[].amount.value: application/json schema, POST /technicalCancel.modificationAmount.value: application/json schema, POST /technicalCancel.splits[].amount.value: application/json schema, POST /authorise3d.additionalAmount.value: application/json schema, POST /authorise3d.installments.value: application/json schema, POST /authorise3d.browserInfo.screenWidth: application/json schema, POST /authorise3d.browserInfo.colorDepth: application/json schema, POST /authorise3d.browserInfo.timeZoneOffset: application/json schema, POST /authorise3d.browserInfo.screenHeight: application/json schema, POST /authorise3d.accountInfo.addCardAttemptsDay: application/json schema, POST /authorise3d.accountInfo.purchasesLast6Months: application/json schema, POST /authorise3d.accountInfo.pastTransactionsYear: application/json schema, POST /authorise3d.accountInfo.pastTransactionsDay: application/json schema, POST /authorise3d.splits[].amount.value: application/json schema, POST /authorise3d.amount.value: application/json schema, POST /authorise3d.fraudOffset: application/json schema, POST /authorise3d.dccQuote.basePoints: application/json schema, POST /authorise3d.dccQuote.buy.value: application/json schema, POST /authorise3d.dccQuote.baseAmount.value: application/json schema, POST /authorise3d.dccQuote.sell.value: application/json schema, POST /authorise3d.dccQuote.interbank.value: application/json schema, POST /authorise3d.merchantRiskIndicator.giftCardCount: application/json schema, POST /authorise3d.merchantRiskIndicator.giftCardAmount.value: application/json schema, POST /authorise3d.threeDS2RequestData.sdkMaxTimeout: application/json schema, POST /authorise3d.captureDelayHours: application/json schema, POST /cancel.splits[].amount.value: application/json schema, POST /donate.modificationAmount.value: application/json schema, POST /refund.modificationAmount.value: application/json schema, POST /refund.splits[].amount.value: application/json schema, POST /adjustAuthorisation.splits[].amount.value: application/json schema, POST /adjustAuthorisation.modificationAmount.value: application/json schema
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
        "POST /authorise3ds2.installments.value: application/json schema",
        "POST /authorise3ds2.fraudOffset: application/json schema",
        "POST /authorise3ds2.additionalAmount.value: application/json schema",
        "POST /authorise3ds2.accountInfo.purchasesLast6Months: application/json schema",
        "POST /authorise3ds2.accountInfo.pastTransactionsYear: application/json schema",
        "POST /authorise3ds2.accountInfo.pastTransactionsDay: application/json schema",
        "POST /authorise3ds2.accountInfo.addCardAttemptsDay: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.giftCardAmount.value: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.giftCardCount: application/json schema",
        "POST /authorise3ds2.dccQuote.interbank.value: application/json schema",
        "POST /authorise3ds2.dccQuote.basePoints: application/json schema",
        "POST /authorise3ds2.dccQuote.baseAmount.value: application/json schema",
        "POST /authorise3ds2.dccQuote.sell.value: application/json schema",
        "POST /authorise3ds2.dccQuote.buy.value: application/json schema",
        "POST /authorise3ds2.splits[].amount.value: application/json schema",
        "POST /authorise3ds2.amount.value: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkMaxTimeout: application/json schema",
        "POST /authorise3ds2.browserInfo.screenWidth: application/json schema",
        "POST /authorise3ds2.browserInfo.screenHeight: application/json schema",
        "POST /authorise3ds2.browserInfo.colorDepth: application/json schema",
        "POST /authorise3ds2.browserInfo.timeZoneOffset: application/json schema",
        "POST /authorise3ds2.captureDelayHours: application/json schema",
        "POST /voidPendingRefund.splits[].amount.value: application/json schema",
        "POST /voidPendingRefund.modificationAmount.value: application/json schema",
        "POST /authorise.amount.value: application/json schema",
        "POST /authorise.fraudOffset: application/json schema",
        "POST /authorise.merchantRiskIndicator.giftCardCount: application/json schema",
        "POST /authorise.merchantRiskIndicator.giftCardAmount.value: application/json schema",
        "POST /authorise.splits[].amount.value: application/json schema",
        "POST /authorise.captureDelayHours: application/json schema",
        "POST /authorise.installments.value: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkMaxTimeout: application/json schema",
        "POST /authorise.accountInfo.pastTransactionsYear: application/json schema",
        "POST /authorise.accountInfo.pastTransactionsDay: application/json schema",
        "POST /authorise.accountInfo.addCardAttemptsDay: application/json schema",
        "POST /authorise.accountInfo.purchasesLast6Months: application/json schema",
        "POST /authorise.browserInfo.timeZoneOffset: application/json schema",
        "POST /authorise.browserInfo.screenHeight: application/json schema",
        "POST /authorise.browserInfo.screenWidth: application/json schema",
        "POST /authorise.browserInfo.colorDepth: application/json schema",
        "POST /authorise.additionalAmount.value: application/json schema",
        "POST /authorise.dccQuote.sell.value: application/json schema",
        "POST /authorise.dccQuote.basePoints: application/json schema",
        "POST /authorise.dccQuote.buy.value: application/json schema",
        "POST /authorise.dccQuote.interbank.value: application/json schema",
        "POST /authorise.dccQuote.baseAmount.value: application/json schema",
        "POST /capture.modificationAmount.value: application/json schema",
        "POST /capture.splits[].amount.value: application/json schema",
        "POST /technicalCancel.modificationAmount.value: application/json schema",
        "POST /technicalCancel.splits[].amount.value: application/json schema",
        "POST /authorise3d.additionalAmount.value: application/json schema",
        "POST /authorise3d.installments.value: application/json schema",
        "POST /authorise3d.browserInfo.screenWidth: application/json schema",
        "POST /authorise3d.browserInfo.colorDepth: application/json schema",
        "POST /authorise3d.browserInfo.timeZoneOffset: application/json schema",
        "POST /authorise3d.browserInfo.screenHeight: application/json schema",
        "POST /authorise3d.accountInfo.addCardAttemptsDay: application/json schema",
        "POST /authorise3d.accountInfo.purchasesLast6Months: application/json schema",
        "POST /authorise3d.accountInfo.pastTransactionsYear: application/json schema",
        "POST /authorise3d.accountInfo.pastTransactionsDay: application/json schema",
        "POST /authorise3d.splits[].amount.value: application/json schema",
        "POST /authorise3d.amount.value: application/json schema",
        "POST /authorise3d.fraudOffset: application/json schema",
        "POST /authorise3d.dccQuote.basePoints: application/json schema",
        "POST /authorise3d.dccQuote.buy.value: application/json schema",
        "POST /authorise3d.dccQuote.baseAmount.value: application/json schema",
        "POST /authorise3d.dccQuote.sell.value: application/json schema",
        "POST /authorise3d.dccQuote.interbank.value: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.giftCardCount: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.giftCardAmount.value: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkMaxTimeout: application/json schema",
        "POST /authorise3d.captureDelayHours: application/json schema",
        "POST /cancel.splits[].amount.value: application/json schema",
        "POST /donate.modificationAmount.value: application/json schema",
        "POST /refund.modificationAmount.value: application/json schema",
        "POST /refund.splits[].amount.value: application/json schema",
        "POST /adjustAuthorisation.splits[].amount.value: application/json schema",
        "POST /adjustAuthorisation.modificationAmount.value: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /authorise3ds2.installments.plan: application/json schema",
        "POST /authorise3ds2.selectedBrand: application/json schema",
        "POST /authorise3ds2.billingAddress.country: application/json schema",
        "POST /authorise3ds2.billingAddress.postalCode: application/json schema",
        "POST /authorise3ds2.billingAddress.stateOrProvince: application/json schema",
        "POST /authorise3ds2.merchantAccount: application/json schema",
        "POST /authorise3ds2.shopperEmail: application/json schema",
        "POST /authorise3ds2.accountInfo.paymentAccountIndicator: application/json schema",
        "POST /authorise3ds2.accountInfo.accountChangeDate: application/json schema",
        "POST /authorise3ds2.accountInfo.passwordChangeIndicator: application/json schema",
        "POST /authorise3ds2.accountInfo.workPhone: application/json schema",
        "POST /authorise3ds2.accountInfo.passwordChangeDate: application/json schema",
        "POST /authorise3ds2.accountInfo.mobilePhone: application/json schema",
        "POST /authorise3ds2.accountInfo.paymentAccountAge: application/json schema",
        "POST /authorise3ds2.accountInfo.accountType: application/json schema",
        "POST /authorise3ds2.accountInfo.accountCreationDate: application/json schema",
        "POST /authorise3ds2.accountInfo.deliveryAddressUsageDate: application/json schema",
        "POST /authorise3ds2.accountInfo.deliveryAddressUsageIndicator: application/json schema",
        "POST /authorise3ds2.accountInfo.accountAgeIndicator: application/json schema",
        "POST /authorise3ds2.accountInfo.accountChangeIndicator: application/json schema",
        "POST /authorise3ds2.accountInfo.homePhone: application/json schema",
        "POST /authorise3ds2.mcc: application/json schema",
        "POST /authorise3ds2.threeDS2Result.transStatus: application/json schema",
        "POST /authorise3ds2.threeDS2Result.authenticationValue: application/json schema",
        "POST /authorise3ds2.threeDS2Result.messageVersion: application/json schema",
        "POST /authorise3ds2.threeDS2Result.threeDSServerTransID: application/json schema",
        "POST /authorise3ds2.threeDS2Result.challengeCancel: application/json schema",
        "POST /authorise3ds2.threeDS2Result.challengeIndicator: application/json schema",
        "POST /authorise3ds2.threeDS2Result.transStatusReason: application/json schema",
        "POST /authorise3ds2.threeDS2Result.cavvAlgorithm: application/json schema",
        "POST /authorise3ds2.threeDS2Result.dsTransID: application/json schema",
        "POST /authorise3ds2.threeDS2Result.exemptionIndicator: application/json schema",
        "POST /authorise3ds2.threeDS2Result.whiteListStatus: application/json schema",
        "POST /authorise3ds2.threeDS2Result.eci: application/json schema",
        "POST /authorise3ds2.threeDS2Result.riskScore: application/json schema",
        "POST /authorise3ds2.threeDS2Result.timestamp: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.giftCardCurr: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.preOrderPurchaseInd: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.reorderItemsInd: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.deliveryAddressIndicator: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.preOrderDate: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.deliveryEmail: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.shipIndicator: application/json schema",
        "POST /authorise3ds2.merchantRiskIndicator.deliveryTimeframe: application/json schema",
        "POST /authorise3ds2.applicationInfo.adyenLibrary.version: application/json schema",
        "POST /authorise3ds2.applicationInfo.adyenLibrary.name: application/json schema",
        "POST /authorise3ds2.applicationInfo.adyenPaymentSource.name: application/json schema",
        "POST /authorise3ds2.applicationInfo.adyenPaymentSource.version: application/json schema",
        "POST /authorise3ds2.applicationInfo.externalPlatform.version: application/json schema",
        "POST /authorise3ds2.applicationInfo.externalPlatform.integrator: application/json schema",
        "POST /authorise3ds2.applicationInfo.externalPlatform.name: application/json schema",
        "POST /authorise3ds2.applicationInfo.merchantApplication.version: application/json schema",
        "POST /authorise3ds2.applicationInfo.merchantApplication.name: application/json schema",
        "POST /authorise3ds2.applicationInfo.merchantDevice.os: application/json schema",
        "POST /authorise3ds2.applicationInfo.merchantDevice.osVersion: application/json schema",
        "POST /authorise3ds2.applicationInfo.merchantDevice.reference: application/json schema",
        "POST /authorise3ds2.applicationInfo.shopperInteractionDevice.os: application/json schema",
        "POST /authorise3ds2.applicationInfo.shopperInteractionDevice.osVersion: application/json schema",
        "POST /authorise3ds2.applicationInfo.shopperInteractionDevice.locale: application/json schema",
        "POST /authorise3ds2.telephoneNumber: application/json schema",
        "POST /authorise3ds2.socialSecurityNumber: application/json schema",
        "POST /authorise3ds2.reference: application/json schema",
        "POST /authorise3ds2.dccQuote.type: application/json schema",
        "POST /authorise3ds2.dccQuote.accountType: application/json schema",
        "POST /authorise3ds2.dccQuote.reference: application/json schema",
        "POST /authorise3ds2.dccQuote.signature: application/json schema",
        "POST /authorise3ds2.dccQuote.source: application/json schema",
        "POST /authorise3ds2.dccQuote.validTill: application/json schema",
        "POST /authorise3ds2.dccQuote.account: application/json schema",
        "POST /authorise3ds2.deliveryAddress.country: application/json schema",
        "POST /authorise3ds2.deliveryAddress.postalCode: application/json schema",
        "POST /authorise3ds2.deliveryAddress.stateOrProvince: application/json schema",
        "POST /authorise3ds2.splits[].reference: application/json schema",
        "POST /authorise3ds2.splits[].type: application/json schema",
        "POST /authorise3ds2.splits[].account: application/json schema",
        "POST /authorise3ds2.splits[].description: application/json schema",
        "POST /authorise3ds2.shopperReference: application/json schema",
        "POST /authorise3ds2.deliveryDate: application/json schema",
        "POST /authorise3ds2.dateOfBirth: application/json schema",
        "POST /authorise3ds2.shopperInteraction: application/json schema",
        "POST /authorise3ds2.orderReference: application/json schema",
        "POST /authorise3ds2.sessionId: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkReferenceNumber: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccChange: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acctInfo.chAccString: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.transactionType: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acquirerBIN: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkAppID: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.messageVersion: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.mcc: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.whiteListStatus: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorID: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.challengeIndicator: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkVersion: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkEncData: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.recurringExpiry: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.x: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.y: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.acquirerMerchantID: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.deviceChannel: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.notificationURL: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.merchantName: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorURL: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSCompInd: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.threeDSRequestorName: application/json schema",
        "POST /authorise3ds2.threeDS2RequestData.sdkTransID: application/json schema",
        "POST /authorise3ds2.selectedRecurringDetailReference: application/json schema",
        "POST /authorise3ds2.browserInfo.userAgent: application/json schema",
        "POST /authorise3ds2.browserInfo.language: application/json schema",
        "POST /authorise3ds2.browserInfo.acceptHeader: application/json schema",
        "POST /authorise3ds2.merchantOrderReference: application/json schema",
        "POST /authorise3ds2.recurringProcessingModel: application/json schema",
        "POST /authorise3ds2.shopperName.lastName: application/json schema",
        "POST /authorise3ds2.shopperName.firstName: application/json schema",
        "POST /authorise3ds2.shopperStatement: application/json schema",
        "POST /authorise3ds2.recurring.recurringDetailName: application/json schema",
        "POST /authorise3ds2.recurring.recurringExpiry: application/json schema",
        "POST /authorise3ds2.recurring.recurringFrequency: application/json schema",
        "POST /authorise3ds2.recurring.tokenService: application/json schema",
        "POST /authorise3ds2.recurring.contract: application/json schema",
        "POST /authorise3ds2.threeDS2Token: application/json schema",
        "POST /authorise3ds2.shopperLocale: application/json schema",
        "POST /authorise3ds2.shopperIP: application/json schema",
        "POST /voidPendingRefund.originalMerchantReference: application/json schema",
        "POST /voidPendingRefund.originalReference: application/json schema",
        "POST /voidPendingRefund.platformChargebackLogic.behavior: application/json schema",
        "POST /voidPendingRefund.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /voidPendingRefund.platformChargebackLogic.targetAccount: application/json schema",
        "POST /voidPendingRefund.reference: application/json schema",
        "POST /voidPendingRefund.merchantAccount: application/json schema",
        "POST /voidPendingRefund.splits[].type: application/json schema",
        "POST /voidPendingRefund.splits[].account: application/json schema",
        "POST /voidPendingRefund.splits[].description: application/json schema",
        "POST /voidPendingRefund.splits[].reference: application/json schema",
        "POST /voidPendingRefund.tenderReference: application/json schema",
        "POST /voidPendingRefund.uniqueTerminalId: application/json schema",
        "POST /voidPendingRefund.mpiData.cavv: application/json schema",
        "POST /voidPendingRefund.mpiData.cavvAlgorithm: application/json schema",
        "POST /voidPendingRefund.mpiData.directoryResponse: application/json schema",
        "POST /voidPendingRefund.mpiData.xid: application/json schema",
        "POST /voidPendingRefund.mpiData.challengeCancel: application/json schema",
        "POST /voidPendingRefund.mpiData.eci: application/json schema",
        "POST /voidPendingRefund.mpiData.transStatusReason: application/json schema",
        "POST /voidPendingRefund.mpiData.authenticationResponse: application/json schema",
        "POST /voidPendingRefund.mpiData.riskScore: application/json schema",
        "POST /voidPendingRefund.mpiData.dsTransID: application/json schema",
        "POST /voidPendingRefund.mpiData.threeDSVersion: application/json schema",
        "POST /voidPendingRefund.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /authorise.bankAccount.bankName: application/json schema",
        "POST /authorise.bankAccount.bic: application/json schema",
        "POST /authorise.bankAccount.ownerName: application/json schema",
        "POST /authorise.bankAccount.countryCode: application/json schema",
        "POST /authorise.bankAccount.taxId: application/json schema",
        "POST /authorise.bankAccount.bankAccountNumber: application/json schema",
        "POST /authorise.bankAccount.bankCity: application/json schema",
        "POST /authorise.bankAccount.bankLocationId: application/json schema",
        "POST /authorise.bankAccount.iban: application/json schema",
        "POST /authorise.dateOfBirth: application/json schema",
        "POST /authorise.billingAddress.postalCode: application/json schema",
        "POST /authorise.billingAddress.stateOrProvince: application/json schema",
        "POST /authorise.billingAddress.country: application/json schema",
        "POST /authorise.applicationInfo.adyenLibrary.name: application/json schema",
        "POST /authorise.applicationInfo.adyenLibrary.version: application/json schema",
        "POST /authorise.applicationInfo.adyenPaymentSource.name: application/json schema",
        "POST /authorise.applicationInfo.adyenPaymentSource.version: application/json schema",
        "POST /authorise.applicationInfo.externalPlatform.integrator: application/json schema",
        "POST /authorise.applicationInfo.externalPlatform.name: application/json schema",
        "POST /authorise.applicationInfo.externalPlatform.version: application/json schema",
        "POST /authorise.applicationInfo.merchantApplication.name: application/json schema",
        "POST /authorise.applicationInfo.merchantApplication.version: application/json schema",
        "POST /authorise.applicationInfo.merchantDevice.os: application/json schema",
        "POST /authorise.applicationInfo.merchantDevice.osVersion: application/json schema",
        "POST /authorise.applicationInfo.merchantDevice.reference: application/json schema",
        "POST /authorise.applicationInfo.shopperInteractionDevice.locale: application/json schema",
        "POST /authorise.applicationInfo.shopperInteractionDevice.os: application/json schema",
        "POST /authorise.applicationInfo.shopperInteractionDevice.osVersion: application/json schema",
        "POST /authorise.merchantOrderReference: application/json schema",
        "POST /authorise.orderReference: application/json schema",
        "POST /authorise.fundSource.billingAddress.country: application/json schema",
        "POST /authorise.fundSource.billingAddress.postalCode: application/json schema",
        "POST /authorise.fundSource.billingAddress.stateOrProvince: application/json schema",
        "POST /authorise.fundSource.shopperEmail: application/json schema",
        "POST /authorise.fundSource.shopperName.firstName: application/json schema",
        "POST /authorise.fundSource.shopperName.lastName: application/json schema",
        "POST /authorise.fundSource.telephoneNumber: application/json schema",
        "POST /authorise.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /authorise.platformChargebackLogic.targetAccount: application/json schema",
        "POST /authorise.platformChargebackLogic.behavior: application/json schema",
        "POST /authorise.deliveryDate: application/json schema",
        "POST /authorise.merchantRiskIndicator.deliveryEmail: application/json schema",
        "POST /authorise.merchantRiskIndicator.preOrderPurchaseInd: application/json schema",
        "POST /authorise.merchantRiskIndicator.reorderItemsInd: application/json schema",
        "POST /authorise.merchantRiskIndicator.shipIndicator: application/json schema",
        "POST /authorise.merchantRiskIndicator.deliveryAddressIndicator: application/json schema",
        "POST /authorise.merchantRiskIndicator.deliveryTimeframe: application/json schema",
        "POST /authorise.merchantRiskIndicator.giftCardCurr: application/json schema",
        "POST /authorise.merchantRiskIndicator.preOrderDate: application/json schema",
        "POST /authorise.mpiData.threeDSVersion: application/json schema",
        "POST /authorise.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /authorise.mpiData.transStatusReason: application/json schema",
        "POST /authorise.mpiData.cavv: application/json schema",
        "POST /authorise.mpiData.cavvAlgorithm: application/json schema",
        "POST /authorise.mpiData.xid: application/json schema",
        "POST /authorise.mpiData.eci: application/json schema",
        "POST /authorise.mpiData.authenticationResponse: application/json schema",
        "POST /authorise.mpiData.directoryResponse: application/json schema",
        "POST /authorise.mpiData.riskScore: application/json schema",
        "POST /authorise.mpiData.challengeCancel: application/json schema",
        "POST /authorise.mpiData.dsTransID: application/json schema",
        "POST /authorise.sessionId: application/json schema",
        "POST /authorise.splits[].reference: application/json schema",
        "POST /authorise.splits[].type: application/json schema",
        "POST /authorise.splits[].account: application/json schema",
        "POST /authorise.splits[].description: application/json schema",
        "POST /authorise.shopperName.firstName: application/json schema",
        "POST /authorise.shopperName.lastName: application/json schema",
        "POST /authorise.shopperStatement: application/json schema",
        "POST /authorise.fundDestination.billingAddress.stateOrProvince: application/json schema",
        "POST /authorise.fundDestination.billingAddress.country: application/json schema",
        "POST /authorise.fundDestination.billingAddress.postalCode: application/json schema",
        "POST /authorise.fundDestination.shopperName.lastName: application/json schema",
        "POST /authorise.fundDestination.shopperName.firstName: application/json schema",
        "POST /authorise.fundDestination.shopperReference: application/json schema",
        "POST /authorise.fundDestination.subMerchant.city: application/json schema",
        "POST /authorise.fundDestination.subMerchant.country: application/json schema",
        "POST /authorise.fundDestination.subMerchant.mcc: application/json schema",
        "POST /authorise.fundDestination.subMerchant.name: application/json schema",
        "POST /authorise.fundDestination.subMerchant.taxId: application/json schema",
        "POST /authorise.fundDestination.telephoneNumber: application/json schema",
        "POST /authorise.fundDestination.shopperEmail: application/json schema",
        "POST /authorise.fundDestination.selectedRecurringDetailReference: application/json schema",
        "POST /authorise.installments.plan: application/json schema",
        "POST /authorise.selectedBrand: application/json schema",
        "POST /authorise.reference: application/json schema",
        "POST /authorise.merchantAccount: application/json schema",
        "POST /authorise.fundingSource: application/json schema",
        "POST /authorise.mandate.startsAt: application/json schema",
        "POST /authorise.mandate.amount: application/json schema",
        "POST /authorise.mandate.amountRule: application/json schema",
        "POST /authorise.mandate.billingAttemptsRule: application/json schema",
        "POST /authorise.mandate.billingDay: application/json schema",
        "POST /authorise.mandate.endsAt: application/json schema",
        "POST /authorise.mandate.frequency: application/json schema",
        "POST /authorise.mandate.remarks: application/json schema",
        "POST /authorise.shopperInteraction: application/json schema",
        "POST /authorise.threeDS2RequestData.recurringExpiry: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkEphemPubKey.x: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkEphemPubKey.y: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema",
        "POST /authorise.threeDS2RequestData.acquirerMerchantID: application/json schema",
        "POST /authorise.threeDS2RequestData.deviceChannel: application/json schema",
        "POST /authorise.threeDS2RequestData.mcc: application/json schema",
        "POST /authorise.threeDS2RequestData.notificationURL: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorName: application/json schema",
        "POST /authorise.threeDS2RequestData.merchantName: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSCompInd: application/json schema",
        "POST /authorise.threeDS2RequestData.challengeIndicator: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorURL: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkAppID: application/json schema",
        "POST /authorise.threeDS2RequestData.messageVersion: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema",
        "POST /authorise.threeDS2RequestData.transactionType: application/json schema",
        "POST /authorise.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema",
        "POST /authorise.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkVersion: application/json schema",
        "POST /authorise.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkEncData: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.chAccString: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.chAccChange: application/json schema",
        "POST /authorise.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkTransID: application/json schema",
        "POST /authorise.threeDS2RequestData.whiteListStatus: application/json schema",
        "POST /authorise.threeDS2RequestData.sdkReferenceNumber: application/json schema",
        "POST /authorise.threeDS2RequestData.acquirerBIN: application/json schema",
        "POST /authorise.threeDS2RequestData.threeDSRequestorID: application/json schema",
        "POST /authorise.entityType: application/json schema",
        "POST /authorise.accountInfo.accountCreationDate: application/json schema",
        "POST /authorise.accountInfo.paymentAccountIndicator: application/json schema",
        "POST /authorise.accountInfo.accountChangeDate: application/json schema",
        "POST /authorise.accountInfo.accountChangeIndicator: application/json schema",
        "POST /authorise.accountInfo.accountType: application/json schema",
        "POST /authorise.accountInfo.deliveryAddressUsageDate: application/json schema",
        "POST /authorise.accountInfo.workPhone: application/json schema",
        "POST /authorise.accountInfo.mobilePhone: application/json schema",
        "POST /authorise.accountInfo.paymentAccountAge: application/json schema",
        "POST /authorise.accountInfo.accountAgeIndicator: application/json schema",
        "POST /authorise.accountInfo.passwordChangeIndicator: application/json schema",
        "POST /authorise.accountInfo.deliveryAddressUsageIndicator: application/json schema",
        "POST /authorise.accountInfo.homePhone: application/json schema",
        "POST /authorise.accountInfo.passwordChangeDate: application/json schema",
        "POST /authorise.shopperReference: application/json schema",
        "POST /authorise.recurring.recurringFrequency: application/json schema",
        "POST /authorise.recurring.tokenService: application/json schema",
        "POST /authorise.recurring.contract: application/json schema",
        "POST /authorise.recurring.recurringDetailName: application/json schema",
        "POST /authorise.recurring.recurringExpiry: application/json schema",
        "POST /authorise.shopperEmail: application/json schema",
        "POST /authorise.telephoneNumber: application/json schema",
        "POST /authorise.browserInfo.userAgent: application/json schema",
        "POST /authorise.browserInfo.language: application/json schema",
        "POST /authorise.browserInfo.acceptHeader: application/json schema",
        "POST /authorise.mcc: application/json schema",
        "POST /authorise.recurringProcessingModel: application/json schema",
        "POST /authorise.deliveryAddress.country: application/json schema",
        "POST /authorise.deliveryAddress.postalCode: application/json schema",
        "POST /authorise.deliveryAddress.stateOrProvince: application/json schema",
        "POST /authorise.dccQuote.validTill: application/json schema",
        "POST /authorise.dccQuote.type: application/json schema",
        "POST /authorise.dccQuote.signature: application/json schema",
        "POST /authorise.dccQuote.source: application/json schema",
        "POST /authorise.dccQuote.account: application/json schema",
        "POST /authorise.dccQuote.accountType: application/json schema",
        "POST /authorise.dccQuote.reference: application/json schema",
        "POST /authorise.shopperLocale: application/json schema",
        "POST /authorise.selectedRecurringDetailReference: application/json schema",
        "POST /authorise.shopperIP: application/json schema",
        "POST /authorise.socialSecurityNumber: application/json schema",
        "POST /capture.platformChargebackLogic.targetAccount: application/json schema",
        "POST /capture.platformChargebackLogic.behavior: application/json schema",
        "POST /capture.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /capture.reference: application/json schema",
        "POST /capture.originalMerchantReference: application/json schema",
        "POST /capture.tenderReference: application/json schema",
        "POST /capture.uniqueTerminalId: application/json schema",
        "POST /capture.originalReference: application/json schema",
        "POST /capture.splits[].type: application/json schema",
        "POST /capture.splits[].account: application/json schema",
        "POST /capture.splits[].description: application/json schema",
        "POST /capture.splits[].reference: application/json schema",
        "POST /capture.merchantAccount: application/json schema",
        "POST /capture.mpiData.eci: application/json schema",
        "POST /capture.mpiData.threeDSVersion: application/json schema",
        "POST /capture.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /capture.mpiData.cavv: application/json schema",
        "POST /capture.mpiData.riskScore: application/json schema",
        "POST /capture.mpiData.challengeCancel: application/json schema",
        "POST /capture.mpiData.dsTransID: application/json schema",
        "POST /capture.mpiData.transStatusReason: application/json schema",
        "POST /capture.mpiData.authenticationResponse: application/json schema",
        "POST /capture.mpiData.cavvAlgorithm: application/json schema",
        "POST /capture.mpiData.directoryResponse: application/json schema",
        "POST /capture.mpiData.xid: application/json schema",
        "POST /getAuthenticationResult.pspReference: application/json schema",
        "POST /getAuthenticationResult.merchantAccount: application/json schema",
        "POST /retrieve3ds2Result.merchantAccount: application/json schema",
        "POST /retrieve3ds2Result.pspReference: application/json schema",
        "POST /technicalCancel.merchantAccount: application/json schema",
        "POST /technicalCancel.mpiData.cavvAlgorithm: application/json schema",
        "POST /technicalCancel.mpiData.riskScore: application/json schema",
        "POST /technicalCancel.mpiData.eci: application/json schema",
        "POST /technicalCancel.mpiData.threeDSVersion: application/json schema",
        "POST /technicalCancel.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /technicalCancel.mpiData.cavv: application/json schema",
        "POST /technicalCancel.mpiData.directoryResponse: application/json schema",
        "POST /technicalCancel.mpiData.xid: application/json schema",
        "POST /technicalCancel.mpiData.challengeCancel: application/json schema",
        "POST /technicalCancel.mpiData.dsTransID: application/json schema",
        "POST /technicalCancel.mpiData.transStatusReason: application/json schema",
        "POST /technicalCancel.mpiData.authenticationResponse: application/json schema",
        "POST /technicalCancel.originalMerchantReference: application/json schema",
        "POST /technicalCancel.platformChargebackLogic.targetAccount: application/json schema",
        "POST /technicalCancel.platformChargebackLogic.behavior: application/json schema",
        "POST /technicalCancel.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /technicalCancel.tenderReference: application/json schema",
        "POST /technicalCancel.uniqueTerminalId: application/json schema",
        "POST /technicalCancel.reference: application/json schema",
        "POST /technicalCancel.splits[].account: application/json schema",
        "POST /technicalCancel.splits[].description: application/json schema",
        "POST /technicalCancel.splits[].reference: application/json schema",
        "POST /technicalCancel.splits[].type: application/json schema",
        "POST /authorise3d.sessionId: application/json schema",
        "POST /authorise3d.installments.plan: application/json schema",
        "POST /authorise3d.applicationInfo.adyenLibrary.version: application/json schema",
        "POST /authorise3d.applicationInfo.adyenLibrary.name: application/json schema",
        "POST /authorise3d.applicationInfo.adyenPaymentSource.name: application/json schema",
        "POST /authorise3d.applicationInfo.adyenPaymentSource.version: application/json schema",
        "POST /authorise3d.applicationInfo.externalPlatform.name: application/json schema",
        "POST /authorise3d.applicationInfo.externalPlatform.version: application/json schema",
        "POST /authorise3d.applicationInfo.externalPlatform.integrator: application/json schema",
        "POST /authorise3d.applicationInfo.merchantApplication.name: application/json schema",
        "POST /authorise3d.applicationInfo.merchantApplication.version: application/json schema",
        "POST /authorise3d.applicationInfo.merchantDevice.os: application/json schema",
        "POST /authorise3d.applicationInfo.merchantDevice.osVersion: application/json schema",
        "POST /authorise3d.applicationInfo.merchantDevice.reference: application/json schema",
        "POST /authorise3d.applicationInfo.shopperInteractionDevice.locale: application/json schema",
        "POST /authorise3d.applicationInfo.shopperInteractionDevice.os: application/json schema",
        "POST /authorise3d.applicationInfo.shopperInteractionDevice.osVersion: application/json schema",
        "POST /authorise3d.shopperStatement: application/json schema",
        "POST /authorise3d.browserInfo.acceptHeader: application/json schema",
        "POST /authorise3d.browserInfo.userAgent: application/json schema",
        "POST /authorise3d.browserInfo.language: application/json schema",
        "POST /authorise3d.merchantAccount: application/json schema",
        "POST /authorise3d.telephoneNumber: application/json schema",
        "POST /authorise3d.billingAddress.country: application/json schema",
        "POST /authorise3d.billingAddress.postalCode: application/json schema",
        "POST /authorise3d.billingAddress.stateOrProvince: application/json schema",
        "POST /authorise3d.mcc: application/json schema",
        "POST /authorise3d.shopperReference: application/json schema",
        "POST /authorise3d.accountInfo.paymentAccountIndicator: application/json schema",
        "POST /authorise3d.accountInfo.accountAgeIndicator: application/json schema",
        "POST /authorise3d.accountInfo.deliveryAddressUsageDate: application/json schema",
        "POST /authorise3d.accountInfo.passwordChangeIndicator: application/json schema",
        "POST /authorise3d.accountInfo.homePhone: application/json schema",
        "POST /authorise3d.accountInfo.mobilePhone: application/json schema",
        "POST /authorise3d.accountInfo.accountChangeIndicator: application/json schema",
        "POST /authorise3d.accountInfo.deliveryAddressUsageIndicator: application/json schema",
        "POST /authorise3d.accountInfo.passwordChangeDate: application/json schema",
        "POST /authorise3d.accountInfo.accountChangeDate: application/json schema",
        "POST /authorise3d.accountInfo.accountType: application/json schema",
        "POST /authorise3d.accountInfo.accountCreationDate: application/json schema",
        "POST /authorise3d.accountInfo.workPhone: application/json schema",
        "POST /authorise3d.accountInfo.paymentAccountAge: application/json schema",
        "POST /authorise3d.recurring.recurringDetailName: application/json schema",
        "POST /authorise3d.recurring.recurringExpiry: application/json schema",
        "POST /authorise3d.recurring.recurringFrequency: application/json schema",
        "POST /authorise3d.recurring.tokenService: application/json schema",
        "POST /authorise3d.recurring.contract: application/json schema",
        "POST /authorise3d.reference: application/json schema",
        "POST /authorise3d.shopperInteraction: application/json schema",
        "POST /authorise3d.splits[].type: application/json schema",
        "POST /authorise3d.splits[].account: application/json schema",
        "POST /authorise3d.splits[].description: application/json schema",
        "POST /authorise3d.splits[].reference: application/json schema",
        "POST /authorise3d.shopperEmail: application/json schema",
        "POST /authorise3d.shopperLocale: application/json schema",
        "POST /authorise3d.socialSecurityNumber: application/json schema",
        "POST /authorise3d.merchantOrderReference: application/json schema",
        "POST /authorise3d.selectedBrand: application/json schema",
        "POST /authorise3d.md: application/json schema",
        "POST /authorise3d.recurringProcessingModel: application/json schema",
        "POST /authorise3d.deliveryAddress.country: application/json schema",
        "POST /authorise3d.deliveryAddress.postalCode: application/json schema",
        "POST /authorise3d.deliveryAddress.stateOrProvince: application/json schema",
        "POST /authorise3d.orderReference: application/json schema",
        "POST /authorise3d.shopperIP: application/json schema",
        "POST /authorise3d.dateOfBirth: application/json schema",
        "POST /authorise3d.deliveryDate: application/json schema",
        "POST /authorise3d.shopperName.firstName: application/json schema",
        "POST /authorise3d.shopperName.lastName: application/json schema",
        "POST /authorise3d.selectedRecurringDetailReference: application/json schema",
        "POST /authorise3d.dccQuote.reference: application/json schema",
        "POST /authorise3d.dccQuote.signature: application/json schema",
        "POST /authorise3d.dccQuote.source: application/json schema",
        "POST /authorise3d.dccQuote.validTill: application/json schema",
        "POST /authorise3d.dccQuote.type: application/json schema",
        "POST /authorise3d.dccQuote.accountType: application/json schema",
        "POST /authorise3d.dccQuote.account: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.shipIndicator: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.deliveryAddressIndicator: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.deliveryTimeframe: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.giftCardCurr: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.preOrderDate: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.reorderItemsInd: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.deliveryEmail: application/json schema",
        "POST /authorise3d.merchantRiskIndicator.preOrderPurchaseInd: application/json schema",
        "POST /authorise3d.threeDS2RequestData.notificationURL: application/json schema",
        "POST /authorise3d.threeDS2RequestData.whiteListStatus: application/json schema",
        "POST /authorise3d.threeDS2RequestData.paymentAuthenticationUseCase: application/json schema",
        "POST /authorise3d.threeDS2RequestData.challengeIndicator: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorID: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSCompInd: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.kty: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.x: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.y: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkEphemPubKey.crv: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorAuthenticationInfo.threeDSReqAuthData: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorURL: application/json schema",
        "POST /authorise3d.threeDS2RequestData.deviceRenderOptions.sdkUiType[]: application/json schema",
        "POST /authorise3d.threeDS2RequestData.deviceRenderOptions.sdkInterface: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkVersion: application/json schema",
        "POST /authorise3d.threeDS2RequestData.messageVersion: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorPriorAuthenticationInfo.threeDSReqPriorAuthData: application/json schema",
        "POST /authorise3d.threeDS2RequestData.merchantName: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorAuthenticationInd: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acquirerMerchantID: application/json schema",
        "POST /authorise3d.threeDS2RequestData.deviceChannel: application/json schema",
        "POST /authorise3d.threeDS2RequestData.recurringExpiry: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkAppID: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkTransID: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkReferenceNumber: application/json schema",
        "POST /authorise3d.threeDS2RequestData.sdkEncData: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorChallengeInd: application/json schema",
        "POST /authorise3d.threeDS2RequestData.threeDSRequestorName: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.shipAddressUsage: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.chAccString: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.nbPurchaseAccount: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.provisionAttemptsDay: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.chAccChange: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.paymentAccAge: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acctInfo.chAccPwChange: application/json schema",
        "POST /authorise3d.threeDS2RequestData.mcc: application/json schema",
        "POST /authorise3d.threeDS2RequestData.transactionType: application/json schema",
        "POST /authorise3d.threeDS2RequestData.acquirerBIN: application/json schema",
        "POST /authorise3d.paResponse: application/json schema",
        "POST /cancel.originalReference: application/json schema",
        "POST /cancel.tenderReference: application/json schema",
        "POST /cancel.mpiData.dsTransID: application/json schema",
        "POST /cancel.mpiData.eci: application/json schema",
        "POST /cancel.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /cancel.mpiData.authenticationResponse: application/json schema",
        "POST /cancel.mpiData.cavv: application/json schema",
        "POST /cancel.mpiData.cavvAlgorithm: application/json schema",
        "POST /cancel.mpiData.riskScore: application/json schema",
        "POST /cancel.mpiData.threeDSVersion: application/json schema",
        "POST /cancel.mpiData.transStatusReason: application/json schema",
        "POST /cancel.mpiData.directoryResponse: application/json schema",
        "POST /cancel.mpiData.xid: application/json schema",
        "POST /cancel.mpiData.challengeCancel: application/json schema",
        "POST /cancel.reference: application/json schema",
        "POST /cancel.uniqueTerminalId: application/json schema",
        "POST /cancel.originalMerchantReference: application/json schema",
        "POST /cancel.platformChargebackLogic.behavior: application/json schema",
        "POST /cancel.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /cancel.platformChargebackLogic.targetAccount: application/json schema",
        "POST /cancel.splits[].description: application/json schema",
        "POST /cancel.splits[].reference: application/json schema",
        "POST /cancel.splits[].type: application/json schema",
        "POST /cancel.splits[].account: application/json schema",
        "POST /cancel.merchantAccount: application/json schema",
        "POST /cancelOrRefund.uniqueTerminalId: application/json schema",
        "POST /cancelOrRefund.originalReference: application/json schema",
        "POST /cancelOrRefund.platformChargebackLogic.behavior: application/json schema",
        "POST /cancelOrRefund.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /cancelOrRefund.platformChargebackLogic.targetAccount: application/json schema",
        "POST /cancelOrRefund.tenderReference: application/json schema",
        "POST /cancelOrRefund.mpiData.xid: application/json schema",
        "POST /cancelOrRefund.mpiData.dsTransID: application/json schema",
        "POST /cancelOrRefund.mpiData.eci: application/json schema",
        "POST /cancelOrRefund.mpiData.authenticationResponse: application/json schema",
        "POST /cancelOrRefund.mpiData.cavv: application/json schema",
        "POST /cancelOrRefund.mpiData.directoryResponse: application/json schema",
        "POST /cancelOrRefund.mpiData.riskScore: application/json schema",
        "POST /cancelOrRefund.mpiData.challengeCancel: application/json schema",
        "POST /cancelOrRefund.mpiData.threeDSVersion: application/json schema",
        "POST /cancelOrRefund.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /cancelOrRefund.mpiData.transStatusReason: application/json schema",
        "POST /cancelOrRefund.mpiData.cavvAlgorithm: application/json schema",
        "POST /cancelOrRefund.originalMerchantReference: application/json schema",
        "POST /cancelOrRefund.reference: application/json schema",
        "POST /cancelOrRefund.merchantAccount: application/json schema",
        "POST /donate.merchantAccount: application/json schema",
        "POST /donate.originalReference: application/json schema",
        "POST /donate.platformChargebackLogic.targetAccount: application/json schema",
        "POST /donate.platformChargebackLogic.behavior: application/json schema",
        "POST /donate.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /donate.reference: application/json schema",
        "POST /donate.donationAccount: application/json schema",
        "POST /refund.platformChargebackLogic.targetAccount: application/json schema",
        "POST /refund.platformChargebackLogic.behavior: application/json schema",
        "POST /refund.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /refund.uniqueTerminalId: application/json schema",
        "POST /refund.mpiData.challengeCancel: application/json schema",
        "POST /refund.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /refund.mpiData.transStatusReason: application/json schema",
        "POST /refund.mpiData.authenticationResponse: application/json schema",
        "POST /refund.mpiData.cavv: application/json schema",
        "POST /refund.mpiData.directoryResponse: application/json schema",
        "POST /refund.mpiData.xid: application/json schema",
        "POST /refund.mpiData.dsTransID: application/json schema",
        "POST /refund.mpiData.eci: application/json schema",
        "POST /refund.mpiData.threeDSVersion: application/json schema",
        "POST /refund.mpiData.cavvAlgorithm: application/json schema",
        "POST /refund.mpiData.riskScore: application/json schema",
        "POST /refund.originalReference: application/json schema",
        "POST /refund.splits[].description: application/json schema",
        "POST /refund.splits[].reference: application/json schema",
        "POST /refund.splits[].type: application/json schema",
        "POST /refund.splits[].account: application/json schema",
        "POST /refund.reference: application/json schema",
        "POST /refund.originalMerchantReference: application/json schema",
        "POST /refund.tenderReference: application/json schema",
        "POST /refund.merchantAccount: application/json schema",
        "POST /adjustAuthorisation.platformChargebackLogic.targetAccount: application/json schema",
        "POST /adjustAuthorisation.platformChargebackLogic.behavior: application/json schema",
        "POST /adjustAuthorisation.platformChargebackLogic.costAllocationAccount: application/json schema",
        "POST /adjustAuthorisation.splits[].description: application/json schema",
        "POST /adjustAuthorisation.splits[].reference: application/json schema",
        "POST /adjustAuthorisation.splits[].type: application/json schema",
        "POST /adjustAuthorisation.splits[].account: application/json schema",
        "POST /adjustAuthorisation.tenderReference: application/json schema",
        "POST /adjustAuthorisation.merchantAccount: application/json schema",
        "POST /adjustAuthorisation.mpiData.threeDSVersion: application/json schema",
        "POST /adjustAuthorisation.mpiData.tokenAuthenticationVerificationValue: application/json schema",
        "POST /adjustAuthorisation.mpiData.authenticationResponse: application/json schema",
        "POST /adjustAuthorisation.mpiData.cavv: application/json schema",
        "POST /adjustAuthorisation.mpiData.directoryResponse: application/json schema",
        "POST /adjustAuthorisation.mpiData.riskScore: application/json schema",
        "POST /adjustAuthorisation.mpiData.xid: application/json schema",
        "POST /adjustAuthorisation.mpiData.dsTransID: application/json schema",
        "POST /adjustAuthorisation.mpiData.eci: application/json schema",
        "POST /adjustAuthorisation.mpiData.transStatusReason: application/json schema",
        "POST /adjustAuthorisation.mpiData.cavvAlgorithm: application/json schema",
        "POST /adjustAuthorisation.mpiData.challengeCancel: application/json schema",
        "POST /adjustAuthorisation.reference: application/json schema",
        "POST /adjustAuthorisation.uniqueTerminalId: application/json schema",
        "POST /adjustAuthorisation.originalMerchantReference: application/json schema",
        "POST /adjustAuthorisation.originalReference: application/json schema"
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

