# API Validation Report

Generated: 2026-05-03T21:05:25+03:00
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
- responses
- schema
- validation
- openapi
- compliance
- documentation
- request
- compatibility
- standards
- versioning
- specification
- quality
- usability
- lifecycle
- errors


## Principle Results

### Specification

#### P001: OpenAPI Specification Compliance (Failed) [critical]

Validates that the API specification follows OpenAPI 3.0/3.1 standards and best practices

- **Status:** Failed
- **Message:** OpenAPI spec validation failed: Specification structure is valid: invalid components: schema "CreateDomainAssociationRequest": error parsing regexp: invalid or unsupported Perl syntax: `(?!`
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
      "Specification structure is valid": "invalid components: schema \"CreateDomainAssociationRequest\": error parsing regexp: invalid or unsupported Perl syntax: `(?!`"
    }
  }
```

---

### Documentation

#### P002: API Documentation Quality (Failed) [critical]

Ensures comprehensive and high-quality API documentation including descriptions, examples, and usage guidelines

- **Status:** Failed
- **Message:** Documentation quality issues found: All operations have clear summaries: GET /apps/{appId}/domains/{domainName}, POST /apps/{appId}/domains/{domainName}, DELETE /apps/{appId}/domains/{domainName}, GET /artifacts/{artifactId}, GET /apps/{appId}/backendenvironments, POST /apps/{appId}/backendenvironments, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop, GET /tags/{resourceArn}, POST /tags/{resourceArn}, GET /apps, POST /apps, POST /apps/{appId}/accesslogs, POST /apps/{appId}/branches/{branchName}/deployments, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts, DELETE /webhooks/{webhookId}, GET /webhooks/{webhookId}, POST /webhooks/{webhookId}, DELETE /apps/{appId}, GET /apps/{appId}, POST /apps/{appId}, DELETE /tags/{resourceArn}#tagKeys, POST /apps/{appId}/domains, GET /apps/{appId}/domains, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}, GET /apps/{appId}/webhooks, POST /apps/{appId}/webhooks, DELETE /apps/{appId}/branches/{branchName}, GET /apps/{appId}/branches/{branchName}, POST /apps/{appId}/branches/{branchName}, GET /apps/{appId}/branches/{branchName}/jobs, POST /apps/{appId}/branches/{branchName}/jobs, DELETE /apps/{appId}/backendenvironments/{environmentName}, GET /apps/{appId}/backendenvironments/{environmentName}, POST /apps/{appId}/branches, GET /apps/{appId}/branches, POST /apps/{appId}/branches/{branchName}/deployments/start; All request/response bodies have examples: GET /apps/{appId}/domains/{domainName}: 200 application/json response, GET /apps/{appId}/domains/{domainName}: 480 application/json response, GET /apps/{appId}/domains/{domainName}: 481 application/json response, GET /apps/{appId}/domains/{domainName}: 482 application/json response, GET /apps/{appId}/domains/{domainName}: 483 application/json response, POST /apps/{appId}/domains/{domainName}: request body, POST /apps/{appId}/domains/{domainName}: application/json request body, POST /apps/{appId}/domains/{domainName}: 482 application/json response, POST /apps/{appId}/domains/{domainName}: 483 application/json response, POST /apps/{appId}/domains/{domainName}: 484 application/json response, POST /apps/{appId}/domains/{domainName}: 200 application/json response, POST /apps/{appId}/domains/{domainName}: 480 application/json response, POST /apps/{appId}/domains/{domainName}: 481 application/json response, DELETE /apps/{appId}/domains/{domainName}: 480 application/json response, DELETE /apps/{appId}/domains/{domainName}: 481 application/json response, DELETE /apps/{appId}/domains/{domainName}: 482 application/json response, DELETE /apps/{appId}/domains/{domainName}: 483 application/json response, DELETE /apps/{appId}/domains/{domainName}: 484 application/json response, DELETE /apps/{appId}/domains/{domainName}: 200 application/json response, GET /artifacts/{artifactId}: 200 application/json response, GET /artifacts/{artifactId}: 480 application/json response, GET /artifacts/{artifactId}: 481 application/json response, GET /artifacts/{artifactId}: 482 application/json response, GET /artifacts/{artifactId}: 483 application/json response, GET /artifacts/{artifactId}: 484 application/json response, GET /apps/{appId}/backendenvironments: 481 application/json response, GET /apps/{appId}/backendenvironments: 482 application/json response, GET /apps/{appId}/backendenvironments: 200 application/json response, GET /apps/{appId}/backendenvironments: 480 application/json response, POST /apps/{appId}/backendenvironments: request body, POST /apps/{appId}/backendenvironments: application/json request body, POST /apps/{appId}/backendenvironments: 481 application/json response, POST /apps/{appId}/backendenvironments: 482 application/json response, POST /apps/{appId}/backendenvironments: 483 application/json response, POST /apps/{appId}/backendenvironments: 484 application/json response, POST /apps/{appId}/backendenvironments: 200 application/json response, POST /apps/{appId}/backendenvironments: 480 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 200 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 480 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 481 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 482 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 483 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 484 application/json response, GET /tags/{resourceArn}: 200 application/json response, GET /tags/{resourceArn}: 480 application/json response, GET /tags/{resourceArn}: 481 application/json response, GET /tags/{resourceArn}: 482 application/json response, POST /tags/{resourceArn}: request body, POST /tags/{resourceArn}: application/json request body, POST /tags/{resourceArn}: 480 application/json response, POST /tags/{resourceArn}: 481 application/json response, POST /tags/{resourceArn}: 482 application/json response, POST /tags/{resourceArn}: 200 application/json response, GET /apps: 481 application/json response, GET /apps: 482 application/json response, GET /apps: 200 application/json response, GET /apps: 480 application/json response, POST /apps: request body, POST /apps: application/json request body, POST /apps: 483 application/json response, POST /apps: 484 application/json response, POST /apps: 200 application/json response, POST /apps: 480 application/json response, POST /apps: 481 application/json response, POST /apps: 482 application/json response, POST /apps/{appId}/accesslogs: request body, POST /apps/{appId}/accesslogs: application/json request body, POST /apps/{appId}/accesslogs: 482 application/json response, POST /apps/{appId}/accesslogs: 483 application/json response, POST /apps/{appId}/accesslogs: 200 application/json response, POST /apps/{appId}/accesslogs: 480 application/json response, POST /apps/{appId}/accesslogs: 481 application/json response, POST /apps/{appId}/branches/{branchName}/deployments: request body, POST /apps/{appId}/branches/{branchName}/deployments: application/json request body, POST /apps/{appId}/branches/{branchName}/deployments: 482 application/json response, POST /apps/{appId}/branches/{branchName}/deployments: 483 application/json response, POST /apps/{appId}/branches/{branchName}/deployments: 200 application/json response, POST /apps/{appId}/branches/{branchName}/deployments: 480 application/json response, POST /apps/{appId}/branches/{branchName}/deployments: 481 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 481 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 482 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 483 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 200 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 480 application/json response, DELETE /webhooks/{webhookId}: 480 application/json response, DELETE /webhooks/{webhookId}: 481 application/json response, DELETE /webhooks/{webhookId}: 482 application/json response, DELETE /webhooks/{webhookId}: 483 application/json response, DELETE /webhooks/{webhookId}: 484 application/json response, DELETE /webhooks/{webhookId}: 200 application/json response, GET /webhooks/{webhookId}: 481 application/json response, GET /webhooks/{webhookId}: 482 application/json response, GET /webhooks/{webhookId}: 483 application/json response, GET /webhooks/{webhookId}: 484 application/json response, GET /webhooks/{webhookId}: 200 application/json response, GET /webhooks/{webhookId}: 480 application/json response, POST /webhooks/{webhookId}: request body, POST /webhooks/{webhookId}: application/json request body, POST /webhooks/{webhookId}: 481 application/json response, POST /webhooks/{webhookId}: 482 application/json response, POST /webhooks/{webhookId}: 483 application/json response, POST /webhooks/{webhookId}: 484 application/json response, POST /webhooks/{webhookId}: 200 application/json response, POST /webhooks/{webhookId}: 480 application/json response, DELETE /apps/{appId}: 481 application/json response, DELETE /apps/{appId}: 482 application/json response, DELETE /apps/{appId}: 483 application/json response, DELETE /apps/{appId}: 484 application/json response, DELETE /apps/{appId}: 200 application/json response, DELETE /apps/{appId}: 480 application/json response, GET /apps/{appId}: 481 application/json response, GET /apps/{appId}: 482 application/json response, GET /apps/{appId}: 483 application/json response, GET /apps/{appId}: 200 application/json response, GET /apps/{appId}: 480 application/json response, POST /apps/{appId}: request body, POST /apps/{appId}: application/json request body, POST /apps/{appId}: 481 application/json response, POST /apps/{appId}: 482 application/json response, POST /apps/{appId}: 483 application/json response, POST /apps/{appId}: 200 application/json response, POST /apps/{appId}: 480 application/json response, DELETE /tags/{resourceArn}#tagKeys: 200 application/json response, DELETE /tags/{resourceArn}#tagKeys: 480 application/json response, DELETE /tags/{resourceArn}#tagKeys: 481 application/json response, DELETE /tags/{resourceArn}#tagKeys: 482 application/json response, POST /apps/{appId}/domains: request body, POST /apps/{appId}/domains: application/json request body, POST /apps/{appId}/domains: 485 application/json response, POST /apps/{appId}/domains: 200 application/json response, POST /apps/{appId}/domains: 480 application/json response, POST /apps/{appId}/domains: 481 application/json response, POST /apps/{appId}/domains: 482 application/json response, POST /apps/{appId}/domains: 483 application/json response, POST /apps/{appId}/domains: 484 application/json response, GET /apps/{appId}/domains: 200 application/json response, GET /apps/{appId}/domains: 480 application/json response, GET /apps/{appId}/domains: 481 application/json response, GET /apps/{appId}/domains: 482 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 application/json response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 200 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 200 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 application/json response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 application/json response, GET /apps/{appId}/webhooks: 482 application/json response, GET /apps/{appId}/webhooks: 483 application/json response, GET /apps/{appId}/webhooks: 200 application/json response, GET /apps/{appId}/webhooks: 480 application/json response, GET /apps/{appId}/webhooks: 481 application/json response, POST /apps/{appId}/webhooks: request body, POST /apps/{appId}/webhooks: application/json request body, POST /apps/{appId}/webhooks: 485 application/json response, POST /apps/{appId}/webhooks: 200 application/json response, POST /apps/{appId}/webhooks: 480 application/json response, POST /apps/{appId}/webhooks: 481 application/json response, POST /apps/{appId}/webhooks: 482 application/json response, POST /apps/{appId}/webhooks: 483 application/json response, POST /apps/{appId}/webhooks: 484 application/json response, DELETE /apps/{appId}/branches/{branchName}: 480 application/json response, DELETE /apps/{appId}/branches/{branchName}: 481 application/json response, DELETE /apps/{appId}/branches/{branchName}: 482 application/json response, DELETE /apps/{appId}/branches/{branchName}: 483 application/json response, DELETE /apps/{appId}/branches/{branchName}: 484 application/json response, DELETE /apps/{appId}/branches/{branchName}: 200 application/json response, GET /apps/{appId}/branches/{branchName}: 200 application/json response, GET /apps/{appId}/branches/{branchName}: 480 application/json response, GET /apps/{appId}/branches/{branchName}: 481 application/json response, GET /apps/{appId}/branches/{branchName}: 482 application/json response, GET /apps/{appId}/branches/{branchName}: 483 application/json response, POST /apps/{appId}/branches/{branchName}: request body, POST /apps/{appId}/branches/{branchName}: application/json request body, POST /apps/{appId}/branches/{branchName}: 481 application/json response, POST /apps/{appId}/branches/{branchName}: 482 application/json response, POST /apps/{appId}/branches/{branchName}: 483 application/json response, POST /apps/{appId}/branches/{branchName}: 484 application/json response, POST /apps/{appId}/branches/{branchName}: 200 application/json response, POST /apps/{appId}/branches/{branchName}: 480 application/json response, GET /apps/{appId}/branches/{branchName}/jobs: 480 application/json response, GET /apps/{appId}/branches/{branchName}/jobs: 481 application/json response, GET /apps/{appId}/branches/{branchName}/jobs: 482 application/json response, GET /apps/{appId}/branches/{branchName}/jobs: 483 application/json response, GET /apps/{appId}/branches/{branchName}/jobs: 200 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: request body, POST /apps/{appId}/branches/{branchName}/jobs: application/json request body, POST /apps/{appId}/branches/{branchName}/jobs: 481 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: 482 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: 483 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: 484 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: 200 application/json response, POST /apps/{appId}/branches/{branchName}/jobs: 480 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 484 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 200 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 480 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 481 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 482 application/json response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 483 application/json response, GET /apps/{appId}/backendenvironments/{environmentName}: 480 application/json response, GET /apps/{appId}/backendenvironments/{environmentName}: 481 application/json response, GET /apps/{appId}/backendenvironments/{environmentName}: 482 application/json response, GET /apps/{appId}/backendenvironments/{environmentName}: 483 application/json response, GET /apps/{appId}/backendenvironments/{environmentName}: 200 application/json response, POST /apps/{appId}/branches: request body, POST /apps/{appId}/branches: application/json request body, POST /apps/{appId}/branches: 483 application/json response, POST /apps/{appId}/branches: 484 application/json response, POST /apps/{appId}/branches: 485 application/json response, POST /apps/{appId}/branches: 200 application/json response, POST /apps/{appId}/branches: 480 application/json response, POST /apps/{appId}/branches: 481 application/json response, POST /apps/{appId}/branches: 482 application/json response, GET /apps/{appId}/branches: 200 application/json response, GET /apps/{appId}/branches: 480 application/json response, GET /apps/{appId}/branches: 481 application/json response, GET /apps/{appId}/branches: 482 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: request body, POST /apps/{appId}/branches/{branchName}/deployments/start: application/json request body, POST /apps/{appId}/branches/{branchName}/deployments/start: 484 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: 200 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: 480 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: 481 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: 482 application/json response, POST /apps/{appId}/branches/{branchName}/deployments/start: 483 application/json response; All schemas have descriptions: Context, ArtifactFileName, JobArn, SubDomains, ThumbnailUrl, WebhookUrl, DefaultDomain, OauthToken, FileUploadUrls, DisplayName, FileMap, WebhookId, DomainName, DeploymentArtifacts, ServiceRoleArn, LogUrl, DomainAssociations, AccessToken, CreateTime, ResourceArn, BackendEnvironments, EnablePerformanceMode, UnauthorizedException, AutoSubDomainCreationPattern, InternalFailureException, EnvKey, EnableAutoSubDomain, EnvironmentVariables, Source, CustomDomains, MD5Hash, Name, ActiveJobId, ArtifactUrl, PullRequestEnvironmentName, ThumbnailName, AutoBranchCreationPatterns, TestArtifactsUrl, JobId, StatusReason, TagKeyList, CommitTime, AutoBranchCreationPattern, NotFoundException, StackName, CommitMessage, CommitId, EnableNotification, AppArn, Stage, CustomHeaders, SubDomainSettings, LimitExceededException, Condition, DomainStatus, EnableBasicAuth, ArtifactId, TotalNumberOfJobs, Steps, RepositoryCloneMethod, TagKey, WebhookArn, EnvironmentName, StartTime, NextToken, EnableAutoBranchCreation, Framework, CertificateVerificationDNSRecord, TagMap, AssociatedResources, BadRequestException, BackendEnvironmentArn, BasicAuthCredentials, EnablePullRequestPreview, MaxResults, JobReason, TagValue, Apps, AutoSubDomainIAMRole, BranchName, EndTime, Screenshots, SourceUrl, Status, Target, ResourceNotFoundException, FileName, Artifacts, EnvValue, DependentServiceFailureException, Verified, GetAppResult, GetJobResult, JobSummaries, GetBranchResult, UploadUrl, Branches, ArtifactsUrl, CreateAppResult, CustomRules, AppId, DomainAssociationArn, EnableAutoBuild, CustomDomain, UpdateTime, StepName, AutoSubDomainCreationPatterns, Repository, JobStatus, DomainPrefix, EnableBranchAutoDeletion, BranchArn, JobType, LastDeployTime, DNSRecord, Webhooks, DeleteDomainAssociationResult, Platform, EnableBranchAutoBuild, TestConfigUrl, Description, AssociatedResource; All enums have descriptions: Stage: enum value PRODUCTION, Stage: enum value BETA, Stage: enum value DEVELOPMENT, Stage: enum value EXPERIMENTAL, Stage: enum value PULL_REQUEST, DomainStatus: enum value PENDING_VERIFICATION, DomainStatus: enum value IN_PROGRESS, DomainStatus: enum value AVAILABLE, DomainStatus: enum value PENDING_DEPLOYMENT, DomainStatus: enum value FAILED, DomainStatus: enum value CREATING, DomainStatus: enum value REQUESTING_CERTIFICATE, DomainStatus: enum value UPDATING, RepositoryCloneMethod: enum value SSH, RepositoryCloneMethod: enum value TOKEN, RepositoryCloneMethod: enum value SIGV4, JobStatus: enum value PENDING, JobStatus: enum value PROVISIONING, JobStatus: enum value RUNNING, JobStatus: enum value FAILED, JobStatus: enum value SUCCEED, JobStatus: enum value CANCELLING, JobStatus: enum value CANCELLED, JobType: enum value RELEASE, JobType: enum value RETRY, JobType: enum value MANUAL, JobType: enum value WEB_HOOK, Platform: enum value WEB, Platform: enum value WEB_DYNAMIC, Platform: enum value WEB_COMPUTE
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
        "Stage: enum value PRODUCTION",
        "Stage: enum value BETA",
        "Stage: enum value DEVELOPMENT",
        "Stage: enum value EXPERIMENTAL",
        "Stage: enum value PULL_REQUEST",
        "DomainStatus: enum value PENDING_VERIFICATION",
        "DomainStatus: enum value IN_PROGRESS",
        "DomainStatus: enum value AVAILABLE",
        "DomainStatus: enum value PENDING_DEPLOYMENT",
        "DomainStatus: enum value FAILED",
        "DomainStatus: enum value CREATING",
        "DomainStatus: enum value REQUESTING_CERTIFICATE",
        "DomainStatus: enum value UPDATING",
        "RepositoryCloneMethod: enum value SSH",
        "RepositoryCloneMethod: enum value TOKEN",
        "RepositoryCloneMethod: enum value SIGV4",
        "JobStatus: enum value PENDING",
        "JobStatus: enum value PROVISIONING",
        "JobStatus: enum value RUNNING",
        "JobStatus: enum value FAILED",
        "JobStatus: enum value SUCCEED",
        "JobStatus: enum value CANCELLING",
        "JobStatus: enum value CANCELLED",
        "JobType: enum value RELEASE",
        "JobType: enum value RETRY",
        "JobType: enum value MANUAL",
        "JobType: enum value WEB_HOOK",
        "Platform: enum value WEB",
        "Platform: enum value WEB_DYNAMIC",
        "Platform: enum value WEB_COMPUTE"
      ],
      "All operations have clear summaries": [
        "GET /apps/{appId}/domains/{domainName}",
        "POST /apps/{appId}/domains/{domainName}",
        "DELETE /apps/{appId}/domains/{domainName}",
        "GET /artifacts/{artifactId}",
        "GET /apps/{appId}/backendenvironments",
        "POST /apps/{appId}/backendenvironments",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop",
        "GET /tags/{resourceArn}",
        "POST /tags/{resourceArn}",
        "GET /apps",
        "POST /apps",
        "POST /apps/{appId}/accesslogs",
        "POST /apps/{appId}/branches/{branchName}/deployments",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts",
        "DELETE /webhooks/{webhookId}",
        "GET /webhooks/{webhookId}",
        "POST /webhooks/{webhookId}",
        "DELETE /apps/{appId}",
        "GET /apps/{appId}",
        "POST /apps/{appId}",
        "DELETE /tags/{resourceArn}#tagKeys",
        "POST /apps/{appId}/domains",
        "GET /apps/{appId}/domains",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}",
        "GET /apps/{appId}/webhooks",
        "POST /apps/{appId}/webhooks",
        "DELETE /apps/{appId}/branches/{branchName}",
        "GET /apps/{appId}/branches/{branchName}",
        "POST /apps/{appId}/branches/{branchName}",
        "GET /apps/{appId}/branches/{branchName}/jobs",
        "POST /apps/{appId}/branches/{branchName}/jobs",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}",
        "GET /apps/{appId}/backendenvironments/{environmentName}",
        "POST /apps/{appId}/branches",
        "GET /apps/{appId}/branches",
        "POST /apps/{appId}/branches/{branchName}/deployments/start"
      ],
      "All request/response bodies have examples": [
        "GET /apps/{appId}/domains/{domainName}: 200 application/json response",
        "GET /apps/{appId}/domains/{domainName}: 480 application/json response",
        "GET /apps/{appId}/domains/{domainName}: 481 application/json response",
        "GET /apps/{appId}/domains/{domainName}: 482 application/json response",
        "GET /apps/{appId}/domains/{domainName}: 483 application/json response",
        "POST /apps/{appId}/domains/{domainName}: request body",
        "POST /apps/{appId}/domains/{domainName}: application/json request body",
        "POST /apps/{appId}/domains/{domainName}: 482 application/json response",
        "POST /apps/{appId}/domains/{domainName}: 483 application/json response",
        "POST /apps/{appId}/domains/{domainName}: 484 application/json response",
        "POST /apps/{appId}/domains/{domainName}: 200 application/json response",
        "POST /apps/{appId}/domains/{domainName}: 480 application/json response",
        "POST /apps/{appId}/domains/{domainName}: 481 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 480 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 481 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 482 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 483 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 484 application/json response",
        "DELETE /apps/{appId}/domains/{domainName}: 200 application/json response",
        "GET /artifacts/{artifactId}: 200 application/json response",
        "GET /artifacts/{artifactId}: 480 application/json response",
        "GET /artifacts/{artifactId}: 481 application/json response",
        "GET /artifacts/{artifactId}: 482 application/json response",
        "GET /artifacts/{artifactId}: 483 application/json response",
        "GET /artifacts/{artifactId}: 484 application/json response",
        "GET /apps/{appId}/backendenvironments: 481 application/json response",
        "GET /apps/{appId}/backendenvironments: 482 application/json response",
        "GET /apps/{appId}/backendenvironments: 200 application/json response",
        "GET /apps/{appId}/backendenvironments: 480 application/json response",
        "POST /apps/{appId}/backendenvironments: request body",
        "POST /apps/{appId}/backendenvironments: application/json request body",
        "POST /apps/{appId}/backendenvironments: 481 application/json response",
        "POST /apps/{appId}/backendenvironments: 482 application/json response",
        "POST /apps/{appId}/backendenvironments: 483 application/json response",
        "POST /apps/{appId}/backendenvironments: 484 application/json response",
        "POST /apps/{appId}/backendenvironments: 200 application/json response",
        "POST /apps/{appId}/backendenvironments: 480 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 200 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 480 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 481 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 482 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 483 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 484 application/json response",
        "GET /tags/{resourceArn}: 200 application/json response",
        "GET /tags/{resourceArn}: 480 application/json response",
        "GET /tags/{resourceArn}: 481 application/json response",
        "GET /tags/{resourceArn}: 482 application/json response",
        "POST /tags/{resourceArn}: request body",
        "POST /tags/{resourceArn}: application/json request body",
        "POST /tags/{resourceArn}: 480 application/json response",
        "POST /tags/{resourceArn}: 481 application/json response",
        "POST /tags/{resourceArn}: 482 application/json response",
        "POST /tags/{resourceArn}: 200 application/json response",
        "GET /apps: 481 application/json response",
        "GET /apps: 482 application/json response",
        "GET /apps: 200 application/json response",
        "GET /apps: 480 application/json response",
        "POST /apps: request body",
        "POST /apps: application/json request body",
        "POST /apps: 483 application/json response",
        "POST /apps: 484 application/json response",
        "POST /apps: 200 application/json response",
        "POST /apps: 480 application/json response",
        "POST /apps: 481 application/json response",
        "POST /apps: 482 application/json response",
        "POST /apps/{appId}/accesslogs: request body",
        "POST /apps/{appId}/accesslogs: application/json request body",
        "POST /apps/{appId}/accesslogs: 482 application/json response",
        "POST /apps/{appId}/accesslogs: 483 application/json response",
        "POST /apps/{appId}/accesslogs: 200 application/json response",
        "POST /apps/{appId}/accesslogs: 480 application/json response",
        "POST /apps/{appId}/accesslogs: 481 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments: request body",
        "POST /apps/{appId}/branches/{branchName}/deployments: application/json request body",
        "POST /apps/{appId}/branches/{branchName}/deployments: 482 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 483 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 200 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 480 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 481 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 481 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 482 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 483 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 200 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 480 application/json response",
        "DELETE /webhooks/{webhookId}: 480 application/json response",
        "DELETE /webhooks/{webhookId}: 481 application/json response",
        "DELETE /webhooks/{webhookId}: 482 application/json response",
        "DELETE /webhooks/{webhookId}: 483 application/json response",
        "DELETE /webhooks/{webhookId}: 484 application/json response",
        "DELETE /webhooks/{webhookId}: 200 application/json response",
        "GET /webhooks/{webhookId}: 481 application/json response",
        "GET /webhooks/{webhookId}: 482 application/json response",
        "GET /webhooks/{webhookId}: 483 application/json response",
        "GET /webhooks/{webhookId}: 484 application/json response",
        "GET /webhooks/{webhookId}: 200 application/json response",
        "GET /webhooks/{webhookId}: 480 application/json response",
        "POST /webhooks/{webhookId}: request body",
        "POST /webhooks/{webhookId}: application/json request body",
        "POST /webhooks/{webhookId}: 481 application/json response",
        "POST /webhooks/{webhookId}: 482 application/json response",
        "POST /webhooks/{webhookId}: 483 application/json response",
        "POST /webhooks/{webhookId}: 484 application/json response",
        "POST /webhooks/{webhookId}: 200 application/json response",
        "POST /webhooks/{webhookId}: 480 application/json response",
        "DELETE /apps/{appId}: 481 application/json response",
        "DELETE /apps/{appId}: 482 application/json response",
        "DELETE /apps/{appId}: 483 application/json response",
        "DELETE /apps/{appId}: 484 application/json response",
        "DELETE /apps/{appId}: 200 application/json response",
        "DELETE /apps/{appId}: 480 application/json response",
        "GET /apps/{appId}: 481 application/json response",
        "GET /apps/{appId}: 482 application/json response",
        "GET /apps/{appId}: 483 application/json response",
        "GET /apps/{appId}: 200 application/json response",
        "GET /apps/{appId}: 480 application/json response",
        "POST /apps/{appId}: request body",
        "POST /apps/{appId}: application/json request body",
        "POST /apps/{appId}: 481 application/json response",
        "POST /apps/{appId}: 482 application/json response",
        "POST /apps/{appId}: 483 application/json response",
        "POST /apps/{appId}: 200 application/json response",
        "POST /apps/{appId}: 480 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 200 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 480 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 481 application/json response",
        "DELETE /tags/{resourceArn}#tagKeys: 482 application/json response",
        "POST /apps/{appId}/domains: request body",
        "POST /apps/{appId}/domains: application/json request body",
        "POST /apps/{appId}/domains: 485 application/json response",
        "POST /apps/{appId}/domains: 200 application/json response",
        "POST /apps/{appId}/domains: 480 application/json response",
        "POST /apps/{appId}/domains: 481 application/json response",
        "POST /apps/{appId}/domains: 482 application/json response",
        "POST /apps/{appId}/domains: 483 application/json response",
        "POST /apps/{appId}/domains: 484 application/json response",
        "GET /apps/{appId}/domains: 200 application/json response",
        "GET /apps/{appId}/domains: 480 application/json response",
        "GET /apps/{appId}/domains: 481 application/json response",
        "GET /apps/{appId}/domains: 482 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 200 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 200 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 application/json response",
        "GET /apps/{appId}/webhooks: 482 application/json response",
        "GET /apps/{appId}/webhooks: 483 application/json response",
        "GET /apps/{appId}/webhooks: 200 application/json response",
        "GET /apps/{appId}/webhooks: 480 application/json response",
        "GET /apps/{appId}/webhooks: 481 application/json response",
        "POST /apps/{appId}/webhooks: request body",
        "POST /apps/{appId}/webhooks: application/json request body",
        "POST /apps/{appId}/webhooks: 485 application/json response",
        "POST /apps/{appId}/webhooks: 200 application/json response",
        "POST /apps/{appId}/webhooks: 480 application/json response",
        "POST /apps/{appId}/webhooks: 481 application/json response",
        "POST /apps/{appId}/webhooks: 482 application/json response",
        "POST /apps/{appId}/webhooks: 483 application/json response",
        "POST /apps/{appId}/webhooks: 484 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 480 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 481 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 482 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 483 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 484 application/json response",
        "DELETE /apps/{appId}/branches/{branchName}: 200 application/json response",
        "GET /apps/{appId}/branches/{branchName}: 200 application/json response",
        "GET /apps/{appId}/branches/{branchName}: 480 application/json response",
        "GET /apps/{appId}/branches/{branchName}: 481 application/json response",
        "GET /apps/{appId}/branches/{branchName}: 482 application/json response",
        "GET /apps/{appId}/branches/{branchName}: 483 application/json response",
        "POST /apps/{appId}/branches/{branchName}: request body",
        "POST /apps/{appId}/branches/{branchName}: application/json request body",
        "POST /apps/{appId}/branches/{branchName}: 481 application/json response",
        "POST /apps/{appId}/branches/{branchName}: 482 application/json response",
        "POST /apps/{appId}/branches/{branchName}: 483 application/json response",
        "POST /apps/{appId}/branches/{branchName}: 484 application/json response",
        "POST /apps/{appId}/branches/{branchName}: 200 application/json response",
        "POST /apps/{appId}/branches/{branchName}: 480 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 480 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 481 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 482 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 483 application/json response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 200 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: request body",
        "POST /apps/{appId}/branches/{branchName}/jobs: application/json request body",
        "POST /apps/{appId}/branches/{branchName}/jobs: 481 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 482 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 483 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 484 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 200 application/json response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 480 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 484 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 200 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 480 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 481 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 482 application/json response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 483 application/json response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 480 application/json response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 481 application/json response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 482 application/json response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 483 application/json response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 200 application/json response",
        "POST /apps/{appId}/branches: request body",
        "POST /apps/{appId}/branches: application/json request body",
        "POST /apps/{appId}/branches: 483 application/json response",
        "POST /apps/{appId}/branches: 484 application/json response",
        "POST /apps/{appId}/branches: 485 application/json response",
        "POST /apps/{appId}/branches: 200 application/json response",
        "POST /apps/{appId}/branches: 480 application/json response",
        "POST /apps/{appId}/branches: 481 application/json response",
        "POST /apps/{appId}/branches: 482 application/json response",
        "GET /apps/{appId}/branches: 200 application/json response",
        "GET /apps/{appId}/branches: 480 application/json response",
        "GET /apps/{appId}/branches: 481 application/json response",
        "GET /apps/{appId}/branches: 482 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: request body",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: application/json request body",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 484 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 200 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 480 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 481 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 482 application/json response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 483 application/json response"
      ],
      "All schemas have descriptions": [
        "Context",
        "ArtifactFileName",
        "JobArn",
        "SubDomains",
        "ThumbnailUrl",
        "WebhookUrl",
        "DefaultDomain",
        "OauthToken",
        "FileUploadUrls",
        "DisplayName",
        "FileMap",
        "WebhookId",
        "DomainName",
        "DeploymentArtifacts",
        "ServiceRoleArn",
        "LogUrl",
        "DomainAssociations",
        "AccessToken",
        "CreateTime",
        "ResourceArn",
        "BackendEnvironments",
        "EnablePerformanceMode",
        "UnauthorizedException",
        "AutoSubDomainCreationPattern",
        "InternalFailureException",
        "EnvKey",
        "EnableAutoSubDomain",
        "EnvironmentVariables",
        "Source",
        "CustomDomains",
        "MD5Hash",
        "Name",
        "ActiveJobId",
        "ArtifactUrl",
        "PullRequestEnvironmentName",
        "ThumbnailName",
        "AutoBranchCreationPatterns",
        "TestArtifactsUrl",
        "JobId",
        "StatusReason",
        "TagKeyList",
        "CommitTime",
        "AutoBranchCreationPattern",
        "NotFoundException",
        "StackName",
        "CommitMessage",
        "CommitId",
        "EnableNotification",
        "AppArn",
        "Stage",
        "CustomHeaders",
        "SubDomainSettings",
        "LimitExceededException",
        "Condition",
        "DomainStatus",
        "EnableBasicAuth",
        "ArtifactId",
        "TotalNumberOfJobs",
        "Steps",
        "RepositoryCloneMethod",
        "TagKey",
        "WebhookArn",
        "EnvironmentName",
        "StartTime",
        "NextToken",
        "EnableAutoBranchCreation",
        "Framework",
        "CertificateVerificationDNSRecord",
        "TagMap",
        "AssociatedResources",
        "BadRequestException",
        "BackendEnvironmentArn",
        "BasicAuthCredentials",
        "EnablePullRequestPreview",
        "MaxResults",
        "JobReason",
        "TagValue",
        "Apps",
        "AutoSubDomainIAMRole",
        "BranchName",
        "EndTime",
        "Screenshots",
        "SourceUrl",
        "Status",
        "Target",
        "ResourceNotFoundException",
        "FileName",
        "Artifacts",
        "EnvValue",
        "DependentServiceFailureException",
        "Verified",
        "GetAppResult",
        "GetJobResult",
        "JobSummaries",
        "GetBranchResult",
        "UploadUrl",
        "Branches",
        "ArtifactsUrl",
        "CreateAppResult",
        "CustomRules",
        "AppId",
        "DomainAssociationArn",
        "EnableAutoBuild",
        "CustomDomain",
        "UpdateTime",
        "StepName",
        "AutoSubDomainCreationPatterns",
        "Repository",
        "JobStatus",
        "DomainPrefix",
        "EnableBranchAutoDeletion",
        "BranchArn",
        "JobType",
        "LastDeployTime",
        "DNSRecord",
        "Webhooks",
        "DeleteDomainAssociationResult",
        "Platform",
        "EnableBranchAutoBuild",
        "TestConfigUrl",
        "Description",
        "AssociatedResource"
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
- **Message:** Request validation issues found: All string fields have length constraints: /apps: parameter X-Amz-Content-Sha256, /apps: parameter X-Amz-Date, /apps: parameter X-Amz-Algorithm, /apps: parameter X-Amz-Credential, /apps: parameter X-Amz-Security-Token, /apps: parameter X-Amz-Signature, /apps: parameter X-Amz-SignedHeaders, POST /apps.platform: application/json schema, /apps/{appId}/accesslogs: parameter X-Amz-Content-Sha256, /apps/{appId}/accesslogs: parameter X-Amz-Date, /apps/{appId}/accesslogs: parameter X-Amz-Algorithm, /apps/{appId}/accesslogs: parameter X-Amz-Credential, /apps/{appId}/accesslogs: parameter X-Amz-Security-Token, /apps/{appId}/accesslogs: parameter X-Amz-Signature, /apps/{appId}/accesslogs: parameter X-Amz-SignedHeaders, POST /apps/{appId}/accesslogs.endTime: application/json schema, POST /apps/{appId}/accesslogs.startTime: application/json schema, /webhooks/{webhookId}: parameter X-Amz-Content-Sha256, /webhooks/{webhookId}: parameter X-Amz-Date, /webhooks/{webhookId}: parameter X-Amz-Algorithm, /webhooks/{webhookId}: parameter X-Amz-Credential, /webhooks/{webhookId}: parameter X-Amz-Security-Token, /webhooks/{webhookId}: parameter X-Amz-Signature, /webhooks/{webhookId}: parameter X-Amz-SignedHeaders, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-SignedHeaders, /apps/{appId}/backendenvironments: parameter X-Amz-Content-Sha256, /apps/{appId}/backendenvironments: parameter X-Amz-Date, /apps/{appId}/backendenvironments: parameter X-Amz-Algorithm, /apps/{appId}/backendenvironments: parameter X-Amz-Credential, /apps/{appId}/backendenvironments: parameter X-Amz-Security-Token, /apps/{appId}/backendenvironments: parameter X-Amz-Signature, /apps/{appId}/backendenvironments: parameter X-Amz-SignedHeaders, /apps/{appId}/branches: parameter X-Amz-Content-Sha256, /apps/{appId}/branches: parameter X-Amz-Date, /apps/{appId}/branches: parameter X-Amz-Algorithm, /apps/{appId}/branches: parameter X-Amz-Credential, /apps/{appId}/branches: parameter X-Amz-Security-Token, /apps/{appId}/branches: parameter X-Amz-Signature, /apps/{appId}/branches: parameter X-Amz-SignedHeaders, POST /apps/{appId}/branches.stage: application/json schema, /apps/{appId}/branches/{branchName}: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}: parameter X-Amz-SignedHeaders, POST /apps/{appId}/branches/{branchName}.stage: application/json schema, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-SignedHeaders, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-SignedHeaders, /apps/{appId}/webhooks: parameter X-Amz-Content-Sha256, /apps/{appId}/webhooks: parameter X-Amz-Date, /apps/{appId}/webhooks: parameter X-Amz-Algorithm, /apps/{appId}/webhooks: parameter X-Amz-Credential, /apps/{appId}/webhooks: parameter X-Amz-Security-Token, /apps/{appId}/webhooks: parameter X-Amz-Signature, /apps/{appId}/webhooks: parameter X-Amz-SignedHeaders, /tags/{resourceArn}: parameter X-Amz-Content-Sha256, /tags/{resourceArn}: parameter X-Amz-Date, /tags/{resourceArn}: parameter X-Amz-Algorithm, /tags/{resourceArn}: parameter X-Amz-Credential, /tags/{resourceArn}: parameter X-Amz-Security-Token, /tags/{resourceArn}: parameter X-Amz-Signature, /tags/{resourceArn}: parameter X-Amz-SignedHeaders, /apps/{appId}/domains: parameter X-Amz-Content-Sha256, /apps/{appId}/domains: parameter X-Amz-Date, /apps/{appId}/domains: parameter X-Amz-Algorithm, /apps/{appId}/domains: parameter X-Amz-Credential, /apps/{appId}/domains: parameter X-Amz-Security-Token, /apps/{appId}/domains: parameter X-Amz-Signature, /apps/{appId}/domains: parameter X-Amz-SignedHeaders, /artifacts/{artifactId}: parameter X-Amz-Content-Sha256, /artifacts/{artifactId}: parameter X-Amz-Date, /artifacts/{artifactId}: parameter X-Amz-Algorithm, /artifacts/{artifactId}: parameter X-Amz-Credential, /artifacts/{artifactId}: parameter X-Amz-Security-Token, /artifacts/{artifactId}: parameter X-Amz-Signature, /artifacts/{artifactId}: parameter X-Amz-SignedHeaders, /tags/{resourceArn}#tagKeys: parameter X-Amz-Content-Sha256, /tags/{resourceArn}#tagKeys: parameter X-Amz-Date, /tags/{resourceArn}#tagKeys: parameter X-Amz-Algorithm, /tags/{resourceArn}#tagKeys: parameter X-Amz-Credential, /tags/{resourceArn}#tagKeys: parameter X-Amz-Security-Token, /tags/{resourceArn}#tagKeys: parameter X-Amz-Signature, /tags/{resourceArn}#tagKeys: parameter X-Amz-SignedHeaders, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-SignedHeaders, /apps/{appId}/domains/{domainName}: parameter X-Amz-Content-Sha256, /apps/{appId}/domains/{domainName}: parameter X-Amz-Date, /apps/{appId}/domains/{domainName}: parameter X-Amz-Algorithm, /apps/{appId}/domains/{domainName}: parameter X-Amz-Credential, /apps/{appId}/domains/{domainName}: parameter X-Amz-Security-Token, /apps/{appId}/domains/{domainName}: parameter X-Amz-Signature, /apps/{appId}/domains/{domainName}: parameter X-Amz-SignedHeaders, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-SignedHeaders, POST /apps/{appId}/branches/{branchName}/jobs.commitTime: application/json schema, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Content-Sha256, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Date, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Algorithm, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Credential, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Security-Token, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Signature, /apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-SignedHeaders, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Content-Sha256, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Date, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Algorithm, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Credential, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Security-Token, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Signature, /apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-SignedHeaders, /apps/{appId}: parameter X-Amz-Content-Sha256, /apps/{appId}: parameter X-Amz-Date, /apps/{appId}: parameter X-Amz-Algorithm, /apps/{appId}: parameter X-Amz-Credential, /apps/{appId}: parameter X-Amz-Security-Token, /apps/{appId}: parameter X-Amz-Signature, /apps/{appId}: parameter X-Amz-SignedHeaders, POST /apps/{appId}.platform: application/json schema; All schemas specify data types: POST /apps.customRules[].condition: application/json schema, POST /apps.customRules[].source: application/json schema, POST /apps.customRules[].status: application/json schema, POST /apps.customRules[].target: application/json schema, POST /apps.autoBranchCreationConfig.enablePerformanceMode: application/json schema, POST /apps.autoBranchCreationConfig.enablePullRequestPreview: application/json schema, POST /apps.autoBranchCreationConfig.environmentVariables: application/json schema, POST /apps.autoBranchCreationConfig.framework: application/json schema, POST /apps.autoBranchCreationConfig.stage: application/json schema, POST /apps.autoBranchCreationConfig.enableAutoBuild: application/json schema, POST /apps.autoBranchCreationConfig.buildSpec: application/json schema, POST /apps.autoBranchCreationConfig.enableBasicAuth: application/json schema, POST /apps.autoBranchCreationConfig.pullRequestEnvironmentName: application/json schema, POST /apps.autoBranchCreationConfig.basicAuthCredentials: application/json schema, POST /apps/{appId}/domains.subDomainSettings[].prefix: application/json schema, POST /apps/{appId}/domains.subDomainSettings[].branchName: application/json schema, POST /apps/{appId}/domains/{domainName}.subDomainSettings[].branchName: application/json schema, POST /apps/{appId}/domains/{domainName}.subDomainSettings[].prefix: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.buildSpec: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.enableAutoBuild: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.environmentVariables: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.pullRequestEnvironmentName: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.basicAuthCredentials: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.framework: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.enableBasicAuth: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.enablePerformanceMode: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.enablePullRequestPreview: application/json schema, POST /apps/{appId}.autoBranchCreationConfig.stage: application/json schema, POST /apps/{appId}.customRules[].target: application/json schema, POST /apps/{appId}.customRules[].condition: application/json schema, POST /apps/{appId}.customRules[].source: application/json schema, POST /apps/{appId}.customRules[].status: application/json schema
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
        "POST /apps.customRules[].condition: application/json schema",
        "POST /apps.customRules[].source: application/json schema",
        "POST /apps.customRules[].status: application/json schema",
        "POST /apps.customRules[].target: application/json schema",
        "POST /apps.autoBranchCreationConfig.enablePerformanceMode: application/json schema",
        "POST /apps.autoBranchCreationConfig.enablePullRequestPreview: application/json schema",
        "POST /apps.autoBranchCreationConfig.environmentVariables: application/json schema",
        "POST /apps.autoBranchCreationConfig.framework: application/json schema",
        "POST /apps.autoBranchCreationConfig.stage: application/json schema",
        "POST /apps.autoBranchCreationConfig.enableAutoBuild: application/json schema",
        "POST /apps.autoBranchCreationConfig.buildSpec: application/json schema",
        "POST /apps.autoBranchCreationConfig.enableBasicAuth: application/json schema",
        "POST /apps.autoBranchCreationConfig.pullRequestEnvironmentName: application/json schema",
        "POST /apps.autoBranchCreationConfig.basicAuthCredentials: application/json schema",
        "POST /apps/{appId}/domains.subDomainSettings[].prefix: application/json schema",
        "POST /apps/{appId}/domains.subDomainSettings[].branchName: application/json schema",
        "POST /apps/{appId}/domains/{domainName}.subDomainSettings[].branchName: application/json schema",
        "POST /apps/{appId}/domains/{domainName}.subDomainSettings[].prefix: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.buildSpec: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.enableAutoBuild: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.environmentVariables: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.pullRequestEnvironmentName: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.basicAuthCredentials: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.framework: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.enableBasicAuth: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.enablePerformanceMode: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.enablePullRequestPreview: application/json schema",
        "POST /apps/{appId}.autoBranchCreationConfig.stage: application/json schema",
        "POST /apps/{appId}.customRules[].target: application/json schema",
        "POST /apps/{appId}.customRules[].condition: application/json schema",
        "POST /apps/{appId}.customRules[].source: application/json schema",
        "POST /apps/{appId}.customRules[].status: application/json schema"
      ],
      "All string fields have length constraints": [
        "/apps: parameter X-Amz-Content-Sha256",
        "/apps: parameter X-Amz-Date",
        "/apps: parameter X-Amz-Algorithm",
        "/apps: parameter X-Amz-Credential",
        "/apps: parameter X-Amz-Security-Token",
        "/apps: parameter X-Amz-Signature",
        "/apps: parameter X-Amz-SignedHeaders",
        "POST /apps.platform: application/json schema",
        "/apps/{appId}/accesslogs: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/accesslogs: parameter X-Amz-Date",
        "/apps/{appId}/accesslogs: parameter X-Amz-Algorithm",
        "/apps/{appId}/accesslogs: parameter X-Amz-Credential",
        "/apps/{appId}/accesslogs: parameter X-Amz-Security-Token",
        "/apps/{appId}/accesslogs: parameter X-Amz-Signature",
        "/apps/{appId}/accesslogs: parameter X-Amz-SignedHeaders",
        "POST /apps/{appId}/accesslogs.endTime: application/json schema",
        "POST /apps/{appId}/accesslogs.startTime: application/json schema",
        "/webhooks/{webhookId}: parameter X-Amz-Content-Sha256",
        "/webhooks/{webhookId}: parameter X-Amz-Date",
        "/webhooks/{webhookId}: parameter X-Amz-Algorithm",
        "/webhooks/{webhookId}: parameter X-Amz-Credential",
        "/webhooks/{webhookId}: parameter X-Amz-Security-Token",
        "/webhooks/{webhookId}: parameter X-Amz-Signature",
        "/webhooks/{webhookId}: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Date",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Algorithm",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Credential",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Security-Token",
        "/apps/{appId}/backendenvironments: parameter X-Amz-Signature",
        "/apps/{appId}/backendenvironments: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches: parameter X-Amz-Date",
        "/apps/{appId}/branches: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches: parameter X-Amz-Credential",
        "/apps/{appId}/branches: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches: parameter X-Amz-Signature",
        "/apps/{appId}/branches: parameter X-Amz-SignedHeaders",
        "POST /apps/{appId}/branches.stage: application/json schema",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}: parameter X-Amz-SignedHeaders",
        "POST /apps/{appId}/branches/{branchName}.stage: application/json schema",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/deployments: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/webhooks: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/webhooks: parameter X-Amz-Date",
        "/apps/{appId}/webhooks: parameter X-Amz-Algorithm",
        "/apps/{appId}/webhooks: parameter X-Amz-Credential",
        "/apps/{appId}/webhooks: parameter X-Amz-Security-Token",
        "/apps/{appId}/webhooks: parameter X-Amz-Signature",
        "/apps/{appId}/webhooks: parameter X-Amz-SignedHeaders",
        "/tags/{resourceArn}: parameter X-Amz-Content-Sha256",
        "/tags/{resourceArn}: parameter X-Amz-Date",
        "/tags/{resourceArn}: parameter X-Amz-Algorithm",
        "/tags/{resourceArn}: parameter X-Amz-Credential",
        "/tags/{resourceArn}: parameter X-Amz-Security-Token",
        "/tags/{resourceArn}: parameter X-Amz-Signature",
        "/tags/{resourceArn}: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/domains: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/domains: parameter X-Amz-Date",
        "/apps/{appId}/domains: parameter X-Amz-Algorithm",
        "/apps/{appId}/domains: parameter X-Amz-Credential",
        "/apps/{appId}/domains: parameter X-Amz-Security-Token",
        "/apps/{appId}/domains: parameter X-Amz-Signature",
        "/apps/{appId}/domains: parameter X-Amz-SignedHeaders",
        "/artifacts/{artifactId}: parameter X-Amz-Content-Sha256",
        "/artifacts/{artifactId}: parameter X-Amz-Date",
        "/artifacts/{artifactId}: parameter X-Amz-Algorithm",
        "/artifacts/{artifactId}: parameter X-Amz-Credential",
        "/artifacts/{artifactId}: parameter X-Amz-Security-Token",
        "/artifacts/{artifactId}: parameter X-Amz-Signature",
        "/artifacts/{artifactId}: parameter X-Amz-SignedHeaders",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Content-Sha256",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Date",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Algorithm",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Credential",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Security-Token",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-Signature",
        "/tags/{resourceArn}#tagKeys: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/deployments/start: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Date",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Algorithm",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Credential",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Security-Token",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-Signature",
        "/apps/{appId}/domains/{domainName}: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/jobs: parameter X-Amz-SignedHeaders",
        "POST /apps/{appId}/branches/{branchName}/jobs.commitTime: application/json schema",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Date",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Algorithm",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Credential",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Security-Token",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-Signature",
        "/apps/{appId}/backendenvironments/{environmentName}: parameter X-Amz-SignedHeaders",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Content-Sha256",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Date",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Algorithm",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Credential",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Security-Token",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-Signature",
        "/apps/{appId}/branches/{branchName}/jobs/{jobId}: parameter X-Amz-SignedHeaders",
        "/apps/{appId}: parameter X-Amz-Content-Sha256",
        "/apps/{appId}: parameter X-Amz-Date",
        "/apps/{appId}: parameter X-Amz-Algorithm",
        "/apps/{appId}: parameter X-Amz-Credential",
        "/apps/{appId}: parameter X-Amz-Security-Token",
        "/apps/{appId}: parameter X-Amz-Signature",
        "/apps/{appId}: parameter X-Amz-SignedHeaders",
        "POST /apps/{appId}.platform: application/json schema"
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
- **Message:** Error handling issues found: Error responses include error details schema: POST /tags/{resourceArn}: 480 response, POST /tags/{resourceArn}: 481 response, POST /tags/{resourceArn}: 482 response, GET /tags/{resourceArn}: 481 response, GET /tags/{resourceArn}: 482 response, GET /tags/{resourceArn}: 480 response, DELETE /apps/{appId}/domains/{domainName}: 482 response, DELETE /apps/{appId}/domains/{domainName}: 483 response, DELETE /apps/{appId}/domains/{domainName}: 484 response, DELETE /apps/{appId}/domains/{domainName}: 480 response, DELETE /apps/{appId}/domains/{domainName}: 481 response, GET /apps/{appId}/domains/{domainName}: 483 response, GET /apps/{appId}/domains/{domainName}: 480 response, GET /apps/{appId}/domains/{domainName}: 481 response, GET /apps/{appId}/domains/{domainName}: 482 response, POST /apps/{appId}/domains/{domainName}: 481 response, POST /apps/{appId}/domains/{domainName}: 482 response, POST /apps/{appId}/domains/{domainName}: 483 response, POST /apps/{appId}/domains/{domainName}: 484 response, POST /apps/{appId}/domains/{domainName}: 480 response, DELETE /webhooks/{webhookId}: 483 response, DELETE /webhooks/{webhookId}: 484 response, DELETE /webhooks/{webhookId}: 480 response, DELETE /webhooks/{webhookId}: 481 response, DELETE /webhooks/{webhookId}: 482 response, GET /webhooks/{webhookId}: 484 response, GET /webhooks/{webhookId}: 480 response, GET /webhooks/{webhookId}: 481 response, GET /webhooks/{webhookId}: 482 response, GET /webhooks/{webhookId}: 483 response, POST /webhooks/{webhookId}: 484 response, POST /webhooks/{webhookId}: 480 response, POST /webhooks/{webhookId}: 481 response, POST /webhooks/{webhookId}: 482 response, POST /webhooks/{webhookId}: 483 response, POST /apps/{appId}/webhooks: 482 response, POST /apps/{appId}/webhooks: 483 response, POST /apps/{appId}/webhooks: 484 response, POST /apps/{appId}/webhooks: 485 response, POST /apps/{appId}/webhooks: 480 response, POST /apps/{appId}/webhooks: 481 response, GET /apps/{appId}/webhooks: 481 response, GET /apps/{appId}/webhooks: 482 response, GET /apps/{appId}/webhooks: 483 response, GET /apps/{appId}/webhooks: 480 response, DELETE /tags/{resourceArn}#tagKeys: 480 response, DELETE /tags/{resourceArn}#tagKeys: 481 response, DELETE /tags/{resourceArn}#tagKeys: 482 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 481 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 482 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 483 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 480 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 480 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 481 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 482 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 483 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 484 response, GET /apps: 480 response, GET /apps: 481 response, GET /apps: 482 response, POST /apps: 481 response, POST /apps: 482 response, POST /apps: 483 response, POST /apps: 484 response, POST /apps: 480 response, GET /apps/{appId}/branches/{branchName}: 480 response, GET /apps/{appId}/branches/{branchName}: 481 response, GET /apps/{appId}/branches/{branchName}: 482 response, GET /apps/{appId}/branches/{branchName}: 483 response, POST /apps/{appId}/branches/{branchName}: 481 response, POST /apps/{appId}/branches/{branchName}: 482 response, POST /apps/{appId}/branches/{branchName}: 483 response, POST /apps/{appId}/branches/{branchName}: 484 response, POST /apps/{appId}/branches/{branchName}: 480 response, DELETE /apps/{appId}/branches/{branchName}: 481 response, DELETE /apps/{appId}/branches/{branchName}: 482 response, DELETE /apps/{appId}/branches/{branchName}: 483 response, DELETE /apps/{appId}/branches/{branchName}: 484 response, DELETE /apps/{appId}/branches/{branchName}: 480 response, POST /apps/{appId}/branches/{branchName}/deployments: 480 response, POST /apps/{appId}/branches/{branchName}/deployments: 481 response, POST /apps/{appId}/branches/{branchName}/deployments: 482 response, POST /apps/{appId}/branches/{branchName}/deployments: 483 response, DELETE /apps/{appId}: 481 response, DELETE /apps/{appId}: 482 response, DELETE /apps/{appId}: 483 response, DELETE /apps/{appId}: 484 response, DELETE /apps/{appId}: 480 response, GET /apps/{appId}: 482 response, GET /apps/{appId}: 483 response, GET /apps/{appId}: 480 response, GET /apps/{appId}: 481 response, POST /apps/{appId}: 483 response, POST /apps/{appId}: 480 response, POST /apps/{appId}: 481 response, POST /apps/{appId}: 482 response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 484 response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 480 response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 481 response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 482 response, DELETE /apps/{appId}/backendenvironments/{environmentName}: 483 response, GET /apps/{appId}/backendenvironments/{environmentName}: 483 response, GET /apps/{appId}/backendenvironments/{environmentName}: 480 response, GET /apps/{appId}/backendenvironments/{environmentName}: 481 response, GET /apps/{appId}/backendenvironments/{environmentName}: 482 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 response, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 response, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 response, POST /apps/{appId}/accesslogs: 480 response, POST /apps/{appId}/accesslogs: 481 response, POST /apps/{appId}/accesslogs: 482 response, POST /apps/{appId}/accesslogs: 483 response, POST /apps/{appId}/branches/{branchName}/deployments/start: 484 response, POST /apps/{appId}/branches/{branchName}/deployments/start: 480 response, POST /apps/{appId}/branches/{branchName}/deployments/start: 481 response, POST /apps/{appId}/branches/{branchName}/deployments/start: 482 response, POST /apps/{appId}/branches/{branchName}/deployments/start: 483 response, POST /apps/{appId}/branches/{branchName}/jobs: 480 response, POST /apps/{appId}/branches/{branchName}/jobs: 481 response, POST /apps/{appId}/branches/{branchName}/jobs: 482 response, POST /apps/{appId}/branches/{branchName}/jobs: 483 response, POST /apps/{appId}/branches/{branchName}/jobs: 484 response, GET /apps/{appId}/branches/{branchName}/jobs: 480 response, GET /apps/{appId}/branches/{branchName}/jobs: 481 response, GET /apps/{appId}/branches/{branchName}/jobs: 482 response, GET /apps/{appId}/branches/{branchName}/jobs: 483 response, GET /apps/{appId}/domains: 480 response, GET /apps/{appId}/domains: 481 response, GET /apps/{appId}/domains: 482 response, POST /apps/{appId}/domains: 483 response, POST /apps/{appId}/domains: 484 response, POST /apps/{appId}/domains: 485 response, POST /apps/{appId}/domains: 480 response, POST /apps/{appId}/domains: 481 response, POST /apps/{appId}/domains: 482 response, GET /artifacts/{artifactId}: 480 response, GET /artifacts/{artifactId}: 481 response, GET /artifacts/{artifactId}: 482 response, GET /artifacts/{artifactId}: 483 response, GET /artifacts/{artifactId}: 484 response, GET /apps/{appId}/backendenvironments: 480 response, GET /apps/{appId}/backendenvironments: 481 response, GET /apps/{appId}/backendenvironments: 482 response, POST /apps/{appId}/backendenvironments: 480 response, POST /apps/{appId}/backendenvironments: 481 response, POST /apps/{appId}/backendenvironments: 482 response, POST /apps/{appId}/backendenvironments: 483 response, POST /apps/{appId}/backendenvironments: 484 response, GET /apps/{appId}/branches: 481 response, GET /apps/{appId}/branches: 482 response, GET /apps/{appId}/branches: 480 response, POST /apps/{appId}/branches: 484 response, POST /apps/{appId}/branches: 485 response, POST /apps/{appId}/branches: 480 response, POST /apps/{appId}/branches: 481 response, POST /apps/{appId}/branches: 482 response, POST /apps/{appId}/branches: 483 response; All operations document 5xx error responses: POST /tags/{resourceArn}, GET /tags/{resourceArn}, DELETE /apps/{appId}/domains/{domainName}, GET /apps/{appId}/domains/{domainName}, POST /apps/{appId}/domains/{domainName}, DELETE /webhooks/{webhookId}, GET /webhooks/{webhookId}, POST /webhooks/{webhookId}, POST /apps/{appId}/webhooks, GET /apps/{appId}/webhooks, DELETE /tags/{resourceArn}#tagKeys, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop, GET /apps, POST /apps, GET /apps/{appId}/branches/{branchName}, POST /apps/{appId}/branches/{branchName}, DELETE /apps/{appId}/branches/{branchName}, POST /apps/{appId}/branches/{branchName}/deployments, DELETE /apps/{appId}, GET /apps/{appId}, POST /apps/{appId}, DELETE /apps/{appId}/backendenvironments/{environmentName}, GET /apps/{appId}/backendenvironments/{environmentName}, DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}, GET /apps/{appId}/branches/{branchName}/jobs/{jobId}, POST /apps/{appId}/accesslogs, POST /apps/{appId}/branches/{branchName}/deployments/start, POST /apps/{appId}/branches/{branchName}/jobs, GET /apps/{appId}/branches/{branchName}/jobs, GET /apps/{appId}/domains, POST /apps/{appId}/domains, GET /artifacts/{artifactId}, GET /apps/{appId}/backendenvironments, POST /apps/{appId}/backendenvironments, GET /apps/{appId}/branches, POST /apps/{appId}/branches
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
        "POST /tags/{resourceArn}",
        "GET /tags/{resourceArn}",
        "DELETE /apps/{appId}/domains/{domainName}",
        "GET /apps/{appId}/domains/{domainName}",
        "POST /apps/{appId}/domains/{domainName}",
        "DELETE /webhooks/{webhookId}",
        "GET /webhooks/{webhookId}",
        "POST /webhooks/{webhookId}",
        "POST /apps/{appId}/webhooks",
        "GET /apps/{appId}/webhooks",
        "DELETE /tags/{resourceArn}#tagKeys",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop",
        "GET /apps",
        "POST /apps",
        "GET /apps/{appId}/branches/{branchName}",
        "POST /apps/{appId}/branches/{branchName}",
        "DELETE /apps/{appId}/branches/{branchName}",
        "POST /apps/{appId}/branches/{branchName}/deployments",
        "DELETE /apps/{appId}",
        "GET /apps/{appId}",
        "POST /apps/{appId}",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}",
        "GET /apps/{appId}/backendenvironments/{environmentName}",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}",
        "POST /apps/{appId}/accesslogs",
        "POST /apps/{appId}/branches/{branchName}/deployments/start",
        "POST /apps/{appId}/branches/{branchName}/jobs",
        "GET /apps/{appId}/branches/{branchName}/jobs",
        "GET /apps/{appId}/domains",
        "POST /apps/{appId}/domains",
        "GET /artifacts/{artifactId}",
        "GET /apps/{appId}/backendenvironments",
        "POST /apps/{appId}/backendenvironments",
        "GET /apps/{appId}/branches",
        "POST /apps/{appId}/branches"
      ],
      "Error responses include error details schema": [
        "POST /tags/{resourceArn}: 480 response",
        "POST /tags/{resourceArn}: 481 response",
        "POST /tags/{resourceArn}: 482 response",
        "GET /tags/{resourceArn}: 481 response",
        "GET /tags/{resourceArn}: 482 response",
        "GET /tags/{resourceArn}: 480 response",
        "DELETE /apps/{appId}/domains/{domainName}: 482 response",
        "DELETE /apps/{appId}/domains/{domainName}: 483 response",
        "DELETE /apps/{appId}/domains/{domainName}: 484 response",
        "DELETE /apps/{appId}/domains/{domainName}: 480 response",
        "DELETE /apps/{appId}/domains/{domainName}: 481 response",
        "GET /apps/{appId}/domains/{domainName}: 483 response",
        "GET /apps/{appId}/domains/{domainName}: 480 response",
        "GET /apps/{appId}/domains/{domainName}: 481 response",
        "GET /apps/{appId}/domains/{domainName}: 482 response",
        "POST /apps/{appId}/domains/{domainName}: 481 response",
        "POST /apps/{appId}/domains/{domainName}: 482 response",
        "POST /apps/{appId}/domains/{domainName}: 483 response",
        "POST /apps/{appId}/domains/{domainName}: 484 response",
        "POST /apps/{appId}/domains/{domainName}: 480 response",
        "DELETE /webhooks/{webhookId}: 483 response",
        "DELETE /webhooks/{webhookId}: 484 response",
        "DELETE /webhooks/{webhookId}: 480 response",
        "DELETE /webhooks/{webhookId}: 481 response",
        "DELETE /webhooks/{webhookId}: 482 response",
        "GET /webhooks/{webhookId}: 484 response",
        "GET /webhooks/{webhookId}: 480 response",
        "GET /webhooks/{webhookId}: 481 response",
        "GET /webhooks/{webhookId}: 482 response",
        "GET /webhooks/{webhookId}: 483 response",
        "POST /webhooks/{webhookId}: 484 response",
        "POST /webhooks/{webhookId}: 480 response",
        "POST /webhooks/{webhookId}: 481 response",
        "POST /webhooks/{webhookId}: 482 response",
        "POST /webhooks/{webhookId}: 483 response",
        "POST /apps/{appId}/webhooks: 482 response",
        "POST /apps/{appId}/webhooks: 483 response",
        "POST /apps/{appId}/webhooks: 484 response",
        "POST /apps/{appId}/webhooks: 485 response",
        "POST /apps/{appId}/webhooks: 480 response",
        "POST /apps/{appId}/webhooks: 481 response",
        "GET /apps/{appId}/webhooks: 481 response",
        "GET /apps/{appId}/webhooks: 482 response",
        "GET /apps/{appId}/webhooks: 483 response",
        "GET /apps/{appId}/webhooks: 480 response",
        "DELETE /tags/{resourceArn}#tagKeys: 480 response",
        "DELETE /tags/{resourceArn}#tagKeys: 481 response",
        "DELETE /tags/{resourceArn}#tagKeys: 482 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 481 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 482 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 483 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}/artifacts: 480 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 480 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 481 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 482 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 483 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}/stop: 484 response",
        "GET /apps: 480 response",
        "GET /apps: 481 response",
        "GET /apps: 482 response",
        "POST /apps: 481 response",
        "POST /apps: 482 response",
        "POST /apps: 483 response",
        "POST /apps: 484 response",
        "POST /apps: 480 response",
        "GET /apps/{appId}/branches/{branchName}: 480 response",
        "GET /apps/{appId}/branches/{branchName}: 481 response",
        "GET /apps/{appId}/branches/{branchName}: 482 response",
        "GET /apps/{appId}/branches/{branchName}: 483 response",
        "POST /apps/{appId}/branches/{branchName}: 481 response",
        "POST /apps/{appId}/branches/{branchName}: 482 response",
        "POST /apps/{appId}/branches/{branchName}: 483 response",
        "POST /apps/{appId}/branches/{branchName}: 484 response",
        "POST /apps/{appId}/branches/{branchName}: 480 response",
        "DELETE /apps/{appId}/branches/{branchName}: 481 response",
        "DELETE /apps/{appId}/branches/{branchName}: 482 response",
        "DELETE /apps/{appId}/branches/{branchName}: 483 response",
        "DELETE /apps/{appId}/branches/{branchName}: 484 response",
        "DELETE /apps/{appId}/branches/{branchName}: 480 response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 480 response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 481 response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 482 response",
        "POST /apps/{appId}/branches/{branchName}/deployments: 483 response",
        "DELETE /apps/{appId}: 481 response",
        "DELETE /apps/{appId}: 482 response",
        "DELETE /apps/{appId}: 483 response",
        "DELETE /apps/{appId}: 484 response",
        "DELETE /apps/{appId}: 480 response",
        "GET /apps/{appId}: 482 response",
        "GET /apps/{appId}: 483 response",
        "GET /apps/{appId}: 480 response",
        "GET /apps/{appId}: 481 response",
        "POST /apps/{appId}: 483 response",
        "POST /apps/{appId}: 480 response",
        "POST /apps/{appId}: 481 response",
        "POST /apps/{appId}: 482 response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 484 response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 480 response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 481 response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 482 response",
        "DELETE /apps/{appId}/backendenvironments/{environmentName}: 483 response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 483 response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 480 response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 481 response",
        "GET /apps/{appId}/backendenvironments/{environmentName}: 482 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 response",
        "DELETE /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 484 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 480 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 481 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 482 response",
        "GET /apps/{appId}/branches/{branchName}/jobs/{jobId}: 483 response",
        "POST /apps/{appId}/accesslogs: 480 response",
        "POST /apps/{appId}/accesslogs: 481 response",
        "POST /apps/{appId}/accesslogs: 482 response",
        "POST /apps/{appId}/accesslogs: 483 response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 484 response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 480 response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 481 response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 482 response",
        "POST /apps/{appId}/branches/{branchName}/deployments/start: 483 response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 480 response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 481 response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 482 response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 483 response",
        "POST /apps/{appId}/branches/{branchName}/jobs: 484 response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 480 response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 481 response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 482 response",
        "GET /apps/{appId}/branches/{branchName}/jobs: 483 response",
        "GET /apps/{appId}/domains: 480 response",
        "GET /apps/{appId}/domains: 481 response",
        "GET /apps/{appId}/domains: 482 response",
        "POST /apps/{appId}/domains: 483 response",
        "POST /apps/{appId}/domains: 484 response",
        "POST /apps/{appId}/domains: 485 response",
        "POST /apps/{appId}/domains: 480 response",
        "POST /apps/{appId}/domains: 481 response",
        "POST /apps/{appId}/domains: 482 response",
        "GET /artifacts/{artifactId}: 480 response",
        "GET /artifacts/{artifactId}: 481 response",
        "GET /artifacts/{artifactId}: 482 response",
        "GET /artifacts/{artifactId}: 483 response",
        "GET /artifacts/{artifactId}: 484 response",
        "GET /apps/{appId}/backendenvironments: 480 response",
        "GET /apps/{appId}/backendenvironments: 481 response",
        "GET /apps/{appId}/backendenvironments: 482 response",
        "POST /apps/{appId}/backendenvironments: 480 response",
        "POST /apps/{appId}/backendenvironments: 481 response",
        "POST /apps/{appId}/backendenvironments: 482 response",
        "POST /apps/{appId}/backendenvironments: 483 response",
        "POST /apps/{appId}/backendenvironments: 484 response",
        "GET /apps/{appId}/branches: 481 response",
        "GET /apps/{appId}/branches: 482 response",
        "GET /apps/{appId}/branches: 480 response",
        "POST /apps/{appId}/branches: 484 response",
        "POST /apps/{appId}/branches: 485 response",
        "POST /apps/{appId}/branches: 480 response",
        "POST /apps/{appId}/branches: 481 response",
        "POST /apps/{appId}/branches: 482 response",
        "POST /apps/{appId}/branches: 483 response"
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
- **Message:** Versioning validation failed: Migration guides are referenced: Info description does not reference migration or upgrade guides; Version follows semantic versioning: Version "2017-07-25" does not match semver format (expected MAJOR.MINOR.PATCH); Versioning strategy is documented: Info description does not mention versioning strategy; Breaking changes are documented: Info description does not reference breaking changes or a changelog; Version compatibility is specified: Info description does not mention version compatibility
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
      "Version follows semantic versioning": "Version \"2017-07-25\" does not match semver format (expected MAJOR.MINOR.PATCH)",
      "Versioning strategy is documented": "Info description does not mention versioning strategy"
    }
  }
```

**Suggested Fix:**
Update the info section with version details, deprecation notices, and migration references

---

