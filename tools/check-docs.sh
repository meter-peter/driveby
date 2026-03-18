#!/usr/bin/env bash
# check-docs.sh — Validate documentation completeness for DriveBy
# Run: bash tools/check-docs.sh
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
errors=0

echo "=== DriveBy Documentation Check ==="
echo ""

# 1. Check that key directories have CLAUDE.md files
echo "--- Checking CLAUDE.md files ---"
required_dirs=(
    "."
    "driveby-cli"
    "driveby-cli/internal/principles"
    "driveby-cli/internal/spec"
    "apis"
    "apis/perfect-api"
    "kubernetes"
    "samples"
    "tools"
    "docs"
    ".github"
    "thesis"
)

for dir in "${required_dirs[@]}"; do
    if [ ! -f "$REPO_ROOT/$dir/CLAUDE.md" ]; then
        echo "  MISSING: $dir/CLAUDE.md"
        errors=$((errors + 1))
    else
        echo "  OK: $dir/CLAUDE.md"
    fi
done
echo ""

# 2. Check that every principle has a docs/principles/ file
echo "--- Checking principle documentation ---"
for i in $(seq 1 8); do
    id=$(printf "P%03d" "$i")
    doc="$REPO_ROOT/docs/principles/${id}.md"
    if [ ! -f "$doc" ]; then
        echo "  MISSING: docs/principles/${id}.md"
        errors=$((errors + 1))
    else
        echo "  OK: docs/principles/${id}.md"
    fi
done
echo ""

# 3. Check core documentation files exist
echo "--- Checking core documentation ---"
core_docs=(
    "docs/ddt-axioms.md"
    "docs/architecture.md"
    "docs/evaluation-methodology.md"
)

for doc in "${core_docs[@]}"; do
    if [ ! -f "$REPO_ROOT/$doc" ]; then
        echo "  MISSING: $doc"
        errors=$((errors + 1))
    else
        echo "  OK: $doc"
    fi
done
echo ""

# 4. Check principle docs don't claim MISSING test coverage when tests exist
echo "--- Cross-checking principle docs vs tests ---"
for i in $(seq 1 8); do
    id=$(printf "P%03d" "$i")
    id_lower=$(printf "p%03d" "$i")
    doc="$REPO_ROOT/docs/principles/${id}.md"
    test_file="$REPO_ROOT/driveby-cli/test/${id_lower}_test.go"

    if [ -f "$doc" ] && [ -f "$test_file" ]; then
        if grep -qi "^\- \*\*Status:\*\* Stub" "$doc" 2>/dev/null; then
            echo "  WARNING: docs/principles/${id}.md claims Stub status — but test file exists at test/${id_lower}_test.go"
        fi
    fi
done
echo ""

# 5. Check thesis chapters that are still empty stubs
echo "--- Checking thesis chapter status ---"
for i in $(seq 1 7); do
    chapter=$(printf "%02d" "$i")
    tex="$REPO_ROOT/thesis/chapters/${chapter}-*.tex"
    # shellcheck disable=SC2086
    for f in $tex; do
        if [ -f "$f" ]; then
            lines=$(wc -l < "$f")
            basename_f=$(basename "$f")
            if [ "$lines" -lt 20 ]; then
                echo "  STUB: thesis/chapters/$basename_f ($lines lines)"
            else
                echo "  OK: thesis/chapters/$basename_f ($lines lines)"
            fi
        fi
    done
done
echo ""

# Summary
echo "=== Summary ==="
if [ "$errors" -eq 0 ]; then
    echo "All documentation checks passed."
    exit 0
else
    echo "$errors documentation issue(s) found."
    exit 1
fi
