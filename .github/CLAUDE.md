# GitHub Directory — CI/CD Pipelines

## Workflows

### `workflows/ci.yml` — Continuous Integration
Triggered on push to `main` and on pull requests to `main`.

**Jobs:**
1. **build-and-test** (ubuntu-latest, Go 1.21):
   - `go build` with ldflags for version/commit injection
   - `go vet ./...` — static analysis
   - `go test -count=1 -coverprofile=coverage.out ./test/...` — tests with coverage
   - Upload coverage artifact on PRs

2. **docker** (ubuntu-latest, depends on build-and-test):
   - Only runs on main branch pushes
   - Builds Docker image from `driveby-cli/`

### `workflows/docker-publish.yml` — Docker Image Publishing
Publishes `ghcr.io/meter-peter/driveby:latest` to GHCR on push to main (when `driveby-cli/` changes) or manual dispatch. Injects version/commit/date build args into the multi-stage Dockerfile. Only triggers on branch pushes — tagged releases are handled by `release.yml`.

### `workflows/release.yml` — Release Pipeline
Triggered on `v*` tag push (e.g., `git tag v0.3.0 && git push origin v0.3.0`). Runs four parallel jobs after tests pass:

1. **test** — `go vet` + `go test` (gate for all other jobs)
2. **goreleaser** — Builds multi-platform CLI binaries (linux/darwin/windows × amd64/arm64), creates GitHub Release with checksums
3. **docker** — Builds and pushes Docker image with semver tags (`0.3.0`, `0.3`, `latest`) to `ghcr.io/meter-peter/driveby`
4. **helm** — Packages and pushes Helm chart to `oci://ghcr.io/meter-peter/charts/driveby`

Depends on `.goreleaser.yaml` at the project root for binary build configuration.

## Thesis Mapping
- **Chapter 5 (Workflow)**: CI/CD pipeline is part of the DDT feedback loop
- The CI pipeline demonstrates the "Observability" axiom — every change produces measurable validation results
- The build-test-publish chain is the simplest DDT integration pattern

## Planned Additions
- **batch-evaluation** workflow: Scheduled run of `tools/run-openapi-batch.sh` against APIs.guru dataset for thesis data collection

## Key Details
- Go version: 1.21
- Tests run from `driveby-cli/test/` directory (not alongside source)
- Coverage is uploaded as artifact, not to an external service
- Docker build context is `driveby-cli/` (not monorepo root)
