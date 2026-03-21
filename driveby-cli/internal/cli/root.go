package cli

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/meter-peter/driveby/driveby-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "driveby",
		Short: "DriveBy - A modern API validation framework",
		Long: `DriveBy is a modern API validation framework that helps you validate, test, and monitor your APIs.
It supports OpenAPI/Swagger specifications and provides comprehensive validation, testing, and rollout capabilities.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Initialize logger — logs always go to stderr so stdout is clean JSON
			logCfg := logger.DefaultConfig()
			logCfg.Level = viper.GetString("log-level")
			logCfg.Format = "json"
			logCfg.Output = "stderr"
			if err := logger.Configure(logCfg); err != nil {
				return fmt.Errorf("failed to configure logger: %w", err)
			}

			// Set up signal-based context for graceful shutdown
			ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT)
			cmd.SetContext(ctx)
			go func() {
				<-ctx.Done()
				stop()
			}()
			return nil
		},
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root command flags
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().String("api-url", "", "Base URL of the API to test (if not provided, will be constructed from protocol, host, and port)")
	rootCmd.PersistentFlags().String("protocol", "http", "Protocol to use (http or https)")
	rootCmd.PersistentFlags().String("port", "8080", "Port to use (defaults to 8080 for http, 443 for https)")
	rootCmd.PersistentFlags().String("openapi", "", "Path or URL to OpenAPI specification (required)")
	rootCmd.PersistentFlags().String("environment", "production", "Environment name (e.g., production, staging)")
	rootCmd.PersistentFlags().String("version", "1.0.0", "API version being tested")
	rootCmd.PersistentFlags().Duration("timeout", 30*time.Second, "Request timeout")
	rootCmd.PersistentFlags().String("validation-mode", "minimal", "validation mode (strict, minimal, test-ready)")
	rootCmd.PersistentFlags().String("report-dir", "/tmp/driveby-reports", "report output directory")
	rootCmd.PersistentFlags().String("host", "", "Host of the API to test (required)")

	// Retry flags
	rootCmd.PersistentFlags().Int("retries", 2, "Number of retries for transient HTTP failures (502/503/504)")

	// Authentication flags
	rootCmd.PersistentFlags().String("auth-token", "", "Authentication token (Bearer token)")
	rootCmd.PersistentFlags().String("auth-token-type", "Bearer", "Token type (default: Bearer)")
	rootCmd.PersistentFlags().String("auth-token-header", "Authorization", "Header name for token (default: Authorization)")
	rootCmd.PersistentFlags().String("auth-username", "", "Username for basic authentication")
	rootCmd.PersistentFlags().String("auth-password", "", "Password for basic authentication")
	rootCmd.PersistentFlags().String("auth-api-key", "", "API key for authentication")
	rootCmd.PersistentFlags().String("auth-api-key-header", "X-API-Key", "Header name for API key (default: X-API-Key)")

	// GitHub integration flags
	rootCmd.PersistentFlags().String("github-token", "", "GitHub token for PR commenting (GITHUB_TOKEN env var)")
	rootCmd.PersistentFlags().String("github-owner", "", "GitHub repository owner")
	rootCmd.PersistentFlags().String("github-repo", "", "GitHub repository name")
	rootCmd.PersistentFlags().Int("github-pr-number", 0, "GitHub PR number for commenting")
	rootCmd.PersistentFlags().Bool("github-comment", false, "Enable GitHub PR commenting")

	// GitHub App authentication flags
	rootCmd.PersistentFlags().Int64("github-app-id", 0, "GitHub App ID")
	rootCmd.PersistentFlags().Int64("github-installation-id", 0, "GitHub App Installation ID")
	rootCmd.PersistentFlags().String("github-private-key", "", "GitHub App private key (file path or PEM content)")
	rootCmd.PersistentFlags().String("github-app-slug", "", "GitHub App slug")

	// Load/performance test flags (shared by load-only and test-only)
	for _, cmd := range []*cobra.Command{loadOnlyCmd, testOnlyCmd} {
		cmd.Flags().Duration("max-latency-p95", 500*time.Millisecond, "Maximum allowed P95 latency")
		cmd.Flags().Float64("min-success-rate", 0.99, "Minimum required success rate (0-1)")
		cmd.Flags().Int("concurrent-users", 10, "Number of concurrent users for load testing")
		cmd.Flags().Duration("test-duration", 5*time.Minute, "Duration of load test")
	}

	// Bind flags to viper
	viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))
	viper.BindPFlag("protocol", rootCmd.PersistentFlags().Lookup("protocol"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("openapi", rootCmd.PersistentFlags().Lookup("openapi"))
	viper.BindPFlag("environment", rootCmd.PersistentFlags().Lookup("environment"))
	viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	viper.BindPFlag("validation-mode", rootCmd.PersistentFlags().Lookup("validation-mode"))
	viper.BindPFlag("report-dir", rootCmd.PersistentFlags().Lookup("report-dir"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("retries", rootCmd.PersistentFlags().Lookup("retries"))

	// Bind authentication flags
	viper.BindPFlag("auth-token", rootCmd.PersistentFlags().Lookup("auth-token"))
	viper.BindPFlag("auth-token-type", rootCmd.PersistentFlags().Lookup("auth-token-type"))
	viper.BindPFlag("auth-token-header", rootCmd.PersistentFlags().Lookup("auth-token-header"))
	viper.BindPFlag("auth-username", rootCmd.PersistentFlags().Lookup("auth-username"))
	viper.BindPFlag("auth-password", rootCmd.PersistentFlags().Lookup("auth-password"))
	viper.BindPFlag("auth-api-key", rootCmd.PersistentFlags().Lookup("auth-api-key"))
	viper.BindPFlag("auth-api-key-header", rootCmd.PersistentFlags().Lookup("auth-api-key-header"))

	// Bind GitHub flags
	viper.BindPFlag("github-token", rootCmd.PersistentFlags().Lookup("github-token"))
	viper.BindPFlag("github-owner", rootCmd.PersistentFlags().Lookup("github-owner"))
	viper.BindPFlag("github-repo", rootCmd.PersistentFlags().Lookup("github-repo"))
	viper.BindPFlag("github-pr-number", rootCmd.PersistentFlags().Lookup("github-pr-number"))
	viper.BindPFlag("github-comment", rootCmd.PersistentFlags().Lookup("github-comment"))

	// Bind GitHub App flags
	viper.BindPFlag("github-app-id", rootCmd.PersistentFlags().Lookup("github-app-id"))
	viper.BindPFlag("github-installation-id", rootCmd.PersistentFlags().Lookup("github-installation-id"))
	viper.BindPFlag("github-private-key", rootCmd.PersistentFlags().Lookup("github-private-key"))
	viper.BindPFlag("github-app-slug", rootCmd.PersistentFlags().Lookup("github-app-slug"))

	// Bind load test flags per-command via PreRunE to avoid cross-command overwrites
	bindPerfFlags := func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag("max-latency-p95", cmd.Flags().Lookup("max-latency-p95"))
		viper.BindPFlag("min-success-rate", cmd.Flags().Lookup("min-success-rate"))
		viper.BindPFlag("concurrent-users", cmd.Flags().Lookup("concurrent-users"))
		viper.BindPFlag("test-duration", cmd.Flags().Lookup("test-duration"))
		return nil
	}
	loadOnlyCmd.PreRunE = bindPerfFlags
	testOnlyCmd.PreRunE = bindPerfFlags

	// Add commands
	rootCmd.AddCommand(validateOnlyCmd)
	rootCmd.AddCommand(functionOnlyCmd)
	rootCmd.AddCommand(loadOnlyCmd)
	rootCmd.AddCommand(testOnlyCmd)
	rootCmd.AddCommand(githubStatusCmd)
	rootCmd.AddCommand(githubCommentCmd)
}
