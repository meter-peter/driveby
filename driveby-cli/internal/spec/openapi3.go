package spec

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
)

// openAPI3Spec is an APISpec implementation that wraps *openapi3.T.
type openAPI3Spec struct {
	doc *openapi3.T
}

// NewOpenAPI3Spec creates a new APISpec adapter from an *openapi3.T document.
func NewOpenAPI3Spec(doc *openapi3.T) APISpec {
	if doc == nil {
		return nil
	}
	return &openAPI3Spec{doc: doc}
}

func (s *openAPI3Spec) Type() SpecType {
	return SpecTypeOpenAPI3
}

func (s *openAPI3Spec) RawVersion() string {
	if s.doc == nil {
		return ""
	}
	return s.doc.OpenAPI
}

func (s *openAPI3Spec) Info() *SpecInfo {
	if s.doc == nil || s.doc.Info == nil {
		return nil
	}
	info := s.doc.Info
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

func (s *openAPI3Spec) Paths() map[string]*PathItem {
	if s.doc == nil || s.doc.Paths == nil {
		return nil
	}
	result := make(map[string]*PathItem)
	for path, pathItem := range s.doc.Paths.Map() {
		if pathItem == nil {
			continue
		}
		pi := &PathItem{
			Parameters: normalizeParameters(pathItem.Parameters),
			Operations: make(map[string]*Operation),
		}
		for method, op := range pathItem.Operations() {
			if op == nil {
				continue
			}
			pi.Operations[method] = normalizeOperation(method, path, op)
		}
		result[path] = pi
	}
	return result
}

func (s *openAPI3Spec) Components() *Components {
	if s.doc == nil || s.doc.Components == nil {
		return nil
	}
	if s.doc.Components.Schemas == nil &&
		s.doc.Components.Responses == nil && s.doc.Components.SecuritySchemes == nil {
		return nil
	}

	out := &Components{
		Schemas:         make(map[string]*Schema),
		Responses:       make(map[string]*Response),
		SecuritySchemes: make(map[string]*SecurityScheme),
	}

	for name, schemaRef := range s.doc.Components.Schemas {
		if schemaRef == nil || schemaRef.Value == nil {
			continue
		}
		out.Schemas[name] = normalizeSchema(schemaRef.Value)
	}

	for name, respRef := range s.doc.Components.Responses {
		if respRef == nil || respRef.Value == nil {
			continue
		}
		out.Responses[name] = normalizeResponseV3(respRef.Value)
	}

	for name, secRef := range s.doc.Components.SecuritySchemes {
		if secRef == nil || secRef.Value == nil {
			continue
		}
		scheme := secRef.Value
		out.SecuritySchemes[name] = &SecurityScheme{
			Type:         scheme.Type,
			Name:         scheme.Name,
			In:           string(scheme.In),
			Scheme:       scheme.Scheme,
			BearerFormat: scheme.BearerFormat,
			Description:  scheme.Description,
		}
	}

	return out
}

func (s *openAPI3Spec) Security() []SecurityRequirement {
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

func (s *openAPI3Spec) ValidateStructure(ctx context.Context) error {
	if s.doc == nil {
		return nil
	}
	return s.doc.Validate(ctx)
}

// --- Normalization helpers ---

func normalizeParameters(params openapi3.Parameters) []*Parameter {
	var result []*Parameter
	for _, p := range params {
		if p == nil || p.Value == nil {
			continue
		}
		result = append(result, &Parameter{
			Name:        p.Value.Name,
			In:          p.Value.In,
			Description: p.Value.Description,
			Required:    p.Value.Required,
			Schema:      normalizeSchemaRef(p.Value.Schema),
		})
	}
	return result
}

func normalizeOperation(method, path string, op *openapi3.Operation) *Operation {
	result := &Operation{
		Method:      method,
		Path:        path,
		Summary:     op.Summary,
		Description: op.Description,
		OperationID: op.OperationID,
		Deprecated:  op.Deprecated,
		Parameters:  normalizeParameters(op.Parameters),
		Responses:   make(map[string]*Response),
	}

	// Request body
	if op.RequestBody != nil && op.RequestBody.Value != nil {
		rb := &RequestBody{
			Description: op.RequestBody.Value.Description,
			Content:     make(map[string]*MediaType),
		}
		for ct, mt := range op.RequestBody.Value.Content {
			rb.Content[ct] = normalizeMediaType(mt)
		}
		result.RequestBody = rb
	}

	// Responses
	if op.Responses != nil {
		for code, respRef := range op.Responses.Map() {
			if respRef == nil || respRef.Value == nil {
				continue
			}
			result.Responses[code] = normalizeResponseV3(respRef.Value)
		}
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

func normalizeSchemaRef(ref *openapi3.SchemaRef) *Schema {
	if ref == nil || ref.Value == nil {
		return nil
	}
	return normalizeSchema(ref.Value)
}

func normalizeSchema(s *openapi3.Schema) *Schema {
	if s == nil {
		return nil
	}

	out := &Schema{
		Type:        s.Type,
		Format:      s.Format,
		MinLength:   s.MinLength,
		MaxLength:   s.MaxLength,
		Pattern:     s.Pattern,
		Min:         s.Min,
		Max:         s.Max,
		Enum:        s.Enum,
		Required:    append([]string{}, s.Required...),
		Description: s.Description,
	}

	if len(s.Properties) > 0 {
		out.Properties = make(map[string]*Schema)
		for name, prop := range s.Properties {
			if prop == nil || prop.Value == nil {
				continue
			}
			out.Properties[name] = normalizeSchema(prop.Value)
		}
	}

	if s.Items != nil && s.Items.Value != nil {
		out.Items = normalizeSchema(s.Items.Value)
	}

	for _, ref := range s.AllOf {
		if ref != nil && ref.Value != nil {
			out.AllOf = append(out.AllOf, normalizeSchema(ref.Value))
		}
	}
	for _, ref := range s.OneOf {
		if ref != nil && ref.Value != nil {
			out.OneOf = append(out.OneOf, normalizeSchema(ref.Value))
		}
	}
	for _, ref := range s.AnyOf {
		if ref != nil && ref.Value != nil {
			out.AnyOf = append(out.AnyOf, normalizeSchema(ref.Value))
		}
	}

	return out
}

func normalizeMediaType(mt *openapi3.MediaType) *MediaType {
	if mt == nil {
		return nil
	}
	out := &MediaType{
		Schema:  normalizeSchemaRef(mt.Schema),
		Example: mt.Example,
	}
	if len(mt.Examples) > 0 {
		out.Examples = make(map[string]interface{})
		for name, ex := range mt.Examples {
			if ex == nil {
				continue
			}
			out.Examples[name] = ex.Value
		}
	}
	return out
}

func normalizeResponseV3(resp *openapi3.Response) *Response {
	if resp == nil {
		return nil
	}
	r := &Response{
		Description: derefString(resp.Description),
	}
	if len(resp.Content) > 0 {
		r.Content = make(map[string]*MediaType)
		for ct, mt := range resp.Content {
			r.Content[ct] = normalizeMediaType(mt)
		}
	}
	return r
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
