#!/usr/bin/env bash
# check-docs.sh — Comprehensive documentation & implementation audit for DriveBy
# Run: bash tools/check-docs.sh
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CLI_ROOT="$REPO_ROOT/driveby-cli"
errors=0
warnings=0

echo "╔══════════════════════════════════════════════════════════════╗"
echo "║          DriveBy Documentation & Code Sync Audit            ║"
echo "╚══════════════════════════════════════════════════════════════╝"
echo ""

# ═══════════════════════════════════════════════════════════════════
# 1. CLAUDE.md File Existence
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 1. CLAUDE.md Files ━━━"
required_dirs=(
    "."
    "driveby-cli"
    "driveby-cli/internal/principles"
    "driveby-cli/internal/spec"
    "apis"
    "apis/perfect-api"
    "kubernetes"
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

# ═══════════════════════════════════════════════════════════════════
# 2. Principle Documentation (docs/principles/P001-P008.md)
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 2. Principle Documentation ━━━"
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

# ═══════════════════════════════════════════════════════════════════
# 3. Core Documentation Files
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 3. Core Documentation ━━━"
core_docs=(
    "docs/ddt-axioms.md"
    "docs/architecture.md"
    "docs/evaluation-methodology.md"
    "docs/CLI_USAGE.md"
    "docs/WORKFLOW.md"
    "docs/MINIMAL_MODE_GUIDE.md"
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

# ═══════════════════════════════════════════════════════════════════
# 4. Principle Checker Implementation Status
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 4. Principle Checker Implementation ━━━"

declare -A principle_names
principle_names[1]="Compliance"
principle_names[2]="Documentation"
principle_names[3]="Errors"
principle_names[4]="Schema"
principle_names[5]="Security"
principle_names[6]="Functional"
principle_names[7]="Performance"
principle_names[8]="Versioning"

declare -A principle_files
principle_files[1]="p001_compliance.go"
principle_files[2]="p002_documentation.go"
principle_files[3]="p003_errors.go"
principle_files[4]="p004_schema.go"
principle_files[5]="p005_security.go"
principle_files[6]="p006_functional.go"
principle_files[7]="p007_performance.go"
principle_files[8]="p008_versioning.go"

for i in $(seq 1 8); do
    id=$(printf "P%03d" "$i")
    name="${principle_names[$i]}"
    src_file="$CLI_ROOT/internal/principles/${principle_files[$i]}"
    test_file="$CLI_ROOT/test/${principle_files[$i]%%.go}_test.go"
    # Some tests use a different naming pattern
    test_file_alt="$CLI_ROOT/test/$(printf "p%03d" "$i")_test.go"

    has_src="no"
    has_test="no"
    in_registry="no"

    if [ -f "$src_file" ]; then has_src="yes"; fi
    if [ -f "$test_file" ] || [ -f "$test_file_alt" ]; then has_test="yes"; fi
    if [ -f "$CLI_ROOT/internal/principles/registry.go" ]; then
        if grep -q "$id\|${principle_names[$i]}" "$CLI_ROOT/internal/principles/registry.go" 2>/dev/null; then
            in_registry="yes"
        fi
    fi

    status="COMPLETE"
    if [ "$has_src" = "no" ]; then
        status="NOT IMPLEMENTED"
        warnings=$((warnings + 1))
    elif [ "$in_registry" = "no" ]; then
        status="NOT REGISTERED"
        warnings=$((warnings + 1))
    fi

    test_status=""
    if [ "$has_test" = "no" ] && [ "$has_src" = "yes" ]; then
        test_status=" (NO TESTS)"
        warnings=$((warnings + 1))
    elif [ "$has_test" = "no" ]; then
        test_status=" (no tests)"
    fi

    echo "  $id ($name): $status$test_status"
    echo "    src: $has_src | test: $has_test | registry: $in_registry"
done
echo ""

# ═══════════════════════════════════════════════════════════════════
# 5. Cross-check: Doc status claims vs actual implementation
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 5. Doc Status vs Reality ━━━"

# Check principles CLAUDE.md for stale status claims
principles_claude="$CLI_ROOT/internal/principles/CLAUDE.md"
if [ -f "$principles_claude" ]; then
    for i in $(seq 1 8); do
        id=$(printf "P%03d" "$i")
        id_lower=$(printf "p%03d" "$i")
        src_file="$CLI_ROOT/internal/principles/${principle_files[$i]}"

        # Check if doc says "Not yet implemented" but file exists
        if [ -f "$src_file" ]; then
            if grep -q "Not yet implemented" "$principles_claude" 2>/dev/null && grep -q "$id.*Not yet implemented" "$principles_claude" 2>/dev/null; then
                echo "  STALE: principles/CLAUDE.md says $id 'Not yet implemented' but source file exists"
                warnings=$((warnings + 1))
            fi
        fi

        # Check if doc says "Complete" but file doesn't exist
        if [ ! -f "$src_file" ]; then
            if grep -q "$id.*Complete" "$principles_claude" 2>/dev/null; then
                echo "  STALE: principles/CLAUDE.md says $id 'Complete' but source file is missing"
                errors=$((errors + 1))
            fi
        fi
    done
fi

# Check docs/ CLAUDE.md for stale "(planned)" references
docs_claude="$REPO_ROOT/docs/CLAUDE.md"
if [ -f "$docs_claude" ]; then
    planned_count=0
    planned_count=$(grep -c "(planned)" "$docs_claude" 2>/dev/null) || true
    if [ "$planned_count" -gt 0 ]; then
        echo "  STALE: docs/CLAUDE.md still has $planned_count '(planned)' references"
        grep "(planned)" "$docs_claude" 2>/dev/null | while IFS= read -r line; do
            filename=$(echo "$line" | grep -oP '`[^`]+`\s*\(planned\)' | head -1)
            if [ -n "$filename" ]; then
                echo "    -> $filename"
            fi
        done
        warnings=$((warnings + 1))
    else
        echo "  OK: docs/CLAUDE.md has no stale '(planned)' references"
    fi
fi

# Check thesis CLAUDE.md for stale "(planned)" references
thesis_claude="$REPO_ROOT/thesis/CLAUDE.md"
if [ -f "$thesis_claude" ]; then
    planned_count=0
    planned_count=$(grep -c "(planned)" "$thesis_claude" 2>/dev/null) || true
    if [ "$planned_count" -gt 0 ]; then
        echo "  STALE: thesis/CLAUDE.md still has $planned_count '(planned)' references"
        warnings=$((warnings + 1))
    else
        echo "  OK: thesis/CLAUDE.md has no stale '(planned)' references"
    fi
fi

# Check root CLAUDE.md principle status table
root_claude="$REPO_ROOT/CLAUDE.md"
if [ -f "$root_claude" ]; then
    for i in $(seq 1 8); do
        id=$(printf "P%03d" "$i")
        src_file="$CLI_ROOT/internal/principles/${principle_files[$i]}"

        if [ -f "$src_file" ] && grep -q "$id.*Not yet" "$root_claude" 2>/dev/null; then
            echo "  STALE: Root CLAUDE.md says $id 'Not yet implemented' but source exists"
            warnings=$((warnings + 1))
        fi
        if [ ! -f "$src_file" ] && grep -q "$id.*| Complete" "$root_claude" 2>/dev/null; then
            echo "  STALE: Root CLAUDE.md says $id 'Complete' but source is missing"
            errors=$((errors + 1))
        fi
    done
fi
echo ""

# ═══════════════════════════════════════════════════════════════════
# 6. Thesis Chapter Status
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 6. Thesis Chapters ━━━"
for i in $(seq 1 7); do
    chapter=$(printf "%02d" "$i")
    # shellcheck disable=SC2086
    for f in "$REPO_ROOT"/thesis/chapters/${chapter}-*.tex; do
        if [ -f "$f" ]; then
            lines=$(wc -l < "$f")
            words=$(wc -w < "$f")
            basename_f=$(basename "$f")
            if [ "$lines" -lt 20 ]; then
                echo "  STUB: thesis/chapters/$basename_f ($lines lines, ~$words words)"
                warnings=$((warnings + 1))
            else
                echo "  OK: thesis/chapters/$basename_f ($lines lines, ~$words words)"
            fi
        fi
    done
done
echo ""

# ═══════════════════════════════════════════════════════════════════
# 7. Go Source Files Without Documentation
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 7. Go Packages vs CLAUDE.md Coverage ━━━"
# Find all Go package directories (that have .go files)
while IFS= read -r pkg_dir; do
    rel_dir="${pkg_dir#"$CLI_ROOT"/}"
    # Skip test/, cmd/, archive/ directories
    case "$rel_dir" in
        test*|cmd*|archive*) continue ;;
    esac
    # Check if parent or self has CLAUDE.md
    has_claude="no"
    if [ -f "$pkg_dir/CLAUDE.md" ]; then
        has_claude="yes"
    fi
    # Check if parent directory has CLAUDE.md
    parent_dir=$(dirname "$pkg_dir")
    if [ -f "$parent_dir/CLAUDE.md" ]; then
        has_claude="yes (parent)"
    fi
    if [ "$has_claude" = "no" ]; then
        echo "  UNDOCUMENTED: driveby-cli/$rel_dir/ — no CLAUDE.md in package or parent"
        warnings=$((warnings + 1))
    fi
done < <(find "$CLI_ROOT/internal" -name "*.go" -exec dirname {} \; | sort -u)
echo ""

# ═══════════════════════════════════════════════════════════════════
# 8. Test Coverage Gaps
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 8. Test Coverage Gaps ━━━"
# List source files that probably should have tests
for src in "$CLI_ROOT"/internal/principles/p[0-9]*.go; do
    if [ -f "$src" ]; then
        basename_src=$(basename "$src" .go)
        # Check for test file with same name or principle ID
        principle_id=$(echo "$basename_src" | grep -oP 'p\d+' | head -1)
        test_found="no"
        for test_pattern in "$CLI_ROOT/test/${basename_src}_test.go" "$CLI_ROOT/test/${principle_id}_test.go"; do
            if [ -f "$test_pattern" ]; then
                test_found="yes"
                break
            fi
        done
        if [ "$test_found" = "no" ]; then
            echo "  NO TEST: principles/$basename_src.go — missing test/${basename_src}_test.go"
            warnings=$((warnings + 1))
        fi
    fi
done
echo ""

# ═══════════════════════════════════════════════════════════════════
# 9. Build Check
# ═══════════════════════════════════════════════════════════════════
echo "━━━ 9. Build Check ━━━"
if command -v go &>/dev/null; then
    if (cd "$CLI_ROOT" && go build ./cmd/driveby 2>&1); then
        echo "  OK: go build ./cmd/driveby succeeded"
    else
        echo "  FAIL: go build ./cmd/driveby failed"
        errors=$((errors + 1))
    fi
    if (cd "$CLI_ROOT" && go vet ./... 2>&1); then
        echo "  OK: go vet ./... passed"
    else
        echo "  FAIL: go vet ./... failed"
        errors=$((errors + 1))
    fi
else
    echo "  SKIP: Go not found on PATH"
fi
echo ""

# ═══════════════════════════════════════════════════════════════════
# SUMMARY
# ═══════════════════════════════════════════════════════════════════
echo "╔══════════════════════════════════════════════════════════════╗"
echo "║  SUMMARY                                                    ║"
echo "╠══════════════════════════════════════════════════════════════╣"
printf "║  Errors:   %-47s ║\n" "$errors"
printf "║  Warnings: %-47s ║\n" "$warnings"
echo "╚══════════════════════════════════════════════════════════════╝"

if [ "$errors" -gt 0 ]; then
    echo ""
    echo "ERRORS require attention — documentation or code is broken."
    exit 1
elif [ "$warnings" -gt 0 ]; then
    echo ""
    echo "WARNINGS found — documentation is out of sync with code."
    exit 0
else
    echo ""
    echo "All checks passed — documentation and code are in sync."
    exit 0
fi
