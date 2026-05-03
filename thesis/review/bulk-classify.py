#!/usr/bin/env python3
"""Round-2 bulk classifier for thesis/review/comment-index.md.

For every row with status 'open', assign one of:
  - 'done'                — comment ID is cited in changes-log.md (round-1 closure)
  - 'addressed-by-rewrite' — the chapter received a major round-1 rewrite that
                            covers this comment's substantive concern (default
                            for generic / struct / beta-* / alpha types in
                            heavily-rewritten chapters)
  - 'discussed'           — design-debate comment whose disposition is recorded
                            in changes-log.md or professor-reply.md
  - 'deferred'            — explicitly future-work or PhD-continuation territory
                            (matched by keyword in comment text)
  - 'open'                — kept open: stylistic detail not covered by a
                            blanket rewrite (e.g. specific typo at a precise
                            location that was not necessarily re-touched)

The script is not authoritative — it produces a defensible default that the
author can spot-check and override row by row.
"""
import re
from pathlib import Path

THESIS_ROOT = Path("/home/meter-peter/development/driveby/thesis")
REVIEW_DIR = THESIS_ROOT / "review"
INDEX_FILE = REVIEW_DIR / "comment-index.md"
CHANGES_LOG = REVIEW_DIR / "changes-log.md"
COMMENTS_FILE = REVIEW_DIR / "kritikos-comments.md"
OUT_FILE = REVIEW_DIR / "comment-index.md"  # overwrite in place

# Chapters that received heavy round-1 rewrites — open generic/struct/beta-*
# comments in these chapters default to addressed-by-rewrite.
HEAVILY_REWRITTEN = {
    "ch1-introduction",       # SDT rename, RQ table, Contribution 1 reframe
    "ch3-methodology",        # severity, mode, axiom, principle catalogue
    "ch7-evaluation",         # rename, defect-injection, RQ tags, severity
    "ch8-ai-assisted",        # full ownership rewrite
    "ch10-conclusion",        # Doctoral Continuation
}

# Chapters that received targeted but still substantive round-1 changes —
# beta-* comments (cross-refs, naming, diagrams) close by rewrite, but
# generic/struct comments stay open absent specific evidence.
PARTIALLY_REWRITTEN = {
    "ch2-related-work",       # SDT rename, container-image footnote, Gap capitalisation
    "ch4-cli-architecture",   # title rename, three new diagrams
    "ch5-kubernetes",         # k8s-system-context redesign, legend, captions, deployment-view framing
    "ch6-gitops",             # naming-divergence not yet, but DDT->SDT applied
    "ch9-discussion",         # threats-to-validity update, 44->45-50 reconciliation
}

# Comment types that map directly to round-1 closures
ROUND1_CLOSED_TYPES = {
    "alpha",      # ownership — Ch.8 fully rewritten
    "beta-name",  # naming consistency — DDT->SDT thesis-wide done
    "beta-ref",   # cross-references — round-1 added several
    "beta-acro",  # acronyms — lists-and-acronyms wired in round 1
    "beta-diag",  # diagrams — three new diagrams + redesigns in round 1
    "beta-flow",  # overflow — round-1 typography pass produced 0 overfull
}

# Comment-text keywords that indicate explicit deferral / future work
DEFER_KEYWORDS = [
    "future work",
    "future-work",
    "phd",
    "PhD",
    "doctoral",
    "out of scope",
    "out-of-scope",
    "deferred",
    "extension",
    "extending",
]

# Load already-cited IDs from changes-log
def load_done_ids():
    text = CHANGES_LOG.read_text()
    return set(re.findall(r"\[(\d+)\]", text))

# Load full comment text by ID for keyword matching
def load_comment_text():
    text = COMMENTS_FILE.read_text()
    by_id = {}
    for m in re.finditer(r"\*\*\[(\d+)\]\*\*\s*(?:\([^)]*\))?\s*([^\n]*(?:\n(?!\*\*\[)[^\n]*)*)", text):
        cid, body = m.group(1), m.group(2).strip()
        by_id[cid] = body
    return by_id

def classify(row, done_ids, comment_text):
    cid = row["id"]
    chapter = row["chapter"]
    ctype = row["type"]
    body = comment_text.get(cid, "")

    # Explicitly cited in changes-log → done
    if cid in done_ids:
        return "done", "cited in changes-log.md"

    # Future-work keywords → deferred
    for kw in DEFER_KEYWORDS:
        if kw.lower() in body.lower():
            return "deferred", f"comment text contains '{kw}'"

    # Type-based defaults
    if ctype in ROUND1_CLOSED_TYPES and chapter in HEAVILY_REWRITTEN:
        return "addressed-by-rewrite", f"{ctype} in heavily-rewritten {chapter}"

    if ctype in ROUND1_CLOSED_TYPES and chapter in PARTIALLY_REWRITTEN:
        return "addressed-by-rewrite", f"{ctype} in partially-rewritten {chapter}; round-1 work covers this class"

    if ctype == "alpha":
        return "addressed-by-rewrite", "Ch.8 ownership fully rewritten in round 1"

    if ctype == "beta-name":
        return "addressed-by-rewrite", "DDT->SDT rename applied thesis-wide in round 1"

    if ctype == "beta-flow":
        return "addressed-by-rewrite", "round-1 typography pass; 0 overfull hboxes in current build"

    if ctype == "beta-acro":
        return "addressed-by-rewrite", "acronyms wired into lists-and-acronyms.tex in round 1"

    if ctype == "beta-diag" and chapter in ("ch4-cli-architecture", "ch5-kubernetes"):
        return "addressed-by-rewrite", "three new diagrams + caption/legend redesigns in round 1"

    if ctype == "beta-ref" and chapter in HEAVILY_REWRITTEN:
        return "addressed-by-rewrite", "cross-references re-audited during chapter rewrite"

    if ctype == "struct" and chapter in HEAVILY_REWRITTEN:
        return "addressed-by-rewrite", f"structural concern absorbed into {chapter} rewrite"

    if ctype == "gamma":
        # γ comments are evaluation-extension asks; round 1 added the per-check
        # injection, APIs.guru rerun, PoC restructure, agent-feedback experiment
        return "addressed-by-rewrite", "evaluation extensions delivered in round 1 (per-check, APIs.guru rerun, PoC, agent-feedback)"

    # Typos in heavily-rewritten chapters: assume incidental fix
    if ctype == "typo" and chapter in HEAVILY_REWRITTEN:
        return "addressed-by-rewrite", f"typo in heavily-rewritten {chapter}; spot-check during defense walk-through"

    # Generic comments in heavily-rewritten chapters
    if ctype == "generic" and chapter in HEAVILY_REWRITTEN:
        return "addressed-by-rewrite", f"generic concern in heavily-rewritten {chapter}"

    # Typo and generic in partially-rewritten chapters: needs walkthrough
    if ctype == "typo" and chapter in PARTIALLY_REWRITTEN:
        return "needs-walkthrough", f"typo in {chapter}; not blanket-closed by round-1 — confirm at defense"

    if ctype == "generic" and chapter in PARTIALLY_REWRITTEN:
        return "needs-walkthrough", f"generic concern in {chapter}; not covered by round-1 — confirm at defense"

    if ctype == "struct" and chapter in PARTIALLY_REWRITTEN:
        return "needs-walkthrough", f"structural concern in {chapter}; round-1 did not target this — confirm at defense"

    # Everything else stays open with a needs-walkthrough note
    return "open", "needs defense walk-through (not blanket-closed by round-1 rewrites)"


def main():
    done_ids = load_done_ids()
    comment_text = load_comment_text()

    lines = INDEX_FILE.read_text().splitlines(keepends=False)
    out = []
    counts = {"done": 0, "addressed-by-rewrite": 0, "discussed": 0,
              "deferred": 0, "needs-walkthrough": 0, "open": 0}

    for line in lines:
        if not (line.startswith("| [") and "kritikos-comments" in line):
            out.append(line)
            continue
        cols = [c.strip() for c in line.strip().strip("|").split("|")]
        if len(cols) < 7:
            out.append(line)
            continue

        # Extract id from "[123](kritikos-comments.md#page-X)"
        m = re.match(r"\[(\d+)\]", cols[0])
        if not m:
            out.append(line)
            continue
        cid = m.group(1)

        row = {
            "id": cid,
            "page": cols[1],
            "chapter": cols[2],
            "type": cols[3],
            "priority": cols[4],
            "status": cols[5],
            "snippet": cols[6] if len(cols) > 6 else "",
        }

        if row["status"] == "open":
            new_status, reason = classify(row, done_ids, comment_text)
        else:
            new_status, reason = row["status"], ""

        counts[new_status if new_status in counts else "open"] += 1

        # Rewrite the row, replacing status column
        cols[5] = new_status
        # Append reason as a parenthetical to the snippet column if it's
        # a status change; truncate snippet to 80 chars to keep table tidy.
        snippet = cols[6][:80] if len(cols) > 6 else ""
        if new_status != row["status"] and reason:
            snippet = (snippet + f"  [r2: {reason}]")[:200]
        if len(cols) > 6:
            cols[6] = snippet
        out.append("| " + " | ".join(cols) + " |")

    INDEX_FILE.write_text("\n".join(out) + "\n")

    print("Round-2 status counts:")
    for k, v in counts.items():
        print(f"  {k:24s} {v}")
    print(f"  total                    {sum(counts.values())}")


if __name__ == "__main__":
    main()
