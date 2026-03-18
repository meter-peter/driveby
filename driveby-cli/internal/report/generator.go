package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/meter-peter/driveby/driveby-cli/internal/types"

	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()

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

// timestampedName returns a filename with an embedded timestamp, e.g.
// "validation-report-20060102T150405Z.json".
func timestampedName(base, ext string) string {
	ts := time.Now().UTC().Format("20060102T150405Z")
	return fmt.Sprintf("%s-%s%s", base, ts, ext)
}

// SaveValidationReport saves a validation report to JSON and Markdown files
func (g *Generator) SaveValidationReport(result *types.ValidationReport) error {
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save timestamped JSON report + latest symlink
	jsonName := timestampedName("validation-report", ".json")
	jsonPath := filepath.Join(g.outputDir, jsonName)
	if err := g.saveJSON(jsonPath, result); err != nil {
		return fmt.Errorf("failed to save JSON report: %w", err)
	}
	copyFile(jsonPath, filepath.Join(g.outputDir, "validation-report-latest.json"))

	// Save timestamped Markdown report + latest symlink
	mdName := timestampedName("validation-report", ".md")
	mdPath := filepath.Join(g.outputDir, mdName)
	if err := g.saveMarkdown(mdPath, result); err != nil {
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}
	copyFile(mdPath, filepath.Join(g.outputDir, "validation-report-latest.md"))

	return nil
}

// SavePerformanceReport saves a performance test report
func (g *Generator) SavePerformanceReport(result *types.ValidationReport) error {
	if len(result.Principles) == 0 {
		return fmt.Errorf("no performance metrics in validation result")
	}

	var perfMetrics *types.PerformanceMetrics
	for _, principle := range result.Principles {
		if principle.Principle.ID == "P007" {
			if details, ok := principle.Details.(*types.PerformanceMetrics); ok {
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

	jsonPath := filepath.Join(g.outputDir, timestampedName("loadtest-report", ".json"))
	if err := g.saveJSON(jsonPath, perfMetrics); err != nil {
		return fmt.Errorf("failed to save JSON report: %w", err)
	}
	copyFile(jsonPath, filepath.Join(g.outputDir, "loadtest-report-latest.json"))

	mdPath := filepath.Join(g.outputDir, timestampedName("loadtest-report", ".md"))
	if err := g.saveMarkdown(mdPath, perfMetrics); err != nil {
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}
	copyFile(mdPath, filepath.Join(g.outputDir, "loadtest-report-latest.md"))

	return nil
}

// SaveFunctionalTestReport saves a functional test report
func (g *Generator) SaveFunctionalTestReport(result *types.ValidationReport) error {
	if result.TestResults == nil || result.TestResults.Functional == nil {
		return fmt.Errorf("no functional test results found in validation result")
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	jsonPath := filepath.Join(g.outputDir, timestampedName("functional-test-report", ".json"))
	if err := g.saveJSON(jsonPath, result.TestResults.Functional); err != nil {
		return fmt.Errorf("failed to save JSON report: %w", err)
	}
	copyFile(jsonPath, filepath.Join(g.outputDir, "functional-test-report-latest.json"))

	mdPath := filepath.Join(g.outputDir, timestampedName("functional-test-report", ".md"))
	if err := g.saveMarkdown(mdPath, result.TestResults.Functional); err != nil {
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}
	copyFile(mdPath, filepath.Join(g.outputDir, "functional-test-report-latest.md"))

	return nil
}

// SaveLoadTestReport saves a load test report
func (g *Generator) SaveLoadTestReport(result *types.ValidationReport) error {
	if len(result.Principles) == 0 {
		return fmt.Errorf("no load test results in validation result")
	}

	var perfMetrics *types.PerformanceMetrics
	for _, principle := range result.Principles {
		if principle.Principle.ID == "P007" {
			if details, ok := principle.Details.(*types.PerformanceMetrics); ok {
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

	jsonPath := filepath.Join(g.outputDir, timestampedName("load-test-report", ".json"))
	if err := g.saveJSON(jsonPath, perfMetrics); err != nil {
		return fmt.Errorf("failed to save JSON report: %w", err)
	}
	copyFile(jsonPath, filepath.Join(g.outputDir, "load-test-report-latest.json"))

	mdPath := filepath.Join(g.outputDir, timestampedName("load-test-report", ".md"))
	if err := g.saveMarkdown(mdPath, perfMetrics); err != nil {
		return fmt.Errorf("failed to save Markdown report: %w", err)
	}
	copyFile(mdPath, filepath.Join(g.outputDir, "load-test-report-latest.md"))

	return nil
}

// saveJSON saves a report in JSON format
func (g *Generator) saveJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

// copyFile copies src to dst (for "latest" convenience copies).
func copyFile(src, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		log.WithError(err).Warnf("failed to read %s for latest copy", src)
		return
	}
	if err := os.WriteFile(dst, data, 0644); err != nil {
		log.WithError(err).Warnf("failed to write latest copy %s", dst)
	}
}
