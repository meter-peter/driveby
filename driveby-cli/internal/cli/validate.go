package cli

import (
	"encoding/json"
	"os"

	"github.com/meter-peter/driveby/driveby-cli/internal/engine"
	"github.com/meter-peter/driveby/driveby-cli/internal/report"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validateOnlyCmd = &cobra.Command{
	Use:   "validate-only",
	Short: "Run only OpenAPI/documentation validation checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := buildValidatorConfig()
		if err != nil {
			return err
		}

		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)

		eng, err := engine.New(cfg)
		if err != nil {
			return logAndReturnError(err)
		}

		result, err := eng.Validate(cmd.Context())
		if err != nil {
			return logAndReturnError(err)
		}

		// Compute top-level status
		result.Status = "passed"
		result.ExitCode = types.ExitSuccess
		for _, principle := range result.Principles {
			if !principle.Passed && principle.Principle.Severity == "critical" {
				result.Status = "failed"
				result.ExitCode = types.ExitValidationFailed
				break
			}
		}

		if err := generator.SaveValidationReport(result); err != nil {
			return logAndReturnError(err)
		}
		json.NewEncoder(os.Stdout).Encode(result)

		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(result, viper.GetString("validation-mode")); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		if result.Status == "failed" {
			return &types.ExitError{Code: types.ExitValidationFailed, Message: "validation failed: critical principle(s) did not pass"}
		}
		return nil
	},
}
