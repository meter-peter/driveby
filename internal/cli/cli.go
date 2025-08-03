package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/meter-peter/driveby/internal/github"
	"github.com/meter-peter/driveby/internal/logger"
	"github.com/meter-peter/driveby/internal/report"
	"github.com/meter-peter/driveby/internal/validation"
	"github.com/sirupsen/logrus"
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
			// Initialize logger with minimal configuration
			logCfg := logger.DefaultConfig()
			logCfg.Level = viper.GetString("log-level")
			logCfg.Format = "json"
			logCfg.Output = "stdout"
			if err := logger.Configure(logCfg); err != nil {
				return fmt.Errorf("failed to configure logger: %w", err)
			}
			return nil
		},
	}
)

// Exit codes
const (
	ExitSuccess          = 0
	ExitValidationFailed = 1 // Tests ran but failed validation
	ExitExecutionError   = 2 // Error executing tests
	ExitInvalidArgs      = 3 // Invalid command line arguments
)

var validateOnlyCmd = &cobra.Command{
	Use:   "validate-only",
	Short: "Run only OpenAPI/documentation validation checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			return fmt.Errorf("--openapi flag is required")
		}

		host := viper.GetString("host")
		if host == "" {
			return fmt.Errorf("--host flag is required")
		}

		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, host, port)
		}

		// Build authentication configuration
		var authConfig *validation.AuthConfig
		if token := viper.GetString("auth-token"); token != "" {
			authConfig = &validation.AuthConfig{
				Token:       token,
				TokenType:   viper.GetString("auth-token-type"),
				TokenHeader: viper.GetString("auth-token-header"),
			}
		} else if username := viper.GetString("auth-username"); username != "" {
			authConfig = &validation.AuthConfig{
				Username: username,
				Password: viper.GetString("auth-password"),
			}
		} else if apiKey := viper.GetString("auth-api-key"); apiKey != "" {
			authConfig = &validation.AuthConfig{
				APIKey:       apiKey,
				APIKeyHeader: viper.GetString("auth-api-key-header"),
			}
		}

		// Validate authentication configuration
		if authConfig != nil {
			if err := validation.ValidateAuthConfig(authConfig); err != nil {
				return fmt.Errorf("authentication configuration error: %w", err)
			}
		}

		cfg := validation.ValidatorConfig{
			BaseURL:        baseURL,
			SpecPath:       openapiPath,
			Environment:    viper.GetString("environment"),
			Version:        viper.GetString("version"),
			Timeout:        viper.GetDuration("timeout"),
			ValidationMode: validation.ValidationMode(viper.GetString("validation-mode")),
			Auth:           authConfig,
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		validator, err := validation.NewAPIValidator(cfg)
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		report, err := validator.Validate(context.Background())
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		if err := generator.SaveValidationReport(report); err != nil {
			logAndExit(err, ExitExecutionError)
		}
		json.NewEncoder(os.Stdout).Encode(report)

		// GitHub integration
		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(report, viper.GetString("validation-mode")); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		// Check if any critical principles failed
		for _, principle := range report.Principles {
			if !principle.Passed && principle.Principle.Severity == "critical" {
				os.Exit(ExitValidationFailed)
			}
		}
		os.Exit(ExitSuccess)
		return nil
	},
}

var functionOnlyCmd = &cobra.Command{
	Use:   "function-only",
	Short: "Run only functional tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			return fmt.Errorf("--openapi flag is required")
		}

		host := viper.GetString("host")
		if host == "" {
			return fmt.Errorf("--host flag is required")
		}

		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, host, port)
		}

		// Build authentication configuration
		var authConfig *validation.AuthConfig
		if token := viper.GetString("auth-token"); token != "" {
			authConfig = &validation.AuthConfig{
				Token:       token,
				TokenType:   viper.GetString("auth-token-type"),
				TokenHeader: viper.GetString("auth-token-header"),
			}
		} else if username := viper.GetString("auth-username"); username != "" {
			authConfig = &validation.AuthConfig{
				Username: username,
				Password: viper.GetString("auth-password"),
			}
		} else if apiKey := viper.GetString("auth-api-key"); apiKey != "" {
			authConfig = &validation.AuthConfig{
				APIKey:       apiKey,
				APIKeyHeader: viper.GetString("auth-api-key-header"),
			}
		}

		// Validate authentication configuration
		if authConfig != nil {
			if err := validation.ValidateAuthConfig(authConfig); err != nil {
				return fmt.Errorf("authentication configuration error: %w", err)
			}
		}

		cfg := validation.ValidatorConfig{
			BaseURL:     baseURL,
			SpecPath:    openapiPath,
			Environment: viper.GetString("environment"),
			Version:     viper.GetString("version"),
			Timeout:     viper.GetDuration("timeout"),
			Auth:        authConfig,
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		tester := validation.NewFunctionalTester(cfg)
		report, err := tester.TestEndpoints(context.Background())
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		if err := generator.SaveFunctionalTestReport(report); err != nil {
			logAndExit(err, ExitExecutionError)
		}
		json.NewEncoder(os.Stdout).Encode(report)

		// GitHub integration
		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(report, "functional-testing"); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		// Check if any endpoints failed
		if report.TestResults != nil && report.TestResults.Functional != nil {
			for _, endpoint := range report.TestResults.Functional.EndpointResults {
				if endpoint.Status == validation.TestStatusFailed {
					os.Exit(ExitValidationFailed)
				}
			}
		}
		os.Exit(ExitSuccess)
		return nil
	},
}

var testOnlyCmd = &cobra.Command{
	Use:   "test-only",
	Short: "Run functional and performance tests without validation",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			return fmt.Errorf("--openapi flag is required")
		}

		host := viper.GetString("host")
		if host == "" {
			return fmt.Errorf("--host flag is required")
		}

		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, host, port)
		}

		// Build authentication configuration
		var authConfig *validation.AuthConfig
		if token := viper.GetString("auth-token"); token != "" {
			authConfig = &validation.AuthConfig{
				Token:       token,
				TokenType:   viper.GetString("auth-token-type"),
				TokenHeader: viper.GetString("auth-token-header"),
			}
		} else if username := viper.GetString("auth-username"); username != "" {
			authConfig = &validation.AuthConfig{
				Username: username,
				Password: viper.GetString("auth-password"),
			}
		} else if apiKey := viper.GetString("auth-api-key"); apiKey != "" {
			authConfig = &validation.AuthConfig{
				APIKey:       apiKey,
				APIKeyHeader: viper.GetString("auth-api-key-header"),
			}
		}

		// Validate authentication configuration
		if authConfig != nil {
			if err := validation.ValidateAuthConfig(authConfig); err != nil {
				return fmt.Errorf("authentication configuration error: %w", err)
			}
		}

		cfg := validation.ValidatorConfig{
			BaseURL:        baseURL,
			SpecPath:       openapiPath,
			Environment:    viper.GetString("environment"),
			Version:        viper.GetString("version"),
			Timeout:        viper.GetDuration("timeout"),
			ValidationMode: validation.ValidationModeTestOnly,
			Auth:           authConfig,
			PerformanceTarget: &validation.PerformanceTargetConfig{
				MaxLatencyP95:   viper.GetDuration("max-latency-p95"),
				MinSuccessRate:  viper.GetFloat64("min-success-rate"),
				ConcurrentUsers: viper.GetInt("concurrent-users"),
				Duration:        viper.GetDuration("test-duration"),
			},
		}

		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)

		// Run functional tests
		logrus.Info("Running functional tests...")
		functionalTester := validation.NewFunctionalTester(cfg)
		functionalReport, err := functionalTester.TestEndpoints(context.Background())
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}

		// Run performance tests
		logrus.Info("Running performance tests...")
		performanceTester, err := validation.NewPerformanceTester(cfg)
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		performanceReport, err := performanceTester.TestPerformance(context.Background())
		if err != nil {
			logAndExit(err, ExitExecutionError)
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
			logAndExit(err, ExitExecutionError)
		}
		if err := generator.SaveLoadTestReport(performanceReport); err != nil {
			logAndExit(err, ExitExecutionError)
		}

		json.NewEncoder(os.Stdout).Encode(combinedReport)

		// GitHub integration
		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(combinedReport, "test-only"); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}

		// Check if any tests failed
		hasFailures := false
		if functionalReport.TestResults != nil && functionalReport.TestResults.Functional != nil {
			for _, endpoint := range functionalReport.TestResults.Functional.EndpointResults {
				if endpoint.Status == validation.TestStatusFailed {
					hasFailures = true
					break
				}
			}
		}
		if performanceReport.TestResults != nil && performanceReport.TestResults.Performance != nil {
			if performanceReport.TestResults.Performance.Status == validation.TestStatusFailed {
				hasFailures = true
			}
		}

		if hasFailures {
			os.Exit(ExitValidationFailed)
		}
		os.Exit(ExitSuccess)
		return nil
	},
}

var loadOnlyCmd = &cobra.Command{
	Use:   "load-only",
	Short: "Run only load/performance tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			return fmt.Errorf("--openapi flag is required")
		}

		host := viper.GetString("host")
		if host == "" {
			return fmt.Errorf("--host flag is required")
		}

		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, host, port)
		}

		// Build authentication configuration
		var authConfig *validation.AuthConfig
		if token := viper.GetString("auth-token"); token != "" {
			authConfig = &validation.AuthConfig{
				Token:       token,
				TokenType:   viper.GetString("auth-token-type"),
				TokenHeader: viper.GetString("auth-token-header"),
			}
		} else if username := viper.GetString("auth-username"); username != "" {
			authConfig = &validation.AuthConfig{
				Username: username,
				Password: viper.GetString("auth-password"),
			}
		} else if apiKey := viper.GetString("auth-api-key"); apiKey != "" {
			authConfig = &validation.AuthConfig{
				APIKey:       apiKey,
				APIKeyHeader: viper.GetString("auth-api-key-header"),
			}
		}

		// Validate authentication configuration
		if authConfig != nil {
			if err := validation.ValidateAuthConfig(authConfig); err != nil {
				return fmt.Errorf("authentication configuration error: %w", err)
			}
		}

		cfg := validation.ValidatorConfig{
			BaseURL:     baseURL,
			SpecPath:    openapiPath,
			Environment: viper.GetString("environment"),
			Version:     viper.GetString("version"),
			Timeout:     viper.GetDuration("timeout"),
			Auth:        authConfig,
			PerformanceTarget: &validation.PerformanceTargetConfig{
				MaxLatencyP95:   viper.GetDuration("max-latency-p95"),
				MinSuccessRate:  viper.GetFloat64("min-success-rate"),
				ConcurrentUsers: viper.GetInt("concurrent-users"),
				Duration:        viper.GetDuration("test-duration"),
			},
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		tester, err := validation.NewPerformanceTester(cfg)
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		report, err := tester.TestPerformance(context.Background())
		if err != nil {
			logAndExit(err, ExitExecutionError)
		}
		if err := generator.SaveLoadTestReport(report); err != nil {
			logAndExit(err, ExitExecutionError)
		}
		json.NewEncoder(os.Stdout).Encode(report)

		// GitHub integration
		if viper.GetBool("github-comment") {
			if err := handleGitHubComment(report, "load-testing"); err != nil {
				logrus.WithError(err).Warn("Failed to comment on GitHub PR")
			}
		}
		return nil
	},
}

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
	rootCmd.PersistentFlags().Duration("timeout", 30, "Request timeout in seconds")
	rootCmd.PersistentFlags().String("validation-mode", "minimal", "validation mode (strict, minimal)")
	rootCmd.PersistentFlags().String("report-dir", "/tmp/driveby-reports", "report output directory")
	rootCmd.PersistentFlags().String("host", "", "Host of the API to test (required)")

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

	// Load test specific flags
	loadOnlyCmd.Flags().Duration("max-latency-p95", 500, "Maximum allowed P95 latency in milliseconds")
	loadOnlyCmd.Flags().Float64("min-success-rate", 0.99, "Minimum required success rate (0-1)")
	loadOnlyCmd.Flags().Int("concurrent-users", 10, "Number of concurrent users for load testing")
	loadOnlyCmd.Flags().Duration("test-duration", 300, "Duration of load test in seconds")

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

	// Bind load test flags
	viper.BindPFlag("max-latency-p95", loadOnlyCmd.Flags().Lookup("max-latency-p95"))
	viper.BindPFlag("min-success-rate", loadOnlyCmd.Flags().Lookup("min-success-rate"))
	viper.BindPFlag("concurrent-users", loadOnlyCmd.Flags().Lookup("concurrent-users"))
	viper.BindPFlag("test-duration", loadOnlyCmd.Flags().Lookup("test-duration"))

	// Add commands
	rootCmd.AddCommand(validateOnlyCmd)
	rootCmd.AddCommand(functionOnlyCmd)
	rootCmd.AddCommand(loadOnlyCmd)
	rootCmd.AddCommand(testOnlyCmd)
}

// logAndExit logs the error and exits with the specified code
func logAndExit(err error, exitCode int) {
	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
		"level": "error",
		"msg":   err.Error(),
	})
	os.Exit(exitCode)
}

// handleGitHubComment handles GitHub PR commenting
func handleGitHubComment(report interface{}, validationMode string) error {
	owner := viper.GetString("github-owner")
	repo := viper.GetString("github-repo")
	prNumber := viper.GetInt("github-pr-number")

	if prNumber <= 0 {
		return fmt.Errorf("GitHub PR number must be greater than 0")
	}

	var client *github.Client
	var err error

	// Check if GitHub App authentication is configured
	appID := viper.GetInt64("github-app-id")
	installationID := viper.GetInt64("github-installation-id")
	privateKey := viper.GetString("github-private-key")
	appSlug := viper.GetString("github-app-slug")

	if appID > 0 && installationID > 0 && privateKey != "" {
		// Use GitHub App authentication
		appConfig := &github.GitHubAppConfig{
			AppID:          appID,
			InstallationID: installationID,
			PrivateKey:     privateKey,
			AppSlug:        appSlug,
		}

		// Validate GitHub App configuration
		if err := github.ValidateAppConfig(appConfig, owner, repo); err != nil {
			return fmt.Errorf("GitHub App configuration error: %w", err)
		}

		client, err = github.NewClient(appConfig, owner, repo)
		if err != nil {
			return fmt.Errorf("failed to create GitHub App client: %w", err)
		}
	} else {
		// Use legacy token authentication
		token := viper.GetString("github-token")
		if token == "" {
			token = os.Getenv("GITHUB_TOKEN")
		}

		// Validate configuration
		if err := github.ValidateConfig(token, owner, repo); err != nil {
			return fmt.Errorf("GitHub configuration error: %w", err)
		}

		client = github.NewClientWithToken(token, owner, repo)
	}

	// Create comment
	comment := client.CreateValidationComment(report, validationMode)

	// Post comment
	return client.CommentOnPR(context.Background(), prNumber, comment)
}
