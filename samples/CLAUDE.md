# Samples Directory — Example Configs and Demo Workflows

## Structure

```
samples/
  configs/
    driveby-minimal.yaml   # Minimal mode configuration (P001 only, fast CI gate)
    driveby-strict.yaml    # Strict mode configuration (all principles, full validation)
  workflows/
    local-demo.sh          # Local demo script: starts perfect-api, runs validation, shows report
  reports/                 # Output directory for generated reports (gitignored)
```

## Configuration Files

### `configs/driveby-minimal.yaml`
Minimal validation configuration -- runs only P001 (OpenAPI Compliance). Intended for:
- Fast CI gate checks (< 5 seconds)
- PR validation where only spec correctness matters
- First-time evaluation of unknown APIs

### `configs/driveby-strict.yaml`
Full strict validation -- runs all implemented principles (P001-P005, P008). Intended for:
- Pre-release quality gates
- Comprehensive API documentation audits
- Thesis evaluation runs

## Demo Workflow

### `workflows/local-demo.sh`
End-to-end local demonstration:
1. Starts the perfect-api via Docker
2. Runs DriveBy CLI in strict mode against it
3. Displays the validation report
4. Cleans up containers

## Thesis Mapping
- **Chapter 5 (Workflow)**: The config files demonstrate DDT usage patterns in CI/CD
- **Chapter 6 (Evaluation)**: The demo workflow is the reproducible evaluation script

## Public API Examples

DriveBy supports validating remote public APIs by URL. Example commands:

```bash
# Validate Swagger Petstore (strict mode)
driveby validate-only \
  --openapi https://petstore3.swagger.io/api/v3/openapi.json \
  --host petstore3.swagger.io \
  --protocol https --port 443 \
  --validation-mode strict

# Batch validate up to 20 public APIs
tools/run-openapi-batch.sh openapis.csv /tmp/driveby-batch
```

See `docs/CLI_USAGE.md` for more examples and `docs/WORKFLOW.md` for the full workflow.

## Rules
- Keep configs minimal and well-commented -- they serve as user-facing documentation
- Reports in `reports/` should be gitignored (generated output)
