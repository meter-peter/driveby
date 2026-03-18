# Tools Directory — Batch Testing Utilities

## Purpose
Python and bash utilities for harvesting, probing, and batch-validating public OpenAPI specifications. These tools generate the evaluation dataset for the thesis (Chapter 6).

## Files

| File | Language | Purpose |
|------|----------|---------|
| `harvest-openapi-apisguru.py` | Python | Harvests OpenAPI specification URLs from the APIs.guru directory. Outputs a CSV of API names + spec URLs. |
| `probe-openapi-endpoints.py` | Python | Probes harvested OpenAPI endpoints for reachability (HTTP 200). Filters out dead/unreachable APIs. Outputs a filtered CSV. |
| `run-openapi-batch.sh` | Bash | Reads a CSV of API URLs and runs the DriveBy CLI against each one. Collects per-API validation reports. |

## Workflow
```
harvest-openapi-apisguru.py  -->  raw CSV (all APIs.guru entries)
        |
probe-openapi-endpoints.py  -->  filtered CSV (reachable APIs only)
        |
run-openapi-batch.sh         -->  per-API validation reports
```

## Usage
```bash
# Step 1: Harvest API URLs from APIs.guru
python3 harvest-openapi-apisguru.py > apis-raw.csv

# Step 2: Filter to reachable endpoints
python3 probe-openapi-endpoints.py apis-raw.csv > apis-reachable.csv

# Step 3: Run DriveBy against all reachable APIs
./run-openapi-batch.sh apis-reachable.csv ./reports/
```

## Thesis Mapping
- **Chapter 6 (Evaluation)**: These tools produce the evaluation dataset
- The harvested APIs.guru dataset provides the "real-world" complement to the controlled `apis/perfect-api/` evaluation
- Batch results feed into the statistical analysis of DDT principle pass/fail rates across public APIs

## Notes
- The CSV format used is: `name,url,host` (header row included, 3 columns)
- `run-openapi-batch.sh` expects the DriveBy CLI binary at `../driveby-cli/driveby` (override via `DRIVEBY_BIN` env var)
- Batch tool is limited to a maximum of 20 entries per run
- Probing can be slow for large datasets; consider running with parallelism
- Batch runs default to `minimal` mode; edit the script to change to `strict`
