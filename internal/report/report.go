package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/meter-peter/driveby/internal/validation"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.Infof("[report] Logger set to DEBUG (verbose) mode")
}

// Generator handles report generation
type Generator struct {
	outputDir string
}

// NewGenerator creates a new report generator
func NewGenerator(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
	}
}

// SaveValidationReport saves a validation report to JSON and Markdown files
func (g *Generator) SaveValidationReport(result *validation.ValidationReport) error {
	log.Debugf("Enter SaveValidationReport with result: %+v", result)
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	jsonPath := filepath.Join(g.outputDir, "validation-report.json")
	if err := g.saveJSON(jsonPath, result); err != nil {
		log.Debugf("Returning from SaveValidationReport with error: %v", err)
		return fmt.Errorf("failed to save JSON report: %w", err)
	}

	// Save Markdown report
	mdPath := filepath.Join(g.outputDir, "validation-report.md")
	if err := g.saveMarkdown(mdPath, result); err != nil {
		log.Debugf("Returning from SaveValidationReport with error: %v", err)
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}

	log.Debugf("Returning from SaveValidationReport with nil")
	return nil
}

// SavePerformanceReport saves a performance test report
func (g *Generator) SavePerformanceReport(result *validation.ValidationReport) error {
	log.Debugf("Enter SavePerformanceReport with result: %+v", result)
	if len(result.Principles) == 0 {
		return fmt.Errorf("no performance metrics in validation result")
	}

	// Find the performance principle result
	var perfMetrics *validation.PerformanceMetrics
	for _, principle := range result.Principles {
		if principle.Principle.ID == "P007" {
			if details, ok := principle.Details.(*validation.PerformanceMetrics); ok {
				perfMetrics = details
				break
			}
		}
	}

	if perfMetrics == nil {
		return fmt.Errorf("no performance metrics found in validation result")
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	jsonPath := filepath.Join(g.outputDir, "loadtest-report.json")
	if err := g.saveJSON(jsonPath, perfMetrics); err != nil {
		log.Debugf("Returning from SavePerformanceReport with error: %v", err)
		return fmt.Errorf("failed to save JSON report: %w", err)
	}

	// Save Markdown report
	mdPath := filepath.Join(g.outputDir, "loadtest-report.md")
	if err := g.saveMarkdown(mdPath, perfMetrics); err != nil {
		log.Debugf("Returning from SavePerformanceReport with error: %v", err)
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}

	log.Debugf("Returning from SavePerformanceReport with nil")
	return nil
}

// SaveFunctionalTestReport saves a functional test report
func (g *Generator) SaveFunctionalTestReport(result *validation.ValidationReport) error {
	log.Debugf("Enter SaveFunctionalTestReport with result: %+v", result)
	if len(result.Principles) == 0 {
		return fmt.Errorf("no functional test results in validation result")
	}

	// Find the functional test principle result
	var endpointResults []validation.EndpointValidation
	for _, principle := range result.Principles {
		if principle.Principle.ID == "P006" {
			if details, ok := principle.Details.([]validation.EndpointValidation); ok {
				endpointResults = details
				break
			}
		}
	}

	if endpointResults == nil {
		return fmt.Errorf("no functional test results found in validation result")
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	jsonPath := filepath.Join(g.outputDir, "functional-test-report.json")
	if err := g.saveJSON(jsonPath, endpointResults); err != nil {
		log.Debugf("Returning from SaveFunctionalTestReport with error: %v", err)
		return fmt.Errorf("failed to save JSON report: %w", err)
	}

	// Save Markdown report
	mdPath := filepath.Join(g.outputDir, "functional-test-report.md")
	if err := g.saveMarkdown(mdPath, endpointResults); err != nil {
		log.Debugf("Returning from SaveFunctionalTestReport with error: %v", err)
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}

	log.Debugf("Returning from SaveFunctionalTestReport with nil")
	return nil
}

// SaveLoadTestReport saves a load test report
func (g *Generator) SaveLoadTestReport(result *validation.ValidationReport) error {
	log.Debugf("Enter SaveLoadTestReport with result: %+v", result)
	if len(result.Principles) == 0 {
		return fmt.Errorf("no load test results in validation result")
	}

	// Find the load test principle result
	var perfMetrics *validation.PerformanceMetrics
	for _, principle := range result.Principles {
		if principle.Principle.ID == "P007" {
			if details, ok := principle.Details.(*validation.PerformanceMetrics); ok {
				perfMetrics = details
				break
			}
		}
	}

	if perfMetrics == nil {
		return fmt.Errorf("no load test results found in validation result")
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	jsonPath := filepath.Join(g.outputDir, "load-test-report.json")
	if err := g.saveJSON(jsonPath, perfMetrics); err != nil {
		log.Debugf("Returning from SaveLoadTestReport with error: %v", err)
		return fmt.Errorf("failed to save JSON report: %w", err)
	}

	// Save Markdown report
	mdPath := filepath.Join(g.outputDir, "load-test-report.md")
	if err := g.saveMarkdown(mdPath, perfMetrics); err != nil {
		log.Debugf("Returning from SaveLoadTestReport with error: %v", err)
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}

	log.Debugf("Returning from SaveLoadTestReport with nil")
	return nil
}

// saveJSON saves a report in JSON format
func (g *Generator) saveJSON(path string, data interface{}) error {
	log.Debugf("Enter saveJSON with path: %s and data: %+v", path, data)
	file, err := os.Create(path)
	if err != nil {
		log.Debugf("Returning from saveJSON with error: %v", err)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Debugf("Returning from saveJSON with error: %v", err)
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	log.Debugf("Returning from saveJSON with nil")
	return nil
}

// saveMarkdown saves a report in Markdown format
func (g *Generator) saveMarkdown(path string, data interface{}) error {
	log.Debugf("Enter saveMarkdown with path: %s and data: %+v", path, data)
	file, err := os.Create(path)
	if err != nil {
		log.Debugf("Returning from saveMarkdown with error: %v", err)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	switch v := data.(type) {
	case *validation.ValidationReport:
		return g.writeValidationMarkdown(file, v)
	case *validation.PerformanceMetrics:
		return g.writeLoadTestMarkdown(file, v)
	case []validation.EndpointValidation:
		return g.writeFunctionalTestMarkdown(file, v)
	default:
		log.Debugf("Returning from saveMarkdown with error: %v", fmt.Errorf("unsupported report type: %T", data))
		return fmt.Errorf("unsupported report type: %T", data)
	}
}

// writeValidationMarkdown writes a validation report in Markdown format
func (g *Generator) writeValidationMarkdown(file *os.File, report *validation.ValidationReport) error {
	log.Debugf("Enter writeValidationMarkdown with report: %+v", report)
	if _, err := fmt.Fprintf(file, `# API Validation Report

Generated: %s
Environment: %s
Version: %s

## Summary

- Total Checks: %d
- Passed Checks: %d
- Failed Checks: %d
- Critical Issues: %d
- Warnings: %d
- Info: %d

### Categories
%s

### Failed Tags
%s

## Principle Results

`, report.Timestamp.Format(time.RFC3339),
		report.Environment,
		report.Version,
		report.TotalChecks,
		report.PassedChecks,
		report.FailedChecks,
		report.Summary.CriticalIssues,
		report.Summary.Warnings,
		report.Summary.Info,
		formatCategories(report.Summary.Categories),
		formatTags(report.Summary.FailedTags)); err != nil {
		return fmt.Errorf("failed to write report header: %w", err)
	}

	// Group principles by category
	principlesByCategory := make(map[string][]validation.PrincipleResult)
	for _, principleResult := range report.Principles {
		category := principleResult.Principle.Category
		principlesByCategory[category] = append(principlesByCategory[category], principleResult)
	}

	// Write results by category
	categories := []string{"Specification", "Documentation", "Schema", "Error Handling", "Security", "Testing", "Performance", "Versioning"}
	for _, category := range categories {
		if principles, ok := principlesByCategory[category]; ok {
			if _, err := fmt.Fprintf(file, "### %s\n\n", category); err != nil {
				return fmt.Errorf("failed to write category header: %w", err)
			}

			for _, principleResult := range principles {
				status := "Passed"
				if !principleResult.Passed {
					status = "Failed"
				}

				// Write principle header with severity
				if _, err := fmt.Fprintf(file, "#### %s: %s (%s) [%s]\n\n",
					principleResult.Principle.ID,
					principleResult.Principle.Name,
					status,
					principleResult.Principle.Severity); err != nil {
					return fmt.Errorf("failed to write principle header: %w", err)
				}

				// Write principle description
				if _, err := fmt.Fprintf(file, "%s\n\n", principleResult.Principle.Description); err != nil {
					return fmt.Errorf("failed to write principle description: %w", err)
				}

				// Write status and message
				if _, err := fmt.Fprintf(file, "- **Status:** %s\n", status); err != nil {
					return fmt.Errorf("failed to write status: %w", err)
				}
				if _, err := fmt.Fprintf(file, "- **Message:** %s\n", principleResult.Message); err != nil {
					return fmt.Errorf("failed to write message: %w", err)
				}

				// Write tags
				if len(principleResult.Principle.Tags) > 0 {
					if _, err := fmt.Fprintf(file, "- **Tags:** %s\n", strings.Join(principleResult.Principle.Tags, ", ")); err != nil {
						return fmt.Errorf("failed to write tags: %w", err)
					}
				}

				// Write checks
				if len(principleResult.Principle.Checks) > 0 {
					if _, err := fmt.Fprintf(file, "\n**Checks Performed:**\n"); err != nil {
						return fmt.Errorf("failed to write checks header: %w", err)
					}
					for _, check := range principleResult.Principle.Checks {
						if _, err := fmt.Fprintf(file, "- %s\n", check); err != nil {
							return fmt.Errorf("failed to write check: %w", err)
						}
					}
				}

				// Write details based on principle type
				if err := g.writePrincipleDetails(file, principleResult); err != nil {
					return err
				}

				// Write suggested fix if available
				if principleResult.SuggestedFix != "" {
					if _, err := fmt.Fprintf(file, "\n**Suggested Fix:**\n%s\n", principleResult.SuggestedFix); err != nil {
						return fmt.Errorf("failed to write suggested fix: %w", err)
					}
				}

				if _, err := fmt.Fprintf(file, "\n---\n\n"); err != nil {
					return fmt.Errorf("failed to write separator: %w", err)
				}
			}
		}
	}

	// Write auto-fixes if any
	if len(report.AutoFixes) > 0 {
		if _, err := fmt.Fprintf(file, "## Automatic Fixes Applied\n\n"); err != nil {
			return fmt.Errorf("failed to write auto-fixes header: %w", err)
		}
		for _, fix := range report.AutoFixes {
			if _, err := fmt.Fprintf(file, "### %s\n\n", fix.PrincipleID); err != nil {
				return fmt.Errorf("failed to write auto-fix header: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- **Time:** %s\n", fix.Timestamp.Format(time.RFC3339)); err != nil {
				return fmt.Errorf("failed to write auto-fix time: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- **Location:** %s\n", fix.Location); err != nil {
				return fmt.Errorf("failed to write auto-fix location: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- **Status:** %s\n", fixStatus(fix.Success)); err != nil {
				return fmt.Errorf("failed to write auto-fix status: %w", err)
			}
			if fix.Message != "" {
				if _, err := fmt.Fprintf(file, "- **Message:** %s\n", fix.Message); err != nil {
					return fmt.Errorf("failed to write auto-fix message: %w", err)
				}
			}
			if fix.Error != "" {
				if _, err := fmt.Fprintf(file, "- **Error:** %s\n", fix.Error); err != nil {
					return fmt.Errorf("failed to write auto-fix error: %w", err)
				}
			}
			if _, err := fmt.Fprintf(file, "\n"); err != nil {
				return fmt.Errorf("failed to write auto-fix separator: %w", err)
			}
		}
	}

	return nil
}

func formatCategories(categories []string) string {
	if len(categories) == 0 {
		return "- None"
	}
	var result strings.Builder
	for _, category := range categories {
		result.WriteString(fmt.Sprintf("- %s\n", category))
	}
	return result.String()
}

func formatTags(tags []string) string {
	if len(tags) == 0 {
		return "- None"
	}
	var result strings.Builder
	for _, tag := range tags {
		result.WriteString(fmt.Sprintf("- %s\n", tag))
	}
	return result.String()
}

func fixStatus(success bool) string {
	if success {
		return "✅ Success"
	}
	return "❌ Failed"
}

func (g *Generator) writePrincipleDetails(file *os.File, principleResult validation.PrincipleResult) error {
	if principleResult.Details == nil {
		return nil
	}

	if _, err := fmt.Fprintf(file, "\n**Details:**\n"); err != nil {
		return fmt.Errorf("failed to write details header: %w", err)
	}

	switch principleResult.Principle.ID {
	case "P006": // API Contract Testing
		if details, ok := principleResult.Details.([]validation.EndpointValidation); ok {
			if _, err := fmt.Fprintf(file, "\n#### Endpoint Test Results\n\n"); err != nil {
				return fmt.Errorf("failed to write endpoint results header: %w", err)
			}
			for _, epVal := range details {
				statusEmoji := "✅"
				if epVal.Status == "error" {
					statusEmoji = "❌"
				} else if epVal.Status == "warning" {
					statusEmoji = "⚠️"
				}
				if _, err := fmt.Fprintf(file, "%s **%s %s**\n", statusEmoji, epVal.Method, epVal.Path); err != nil {
					return fmt.Errorf("failed to write endpoint header: %w", err)
				}
				if _, err := fmt.Fprintf(file, "  - Status: %s\n", epVal.Status); err != nil {
					return fmt.Errorf("failed to write endpoint status: %w", err)
				}
				if _, err := fmt.Fprintf(file, "  - Code: %d\n", epVal.StatusCode); err != nil {
					return fmt.Errorf("failed to write endpoint code: %w", err)
				}
				if _, err := fmt.Fprintf(file, "  - Response Time: %s\n", epVal.ResponseTime); err != nil {
					return fmt.Errorf("failed to write endpoint response time: %w", err)
				}
				if len(epVal.Errors) > 0 {
					if _, err := fmt.Fprintf(file, "  - Errors:\n"); err != nil {
						return fmt.Errorf("failed to write endpoint errors header: %w", err)
					}
					for _, err := range epVal.Errors {
						if _, err := fmt.Fprintf(file, "    - %s\n", err); err != nil {
							return fmt.Errorf("failed to write endpoint error: %w", err)
						}
					}
				}
				if _, err := fmt.Fprintf(file, "\n"); err != nil {
					return fmt.Errorf("failed to write endpoint separator: %w", err)
				}
			}
		}
	case "P007": // Performance Requirements
		if details, ok := principleResult.Details.(*validation.PerformanceMetrics); ok {
			if _, err := fmt.Fprintf(file, "\n#### Performance Metrics\n\n"); err != nil {
				return fmt.Errorf("failed to write performance metrics header: %w", err)
			}
			if _, err := fmt.Fprintf(file, "**Request Statistics**\n"); err != nil {
				return fmt.Errorf("failed to write request stats header: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- Total Requests: %d\n", details.TotalRequests); err != nil {
				return fmt.Errorf("failed to write total requests: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- Success Rate: %.2f%%\n", (1-details.ErrorRate)*100); err != nil {
				return fmt.Errorf("failed to write success rate: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- Error Rate: %.2f%%\n", details.ErrorRate*100); err != nil {
				return fmt.Errorf("failed to write error rate: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- Requests/sec: %.2f\n\n", details.RequestsPerSec); err != nil {
				return fmt.Errorf("failed to write requests per second: %w", err)
			}

			if _, err := fmt.Fprintf(file, "**Latency Percentiles**\n"); err != nil {
				return fmt.Errorf("failed to write latency header: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- P50: %s\n", details.LatencyP50); err != nil {
				return fmt.Errorf("failed to write latency P50: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- P95: %s\n", details.LatencyP95); err != nil {
				return fmt.Errorf("failed to write latency P95: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- P99: %s\n\n", details.LatencyP99); err != nil {
				return fmt.Errorf("failed to write latency P99: %w", err)
			}

			if _, err := fmt.Fprintf(file, "**Test Duration**\n"); err != nil {
				return fmt.Errorf("failed to write duration header: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- Start: %s\n", details.StartTime.Format(time.RFC3339)); err != nil {
				return fmt.Errorf("failed to write start time: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- End: %s\n", details.EndTime.Format(time.RFC3339)); err != nil {
				return fmt.Errorf("failed to write end time: %w", err)
			}
		}
	default:
		// For other principles, format details as JSON
		detailsJSON, err := json.MarshalIndent(principleResult.Details, "  ", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal details: %w", err)
		}
		if _, err := fmt.Fprintf(file, "```json\n%s\n```\n", string(detailsJSON)); err != nil {
			return fmt.Errorf("failed to write details JSON: %w", err)
		}
	}

	return nil
}

// writeLoadTestMarkdown writes a load test report in Markdown format
func (g *Generator) writeLoadTestMarkdown(file *os.File, metrics *validation.PerformanceMetrics) error {
	log.Debugf("Enter writeLoadTestMarkdown with metrics: %+v", metrics)
	if _, err := fmt.Fprintf(file, `# Load Test Report

## Summary

- Start Time: %s
- End Time: %s
- Total Requests: %d
- Success Count: %d
- Error Count: %d
- Error Rate: %.2f%%
- Requests/sec: %.2f

## Latency Percentiles

- P50: %s
- P95: %s
- P99: %s

## Test Duration

- Start: %s
- End: %s
- Duration: %s

`,
		metrics.StartTime.Format(time.RFC3339),
		metrics.EndTime.Format(time.RFC3339),
		metrics.TotalRequests,
		metrics.SuccessCount,
		metrics.ErrorCount,
		metrics.ErrorRate*100,
		metrics.RequestsPerSec,
		metrics.LatencyP50,
		metrics.LatencyP95,
		metrics.LatencyP99,
		metrics.StartTime.Format(time.RFC3339),
		metrics.EndTime.Format(time.RFC3339),
		metrics.EndTime.Sub(metrics.StartTime)); err != nil {
		log.Debugf("Returning from writeLoadTestMarkdown with error: %v", err)
		return fmt.Errorf("failed to write load test report header: %w", err)
	}
	log.Debugf("Returning from writeLoadTestMarkdown with nil")
	return nil
}

// writeFunctionalTestMarkdown writes a functional test report in Markdown format
func (g *Generator) writeFunctionalTestMarkdown(file *os.File, results []validation.EndpointValidation) error {
	log.Debugf("Enter writeFunctionalTestMarkdown with results: %+v", results)
	if _, err := fmt.Fprintf(file, `# Functional Test Report

## Summary

- Total Endpoints Tested: %d
- Successful: %d
- Failed: %d
- Warnings: %d

## Test Results

`, len(results),
		countStatus(results, "success"),
		countStatus(results, "error"),
		countStatus(results, "warning")); err != nil {
		return fmt.Errorf("failed to write functional test report header: %w", err)
	}

	// Group results by status
	resultsByStatus := make(map[string][]validation.EndpointValidation)
	for _, result := range results {
		resultsByStatus[result.Status] = append(resultsByStatus[result.Status], result)
	}

	// Write results by status
	statuses := []string{"success", "error", "warning"}
	for _, status := range statuses {
		if endpoints, ok := resultsByStatus[status]; ok {
			statusEmoji := "✅"
			if status == "error" {
				statusEmoji = "❌"
			} else if status == "warning" {
				statusEmoji = "⚠️"
			}

			if _, err := fmt.Fprintf(file, "\n### %s %s (%d)\n\n", statusEmoji, strings.Title(status), len(endpoints)); err != nil {
				return fmt.Errorf("failed to write status header: %w", err)
			}

			for _, endpoint := range endpoints {
				if _, err := fmt.Fprintf(file, "#### %s %s\n", endpoint.Method, endpoint.Path); err != nil {
					return fmt.Errorf("failed to write endpoint header: %w", err)
				}
				if _, err := fmt.Fprintf(file, "- Status Code: %d\n", endpoint.StatusCode); err != nil {
					return fmt.Errorf("failed to write status code: %w", err)
				}
				if _, err := fmt.Fprintf(file, "- Response Time: %s\n", endpoint.ResponseTime); err != nil {
					return fmt.Errorf("failed to write response time: %w", err)
				}
				if len(endpoint.Errors) > 0 {
					if _, err := fmt.Fprintf(file, "- Errors:\n"); err != nil {
						return fmt.Errorf("failed to write errors header: %w", err)
					}
					for _, err := range endpoint.Errors {
						if _, err := fmt.Fprintf(file, "  - %s\n", err); err != nil {
							return fmt.Errorf("failed to write error: %w", err)
						}
					}
				}
				if _, err := fmt.Fprintf(file, "\n"); err != nil {
					return fmt.Errorf("failed to write endpoint separator: %w", err)
				}
			}
		}
	}

	return nil
}

// countStatus counts the number of endpoints with a given status
func countStatus(results []validation.EndpointValidation, status string) int {
	count := 0
	for _, result := range results {
		if result.Status == status {
			count++
		}
	}
	return count
}

// formatError formats an error message for the report
func formatError(errMsg string) string {
	if errMsg == "" {
		return ""
	}
	return fmt.Sprintf("- Error: %s", errMsg)
}
