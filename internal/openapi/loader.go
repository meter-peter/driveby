package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/util"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[openapi] Logger set to DEBUG (verbose) mode")
}

// Loader handles loading and validating OpenAPI specifications
type Loader struct {
	doc *openapi3.T
}

// NewLoader creates a new OpenAPI loader
func NewLoader() *Loader {
	return &Loader{}
}

// LoadFromFile loads an OpenAPI specification from a file
func (l *Loader) LoadFromFile(path string) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("failed to load OpenAPI spec from file: %w", err)
	}
	l.doc = doc
	return nil
}

// LoadFromURL loads an OpenAPI specification from a URL
func (l *Loader) LoadFromURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch OpenAPI spec from URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch OpenAPI spec: status %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read OpenAPI spec from response: %w", err)
	}

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return fmt.Errorf("failed to load OpenAPI spec from data: %w", err)
	}
	l.doc = doc
	return nil
}

// LoadFromFileOrURL loads an OpenAPI spec from a local file or a URL
func (l *Loader) LoadFromFileOrURL(path string) error {
	var data []byte
	var err error
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to fetch OpenAPI spec: status %d", resp.StatusCode)
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	} else {
		data, err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}
	}
	// Preprocess exclusiveMinimum/exclusiveMaximum
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err == nil {
		util.PreprocessExclusiveMinMax(raw)
		if data, err = json.Marshal(raw); err != nil {
			return err
		}
	}
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return err
	}
	l.doc = doc
	return nil
}

// Validate validates the loaded OpenAPI specification
func (l *Loader) Validate() error {
	if l.doc == nil {
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	if err := l.doc.Validate(context.Background()); err != nil {
		return fmt.Errorf("invalid OpenAPI specification: %w", err)
	}

	return nil
}

// GetDocument returns the loaded OpenAPI document
func (l *Loader) GetDocument() *openapi3.T {
	return l.doc
}

// GetEndpoints returns a list of all endpoints in the specification
func (l *Loader) GetEndpoints() []string {
	if l.doc == nil || l.doc.Paths == nil {
		return nil
	}

	var endpoints []string
	for path, pathItem := range l.doc.Paths.Map() {
		for method := range pathItem.Operations() {
			endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
		}
	}
	return endpoints
}

// GetExampleValues generates example values for a schema
func (l *Loader) GetExampleValues(schema *openapi3.Schema) map[string]interface{} {
	if schema == nil {
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
				obj[propName] = l.GetExampleValues(propSchema.Value)["value"]
			}
			examples["value"] = obj
		} else {
			examples["value"] = map[string]interface{}{}
		}
	}

	return examples
}

// SaveToFile saves the OpenAPI specification to a file
func (l *Loader) SaveToFile(path string) error {
	if l.doc == nil {
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	data, err := l.doc.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal OpenAPI spec: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write OpenAPI spec to file: %w", err)
	}

	return nil
}
