package spec

import (
	"context"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
)

// swagger2Spec is an APISpec implementation that wraps *openapi2.T.
type swagger2Spec struct {
	doc *openapi2.T
}

// NewSwagger2Spec creates a new APISpec adapter from an *openapi2.T document.
func NewSwagger2Spec(doc *openapi2.T) APISpec {
	if doc == nil {
		return nil
	}
	return &swagger2Spec{doc: doc}
}

func (s *swagger2Spec) Type() SpecType {
	return SpecTypeSwagger2
}

func (s *swagger2Spec) RawVersion() string {
	if s.doc == nil {
		return ""
	}
	return s.doc.Swagger
}

func (s *swagger2Spec) Info() *SpecInfo {
	if s.doc == nil {
		return nil
	}
	info := &s.doc.Info
	if info.Title == "" && info.Version == "" && info.Description == "" {
		return nil
	}
	result := &SpecInfo{
		Title:       info.Title,
		Version:     info.Version,
		Description: info.Description,
	}
	if info.Contact != nil {
		result.Contact = &SpecContact{
			Name:  info.Contact.Name,
			Email: info.Contact.Email,
			URL:   info.Contact.URL,
		}
	}
	if info.License != nil {
		result.License = &SpecLicense{
			Name: info.License.Name,
			URL:  info.License.URL,
		}
	}
	return result
}

func (s *swagger2Spec) Paths() map[string]*PathItem {
	if s.doc == nil || s.doc.Paths == nil {
		return nil
	}
	result := make(map[string]*PathItem)
	for path, pathItem := range s.doc.Paths {
		if pathItem == nil {
			continue
		}
		pi := &PathItem{
			Parameters: normalizeV2Parameters(pathItem.Parameters),
			Operations: make(map[string]*Operation),
		}

		if pathItem.Get != nil {
			pi.Operations["GET"] = normalizeV2Operation("GET", path, pathItem.Get)
		}
		if pathItem.Post != nil {
			pi.Operations["POST"] = normalizeV2Operation("POST", path, pathItem.Post)
		}
		if pathItem.Put != nil {
			pi.Operations["PUT"] = normalizeV2Operation("PUT", path, pathItem.Put)
		}
		if pathItem.Delete != nil {
			pi.Operations["DELETE"] = normalizeV2Operation("DELETE", path, pathItem.Delete)
		}
		if pathItem.Patch != nil {
			pi.Operations["PATCH"] = normalizeV2Operation("PATCH", path, pathItem.Patch)
		}
		if pathItem.Head != nil {
			pi.Operations["HEAD"] = normalizeV2Operation("HEAD", path, pathItem.Head)
		}
		if pathItem.Options != nil {
			pi.Operations["OPTIONS"] = normalizeV2Operation("OPTIONS", path, pathItem.Options)
		}

		result[path] = pi
	}
	return result
}

func (s *swagger2Spec) Components() *Components {
	if s.doc == nil {
		return nil
	}

	out := &Components{
		Schemas:         make(map[string]*Schema),
		Responses:       make(map[string]*Response),
		SecuritySchemes: make(map[string]*SecurityScheme),
	}

	// Definitions -> Schemas (openapi2.Definitions is map[string]*openapi3.SchemaRef)
	for name, schemaRef := range s.doc.Definitions {
		if schemaRef == nil || schemaRef.Value == nil {
			continue
		}
		out.Schemas[name] = normalizeSchema(schemaRef.Value)
	}

	// Responses
	for name, resp := range s.doc.Responses {
		if resp == nil {
			continue
		}
		out.Responses[name] = normalizeV2Response(resp)
	}

	// SecurityDefinitions -> SecuritySchemes
	for name, scheme := range s.doc.SecurityDefinitions {
		if scheme == nil {
			continue
		}
		out.SecuritySchemes[name] = &SecurityScheme{
			Type:         scheme.Type,
			Name:         scheme.Name,
			In:           scheme.In,
			Scheme:       "", // Swagger 2.0 does not have a "scheme" field
			BearerFormat: "",
			Description:  scheme.Description,
		}
	}

	if len(out.Schemas) == 0 && len(out.Responses) == 0 && len(out.SecuritySchemes) == 0 {
		return nil
	}
	return out
}

func (s *swagger2Spec) Security() []SecurityRequirement {
	if s.doc == nil || s.doc.Security == nil {
		return nil
	}
	var result []SecurityRequirement
	for _, req := range s.doc.Security {
		sr := SecurityRequirement{}
		for name, scopes := range req {
			copied := make([]string, len(scopes))
			copy(copied, scopes)
			sr[name] = copied
		}
		result = append(result, sr)
	}
	return result
}

func (s *swagger2Spec) ValidateStructure(ctx context.Context) error {
	if s.doc == nil {
		return nil
	}
	if s.doc.Swagger != "2.0" {
		return &ValidationError{"swagger field must be \"2.0\""}
	}
	if s.doc.Info.Title == "" || s.doc.Info.Version == "" {
		return &ValidationError{"info.title and info.version are required"}
	}
	if len(s.doc.Paths) == 0 {
		return &ValidationError{"at least one path must be defined"}
	}
	return nil
}

// ValidationError is a lightweight error type for structural validation issues.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// --- Normalization helpers for Swagger 2.0 ---

func normalizeV2Parameters(params []*openapi2.Parameter) []*Parameter {
	var result []*Parameter
	for _, p := range params {
		if p == nil {
			continue
		}
		var schema *Schema
		if p.Schema != nil && p.Schema.Value != nil {
			schema = normalizeSchema(p.Schema.Value)
		} else {
			schema = &Schema{
				Type:   p.Type,
				Format: p.Format,
			}
		}

		result = append(result, &Parameter{
			Name:        p.Name,
			In:          p.In,
			Description: p.Description,
			Required:    p.Required,
			Schema:      schema,
		})
	}
	return result
}

func normalizeV2Operation(method, path string, op *openapi2.Operation) *Operation {
	result := &Operation{
		Method:      method,
		Path:        path,
		Summary:     op.Summary,
		Description: op.Description,
		OperationID: op.OperationID,
		Deprecated:  op.Deprecated,
		Parameters:  normalizeV2Parameters(op.Parameters),
		Responses:   make(map[string]*Response),
	}

	// In Swagger 2.0, a body parameter represents the request body.
	for _, p := range op.Parameters {
		if p != nil && p.In == "body" && p.Schema != nil && p.Schema.Value != nil {
			rb := &RequestBody{
				Description: p.Description,
				Content: map[string]*MediaType{
					"application/json": {
						Schema: normalizeSchema(p.Schema.Value),
					},
				},
			}
			result.RequestBody = rb
			break
		}
	}

	// Responses
	for code, resp := range op.Responses {
		if resp == nil {
			continue
		}
		result.Responses[code] = normalizeV2Response(resp)
	}

	// Security
	if op.Security != nil {
		for _, req := range *op.Security {
			sr := SecurityRequirement{}
			for name, scopes := range req {
				copied := make([]string, len(scopes))
				copy(copied, scopes)
				sr[name] = copied
			}
			result.Security = append(result.Security, sr)
		}
	}

	return result
}

func normalizeV2Response(resp *openapi2.Response) *Response {
	if resp == nil {
		return nil
	}
	r := &Response{
		Description: resp.Description,
	}
	if resp.Schema != nil && resp.Schema.Value != nil {
		r.Content = map[string]*MediaType{
			"application/json": {
				Schema: normalizeSchema(resp.Schema.Value),
			},
		}
	}
	return r
}

// ensure openapi3 import is used
var _ *openapi3.Schema
