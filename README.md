# DriveBy - Modern API Validation Framework

DriveBy is a comprehensive API validation framework that helps you validate, test, and monitor your APIs. It supports OpenAPI/Swagger specifications and provides extensive validation, testing, and monitoring capabilities.

## Features

- **OpenAPI Validation**: Validates API specifications against OpenAPI 3.0 standards
- **Functional Testing**: Tests API endpoints for functionality and correctness
- **Performance Testing**: Load tests APIs with configurable targets
- **Documentation Validation**: Ensures API documentation is complete and accurate
- **Auto-fixing**: Automatically fixes common documentation issues (TODO)
- **Comprehensive Reporting**: Generates detailed reports in JSON and Markdown formats
- **Authentication Support**: Supports various authentication methods
- **Configurable**: Highly configurable through YAML configuration
- **Validation Modes**: Supports different validation levels (minimal/strict) for different use cases

## Validation Modes

DriveBy supports two validation modes to accommodate different use cases:

1. **Minimal Mode** (default)
   - Focuses on essential validation only
   - Validates basic schema existence and structure
   - Only checks documentation for present error codes (doesn't require specific codes to be present)
   - Skips functional testing and performance testing
   - No load testing or functional testing reports are generated
   - Faster execution with fewer checks
   - Suitable for development and CI/CD pipelines
   - Validates:
     - OpenAPI Specification Compliance (P001)
     - Basic Request Schema Validation (P004)
     - Documentation for any present error responses

2. **Strict Mode**
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
- Config file: `validation.mode: strict|minimal`
- Environment variable: `DRIVEBY_VALIDATION_MODE=strict|minimal`

Example config.yaml:
```yaml
validation:
  mode: "minimal"  # or "strict"
  openapi_path: "openapi.json"
  environment: "development"
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

1. Create a configuration file (config.yaml):
```yaml
api:
  base_url: "http://localhost:8080"
  host: "localhost"
  port: "8080"
  base_path: "/api/v1"

validation:
  openapi_path: "openapi.json"
  environment: "development"
```

2. Run validation:
```bash
# Run all validations
driveby validate

# Run only documentation validation
driveby validate-only

# Run only functional tests
driveby function-only

# Run only performance tests
driveby load-only
```

## Configuration

DriveBy supports configuration via environment variables or command-line flags. The following variables are available:

- **DRIVEBY_OPENAPI** (or --openapi):  
  The path or URL (with protocol, e.g. "https://..." or "http://...") to your OpenAPI specification.  
  If the value starts with "http://" or "https://", DriveBy treats it as a remote URL and fetches the spec from the network; otherwise, it is treated as a local file path.  
  If the --openapi flag is not set, DriveBy falls back to the DRIVEBY_OPENAPI environment variable (and logs a debug message).

- **DRIVEBY_PROTOCOL** (or --protocol):  
  The protocol (e.g. "http" or "https") used to construct the API's base URL for testing.  
  (This is separate from the protocol in DRIVEBY_OPENAPI.)

- **DRIVEBY_HOST** (or --host):  
  The host (e.g. "api.example.com") used to construct the API's base URL.

- **DRIVEBY_PORT** (or --port):  
  The port (e.g. "443" or "8080") used to construct the API's base URL.

- **DRIVEBY_API_URL** (or --api-url):  
  (Optional) If provided, this overrides the constructed base URL (from protocol, host, and port).

- **DRIVEBY_ENVIRONMENT** (or --environment):  
  The environment name (e.g. "production", "staging") for reporting.

- **DRIVEBY_VERSION** (or --version):  
  The API version being tested.

- **DRIVEBY_TIMEOUT** (or --timeout):  
  Request timeout (in seconds).

- **DRIVEBY_VALIDATION_MODE** (or --validation-mode):  
  Validation mode (e.g. "strict", "minimal").

- **DRIVEBY_REPORT_DIR** (or --report-dir):  
  Directory where validation reports are saved.

- **DRIVEBY_MAX_LATENCY_P95** (or --max-latency-p95):  
  (Load test only) Maximum allowed P95 latency (in milliseconds).

- **DRIVEBY_MIN_SUCCESS_RATE** (or --min-success-rate):  
  (Load test only) Minimum required success rate (0–1).

- **DRIVEBY_CONCURRENT_USERS** (or --concurrent-users):  
  (Load test only) Number of concurrent users for load testing.

- **DRIVEBY_TEST_DURATION** (or --test-duration):  
  (Load test only) Duration (in seconds) of the load test.

---

## Workflow

### Sample Workflow Using Environment Variables

1. **Set your environment variables** (or use command-line flags) to configure DriveBy. For example:

   ```bash
   # Set the OpenAPI spec path or URL – if you intend to load a local file, omit the protocol (http:// or https://).
   # If you include a protocol (e.g. "https://..."), DriveBy treats it as a remote URL.
   export DRIVEBY_OPENAPI="https://docs.example.com/openapi.json"
   # (Alternatively, for a local file, you can omit the protocol, e.g. "/path/to/openapi.json".)

   # Set the API's base URL (used for testing endpoints) – these are separate from the OpenAPI spec URL.
   export DRIVEBY_PROTOCOL="https"
   export DRIVEBY_HOST="api.example.com"
   export DRIVEBY_PORT="443"

   # (Optional) Override the constructed base URL if needed.
   # export DRIVEBY_API_URL="https://api.example.com:443"

   # Set additional configuration (e.g. environment, version, timeout, etc.)
   export DRIVEBY_ENVIRONMENT="production"
   export DRIVEBY_VERSION="1.0.0"
   export DRIVEBY_TIMEOUT="30"
   export DRIVEBY_VALIDATION_MODE="minimal"
   export DRIVEBY_REPORT_DIR="/tmp/driveby-reports"
   ```

2. **Run DriveBy** (for example, to run only OpenAPI validation):

   ```bash
   ./driveby validate-only
   ```

   (You can also run "function-only" or "load-only" commands as needed.)

3. **Review the validation report** (saved in the report directory) to see if your API spec meets the validation criteria.

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