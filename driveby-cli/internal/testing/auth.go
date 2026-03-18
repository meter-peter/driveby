package testing

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/meter-peter/driveby/driveby-cli/internal/types"
)

// AddAuthHeaders adds authentication headers to the request based on the configured auth method.
func AddAuthHeaders(req *http.Request, auth *types.AuthConfig) error {
	if auth == nil {
		return nil
	}

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
	if authMethods > 1 {
		return fmt.Errorf("only one authentication method can be specified")
	}

	if auth.Token != "" {
		headerName := auth.TokenHeader
		if headerName == "" {
			headerName = "Authorization"
		}
		tokenType := auth.TokenType
		if tokenType == "" {
			tokenType = "Bearer"
		}
		req.Header.Set(headerName, fmt.Sprintf("%s %s", tokenType, auth.Token))
	} else if auth.APIKey != "" {
		headerName := auth.APIKeyHeader
		if headerName == "" {
			headerName = "X-API-Key"
		}
		req.Header.Set(headerName, auth.APIKey)
	} else if auth.Username != "" {
		encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", auth.Username, auth.Password)))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	}

	return nil
}
