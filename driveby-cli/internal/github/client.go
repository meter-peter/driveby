package github

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"sort"
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

// SetCommitStatus creates or updates a commit status on a given SHA
func (c *Client) SetCommitStatus(ctx context.Context, sha, state, description, statusContext, targetURL string) error {
	status := &github.RepoStatus{
		State:       github.String(state),
		Description: github.String(description),
		Context:     github.String(statusContext),
		TargetURL:   github.String(targetURL),
	}
	_, _, err := c.client.Repositories.CreateStatus(ctx, c.owner, c.repo, sha, status)
	if err != nil {
		return fmt.Errorf("failed to create commit status: %w", err)
	}
	logrus.Infof("Successfully set commit status on %s: state=%s context=%s", sha[:7], state, statusContext)
	return nil
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

// CreateValidationComment creates a formatted Markdown comment for validation results.
// gateCtx is optional — nil for standalone CLI, non-nil for XSDLC workflow mode.
func (c *Client) CreateValidationComment(report interface{}, validationMode string, gateCtx *types.GateContext) string {
	var sb strings.Builder

	switch r := report.(type) {
	case *types.ValidationReport:
		writeValidationReportComment(&sb, r, validationMode, gateCtx)
	default:
		sb.WriteString("## DriveBy API Validation Report\n\n")
		sb.WriteString(fmt.Sprintf("**Validation Mode:** `%s`\n\n", validationMode))
		// Fallback: encode as JSON
		if data, err := json.MarshalIndent(report, "", "  "); err == nil {
			sb.WriteString("### Results\n\n```json\n")
			sb.WriteString(string(data))
			sb.WriteString("\n```\n\n")
		}
	}

	sb.WriteString("\n---\n*Generated by [DriveBy](https://github.com/meter-peter/driveby) — Documentation-Driven Testing for the GitOps era.*\n")

	return sb.String()
}

// writeValidationReportComment writes the full enhanced validation report as a GitHub PR comment.
// When gateCtx is non-nil, the comment includes gate-specific headers, details, and "How to Pass" guidance.
func writeValidationReportComment(sb *strings.Builder, r *types.ValidationReport, validationMode string, gateCtx *types.GateContext) {
	// --- Header with score badge ---
	// Status reflects ALL checks: validation principles + functional + performance
	overallFailed := r.Status == "failed"
	if r.TestResults != nil {
		if r.TestResults.Functional != nil && r.TestResults.Functional.FailedEndpoints > 0 {
			overallFailed = true
		}
	}
	// Check if load test failed via P007 principle result or performance metrics
	for _, p := range r.Principles {
		if p.Principle.ID == "P007" && !p.Passed {
			overallFailed = true
		}
	}
	if r.TestResults != nil && r.TestResults.Performance != nil && r.TestResults.Performance.ErrorCount > 0 {
		// Also check if error rate is significant
		if r.TestResults.Performance.ErrorRate > 0.01 {
			overallFailed = true
		}
	}
	statusIcon := "🟢"
	if overallFailed {
		statusIcon = "🔴"
	}
	score := 0
	if r.TotalChecks > 0 {
		score = r.PassedChecks * 100 / r.TotalChecks
	}

	if gateCtx != nil {
		// Gate-aware header
		sb.WriteString(fmt.Sprintf("## %s %s — %s\n\n",
			statusIcon, humanizeGateName(gateCtx.GateName), gateCtx.AppName))
		sb.WriteString(fmt.Sprintf("**%d/%d principles passed (%d%%)** | Mode: `%s` | Checks: `%s`",
			r.PassedChecks, r.TotalChecks, score, validationMode, gateCtx.CheckTypes))
	} else {
		// Standalone header
		sb.WriteString(fmt.Sprintf("## %s DriveBy API Validation Report\n\n", statusIcon))
		sb.WriteString(fmt.Sprintf("**%d/%d principles passed (%d%%)** | Mode: `%s`",
			r.PassedChecks, r.TotalChecks, score, validationMode))
	}
	if r.Version != "" {
		sb.WriteString(fmt.Sprintf(" | Version: `%s`", r.Version))
	}
	if r.Environment != "" {
		sb.WriteString(fmt.Sprintf(" | Env: `%s`", r.Environment))
	}
	sb.WriteString("\n\n")

	// --- Severity breakdown ---
	writeSeverityBreakdown(sb, r)

	// --- Gate details + How to Pass (only in gate mode) ---
	if gateCtx != nil {
		writeGateDetails(sb, gateCtx)
		if overallFailed {
			writeHowToPass(sb, r)
		}
	}

	// --- Principle results: passed summary + failed details ---
	writePrincipleResults(sb, r)

	// --- Quick fixes section ---
	writeQuickFixes(sb, r)

	// --- Functional test results ---
	if r.TestResults != nil && r.TestResults.Functional != nil {
		writeFunctionalResults(sb, r.TestResults.Functional)
	}

	// --- Performance results ---
	if r.TestResults != nil && r.TestResults.Performance != nil {
		writePerformanceResults(sb, r.TestResults.Performance)
	} else {
		// Fallback: extract P007 PerformanceMetrics from PrincipleResult.Details
		if pr := extractPerformanceFromPrinciples(r); pr != nil {
			writePerformanceResults(sb, pr)
		}
	}
}

// humanizeGateName converts "staging-gate" to "Staging Gate".
func humanizeGateName(name string) string {
	parts := strings.Split(name, "-")
	for i, p := range parts {
		parts[i] = capitalizeFirst(p)
	}
	return strings.Join(parts, " ")
}

// writeGateDetails renders a table with quality gate metadata.
func writeGateDetails(sb *strings.Builder, ctx *types.GateContext) {
	sb.WriteString("### Gate Details\n\n")
	sb.WriteString("| Property | Value |\n")
	sb.WriteString("|----------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Environment | `%s` |\n", ctx.Environment))
	sb.WriteString(fmt.Sprintf("| Checks | `%s` |\n", ctx.CheckTypes))
	sb.WriteString(fmt.Sprintf("| Validation Mode | `%s` |\n", ctx.ValidationMode))
	if ctx.WorkflowURL != "" {
		sb.WriteString(fmt.Sprintf("| Workflow | [View Run](%s) |\n", ctx.WorkflowURL))
	}
	sb.WriteString("\n")
}

// writeHowToPass renders actionable guidance for failing gates.
func writeHowToPass(sb *strings.Builder, r *types.ValidationReport) {
	sb.WriteString("### How to Pass This Gate\n\n")

	// Identify critical blockers
	var blockers []string
	for _, p := range r.Principles {
		if !p.Passed && p.Principle.Severity == "critical" {
			fix := p.SuggestedFix
			if fix == "" {
				fix = p.Message
			}
			blockers = append(blockers, fmt.Sprintf("**%s: %s** — %s",
				p.Principle.ID, p.Principle.Name, fix))
		}
	}

	if len(blockers) > 0 {
		sb.WriteString("**Critical blockers** (must fix to unblock promotion):\n\n")
		for _, b := range blockers {
			sb.WriteString(fmt.Sprintf("- 🚨 %s\n", b))
		}
		sb.WriteString("\n")
	}

	// Show warning-level failures (informational, don't block)
	var warnings []string
	for _, p := range r.Principles {
		if !p.Passed && p.Principle.Severity != "critical" {
			warnings = append(warnings, fmt.Sprintf("**%s: %s**", p.Principle.ID, p.Principle.Name))
		}
	}
	if len(warnings) > 0 {
		sb.WriteString("**Warnings** (do not block promotion, but should be addressed):\n\n")
		for _, w := range warnings {
			sb.WriteString(fmt.Sprintf("- ⚠️ %s\n", w))
		}
		sb.WriteString("\n")
	}

	// Check functional/load test failures
	if r.TestResults != nil {
		if r.TestResults.Functional != nil && r.TestResults.Functional.FailedEndpoints > 0 {
			sb.WriteString(fmt.Sprintf("**Functional test failure** (P006): %d/%d endpoints failed — fix implementation to match specification.\n\n",
				r.TestResults.Functional.FailedEndpoints, r.TestResults.Functional.TestedEndpoints))
		}
		if r.TestResults.Performance != nil && r.TestResults.Performance.Status == types.TestStatusFailed {
			sb.WriteString(fmt.Sprintf("**Load test failure** (P007): P95 latency %s exceeded target — optimize API performance.\n\n",
				r.TestResults.Performance.LatencyP95))
		}
	}

	if len(blockers) == 0 && (r.TestResults == nil || (r.TestResults.Functional == nil && r.TestResults.Performance == nil)) {
		sb.WriteString("> Check test results above for specific failures.\n\n")
	}
}

// writeSeverityBreakdown writes severity summary table.
func writeSeverityBreakdown(sb *strings.Builder, r *types.ValidationReport) {
	critPass, critFail := 0, 0
	warnPass, warnFail := 0, 0
	infoPass, infoFail := 0, 0

	for _, p := range r.Principles {
		switch p.Principle.Severity {
		case "critical":
			if p.Passed {
				critPass++
			} else {
				critFail++
			}
		case "warning":
			if p.Passed {
				warnPass++
			} else {
				warnFail++
			}
		default:
			if p.Passed {
				infoPass++
			} else {
				infoFail++
			}
		}
	}

	sb.WriteString("### Summary\n\n")
	sb.WriteString("| Severity | Passed | Failed |\n")
	sb.WriteString("|----------|--------|--------|\n")
	if critPass+critFail > 0 {
		sb.WriteString(fmt.Sprintf("| Critical | %d | %d |\n", critPass, critFail))
	}
	if warnPass+warnFail > 0 {
		sb.WriteString(fmt.Sprintf("| Warning | %d | %d |\n", warnPass, warnFail))
	}
	if infoPass+infoFail > 0 {
		sb.WriteString(fmt.Sprintf("| Info | %d | %d |\n", infoPass, infoFail))
	}
	sb.WriteString("\n")
}

// writePrincipleResults writes passed principles as a compact list and failed ones as expandable details blocks.
func writePrincipleResults(sb *strings.Builder, r *types.ValidationReport) {
	var passed, failed []types.PrincipleResult
	for _, p := range r.Principles {
		if p.Passed {
			passed = append(passed, p)
		} else {
			failed = append(failed, p)
		}
	}

	// Passed principles — compact
	if len(passed) > 0 {
		sb.WriteString("### Passed Principles\n\n")
		for _, p := range passed {
			sb.WriteString(fmt.Sprintf("- ✅ **%s: %s** — %s\n", p.Principle.ID, p.Principle.Name, p.Message))
		}
		sb.WriteString("\n")
	}

	// Failed principles — expandable with full error details
	if len(failed) > 0 {
		sb.WriteString("### Failed Principles\n\n")
		for _, p := range failed {
			severityBadge := "⚠️"
			if p.Principle.Severity == "critical" {
				severityBadge = "🚨"
			}
			sb.WriteString(fmt.Sprintf("<details><summary>%s <strong>%s: %s</strong> (%s)</summary>\n\n",
				severityBadge, p.Principle.ID, p.Principle.Name, p.Principle.Severity))

			// Message
			if p.Message != "" {
				sb.WriteString(fmt.Sprintf("**Result:** %s\n\n", p.Message))
			}

			// Explanation
			if p.Explanation != "" {
				sb.WriteString(fmt.Sprintf("**What went wrong:** %s\n\n", p.Explanation))
			}

			// Details — try to extract actionable error info
			writePrincipleDetails(sb, p.Details)

			// Suggested fix
			if p.SuggestedFix != "" {
				sb.WriteString(fmt.Sprintf("**How to fix:** %s\n\n", p.SuggestedFix))
			}

			// Test impact
			if p.TestImpact != nil {
				writeTestImpact(sb, p.TestImpact)
			}

			sb.WriteString("</details>\n\n")
		}
	}
}

// writePrincipleDetails extracts and formats principle-specific error details.
func writePrincipleDetails(sb *strings.Builder, details interface{}) {
	if details == nil {
		return
	}

	switch d := details.(type) {
	case []types.EndpointValidation:
		if len(d) == 0 {
			return
		}
		sb.WriteString("**Endpoint Issues:**\n\n")
		sb.WriteString("| Method | Path | Status | Errors |\n")
		sb.WriteString("|--------|------|--------|--------|\n")
		for _, ep := range d {
			errStr := "—"
			if len(ep.Errors) > 0 {
				errStr = strings.Join(ep.Errors, "; ")
			}
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s |\n",
				ep.Method, ep.Path, ep.Status, errStr))
		}
		sb.WriteString("\n")

	case *types.EndpointValidationResult:
		if d == nil || len(d.Endpoints) == 0 {
			return
		}
		sb.WriteString("**Endpoint Issues:**\n\n")
		sb.WriteString("| Method | Path | Status | Errors |\n")
		sb.WriteString("|--------|------|--------|--------|\n")
		for _, ep := range d.Endpoints {
			errStr := "—"
			if len(ep.Errors) > 0 {
				errStr = strings.Join(ep.Errors, "; ")
			}
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s |\n",
				ep.Method, ep.Path, ep.Status, errStr))
		}
		sb.WriteString("\n")

	case map[string]interface{}:
		// Generic map — render as key-value pairs or nested errors
		writeMapDetails(sb, d)

	case []interface{}:
		// Generic list — render each item
		if len(d) == 0 {
			return
		}
		sb.WriteString("**Issues found:**\n\n")
		for _, item := range d {
			if m, ok := item.(map[string]interface{}); ok {
				writeDetailItem(sb, m)
			} else {
				sb.WriteString(fmt.Sprintf("- %v\n", item))
			}
		}
		sb.WriteString("\n")

	default:
		// Last resort: JSON block
		if data, err := json.MarshalIndent(details, "", "  "); err == nil {
			rendered := string(data)
			if rendered != "null" && rendered != "{}" && rendered != "[]" {
				sb.WriteString("<details><summary>Raw details</summary>\n\n```json\n")
				sb.WriteString(rendered)
				sb.WriteString("\n```\n\n</details>\n\n")
			}
		}
	}
}

// writeMapDetails renders a map[string]interface{} as a readable list of issues.
func writeMapDetails(sb *strings.Builder, m map[string]interface{}) {
	if len(m) == 0 {
		return
	}

	rendered := false

	// DDT principle checker output: "checks" map with pass/fail booleans
	if checks, ok := m["checks"].(map[string]interface{}); ok {
		messages, _ := m["messages"].(map[string]interface{})

		var failed []string
		var passed []string
		for name, val := range checks {
			if b, ok := val.(bool); ok {
				if b {
					passed = append(passed, name)
				} else {
					failed = append(failed, name)
				}
			}
		}

		sort.Strings(failed)
		sort.Strings(passed)

		if len(failed) > 0 {
			sb.WriteString("**Failed checks:**\n\n")
			for _, name := range failed {
				sb.WriteString(fmt.Sprintf("- ❌ %s", name))
				if messages != nil {
					if msg, ok := messages[name].(string); ok && msg != "" {
						sb.WriteString(fmt.Sprintf(" — %s", msg))
					}
				}
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}

		if len(passed) > 0 {
			sb.WriteString("<details><summary>Passed checks (" + fmt.Sprintf("%d", len(passed)) + ")</summary>\n\n")
			for _, name := range passed {
				sb.WriteString(fmt.Sprintf("- ✅ %s\n", name))
			}
			sb.WriteString("\n</details>\n\n")
		}

		// Render any additional detail keys (missing_docs, missing_errors, etc.)
		for key, val := range m {
			if key == "checks" || key == "messages" {
				continue
			}
			writeDetailKey(sb, key, val)
		}

		rendered = true
	}

	// Legacy/generic: look for known error list keys
	if !rendered {
		for _, key := range []string{"errors", "issues", "failures", "missing", "violations"} {
			if items, ok := m[key]; ok {
				if list, ok := items.([]interface{}); ok && len(list) > 0 {
					rendered = true
					sb.WriteString(fmt.Sprintf("**%s:**\n\n", capitalizeFirst(key)))
					for _, item := range list {
						if s, ok := item.(string); ok {
							sb.WriteString(fmt.Sprintf("- %s\n", s))
						} else if im, ok := item.(map[string]interface{}); ok {
							writeDetailItem(sb, im)
						} else {
							sb.WriteString(fmt.Sprintf("- %v\n", item))
						}
					}
					sb.WriteString("\n")
				}
			}
		}
	}

	if !rendered {
		// Last resort: JSON block
		if data, err := json.MarshalIndent(m, "", "  "); err == nil {
			sb.WriteString("<details><summary>Raw details</summary>\n\n```json\n")
			sb.WriteString(string(data))
			sb.WriteString("\n```\n\n</details>\n\n")
		}
	}
}

// writeDetailKey renders a single detail key (like missing_docs, missing_errors) as a readable list.
func writeDetailKey(sb *strings.Builder, key string, val interface{}) {
	label := strings.ReplaceAll(key, "_", " ")
	label = capitalizeFirst(label)

	switch v := val.(type) {
	case map[string]interface{}:
		// Map of check name → list of items (e.g., missing_docs: {"check": ["item1", ...]})
		hasContent := false
		for checkName, items := range v {
			if list, ok := items.([]interface{}); ok && len(list) > 0 {
				if !hasContent {
					sb.WriteString(fmt.Sprintf("**%s:**\n\n", label))
					hasContent = true
				}
				sb.WriteString(fmt.Sprintf("_%s:_\n", checkName))
				for _, item := range list {
					sb.WriteString(fmt.Sprintf("  - %v\n", item))
				}
			}
		}
		if hasContent {
			sb.WriteString("\n")
		}
	case []interface{}:
		if len(v) > 0 {
			sb.WriteString(fmt.Sprintf("**%s:**\n\n", label))
			for _, item := range v {
				sb.WriteString(fmt.Sprintf("- %v\n", item))
			}
			sb.WriteString("\n")
		}
	}
}

// writeDetailItem renders a single detail map item as a bullet point.
func writeDetailItem(sb *strings.Builder, m map[string]interface{}) {
	// Try common patterns: path+method, message, error
	parts := []string{}
	if method, ok := m["method"].(string); ok {
		if path, ok := m["path"].(string); ok {
			parts = append(parts, fmt.Sprintf("`%s %s`", method, path))
		}
	}
	for _, key := range []string{"message", "error", "description", "reason"} {
		if v, ok := m[key].(string); ok && v != "" {
			parts = append(parts, v)
			break
		}
	}
	if len(parts) > 0 {
		sb.WriteString(fmt.Sprintf("- %s\n", strings.Join(parts, " — ")))
	} else {
		if data, err := json.Marshal(m); err == nil {
			sb.WriteString(fmt.Sprintf("- `%s`\n", string(data)))
		}
	}
}

// writeTestImpact renders test impact information.
func writeTestImpact(sb *strings.Builder, impact *types.TestImpact) {
	if impact == nil {
		return
	}
	icon := "ℹ️"
	switch impact.ImpactLevel {
	case types.ImpactLevelHigh, types.ImpactLevelCritical:
		icon = "🚨"
	case types.ImpactLevelMedium:
		icon = "⚠️"
	}
	sb.WriteString(fmt.Sprintf("%s **Test Impact:** %s\n", icon, impact.ImpactLevel))
	if len(impact.Recommendations) > 0 {
		for _, rec := range impact.Recommendations {
			sb.WriteString(fmt.Sprintf("  - %s\n", rec))
		}
	}
	sb.WriteString("\n")
}

// writeQuickFixes aggregates all suggested fixes into a single actionable section.
func writeQuickFixes(sb *strings.Builder, r *types.ValidationReport) {
	var fixes []string
	for _, p := range r.Principles {
		if !p.Passed && p.SuggestedFix != "" {
			fixes = append(fixes, fmt.Sprintf("- **%s:** %s", p.Principle.ID, p.SuggestedFix))
		}
	}
	if len(fixes) == 0 {
		return
	}
	sb.WriteString("<details><summary>📋 <strong>Quick Fixes</strong></summary>\n\n")
	for _, fix := range fixes {
		sb.WriteString(fix + "\n")
	}
	sb.WriteString("\n</details>\n\n")
}

// writeFunctionalResults writes functional test results with error details.
func writeFunctionalResults(sb *strings.Builder, fr *types.FunctionalTestResults) {
	sb.WriteString("### Functional Test Results\n\n")
	sb.WriteString(fmt.Sprintf("**Endpoints:** %d tested, %d passed, %d failed",
		fr.TestedEndpoints, fr.PassedEndpoints, fr.FailedEndpoints))
	if fr.SkippedEndpoints > 0 {
		sb.WriteString(fmt.Sprintf(", %d skipped", fr.SkippedEndpoints))
	}
	sb.WriteString("\n")
	if fr.AverageResponseTime > 0 {
		sb.WriteString(fmt.Sprintf("**Avg Response Time:** %s | **Min:** %s | **Max:** %s\n",
			fr.AverageResponseTime, fr.MinResponseTime, fr.MaxResponseTime))
	}
	sb.WriteString("\n")

	if len(fr.EndpointResults) == 0 {
		return
	}

	// Separate failures from successes
	var failures, successes []types.EndpointTestResult
	for _, ep := range fr.EndpointResults {
		if ep.Status == types.TestStatusFailed || len(ep.Errors) > 0 {
			failures = append(failures, ep)
		} else {
			successes = append(successes, ep)
		}
	}

	// Failed endpoints — always shown
	if len(failures) > 0 {
		sb.WriteString("**Failed Endpoints:**\n\n")
		sb.WriteString("| Method | Path | Code | Response Time | Errors |\n")
		sb.WriteString("|--------|------|------|---------------|--------|\n")
		for _, ep := range failures {
			errStr := "—"
			if len(ep.Errors) > 0 {
				errStr = strings.Join(ep.Errors, "; ")
			}
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %d | %s | %s |\n",
				ep.Method, ep.Path, ep.StatusCode, ep.ResponseTime, errStr))
		}
		sb.WriteString("\n")

		// Warnings on failed endpoints
		for _, ep := range failures {
			if len(ep.Warnings) > 0 {
				sb.WriteString(fmt.Sprintf("⚠️ `%s %s`: %s\n", ep.Method, ep.Path, strings.Join(ep.Warnings, "; ")))
			}
		}
	}

	// Passed endpoints — collapsible
	if len(successes) > 0 {
		sb.WriteString(fmt.Sprintf("<details><summary>✅ <strong>%d Passed Endpoints</strong></summary>\n\n", len(successes)))
		sb.WriteString("| Method | Path | Code | Response Time |\n")
		sb.WriteString("|--------|------|------|---------------|\n")
		for _, ep := range successes {
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %d | %s |\n",
				ep.Method, ep.Path, ep.StatusCode, ep.ResponseTime))
		}
		sb.WriteString("\n</details>\n\n")
	}
}

// writePerformanceResults writes performance test results with metrics.
func writePerformanceResults(sb *strings.Builder, pr *types.PerformanceTestResults) {
	sb.WriteString("### Performance Results\n\n")

	sb.WriteString("| Metric | Value |\n")
	sb.WriteString("|--------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Total Requests | %d |\n", pr.TotalRequests))
	sb.WriteString(fmt.Sprintf("| Success Rate | %.1f%% |\n", (1-pr.ErrorRate)*100))
	sb.WriteString(fmt.Sprintf("| Error Rate | %.2f%% |\n", pr.ErrorRate*100))
	sb.WriteString(fmt.Sprintf("| Requests/sec | %.2f |\n", pr.RequestsPerSecond))
	sb.WriteString(fmt.Sprintf("| P50 Latency | %s |\n", pr.LatencyP50))
	sb.WriteString(fmt.Sprintf("| P95 Latency | %s |\n", pr.LatencyP95))
	sb.WriteString(fmt.Sprintf("| P99 Latency | %s |\n", pr.LatencyP99))
	if pr.Duration > 0 {
		sb.WriteString(fmt.Sprintf("| Duration | %s |\n", pr.Duration))
	}
	sb.WriteString("\n")

	// Failed requests — show if any
	if len(pr.FailedRequests) > 0 {
		limit := len(pr.FailedRequests)
		if limit > 10 {
			limit = 10
		}
		sb.WriteString(fmt.Sprintf("<details><summary>❌ <strong>%d Failed Requests</strong> (showing %d)</summary>\n\n",
			len(pr.FailedRequests), limit))
		sb.WriteString("| Method | Path | Code | Error | Latency |\n")
		sb.WriteString("|--------|------|------|-------|---------|\n")
		for _, fr := range pr.FailedRequests[:limit] {
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %d | %s | %s |\n",
				fr.Method, fr.Path, fr.StatusCode, fr.Error, fr.Latency))
		}
		sb.WriteString("\n</details>\n\n")
	}
}

// extractPerformanceFromPrinciples scans P007 Details for *PerformanceMetrics
// and converts it to *PerformanceTestResults for comment rendering.
func extractPerformanceFromPrinciples(r *types.ValidationReport) *types.PerformanceTestResults {
	for _, p := range r.Principles {
		if p.Principle.ID != "P007" {
			continue
		}
		if pm, ok := p.Details.(*types.PerformanceMetrics); ok {
			return &types.PerformanceTestResults{
				TotalRequests:     int64(pm.TotalRequests),
				SuccessCount:      int64(pm.SuccessCount),
				ErrorCount:        int64(pm.ErrorCount),
				ErrorRate:         pm.ErrorRate,
				LatencyP50:        pm.LatencyP50,
				LatencyP95:        pm.LatencyP95,
				LatencyP99:        pm.LatencyP99,
				RequestsPerSecond: pm.RequestsPerSec,
				Duration:          pm.EndTime.Sub(pm.StartTime),
			}
		}
	}
	return nil
}

// capitalizeFirst returns the string with the first letter uppercased.
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
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
