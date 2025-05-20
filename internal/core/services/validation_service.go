package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"driveby/internal/config"
	"driveby/internal/core/models"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/sirupsen/logrus"
)

// ValidationServiceImpl implements the ValidationService interface
type ValidationServiceImpl struct {
	config        *config.Config
	logger        *logrus.Logger
	githubService GitHubService
}

// NewValidationService creates a new validation service
func NewValidationService(
	cfg *config.Config,
	logger *logrus.Logger,
	githubService GitHubService,
) ValidationService {
	return &ValidationServiceImpl{
		config:        cfg,
		logger:        logger,
		githubService: githubService,
	}
}

// ValidateOpenAPI performs validation on an OpenAPI specification
func (s *ValidationServiceImpl) ValidateOpenAPI(ctx context.Context, test *models.ValidationTest) (*models.ValidationResult, error) {
	s.logger.WithField("test_id", test.ID).Info("Starting OpenAPI validation")

	// Create or get result
	var result *models.ValidationResult
	if test.Result == nil {
		result = models.NewValidationResult(test.ID)
	} else {
		result = test.Result
	}

	// Fetch OpenAPI spec
	doc, err := s.fetchOpenAPI(ctx, test.OpenAPIURL)
	if err != nil {
		result.Status = models.TestStatusFailed
		result.ErrorDetail = fmt.Sprintf("Failed to fetch OpenAPI spec: %v", err)
		test.Result = result
		return result, err
	}

	// Validate documentation
	validationReport, validationErrors := s.validateAPIDocumentation(doc)

	// Update result
	result.ComplianceScore = validationReport.ComplianceScore
	result.MissingExamples = validationReport.MissingExamples
	result.UndocumentedEndpoints = validationReport.UndocumentedEndpoints
	result.ErrorResponses = validationReport.ErrorResponses

	// Convert errors to ValidationError objects
	for _, err := range validationErrors {
		var endpoint, message string
		parts := strings.SplitN(err.Error(), ":", 2)
		if len(parts) > 1 {
			endpoint = strings.TrimSpace(parts[0])
			message = strings.TrimSpace(parts[1])
		} else {
			message = err.Error()
		}

		result.ValidationErrors = append(result.ValidationErrors, models.ValidationError{
			EndpointID: endpoint,
			Message:    message,
			Severity:   "error",
		})
	}

	// Set status
	now := time.Now()
	result.EndTime = now
	result.Duration = now.Sub(result.StartTime).String()

	if result.ComplianceScore < test.ComplianceThreshold && test.FailOnValidation {
		result.Status = models.TestStatusFailed
		result.ErrorDetail = fmt.Sprintf("Compliance score %.2f%% is below threshold %.2f%%",
			result.ComplianceScore, test.ComplianceThreshold)
	} else {
		result.Status = models.TestStatusCompleted
	}

	// Update test with result
	test.Result = result

	s.logger.WithFields(logrus.Fields{
		"test_id":          test.ID,
		"compliance_score": result.ComplianceScore,
		"status":           result.Status,
	}).Info("OpenAPI validation completed")

	return result, nil
}

// GetValidationTest retrieves a validation test by ID
func (s *ValidationServiceImpl) GetValidationTest(ctx context.Context, testID string) (*models.ValidationTest, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListValidationTests retrieves all validation tests
func (s *ValidationServiceImpl) ListValidationTests(ctx context.Context) ([]*models.ValidationTest, error) {
	return nil, fmt.Errorf("not implemented")
}

// QueueValidationTest queues a validation test for asynchronous processing
func (s *ValidationServiceImpl) QueueValidationTest(ctx context.Context, test *models.ValidationTest) error {
	return fmt.Errorf("not implemented")
}

// GenerateReport creates a validation report for a completed test
func (s *ValidationServiceImpl) GenerateReport(ctx context.Context, testID string) (string, error) {
	s.logger.WithField("test_id", testID).Info("Generating validation report")

	// Get test
	test, err := s.GetValidationTest(ctx, testID)
	if err != nil {
		return "", fmt.Errorf("failed to get validation test: %w", err)
	}

	// Check if test has result
	if test.Result == nil {
		return "", fmt.Errorf("test has no result")
	}

	// Build report
	reportContent := s.buildValidationReport(test)

	return reportContent, nil
}

// fetchOpenAPI fetches an OpenAPI specification from a URL
func (s *ValidationServiceImpl) fetchOpenAPI(ctx context.Context, url string) (*openapi3.T, error) {
	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch OpenAPI spec: %s", resp.Status)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse OpenAPI spec
	var spec map[string]interface{}
	if err := json.Unmarshal(body, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAPI spec: %w", err)
	}

	// Load OpenAPI spec
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(body)
	if err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	return doc, nil
}

// validateAPIDocumentation validates an OpenAPI specification
func (s *ValidationServiceImpl) validateAPIDocumentation(doc *openapi3.T) (models.DocumentationReport, []error) {
	report := models.DocumentationReport{
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
			if operation.Responses == nil || len(operation.Responses.Map()) == 0 {
				errors = append(errors, fmt.Errorf("%s: missing response documentation", endpointId))
				endpointCompliant = false
			}

			// Check all response status codes have documentation
			for statusCode, response := range operation.Responses.Map() {
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
				s.logger.Warnf("%s: Missing metadata (summary or tags)", endpointId)
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

// buildValidationReport builds a validation report in markdown format
func (s *ValidationServiceImpl) buildValidationReport(test *models.ValidationTest) string {
	result := test.Result
	if result == nil {
		return "No validation results available."
	}

	report := fmt.Sprintf(`
## API Documentation Validation Report

**Test:** %s  
**Open API URL:** %s  
**Compliance Score:** %.2f%%  
**Threshold:** %.2f%%  
**Missing Examples:** %d  
**Error Responses:** %d

`,
		test.Name,
		test.OpenAPIURL,
		result.ComplianceScore,
		test.ComplianceThreshold,
		result.MissingExamples,
		len(result.ErrorResponses),
	)

	// Add validation status
	if result.ComplianceScore >= test.ComplianceThreshold {
		report += "**Status:** ✅ Validation Passed\n\n"
	} else {
		report += "**Status:** ❌ Validation Failed\n\n"
	}

	// Add validation errors
	if len(result.ValidationErrors) > 0 {
		report += "### Critical Issues:\n\n"
		for _, err := range result.ValidationErrors {
			severity := ""
			if err.Severity != "" {
				severity = fmt.Sprintf(" [%s]", err.Severity)
			}
			if err.EndpointID != "" {
				report += fmt.Sprintf("- **%s**%s: %s\n", err.EndpointID, severity, err.Message)
			} else {
				report += fmt.Sprintf("- %s%s\n", err.Message, severity)
			}
		}
		report += "\n"
	}

	// Add undocumented endpoints
	if len(result.UndocumentedEndpoints) > 0 {
		report += "### Undocumented Endpoints:\n\n"
		for _, endpoint := range result.UndocumentedEndpoints {
			report += fmt.Sprintf("- `%s`\n", endpoint)
		}
		report += "\n"
	}

	// Add error responses
	if len(result.ErrorResponses) > 0 {
		report += "### Error Response Codes:\n\n"
		for code, count := range result.ErrorResponses {
			report += fmt.Sprintf("- **%s**: %d occurrences\n", code, count)
		}
		report += "\n"
	}

	// Add test details
	report += fmt.Sprintf(`
### Test Details

**Test ID:** %s  
**Created:** %s  
**Completed:** %s  
**Duration:** %s
`,
		test.ID,
		test.CreatedAt.Format(time.RFC3339),
		result.EndTime.Format(time.RFC3339),
		result.Duration,
	)

	return report
}
