package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/meter-peter/driveby/driveby-cli/internal/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var githubStatusCmd = &cobra.Command{
	Use:   "github-status",
	Short: "Set a GitHub commit status using GitHub App or PAT authentication",
	Long: `Set a GitHub commit status on a specific SHA.

Uses GitHub App authentication (preferred) when --github-app-id, --github-installation-id,
and --github-private-key are provided. Falls back to PAT via --github-token or GITHUB_TOKEN.`,
	RunE: runGitHubStatus,
}

func init() {
	githubStatusCmd.Flags().String("sha", "", "Commit SHA to set status on (required)")
	githubStatusCmd.Flags().String("state", "", "Status state: pending, success, failure, error (required)")
	githubStatusCmd.Flags().String("description", "", "Human-readable status description")
	githubStatusCmd.Flags().String("context", "", "Commit status context key (required)")
	githubStatusCmd.Flags().String("target-url", "", "URL to link from the status")

	githubStatusCmd.MarkFlagRequired("sha")
	githubStatusCmd.MarkFlagRequired("state")
	githubStatusCmd.MarkFlagRequired("context")
}

func runGitHubStatus(cmd *cobra.Command, args []string) error {
	sha, _ := cmd.Flags().GetString("sha")
	state, _ := cmd.Flags().GetString("state")
	description, _ := cmd.Flags().GetString("description")
	statusContext, _ := cmd.Flags().GetString("context")
	targetURL, _ := cmd.Flags().GetString("target-url")

	owner, _ := cmd.Flags().GetString("github-owner")
	repo, _ := cmd.Flags().GetString("github-repo")

	// Validate state
	validStates := map[string]bool{"pending": true, "success": true, "failure": true, "error": true}
	if !validStates[state] {
		return fmt.Errorf("invalid state %q: must be one of pending, success, failure, error", state)
	}

	if owner == "" || repo == "" {
		return fmt.Errorf("--github-owner and --github-repo are required")
	}

	// Build GitHub client — prefer App auth, fall back to PAT
	var client *github.Client

	appID, _ := cmd.Flags().GetInt64("github-app-id")
	installationID, _ := cmd.Flags().GetInt64("github-installation-id")
	privateKey, _ := cmd.Flags().GetString("github-private-key")

	if appID > 0 && installationID > 0 && privateKey != "" {
		// Resolve private key: if it looks like a file path, read it
		if !strings.HasPrefix(privateKey, "-----") {
			if data, err := os.ReadFile(privateKey); err == nil {
				privateKey = string(data)
			}
		}

		appConfig := &github.GitHubAppConfig{
			AppID:          appID,
			InstallationID: installationID,
			PrivateKey:     privateKey,
		}

		if err := github.ValidateAppConfig(appConfig, owner, repo); err != nil {
			return fmt.Errorf("GitHub App configuration error: %w", err)
		}

		var err error
		client, err = github.NewClient(appConfig, owner, repo)
		if err != nil {
			return fmt.Errorf("failed to create GitHub App client: %w", err)
		}
		logrus.Info("Using GitHub App authentication")
	} else {
		// Fall back to PAT
		token, _ := cmd.Flags().GetString("github-token")
		if token == "" {
			token = os.Getenv("GITHUB_TOKEN")
		}
		if err := github.ValidateConfig(token, owner, repo); err != nil {
			return fmt.Errorf("GitHub PAT configuration error: %w", err)
		}
		client = github.NewClientWithToken(token, owner, repo)
		logrus.Info("Using GitHub PAT authentication")
	}

	ctx := context.Background()
	if err := client.SetCommitStatus(ctx, sha, state, description, statusContext, targetURL); err != nil {
		return logAndReturnError(err)
	}

	logrus.Infof("Commit status set: sha=%s state=%s context=%s", sha[:7], state, statusContext)
	return nil
}
