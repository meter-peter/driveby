package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/meter-peter/driveby/driveby-cli/internal/github"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/spf13/viper"
)

// logAndReturnError logs the error as JSON to stderr and returns it.
func logAndReturnError(err error) error {
	json.NewEncoder(os.Stderr).Encode(map[string]interface{}{
		"level": "error",
		"msg":   err.Error(),
	})
	return err
}

// handleGitHubComment handles GitHub PR commenting.
// gateCtx is optional — nil for standalone CLI, non-nil for XSDLC workflow mode.
func handleGitHubComment(report interface{}, validationMode string, gateCtx *types.GateContext) error {
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
	comment := client.CreateValidationComment(report, validationMode, gateCtx)

	// Post comment
	return client.CommentOnPR(context.Background(), prNumber, comment)
}
