# API Validation Report

Generated: 2026-05-03T21:05:23+03:00
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
- Versioning
- Specification
- Documentation
- Error Handling


### Failed Tags
- schema
- compatibility
- openapi
- compliance
- quality
- usability
- errors
- responses
- versioning
- lifecycle
- validation
- specification
- documentation
- standards
- request


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "CertificateDetail": error parsing regexp: invalid or unsupported Perl syntax: `(?!`
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
      "Specification structure is valid": "invalid components: schema \"CertificateDetail\": error parsing regexp: invalid or unsupported Perl syntax: `(?!`"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All operations have clear summaries: POST /#X-Amz-Target=CertificateManager.RenewCertificate, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration, POST /#X-Amz-Target=CertificateManager.GetCertificate, POST /#X-Amz-Target=CertificateManager.DescribeCertificate, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate, POST /#X-Amz-Target=CertificateManager.ExportCertificate, POST /#X-Amz-Target=CertificateManager.ImportCertificate, POST /#X-Amz-Target=CertificateManager.ListCertificates, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions, POST /#X-Amz-Target=CertificateManager.RequestCertificate, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail, POST /#X-Amz-Target=CertificateManager.DeleteCertificate; All parameters have descriptions: POST /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Target, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Target; All request/response bodies have examples: POST /#X-Amz-Target=CertificateManager.RenewCertificate: request body, POST /#X-Amz-Target=CertificateManager.RenewCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.RenewCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.RenewCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: request body, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: request body, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: application/json request body, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 480 application/json response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 481 application/json response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 482 application/json response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 483 application/json response, POST /#X-Amz-Target=CertificateManager.GetCertificate: request body, POST /#X-Amz-Target=CertificateManager.GetCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.GetCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: request body, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: request body, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 483 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 484 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 485 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: request body, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 483 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 484 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 485 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 486 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: request body, POST /#X-Amz-Target=CertificateManager.ExportCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: request body, POST /#X-Amz-Target=CertificateManager.ImportCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 483 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 484 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 485 application/json response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 486 application/json response, POST /#X-Amz-Target=CertificateManager.ListCertificates: request body, POST /#X-Amz-Target=CertificateManager.ListCertificates: application/json request body, POST /#X-Amz-Target=CertificateManager.ListCertificates: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ListCertificates: 200 application/json response, POST /#X-Amz-Target=CertificateManager.ListCertificates: 480 application/json response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: request body, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: application/json request body, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 480 application/json response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 481 application/json response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 482 application/json response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 483 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: request body, POST /#X-Amz-Target=CertificateManager.RequestCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 486 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 200 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 483 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 484 application/json response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 485 application/json response, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 200 application/json response, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 480 application/json response, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: request body, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: application/json request body, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 480 application/json response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 481 application/json response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 482 application/json response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 483 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: request body, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: application/json request body, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 485 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 480 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 481 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 482 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 483 application/json response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 484 application/json response; All schemas have descriptions: ExtendedKeyUsageFilterList, DomainList, TagList, KeyUsageList, FailureReason, UpdateCertificateOptionsRequest, RenewalStatus, TagKey, ExportCertificateResponse, KeyAlgorithm, MaxItems, DescribeCertificateRequest, DomainValidationList, ResourceNotFoundException, AddTagsToCertificateRequest, ThrottlingException, CertificateStatuses, ExtendedKeyUsageList, InvalidStateException, CertificateSummaryList, CertificateType, CertificateTransparencyLoggingPreference, Arn, RenewalEligibility, TagValue, KeyUsageFilterList, TStamp, InUseList, RequestInProgressException, ImportCertificateRequest, RemoveTagsFromCertificateRequest, CertificateBody, DeleteCertificateRequest, PutAccountConfigurationRequest, ExtendedKeyUsageName, ListTagsForCertificateRequest, NullableBoolean, DomainStatus, TagPolicyException, NextToken, CertificateBodyBlob, TooManyTagsException, ValidationMethod, RevocationReason, ValidationException, ListTagsForCertificateResponse, ImportCertificateResponse, PcaArn, InvalidTagException, ExportCertificateRequest, SortBy, ResendValidationEmailRequest, InvalidArnException, KeyUsageNames, InvalidDomainValidationOptionsException, SortOrder, InvalidArgsException, AccessDeniedException, PassphraseBlob, CertificateChain, String, InvalidParameterException, LimitExceededException, KeyUsageName, ConflictException, RecordType, DescribeCertificateResponse, DomainNameString, ListCertificatesResponse, CertificateStatus, GetCertificateRequest, ListCertificatesRequest, IdempotencyToken, PrivateKey, CertificateChainBlob, RequestCertificateRequest, RenewCertificateRequest, GetCertificateResponse, PrivateKeyBlob, RequestCertificateResponse, GetAccountConfigurationResponse, ValidationEmailList, PositiveInteger, ResourceInUseException, KeyAlgorithmList, DomainValidationOptionList, ExtendedKeyUsageNames; All enums have descriptions: FailureReason: enum value NO_AVAILABLE_CONTACTS, FailureReason: enum value ADDITIONAL_VERIFICATION_REQUIRED, FailureReason: enum value DOMAIN_NOT_ALLOWED, FailureReason: enum value INVALID_PUBLIC_DOMAIN, FailureReason: enum value DOMAIN_VALIDATION_DENIED, FailureReason: enum value CAA_ERROR, FailureReason: enum value PCA_LIMIT_EXCEEDED, FailureReason: enum value PCA_INVALID_ARN, FailureReason: enum value PCA_INVALID_STATE, FailureReason: enum value PCA_REQUEST_FAILED, FailureReason: enum value PCA_NAME_CONSTRAINTS_VALIDATION, FailureReason: enum value PCA_RESOURCE_NOT_FOUND, FailureReason: enum value PCA_INVALID_ARGS, FailureReason: enum value PCA_INVALID_DURATION, FailureReason: enum value PCA_ACCESS_DENIED, FailureReason: enum value SLR_NOT_FOUND, FailureReason: enum value OTHER, RenewalStatus: enum value PENDING_AUTO_RENEWAL, RenewalStatus: enum value PENDING_VALIDATION, RenewalStatus: enum value SUCCESS, RenewalStatus: enum value FAILED, KeyAlgorithm: enum value RSA_1024, KeyAlgorithm: enum value RSA_2048, KeyAlgorithm: enum value RSA_3072, KeyAlgorithm: enum value RSA_4096, KeyAlgorithm: enum value EC_prime256v1, KeyAlgorithm: enum value EC_secp384r1, KeyAlgorithm: enum value EC_secp521r1, CertificateType: enum value IMPORTED, CertificateType: enum value AMAZON_ISSUED, CertificateType: enum value PRIVATE, CertificateTransparencyLoggingPreference: enum value ENABLED, CertificateTransparencyLoggingPreference: enum value DISABLED, RenewalEligibility: enum value ELIGIBLE, RenewalEligibility: enum value INELIGIBLE, ExtendedKeyUsageName: enum value TLS_WEB_SERVER_AUTHENTICATION, ExtendedKeyUsageName: enum value TLS_WEB_CLIENT_AUTHENTICATION, ExtendedKeyUsageName: enum value CODE_SIGNING, ExtendedKeyUsageName: enum value EMAIL_PROTECTION, ExtendedKeyUsageName: enum value TIME_STAMPING, ExtendedKeyUsageName: enum value OCSP_SIGNING, ExtendedKeyUsageName: enum value IPSEC_END_SYSTEM, ExtendedKeyUsageName: enum value IPSEC_TUNNEL, ExtendedKeyUsageName: enum value IPSEC_USER, ExtendedKeyUsageName: enum value ANY, ExtendedKeyUsageName: enum value NONE, ExtendedKeyUsageName: enum value CUSTOM, DomainStatus: enum value PENDING_VALIDATION, DomainStatus: enum value SUCCESS, DomainStatus: enum value FAILED, ValidationMethod: enum value EMAIL, ValidationMethod: enum value DNS, RevocationReason: enum value UNSPECIFIED, RevocationReason: enum value KEY_COMPROMISE, RevocationReason: enum value CA_COMPROMISE, RevocationReason: enum value AFFILIATION_CHANGED, RevocationReason: enum value SUPERCEDED, RevocationReason: enum value CESSATION_OF_OPERATION, RevocationReason: enum value CERTIFICATE_HOLD, RevocationReason: enum value REMOVE_FROM_CRL, RevocationReason: enum value PRIVILEGE_WITHDRAWN, RevocationReason: enum value A_A_COMPROMISE, SortBy: enum value CREATED_AT, SortOrder: enum value ASCENDING, SortOrder: enum value DESCENDING, KeyUsageName: enum value DIGITAL_SIGNATURE, KeyUsageName: enum value NON_REPUDIATION, KeyUsageName: enum value KEY_ENCIPHERMENT, KeyUsageName: enum value DATA_ENCIPHERMENT, KeyUsageName: enum value KEY_AGREEMENT, KeyUsageName: enum value CERTIFICATE_SIGNING, KeyUsageName: enum value CRL_SIGNING, KeyUsageName: enum value ENCIPHER_ONLY, KeyUsageName: enum value DECIPHER_ONLY, KeyUsageName: enum value ANY, KeyUsageName: enum value CUSTOM, RecordType: enum value CNAME, CertificateStatus: enum value PENDING_VALIDATION, CertificateStatus: enum value ISSUED, CertificateStatus: enum value INACTIVE, CertificateStatus: enum value EXPIRED, CertificateStatus: enum value VALIDATION_TIMED_OUT, CertificateStatus: enum value REVOKED, CertificateStatus: enum value FAILED
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
      "All operations have clear summaries": false,
      "All parameters have descriptions": false,
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All enums have descriptions": [
        "FailureReason: enum value NO_AVAILABLE_CONTACTS",
        "FailureReason: enum value ADDITIONAL_VERIFICATION_REQUIRED",
        "FailureReason: enum value DOMAIN_NOT_ALLOWED",
        "FailureReason: enum value INVALID_PUBLIC_DOMAIN",
        "FailureReason: enum value DOMAIN_VALIDATION_DENIED",
        "FailureReason: enum value CAA_ERROR",
        "FailureReason: enum value PCA_LIMIT_EXCEEDED",
        "FailureReason: enum value PCA_INVALID_ARN",
        "FailureReason: enum value PCA_INVALID_STATE",
        "FailureReason: enum value PCA_REQUEST_FAILED",
        "FailureReason: enum value PCA_NAME_CONSTRAINTS_VALIDATION",
        "FailureReason: enum value PCA_RESOURCE_NOT_FOUND",
        "FailureReason: enum value PCA_INVALID_ARGS",
        "FailureReason: enum value PCA_INVALID_DURATION",
        "FailureReason: enum value PCA_ACCESS_DENIED",
        "FailureReason: enum value SLR_NOT_FOUND",
        "FailureReason: enum value OTHER",
        "RenewalStatus: enum value PENDING_AUTO_RENEWAL",
        "RenewalStatus: enum value PENDING_VALIDATION",
        "RenewalStatus: enum value SUCCESS",
        "RenewalStatus: enum value FAILED",
        "KeyAlgorithm: enum value RSA_1024",
        "KeyAlgorithm: enum value RSA_2048",
        "KeyAlgorithm: enum value RSA_3072",
        "KeyAlgorithm: enum value RSA_4096",
        "KeyAlgorithm: enum value EC_prime256v1",
        "KeyAlgorithm: enum value EC_secp384r1",
        "KeyAlgorithm: enum value EC_secp521r1",
        "CertificateType: enum value IMPORTED",
        "CertificateType: enum value AMAZON_ISSUED",
        "CertificateType: enum value PRIVATE",
        "CertificateTransparencyLoggingPreference: enum value ENABLED",
        "CertificateTransparencyLoggingPreference: enum value DISABLED",
        "RenewalEligibility: enum value ELIGIBLE",
        "RenewalEligibility: enum value INELIGIBLE",
        "ExtendedKeyUsageName: enum value TLS_WEB_SERVER_AUTHENTICATION",
        "ExtendedKeyUsageName: enum value TLS_WEB_CLIENT_AUTHENTICATION",
        "ExtendedKeyUsageName: enum value CODE_SIGNING",
        "ExtendedKeyUsageName: enum value EMAIL_PROTECTION",
        "ExtendedKeyUsageName: enum value TIME_STAMPING",
        "ExtendedKeyUsageName: enum value OCSP_SIGNING",
        "ExtendedKeyUsageName: enum value IPSEC_END_SYSTEM",
        "ExtendedKeyUsageName: enum value IPSEC_TUNNEL",
        "ExtendedKeyUsageName: enum value IPSEC_USER",
        "ExtendedKeyUsageName: enum value ANY",
        "ExtendedKeyUsageName: enum value NONE",
        "ExtendedKeyUsageName: enum value CUSTOM",
        "DomainStatus: enum value PENDING_VALIDATION",
        "DomainStatus: enum value SUCCESS",
        "DomainStatus: enum value FAILED",
        "ValidationMethod: enum value EMAIL",
        "ValidationMethod: enum value DNS",
        "RevocationReason: enum value UNSPECIFIED",
        "RevocationReason: enum value KEY_COMPROMISE",
        "RevocationReason: enum value CA_COMPROMISE",
        "RevocationReason: enum value AFFILIATION_CHANGED",
        "RevocationReason: enum value SUPERCEDED",
        "RevocationReason: enum value CESSATION_OF_OPERATION",
        "RevocationReason: enum value CERTIFICATE_HOLD",
        "RevocationReason: enum value REMOVE_FROM_CRL",
        "RevocationReason: enum value PRIVILEGE_WITHDRAWN",
        "RevocationReason: enum value A_A_COMPROMISE",
        "SortBy: enum value CREATED_AT",
        "SortOrder: enum value ASCENDING",
        "SortOrder: enum value DESCENDING",
        "KeyUsageName: enum value DIGITAL_SIGNATURE",
        "KeyUsageName: enum value NON_REPUDIATION",
        "KeyUsageName: enum value KEY_ENCIPHERMENT",
        "KeyUsageName: enum value DATA_ENCIPHERMENT",
        "KeyUsageName: enum value KEY_AGREEMENT",
        "KeyUsageName: enum value CERTIFICATE_SIGNING",
        "KeyUsageName: enum value CRL_SIGNING",
        "KeyUsageName: enum value ENCIPHER_ONLY",
        "KeyUsageName: enum value DECIPHER_ONLY",
        "KeyUsageName: enum value ANY",
        "KeyUsageName: enum value CUSTOM",
        "RecordType: enum value CNAME",
        "CertificateStatus: enum value PENDING_VALIDATION",
        "CertificateStatus: enum value ISSUED",
        "CertificateStatus: enum value INACTIVE",
        "CertificateStatus: enum value EXPIRED",
        "CertificateStatus: enum value VALIDATION_TIMED_OUT",
        "CertificateStatus: enum value REVOKED",
        "CertificateStatus: enum value FAILED"
      ],
      "All operations have clear summaries": [
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate"
      ],
      "All parameters have descriptions": [
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Target",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Target"
      ],
      "All request/response bodies have examples": [
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: request body",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 484 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 485 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 484 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 485 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 486 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 484 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 485 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 486 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: request body",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: request body",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 486 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 484 application/json response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 485 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 200 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: request body",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: request body",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: application/json request body",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 485 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 480 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 481 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 482 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 483 application/json response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 484 application/json response"
      ],
      "All schemas have descriptions": [
        "ExtendedKeyUsageFilterList",
        "DomainList",
        "TagList",
        "KeyUsageList",
        "FailureReason",
        "UpdateCertificateOptionsRequest",
        "RenewalStatus",
        "TagKey",
        "ExportCertificateResponse",
        "KeyAlgorithm",
        "MaxItems",
        "DescribeCertificateRequest",
        "DomainValidationList",
        "ResourceNotFoundException",
        "AddTagsToCertificateRequest",
        "ThrottlingException",
        "CertificateStatuses",
        "ExtendedKeyUsageList",
        "InvalidStateException",
        "CertificateSummaryList",
        "CertificateType",
        "CertificateTransparencyLoggingPreference",
        "Arn",
        "RenewalEligibility",
        "TagValue",
        "KeyUsageFilterList",
        "TStamp",
        "InUseList",
        "RequestInProgressException",
        "ImportCertificateRequest",
        "RemoveTagsFromCertificateRequest",
        "CertificateBody",
        "DeleteCertificateRequest",
        "PutAccountConfigurationRequest",
        "ExtendedKeyUsageName",
        "ListTagsForCertificateRequest",
        "NullableBoolean",
        "DomainStatus",
        "TagPolicyException",
        "NextToken",
        "CertificateBodyBlob",
        "TooManyTagsException",
        "ValidationMethod",
        "RevocationReason",
        "ValidationException",
        "ListTagsForCertificateResponse",
        "ImportCertificateResponse",
        "PcaArn",
        "InvalidTagException",
        "ExportCertificateRequest",
        "SortBy",
        "ResendValidationEmailRequest",
        "InvalidArnException",
        "KeyUsageNames",
        "InvalidDomainValidationOptionsException",
        "SortOrder",
        "InvalidArgsException",
        "AccessDeniedException",
        "PassphraseBlob",
        "CertificateChain",
        "String",
        "InvalidParameterException",
        "LimitExceededException",
        "KeyUsageName",
        "ConflictException",
        "RecordType",
        "DescribeCertificateResponse",
        "DomainNameString",
        "ListCertificatesResponse",
        "CertificateStatus",
        "GetCertificateRequest",
        "ListCertificatesRequest",
        "IdempotencyToken",
        "PrivateKey",
        "CertificateChainBlob",
        "RequestCertificateRequest",
        "RenewCertificateRequest",
        "GetCertificateResponse",
        "PrivateKeyBlob",
        "RequestCertificateResponse",
        "GetAccountConfigurationResponse",
        "ValidationEmailList",
        "PositiveInteger",
        "ResourceInUseException",
        "KeyAlgorithmList",
        "DomainValidationOptionList",
        "ExtendedKeyUsageNames"
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
- **Message:** Request validation issues found: All string fields have length constraints: /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter MaxItems, POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter NextToken, POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Target, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Date, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Credential, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Signature, /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Target; All schemas specify data types: POST /#X-Amz-Target=CertificateManager.GetCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions.Options: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.SubjectAlternativeNames: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.KeyAlgorithm: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.Options: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.DomainName: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.ValidationMethod: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.IdempotencyToken: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.Tags: application/json schema, POST /#X-Amz-Target=CertificateManager.RequestCertificate.DomainValidationOptions: application/json schema, POST /#X-Amz-Target=CertificateManager.RenewCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.ExportCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.ExportCertificate.Passphrase: application/json schema, POST /#X-Amz-Target=CertificateManager.DeleteCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate.Tags: application/json schema, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate.Tags: application/json schema, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration.ExpiryEvents: application/json schema, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration.IdempotencyToken: application/json schema, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.Domain: application/json schema, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.ValidationDomain: application/json schema, POST /#X-Amz-Target=CertificateManager.ImportCertificate.CertificateChain: application/json schema, POST /#X-Amz-Target=CertificateManager.ImportCertificate.PrivateKey: application/json schema, POST /#X-Amz-Target=CertificateManager.ImportCertificate.Tags: application/json schema, POST /#X-Amz-Target=CertificateManager.ImportCertificate.Certificate: application/json schema, POST /#X-Amz-Target=CertificateManager.ImportCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.CertificateStatuses: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.Includes: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.MaxItems: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.NextToken: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.SortBy: application/json schema, POST /#X-Amz-Target=CertificateManager.ListCertificates.SortOrder: application/json schema, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=CertificateManager.DescribeCertificate.CertificateArn: application/json schema
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
      "All numeric fields have min/max values": true,
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
      "All schemas specify data types": [
        "POST /#X-Amz-Target=CertificateManager.GetCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions.Options: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.SubjectAlternativeNames: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.KeyAlgorithm: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.Options: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.DomainName: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.ValidationMethod: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.IdempotencyToken: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.Tags: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate.DomainValidationOptions: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate.Passphrase: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate.Tags: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate.Tags: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration.ExpiryEvents: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration.IdempotencyToken: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.Domain: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail.ValidationDomain: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate.CertificateChain: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate.PrivateKey: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate.Tags: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate.Certificate: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.CertificateStatuses: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.Includes: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.MaxItems: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.NextToken: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.SortBy: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates.SortOrder: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate.CertificateArn: application/json schema"
      ],
      "All string fields have length constraints": [
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter MaxItems",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter NextToken",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: parameter X-Amz-Target"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: POST /#X-Amz-Target=CertificateManager.ExportCertificate, POST /#X-Amz-Target=CertificateManager.ImportCertificate, POST /#X-Amz-Target=CertificateManager.ListCertificates, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate, POST /#X-Amz-Target=CertificateManager.RenewCertificate, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration, POST /#X-Amz-Target=CertificateManager.GetCertificate, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate, POST /#X-Amz-Target=CertificateManager.DeleteCertificate, POST /#X-Amz-Target=CertificateManager.DescribeCertificate, POST /#X-Amz-Target=CertificateManager.RequestCertificate; Error responses include error details schema: POST /#X-Amz-Target=CertificateManager.ExportCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.ExportCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 483 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 484 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 485 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 486 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.ImportCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.ListCertificates: 480 response, POST /#X-Amz-Target=CertificateManager.ListCertificates: 481 response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 480 response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 481 response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 482 response, POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 483 response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 483 response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 480 response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 481 response, POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 482 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 483 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 484 response, POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 485 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 484 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 485 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 486 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 483 response, POST /#X-Amz-Target=CertificateManager.RenewCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.RenewCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 483 response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 480 response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 481 response, POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 482 response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.GetCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 480 response, POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 481 response, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 483 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 484 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 485 response, POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 485 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 486 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 480 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 481 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 482 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 483 response, POST /#X-Amz-Target=CertificateManager.RequestCertificate: 484 response
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
      "All operations document 5xx error responses": false,
      "Common error responses are defined in components": false,
      "Error responses follow consistent format": true,
      "Error responses include error details schema": false
    },
    "messages": {
      "Common error responses are defined in components": "No common error responses defined in components"
    },
    "missing_errors": {
      "All operations document 5xx error responses": [
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate"
      ],
      "Error responses include error details schema": [
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.ExportCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 483 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 484 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 485 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 486 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.ImportCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: 480 response",
        "POST /#X-Amz-Target=CertificateManager.ListCertificates: 481 response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 480 response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 481 response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 482 response",
        "POST /#X-Amz-Target=CertificateManager.ResendValidationEmail: 483 response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 483 response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 480 response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 481 response",
        "POST /#X-Amz-Target=CertificateManager.UpdateCertificateOptions: 482 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 483 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 484 response",
        "POST /#X-Amz-Target=CertificateManager.RemoveTagsFromCertificate: 485 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 484 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 485 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 486 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.AddTagsToCertificate: 483 response",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.RenewCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 483 response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 480 response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 481 response",
        "POST /#X-Amz-Target=CertificateManager.PutAccountConfiguration: 482 response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.GetCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 480 response",
        "POST /#X-Amz-Target=CertificateManager.GetAccountConfiguration: 481 response",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.ListTagsForCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 483 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 484 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 485 response",
        "POST /#X-Amz-Target=CertificateManager.DeleteCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.DescribeCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 485 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 486 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 480 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 481 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 482 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 483 response",
        "POST /#X-Amz-Target=CertificateManager.RequestCertificate: 484 response"
      ]
    }
  }
```

**Suggested Fix:**
Add comprehensive error response documentation including codes, messages, and consistent error schemas

---

### Security

#### P005: Security Standards (Passed) [critical]

Validates comprehensive security requirements and authentication mechanisms

- **Status:** Passed
- **Message:** All security requirements are properly defined and consistent
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
      "API keys are properly described": true,
      "Authentication headers are specified": true,
      "Global security requirements are set": true,
      "OAuth2 scopes are documented": true,
      "Operation-level security is defined": true,
      "Security requirements are consistent": true,
      "Security schemes are defined": true
    },
    "messages": {}
  }
```

---

### Versioning

#### P008: API Versioning Strategy (Failed) [warning]

Validates proper API versioning implementation and documentation

- **Status:** Failed
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "2015-12-08" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy
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
      "Version follows semantic versioning": "Version \"2015-12-08\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

