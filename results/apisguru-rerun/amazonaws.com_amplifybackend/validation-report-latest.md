# API Validation Report

Generated: 2026-05-03T21:05:25+03:00
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
- Schema
- Versioning
- Documentation
- Error Handling


### Failed Tags
- versioning
- standards
- schema
- validation
- request
- compatibility
- lifecycle
- documentation
- quality
- usability
- errors
- responses


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /backend/{appId}/api/{backendEnvironmentName}/remove: request body, POST /backend/{appId}/api/{backendEnvironmentName}/remove: application/json request body, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 200 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 480 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 483 application/json response, POST /backend/{appId}/details: request body, POST /backend/{appId}/details: application/json request body, POST /backend/{appId}/details: 481 application/json response, POST /backend/{appId}/details: 482 application/json response, POST /backend/{appId}/details: 483 application/json response, POST /backend/{appId}/details: 200 application/json response, POST /backend/{appId}/details: 480 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: request body, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: application/json request body, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 200 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 480 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 481 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 482 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 483 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}: request body, POST /backend/{appId}/auth/{backendEnvironmentName}: application/json request body, POST /backend/{appId}/auth/{backendEnvironmentName}: 482 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}: 483 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}: 200 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}: 480 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}: 481 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: request body, POST /backend/{appId}/storage/{backendEnvironmentName}/details: application/json request body, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 483 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 200 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 480 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 481 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: request body, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: application/json request body, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 480 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 483 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 200 application/json response, POST /backend/{appId}/config: request body, POST /backend/{appId}/config: application/json request body, POST /backend/{appId}/config: 200 application/json response, POST /backend/{appId}/config: 480 application/json response, POST /backend/{appId}/config: 481 application/json response, POST /backend/{appId}/config: 482 application/json response, POST /backend/{appId}/config: 483 application/json response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 application/json response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 application/json response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 application/json response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 200 application/json response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: request body, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: application/json request body, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 200 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 application/json response, POST /backend/{appId}/storage: request body, POST /backend/{appId}/storage: application/json request body, POST /backend/{appId}/storage: 483 application/json response, POST /backend/{appId}/storage: 200 application/json response, POST /backend/{appId}/storage: 480 application/json response, POST /backend/{appId}/storage: 481 application/json response, POST /backend/{appId}/storage: 482 application/json response, POST /backend: request body, POST /backend: application/json request body, POST /backend: 480 application/json response, POST /backend: 481 application/json response, POST /backend: 482 application/json response, POST /backend: 483 application/json response, POST /backend: 200 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: request body, POST /backend/{appId}/auth/{backendEnvironmentName}/import: application/json request body, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 482 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 483 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 200 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 480 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 481 application/json response, POST /backend/{appId}/challenge: 480 application/json response, POST /backend/{appId}/challenge: 481 application/json response, POST /backend/{appId}/challenge: 482 application/json response, POST /backend/{appId}/challenge: 483 application/json response, POST /backend/{appId}/challenge: 200 application/json response, GET /backend/{appId}/challenge/{sessionId}: 482 application/json response, GET /backend/{appId}/challenge/{sessionId}: 483 application/json response, GET /backend/{appId}/challenge/{sessionId}: 200 application/json response, GET /backend/{appId}/challenge/{sessionId}: 480 application/json response, GET /backend/{appId}/challenge/{sessionId}: 481 application/json response, POST /s3Buckets: request body, POST /s3Buckets: application/json request body, POST /s3Buckets: 483 application/json response, POST /s3Buckets: 200 application/json response, POST /s3Buckets: 480 application/json response, POST /s3Buckets: 481 application/json response, POST /s3Buckets: 482 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: request body, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: application/json request body, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 481 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 482 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 483 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 200 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 480 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: request body, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: application/json request body, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 200 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 480 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 483 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/details: request body, POST /backend/{appId}/api/{backendEnvironmentName}/details: application/json request body, POST /backend/{appId}/api/{backendEnvironmentName}/details: 480 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 483 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 200 application/json response, POST /backend/{appId}/config/remove: 483 application/json response, POST /backend/{appId}/config/remove: 200 application/json response, POST /backend/{appId}/config/remove: 480 application/json response, POST /backend/{appId}/config/remove: 481 application/json response, POST /backend/{appId}/config/remove: 482 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}: request body, POST /backend/{appId}/storage/{backendEnvironmentName}: application/json request body, POST /backend/{appId}/storage/{backendEnvironmentName}: 483 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}: 200 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}: 480 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}: 481 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}: 482 application/json response, POST /backend/{appId}/config/update: request body, POST /backend/{appId}/config/update: application/json request body, POST /backend/{appId}/config/update: 482 application/json response, POST /backend/{appId}/config/update: 483 application/json response, POST /backend/{appId}/config/update: 200 application/json response, POST /backend/{appId}/config/update: 480 application/json response, POST /backend/{appId}/config/update: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}: request body, POST /backend/{appId}/api/{backendEnvironmentName}: application/json request body, POST /backend/{appId}/api/{backendEnvironmentName}: 481 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}: 482 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}: 483 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}: 200 application/json response, POST /backend/{appId}/api/{backendEnvironmentName}: 480 application/json response, POST /backend/{appId}/api: request body, POST /backend/{appId}/api: application/json request body, POST /backend/{appId}/api: 200 application/json response, POST /backend/{appId}/api: 480 application/json response, POST /backend/{appId}/api: 481 application/json response, POST /backend/{appId}/api: 482 application/json response, POST /backend/{appId}/api: 483 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 200 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 480 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 481 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 482 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 483 application/json response, POST /backend/{appId}/remove: request body, POST /backend/{appId}/remove: application/json request body, POST /backend/{appId}/remove: 483 application/json response, POST /backend/{appId}/remove: 200 application/json response, POST /backend/{appId}/remove: 480 application/json response, POST /backend/{appId}/remove: 481 application/json response, POST /backend/{appId}/remove: 482 application/json response, POST /backend/{appId}/challenge/{sessionId}/remove: 200 application/json response, POST /backend/{appId}/challenge/{sessionId}/remove: 480 application/json response, POST /backend/{appId}/challenge/{sessionId}/remove: 481 application/json response, POST /backend/{appId}/challenge/{sessionId}/remove: 482 application/json response, POST /backend/{appId}/challenge/{sessionId}/remove: 483 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}: request body, POST /backend/{appId}/job/{backendEnvironmentName}: application/json request body, POST /backend/{appId}/job/{backendEnvironmentName}: 481 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}: 482 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}: 483 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}: 200 application/json response, POST /backend/{appId}/job/{backendEnvironmentName}: 480 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: request body, POST /backend/{appId}/storage/{backendEnvironmentName}/import: application/json request body, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 200 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 480 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 481 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 482 application/json response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 483 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: request body, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: application/json request body, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 483 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 200 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 480 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 481 application/json response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 482 application/json response, POST /backend/{appId}/auth: request body, POST /backend/{appId}/auth: application/json request body, POST /backend/{appId}/auth: 482 application/json response, POST /backend/{appId}/auth: 483 application/json response, POST /backend/{appId}/auth: 200 application/json response, POST /backend/{appId}/auth: 480 application/json response, POST /backend/{appId}/auth: 481 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: request body, POST /backend/{appId}/auth/{backendEnvironmentName}/details: application/json request body, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 481 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 482 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 483 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 200 application/json response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 480 application/json response; All schemas have descriptions: GetTokenRequest, ListOfAdditionalConstraintsElement, RequiredSignUpAttributesElement, UpdateBackendAPIResponse, AuthenticatedElement, OAuthScopesElement, AuthResources, DeleteBackendRequest, ListOfUnAuthenticatedElement, MfaTypesElement, GetTokenResponse, GetBackendAPIModelsResponse, UpdateBackendJobResponse, UpdateBackendStorageResponse, GatewayTimeoutException, ListOfAuthenticatedElement, Service, ListOfBackendJobRespObj, __string, DeleteBackendResponse, __double, DeleteBackendAuthResponse, RemoveBackendConfigRequest, OAuthGrantType, GenerateBackendAPIModelsResponse, RemoveAllBackendsResponse, AdditionalConstraintsElement, CloneBackendResponse, ListBackendJobsResponse, ImportBackendAuthResponse, RemoveBackendConfigResponse, TooManyRequestsException, CreateBackendAPIResponse, MFAMode, BadRequestException, CreateBackendConfigResponse, __integerMin1Max25, ListOfRequiredSignUpAttributesElement, ResolutionStrategy, CreateTokenRequest, GetBackendJobRequest, GetBackendJobResponse, DeleteTokenResponse, DeleteBackendStorageResponse, ListS3BucketsResponse, CreateBackendResponse, NotFoundException, __boolean, CreateBackendStorageResponse, Status, ImportBackendStorageResponse, DeleteBackendAPIResponse, GetBackendResponse, GetBackendStorageResponse, GetBackendAuthResponse, ListOfS3BucketInfo, Mode, UpdateBackendAuthResponse, ListOfMfaTypesElement, ServiceName, SignInMethod, CreateBackendAuthResponse, ListOfOAuthScopesElement, DeleteTokenRequest, CreateTokenResponse, ListOf__string, UnAuthenticatedElement, UpdateBackendConfigResponse, GetBackendAPIResponse, ListOfBackendAPIAuthType; All enums have descriptions: RequiredSignUpAttributesElement: enum value ADDRESS, RequiredSignUpAttributesElement: enum value BIRTHDATE, RequiredSignUpAttributesElement: enum value EMAIL, RequiredSignUpAttributesElement: enum value FAMILY_NAME, RequiredSignUpAttributesElement: enum value GENDER, RequiredSignUpAttributesElement: enum value GIVEN_NAME, RequiredSignUpAttributesElement: enum value LOCALE, RequiredSignUpAttributesElement: enum value MIDDLE_NAME, RequiredSignUpAttributesElement: enum value NAME, RequiredSignUpAttributesElement: enum value NICKNAME, RequiredSignUpAttributesElement: enum value PHONE_NUMBER, RequiredSignUpAttributesElement: enum value PICTURE, RequiredSignUpAttributesElement: enum value PREFERRED_USERNAME, RequiredSignUpAttributesElement: enum value PROFILE, RequiredSignUpAttributesElement: enum value UPDATED_AT, RequiredSignUpAttributesElement: enum value WEBSITE, RequiredSignUpAttributesElement: enum value ZONE_INFO, AuthenticatedElement: enum value READ, AuthenticatedElement: enum value CREATE_AND_UPDATE, AuthenticatedElement: enum value DELETE, OAuthScopesElement: enum value PHONE, OAuthScopesElement: enum value EMAIL, OAuthScopesElement: enum value OPENID, OAuthScopesElement: enum value PROFILE, OAuthScopesElement: enum value AWS_COGNITO_SIGNIN_USER_ADMIN, AuthResources: enum value USER_POOL_ONLY, AuthResources: enum value IDENTITY_POOL_AND_USER_POOL, DeliveryMethod: enum value EMAIL, DeliveryMethod: enum value SMS, MfaTypesElement: enum value SMS, MfaTypesElement: enum value TOTP, Service: enum value COGNITO, OAuthGrantType: enum value CODE, OAuthGrantType: enum value IMPLICIT, AdditionalConstraintsElement: enum value REQUIRE_DIGIT, AdditionalConstraintsElement: enum value REQUIRE_LOWERCASE, AdditionalConstraintsElement: enum value REQUIRE_SYMBOL, AdditionalConstraintsElement: enum value REQUIRE_UPPERCASE, MFAMode: enum value ON, MFAMode: enum value OFF, MFAMode: enum value OPTIONAL, ResolutionStrategy: enum value OPTIMISTIC_CONCURRENCY, ResolutionStrategy: enum value LAMBDA, ResolutionStrategy: enum value AUTOMERGE, ResolutionStrategy: enum value NONE, Status: enum value LATEST, Status: enum value STALE, Mode: enum value API_KEY, Mode: enum value AWS_IAM, Mode: enum value AMAZON_COGNITO_USER_POOLS, Mode: enum value OPENID_CONNECT, ServiceName: enum value S3, SignInMethod: enum value EMAIL, SignInMethod: enum value EMAIL_AND_PHONE_NUMBER, SignInMethod: enum value PHONE_NUMBER, SignInMethod: enum value USERNAME, UnAuthenticatedElement: enum value READ, UnAuthenticatedElement: enum value CREATE_AND_UPDATE, UnAuthenticatedElement: enum value DELETE; All operations have clear summaries: POST /backend/{appId}/api/{backendEnvironmentName}/remove, POST /backend/{appId}/details, POST /backend/{appId}/auth/{backendEnvironmentName}/remove, POST /backend/{appId}/auth/{backendEnvironmentName}, POST /backend/{appId}/storage/{backendEnvironmentName}/details, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels, POST /backend/{appId}/config, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}, POST /backend/{appId}/storage, POST /backend, POST /backend/{appId}/auth/{backendEnvironmentName}/import, POST /backend/{appId}/challenge, GET /backend/{appId}/challenge/{sessionId}, POST /s3Buckets, POST /backend/{appId}/storage/{backendEnvironmentName}/remove, POST /backend/{appId}/api/{backendEnvironmentName}/getModels, POST /backend/{appId}/api/{backendEnvironmentName}/details, POST /backend/{appId}/config/remove, POST /backend/{appId}/storage/{backendEnvironmentName}, POST /backend/{appId}/config/update, POST /backend/{appId}/api/{backendEnvironmentName}, POST /backend/{appId}/api, POST /backend/{appId}/environments/{backendEnvironmentName}/remove, POST /backend/{appId}/remove, POST /backend/{appId}/challenge/{sessionId}/remove, POST /backend/{appId}/job/{backendEnvironmentName}, POST /backend/{appId}/storage/{backendEnvironmentName}/import, POST /backend/{appId}/environments/{backendEnvironmentName}/clone, POST /backend/{appId}/auth, POST /backend/{appId}/auth/{backendEnvironmentName}/details
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
      "All request/response bodies have examples": false,
      "All schemas have descriptions": false,
      "Contact information is provided": true,
      "License information is provided": true
    },
    "messages": {},
    "missing_docs": {
      "All enums have descriptions": [
        "RequiredSignUpAttributesElement: enum value ADDRESS",
        "RequiredSignUpAttributesElement: enum value BIRTHDATE",
        "RequiredSignUpAttributesElement: enum value EMAIL",
        "RequiredSignUpAttributesElement: enum value FAMILY_NAME",
        "RequiredSignUpAttributesElement: enum value GENDER",
        "RequiredSignUpAttributesElement: enum value GIVEN_NAME",
        "RequiredSignUpAttributesElement: enum value LOCALE",
        "RequiredSignUpAttributesElement: enum value MIDDLE_NAME",
        "RequiredSignUpAttributesElement: enum value NAME",
        "RequiredSignUpAttributesElement: enum value NICKNAME",
        "RequiredSignUpAttributesElement: enum value PHONE_NUMBER",
        "RequiredSignUpAttributesElement: enum value PICTURE",
        "RequiredSignUpAttributesElement: enum value PREFERRED_USERNAME",
        "RequiredSignUpAttributesElement: enum value PROFILE",
        "RequiredSignUpAttributesElement: enum value UPDATED_AT",
        "RequiredSignUpAttributesElement: enum value WEBSITE",
        "RequiredSignUpAttributesElement: enum value ZONE_INFO",
        "AuthenticatedElement: enum value READ",
        "AuthenticatedElement: enum value CREATE_AND_UPDATE",
        "AuthenticatedElement: enum value DELETE",
        "OAuthScopesElement: enum value PHONE",
        "OAuthScopesElement: enum value EMAIL",
        "OAuthScopesElement: enum value OPENID",
        "OAuthScopesElement: enum value PROFILE",
        "OAuthScopesElement: enum value AWS_COGNITO_SIGNIN_USER_ADMIN",
        "AuthResources: enum value USER_POOL_ONLY",
        "AuthResources: enum value IDENTITY_POOL_AND_USER_POOL",
        "DeliveryMethod: enum value EMAIL",
        "DeliveryMethod: enum value SMS",
        "MfaTypesElement: enum value SMS",
        "MfaTypesElement: enum value TOTP",
        "Service: enum value COGNITO",
        "OAuthGrantType: enum value CODE",
        "OAuthGrantType: enum value IMPLICIT",
        "AdditionalConstraintsElement: enum value REQUIRE_DIGIT",
        "AdditionalConstraintsElement: enum value REQUIRE_LOWERCASE",
        "AdditionalConstraintsElement: enum value REQUIRE_SYMBOL",
        "AdditionalConstraintsElement: enum value REQUIRE_UPPERCASE",
        "MFAMode: enum value ON",
        "MFAMode: enum value OFF",
        "MFAMode: enum value OPTIONAL",
        "ResolutionStrategy: enum value OPTIMISTIC_CONCURRENCY",
        "ResolutionStrategy: enum value LAMBDA",
        "ResolutionStrategy: enum value AUTOMERGE",
        "ResolutionStrategy: enum value NONE",
        "Status: enum value LATEST",
        "Status: enum value STALE",
        "Mode: enum value API_KEY",
        "Mode: enum value AWS_IAM",
        "Mode: enum value AMAZON_COGNITO_USER_POOLS",
        "Mode: enum value OPENID_CONNECT",
        "ServiceName: enum value S3",
        "SignInMethod: enum value EMAIL",
        "SignInMethod: enum value EMAIL_AND_PHONE_NUMBER",
        "SignInMethod: enum value PHONE_NUMBER",
        "SignInMethod: enum value USERNAME",
        "UnAuthenticatedElement: enum value READ",
        "UnAuthenticatedElement: enum value CREATE_AND_UPDATE",
        "UnAuthenticatedElement: enum value DELETE"
      ],
      "All operations have clear summaries": [
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/details",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/auth/{backendEnvironmentName}",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels",
        "POST /backend/{appId}/config",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}",
        "POST /backend/{appId}/storage",
        "POST /backend",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import",
        "POST /backend/{appId}/challenge",
        "GET /backend/{appId}/challenge/{sessionId}",
        "POST /s3Buckets",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details",
        "POST /backend/{appId}/config/remove",
        "POST /backend/{appId}/storage/{backendEnvironmentName}",
        "POST /backend/{appId}/config/update",
        "POST /backend/{appId}/api/{backendEnvironmentName}",
        "POST /backend/{appId}/api",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/remove",
        "POST /backend/{appId}/challenge/{sessionId}/remove",
        "POST /backend/{appId}/job/{backendEnvironmentName}",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone",
        "POST /backend/{appId}/auth",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details"
      ],
      "All request/response bodies have examples": [
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: application/json request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 200 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 480 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 483 application/json response",
        "POST /backend/{appId}/details: request body",
        "POST /backend/{appId}/details: application/json request body",
        "POST /backend/{appId}/details: 481 application/json response",
        "POST /backend/{appId}/details: 482 application/json response",
        "POST /backend/{appId}/details: 483 application/json response",
        "POST /backend/{appId}/details: 200 application/json response",
        "POST /backend/{appId}/details: 480 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: application/json request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 200 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 480 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 481 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 482 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 483 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: application/json request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 482 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 483 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 200 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 480 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 481 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: application/json request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 483 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 200 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 480 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 481 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: application/json request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 480 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 483 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 200 application/json response",
        "POST /backend/{appId}/config: request body",
        "POST /backend/{appId}/config: application/json request body",
        "POST /backend/{appId}/config: 200 application/json response",
        "POST /backend/{appId}/config: 480 application/json response",
        "POST /backend/{appId}/config: 481 application/json response",
        "POST /backend/{appId}/config: 482 application/json response",
        "POST /backend/{appId}/config: 483 application/json response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 application/json response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 application/json response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 application/json response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 200 application/json response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: request body",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: application/json request body",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 200 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 application/json response",
        "POST /backend/{appId}/storage: request body",
        "POST /backend/{appId}/storage: application/json request body",
        "POST /backend/{appId}/storage: 483 application/json response",
        "POST /backend/{appId}/storage: 200 application/json response",
        "POST /backend/{appId}/storage: 480 application/json response",
        "POST /backend/{appId}/storage: 481 application/json response",
        "POST /backend/{appId}/storage: 482 application/json response",
        "POST /backend: request body",
        "POST /backend: application/json request body",
        "POST /backend: 480 application/json response",
        "POST /backend: 481 application/json response",
        "POST /backend: 482 application/json response",
        "POST /backend: 483 application/json response",
        "POST /backend: 200 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: application/json request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 482 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 483 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 200 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 480 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 481 application/json response",
        "POST /backend/{appId}/challenge: 480 application/json response",
        "POST /backend/{appId}/challenge: 481 application/json response",
        "POST /backend/{appId}/challenge: 482 application/json response",
        "POST /backend/{appId}/challenge: 483 application/json response",
        "POST /backend/{appId}/challenge: 200 application/json response",
        "GET /backend/{appId}/challenge/{sessionId}: 482 application/json response",
        "GET /backend/{appId}/challenge/{sessionId}: 483 application/json response",
        "GET /backend/{appId}/challenge/{sessionId}: 200 application/json response",
        "GET /backend/{appId}/challenge/{sessionId}: 480 application/json response",
        "GET /backend/{appId}/challenge/{sessionId}: 481 application/json response",
        "POST /s3Buckets: request body",
        "POST /s3Buckets: application/json request body",
        "POST /s3Buckets: 483 application/json response",
        "POST /s3Buckets: 200 application/json response",
        "POST /s3Buckets: 480 application/json response",
        "POST /s3Buckets: 481 application/json response",
        "POST /s3Buckets: 482 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: application/json request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 481 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 482 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 483 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 200 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 480 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: application/json request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 200 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 480 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 483 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: application/json request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 480 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 483 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 200 application/json response",
        "POST /backend/{appId}/config/remove: 483 application/json response",
        "POST /backend/{appId}/config/remove: 200 application/json response",
        "POST /backend/{appId}/config/remove: 480 application/json response",
        "POST /backend/{appId}/config/remove: 481 application/json response",
        "POST /backend/{appId}/config/remove: 482 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: application/json request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 483 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 200 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 480 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 481 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 482 application/json response",
        "POST /backend/{appId}/config/update: request body",
        "POST /backend/{appId}/config/update: application/json request body",
        "POST /backend/{appId}/config/update: 482 application/json response",
        "POST /backend/{appId}/config/update: 483 application/json response",
        "POST /backend/{appId}/config/update: 200 application/json response",
        "POST /backend/{appId}/config/update: 480 application/json response",
        "POST /backend/{appId}/config/update: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}: application/json request body",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 481 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 482 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 483 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 200 application/json response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 480 application/json response",
        "POST /backend/{appId}/api: request body",
        "POST /backend/{appId}/api: application/json request body",
        "POST /backend/{appId}/api: 200 application/json response",
        "POST /backend/{appId}/api: 480 application/json response",
        "POST /backend/{appId}/api: 481 application/json response",
        "POST /backend/{appId}/api: 482 application/json response",
        "POST /backend/{appId}/api: 483 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 200 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 480 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 481 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 482 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 483 application/json response",
        "POST /backend/{appId}/remove: request body",
        "POST /backend/{appId}/remove: application/json request body",
        "POST /backend/{appId}/remove: 483 application/json response",
        "POST /backend/{appId}/remove: 200 application/json response",
        "POST /backend/{appId}/remove: 480 application/json response",
        "POST /backend/{appId}/remove: 481 application/json response",
        "POST /backend/{appId}/remove: 482 application/json response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 200 application/json response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 480 application/json response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 481 application/json response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 482 application/json response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 483 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: request body",
        "POST /backend/{appId}/job/{backendEnvironmentName}: application/json request body",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 481 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 482 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 483 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 200 application/json response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 480 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: application/json request body",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 200 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 480 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 481 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 482 application/json response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 483 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: request body",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: application/json request body",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 483 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 200 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 480 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 481 application/json response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 482 application/json response",
        "POST /backend/{appId}/auth: request body",
        "POST /backend/{appId}/auth: application/json request body",
        "POST /backend/{appId}/auth: 482 application/json response",
        "POST /backend/{appId}/auth: 483 application/json response",
        "POST /backend/{appId}/auth: 200 application/json response",
        "POST /backend/{appId}/auth: 480 application/json response",
        "POST /backend/{appId}/auth: 481 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: application/json request body",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 481 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 482 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 483 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 200 application/json response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 480 application/json response"
      ],
      "All schemas have descriptions": [
        "GetTokenRequest",
        "ListOfAdditionalConstraintsElement",
        "RequiredSignUpAttributesElement",
        "UpdateBackendAPIResponse",
        "AuthenticatedElement",
        "OAuthScopesElement",
        "AuthResources",
        "DeleteBackendRequest",
        "ListOfUnAuthenticatedElement",
        "MfaTypesElement",
        "GetTokenResponse",
        "GetBackendAPIModelsResponse",
        "UpdateBackendJobResponse",
        "UpdateBackendStorageResponse",
        "GatewayTimeoutException",
        "ListOfAuthenticatedElement",
        "Service",
        "ListOfBackendJobRespObj",
        "__string",
        "DeleteBackendResponse",
        "__double",
        "DeleteBackendAuthResponse",
        "RemoveBackendConfigRequest",
        "OAuthGrantType",
        "GenerateBackendAPIModelsResponse",
        "RemoveAllBackendsResponse",
        "AdditionalConstraintsElement",
        "CloneBackendResponse",
        "ListBackendJobsResponse",
        "ImportBackendAuthResponse",
        "RemoveBackendConfigResponse",
        "TooManyRequestsException",
        "CreateBackendAPIResponse",
        "MFAMode",
        "BadRequestException",
        "CreateBackendConfigResponse",
        "__integerMin1Max25",
        "ListOfRequiredSignUpAttributesElement",
        "ResolutionStrategy",
        "CreateTokenRequest",
        "GetBackendJobRequest",
        "GetBackendJobResponse",
        "DeleteTokenResponse",
        "DeleteBackendStorageResponse",
        "ListS3BucketsResponse",
        "CreateBackendResponse",
        "NotFoundException",
        "__boolean",
        "CreateBackendStorageResponse",
        "Status",
        "ImportBackendStorageResponse",
        "DeleteBackendAPIResponse",
        "GetBackendResponse",
        "GetBackendStorageResponse",
        "GetBackendAuthResponse",
        "ListOfS3BucketInfo",
        "Mode",
        "UpdateBackendAuthResponse",
        "ListOfMfaTypesElement",
        "ServiceName",
        "SignInMethod",
        "CreateBackendAuthResponse",
        "ListOfOAuthScopesElement",
        "DeleteTokenRequest",
        "CreateTokenResponse",
        "ListOf__string",
        "UnAuthenticatedElement",
        "UpdateBackendConfigResponse",
        "GetBackendAPIResponse",
        "ListOfBackendAPIAuthType"
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
- **Message:** Request validation issues found: All string fields have length constraints: /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Content-Sha256, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Date, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Algorithm, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Credential, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Security-Token, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Signature, /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-SignedHeaders, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter appId, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter backendEnvironmentName, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter jobId, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter appId, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter backendEnvironmentName, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter jobId, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}.operation: application/json schema, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}.status: application/json schema, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Content-Sha256, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Date, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Algorithm, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Credential, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Security-Token, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Signature, /backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-SignedHeaders, POST /backend/{appId}/auth/{backendEnvironmentName}/import: parameter appId, POST /backend/{appId}/auth/{backendEnvironmentName}/import: parameter backendEnvironmentName, POST /backend/{appId}/auth/{backendEnvironmentName}/import.identityPoolId: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}/import.nativeClientId: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}/import.userPoolId: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}/import.webClientId: application/json schema, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Date, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Algorithm, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Credential, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Security-Token, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Signature, /backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api/{backendEnvironmentName}/details: parameter appId, POST /backend/{appId}/api/{backendEnvironmentName}/details: parameter backendEnvironmentName, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceName: application/json schema, /backend/{appId}/auth: parameter X-Amz-Content-Sha256, /backend/{appId}/auth: parameter X-Amz-Date, /backend/{appId}/auth: parameter X-Amz-Algorithm, /backend/{appId}/auth: parameter X-Amz-Credential, /backend/{appId}/auth: parameter X-Amz-Security-Token, /backend/{appId}/auth: parameter X-Amz-Signature, /backend/{appId}/auth: parameter X-Amz-SignedHeaders, POST /backend/{appId}/auth: parameter appId, POST /backend/{appId}/auth.resourceName: application/json schema, POST /backend/{appId}/auth.backendEnvironmentName: application/json schema, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Date, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Credential, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Signature, /backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api/{backendEnvironmentName}/remove: parameter appId, POST /backend/{appId}/api/{backendEnvironmentName}/remove: parameter backendEnvironmentName, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceName: application/json schema, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Content-Sha256, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Date, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Algorithm, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Credential, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Security-Token, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Signature, /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-SignedHeaders, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter appId, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter backendEnvironmentName, POST /backend/{appId}/environments/{backendEnvironmentName}/clone.targetEnvironmentName: application/json schema, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Content-Sha256, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Date, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Algorithm, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Credential, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Security-Token, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Signature, /backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-SignedHeaders, POST /backend/{appId}/job/{backendEnvironmentName}: parameter appId, POST /backend/{appId}/job/{backendEnvironmentName}: parameter backendEnvironmentName, POST /backend/{appId}/job/{backendEnvironmentName}: parameter MaxResults, POST /backend/{appId}/job/{backendEnvironmentName}: parameter NextToken, POST /backend/{appId}/job/{backendEnvironmentName}.nextToken: application/json schema, POST /backend/{appId}/job/{backendEnvironmentName}.operation: application/json schema, POST /backend/{appId}/job/{backendEnvironmentName}.status: application/json schema, POST /backend/{appId}/job/{backendEnvironmentName}.jobId: application/json schema, /backend: parameter X-Amz-Content-Sha256, /backend: parameter X-Amz-Date, /backend: parameter X-Amz-Algorithm, /backend: parameter X-Amz-Credential, /backend: parameter X-Amz-Security-Token, /backend: parameter X-Amz-Signature, /backend: parameter X-Amz-SignedHeaders, POST /backend.backendEnvironmentName: application/json schema, POST /backend.resourceName: application/json schema, POST /backend.appId: application/json schema, POST /backend.appName: application/json schema, /backend/{appId}/api: parameter X-Amz-Content-Sha256, /backend/{appId}/api: parameter X-Amz-Date, /backend/{appId}/api: parameter X-Amz-Algorithm, /backend/{appId}/api: parameter X-Amz-Credential, /backend/{appId}/api: parameter X-Amz-Security-Token, /backend/{appId}/api: parameter X-Amz-Signature, /backend/{appId}/api: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api: parameter appId, POST /backend/{appId}/api.resourceName: application/json schema, POST /backend/{appId}/api.backendEnvironmentName: application/json schema, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Date, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Credential, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Signature, /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter appId, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter backendEnvironmentName, POST /backend/{appId}/storage/{backendEnvironmentName}/remove.resourceName: application/json schema, POST /backend/{appId}/storage/{backendEnvironmentName}/remove.serviceName: application/json schema, /backend/{appId}/config/update: parameter X-Amz-Content-Sha256, /backend/{appId}/config/update: parameter X-Amz-Date, /backend/{appId}/config/update: parameter X-Amz-Algorithm, /backend/{appId}/config/update: parameter X-Amz-Credential, /backend/{appId}/config/update: parameter X-Amz-Security-Token, /backend/{appId}/config/update: parameter X-Amz-Signature, /backend/{appId}/config/update: parameter X-Amz-SignedHeaders, POST /backend/{appId}/config/update: parameter appId, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Content-Sha256, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Date, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Algorithm, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Credential, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Security-Token, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Signature, /backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-SignedHeaders, POST /backend/{appId}/auth/{backendEnvironmentName}: parameter appId, POST /backend/{appId}/auth/{backendEnvironmentName}: parameter backendEnvironmentName, POST /backend/{appId}/auth/{backendEnvironmentName}.resourceName: application/json schema, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Content-Sha256, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Date, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Algorithm, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Credential, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Security-Token, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Signature, /backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-SignedHeaders, POST /backend/{appId}/storage/{backendEnvironmentName}/import: parameter appId, POST /backend/{appId}/storage/{backendEnvironmentName}/import: parameter backendEnvironmentName, POST /backend/{appId}/storage/{backendEnvironmentName}/import.serviceName: application/json schema, POST /backend/{appId}/storage/{backendEnvironmentName}/import.bucketName: application/json schema, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Content-Sha256, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Date, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Algorithm, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Credential, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Security-Token, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Signature, /backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-SignedHeaders, POST /backend/{appId}/storage/{backendEnvironmentName}: parameter appId, POST /backend/{appId}/storage/{backendEnvironmentName}: parameter backendEnvironmentName, POST /backend/{appId}/storage/{backendEnvironmentName}.resourceName: application/json schema, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Date, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Algorithm, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Credential, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Security-Token, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Signature, /backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders, POST /backend/{appId}/storage/{backendEnvironmentName}/details: parameter appId, POST /backend/{appId}/storage/{backendEnvironmentName}/details: parameter backendEnvironmentName, POST /backend/{appId}/storage/{backendEnvironmentName}/details.resourceName: application/json schema, /backend/{appId}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/remove: parameter X-Amz-Date, /backend/{appId}/remove: parameter X-Amz-Algorithm, /backend/{appId}/remove: parameter X-Amz-Credential, /backend/{appId}/remove: parameter X-Amz-Security-Token, /backend/{appId}/remove: parameter X-Amz-Signature, /backend/{appId}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/remove: parameter appId, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Content-Sha256, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Date, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Algorithm, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Credential, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Security-Token, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Signature, /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter appId, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter backendEnvironmentName, POST /backend/{appId}/api/{backendEnvironmentName}/getModels.resourceName: application/json schema, /backend/{appId}/challenge: parameter X-Amz-Content-Sha256, /backend/{appId}/challenge: parameter X-Amz-Date, /backend/{appId}/challenge: parameter X-Amz-Algorithm, /backend/{appId}/challenge: parameter X-Amz-Credential, /backend/{appId}/challenge: parameter X-Amz-Security-Token, /backend/{appId}/challenge: parameter X-Amz-Signature, /backend/{appId}/challenge: parameter X-Amz-SignedHeaders, POST /backend/{appId}/challenge: parameter appId, /backend/{appId}/details: parameter X-Amz-Content-Sha256, /backend/{appId}/details: parameter X-Amz-Date, /backend/{appId}/details: parameter X-Amz-Algorithm, /backend/{appId}/details: parameter X-Amz-Credential, /backend/{appId}/details: parameter X-Amz-Security-Token, /backend/{appId}/details: parameter X-Amz-Signature, /backend/{appId}/details: parameter X-Amz-SignedHeaders, POST /backend/{appId}/details: parameter appId, POST /backend/{appId}/details.backendEnvironmentName: application/json schema, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Content-Sha256, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Date, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Algorithm, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Credential, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Security-Token, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Signature, /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter appId, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter backendEnvironmentName, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels.resourceName: application/json schema, /backend/{appId}/storage: parameter X-Amz-Content-Sha256, /backend/{appId}/storage: parameter X-Amz-Date, /backend/{appId}/storage: parameter X-Amz-Algorithm, /backend/{appId}/storage: parameter X-Amz-Credential, /backend/{appId}/storage: parameter X-Amz-Security-Token, /backend/{appId}/storage: parameter X-Amz-Signature, /backend/{appId}/storage: parameter X-Amz-SignedHeaders, POST /backend/{appId}/storage: parameter appId, POST /backend/{appId}/storage.backendEnvironmentName: application/json schema, POST /backend/{appId}/storage.resourceName: application/json schema, /s3Buckets: parameter X-Amz-Content-Sha256, /s3Buckets: parameter X-Amz-Date, /s3Buckets: parameter X-Amz-Algorithm, /s3Buckets: parameter X-Amz-Credential, /s3Buckets: parameter X-Amz-Security-Token, /s3Buckets: parameter X-Amz-Signature, /s3Buckets: parameter X-Amz-SignedHeaders, POST /s3Buckets.nextToken: application/json schema, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Date, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Algorithm, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Credential, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Security-Token, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Signature, /backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders, POST /backend/{appId}/auth/{backendEnvironmentName}/details: parameter appId, POST /backend/{appId}/auth/{backendEnvironmentName}/details: parameter backendEnvironmentName, POST /backend/{appId}/auth/{backendEnvironmentName}/details.resourceName: application/json schema, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Date, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Algorithm, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Credential, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Security-Token, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Signature, /backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/challenge/{sessionId}/remove: parameter appId, POST /backend/{appId}/challenge/{sessionId}/remove: parameter sessionId, /backend/{appId}/config: parameter X-Amz-Content-Sha256, /backend/{appId}/config: parameter X-Amz-Date, /backend/{appId}/config: parameter X-Amz-Algorithm, /backend/{appId}/config: parameter X-Amz-Credential, /backend/{appId}/config: parameter X-Amz-Security-Token, /backend/{appId}/config: parameter X-Amz-Signature, /backend/{appId}/config: parameter X-Amz-SignedHeaders, POST /backend/{appId}/config: parameter appId, POST /backend/{appId}/config.backendManagerAppId: application/json schema, /backend/{appId}/config/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/config/remove: parameter X-Amz-Date, /backend/{appId}/config/remove: parameter X-Amz-Algorithm, /backend/{appId}/config/remove: parameter X-Amz-Credential, /backend/{appId}/config/remove: parameter X-Amz-Security-Token, /backend/{appId}/config/remove: parameter X-Amz-Signature, /backend/{appId}/config/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/config/remove: parameter appId, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Content-Sha256, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Date, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Algorithm, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Credential, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Security-Token, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Signature, /backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-SignedHeaders, POST /backend/{appId}/api/{backendEnvironmentName}: parameter appId, POST /backend/{appId}/api/{backendEnvironmentName}: parameter backendEnvironmentName, POST /backend/{appId}/api/{backendEnvironmentName}.resourceName: application/json schema, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Date, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Credential, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Signature, /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter appId, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter backendEnvironmentName, POST /backend/{appId}/auth/{backendEnvironmentName}/remove.resourceName: application/json schema, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Content-Sha256, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Date, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Algorithm, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Credential, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Security-Token, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-Signature, /backend/{appId}/challenge/{sessionId}: parameter X-Amz-SignedHeaders, GET /backend/{appId}/challenge/{sessionId}: parameter appId, GET /backend/{appId}/challenge/{sessionId}: parameter sessionId, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Date, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Credential, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Signature, /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter appId, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter backendEnvironmentName; All schemas specify data types: POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.ConflictResolution: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.DefaultAuthType: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.Service: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.TransformSchema: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.AdditionalAuthTypes: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.ApiName: application/json schema, POST /backend/{appId}/auth.resourceConfig.AuthResources: application/json schema, POST /backend/{appId}/auth.resourceConfig.IdentityPoolConfigs: application/json schema, POST /backend/{appId}/auth.resourceConfig.Service: application/json schema, POST /backend/{appId}/auth.resourceConfig.UserPoolConfigs: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.Service: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.TransformSchema: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.AdditionalAuthTypes: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.ApiName: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.ConflictResolution: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.DefaultAuthType: application/json schema, POST /backend/{appId}/api.resourceConfig.AdditionalAuthTypes: application/json schema, POST /backend/{appId}/api.resourceConfig.ApiName: application/json schema, POST /backend/{appId}/api.resourceConfig.ConflictResolution: application/json schema, POST /backend/{appId}/api.resourceConfig.DefaultAuthType: application/json schema, POST /backend/{appId}/api.resourceConfig.Service: application/json schema, POST /backend/{appId}/api.resourceConfig.TransformSchema: application/json schema, POST /backend/{appId}/config/update.loginAuthConfig.AwsCognitoIdentityPoolId: application/json schema, POST /backend/{appId}/config/update.loginAuthConfig.AwsCognitoRegion: application/json schema, POST /backend/{appId}/config/update.loginAuthConfig.AwsUserPoolsId: application/json schema, POST /backend/{appId}/config/update.loginAuthConfig.AwsUserPoolsWebClientId: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.UserPoolConfigs: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.AuthResources: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.IdentityPoolConfigs: application/json schema, POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.Service: application/json schema, POST /backend/{appId}/storage/{backendEnvironmentName}.resourceConfig.Permissions: application/json schema, POST /backend/{appId}/storage/{backendEnvironmentName}.resourceConfig.ServiceName: application/json schema, POST /backend/{appId}/storage.resourceConfig.BucketName: application/json schema, POST /backend/{appId}/storage.resourceConfig.Permissions: application/json schema, POST /backend/{appId}/storage.resourceConfig.ServiceName: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.TransformSchema: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.AdditionalAuthTypes: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.ApiName: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.ConflictResolution: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.DefaultAuthType: application/json schema, POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.Service: application/json schema
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
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.ConflictResolution: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.DefaultAuthType: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.Service: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.TransformSchema: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.AdditionalAuthTypes: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceConfig.ApiName: application/json schema",
        "POST /backend/{appId}/auth.resourceConfig.AuthResources: application/json schema",
        "POST /backend/{appId}/auth.resourceConfig.IdentityPoolConfigs: application/json schema",
        "POST /backend/{appId}/auth.resourceConfig.Service: application/json schema",
        "POST /backend/{appId}/auth.resourceConfig.UserPoolConfigs: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.Service: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.TransformSchema: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.AdditionalAuthTypes: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.ApiName: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.ConflictResolution: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceConfig.DefaultAuthType: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.AdditionalAuthTypes: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.ApiName: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.ConflictResolution: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.DefaultAuthType: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.Service: application/json schema",
        "POST /backend/{appId}/api.resourceConfig.TransformSchema: application/json schema",
        "POST /backend/{appId}/config/update.loginAuthConfig.AwsCognitoIdentityPoolId: application/json schema",
        "POST /backend/{appId}/config/update.loginAuthConfig.AwsCognitoRegion: application/json schema",
        "POST /backend/{appId}/config/update.loginAuthConfig.AwsUserPoolsId: application/json schema",
        "POST /backend/{appId}/config/update.loginAuthConfig.AwsUserPoolsWebClientId: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.UserPoolConfigs: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.AuthResources: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.IdentityPoolConfigs: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}.resourceConfig.Service: application/json schema",
        "POST /backend/{appId}/storage/{backendEnvironmentName}.resourceConfig.Permissions: application/json schema",
        "POST /backend/{appId}/storage/{backendEnvironmentName}.resourceConfig.ServiceName: application/json schema",
        "POST /backend/{appId}/storage.resourceConfig.BucketName: application/json schema",
        "POST /backend/{appId}/storage.resourceConfig.Permissions: application/json schema",
        "POST /backend/{appId}/storage.resourceConfig.ServiceName: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.TransformSchema: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.AdditionalAuthTypes: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.ApiName: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.ConflictResolution: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.DefaultAuthType: application/json schema",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceConfig.Service: application/json schema"
      ],
      "All string fields have length constraints": [
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Date",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Algorithm",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Credential",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Security-Token",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-Signature",
        "/backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter X-Amz-SignedHeaders",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter appId",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter backendEnvironmentName",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter jobId",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter appId",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter backendEnvironmentName",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: parameter jobId",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}.operation: application/json schema",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}.status: application/json schema",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Date",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Algorithm",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Credential",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Security-Token",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-Signature",
        "/backend/{appId}/auth/{backendEnvironmentName}/import: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: parameter appId",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: parameter backendEnvironmentName",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import.identityPoolId: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import.nativeClientId: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import.userPoolId: application/json schema",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import.webClientId: application/json schema",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Date",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Algorithm",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Credential",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Security-Token",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-Signature",
        "/backend/{appId}/api/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: parameter appId",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: parameter backendEnvironmentName",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details.resourceName: application/json schema",
        "/backend/{appId}/auth: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/auth: parameter X-Amz-Date",
        "/backend/{appId}/auth: parameter X-Amz-Algorithm",
        "/backend/{appId}/auth: parameter X-Amz-Credential",
        "/backend/{appId}/auth: parameter X-Amz-Security-Token",
        "/backend/{appId}/auth: parameter X-Amz-Signature",
        "/backend/{appId}/auth: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/auth: parameter appId",
        "POST /backend/{appId}/auth.resourceName: application/json schema",
        "POST /backend/{appId}/auth.backendEnvironmentName: application/json schema",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Date",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/api/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: parameter appId",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: parameter backendEnvironmentName",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove.resourceName: application/json schema",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Date",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Algorithm",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Credential",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Security-Token",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-Signature",
        "/backend/{appId}/environments/{backendEnvironmentName}/clone: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter appId",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: parameter backendEnvironmentName",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone.targetEnvironmentName: application/json schema",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Date",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Algorithm",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Credential",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Security-Token",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-Signature",
        "/backend/{appId}/job/{backendEnvironmentName}: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/job/{backendEnvironmentName}: parameter appId",
        "POST /backend/{appId}/job/{backendEnvironmentName}: parameter backendEnvironmentName",
        "POST /backend/{appId}/job/{backendEnvironmentName}: parameter MaxResults",
        "POST /backend/{appId}/job/{backendEnvironmentName}: parameter NextToken",
        "POST /backend/{appId}/job/{backendEnvironmentName}.nextToken: application/json schema",
        "POST /backend/{appId}/job/{backendEnvironmentName}.operation: application/json schema",
        "POST /backend/{appId}/job/{backendEnvironmentName}.status: application/json schema",
        "POST /backend/{appId}/job/{backendEnvironmentName}.jobId: application/json schema",
        "/backend: parameter X-Amz-Content-Sha256",
        "/backend: parameter X-Amz-Date",
        "/backend: parameter X-Amz-Algorithm",
        "/backend: parameter X-Amz-Credential",
        "/backend: parameter X-Amz-Security-Token",
        "/backend: parameter X-Amz-Signature",
        "/backend: parameter X-Amz-SignedHeaders",
        "POST /backend.backendEnvironmentName: application/json schema",
        "POST /backend.resourceName: application/json schema",
        "POST /backend.appId: application/json schema",
        "POST /backend.appName: application/json schema",
        "/backend/{appId}/api: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api: parameter X-Amz-Date",
        "/backend/{appId}/api: parameter X-Amz-Algorithm",
        "/backend/{appId}/api: parameter X-Amz-Credential",
        "/backend/{appId}/api: parameter X-Amz-Security-Token",
        "/backend/{appId}/api: parameter X-Amz-Signature",
        "/backend/{appId}/api: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api: parameter appId",
        "POST /backend/{appId}/api.resourceName: application/json schema",
        "POST /backend/{appId}/api.backendEnvironmentName: application/json schema",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Date",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/storage/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter appId",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: parameter backendEnvironmentName",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove.resourceName: application/json schema",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove.serviceName: application/json schema",
        "/backend/{appId}/config/update: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/config/update: parameter X-Amz-Date",
        "/backend/{appId}/config/update: parameter X-Amz-Algorithm",
        "/backend/{appId}/config/update: parameter X-Amz-Credential",
        "/backend/{appId}/config/update: parameter X-Amz-Security-Token",
        "/backend/{appId}/config/update: parameter X-Amz-Signature",
        "/backend/{appId}/config/update: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/config/update: parameter appId",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Date",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Algorithm",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Credential",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Security-Token",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-Signature",
        "/backend/{appId}/auth/{backendEnvironmentName}: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: parameter appId",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: parameter backendEnvironmentName",
        "POST /backend/{appId}/auth/{backendEnvironmentName}.resourceName: application/json schema",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Date",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Algorithm",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Credential",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Security-Token",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-Signature",
        "/backend/{appId}/storage/{backendEnvironmentName}/import: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: parameter appId",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: parameter backendEnvironmentName",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import.serviceName: application/json schema",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import.bucketName: application/json schema",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Date",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Algorithm",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Credential",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Security-Token",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-Signature",
        "/backend/{appId}/storage/{backendEnvironmentName}: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: parameter appId",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: parameter backendEnvironmentName",
        "POST /backend/{appId}/storage/{backendEnvironmentName}.resourceName: application/json schema",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Date",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Algorithm",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Credential",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Security-Token",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-Signature",
        "/backend/{appId}/storage/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: parameter appId",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: parameter backendEnvironmentName",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details.resourceName: application/json schema",
        "/backend/{appId}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/remove: parameter X-Amz-Date",
        "/backend/{appId}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/remove: parameter appId",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Date",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Algorithm",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Credential",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Security-Token",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-Signature",
        "/backend/{appId}/api/{backendEnvironmentName}/getModels: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter appId",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: parameter backendEnvironmentName",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels.resourceName: application/json schema",
        "/backend/{appId}/challenge: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/challenge: parameter X-Amz-Date",
        "/backend/{appId}/challenge: parameter X-Amz-Algorithm",
        "/backend/{appId}/challenge: parameter X-Amz-Credential",
        "/backend/{appId}/challenge: parameter X-Amz-Security-Token",
        "/backend/{appId}/challenge: parameter X-Amz-Signature",
        "/backend/{appId}/challenge: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/challenge: parameter appId",
        "/backend/{appId}/details: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/details: parameter X-Amz-Date",
        "/backend/{appId}/details: parameter X-Amz-Algorithm",
        "/backend/{appId}/details: parameter X-Amz-Credential",
        "/backend/{appId}/details: parameter X-Amz-Security-Token",
        "/backend/{appId}/details: parameter X-Amz-Signature",
        "/backend/{appId}/details: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/details: parameter appId",
        "POST /backend/{appId}/details.backendEnvironmentName: application/json schema",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Date",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Algorithm",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Credential",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Security-Token",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-Signature",
        "/backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter appId",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: parameter backendEnvironmentName",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels.resourceName: application/json schema",
        "/backend/{appId}/storage: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/storage: parameter X-Amz-Date",
        "/backend/{appId}/storage: parameter X-Amz-Algorithm",
        "/backend/{appId}/storage: parameter X-Amz-Credential",
        "/backend/{appId}/storage: parameter X-Amz-Security-Token",
        "/backend/{appId}/storage: parameter X-Amz-Signature",
        "/backend/{appId}/storage: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/storage: parameter appId",
        "POST /backend/{appId}/storage.backendEnvironmentName: application/json schema",
        "POST /backend/{appId}/storage.resourceName: application/json schema",
        "/s3Buckets: parameter X-Amz-Content-Sha256",
        "/s3Buckets: parameter X-Amz-Date",
        "/s3Buckets: parameter X-Amz-Algorithm",
        "/s3Buckets: parameter X-Amz-Credential",
        "/s3Buckets: parameter X-Amz-Security-Token",
        "/s3Buckets: parameter X-Amz-Signature",
        "/s3Buckets: parameter X-Amz-SignedHeaders",
        "POST /s3Buckets.nextToken: application/json schema",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Date",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Algorithm",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Credential",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Security-Token",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-Signature",
        "/backend/{appId}/auth/{backendEnvironmentName}/details: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: parameter appId",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: parameter backendEnvironmentName",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details.resourceName: application/json schema",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Date",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/challenge/{sessionId}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/challenge/{sessionId}/remove: parameter appId",
        "POST /backend/{appId}/challenge/{sessionId}/remove: parameter sessionId",
        "/backend/{appId}/config: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/config: parameter X-Amz-Date",
        "/backend/{appId}/config: parameter X-Amz-Algorithm",
        "/backend/{appId}/config: parameter X-Amz-Credential",
        "/backend/{appId}/config: parameter X-Amz-Security-Token",
        "/backend/{appId}/config: parameter X-Amz-Signature",
        "/backend/{appId}/config: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/config: parameter appId",
        "POST /backend/{appId}/config.backendManagerAppId: application/json schema",
        "/backend/{appId}/config/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/config/remove: parameter X-Amz-Date",
        "/backend/{appId}/config/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/config/remove: parameter X-Amz-Credential",
        "/backend/{appId}/config/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/config/remove: parameter X-Amz-Signature",
        "/backend/{appId}/config/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/config/remove: parameter appId",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Date",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Algorithm",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Credential",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Security-Token",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-Signature",
        "/backend/{appId}/api/{backendEnvironmentName}: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/api/{backendEnvironmentName}: parameter appId",
        "POST /backend/{appId}/api/{backendEnvironmentName}: parameter backendEnvironmentName",
        "POST /backend/{appId}/api/{backendEnvironmentName}.resourceName: application/json schema",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Date",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/auth/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter appId",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: parameter backendEnvironmentName",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove.resourceName: application/json schema",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Date",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Algorithm",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Credential",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Security-Token",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-Signature",
        "/backend/{appId}/challenge/{sessionId}: parameter X-Amz-SignedHeaders",
        "GET /backend/{appId}/challenge/{sessionId}: parameter appId",
        "GET /backend/{appId}/challenge/{sessionId}: parameter sessionId",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Content-Sha256",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Date",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Algorithm",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Credential",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Security-Token",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-Signature",
        "/backend/{appId}/environments/{backendEnvironmentName}/remove: parameter X-Amz-SignedHeaders",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter appId",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: parameter backendEnvironmentName"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: GET /backend/{appId}/challenge/{sessionId}, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}, POST /backend/{appId}/config/update, POST /backend/{appId}/challenge, POST /backend/{appId}/storage/{backendEnvironmentName}/details, POST /backend/{appId}/environments/{backendEnvironmentName}/clone, POST /backend/{appId}/remove, POST /backend/{appId}/challenge/{sessionId}/remove, POST /backend/{appId}/auth, POST /backend/{appId}/auth/{backendEnvironmentName}, POST /backend/{appId}/storage/{backendEnvironmentName}/import, POST /backend/{appId}/job/{backendEnvironmentName}, POST /backend, POST /backend/{appId}/api/{backendEnvironmentName}/details, POST /backend/{appId}/auth/{backendEnvironmentName}/details, POST /backend/{appId}/storage/{backendEnvironmentName}/remove, POST /backend/{appId}/details, POST /backend/{appId}/api/{backendEnvironmentName}, POST /backend/{appId}/storage/{backendEnvironmentName}, POST /backend/{appId}/environments/{backendEnvironmentName}/remove, POST /backend/{appId}/api/{backendEnvironmentName}/getModels, POST /backend/{appId}/config, POST /backend/{appId}/storage, POST /backend/{appId}/config/remove, POST /backend/{appId}/auth/{backendEnvironmentName}/import, POST /backend/{appId}/api, POST /backend/{appId}/auth/{backendEnvironmentName}/remove, POST /s3Buckets, POST /backend/{appId}/api/{backendEnvironmentName}/remove, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels; Error responses include error details schema: GET /backend/{appId}/challenge/{sessionId}: 481 response, GET /backend/{appId}/challenge/{sessionId}: 482 response, GET /backend/{appId}/challenge/{sessionId}: 483 response, GET /backend/{appId}/challenge/{sessionId}: 480 response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 response, GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 response, POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 response, POST /backend/{appId}/config/update: 480 response, POST /backend/{appId}/config/update: 481 response, POST /backend/{appId}/config/update: 482 response, POST /backend/{appId}/config/update: 483 response, POST /backend/{appId}/challenge: 483 response, POST /backend/{appId}/challenge: 480 response, POST /backend/{appId}/challenge: 481 response, POST /backend/{appId}/challenge: 482 response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 482 response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 480 response, POST /backend/{appId}/storage/{backendEnvironmentName}/details: 481 response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 480 response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 481 response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 482 response, POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 483 response, POST /backend/{appId}/remove: 481 response, POST /backend/{appId}/remove: 482 response, POST /backend/{appId}/remove: 483 response, POST /backend/{appId}/remove: 480 response, POST /backend/{appId}/challenge/{sessionId}/remove: 480 response, POST /backend/{appId}/challenge/{sessionId}/remove: 481 response, POST /backend/{appId}/challenge/{sessionId}/remove: 482 response, POST /backend/{appId}/challenge/{sessionId}/remove: 483 response, POST /backend/{appId}/auth: 481 response, POST /backend/{appId}/auth: 482 response, POST /backend/{appId}/auth: 483 response, POST /backend/{appId}/auth: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}: 481 response, POST /backend/{appId}/auth/{backendEnvironmentName}: 482 response, POST /backend/{appId}/auth/{backendEnvironmentName}: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 480 response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 481 response, POST /backend/{appId}/storage/{backendEnvironmentName}/import: 482 response, POST /backend/{appId}/job/{backendEnvironmentName}: 481 response, POST /backend/{appId}/job/{backendEnvironmentName}: 482 response, POST /backend/{appId}/job/{backendEnvironmentName}: 483 response, POST /backend/{appId}/job/{backendEnvironmentName}: 480 response, POST /backend: 482 response, POST /backend: 483 response, POST /backend: 480 response, POST /backend: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 482 response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}/details: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 481 response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 482 response, POST /backend/{appId}/auth/{backendEnvironmentName}/details: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 481 response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 482 response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 480 response, POST /backend/{appId}/details: 480 response, POST /backend/{appId}/details: 481 response, POST /backend/{appId}/details: 482 response, POST /backend/{appId}/details: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}: 480 response, POST /backend/{appId}/api/{backendEnvironmentName}: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}: 482 response, POST /backend/{appId}/storage/{backendEnvironmentName}: 483 response, POST /backend/{appId}/storage/{backendEnvironmentName}: 480 response, POST /backend/{appId}/storage/{backendEnvironmentName}: 481 response, POST /backend/{appId}/storage/{backendEnvironmentName}: 482 response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 482 response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 483 response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 480 response, POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 482 response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 480 response, POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 481 response, POST /backend/{appId}/config: 480 response, POST /backend/{appId}/config: 481 response, POST /backend/{appId}/config: 482 response, POST /backend/{appId}/config: 483 response, POST /backend/{appId}/storage: 482 response, POST /backend/{appId}/storage: 483 response, POST /backend/{appId}/storage: 480 response, POST /backend/{appId}/storage: 481 response, POST /backend/{appId}/config/remove: 481 response, POST /backend/{appId}/config/remove: 482 response, POST /backend/{appId}/config/remove: 483 response, POST /backend/{appId}/config/remove: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 483 response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 481 response, POST /backend/{appId}/auth/{backendEnvironmentName}/import: 482 response, POST /backend/{appId}/api: 481 response, POST /backend/{appId}/api: 482 response, POST /backend/{appId}/api: 483 response, POST /backend/{appId}/api: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 483 response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 480 response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 481 response, POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 482 response, POST /s3Buckets: 480 response, POST /s3Buckets: 481 response, POST /s3Buckets: 482 response, POST /s3Buckets: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 483 response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 480 response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}/remove: 482 response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 480 response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 481 response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 482 response, POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 483 response
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
        "GET /backend/{appId}/challenge/{sessionId}",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}",
        "POST /backend/{appId}/config/update",
        "POST /backend/{appId}/challenge",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone",
        "POST /backend/{appId}/remove",
        "POST /backend/{appId}/challenge/{sessionId}/remove",
        "POST /backend/{appId}/auth",
        "POST /backend/{appId}/auth/{backendEnvironmentName}",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import",
        "POST /backend/{appId}/job/{backendEnvironmentName}",
        "POST /backend",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/details",
        "POST /backend/{appId}/api/{backendEnvironmentName}",
        "POST /backend/{appId}/storage/{backendEnvironmentName}",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels",
        "POST /backend/{appId}/config",
        "POST /backend/{appId}/storage",
        "POST /backend/{appId}/config/remove",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import",
        "POST /backend/{appId}/api",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove",
        "POST /s3Buckets",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels"
      ],
      "Error responses include error details schema": [
        "GET /backend/{appId}/challenge/{sessionId}: 481 response",
        "GET /backend/{appId}/challenge/{sessionId}: 482 response",
        "GET /backend/{appId}/challenge/{sessionId}: 483 response",
        "GET /backend/{appId}/challenge/{sessionId}: 480 response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 response",
        "GET /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 480 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 481 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 482 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}/{jobId}: 483 response",
        "POST /backend/{appId}/config/update: 480 response",
        "POST /backend/{appId}/config/update: 481 response",
        "POST /backend/{appId}/config/update: 482 response",
        "POST /backend/{appId}/config/update: 483 response",
        "POST /backend/{appId}/challenge: 483 response",
        "POST /backend/{appId}/challenge: 480 response",
        "POST /backend/{appId}/challenge: 481 response",
        "POST /backend/{appId}/challenge: 482 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 482 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 480 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/details: 481 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 480 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 481 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 482 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/clone: 483 response",
        "POST /backend/{appId}/remove: 481 response",
        "POST /backend/{appId}/remove: 482 response",
        "POST /backend/{appId}/remove: 483 response",
        "POST /backend/{appId}/remove: 480 response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 480 response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 481 response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 482 response",
        "POST /backend/{appId}/challenge/{sessionId}/remove: 483 response",
        "POST /backend/{appId}/auth: 481 response",
        "POST /backend/{appId}/auth: 482 response",
        "POST /backend/{appId}/auth: 483 response",
        "POST /backend/{appId}/auth: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 481 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 482 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 480 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 481 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/import: 482 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 481 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 482 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 483 response",
        "POST /backend/{appId}/job/{backendEnvironmentName}: 480 response",
        "POST /backend: 482 response",
        "POST /backend: 483 response",
        "POST /backend: 480 response",
        "POST /backend: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 482 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/details: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 481 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 482 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/details: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 481 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 482 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}/remove: 480 response",
        "POST /backend/{appId}/details: 480 response",
        "POST /backend/{appId}/details: 481 response",
        "POST /backend/{appId}/details: 482 response",
        "POST /backend/{appId}/details: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 480 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}: 482 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 483 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 480 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 481 response",
        "POST /backend/{appId}/storage/{backendEnvironmentName}: 482 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 482 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 483 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 480 response",
        "POST /backend/{appId}/environments/{backendEnvironmentName}/remove: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 482 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 480 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/getModels: 481 response",
        "POST /backend/{appId}/config: 480 response",
        "POST /backend/{appId}/config: 481 response",
        "POST /backend/{appId}/config: 482 response",
        "POST /backend/{appId}/config: 483 response",
        "POST /backend/{appId}/storage: 482 response",
        "POST /backend/{appId}/storage: 483 response",
        "POST /backend/{appId}/storage: 480 response",
        "POST /backend/{appId}/storage: 481 response",
        "POST /backend/{appId}/config/remove: 481 response",
        "POST /backend/{appId}/config/remove: 482 response",
        "POST /backend/{appId}/config/remove: 483 response",
        "POST /backend/{appId}/config/remove: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 483 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 481 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/import: 482 response",
        "POST /backend/{appId}/api: 481 response",
        "POST /backend/{appId}/api: 482 response",
        "POST /backend/{appId}/api: 483 response",
        "POST /backend/{appId}/api: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 483 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 480 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 481 response",
        "POST /backend/{appId}/auth/{backendEnvironmentName}/remove: 482 response",
        "POST /s3Buckets: 480 response",
        "POST /s3Buckets: 481 response",
        "POST /s3Buckets: 482 response",
        "POST /s3Buckets: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 483 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 480 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/remove: 482 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 480 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 481 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 482 response",
        "POST /backend/{appId}/api/{backendEnvironmentName}/generateModels: 483 response"
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
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "2020-08-11" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy
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
      "Version follows semantic versioning": "Version \"2020-08-11\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

