# DriveBy - Documentation-Driven Testing (DDT)

DriveBy is a thesis-ready API validation framework implementing **Documentation-Driven Testing (DDT)** — a methodology for automated API quality assurance in the GitOps era. It validates OpenAPI 3.x/Swagger 2.0 specifications and tests live API endpoints.

## Monorepo Structure

```
driveby-cli/     Go CLI tool — the core validation engine
apis/            Sample APIs for testing (perfect-api)
kubernetes/      Helm chart + raw YAML K8s deployment examples
samples/         Example configs, reports, demo workflows
thesis/          LaTeX thesis document
tools/           Python/bash utilities for batch testing
docs/            CLI usage guides and documentation
```

## Quick Start

```bash
make build                # Build the CLI
make up                   # Start docker-compose (perfect-api)
make validate             # Run validation against perfect-api
make test                 # Run unit tests
make integration-test     # Run integration tests (docker-compose)
```

## Features

- **OpenAPI Validation**: Validates API specifications against OpenAPI 3.0/3.1 standards
- **Functional Testing**: Tests API endpoints for functionality and correctness
- **Performance Testing**: Load tests APIs with configurable targets
- **Documentation Validation**: Ensures API documentation is complete and accurate
- **Auto-fixing**: Automatically fixes common documentation issues (TODO)
- **Comprehensive Reporting**: Generates detailed reports in JSON and Markdown formats
- **Authentication Support**: Supports various authentication methods (Bearer tokens, API keys, Basic auth)
- **GitHub Integration**: Automatically comments on pull requests with validation results
- **Configurable**: Highly configurable through YAML configuration
- **Validation Modes**: Supports different validation levels (minimal/strict) for different use cases

## Authentication Support

DriveBy supports multiple authentication methods for testing protected APIs:

### Bearer Token Authentication
```bash
driveby validate-only \
  --openapi openapi.json \
  --host api.example.com \
  --auth-token "your-bearer-token" \
  --auth-token-type "Bearer" \
  --auth-token-header "Authorization"
```

### API Key Authentication
```bash
driveby function-only \
  --openapi openapi.json \
  --host api.example.com \
  --auth-api-key "your-api-key" \
  --auth-api-key-header "X-API-Key"
```

### Basic Authentication
```bash
driveby load-only \
  --openapi openapi.json \
  --host api.example.com \
  --auth-username "username" \
  --auth-password "password"
```

**Note**: Only one authentication method can be used at a time. The tool will validate this and provide clear error messages if multiple methods are specified.

## GitHub Integration

DriveBy can automatically comment on GitHub pull requests with validation results. It supports both GitHub App authentication (recommended) and personal access tokens.

### GitHub App Authentication (Recommended)

For better security and granular permissions, use GitHub App authentication:

```bash
driveby validate-only \
  --openapi openapi.json \
  --host api.example.com \
  --github-comment \
  --github-app-id 123456 \
  --github-installation-id 789012 \
  --github-private-key "-----BEGIN RSA PRIVATE KEY-----\n..." \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

You can also provide the private key from a file:

```bash
driveby validate-only \
  --openapi openapi.json \
  --host api.example.com \
  --github-comment \
  --github-app-id 123456 \
  --github-installation-id 789012 \
  --github-private-key "$(cat /path/to/private-key.pem)" \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

### Personal Access Token Authentication (Legacy)

For backward compatibility, personal access tokens are still supported:

```bash
driveby validate-only \
  --openapi openapi.json \
  --host api.example.com \
  --github-comment \
  --github-token "your-github-token" \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

The GitHub token can also be provided via the `GITHUB_TOKEN` environment variable:

```bash
export GITHUB_TOKEN="your-github-token"
driveby validate-only \
  --openapi openapi.json \
  --host api.example.com \
  --github-comment \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number 123
```

## Validation Modes

DriveBy supports multiple validation modes to accommodate different use cases:

1. **Test-Only Mode** (new)
   - Skips all validation completely
   - Runs only functional and performance tests
   - Ideal for testing APIs without validation overhead
   - Fastest execution for pure testing scenarios
   - Suitable for CI/CD pipelines focused on testing
   - Command: `driveby test-only`

2. **Minimal Mode** (default)
   - Focuses on essential validation only
   - Validates basic OpenAPI structure and compliance
   - Skips functional testing and performance testing
   - Faster execution with minimal checks
   - Suitable for development and basic validation
   - Validates:
     - OpenAPI Specification Compliance (P001) - basic structure only

3. **Strict Mode**
   - Comprehensive validation of all aspects
   - Validates schema constraints, types, and formats
   - Requires complete documentation including all standard error codes
   - Enforces all validation principles
   - Generates full reports including load testing and functional testing
   - Suitable for production readiness checks
   - Validates all principles (P001-P008):
     - OpenAPI Specification Compliance
     - API Documentation Completeness
     - Error Response Documentation
     - Request Validation
     - Authentication Requirements
     - Endpoint Functional Testing
     - API Performance Compliance
     - API Versioning

You can set the validation mode through:
- Command line: `--validation-mode=strict|minimal`

Example usage:
```bash
driveby validate-only --validation-mode strict --openapi openapi.json --host localhost
```

Note: In minimal mode, the focus is on ensuring that any documented endpoints and responses are properly documented, rather than enforcing a complete set of documentation. This makes it ideal for development and test generation scenarios where you want to validate what's present without requiring comprehensive documentation.

## Installation

```bash
# Clone the repository
git clone https://github.com/meter-peter/driveby.git
cd driveby

# Build the project
go build -o driveby ./cmd/driveby

# Install globally (optional)
go install ./cmd/driveby
```

## Quick Start

Run validation using command-line flags:

```bash
# Run only documentation validation
driveby validate-only --openapi openapi.json --host localhost --port 8080

# Run only functional tests
driveby function-only --openapi openapi.json --host localhost --port 8080

# Run only performance tests
driveby load-only --openapi openapi.json --host localhost --port 8080

# Run functional and performance tests without validation (fastest)
driveby test-only --openapi openapi.json --host localhost --port 8080
```

## Configuration

DriveBy is configured entirely through command-line flags. No environment variables or configuration files are required.

### Required Flags
- **--openapi**: Path or URL to OpenAPI specification (required)
- **--host**: Host of the API to test (required)

### Optional Flags
- **--api-url**: Base URL of the API to test (if not provided, will be constructed from protocol, host, and port)
- **--protocol**: Protocol to use (http or https, default: http)
- **--port**: Port to use (defaults to 8080 for http, 443 for https)
- **--environment**: Environment name (default: production)
- **--version**: API version being tested (default: 1.0.0)
- **--timeout**: Request timeout in seconds (default: 30)
- **--validation-mode**: Validation mode (strict, minimal, default: minimal)
- **--report-dir**: Report output directory (default: /tmp/driveby-reports)
- **--log-level**: Log level (debug, info, warn, error, fatal, default: info)

### Authentication Flags
- **--auth-token**: Authentication token (Bearer token)
- **--auth-token-type**: Token type (default: Bearer)
- **--auth-token-header**: Header name for token (default: Authorization)
- **--auth-username**: Username for basic authentication
- **--auth-password**: Password for basic authentication
- **--auth-api-key**: API key for authentication
- **--auth-api-key-header**: Header name for API key (default: X-API-Key)

### GitHub Integration Flags
- **--github-token**: GitHub token for PR commenting (GITHUB_TOKEN env var) - Legacy
- **--github-owner**: GitHub repository owner
- **--github-repo**: GitHub repository name
- **--github-pr-number**: GitHub PR number for commenting
- **--github-comment**: Enable GitHub PR commenting

### GitHub App Authentication Flags (Recommended)
- **--github-app-id**: GitHub App ID
- **--github-installation-id**: GitHub App Installation ID
- **--github-private-key**: GitHub App private key (file path or PEM content)
- **--github-app-slug**: GitHub App slug

### Load Test Specific Flags
- **--max-latency-p95**: Maximum allowed P95 latency in milliseconds (default: 500)
- **--min-success-rate**: Minimum required success rate 0-1 (default: 0.99)
- **--concurrent-users**: Number of concurrent users for load testing (default: 10)
- **--test-duration**: Duration of load test in seconds (default: 300)

---

## Workflow

### Sample Workflow Using Command-Line Flags

1. **Run DriveBy with required flags** (for example, to run only OpenAPI validation):

   ```bash
   ./driveby validate-only \
     --openapi https://docs.example.com/openapi.json \
     --host api.example.com \
     --protocol https \
     --port 443 \
     --environment production \
     --version 1.0.0 \
     --timeout 30 \
     --validation-mode minimal \
     --auth-token "your-bearer-token"
   ```

   (You can also run "function-only" or "load-only" commands as needed.)

2. **Review the validation report** (saved in the report directory) to see if your API spec meets the validation criteria.

### Example Scripts

For repeated use, you can create shell scripts:

```bash
#!/bin/bash
# validate-api.sh
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host localhost \
  --port 8080 \
  --environment development \
  --auth-token "your-token"
```

```bash
#!/bin/bash
# load-test.sh
driveby load-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --protocol https \
  --concurrent-users 50 \
  --test-duration 600 \
  --max-latency-p95 1000 \
  --auth-api-key "your-api-key"
```

```bash
#!/bin/bash
# test-only.sh
driveby test-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --protocol https \
  --concurrent-users 10 \
  --test-duration 60 \
  --max-latency-p95 500 \
  --auth-api-key "your-api-key"
```

### GitHub Integration Example

```bash
#!/bin/bash
# validate-and-comment.sh (with GitHub App)
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --github-comment \
  --github-app-id "$GITHUB_APP_ID" \
  --github-installation-id "$GITHUB_INSTALLATION_ID" \
  --github-private-key "$GITHUB_PRIVATE_KEY" \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number "$PR_NUMBER"
```

```bash
#!/bin/bash
# validate-and-comment.sh (with legacy token)
driveby validate-only \
  --openapi ./api-spec.yaml \
  --host api.example.com \
  --github-comment \
  --github-token "$GITHUB_TOKEN" \
  --github-owner "your-org" \
  --github-repo "your-repo" \
  --github-pr-number "$PR_NUMBER"
```

---

## Validation Principles

DriveBy implements several validation principles (P001-P008):

1. **P001**: OpenAPI Specification Compliance
2. **P002**: Response Time Performance
3. **P003**: Error Response Documentation
4. **P004**: Request Validation
5. **P005**: Authentication Requirements
6. **P006**: Endpoint Functional Testing
7. **P007**: API Performance Compliance
8. **P008**: API Versioning

## Reports

DriveBy generates detailed reports in both JSON and Markdown formats, including:

- Validation results for each principle
- Performance metrics
- Documentation quality scores
- Auto-fix attempts and results
- Summary statistics

Reports are saved in the configured output directory (default: `./reports`).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 