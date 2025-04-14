# DriveBy API: Testing and Validation Platform

DriveBy is a Kubernetes-native API for managing testing of APIs in your organization. It provides three core testing capabilities:

1. **Documentation Validation** - Validate OpenAPI documentation for compliance and completeness
2. **Load Testing** - Perform load tests against your services with detailed reporting
3. **Acceptance Testing** - Run comprehensive acceptance test suites against your APIs

All test failures can automatically create GitHub issues in the appropriate repositories, providing a seamless feedback loop for your development workflow.

## Architecture

DriveBy follows a modular, microservices-oriented architecture with clean separation of concerns:

![Architecture Diagram](architecture.png)

### Core Components

- **API Server**: RESTful API built with Gin framework
- **Queue System**: Redis-based asynchronous task processing
- **Storage**: Minio (S3-compatible) storage for test results and reports
- **Service Layer**: Modular services for different testing capabilities
- **GitHub Integration**: Automatic issue creation for test failures

### Design Principles

The system is built using the following design principles:

- **Domain-Driven Design**: Core business logic is organized around domain models
- **Clean Architecture**: Dependency inversion and clear boundaries between layers
- **Interface-Based Design**: All services are defined by interfaces for testability
- **Modular Components**: Each testing capability is a separate module
- **Kubernetes-Native**: Designed to run in Kubernetes from the ground up

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Kubernetes cluster (for production deployment)
- Go 1.19+ (for local development)
- Redis instance
- Minio instance (or S3 bucket)
- GitHub token (for issue creation)

### Configuration

DriveBy is configured through environment variables and/or a configuration file. See [Configuration](#configuration) section for details.

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/driveby
   cd driveby
   ```

2. Run with Docker Compose:
   ```bash
   docker-compose up
   ```

3. Or build and run locally:
   ```bash
   go build -o driveby ./cmd/driveby
   ./driveby
   ```

### Kubernetes Deployment

1. Apply the Kubernetes manifests:
   ```bash
   kubectl apply -f k8s/deployment.yaml
   ```

2. Update the GitHub token secret:
   ```bash
   kubectl create secret generic driveby-github \
     --from-literal=token=your-github-token
   ```

3. Access the API through the defined ingress.

## Testing Capabilities

### API Documentation Validation

Validates OpenAPI specifications for compliance and documentation quality:

```http
POST /api/v1/validation
Content-Type: application/json

{
  "name": "User Service API Validation",
  "description": "Validation of User Service OpenAPI documentation",
  "openapi_url": "https://api.example.com/swagger/doc.json",
  "compliance_threshold": 95.0,
  "create_github_issue": true,
  "github_repo": {
    "owner": "your-org",
    "repository": "user-service"
  }
}
```

### Load Testing

Performs load testing against your services:

```http
POST /api/v1/loadtest
Content-Type: application/json

{
  "name": "User Service Load Test",
  "description": "Load test for User Service API",
  "target_url": "https://api.example.com",
  "request_rate": 50,
  "duration": 300,
  "timeout": 5,
  "method": "GET",
  "create_github_issue": true,
  "github_repo": {
    "owner": "your-org",
    "repository": "user-service"
  }
}
```

### Acceptance Testing

Runs acceptance test suites against your APIs:

```http
POST /api/v1/acceptance
Content-Type: application/json

{
  "name": "User Service Acceptance Tests",
  "description": "Acceptance tests for User Service API",
  "base_url": "https://api.example.com",
  "test_cases": [
    {
      "name": "Get User Profile",
      "description": "Verify user profile retrieval",
      "path": "/users/123",
      "method": "GET",
      "assertions": [
        {
          "type": "status",
          "target": "status_code",
          "value": 200,
          "command": "eq"
        },
        {
          "type": "json",
          "target": "$.name",
          "value": "John Doe",
          "command": "eq"
        }
      ]
    }
  ],
  "create_github_issue": true,
  "github_repo": {
    "owner": "your-org",
    "repository": "user-service"
  }
}
```

## API Reference

Full API documentation is available at `/api/v1/docs/index.html` when the server is running.

## Configuration

Configuration can be provided via environment variables, config file, or both.

### Environment Variables

All environment variables are prefixed with `DRIVEBY_`. Examples:

```
DRIVEBY_SERVER_PORT=8080
DRIVEBY_REDIS_HOST=localhost
DRIVEBY_MINIO_ENDPOINT=localhost:9000
DRIVEBY_GITHUB_TOKEN=your-github-token
```

### Configuration File

The configuration file uses YAML format. Example:

```yaml
server:
  host: 0.0.0.0
  port: 8080
  mode: release
  timeout: 30s
  shutdown_timeout: 10s

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
  enabled: true

minio:
  endpoint: localhost:9000
  access_key_id: minioadmin
  secret_access_key: minioadmin
  use_ssl: false
  bucket_name: driveby
  region: us-east-1
  enabled: true

github:
  api_base_url: https://api.github.com
  token: your-github-token

logging:
  level: info
  format: json

testing:
  validation:
    compliance_threshold: 95.0
    fail_on_validation: true
  load_test:
    default_rps: 10
    default_duration: 30s
    default_timeout: 5s
  acceptance:
    default_timeout: 30s

features:
  enable_validation: true
  enable_load_test: true
  enable_acceptance: true
  enable_github: true
  enable_workers: true
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
