package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/example/driveby/internal/config"
	"github.com/example/driveby/internal/core/models"
	"github.com/sirupsen/logrus"
)

// GitHubServiceImpl implements the GitHubService interface
type GitHubServiceImpl struct {
	config *config.Config
	logger *logrus.Logger
	client *http.Client
}

// NewGitHubService creates a new GitHub service
func NewGitHubService(cfg *config.Config, logger *logrus.Logger) GitHubService {
	return &GitHubServiceImpl{
		config: cfg,
		logger: logger,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// CreateIssue creates a GitHub issue
func (s *GitHubServiceImpl) CreateIssue(ctx context.Context, request *models.GitHubIssueRequest) (*models.GitHubIssueResponse, error) {
	s.logger.WithFields(logrus.Fields{
		"owner": request.Owner,
		"repo":  request.Repository,
		"title": request.Title,
	}).Info("Creating GitHub issue")

	// Check if GitHub token is set
	githubToken := s.config.GitHub.Token
	if githubToken == "" {
		return nil, fmt.Errorf("GitHub token not set")
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"title": request.Title,
		"body":  request.Body,
	}

	// Add labels if provided
	if len(request.Labels) > 0 {
		payload["labels"] = request.Labels
	}

	// Marshal payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal issue payload: %w", err)
	}

	// Create API URL
	apiBaseURL := s.config.GitHub.APIBaseURL
	if apiBaseURL == "" {
		apiBaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues", apiBaseURL, request.Owner, request.Repository)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// Send request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create GitHub issue: %w", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusCreated {
		var errorResp map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err == nil {
			s.logger.WithField("error", errorResp).Error("GitHub API error response")
		}
		return nil, fmt.Errorf("failed to create GitHub issue. Status: %s", resp.Status)
	}

	// Parse response
	var issueResp struct {
		Number int    `json:"number"`
		URL    string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&issueResp); err != nil {
		return nil, fmt.Errorf("failed to parse GitHub issue response: %w", err)
	}

	s.logger.WithFields(logrus.Fields{
		"issue_number": issueResp.Number,
		"issue_url":    issueResp.URL,
	}).Info("GitHub issue created successfully")

	return &models.GitHubIssueResponse{
		IssueNumber: issueResp.Number,
		IssueURL:    issueResp.URL,
	}, nil
}

// GetIssue retrieves a GitHub issue by number
func (s *GitHubServiceImpl) GetIssue(ctx context.Context, owner, repo string, issueNumber int) (*models.GitHubIssueResponse, error) {
	s.logger.WithFields(logrus.Fields{
		"owner":        owner,
		"repo":         repo,
		"issue_number": issueNumber,
	}).Info("Getting GitHub issue")

	// Check if GitHub token is set
	githubToken := s.config.GitHub.Token
	if githubToken == "" {
		return nil, fmt.Errorf("GitHub token not set")
	}

	// Create API URL
	apiBaseURL := s.config.GitHub.APIBaseURL
	if apiBaseURL == "" {
		apiBaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues/%d", apiBaseURL, owner, repo, issueNumber)

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubToken))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// Send request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get GitHub issue: %w", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err == nil {
			s.logger.WithField("error", errorResp).Error("GitHub API error response")
		}
		return nil, fmt.Errorf("failed to get GitHub issue. Status: %s", resp.Status)
	}

	// Parse response
	var issueResp struct {
		Number int    `json:"number"`
		URL    string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&issueResp); err != nil {
		return nil, fmt.Errorf("failed to parse GitHub issue response: %w", err)
	}

	return &models.GitHubIssueResponse{
		IssueNumber: issueResp.Number,
		IssueURL:    issueResp.URL,
	}, nil
}