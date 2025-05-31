# Report Formats

## JSON Reports
All validation, functional, and performance results are output as JSON to stdout and saved in the report output directory (default: ./reports). Example:

```json
{
  "timestamp": "2025-05-31T16:43:30.975+03:00",
  "version": "1.0.0",
  "environment": "development",
  "principles": [ ... ],
  "summary": { ... }
}
```

## Markdown Reports
A human-readable Markdown report is also generated for each run. Example:

```markdown
# API Validation Report

Generated: 2025-05-31T16:43:30.975+03:00

## Summary
| Total Checks | Passed | Failed | Critical | Warnings | Info |
|--------------|--------|--------|----------|----------|------|
| 7            | 5      | 2      | 1        | 1        | 0    |

## Principle Results
| Principle | Status | Message |
|-----------|--------|---------|
| P001      | Passed | ...     |
| P003      | Failed | ...     |
```

## Configuring Output Directory

You can set the report output directory with the `--report-dir` flag:

```
driveby validate-only --report-dir ./my-reports
```

All reports will be saved in the specified directory. 