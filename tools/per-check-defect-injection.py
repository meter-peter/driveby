#!/usr/bin/env python3
"""Per-check defect injection runner.

For each check inside each implemented principle, generate a mutation of the
baseline spec that violates that single check (when feasible) and run DriveBy
against the mutated spec. Produces a confusion matrix used by Ch.7 §7.2.

Per Kritikos comments [551], [463] — 1-2 checks per principle is insufficient;
this script extends coverage and adds combinations of 2-3 simultaneous defects.
"""
from __future__ import annotations
import copy
import json
import random
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
BASELINE = ROOT / "apis" / "perfect-api" / "openapi.json"
DRIVEBY = ROOT / "driveby-cli" / "driveby"
OUT_DIR = ROOT / "results" / "per-check-injection"
OUT_DIR.mkdir(parents=True, exist_ok=True)

random.seed(42)  # reproducibility


def mutator_p001_duplicate_opid(spec):
    """P001: duplicate operationId across two operations."""
    ops = []
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict) and "operationId" in op:
                ops.append((path, method, op))
    if len(ops) < 2:
        return None
    ops[1][2]["operationId"] = ops[0][2]["operationId"]
    return spec


def mutator_p001_invalid_method(spec):
    """P001: invalid HTTP method."""
    for path, methods in spec.get("paths", {}).items():
        if "get" in methods:
            methods["FOO"] = methods.pop("get")
            return spec
    return None


def mutator_p002_strip_descriptions(spec):
    """P002: strip all operation/schema descriptions."""
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict):
                op.pop("description", None)
                op.pop("summary", None)
    for s in spec.get("components", {}).get("schemas", {}).values():
        if isinstance(s, dict):
            s.pop("description", None)
    return spec


def mutator_p002_strip_examples(spec):
    """P002: strip all request/response examples."""
    def strip(d):
        if isinstance(d, dict):
            d.pop("example", None)
            d.pop("examples", None)
            for v in d.values():
                strip(v)
        elif isinstance(d, list):
            for v in d:
                strip(v)
    strip(spec)
    return spec


def mutator_p002_strip_contact(spec):
    """P002: remove contact and license."""
    info = spec.get("info", {})
    info.pop("contact", None)
    info.pop("license", None)
    return spec


def mutator_p003_strip_4xx(spec):
    """P003: remove all 4xx/5xx error responses."""
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict) and "responses" in op:
                op["responses"] = {
                    k: v for k, v in op["responses"].items()
                    if not (k.startswith("4") or k.startswith("5"))
                }
    return spec


def mutator_p003_strip_components_responses(spec):
    """P003: remove the shared components.responses block."""
    spec.get("components", {}).pop("responses", None)
    return spec


def mutator_p004_untype_request_bodies(spec):
    """P004: remove `type` from all request body schemas."""
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict):
                rb = op.get("requestBody", {}).get("content", {})
                for media in rb.values():
                    sch = media.get("schema", {})
                    sch.pop("type", None)
    return spec


def mutator_p004_strip_constraints(spec):
    """P004: strip min/max/length constraints from schemas."""
    def strip(d):
        if isinstance(d, dict):
            for k in ["minLength", "maxLength", "minimum", "maximum",
                     "pattern", "format"]:
                d.pop(k, None)
            for v in d.values():
                strip(v)
        elif isinstance(d, list):
            for v in d:
                strip(v)
    strip(spec.get("components", {}).get("schemas", {}))
    return spec


def mutator_p005_strip_security_schemes(spec):
    """P005: remove components.securitySchemes entirely."""
    spec.get("components", {}).pop("securitySchemes", None)
    return spec


def mutator_p005_strip_global_security(spec):
    """P005: remove top-level security requirement."""
    spec.pop("security", None)
    return spec


def mutator_p005_strip_op_security(spec):
    """P005: remove per-operation security."""
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict):
                op.pop("security", None)
    return spec


def mutator_p008_invalid_version(spec):
    """P008: change version to non-semver string."""
    spec.setdefault("info", {})["version"] = "latest"
    return spec


def mutator_p008_strip_strategy_doc(spec):
    """P008: strip versioning-strategy keywords from info.description."""
    info = spec.get("info", {})
    if "description" in info:
        # Replace any version/migration/breaking-change wording
        for word in ["version", "Version", "VERSION", "migration",
                     "Migration", "breaking", "Breaking"]:
            info["description"] = info["description"].replace(word, "x")
    return spec


def mutator_p009_strip_param_examples(spec):
    """P009: strip parameter-level examples."""
    for path, methods in spec.get("paths", {}).items():
        for method, op in methods.items():
            if isinstance(op, dict):
                for p in op.get("parameters", []):
                    if isinstance(p, dict):
                        p.pop("example", None)
                        p.pop("examples", None)
    return spec


MUTATORS = {
    "p001-duplicate-opid": mutator_p001_duplicate_opid,
    "p001-invalid-method": mutator_p001_invalid_method,
    "p002-strip-descriptions": mutator_p002_strip_descriptions,
    "p002-strip-examples": mutator_p002_strip_examples,
    "p002-strip-contact": mutator_p002_strip_contact,
    "p003-strip-4xx-5xx": mutator_p003_strip_4xx,
    "p003-strip-shared-responses": mutator_p003_strip_components_responses,
    "p004-untype-request-bodies": mutator_p004_untype_request_bodies,
    "p004-strip-constraints": mutator_p004_strip_constraints,
    "p005-strip-security-schemes": mutator_p005_strip_security_schemes,
    "p005-strip-global-security": mutator_p005_strip_global_security,
    "p005-strip-op-security": mutator_p005_strip_op_security,
    "p008-invalid-version": mutator_p008_invalid_version,
    "p008-strip-strategy-doc": mutator_p008_strip_strategy_doc,
    "p009-strip-param-examples": mutator_p009_strip_param_examples,
}


def expected_principles(name):
    return [name.split("-")[0].upper()]


def run_driveby(spec_path, report_dir):
    report_dir.mkdir(parents=True, exist_ok=True)
    cmd = [
        str(DRIVEBY), "validate-only",
        "--openapi", str(spec_path),
        "--host", "localhost", "--protocol", "https", "--port", "443",
        "--validation-mode", "strict",
        "--report-dir", str(report_dir),
    ]
    proc = subprocess.run(cmd, capture_output=True, text=True, timeout=60)
    latest = report_dir / "validation-report-latest.json"
    if not latest.exists():
        return None, proc.stderr
    return json.loads(latest.read_text()), None


def parse_outcomes(report):
    """Return {pid: passed} and {pid: list of failed-check names}."""
    out, failed = {}, {}
    for p in report.get("principles", []):
        pid = p["Principle"]["id"]
        out[pid] = bool(p["Passed"])
        checks = p.get("Details", {}).get("checks", {}) or {}
        failed[pid] = [c for c, v in checks.items() if not v]
    return out, failed


def main():
    if not DRIVEBY.exists():
        print(f"DriveBy binary not found: {DRIVEBY}", file=sys.stderr)
        return 1
    if not BASELINE.exists():
        print(f"Baseline spec not found: {BASELINE}", file=sys.stderr)
        return 1

    baseline_spec = json.loads(BASELINE.read_text())

    # Baseline run
    baseline_dir = OUT_DIR / "_baseline"
    baseline_path = baseline_dir / "openapi.json"
    baseline_dir.mkdir(parents=True, exist_ok=True)
    baseline_path.write_text(json.dumps(baseline_spec, indent=2))
    baseline_report, err = run_driveby(baseline_path, baseline_dir)
    if baseline_report is None:
        print(f"Baseline DriveBy failed: {err}", file=sys.stderr)
        return 1
    baseline_outcomes, baseline_failed = parse_outcomes(baseline_report)
    print(f"Baseline: {baseline_outcomes}")

    # Per-mutation runs
    rows = []
    for name, mut in MUTATORS.items():
        spec = mut(copy.deepcopy(baseline_spec))
        if spec is None:
            rows.append((name, "n/a", "skip", []))
            continue
        run_dir = OUT_DIR / name
        run_dir.mkdir(parents=True, exist_ok=True)
        spec_path = run_dir / "openapi.json"
        spec_path.write_text(json.dumps(spec, indent=2))
        report, err = run_driveby(spec_path, run_dir)
        if report is None:
            rows.append((name, "n/a", f"err: {err[:80]}", []))
            continue
        outcomes, failed = parse_outcomes(report)
        target = name.split("-")[0].upper()
        # Detection success = target principle now fails AND was passing in baseline,
        #                     OR target principle has a new failed check it didn't have
        baseline_target_passed = baseline_outcomes.get(target, True)
        revised_target_passed = outcomes.get(target, True)
        baseline_failed_set = set(baseline_failed.get(target, []))
        revised_failed_set = set(failed.get(target, []))
        new_failures = revised_failed_set - baseline_failed_set
        if not baseline_target_passed:
            # baseline already fails — detection requires either still failing OR new failures appearing
            detected = (not revised_target_passed) or bool(new_failures)
        else:
            detected = not revised_target_passed
        # collateral: any *other* principle that newly fails
        collateral = []
        for pid, was_pass in baseline_outcomes.items():
            if pid == target:
                continue
            if was_pass and not outcomes.get(pid, True):
                collateral.append(pid)
        rows.append((name, target, "DETECTED" if detected else "MISSED",
                     collateral, list(new_failures)))

    # Combinations: 3 random pairs + 1 triple (per plan risk-register cap)
    mutator_names = [n for n in MUTATORS if MUTATORS[n] is not None]
    pairs = random.sample(
        [(a, b) for i, a in enumerate(mutator_names)
         for b in mutator_names[i + 1:]
         if a.split("-")[0] != b.split("-")[0]],
        3,
    )
    triples = random.sample(
        [(a, b, c) for i, a in enumerate(mutator_names)
         for j, b in enumerate(mutator_names[i + 1:], i + 1)
         for c in mutator_names[j + 1:]
         if len({a.split("-")[0], b.split("-")[0], c.split("-")[0]}) == 3],
        1,
    )
    combos = [("pair", p) for p in pairs] + [("triple", t) for t in triples]
    combo_rows = []
    for kind, combo in combos:
        spec = copy.deepcopy(baseline_spec)
        for n in combo:
            res = MUTATORS[n](spec)
            if res is None:
                spec = None
                break
            spec = res
        if spec is None:
            continue
        name = "+".join(combo)
        run_dir = OUT_DIR / f"_combo_{kind}_{abs(hash(name)) % 100000}"
        run_dir.mkdir(parents=True, exist_ok=True)
        spec_path = run_dir / "openapi.json"
        spec_path.write_text(json.dumps(spec, indent=2))
        report, err = run_driveby(spec_path, run_dir)
        if report is None:
            continue
        outcomes, failed = parse_outcomes(report)
        targets = sorted({n.split("-")[0].upper() for n in combo})
        all_detected = all(
            (not outcomes.get(t, True))
            or (set(failed.get(t, [])) - set(baseline_failed.get(t, [])))
            for t in targets
        )
        combo_rows.append((kind, combo, targets,
                           "ALL DETECTED" if all_detected else "PARTIAL",
                           outcomes))

    # Summary CSV-ish output
    summary = {
        "baseline": baseline_outcomes,
        "single_check_runs": [
            {"mutation": r[0], "target": r[1], "outcome": r[2],
             "collateral": r[3] if len(r) > 3 else [],
             "new_failed_checks": r[4] if len(r) > 4 else []}
            for r in rows
        ],
        "combination_runs": [
            {"kind": k, "mutations": list(c), "targets": t,
             "outcome": o, "principles": op}
            for (k, c, t, o, op) in combo_rows
        ],
    }
    summary_path = OUT_DIR / "summary.json"
    summary_path.write_text(json.dumps(summary, indent=2))
    print(f"\nSummary written to: {summary_path}")
    print(f"\nSingle-check detection: "
          f"{sum(1 for r in rows if r[2] == 'DETECTED')}/{len(rows)}")
    print(f"Combination detection (all targets caught): "
          f"{sum(1 for r in combo_rows if 'ALL' in r[3])}/{len(combo_rows)}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
