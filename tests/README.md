# DriveBy Test Suites

This folder groups higher-level test suites and scenarios for exercising DriveBy itself (not your product APIs). Each suite documents its purpose, scope, and how it should be run in CI/CD or Kubernetes/Argo workflows.

## 1. Public OpenAPI Batch (Max 20 Specs, includes Petstore)

- **Location**: `tools/run-openapi-batch.sh`
- **Type**: Corpus / compatibility testing
- **Limit**: Maximum of 20 OpenAPI specs per run (enforced by the script).

### Purpose

- Exercise DriveBy against a small, curated batch of real-world public OpenAPI specs (including the Swagger Petstore). Specs are expected to be OpenAPI 3.x documents.
- Discover schema shapes, edge cases, and performance characteristics that differ from internal APIs.
- Harden validation rules and error handling before integrating DriveBy into production quality gates.

### Inputs

- A CSV file with at most 20 rows (excluding the header) in the format:

```text
name,url,host
petstore,https://petstore3.swagger.io/api/v3/openapi.json,petstore3.swagger.io
github,https://example.com/github-openapi.json,api.github.com
```

### How to Run Locally or in CI

```bash
tools/run-openapi-batch.sh tests/openapis-sample.csv /tmp/driveby-openapi-batch
```

This will:

- Run `driveby validate-only --validation-mode minimal` once per row.
- Store reports under `/tmp/driveby-openapi-batch/<name>-<timestamp>/`.
- Generate a summary CSV at `/tmp/driveby-openapi-batch/summary-<timestamp>.csv`.

### Kubernetes / Argo Workflow Usage (Concept)

- Mount or inject the CSV file as a ConfigMap or artifact.
- Run a single pod using the official DriveBy image.
- Use the container command:

```bash
tools/run-openapi-batch.sh /data/openapis.csv /tmp/driveby-openapi-batch
```

- Archive the `/tmp/driveby-openapi-batch` directory as workflow artifacts for later analysis.

## 2. Public OpenAPI Endpoint Probing (Spec + Live Behavior)

- **Location**: `tools/probe-openapi-endpoints.py`
- **Type**: Endpoint probing / contract sanity check

### Purpose

- From a small CSV of candidate public OpenAPI specs, automatically keep only those where:
  - The spec is reachable as JSON.
  - At least one simple unauthenticated `GET` endpoint returns:
    - A status code documented in the spec (typically 2xx/3xx), and
    - A response `Content-Type` that matches the spec.
- Use the filtered CSV as input to `tools/run-openapi-batch.sh` so DriveBy runs mostly against APIs whose endpoints actually work.

### Inputs

- Input CSV format (same as the batch tool):

```text
name,url,host
petstore,https://petstore3.swagger.io/api/v3/openapi.json,petstore3.swagger.io
someapi,https://example.com/openapi.json,api.example.com
```

### How to Run Locally or in CI

```bash
# 1) Start with up to 20 candidate APIs
python tools/probe-openapi-endpoints.py tests/openapis-sample.csv /tmp/filtered-openapis.csv

# 2) Run DriveBy only against the filtered set
tools/run-openapi-batch.sh /tmp/filtered-openapis.csv /tmp/driveby-openapi-good
```

If no APIs pass the probe, `/tmp/filtered-openapis.csv` will only contain the header row.

## 3. Future Suites

Additional suites should follow the same pattern:

- Document the **intent** (e.g., load, functional, regression).
- Specify **inputs**, **commands**, and **expected outputs**.
- Include guidance for both **local/CI** and **Kubernetes/Argo** execution.

