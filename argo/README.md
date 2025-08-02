# DriveBy Argo Workflows for Quality Gates

This directory contains three Argo workflows designed for quality gates in CI/CD pipelines. These workflows use the enhanced minimal mode and GitHub App authentication for optimal performance and security.

## 📋 Workflow Overview

### 1. Essential Validation Workflow (`01-essential-validation-workflow.yaml`)
- **Purpose**: Validates essential OpenAPI structure for functional and performance testing
- **Mode**: Minimal validation (P001 only)
- **Use Case**: Quality gate before functional testing
- **Execution Time**: ~10-20 seconds

### 2. Functional Testing Workflow (`02-functional-testing-workflow.yaml`)
- **Purpose**: Runs functional tests on API endpoints
- **Mode**: Pure functional testing
- **Use Case**: Quality gate for API functionality
- **Execution Time**: ~30-60 seconds

### 3. Performance Testing Workflow (`03-performance-testing-workflow.yaml`)
- **Purpose**: Runs load/performance tests on API endpoints
- **Mode**: Pure performance testing
- **Use Case**: Quality gate for API performance
- **Execution Time**: ~5-10 minutes (configurable)

## 🚀 Quick Start

### Prerequisites

1. **Argo Workflows** installed in your cluster
2. **DriveBy Docker image** available: `driveby:latest`
3. **Persistent Volume Claim** for reports: `driveby-reports-pvc`
4. **GitHub App** configured with repository access

### Required Parameters

All workflows require these essential parameters:

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

## 📊 Workflow Usage Examples

### 1. Essential Validation

```bash
# Submit essential validation workflow
argo submit argo/01-essential-validation-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter api-port="443" \
  --parameter api-protocol="https" \
  --parameter environment="staging" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)"
```

### 2. Functional Testing

```bash
# Submit functional testing workflow
argo submit argo/02-functional-testing-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter api-port="443" \
  --parameter api-protocol="https" \
  --parameter environment="staging" \
  --parameter timeout="30" \
  --parameter auth-token="your-bearer-token" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)"
```

### 3. Performance Testing

```bash
# Submit performance testing workflow
argo submit argo/03-performance-testing-workflow.yaml \
  --parameter openapi-spec-url="https://api.example.com/openapi.json" \
  --parameter api-host="api.example.com" \
  --parameter api-port="443" \
  --parameter api-protocol="https" \
  --parameter environment="staging" \
  --parameter timeout="30" \
  --parameter concurrent-users="20" \
  --parameter test-duration="600" \
  --parameter max-latency-p95="1000" \
  --parameter min-success-rate="0.99" \
  --parameter auth-api-key="your-api-key" \
  --parameter github-owner="your-org" \
  --parameter github-repo="your-repo" \
  --parameter github-pr-number="123" \
  --parameter github-app-id="123456" \
  --parameter github-installation-id="789012" \
  --parameter github-private-key="$(cat private-key.pem | base64)"
```

## 🔄 Quality Gate Integration

### Sequential Execution

For comprehensive quality gates, run workflows in sequence:

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

## 🔧 Configuration

### Resource Requirements

| Workflow | CPU Request | CPU Limit | Memory Request | Memory Limit |
|----------|-------------|-----------|----------------|--------------|
| Essential Validation | 250m | 500m | 256Mi | 512Mi |
| Functional Testing | 500m | 1000m | 512Mi | 1Gi |
| Performance Testing | 1000m | 2000m | 1Gi | 2Gi |

### Default Parameters

| Parameter | Essential | Functional | Performance | Default Value |
|-----------|-----------|------------|-------------|---------------|
| `api-port` | ✅ | ✅ | ✅ | `8080` |
| `api-protocol` | ✅ | ✅ | ✅ | `http` |
| `environment` | ✅ | ✅ | ✅ | `production` |
| `timeout` | ❌ | ✅ | ✅ | `30` |
| `concurrent-users` | ❌ | ❌ | ✅ | `10` |
| `test-duration` | ❌ | ❌ | ✅ | `300` |
| `max-latency-p95` | ❌ | ❌ | ✅ | `500` |
| `min-success-rate` | ❌ | ❌ | ✅ | `0.99` |

## 📈 Monitoring and Reporting

### Workflow Status

```bash
# Check workflow status
argo get <workflow-name>

# Get workflow logs
argo logs <workflow-name>

# Get workflow outputs
argo get <workflow-name> -o jsonpath='{.status.outputs.parameters}'
```

### Quality Gate Metrics

```bash
# Essential validation metrics
VALIDATION_STATUS=$(argo get <validation-wf> -o jsonpath='{.status.outputs.parameters[0].value}')
echo "Validation Status: $VALIDATION_STATUS"

# Functional testing metrics
FUNCTIONAL_STATUS=$(argo get <functional-wf> -o jsonpath='{.status.outputs.parameters[0].value}')
echo "Functional Test Status: $FUNCTIONAL_STATUS"

# Performance testing metrics
PERFORMANCE_STATUS=$(argo get <performance-wf> -o jsonpath='{.status.outputs.parameters[0].value}')
LATENCY_P95=$(argo get <performance-wf> -o jsonpath='{.status.outputs.parameters[2].value}')
SUCCESS_RATE=$(argo get <performance-wf> -o jsonpath='{.status.outputs.parameters[3].value}')
RPS=$(argo get <performance-wf> -o jsonpath='{.status.outputs.parameters[4].value}')

echo "Performance Status: $PERFORMANCE_STATUS"
echo "P95 Latency: $LATENCY_P95"
echo "Success Rate: $SUCCESS_RATE"
echo "Requests/Second: $RPS"
```

## 🚨 Troubleshooting

### Common Issues

1. **Workflow Fails with Image Pull Error**
   ```bash
   # Ensure DriveBy image is available
   docker pull driveby:latest
   ```

2. **GitHub App Authentication Fails**
   ```bash
   # Check GitHub App configuration
   echo "App ID: $GITHUB_APP_ID"
   echo "Installation ID: $GITHUB_INSTALLATION_ID"
   echo "Private Key: ${GITHUB_PRIVATE_KEY:0:50}..."
   ```

3. **Persistent Volume Claim Not Found**
   ```bash
   # Create PVC for reports
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

4. **Workflow Times Out**
   ```bash
   # Increase timeout for performance testing
   --parameter test-duration="900"  # 15 minutes
   ```

### Debug Mode

Enable debug logging for troubleshooting:

```bash
# Add debug parameter to any workflow
--parameter log-level="debug"
```

## 📚 Additional Resources

- [Argo Workflows Documentation](https://argoproj.github.io/argo-workflows/)
- [DriveBy Minimal Mode Guide](../MINIMAL_MODE_GUIDE.md)
- [DriveBy CLI Usage](../CLI_USAGE.md)
- [GitHub App Setup Guide](https://docs.github.com/en/apps)
