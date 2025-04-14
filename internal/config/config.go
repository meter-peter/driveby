package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Redis    RedisConfig    `mapstructure:"redis"`
	GitHub   GitHubConfig   `mapstructure:"github"`
	Minio    MinioConfig    `mapstructure:"minio"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Testing  TestingConfig  `mapstructure:"testing"`
	Features FeaturesConfig `mapstructure:"features"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Mode            string        `mapstructure:"mode"` // "debug", "release", "test"
	Timeout         time.Duration `mapstructure:"timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

// RedisConfig holds Redis-related configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Enabled  bool   `mapstructure:"enabled"`
}

// GitHubConfig holds GitHub-related configuration
type GitHubConfig struct {
	APIBaseURL  string `mapstructure:"api_base_url"`
	DefaultOrg  string `mapstructure:"default_org"`
	DefaultRepo string `mapstructure:"default_repo"`
	Token       string `mapstructure:"token"`
}

// MinioConfig holds Minio storage configuration
type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UseSSL          bool   `mapstructure:"use_ssl"`
	BucketName      string `mapstructure:"bucket_name"`
	Region          string `mapstructure:"region"`
	Enabled         bool   `mapstructure:"enabled"`
}

// LoggingConfig holds logging-related configuration
type LoggingConfig struct {
	Level  string `mapstructure:"level"`  // "debug", "info", "warn", "error"
	Format string `mapstructure:"format"` // "json", "text"
}

// TestingConfig holds testing-related configuration
type TestingConfig struct {
	Validation ValidationConfig `mapstructure:"validation"`
	LoadTest   LoadTestConfig   `mapstructure:"load_test"`
	Acceptance AcceptanceConfig `mapstructure:"acceptance"`
}

// ValidationConfig holds validation-related configuration
type ValidationConfig struct {
	ComplianceThreshold float64 `mapstructure:"compliance_threshold"`
	FailOnValidation    bool    `mapstructure:"fail_on_validation"`
}

// LoadTestConfig holds load test-related configuration
type LoadTestConfig struct {
	DefaultRPS      int           `mapstructure:"default_rps"`
	DefaultDuration time.Duration `mapstructure:"default_duration"`
	DefaultTimeout  time.Duration `mapstructure:"default_timeout"`
}

// AcceptanceConfig holds acceptance test-related configuration
type AcceptanceConfig struct {
	DefaultTimeout time.Duration `mapstructure:"default_timeout"`
}

// FeaturesConfig holds feature flags
type FeaturesConfig struct {
	EnableValidation bool `mapstructure:"enable_validation"`
	EnableLoadTest   bool `mapstructure:"enable_load_test"`
	EnableAcceptance bool `mapstructure:"enable_acceptance"`
	EnableGitHub     bool `mapstructure:"enable_github"`
	EnableWorkers    bool `mapstructure:"enable_workers"`
}

// Load loads the configuration from environment variables and configuration files
func Load() (*Config, error) {
	v := viper.New()

	// Set defaults
	setDefaults(v)

	// Read from config file
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("/etc/driveby")

	// Read environment variables
	v.AutomaticEnv()
	v.SetEnvPrefix("DRIVEBY")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Explicitly bind critical connection environment variables
	v.BindEnv("redis.host", "DRIVEBY_REDIS_HOST")
	v.BindEnv("redis.port", "DRIVEBY_REDIS_PORT")
	v.BindEnv("minio.endpoint", "DRIVEBY_MINIO_ENDPOINT")
	v.BindEnv("minio.access_key_id", "DRIVEBY_MINIO_ACCESS_KEY_ID")
	v.BindEnv("minio.secret_access_key", "DRIVEBY_MINIO_SECRET_ACCESS_KEY")

	// Read config file (optional)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		// Config file not found, using defaults and environment variables
	}

	// Unmarshal config
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// setDefaults sets default values for the configuration
func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.mode", "release")
	v.SetDefault("server.timeout", 30*time.Second)
	v.SetDefault("server.shutdown_timeout", 10*time.Second)

	// Redis defaults
	v.SetDefault("redis.host", "redis")
	v.SetDefault("redis.port", 6379)
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)
	v.SetDefault("redis.enabled", true)

	// GitHub defaults
	v.SetDefault("github.api_base_url", "https://api.github.com")
	v.SetDefault("github.default_org", "")
	v.SetDefault("github.default_repo", "")
	v.SetDefault("github.token", "")

	// Minio defaults
	v.SetDefault("minio.endpoint", "localhost:9000")
	v.SetDefault("minio.access_key_id", "minioadmin")
	v.SetDefault("minio.secret_access_key", "minioadmin")
	v.SetDefault("minio.use_ssl", false)
	v.SetDefault("minio.bucket_name", "driveby")
	v.SetDefault("minio.region", "us-east-1")
	v.SetDefault("minio.enabled", true)

	// Logging defaults
	v.SetDefault("logging.level", "info")
	v.SetDefault("logging.format", "json")

	// Validation defaults
	v.SetDefault("testing.validation.compliance_threshold", 95.0)
	v.SetDefault("testing.validation.fail_on_validation", true)

	// Load test defaults
	v.SetDefault("testing.load_test.default_rps", 10)
	v.SetDefault("testing.load_test.default_duration", 30*time.Second)
	v.SetDefault("testing.load_test.default_timeout", 5*time.Second)

	// Acceptance test defaults
	v.SetDefault("testing.acceptance.default_timeout", 30*time.Second)

	// Feature flags defaults
	v.SetDefault("features.enable_validation", true)
	v.SetDefault("features.enable_load_test", true)
	v.SetDefault("features.enable_acceptance", true)
	v.SetDefault("features.enable_github", true)
	v.SetDefault("features.enable_workers", true)
}
