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
- Specification
- Documentation
- Error Handling
- Schema
- Versioning


### Failed Tags
- compatibility
- openapi
- specification
- documentation
- schema
- validation
- lifecycle
- quality
- standards
- versioning
- request
- compliance
- errors
- responses
- usability


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "ApplicationId": error parsing regexp: invalid repeat count: `{1,1600}`
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
      "Specification structure is valid": "invalid components: schema \"ApplicationId\": error parsing regexp: invalid repeat count: `{1,1600}`"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All operations have clear summaries: POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState; All parameters have descriptions: POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Target, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Target; All request/response bodies have examples: POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: request body, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 489 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: request body, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: request body, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: request body, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: request body, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: request body, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: request body, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 489 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: request body, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: request body, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: request body, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: request body, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: request body, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: request body, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: request body, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: request body, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: request body, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 482 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: request body, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: application/json request body, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 483 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 484 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 487 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 481 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 485 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 486 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 488 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 200 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 480 application/json response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 482 application/json response; All schemas have descriptions: ProgressUpdateStream, MaxResultsCreatedArtifacts, NextUpdateSeconds, ImportMigrationTaskRequest, DescribeMigrationTaskRequest, DescribeMigrationTaskResult, DryRun, ApplicationStatus, UpdateDateTime, DisassociateDiscoveredResourceResult, ProgressPercent, ResourceAttributeList, InternalServerError, ApplicationId, NotifyApplicationStateRequest, CreatedArtifactName, ListCreatedArtifactsResult, ResourceAttributeType, ListProgressUpdateStreamsRequest, ResourceName, CreatedArtifactDescription, MigrationTaskName, ListDiscoveredResourcesResult, Token, ListMigrationTasksResult, NotifyApplicationStateResult, ListMigrationTasksRequest, NotifyMigrationTaskStateRequest, ApplicationIds, ConfigurationId, DeleteProgressUpdateStreamRequest, ServiceUnavailableException, AssociateDiscoveredResourceResult, DisassociateCreatedArtifactRequest, ListCreatedArtifactsRequest, PolicyErrorException, ProgressUpdateStreamSummaryList, UnauthorizedOperation, DryRunOperation, PutResourceAttributesResult, Status, ThrottlingException, DeleteProgressUpdateStreamResult, CreateProgressUpdateStreamResult, DisassociateDiscoveredResourceRequest, ListApplicationStatesResult, HomeRegionNotSetException, LatestResourceAttributeList, StatusDetail, CreateProgressUpdateStreamRequest, ListDiscoveredResourcesRequest, MigrationTaskSummaryList, ResourceNotFoundException, ResourceAttributeValue, ListProgressUpdateStreamsResult, ApplicationStateList, AccessDeniedException, DiscoveredResourceDescription, PutResourceAttributesRequest, AssociateDiscoveredResourceRequest, NotifyMigrationTaskStateResult, InvalidInputException, AssociateCreatedArtifactRequest, DescribeApplicationStateRequest, DisassociateCreatedArtifactResult, CreatedArtifactList, MaxResults, DiscoveredResourceList, ListApplicationStatesRequest, ImportMigrationTaskResult, DescribeApplicationStateResult, MaxResultsResources, AssociateCreatedArtifactResult; All enums have descriptions: ApplicationStatus: enum value NOT_STARTED, ApplicationStatus: enum value IN_PROGRESS, ApplicationStatus: enum value COMPLETED, ResourceAttributeType: enum value IPV4_ADDRESS, ResourceAttributeType: enum value IPV6_ADDRESS, ResourceAttributeType: enum value MAC_ADDRESS, ResourceAttributeType: enum value FQDN, ResourceAttributeType: enum value VM_MANAGER_ID, ResourceAttributeType: enum value VM_MANAGED_OBJECT_REFERENCE, ResourceAttributeType: enum value VM_NAME, ResourceAttributeType: enum value VM_PATH, ResourceAttributeType: enum value BIOS_ID, ResourceAttributeType: enum value MOTHERBOARD_SERIAL_NUMBER, Status: enum value NOT_STARTED, Status: enum value IN_PROGRESS, Status: enum value FAILED, Status: enum value COMPLETED
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
        "ApplicationStatus: enum value NOT_STARTED",
        "ApplicationStatus: enum value IN_PROGRESS",
        "ApplicationStatus: enum value COMPLETED",
        "ResourceAttributeType: enum value IPV4_ADDRESS",
        "ResourceAttributeType: enum value IPV6_ADDRESS",
        "ResourceAttributeType: enum value MAC_ADDRESS",
        "ResourceAttributeType: enum value FQDN",
        "ResourceAttributeType: enum value VM_MANAGER_ID",
        "ResourceAttributeType: enum value VM_MANAGED_OBJECT_REFERENCE",
        "ResourceAttributeType: enum value VM_NAME",
        "ResourceAttributeType: enum value VM_PATH",
        "ResourceAttributeType: enum value BIOS_ID",
        "ResourceAttributeType: enum value MOTHERBOARD_SERIAL_NUMBER",
        "Status: enum value NOT_STARTED",
        "Status: enum value IN_PROGRESS",
        "Status: enum value FAILED",
        "Status: enum value COMPLETED"
      ],
      "All operations have clear summaries": [
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState"
      ],
      "All parameters have descriptions": [
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Target",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Target"
      ],
      "All request/response bodies have examples": [
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 489 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 489 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 482 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: request body",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: application/json request body",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 483 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 484 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 487 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 481 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 485 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 486 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 488 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 200 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 480 application/json response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 482 application/json response"
      ],
      "All schemas have descriptions": [
        "ProgressUpdateStream",
        "MaxResultsCreatedArtifacts",
        "NextUpdateSeconds",
        "ImportMigrationTaskRequest",
        "DescribeMigrationTaskRequest",
        "DescribeMigrationTaskResult",
        "DryRun",
        "ApplicationStatus",
        "UpdateDateTime",
        "DisassociateDiscoveredResourceResult",
        "ProgressPercent",
        "ResourceAttributeList",
        "InternalServerError",
        "ApplicationId",
        "NotifyApplicationStateRequest",
        "CreatedArtifactName",
        "ListCreatedArtifactsResult",
        "ResourceAttributeType",
        "ListProgressUpdateStreamsRequest",
        "ResourceName",
        "CreatedArtifactDescription",
        "MigrationTaskName",
        "ListDiscoveredResourcesResult",
        "Token",
        "ListMigrationTasksResult",
        "NotifyApplicationStateResult",
        "ListMigrationTasksRequest",
        "NotifyMigrationTaskStateRequest",
        "ApplicationIds",
        "ConfigurationId",
        "DeleteProgressUpdateStreamRequest",
        "ServiceUnavailableException",
        "AssociateDiscoveredResourceResult",
        "DisassociateCreatedArtifactRequest",
        "ListCreatedArtifactsRequest",
        "PolicyErrorException",
        "ProgressUpdateStreamSummaryList",
        "UnauthorizedOperation",
        "DryRunOperation",
        "PutResourceAttributesResult",
        "Status",
        "ThrottlingException",
        "DeleteProgressUpdateStreamResult",
        "CreateProgressUpdateStreamResult",
        "DisassociateDiscoveredResourceRequest",
        "ListApplicationStatesResult",
        "HomeRegionNotSetException",
        "LatestResourceAttributeList",
        "StatusDetail",
        "CreateProgressUpdateStreamRequest",
        "ListDiscoveredResourcesRequest",
        "MigrationTaskSummaryList",
        "ResourceNotFoundException",
        "ResourceAttributeValue",
        "ListProgressUpdateStreamsResult",
        "ApplicationStateList",
        "AccessDeniedException",
        "DiscoveredResourceDescription",
        "PutResourceAttributesRequest",
        "AssociateDiscoveredResourceRequest",
        "NotifyMigrationTaskStateResult",
        "InvalidInputException",
        "AssociateCreatedArtifactRequest",
        "DescribeApplicationStateRequest",
        "DisassociateCreatedArtifactResult",
        "CreatedArtifactList",
        "MaxResults",
        "DiscoveredResourceList",
        "ListApplicationStatesRequest",
        "ImportMigrationTaskResult",
        "DescribeApplicationStateResult",
        "MaxResultsResources",
        "AssociateCreatedArtifactResult"
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
- **Message:** Request validation issues found: All string fields have length constraints: /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter MaxResults, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter NextToken, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter MaxResults, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter NextToken, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter MaxResults, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter NextToken, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter MaxResults, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter NextToken, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter MaxResults, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter NextToken, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Target, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Content-Sha256, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Date, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Algorithm, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Credential, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Security-Token, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Signature, /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-SignedHeaders, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Target; All schemas specify data types: POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.UpdateDateTime: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.ApplicationId: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.Status: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.ApplicationIds: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.MaxResults: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.NextToken: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.CreatedArtifact: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.CreatedArtifactName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream.ProgressUpdateStreamName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.MaxResults: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.NextToken: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.ResourceName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState.ApplicationId: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.DiscoveredResource: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream.ProgressUpdateStreamName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.ConfigurationId: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.NextToken: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.MaxResults: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams.MaxResults: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams.NextToken: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.NextToken: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.MaxResults: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.ResourceAttributeList: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.MigrationTaskName: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.NextUpdateSeconds: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.ProgressUpdateStream: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.Task: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.UpdateDateTime: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.DryRun: application/json schema, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.MigrationTaskName: application/json schema
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
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.UpdateDateTime: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.ApplicationId: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState.Status: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.ApplicationIds: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.MaxResults: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates.NextToken: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.CreatedArtifact: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.CreatedArtifactName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream.ProgressUpdateStreamName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.MaxResults: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.NextToken: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks.ResourceName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState.ApplicationId: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.DiscoveredResource: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream.ProgressUpdateStreamName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.ConfigurationId: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.NextToken: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources.MaxResults: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams.MaxResults: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams.NextToken: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.NextToken: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.MaxResults: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.ResourceAttributeList: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes.MigrationTaskName: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.NextUpdateSeconds: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.ProgressUpdateStream: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.Task: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.UpdateDateTime: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.DryRun: application/json schema",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState.MigrationTaskName: application/json schema"
      ],
      "All string fields have length constraints": [
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter MaxResults",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter NextToken",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter MaxResults",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter NextToken",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter MaxResults",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter NextToken",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter MaxResults",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter NextToken",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter MaxResults",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter NextToken",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: parameter X-Amz-Target",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Content-Sha256",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Date",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Algorithm",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Credential",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Security-Token",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Signature",
        "/#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-SignedHeaders",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: parameter X-Amz-Target"
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
- **Message:** Error handling issues found: All operations document 5xx error responses: POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams; Error responses include error details schema: POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 485 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 486 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 487 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 481 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 488 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 482 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 484 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 488 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 482 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 484 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 485 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 486 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 481 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 483 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 487 response, POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 482 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 485 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 487 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 481 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 484 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 486 response, POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 488 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 484 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 484 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 488 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 480 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 485 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 487 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 489 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 481 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 486 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 482 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 486 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 487 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 481 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 484 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 485 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 482 response, POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 488 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 485 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 486 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 481 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 482 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 484 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 487 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 484 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 488 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 486 response, POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 486 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 484 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 487 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 481 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 483 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 486 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 480 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 482 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 485 response, POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 484 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 486 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 484 response, POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 485 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 481 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 483 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 485 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 487 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 482 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 484 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 486 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 488 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 480 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 483 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 485 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 487 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 484 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 486 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 480 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 481 response, POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 484 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 486 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 487 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 483 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 484 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 485 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 486 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 488 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 482 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 480 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 481 response, POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 487 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 488 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 489 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 480 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 481 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 483 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 484 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 486 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 487 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 482 response, POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 485 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 480 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 481 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 482 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 483 response, POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 484 response
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
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams"
      ],
      "Error responses include error details schema": [
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateDiscoveredResource: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.PutResourceAttributes: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DisassociateCreatedArtifact: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListApplicationStates: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 489 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateDiscoveredResource: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DeleteProgressUpdateStream: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeMigrationTask: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ImportMigrationTask: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListCreatedArtifacts: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.DescribeApplicationState: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListDiscoveredResources: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyMigrationTaskState: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.CreateProgressUpdateStream: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListMigrationTasks: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.AssociateCreatedArtifact: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 488 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 489 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 484 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 486 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 487 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.NotifyApplicationState: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 485 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 480 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 481 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 482 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 483 response",
        "POST /#X-Amz-Target=AWSMigrationHub.ListProgressUpdateStreams: 484 response"
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
- **Message:** Versioning validation failed: Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility; Version follows semantic versioning: Version "2017-05-31" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy
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
      "Migration guides are referenced": true,
      "Version compatibility is specified": false,
      "Version follows semantic versioning": false,
      "Versioning strategy is documented": false
    },
    "messages": {
      "Breaking changes are documented": "Info description does not reference breaking changes or a changelog",
      "Version compatibility is specified": "Info description does not mention version compatibility",
      "Version follows semantic versioning": "Version \"2017-05-31\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

