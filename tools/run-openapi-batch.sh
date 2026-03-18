#!/usr/bin/env bash
set -euo pipefail

# run-openapi-batch.sh
# Small helper tool to run driveby against a curated batch
# of public OpenAPI specs (maximum of 20 entries per run).
#
# Usage:
#   tools/run-openapi-batch.sh path/to/openapis.csv [report-root]
#
# CSV format (header required):
#   name,url,host
#
# Example row:
#   petstore,https://petstore3.swagger.io/api/v3/openapi.json,petstore3.swagger.io

if [[ $# -lt 1 || $# -gt 2 ]]; then
  echo "Usage: $0 path/to/openapis.csv [report-root]" >&2
  exit 3
fi

CSV_FILE="$1"
REPORT_ROOT="${2:-/tmp/driveby-openapi-batch}"
REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
DRIVEBY_BIN="${DRIVEBY_BIN:-$REPO_ROOT/driveby-cli/driveby}"

if [[ ! -f "$CSV_FILE" ]]; then
  echo "Error: CSV file not found: $CSV_FILE" >&2
  exit 3
fi

mkdir -p "$REPORT_ROOT"
timestamp="$(date -u +%Y%m%dT%H%M%SZ)"
summary_file="$REPORT_ROOT/summary-$timestamp.csv"

# Count data rows (excluding header)
data_rows=$(tail -n +2 "$CSV_FILE" | sed '/^\s*$/d' | wc -l | tr -d ' ')
if (( data_rows == 0 )); then
  echo "Error: CSV contains no data rows (only header)." >&2
  exit 3
fi

if (( data_rows > 20 )); then
  echo "Error: CSV has $data_rows entries; this tool is limited to a maximum of 20." >&2
  exit 3
fi

echo "name,url,host,exit_code,mode,report_dir" > "$summary_file"

tail -n +2 "$CSV_FILE" | sed '/^\s*$/d' | while IFS=',' read -r name url host; do
  if [[ -z "${name:-}" || -z "${url:-}" || -z "${host:-}" ]]; then
    echo "Skipping row with missing fields: name='$name' url='$url' host='$host'" >&2
    continue
  fi

  echo "Running driveby for '$name' ($url) against host '$host'..."

  report_dir="$REPORT_ROOT/${name}-${timestamp}"
  mkdir -p "$report_dir"

  set +e
  "$DRIVEBY_BIN" validate-only \
    --openapi "$url" \
    --host "$host" \
    --protocol https \
    --timeout 30s \
    --validation-mode minimal \
    --report-dir "$report_dir"
  exit_code=$?
  set -e

  echo "$name,$url,$host,$exit_code,minimal,$report_dir" >> "$summary_file"
done

echo "Batch run complete. Summary written to: $summary_file"

