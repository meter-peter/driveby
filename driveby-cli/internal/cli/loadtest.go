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

var loadOnlyCmd = &cobra.Command{
	Use:   "load-only",
	Short: "Run only load/performance tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := buildValidatorConfigWithPerf()
		if err != nil {
			return err
		}

		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)

		tester, err := testing.NewPerformanceTester(cfg)
		if err != nil {
			return logAndReturnError(err)
		}

		result, err := tester.TestPerformance(cmd.Context())
		if err != nil {
			return logAndReturnError(err)
		}

		// Compute top-level status
		result.Status = "passed"
		if result.FailedChecks > 0 {
			result.Status = "failed"
		}

		if err := generator.SaveLoadTestReport(result); err != nil {
			return logAndReturnError(err)
		}
		json.NewEncoder(os.Stdout).Encode(result)

		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(result, "load-testing", nil); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		if result.Status == "failed" {
			return &types.ExitError{Code: types.ExitValidationFailed, Message: "load testing failed: performance targets not met"}
		}
		return nil
	},
}
