package validation

import (
	"fmt"
)

// ValidateAuthConfig validates authentication configuration
func ValidateAuthConfig(auth *AuthConfig) error {
	if auth == nil {
		return nil // No auth is valid
	}

	// Count authentication methods
	authMethods := 0
	if auth.Token != "" {
		authMethods++
	}
	if auth.APIKey != "" {
		authMethods++
	}
	if auth.Username != "" {
		authMethods++
	}

	// Only one authentication method should be used
	if authMethods > 1 {
		return fmt.Errorf("only one authentication method can be specified (token, api-key, or username/password)")
	}

	// Validate specific auth methods
	if auth.Token != "" {
		if auth.TokenType == "" {
			auth.TokenType = "Bearer"
		}
		if auth.TokenHeader == "" {
			auth.TokenHeader = "Authorization"
		}
	} else if auth.APIKey != "" {
		if auth.APIKeyHeader == "" {
			auth.APIKeyHeader = "X-API-Key"
		}
	} else if auth.Username != "" {
		if auth.Password == "" {
			return fmt.Errorf("password is required when username is specified")
		}
	}

	return nil
}

// GetAuthMethod returns the authentication method being used
func GetAuthMethod(auth *AuthConfig) string {
	if auth == nil {
		return "none"
	}
	if auth.Token != "" {
		return "token"
	}
	if auth.APIKey != "" {
		return "api-key"
	}
	if auth.Username != "" {
		return "basic"
	}
	return "none"
}
