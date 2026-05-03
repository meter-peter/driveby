# API Validation Report

Generated: 2026-05-03T21:05:13+03:00
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
- Security
- Versioning
- Specification
- Documentation
- Error Handling
- Schema


### Failed Tags
- authorization
- lifecycle
- specification
- documentation
- responses
- compliance
- quality
- usability
- errors
- standards
- versioning
- schema
- validation
- request
- security
- compatibility
- openapi
- authentication


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "GetTaxFormResponse": extra sibling fields: [contentEncoding]
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
      "Specification structure is valid": "invalid components: schema \"GetTaxFormResponse\": extra sibling fields: [contentEncoding]"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /closeAccount: request body, POST /closeAccount: 500 application/json response, POST /closeAccount: 200 application/json response, POST /closeAccount: 202 application/json response, POST /closeAccount: 401 application/json response, POST /closeAccount: 422 application/json response, POST /createAccount: request body, POST /createAccount: 422 application/json response, POST /createAccount: 500 application/json response, POST /createAccount: 200 application/json response, POST /createAccount: 202 application/json response, POST /createAccount: 401 application/json response, POST /updateAccount: request body, POST /updateAccount: 401 application/json response, POST /updateAccount: 422 application/json response, POST /updateAccount: 500 application/json response, POST /updateAccount: 200 application/json response, POST /updateAccount: 202 application/json response, POST /checkAccountHolder: request body, POST /checkAccountHolder: 422 application/json response, POST /checkAccountHolder: 500 application/json response, POST /checkAccountHolder: 200 application/json response, POST /checkAccountHolder: 202 application/json response, POST /checkAccountHolder: 401 application/json response, POST /createAccountHolder: request body, POST /createAccountHolder: 500 application/json response, POST /createAccountHolder: 200 application/json response, POST /createAccountHolder: 401 application/json response, POST /createAccountHolder: 422 application/json response, POST /deleteBankAccounts: request body, POST /deleteBankAccounts: 200 application/json response, POST /deleteBankAccounts: 202 application/json response, POST /deleteBankAccounts: 401 application/json response, POST /deleteBankAccounts: 422 application/json response, POST /deleteBankAccounts: 500 application/json response, POST /deleteLegalArrangements: request body, POST /deleteLegalArrangements: 401 application/json response, POST /deleteLegalArrangements: 422 application/json response, POST /deleteLegalArrangements: 500 application/json response, POST /deleteLegalArrangements: 202 application/json response, POST /updateAccountHolderState: request body, POST /updateAccountHolderState: 200 application/json response, POST /updateAccountHolderState: 202 application/json response, POST /updateAccountHolderState: 401 application/json response, POST /updateAccountHolderState: 422 application/json response, POST /updateAccountHolderState: 500 application/json response, POST /unSuspendAccountHolder: request body, POST /unSuspendAccountHolder: 401 application/json response, POST /unSuspendAccountHolder: 422 application/json response, POST /unSuspendAccountHolder: 500 application/json response, POST /unSuspendAccountHolder: 200 application/json response, POST /unSuspendAccountHolder: 202 application/json response, POST /updateAccountHolder: request body, POST /updateAccountHolder: 202 application/json response, POST /updateAccountHolder: 401 application/json response, POST /updateAccountHolder: 422 application/json response, POST /updateAccountHolder: 500 application/json response, POST /updateAccountHolder: 200 application/json response, POST /getAccountHolder: request body, POST /getAccountHolder: 401 application/json response, POST /getAccountHolder: 422 application/json response, POST /getAccountHolder: 500 application/json response, POST /getAccountHolder: 200 application/json response, POST /getAccountHolder: 202 application/json response, POST /closeAccountHolder: request body, POST /closeAccountHolder: 202 application/json response, POST /closeAccountHolder: 401 application/json response, POST /closeAccountHolder: 422 application/json response, POST /closeAccountHolder: 500 application/json response, POST /closeAccountHolder: 200 application/json response, POST /uploadDocument: request body, POST /uploadDocument: 401 application/json response, POST /uploadDocument: 422 application/json response, POST /uploadDocument: 500 application/json response, POST /uploadDocument: 200 application/json response, POST /uploadDocument: 202 application/json response, POST /deleteSignatories: request body, POST /deleteSignatories: application/json request body, POST /deleteSignatories: 202 application/json response, POST /deleteSignatories: 401 application/json response, POST /deleteSignatories: 422 application/json response, POST /deleteSignatories: 500 application/json response, POST /deleteSignatories: 200 application/json response, POST /deleteShareholders: request body, POST /deleteShareholders: 401 application/json response, POST /deleteShareholders: 422 application/json response, POST /deleteShareholders: 500 application/json response, POST /deleteShareholders: 200 application/json response, POST /deleteShareholders: 202 application/json response, POST /closeStores: request body, POST /closeStores: application/json request body, POST /closeStores: 422 application/json response, POST /closeStores: 500 application/json response, POST /closeStores: 200 application/json response, POST /closeStores: 401 application/json response, POST /deletePayoutMethods: request body, POST /deletePayoutMethods: 401 application/json response, POST /deletePayoutMethods: 422 application/json response, POST /deletePayoutMethods: 500 application/json response, POST /deletePayoutMethods: 200 application/json response, POST /deletePayoutMethods: 202 application/json response, POST /getTaxForm: request body, POST /getTaxForm: 422 application/json response, POST /getTaxForm: 500 application/json response, POST /getTaxForm: 200 application/json response, POST /getTaxForm: 401 application/json response, POST /getUploadedDocuments: request body, POST /getUploadedDocuments: 401 application/json response, POST /getUploadedDocuments: 422 application/json response, POST /getUploadedDocuments: 500 application/json response, POST /getUploadedDocuments: 200 application/json response, POST /suspendAccountHolder: request body, POST /suspendAccountHolder: 500 application/json response, POST /suspendAccountHolder: 200 application/json response, POST /suspendAccountHolder: 202 application/json response, POST /suspendAccountHolder: 401 application/json response, POST /suspendAccountHolder: 422 application/json response; All schemas have descriptions: KYCLegalArrangementCheckResult, SuspendAccountHolderRequest, AccountHolderStatus, GetTaxFormResponse, DeleteSignatoriesRequest, PayoutMethod, ViasPersonalData, CreateAccountHolderRequest, GetTaxFormRequest, DeletePayoutMethodRequest, AccountPayoutState, KYCShareholderCheckResult, UltimateParentCompanyBusinessDetails, CloseStoresRequest, ViasPhoneNumber, IndividualDetails, MigratedAccounts, UpdateAccountHolderStateRequest, CloseAccountHolderResponse, KYCPayoutMethodCheckResult, UnSuspendAccountHolderRequest, KYCVerificationResult, LegalArrangementDetail, GenericResponse, GetUploadedDocumentsRequest, ServiceError, GetUploadedDocumentsResponse, ErrorFieldType, UnSuspendAccountHolderResponse, Amount, DocumentDetail, CreateAccountRequest, SignatoryContact, DeleteLegalArrangementRequest, KYCUltimateParentCompanyCheckResult, CloseAccountRequest, UltimateParentCompany, GetAccountHolderResponse, KYCSignatoryCheckResult, PersonalDocumentData, LegalArrangementEntityDetail, AccountProcessingState, DeleteShareholderRequest, ShareholderContact, PayoutScheduleResponse, GetAccountHolderStatusResponse, FieldType, StoreDetail, Account, MigrationData, KYCCheckResult, LegalArrangementRequest, CloseAccountResponse, KYCCheckSummary, KYCLegalArrangementEntityCheckResult, UpdateAccountResponse, UpdatePayoutScheduleRequest, GetAccountHolderRequest, DeleteBankAccountRequest, SuspendAccountHolderResponse, AccountHolderDetails, BusinessDetails, BankAccountDetail, UpdateAccountHolderRequest, MigratedStores, CreateAccountHolderResponse, ViasAddress, AccountEvent, UpdateAccountHolderResponse, ViasName, CreateAccountResponse, UpdateAccountRequest, UploadDocumentRequest, PerformVerificationRequest, CloseAccountHolderRequest, KYCCheckStatusData
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
        "POST /closeAccount: request body",
        "POST /closeAccount: 500 application/json response",
        "POST /closeAccount: 200 application/json response",
        "POST /closeAccount: 202 application/json response",
        "POST /closeAccount: 401 application/json response",
        "POST /closeAccount: 422 application/json response",
        "POST /createAccount: request body",
        "POST /createAccount: 422 application/json response",
        "POST /createAccount: 500 application/json response",
        "POST /createAccount: 200 application/json response",
        "POST /createAccount: 202 application/json response",
        "POST /createAccount: 401 application/json response",
        "POST /updateAccount: request body",
        "POST /updateAccount: 401 application/json response",
        "POST /updateAccount: 422 application/json response",
        "POST /updateAccount: 500 application/json response",
        "POST /updateAccount: 200 application/json response",
        "POST /updateAccount: 202 application/json response",
        "POST /checkAccountHolder: request body",
        "POST /checkAccountHolder: 422 application/json response",
        "POST /checkAccountHolder: 500 application/json response",
        "POST /checkAccountHolder: 200 application/json response",
        "POST /checkAccountHolder: 202 application/json response",
        "POST /checkAccountHolder: 401 application/json response",
        "POST /createAccountHolder: request body",
        "POST /createAccountHolder: 500 application/json response",
        "POST /createAccountHolder: 200 application/json response",
        "POST /createAccountHolder: 401 application/json response",
        "POST /createAccountHolder: 422 application/json response",
        "POST /deleteBankAccounts: request body",
        "POST /deleteBankAccounts: 200 application/json response",
        "POST /deleteBankAccounts: 202 application/json response",
        "POST /deleteBankAccounts: 401 application/json response",
        "POST /deleteBankAccounts: 422 application/json response",
        "POST /deleteBankAccounts: 500 application/json response",
        "POST /deleteLegalArrangements: request body",
        "POST /deleteLegalArrangements: 401 application/json response",
        "POST /deleteLegalArrangements: 422 application/json response",
        "POST /deleteLegalArrangements: 500 application/json response",
        "POST /deleteLegalArrangements: 202 application/json response",
        "POST /updateAccountHolderState: request body",
        "POST /updateAccountHolderState: 200 application/json response",
        "POST /updateAccountHolderState: 202 application/json response",
        "POST /updateAccountHolderState: 401 application/json response",
        "POST /updateAccountHolderState: 422 application/json response",
        "POST /updateAccountHolderState: 500 application/json response",
        "POST /unSuspendAccountHolder: request body",
        "POST /unSuspendAccountHolder: 401 application/json response",
        "POST /unSuspendAccountHolder: 422 application/json response",
        "POST /unSuspendAccountHolder: 500 application/json response",
        "POST /unSuspendAccountHolder: 200 application/json response",
        "POST /unSuspendAccountHolder: 202 application/json response",
        "POST /updateAccountHolder: request body",
        "POST /updateAccountHolder: 202 application/json response",
        "POST /updateAccountHolder: 401 application/json response",
        "POST /updateAccountHolder: 422 application/json response",
        "POST /updateAccountHolder: 500 application/json response",
        "POST /updateAccountHolder: 200 application/json response",
        "POST /getAccountHolder: request body",
        "POST /getAccountHolder: 401 application/json response",
        "POST /getAccountHolder: 422 application/json response",
        "POST /getAccountHolder: 500 application/json response",
        "POST /getAccountHolder: 200 application/json response",
        "POST /getAccountHolder: 202 application/json response",
        "POST /closeAccountHolder: request body",
        "POST /closeAccountHolder: 202 application/json response",
        "POST /closeAccountHolder: 401 application/json response",
        "POST /closeAccountHolder: 422 application/json response",
        "POST /closeAccountHolder: 500 application/json response",
        "POST /closeAccountHolder: 200 application/json response",
        "POST /uploadDocument: request body",
        "POST /uploadDocument: 401 application/json response",
        "POST /uploadDocument: 422 application/json response",
        "POST /uploadDocument: 500 application/json response",
        "POST /uploadDocument: 200 application/json response",
        "POST /uploadDocument: 202 application/json response",
        "POST /deleteSignatories: request body",
        "POST /deleteSignatories: application/json request body",
        "POST /deleteSignatories: 202 application/json response",
        "POST /deleteSignatories: 401 application/json response",
        "POST /deleteSignatories: 422 application/json response",
        "POST /deleteSignatories: 500 application/json response",
        "POST /deleteSignatories: 200 application/json response",
        "POST /deleteShareholders: request body",
        "POST /deleteShareholders: 401 application/json response",
        "POST /deleteShareholders: 422 application/json response",
        "POST /deleteShareholders: 500 application/json response",
        "POST /deleteShareholders: 200 application/json response",
        "POST /deleteShareholders: 202 application/json response",
        "POST /closeStores: request body",
        "POST /closeStores: application/json request body",
        "POST /closeStores: 422 application/json response",
        "POST /closeStores: 500 application/json response",
        "POST /closeStores: 200 application/json response",
        "POST /closeStores: 401 application/json response",
        "POST /deletePayoutMethods: request body",
        "POST /deletePayoutMethods: 401 application/json response",
        "POST /deletePayoutMethods: 422 application/json response",
        "POST /deletePayoutMethods: 500 application/json response",
        "POST /deletePayoutMethods: 200 application/json response",
        "POST /deletePayoutMethods: 202 application/json response",
        "POST /getTaxForm: request body",
        "POST /getTaxForm: 422 application/json response",
        "POST /getTaxForm: 500 application/json response",
        "POST /getTaxForm: 200 application/json response",
        "POST /getTaxForm: 401 application/json response",
        "POST /getUploadedDocuments: request body",
        "POST /getUploadedDocuments: 401 application/json response",
        "POST /getUploadedDocuments: 422 application/json response",
        "POST /getUploadedDocuments: 500 application/json response",
        "POST /getUploadedDocuments: 200 application/json response",
        "POST /suspendAccountHolder: request body",
        "POST /suspendAccountHolder: 500 application/json response",
        "POST /suspendAccountHolder: 200 application/json response",
        "POST /suspendAccountHolder: 202 application/json response",
        "POST /suspendAccountHolder: 401 application/json response",
        "POST /suspendAccountHolder: 422 application/json response"
      ],
      "All schemas have descriptions": [
        "KYCLegalArrangementCheckResult",
        "SuspendAccountHolderRequest",
        "AccountHolderStatus",
        "GetTaxFormResponse",
        "DeleteSignatoriesRequest",
        "PayoutMethod",
        "ViasPersonalData",
        "CreateAccountHolderRequest",
        "GetTaxFormRequest",
        "DeletePayoutMethodRequest",
        "AccountPayoutState",
        "KYCShareholderCheckResult",
        "UltimateParentCompanyBusinessDetails",
        "CloseStoresRequest",
        "ViasPhoneNumber",
        "IndividualDetails",
        "MigratedAccounts",
        "UpdateAccountHolderStateRequest",
        "CloseAccountHolderResponse",
        "KYCPayoutMethodCheckResult",
        "UnSuspendAccountHolderRequest",
        "KYCVerificationResult",
        "LegalArrangementDetail",
        "GenericResponse",
        "GetUploadedDocumentsRequest",
        "ServiceError",
        "GetUploadedDocumentsResponse",
        "ErrorFieldType",
        "UnSuspendAccountHolderResponse",
        "Amount",
        "DocumentDetail",
        "CreateAccountRequest",
        "SignatoryContact",
        "DeleteLegalArrangementRequest",
        "KYCUltimateParentCompanyCheckResult",
        "CloseAccountRequest",
        "UltimateParentCompany",
        "GetAccountHolderResponse",
        "KYCSignatoryCheckResult",
        "PersonalDocumentData",
        "LegalArrangementEntityDetail",
        "AccountProcessingState",
        "DeleteShareholderRequest",
        "ShareholderContact",
        "PayoutScheduleResponse",
        "GetAccountHolderStatusResponse",
        "FieldType",
        "StoreDetail",
        "Account",
        "MigrationData",
        "KYCCheckResult",
        "LegalArrangementRequest",
        "CloseAccountResponse",
        "KYCCheckSummary",
        "KYCLegalArrangementEntityCheckResult",
        "UpdateAccountResponse",
        "UpdatePayoutScheduleRequest",
        "GetAccountHolderRequest",
        "DeleteBankAccountRequest",
        "SuspendAccountHolderResponse",
        "AccountHolderDetails",
        "BusinessDetails",
        "BankAccountDetail",
        "UpdateAccountHolderRequest",
        "MigratedStores",
        "CreateAccountHolderResponse",
        "ViasAddress",
        "AccountEvent",
        "UpdateAccountHolderResponse",
        "ViasName",
        "CreateAccountResponse",
        "UpdateAccountRequest",
        "UploadDocumentRequest",
        "PerformVerificationRequest",
        "CloseAccountHolderRequest",
        "KYCCheckStatusData"
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
- **Message:** Request validation issues found: All string fields have length constraints: POST /closeAccountHolder.accountHolderCode: application/json schema, POST /checkAccountHolder.accountHolderCode: application/json schema, POST /checkAccountHolder.accountStateType: application/json schema, POST /suspendAccountHolder.accountHolderCode: application/json schema, POST /unSuspendAccountHolder.accountHolderCode: application/json schema, POST /closeAccount.accountCode: application/json schema, POST /deleteShareholders.accountHolderCode: application/json schema, POST /deleteShareholders.shareholderCodes[]: application/json schema, POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].storeName: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantAccount: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].storeReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].status: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].store: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].shopperInteraction: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].virtualAccount: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].splitConfigurationUUID: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].logo: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantCategoryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantHouseNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerNationality: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountName: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].currencyCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerHouseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].taxId: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountUUID: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].branchCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].accountNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].urlForVerification: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankCity: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].countryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCity: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerDateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankBicSwift: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankName: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerPostalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].accountType: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].iban: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerName: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerStreet: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].checkCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].email: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementMembers[]: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalEntityType: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.legalBusinessName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].email: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].jobTitle: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderType: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockExchange: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockTicker: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.doingBusinessAs: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.taxId: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].jobTitle: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].email: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.registrationNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].taxNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].name: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].registrationNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalForm: application/json schema, POST /updateAccountHolder.accountHolderDetails.legalArrangements[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.lastReviewDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.payoutMethods[].merchantAccount: application/json schema, POST /updateAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.payoutMethods[].recurringDetailReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.payoutMethods[].shopperReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.bankAggregatorDataReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.merchantCategoryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.email: application/json schema, POST /updateAccountHolder.accountHolderDetails.fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.doingBusinessAs: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.registrationNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.legalBusinessName: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderType: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].email: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].jobTitle: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.city: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.country: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.houseNumberOrName: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.postalCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.stateOrProvince: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.street: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].email: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneType: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryCode: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryReference: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.dateOfBirth: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].number: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].type: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].fullPhoneNumber: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].jobTitle: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].webAddress: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.stockExchange: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.stockTicker: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.taxId: application/json schema, POST /updateAccountHolder.accountHolderDetails.businessDetails.stockNumber: application/json schema, POST /updateAccountHolder.description: application/json schema, POST /updateAccountHolder.legalEntity: application/json schema, POST /updateAccountHolder.primaryCurrency: application/json schema, POST /updateAccountHolder.verificationProfile: application/json schema, POST /updateAccountHolder.accountHolderCode: application/json schema, POST /updateAccountHolderState.reason: application/json schema, POST /updateAccountHolderState.stateType: application/json schema, POST /updateAccountHolderState.accountHolderCode: application/json schema, POST /getTaxForm.accountHolderCode: application/json schema, POST /getTaxForm.formType: application/json schema, POST /getUploadedDocuments.accountHolderCode: application/json schema, POST /getUploadedDocuments.bankAccountUUID: application/json schema, POST /getUploadedDocuments.shareholderCode: application/json schema, POST /createAccount.accountHolderCode: application/json schema, POST /createAccount.bankAccountUUID: application/json schema, POST /createAccount.description: application/json schema, POST /createAccount.payoutMethodCode: application/json schema, POST /createAccount.payoutSchedule: application/json schema, POST /createAccount.payoutScheduleReason: application/json schema, POST /createAccount.payoutSpeed: application/json schema, POST /deleteBankAccounts.bankAccountUUIDs[]: application/json schema, POST /deleteBankAccounts.accountHolderCode: application/json schema, POST /getAccountHolder.accountCode: application/json schema, POST /getAccountHolder.accountHolderCode: application/json schema, POST /uploadDocument.documentContent: application/json schema, POST /uploadDocument.documentDetail.documentType: application/json schema, POST /uploadDocument.documentDetail.legalArrangementCode: application/json schema, POST /uploadDocument.documentDetail.legalArrangementEntityCode: application/json schema, POST /uploadDocument.documentDetail.signatoryCode: application/json schema, POST /uploadDocument.documentDetail.bankAccountUUID: application/json schema, POST /uploadDocument.documentDetail.shareholderCode: application/json schema, POST /uploadDocument.documentDetail.accountHolderCode: application/json schema, POST /uploadDocument.documentDetail.filename: application/json schema, POST /uploadDocument.documentDetail.description: application/json schema, POST /deletePayoutMethods.payoutMethodCodes[]: application/json schema, POST /deletePayoutMethods.accountHolderCode: application/json schema, POST /createAccountHolder.description: application/json schema, POST /createAccountHolder.legalEntity: application/json schema, POST /createAccountHolder.primaryCurrency: application/json schema, POST /createAccountHolder.verificationProfile: application/json schema, POST /createAccountHolder.accountHolderCode: application/json schema, POST /createAccountHolder.accountHolderDetails.merchantCategoryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.city: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.country: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.street: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.legalBusinessName: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderType: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].jobTitle: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderReference: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].email: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryReference: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].email: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].jobTitle: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.stockExchange: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.stockTicker: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.doingBusinessAs: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.registrationNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.stockNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.businessDetails.taxId: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].store: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].shopperInteraction: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].virtualAccount: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].status: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].storeName: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].logo: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].splitConfigurationUUID: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].storeReference: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantAccount: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantCategoryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantHouseNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.storeDetails[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAggregatorDataReference: application/json schema, POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.lastReviewDate: application/json schema, POST /createAccountHolder.accountHolderDetails.payoutMethods[].merchantAccount: application/json schema, POST /createAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodCode: application/json schema, POST /createAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodReference: application/json schema, POST /createAccountHolder.accountHolderDetails.payoutMethods[].recurringDetailReference: application/json schema, POST /createAccountHolder.accountHolderDetails.payoutMethods[].shopperReference: application/json schema, POST /createAccountHolder.accountHolderDetails.email: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerPostalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountReference: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankCity: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerName: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].countryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].branchCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerStreet: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].taxId: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankName: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].accountNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].checkCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCity: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerDateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerHouseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].accountType: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].iban: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountUUID: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountName: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].currencyCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerState: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerNationality: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankBicSwift: application/json schema, POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].urlForVerification: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].name: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].taxNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].registrationNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityReference: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].email: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalEntityType: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementMembers[]: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.doingBusinessAs: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.taxId: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockExchange: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.registrationNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.legalBusinessName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderType: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].email: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderReference: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].jobTitle: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneType: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].jobTitle: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].email: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.dateOfBirth: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].number: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryCode: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryReference: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].fullPhoneNumber: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].webAddress: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockTicker: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementReference: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalForm: application/json schema, POST /createAccountHolder.accountHolderDetails.legalArrangements[].type: application/json schema, POST /createAccountHolder.accountHolderDetails.address.country: application/json schema, POST /createAccountHolder.accountHolderDetails.address.houseNumberOrName: application/json schema, POST /createAccountHolder.accountHolderDetails.address.postalCode: application/json schema, POST /createAccountHolder.accountHolderDetails.address.stateOrProvince: application/json schema, POST /createAccountHolder.accountHolderDetails.address.street: application/json schema, POST /createAccountHolder.accountHolderDetails.address.city: application/json schema, POST /createAccountHolder.accountHolderDetails.fullPhoneNumber: application/json schema, POST /closeStores.stores[]: application/json schema, POST /closeStores.accountHolderCode: application/json schema, POST /deleteLegalArrangements.accountHolderCode: application/json schema, POST /deleteLegalArrangements.legalArrangements[].legalArrangementCode: application/json schema, POST /deleteLegalArrangements.legalArrangements[].legalArrangementEntityCodes[]: application/json schema, POST /updateAccount.payoutSchedule.action: application/json schema, POST /updateAccount.payoutSchedule.reason: application/json schema, POST /updateAccount.payoutSchedule.schedule: application/json schema, POST /updateAccount.payoutSpeed: application/json schema, POST /updateAccount.accountCode: application/json schema, POST /updateAccount.bankAccountUUID: application/json schema, POST /updateAccount.description: application/json schema, POST /updateAccount.payoutMethodCode: application/json schema, POST /deleteSignatories.signatoryCodes[]: application/json schema, POST /deleteSignatories.accountHolderCode: application/json schema; All numeric fields have min/max values: POST /checkAccountHolder.tier: application/json schema, POST /updateAccountHolder.processingTier: application/json schema, POST /getTaxForm.year: application/json schema, POST /createAccountHolder.processingTier: application/json schema
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
        "POST /checkAccountHolder.tier: application/json schema",
        "POST /updateAccountHolder.processingTier: application/json schema",
        "POST /getTaxForm.year: application/json schema",
        "POST /createAccountHolder.processingTier: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /closeAccountHolder.accountHolderCode: application/json schema",
        "POST /checkAccountHolder.accountHolderCode: application/json schema",
        "POST /checkAccountHolder.accountStateType: application/json schema",
        "POST /suspendAccountHolder.accountHolderCode: application/json schema",
        "POST /unSuspendAccountHolder.accountHolderCode: application/json schema",
        "POST /closeAccount.accountCode: application/json schema",
        "POST /deleteShareholders.accountHolderCode: application/json schema",
        "POST /deleteShareholders.shareholderCodes[]: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.individualDetails.personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].storeName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantAccount: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].storeReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].status: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].store: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].shopperInteraction: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].virtualAccount: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].splitConfigurationUUID: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].logo: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantCategoryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.storeDetails[].merchantHouseNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerNationality: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].currencyCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerHouseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].taxId: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountUUID: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].branchCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].accountNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].urlForVerification: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankCity: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].countryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCity: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerDateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankBicSwift: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].bankName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerPostalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].accountType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].iban: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].ownerStreet: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAccountDetails[].checkCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementMembers[]: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalEntityType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.legalBusinessName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].jobTitle: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockExchange: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockTicker: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.doingBusinessAs: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.taxId: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].jobTitle: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.registrationNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].taxNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].name: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].registrationNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].legalForm: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.legalArrangements[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.lastReviewDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.payoutMethods[].merchantAccount: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.payoutMethods[].recurringDetailReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.payoutMethods[].shopperReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.principalBusinessAddress.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.bankAggregatorDataReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.merchantCategoryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.doingBusinessAs: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.registrationNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.legalBusinessName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].jobTitle: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.city: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.country: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.houseNumberOrName: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.postalCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.stateOrProvince: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].address.street: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].email: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneType: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryCode: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryReference: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.dateOfBirth: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].number: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].type: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].fullPhoneNumber: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].jobTitle: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.signatories[].webAddress: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.stockExchange: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.stockTicker: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.taxId: application/json schema",
        "POST /updateAccountHolder.accountHolderDetails.businessDetails.stockNumber: application/json schema",
        "POST /updateAccountHolder.description: application/json schema",
        "POST /updateAccountHolder.legalEntity: application/json schema",
        "POST /updateAccountHolder.primaryCurrency: application/json schema",
        "POST /updateAccountHolder.verificationProfile: application/json schema",
        "POST /updateAccountHolder.accountHolderCode: application/json schema",
        "POST /updateAccountHolderState.reason: application/json schema",
        "POST /updateAccountHolderState.stateType: application/json schema",
        "POST /updateAccountHolderState.accountHolderCode: application/json schema",
        "POST /getTaxForm.accountHolderCode: application/json schema",
        "POST /getTaxForm.formType: application/json schema",
        "POST /getUploadedDocuments.accountHolderCode: application/json schema",
        "POST /getUploadedDocuments.bankAccountUUID: application/json schema",
        "POST /getUploadedDocuments.shareholderCode: application/json schema",
        "POST /createAccount.accountHolderCode: application/json schema",
        "POST /createAccount.bankAccountUUID: application/json schema",
        "POST /createAccount.description: application/json schema",
        "POST /createAccount.payoutMethodCode: application/json schema",
        "POST /createAccount.payoutSchedule: application/json schema",
        "POST /createAccount.payoutScheduleReason: application/json schema",
        "POST /createAccount.payoutSpeed: application/json schema",
        "POST /deleteBankAccounts.bankAccountUUIDs[]: application/json schema",
        "POST /deleteBankAccounts.accountHolderCode: application/json schema",
        "POST /getAccountHolder.accountCode: application/json schema",
        "POST /getAccountHolder.accountHolderCode: application/json schema",
        "POST /uploadDocument.documentContent: application/json schema",
        "POST /uploadDocument.documentDetail.documentType: application/json schema",
        "POST /uploadDocument.documentDetail.legalArrangementCode: application/json schema",
        "POST /uploadDocument.documentDetail.legalArrangementEntityCode: application/json schema",
        "POST /uploadDocument.documentDetail.signatoryCode: application/json schema",
        "POST /uploadDocument.documentDetail.bankAccountUUID: application/json schema",
        "POST /uploadDocument.documentDetail.shareholderCode: application/json schema",
        "POST /uploadDocument.documentDetail.accountHolderCode: application/json schema",
        "POST /uploadDocument.documentDetail.filename: application/json schema",
        "POST /uploadDocument.documentDetail.description: application/json schema",
        "POST /deletePayoutMethods.payoutMethodCodes[]: application/json schema",
        "POST /deletePayoutMethods.accountHolderCode: application/json schema",
        "POST /createAccountHolder.description: application/json schema",
        "POST /createAccountHolder.legalEntity: application/json schema",
        "POST /createAccountHolder.primaryCurrency: application/json schema",
        "POST /createAccountHolder.verificationProfile: application/json schema",
        "POST /createAccountHolder.accountHolderCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.merchantCategoryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.principalBusinessAddress.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.legalBusinessName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].jobTitle: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].shareholderCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.shareholders[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].signatoryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].jobTitle: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.stockExchange: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.stockTicker: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.doingBusinessAs: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.registrationNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.stockNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.businessDetails.taxId: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].store: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].shopperInteraction: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].virtualAccount: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].status: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].storeName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].logo: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].splitConfigurationUUID: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].storeReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantAccount: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantCategoryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].merchantHouseNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.storeDetails[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAggregatorDataReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.individualDetails.personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.lastReviewDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.payoutMethods[].merchantAccount: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.payoutMethods[].payoutMethodReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.payoutMethods[].recurringDetailReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.payoutMethods[].shopperReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerPostalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankCity: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].countryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].branchCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerStreet: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].taxId: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].accountNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].checkCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCity: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerDateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerHouseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].accountType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].iban: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountUUID: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankAccountName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].currencyCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].ownerNationality: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].bankBicSwift: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.bankAccountDetails[].urlForVerification: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].name: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].taxNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].registrationNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].individualDetails.personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementEntityReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalEntityType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].legalArrangementMembers[]: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.doingBusinessAs: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.taxId: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.registrationNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockExchange: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.stockTicker: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].businessDetails.legalBusinessName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.listedUltimateParentCompany[].ultimateParentCompanyCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockExchange: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.registrationNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.legalBusinessName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].shareholderReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].jobTitle: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.shareholders[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneType: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].phoneNumber.phoneCountryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].jobTitle: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].email: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.dateOfBirth: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].issuerState: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].number: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].personalData.documentData[].expirationDate: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].signatoryReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].fullPhoneNumber: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.signatories[].webAddress: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementEntities[].businessDetails.stockTicker: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalArrangementReference: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].legalForm: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.legalArrangements[].type: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.country: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.houseNumberOrName: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.postalCode: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.stateOrProvince: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.street: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.address.city: application/json schema",
        "POST /createAccountHolder.accountHolderDetails.fullPhoneNumber: application/json schema",
        "POST /closeStores.stores[]: application/json schema",
        "POST /closeStores.accountHolderCode: application/json schema",
        "POST /deleteLegalArrangements.accountHolderCode: application/json schema",
        "POST /deleteLegalArrangements.legalArrangements[].legalArrangementCode: application/json schema",
        "POST /deleteLegalArrangements.legalArrangements[].legalArrangementEntityCodes[]: application/json schema",
        "POST /updateAccount.payoutSchedule.action: application/json schema",
        "POST /updateAccount.payoutSchedule.reason: application/json schema",
        "POST /updateAccount.payoutSchedule.schedule: application/json schema",
        "POST /updateAccount.payoutSpeed: application/json schema",
        "POST /updateAccount.accountCode: application/json schema",
        "POST /updateAccount.bankAccountUUID: application/json schema",
        "POST /updateAccount.description: application/json schema",
        "POST /updateAccount.payoutMethodCode: application/json schema",
        "POST /deleteSignatories.signatoryCodes[]: application/json schema",
        "POST /deleteSignatories.accountHolderCode: application/json schema"
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

