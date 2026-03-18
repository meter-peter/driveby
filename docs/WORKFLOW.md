# Driveby API Validation in Kubernetes Workflows

This document describes how to use Driveby for API validation in Kubernetes/Argo workflows.

## Environment Variables

The following environment variables can be used to configure Driveby:

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `DRIVEBY_API_URL` | Base URL of the API to test | Yes | - |
| `DRIVEBY_OPENAPI_URL` | URL to the OpenAPI specification | Yes | - |
| `DRIVEBY_ENVIRONMENT` | Environment name (e.g., "dev", "staging") | No | "default" |
| `DRIVEBY_VERSION` | API version to validate | No | "1.0.0" |
| `DRIVEBY_TIMEOUT` | Request timeout duration | No | "30s" |
| `DRIVEBY_VALIDATION_MODE` | Validation mode ("minimal" or "strict") | No | "minimal" |
| `DRIVEBY_AUTH_TOKEN` | Authentication token for API access | No | - |
| `DRIVEBY_MAX_LATENCY_P95` | Maximum allowed 95th percentile latency | No | "500ms" |
| `DRIVEBY_MIN_SUCCESS_RATE` | Minimum required success rate (0-1) | No | "0.99" |
| `DRIVEBY_CONCURRENT_USERS` | Number of concurrent users for load testing | No | "10" |
| `DRIVEBY_TEST_DURATION` | Duration of load tests | No | "5m" |

## Example Argo Workflow

Here's an example Argo workflow that runs API validation:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: driveby-validation-
spec:
  entrypoint: validate-api
  templates:
  - name: validate-api
    inputs:
      parameters:
      - name: api-url
        value: "https://api.example.com"
      - name: openapi-url
        value: "https://api.example.com/openapi.json"
      - name: environment
        value: "staging"
      - name: auth-token
        valueFrom:
          secretKeyRef:
            name: api-credentials
            key: token
    container:
      image: your-driveby-image:latest
      command: ["driveby", "validate-only"]
      env:
        - name: DRIVEBY_API_URL
          value: "{{inputs.parameters.api-url}}"
        - name: DRIVEBY_OPENAPI_URL
          value: "{{inputs.parameters.openapi-url}}"
        - name: DRIVEBY_ENVIRONMENT
          value: "{{inputs.parameters.environment}}"
        - name: DRIVEBY_AUTH_TOKEN
          valueFrom:
            secretKeyRef:
              name: api-credentials
              key: token
        - name: DRIVEBY_VALIDATION_MODE
          value: "strict"
      volumeMounts:
        - name: reports
          mountPath: /tmp/driveby-reports
    volumes:
      - name: reports
        emptyDir: {}
    outputs:
      artifacts:
        - name: validation-report
          path: /tmp/driveby-reports/validation-report.json
```

## Validation Modes

### Minimal Mode
- Runs basic validation principles (P001, P004)
- Checks OpenAPI specification compliance
- Validates basic request schemas
- Suitable for quick validation in CI/CD pipelines

### Strict Mode
- Runs all validation principles
- Includes comprehensive documentation checks
- Validates error handling and authentication
- Performs detailed schema validation
- Suitable for thorough API reviews

## Performance Testing

For load testing, use the `load-only` command:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: driveby-load-test-
spec:
  entrypoint: load-test
  templates:
  - name: load-test
    container:
      image: your-driveby-image:latest
      command: ["driveby", "load-only"]
      env:
        - name: DRIVEBY_API_URL
          value: "https://api.example.com"
        - name: DRIVEBY_OPENAPI_URL
          value: "https://api.example.com/openapi.json"
        - name: DRIVEBY_MAX_LATENCY_P95
          value: "500ms"
        - name: DRIVEBY_MIN_SUCCESS_RATE
          value: "0.99"
        - name: DRIVEBY_CONCURRENT_USERS
          value: "10"
        - name: DRIVEBY_TEST_DURATION
          value: "5m"
      volumeMounts:
        - name: reports
          mountPath: /tmp/driveby-reports
    volumes:
      - name: reports
        emptyDir: {}
```

## Validation Report

The validation report is generated in JSON format and includes:
- Validation results for each principle
- Performance metrics (if load testing)
- Summary of passed/failed checks
- Detailed error messages and suggestions

Example report structure:
```json
{
  "version": "1.0.0",
  "environment": "staging",
  "timestamp": "2024-03-14T12:00:00Z",
  "principles": [
    {
      "id": "P001",
      "name": "OpenAPI Specification Compliance",
      "passed": true,
      "message": "OpenAPI specification is fully compliant"
    }
  ],
  "summary": {
    "total_checks": 8,
    "passed_checks": 7,
    "failed_checks": 1,
    "critical_issues": 0,
    "warnings": 1
  }
}
```

## Best Practices

1. **Secrets Management**
   - Store authentication tokens in Kubernetes secrets
   - Use Argo's secret management features
   - Never hardcode credentials in workflows

2. **Resource Management**
   - Set appropriate resource limits for the validation pod
   - Consider the impact of load testing on the target API
   - Use appropriate concurrent user counts

3. **Validation Strategy**
   - Use minimal mode for CI/CD pipelines
   - Run strict validation in staging environments
   - Schedule regular load tests during off-peak hours

4. **Report Handling**
   - Store validation reports as artifacts
   - Set up notifications for failed validations
   - Archive reports for historical analysis

## Troubleshooting

Common issues and solutions:

1. **Authentication Failures**
   - Verify the auth token is correctly set in secrets
   - Check token expiration
   - Ensure the token has necessary permissions

2. **Timeout Issues**
   - Increase `DRIVEBY_TIMEOUT` for slow APIs
   - Check network connectivity
   - Verify API availability

3. **Load Test Failures**
   - Adjust concurrent users if rate limited
   - Verify API can handle the load
   - Check resource limits on the validation pod

4. **OpenAPI Spec Issues**
   - Verify the spec URL is accessible
   - Check spec validity
   - Ensure spec version is supported (3.0.x or 3.1.0) 