#!/usr/bin/env python3
"""Detect right-margin overflow in a PDF: text glyphs past the body's right edge.

Robust to footers/headers: the type area's right edge is inferred from the
modal (most frequent) right x-coordinate of full body-text lines, not from
the 95th percentile of every glyph (which would include figures' captions
that bleed wider than the column).

This catches keyword/long-string/listing/TikZ overflows that pdflatex's
\\hbox warnings missed — for the thesis at hand, pdflatex reports zero
overfulls but visual overflow can still occur in cells that aren't \\hbox.
"""
import sys
from collections import Counter, defaultdict
import fitz

PDF = sys.argv[1] if len(sys.argv) > 1 else "main.pdf"
TOL = float(sys.argv[2]) if len(sys.argv) > 2 else 3.0

doc = fitz.open(PDF)

# Pass 1: collect right-edge x of *long* body lines (>50% of page width) to
# infer the modal right margin. Bin to 1pt to absorb tiny kerning differences.
right_edges = Counter()
left_edges = Counter()
for page in doc:
    pw = page.rect.width
    for block in page.get_text("dict")["blocks"]:
        if block.get("type", 0) != 0:
            continue
        for line in block.get("lines", []):
            spans = line.get("spans", [])
            if not spans:
                continue
            x0 = min(s["bbox"][0] for s in spans)
            x1 = max(s["bbox"][2] for s in spans)
            if (x1 - x0) > pw * 0.5:
                right_edges[round(x1)] += 1
                left_edges[round(x0)] += 1

right_edge = right_edges.most_common(1)[0][0]
left_edge = left_edges.most_common(1)[0][0]

print(f"# PDF: {PDF}  pages={doc.page_count}")
print(f"# Modal text body: x=[{left_edge}, {right_edge}]  tolerance={TOL}pt")
print()

# Pass 2: flag any span whose right x exceeds right_edge + TOL.
# Skip the page-number footer zone (bottom ~30pt of the page).
overflows = []
for pageno, page in enumerate(doc, start=1):
    ph = page.rect.height
    footer_y = ph - 40   # ignore anything below this (page numbers)
    header_y = 50        # ignore anything above this (running heads)
    for block in page.get_text("dict")["blocks"]:
        if block.get("type", 0) != 0:
            continue
        for line in block.get("lines", []):
            for span in line.get("spans", []):
                x0, y0, x1, y1 = span["bbox"]
                text = span["text"].rstrip()
                if not text.strip():
                    continue
                if y1 > footer_y or y0 < header_y:
                    continue
                over = x1 - right_edge
                if over > TOL:
                    overflows.append({
                        "page": pageno,
                        "over": over,
                        "x0": x0, "x1": x1, "y": y0,
                        "font": span.get("font", "?"),
                        "size": span.get("size", 0),
                        "text": text,
                    })

by_page = defaultdict(list)
for o in overflows:
    by_page[o["page"]].append(o)

print(f"# {len(overflows)} overflowing spans across {len(by_page)} pages "
      f"(tolerance {TOL}pt)")
print()
for page in sorted(by_page):
    spans = sorted(by_page[page], key=lambda o: o["over"], reverse=True)
    print(f"=== page {page} — worst {min(len(spans), 5)} of {len(spans)} ===")
    for o in spans[:5]:
        snippet = o["text"][:100].replace("\n", " ")
        mono = "TT" in o["font"].upper() or "MONO" in o["font"].upper()
        tag = "mono" if mono else "prose"
        print(f"  +{o['over']:5.1f}pt  y={o['y']:.0f}  font={o['font']}@{o['size']:.1f} [{tag}]")
        print(f"    {snippet!r}")
