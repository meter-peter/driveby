package validation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/meter-peter/driveby/internal/openapi"

	"github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

var log = logrus.New()

// ValidationResult represents the result of an API validation
type ValidationResult struct {
	Timestamp     time.Time            `json:"timestamp"`
	Compliance    float64              `json:"compliance_score"`
	Endpoints     []EndpointValidation `json:"endpoints"`
	Documentation DocumentationReport  `json:"documentation"`
	Performance   *PerformanceMetrics  `json:"performance,omitempty"`
}

// EndpointValidation represents the validation result for a single endpoint
type EndpointValidation struct {
	Path         string            `json:"path"`
	Method       string            `json:"method"`
	Status       string            `json:"status"`
	Errors       []string          `json:"errors,omitempty"`
	ResponseTime time.Duration     `json:"response_time,omitempty"`
	StatusCode   int               `json:"status_code,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
}

// DocumentationReport holds metrics related to API documentation quality
type DocumentationReport struct {
	ComplianceScore       float64        `json:"compliance_score"`
	MissingExamples       []string       `json:"missing_examples"`
	UndocumentedEndpoints []string       `json:"undocumented_endpoints"`
	ErrorResponses        map[string]int `json:"error_responses"`
	UndocumentedErrors    []string       `json:"undocumented_errors"`
}

// PerformanceMetrics holds performance testing results
type PerformanceMetrics struct {
	StartTime     time.Time     `json:"start_time"`
	EndTime       time.Time     `json:"end_time"`
	TotalRequests int64         `json:"total_requests"`
	SuccessCount  int64         `json:"success_count"`
	ErrorCount    int64         `json:"error_count"`
	ErrorRate     float64       `json:"error_rate"`
	LatencyP50    time.Duration `json:"latency_p50"`
	LatencyP95    time.Duration `json:"latency_p95"`
	LatencyP99    time.Duration `json:"latency_p99"`
}

// Validator handles API validation
type Validator struct {
	loader *openapi.Loader
	client *http.Client
	config struct {
		BaseURL string
	}
}

// NewValidator creates a new API validator
func NewValidator() *Validator {
	return &Validator{
		loader: openapi.NewLoader(),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SetBaseURL sets the base URL for API requests
func (v *Validator) SetBaseURL(baseURL string) {
	v.config.BaseURL = baseURL
}

// preprocessNullTypes recursively converts "type": "null" or ["string", "null"] to "type": "string", "nullable": true
func preprocessNullTypes(m map[string]interface{}) {
	log.Debugf("Enter preprocessNullTypes with m: %+v", m)
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			preprocessNullTypes(val)
		case []interface{}:
			if k == "type" {
				// Convert ["string", "null"] to "string" with nullable
				hasNull := false
				otherType := ""
				for _, typeVal := range val {
					if typeStr, ok := typeVal.(string); ok {
						if typeStr == "null" {
							hasNull = true
						} else {
							otherType = typeStr
						}
					}
				}
				if hasNull && otherType != "" {
					m["type"] = otherType
					m["nullable"] = true
				}
			} else {
				// Recursively process array items
				for _, item := range val {
					if sub, ok := item.(map[string]interface{}); ok {
						preprocessNullTypes(sub)
					}
				}
			}
		case string:
			if k == "type" && val == "null" {
				m["type"] = "string"
				m["nullable"] = true
			}
		}
	}
	log.Debugf("Returning from preprocessNullTypes with m: %+v", m)
}

// removeTitleFields recursively removes any 'title' field from the spec, except in the root 'info' object
func removeTitleFields(m map[string]interface{}) {
	log.Debugf("Enter removeTitleFields with m: %+v", m)
	for k, v := range m {
		if k == "paths" {
			if paths, ok := v.(map[string]interface{}); ok {
				for _, pathItem := range paths {
					if pathMap, ok := pathItem.(map[string]interface{}); ok {
						removeTitleFieldsFromPathItem(pathMap)
					}
				}
			}
		}
		// Continue recursion for all other fields
		switch val := v.(type) {
		case map[string]interface{}:
			removeTitleFields(val)
		case []interface{}:
			for _, item := range val {
				if sub, ok := item.(map[string]interface{}); ok {
					removeTitleFields(sub)
				}
			}
		}
	}
	log.Debugf("Returning from removeTitleFields with m: %+v", m)
}

// removeTitleFieldsFromPathItem removes 'title' fields from a path item and its operations
func removeTitleFieldsFromPathItem(m map[string]interface{}) {
	log.Debugf("Enter removeTitleFieldsFromPathItem with m: %+v", m)
	for k, v := range m {
		if k == "title" {
			delete(m, k)
			continue
		}
		// Remove from operations
		if k == "get" || k == "put" || k == "post" || k == "delete" || k == "options" || k == "head" || k == "patch" || k == "trace" {
			if opMap, ok := v.(map[string]interface{}); ok {
				if _, exists := opMap["title"]; exists {
					delete(opMap, "title")
				}
			}
		}
	}
	log.Debugf("Returning from removeTitleFieldsFromPathItem with m: %+v", m)
}

// stripInvalidPathItemFields removes fields not allowed in OpenAPI path item objects
func stripInvalidPathItemFields(m map[string]interface{}) {
	log.Debugf("Enter stripInvalidPathItemFields with m: %+v", m)
	if paths, ok := m["paths"].(map[string]interface{}); ok {
		for _, pathItem := range paths {
			if pathMap, ok := pathItem.(map[string]interface{}); ok {
				valid := map[string]struct{}{
					"get": {}, "put": {}, "post": {}, "delete": {}, "options": {}, "head": {}, "patch": {}, "trace": {},
					"summary": {}, "description": {}, "servers": {}, "parameters": {},
				}
				for k := range pathMap {
					_, ok := valid[k]
					if k == "title" || !ok {
						delete(pathMap, k)
					}
				}
			}
		}
	}
	log.Debugf("Returning from stripInvalidPathItemFields with m: %+v", m)
}

// stripInvalidOperationFields removes fields not allowed in OpenAPI operation objects
func stripInvalidOperationFields(m map[string]interface{}) {
	log.Debugf("Enter stripInvalidOperationFields with m: %+v", m)
	if paths, ok := m["paths"].(map[string]interface{}); ok {
		for _, pathItem := range paths {
			if pathMap, ok := pathItem.(map[string]interface{}); ok {
				for method, operation := range pathMap {
					if method == "get" || method == "put" || method == "post" || method == "delete" || method == "options" || method == "head" || method == "patch" || method == "trace" {
						if opMap, ok := operation.(map[string]interface{}); ok {
							valid := map[string]struct{}{
								"tags": {}, "summary": {}, "description": {}, "externalDocs": {}, "operationId": {}, "parameters": {}, "requestBody": {}, "responses": {}, "callbacks": {}, "deprecated": {}, "security": {}, "servers": {},
							}
							for k := range opMap {
								_, ok := valid[k]
								if k == "title" || !ok {
									delete(opMap, k)
								}
							}
						}
					}
				}
			}
		}
	}
	log.Debugf("Returning from stripInvalidOperationFields with m: %+v", m)
}

// Helper to check if a string is a valid HTTP response code or 'default'
func isValidResponseCode(key string) bool {
	log.Debugf("Enter isValidResponseCode with key: %s", key)
	if key == "default" {
		return true
	}
	if len(key) == 3 {
		for _, c := range key {
			if c < '0' || c > '9' {
				return false
			}
		}
		return true
	}
	return false
}

// cleanOpenAPIPaths ensures all path items and operations only have valid fields
// Now, only logs unexpected fields but does not delete them unless they are known to break the validator.
func cleanOpenAPIPaths(m map[string]interface{}) {
	log.Debugf("Enter cleanOpenAPIPaths with m: %+v", m)
	if paths, ok := m["paths"].(map[string]interface{}); ok {
		for path, pathItem := range paths {
			if pathMap, ok := pathItem.(map[string]interface{}); ok {
				for method, op := range pathMap {
					if opMap, ok := op.(map[string]interface{}); ok {
						// Only log unexpected fields, do not delete
						for k := range opMap {
							// List of valid operation fields from OpenAPI 3.0.3
							valid := map[string]struct{}{
								"tags": {}, "summary": {}, "description": {}, "externalDocs": {}, "operationId": {}, "parameters": {},
								"requestBody": {}, "responses": {}, "callbacks": {}, "deprecated": {}, "security": {}, "servers": {},
							}
							if _, ok := valid[k]; !ok {
								log.WithFields(logrus.Fields{
									"path":   path,
									"method": method,
									"field":  k,
								}).Warn("Unexpected field in operation; not deleting, just warning")
							}
						}
						// For responses, log unexpected fields but do not delete
						if responses, ok := opMap["responses"].(map[string]interface{}); ok {
							for k := range responses {
								if !isValidResponseCode(k) {
									log.WithFields(logrus.Fields{
										"path":   path,
										"method": method,
										"field":  k,
									}).Warn("Unexpected field in responses; not deleting, just warning")
								}
							}
						}
						// For requestBody, log unexpected fields but do not delete
						if requestBody, ok := opMap["requestBody"].(map[string]interface{}); ok {
							for k := range requestBody {
								if k != "description" && k != "content" && k != "required" {
									log.WithFields(logrus.Fields{
										"path":   path,
										"method": method,
										"field":  k,
									}).Warn("Unexpected field in requestBody; not deleting, just warning")
								}
							}
						}
					}
				}
			}
		}
	}
	log.Debugf("Returning from cleanOpenAPIPaths with m: %+v", m)
}

// Helper function to get map keys for logging
func getKeys(m map[string]interface{}) []string {
	log.Debugf("Enter getKeys with m: %+v", m)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	log.Debugf("Returning from getKeys with keys: %+v", keys)
	return keys
}

// ValidateDocumentation validates the API documentation
func (v *Validator) ValidateDocumentation(doc *openapi3.T) (*ValidationResult, error) {
	log.Debugf("Enter ValidateDocumentation with doc: %+v", doc)
	result := &ValidationResult{
		Timestamp: time.Now(),
		Documentation: DocumentationReport{
			ErrorResponses: make(map[string]int),
		},
	}

	totalEndpoints := 0
	compliantEndpoints := 0

	// Iterate over paths
	for path, pathItem := range doc.Paths.Map() {
		log.Debugf("Checking path: %s", path)
		for method, operation := range pathItem.Operations() {
			totalEndpoints++
			endpointCompliant := true
			endpointId := fmt.Sprintf("%s %s", method, path)
			log.Debugf("Checking endpoint: %s", endpointId)

			// Check for operation description
			if operation.Description == "" {
				log.Warnf("Endpoint %s missing description", endpointId)
				result.Documentation.UndocumentedEndpoints = append(result.Documentation.UndocumentedEndpoints, endpointId)
				endpointCompliant = false
			}

			// Check for examples
			if operation.RequestBody != nil && operation.RequestBody.Value != nil {
				for _, mediaType := range operation.RequestBody.Value.Content {
					if mediaType.Example == nil && len(mediaType.Examples) == 0 {
						log.Warnf("Endpoint %s missing request body example", endpointId)
						result.Documentation.MissingExamples = append(result.Documentation.MissingExamples, endpointId)
						endpointCompliant = false
					}
				}
			}

			// Check for error responses
			if operation.Responses != nil {
				for status, response := range operation.Responses.Map() {
					if strings.HasPrefix(status, "4") || strings.HasPrefix(status, "5") {
						log.Debugf("Endpoint %s has error response: %s", endpointId, status)
						result.Documentation.ErrorResponses[status]++
						if response.Value == nil || response.Value.Description == nil || *response.Value.Description == "" {
							log.Warnf("Endpoint %s error response %s undocumented", endpointId, status)
							result.Documentation.UndocumentedErrors = append(result.Documentation.UndocumentedErrors, endpointId)
							endpointCompliant = false
						}
					}
				}
			}

			if endpointCompliant {
				log.Infof("Endpoint %s is compliant", endpointId)
				compliantEndpoints++
			} else {
				log.Infof("Endpoint %s is NOT compliant", endpointId)
			}
		}
	}

	if totalEndpoints > 0 {
		result.Compliance = float64(compliantEndpoints) / float64(totalEndpoints) * 100
		result.Documentation.ComplianceScore = result.Compliance
		log.Infof("Compliance: %.2f%% (%d/%d endpoints)", result.Compliance, compliantEndpoints, totalEndpoints)
	} else {
		log.Warn("No endpoints found in OpenAPI document")
	}

	log.Debugf("Validation result: %+v", result)
	log.Debugf("Returning from ValidateDocumentation with result: %+v, error: %v", result, nil)
	return result, nil
}

// generateParameterValue generates an example value for an OpenAPI parameter schema.
func (v *Validator) generateParameterValue(schema *openapi3.Schema) (interface{}, error) {
	log.Debugf("Enter generateParameterValue with schema: %+v", schema)
	if schema == nil {
		return nil, fmt.Errorf("schema is nil")
	}

	// Use existing GetExampleValues for basic types and objects/arrays
	exampleMap := v.loader.GetExampleValues(schema)
	if value, ok := exampleMap["value"]; ok {
		return value, nil
	}

	// Fallback or handle specific cases not covered by GetExampleValues
	switch strings.ToLower(schema.Type) {
	case "string":
		if schema.Example != nil {
			return schema.Example, nil
		}
		if len(schema.Enum) > 0 {
			return schema.Enum[0], nil
		}
		return "string_value", nil
	case "number":
		if schema.Example != nil {
			return schema.Example, nil
		}
		return 123.45, nil
	case "integer":
		if schema.Example != nil {
			return schema.Example, nil
		}
		return 123, nil
	case "boolean":
		if schema.Example != nil {
			return schema.Example, nil
		}
		return true, nil
	case "array":
		if schema.Example != nil {
			return schema.Example, nil
		}
		if schema.Items != nil && schema.Items.Value != nil {
			item, err := v.generateParameterValue(schema.Items.Value)
			if err != nil {
				return nil, fmt.Errorf("failed to generate array item: %w", err)
			}
			return []interface{}{item}, nil
		}
		return []interface{}{}, nil
	case "object":
		if schema.Example != nil {
			return schema.Example, nil
		}
		// Recursively generate properties for object
		obj := make(map[string]interface{})
		if schema.Properties != nil {
			for name, propRef := range schema.Properties {
				if propRef != nil && propRef.Value != nil {
					propValue, err := v.generateParameterValue(propRef.Value)
					if err != nil {
						// Log the error but continue with other properties
						log.WithError(err).Warnf("Failed to generate value for property '%s'", name)
						continue
					}
					obj[name] = propValue
				}
			}
		}
		return obj, nil
	default:
		return nil, fmt.Errorf("unsupported schema type for parameter generation: %s", schema.Type)
	}
}

// generateRequestBody generates an example request body based on the OpenAPI content object.
func (v *Validator) generateRequestBody(content openapi3.Content) (io.Reader, string, error) {
	log.Debugf("Enter generateRequestBody with content: %+v", content)
	// Prioritize JSON content
	if mediaType, ok := content["application/json"]; ok && mediaType.Schema != nil && mediaType.Schema.Value != nil {
		exampleData, err := v.generateParameterValue(mediaType.Schema.Value)
		if err != nil {
			return nil, "", fmt.Errorf("failed to generate example data for JSON body: %w", err)
		}
		jsonData, err := json.Marshal(exampleData)
		if err != nil {
			return nil, "", fmt.Errorf("failed to marshal example data to JSON: %w", err)
		}
		return bytes.NewReader(jsonData), "application/json", nil
	}

	// Handle other content types like form data if needed

	return nil, "", fmt.Errorf("unsupported or missing request body content type for generation")
}

// ValidateEndpoints tests the API endpoints
func (v *Validator) ValidateEndpoints(ctx context.Context, doc *openapi3.T, baseURL string) (*ValidationResult, error) {
	log.Debugf("Enter ValidateEndpoints with doc: %+v, baseURL: %s", doc, baseURL)
	result := &ValidationResult{
		Timestamp: time.Now(),
	}

	total := 0
	passed := 0
	failed := 0
	authFailed := 0
	serverError := 0
	clientError := 0
	undocumented := 0

	for path, pathItem := range doc.Paths.Map() {
		log.Debugf("Validating path: %s", path)
		// Replace path parameters with generated values for URL construction
		templatedPath := path
		pathParams := make(map[string]interface{})
		for _, paramRef := range pathItem.Parameters {
			if paramRef != nil && paramRef.Value != nil && paramRef.Value.In == openapi3.ParameterInPath {
				paramName := paramRef.Value.Name
				paramValue, err := v.generateParameterValue(paramRef.Value.Schema.Value)
				if err != nil {
					log.WithError(err).Warnf("Failed to generate value for path parameter '%s' in 'pathItem' %s", paramName, path)
					paramValue = fmt.Sprintf("{%s}", paramName)
				}
				log.Debugf("Path parameter: %s = %v", paramName, paramValue)
				templatedPath = strings.ReplaceAll(templatedPath, fmt.Sprintf("{%s}", paramName), fmt.Sprintf("%v", paramValue))
				pathParams[paramName] = paramValue
			}
		}

		for method, operation := range pathItem.Operations() {
			endpointId := fmt.Sprintf("%s %s", method, path)
			log.Debugf("Validating endpoint: %s", endpointId)
			validation := EndpointValidation{
				Path:   path,
				Method: method,
				Status: "pending",
			}

			url := fmt.Sprintf("%s%s", baseURL, templatedPath)
			log.Debugf("Request URL: %s", url)

			// Add query parameters
			req, err := http.NewRequestWithContext(ctx, method, url, nil)
			if err != nil {
				log.WithError(err).Errorf("Failed to create request for endpoint %s", endpointId)
				validation.Status = "failed"
				validation.Errors = append(validation.Errors, fmt.Sprintf("failed to create request: %v", err))
				failed++
				log.Warnf("Endpoint %s failed: %v", endpointId, err)
				result.Endpoints = append(result.Endpoints, validation)
				continue
			}

			queryParams := req.URL.Query()
			for _, paramRef := range operation.Parameters {
				if paramRef != nil && paramRef.Value != nil {
					param := paramRef.Value
					if param.In == openapi3.ParameterInQuery {
						paramValue, err := v.generateParameterValue(param.Schema.Value)
						if err != nil {
							log.WithError(err).Warnf("Failed to generate value for query parameter '%s' in '%s %s'", param.Name, method, path)
							continue
						}
						log.Debugf("Query parameter: %s = %v", param.Name, paramValue)
						queryParams.Add(param.Name, fmt.Sprintf("%v", paramValue))
					} else if param.In == openapi3.ParameterInHeader {
						paramValue, err := v.generateParameterValue(param.Schema.Value)
						if err != nil {
							log.WithError(err).Warnf("Failed to generate value for header parameter '%s' in '%s %s'", param.Name, method, path)
							continue
						}
						log.Debugf("Header parameter: %s = %v", param.Name, paramValue)
						req.Header.Add(param.Name, fmt.Sprintf("%v", paramValue))
					} else if param.In == openapi3.ParameterInCookie {
						log.Warnf("Cookie parameters are not fully supported for validation: '%s' in '%s %s'", param.Name, method, path)
						continue
					} else if param.In == openapi3.ParameterInPath {
						// Path parameters are handled when building the URL.
						continue
					} else {
						log.Warnf("Unsupported parameter location '%s' for parameter '%s' in '%s %s'", param.In, param.Name, method, path)
					}
				}
			}
			req.URL.RawQuery = queryParams.Encode()
			log.Debugf("Final request URL: %s", req.URL.String())

			// Add request body
			if operation.RequestBody != nil && operation.RequestBody.Value != nil {
				body, contentType, err := v.generateRequestBody(operation.RequestBody.Value.Content)
				if err != nil {
					log.WithError(err).Warnf("Failed to generate request body for '%s %s'", method, path)
					// Continue without a request body
				} else {
					log.Debugf("Generated request body for '%s %s' with content-type %s", method, path, contentType)
					req.Body = ioutil.NopCloser(body)
					req.Header.Set("Content-Type", contentType)
				}
			}

			// Send request and measure time
			start := time.Now()
			resp, err := v.client.Do(req)
			validation.ResponseTime = time.Since(start)
			total++
			if err != nil {
				validation.Status = "failed"
				validation.Errors = append(validation.Errors, fmt.Sprintf("request failed: %v", err))
				failed++
				log.Warnf("Endpoint %s failed: %v", endpointId, err)
				result.Endpoints = append(result.Endpoints, validation)
				continue
			}
			defer resp.Body.Close()

			// Validate response
			validation.StatusCode = resp.StatusCode
			validation.Headers = make(map[string]string)
			for k, v := range resp.Header {
				if len(v) > 0 {
					validation.Headers[k] = v[0]
				}
			}

			statusCodeStr := fmt.Sprintf("%d", resp.StatusCode)
			if resp.StatusCode == 401 || resp.StatusCode == 403 {
				validation.Status = "auth_failed"
				validation.Errors = append(validation.Errors, fmt.Sprintf("authentication failed: status %d", resp.StatusCode))
				authFailed++
				log.Warnf("Endpoint %s authentication failed (status %d)", endpointId, resp.StatusCode)
			} else if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				validation.Status = "server_error"
				validation.Errors = append(validation.Errors, fmt.Sprintf("server error: status %d", resp.StatusCode))
				serverError++
				log.Warnf("Endpoint %s server error (status %d)", endpointId, resp.StatusCode)
			} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
				validation.Status = "client_error"
				validation.Errors = append(validation.Errors, fmt.Sprintf("client error: status %d", resp.StatusCode))
				clientError++
				log.Warnf("Endpoint %s client error (status %d)", endpointId, resp.StatusCode)
			} else if operation.Responses != nil {
				if _, ok := operation.Responses.Map()[statusCodeStr]; ok {
					validation.Status = "success"
					passed++
					log.Infof("Endpoint %s passed (status %d)", endpointId, resp.StatusCode)
				} else {
					validation.Status = "undocumented"
					validation.Errors = append(validation.Errors, fmt.Sprintf("response status code %d not documented in OpenAPI spec", resp.StatusCode))
					undocumented++
					log.Warnf("Endpoint %s returned undocumented status code %d", endpointId, resp.StatusCode)
				}
			} else {
				validation.Status = "undocumented"
				validation.Errors = append(validation.Errors, "no responses documented in OpenAPI spec")
				undocumented++
				log.Warnf("Endpoint %s has no responses documented in OpenAPI spec", endpointId)
			}

			result.Endpoints = append(result.Endpoints, validation)
		}
	}

	log.Infof("Validation summary: total=%d, passed=%d, failed=%d, auth_failed=%d, server_error=%d, client_error=%d, undocumented=%d", total, passed, failed, authFailed, serverError, clientError, undocumented)
	// Optionally, add a summary to result (if struct allows)
	// result.Summary = ...

	log.Debugf("Returning from ValidateEndpoints with result: %+v, error: %v", result, nil)
	return result, nil
}

// RunPerformanceTests runs load tests against the API
func (v *Validator) RunPerformanceTests(targets []vegeta.Target, rate float64, duration time.Duration) (*ValidationResult, error) {
	log.Debugf("Enter RunPerformanceTests with targets: %+v, rate: %.2f, duration: %s", targets, rate, duration)
	metrics := &PerformanceMetrics{
		StartTime: time.Now(),
	}

	attacker := vegeta.NewAttacker()
	pacer := vegeta.Rate{Freq: int(rate), Per: time.Second}
	targeter := vegeta.NewStaticTargeter(targets...)
	vegetaMetrics := &vegeta.Metrics{}

	for res := range attacker.Attack(targeter, pacer, duration, "Load Test") {
		vegetaMetrics.Add(res)
	}
	vegetaMetrics.Close()

	metrics.EndTime = time.Now()
	metrics.TotalRequests = int64(vegetaMetrics.Requests)
	metrics.SuccessCount = int64(float64(vegetaMetrics.Requests) * vegetaMetrics.Success)
	metrics.ErrorCount = metrics.TotalRequests - metrics.SuccessCount
	metrics.ErrorRate = 1 - vegetaMetrics.Success
	metrics.LatencyP50 = vegetaMetrics.Latencies.P50
	metrics.LatencyP95 = vegetaMetrics.Latencies.P95
	metrics.LatencyP99 = vegetaMetrics.Latencies.P99

	log.Debugf("Performance metrics: %+v", metrics)

	return &ValidationResult{
		Timestamp:   time.Now(),
		Performance: metrics,
	}, nil
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[validation] Logger set to DEBUG (verbose) mode")
}

// RunValidation runs only OpenAPI/documentation validation checks (P001, P003, P004, P005)
func RunValidation(ctx context.Context, config ValidatorConfig) (*ValidationReport, error) {
	log.Debugf("Starting RunValidation with config: %+v", config)
	loader := openapi.NewLoader()
	if err := loader.LoadFromFileOrURL(config.SpecPath); err != nil {
		log.WithError(err).Errorf("Failed to load OpenAPI spec from %s", config.SpecPath)
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}
	doc := loader.GetDocument()
	if doc == nil {
		log.Error("Failed to get OpenAPI document after loading")
		return nil, fmt.Errorf("failed to get OpenAPI document")
	}

	principles := []string{"P001", "P003", "P004", "P005"}
	var results []PrincipleResult
	for _, pid := range principles {
		var res PrincipleResult
		log.Debugf("Validating principle: %s", pid)
		switch pid {
		case "P001":
			res = ValidateOpenAPICompliance(doc)
		case "P003":
			res = ValidateErrorDocumentation(doc)
		case "P004":
			res = ValidateRequestValidation(doc)
		case "P005":
			res = ValidateAuthentication(doc)
		}
		log.Debugf("Result for principle %s: %+v", pid, res)
		results = append(results, res)
	}

	report := &ValidationReport{
		Version:     config.Version,
		Environment: config.Environment,
		Timestamp:   time.Now(),
		Principles:  results,
	}
	log.Debugf("Final validation report: %+v", report)
	return report, nil
}
