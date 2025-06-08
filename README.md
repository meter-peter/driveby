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

DriveBy is configured through a YAML file. See `config.yaml` for all available options:

- **API Configuration**: Base URL, host, port, and base path
- **Validation Configuration**: OpenAPI spec path, environment, version, etc.
- **Performance Configuration**: Latency targets, success rates, test duration
- **Authentication**: Token-based authentication settings
- **Logging**: Log level, format, and output settings
- **Reporting**: Report output directory and formats

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