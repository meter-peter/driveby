package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var githubCommentCmd = &cobra.Command{
	Use:   "github-comment",
	Short: "Post a combined PR comment from pre-saved report files",
	Long: `Reads validation, functional test, and load test reports from a shared directory
and posts a combined rich PR comment. Unlike validate-only --github-comment, this command
does NOT re-run any tests — it assembles results from prior pipeline steps.

Used by Argo Workflow DAGs where each step runs in a separate pod but shares a volume.`,
	RunE: runGitHubComment,
}

func runGitHubComment(cmd *cobra.Command, args []string) error {
	reportDir, _ := cmd.Flags().GetString("report-dir")
	if reportDir == "" {
		reportDir = "/tmp/reports"
	}

	// Load validation report (base)
	report, err := loadValidationReport(reportDir)
	if err != nil {
		logrus.WithError(err).Warn("No validation report found, creating synthetic report")
		report = &types.ValidationReport{
			Status:  "passed",
			Summary: types.ValidationSummary{},
		}
	}

	// Load functional test results and attach
	if funcResults, err := loadFunctionalReport(reportDir); err == nil {
		if report.TestResults == nil {
			report.TestResults = &types.TestResults{}
		}
		report.TestResults.Functional = funcResults
		logrus.Info("Loaded functional test results")
	} else {
		logrus.WithError(err).Debug("No functional test report found")
	}

	// Load load test results and attach
	if perfMetrics, err := loadLoadtestReport(reportDir); err == nil {
		if report.TestResults == nil {
			report.TestResults = &types.TestResults{}
		}
		report.TestResults.Performance = &types.PerformanceTestResults{
			TotalRequests:     int64(perfMetrics.TotalRequests),
			SuccessCount:      int64(perfMetrics.SuccessCount),
			ErrorCount:        int64(perfMetrics.ErrorCount),
			ErrorRate:         perfMetrics.ErrorRate,
			LatencyP50:        perfMetrics.LatencyP50,
			LatencyP95:        perfMetrics.LatencyP95,
			LatencyP99:        perfMetrics.LatencyP99,
			RequestsPerSecond: perfMetrics.RequestsPerSec,
			Duration:          perfMetrics.EndTime.Sub(perfMetrics.StartTime),
		}
		logrus.Info("Loaded load test results")
	} else {
		logrus.WithError(err).Debug("No load test report found")
	}

	// Determine validation mode: infer from report, override with flag
	validationMode := "strict"
	if len(report.Principles) > 0 {
		validationMode = inferValidationMode(report)
	}
	if m, _ := cmd.Flags().GetString("validation-mode"); m != "" {
		validationMode = m
	}

	return handleGitHubComment(report, validationMode)
}

func inferValidationMode(r *types.ValidationReport) string {
	hasP009 := false
	principleCount := len(r.Principles)
	for _, p := range r.Principles {
		if p.Principle.ID == "P009" {
			hasP009 = true
		}
	}
	if hasP009 && principleCount <= 3 {
		return "test-ready"
	}
	if principleCount == 1 {
		return "minimal"
	}
	return "strict"
}

func loadValidationReport(dir string) (*types.ValidationReport, error) {
	path := filepath.Join(dir, "validation-report-latest.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read validation report: %w", err)
	}
	var report types.ValidationReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("parse validation report: %w", err)
	}
	return &report, nil
}

func loadFunctionalReport(dir string) (*types.FunctionalTestResults, error) {
	path := filepath.Join(dir, "functional-test-report-latest.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read functional report: %w", err)
	}
	var results types.FunctionalTestResults
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("parse functional report: %w", err)
	}
	return &results, nil
}

func loadLoadtestReport(dir string) (*types.PerformanceMetrics, error) {
	// Try both naming conventions (load-test and loadtest)
	for _, name := range []string{"loadtest-report-latest.json", "load-test-report-latest.json"} {
		path := filepath.Join(dir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var metrics types.PerformanceMetrics
		if err := json.Unmarshal(data, &metrics); err != nil {
			continue
		}
		return &metrics, nil
	}
	return nil, fmt.Errorf("no load test report found in %s", dir)
}
