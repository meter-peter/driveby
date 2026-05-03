#!/usr/bin/env python3
"""
Extract every PDF Text annotation (sticky note) from a reviewed thesis PDF.

Usage:
    python3 extract_comments.py <input.pdf> [output.md]

Output: a Markdown file grouped by page, sorted top-to-bottom within each page.
Each comment is numbered globally so it can be referenced as "[123]".

The script avoids external dependencies — it parses PDF objects directly,
including those bundled inside compressed FlateDecode object streams (PDF 1.5+).

Tested on Poppler-produced PDFs (pdfTeX, hyperref). Adobe Reader-style
sticky notes appear as /Subtype /Text annotations with a /Contents string.
"""
from __future__ import annotations

import argparse
import re
import sys
import zlib
from pathlib import Path

OBJ_PATTERN = re.compile(rb"(\d+)\s+(\d+)\s+obj\b(.*?)\bendobj\b", re.DOTALL)


def parse_top_objects(data: bytes) -> dict[int, bytes]:
    out: dict[int, bytes] = {}
    for m in OBJ_PATTERN.finditer(data):
        out[int(m.group(1))] = m.group(3)
    return out


def get_stream_bytes(body: bytes) -> bytes | None:
    sm = re.search(rb"stream\r?\n(.*?)\r?\nendstream", body, re.DOTALL)
    return sm.group(1) if sm else None


def extract_objstm_objects(top_objects: dict[int, bytes]) -> dict[int, bytes]:
    """PDF 1.5+ packs many objects into compressed object streams.
    For each /Type /ObjStm, decompress and pull out the bundled objects."""
    extra: dict[int, bytes] = {}
    for body in top_objects.values():
        if not (b"/Type /ObjStm" in body or b"/Type/ObjStm" in body):
            continue
        raw = get_stream_bytes(body)
        if raw is None:
            continue
        if b"/FlateDecode" in body:
            try:
                stream_bytes = zlib.decompress(raw)
            except zlib.error:
                continue
        else:
            stream_bytes = raw
        n_match = re.search(rb"/N\s+(\d+)", body)
        first_match = re.search(rb"/First\s+(\d+)", body)
        if not n_match or not first_match:
            continue
        n = int(n_match.group(1))
        first = int(first_match.group(1))
        header, body_part = stream_bytes[:first], stream_bytes[first:]
        pairs = re.findall(rb"(\d+)\s+(\d+)", header)[:n]
        for i, (sub_onum, sub_off) in enumerate(pairs):
            sub_off = int(sub_off)
            sub_end = int(pairs[i + 1][1]) if i + 1 < n else len(body_part)
            extra[int(sub_onum)] = body_part[sub_off:sub_end]
    return extra


def get_ref(body: bytes, key: str) -> int | None:
    m = re.search(rb"/" + key.encode() + rb"\s+(\d+)\s+(\d+)\s+R", body)
    return int(m.group(1)) if m else None


def get_kids(body: bytes) -> list[int]:
    m = re.search(rb"/Kids\s*\[(.*?)\]", body, re.DOTALL)
    if not m:
        return []
    return [int(x) for x in re.findall(rb"(\d+)\s+\d+\s+R", m.group(1))]


def walk_pages(node_id: int, all_objects: dict[int, bytes], out: list[int], depth: int = 0) -> None:
    if depth > 30 or node_id not in all_objects:
        return
    body = all_objects[node_id]
    if b"/Type /Pages" in body or b"/Type/Pages" in body:
        for kid in get_kids(body):
            walk_pages(kid, all_objects, out, depth + 1)
    elif b"/Type /Page" in body or b"/Type/Page" in body:
        out.append(node_id)


def parse_annot_refs(body: bytes, all_objects: dict[int, bytes]) -> list[int]:
    m = re.search(rb"/Annots\s*\[(.*?)\]", body, re.DOTALL)
    if m:
        return [int(x) for x in re.findall(rb"(\d+)\s+\d+\s+R", m.group(1))]
    m2 = re.search(rb"/Annots\s+(\d+)\s+\d+\s+R", body)
    if m2:
        ref_obj = all_objects.get(int(m2.group(1)))
        if ref_obj:
            m3 = re.search(rb"\[(.*?)\]", ref_obj, re.DOTALL)
            if m3:
                return [int(x) for x in re.findall(rb"(\d+)\s+\d+\s+R", m3.group(1))]
    return []


def extract_contents(body: bytes) -> str | None:
    m = re.search(rb"/Contents\s*\(((?:[^()\\]|\\.|\\\d{1,3})*)\)", body, re.DOTALL)
    if m:
        text = m.group(1).decode("latin-1", errors="replace")
        text = text.replace("\\(", "(").replace("\\)", ")").replace("\\\\", "\\")
        return text.replace("\\r", "\n").replace("\\n", "\n").replace("\r", "\n")
    m2 = re.search(rb"/Contents\s*<([0-9A-Fa-f\s]+)>", body)
    if m2:
        hex_str = re.sub(rb"\s", b"", m2.group(1))
        try:
            raw = bytes.fromhex(hex_str.decode())
            if raw[:2] == b"\xfe\xff":
                return raw[2:].decode("utf-16-be", errors="replace")
            return raw.decode("utf-8", errors="replace")
        except (ValueError, UnicodeDecodeError):
            return None
    return None


def extract_rect(body: bytes) -> tuple[float, ...] | None:
    m = re.search(
        rb"/Rect\s*\[\s*([\d.\-]+)\s+([\d.\-]+)\s+([\d.\-]+)\s+([\d.\-]+)\s*\]",
        body,
    )
    return tuple(float(m.group(i)) for i in range(1, 5)) if m else None


def extract_subtype(body: bytes) -> str | None:
    m = re.search(rb"/Subtype\s*/(\w+)", body)
    return m.group(1).decode() if m else None


def extract_all(pdf_path: Path) -> list[dict]:
    data = pdf_path.read_bytes()
    top_objects = parse_top_objects(data)
    all_objects = {**top_objects, **extract_objstm_objects(top_objects)}

    catalog_id = next(
        (
            onum
            for onum, body in all_objects.items()
            if b"/Type /Catalog" in body or b"/Type/Catalog" in body
        ),
        None,
    )
    page_order: list[int] = []
    if catalog_id is not None:
        pages_root = get_ref(all_objects[catalog_id], "Pages")
        if pages_root is not None:
            walk_pages(pages_root, all_objects, page_order)

    comments: list[dict] = []
    for page_idx, page_obj_id in enumerate(page_order, start=1):
        page_body = all_objects[page_obj_id]
        for aid in parse_annot_refs(page_body, all_objects):
            abody = all_objects.get(aid)
            if not abody:
                continue
            subtype = extract_subtype(abody)
            if subtype not in {"Text", "FreeText", "Highlight", "Square", "Note"}:
                continue
            contents = extract_contents(abody)
            if not contents:
                continue
            rect = extract_rect(abody)
            comments.append(
                {
                    "page": page_idx,
                    "subtype": subtype,
                    "rect": rect,
                    "ykey": -rect[3] if rect else 0.0,
                    "text": contents.strip(),
                }
            )

    comments.sort(key=lambda c: (c["page"], c["ykey"]))
    return comments


def write_markdown(comments: list[dict], pdf_path: Path, out_path: Path) -> None:
    lines = [
        "# Prof. Kritikos — Inline PDF Comments (extracted)",
        "",
        f"Source: `{pdf_path}`",
        f"Total comments: {len(comments)}",
        "",
        "Regenerate with: `python3 extract_comments.py <pdf> <md>`",
        "",
    ]
    cur_page = None
    for i, c in enumerate(comments, 1):
        if c["page"] != cur_page:
            cur_page = c["page"]
            lines.append("")
            lines.append(f"## Page {cur_page}")
            lines.append("")
        lines.append(f"**[{i}]** ({c['subtype']}) {c['text']}")
        lines.append("")
    out_path.write_text("\n".join(lines), encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("pdf", type=Path, help="annotated PDF")
    parser.add_argument(
        "output",
        type=Path,
        nargs="?",
        default=Path("kritikos-comments.md"),
        help="output markdown file (default: kritikos-comments.md)",
    )
    args = parser.parse_args()
    if not args.pdf.exists():
        print(f"error: {args.pdf} not found", file=sys.stderr)
        return 1
    comments = extract_all(args.pdf)
    write_markdown(comments, args.pdf, args.output)
    print(f"Extracted {len(comments)} comments → {args.output}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
