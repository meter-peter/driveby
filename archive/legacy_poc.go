package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// DocumentationReport holds metrics related to API documentation quality
type DocumentationReport struct {
	ComplianceScore       float64        `json:"compliance_score"`
	MissingExamples       int            `json:"missing_examples"`
	UndocumentedEndpoints []string       `json:"undocumented_endpoints"`
	ErrorResponses        map[string]int `json:"error_responses"` // count per status code
}

// ValidateAPIDocumentation performs comprehensive validation of the OpenAPI documentation
func ValidateAPIDocumentation(doc *openapi3.T) (DocumentationReport, []error) {
	report := DocumentationReport{
		ErrorResponses: make(map[string]int),
	}
	var errors []error

	totalEndpoints := 0
	compliantEndpoints := 0

	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			totalEndpoints++
			endpointCompliant := true
			endpointId := fmt.Sprintf("%s %s", method, path)

			// Check response documentation
			if len(operation.Responses) == 0 {
				errors = append(errors, fmt.Errorf("%s: missing response documentation", endpointId))
				endpointCompliant = false
			}

			// Check all response status codes have documentation
			for statusCode, response := range operation.Responses {
				if response.Value.Description == nil || *response.Value.Description == "" {
					errors = append(errors, fmt.Errorf("%s: missing description for status code %s",
						endpointId, statusCode))
					endpointCompliant = false
				}

				// Count error responses
				if strings.HasPrefix(statusCode, "4") || strings.HasPrefix(statusCode, "5") {
					report.ErrorResponses[statusCode]++
				}

				// Check response examples
				if response.Value.Content != nil {
					jsonContent := response.Value.Content.Get("application/json")
					if jsonContent != nil && jsonContent.Example == nil && len(jsonContent.Examples) == 0 {
						errors = append(errors, fmt.Errorf("%s: missing examples for response with status %s",
							endpointId, statusCode))
						report.MissingExamples++
						endpointCompliant = false
					}
				}
			}

			// Check parameters
			for _, param := range operation.Parameters {
				if param.Value.Required && param.Value.Example == nil {
					errors = append(errors, fmt.Errorf("%s: missing example for required parameter '%s'",
						endpointId, param.Value.Name))
					report.MissingExamples++
					endpointCompliant = false
				}
			}

			// Check request body examples
			if operation.RequestBody != nil && operation.RequestBody.Value.Required {
				hasExamples := false
				for contentType, content := range operation.RequestBody.Value.Content {
					if contentType == "application/json" {
						if content.Example != nil || len(content.Examples) > 0 {
							hasExamples = true
							break
						} else if content.Schema != nil && content.Schema.Value != nil && content.Schema.Value.Example != nil {
							hasExamples = true
							break
						}
					}
				}

				if !hasExamples {
					errors = append(errors, fmt.Errorf("%s: missing request body examples", endpointId))
					report.MissingExamples++
					endpointCompliant = false
				}
			}

			// Check metadata (this is a warning, not blocking)
			if operation.Summary == "" || len(operation.Tags) == 0 {
				logrus.Warnf("%s: Missing metadata (summary or tags)", endpointId)
				// Don't fail compliance for this, just warn
			}

			if endpointCompliant {
				compliantEndpoints++
			} else {
				report.UndocumentedEndpoints = append(report.UndocumentedEndpoints, endpointId)
			}
		}
	}

	if totalEndpoints > 0 {
		report.ComplianceScore = float64(compliantEndpoints) / float64(totalEndpoints) * 100
	}

	return report, errors
}

// Configuration with default values, can be overridden with environment variables
var (
	apiHost        = getEnv("API_HOST", "localhost")
	apiPort        = getEnv("API_PORT", "8080")
	apiBasePath    = getEnv("API_BASE_PATH", "") // Empty by default, set to "/api/v1" if needed
	openAPIPath    = getEnv("OPENAPI_PATH", "/swagger/doc.json")
	requestRate    = getEnvInt("REQUEST_RATE", 1)    // Requests per second
	testDuration   = getEnvInt("TEST_DURATION", 10)  // Test duration in seconds
	requestTimeout = getEnvInt("REQUEST_TIMEOUT", 5) // Request timeout in seconds
)

// Helper function to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// Helper function to get integer environment variables with defaults
func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, fmt.Sprintf("%d", defaultValue))
	value := defaultValue
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		logrus.Warnf("Invalid value for %s: %s, using default: %d", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

type EndpointTest struct {
	Method   string
	Path     string
	Params   url.Values
	Body     interface{}
	Examples map[string]interface{}
}

// Validation configuration
var (
	validationThreshold = getEnvFloat("VALIDATION_THRESHOLD", 95.0) // Minimum compliance score to pass validation
	failOnValidation    = getEnvBool("FAIL_ON_VALIDATION", true)    // Whether to exit if validation fails
)

func getEnvFloat(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, fmt.Sprintf("%f", defaultValue))
	value := defaultValue
	_, err := fmt.Sscanf(valueStr, "%f", &value)
	if err != nil {
		logrus.Warnf("Invalid value for %s: %s, using default: %f", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, fmt.Sprintf("%t", defaultValue))
	if valueStr == "true" || valueStr == "1" || valueStr == "yes" {
		return true
	} else if valueStr == "false" || valueStr == "0" || valueStr == "no" {
		return false
	}
	return defaultValue
}

// WriteValidationReport writes a detailed validation report to a markdown file
func WriteValidationReport(report DocumentationReport, validationErrors []error) error {
	f, err := os.Create("validation-report.md")
	if err != nil {
		return err
	}
	defer f.Close()

	// Write report header
	_, err = io.WriteString(f, fmt.Sprintf(`
## API Documentation Validation Report

**Compliance Score:** %.2f%%  
**Missing Examples:** %d  
**Undocumented Error Responses:** %d

`,
		report.ComplianceScore,
		report.MissingExamples,
		len(report.ErrorResponses),
	))
	if err != nil {
		return err
	}

	// Write critical issues
	if len(validationErrors) > 0 {
		_, err = io.WriteString(f, "### Critical Issues:\n")
		if err != nil {
			return err
		}

		for _, errMsg := range validationErrors {
			_, err = io.WriteString(f, fmt.Sprintf("- %s\n", errMsg.Error()))
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(f, "\n")
		if err != nil {
			return err
		}
	}

	// Write undocumented endpoints
	if len(report.UndocumentedEndpoints) > 0 {
		_, err = io.WriteString(f, "### Undocumented Endpoints:\n")
		if err != nil {
			return err
		}

		for _, endpoint := range report.UndocumentedEndpoints {
			_, err = io.WriteString(f, fmt.Sprintf("- `%s`\n", endpoint))
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(f, "\n")
		if err != nil {
			return err
		}
	}

	// Write validation result
	result := "✅ Validation Passed"
	if report.ComplianceScore < validationThreshold {
		result = "❌ Validation Failed"
	}

	_, err = io.WriteString(f, fmt.Sprintf(`
**Validation Threshold:** %.2f%%  
**Result:** %s
`, validationThreshold, result))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Define command line flags
	urlFlag := flag.String("u", "", "URL to the OpenAPI specification (e.g., https://example.com/openapi.json)")
	validateOnlyFlag := flag.Bool("validate", false, "Run only validation (skip load testing)")

	// Parse command line flags
	flag.Parse()

	// Set up logging with more verbose output
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel) // Enable detailed logging

	// Determine OpenAPI URL to use
	var openAPIURL string
	if *urlFlag != "" {
		// Use the URL provided via command line flag
		openAPIURL = *urlFlag
		logrus.Infof("Using OpenAPI spec from command line URL: %s", openAPIURL)
	} else {
		// Fall back to the configuration from environment variables
		openAPIURL = fmt.Sprintf("http://%s:%s%s", apiHost, apiPort, openAPIPath)
		logrus.Infof("Using OpenAPI spec from environment settings: %s", openAPIURL)
	}

	if !*validateOnlyFlag {
		logrus.Infof("Load test configuration: %d req/sec for %d seconds, timeout: %ds",
			requestRate, testDuration, requestTimeout)
	}

	// 1. Fetch OpenAPI spec
	logrus.Info("Fetching OpenAPI specification...")
	doc, err := fetchOpenAPI(openAPIURL)
	if err != nil {
		logrus.Fatalf("Failed to fetch OpenAPI: %v", err)
	}
	logrus.Infof("Successfully fetched OpenAPI spec with %d paths", len(doc.Paths.Map()))

	// 2. Validate API documentation
	logrus.Info("Validating API documentation...")
	report, validationErrors := ValidateAPIDocumentation(doc)

	// Write validation report
	if err := WriteValidationReport(report, validationErrors); err != nil {
		logrus.Errorf("Failed to write validation report: %v", err)
	}

	// Log validation results
	logrus.Infof("Documentation Compliance Score: %.2f%% (threshold: %.2f%%)",
		report.ComplianceScore, validationThreshold)

	if len(validationErrors) > 0 {
		logrus.Warn("Documentation validation found issues:")
		for _, err := range validationErrors {
			logrus.Warnf("- %v", err)
		}
	}

	// Check if we should fail on validation issues
	if failOnValidation && report.ComplianceScore < validationThreshold {
		logrus.Fatalf("Documentation validation failed. Score %.2f%% is below threshold %.2f%%",
			report.ComplianceScore, validationThreshold)
	}

	// Exit early if only validation was requested
	if *validateOnlyFlag {
		logrus.Info("Validation completed. Skipping load testing as requested.")
		return
	}

	// 3. Discover endpoints for load testing
	logrus.Info("Discovering endpoints for load testing...")
	endpoints, discoveryErrors := discoverEndpoints(doc)
	if len(discoveryErrors) > 0 {
		for _, err := range discoveryErrors {
			logrus.Errorf("Endpoint discovery issue: %v", err)
		}
		// Log a warning but continue with the test
		logrus.Warn("Some endpoints could not be processed due to missing examples or defaults. Continuing with the rest.")
	}

	// Log the list of endpoints being tested with more details
	logrus.Info("Endpoints to be tested:")
	for i, endpoint := range endpoints {
		logrus.Infof("[%d/%d] %s %s", i+1, len(endpoints), endpoint.Method, endpoint.Path)

		// Log parameters
		if len(endpoint.Examples) > 0 {
			logrus.Info("  Parameters:")
			for name, value := range endpoint.Examples {
				logrus.Infof("    %s: %v", name, value)
			}
		}

		// Log request body
		if endpoint.Body != nil {
			bodyBytes, _ := json.MarshalIndent(endpoint.Body, "", "  ")
			logrus.Infof("  Request Body: %s", string(bodyBytes))
		}
	}

	// 3. Create Vegeta targets
	targets := createTargets(endpoints)

	// 4. Configure load test
	logrus.Info("Configuring load test...")
	rate := vegeta.Rate{Freq: requestRate, Per: time.Second}
	duration := time.Duration(testDuration) * time.Second
	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker(vegeta.Timeout(time.Duration(requestTimeout) * time.Second))

	// 5. Run test with detailed logging of each request/response
	logrus.Info("Starting load test...")

	// Create a metrics collector
	var metrics vegeta.Metrics
	resultChan := attacker.Attack(targeter, rate, duration, "API Load Test")

	// Track stats during test for real-time feedback
	requestCount := 0
	successCount := 0
	failureCount := 0
	statusCodes := make(map[int]int)

	// Process each request result with detailed logging
	for res := range resultChan {
		metrics.Add(res)
		requestCount++

		// Extract status code for tracking and convert to int for map key
		statusCode := res.Code
		statusCodes[int(statusCode)]++

		// Prepare a detailed log message for each request
		if res.Error != "" {
			// Error occurred during request
			failureCount++
			logrus.Warnf("Request #%d FAILED: %s %s - Error: %s",
				requestCount, res.Method, res.URL, res.Error)
		} else if statusCode >= 400 {
			// Request completed but returned an error status
			failureCount++
			logrus.Warnf("Request #%d: %s %s - Status: %d [FAILED]",
				requestCount, res.Method, res.URL, statusCode)

			// Log response body for failed requests
			responseBody := string(res.Body)
			if len(responseBody) > 0 {
				// Truncate long responses
				if len(responseBody) > 500 {
					responseBody = responseBody[:500] + "... [truncated]"
				}
				logrus.Warnf("Response body: %s", responseBody)
			}
		} else {
			// Successful request
			successCount++
			logrus.Infof("Request #%d: %s %s - Status: %d, Latency: %s",
				requestCount, res.Method, res.URL, statusCode, res.Latency)
		}

		// Show running stats every few requests
		if requestCount%1 == 0 || requestCount == int(requestRate*testDuration) {
			logrus.Infof("Progress: %d/%d requests, %d successes, %d failures",
				requestCount, int(requestRate*testDuration), successCount, failureCount)

			// Print status code distribution
			logrus.Info("Current status codes:")
			for code, count := range statusCodes {
				logrus.Infof("  %d: %d requests (%.1f%%)",
					code, count, float64(count)/float64(requestCount)*100)
			}
		}
	}
	metrics.Close()

	// 6. Generate report
	logrus.Info("Load test completed. Generating detailed report...")
	printReport(&metrics)

	// 7. Optionally, open a GitHub issue with the discovery errors
	if len(discoveryErrors) > 0 {
		openGitHubIssue(discoveryErrors)
	}

	logrus.Info("Load test process complete")
}

func openGitHubIssue(errors []error) {
	// Replace with your GitHub repository details
	repoOwner := "your-org"
	repoName := "your-repo"
	githubToken := os.Getenv("GITHUB_TOKEN") // Ensure this is set in your environment

	if githubToken == "" {
		logrus.Warn("GITHUB_TOKEN not set. Skipping GitHub issue creation.")
		return
	}

	// Create the issue body
	var issueBody strings.Builder
	issueBody.WriteString("The following issues were found during endpoint discovery:\n\n")
	for _, err := range errors {
		issueBody.WriteString(fmt.Sprintf("- %s\n", err.Error()))
	}

	// Create the GitHub issue
	client := &http.Client{}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", repoOwner, repoName)
	payload := map[string]string{
		"title": "Endpoint Discovery Issues",
		"body":  issueBody.String(),
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logrus.Errorf("Failed to create GitHub issue request: %v", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Failed to create GitHub issue: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		logrus.Errorf("Failed to create GitHub issue. Status: %s", resp.Status)
		return
	}

	logrus.Info("GitHub issue created successfully.")
}

func fetchOpenAPI(url string) (*openapi3.T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spec map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&spec); err != nil {
		return nil, err
	}

	loader := openapi3.NewLoader()
	return loader.LoadFromData(func() []byte {
		b, _ := json.Marshal(spec)
		return b
	}())
}

// Helper function to generate an example value based on schema
func generateExampleFromSchema(schema *openapi3.Schema, paramName string) interface{} {
	// Check for enum values first
	if schema.Enum != nil && len(schema.Enum) > 0 {
		return schema.Enum[0] // Return first enum value
	}

	// Work with schema.Type directly as a string
	typeStr := schema.Type // It's already a string

	// Check type and generate examples
	if strings.Contains(typeStr, "string") {
		if schema.Format != "" {
			switch schema.Format {
			case "date":
				return "2023-01-01"
			case "date-time":
				return "2023-01-01T00:00:00Z"
			case "email":
				return "user@example.com"
			case "uuid":
				return "f7cfc49d-824b-4728-a4c4-45e5901e3d42"
			}
		}
		return "example-string"
	} else if strings.Contains(typeStr, "integer") {
		return 1
	} else if strings.Contains(typeStr, "number") {
		return 1.0
	} else if strings.Contains(typeStr, "boolean") {
		return true
	} else if strings.Contains(typeStr, "array") {
		return []interface{}{}
	} else if strings.Contains(typeStr, "object") {
		return map[string]interface{}{}
	}

	// Check schema structure for implicit types
	if len(schema.Properties) > 0 {
		return map[string]interface{}{} // Object
	} else if schema.Items != nil {
		return []interface{}{} // Array
	}

	// Default fallback
	return "example-value"
}

func discoverEndpoints(doc *openapi3.T) ([]EndpointTest, []error) {
	var endpoints []EndpointTest
	var errors []error

	// Iterate directly over the paths
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			test := EndpointTest{
				Method:   method,
				Path:     path,
				Examples: make(map[string]interface{}),
			}

			logrus.Infof("Processing endpoint: %s %s", method, path)

			// Extract parameters
			for _, param := range operation.Parameters {
				if param.Value.Required {
					paramName := param.Value.Name
					// Try to find an example in several possible locations
					if param.Value.Example != nil {
						test.Examples[paramName] = param.Value.Example
						logrus.Infof("Found direct example for parameter '%s': %v", paramName, param.Value.Example)
					} else if param.Value.Schema != nil && param.Value.Schema.Value.Example != nil {
						test.Examples[paramName] = param.Value.Schema.Value.Example
						logrus.Infof("Found schema example for parameter '%s': %v", paramName, param.Value.Schema.Value.Example)
					} else if param.Value.Schema != nil && param.Value.Schema.Value.Default != nil {
						test.Examples[paramName] = param.Value.Schema.Value.Default
						logrus.Infof("Found default value for parameter '%s': %v", paramName, param.Value.Schema.Value.Default)
					} else if operation.Extensions != nil {
						// Some APIs use the x-examples extension
						if examples, ok := operation.Extensions["x-examples"].(map[string]interface{}); ok {
							if paramExample, ok := examples[paramName]; ok {
								test.Examples[paramName] = paramExample
								logrus.Infof("Found x-examples value for parameter '%s': %v", paramName, paramExample)
							}
						}
					} else {
						// Try to generate a reasonable example based on schema properties
						if param.Value.Schema != nil && param.Value.Schema.Value != nil {
							// Use our helper function to generate examples
							example := generateExampleFromSchema(param.Value.Schema.Value, paramName)
							test.Examples[paramName] = example
							logrus.Infof("Generated example value for parameter '%s': %v", paramName, example)
						} else {
							// Log the error but don't fail
							err := fmt.Errorf("missing example for required parameter '%s' in %s %s",
								paramName, method, path)
							errors = append(errors, err)
							logrus.Warnf("Skipping parameter '%s' in %s %s: %v", paramName, method, path, err)
							continue // Skip this parameter but continue processing the endpoint
						}
					}
				}
			}

			// Check request body examples (with enhanced support for different example formats)
			if operation.RequestBody != nil {
				content := operation.RequestBody.Value.Content.Get("application/json")
				if content != nil {
					// First try direct example
					if content.Example != nil {
						test.Body = content.Example
						logrus.Infof("Found direct request body example for %s %s", method, path)
					} else if len(content.Examples) > 0 {
						// Then try the examples map - take the first one we find
						for name, example := range content.Examples {
							if example.Value != nil && example.Value.Value != nil {
								test.Body = example.Value.Value
								logrus.Infof("Found named example '%s' for request body in %s %s", name, method, path)
								break
							}
						}
					} else if content.Schema != nil && content.Schema.Value != nil {
						// Then try schema example
						if content.Schema.Value.Example != nil {
							test.Body = content.Schema.Value.Example
							logrus.Infof("Found schema example for request body in %s %s", method, path)
						}
					}
				}
			}

			endpoints = append(endpoints, test)
			logrus.Infof("Added endpoint for testing: %s %s", method, path)
		}
	}

	logrus.Infof("Total endpoints discovered for testing: %d", len(endpoints))
	return endpoints, errors
}
func createTargets(endpoints []EndpointTest) []vegeta.Target {
	var targets []vegeta.Target

	logrus.Info("Creating API request targets...")

	for _, endpoint := range endpoints {
		// Create base target with common headers
		target := vegeta.Target{
			Method: endpoint.Method,
			URL:    buildURL(endpoint),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
				"Accept":       []string{"application/json"},
				"User-Agent":   []string{"LoadTester/1.0"},
			},
		}

		// Check for API key and add it as a header if found
		// This fixes the 422 errors by correctly providing the API key in the header
		if apiKey, exists := endpoint.Examples["api-key"]; exists {
			target.Header.Set("api-key", fmt.Sprintf("%v", apiKey))
			logrus.Infof("  Adding API key header for %s %s", endpoint.Method, endpoint.Path)

			// Remove from URL parameters since it's now in headers
			delete(endpoint.Examples, "api-key")
		}

		// Add request body if available with proper error handling
		if endpoint.Body != nil {
			body, err := json.Marshal(endpoint.Body)
			if err != nil {
				logrus.Warnf("Failed to marshal request body for %s %s: %v", endpoint.Method, endpoint.Path, err)
			} else {
				target.Body = body
				logrus.Debugf("Target %s %s has body: %s", endpoint.Method, target.URL, string(body))
			}
		}

		// Log detailed target info
		logrus.Infof("Created target: %s %s", target.Method, target.URL)
		if len(endpoint.Examples) > 0 {
			logrus.Debugf("  Using URL parameters: %v", endpoint.Examples)
		}
		logrus.Debugf("  Using headers: %v", target.Header)
		if target.Body != nil {
			bodyPreview := string(target.Body)
			if len(bodyPreview) > 100 {
				bodyPreview = bodyPreview[:100] + "..."
			}
			logrus.Debugf("  Request body: %s", bodyPreview)
		}

		targets = append(targets, target)
	}

	logrus.Infof("Total targets created: %d", len(targets))
	return targets
}

func buildURL(endpoint EndpointTest) string {
	u := endpoint.Path
	for param, value := range endpoint.Examples {
		u = strings.ReplaceAll(u, fmt.Sprintf("{%s}", param), fmt.Sprintf("%v", value))
	}

	// Use the configured API host, port and base path
	return fmt.Sprintf("http://%s:%s%s%s", apiHost, apiPort, apiBasePath, u)
}

func printReport(metrics *vegeta.Metrics) {
	// Add detailed error breakdown
	errorCounts := make(map[string]int)
	statusCodeCounts := make(map[string]int)

	// Count error types
	for _, err := range metrics.Errors {
		errorCounts[err]++
	}

	// Count status codes
	for code, count := range metrics.StatusCodes {
		statusCodeCounts[code] = count
	}

	// Print detailed console report
	fmt.Printf(`
Load Test Report
================
Date: %s

Requests: %d
Success Rate: %.2f%%
Latency (p95): %s

Status Code Breakdown:
`,
		time.Now().Format(time.RFC3339),
		metrics.Requests,
		metrics.Success*100,
		metrics.Latencies.P95,
	)

	// Print status code breakdown
	for code, count := range statusCodeCounts {
		fmt.Printf("- %s: %d requests (%.2f%%)\n",
			code, count, float64(count)/float64(metrics.Requests)*100)
	}

	// Print error breakdown
	fmt.Println("\nDetailed Errors:")
	if len(metrics.Errors) == 0 {
		fmt.Println("- No errors")
	} else {
		for err, count := range errorCounts {
			fmt.Printf("- %s: %d occurrences (%.2f%%)\n",
				err, count, float64(count)/float64(metrics.Requests)*100)
		}
	}

	// Create a markdown report
	f, err := os.Create("loadtest-report.md")
	if err != nil {
		logrus.Errorf("Failed to create report file: %v", err)
		return
	}
	defer f.Close()

	// Write the report to the file with detailed information
	_, err = io.WriteString(f, fmt.Sprintf(`
# Load Test Report

## Summary
- **Date**: %s
- **Requests**: %d
- **Success Rate**: %.2f%%
- **Latency (p95)**: %s

## Status Code Breakdown
`,
		time.Now().Format(time.RFC3339),
		metrics.Requests,
		metrics.Success*100,
		metrics.Latencies.P95,
	))
	if err != nil {
		logrus.Errorf("Failed to write report: %v", err)
		return
	}

	// Add status code breakdown to the report
	for code, count := range statusCodeCounts {
		_, err = io.WriteString(f, fmt.Sprintf("- **%s**: %d requests (%.2f%%)\n",
			code, count, float64(count)/float64(metrics.Requests)*100))
		if err != nil {
			logrus.Errorf("Failed to write report: %v", err)
			return
		}
	}

	// Add error breakdown to the report
	_, err = io.WriteString(f, "\n## Detailed Errors\n")
	if err != nil {
		logrus.Errorf("Failed to write report: %v", err)
		return
	}

	if len(metrics.Errors) == 0 {
		_, err = io.WriteString(f, "- No errors encountered\n")
		if err != nil {
			logrus.Errorf("Failed to write report: %v", err)
		}
	} else {
		for err, count := range errorCounts {
			_, err2 := io.WriteString(f, fmt.Sprintf("- **%s**: %d occurrences (%.2f%%)\n",
				err, count, float64(count)/float64(metrics.Requests)*100))
			if err2 != nil {
				logrus.Errorf("Failed to write report: %v", err2)
				return
			}
		}
	}

	// Add latency statistics
	_, err = io.WriteString(f, "\n## Latency Statistics\n")
	if err != nil {
		logrus.Errorf("Failed to write report: %v", err)
		return
	}

	_, err = io.WriteString(f, fmt.Sprintf(`
- **Min**: %s
- **Mean**: %s
- **50%%**: %s
- **90%%**: %s
- **95%%**: %s
- **99%%**: %s
- **Max**: %s
`,
		metrics.Latencies.Min,
		metrics.Latencies.Mean,
		metrics.Latencies.P50,
		metrics.Latencies.P90,
		metrics.Latencies.P95,
		metrics.Latencies.P99,
		metrics.Latencies.Max,
	))
	if err != nil {
		logrus.Errorf("Failed to write report: %v", err)
	}

	logrus.Info("Detailed load test report generated")
}
