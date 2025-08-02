#!/bin/bash

# Test script for DriveBy minimal mode and GitHub App authentication
# This script demonstrates the new features for functional and performance testing

echo "🚀 DriveBy Enhanced Testing Script"
echo "=================================="

# Configuration
OPENAPI_SPEC="./openapi.json"
API_HOST="localhost"
API_PORT="8080"
REPORT_DIR="./reports"

# Create reports directory
mkdir -p "$REPORT_DIR"

echo ""
echo "1. Testing Minimal Mode (Essential validation only)"
echo "---------------------------------------------------"
./driveby validate-only \
  --openapi "$OPENAPI_SPEC" \
  --host "$API_HOST" \
  --port "$API_PORT" \
  --validation-mode minimal \
  --report-dir "$REPORT_DIR" \
  --log-level info

echo ""
echo "2. Testing Test-Only Mode (No validation, pure testing)"
echo "-------------------------------------------------------"
./driveby test-only \
  --openapi "$OPENAPI_SPEC" \
  --host "$API_HOST" \
  --port "$API_PORT" \
  --report-dir "$REPORT_DIR" \
  --concurrent-users 5 \
  --test-duration 30 \
  --max-latency-p95 1000 \
  --log-level info

echo ""
echo "3. Testing with GitHub App Authentication (if configured)"
echo "--------------------------------------------------------"
if [ -n "$GITHUB_APP_ID" ] && [ -n "$GITHUB_INSTALLATION_ID" ] && [ -n "$GITHUB_PRIVATE_KEY" ]; then
    echo "GitHub App credentials detected, testing with GitHub integration..."
    ./driveby test-only \
      --openapi "$OPENAPI_SPEC" \
      --host "$API_HOST" \
      --port "$API_PORT" \
      --github-comment \
      --github-app-id "$GITHUB_APP_ID" \
      --github-installation-id "$GITHUB_INSTALLATION_ID" \
      --github-private-key "$GITHUB_PRIVATE_KEY" \
      --github-owner "$GITHUB_OWNER" \
      --github-repo "$GITHUB_REPO" \
      --github-pr-number "$GITHUB_PR_NUMBER" \
      --report-dir "$REPORT_DIR" \
      --log-level info
else
    echo "GitHub App credentials not configured. Set the following environment variables:"
    echo "  GITHUB_APP_ID"
    echo "  GITHUB_INSTALLATION_ID" 
    echo "  GITHUB_PRIVATE_KEY"
    echo "  GITHUB_OWNER"
    echo "  GITHUB_REPO"
    echo "  GITHUB_PR_NUMBER"
fi

echo ""
echo "4. Testing with Legacy GitHub Token (if configured)"
echo "--------------------------------------------------"
if [ -n "$GITHUB_TOKEN" ]; then
    echo "GitHub token detected, testing with legacy authentication..."
    ./driveby test-only \
      --openapi "$OPENAPI_SPEC" \
      --host "$API_HOST" \
      --port "$API_PORT" \
      --github-comment \
      --github-token "$GITHUB_TOKEN" \
      --github-owner "$GITHUB_OWNER" \
      --github-repo "$GITHUB_REPO" \
      --github-pr-number "$GITHUB_PR_NUMBER" \
      --report-dir "$REPORT_DIR" \
      --log-level info
else
    echo "GitHub token not configured. Set GITHUB_TOKEN environment variable for legacy authentication."
fi

echo ""
echo "✅ Testing completed! Check reports in: $REPORT_DIR"
echo ""
echo "Summary of new features:"
echo "- Minimal mode: Essential validation only (P001)"
echo "- Test-only mode: Pure functional and performance testing"
echo "- GitHub App authentication: Secure, granular permissions"
echo "- Legacy token support: Backward compatibility" 