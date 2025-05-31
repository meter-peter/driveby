package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/meter-peter/driveby/internal/logger"
	"github.com/meter-peter/driveby/internal/report"
	"github.com/meter-peter/driveby/internal/validation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "driveby",
		Short: "DriveBy - A modern API validation framework",
		Long: `DriveBy is a modern API validation framework that helps you validate, test, and monitor your APIs.
It supports OpenAPI/Swagger specifications and provides comprehensive validation, testing, and rollout capabilities.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Initialize logger with configuration
			logCfg := logger.DefaultConfig()
			if err := viper.UnmarshalKey("logger", &logCfg); err != nil {
				return fmt.Errorf("failed to unmarshal logger config: %w", err)
			}
			if err := logger.Configure(logCfg); err != nil {
				return fmt.Errorf("failed to configure logger: %w", err)
			}
			return nil
		},
	}
)

var validateOnlyCmd = &cobra.Command{
	Use:   "validate-only",
	Short: "Run only OpenAPI/documentation validation checks",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Build OpenAPI URL if needed
		openapiPath := viper.GetString("validation.openapi_path")
		host := viper.GetString("api.host")
		port := viper.GetString("api.port")
		basePath := viper.GetString("api.base_path")
		if openapiPath != "" && !startsWithHTTP(openapiPath) && host != "" {
			proto := "http"
			if port == "443" || port == "8443" {
				proto = "https"
			}
			openapiPath = fmt.Sprintf("%s://%s:%s%s/%s", proto, host, port, basePath, openapiPath)
			openapiPath = strings.Replace(openapiPath, "//", "/", -1) // Clean up double slashes
			if strings.HasPrefix(openapiPath, "http:/") && !strings.HasPrefix(openapiPath, "http://") {
				openapiPath = strings.Replace(openapiPath, "http:/", "http://", 1)
			}
			if strings.HasPrefix(openapiPath, "https:/") && !strings.HasPrefix(openapiPath, "https://") {
				openapiPath = strings.Replace(openapiPath, "https:/", "https://", 1)
			}
		}
		cfg := validation.ValidatorConfig{
			BaseURL:     viper.GetString("api.base_url"),
			SpecPath:    openapiPath,
			LogPath:     viper.GetString("validation.log_path"),
			Environment: viper.GetString("validation.environment"),
			Version:     viper.GetString("validation.version"),
			Timeout:     viper.GetDuration("validation.timeout"),
			PerformanceTarget: validation.PerformanceTargetConfig{
				MaxLatencyP95:  viper.GetDuration("performance.max_latency_p95"),
				MinSuccessRate: viper.GetFloat64("performance.min_success_rate"),
			},
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		validator, err := validation.NewAPIValidator(cfg)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		report, err := validator.Validate(context.Background())
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		err = generator.SaveValidationReport(report)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		json.NewEncoder(os.Stdout).Encode(report)
		return nil
	},
}

var functionOnlyCmd = &cobra.Command{
	Use:   "function-only",
	Short: "Run only endpoint functional tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Build OpenAPI URL if needed
		openapiPath := viper.GetString("validation.openapi_path")
		host := viper.GetString("api.host")
		port := viper.GetString("api.port")
		basePath := viper.GetString("api.base_path")
		if openapiPath != "" && !startsWithHTTP(openapiPath) && host != "" {
			proto := "http"
			if port == "443" || port == "8443" {
				proto = "https"
			}
			openapiPath = fmt.Sprintf("%s://%s:%s%s/%s", proto, host, port, basePath, openapiPath)
			openapiPath = strings.Replace(openapiPath, "//", "/", -1) // Clean up double slashes
			if strings.HasPrefix(openapiPath, "http:/") && !strings.HasPrefix(openapiPath, "http://") {
				openapiPath = strings.Replace(openapiPath, "http:/", "http://", 1)
			}
			if strings.HasPrefix(openapiPath, "https:/") && !strings.HasPrefix(openapiPath, "https://") {
				openapiPath = strings.Replace(openapiPath, "https:/", "https://", 1)
			}
		}
		cfg := validation.ValidatorConfig{
			BaseURL:     viper.GetString("api.base_url"),
			SpecPath:    openapiPath,
			LogPath:     viper.GetString("validation.log_path"),
			Environment: viper.GetString("validation.environment"),
			Version:     viper.GetString("validation.version"),
			Timeout:     viper.GetDuration("validation.timeout"),
			PerformanceTarget: validation.PerformanceTargetConfig{
				MaxLatencyP95:  viper.GetDuration("performance.max_latency_p95"),
				MinSuccessRate: viper.GetFloat64("performance.min_success_rate"),
			},
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		validator, err := validation.NewAPIValidator(cfg)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		report, err := validator.Validate(context.Background())
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		err = generator.SaveValidationReport(report)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		json.NewEncoder(os.Stdout).Encode(report)
		return nil
	},
}

var loadOnlyCmd = &cobra.Command{
	Use:   "load-only",
	Short: "Run only load/performance tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Build OpenAPI URL if needed
		openapiPath := viper.GetString("validation.openapi_path")
		host := viper.GetString("api.host")
		port := viper.GetString("api.port")
		basePath := viper.GetString("api.base_path")
		if openapiPath != "" && !startsWithHTTP(openapiPath) && host != "" {
			proto := "http"
			if port == "443" || port == "8443" {
				proto = "https"
			}
			openapiPath = fmt.Sprintf("%s://%s:%s%s/%s", proto, host, port, basePath, openapiPath)
			openapiPath = strings.Replace(openapiPath, "//", "/", -1) // Clean up double slashes
			if strings.HasPrefix(openapiPath, "http:/") && !strings.HasPrefix(openapiPath, "http://") {
				openapiPath = strings.Replace(openapiPath, "http:/", "http://", 1)
			}
			if strings.HasPrefix(openapiPath, "https:/") && !strings.HasPrefix(openapiPath, "https://") {
				openapiPath = strings.Replace(openapiPath, "https:/", "https://", 1)
			}
		}
		cfg := validation.ValidatorConfig{
			BaseURL:     viper.GetString("api.base_url"),
			SpecPath:    openapiPath,
			LogPath:     viper.GetString("validation.log_path"),
			Environment: viper.GetString("validation.environment"),
			Version:     viper.GetString("validation.version"),
			Timeout:     viper.GetDuration("validation.timeout"),
			PerformanceTarget: validation.PerformanceTargetConfig{
				MaxLatencyP95:  viper.GetDuration("performance.max_latency_p95"),
				MinSuccessRate: viper.GetFloat64("performance.min_success_rate"),
			},
		}
		reportDir := viper.GetString("report-dir")
		generator := report.NewGenerator(reportDir)
		validator, err := validation.NewAPIValidator(cfg)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		report, err := validator.Validate(context.Background())
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
		}
		err = generator.SaveValidationReport(report)
		if err != nil {
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
				"level": "error",
				"msg":   err.Error(),
			})
			os.Exit(1)
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().String("log-format", "json", "log format (json, text)")
	rootCmd.PersistentFlags().String("log-output", "stdout", "log output (stdout, stderr, or file path)")

	// Bind flags to viper
	viper.BindPFlag("logger.level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("logger.format", rootCmd.PersistentFlags().Lookup("log-format"))
	viper.BindPFlag("logger.output", rootCmd.PersistentFlags().Lookup("log-output"))

	// Add commands
	rootCmd.AddCommand(validateOnlyCmd)
	rootCmd.AddCommand(functionOnlyCmd)
	rootCmd.AddCommand(loadOnlyCmd)

	// Add command flags
	validateOnlyCmd.Flags().String("openapi", "", "Path to OpenAPI specification")
	functionOnlyCmd.Flags().String("openapi", "", "Path to OpenAPI specification")
	loadOnlyCmd.Flags().String("openapi", "", "Path to OpenAPI specification")

	// Add a --report-dir flag to the root command, defaulting to ./reports
	rootCmd.PersistentFlags().String("report-dir", "./reports", "report output directory")
	viper.BindPFlag("report-dir", rootCmd.PersistentFlags().Lookup("report-dir"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Fatalf("Error reading config file: %v", err)
		}
		logger.Warn("No config file found, using defaults")
	} else {
		logger.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

// Helper
func startsWithHTTP(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
