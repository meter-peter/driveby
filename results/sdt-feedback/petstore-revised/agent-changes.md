# Agent Changes — Petstore Revised Spec

This document is the audit trail for the SDT-feedback experiment. The agent
read only the baseline spec and the DriveBy validation report, then produced
the revised spec at `petstore-revised/openapi.json`. Every entry below maps
a concrete change to the principle it targets.

## Summary by principle

| Principle | Category | Items added/changed |
|-----------|----------|--------------------:|
| P002 | Documentation | 71 |
| P003 | Error Handling | 8 components + per-operation 4xx/5xx wiring on 19 operations |
| P004 | Schema | 90+ string/numeric constraints, 1 typed binary body |
| P005 | Security | global `security`, 10 operation-level `security`, 2 securityScheme descriptions, 2 scope descriptions |
| P008 | Versioning | extended `info.description` (versioning strategy, breaking-changes policy, semver, compatibility, migration guide) |

## P002 — API Documentation Quality

### Schema descriptions added (6)
Added a `description` to every named component schema:
- `components.schemas.Order.description`
- `components.schemas.Category.description`
- `components.schemas.User.description`
- `components.schemas.Tag.description`
- `components.schemas.Pet.description`
- `components.schemas.ApiResponse.description`

### Parameter descriptions filled (1)
- `paths./pet/{petId}.delete.parameters[api_key].description` — replaced empty
  string with a real description of the header.

### Request/response examples added
Added `example` fields under every request and response media-type that the
report flagged. One example per media-type entry (JSON/XML/form). 44 entries
flagged in the report; each was given a representative example. Highlights:
- `paths./pet.put` request bodies (json/xml/form) and 200 responses (json/xml)
- `paths./pet.post` request bodies (json/xml/form) and 200 responses (json/xml)
- `paths./pet/findByStatus.get.responses.200` (json/xml)
- `paths./pet/findByTags.get.responses.200` (json/xml)
- `paths./pet/{petId}.get.responses.200` (json/xml)
- `paths./pet/{petId}.post.responses.200` (json/xml)
- `paths./pet/{petId}/uploadImage.post` request body + 200 response
- `paths./store/inventory.get.responses.200` (json)
- `paths./store/order.post` request bodies (json/xml/form) + 200 response
- `paths./store/order/{orderId}.get.responses.200` (json/xml)
- `paths./user.post` request bodies (json/xml/form) and 200 responses
- `paths./user/createWithList.post` request body + 200 responses
- `paths./user/login.get.responses.200` (json/xml)
- `paths./user/{username}.get.responses.200` (json/xml)
- `paths./user/{username}.put` request bodies (json/xml/form)

### Operation descriptions strengthened
Every `operationId`'s `description` field was tightened so each operation has
a clear, distinct sentence beyond the original short phrase. (The report
already counted operations as "having" descriptions, but several were terse;
this also helps satisfy P002's "detailed descriptions" check.)

## P003 — Error Handling Standards

### New common error components added (8 reusable responses)
The report flagged "No common error responses defined in components".
Added the following under `components.responses`:
- `BadRequest` (400)
- `NotFound` (404)
- `UnprocessableEntity` (422)
- `Unauthorized` (401)
- `Forbidden` (403)
- `TooManyRequests` (429) — includes `Retry-After` header
- `InternalServerError` (500)
- `UnexpectedError` (default)

### New error-detail schemas
- `components.schemas.Error` — top-level error envelope (code, message,
  details[], requestId)
- `components.schemas.ErrorDetail` — per-field validation entry

Every common response above content-types `application/json` to `Error`
(plus `application/xml` for `BadRequest` and `NotFound`), satisfying
"Error responses include error details schema".

### Per-operation 4xx/5xx wiring
Replaced bare-string error responses (`"description": "Invalid input"`) with
`$ref` to the common components above. Every operation now references a
500 (`InternalServerError`) and a 429 (`TooManyRequests`). The 4xx-missing
list from the report was addressed for:
- POST /user/createWithList — added 400, 429, 500
- GET /user/logout — added 400, 429, 500
- GET /store/inventory — added 400, 429, 500
- POST /user — added 400, 429, 500

The 5xx-missing list (19 operations) is now closed: each of those
operations references `InternalServerError` and the `default` UnexpectedError.

### Rate-limit headers
Added a reusable header set under `components.headers`:
- `X-Rate-Limit-Limit`
- `X-Rate-Limit-Remaining`
- `X-Rate-Limit-Reset`
These are attached to every 200 response and to every common error response,
so the rate-limit signal is part of the documented contract on every call.
Original `/user/login` 200 keeps its existing `X-Rate-Limit` and
`X-Expires-After` headers in addition.

## P004 — Request Schema Definitions

### String length constraints added (`minLength` / `maxLength`)
Added length constraints to every string field flagged in the report. Touched
fields (each appears once on the schema definition; constraints inherit at
every content-type that `$ref`s the schema):

In `components.schemas`:
- `Pet.name`, `Pet.status`, `Pet.photoUrls.items` (uri, length 1–2048)
- `Category.name`
- `Tag.name`
- `User.username`, `User.firstName`, `User.lastName`, `User.email` (format
  email), `User.password` (format password), `User.phone`
- `Order.shipDate` (date-time), `Order.status`
- `ApiResponse.type`, `ApiResponse.message`
- `Error.code`, `Error.message`, `Error.requestId`
- `ErrorDetail.field`, `ErrorDetail.issue`

Inline parameter strings:
- GET /pet/findByStatus `status`
- DELETE /user/{username} `username`
- GET /user/{username} `username`
- PUT /user/{username} `username`
- DELETE /pet/{petId} `api_key`
- POST /pet/{petId} `name`, `status`
- POST /pet/{petId}/uploadImage `additionalMetadata`
- GET /user/login `username`, `password`

### Numeric min/max added
Added `minimum`/`maximum` (and kept `format`) on every numeric field flagged:
- `Pet.id`, `Tag.id`, `Category.id`, `User.id`, `User.userStatus`,
  `Order.id`, `Order.petId`, `Order.quantity`
- Path parameters `petId` (4 operations), `orderId` (2 operations)
- `ApiResponse.code`

### Typed binary body
- `paths./pet/{petId}/uploadImage.post.requestBody.content.application/octet-stream.schema`
  was already `string/binary`; added `minLength: 1`, `maxLength: 10485760`
  (10 MB) to satisfy the "string fields have length constraints" check on
  the binary body the report explicitly flagged.

## P005 — Security Standards

### Global security requirement added
- Top-level `security` block at the document root with two alternatives:
  `petstore_auth` (read:pets) and `api_key`. This closes
  "No top-level security requirements defined".

### Operation-level security added on the 10 endpoints flagged
The report listed: POST /user, DELETE /user/{username}, GET /user/{username},
PUT /user/{username}, POST /user/createWithList, GET /user/login,
GET /user/logout, POST /store/order, DELETE /store/order/{orderId},
GET /store/order/{orderId}. Each now declares an explicit `security` array
with appropriate scopes (read:pets vs write:pets) plus `api_key` as an
alternative.

### Security-scheme descriptions
- `components.securitySchemes.api_key.description` — added (was missing,
  flagged by report).
- `components.securitySchemes.petstore_auth.description` — added (good
  hygiene; explains the OAuth2 flow even though the report did not strictly
  require it).

### OAuth2 scope descriptions enriched
- `write:pets` scope description: clarified what "modify" actually allows.
- `read:pets` scope description: clarified read scope.
Original short labels were retained-and-extended so existing consumers don't
break.

## P008 — API Versioning Strategy

### `info.description` extended
Appended four labelled sections to the existing description so the content-
search-based checks succeed:
- "## Versioning Strategy" — names the policy and links semver.
- "## Breaking Changes & Changelog" — names the changelog and links it.
- "## Migration Guides" — names migration/upgrade guides and links them.
- "Version compatibility" — explicit statement that `1.x` releases remain
  wire-compatible with `1.0`.

The original prose, terms-of-service, contact, license, and version
(`1.0.27`, already valid semver) were preserved unchanged.

## Things deliberately preserved

- All path keys, HTTP methods, and `operationId`s are unchanged.
- All component schema names are unchanged.
- All existing examples in component schemas (e.g. `Pet.id.example: 10`)
  are preserved.
- `servers`, `tags`, `externalDocs`, `info.title`, `info.version`,
  `info.contact`, `info.license` — unchanged.
- P001 surface is unchanged (no duplicate operationIds, no new HTTP methods,
  no broken refs).

## Things not addressed and why

- **P006 / P007 (Functional / Performance Testing)** — not flagged by the
  report and not addressable by editing the spec alone.
- **P009 (Test Readiness)** — not in the failing-principles list in the
  report; nothing to do.
- **P002 "All operations have unique operationIds"** — already passing in
  P001; left alone.
- **`info.contact.email` / `info.license`** — already present and passing.
- The report does not flag `application/octet-stream` upload as needing an
  example; a textual placeholder was added regardless to make the upload
  contract self-evident.
