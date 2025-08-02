#!/bin/bash

echo "Testing DriveBy Authentication Validation"
echo "========================================"

# Test 1: Multiple auth methods (should fail)
echo -e "\n1. Testing multiple auth methods (should fail):"
./driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-token "test-token" \
  --auth-username "test-user" \
  --auth-password "test-pass" \
  2>&1 | head -5

# Test 2: Missing password for basic auth (should fail)
echo -e "\n2. Testing missing password for basic auth (should fail):"
./driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-username "test-user" \
  2>&1 | head -5

# Test 3: Valid bearer token auth (should work)
echo -e "\n3. Testing valid bearer token auth (should work):"
./driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-token "test-token" \
  2>&1 | head -5

# Test 4: Valid API key auth (should work)
echo -e "\n4. Testing valid API key auth (should work):"
./driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-api-key "test-api-key" \
  2>&1 | head -5

# Test 5: Valid basic auth (should work)
echo -e "\n5. Testing valid basic auth (should work):"
./driveby validate-only \
  --openapi openapi.json \
  --host localhost \
  --auth-username "test-user" \
  --auth-password "test-pass" \
  2>&1 | head -5

echo -e "\nAuthentication validation tests completed!" 