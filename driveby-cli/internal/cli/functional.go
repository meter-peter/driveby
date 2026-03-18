package cli

import (
	"encoding/json"
	"os"

	"github.com/meter-peter/driveby/driveby-cli/internal/report"
	"github.com/meter-peter/driveby/driveby-cli/internal/testing"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var functionOnlyCmd = &cobra.Command{
	Use:   "function-only",
	Short: "Run only functional tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := buildValidatorConfig()
		if err != nil {
			return err
		}

		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)

		tester := testing.NewFunctionalTester(cfg)
		result, err := tester.TestEndpoints(cmd.Context())
		if err != nil {
			return logAndReturnError(err)
		}

		// Compute top-level status
		result.Status = "passed"
		result.ExitCode = types.ExitSuccess
		if result.TestResults != nil && result.TestResults.Functional != nil {
			for _, endpoint := range result.TestResults.Functional.EndpointResults {
				if endpoint.Status == types.TestStatusFailed {
					result.Status = "failed"
					result.ExitCode = types.ExitValidationFailed
					break
				}
			}
		}

		if err := generator.SaveFunctionalTestReport(result); err != nil {
			return logAndReturnError(err)
		}
		json.NewEncoder(os.Stdout).Encode(result)

		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(result, "functional-testing"); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		if result.Status == "failed" {
			return &types.ExitError{Code: types.ExitValidationFailed, Message: "functional testing failed: one or more endpoints did not pass"}
		}
		return nil
	},
}
