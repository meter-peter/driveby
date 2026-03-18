package cli

import (
	"encoding/json"
	"os"
	"time"

	"github.com/meter-peter/driveby/driveby-cli/internal/report"
	"github.com/meter-peter/driveby/driveby-cli/internal/testing"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testOnlyCmd = &cobra.Command{
	Use:   "test-only",
	Short: "Run functional and performance tests without validation",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := buildValidatorConfigWithPerf()
		if err != nil {
			return err
		}
		cfg.ValidationMode = types.ValidationModeTestOnly

		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)

		// Run functional tests
		logrus.Info("Running functional tests...")
		functionalTester := testing.NewFunctionalTester(cfg)
		functionalReport, err := functionalTester.TestEndpoints(cmd.Context())
		if err != nil {
			return logAndReturnError(err)
		}

		// Run performance tests
		logrus.Info("Running performance tests...")
		performanceTester, err := testing.NewPerformanceTester(cfg)
		if err != nil {
			return logAndReturnError(err)
		}
		performanceReport, err := performanceTester.TestPerformance(cmd.Context())
		if err != nil {
			return logAndReturnError(err)
		}

		// Combine results
		combinedReport := map[string]interface{}{
			"functional":  functionalReport,
			"performance": performanceReport,
			"mode":        "test-only",
			"timestamp":   time.Now(),
		}

		// Save reports
		if err := generator.SaveFunctionalTestReport(functionalReport); err != nil {
			return logAndReturnError(err)
		}
		if err := generator.SaveLoadTestReport(performanceReport); err != nil {
			return logAndReturnError(err)
		}

		// Compute top-level status
		status := "passed"
		hasFailures := false
		if functionalReport.TestResults != nil && functionalReport.TestResults.Functional != nil {
			for _, endpoint := range functionalReport.TestResults.Functional.EndpointResults {
				if endpoint.Status == types.TestStatusFailed {
					hasFailures = true
					break
				}
			}
		}
		if performanceReport.TestResults != nil && performanceReport.TestResults.Performance != nil {
			if performanceReport.TestResults.Performance.Status == types.TestStatusFailed {
				hasFailures = true
			}
		}
		if hasFailures {
			status = "failed"
		}
		combinedReport["status"] = status

		json.NewEncoder(os.Stdout).Encode(combinedReport)

		// GitHub integration
		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(combinedReport, "test-only"); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		if hasFailures {
			return &types.ExitError{Code: types.ExitValidationFailed, Message: "testing failed: one or more tests did not pass"}
		}
		return nil
	},
}
