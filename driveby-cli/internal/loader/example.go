package loader

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// GetExampleValues generates example values for a schema
func (l *Loader) GetExampleValues(schema *openapi3.Schema) map[string]interface{} {
	log.Debugf("[loader] Enter GetExampleValues with schema: %+v", schema)
	if schema == nil {
		log.Warn("[loader] Schema is nil")
		return nil
	}

	examples := make(map[string]interface{})

	switch schema.Type {
	case "string":
		if schema.Enum != nil && len(schema.Enum) > 0 {
			examples["value"] = schema.Enum[0]
		} else if schema.Format == "date-time" {
			examples["value"] = "2024-01-01T00:00:00Z"
		} else if schema.Format == "date" {
			examples["value"] = "2024-01-01"
		} else if schema.Format == "email" {
			examples["value"] = "example@example.com"
		} else if schema.Format == "uuid" {
			examples["value"] = "123e4567-e89b-12d3-a456-426614174000"
		} else {
			examples["value"] = "example string"
		}
	case "number", "integer":
		examples["value"] = 42
	case "boolean":
		examples["value"] = true
	case "array":
		if schema.Items != nil {
			examples["value"] = []interface{}{l.GetExampleValues(schema.Items.Value)}
		} else {
			examples["value"] = []interface{}{}
		}
	case "object":
		if len(schema.Properties) > 0 {
			obj := make(map[string]interface{})
			for propName, propSchema := range schema.Properties {
				log.Debugf("[loader] Object property: %s", propName)
				obj[propName] = l.GetExampleValues(propSchema.Value)["value"]
			}
			examples["value"] = obj
		} else {
			examples["value"] = map[string]interface{}{}
		}
	}

	log.Debugf("[loader] Example values generated: %v", examples)
	return examples
}
