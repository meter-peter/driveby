# API Validation Report

Generated: 2026-05-03T21:05:24+03:00
Environment: production
Version: 1.0.0

## Summary

- Total Checks: 6
- Passed Checks: 2
- Failed Checks: 4
- Critical Issues: 3
- Warnings: 1
- Info: 0

### Categories
- Documentation
- Error Handling
- Schema
- Versioning


### Failed Tags
- validation
- request
- versioning
- quality
- responses
- schema
- compatibility
- lifecycle
- documentation
- usability
- errors
- standards


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
- **Message:** Documentation quality issues found: All enums have descriptions: CertificateAuthorityType: enum value ROOT, CertificateAuthorityType: enum value SUBORDINATE, S3ObjectAcl: enum value PUBLIC_READ, S3ObjectAcl: enum value BUCKET_OWNER_FULL_CONTROL, RevocationReason: enum value UNSPECIFIED, RevocationReason: enum value KEY_COMPROMISE, RevocationReason: enum value CERTIFICATE_AUTHORITY_COMPROMISE, RevocationReason: enum value AFFILIATION_CHANGED, RevocationReason: enum value SUPERSEDED, RevocationReason: enum value CESSATION_OF_OPERATION, RevocationReason: enum value PRIVILEGE_WITHDRAWN, RevocationReason: enum value A_A_COMPROMISE, AuditReportStatus: enum value CREATING, AuditReportStatus: enum value SUCCESS, AuditReportStatus: enum value FAILED, AuditReportResponseFormat: enum value JSON, AuditReportResponseFormat: enum value CSV, SigningAlgorithm: enum value SHA256WITHECDSA, SigningAlgorithm: enum value SHA384WITHECDSA, SigningAlgorithm: enum value SHA512WITHECDSA, SigningAlgorithm: enum value SHA256WITHRSA, SigningAlgorithm: enum value SHA384WITHRSA, SigningAlgorithm: enum value SHA512WITHRSA, KeyStorageSecurityStandard: enum value FIPS_140_2_LEVEL_2_OR_HIGHER, KeyStorageSecurityStandard: enum value FIPS_140_2_LEVEL_3_OR_HIGHER, ExtendedKeyUsageType: enum value SERVER_AUTH, ExtendedKeyUsageType: enum value CLIENT_AUTH, ExtendedKeyUsageType: enum value CODE_SIGNING, ExtendedKeyUsageType: enum value EMAIL_PROTECTION, ExtendedKeyUsageType: enum value TIME_STAMPING, ExtendedKeyUsageType: enum value OCSP_SIGNING, ExtendedKeyUsageType: enum value SMART_CARD_LOGIN, ExtendedKeyUsageType: enum value DOCUMENT_SIGNING, ExtendedKeyUsageType: enum value CERTIFICATE_TRANSPARENCY, ResourceOwner: enum value SELF, ResourceOwner: enum value OTHER_ACCOUNTS, FailureReason: enum value REQUEST_TIMED_OUT, FailureReason: enum value UNSUPPORTED_ALGORITHM, FailureReason: enum value OTHER, AccessMethodType: enum value CA_REPOSITORY, AccessMethodType: enum value RESOURCE_PKI_MANIFEST, AccessMethodType: enum value RESOURCE_PKI_NOTIFY, CertificateAuthorityUsageMode: enum value GENERAL_PURPOSE, CertificateAuthorityUsageMode: enum value SHORT_LIVED_CERTIFICATE, ValidityPeriodType: enum value END_DATE, ValidityPeriodType: enum value ABSOLUTE, ValidityPeriodType: enum value DAYS, ValidityPeriodType: enum value MONTHS, ValidityPeriodType: enum value YEARS, KeyAlgorithm: enum value RSA_2048, KeyAlgorithm: enum value RSA_4096, KeyAlgorithm: enum value EC_prime256v1, KeyAlgorithm: enum value EC_secp384r1, CertificateAuthorityStatus: enum value CREATING, CertificateAuthorityStatus: enum value PENDING_CERTIFICATE, CertificateAuthorityStatus: enum value ACTIVE, CertificateAuthorityStatus: enum value DELETED, CertificateAuthorityStatus: enum value DISABLED, CertificateAuthorityStatus: enum value EXPIRED, CertificateAuthorityStatus: enum value FAILED, ActionType: enum value IssueCertificate, ActionType: enum value GetCertificate, ActionType: enum value ListPermissions, PolicyQualifierId: enum value CPS; All operations have clear summaries: POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate, POST /#X-Amz-Target=ACMPrivateCA.ListTags, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport; All parameters have descriptions: POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Target, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Target; All request/response bodies have examples: POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: request body, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: request body, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 486 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 487 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 488 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: request body, POST /#X-Amz-Target=ACMPrivateCA.ListTags: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: request body, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: request body, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: request body, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: request body, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: request body, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: request body, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: request body, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: request body, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 486 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: request body, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 485 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 487 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 488 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 486 application/json response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: request body, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 483 application/json response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 484 application/json response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 480 application/json response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: request body, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: application/json request body, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 481 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 482 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 200 application/json response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 480 application/json response; All schemas have descriptions: DescribeCertificateAuthorityAuditReportResponse, CertificateChainBlob, S3BucketName, UpdateCertificateAuthorityRequest, CreateCertificateAuthorityRequest, CertificateBody, GetCertificateResponse, CustomExtensionList, GetPolicyRequest, CustomObjectIdentifier, TStamp, CertificateAuthorityType, String64, ListCertificateAuthoritiesRequest, DescribeCertificateAuthorityRequest, S3ObjectAcl, RevocationReason, AuditReportId, AuditReportStatus, RevokeCertificateRequest, InvalidNextTokenException, CertificateMismatchException, TagKey, String3, ActionList, NextToken, AuditReportResponseFormat, IssueCertificateRequest, CertificatePolicyList, RestoreCertificateAuthorityRequest, RequestAlreadyProcessedException, CreateCertificateAuthorityResponse, UntagCertificateAuthorityRequest, CountryCodeString, InvalidArnException, SigningAlgorithm, GetCertificateAuthorityCsrResponse, KeyStorageSecurityStandard, PutPolicyRequest, DeletePermissionRequest, PolicyQualifierInfoList, DeletePolicyRequest, CustomAttributeList, Boolean, DescribeCertificateAuthorityResponse, ExtendedKeyUsageType, String128, MalformedCSRException, String, GeneralNameList, GetCertificateAuthorityCertificateRequest, PositiveLong, CreateCertificateAuthorityAuditReportRequest, AccessDescriptionList, GetCertificateRequest, DescribeCertificateAuthorityAuditReportRequest, TagList, TagCertificateAuthorityRequest, IssueCertificateResponse, Base64String1To4096, CertificateAuthorities, ListCertificateAuthoritiesResponse, ListTagsRequest, AccountId, TooManyTagsException, ResourceOwner, ConcurrentModificationException, Integer1To5000, MaxResults, RequestFailedException, String16, CertificateBodyBlob, TagValue, FailureReason, IdempotencyToken, String39, AccessMethodType, CertificateAuthorityUsageMode, InvalidStateException, String253, InvalidRequestException, PermanentDeletionTimeInDays, AWSPolicy, ValidityPeriodType, ExtendedKeyUsageList, InvalidTagException, CsrBody, String1To256, GetPolicyResponse, KeyAlgorithm, ImportCertificateAuthorityCertificateRequest, MalformedCertificateException, CnameString, CreateCertificateAuthorityAuditReportResponse, DeleteCertificateAuthorityRequest, CertificateAuthorityStatus, Arn, S3BucketName3To255, LimitExceededException, PermissionList, ListPermissionsResponse, InvalidArgsException, ActionType, CsrBlob, PermissionAlreadyExistsException, S3Key, ASN1PrintableString64, String256, Principal, CreatePermissionRequest, InvalidPolicyException, PolicyQualifierId, ListPermissionsRequest, ResourceNotFoundException, GetCertificateAuthorityCsrRequest, String5, GetCertificateAuthorityCertificateResponse, String40, RequestInProgressException, LockoutPreventedException, CertificateChain, ListTagsResponse
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
        "CertificateAuthorityType: enum value ROOT",
        "CertificateAuthorityType: enum value SUBORDINATE",
        "S3ObjectAcl: enum value PUBLIC_READ",
        "S3ObjectAcl: enum value BUCKET_OWNER_FULL_CONTROL",
        "RevocationReason: enum value UNSPECIFIED",
        "RevocationReason: enum value KEY_COMPROMISE",
        "RevocationReason: enum value CERTIFICATE_AUTHORITY_COMPROMISE",
        "RevocationReason: enum value AFFILIATION_CHANGED",
        "RevocationReason: enum value SUPERSEDED",
        "RevocationReason: enum value CESSATION_OF_OPERATION",
        "RevocationReason: enum value PRIVILEGE_WITHDRAWN",
        "RevocationReason: enum value A_A_COMPROMISE",
        "AuditReportStatus: enum value CREATING",
        "AuditReportStatus: enum value SUCCESS",
        "AuditReportStatus: enum value FAILED",
        "AuditReportResponseFormat: enum value JSON",
        "AuditReportResponseFormat: enum value CSV",
        "SigningAlgorithm: enum value SHA256WITHECDSA",
        "SigningAlgorithm: enum value SHA384WITHECDSA",
        "SigningAlgorithm: enum value SHA512WITHECDSA",
        "SigningAlgorithm: enum value SHA256WITHRSA",
        "SigningAlgorithm: enum value SHA384WITHRSA",
        "SigningAlgorithm: enum value SHA512WITHRSA",
        "KeyStorageSecurityStandard: enum value FIPS_140_2_LEVEL_2_OR_HIGHER",
        "KeyStorageSecurityStandard: enum value FIPS_140_2_LEVEL_3_OR_HIGHER",
        "ExtendedKeyUsageType: enum value SERVER_AUTH",
        "ExtendedKeyUsageType: enum value CLIENT_AUTH",
        "ExtendedKeyUsageType: enum value CODE_SIGNING",
        "ExtendedKeyUsageType: enum value EMAIL_PROTECTION",
        "ExtendedKeyUsageType: enum value TIME_STAMPING",
        "ExtendedKeyUsageType: enum value OCSP_SIGNING",
        "ExtendedKeyUsageType: enum value SMART_CARD_LOGIN",
        "ExtendedKeyUsageType: enum value DOCUMENT_SIGNING",
        "ExtendedKeyUsageType: enum value CERTIFICATE_TRANSPARENCY",
        "ResourceOwner: enum value SELF",
        "ResourceOwner: enum value OTHER_ACCOUNTS",
        "FailureReason: enum value REQUEST_TIMED_OUT",
        "FailureReason: enum value UNSUPPORTED_ALGORITHM",
        "FailureReason: enum value OTHER",
        "AccessMethodType: enum value CA_REPOSITORY",
        "AccessMethodType: enum value RESOURCE_PKI_MANIFEST",
        "AccessMethodType: enum value RESOURCE_PKI_NOTIFY",
        "CertificateAuthorityUsageMode: enum value GENERAL_PURPOSE",
        "CertificateAuthorityUsageMode: enum value SHORT_LIVED_CERTIFICATE",
        "ValidityPeriodType: enum value END_DATE",
        "ValidityPeriodType: enum value ABSOLUTE",
        "ValidityPeriodType: enum value DAYS",
        "ValidityPeriodType: enum value MONTHS",
        "ValidityPeriodType: enum value YEARS",
        "KeyAlgorithm: enum value RSA_2048",
        "KeyAlgorithm: enum value RSA_4096",
        "KeyAlgorithm: enum value EC_prime256v1",
        "KeyAlgorithm: enum value EC_secp384r1",
        "CertificateAuthorityStatus: enum value CREATING",
        "CertificateAuthorityStatus: enum value PENDING_CERTIFICATE",
        "CertificateAuthorityStatus: enum value ACTIVE",
        "CertificateAuthorityStatus: enum value DELETED",
        "CertificateAuthorityStatus: enum value DISABLED",
        "CertificateAuthorityStatus: enum value EXPIRED",
        "CertificateAuthorityStatus: enum value FAILED",
        "ActionType: enum value IssueCertificate",
        "ActionType: enum value GetCertificate",
        "ActionType: enum value ListPermissions",
        "PolicyQualifierId: enum value CPS"
      ],
      "All operations have clear summaries": [
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport"
      ],
      "All parameters have descriptions": [
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Target",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Target"
      ],
      "All request/response bodies have examples": [
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 486 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 487 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 488 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 486 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 485 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 487 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 488 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 486 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 483 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 484 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 480 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: application/json request body",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 481 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 482 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 200 application/json response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 480 application/json response"
      ],
      "All schemas have descriptions": [
        "DescribeCertificateAuthorityAuditReportResponse",
        "CertificateChainBlob",
        "S3BucketName",
        "UpdateCertificateAuthorityRequest",
        "CreateCertificateAuthorityRequest",
        "CertificateBody",
        "GetCertificateResponse",
        "CustomExtensionList",
        "GetPolicyRequest",
        "CustomObjectIdentifier",
        "TStamp",
        "CertificateAuthorityType",
        "String64",
        "ListCertificateAuthoritiesRequest",
        "DescribeCertificateAuthorityRequest",
        "S3ObjectAcl",
        "RevocationReason",
        "AuditReportId",
        "AuditReportStatus",
        "RevokeCertificateRequest",
        "InvalidNextTokenException",
        "CertificateMismatchException",
        "TagKey",
        "String3",
        "ActionList",
        "NextToken",
        "AuditReportResponseFormat",
        "IssueCertificateRequest",
        "CertificatePolicyList",
        "RestoreCertificateAuthorityRequest",
        "RequestAlreadyProcessedException",
        "CreateCertificateAuthorityResponse",
        "UntagCertificateAuthorityRequest",
        "CountryCodeString",
        "InvalidArnException",
        "SigningAlgorithm",
        "GetCertificateAuthorityCsrResponse",
        "KeyStorageSecurityStandard",
        "PutPolicyRequest",
        "DeletePermissionRequest",
        "PolicyQualifierInfoList",
        "DeletePolicyRequest",
        "CustomAttributeList",
        "Boolean",
        "DescribeCertificateAuthorityResponse",
        "ExtendedKeyUsageType",
        "String128",
        "MalformedCSRException",
        "String",
        "GeneralNameList",
        "GetCertificateAuthorityCertificateRequest",
        "PositiveLong",
        "CreateCertificateAuthorityAuditReportRequest",
        "AccessDescriptionList",
        "GetCertificateRequest",
        "DescribeCertificateAuthorityAuditReportRequest",
        "TagList",
        "TagCertificateAuthorityRequest",
        "IssueCertificateResponse",
        "Base64String1To4096",
        "CertificateAuthorities",
        "ListCertificateAuthoritiesResponse",
        "ListTagsRequest",
        "AccountId",
        "TooManyTagsException",
        "ResourceOwner",
        "ConcurrentModificationException",
        "Integer1To5000",
        "MaxResults",
        "RequestFailedException",
        "String16",
        "CertificateBodyBlob",
        "TagValue",
        "FailureReason",
        "IdempotencyToken",
        "String39",
        "AccessMethodType",
        "CertificateAuthorityUsageMode",
        "InvalidStateException",
        "String253",
        "InvalidRequestException",
        "PermanentDeletionTimeInDays",
        "AWSPolicy",
        "ValidityPeriodType",
        "ExtendedKeyUsageList",
        "InvalidTagException",
        "CsrBody",
        "String1To256",
        "GetPolicyResponse",
        "KeyAlgorithm",
        "ImportCertificateAuthorityCertificateRequest",
        "MalformedCertificateException",
        "CnameString",
        "CreateCertificateAuthorityAuditReportResponse",
        "DeleteCertificateAuthorityRequest",
        "CertificateAuthorityStatus",
        "Arn",
        "S3BucketName3To255",
        "LimitExceededException",
        "PermissionList",
        "ListPermissionsResponse",
        "InvalidArgsException",
        "ActionType",
        "CsrBlob",
        "PermissionAlreadyExistsException",
        "S3Key",
        "ASN1PrintableString64",
        "String256",
        "Principal",
        "CreatePermissionRequest",
        "InvalidPolicyException",
        "PolicyQualifierId",
        "ListPermissionsRequest",
        "ResourceNotFoundException",
        "GetCertificateAuthorityCsrRequest",
        "String5",
        "GetCertificateAuthorityCertificateResponse",
        "String40",
        "RequestInProgressException",
        "LockoutPreventedException",
        "CertificateChain",
        "ListTagsResponse"
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
- **Message:** Request validation issues found: All string fields have length constraints: /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter MaxResults, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter NextToken, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter MaxResults, POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter NextToken, POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter MaxResults, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter NextToken, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Target, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Content-Sha256, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Date, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Algorithm, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Credential, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Security-Token, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Signature, /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Target; All schemas specify data types: POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.CertificateSerial: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.RevocationReason: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.MaxResults: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.NextToken: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListTags.MaxResults: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListTags.NextToken: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListTags.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority.Tags: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy.ResourceArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.S3BucketName: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.AuditReportResponseFormat: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.Tags: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.UsageMode: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.CertificateAuthorityConfiguration: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.CertificateAuthorityType: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.IdempotencyToken: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.KeyStorageSecurityStandard: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.RevocationConfiguration: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.ResourceOwner: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.MaxResults: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.NextToken: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.Principal: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.SourceAccount: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.Actions: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate.CertificateArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.TemplateArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.Validity: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.ValidityNotBefore: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.ApiPassthrough: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.Csr: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.IdempotencyToken: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.SigningAlgorithm: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority.Tags: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.Certificate: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.CertificateChain: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy.Policy: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy.ResourceArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport.AuditReportId: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority.PermanentDeletionTimeInDays: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.Principal: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.SourceAccount: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.CertificateAuthorityArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.RevocationConfiguration: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.Status: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy.ResourceArn: application/json schema, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr.CertificateAuthorityArn: application/json schema
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
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.CertificateSerial: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate.RevocationReason: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.MaxResults: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.NextToken: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags.MaxResults: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags.NextToken: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority.Tags: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy.ResourceArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.S3BucketName: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.AuditReportResponseFormat: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.Tags: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.UsageMode: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.CertificateAuthorityConfiguration: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.CertificateAuthorityType: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.IdempotencyToken: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.KeyStorageSecurityStandard: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority.RevocationConfiguration: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.ResourceOwner: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.MaxResults: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities.NextToken: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.Principal: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.SourceAccount: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission.Actions: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate.CertificateArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.TemplateArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.Validity: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.ValidityNotBefore: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.ApiPassthrough: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.Csr: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.IdempotencyToken: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate.SigningAlgorithm: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority.Tags: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.Certificate: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate.CertificateChain: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy.Policy: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy.ResourceArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport.AuditReportId: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority.PermanentDeletionTimeInDays: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.Principal: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission.SourceAccount: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.CertificateAuthorityArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.RevocationConfiguration: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority.Status: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy.ResourceArn: application/json schema",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr.CertificateAuthorityArn: application/json schema"
      ],
      "All string fields have length constraints": [
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter MaxResults",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter NextToken",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter MaxResults",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter NextToken",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter MaxResults",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter NextToken",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: parameter X-Amz-Target",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Date",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Credential",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Signature",
        "/#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: parameter X-Amz-Target"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate, POST /#X-Amz-Target=ACMPrivateCA.ListTags, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority; Error responses include error details schema: POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 484 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 480 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 481 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 482 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 483 response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 483 response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 483 response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 481 response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 482 response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 480 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 485 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 480 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 481 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 482 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 483 response, POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 484 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 482 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 483 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 484 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 485 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 480 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 481 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 487 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 481 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 488 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 484 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 486 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 480 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 482 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 485 response, POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 483 response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 480 response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 481 response, POST /#X-Amz-Target=ACMPrivateCA.ListTags: 482 response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 484 response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 483 response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 481 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 482 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 483 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 484 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 485 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 480 response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 483 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 480 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 481 response, POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 482 response, POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 480 response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 480 response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 481 response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 482 response, POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 483 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 482 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 483 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 484 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 480 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 481 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 485 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 480 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 481 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 482 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 483 response, POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 484 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 486 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 487 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 481 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 482 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 483 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 485 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 488 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 480 response, POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 484 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 483 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 484 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 485 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 486 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 480 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 481 response, POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 482 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 481 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 483 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 480 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 481 response, POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 482 response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 481 response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 482 response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 483 response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 484 response, POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 480 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 482 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 483 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 484 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 485 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 480 response, POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 481 response
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
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority"
      ],
      "Error responses include error details schema": [
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCsr: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UntagCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeleteCertificateAuthority: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthorityAuditReport: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreatePermission: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePolicy: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 487 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 488 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 486 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RevokeCertificate: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListTags: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.TagCertificateAuthority: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DescribeCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthorityAuditReport: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.RestoreCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.DeletePermission: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListCertificateAuthorities: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetPolicy: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificate: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.IssueCertificate: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 486 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 487 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 488 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ImportCertificateAuthorityCertificate: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 486 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.PutPolicy: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.CreateCertificateAuthority: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.GetCertificateAuthorityCertificate: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 481 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.ListPermissions: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 482 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 483 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 484 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 485 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 480 response",
        "POST /#X-Amz-Target=ACMPrivateCA.UpdateCertificateAuthority: 481 response"
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
- **Message:** Versioning validation failed: Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "2017-08-22" does not match semver format (expected MAJOR.MINOR.PATCH)
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
      "Version follows semantic versioning": "Version \"2017-08-22\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

