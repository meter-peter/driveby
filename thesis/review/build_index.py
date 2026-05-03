#!/usr/bin/env python3
"""
Build a triage index from the extracted Kritikos comments.

Reads kritikos-comments.md and produces comment-index.md with one row per
comment: id, page, chapter, type-tag, priority, status.

Type tags (mutually exclusive, first match wins):
  alpha     — Chapter 8 ownership (any p.134-145 comment, or keyword "AI/agent/authorship")
  beta-name — Methodology renaming / Contribution 1 reframing (Ch.3, ids ~155-242)
  beta-diag — Diagram requests / redesigns ("diagram", "figure", "context", "class", "package")
  beta-flow — Page-overflow complaints (literal "overflow")
  beta-ref  — Missing reference / footnote / URL
  beta-acro — Acronym not introduced
  gamma     — Evaluation re-runs / inconsistencies (Ch.7, p.119-133)
  typo      — Tiny single-word substitutions ("an", "also", "a", "software", etc.)
  struct    — Lists, abstract, committee names (cover/TOC level)
  generic   — Anything else; default

Priority:
  must     — Email's three buckets (alpha/beta-name/beta-diag/gamma) + alpha + struct items
            from cover (lists/abstracts).
  should   — Other beta comments (overflow, refs, acronyms), and typo/struct.
  defer    — Generic discussion / "future direction" comments — escalate to PhD/future-work
            sections rather than addressing in this revision.

Status starts as "open". A separate `changes-log.md` file (maintained by the
human) records when each comment is closed.
"""
from __future__ import annotations

import re
import sys
from pathlib import Path

CHAPTER_BOUNDS = [
    (1, 6, "front-matter"),
    (7, 17, "ch1-introduction"),
    (18, 33, "ch2-related-work"),
    (34, 46, "ch3-methodology"),
    (47, 67, "ch4-cli-architecture"),
    (68, 97, "ch5-kubernetes"),
    (98, 118, "ch6-gitops"),
    (119, 133, "ch7-evaluation"),
    (134, 145, "ch8-ai-assisted"),
    (146, 153, "ch9-discussion"),
    (154, 165, "ch10-conclusion"),
]


def chapter_for(page: int) -> str:
    for lo, hi, name in CHAPTER_BOUNDS:
        if lo <= page <= hi:
            return name
    return "unknown"


TYPO_TEXTS = {
    "an", "a", "also", "software", "the", "in this thesis",
    "respectively,", ", respectively,", "however, the",
    "their current work", "structural and informative?",
    "?", "structural and informative?", "(developed by the platform team)",
    "(see Section 5.6.2 below)", "(PersistentVolumeClaim)", "(Related to RQ1)",
    "(between the agents)", "an\n", "the development of these ...",
}


def classify(comment: dict) -> tuple[str, str]:
    """Return (type_tag, priority)."""
    text = comment["text"]
    text_lc = text.lower()
    chapter = comment["chapter"]
    page = comment["page"]

    # Page-overflow complaints — mechanical fix, classify regardless of chapter
    if "overflow" in text_lc or "overflows" in text_lc:
        return ("beta-flow", "must")

    # Front-matter housekeeping (lists, abstract, committee)
    if page <= 6:
        return ("struct", "must")

    # Tiny single-token edits
    if len(text) < 25 and not text.endswith("?") and "\n" not in text:
        return ("typo", "should")
    if text.strip().lower() in TYPO_TEXTS:
        return ("typo", "should")

    # Diagram requests / redesigns
    if any(
        kw in text_lc
        for kw in [
            "diagram", "figure", "image", "depict", "depicted",
            "package diagram", "class diagram", "context diagram",
            "uml", "data flow", "architecture and data flow",
        ]
    ):
        return ("beta-diag", "must")

    # Methodology naming / Contribution 1 reframing — heavily clustered in Ch.3
    if chapter == "ch3-methodology":
        # Strong signals
        if any(
            kw in text_lc
            for kw in [
                "methodology", "specification-driven", "specification-first",
                "sdt", "sft", "axiom", "principle", "framework",
                "validation theory", "conceptual framework",
                "completeness", "determinism", "observability",
                "minimal", "strict", "test-only", "test-ready",
            ]
        ):
            return ("beta-name", "must")
        return ("beta-name", "should")

    # Chapter 8: thesis ownership
    if chapter == "ch8-ai-assisted":
        return ("alpha", "must")

    # Chapter 7: evaluation
    if chapter == "ch7-evaluation":
        return ("gamma", "must")

    # Missing reference / footnote / URL
    if any(
        kw in text_lc
        for kw in ["footnote", "reference", "please provide a reference", "url", "github"]
    ):
        return ("beta-ref", "should")

    # Acronyms
    if "acronym" in text_lc:
        return ("beta-acro", "should")

    # Structural — research questions, contributions table, etc.
    if any(
        kw in text_lc
        for kw in [
            "list of figures", "list of tables", "abstract", "research question",
            "contribution", "mapping table", "examination committee",
        ]
    ):
        return ("struct", "must")

    # Long discussion / "future direction" prompts
    if len(text) > 400 and any(
        kw in text_lc
        for kw in [
            "future work", "ph.d", "doctoral", "interesting direction",
            "could be extended", "extension", "could explore", "in the future",
        ]
    ):
        return ("generic", "defer")

    # Default
    return ("generic", "should")


def parse_comments(md_path: Path) -> list[dict]:
    text = md_path.read_text(encoding="utf-8")
    comments: list[dict] = []
    current_page: int | None = None
    page_re = re.compile(r"^## Page (\d+)\s*$")
    cmt_re = re.compile(r"^\*\*\[(\d+)\]\*\*\s+\((\w+)\)\s+(.*)$", re.DOTALL)
    buf: list[str] = []
    cur_id: int | None = None
    cur_subtype: str | None = None
    cur_first_line: str | None = None

    def flush() -> None:
        if cur_id is None or current_page is None:
            return
        body = (cur_first_line or "") + "\n" + "\n".join(buf) if buf else (cur_first_line or "")
        comments.append(
            {
                "id": cur_id,
                "page": current_page,
                "subtype": cur_subtype,
                "text": body.strip(),
                "chapter": chapter_for(current_page),
            }
        )

    for line in text.splitlines():
        pm = page_re.match(line)
        if pm:
            flush()
            cur_id = None
            buf = []
            current_page = int(pm.group(1))
            continue
        cm = cmt_re.match(line)
        if cm:
            flush()
            cur_id = int(cm.group(1))
            cur_subtype = cm.group(2)
            cur_first_line = cm.group(3)
            buf = []
            continue
        if cur_id is not None:
            if line.strip() or buf:
                buf.append(line)
    flush()

    for c in comments:
        c["type"], c["priority"] = classify(c)
    return comments


def write_index(comments: list[dict], out_path: Path) -> None:
    by_priority = {"must": 0, "should": 0, "defer": 0}
    by_type: dict[str, int] = {}
    by_chapter: dict[str, int] = {}
    for c in comments:
        by_priority[c["priority"]] = by_priority.get(c["priority"], 0) + 1
        by_type[c["type"]] = by_type.get(c["type"], 0) + 1
        by_chapter[c["chapter"]] = by_chapter.get(c["chapter"], 0) + 1

    lines = [
        "# Kritikos Comment Index",
        "",
        f"Total: **{len(comments)}** comments across "
        f"**{len({c['page'] for c in comments})}** pages.",
        "",
        "## Distribution",
        "",
        "### By priority",
        "",
        "| priority | count | meaning |",
        "|---|---|---|",
        f"| must | {by_priority.get('must', 0)} | address before resubmission |",
        f"| should | {by_priority.get('should', 0)} | address if time permits |",
        f"| defer | {by_priority.get('defer', 0)} | route to PhD-continuation / future-work sections |",
        "",
        "### By type",
        "",
        "| type | count | description |",
        "|---|---|---|",
    ]
    type_descriptions = {
        "alpha": "Ch.8 thesis ownership (group α from cover email)",
        "beta-name": "Methodology renaming / Contribution 1 reframing",
        "beta-diag": "Diagram request / redesign",
        "beta-flow": "Page-overflow complaint",
        "beta-ref": "Missing reference / footnote / URL",
        "beta-acro": "Acronym not introduced",
        "gamma": "Evaluation re-run / inconsistency (Ch.7)",
        "typo": "Tiny single-word fix",
        "struct": "Front-matter / structural element (lists, abstract, committee)",
        "generic": "Other (often discussion / future-work prompts)",
    }
    for t in [
        "alpha", "beta-name", "beta-diag", "beta-flow", "beta-ref", "beta-acro",
        "gamma", "typo", "struct", "generic",
    ]:
        lines.append(f"| {t} | {by_type.get(t, 0)} | {type_descriptions[t]} |")
    lines += [
        "",
        "### By chapter",
        "",
        "| chapter | count |",
        "|---|---|",
    ]
    for _lo, _hi, name in CHAPTER_BOUNDS:
        lines.append(f"| {name} | {by_chapter.get(name, 0)} |")
    lines += [
        "",
        "## Index",
        "",
        "Each row links back to the full comment text in `kritikos-comments.md`.",
        "Status starts as `open`; record closures in `changes-log.md` and update here.",
        "",
        "| id | page | chapter | type | priority | status | excerpt |",
        "|---:|---:|---|---|---|---|---|",
    ]
    for c in comments:
        excerpt = c["text"].replace("\n", " ").replace("|", "\\|")
        if len(excerpt) > 110:
            excerpt = excerpt[:107] + "..."
        lines.append(
            f"| [{c['id']}](kritikos-comments.md#page-{c['page']}) | "
            f"{c['page']} | {c['chapter']} | {c['type']} | {c['priority']} | open | {excerpt} |"
        )
    out_path.write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> int:
    here = Path(__file__).parent
    comments = parse_comments(here / "kritikos-comments.md")
    write_index(comments, here / "comment-index.md")
    print(f"Indexed {len(comments)} comments → comment-index.md")
    return 0


if __name__ == "__main__":
    sys.exit(main())
