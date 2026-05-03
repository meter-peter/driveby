# API Validation Report

Generated: 2026-05-03T21:05:17+03:00
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
- Versioning
- Documentation
- Error Handling
- Schema
- Security


### Failed Tags
- errors
- validation
- request
- schema
- security
- versioning
- compatibility
- usability
- responses
- standards
- lifecycle
- documentation
- quality
- authentication
- authorization


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
- **Message:** Documentation quality issues found: All request/response bodies have examples: POST /getNotificationConfigurationList: request body, POST /getNotificationConfigurationList: 400 application/json response, POST /getNotificationConfigurationList: 401 application/json response, POST /getNotificationConfigurationList: 403 application/json response, POST /getNotificationConfigurationList: 422 application/json response, POST /getNotificationConfigurationList: 500 application/json response, POST /testNotificationConfiguration: request body, POST /testNotificationConfiguration: 422 application/json response, POST /testNotificationConfiguration: 500 application/json response, POST /testNotificationConfiguration: 400 application/json response, POST /testNotificationConfiguration: 401 application/json response, POST /testNotificationConfiguration: 403 application/json response, POST /updateNotificationConfiguration: request body, POST /updateNotificationConfiguration: 400 application/json response, POST /updateNotificationConfiguration: 401 application/json response, POST /updateNotificationConfiguration: 403 application/json response, POST /updateNotificationConfiguration: 422 application/json response, POST /updateNotificationConfiguration: 500 application/json response, POST /createNotificationConfiguration: request body, POST /createNotificationConfiguration: 401 application/json response, POST /createNotificationConfiguration: 403 application/json response, POST /createNotificationConfiguration: 422 application/json response, POST /createNotificationConfiguration: 500 application/json response, POST /createNotificationConfiguration: 400 application/json response, POST /deleteNotificationConfigurations: request body, POST /deleteNotificationConfigurations: 400 application/json response, POST /deleteNotificationConfigurations: 401 application/json response, POST /deleteNotificationConfigurations: 403 application/json response, POST /deleteNotificationConfigurations: 422 application/json response, POST /deleteNotificationConfigurations: 500 application/json response, POST /getNotificationConfiguration: request body, POST /getNotificationConfiguration: 401 application/json response, POST /getNotificationConfiguration: 403 application/json response, POST /getNotificationConfiguration: 422 application/json response, POST /getNotificationConfiguration: 500 application/json response, POST /getNotificationConfiguration: 400 application/json response; All schemas have descriptions: TestNotificationConfigurationRequest, TestNotificationConfigurationResponse, CreateNotificationConfigurationRequest, GetNotificationConfigurationListResponse, GenericResponse, NotificationConfigurationDetails, EmptyRequest, ExchangeMessage, DeleteNotificationConfigurationRequest, NotificationEventConfiguration, FieldType, UpdateNotificationConfigurationRequest, GetNotificationConfigurationResponse, ErrorFieldType, GetNotificationConfigurationRequest, ServiceError
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
        "POST /getNotificationConfigurationList: request body",
        "POST /getNotificationConfigurationList: 400 application/json response",
        "POST /getNotificationConfigurationList: 401 application/json response",
        "POST /getNotificationConfigurationList: 403 application/json response",
        "POST /getNotificationConfigurationList: 422 application/json response",
        "POST /getNotificationConfigurationList: 500 application/json response",
        "POST /testNotificationConfiguration: request body",
        "POST /testNotificationConfiguration: 422 application/json response",
        "POST /testNotificationConfiguration: 500 application/json response",
        "POST /testNotificationConfiguration: 400 application/json response",
        "POST /testNotificationConfiguration: 401 application/json response",
        "POST /testNotificationConfiguration: 403 application/json response",
        "POST /updateNotificationConfiguration: request body",
        "POST /updateNotificationConfiguration: 400 application/json response",
        "POST /updateNotificationConfiguration: 401 application/json response",
        "POST /updateNotificationConfiguration: 403 application/json response",
        "POST /updateNotificationConfiguration: 422 application/json response",
        "POST /updateNotificationConfiguration: 500 application/json response",
        "POST /createNotificationConfiguration: request body",
        "POST /createNotificationConfiguration: 401 application/json response",
        "POST /createNotificationConfiguration: 403 application/json response",
        "POST /createNotificationConfiguration: 422 application/json response",
        "POST /createNotificationConfiguration: 500 application/json response",
        "POST /createNotificationConfiguration: 400 application/json response",
        "POST /deleteNotificationConfigurations: request body",
        "POST /deleteNotificationConfigurations: 400 application/json response",
        "POST /deleteNotificationConfigurations: 401 application/json response",
        "POST /deleteNotificationConfigurations: 403 application/json response",
        "POST /deleteNotificationConfigurations: 422 application/json response",
        "POST /deleteNotificationConfigurations: 500 application/json response",
        "POST /getNotificationConfiguration: request body",
        "POST /getNotificationConfiguration: 401 application/json response",
        "POST /getNotificationConfiguration: 403 application/json response",
        "POST /getNotificationConfiguration: 422 application/json response",
        "POST /getNotificationConfiguration: 500 application/json response",
        "POST /getNotificationConfiguration: 400 application/json response"
      ],
      "All schemas have descriptions": [
        "TestNotificationConfigurationRequest",
        "TestNotificationConfigurationResponse",
        "CreateNotificationConfigurationRequest",
        "GetNotificationConfigurationListResponse",
        "GenericResponse",
        "NotificationConfigurationDetails",
        "EmptyRequest",
        "ExchangeMessage",
        "DeleteNotificationConfigurationRequest",
        "NotificationEventConfiguration",
        "FieldType",
        "UpdateNotificationConfigurationRequest",
        "GetNotificationConfigurationResponse",
        "ErrorFieldType",
        "GetNotificationConfigurationRequest",
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
- **Message:** Request validation issues found: All numeric fields have min/max values: POST /getNotificationConfiguration.notificationId: application/json schema, POST /testNotificationConfiguration.notificationId: application/json schema, POST /updateNotificationConfiguration.configurationDetails.notificationId: application/json schema, POST /updateNotificationConfiguration.configurationDetails.apiVersion: application/json schema, POST /createNotificationConfiguration.configurationDetails.apiVersion: application/json schema, POST /createNotificationConfiguration.configurationDetails.notificationId: application/json schema, POST /deleteNotificationConfigurations.notificationIds[]: application/json schema; All string fields have length constraints: POST /testNotificationConfiguration.eventTypes[]: application/json schema, POST /updateNotificationConfiguration.configurationDetails.description: application/json schema, POST /updateNotificationConfiguration.configurationDetails.notifyURL: application/json schema, POST /updateNotificationConfiguration.configurationDetails.eventConfigs[].eventType: application/json schema, POST /updateNotificationConfiguration.configurationDetails.eventConfigs[].includeMode: application/json schema, POST /updateNotificationConfiguration.configurationDetails.notifyPassword: application/json schema, POST /updateNotificationConfiguration.configurationDetails.sslProtocol: application/json schema, POST /updateNotificationConfiguration.configurationDetails.hmacSignatureKey: application/json schema, POST /updateNotificationConfiguration.configurationDetails.notifyUsername: application/json schema, POST /createNotificationConfiguration.configurationDetails.description: application/json schema, POST /createNotificationConfiguration.configurationDetails.eventConfigs[].includeMode: application/json schema, POST /createNotificationConfiguration.configurationDetails.eventConfigs[].eventType: application/json schema, POST /createNotificationConfiguration.configurationDetails.notifyPassword: application/json schema, POST /createNotificationConfiguration.configurationDetails.sslProtocol: application/json schema, POST /createNotificationConfiguration.configurationDetails.notifyUsername: application/json schema, POST /createNotificationConfiguration.configurationDetails.notifyURL: application/json schema, POST /createNotificationConfiguration.configurationDetails.hmacSignatureKey: application/json schema
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
        "POST /getNotificationConfiguration.notificationId: application/json schema",
        "POST /testNotificationConfiguration.notificationId: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.notificationId: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.apiVersion: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.apiVersion: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.notificationId: application/json schema",
        "POST /deleteNotificationConfigurations.notificationIds[]: application/json schema"
      ],
      "All string fields have length constraints": [
        "POST /testNotificationConfiguration.eventTypes[]: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.description: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.notifyURL: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.eventConfigs[].eventType: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.eventConfigs[].includeMode: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.notifyPassword: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.sslProtocol: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.hmacSignatureKey: application/json schema",
        "POST /updateNotificationConfiguration.configurationDetails.notifyUsername: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.description: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.eventConfigs[].includeMode: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.eventConfigs[].eventType: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.notifyPassword: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.sslProtocol: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.notifyUsername: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.notifyURL: application/json schema",
        "POST /createNotificationConfiguration.configurationDetails.hmacSignatureKey: application/json schema"
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

