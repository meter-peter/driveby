# Driveby API Testing Framework

A documentation-driven API testing framework that validates OpenAPI documentation, runs integration and load tests, and enforces quality gates.

## Principles of a Testable API

To be effectively validated and tested by Driveby (or any documentation-driven testing tool), an API should adhere to the following principles:

### 1. **Examplable**
- **Every request and response must include examples** in the OpenAPI schema.
- Examples should be realistic and cover typical as well as edge cases.
- Parameter, request body, and response examples are essential for automated test generation and validation.

### 2. **Describable**
- **All endpoints, parameters, and schemas must have clear descriptions.**
- Descriptions help both humans and machines understand the intent and usage of each part of the API.
- Error responses and edge cases should be described, not just success paths.

### 3. **Deterministic**
- **Given the same input, the API should produce the same output.**
- Determinism is crucial for reliable automated testing and reproducibility.

### 4. **Observable**
- **The API should provide clear, observable outputs for all operations.**
- Error conditions, validation failures, and business logic errors should be surfaced in responses, not hidden.

### 5. **Complete**
- **All possible responses (success, error, validation, etc.) must be documented.**
- Status codes, error structures, and edge cases should be included in the OpenAPI spec.

### 6. **Consistent**
- **Naming, status codes, and error formats should be consistent across the API.**
- Consistency reduces ambiguity and makes automated validation more effective.

### 7. **Versioned**
- **The API and its documentation should be versioned.**
- This allows for safe evolution and backward compatibility.

### 8. **Discoverable**
- **The OpenAPI spec should be available at a standard endpoint (e.g., `/openapi.json`).**
- Interactive documentation (e.g., Swagger UI at `/docs`) should be provided for human users.

---

## API Documentation

### Driveby Testing API Documentation
- Interactive Documentation (Swagger UI): `http://localhost:8081/api/v1/docs`
- OpenAPI Specification: `http://localhost:8081/api/v1/openapi.json`

### Example API Documentation
- Interactive Documentation (Swagger UI): `http://localhost:8082/docs`
- OpenAPI Specification: `http://localhost:8082/openapi.json`

## API Endpoints

### Testing Endpoints

#### POST /api/v1/tests
Executes a complete test suite including documentation validation, integration tests, and load tests.

Example Request:
```json
{
    "openapi_spec": "http://localhost:8082/openapi.json",
    "thresholds": {
        "documentation": {
            "threshold": 0.95
        },
        "load_test": {
            "success_rate": 0.99,
            "max_latency": "500ms"
        }
  }
}
```

Example Response:
```json
{
    "test_id": "test_123456",
    "timestamp": "2024-03-20T10:00:00Z",
    "results": {
        "documentation": {
            "compliance_score": 0.98,
            "missing_examples": 1,
            "undocumented_endpoints": ["/api/v1/health"],
            "errors": []
        },
        "integration": {
            "total_tests": 25,
            "passed": 24,
            "failed": 1,
            "failed_endpoints": [
                {
                    "endpoint": "/api/v1/users",
                    "error": "Response schema mismatch"
                }
            ]
        },
        "load_test": {
            "total_requests": 1000,
            "success_rate": 0.995,
            "latency_p95": "450ms",
            "status_codes": {
                "200": 995,
                "500": 5
            },
            "errors": []
        }
  }
}
```

#### POST /api/v1/validation
Runs validation tests against the OpenAPI specification.

Example Request:
```json
{
    "openapi_spec": "http://localhost:8082/openapi.json"
}
```

Example Response:
```json
{
    "compliance_score": 0.98,
    "missing_examples": 1,
    "undocumented_endpoints": ["/api/v1/health"],
    "errors": [
        "Missing response example for POST /api/v1/users"
    ]
}
```

#### POST /api/v1/loadtest
Conducts load tests with configurable parameters.

Example Request:
```json
{
    "openapi_spec": "http://localhost:8082/openapi.json",
    "request_rate": 100,
    "duration": "30s"
}
```

Example Response:
```json
{
    "total_requests": 3000,
    "success_rate": 0.995,
    "latency_p95": "450ms",
    "status_codes": {
        "200": 2985,
        "500": 15
    },
    "errors": [
        "Connection timeout on POST /api/v1/users"
    ]
}
```

#### POST /api/v1/acceptance
Runs acceptance tests to validate business requirements.

Example Request:
```json
{
    "openapi_spec": "http://localhost:8082/openapi.json"
}
```

Example Response:
```json
{
    "passed": true,
    "details": "All business requirements met. User creation, authentication, and data retrieval workflows validated successfully."
}
```

## Testing Methods

### Documentation Validation
The `validateDocumentation` method performs comprehensive validation of API documentation:

#### Validation Checks
- **Response Documentation**
  - Presence of response documentation for each endpoint
  - Completeness of response descriptions
  - Example presence for all responses
  - Error response documentation (4xx and 5xx status codes)

- **Request Documentation**
  - Request body examples
  - Parameter examples
  - Query parameter documentation
  - Path parameter documentation

#### Compliance Metrics
- **Compliance Score**
  - Calculated as: (Compliant Endpoints / Total Endpoints) * 100
  - Minimum threshold: 95%

- **Documentation Gaps**
  - Missing examples count
  - Undocumented endpoints list
  - Incomplete response documentation
  - Missing error documentation

### Integration Testing
The `runIntegrationTests` method performs automated integration testing:

#### Test Discovery
- Discovers testable endpoints from OpenAPI spec
- Extracts examples from parameters
- Gets request body examples
- Validates endpoint implementation against specification

#### Test Tracking
- **Test Metrics**
  - Total tests run
  - Passed tests count
  - Failed tests count
  - Failed endpoints with error messages

#### Success Criteria
- Minimum pass rate: 95%
- All critical endpoints must pass
- No schema mismatches
- All examples must validate

### Load Testing
The `runLoadTests` method performs performance testing using Vegeta:

#### Performance Metrics
- **Success Rate**
  - Minimum threshold: 99%
  - Tracks failed requests
  - Monitors error distribution

- **Latency**
  - P95 latency threshold: 500ms
  - Tracks response time distribution
  - Monitors slow endpoints

- **Error Rate**
  - Maximum threshold: 1%
  - Tracks status code distribution
  - Monitors error patterns

#### Test Configuration
- Configurable request rate
- Adjustable test duration
- Customizable endpoints
- Threshold configuration

## Quality Standards

### Documentation
- Complete OpenAPI documentation
- All endpoints documented
- Request/response examples
- Error documentation
- Parameter descriptions

### Tests
- 95% test coverage
- Integration test suite
- Load test suite
- Acceptance test suite
- Automated test execution

### API Implementation
- Response time < 500ms
- 99% success rate
- Error handling
- Input validation
- Security measures

## Getting Started

1. Clone the repository
2. Install dependencies
3. Configure the API endpoints
4. Run the test suite

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
