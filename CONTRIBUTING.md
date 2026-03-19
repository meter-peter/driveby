# Contributing to DriveBy

## Development Setup

```bash
# Clone the repository
git clone https://github.com/meter-peter/driveby.git
cd driveby

# Build the CLI
cd driveby-cli
go build ./cmd/driveby

# Run tests
go test -count=1 ./test/...

# Lint
go vet ./...
```

## Code Standards

- **Go version**: 1.21 (match `go.mod`)
- **`spec.APISpec` abstraction is mandatory** — no raw `*openapi3.T` in principle checkers
- **No `os.Exit()` in command handlers** — return errors from `RunE`
- **Types live in `internal/types/`** — standalone, no circular dependencies
- **CLI config shared via `config.go`** — no flag duplication across commands
- **Dependency flow**: `types` <- `spec` <- `loader` <- `principles` <- `testing` <- `engine` <- `cli`

## Adding a New Principle

1. Create `driveby-cli/internal/principles/pXXX_name.go` implementing `PrincipleChecker`
2. Register the checker in the principle registry
3. Create `driveby-cli/test/pXXX_test.go` with unit tests
4. Create `docs/principles/PXXX.md` with documentation
5. Update `driveby-cli/internal/principles/CLAUDE.md`
6. Update the root `CLAUDE.md` principles table

## Pull Request Process

1. Fork the repository and create a feature branch from `main`
2. Make your changes following the code standards above
3. Ensure all tests pass: `cd driveby-cli && go test -count=1 ./test/...`
4. Ensure linting passes: `cd driveby-cli && go vet ./...`
5. Update documentation if your change affects documented behavior
6. Submit a pull request against `main`

## Project Structure

See the root [CLAUDE.md](CLAUDE.md) for the full monorepo layout and architecture details.
