package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/meter-peter/driveby/internal/logger"
	"github.com/meter-peter/driveby/internal/report"
	"github.com/meter-peter/driveby/internal/validation"
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
			logCfg.Format = "json"   // Force JSON for Kubernetes environment
			logCfg.Output = "stdout" // Force stdout for Kubernetes environment
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
)

var validateOnlyCmd = &cobra.Command{
	Use:   "validate-only",
	Short: "Run only OpenAPI/documentation validation checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			openapiPath = os.Getenv("DRIVEBY_OPENAPI")
			fmt.Fprintf(os.Stderr, "[DEBUG] Fallback: using DRIVEBY_OPENAPI env var: %s\n", openapiPath)
		}
		osOpenapi := os.Getenv("DRIVEBY_OPENAPI")
		fmt.Fprintf(os.Stderr, "[DEBUG] os.Getenv(DRIVEBY_OPENAPI): %s\n", osOpenapi)
		fmt.Fprintf(os.Stderr, "[DEBUG] viper.GetString(openapi): %s\n", viper.GetString("openapi"))
		if openapiPath == "" {
			fmt.Fprintln(os.Stderr, "[ERROR] --openapi flag or DRIVEBY_OPENAPI env variable must be set")
			os.Exit(2)
		}
		// Debug print for openapi path
		fmt.Fprintf(os.Stderr, "[DEBUG] openapi path: %s\n", openapiPath)
		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, viper.GetString("host"), port)
		}

		cfg := validation.ValidatorConfig{
			BaseURL:        baseURL,
			SpecPath:       openapiPath,
			Environment:    viper.GetString("environment"),
			Version:        viper.GetString("version"),
			Timeout:        viper.GetDuration("timeout"),
			ValidationMode: validation.ValidationMode(viper.GetString("validation-mode")),
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
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			openapiPath = os.Getenv("DRIVEBY_OPENAPI")
			fmt.Fprintf(os.Stderr, "[DEBUG] Fallback: using DRIVEBY_OPENAPI env var: %s\n", openapiPath)
		}
		osOpenapi := os.Getenv("DRIVEBY_OPENAPI")
		fmt.Fprintf(os.Stderr, "[DEBUG] os.Getenv(DRIVEBY_OPENAPI): %s\n", osOpenapi)
		fmt.Fprintf(os.Stderr, "[DEBUG] viper.GetString(openapi): %s\n", viper.GetString("openapi"))
		if openapiPath == "" {
			fmt.Fprintln(os.Stderr, "[ERROR] --openapi flag or DRIVEBY_OPENAPI env variable must be set")
			os.Exit(2)
		}
		// Debug print for openapi path
		fmt.Fprintf(os.Stderr, "[DEBUG] openapi path: %s\n", openapiPath)
		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, viper.GetString("host"), port)
		}

		cfg := validation.ValidatorConfig{
			BaseURL:     baseURL,
			SpecPath:    openapiPath,
			Environment: viper.GetString("environment"),
			Version:     viper.GetString("version"),
			Timeout:     viper.GetDuration("timeout"),
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

		// Check if any endpoints failed
		for _, endpoint := range report.TestResults.Functional.EndpointResults {
			if endpoint.Status == validation.TestStatusFailed {
				os.Exit(ExitValidationFailed)
			}
		}
		os.Exit(ExitSuccess)
		return nil
	},
}

var loadOnlyCmd = &cobra.Command{
	Use:   "load-only",
	Short: "Run only load/performance tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		openapiPath := viper.GetString("openapi")
		if openapiPath == "" {
			openapiPath = os.Getenv("DRIVEBY_OPENAPI")
			fmt.Fprintf(os.Stderr, "[DEBUG] Fallback: using DRIVEBY_OPENAPI env var: %s\n", openapiPath)
		}
		osOpenapi := os.Getenv("DRIVEBY_OPENAPI")
		fmt.Fprintf(os.Stderr, "[DEBUG] os.Getenv(DRIVEBY_OPENAPI): %s\n", osOpenapi)
		fmt.Fprintf(os.Stderr, "[DEBUG] viper.GetString(openapi): %s\n", viper.GetString("openapi"))
		if openapiPath == "" {
			fmt.Fprintln(os.Stderr, "[ERROR] --openapi flag or DRIVEBY_OPENAPI env variable must be set")
			os.Exit(2)
		}
		// Debug print for openapi path
		fmt.Fprintf(os.Stderr, "[DEBUG] openapi path: %s\n", openapiPath)
		protocol := viper.GetString("protocol")
		port := viper.GetString("port")
		if protocol == "https" && port == "8080" {
			port = "443"
		}

		baseURL := viper.GetString("api-url")
		if baseURL == "" {
			baseURL = fmt.Sprintf("%s://%s:%s", protocol, viper.GetString("host"), port)
		}

		cfg := validation.ValidatorConfig{
			BaseURL:     baseURL,
			SpecPath:    openapiPath,
			Environment: viper.GetString("environment"),
			Version:     viper.GetString("version"),
			Timeout:     viper.GetDuration("timeout"),
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
	rootCmd.PersistentFlags().String("api-url", "", "Base URL of the API to test")
	rootCmd.PersistentFlags().String("protocol", "http", "Protocol to use (http or https)")
	rootCmd.PersistentFlags().String("port", "8080", "Port to use (defaults to 8080 for http, 443 for https)")
	rootCmd.PersistentFlags().String("openapi", "", "Path or URL to OpenAPI specification")
	rootCmd.PersistentFlags().String("environment", "production", "Environment name (e.g., production, staging)")
	rootCmd.PersistentFlags().String("version", "1.0.0", "API version being tested")
	rootCmd.PersistentFlags().Duration("timeout", 30, "Request timeout in seconds")
	rootCmd.PersistentFlags().String("validation-mode", "minimal", "validation mode (strict, minimal)")
	rootCmd.PersistentFlags().String("report-dir", "/tmp/driveby-reports", "report output directory")
	rootCmd.PersistentFlags().String("host", "", "Host of the API to test")

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

	// Bind load test flags
	viper.BindPFlag("max-latency-p95", loadOnlyCmd.Flags().Lookup("max-latency-p95"))
	viper.BindPFlag("min-success-rate", loadOnlyCmd.Flags().Lookup("min-success-rate"))
	viper.BindPFlag("concurrent-users", loadOnlyCmd.Flags().Lookup("concurrent-users"))
	viper.BindPFlag("test-duration", loadOnlyCmd.Flags().Lookup("test-duration"))

	// Add commands
	rootCmd.AddCommand(validateOnlyCmd)
	rootCmd.AddCommand(functionOnlyCmd)
	rootCmd.AddCommand(loadOnlyCmd)

	// Set up environment variable bindings
	viper.BindEnv("api-url", "DRIVEBY_API_URL")
	viper.BindEnv("protocol", "DRIVEBY_PROTOCOL")
	viper.BindEnv("port", "DRIVEBY_PORT")
	viper.BindEnv("host", "DRIVEBY_HOST")
	viper.BindEnv("openapi", "DRIVEBY_OPENAPI")
	viper.BindEnv("environment", "DRIVEBY_ENVIRONMENT")
	viper.BindEnv("version", "DRIVEBY_VERSION")
	viper.BindEnv("timeout", "DRIVEBY_TIMEOUT")
	viper.BindEnv("validation-mode", "DRIVEBY_VALIDATION_MODE")
	viper.BindEnv("report-dir", "DRIVEBY_REPORT_DIR")
	viper.BindEnv("max-latency-p95", "DRIVEBY_MAX_LATENCY_P95")
	viper.BindEnv("min-success-rate", "DRIVEBY_MIN_SUCCESS_RATE")
	viper.BindEnv("concurrent-users", "DRIVEBY_CONCURRENT_USERS")
	viper.BindEnv("test-duration", "DRIVEBY_TEST_DURATION")

	viper.AutomaticEnv()
}

// logAndExit logs the error and exits with the specified code
func logAndExit(err error, exitCode int) {
	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
		"level": "error",
		"msg":   err.Error(),
	})
	os.Exit(exitCode)
}
