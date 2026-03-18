# Spec Package — APISpec Abstraction Layer

## Purpose
Provides a **normalized interface** (`APISpec`) over OpenAPI 3.x and Swagger 2.x specifications. This is the mandatory abstraction boundary -- all principle checkers and engine code MUST use `spec.APISpec`, never raw kin-openapi types (`*openapi3.T`, `*openapi2.T`).

## Files

| File | Purpose |
|------|---------|
| `spec.go` | `APISpec` interface definition + normalized types (Schema, Operation, PathItem, Response, Parameter, SecurityScheme, Components) |
| `openapi3.go` | Adapter wrapping `*openapi3.T` from kin-openapi into `APISpec` |
| `swagger2.go` | Adapter wrapping `*openapi2.T` from kin-openapi into `APISpec` |

## APISpec Interface
The interface exposes normalized accessors:
- `GetPaths() map[string]PathItem` — all path items keyed by path string
- `GetComponents() Components` — schemas, security schemes, responses
- `GetInfo()` — title, version, description
- `GetServers()` — server URLs
- `GetSecurity()` — global security requirements

## Normalized Types
These types decouple principle checkers from spec version details:
- `Schema` — type, properties, required fields, allOf/oneOf/anyOf
- `Operation` — method, summary, description, parameters, responses, security
- `PathItem` — operations keyed by HTTP method
- `Response` — status code, description, content/schema
- `Parameter` — name, in (query/header/path/cookie), required, schema
- `SecurityScheme` — type (apiKey/http/oauth2/openIdConnect), details
- `Components` — named schemas, security schemes, responses

## Key Design Decisions
- **Adapter pattern**: `openapi3.go` and `swagger2.go` each implement `APISpec` by wrapping their respective kin-openapi structs
- **FlattenAllOf helper**: Merges `allOf` compositions into a single schema for principle checkers that need flat property access
- **No version-specific logic leaks**: If a principle needs something version-specific, it goes in the adapter, not the checker

## Thesis Mapping
- **Chapter 4 (Architecture)**: The spec abstraction layer is a key architectural contribution, enabling version-agnostic validation

## Rules
1. **NEVER** pass `*openapi3.T` or `*openapi2.T` to principle checkers
2. When adding new spec accessors, add to the interface in `spec.go` and implement in BOTH adapters
3. Normalized types must remain version-agnostic -- no OpenAPI-version-specific fields
