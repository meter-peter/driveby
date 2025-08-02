# DriveBy Argo Workflows Summary

## 🎯 Overview

This directory contains **3 Argo workflows** designed specifically for **quality gates** in CI/CD pipelines. These workflows leverage DriveBy's enhanced minimal mode and GitHub App authentication for optimal performance and security.

## 📋 The Three Workflows

### 1. Essential Validation Workflow
**File**: `01-essential-validation-workflow.yaml`
- **Purpose**: Validates essential OpenAPI structure (P001 only)
- **Execution Time**: ~10-20 seconds
- **Resource Usage**: Minimal (256Mi RAM, 250m CPU)
- **Use Case**: Fast quality gate before functional testing

### 2. Functional Testing Workflow
**File**: `02-functional-testing-workflow.yaml`
- **Purpose**: Runs functional tests on API endpoints
- **Execution Time**: ~30-60 seconds
- **Resource Usage**: Medium (512Mi RAM, 500m CPU)
- **Use Case**: Quality gate for API functionality

### 3. Performance Testing Workflow
**File**: `03-performance-testing-workflow.yaml`
- **Purpose**: Runs load/performance tests on API endpoints
- **Execution Time**: ~5-10 minutes (configurable)
- **Resource Usage**: High (1Gi RAM, 1000m CPU)
- **Use Case**: Quality gate for API performance

## 🚀 Quick Usage

### Minimal Input Required

All workflows require only these essential parameters:

```bash
# Required for all workflows
openapi-spec-url    # URL to OpenAPI specification
api-host           # API host to test
github-owner       # GitHub repository owner
github-repo        # GitHub repository name
github-pr-number   # GitHub PR number for commenting
github-app-id      # GitHub App ID
github-installation-id  # GitHub App Installation ID
github-private-key # GitHub App private key (base64 encoded)
```

### Example: Essential Validation

```bash
argo submit argo/01-essential-validation-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)"
```

### Example: Functional Testing

```bash
argo submit argo/02-functional-testing-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)"
```

### Example: Performance Testing

```bash
argo submit argo/03-performance-testing-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)" \
  --parameter concurrent-users="20" \
  --parameter test-duration="600" \
  --parameter max-latency-p95="1000"
```

## 🔄 Quality Gate Integration

### Sequential Execution (Recommended)

```bash
# 1. Essential validation (fastest)
VALIDATION_WF=$(argo submit argo/01-essential-validation-workflow.yaml \
  --parameter openapi-spec-url="..." \
  --parameter api-host="..." \
  --parameter github-owner="..." \
  --parameter github-repo="..." \
  --parameter github-pr-number="..." \
  --parameter github-app-id="..." \
  --parameter github-installation-id="..." \
  --parameter github-private-key="..." \
  --wait)

# Check validation result
if [ "$(argo get $VALIDATION_WF -o jsonpath='{.status.phase}')" != "Succeeded" ]; then
  echo "❌ Essential validation failed"
  exit 1
fi

# 2. Functional testing (if validation passes)
FUNCTIONAL_WF=$(argo submit argo/02-functional-testing-workflow.yaml \
  --parameter openapi-spec-url="..." \
  --parameter api-host="..." \
  --parameter github-owner="..." \
  --parameter github-repo="..." \
  --parameter github-pr-number="..." \
  --parameter github-app-id="..." \
  --parameter github-installation-id="..." \
  --parameter github-private-key="..." \
  --wait)

# Check functional test result
if [ "$(argo get $FUNCTIONAL_WF -o jsonpath='{.status.phase}')" != "Succeeded" ]; then
  echo "❌ Functional testing failed"
  exit 1
fi

# 3. Performance testing (if functional tests pass)
PERFORMANCE_WF=$(argo submit argo/03-performance-testing-workflow.yaml \
  --parameter openapi-spec-url="..." \
  --parameter api-host="..." \
  --parameter github-owner="..." \
  --parameter github-repo="..." \
  --parameter github-pr-number="..." \
  --parameter github-app-id="..." \
  --parameter github-installation-id="..." \
  --parameter github-private-key="..." \
  --wait)

# Check performance test result
if [ "$(argo get $PERFORMANCE_WF -o jsonpath='{.status.phase}')" != "Succeeded" ]; then
  echo "❌ Performance testing failed"
  exit 1
fi

echo "✅ All quality gates passed"
```

### Using the Runner Script

```bash
# Set environment variables
export OPENAPI_SPEC_URL="https://api.example.com/openapi.json"
export API_HOST="api.example.com"
export GITHUB_OWNER="your-org"
export GITHUB_REPO="your-repo"
export GITHUB_PR_NUMBER="123"
export GITHUB_APP_ID="123456"
export GITHUB_INSTALLATION_ID="789012"
export GITHUB_PRIVATE_KEY="$(cat private-key.pem | base64)"

# Run all quality gates
./argo/run-quality-gates.sh
```

## 📊 Performance Characteristics

| Workflow | Execution Time | CPU | Memory | Use Case |
|----------|----------------|-----|--------|----------|
| Essential Validation | ~10-20s | 250m | 256Mi | Fast validation |
| Functional Testing | ~30-60s | 500m | 512Mi | API functionality |
| Performance Testing | ~5-10min | 1000m | 1Gi | Load testing |

## 🔧 Configuration Options

### Default Values

| Parameter | Essential | Functional | Performance | Default |
|-----------|-----------|------------|-------------|---------|
| `api-port` | ✅ | ✅ | ✅ | `8080` |
| `api-protocol` | ✅ | ✅ | ✅ | `http` |
| `environment` | ✅ | ✅ | ✅ | `production` |
| `timeout` | ❌ | ✅ | ✅ | `30` |
| `concurrent-users` | ❌ | ❌ | ✅ | `10` |
| `test-duration` | ❌ | ❌ | ✅ | `300` |
| `max-latency-p95` | ❌ | ❌ | ✅ | `500` |
| `min-success-rate` | ❌ | ❌ | ✅ | `0.99` |

### Optional Authentication

```bash
# Bearer token authentication
--parameter auth-token="your-bearer-token"

# API key authentication
--parameter auth-api-key="your-api-key"
```

## 📈 Outputs and Metrics

### Essential Validation Outputs
- `validation-status`: Success/Failure
- `validation-report-path`: Path to validation report

### Functional Testing Outputs
- `functional-test-status`: Success/Failure
- `functional-report-path`: Path to functional test report

### Performance Testing Outputs
- `performance-test-status`: Success/Failure
- `performance-report-path`: Path to performance test report
- `latency-p95`: P95 latency in milliseconds
- `success-rate`: Success rate (0-1)
- `requests-per-second`: Requests per second

## 🚨 Prerequisites

1. **Argo Workflows** installed in your cluster
2. **DriveBy Docker image** available: `driveby:latest`
3. **Persistent Volume Claim** for reports: `driveby-reports-pvc`
4. **GitHub App** configured with repository access

### Create PVC for Reports

```bash
kubectl apply -f - <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: driveby-reports-pvc
  namespace: argo
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
EOF
```

## 📚 Files in This Directory

- `01-essential-validation-workflow.yaml` - Essential validation workflow
- `02-functional-testing-workflow.yaml` - Functional testing workflow
- `03-performance-testing-workflow.yaml` - Performance testing workflow
- `run-quality-gates.sh` - Runner script for all quality gates
- `README.md` - Detailed documentation
- `SUMMARY.md` - This summary document

## 🎯 Key Benefits

1. **Minimal Input**: Only essential parameters required
2. **Fast Execution**: Optimized for CI/CD pipelines
3. **GitHub Integration**: Automatic PR commenting
4. **Quality Gates**: Sequential execution with failure handling
5. **Resource Efficient**: Minimal resource usage
6. **Secure**: GitHub App authentication
7. **Configurable**: Flexible parameters for different environments

## 🔗 Related Documentation

- [DriveBy Minimal Mode Guide](../MINIMAL_MODE_GUIDE.md)
- [DriveBy CLI Usage](../CLI_USAGE.md)
- [Argo Workflows Documentation](https://argoproj.github.io/argo-workflows/)
- [GitHub App Setup Guide](https://docs.github.com/en/apps) 