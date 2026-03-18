package loader

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	stdlog "log"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/driveby-cli/internal/spec"
	"github.com/meter-peter/driveby/driveby-cli/internal/util"
	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()

func init() {
	stdlog.SetOutput(log.Writer())
}

// Loader handles loading and validating OpenAPI/Swagger specifications
type Loader struct {
	doc spec.APISpec
}

// NewLoader creates a new OpenAPI loader
func NewLoader() *Loader {
	log.Debug("[loader] Creating new OpenAPI Loader")
	return &Loader{}
}

// LoadFromFile loads an OpenAPI specification from a file
func (l *Loader) LoadFromFile(path string) error {
	log.Debugf("[loader] Enter LoadFromFile with path: %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	if err := l.loadFromData(path, data); err != nil {
		return err
	}
	log.Infof("[loader] Successfully loaded OpenAPI spec from file: %s", path)
	return nil
}

// LoadFromURL loads an OpenAPI specification from a URL
func (l *Loader) LoadFromURL(url string) error {
	log.Debugf("[loader] Enter LoadFromURL with url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Errorf("[loader] Failed to fetch OpenAPI spec from URL: %s", url)
		return fmt.Errorf("failed to fetch OpenAPI spec from URL: %w", err)
	}
	defer resp.Body.Close()

	log.Debugf("[loader] HTTP status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		log.Errorf("[loader] Failed to fetch OpenAPI spec: status %d", resp.StatusCode)
		return fmt.Errorf("failed to fetch OpenAPI spec: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Errorf("[loader] Failed to read OpenAPI spec from response: %s", url)
		return fmt.Errorf("failed to read OpenAPI spec from response: %w", err)
	}
	log.Debugf("[loader] Read %d bytes from response", len(data))

	if err := l.loadFromData(url, data); err != nil {
		return err
	}
	log.Infof("[loader] Successfully loaded OpenAPI spec from URL: %s", url)
	return nil
}

// LoadFromFileOrURL loads an OpenAPI spec from a local file or a URL
func (l *Loader) LoadFromFileOrURL(path string) error {
	if path == "" {
		return fmt.Errorf("OpenAPI spec path is empty")
	}
	log.Debugf("[loader] Enter LoadFromFileOrURL with path: %s", path)
	var data []byte
	var err error
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		log.Debugf("[loader] Detected URL, fetching: %s", path)
		resp, err := http.Get(path)
		if err != nil {
			log.WithError(err).Errorf("[loader] Failed to fetch OpenAPI spec from URL: %s", path)
			return err
		}
		defer resp.Body.Close()
		log.Debugf("[loader] HTTP status: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusOK {
			log.Errorf("[loader] Failed to fetch OpenAPI spec: status %d", resp.StatusCode)
			return fmt.Errorf("failed to fetch OpenAPI spec: status %d", resp.StatusCode)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			log.WithError(err).Errorf("[loader] Failed to read OpenAPI spec from response: %s", path)
			return err
		}
		log.Debugf("[loader] Read %d bytes from response", len(data))
	} else {
		log.Debugf("[loader] Detected file, reading: %s", path)
		data, err = os.ReadFile(path)
		if err != nil {
			log.WithError(err).Errorf("[loader] Failed to read OpenAPI spec from file: %s", path)
			return err
		}
		log.Debugf("[loader] Read %d bytes from file", len(data))
	}

	if err := l.loadFromData(path, data); err != nil {
		return err
	}
	log.Infof("[loader] Successfully loaded OpenAPI spec from: %s", path)
	return nil
}

// loadFromData inspects the raw document, detects the specification flavour,
// and populates the Loader with the appropriate APISpec adapter.
func (l *Loader) loadFromData(source string, data []byte) error {
	// First, try to unmarshal into a generic map to detect "openapi" vs "swagger".
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.WithError(err).Errorf("[loader] Failed to unmarshal spec for detection: %s", source)
		return fmt.Errorf("failed to parse spec for detection: %w", err)
	}

	// OpenAPI 3.x: has "openapi" field.
	if ver, ok := raw["openapi"].(string); ok && ver != "" {
		log.Debugf("[loader] Detected OpenAPI spec (version=%s) for: %s", ver, source)

		// Preprocess OpenAPI 3.1.0 nullable anyOf patterns before kin-openapi loads them.
		util.PreprocessNullableAnyOf(raw)
		// Preprocess exclusiveMinimum/exclusiveMaximum for native OpenAPI 3.x specs.
		util.PreprocessExclusiveMinMax(raw)
		processed, err := json.Marshal(raw)
		if err != nil {
			log.WithError(err).Errorf("[loader] Failed to marshal preprocessed OpenAPI data for: %s", source)
			return err
		}

		kinLoader := openapi3.NewLoader()
		doc, err := kinLoader.LoadFromData(processed)
		if err != nil {
			log.WithError(err).Errorf("[loader] Failed to load OpenAPI 3.x spec from data: %s", source)
			return err
		}
		l.doc = spec.NewOpenAPI3Spec(doc)
		log.Debugf("[loader] Loaded OpenAPI 3.x doc: %+v", doc)
		return nil
	}

	// Swagger 2.0: has "swagger" field with value "2.0".
	if ver, ok := raw["swagger"].(string); ok && ver == "2.0" {
		log.Debugf("[loader] Detected Swagger 2.0 spec for: %s", source)
		var doc openapi2.T
		if err := json.Unmarshal(data, &doc); err != nil {
			log.WithError(err).Errorf("[loader] Failed to unmarshal Swagger 2.0 spec: %s", source)
			return fmt.Errorf("failed to unmarshal Swagger 2.0 spec: %w", err)
		}
		l.doc = spec.NewSwagger2Spec(&doc)
		log.Debugf("[loader] Loaded Swagger 2.0 doc")
		return nil
	}

	return fmt.Errorf("unsupported API specification format for: %s (missing 'openapi' or 'swagger: \"2.0\"')", source)
}

// Validate validates the loaded OpenAPI specification
func (l *Loader) Validate() error {
	log.Debug("[loader] Enter Validate")
	if l.doc == nil {
		log.Error("[loader] No OpenAPI specification loaded")
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	if err := l.doc.ValidateStructure(context.Background()); err != nil {
		log.WithError(err).Error("[loader] Invalid API specification")
		return fmt.Errorf("invalid API specification: %w", err)
	}
	log.Debug("[loader] API specification is valid")
	return nil
}

// GetDocument returns the loaded API specification in normalized form.
func (l *Loader) GetDocument() spec.APISpec {
	log.Debugf("[loader] GetDocument called, spec type: %T", l.doc)
	return l.doc
}

// GetEndpoints returns a list of all endpoints in the specification
func (l *Loader) GetEndpoints() []string {
	log.Debug("[loader] Enter GetEndpoints")
	if l.doc == nil {
		log.Warn("[loader] No document loaded")
		return nil
	}

	paths := l.doc.Paths()
	if paths == nil {
		log.Warn("[loader] No paths loaded")
		return nil
	}

	var endpoints []string
	for path, pathItem := range paths {
		if pathItem == nil || pathItem.Operations == nil {
			continue
		}
		log.Debugf("[loader] Path: %s", path)
		for method := range pathItem.Operations {
			log.Debugf("[loader] Method: %s for path %s", method, path)
			endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
		}
	}
	log.Debugf("[loader] Endpoints found: %v", endpoints)
	return endpoints
}

// SaveToFile saves the OpenAPI specification to a file
func (l *Loader) SaveToFile(path string) error {
	log.Debugf("[loader] Enter SaveToFile with path: %s", path)
	if l.doc == nil {
		log.Error("[loader] No OpenAPI specification loaded")
		return fmt.Errorf("no OpenAPI specification loaded")
	}

	// Persist the original JSON form is not currently tracked; for now,
	// marshal the normalized OpenAPI 3.x representation when available.
	// Swagger 2.0 documents will be re-encoded from the adapter.
	var data []byte
	var err error
	switch d := l.doc.(type) {
	case interface{ MarshalJSON() ([]byte, error) }:
		data, err = d.MarshalJSON()
	default:
		data, err = json.Marshal(d)
	}
	if err != nil {
		log.WithError(err).Errorf("[loader] Failed to marshal API spec for saving: %s", path)
		return fmt.Errorf("failed to marshal API spec: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.WithError(err).Errorf("[loader] Failed to create directory for saving: %s", path)
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		log.WithError(err).Errorf("[loader] Failed to write OpenAPI spec to file: %s", path)
		return fmt.Errorf("failed to write OpenAPI spec to file: %w", err)
	}

	log.Infof("[loader] Successfully saved OpenAPI spec to file: %s", path)
	return nil
}
