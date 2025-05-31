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
func (g *Generator) SavePerformanceReport(result *validation.ValidationResult) error {
	log.Debugf("Enter SavePerformanceReport with result: %+v", result)
	if result.Performance == nil {
		return fmt.Errorf("no performance metrics in validation result")
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	jsonPath := filepath.Join(g.outputDir, "loadtest-report.json")
	if err := g.saveJSON(jsonPath, result.Performance); err != nil {
		log.Debugf("Returning from SavePerformanceReport with error: %v", err)
		return fmt.Errorf("failed to save JSON report: %w", err)
	}

	// Save Markdown report
	mdPath := filepath.Join(g.outputDir, "loadtest-report.md")
	if err := g.saveMarkdown(mdPath, result.Performance); err != nil {
		log.Debugf("Returning from SavePerformanceReport with error: %v", err)
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}

	log.Debugf("Returning from SavePerformanceReport with nil")
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
		return g.writePerformanceMarkdown(file, v)
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

## Summary

- Total Checks: %d
- Passed Checks: %d
- Failed Checks: %d
- Critical Issues: %d
- Warnings: %d
- Info: %d

`, report.Timestamp.Format(time.RFC3339),
		report.TotalChecks,
		report.PassedChecks,
		report.FailedChecks,
		report.Summary.CriticalIssues,
		report.Summary.Warnings,
		report.Summary.Info,
	); err != nil {
		log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
		return fmt.Errorf("failed to write report summary: %w", err)
	}

	// Principle Results
	if len(report.Principles) > 0 {
		if _, err := fmt.Fprintf(file, "## Principle Results\n\n"); err != nil {
			log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
			return fmt.Errorf("failed to write principle results header: %w", err)
		}

		for _, principleResult := range report.Principles {
			status := "Passed"
			if !principleResult.Passed {
				status = "Failed"
			}
			if _, err := fmt.Fprintf(file, "### %s: %s (%s)\n\n", principleResult.Principle.ID, principleResult.Principle.Name, status); err != nil {
				log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
				return fmt.Errorf("failed to write principle summary: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- **Status:** %s\n", status); err != nil {
				log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
				return fmt.Errorf("failed to write principle status: %w", err)
			}
			if _, err := fmt.Fprintf(file, "- **Message:** %s\n", principleResult.Message); err != nil {
				log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
				return fmt.Errorf("failed to write principle message: %w", err)
			}

			// Display details based on principle ID
			switch principleResult.Principle.ID {
			case "P006": // Endpoint Functional Testing Details
				if details, ok := principleResult.Details.([]validation.EndpointValidation); ok {
					if _, err := fmt.Fprintf(file, "\n#### Endpoint Validation Results\n\n"); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write endpoint results header: %w", err)
					}
					for _, epVal := range details {
						if _, err := fmt.Fprintf(file, "- **%s %s:** Status: %s, Code: %d, Time: %s\n", epVal.Method, epVal.Path, epVal.Status, epVal.StatusCode, epVal.ResponseTime); err != nil {
							log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
							return fmt.Errorf("failed to write endpoint result: %w", err)
						}
						if len(epVal.Errors) > 0 {
							if _, err := fmt.Fprintf(file, "  - Errors: %s\n", strings.Join(epVal.Errors, "; ")); err != nil {
								log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
								return fmt.Errorf("failed to write endpoint errors: %w", err)
							}
						}
					}
				}
			case "P007": // API Performance Compliance Details
				if details, ok := principleResult.Details.(*validation.PerformanceMetrics); ok {
					if _, err := fmt.Fprintf(file, "\n#### Performance Metrics\n\n"); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write performance metrics header: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Total Requests:** %d\n", details.TotalRequests); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write total requests: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Success Count:** %d\n", details.SuccessCount); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write success count: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Error Count:** %d\n", details.ErrorCount); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write error count: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Error Rate:** %.2f%%\n", details.ErrorRate*100); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write error rate: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Latency (P50):** %s\n", details.LatencyP50); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write latency P50: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Latency (P95):** %s\n", details.LatencyP95); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write latency P95: %w", err)
					}
					if _, err := fmt.Fprintf(file, "- **Latency (P99):** %s\n\n", details.LatencyP99); err != nil {
						log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
						return fmt.Errorf("failed to write latency P99: %w", err)
					}
				}
			default:
				// For other principles, just print the message and any simple details
				if principleResult.Details != nil {
					detailsJSON, err := json.MarshalIndent(principleResult.Details, "  ", "  ")
					if err != nil {
						log.WithError(err).Warnf("Failed to marshal details for principle %s", principleResult.Principle.ID)
					} else {
						if _, err := fmt.Fprintf(file, "\n#### Details\n\n```json\n%s\n```\n\n", string(detailsJSON)); err != nil {
							log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
							return fmt.Errorf("failed to write principle details: %w", err)
						}
					}
				}
			}

			if _, err := fmt.Fprintf(file, "\n"); err != nil {
				log.Debugf("Returning from writeValidationMarkdown with error: %v", err)
				return fmt.Errorf("failed to write principle spacing: %w", err)
			}
		}
	}

	log.Debugf("Returning from writeValidationMarkdown with nil")
	return nil
}

// writePerformanceMarkdown writes a performance test report in Markdown format
func (g *Generator) writePerformanceMarkdown(file *os.File, metrics *validation.PerformanceMetrics) error {
	log.Debugf("Enter writePerformanceMarkdown with metrics: %+v", metrics)
	if _, err := fmt.Fprintf(file, `# Performance Test Report

## Summary

- Start Time: %s
- End Time: %s
- Total Requests: %d
- Success Count: %d
- Error Count: %d
- Error Rate: %.2f%%

## Latency

- P50: %s
- P95: %s
- P99: %s

`,
		metrics.StartTime.Format(time.RFC3339),
		metrics.EndTime.Format(time.RFC3339),
		metrics.TotalRequests,
		metrics.SuccessCount,
		metrics.ErrorCount,
		metrics.ErrorRate*100,
		metrics.LatencyP50,
		metrics.LatencyP95,
		metrics.LatencyP99); err != nil {
		log.Debugf("Returning from writePerformanceMarkdown with error: %v", err)
		return fmt.Errorf("failed to write performance report header: %w", err)
	}
	log.Debugf("Returning from writePerformanceMarkdown with nil")
	return nil
}

// formatError formats an error message for the report
func formatError(errMsg string) string {
	if errMsg == "" {
		return ""
	}
	return fmt.Sprintf("- Error: %s", errMsg)
}
