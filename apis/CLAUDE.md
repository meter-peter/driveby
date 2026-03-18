# APIs Directory — Sample APIs for Testing

## Purpose
Contains sample API implementations used for controlled evaluation of DriveBy's validation principles. These APIs serve as the ground truth for thesis evaluation (Chapter 6).

## Inventory

| Directory | Description | Spec Version | Status |
|-----------|-------------|-------------|--------|
| `perfect-api/` | FastAPI reference implementation — designed to pass all static DDT principles | OpenAPI 3.1.0 | Complete |

## Evaluation Purpose
- **Thesis Chapter 6**: These sample APIs provide the controlled evaluation environment
- The perfect-api is the "known-good" baseline: DriveBy should report high scores against it
- Real-world APIs from APIs.guru (via `tools/harvest-openapi-apisguru.py`) provide the "unknown" evaluation set

## Public API Validation

DriveBy can also validate remote public APIs by URL -- no local API required. This capability is used for:
- **Thesis evaluation** (Chapter 6, Wild Evaluation): batch-validate APIs from APIs.guru
- **Benchmarking**: compare DDT scores of real-world APIs against the perfect-api baseline
- **Ad-hoc testing**: validate any public OpenAPI spec before integrating a third-party API

Example: `driveby validate-only --openapi https://petstore3.swagger.io/api/v3/openapi.json --host petstore3.swagger.io --protocol https --port 443 --validation-mode strict`

Reference result: Swagger Petstore v3 scores 1/6 in strict mode (only P001 passes).

## Gaps
- Currently only 1 local sample API (perfect-api)
- Consider adding: a deliberately flawed API (failing P001-P005), a minimal API (edge cases), a Swagger 2.0 API (testing swagger2.go adapter)

## Rules
- Each API directory must contain its own `openapi.json` (or equivalent spec file)
- Each API must be independently runnable via Docker
- Document which principles each API is expected to pass/fail in its local CLAUDE.md
