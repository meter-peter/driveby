package cli

import (
	"fmt"

	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/spf13/viper"
)

// buildBaseURL constructs the API base URL from flags.
func buildBaseURL() string {
	protocol := viper.GetString("protocol")
	port := viper.GetString("port")
	if protocol == "https" && port == "8080" {
		port = "443"
	}
	baseURL := viper.GetString("api-url")
	if baseURL == "" {
		baseURL = fmt.Sprintf("%s://%s:%s", protocol, viper.GetString("host"), port)
	}
	return baseURL
}

// buildAuthConfig builds authentication configuration from flags.
func buildAuthConfig() *types.AuthConfig {
	if token := viper.GetString("auth-token"); token != "" {
		return &types.AuthConfig{
			Token:       token,
			TokenType:   viper.GetString("auth-token-type"),
			TokenHeader: viper.GetString("auth-token-header"),
		}
	}
	if username := viper.GetString("auth-username"); username != "" {
		return &types.AuthConfig{
			Username: username,
			Password: viper.GetString("auth-password"),
		}
	}
	if apiKey := viper.GetString("auth-api-key"); apiKey != "" {
		return &types.AuthConfig{
			APIKey:       apiKey,
			APIKeyHeader: viper.GetString("auth-api-key-header"),
		}
	}
	return nil
}

// buildValidatorConfig builds the full validator configuration.
func buildValidatorConfig() (types.ValidatorConfig, error) {
	openapiPath := viper.GetString("openapi")
	if openapiPath == "" {
		return types.ValidatorConfig{}, fmt.Errorf("--openapi flag is required")
	}
	host := viper.GetString("host")
	if host == "" {
		return types.ValidatorConfig{}, fmt.Errorf("--host flag is required")
	}

	authConfig := buildAuthConfig()
	if authConfig != nil {
		if err := types.ValidateAuthConfig(authConfig); err != nil {
			return types.ValidatorConfig{}, fmt.Errorf("authentication configuration error: %w", err)
		}
	}

	return types.ValidatorConfig{
		BaseURL:        buildBaseURL(),
		SpecPath:       openapiPath,
		Environment:    viper.GetString("environment"),
		Version:        viper.GetString("version"),
		Timeout:        viper.GetDuration("timeout"),
		ValidationMode: types.ValidationMode(viper.GetString("validation-mode")),
		Auth:           authConfig,
		Retries:        viper.GetInt("retries"),
	}, nil
}

// buildValidatorConfigWithPerf builds config including performance targets.
func buildValidatorConfigWithPerf() (types.ValidatorConfig, error) {
	cfg, err := buildValidatorConfig()
	if err != nil {
		return cfg, err
	}
	cfg.PerformanceTarget = &types.PerformanceTargetConfig{
		MaxLatencyP95:   viper.GetDuration("max-latency-p95"),
		MinSuccessRate:  viper.GetFloat64("min-success-rate"),
		ConcurrentUsers: viper.GetInt("concurrent-users"),
		Duration:        viper.GetDuration("test-duration"),
	}
	return cfg, nil
}
