package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig  `mapstructure:"server"`
	GitHub   GitHubConfig  `mapstructure:"github"`
	LogLevel string        `mapstructure:"log_level"`
	Testing  TestingConfig `mapstructure:"testing"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
}

// GitHubConfig holds GitHub-related configuration
type GitHubConfig struct {
	Token       string `mapstructure:"token"`
	APIBaseURL  string `mapstructure:"api_base_url"`
	DefaultOrg  string `mapstructure:"default_org"`
	DefaultRepo string `mapstructure:"default_repo"`
}

// TestingConfig holds testing-related configuration
type TestingConfig struct {
	Validation ValidationConfig `mapstructure:"validation"`
	LoadTest   LoadTestConfig   `mapstructure:"load_test"`
	Acceptance AcceptanceConfig `mapstructure:"acceptance"`
}

// ValidationConfig holds validation test configuration
type ValidationConfig struct {
	ComplianceThreshold float64 `mapstructure:"compliance_threshold"`
	FailOnValidation    bool    `mapstructure:"fail_on_validation"`
}

// LoadTestConfig holds load test configuration
type LoadTestConfig struct {
	DefaultRPS      int           `mapstructure:"default_rps"`
	DefaultDuration time.Duration `mapstructure:"default_duration"`
	DefaultTimeout  time.Duration `mapstructure:"default_timeout"`
}

// AcceptanceConfig holds acceptance test configuration
type AcceptanceConfig struct {
	DefaultTimeout time.Duration `mapstructure:"default_timeout"`
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// Set default values
	setDefaults(v)

	// Set config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// Look for config in default locations
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
		v.AddConfigPath("$HOME/.driveby")
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Bind environment variables
	bindEnvVars(v)

	// Unmarshal config
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.base_path", "/api/v1")

	// GitHub defaults
	v.SetDefault("github.api_base_url", "https://api.github.com")

	// Testing defaults
	v.SetDefault("testing.validation.compliance_threshold", 95.0)
	v.SetDefault("testing.validation.fail_on_validation", true)
	v.SetDefault("testing.load_test.default_rps", 10)
	v.SetDefault("testing.load_test.default_duration", "30s")
	v.SetDefault("testing.load_test.default_timeout", "5s")
	v.SetDefault("testing.acceptance.default_timeout", "30s")

	// Logging defaults
	v.SetDefault("log_level", "info")
}

// bindEnvVars binds environment variables to configuration
func bindEnvVars(v *viper.Viper) {
	// Server environment variables
	v.BindEnv("server.host", "DRIVEBY_SERVER_HOST")
	v.BindEnv("server.port", "DRIVEBY_SERVER_PORT")
	v.BindEnv("server.base_path", "DRIVEBY_SERVER_BASE_PATH")

	// GitHub environment variables
	v.BindEnv("github.token", "DRIVEBY_GITHUB_TOKEN")
	v.BindEnv("github.api_base_url", "DRIVEBY_GITHUB_API_BASE_URL")
	v.BindEnv("github.default_org", "DRIVEBY_GITHUB_DEFAULT_ORG")
	v.BindEnv("github.default_repo", "DRIVEBY_GITHUB_DEFAULT_REPO")

	// Testing environment variables
	v.BindEnv("testing.validation.compliance_threshold", "DRIVEBY_VALIDATION_THRESHOLD")
	v.BindEnv("testing.validation.fail_on_validation", "DRIVEBY_FAIL_ON_VALIDATION")
	v.BindEnv("testing.load_test.default_rps", "DRIVEBY_LOAD_TEST_RPS")
	v.BindEnv("testing.load_test.default_duration", "DRIVEBY_LOAD_TEST_DURATION")
	v.BindEnv("testing.load_test.default_timeout", "DRIVEBY_LOAD_TEST_TIMEOUT")
	v.BindEnv("testing.acceptance.default_timeout", "DRIVEBY_ACCEPTANCE_TIMEOUT")

	// Logging environment variables
	v.BindEnv("log_level", "DRIVEBY_LOG_LEVEL")
}
