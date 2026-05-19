## 🔴 Staging Gate — bad-docs-api

**2/6 principles passed (33%)** | Mode: `strict` | Checks: `validate-only` | Version: `1.0.0` | Env: `production`

### Summary

| Severity | Passed | Failed |
|----------|--------|--------|
| Critical | 2 | 3 |
| Warning | 0 | 1 |

### Gate Details

| Property | Value |
|----------|-------|
| Environment | `staging` |
| Checks | `validate-only` |
| Validation Mode | `strict` |
| Workflow | [View Run](https://workflow.private.novelcore.org/workflows/driveby/bad-docs-api-staging-gate-pipeline-k2qqt) |

### How to Pass This Gate

**Critical blockers** (must fix to unblock promotion):

- 🚨 **P002: API Documentation Quality** — Add missing documentation including descriptions, examples, and operation details
- 🚨 **P003: Error Handling Standards** — Add comprehensive error response documentation including codes, messages, and consistent error schemas
- 🚨 **P004: Request Schema Definitions** — Add comprehensive schema validation including data types, constraints, and required fields

**Warnings** (do not block promotion, but should be addressed):

- ⚠️ **P008: API Versioning Strategy**

### Passed Principles

- ✅ **P001: OpenAPI Specification Compliance** — OpenAPI specification is fully compliant with 3.0/3.1 standards
- ✅ **P005: Security Standards** — All security requirements are properly defined and consistent

### Failed Principles

<details><summary>🚨 <strong>P002: API Documentation Quality</strong> (critical)</summary>

**Result:** Documentation quality issues found: All request/response bodies have examples: GET /products: 422 application/json response, GET /products: 200 application/json response, POST /products: request body, POST /products: application/json request body, POST /products: 201 application/json response, POST /products: 422 application/json response, GET /products/{product_id}: 200 application/json response, GET /products/{product_id}: 422 application/json response, PUT /products/{product_id}: request body, PUT /products/{product_id}: application/json request body, PUT /products/{product_id}: 200 application/json response, PUT /products/{product_id}: 422 application/json response, DELETE /products/{product_id}: 422 application/json response, POST /tasks: request body, POST /tasks: application/json request body, POST /tasks: 201 application/json response, POST /tasks: 422 application/json response, POST /test/echo: request body, POST /test/echo: application/json request body, POST /test/echo: 200 application/json response, POST /test/echo: 422 application/json response, GET /test/health: 200 application/json response, GET /legacy/products: 200 application/json response; All schemas have descriptions: HTTPValidationError, ProductBase, Product, ProductCreate, ValidationError, ProductCategory, TaskCreate, Task; All enums have descriptions: ProductCategory: enum value electronics, ProductCategory: enum value clothing, ProductCategory: enum value food, ProductCategory: enum value books, ProductCategory: enum value other; All operations have detailed descriptions: GET /products, POST /products, GET /products/{product_id}, PUT /products/{product_id}, DELETE /products/{product_id}, POST /tasks, POST /test/echo, GET /test/health, GET /legacy/products; All parameters have descriptions: GET /products: parameter category, GET /products: parameter min_price, GET /products: parameter max_price, GET /products: parameter in_stock, GET /products/{product_id}: parameter product_id, PUT /products/{product_id}: parameter product_id, DELETE /products/{product_id}: parameter product_id

**Failed checks:**

- ❌ All enums have descriptions
- ❌ All operations have detailed descriptions
- ❌ All parameters have descriptions
- ❌ All request/response bodies have examples
- ❌ All schemas have descriptions
- ❌ Contact information is provided — Contact information is missing
- ❌ License information is provided — License information is missing

<details><summary>Passed checks (1)</summary>

- ✅ API has a general description

</details>

**Missing docs:**

_All enums have descriptions:_
  - ProductCategory: enum value electronics
  - ProductCategory: enum value clothing
  - ProductCategory: enum value food
  - ProductCategory: enum value books
  - ProductCategory: enum value other
_All operations have detailed descriptions:_
  - GET /products
  - POST /products
  - GET /products/{product_id}
  - PUT /products/{product_id}
  - DELETE /products/{product_id}
  - POST /tasks
  - POST /test/echo
  - GET /test/health
  - GET /legacy/products
_All parameters have descriptions:_
  - GET /products: parameter category
  - GET /products: parameter min_price
  - GET /products: parameter max_price
  - GET /products: parameter in_stock
  - GET /products/{product_id}: parameter product_id
  - PUT /products/{product_id}: parameter product_id
  - DELETE /products/{product_id}: parameter product_id
_All request/response bodies have examples:_
  - GET /products: 422 application/json response
  - GET /products: 200 application/json response
  - POST /products: request body
  - POST /products: application/json request body
  - POST /products: 201 application/json response
  - POST /products: 422 application/json response
  - GET /products/{product_id}: 200 application/json response
  - GET /products/{product_id}: 422 application/json response
  - PUT /products/{product_id}: request body
  - PUT /products/{product_id}: application/json request body
  - PUT /products/{product_id}: 200 application/json response
  - PUT /products/{product_id}: 422 application/json response
  - DELETE /products/{product_id}: 422 application/json response
  - POST /tasks: request body
  - POST /tasks: application/json request body
  - POST /tasks: 201 application/json response
  - POST /tasks: 422 application/json response
  - POST /test/echo: request body
  - POST /test/echo: application/json request body
  - POST /test/echo: 200 application/json response
  - POST /test/echo: 422 application/json response
  - GET /test/health: 200 application/json response
  - GET /legacy/products: 200 application/json response
_All schemas have descriptions:_
  - HTTPValidationError
  - ProductBase
  - Product
  - ProductCreate
  - ValidationError
  - ProductCategory
  - TaskCreate
  - Task

**How to fix:** Add missing documentation including descriptions, examples, and operation details

</details>

<details><summary>🚨 <strong>P003: Error Handling Standards</strong> (critical)</summary>

**Result:** Error handling issues found: All operations document 5xx error responses: GET /products, POST /products, DELETE /products/{product_id}, GET /products/{product_id}, PUT /products/{product_id}, POST /tasks, POST /test/echo, GET /test/health, GET /legacy/products; Error responses include error details schema: GET /products: 422 response, POST /products: 422 response, DELETE /products/{product_id}: 422 response, GET /products/{product_id}: 422 response, PUT /products/{product_id}: 422 response, POST /tasks: 422 response, POST /test/echo: 422 response; All operations document 4xx error responses: GET /test/health, GET /legacy/products

**Failed checks:**

- ❌ All operations document 4xx error responses
- ❌ All operations document 5xx error responses
- ❌ Common error responses are defined in components — No common error responses defined in components
- ❌ Error responses include error details schema

<details><summary>Passed checks (1)</summary>

- ✅ Error responses follow consistent format

</details>

**Missing errors:**

_All operations document 4xx error responses:_
  - GET /test/health
  - GET /legacy/products
_All operations document 5xx error responses:_
  - GET /products
  - POST /products
  - DELETE /products/{product_id}
  - GET /products/{product_id}
  - PUT /products/{product_id}
  - POST /tasks
  - POST /test/echo
  - GET /test/health
  - GET /legacy/products
_Error responses include error details schema:_
  - GET /products: 422 response
  - POST /products: 422 response
  - DELETE /products/{product_id}: 422 response
  - GET /products/{product_id}: 422 response
  - PUT /products/{product_id}: 422 response
  - POST /tasks: 422 response
  - POST /test/echo: 422 response

**How to fix:** Add comprehensive error response documentation including codes, messages, and consistent error schemas

</details>

<details><summary>🚨 <strong>P004: Request Schema Definitions</strong> (critical)</summary>

**Result:** Request validation issues found: All string fields have length constraints: POST /tasks.description: application/json schema, POST /products.tags[]: application/json schema, POST /products.category: application/json schema, POST /products.description: application/json schema, PUT /products/{product_id}.tags[]: application/json schema, PUT /products/{product_id}.category: application/json schema, PUT /products/{product_id}.description: application/json schema

**Failed checks:**

- ❌ All string fields have length constraints

<details><summary>Passed checks (9)</summary>

- ✅ All enums have valid values
- ✅ All header parameters have schemas
- ✅ All numeric fields have min/max values
- ✅ All path parameters have schemas
- ✅ All query parameters have schemas
- ✅ All request bodies have content schemas
- ✅ All required fields are marked
- ✅ All schemas have appropriate constraints
- ✅ All schemas specify data types

</details>

**Missing validation:**

_All string fields have length constraints:_
  - POST /tasks.description: application/json schema
  - POST /products.tags[]: application/json schema
  - POST /products.category: application/json schema
  - POST /products.description: application/json schema
  - PUT /products/{product_id}.tags[]: application/json schema
  - PUT /products/{product_id}.category: application/json schema
  - PUT /products/{product_id}.description: application/json schema

**How to fix:** Add comprehensive schema validation including data types, constraints, and required fields

</details>

<details><summary>⚠️ <strong>P008: API Versioning Strategy</strong> (warning)</summary>

**Result:** Versioning validation failed: Deprecation notices are present: Deprecated operations missing deprecation details in description: GET /legacy/products

**Failed checks:**

- ❌ Deprecation notices are present — Deprecated operations missing deprecation details in description: GET /legacy/products

<details><summary>Passed checks (6)</summary>

- ✅ API version is specified
- ✅ Breaking changes are documented
- ✅ Migration guides are referenced
- ✅ Version compatibility is specified
- ✅ Version follows semantic versioning
- ✅ Versioning strategy is documented

</details>

**How to fix:** Update the info section with version details, deprecation notices, and migration references

</details>

<details><summary>📋 <strong>Quick Fixes</strong></summary>

- **P002:** Add missing documentation including descriptions, examples, and operation details
- **P003:** Add comprehensive error response documentation including codes, messages, and consistent error schemas
- **P004:** Add comprehensive schema validation including data types, constraints, and required fields
- **P008:** Update the info section with version details, deprecation notices, and migration references

</details>


---
*Generated by [DriveBy](https://github.com/meter-peter/driveby) — Documentation-Driven Testing for the GitOps era.*

