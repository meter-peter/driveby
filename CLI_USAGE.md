# DriveBy CLI Usage

DriveBy is now a pure CLI tool that follows standard CLI principles. All configuration is done through command-line flags - no environment variables are required or supported. DriveBy supports OpenAPI 3.0/3.1 specifications.

## Basic Usage

### Validation Only
```bash
driveby validate-only \
  --openapi /path/to/openapi.yaml \
  --host api.example.com \
  --protocol https \
  --port 443
```

### Functional Testing Only
```bash
driveby function-only \
  --openapi /path/to/openapi.yaml \
  --host api.example.com \
  --protocol https \
  --port 443
```

### Load Testing Only
```bash
driveby load-only \
  --openapi /path/to/openapi.yaml \
  --host api.example.com \
  --protocol https \
  --port 443 \
  --concurrent-users 50 \
  --test-duration 600 \
  --max-latency-p95 1000
```

## Required Flags

- `--openapi`: Path or URL to OpenAPI specification (required)
- `--host`: Host of the API to test (required)

## Optional Flags

### Global Flags
- `--api-url`: Base URL of the API to test (if not provided, will be constructed from protocol, host, and port)
- `--protocol`: Protocol to use (http or https, default: http)
- `--port`: Port to use (defaults to 8080 for http, 443 for https)
- `--environment`: Environment name (default: production)
- `--version`: API version being tested (default: 1.0.0)
- `--timeout`: Request timeout in seconds (default: 30)
- `--validation-mode`: Validation mode (strict, minimal, default: minimal)
- `--report-dir`: Report output directory (default: /tmp/driveby-reports)
- `--log-level`: Log level (debug, info, warn, error, fatal, default: info)

### Authentication Flags
- `--auth-token`: Authentication token (Bearer token)
- `--auth-token-type`: Token type (default: Bearer)
- `--auth-token-header`: Header name for token (default: Authorization)
- `--auth-username`: Username for basic authentication
- `--auth-password`: Password for basic authentication
- `--auth-api-key`: API key for authentication
- `--auth-api-key-header`: Header name for API key (default: X-API-Key)

### GitHub Integration Flags
- `--github-token`: GitHub token for PR commenting (GITHUB_TOKEN env var)
- `--github-owner`: GitHub repository owner
- `--github-repo`: GitHub repository name
- `--github-pr-number`: GitHub PR number for commenting
- `--github-comment`: Enable GitHub PR commenting

### Load Test Specific Flags
- `--max-latency-p95`: Maximum allowed P95 latency in milliseconds (default: 500)
- `--min-success-rate`: Minimum required success rate (0-1) (default: 0.99)
- `--concurrent-users`: Number of concurrent users for load testing (default: 10)
- `--test-duration`: Duration of load test in seconds (default: 300)

## Examples

### Simple HTTP API
```bash
driveby validate-only \
  --openapi ./swagger.yaml \
  --host localhost \
  --port 8080
```

### HTTPS API with Custom Configuration
```bash
driveby function-only \
  --openapi https://api.example.com/swagger.json \
  --host api.example.com \
  --protocol https \
  --port 443 \
  --timeout 60 \
  --validation-mode strict \
  --environment staging
```

### Public Example: Swagger Petstore

The public Swagger Petstore is used as a reference external API for DriveBy examples and internal test suites:

```bash
# Minimal OpenAPI validation against Petstore
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io \
  --protocol https \
  --port 443 \
  --validation-mode minimal

# Strict validation against Petstore
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io \
  --protocol https \
  --port 443 \
  --validation-mode strict

# Functional testing against Petstore
driveby function-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io \
  --protocol https \
  --port 443 \
  --timeout 30s
```

### Load Testing with High Concurrency
```bash
driveby load-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --protocol https \
  --concurrent-users 100 \
  --test-duration 900 \
  --max-latency-p95 2000 \
  --min-success-rate 0.95
```

### Authentication Examples

#### Bearer Token Authentication
```bash
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --auth-token "your-bearer-token" \
  --auth-token-type "Bearer" \
  --auth-token-header "Authorization"
```

#### API Key Authentication
```bash
driveby function-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --auth-api-key "your-api-key" \
  --auth-api-key-header "X-API-Key"
```

#### Basic Authentication
```bash
driveby load-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --auth-username "username" \
  --auth-password "password"
```

### GitHub Integration Examples

#### Comment on PR with Validation Results
```bash
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --github-comment \
  --github-token "your-github-token" \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

#### Using Environment Variable for GitHub Token
```bash
export GITHUB_TOKEN="your-github-token"
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --github-comment \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

## Exit Codes

- `0`: Success
- `1`: Tests ran but failed validation
- `2`: Error executing tests
- `3`: Invalid command line arguments

## Authentication Validation

DriveBy validates authentication configuration to ensure:

1. **Single Method**: Only one authentication method can be used at a time
2. **Required Fields**: All required fields for the chosen method are provided
3. **Valid Headers**: Default headers are used if not specified

### Validation Rules

- **Bearer Token**: Requires `--auth-token`, optional `--auth-token-type` and `--auth-token-header`
- **API Key**: Requires `--auth-api-key`, optional `--auth-api-key-header`
- **Basic Auth**: Requires both `--auth-username` and `--auth-password`

### Error Messages

The tool provides clear error messages for authentication issues:

```bash
# Multiple auth methods (will fail)
driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-token "token" \
  --auth-username "user" \
  --auth-password "pass"
# Error: only one authentication method can be specified (token, api-key, or username/password)

# Missing password for basic auth (will fail)
driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-username "user"
# Error: password is required when username is specified
```

## GitHub Integration Validation

When using GitHub integration, the following validation occurs:

1. **Required Fields**: `--github-owner`, `--github-repo`, and `--github-pr-number` are required
2. **Token Source**: GitHub token can be provided via `--github-token` flag or `GITHUB_TOKEN` environment variable
3. **PR Number**: Must be greater than 0

### Error Messages

```bash
# Missing required fields (will fail)
driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --github-comment
# Error: GitHub configuration error: GitHub owner is required

# Invalid PR number (will fail)
driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --github-comment \
  --github-owner "org" \
  --github-repo "repo" \
  --github-pr-number 0
# Error: GitHub PR number must be greater than 0
```

## Migration from Environment Variables

If you were previously using environment variables, replace them with the corresponding flags:

| Old Environment Variable | New Flag |
|-------------------------|----------|
| `DRIVEBY_OPENAPI` | `--openapi` |
| `DRIVEBY_HOST` | `--host` |
| `DRIVEBY_PROTOCOL` | `--protocol` |
| `DRIVEBY_PORT` | `--port` |
| `DRIVEBY_API_URL` | `--api-url` |
| `DRIVEBY_ENVIRONMENT` | `--environment` |
| `DRIVEBY_VERSION` | `--version` |
| `DRIVEBY_TIMEOUT` | `--timeout` |
| `DRIVEBY_VALIDATION_MODE` | `--validation-mode` |
| `DRIVEBY_REPORT_DIR` | `--report-dir` |
| `DRIVEBY_MAX_LATENCY_P95` | `--max-latency-p95` |
| `DRIVEBY_MIN_SUCCESS_RATE` | `--min-success-rate` |
| `DRIVEBY_CONCURRENT_USERS` | `--concurrent-users` |
| `DRIVEBY_TEST_DURATION` | `--test-duration` |

## Benefits of This Approach

1. **Explicit Configuration**: All settings are visible in the command line
2. **No Hidden Dependencies**: No reliance on environment variables
3. **Better Debugging**: Easy to see what configuration is being used
4. **Portable**: Commands work the same way across different environments
5. **Scriptable**: Easy to automate and version control
6. **Self-Documenting**: Help text and examples are built into the CLI
7. **Authentication Support**: Multiple auth methods with validation
8. **GitHub Integration**: Automatic PR commenting with validation results 

## Tools: Running a Small Public OpenAPI Batch (Max 20)

For smoke-testing DriveBy against a small, curated set of public OpenAPI specifications, this repository includes a helper script limited to **20 entries per run**:

- **Script path**: `tools/run-openapi-batch.sh`
- **Purpose**: run `driveby validate-only` in `minimal` mode against up to 20 OpenAPI URLs and collect a CSV summary.

### CSV Format

Create a CSV file with at most 20 rows (excluding the header):

```text
name,url,host
petstore,https://petstore3.swagger.io/api/v3/openapi.json,petstore3.swagger.io
github,https://example.com/github-openapi.json,api.github.com
```

The script enforces the **max 20 OpenAPIs** constraint and will fail fast if the CSV contains more than 20 non-empty data rows.

### Running the Batch Tool

```bash
tools/run-openapi-batch.sh openapis-sample.csv /tmp/driveby-openapi-batch
```

This will:

- Run `driveby validate-only --validation-mode minimal` once per row.
- Store individual reports under `/tmp/driveby-openapi-batch/<name>-<timestamp>/`.
- Write a summary CSV to `/tmp/driveby-openapi-batch/summary-<timestamp>.csv` containing:
  - `name,url,host,exit_code,mode,report_dir`.