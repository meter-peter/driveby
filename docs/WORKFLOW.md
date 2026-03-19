# DriveBy End-to-End Workflow

This document describes how to use DriveBy for API validation -- from local development through CI/CD pipelines and Kubernetes deployments.

## Workflow Overview

DriveBy supports three input modes:
1. **Local file** -- validate a spec from disk (`--openapi ./openapi.json`)
2. **Remote URL** -- validate a spec hosted anywhere on the internet (`--openapi https://...`)
3. **Live API** -- combine spec validation with runtime testing against a live endpoint

```
┌─────────────┐     ┌──────────────┐     ┌───────────────┐     ┌────────────┐
│ OpenAPI Spec │────>│ DriveBy CLI  │────>│ Principle     │────>│ JSON + MD  │
│ (file / URL) │     │ (loader)     │     │ Checkers      │     │ Reports    │
└─────────────┘     └──────────────┘     │ P001-P008     │     └────────────┘
                                          └───────────────┘
```

## Local Development

### Validate a local spec (perfect-api)
```bash
make up                   # Start docker-compose (perfect-api)
make validate             # Run validation against perfect-api
```

### Validate a public API spec by URL
```bash
# Swagger Petstore -- the canonical OpenAPI example
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io \
  --protocol https \
  --port 443 \
  --validation-mode strict
```

DriveBy auto-detects whether `--openapi` points to a file or URL and fetches accordingly. Both OpenAPI 3.x and Swagger 2.0 specs are supported.

### Batch validation against public APIs
```bash
# Create a CSV (max 20 entries)
cat > openapis.csv <<EOF
name,url,host
petstore,https://petstore3.swagger.io/api/v3/openapi.json,petstore3.swagger.io
EOF

# Run batch validation
tools/run-openapi-batch.sh openapis.csv /tmp/driveby-batch
```

See `tools/CLAUDE.md` for the full harvest-probe-validate workflow against APIs.guru.

## Validation Modes

### Minimal Mode (default)
- Runs P001 (OpenAPI Compliance) only
- Fast CI gate (< 5 seconds)
- Suitable for first-time evaluation of unknown APIs

### Strict Mode
- Runs P001-P005, P008 (all static analysis principles)
- Comprehensive documentation, schema, error handling, security, and versioning checks
- Suitable for pre-release quality gates and thesis evaluation

### Test-Only Mode
- Skips static validation entirely
- Runs runtime tests only (P006 functional, P007 performance when implemented)
- Fastest execution for pure endpoint testing

### Flexible Mode
- Same as minimal (P001 only) but allows tests to proceed even with some validation failures

## Kubernetes / Argo Workflow Integration

DriveBy is configured entirely through CLI flags. In Kubernetes, pass flags via container args:

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
      - name: openapi-url
        value: "https://api.example.com/openapi.json"
      - name: host
        value: "api.example.com"
    container:
      image: ghcr.io/meter-peter/driveby:latest
      command: ["driveby", "validate-only"]
      args:
        - "--openapi={{inputs.parameters.openapi-url}}"
        - "--host={{inputs.parameters.host}}"
        - "--protocol=https"
        - "--port=443"
        - "--validation-mode=strict"
        - "--report-dir=/tmp/driveby-reports"
      env:
        - name: DRIVEBY_AUTH_TOKEN
          valueFrom:
            secretKeyRef:
              name: api-credentials
              key: token
      volumeMounts:
        - name: reports
          mountPath: /tmp/driveby-reports
    volumes:
      - name: reports
        emptyDir: {}
    outputs:
      artifacts:
        - name: validation-report
          path: /tmp/driveby-reports/validation-report-latest.json
```

### Load Testing Workflow

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
      image: ghcr.io/meter-peter/driveby:latest
      command: ["driveby", "load-only"]
      args:
        - "--openapi=https://api.example.com/openapi.json"
        - "--host=api.example.com"
        - "--protocol=https"
        - "--port=443"
        - "--max-latency-p95=500ms"
        - "--min-success-rate=0.99"
        - "--concurrent-users=10"
        - "--test-duration=5m"
        - "--report-dir=/tmp/driveby-reports"
      volumeMounts:
        - name: reports
          mountPath: /tmp/driveby-reports
    volumes:
      - name: reports
        emptyDir: {}
```

## Automated GitOps Pipeline (Argo Events)

The production deployment uses Argo Events to automatically trigger validation when a PR is opened on the `perfect-api` repository. This is the primary workflow for the thesis demonstration.

```
PR on perfect-api  ──>  GitHub webhook  ──>  EventSource  ──>  Sensor  ──>  Workflow
                                                                              |
                                                   ┌──────────────────────────┘
                                                   v
                                          set-pending-status
                                                   |
                                               db-sync
                                                   |
                                           spec-sync-check
                                                   |
                                          validate-staging
                                                   |
                                          functional-test
                                                   |
                                     report-success + comment-pr
```

### How it works

1. A developer opens (or updates) a PR on `meter-peter/perfect-api`
2. GitHub sends a `pull_request` webhook to `https://driveby-webhook.private.novelcore.org/github/driveby`
3. The Argo Events **EventSource** receives it, publishes to the JetStream **EventBus**
4. The **Sensor** filters for `opened`/`reopened`/`synchronize` actions
5. The Sensor submits a `driveby-staging-promotion` Argo **Workflow** with PR metadata
6. The workflow DAG runs: pending status, DB sync, spec check, validation, functional test
7. On success: commit status set to `success`, validation report posted as PR comment
8. On failure: commit status set to `failure`, exit handler reports the error

### Verify the pipeline

```bash
# Check Argo Events resources
kubectl get eventbus,eventsource,sensor -n driveby

# Check pods are running
kubectl get pods -n driveby | grep -E 'eventbus|eventsource|sensor'

# Check webhook exists
gh api repos/meter-peter/perfect-api/hooks

# Watch for triggered workflows
kubectl get workflows -n driveby --watch
```

For full details, see [gitops-pipeline.md](gitops-pipeline.md).

## Validation Report

Reports are generated in both JSON and Markdown formats. Output directory defaults to `/tmp/driveby-reports/`.

Files generated per run:
- `validation-report-<timestamp>.json` -- full structured results
- `validation-report-<timestamp>.md` -- human-readable Markdown
- `validation-report-latest.json` / `.md` -- symlinks to most recent

Example report structure:
```json
{
  "status": "failed",
  "exit_code": 1,
  "version": "1.0.0",
  "environment": "production",
  "timestamp": "2026-03-19T00:53:08+02:00",
  "principles": [
    {
      "Principle": { "id": "P001", "name": "OpenAPI Specification Compliance", ... },
      "Passed": true,
      "Message": "OpenAPI specification is fully compliant with 3.0/3.1 standards",
      "Details": { "checks": { ... } }
    }
  ]
}
```

## Public API Validation Reference

DriveBy can validate any publicly available OpenAPI specification. This is useful for:
- Benchmarking DDT principles against real-world APIs
- Thesis evaluation (Chapter 6: Wild Evaluation)
- Demonstrating the gap between spec compliance and documentation quality

### Petstore API Results (March 2026)

The Swagger Petstore v3 was validated in strict mode as a reference data point:

| Principle | Result | Key Findings |
|-----------|--------|--------------|
| P001 OpenAPI Compliance | PASSED | Fully compliant OpenAPI 3.0 |
| P002 Documentation Quality | FAILED | 44 bodies missing examples, 6 schemas undescribed |
| P003 Error Handling | FAILED | Zero 5xx responses, 4 endpoints missing 4xx |
| P004 Schema Definitions | FAILED | No string length constraints, no numeric min/max |
| P005 Security Standards | FAILED | 10 endpoints unsecured, no global security |
| P008 Versioning Strategy | FAILED | No versioning strategy or changelog documented |

**Result: 1/6 passed** -- even the canonical OpenAPI reference API fails strict DDT validation.

## Best Practices

1. **Secrets Management**
   - Store authentication tokens in Kubernetes secrets
   - Use `--auth-token`, `--auth-api-key`, or `--auth-username`/`--auth-password` flags
   - Never hardcode credentials in workflows

2. **Resource Management**
   - Set appropriate resource limits for the validation pod
   - Consider the impact of load testing on the target API
   - Use appropriate concurrent user counts

3. **Validation Strategy**
   - Use minimal mode for CI/CD gates (fast, P001 only)
   - Run strict mode in staging for full quality assessment
   - Use batch tools for evaluating third-party API specs before integration
   - Schedule regular load tests during off-peak hours

4. **Report Handling**
   - Store validation reports as artifacts
   - Set up notifications for failed validations
   - Archive reports for historical analysis

## Troubleshooting

1. **Remote spec fetch fails**
   - Verify the spec URL is accessible (try `curl <url>`)
   - Check for redirects -- DriveBy follows standard HTTP redirects
   - Ensure the response is valid JSON (YAML support via URL is not yet available)

2. **Authentication Failures**
   - Verify the auth token via `--auth-token` flag
   - Check token expiration
   - Ensure the token has necessary permissions

3. **Timeout Issues**
   - Increase `--timeout` for slow APIs or large specs
   - Check network connectivity
   - Verify API availability

4. **Load Test Failures**
   - Adjust `--concurrent-users` if rate limited
   - Verify API can handle the load
   - Check resource limits on the validation pod

5. **OpenAPI Spec Issues**
   - Ensure spec version is supported (OpenAPI 3.0.x, 3.1.0, or Swagger 2.0)
   - Check spec validity with `driveby validate-only --validation-mode minimal`
   - Use `--log-level debug` for detailed parsing diagnostics