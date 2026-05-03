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
- compatibility
- errors
- responses
- standards
- lifecycle
- documentation
- quality
- usability
- schema


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
- **Message:** Documentation quality issues found: All operations have clear summaries: POST /workspaces, GET /workspaces, DELETE /workspaces/{workspaceId}/alertmanager/definition, GET /workspaces/{workspaceId}/alertmanager/definition, POST /workspaces/{workspaceId}/alertmanager/definition, PUT /workspaces/{workspaceId}/alertmanager/definition, GET /workspaces/{workspaceId}/rulegroupsnamespaces, POST /workspaces/{workspaceId}/rulegroupsnamespaces, DELETE /tags/{resourceArn}#tagKeys, DELETE /workspaces/{workspaceId}, GET /workspaces/{workspaceId}, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, GET /tags/{resourceArn}, POST /tags/{resourceArn}, POST /workspaces/{workspaceId}/alias, POST /workspaces/{workspaceId}/logging, PUT /workspaces/{workspaceId}/logging, DELETE /workspaces/{workspaceId}/logging, GET /workspaces/{workspaceId}/logging; All request/response bodies have examples: POST /workspaces: request body, POST /workspaces: application/json request body, POST /workspaces: 481 application/json response, POST /workspaces: 482 application/json response, POST /workspaces: 483 application/json response, POST /workspaces: 484 application/json response, POST /workspaces: 485 application/json response, POST /workspaces: 202 application/json response, POST /workspaces: 480 application/json response, GET /workspaces: 480 application/json response, GET /workspaces: 481 application/json response, GET /workspaces: 482 application/json response, GET /workspaces: 483 application/json response, GET /workspaces: 200 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 200 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response, GET /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: request body, POST /workspaces/{workspaceId}/alertmanager/definition: application/json request body, POST /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 486 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 202 application/json response, POST /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: request body, PUT /workspaces/{workspaceId}/alertmanager/definition: application/json request body, PUT /workspaces/{workspaceId}/alertmanager/definition: 202 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response, PUT /workspaces/{workspaceId}/alertmanager/definition: 486 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 482 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 483 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 484 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 200 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 480 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 481 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: request body, POST /workspaces/{workspaceId}/rulegroupsnamespaces: application/json request body, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 486 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 202 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 480 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 481 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 482 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 483 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 484 application/json response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 485 application/json response, DELETE /tags/{resourceArn}#tagKeys: 200 application/json response, DELETE /tags/{resourceArn}#tagKeys: 480 application/json response, DELETE /tags/{resourceArn}#tagKeys: 481 application/json response, DELETE /tags/{resourceArn}#tagKeys: 482 application/json response, DELETE /tags/{resourceArn}#tagKeys: 483 application/json response, DELETE /tags/{resourceArn}#tagKeys: 484 application/json response, DELETE /workspaces/{workspaceId}: 484 application/json response, DELETE /workspaces/{workspaceId}: 485 application/json response, DELETE /workspaces/{workspaceId}: 480 application/json response, DELETE /workspaces/{workspaceId}: 481 application/json response, DELETE /workspaces/{workspaceId}: 482 application/json response, DELETE /workspaces/{workspaceId}: 483 application/json response, GET /workspaces/{workspaceId}: 484 application/json response, GET /workspaces/{workspaceId}: 200 application/json response, GET /workspaces/{workspaceId}: 480 application/json response, GET /workspaces/{workspaceId}: 481 application/json response, GET /workspaces/{workspaceId}: 482 application/json response, GET /workspaces/{workspaceId}: 483 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 200 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: request body, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: application/json request body, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 202 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 application/json response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 486 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response, GET /tags/{resourceArn}: 481 application/json response, GET /tags/{resourceArn}: 482 application/json response, GET /tags/{resourceArn}: 483 application/json response, GET /tags/{resourceArn}: 484 application/json response, GET /tags/{resourceArn}: 200 application/json response, GET /tags/{resourceArn}: 480 application/json response, POST /tags/{resourceArn}: request body, POST /tags/{resourceArn}: application/json request body, POST /tags/{resourceArn}: 200 application/json response, POST /tags/{resourceArn}: 480 application/json response, POST /tags/{resourceArn}: 481 application/json response, POST /tags/{resourceArn}: 482 application/json response, POST /tags/{resourceArn}: 483 application/json response, POST /tags/{resourceArn}: 484 application/json response, POST /workspaces/{workspaceId}/alias: request body, POST /workspaces/{workspaceId}/alias: application/json request body, POST /workspaces/{workspaceId}/alias: 484 application/json response, POST /workspaces/{workspaceId}/alias: 485 application/json response, POST /workspaces/{workspaceId}/alias: 486 application/json response, POST /workspaces/{workspaceId}/alias: 480 application/json response, POST /workspaces/{workspaceId}/alias: 481 application/json response, POST /workspaces/{workspaceId}/alias: 482 application/json response, POST /workspaces/{workspaceId}/alias: 483 application/json response, POST /workspaces/{workspaceId}/logging: request body, POST /workspaces/{workspaceId}/logging: application/json request body, POST /workspaces/{workspaceId}/logging: 483 application/json response, POST /workspaces/{workspaceId}/logging: 202 application/json response, POST /workspaces/{workspaceId}/logging: 480 application/json response, POST /workspaces/{workspaceId}/logging: 481 application/json response, POST /workspaces/{workspaceId}/logging: 482 application/json response, PUT /workspaces/{workspaceId}/logging: request body, PUT /workspaces/{workspaceId}/logging: application/json request body, PUT /workspaces/{workspaceId}/logging: 480 application/json response, PUT /workspaces/{workspaceId}/logging: 481 application/json response, PUT /workspaces/{workspaceId}/logging: 482 application/json response, PUT /workspaces/{workspaceId}/logging: 483 application/json response, PUT /workspaces/{workspaceId}/logging: 484 application/json response, PUT /workspaces/{workspaceId}/logging: 202 application/json response, DELETE /workspaces/{workspaceId}/logging: 483 application/json response, DELETE /workspaces/{workspaceId}/logging: 484 application/json response, DELETE /workspaces/{workspaceId}/logging: 480 application/json response, DELETE /workspaces/{workspaceId}/logging: 481 application/json response, DELETE /workspaces/{workspaceId}/logging: 482 application/json response, GET /workspaces/{workspaceId}/logging: 482 application/json response, GET /workspaces/{workspaceId}/logging: 483 application/json response, GET /workspaces/{workspaceId}/logging: 200 application/json response, GET /workspaces/{workspaceId}/logging: 480 application/json response, GET /workspaces/{workspaceId}/logging: 481 application/json response; All schemas have descriptions: ConflictException, ValidationException, ListWorkspacesRequestMaxResultsInteger, TagResourceRequest, TagResourceResponse, UntagResourceRequest, Timestamp, TagKeys, ListTagsForResourceResponse, ResourceNotFoundException, ServiceQuotaExceededException, TagValue, UntagResourceResponse, ThrottlingException, String, LogGroupArn, ListTagsForResourceRequest, ListRuleGroupsNamespacesRequestMaxResultsInteger, TagKey, Uri, InternalServerException, AccessDeniedException; All enums have descriptions: AlertManagerDefinitionStatusCode: enum value CREATING, AlertManagerDefinitionStatusCode: enum value ACTIVE, AlertManagerDefinitionStatusCode: enum value UPDATING, AlertManagerDefinitionStatusCode: enum value DELETING, AlertManagerDefinitionStatusCode: enum value CREATION_FAILED, AlertManagerDefinitionStatusCode: enum value UPDATE_FAILED, WorkspaceStatusCode: enum value CREATING, WorkspaceStatusCode: enum value ACTIVE, WorkspaceStatusCode: enum value UPDATING, WorkspaceStatusCode: enum value DELETING, WorkspaceStatusCode: enum value CREATION_FAILED, RuleGroupsNamespaceStatusCode: enum value CREATING, RuleGroupsNamespaceStatusCode: enum value ACTIVE, RuleGroupsNamespaceStatusCode: enum value UPDATING, RuleGroupsNamespaceStatusCode: enum value DELETING, RuleGroupsNamespaceStatusCode: enum value CREATION_FAILED, RuleGroupsNamespaceStatusCode: enum value UPDATE_FAILED, LoggingConfigurationStatusCode: enum value CREATING, LoggingConfigurationStatusCode: enum value ACTIVE, LoggingConfigurationStatusCode: enum value UPDATING, LoggingConfigurationStatusCode: enum value DELETING, LoggingConfigurationStatusCode: enum value CREATION_FAILED, LoggingConfigurationStatusCode: enum value UPDATE_FAILED
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
        "AlertManagerDefinitionStatusCode: enum value CREATING",
        "AlertManagerDefinitionStatusCode: enum value ACTIVE",
        "AlertManagerDefinitionStatusCode: enum value UPDATING",
        "AlertManagerDefinitionStatusCode: enum value DELETING",
        "AlertManagerDefinitionStatusCode: enum value CREATION_FAILED",
        "AlertManagerDefinitionStatusCode: enum value UPDATE_FAILED",
        "WorkspaceStatusCode: enum value CREATING",
        "WorkspaceStatusCode: enum value ACTIVE",
        "WorkspaceStatusCode: enum value UPDATING",
        "WorkspaceStatusCode: enum value DELETING",
        "WorkspaceStatusCode: enum value CREATION_FAILED",
        "RuleGroupsNamespaceStatusCode: enum value CREATING",
        "RuleGroupsNamespaceStatusCode: enum value ACTIVE",
        "RuleGroupsNamespaceStatusCode: enum value UPDATING",
        "RuleGroupsNamespaceStatusCode: enum value DELETING",
        "RuleGroupsNamespaceStatusCode: enum value CREATION_FAILED",
        "RuleGroupsNamespaceStatusCode: enum value UPDATE_FAILED",
        "LoggingConfigurationStatusCode: enum value CREATING",
        "LoggingConfigurationStatusCode: enum value ACTIVE",
        "LoggingConfigurationStatusCode: enum value UPDATING",
        "LoggingConfigurationStatusCode: enum value DELETING",
        "LoggingConfigurationStatusCode: enum value CREATION_FAILED",
        "LoggingConfigurationStatusCode: enum value UPDATE_FAILED"
      ],
      "All operations have clear summaries": [
        "POST /workspaces",
        "GET /workspaces",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition",
        "GET /workspaces/{workspaceId}/alertmanager/definition",
        "POST /workspaces/{workspaceId}/alertmanager/definition",
        "PUT /workspaces/{workspaceId}/alertmanager/definition",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces",
        "DELETE /tags/{resourceArn}#tagKeys",
        "DELETE /workspaces/{workspaceId}",
        "GET /workspaces/{workspaceId}",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "GET /tags/{resourceArn}",
        "POST /tags/{resourceArn}",
        "POST /workspaces/{workspaceId}/alias",
        "POST /workspaces/{workspaceId}/logging",
        "PUT /workspaces/{workspaceId}/logging",
        "DELETE /workspaces/{workspaceId}/logging",
        "GET /workspaces/{workspaceId}/logging"
      ],
      "All request/response bodies have examples": [
        "POST /workspaces: request body",
        "POST /workspaces: application/json request body",
        "POST /workspaces: 481 application/json response",
        "POST /workspaces: 482 application/json response",
        "POST /workspaces: 483 application/json response",
        "POST /workspaces: 484 application/json response",
        "POST /workspaces: 485 application/json response",
        "POST /workspaces: 202 application/json response",
        "POST /workspaces: 480 application/json response",
        "GET /workspaces: 480 application/json response",
        "GET /workspaces: 481 application/json response",
        "GET /workspaces: 482 application/json response",
        "GET /workspaces: 483 application/json response",
        "GET /workspaces: 200 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 200 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: request body",
        "POST /workspaces/{workspaceId}/alertmanager/definition: application/json request body",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 486 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 202 application/json response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: request body",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: application/json request body",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 202 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 480 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 481 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 482 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 483 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 484 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 485 application/json response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 486 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 482 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 483 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 484 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 200 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 480 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 481 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: request body",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: application/json request body",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 486 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 202 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 480 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 481 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 482 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 483 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 484 application/json response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 485 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 200 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 480 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 481 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 482 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 483 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 484 application/json response",
        "DELETE /workspaces/{workspaceId}: 484 application/json response",
        "DELETE /workspaces/{workspaceId}: 485 application/json response",
        "DELETE /workspaces/{workspaceId}: 480 application/json response",
        "DELETE /workspaces/{workspaceId}: 481 application/json response",
        "DELETE /workspaces/{workspaceId}: 482 application/json response",
        "DELETE /workspaces/{workspaceId}: 483 application/json response",
        "GET /workspaces/{workspaceId}: 484 application/json response",
        "GET /workspaces/{workspaceId}: 200 application/json response",
        "GET /workspaces/{workspaceId}: 480 application/json response",
        "GET /workspaces/{workspaceId}: 481 application/json response",
        "GET /workspaces/{workspaceId}: 482 application/json response",
        "GET /workspaces/{workspaceId}: 483 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 200 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: request body",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: application/json request body",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 202 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 application/json response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 486 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 application/json response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 application/json response",
        "GET /tags/{resourceArn}: 481 application/json response",
        "GET /tags/{resourceArn}: 482 application/json response",
        "GET /tags/{resourceArn}: 483 application/json response",
        "GET /tags/{resourceArn}: 484 application/json response",
        "GET /tags/{resourceArn}: 200 application/json response",
        "GET /tags/{resourceArn}: 480 application/json response",
        "POST /tags/{resourceArn}: request body",
        "POST /tags/{resourceArn}: application/json request body",
        "POST /tags/{resourceArn}: 200 application/json response",
        "POST /tags/{resourceArn}: 480 application/json response",
        "POST /tags/{resourceArn}: 481 application/json response",
        "POST /tags/{resourceArn}: 482 application/json response",
        "POST /tags/{resourceArn}: 483 application/json response",
        "POST /tags/{resourceArn}: 484 application/json response",
        "POST /workspaces/{workspaceId}/alias: request body",
        "POST /workspaces/{workspaceId}/alias: application/json request body",
        "POST /workspaces/{workspaceId}/alias: 484 application/json response",
        "POST /workspaces/{workspaceId}/alias: 485 application/json response",
        "POST /workspaces/{workspaceId}/alias: 486 application/json response",
        "POST /workspaces/{workspaceId}/alias: 480 application/json response",
        "POST /workspaces/{workspaceId}/alias: 481 application/json response",
        "POST /workspaces/{workspaceId}/alias: 482 application/json response",
        "POST /workspaces/{workspaceId}/alias: 483 application/json response",
        "POST /workspaces/{workspaceId}/logging: request body",
        "POST /workspaces/{workspaceId}/logging: application/json request body",
        "POST /workspaces/{workspaceId}/logging: 483 application/json response",
        "POST /workspaces/{workspaceId}/logging: 202 application/json response",
        "POST /workspaces/{workspaceId}/logging: 480 application/json response",
        "POST /workspaces/{workspaceId}/logging: 481 application/json response",
        "POST /workspaces/{workspaceId}/logging: 482 application/json response",
        "PUT /workspaces/{workspaceId}/logging: request body",
        "PUT /workspaces/{workspaceId}/logging: application/json request body",
        "PUT /workspaces/{workspaceId}/logging: 480 application/json response",
        "PUT /workspaces/{workspaceId}/logging: 481 application/json response",
        "PUT /workspaces/{workspaceId}/logging: 482 application/json response",
        "PUT /workspaces/{workspaceId}/logging: 483 application/json response",
        "PUT /workspaces/{workspaceId}/logging: 484 application/json response",
        "PUT /workspaces/{workspaceId}/logging: 202 application/json response",
        "DELETE /workspaces/{workspaceId}/logging: 483 application/json response",
        "DELETE /workspaces/{workspaceId}/logging: 484 application/json response",
        "DELETE /workspaces/{workspaceId}/logging: 480 application/json response",
        "DELETE /workspaces/{workspaceId}/logging: 481 application/json response",
        "DELETE /workspaces/{workspaceId}/logging: 482 application/json response",
        "GET /workspaces/{workspaceId}/logging: 482 application/json response",
        "GET /workspaces/{workspaceId}/logging: 483 application/json response",
        "GET /workspaces/{workspaceId}/logging: 200 application/json response",
        "GET /workspaces/{workspaceId}/logging: 480 application/json response",
        "GET /workspaces/{workspaceId}/logging: 481 application/json response"
      ],
      "All schemas have descriptions": [
        "ConflictException",
        "ValidationException",
        "ListWorkspacesRequestMaxResultsInteger",
        "TagResourceRequest",
        "TagResourceResponse",
        "UntagResourceRequest",
        "Timestamp",
        "TagKeys",
        "ListTagsForResourceResponse",
        "ResourceNotFoundException",
        "ServiceQuotaExceededException",
        "TagValue",
        "UntagResourceResponse",
        "ThrottlingException",
        "String",
        "LogGroupArn",
        "ListTagsForResourceRequest",
        "ListRuleGroupsNamespacesRequestMaxResultsInteger",
        "TagKey",
        "Uri",
        "InternalServerException",
        "AccessDeniedException"
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
- **Message:** Request validation issues found: All string fields have length constraints: /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Date, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Algorithm, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Credential, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Security-Token, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Signature, /workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-SignedHeaders, GET /workspaces/{workspaceId}/rulegroupsnamespaces: parameter nextToken, POST /workspaces/{workspaceId}/rulegroupsnamespaces.data: application/json schema, /tags/{resourceArn}#tagKeys: parameter X-Amz-Content-Sha256, /tags/{resourceArn}#tagKeys: parameter X-Amz-Date, /tags/{resourceArn}#tagKeys: parameter X-Amz-Algorithm, /tags/{resourceArn}#tagKeys: parameter X-Amz-Credential, /tags/{resourceArn}#tagKeys: parameter X-Amz-Security-Token, /tags/{resourceArn}#tagKeys: parameter X-Amz-Signature, /tags/{resourceArn}#tagKeys: parameter X-Amz-SignedHeaders, DELETE /tags/{resourceArn}#tagKeys: parameter resourceArn, /workspaces/{workspaceId}/alias: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}/alias: parameter X-Amz-Date, /workspaces/{workspaceId}/alias: parameter X-Amz-Algorithm, /workspaces/{workspaceId}/alias: parameter X-Amz-Credential, /workspaces/{workspaceId}/alias: parameter X-Amz-Security-Token, /workspaces/{workspaceId}/alias: parameter X-Amz-Signature, /workspaces/{workspaceId}/alias: parameter X-Amz-SignedHeaders, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Date, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Algorithm, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Credential, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Security-Token, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Signature, /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-SignedHeaders, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}.data: application/json schema, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Date, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Algorithm, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Credential, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Security-Token, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Signature, /workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-SignedHeaders, PUT /workspaces/{workspaceId}/alertmanager/definition.data: application/json schema, POST /workspaces/{workspaceId}/alertmanager/definition.data: application/json schema, /tags/{resourceArn}: parameter X-Amz-Content-Sha256, /tags/{resourceArn}: parameter X-Amz-Date, /tags/{resourceArn}: parameter X-Amz-Algorithm, /tags/{resourceArn}: parameter X-Amz-Credential, /tags/{resourceArn}: parameter X-Amz-Security-Token, /tags/{resourceArn}: parameter X-Amz-Signature, /tags/{resourceArn}: parameter X-Amz-SignedHeaders, POST /tags/{resourceArn}: parameter resourceArn, GET /tags/{resourceArn}: parameter resourceArn, /workspaces/{workspaceId}: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}: parameter X-Amz-Date, /workspaces/{workspaceId}: parameter X-Amz-Algorithm, /workspaces/{workspaceId}: parameter X-Amz-Credential, /workspaces/{workspaceId}: parameter X-Amz-Security-Token, /workspaces/{workspaceId}: parameter X-Amz-Signature, /workspaces/{workspaceId}: parameter X-Amz-SignedHeaders, /workspaces/{workspaceId}/logging: parameter X-Amz-Content-Sha256, /workspaces/{workspaceId}/logging: parameter X-Amz-Date, /workspaces/{workspaceId}/logging: parameter X-Amz-Algorithm, /workspaces/{workspaceId}/logging: parameter X-Amz-Credential, /workspaces/{workspaceId}/logging: parameter X-Amz-Security-Token, /workspaces/{workspaceId}/logging: parameter X-Amz-Signature, /workspaces/{workspaceId}/logging: parameter X-Amz-SignedHeaders, /workspaces: parameter X-Amz-Content-Sha256, /workspaces: parameter X-Amz-Date, /workspaces: parameter X-Amz-Algorithm, /workspaces: parameter X-Amz-Credential, /workspaces: parameter X-Amz-Security-Token, /workspaces: parameter X-Amz-Signature, /workspaces: parameter X-Amz-SignedHeaders, GET /workspaces: parameter nextToken
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
      "All schemas specify data types": true,
      "All string fields have length constraints": false
    },
    "messages": {},
    "missing_validation": {
      "All string fields have length constraints": [
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Date",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}/rulegroupsnamespaces: parameter X-Amz-SignedHeaders",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: parameter nextToken",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces.data: application/json schema",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Content-Sha256",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Date",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Algorithm",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Credential",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Security-Token",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Signature",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-SignedHeaders",
        "DELETE /tags/{resourceArn}#tagKeys: parameter resourceArn",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Date",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}/alias: parameter X-Amz-SignedHeaders",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Date",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}/rulegroupsnamespaces/{name}: parameter X-Amz-SignedHeaders",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}.data: application/json schema",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Date",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}/alertmanager/definition: parameter X-Amz-SignedHeaders",
        "PUT /workspaces/{workspaceId}/alertmanager/definition.data: application/json schema",
        "POST /workspaces/{workspaceId}/alertmanager/definition.data: application/json schema",
        "/tags/{resourceArn}: parameter X-Amz-Content-Sha256",
        "/tags/{resourceArn}: parameter X-Amz-Date",
        "/tags/{resourceArn}: parameter X-Amz-Algorithm",
        "/tags/{resourceArn}: parameter X-Amz-Credential",
        "/tags/{resourceArn}: parameter X-Amz-Security-Token",
        "/tags/{resourceArn}: parameter X-Amz-Signature",
        "/tags/{resourceArn}: parameter X-Amz-SignedHeaders",
        "POST /tags/{resourceArn}: parameter resourceArn",
        "GET /tags/{resourceArn}: parameter resourceArn",
        "/workspaces/{workspaceId}: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}: parameter X-Amz-Date",
        "/workspaces/{workspaceId}: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}: parameter X-Amz-SignedHeaders",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Content-Sha256",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Date",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Algorithm",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Credential",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Security-Token",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-Signature",
        "/workspaces/{workspaceId}/logging: parameter X-Amz-SignedHeaders",
        "/workspaces: parameter X-Amz-Content-Sha256",
        "/workspaces: parameter X-Amz-Date",
        "/workspaces: parameter X-Amz-Algorithm",
        "/workspaces: parameter X-Amz-Credential",
        "/workspaces: parameter X-Amz-Security-Token",
        "/workspaces: parameter X-Amz-Signature",
        "/workspaces: parameter X-Amz-SignedHeaders",
        "GET /workspaces: parameter nextToken"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: GET /workspaces/{workspaceId}/rulegroupsnamespaces, POST /workspaces/{workspaceId}/rulegroupsnamespaces, GET /tags/{resourceArn}, POST /tags/{resourceArn}, DELETE /tags/{resourceArn}#tagKeys, POST /workspaces/{workspaceId}/logging, PUT /workspaces/{workspaceId}/logging, DELETE /workspaces/{workspaceId}/logging, GET /workspaces/{workspaceId}/logging, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}, GET /workspaces, POST /workspaces, DELETE /workspaces/{workspaceId}, GET /workspaces/{workspaceId}, POST /workspaces/{workspaceId}/alias, DELETE /workspaces/{workspaceId}/alertmanager/definition, GET /workspaces/{workspaceId}/alertmanager/definition, POST /workspaces/{workspaceId}/alertmanager/definition, PUT /workspaces/{workspaceId}/alertmanager/definition; Error responses include error details schema: GET /workspaces/{workspaceId}/rulegroupsnamespaces: 481 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 482 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 483 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 484 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces: 480 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 481 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 482 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 483 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 484 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 485 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 486 response, POST /workspaces/{workspaceId}/rulegroupsnamespaces: 480 response, GET /tags/{resourceArn}: 484 response, GET /tags/{resourceArn}: 480 response, GET /tags/{resourceArn}: 481 response, GET /tags/{resourceArn}: 482 response, GET /tags/{resourceArn}: 483 response, POST /tags/{resourceArn}: 483 response, POST /tags/{resourceArn}: 484 response, POST /tags/{resourceArn}: 480 response, POST /tags/{resourceArn}: 481 response, POST /tags/{resourceArn}: 482 response, DELETE /tags/{resourceArn}#tagKeys: 481 response, DELETE /tags/{resourceArn}#tagKeys: 482 response, DELETE /tags/{resourceArn}#tagKeys: 483 response, DELETE /tags/{resourceArn}#tagKeys: 484 response, DELETE /tags/{resourceArn}#tagKeys: 480 response, POST /workspaces/{workspaceId}/logging: 480 response, POST /workspaces/{workspaceId}/logging: 481 response, POST /workspaces/{workspaceId}/logging: 482 response, POST /workspaces/{workspaceId}/logging: 483 response, PUT /workspaces/{workspaceId}/logging: 481 response, PUT /workspaces/{workspaceId}/logging: 482 response, PUT /workspaces/{workspaceId}/logging: 483 response, PUT /workspaces/{workspaceId}/logging: 484 response, PUT /workspaces/{workspaceId}/logging: 480 response, DELETE /workspaces/{workspaceId}/logging: 484 response, DELETE /workspaces/{workspaceId}/logging: 480 response, DELETE /workspaces/{workspaceId}/logging: 481 response, DELETE /workspaces/{workspaceId}/logging: 482 response, DELETE /workspaces/{workspaceId}/logging: 483 response, GET /workspaces/{workspaceId}/logging: 480 response, GET /workspaces/{workspaceId}/logging: 481 response, GET /workspaces/{workspaceId}/logging: 482 response, GET /workspaces/{workspaceId}/logging: 483 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 response, PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 486 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response, DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response, GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response, GET /workspaces: 480 response, GET /workspaces: 481 response, GET /workspaces: 482 response, GET /workspaces: 483 response, POST /workspaces: 480 response, POST /workspaces: 481 response, POST /workspaces: 482 response, POST /workspaces: 483 response, POST /workspaces: 484 response, POST /workspaces: 485 response, DELETE /workspaces/{workspaceId}: 480 response, DELETE /workspaces/{workspaceId}: 481 response, DELETE /workspaces/{workspaceId}: 482 response, DELETE /workspaces/{workspaceId}: 483 response, DELETE /workspaces/{workspaceId}: 484 response, DELETE /workspaces/{workspaceId}: 485 response, GET /workspaces/{workspaceId}: 480 response, GET /workspaces/{workspaceId}: 481 response, GET /workspaces/{workspaceId}: 482 response, GET /workspaces/{workspaceId}: 483 response, GET /workspaces/{workspaceId}: 484 response, POST /workspaces/{workspaceId}/alias: 482 response, POST /workspaces/{workspaceId}/alias: 483 response, POST /workspaces/{workspaceId}/alias: 484 response, POST /workspaces/{workspaceId}/alias: 485 response, POST /workspaces/{workspaceId}/alias: 486 response, POST /workspaces/{workspaceId}/alias: 480 response, POST /workspaces/{workspaceId}/alias: 481 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 483 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 484 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 485 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 480 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 481 response, DELETE /workspaces/{workspaceId}/alertmanager/definition: 482 response, GET /workspaces/{workspaceId}/alertmanager/definition: 480 response, GET /workspaces/{workspaceId}/alertmanager/definition: 481 response, GET /workspaces/{workspaceId}/alertmanager/definition: 482 response, GET /workspaces/{workspaceId}/alertmanager/definition: 483 response, GET /workspaces/{workspaceId}/alertmanager/definition: 484 response, POST /workspaces/{workspaceId}/alertmanager/definition: 482 response, POST /workspaces/{workspaceId}/alertmanager/definition: 483 response, POST /workspaces/{workspaceId}/alertmanager/definition: 484 response, POST /workspaces/{workspaceId}/alertmanager/definition: 485 response, POST /workspaces/{workspaceId}/alertmanager/definition: 486 response, POST /workspaces/{workspaceId}/alertmanager/definition: 480 response, POST /workspaces/{workspaceId}/alertmanager/definition: 481 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 480 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 481 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 482 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 483 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 484 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 485 response, PUT /workspaces/{workspaceId}/alertmanager/definition: 486 response
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
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces",
        "GET /tags/{resourceArn}",
        "POST /tags/{resourceArn}",
        "DELETE /tags/{resourceArn}#tagKeys",
        "POST /workspaces/{workspaceId}/logging",
        "PUT /workspaces/{workspaceId}/logging",
        "DELETE /workspaces/{workspaceId}/logging",
        "GET /workspaces/{workspaceId}/logging",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}",
        "GET /workspaces",
        "POST /workspaces",
        "DELETE /workspaces/{workspaceId}",
        "GET /workspaces/{workspaceId}",
        "POST /workspaces/{workspaceId}/alias",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition",
        "GET /workspaces/{workspaceId}/alertmanager/definition",
        "POST /workspaces/{workspaceId}/alertmanager/definition",
        "PUT /workspaces/{workspaceId}/alertmanager/definition"
      ],
      "Error responses include error details schema": [
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 481 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 482 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 483 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 484 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces: 480 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 481 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 482 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 483 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 484 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 485 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 486 response",
        "POST /workspaces/{workspaceId}/rulegroupsnamespaces: 480 response",
        "GET /tags/{resourceArn}: 484 response",
        "GET /tags/{resourceArn}: 480 response",
        "GET /tags/{resourceArn}: 481 response",
        "GET /tags/{resourceArn}: 482 response",
        "GET /tags/{resourceArn}: 483 response",
        "POST /tags/{resourceArn}: 483 response",
        "POST /tags/{resourceArn}: 484 response",
        "POST /tags/{resourceArn}: 480 response",
        "POST /tags/{resourceArn}: 481 response",
        "POST /tags/{resourceArn}: 482 response",
        "DELETE /tags/{resourceArn}#tagKeys: 481 response",
        "DELETE /tags/{resourceArn}#tagKeys: 482 response",
        "DELETE /tags/{resourceArn}#tagKeys: 483 response",
        "DELETE /tags/{resourceArn}#tagKeys: 484 response",
        "DELETE /tags/{resourceArn}#tagKeys: 480 response",
        "POST /workspaces/{workspaceId}/logging: 480 response",
        "POST /workspaces/{workspaceId}/logging: 481 response",
        "POST /workspaces/{workspaceId}/logging: 482 response",
        "POST /workspaces/{workspaceId}/logging: 483 response",
        "PUT /workspaces/{workspaceId}/logging: 481 response",
        "PUT /workspaces/{workspaceId}/logging: 482 response",
        "PUT /workspaces/{workspaceId}/logging: 483 response",
        "PUT /workspaces/{workspaceId}/logging: 484 response",
        "PUT /workspaces/{workspaceId}/logging: 480 response",
        "DELETE /workspaces/{workspaceId}/logging: 484 response",
        "DELETE /workspaces/{workspaceId}/logging: 480 response",
        "DELETE /workspaces/{workspaceId}/logging: 481 response",
        "DELETE /workspaces/{workspaceId}/logging: 482 response",
        "DELETE /workspaces/{workspaceId}/logging: 483 response",
        "GET /workspaces/{workspaceId}/logging: 480 response",
        "GET /workspaces/{workspaceId}/logging: 481 response",
        "GET /workspaces/{workspaceId}/logging: 482 response",
        "GET /workspaces/{workspaceId}/logging: 483 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 response",
        "PUT /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 486 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 485 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response",
        "DELETE /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 481 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 482 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 483 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 484 response",
        "GET /workspaces/{workspaceId}/rulegroupsnamespaces/{name}: 480 response",
        "GET /workspaces: 480 response",
        "GET /workspaces: 481 response",
        "GET /workspaces: 482 response",
        "GET /workspaces: 483 response",
        "POST /workspaces: 480 response",
        "POST /workspaces: 481 response",
        "POST /workspaces: 482 response",
        "POST /workspaces: 483 response",
        "POST /workspaces: 484 response",
        "POST /workspaces: 485 response",
        "DELETE /workspaces/{workspaceId}: 480 response",
        "DELETE /workspaces/{workspaceId}: 481 response",
        "DELETE /workspaces/{workspaceId}: 482 response",
        "DELETE /workspaces/{workspaceId}: 483 response",
        "DELETE /workspaces/{workspaceId}: 484 response",
        "DELETE /workspaces/{workspaceId}: 485 response",
        "GET /workspaces/{workspaceId}: 480 response",
        "GET /workspaces/{workspaceId}: 481 response",
        "GET /workspaces/{workspaceId}: 482 response",
        "GET /workspaces/{workspaceId}: 483 response",
        "GET /workspaces/{workspaceId}: 484 response",
        "POST /workspaces/{workspaceId}/alias: 482 response",
        "POST /workspaces/{workspaceId}/alias: 483 response",
        "POST /workspaces/{workspaceId}/alias: 484 response",
        "POST /workspaces/{workspaceId}/alias: 485 response",
        "POST /workspaces/{workspaceId}/alias: 486 response",
        "POST /workspaces/{workspaceId}/alias: 480 response",
        "POST /workspaces/{workspaceId}/alias: 481 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 483 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 484 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 485 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 480 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 481 response",
        "DELETE /workspaces/{workspaceId}/alertmanager/definition: 482 response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 480 response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 481 response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 482 response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 483 response",
        "GET /workspaces/{workspaceId}/alertmanager/definition: 484 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 482 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 483 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 484 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 485 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 486 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 480 response",
        "POST /workspaces/{workspaceId}/alertmanager/definition: 481 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 480 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 481 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 482 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 483 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 484 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 485 response",
        "PUT /workspaces/{workspaceId}/alertmanager/definition: 486 response"
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
- **Message:** Versioning validation failed: Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "2020-08-01" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility
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
      "Version follows semantic versioning": "Version \"2020-08-01\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

