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

	stdlog "log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/util"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[openapi] Logger set to DEBUG (verbose) mode")
	// Remove viper.SetLogger as it's undefined
	stdlog.SetOutput(logrus.StandardLogger().Writer())
}

// Loader handles loading and validating OpenAPI specifications
type Loader struct {
	doc *openapi3.T
}

// NewLoader creates a new OpenAPI loader
func NewLoader() *Loader {
	log.Debug("[openapi] Creating new OpenAPI Loader")
	return &Loader{}
}

// LoadFromFile loads an OpenAPI specification from a file
func (l *Loader) LoadFromFile(path string) error {
	log.Debugf("[openapi] Enter LoadFromFile with path: %s", path)
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(path)
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to load OpenAPI spec from file: %s", path)
		return fmt.Errorf("failed to load OpenAPI spec from file: %w", err)
	}
	l.doc = doc
	log.Debugf("[openapi] Loaded OpenAPI doc: %+v", doc)
	log.Infof("[openapi] Successfully loaded OpenAPI spec from file: %s", path)
	return nil
}

// LoadFromURL loads an OpenAPI specification from a URL
func (l *Loader) LoadFromURL(url string) error {
	log.Debugf("[openapi] Enter LoadFromURL with url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to fetch OpenAPI spec from URL: %s", url)
		return fmt.Errorf("failed to fetch OpenAPI spec from URL: %w", err)
	}
	defer resp.Body.Close()

	log.Debugf("[openapi] HTTP status: %s", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Errorf("[openapi] Failed to fetch OpenAPI spec: status %s", resp.Status)
		return fmt.Errorf("failed to fetch OpenAPI spec: status %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to read OpenAPI spec from response: %s", url)
		return fmt.Errorf("failed to read OpenAPI spec from response: %w", err)
	}
	log.Debugf("[openapi] Read %d bytes from response", len(data))

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to load OpenAPI spec from data: %s", url)
		return fmt.Errorf("failed to load OpenAPI spec from data: %w", err)
	}
	l.doc = doc
	log.Debugf("[openapi] Loaded OpenAPI doc: %+v", doc)
	log.Infof("[openapi] Successfully loaded OpenAPI spec from URL: %s", url)
	return nil
}

// LoadFromFileOrURL loads an OpenAPI spec from a local file or a URL
func (l *Loader) LoadFromFileOrURL(path string) error {
	if path == "" {
		return fmt.Errorf("OpenAPI spec path is empty")
	}
	log.Debugf("[openapi] Enter LoadFromFileOrURL with path: %s", path)
	var data []byte
	var err error
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		log.Debugf("[openapi] Detected URL, fetching: %s", path)
		resp, err := http.Get(path)
		if err != nil {
			log.WithError(err).Errorf("[openapi] Failed to fetch OpenAPI spec from URL: %s", path)
			return err
		}
		defer resp.Body.Close()
		log.Debugf("[openapi] HTTP status: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusOK {
			log.Errorf("[openapi] Failed to fetch OpenAPI spec: status %d", resp.StatusCode)
			return fmt.Errorf("failed to fetch OpenAPI spec: status %d", resp.StatusCode)
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithError(err).Errorf("[openapi] Failed to read OpenAPI spec from response: %s", path)
			return err
		}
		log.Debugf("[openapi] Read %d bytes from response", len(data))
	} else {
		log.Debugf("[openapi] Detected file, reading: %s", path)
		data, err = ioutil.ReadFile(path)
		if err != nil {
			log.WithError(err).Errorf("[openapi] Failed to read OpenAPI spec from file: %s", path)
			return err
		}
		log.Debugf("[openapi] Read %d bytes from file", len(data))
	}
	// Preprocess exclusiveMinimum/exclusiveMaximum
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err == nil {
		log.Debugf("[openapi] Preprocessing exclusiveMinimum/exclusiveMaximum for: %s", path)
		util.PreprocessExclusiveMinMax(raw)
		if data, err = json.Marshal(raw); err != nil {
			log.WithError(err).Errorf("[openapi] Failed to marshal preprocessed data for: %s", path)
			return err
		}
		log.Debugf("[openapi] Marshaled preprocessed data, length: %d", len(data))
	}
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to load OpenAPI spec from data: %s", path)
		return err
	}
	l.doc = doc
	log.Debugf("[openapi] Loaded OpenAPI doc: %+v", doc)
	log.Infof("[openapi] Successfully loaded OpenAPI spec from: %s", path)
	return nil
}

// Validate validates the loaded OpenAPI specification
func (l *Loader) Validate() error {
	log.Debug("[openapi] Enter Validate")
	if l.doc == nil {
		log.Error("[openapi] No OpenAPI specification loaded")
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	if err := l.doc.Validate(context.Background()); err != nil {
		log.WithError(err).Error("[openapi] Invalid OpenAPI specification")
		return fmt.Errorf("invalid OpenAPI specification: %w", err)
	}
	log.Debug("[openapi] OpenAPI specification is valid")
	return nil
}

// GetDocument returns the loaded OpenAPI document
func (l *Loader) GetDocument() *openapi3.T {
	log.Debugf("[openapi] GetDocument called, doc: %+v", l.doc)
	return l.doc
}

// GetEndpoints returns a list of all endpoints in the specification
func (l *Loader) GetEndpoints() []string {
	log.Debug("[openapi] Enter GetEndpoints")
	if l.doc == nil || l.doc.Paths == nil {
		log.Warn("[openapi] No document or paths loaded")
		return nil
	}

	var endpoints []string
	for path, pathItem := range l.doc.Paths.Map() {
		log.Debugf("[openapi] Path: %s", path)
		for method := range pathItem.Operations() {
			log.Debugf("[openapi] Method: %s for path %s", method, path)
			endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
		}
	}
	log.Debugf("[openapi] Endpoints found: %v", endpoints)
	return endpoints
}

// GetExampleValues generates example values for a schema
func (l *Loader) GetExampleValues(schema *openapi3.Schema) map[string]interface{} {
	log.Debugf("[openapi] Enter GetExampleValues with schema: %+v", schema)
	if schema == nil {
		log.Warn("[openapi] Schema is nil")
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
				log.Debugf("[openapi] Object property: %s", propName)
				obj[propName] = l.GetExampleValues(propSchema.Value)["value"]
			}
			examples["value"] = obj
		} else {
			examples["value"] = map[string]interface{}{}
		}
	}

	log.Debugf("[openapi] Example values generated: %v", examples)
	return examples
}

// SaveToFile saves the OpenAPI specification to a file
func (l *Loader) SaveToFile(path string) error {
	log.Debugf("[openapi] Enter SaveToFile with path: %s", path)
	if l.doc == nil {
		log.Error("[openapi] No OpenAPI specification loaded")
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	data, err := l.doc.MarshalJSON()
	if err != nil {
		log.WithError(err).Errorf("[openapi] Failed to marshal OpenAPI spec for saving: %s", path)
		return fmt.Errorf("failed to marshal OpenAPI spec: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.WithError(err).Errorf("[openapi] Failed to create directory for saving: %s", path)
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		log.WithError(err).Errorf("[openapi] Failed to write OpenAPI spec to file: %s", path)
		return fmt.Errorf("failed to write OpenAPI spec to file: %w", err)
	}

	log.Infof("[openapi] Successfully saved OpenAPI spec to file: %s", path)
	return nil
}
