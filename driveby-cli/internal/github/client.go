package github

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v62/github"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
)

// Client represents a GitHub API client
type Client struct {
	client *github.Client
	owner  string
	repo   string
}

// GitHubAppConfig represents GitHub App configuration
type GitHubAppConfig struct {
	AppID          int64  `json:"app_id"`
	InstallationID int64  `json:"installation_id"`
	PrivateKey     string `json:"private_key"`
	AppSlug        string `json:"app_slug"`
}

// NewClient creates a new GitHub client with GitHub App authentication
func NewClient(config *GitHubAppConfig, owner, repo string) (*Client, error) {
	// Generate JWT token for GitHub App
	jwtToken, err := generateJWT(config.AppID, config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %w", err)
	}

	// Get installation token
	installationToken, err := getInstallationToken(jwtToken, config.InstallationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation token: %w", err)
	}

	// Create GitHub client with installation token
	ts := github.BasicAuthTransport{
		Username: "driveby-app",
		Password: installationToken,
	}

	client := github.NewClient(ts.Client())

	return &Client{
		client: client,
		owner:  owner,
		repo:   repo,
	}, nil
}

// NewClientWithToken creates a new GitHub client with personal access token (legacy)
func NewClientWithToken(token, owner, repo string) *Client {
	ts := github.BasicAuthTransport{
		Username: "driveby",
		Password: token,
	}

	client := github.NewClient(ts.Client())

	return &Client{
		client: client,
		owner:  owner,
		repo:   repo,
	}
}

// generateJWT generates a JWT token for GitHub App authentication
func generateJWT(appID int64, privateKeyPEM string) (string, error) {
	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to decode private key PEM")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create JWT claims
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    fmt.Sprintf("%d", appID),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)), // JWT expires in 10 minutes
	}

	// Sign JWT
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return signedToken, nil
}

// getInstallationToken retrieves an installation token from GitHub
func getInstallationToken(jwtToken string, installationID int64) (string, error) {
	url := fmt.Sprintf("https://api.github.com/app/installations/%d/access_tokens", installationID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "driveby-app")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get installation token: %s - %s", resp.Status, string(body))
	}

	var result struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	logrus.Debugf("Retrieved installation token, expires at: %s", result.ExpiresAt)
	return result.Token, nil
}

// CommentOnPR adds a comment to a pull request
func (c *Client) CommentOnPR(ctx context.Context, prNumber int, comment string) error {
	commentBody := &github.IssueComment{
		Body: github.String(comment),
	}

	_, _, err := c.client.Issues.CreateComment(ctx, c.owner, c.repo, prNumber, commentBody)
	if err != nil {
		return fmt.Errorf("failed to create PR comment: %w", err)
	}

	logrus.Infof("Successfully commented on PR #%d", prNumber)
	return nil
}

// CreateValidationComment creates a formatted Markdown comment for validation results
func (c *Client) CreateValidationComment(report interface{}, validationMode string) string {
	var sb strings.Builder

	sb.WriteString("## DriveBy API Validation Report\n\n")
	sb.WriteString(fmt.Sprintf("**Validation Mode:** `%s`\n\n", validationMode))

	switch r := report.(type) {
	case *types.ValidationReport:
		sb.WriteString(fmt.Sprintf("**Status:** %s | **Checks:** %d passed, %d failed\n\n", r.Status, r.PassedChecks, r.FailedChecks))

		if len(r.Principles) > 0 {
			sb.WriteString("### Principle Results\n\n")
			sb.WriteString("| Principle | Severity | Status | Message |\n")
			sb.WriteString("|-----------|----------|--------|---------|\n")
			for _, p := range r.Principles {
				status := "PASS"
				if !p.Passed {
					status = "FAIL"
				}
				sb.WriteString(fmt.Sprintf("| %s: %s | %s | %s | %s |\n",
					p.Principle.ID, p.Principle.Name, p.Principle.Severity, status, p.Message))
			}
			sb.WriteString("\n")
		}

		if r.TestResults != nil && r.TestResults.Functional != nil {
			fr := r.TestResults.Functional
			sb.WriteString("### Functional Test Results\n\n")
			sb.WriteString(fmt.Sprintf("**Endpoints:** %d tested, %d passed, %d failed\n\n", fr.TestedEndpoints, fr.PassedEndpoints, fr.FailedEndpoints))
			if len(fr.EndpointResults) > 0 {
				sb.WriteString("| Method | Path | Status | Code | Response Time |\n")
				sb.WriteString("|--------|------|--------|------|---------------|\n")
				for _, ep := range fr.EndpointResults {
					sb.WriteString(fmt.Sprintf("| %s | %s | %s | %d | %s |\n",
						ep.Method, ep.Path, ep.Status, ep.StatusCode, ep.ResponseTime))
				}
				sb.WriteString("\n")
			}
		}

		if r.TestResults != nil && r.TestResults.Performance != nil {
			pr := r.TestResults.Performance
			sb.WriteString("### Performance Results\n\n")
			sb.WriteString(fmt.Sprintf("| Metric | Value |\n|--------|-------|\n"))
			sb.WriteString(fmt.Sprintf("| P50 Latency | %s |\n", pr.LatencyP50))
			sb.WriteString(fmt.Sprintf("| P95 Latency | %s |\n", pr.LatencyP95))
			sb.WriteString(fmt.Sprintf("| P99 Latency | %s |\n", pr.LatencyP99))
			sb.WriteString(fmt.Sprintf("| Requests/sec | %.2f |\n", pr.RequestsPerSecond))
			sb.WriteString(fmt.Sprintf("| Error Rate | %.2f%% |\n\n", pr.ErrorRate*100))
		}

	default:
		// Fallback: encode as JSON
		if data, err := json.MarshalIndent(report, "", "  "); err == nil {
			sb.WriteString("### Results\n\n```json\n")
			sb.WriteString(string(data))
			sb.WriteString("\n```\n\n")
		}
	}

	sb.WriteString("---\n*Generated by [DriveBy](https://github.com/meter-peter/driveby) API validation framework.*\n")

	return sb.String()
}

// ValidateConfig validates GitHub configuration
func ValidateConfig(token, owner, repo string) error {
	if token == "" {
		return fmt.Errorf("GitHub token is required")
	}
	if owner == "" {
		return fmt.Errorf("GitHub owner is required")
	}
	if repo == "" {
		return fmt.Errorf("GitHub repository is required")
	}
	return nil
}

// ValidateAppConfig validates GitHub App configuration
func ValidateAppConfig(config *GitHubAppConfig, owner, repo string) error {
	if config == nil {
		return fmt.Errorf("GitHub App configuration is required")
	}
	if config.AppID <= 0 {
		return fmt.Errorf("GitHub App ID must be greater than 0")
	}
	if config.InstallationID <= 0 {
		return fmt.Errorf("GitHub Installation ID must be greater than 0")
	}
	if config.PrivateKey == "" {
		return fmt.Errorf("GitHub App private key is required")
	}
	if owner == "" {
		return fmt.Errorf("GitHub owner is required")
	}
	if repo == "" {
		return fmt.Errorf("GitHub repository is required")
	}
	return nil
}
