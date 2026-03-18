package spec

import "context"

// SpecType identifies the underlying specification flavour.
//
// Currently supported values:
//   - "openapi3"  -> OpenAPI 3.0 / 3.1
//   - "swagger2"  -> Swagger 2.0
type SpecType string

const (
	SpecTypeOpenAPI3 SpecType = "openapi3"
	SpecTypeSwagger2 SpecType = "swagger2"
)

// APISpec is a version-agnostic view over an API specification.
//
// Implementations (adapters) wrap concrete spec structs (e.g. *openapi3.T, *openapi2.T)
// and expose only the concepts that DriveBy's validators and testers need.
type APISpec interface {
	// Type returns which concrete spec flavour is wrapped.
	Type() SpecType

	// RawVersion returns the version string from the underlying document,
	// e.g. "3.0.3", "3.1.0", or "2.0".
	RawVersion() string

	// Info returns high-level API metadata.
	Info() *SpecInfo

	// Paths returns all paths in the API, keyed by raw path string (e.g. "/pets").
	Paths() map[string]*PathItem

	// Components returns reusable schemas, responses, and security schemes, if present.
	Components() *Components

	// Security returns the top-level security requirements (if any).
	Security() []SecurityRequirement

	// ValidateStructure performs structural validation appropriate for the
	// underlying specification flavour (e.g. OpenAPI 3.x schema validation,
	// basic sanity checks for Swagger 2.0).
	ValidateStructure(ctx context.Context) error
}

// SpecInfo is a version-agnostic representation of the top-level info section.
type SpecInfo struct {
	Title       string
	Version     string
	Description string
	Contact     *SpecContact
	License     *SpecLicense
}

type SpecContact struct {
	Name  string
	Email string
	URL   string
}

type SpecLicense struct {
	Name string
	URL  string
}

// Schema is a minimal, normalized view over a schema definition used by
// DriveBy's validation principles. It intentionally models only the fields
// that are currently accessed by validators.
type Schema struct {
	Type        string
	Format      string
	MinLength   uint64
	MaxLength   *uint64
	Pattern     string
	Min         *float64
	Max         *float64
	Enum        []interface{}
	Required    []string
	Properties  map[string]*Schema
	Items       *Schema
	Description string
}

// MediaType represents a content-type specific payload with an optional schema
// and examples.
type MediaType struct {
	Schema   *Schema
	Example  interface{}
	Examples map[string]interface{}
}

// Response is a normalized response definition keyed by HTTP status code in
// the APISpec interface.
type Response struct {
	Description string
	Content     map[string]*MediaType
}

// Parameter models a request parameter in any location (path, query, header, cookie/body).
type Parameter struct {
	Name        string
	In          string
	Description string
	Required    bool
	Schema      *Schema
}

// RequestBody is the unified representation of a request payload schema.
type RequestBody struct {
	Description string
	Content     map[string]*MediaType
}

// SecurityRequirement represents a single security requirement mapping scheme
// names to scopes.
type SecurityRequirement map[string][]string

// Operation is a normalized view of an HTTP operation on a path.
type Operation struct {
	Method      string
	Path        string
	Summary     string
	Description string
	OperationID string
	Deprecated  bool

	Parameters  []*Parameter
	RequestBody *RequestBody
	Responses   map[string]*Response

	Security []SecurityRequirement
}

// PathItem groups operations and parameters for a single path.
type PathItem struct {
	Parameters []*Parameter
	Operations map[string]*Operation
}

// SecurityScheme is a minimal representation of an auth scheme used by P005.
type SecurityScheme struct {
	Type        string
	Name        string
	In          string
	Scheme      string
	BearerFormat string
	Description string
}

// Components aggregates reusable definitions referenced from operations.
type Components struct {
	Schemas         map[string]*Schema
	Responses       map[string]*Response
	SecuritySchemes map[string]*SecurityScheme
}

