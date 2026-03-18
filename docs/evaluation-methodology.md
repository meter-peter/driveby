# DriveBy Evaluation Methodology

## Overview

This document defines the evaluation protocol for the DriveBy DDT framework, corresponding to Chapter 6 of the thesis. The evaluation demonstrates that Documentation-Driven Testing can automatically assess API quality with measurable accuracy, and compares DDT effort against traditional QA approaches.

## Evaluation Dimensions

The evaluation addresses three research questions:

1. **Effectiveness:** Can DDT principles detect real API quality issues?
2. **Efficiency:** Does DDT reduce quality assurance effort compared to manual approaches?
3. **Generalizability:** Does DDT work across diverse, real-world APIs (not just the reference API)?

## Evaluation 1: Controlled Evaluation (Perfect API)

### Objective
Validate all 8 DDT principles against a single, fully-controlled reference API designed to pass every check. Then introduce known defects to verify detection accuracy.

### Setup
- **Target:** `perfect-api` (Python/FastAPI application in `apis/perfect-api/`)
- **Deployment:** `docker-compose up` (local) or Kubernetes CronJob
- **Spec:** `apis/perfect-api/openapi.json` (hand-crafted to be complete)
- **Modes tested:** minimal, strict, test-only

### Protocol

#### Phase 1: Baseline (All Principles Pass)
1. Start `perfect-api` via docker-compose
2. Run `driveby validate --spec apis/perfect-api/openapi.json --mode strict --base-url http://localhost:8000`
3. Verify all 8 principles pass
4. Record execution time

#### Phase 2: Defect Injection (Detection Accuracy)
For each principle P001-P008, introduce a known defect into the specification or implementation:

| Principle | Injected Defect | Expected Detection |
|-----------|-----------------|-------------------|
| P001 | Remove `info.version` field | FAIL: missing required field |
| P002 | Remove all operation summaries | FAIL: missing summaries |
| P003 | Remove all 4xx error responses | FAIL: missing error documentation |
| P004 | Remove type from path parameter schema | FAIL: missing type definition |
| P005 | Remove security schemes entirely | FAIL: no security defined |
| P006 | Change response schema (add field not in spec) | FAIL: schema mismatch |
| P007 | Add artificial 2-second delay to endpoints | FAIL: latency exceeds threshold |
| P008 | Change version to "latest" (non-semver) | FAIL: invalid version format |

Each defect is injected individually, validated, then reverted.

#### Phase 3: Mode Comparison
Run the same specification in all three modes and compare:
- Number of checks executed
- Number of findings
- Execution time
- Pass/fail outcome

### Metrics Collected

| Metric | Description | Unit |
|--------|-------------|------|
| Pass rate | Principles passing / total principles | Percentage |
| Detection rate | Injected defects detected / total injected | Percentage |
| False positive rate | Non-defects flagged as failures / total non-defect checks | Percentage |
| Execution time | Wall clock time for full validation | Seconds |
| Findings count | Total findings across all principles | Integer |

### Result Table: Controlled Evaluation

| Principle | Baseline | Defect Injected | Detected | False Positives | Time (ms) |
|-----------|----------|-----------------|----------|-----------------|-----------|
| P001 | | | | | |
| P002 | | | | | |
| P003 | | | | | |
| P004 | | | | | |
| P005 | | | | | |
| P006 | | | | | |
| P007 | | | | | |
| P008 | | | | | |
| **Total** | | | | | |

## Evaluation 2: Wild Evaluation (Public APIs)

### Objective
Assess DDT effectiveness against real-world, uncontrolled API specifications harvested from public registries. This tests generalizability — whether principles designed against a reference API apply to diverse specifications in the wild.

### Setup
- **Source:** APIs.guru registry (https://apis.guru/)
- **Harvest tool:** `tools/harvest-openapi-apisguru.py` — downloads OpenAPI specs from the registry
- **Probe tool:** Validates that specs are parseable before batch testing
- **Batch runner:** `tools/run-openapi-batch.sh` — runs DriveBy against each spec
- **Sample file:** `tests/openapis-sample.csv` — curated list of public API specs

### Protocol

#### Phase 1: Harvest
1. Run harvest script to download N OpenAPI specifications (target: 100+)
2. Filter to parseable specs (discard specs that fail basic JSON/YAML parse)
3. Record spec metadata: name, version, format (3.0 vs 3.1), endpoint count

#### Phase 2: Batch Validation
1. Run DriveBy in `minimal` mode against all specs
2. Run DriveBy in `strict` mode against all specs
3. Record per-principle pass/fail for each spec
4. Note: P006 and P007 are skipped (no live endpoints available for public APIs)

#### Phase 3: Analysis
1. Calculate per-principle pass rate across all specs
2. Identify most common failure patterns
3. Categorize failures: genuine quality issue vs. DDT false positive
4. Manual review of a random sample (20 specs) to validate DDT findings

### Metrics Collected

| Metric | Description | Unit |
|--------|-------------|------|
| Spec count | Total parseable specs evaluated | Integer |
| Per-principle pass rate | Specs passing each principle / total specs | Percentage |
| Overall pass rate | Specs passing all principles / total specs | Percentage |
| Common failure patterns | Top 5 most frequent finding types | Ranked list |
| False positive rate | Manually verified false positives / total findings (sample) | Percentage |
| Execution time per spec | Average validation time | Milliseconds |
| Total batch time | Wall clock time for full batch | Minutes |

### Result Table: Wild Evaluation (Per-Principle)

| Principle | Specs Tested | Pass | Fail | Pass Rate | Top Failure Reason |
|-----------|-------------|------|------|-----------|-------------------|
| P001 | | | | | |
| P002 | | | | | |
| P003 | | | | | |
| P004 | | | | | |
| P005 | | | | | |
| P008 | | | | | |
| **Overall** | | | | | |

### Result Table: Wild Evaluation (Mode Comparison)

| Mode | Avg Checks/Spec | Avg Findings/Spec | Avg Time/Spec (ms) | Overall Pass Rate |
|------|----------------|-------------------|--------------------|--------------------|
| Minimal | | | | |
| Strict | | | | |

## Evaluation 3: Effort Comparison (DDT vs Traditional QA)

### Objective
Compare the effort required to achieve equivalent API quality assurance coverage using DDT versus traditional manual/semi-automated approaches.

### Methodology
For each of the 8 DDT principles, estimate the equivalent traditional QA activity and its effort:

| DDT Principle | Traditional Equivalent | Traditional Effort | DDT Effort | Savings |
|---------------|----------------------|-------------------|------------|---------|
| P001 Compliance | Manual spec review | Hours/spec | Seconds/spec | |
| P002 Documentation | Documentation audit | Hours/spec | Seconds/spec | |
| P003 Error Handling | Negative test writing | Days/API | Seconds/spec | |
| P004 Schema | Schema review + test writing | Hours/spec | Seconds/spec | |
| P005 Security | Security audit | Days/API | Seconds/spec | |
| P006 Contract | Manual integration testing | Days/API | Minutes/spec | |
| P007 Performance | Performance test suite | Days/API | Minutes/spec | |
| P008 Versioning | Version policy review | Hours/spec | Seconds/spec | |
| **Total** | | | | |

### Assumptions and Limitations
- Traditional effort estimates are based on industry literature and author experience
- DDT effort includes only execution time, not initial framework development
- The comparison is per-API: DDT's advantage grows with the number of APIs validated
- P006 and P007 DDT effort is higher than static principles due to live API requirement

## Evaluation 4: CI/CD Integration Assessment

### Objective
Demonstrate that DDT integrates into GitOps workflows with minimal overhead.

### Protocol
1. Measure DriveBy execution time as a CI pipeline step
2. Measure Docker image size and startup time
3. Assess Kubernetes CronJob resource utilization
4. Compare against baseline pipeline time (without DDT step)

### Metrics

| Metric | Value |
|--------|-------|
| CLI binary size | |
| Docker image size | |
| Startup time (cold) | |
| Validation time (minimal, perfect-api) | |
| Validation time (strict, perfect-api) | |
| CI pipeline overhead | |
| K8s CronJob CPU request | |
| K8s CronJob memory request | |

## Statistical Analysis

For the wild evaluation, the following statistical measures will be reported:
- **Mean and standard deviation** of principle pass rates
- **Confidence intervals** (95%) for pass rates
- **Correlation analysis** between spec size (endpoint count) and validation findings
- **Distribution analysis** of execution times (histogram, percentiles)

## Threats to Validity

### Internal Validity
- **Perfect API bias:** The reference API was designed to pass all checks, which may not represent real-world complexity.
- **Defect injection artificiality:** Injected defects are known; real defects may be more subtle.

### External Validity
- **APIs.guru sample bias:** Public APIs in the registry may not represent enterprise/private APIs.
- **OpenAPI version skew:** Most public APIs use 3.0.x; results may not generalize to 3.1.0.
- **No live endpoint testing for wild APIs:** P006 and P007 cannot be evaluated against public APIs without authentication.

### Construct Validity
- **Effort estimation subjectivity:** Traditional QA effort estimates rely on literature and experience, not direct measurement.
- **False positive assessment:** Manual review of a sample (not all) findings introduces sampling bias.

## Source References

- Batch testing tools: `tools/run-openapi-batch.sh`, `tools/harvest-openapi-apisguru.py`
- Sample API list: `tests/openapis-sample.csv`
- Perfect API: `apis/perfect-api/`
- Thesis Chapter 6: Evaluation
- Docker Compose: `docker-compose.yml`
- Kubernetes deployment: `kubernetes/`
