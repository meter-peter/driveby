# driveby

`driveby` is a modern API validation and testing framework built on the principle of **Documentation Driven Testing**. This approach treats your API documentation, specifically OpenAPI (Swagger) specifications, as the single source of truth for validating and testing your API.

## Concept and Purpose

The core concept behind `driveby` is to shift the focus of API quality assurance from manually written tests to automated validation based on comprehensive and accurate documentation. By using your OpenAPI specification, `driveby` can automatically verify that your API implementation adheres to its documented contract.

The primary purpose of `driveby` is to help development teams ensure their APIs are consistently reliable, performant, and well-documented. It aims to catch discrepancies between the documentation and the implementation early in the development lifecycle.

## Principles of a Testable API

To be effectively validated and tested by Driveby (or any documentation-driven testing tool), an API should adhere to the following principles:

1.  **Examplable:** Every request and response must include realistic and representative examples in the OpenAPI schema. Examples for parameters, request bodies, and responses are crucial for automated test generation and validation.
2.  **Describable:** All endpoints, parameters, and schemas must have clear and informative descriptions. Descriptions aid both human and machine understanding of the API's intended behavior, including error responses and edge cases.
3.  **Deterministic:** Given the same valid input under the same conditions, the API should consistently produce the same output. Determinism is fundamental for reliable and reproducible automated testing.
4.  **Observable:** The API should provide clear and discernible outputs for all operations. This includes surfacing error conditions, validation failures, and business logic outcomes in responses rather than hiding them.
5.  **Complete:** All possible responses, including success, various error states (validation, business logic), and edge cases, must be comprehensively documented in the OpenAPI specification.
6.  **Consistent:** Adhering to consistent naming conventions, status code usage, and error response formats across the API reduces ambiguity and enhances the effectiveness of automated validation.
7.  **Versioned:** The API and its corresponding documentation should be clearly versioned. This practice is essential for managing API evolution and ensuring backward compatibility while allowing for safe testing of different versions.
8.  **Discoverable:** The OpenAPI specification should be readily available, ideally at a standard, well-known endpoint (e.g., `/openapi.json`). Providing interactive documentation (e.g., Swagger UI at `/docs`) further aids human discoverability and understanding.

## Driveby's Operational Principles

`driveby` itself operates based on a set of core principles that guide its validation and testing processes. These principles, embedded within the framework, define what `driveby` checks for when analyzing an API against its documentation:

*   **Documentation as the Source of Truth:** Reinforcing the core concept, the OpenAPI specification is treated as the definitive source for validation criteria.
*   **Comprehensive Validation:** `driveby` aims to cover various aspects beyond just schema compliance, including functional correctness and performance characteristics derived from or implied by the documentation.
*   **Early Detection of Discrepancies:** The framework is designed to quickly identify divergences between the live API's behavior and its documented contract.
*   **Actionable Reporting:** Providing clear, detailed, and easy-to-understand reports on validation failures is crucial for efficient issue resolution.
*   **Support for Automation:** `driveby` is built to be easily integrated into CI/CD pipelines for continuous and automated quality assurance.

## Core Principles and Management within Driveby

`driveby`'s validation process is driven by a set of internal principles that represent key aspects of API quality and documentation adherence. These principles are implemented as specific checks within the framework, automatically performed against your live API based on its OpenAPI specification. The key internal principles include:

*   **P001: OpenAPI Specification Compliance:** Ensures the API's structure, paths, operations, and schemas strictly adhere to the OpenAPI 3.0 specification.
*   **P002: Response Time Performance:** Validates that API endpoints respond within acceptable time limits, often derived from or compared against configured performance targets.
*   **P003: Error Response Documentation:** Checks if the OpenAPI spec documents possible error responses (status codes and schemas) for all relevant endpoints.
*   **P004: Request Validation:** Verifies, where possible through documentation analysis and potential interaction, that the API properly validates incoming request parameters as defined in the documentation.
*   **P005: Authentication Requirements:** Confirms that the OpenAPI spec clearly specifies the authentication and authorization requirements for endpoints.
*   **P006: Endpoint Functional Testing:** Performs basic functional tests by sending requests to documented endpoints and verifying that the responses align with the documented status codes and potentially schemas.
*   **P007: API Performance Compliance:** Assesses the overall performance of the API, measuring metrics like P95 latency and error rates, and comparing them against configured benchmarks.
*   **P008: API Versioning:** Checks for the presence and format of versioning information within the API documentation (e.g., in the `info.version` field).

These internal principles are managed and applied by the `Validator` component. The `Validator` parses the OpenAPI specification and uses the information to execute the checks corresponding to each principle against the live API. The results are compiled into a `ValidationReport`, providing a clear assessment of the API's compliance with its documentation and the defined quality standards.

## General Idea

The general idea is simple: provide `driveby` with the endpoint of your running API and its OpenAPI specification. `driveby` will then:

1.  Fetch and parse the OpenAPI specification.
2.  Apply the core validation principles against the live API endpoint, using the documentation to guide the tests.
3.  Perform functional tests by making requests based on the documented paths, parameters, and expected responses.
4.  Execute performance checks to ensure endpoints meet documented or configured performance targets.
5.  Generate detailed reports highlighting any violations of the principles or discrepancies found, outputting results in formats like JSON and Markdown.

This process ensures that your API not only functions correctly but also accurately reflects its documentation, reinforcing the Documentation Driven Testing methodology.

## Functional Testing in Practice

`driveby` incorporates functional testing (aligned with Principle P006) by automatically verifying the reachability and expected responses of the endpoints defined in your OpenAPI specification. By parsing the documented paths and HTTP methods, `driveby` sends requests to your live API to ensure that the endpoints are active and return the documented status codes. This practical application helps confirm that your API's basic functionality aligns with its documentation, serving as a continuous integration check.

## Performance Testing in Practice

Addressing performance concerns (aligned with Principles P002 and P007), `driveby` utilizes your OpenAPI specification to identify API endpoints and generate load targets. It then performs performance tests to measure key metrics such as P95 latency and error rate using tools like `vegeta`. These measured metrics are compared against performance targets that you can configure, allowing you to ensure that your API not only functions correctly but also consistently meets the necessary performance benchmarks as expected or required for a production environment.

# Report Formats

## JSON Reports
All validation, functional, and performance results are output as JSON to stdout and saved in the report output directory (default: ./reports). Example:

```json
{
  "timestamp": "2025-05-31T16:43:30.975+03:00",
  "version": "1.0.0",
  "environment": "development",
  "principles": [ ... ],
  "summary": { ... }
}
```

## Markdown Reports
A human-readable Markdown report is also generated for each run. Example:

```markdown
# API Validation Report

Generated: 2025-05-31T16:43:30.975+03:00

## Summary
| Total Checks | Passed | Failed | Critical | Warnings | Info |
|--------------|--------|--------|----------|----------|------|
| 7            | 5      | 2      | 1        | 1        | 0    |

## Principle Results
| Principle | Status | Message |
|-----------|--------|---------|
| P001      | Passed | ...     |
| P003      | Failed | ...     |
```

## Configuring Output Directory

You can set the report output directory with the `--report-dir` flag:

```
driveby validate-only --report-dir ./my-reports
```

All reports will be saved in the specified directory. 