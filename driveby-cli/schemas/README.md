# DriveBy report schemas

This directory holds the **JSON Schema (Draft 7)** describing the validation
report format emitted by the DriveBy CLI.

## Files

- `report.schema.json` — schema for `ValidationReport`, the top-level JSON
  document produced by `driveby validate` and related commands.

## Validating a report

Any standard JSON Schema validator works. For example:

```bash
# python
pip install jsonschema
python -c "import json,jsonschema; \
  jsonschema.validate( \
    json.load(open('results/local/perfect-api-validate-strict.json')), \
    json.load(open('driveby-cli/schemas/report.schema.json')))"

# go
go run github.com/santhosh-tekuri/jsonschema/v5/cmd/jv \
  driveby-cli/schemas/report.schema.json \
  results/local/perfect-api-validate-strict.json
```

The schema is verified against every report in `results/local/*-validate-*.json`
during the release process.

## Known wire-format issue: mixed casing

The current wire format mixes naming conventions:

- **Top-level** keys are `snake_case` (`status`, `exit_code`, `total_checks`, …).
- **Nested `PrincipleResult` and `ValidationSummary`** fields are `PascalCase`
  (`Principle`, `Passed`, `Message`, `CriticalIssues`, `Categories`, …).

This is because `internal/types/report.go` is missing `json:"…"` struct tags on
those two types, so Go's encoder falls back to the field's Go identifier.

The schema documents the format **as currently shipped**, so consumers can rely
on it today. Aligning the casing is tracked as a follow-up; when those tags are
added the schema will be updated in lockstep.

## Versioning

The schema is versioned alongside the CLI binary. The `$id` in
`report.schema.json` points to the canonical path in the repository. Breaking
changes to the report format are accompanied by a bump in the major version of
the DriveBy CLI and an updated schema.
