#!/usr/bin/env python3
"""Generate DDT thesis presentations in Greek and English.

Content sourced directly from thesis chapters 01-10.
"""

from pptx import Presentation
from pptx.util import Inches, Pt
from pptx.dml.color import RGBColor
from pptx.enum.text import PP_ALIGN
from pptx.enum.shapes import MSO_SHAPE
import os

# ── Color palette ──────────────────────────────────────────────────────
DARK_BG       = RGBColor(0x1A, 0x1A, 0x2E)
ACCENT_BLUE   = RGBColor(0x00, 0x96, 0xC7)
ACCENT_GREEN  = RGBColor(0x2E, 0xCC, 0x71)
ACCENT_RED    = RGBColor(0xE7, 0x4C, 0x3C)
ACCENT_ORANGE = RGBColor(0xF3, 0x9C, 0x12)
WHITE         = RGBColor(0xFF, 0xFF, 0xFF)
LIGHT_GRAY    = RGBColor(0xBD, 0xBD, 0xBD)
MID_GRAY      = RGBColor(0x66, 0x66, 0x66)
DARK_TEXT      = RGBColor(0x2C, 0x2C, 0x2C)
LIGHT_BG      = RGBColor(0xF5, 0xF7, 0xFA)
CODE_BG       = RGBColor(0x28, 0x2C, 0x34)
TABLE_ALT     = RGBColor(0xE8, 0xEA, 0xED)
TABLE_ALT2    = RGBColor(0xF5, 0xF5, 0xF5)

SLIDE_WIDTH  = Inches(13.333)
SLIDE_HEIGHT = Inches(7.5)


def set_slide_bg(slide, color):
    bg = slide.background
    fill = bg.fill
    fill.solid()
    fill.fore_color.rgb = color


def add_text(slide, left, top, width, height, text, size=18,
             color=WHITE, bold=False, align=PP_ALIGN.LEFT,
             font="Segoe UI"):
    box = slide.shapes.add_textbox(left, top, width, height)
    tf = box.text_frame
    tf.word_wrap = True
    p = tf.paragraphs[0]
    p.text = text
    p.font.size = Pt(size)
    p.font.color.rgb = color
    p.font.bold = bold
    p.font.name = font
    p.alignment = align
    return box


def add_multiline(slide, left, top, width, height, lines, size=16,
                  color=WHITE, font="Segoe UI", spacing=Pt(6), bold=False):
    """Add multiple lines as separate paragraphs."""
    box = slide.shapes.add_textbox(left, top, width, height)
    tf = box.text_frame
    tf.word_wrap = True
    for i, line in enumerate(lines):
        p = tf.paragraphs[0] if i == 0 else tf.add_paragraph()
        p.text = line
        p.font.size = Pt(size)
        p.font.color.rgb = color
        p.font.name = font
        p.font.bold = bold
        p.space_after = spacing
    return box


def add_box(slide, left, top, width, height, fill_color, text="",
            size=14, text_color=WHITE, bold=True, align=PP_ALIGN.CENTER,
            font="Segoe UI"):
    shape = slide.shapes.add_shape(MSO_SHAPE.ROUNDED_RECTANGLE,
                                   left, top, width, height)
    shape.fill.solid()
    shape.fill.fore_color.rgb = fill_color
    shape.line.fill.background()
    if hasattr(shape, 'adjustments') and len(shape.adjustments) > 0:
        shape.adjustments[0] = 0.05
    tf = shape.text_frame
    tf.word_wrap = True
    p = tf.paragraphs[0]
    p.text = text
    p.font.size = Pt(size)
    p.font.color.rgb = text_color
    p.font.bold = bold
    p.font.name = font
    p.alignment = align
    return shape


def add_line(slide, left, top, width, color=ACCENT_BLUE):
    shape = slide.shapes.add_shape(MSO_SHAPE.RECTANGLE, left, top, width, Pt(3))
    shape.fill.solid()
    shape.fill.fore_color.rgb = color
    shape.line.fill.background()


def add_code_block(slide, left, top, width, height, code, size=12):
    """Code block with monospace font."""
    add_box(slide, left, top, width, height, CODE_BG, code,
            size=size, text_color=ACCENT_GREEN, bold=False,
            align=PP_ALIGN.LEFT, font="Consolas")


def build(lang="en"):
    prs = Presentation()
    prs.slide_width = SLIDE_WIDTH
    prs.slide_height = SLIDE_HEIGHT
    blank = prs.slide_layouts[6]
    t = TR[lang]

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 1: TITLE
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_line(s, Inches(1), Inches(1.5), Inches(3))
    add_text(s, Inches(1), Inches(1.7), Inches(11), Inches(1.2),
             "Documentation-Driven Testing (DDT)", size=36, bold=True)
    add_text(s, Inches(1), Inches(3.0), Inches(11), Inches(0.8),
             t["subtitle"], size=22, color=ACCENT_BLUE)
    add_text(s, Inches(1), Inches(4.2), Inches(11), Inches(0.5),
             t["author"], size=16, color=LIGHT_GRAY)
    add_text(s, Inches(1), Inches(4.8), Inches(11), Inches(0.5),
             t["institution"], size=14, color=MID_GRAY)
    add_text(s, Inches(1), Inches(5.3), Inches(11), Inches(0.5),
             "2026", size=14, color=MID_GRAY)
    add_text(s, Inches(9), Inches(5.8), Inches(3.5), Inches(0.8),
             "DriveBy", size=28, color=ACCENT_BLUE, bold=True,
             align=PP_ALIGN.RIGHT, font="Consolas")

    # University logo
    logo_path = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                             "..", "thesis", "figures", "uoa-logo.png")
    if os.path.exists(logo_path):
        s.shapes.add_picture(logo_path, Inches(1), Inches(6.0),
                             width=Inches(2.5))

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 2: THE PROBLEM (with real thesis quotes)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s2_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_RED)

    # Real thesis quote (Ch.1)
    add_box(s, Inches(0.8), Inches(1.5), Inches(11.5), Inches(1.3),
            CODE_BG,
            t["s2_quote"],
            size=15, text_color=ACCENT_BLUE, bold=False)

    # 3 real problem statements from Ch.1 + Ch.2
    problems = t["s2_problems"]
    for i, (title, desc) in enumerate(problems):
        x = Inches(0.8) + i * Inches(4.05)
        add_box(s, x, Inches(3.2), Inches(3.7), Inches(0.7),
                ACCENT_RED, title, size=14)
        add_text(s, x + Inches(0.1), Inches(4.0), Inches(3.5), Inches(2.5),
                 desc, size=12, color=DARK_TEXT, align=PP_ALIGN.LEFT)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 3: KEY INSIGHT (real Ch.1 thesis statement)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(1), Inches(2.0), Inches(11), Inches(1.5),
             t["s3_question"], size=36, bold=True, align=PP_ALIGN.CENTER)
    add_line(s, Inches(5), Inches(3.7), Inches(3.3))
    add_text(s, Inches(1.5), Inches(4.2), Inches(10), Inches(1.5),
             t["s3_answer"], size=20, color=ACCENT_BLUE, align=PP_ALIGN.CENTER)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 4: THREE AXIOMS (real definitions from Ch.3)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s4_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    axioms = t["s4_axioms"]
    colors = [ACCENT_BLUE, ACCENT_GREEN, ACCENT_ORANGE]
    for i, (name, formal, principles) in enumerate(axioms):
        x = Inches(0.8) + i * Inches(4.05)
        add_box(s, x, Inches(1.5), Inches(3.7), Inches(0.8),
                colors[i], name, size=18)
        add_text(s, x + Inches(0.15), Inches(2.5), Inches(3.4), Inches(1.5),
                 formal, size=12, color=DARK_TEXT, bold=True)
        add_text(s, x + Inches(0.15), Inches(3.8), Inches(3.4), Inches(2),
                 principles, size=12, color=MID_GRAY)

    add_text(s, Inches(0.8), Inches(6.0), Inches(11), Inches(0.5),
             t["s4_bottom"], size=14, color=MID_GRAY, align=PP_ALIGN.CENTER)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 5: APISPEC INTERFACE (real Go code from Ch.4)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s5_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Real APISpec interface from spec/spec.go
    add_code_block(s, Inches(0.8), Inches(1.5), Inches(5.8), Inches(3.5),
        "type APISpec interface {\n"
        "    Type() SpecType\n"
        "    RawVersion() string\n"
        "    Info() *SpecInfo\n"
        "    Paths() map[string]*PathItem\n"
        "    Components() *Components\n"
        "    Security() []SecurityRequirement\n"
        "    ValidateStructure(ctx) error\n"
        "}\n\n"
        "// Adapter pattern (GoF):\n"
        "// openapi3.go  -> OpenAPI 3.0.x, 3.1.0\n"
        "// swagger2.go  -> Swagger 2.0\n"
        "// No raw *openapi3.T in validators",
        size=13)

    # Why it matters
    add_text(s, Inches(7.2), Inches(1.5), Inches(5.5), Inches(0.5),
             t["s5_right_title"], size=20, color=ACCENT_BLUE, bold=True)
    add_multiline(s, Inches(7.2), Inches(2.2), Inches(5.5), Inches(4),
                  t["s5_points"], size=14, color=LIGHT_GRAY)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 6: PRINCIPLE CHECKER PATTERN (real Go code from Ch.4)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s6_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # PrincipleChecker interface
    add_code_block(s, Inches(0.8), Inches(1.5), Inches(5.8), Inches(2),
        "type PrincipleChecker interface {\n"
        "    ID() string\n"
        "    Check(ctx, spec.APISpec,\n"
        "          mode ValidationMode) PrincipleResult\n"
        "}",
        size=13)

    # Registry ForMode — real code from registry.go
    add_code_block(s, Inches(0.8), Inches(3.8), Inches(5.8), Inches(3),
        "func (r *Registry) ForMode(mode) []PrincipleChecker {\n"
        "  switch mode {\n"
        "  case Minimal:   return {P001}\n"
        "  case TestReady: return {P001, P004, P009}\n"
        "  case Strict:    return {P001, P002, P003,\n"
        "                          P004, P005, P008}\n"
        "  case TestOnly:  return nil // runtime only\n"
        "  }\n"
        "}",
        size=12)

    # Real checker table from Ch.4
    add_text(s, Inches(7.2), Inches(1.5), Inches(5.3), Inches(0.5),
             t["s6_table_title"], size=16, color=ACCENT_BLUE, bold=True)
    checkers = [
        ("p001_compliance.go", "P001", "7", "180"),
        ("p002_documentation.go", "P002", "10", "197"),
        ("p003_errors.go", "P003", "7", "234"),
        ("p004_schema.go", "P004", "10", "432"),
        ("p005_security.go", "P005", "7", "221"),
        ("p008_versioning.go", "P008", "7", "163"),
        ("p009_test_readiness.go", "P009", "4", "188"),
    ]
    y = Inches(2.1)
    for f, pid, checks, loc in checkers:
        add_box(s, Inches(7.2), y, Inches(1.0), Inches(0.4),
                ACCENT_BLUE, pid, size=11)
        add_box(s, Inches(8.3), y, Inches(2.8), Inches(0.4),
                TABLE_ALT, f, size=10, text_color=DARK_TEXT, bold=False,
                font="Consolas")
        add_box(s, Inches(11.2), y, Inches(0.7), Inches(0.4),
                TABLE_ALT, checks, size=10, text_color=DARK_TEXT, bold=False)
        add_box(s, Inches(12.0), y, Inches(0.7), Inches(0.4),
                TABLE_ALT, loc, size=10, text_color=DARK_TEXT, bold=False)
        y += Inches(0.48)

    add_text(s, Inches(7.2), Inches(5.6), Inches(5.3), Inches(1),
             t["s6_total"], size=13, color=LIGHT_GRAY)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 7: VALIDATION ENGINE (real code from Ch.4)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s7_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Real ValidateSpec pipeline from engine.go
    add_code_block(s, Inches(0.8), Inches(1.5), Inches(7), Inches(4.2),
        "func (e *Engine) ValidateSpec(ctx) (*ValidationReport, error) {\n"
        "    // Step 1: Load and normalize the specification\n"
        "    err := e.loader.LoadFromFileOrURL(e.config.SpecPath)\n"
        "    doc := e.loader.GetDocument()\n"
        "\n"
        "    // Step 2: Initialize report\n"
        "    report := &ValidationReport{\n"
        "        Version: e.config.Version,\n"
        "        Timestamp: time.Now(),\n"
        "    }\n"
        "\n"
        "    // Step 3: Select and run principle checkers\n"
        "    checkers := e.registry.ForMode(e.config.ValidationMode)\n"
        "    for _, checker := range checkers {\n"
        "        result := checker.Check(ctx, doc, mode)\n"
        "        report.Principles = append(report.Principles, result)\n"
        "    }\n"
        "\n"
        "    // Step 4: Compute summary\n"
        "    updateSummary(report)\n"
        "    return report, nil\n"
        "}",
        size=12)

    # Dependency flow from Ch.4
    add_text(s, Inches(8.3), Inches(1.5), Inches(4.5), Inches(0.5),
             t["s7_dep_title"], size=18, color=DARK_TEXT, bold=True)
    deps = ["types    (655 LOC)  -- zero deps",
            "spec     (831 LOC)  -- APISpec + adapters",
            "loader   (322 LOC)  -- version detection",
            "principles (1,690 LOC) -- 7 checkers",
            "testing  (860 LOC)  -- P006 + P007",
            "engine   (117 LOC)  -- orchestrator",
            "cli      (907 LOC)  -- Cobra commands"]
    add_multiline(s, Inches(8.3), Inches(2.2), Inches(4.5), Inches(3.5),
                  deps, size=12, color=DARK_TEXT, font="Consolas",
                  spacing=Pt(8))

    add_text(s, Inches(8.3), Inches(5.2), Inches(4.5), Inches(1),
             t["s7_note"], size=12, color=MID_GRAY)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 8: FROM CLI TO CUSTOM RESOURCE (5 stages from Ch.5)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s8_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Real 5-stage evolution from Ch.5 sec:cli-to-cr-evolution
    stages = t["s8_stages"]
    stage_colors = [MID_GRAY, ACCENT_BLUE, ACCENT_BLUE, ACCENT_GREEN, ACCENT_GREEN]
    for i, (label, desc) in enumerate(stages):
        y = Inches(1.5) + i * Inches(1.1)
        add_box(s, Inches(0.8), y, Inches(0.7), Inches(0.7),
                stage_colors[i], str(i + 1), size=22)
        add_box(s, Inches(1.7), y, Inches(3), Inches(0.7),
                stage_colors[i], label, size=13)
        add_text(s, Inches(4.9), y + Inches(0.1), Inches(7.8), Inches(0.6),
                 desc, size=12, color=LIGHT_GRAY)

    # Key insight quote from Ch.5
    add_box(s, Inches(0.8), Inches(6.3), Inches(11.5), Inches(0.7),
            CODE_BG, t["s8_quote"], size=13, text_color=ACCENT_BLUE, bold=False)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 9: XSDLC YAML + GENERATED RESOURCES (real from Ch.5)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s9_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_GREEN)

    # Real XSDLC CR from Ch.5 Listing 10
    add_code_block(s, Inches(0.8), Inches(1.5), Inches(5.5), Inches(5.0),
        "apiVersion: driveby.io/v1alpha1\n"
        "kind: XSDLC\n"
        "metadata:\n"
        "  name: perfect-api\n"
        "  namespace: driveby\n"
        "spec:\n"
        "  gitopsRepository:\n"
        "    owner: novelcore\n"
        "    name: perfect-api-gitops\n"
        "  environments:\n"
        "    - name: dev\n"
        "    - name: staging\n"
        "      gate:\n"
        "        checks:\n"
        "          - type: validate-only\n"
        "          - type: functional-test\n"
        "    - name: prod\n"
        "      autoMerge: false\n"
        "      gate:\n"
        "        checks:\n"
        "          - type: validate-only\n"
        "          - type: load-test",
        size=12)

    # Generated resources breakdown (real from Ch.5)
    add_text(s, Inches(6.8), Inches(1.5), Inches(5.8), Inches(0.5),
             t["s9_gen_title"], size=18, color=ACCENT_GREEN, bold=True)
    resources = t["s9_resources"]
    y = Inches(2.2)
    for category, items in resources:
        add_box(s, Inches(6.8), y, Inches(5.8), Inches(0.45),
                ACCENT_BLUE, category, size=11)
        y += Inches(0.5)
        for item in items:
            add_text(s, Inches(7.0), y, Inches(5.5), Inches(0.35),
                     item, size=10, color=LIGHT_GRAY, font="Consolas")
            y += Inches(0.3)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 10: RECONCILIATION LOOP (real from Ch.5)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s10_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Real reconciliation explanation from Ch.5
    # Left: level-triggered
    add_box(s, Inches(0.8), Inches(1.5), Inches(5.5), Inches(0.7),
            ACCENT_RED, t["s10_edge"], size=15)
    add_text(s, Inches(0.8), Inches(2.4), Inches(5.5), Inches(1.5),
             t["s10_edge_desc"], size=13, color=DARK_TEXT)

    add_box(s, Inches(7), Inches(1.5), Inches(5.5), Inches(0.7),
            ACCENT_GREEN, t["s10_level"], size=15)
    add_text(s, Inches(7), Inches(2.4), Inches(5.5), Inches(1.5),
             t["s10_level_desc"], size=13, color=DARK_TEXT)

    # Day 0-2-N lifecycle from Ch.5
    add_text(s, Inches(0.8), Inches(4.2), Inches(11), Inches(0.5),
             t["s10_lifecycle_title"], size=18, color=DARK_TEXT, bold=True)
    days = t["s10_days"]
    day_colors = [ACCENT_BLUE, ACCENT_GREEN, ACCENT_ORANGE, MID_GRAY]
    for i, (day, desc) in enumerate(days):
        x = Inches(0.8) + i * Inches(3.1)
        add_box(s, x, Inches(4.9), Inches(2.8), Inches(0.6),
                day_colors[i], day, size=13)
        add_text(s, x + Inches(0.1), Inches(5.6), Inches(2.7), Inches(1.5),
                 desc, size=11, color=DARK_TEXT)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 11: BYOCI DELIVERY PIPELINE (real from Ch.5/6)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s11_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # BYOCI boundary — the key concept
    add_box(s, Inches(0.8), Inches(1.4), Inches(11.5), Inches(0.8),
            CODE_BG, t["s11_byoci"], size=14, text_color=ACCENT_BLUE, bold=True)

    # XSDLC creates ONLY the GitOps repo
    add_box(s, Inches(0.8), Inches(2.6), Inches(5.5), Inches(0.8),
            ACCENT_GREEN, t["s11_gitops_repo"], size=15)
    add_multiline(s, Inches(0.8), Inches(3.6), Inches(5.5), Inches(2),
                  t["s11_gitops_desc"], size=12, color=LIGHT_GRAY)

    # Developer's responsibility (NOT managed by XSDLC)
    add_box(s, Inches(7), Inches(2.6), Inches(5.5), Inches(0.8),
            MID_GRAY, t["s11_dev_owns"], size=15)
    add_multiline(s, Inches(7), Inches(3.6), Inches(5.5), Inches(2),
                  t["s11_dev_desc"], size=12, color=LIGHT_GRAY)

    # Promotion flow — starts from developer committing to gitops repo
    flow = t["s11_flow"]
    flow_colors = [ACCENT_BLUE, ACCENT_ORANGE, ACCENT_GREEN, ACCENT_BLUE, ACCENT_GREEN]
    x_pos = [Inches(0.5), Inches(2.8), Inches(5.1), Inches(7.4), Inches(9.7)]
    for txt, c, x in zip(flow, flow_colors, x_pos):
        add_box(s, x, Inches(5.8), Inches(2.2), Inches(1.0), c, txt, size=11)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 12: CONTROLLED EVALUATION (real data from Ch.7)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s12_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_GREEN)

    # Defect injection matrix — real from Table 7.2
    add_text(s, Inches(0.8), Inches(1.4), Inches(5.5), Inches(0.5),
             t["s12_inject_title"], size=16, color=DARK_TEXT, bold=True)
    defects = [
        ("P001", t["s12_d1"], ACCENT_GREEN),
        ("P002", t["s12_d2"], ACCENT_GREEN),
        ("P003", t["s12_d3"], ACCENT_GREEN),
        ("P004", t["s12_d4"], ACCENT_GREEN),
        ("P005", t["s12_d5"], ACCENT_GREEN),
        ("P008", t["s12_d6"], ACCENT_GREEN),
        ("P009", t["s12_d7"], ACCENT_GREEN),
    ]
    y = Inches(1.9)
    for pid, desc, c in defects:
        add_box(s, Inches(0.8), y, Inches(0.8), Inches(0.38),
                ACCENT_BLUE, pid, size=10)
        bg = TABLE_ALT if defects.index((pid, desc, c)) % 2 == 0 else TABLE_ALT2
        add_box(s, Inches(1.7), y, Inches(4.0), Inches(0.38),
                bg, desc, size=10, text_color=DARK_TEXT, bold=False,
                align=PP_ALIGN.LEFT)
        add_box(s, Inches(5.8), y, Inches(0.5), Inches(0.38),
                c, "v", size=10)
        y += Inches(0.42)

    # Baseline table — real from Table 7.1
    add_text(s, Inches(7), Inches(1.4), Inches(5.5), Inches(0.5),
             t["s12_baseline_title"], size=16, color=DARK_TEXT, bold=True)
    baselines = [
        ("Minimal", "P001", "1/1"),
        ("Test-ready", "P001, P004, P009", "3/3"),
        ("Strict", "P001-P005, P008", "6/6"),
    ]
    y = Inches(1.9)
    for mode, checks, result in baselines:
        add_box(s, Inches(7), y, Inches(2), Inches(0.5),
                ACCENT_BLUE, mode, size=12)
        add_box(s, Inches(9.1), y, Inches(2.4), Inches(0.5),
                TABLE_ALT, checks, size=10, text_color=DARK_TEXT, bold=False)
        add_box(s, Inches(11.6), y, Inches(1), Inches(0.5),
                ACCENT_GREEN, result, size=12)
        y += Inches(0.6)

    # Key result
    add_box(s, Inches(7), Inches(3.8), Inches(5.5), Inches(0.8),
            ACCENT_GREEN,
            t["s12_result"], size=14)

    # Severity model box
    add_text(s, Inches(0.8), Inches(5.2), Inches(11.5), Inches(0.5),
             t["s12_severity_title"], size=16, color=DARK_TEXT, bold=True)
    add_multiline(s, Inches(0.8), Inches(5.7), Inches(11.5), Inches(1.5),
                  t["s12_severity"], size=12, color=DARK_TEXT)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 13: MULTI-API EVALUATION (real from Ch.7 Tables 7.4-7.6)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s13_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Headers
    add_box(s, Inches(0.8), Inches(1.4), Inches(2.2), Inches(0.5),
            MID_GRAY, "API", size=12)
    add_box(s, Inches(3.1), Inches(1.4), Inches(2), Inches(0.5),
            MID_GRAY, t["s13_defect"], size=11)
    add_box(s, Inches(5.2), Inches(1.4), Inches(3.2), Inches(0.5),
            MID_GRAY, "Staging Gate", size=12)
    add_box(s, Inches(8.5), Inches(1.4), Inches(3.5), Inches(0.5),
            MID_GRAY, "Production Gate", size=12)

    # Real data from Tables 7.4, 7.5, 7.6
    apis = t["s13_apis"]
    y = Inches(2.0)
    for api, defect, staging, prod, api_color in apis:
        add_box(s, Inches(0.8), y, Inches(2.2), Inches(0.55),
                api_color, api, size=11)
        add_box(s, Inches(3.1), y, Inches(2), Inches(0.55),
                CODE_BG, defect, size=9, text_color=LIGHT_GRAY, bold=False)
        sc = ACCENT_GREEN if "PASS" in staging else ACCENT_RED
        add_box(s, Inches(5.2), y, Inches(3.2), Inches(0.55),
                sc, staging, size=10)
        pc = ACCENT_GREEN if "PASS" in prod else (ACCENT_RED if "FAIL" in prod else MID_GRAY)
        add_box(s, Inches(8.5), y, Inches(3.5), Inches(0.55),
                pc, prod, size=10)
        y += Inches(0.62)

    # Bottom stat
    add_box(s, Inches(0.8), Inches(5.4), Inches(11.2), Inches(0.8),
            CODE_BG, t["s13_bottom"], size=15, text_color=ACCENT_BLUE)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 14: LAYERED DEFENCE (real from Ch.7 Table 7.7)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s14_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    layers = t["s14_layers"]
    layer_colors = [ACCENT_BLUE, ACCENT_GREEN, ACCENT_ORANGE]
    for i, (layer, checks, defect_class, caught) in enumerate(layers):
        y = Inches(1.6) + i * Inches(1.6)
        add_box(s, Inches(0.8), y, Inches(1.5), Inches(1.2),
                layer_colors[i], layer, size=18)
        add_box(s, Inches(2.5), y, Inches(3.5), Inches(1.2),
                TABLE_ALT, checks, size=13, text_color=DARK_TEXT, bold=False)
        add_box(s, Inches(6.2), y, Inches(2.8), Inches(1.2),
                TABLE_ALT, defect_class, size=13, text_color=DARK_TEXT, bold=False)
        add_box(s, Inches(9.2), y, Inches(3.3), Inches(1.2),
                layer_colors[i], caught, size=13)

    # Real quote from Ch.7
    add_box(s, Inches(0.8), Inches(6.0), Inches(11.5), Inches(0.7),
            CODE_BG, t["s14_quote"], size=13, text_color=ACCENT_BLUE, bold=False)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 15: APIS.GURU + PETSTORE (real from Ch.7)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s15_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_ORANGE)

    # Petstore result — real from Ch.7 sec 7.4.2
    add_box(s, Inches(0.8), Inches(1.4), Inches(11.5), Inches(1.3),
            ACCENT_RED,
            t["s15_petstore"],
            size=14)

    # Real per-principle failures for Petstore
    add_text(s, Inches(0.8), Inches(3.0), Inches(11), Inches(0.5),
             t["s15_failures_title"], size=16, color=WHITE, bold=True)
    failures = t["s15_failures"]
    y = Inches(3.5)
    for pid, reason in failures:
        add_box(s, Inches(0.8), y, Inches(0.8), Inches(0.4),
                ACCENT_RED, pid, size=11)
        add_text(s, Inches(1.8), y, Inches(10.5), Inches(0.4),
                 reason, size=12, color=LIGHT_GRAY)
        y += Inches(0.45)

    # Finding — real from thesis
    add_box(s, Inches(0.8), Inches(5.9), Inches(11.5), Inches(0.8),
            CODE_BG, t["s15_finding"], size=14, text_color=ACCENT_BLUE, bold=False)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 16: CI/CD METRICS (real from Ch.7 Table 7.11)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s16_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_GREEN)

    # Real metrics from Table 7.11
    metrics = [
        (t["s16_m1"], "~15 MB", t["s16_m1n"]),
        (t["s16_m2"], "~25 MB", t["s16_m2n"]),
        (t["s16_m3"], "<1 s", t["s16_m3n"]),
        (t["s16_m4"], "2-5 s", t["s16_m4n"]),
        (t["s16_m5"], "~15 s", t["s16_m5n"]),
        (t["s16_m6"], "<90 s", t["s16_m6n"]),
        (t["s16_m7"], "~44", t["s16_m7n"]),
        (t["s16_m8"], "<50 KB", t["s16_m8n"]),
    ]
    y = Inches(1.5)
    for label, value, note in metrics:
        add_box(s, Inches(0.8), y, Inches(3.5), Inches(0.5),
                TABLE_ALT, label, size=12, text_color=DARK_TEXT, bold=False,
                align=PP_ALIGN.LEFT)
        add_box(s, Inches(4.5), y, Inches(1.8), Inches(0.5),
                ACCENT_GREEN, value, size=14)
        add_box(s, Inches(6.5), y, Inches(6), Inches(0.5),
                TABLE_ALT, note, size=11, text_color=MID_GRAY, bold=False,
                align=PP_ALIGN.LEFT)
        y += Inches(0.58)

    # Operational edge cases summary
    add_text(s, Inches(0.8), Inches(6.2), Inches(11), Inches(0.5),
             t["s16_edge"], size=13, color=MID_GRAY)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 17: AI-ASSISTED DEVELOPMENT (real from Ch.8)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s17_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    # Real stats from Ch.8
    stats = [
        ("57", t["s17_stat1"]),
        ("41", t["s17_stat2"]),
        ("17 (41%)", t["s17_stat3"]),
        ("16", t["s17_stat4"]),
    ]
    for i, (num, label) in enumerate(stats):
        x = Inches(0.8) + (i % 2) * Inches(6.2)
        y = Inches(1.5) + (i // 2) * Inches(1.1)
        add_box(s, x, y, Inches(1.5), Inches(0.8), ACCENT_BLUE, num, size=22)
        add_text(s, x + Inches(1.7), y + Inches(0.1), Inches(4), Inches(0.6),
                 label, size=14, color=LIGHT_GRAY)

    # CLAUDE.md hierarchy — real from Ch.8
    add_text(s, Inches(0.8), Inches(3.8), Inches(5.5), Inches(0.5),
             "CLAUDE.md Knowledge Protocol", size=16, color=ACCENT_BLUE, bold=True)
    add_code_block(s, Inches(0.8), Inches(4.4), Inches(5.5), Inches(2.5),
        "Root CLAUDE.md (project vision, axioms)\n"
        "  +-- driveby-cli/CLAUDE.md (packages)\n"
        "  |   +-- principles/CLAUDE.md (interface)\n"
        "  |   +-- spec/CLAUDE.md (adapters)\n"
        "  +-- kubernetes/CLAUDE.md (XRD, Helm)\n"
        "  +-- apis/perfect-api/CLAUDE.md\n"
        "  +-- thesis/CLAUDE.md (chapter specs)\n"
        "  +-- tools/, results/, docs/, .github/",
        size=11)

    # Agent loop — real from Ch.8
    add_text(s, Inches(7), Inches(3.8), Inches(5.5), Inches(0.5),
             t["s17_loop_title"], size=16, color=ACCENT_GREEN, bold=True)
    add_multiline(s, Inches(7), Inches(4.4), Inches(5.5), Inches(2.5),
                  t["s17_loop"], size=13, color=LIGHT_GRAY, spacing=Pt(10))

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 18: RESEARCH QUESTIONS ANSWERED (real from Ch.10)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s18_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_GREEN)

    rqs = t["s18_rqs"]
    for i, (q, a, evidence) in enumerate(rqs):
        y = Inches(1.4) + i * Inches(1.4)
        add_box(s, Inches(0.8), y, Inches(4.5), Inches(1.1),
                ACCENT_BLUE, q, size=11)
        add_box(s, Inches(5.5), y, Inches(1.5), Inches(1.1),
                ACCENT_GREEN, a, size=16)
        add_box(s, Inches(7.2), y, Inches(5.3), Inches(1.1),
                TABLE_ALT, evidence, size=10, text_color=DARK_TEXT, bold=False,
                align=PP_ALIGN.LEFT)

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 19: CONTRIBUTIONS (real from Ch.10)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s19_title"], size=32, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5))

    contribs = t["s19_contribs"]
    c_colors = [ACCENT_BLUE, ACCENT_GREEN, ACCENT_ORANGE, ACCENT_BLUE]
    for i, (title, details) in enumerate(contribs):
        x = Inches(0.8) + (i % 2) * Inches(6.2)
        y = Inches(1.5) + (i // 2) * Inches(2.6)
        add_box(s, x, y, Inches(5.5), Inches(0.7),
                c_colors[i], title, size=16)
        add_multiline(s, x + Inches(0.2), y + Inches(0.8), Inches(5.2), Inches(1.8),
                      details, size=12, color=LIGHT_GRAY, spacing=Pt(5))

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 20: LIMITATIONS & FUTURE (real from Ch.10)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, LIGHT_BG)
    add_text(s, Inches(0.8), Inches(0.4), Inches(11), Inches(0.8),
             t["s20_title"], size=32, color=DARK_TEXT, bold=True)
    add_line(s, Inches(0.8), Inches(1.1), Inches(2.5), ACCENT_ORANGE)

    add_text(s, Inches(0.8), Inches(1.4), Inches(5.5), Inches(0.5),
             t["s20_lim_title"], size=18, color=ACCENT_RED, bold=True)
    add_multiline(s, Inches(0.8), Inches(1.9), Inches(5.5), Inches(4),
                  t["s20_lims"], size=13, color=DARK_TEXT, spacing=Pt(8))

    add_text(s, Inches(7), Inches(1.4), Inches(5.5), Inches(0.5),
             t["s20_fut_title"], size=18, color=ACCENT_GREEN, bold=True)
    add_multiline(s, Inches(7), Inches(1.9), Inches(5.5), Inches(4),
                  t["s20_future"], size=13, color=DARK_TEXT, spacing=Pt(8))

    # ═══════════════════════════════════════════════════════════════════
    # SLIDE 21: CLOSING (real thesis quotes)
    # ═══════════════════════════════════════════════════════════════════
    s = prs.slides.add_slide(blank)
    set_slide_bg(s, DARK_BG)

    add_text(s, Inches(1), Inches(1.2), Inches(11), Inches(1.2),
             t["s21_quote1"], size=26, bold=True, align=PP_ALIGN.CENTER)
    add_line(s, Inches(5), Inches(2.6), Inches(3.3))
    add_text(s, Inches(1), Inches(3.0), Inches(11), Inches(1),
             t["s21_quote2"], size=16, color=ACCENT_BLUE, align=PP_ALIGN.CENTER)
    add_text(s, Inches(1), Inches(4.5), Inches(11), Inches(1),
             t["s21_quote3"], size=16, color=LIGHT_GRAY, align=PP_ALIGN.CENTER)

    add_text(s, Inches(1), Inches(5.8), Inches(11), Inches(0.5),
             t["s21_thanks"], size=24, bold=True, align=PP_ALIGN.CENTER)
    add_text(s, Inches(1), Inches(6.4), Inches(11), Inches(0.5),
             t["s21_contact"], size=14, color=LIGHT_GRAY, align=PP_ALIGN.CENTER)

    return prs


# ═══════════════════════════════════════════════════════════════════════
# TRANSLATIONS — All content sourced from thesis chapters
# ═══════════════════════════════════════════════════════════════════════

TR = {}

TR["en"] = {
    # S1
    "subtitle": "A Paradigm for Automated API Quality Assurance\nin the GitOps Era",
    "author": "Petros Evangelos Triantafyllis  |  Supervisor: Kyriakos Kritikos",
    "institution": "University of the Aegean -- Dept. of Information and Communication Systems Engineering",

    # S2 — real thesis quotes from Ch.1
    "s2_title": "The Problem",
    "s2_quote": (
        '"Documentation drift is not a symptom of laziness. It is a systemic failure '
        'rooted in the separation of two artifacts -- the specification and the '
        'implementation -- that should be governed by a single source of truth."  '
        '-- Chapter 1'
    ),
    "s2_problems": [
        ("Form, Not Substance",
         "Existing tools (Spectral, Vacuum, Redocly) "
         "enforce structural rules but cannot verify "
         "described behavior matches implementation. "
         "No tool combines static + runtime validation."),
        ("Tribal Knowledge",
         "Configuration lives in scripts, CI pipelines, "
         "or developer knowledge. Only structured, "
         "machine-readable knowledge can be automated "
         "or consumed by AI agents."),
        ("Hours of ClickOps",
         "Setting up quality gates requires navigating "
         "ArgoCD UI, GitHub UI, and kubectl manually. "
         "Non-reproducible, non-versionable, "
         "error-prone process."),
    ],

    # S3
    "s3_question": "What if the documentation itself\ncould BE the testing infrastructure?",
    "s3_answer": (
        "The specification is simultaneously the documentation,\n"
        "the contract, and the test oracle.\n"
        "-- Chapter 4, CLI Architecture"
    ),

    # S4 — real axiom definitions from Ch.3
    "s4_title": "DDT Methodology -- Three Axioms",
    "s4_axioms": [
        ("Completeness",
         '"The API specification must fully describe the API contract"',
         "Operationalized by:\nP001 OpenAPI Compliance\nP002 Documentation Quality\n"
         "P003 Error Handling\nP004 Schema Definitions"),
        ("Determinism",
         '"Same specification + same API = same validation results"',
         "Operationalized by:\nP006 Functional Testing\nP009 Test Readiness"),
        ("Observability",
         '"Validation produces measurable, actionable results"',
         "Operationalized by:\nP005 Security Standards\nP007 Performance Testing\n"
         "P008 Versioning Strategy"),
    ],
    "s4_bottom": "3 Axioms  -->  9 Principles (P001-P009)  -->  4 Validation Modes (minimal, test-ready, strict, test-only)",

    # S5 — real code from Ch.4
    "s5_title": "The APISpec Abstraction Layer (Ch. 4)",
    "s5_right_title": "Why This Matters",
    "s5_points": [
        "Adapter pattern (GoF) -- version-agnostic validation",
        "OpenAPI 3.0.x, 3.1.0, Swagger 2.0 -- same checkers",
        "No raw *openapi3.T in any validator -- enforced by interface",
        "Preprocessing: nullable anyOf rewrite, exclusive min/max",
        "Recursive schema normalization (allOf, oneOf, anyOf)",
        "FlattenAllOf utility for composed schemas",
        "Adding AsyncAPI/gRPC = implement interface, no checker changes",
    ],

    # S6 — real code + table from Ch.4
    "s6_title": "The PrincipleChecker Pattern (Ch. 4)",
    "s6_table_title": "Checker Implementations",
    "s6_total": "Total: 7 checkers, 52 checks, 1,615 LOC\nStrategy pattern (GoF): each checker interchangeable",

    # S7
    "s7_title": "The Validation Engine (Ch. 4)",
    "s7_dep_title": "Dependency Flow",
    "s7_note": "Deterministic by construction: same spec + mode =\nsame checkers in same order = same results.\nOnly non-deterministic element: timestamp.",

    # S8 — real from Ch.5
    "s8_title": "From CLI to Custom Resource (Ch. 5)",
    "s8_stages": [
        ("Shell Script", "Ad-hoc string matches, no structure, no reuse, no determinism"),
        ("Go CLI", "Typed principles, APISpec abstraction, deterministic engine, JSON output"),
        ("Container Image", "Same binary in any container runtime (Docker, containerd, CRI-O)"),
        ("Argo Workflow Step", "Event-driven execution, DAG ordering, artifact persistence, exit handling"),
        ("Crossplane CRD", "~25 YAML lines --> ~44 resources in <15s. Reconciled. Self-healing."),
    ],
    "s8_quote": '"Each stage subsumes the previous. The CR still runs the same CLI, inside the same container, as a step in the same DAG." -- Ch. 5',

    # S9
    "s9_title": "XSDLC: 25 Lines of YAML (Ch. 5)",
    "s9_gen_title": "Generated Resources (~44 total)",
    "s9_resources": [
        ("App-level (once per CR)", [
            "GitOps Repository (auto-created, NO software repo ref)",
            "RepositoryFiles (dry/base/ boilerplate manifests)",
            "ServiceAccount + EventBus (NATS JetStream)",
            "ArgoCD push secret + ScmProvider + GitRepository",
            "BranchProtection (*-next branches)",
        ]),
        ("Per-environment", [
            "PromotionStrategy + 2 Branches + Kustomize overlay",
            "ArgoCD Application (sourceHydrator per overlay)",
            "ghcr-creds secret (image pull auth)",
        ]),
        ("Per-gate", [
            "Role + RoleBinding (RBAC)",
            "WorkflowTemplate (validation DAG)",
            "EventSource + Sensor + Service + Ingress",
            "BranchProtection (require gate status check)",
        ]),
    ],

    # S10
    "s10_title": "The Reconciliation Loop (Ch. 5)",
    "s10_edge": "Edge-triggered (Events)",
    "s10_edge_desc": "Reacts to state changes as they happen.\nIf it misses one (crash, restart, partition)\nthe missed event is lost forever.\nMust record every event.",
    "s10_level": "Level-triggered (State)",
    "s10_level_desc": "Periodically re-evaluates current condition.\nIf controller crashes and restarts, it re-reads\ncurrent state and reconciles from scratch.\nNo events need to be recorded.",
    "s10_lifecycle_title": "Declarative Lifecycle",
    "s10_days": [
        ("Day 0: Install", "Helm chart installs XRD,\ncomposition, ProviderConfigs.\nNo XSDLC instances yet."),
        ("Day 1: Create", "Apply XSDLC CR.\n~44 resources generated\nin <15 seconds."),
        ("Day 2: Reconcile", "Drift detection: overwrites\nmanual changes. Self-healing:\nre-creates deleted resources."),
        ("Day N: Update/Delete", "Modify spec --> delta reconcile.\nDelete CR --> cascade delete\n(orphan shared resources)."),
    ],

    # S11
    "s11_title": "BYOCI Delivery Pipeline (Ch. 5/6)",
    "s11_byoci": "Bring Your Own CI: XSDLC owns ONLY delivery (promotion, gates, environments). Developer owns everything else (build, test, image, CI).",
    "s11_gitops_repo": "GitOps Repo (auto-created by XSDLC)",
    "s11_gitops_desc": [
        "XSDLC auto-creates a dedicated gitops repo per app",
        "dry/base/ + dry/overlays/<env>/ (Kustomize)",
        "ArgoCD Source Hydrator renders per-env overlays",
        "Promoter creates merge PRs: staging-next -> staging",
        "Developer commits manifests to main -- that's the contract",
    ],
    "s11_dev_owns": "Developer's Responsibility (NOT XSDLC)",
    "s11_dev_desc": [
        "Source code, application logic, unit tests",
        "CI pipeline: build, test, publish image",
        "Any CI system: GitHub Actions, GitLab, Jenkins, bash",
        "XSDLC has ZERO visibility into developer's CI",
        "No software repo reference needed in XSDLC CR",
    ],
    "s11_flow": [
        "Dev commits\nmanifests\nto main",
        "Hydrator\nrenders\noverlays",
        "Promoter\nopens PR\n(env-next)",
        "Gate fires\nDriveBy\nvalidation",
        "Auto-merge\n& ArgoCD\nsync",
    ],

    # S12
    "s12_title": "Controlled Evaluation (Ch. 7)",
    "s12_inject_title": "Defect Injection Matrix",
    "s12_d1": "Duplicate operationId",
    "s12_d2": "Remove all descriptions",
    "s12_d3": "Remove all 4xx/5xx responses",
    "s12_d4": "Remove schema type specifications",
    "s12_d5": "Remove securitySchemes entirely",
    "s12_d6": 'Version = "latest" (non-semver)',
    "s12_d7": "Remove all examples + types",
    "s12_baseline_title": "Baseline: perfect-api",
    "s12_result": "100% detection, 0 false positives\n0 false negatives, deterministic",
    "s12_severity_title": "Severity Model",
    "s12_severity": [
        "Critical (2/9): P001 OpenAPI Compliance, P005 Security Standards --> exit code 1, blocks promotion",
        "Warning (7/9): P002-P004, P008, P009 --> exit code 0, reported but non-blocking",
        "Asymmetry by design: slow-api passed staging (3/6, all warnings), caught at production by load test",
    ],

    # S13
    "s13_title": "Multi-API Pipeline Results (Ch. 7)",
    "s13_defect": "Defect Class",
    "s13_apis": [
        ("perfect-api", "None", "PASS (P001, P006)", "PASS (P001, P007)", ACCENT_GREEN),
        ("bad-docs-api", "No security", "FAIL (P005 critical)", "-- blocked --", ACCENT_RED),
        ("no-auth-api", "No auth", "FAIL (P005 critical)", "-- blocked --", ACCENT_RED),
        ("slow-api", "500ms delay", "PASS (3/6 warnings)", "FAIL (P95: 502ms > 200ms)", ACCENT_ORANGE),
        ("broken-api", "Spec mismatch", "FAIL (P006: 418!=200)", "-- blocked --", ACCENT_RED),
    ],
    "s13_bottom": "Only 1 of 5 APIs (20%) completed the full staging-to-production pipeline",

    # S14
    "s14_title": "Three-Layer Defence Model (Ch. 7)",
    "s14_layers": [
        ("Layer 1", "Static Validation\n(P001-P005, P008)", "Specification\ndefects", "bad-docs-api\nno-auth-api"),
        ("Layer 2", "Functional Testing\n(P006)", "Implementation\ndefects", "broken-api"),
        ("Layer 3", "Load Testing\n(P007)", "Performance\ndefects", "slow-api"),
    ],
    "s14_quote": '"Each layer was necessary. Static validation alone would have passed slow-api and broken-api. Only the combination of all three layers provided the discrimination required." -- Ch. 7',

    # S15
    "s15_title": "Large-Scale: APIs.guru Dataset (Ch. 7)",
    "s15_petstore": "Swagger Petstore v3 (canonical OpenAPI example, maintained by OpenAPI Initiative)\nStrict mode: 1/6 principles pass -- only P001 (structural compliance)",
    "s15_failures_title": "Why Petstore Fails (5 principles):",
    "s15_failures": [
        ("P002", "Missing operation descriptions, parameter descriptions, response examples"),
        ("P003", "No 5xx error responses defined; error schemas lack structured fields"),
        ("P004", "String parameters lack length constraints; numeric fields lack range constraints"),
        ("P005", "OAuth2 scopes defined but not referenced in security requirements"),
        ("P008", "Version present + semver, but no versioning strategy / deprecation / migration docs"),
    ],
    "s15_finding": '"Even the canonical example specification maintained by the OpenAPI Initiative fails DDT\'s strict validation. Existing tooling validates form, not substance." -- Ch. 7',

    # S16
    "s16_title": "CI/CD Integration Metrics (Ch. 7)",
    "s16_m1": "CLI binary size", "s16_m1n": "Cross-compiled Go binary",
    "s16_m2": "Docker image size", "s16_m2n": "Distroless base image",
    "s16_m3": "Startup time", "s16_m3n": "No JVM, no runtime dependencies",
    "s16_m4": "Validation time (strict)", "s16_m4n": "Depends on specification size",
    "s16_m5": "XSDLC provisioning", "s16_m5n": "From CR apply to all resources ready",
    "s16_m6": "End-to-end gate latency", "s16_m6n": "Webhook to commit status posted",
    "s16_m7": "Resources per XSDLC CR", "s16_m7n": "For a 3-environment pipeline",
    "s16_m8": "Helm chart size", "s16_m8n": "OCI-distributed",
    "s16_edge": "Operational edge cases: 6/7 passed. Sole failure: ArgoCD selfHeal restored deployment faster than test could trigger.",

    # S17
    "s17_title": "AI-Assisted Development (Ch. 8)",
    "s17_stat1": "Claude Code sessions",
    "s17_stat2": "total commits in repository",
    "s17_stat3": "commits Co-Authored by AI",
    "s17_stat4": "CLAUDE.md files (knowledge protocol)",
    "s17_loop_title": "DDT as Agent Feedback Loop",
    "s17_loop": [
        "1. validate  --> parse JSON report",
        "2. diagnose  --> identify failing principle + check",
        "3. remediate --> generate spec fix (e.g. add maxLength)",
        "4. re-validate --> confirm fix (deterministic)",
        "",
        '"Agent effectiveness scales with quality',
        'of knowledge infrastructure, not model',
        'capability alone." -- Ch. 8',
    ],

    # S18
    "s18_title": "Research Questions -- Answered (Ch. 10)",
    "s18_rqs": [
        ("RQ1: Can DDT provide\nautomated QA comparable\nto manual approaches?",
         "YES",
         "100% detection, 0 false positives.\nLayered discrimination: each layer\ncatches distinct defect class."),
        ("RQ2: Can validation rules\nbe derived automatically\nfrom specifications?",
         "FULLY",
         "7/9 principle checkers derive all\nvalidation from spec. Zero manual\ntest authorship confirmed at scale."),
        ("RQ3: How effectively does\nDDT integrate into\nGitOps pipelines?",
         "EFFECTIVELY",
         "1 CR = ~44 resources in <15s.\nSurgical coexistence on production\ncluster. 5 APIs simultaneously."),
        ("RQ4: Can quality gates\nbe expressed as ontological\nspecifications?",
         "YES",
         "XSDLC = discoverable, inspectable,\ncomposable Kubernetes-native\nontological specification."),
    ],

    # S19
    "s19_title": "Four Contributions (Ch. 10)",
    "s19_contribs": [
        ("1. DDT Methodology", [
            "3 axioms, 9 principles, 4 validation modes",
            "Zero manual test authorship",
            "Stable axiomatic framework, evolvable principles",
            "Severity model: 2 critical, 7 warning",
        ]),
        ("2. DriveBy Framework", [
            "Go CLI: 11 packages, 7,170 LOC",
            "APISpec abstraction (Adapter pattern)",
            "7 principle checkers (Strategy pattern)",
            "JSON + Markdown reports, 4-code exit protocol",
        ]),
        ("3. XSDLC Pipeline", [
            "Crossplane CRD: ~25 YAML --> ~44 resources",
            "Dual-provider: provider-kubernetes + provider-upjet-github",
            "BYOCI: developer owns CI, XSDLC owns delivery",
            "Level-triggered reconciliation, self-healing",
        ]),
        ("4. Agent-Driven Development", [
            "CLAUDE.md knowledge protocol (16 files)",
            "57 sessions, CI-enforced documentation",
            "Bidirectional code <--> thesis traceability",
            "DDT output as agent feedback infrastructure",
        ]),
    ],

    # S20
    "s20_title": "Limitations & Future Work (Ch. 10)",
    "s20_lim_title": "Limitations",
    "s20_lims": [
        "P006, P007: logic exists but no PrincipleChecker wrapper",
        "P008: keyword-based matching in info.description (brittle)",
        "OpenAPI 3.x + Swagger 2.0 only (no GraphQL, gRPC, AsyncAPI)",
        "Static load test thresholds (no adaptive adjustment)",
        "Evaluation: single IDP (Novelcore), public APIs may not represent enterprise",
        "Binary pass/fail scoring (no percentage-based granularity)",
    ],
    "s20_fut_title": "Future Work",
    "s20_future": [
        "Adaptive load thresholds: controller watches Prometheus metrics",
        "Agent-driven remediation: AI parses DDT reports, fixes specs",
        "P006/P007 as PrincipleCheckers in unified pipeline",
        "Multi-format: GraphQL, gRPC, AsyncAPI adapters",
        "Enterprise multi-cluster delivery patterns",
        "XSDLC as child resource of microservice provisioning operator",
    ],

    # S21
    "s21_quote1": "You stop telling the system what to do.\nYou start telling it how the world should BE.",
    "s21_quote2": (
        "From imperative scripts to ontological specifications.\n"
        "The system reconciles toward the declared state continuously,\n"
        "without human initiation, without human interpretation.\n"
        "-- Chapter 10"
    ),
    "s21_quote3": (
        "Quality assurance becomes a platform capability\n"
        "rather than a per-team responsibility."
    ),
    "s21_thanks": "Thank You",
    "s21_contact": "Petros Evangelos Triantafyllis  --  University of the Aegean  --  github.com/meter-peter/driveby",
}

# ── Greek translation ─────────────────────────────────────────────────
TR["gr"] = {
    # S1
    "subtitle": "Ενα Παραδειγμα για Αυτοματοποιημενη Διασφαλιση\nΠοιοτητας API στην Εποχη του GitOps",
    "author": "Πετρος Ευαγγελος Τριανταφυλλης  |  Επιβλεπων: Κυριακος Κριτικος",
    "institution": "Πανεπιστημιο Αιγαιου -- Τμημα Μηχανικων Πληροφοριακων & Επικοινωνιακων Συστηματων",

    # S2
    "s2_title": "Το Προβλημα",
    "s2_quote": (
        '"Το documentation drift δεν ειναι συμπτωμα τεμπελιας. Ειναι συστημικη αποτυχια '
        'που εδραζεται στον διαχωρισμο δυο artifacts -- του specification και της '
        'υλοποιησης -- που θα επρεπε να διεπονται απο μια μοναδικη πηγη αληθειας."  '
        '-- Κεφαλαιο 1'
    ),
    "s2_problems": [
        ("Μορφη, Οχι Ουσια",
         "Τα υπαρχοντα εργαλεια (Spectral, Vacuum) "
         "ελεγχουν δομικους κανονες αλλα δεν "
         "επαληθευουν αν η περιγραφομενη συμπεριφορα "
         "αντιστοιχει στην υλοποιηση."),
        ("Tribal Knowledge",
         "Η γνωση ζει σε scripts, CI pipelines "
         "η στη μνημη των developers. Μονο "
         "δομημενη, machine-readable γνωση μπορει "
         "να αυτοματοποιηθει η να καταναλωθει απο AI."),
        ("Ωρες ClickOps",
         "Η δημιουργια quality gates απαιτει πλοηγηση "
         "στο ArgoCD UI, GitHub UI, και kubectl. "
         "Μη αναπαραγωγιμο, μη versionable, "
         "επιρρεπες σε σφαλματα."),
    ],

    # S3
    "s3_question": "Τι θα γινοταν αν η τεκμηριωση\nμπορουσε να ΕΙΝΑΙ η υποδομη testing;",
    "s3_answer": (
        "Το specification ειναι ταυτοχρονα η τεκμηριωση,\n"
        "το contract, και το test oracle.\n"
        "-- Κεφαλαιο 4, Αρχιτεκτονικη CLI"
    ),

    # S4
    "s4_title": "Μεθοδολογια DDT -- Τρια Αξιωματα",
    "s4_axioms": [
        ("Πληροτητα (Completeness)",
         '"Το API specification πρεπει να περιγραφει πληρως το API contract"',
         "Λειτουργικοποιειται απο:\nP001 OpenAPI Compliance\nP002 Documentation Quality\n"
         "P003 Error Handling\nP004 Schema Definitions"),
        ("Ντετερμινισμος (Determinism)",
         '"Ιδιο specification + ιδιο API = ιδια αποτελεσματα"',
         "Λειτουργικοποιειται απο:\nP006 Functional Testing\nP009 Test Readiness"),
        ("Παρατηρησιμοτητα (Observability)",
         '"Το validation παραγει μετρησιμα, ενεργοποιησιμα αποτελεσματα"',
         "Λειτουργικοποιειται απο:\nP005 Security Standards\nP007 Performance Testing\n"
         "P008 Versioning Strategy"),
    ],
    "s4_bottom": "3 Αξιωματα  -->  9 Αρχες (P001-P009)  -->  4 Modes (minimal, test-ready, strict, test-only)",

    # S5
    "s5_title": "Το APISpec Abstraction Layer (Κεφ. 4)",
    "s5_right_title": "Γιατι Εχει Σημασια",
    "s5_points": [
        "Adapter pattern (GoF) -- version-agnostic validation",
        "OpenAPI 3.0.x, 3.1.0, Swagger 2.0 -- ιδιοι checkers",
        "Κανενα raw *openapi3.T σε validators -- enforced by interface",
        "Preprocessing: nullable anyOf rewrite, exclusive min/max",
        "Αναδρομικη κανονικοποιηση schema (allOf, oneOf, anyOf)",
        "FlattenAllOf utility για composed schemas",
        "Προσθηκη AsyncAPI/gRPC = implement interface, χωρις αλλαγες",
    ],

    # S6
    "s6_title": "Το PrincipleChecker Pattern (Κεφ. 4)",
    "s6_table_title": "Υλοποιησεις Checkers",
    "s6_total": "Συνολο: 7 checkers, 52 checks, 1,615 LOC\nStrategy pattern (GoF): καθε checker εναλλασσιμο",

    # S7
    "s7_title": "Η Validation Engine (Κεφ. 4)",
    "s7_dep_title": "Dependency Flow",
    "s7_note": "Ντετερμινιστικη κατασκευαστικα: ιδιο spec + mode =\nιδιοι checkers με ιδια σειρα = ιδια αποτελεσματα.\nΜοναδικο μη-ντετερμινιστικο στοιχειο: timestamp.",

    # S8
    "s8_title": "Απο CLI σε Custom Resource (Κεφ. 5)",
    "s8_stages": [
        ("Shell Script", "Ad-hoc string matches, χωρις δομη, χωρις επαναχρηση, χωρις ντετερμινισμο"),
        ("Go CLI", "Typed principles, APISpec abstraction, deterministic engine, JSON output"),
        ("Container Image", "Ιδιο binary σε οποιοδηποτε container runtime (Docker, containerd, CRI-O)"),
        ("Argo Workflow Step", "Event-driven εκτελεση, DAG ordering, artifact persistence, exit handling"),
        ("Crossplane CRD", "~25 YAML lines --> ~44 resources σε <15s. Reconciled. Self-healing."),
    ],
    "s8_quote": '"Καθε σταδιο ενσωματωνει το προηγουμενο. Το CR τρεχει ακομα το ιδιο CLI, στο ιδιο container, ως step στο ιδιο DAG." -- Κεφ. 5',

    # S9
    "s9_title": "XSDLC: 25 Γραμμες YAML (Κεφ. 5)",
    "s9_gen_title": "Παραγομενοι Ποροι (~44 συνολο)",
    "s9_resources": [
        ("Σε επιπεδο εφαρμογης (1 φορα ανα CR)", [
            "GitOps Repository (auto-created, ΧΩΡΙΣ software repo ref)",
            "RepositoryFiles (dry/base/ boilerplate manifests)",
            "ServiceAccount + EventBus (NATS JetStream)",
            "ArgoCD push secret + ScmProvider + GitRepository",
            "BranchProtection (*-next branches)",
        ]),
        ("Ανα περιβαλλον", [
            "PromotionStrategy + 2 Branches + Kustomize overlay",
            "ArgoCD Application (sourceHydrator ανα overlay)",
            "ghcr-creds secret (image pull auth)",
        ]),
        ("Ανα gate", [
            "Role + RoleBinding (RBAC)",
            "WorkflowTemplate (validation DAG)",
            "EventSource + Sensor + Service + Ingress",
            "BranchProtection (require gate status check)",
        ]),
    ],

    # S10
    "s10_title": "Ο Reconciliation Loop (Κεφ. 5)",
    "s10_edge": "Edge-triggered (Events)",
    "s10_edge_desc": "Αντιδρα σε αλλαγες καταστασης καθως συμβαινουν.\nΑν χασει ενα event (crash, restart, partition)\nτο χαμενο event χανεται για παντα.\nΠρεπει να καταγραφει καθε event.",
    "s10_level": "Level-triggered (State)",
    "s10_level_desc": "Επαναξιολογει περιοδικα την τρεχουσα κατασταση.\nΑν ο controller κρασαρει και ξαναρχισει, διαβαζει\nτην τρεχουσα κατασταση και κανει reconcile απο την αρχη.\nΚανενα event δεν χρειαζεται καταγραφη.",
    "s10_lifecycle_title": "Δηλωτικος Κυκλος Ζωης",
    "s10_days": [
        ("Day 0: Εγκατασταση", "Helm chart εγκαθιστα XRD,\ncomposition, ProviderConfigs.\nΚανενα XSDLC instance ακομα."),
        ("Day 1: Δημιουργια", "Apply XSDLC CR.\n~44 resources δημιουργουνται\nσε <15 δευτερολεπτα."),
        ("Day 2: Reconcile", "Drift detection: αντικαθιστα\nχειροκινητες αλλαγες. Self-healing:\nξαναδημιουργει σβησμενους πορους."),
        ("Day N: Update/Delete", "Τροποποιηση spec --> delta reconcile.\nΔιαγραφη CR --> cascade delete\n(orphan shared resources)."),
    ],

    # S11
    "s11_title": "BYOCI Delivery Pipeline (Κεφ. 5/6)",
    "s11_byoci": "Bring Your Own CI: Ο XSDLC κατεχει ΜΟΝΟ το delivery (promotion, gates, environments). Ο developer κατεχει ολα τα αλλα (build, test, image, CI).",
    "s11_gitops_repo": "GitOps Repo (auto-created by XSDLC)",
    "s11_gitops_desc": [
        "Ο XSDLC δημιουργει αυτοματα ενα gitops repo ανα app",
        "dry/base/ + dry/overlays/<env>/ (Kustomize)",
        "ArgoCD Source Hydrator renders ανα env overlay",
        "Promoter δημιουργει merge PRs: staging-next -> staging",
        "Ο developer κανει commit manifests στο main -- αυτο ειναι το contract",
    ],
    "s11_dev_owns": "Ευθυνη Developer (ΟΧΙ XSDLC)",
    "s11_dev_desc": [
        "Source code, application logic, unit tests",
        "CI pipeline: build, test, publish image",
        "Οποιοδηποτε CI: GitHub Actions, GitLab, Jenkins, bash",
        "Ο XSDLC εχει ΜΗΔΕΝ ορατοτητα στο CI του developer",
        "Κανενα software repo reference στο XSDLC CR",
    ],
    "s11_flow": [
        "Dev κανει\ncommit\nmanifests",
        "Hydrator\nrenders\noverlays",
        "Promoter\nανοιγει PR\n(env-next)",
        "Gate fires\nDriveBy\nvalidation",
        "Auto-merge\n& ArgoCD\nsync",
    ],

    # S12
    "s12_title": "Ελεγχομενη Αξιολογηση (Κεφ. 7)",
    "s12_inject_title": "Πινακας Εγχυσης Ελαττωματων",
    "s12_d1": "Duplicate operationId",
    "s12_d2": "Αφαιρεση ολων των descriptions",
    "s12_d3": "Αφαιρεση ολων των 4xx/5xx responses",
    "s12_d4": "Αφαιρεση schema type specifications",
    "s12_d5": "Αφαιρεση securitySchemes ολοκληρωτικα",
    "s12_d6": 'Version = "latest" (μη-semver)',
    "s12_d7": "Αφαιρεση ολων examples + types",
    "s12_baseline_title": "Baseline: perfect-api",
    "s12_result": "100% ανιχνευση, 0 false positives\n0 false negatives, ντετερμινιστικη",
    "s12_severity_title": "Μοντελο Σοβαροτητας",
    "s12_severity": [
        "Critical (2/9): P001 OpenAPI Compliance, P005 Security Standards --> exit code 1, μπλοκαρει promotion",
        "Warning (7/9): P002-P004, P008, P009 --> exit code 0, αναφερεται αλλα δεν μπλοκαρει",
        "Ασυμμετρια σκοπιμη: slow-api περασε staging (3/6, ολα warnings), πιαστηκε στο production απο load test",
    ],

    # S13
    "s13_title": "Αποτελεσματα Multi-API Pipeline (Κεφ. 7)",
    "s13_defect": "Ελαττωμα",
    "s13_apis": [
        ("perfect-api", "Κανενα", "PASS (P001, P006)", "PASS (P001, P007)", ACCENT_GREEN),
        ("bad-docs-api", "Χωρις security", "FAIL (P005 critical)", "-- blocked --", ACCENT_RED),
        ("no-auth-api", "Χωρις auth", "FAIL (P005 critical)", "-- blocked --", ACCENT_RED),
        ("slow-api", "500ms delay", "PASS (3/6 warnings)", "FAIL (P95: 502ms > 200ms)", ACCENT_ORANGE),
        ("broken-api", "Spec mismatch", "FAIL (P006: 418!=200)", "-- blocked --", ACCENT_RED),
    ],
    "s13_bottom": "Μονο 1 απο 5 APIs (20%) ολοκληρωσε πληρως το pipeline staging-to-production",

    # S14
    "s14_title": "Μοντελο Πολυεπιπεδης Αμυνας (Κεφ. 7)",
    "s14_layers": [
        ("Layer 1", "Στατικη Αναλυση\n(P001-P005, P008)", "Ελαττωματα\nspecification", "bad-docs-api\nno-auth-api"),
        ("Layer 2", "Functional Testing\n(P006)", "Ελαττωματα\nυλοποιησης", "broken-api"),
        ("Layer 3", "Load Testing\n(P007)", "Ελαττωματα\nαποδοσης", "slow-api"),
    ],
    "s14_quote": '"Καθε επιπεδο ηταν αναγκαιο. Η στατικη αναλυση μονη της θα αφηνε να περασουν slow-api και broken-api. Μονο ο συνδυασμος και των τριων εδωσε τη διακριση." -- Κεφ. 7',

    # S15
    "s15_title": "Large-Scale: APIs.guru Dataset (Κεφ. 7)",
    "s15_petstore": "Swagger Petstore v3 (canonical OpenAPI παραδειγμα, συντηρειται απο το OpenAPI Initiative)\nStrict mode: 1/6 αρχες περνανε -- μονο P001 (δομικη συμμορφωση)",
    "s15_failures_title": "Γιατι Αποτυγχανει το Petstore (5 αρχες):",
    "s15_failures": [
        ("P002", "Ελλειπεις περιγραφες operations, parameters, response examples"),
        ("P003", "Κανενα 5xx error response; τα error schemas χωρις δομημενα πεδια"),
        ("P004", "String parameters χωρις length constraints; numeric fields χωρις range"),
        ("P005", "OAuth2 scopes ορισμενα αλλα δεν αναφερονται στα security requirements"),
        ("P008", "Version + semver, αλλα χωρις versioning strategy / deprecation / migration docs"),
    ],
    "s15_finding": '"Ακομα και το canonical specification που συντηρειται απο το OpenAPI Initiative αποτυγχανει στο strict mode. Τα υπαρχοντα εργαλεια ελεγχουν μορφη, οχι ουσια." -- Κεφ. 7',

    # S16
    "s16_title": "Μετρικες CI/CD Ενσωματωσης (Κεφ. 7)",
    "s16_m1": "CLI binary size", "s16_m1n": "Cross-compiled Go binary",
    "s16_m2": "Docker image size", "s16_m2n": "Distroless base image",
    "s16_m3": "Startup time", "s16_m3n": "Χωρις JVM, χωρις runtime dependencies",
    "s16_m4": "Validation time (strict)", "s16_m4n": "Εξαρταται απο μεγεθος specification",
    "s16_m5": "XSDLC provisioning", "s16_m5n": "Απο CR apply εως ολα ready",
    "s16_m6": "End-to-end gate latency", "s16_m6n": "Webhook εως commit status posted",
    "s16_m7": "Resources per XSDLC CR", "s16_m7n": "Για pipeline 3 περιβαλλοντων",
    "s16_m8": "Helm chart size", "s16_m8n": "OCI-distributed",
    "s16_edge": "Edge cases: 6/7 επιτυχια. Μοναδικη αποτυχια: ArgoCD selfHeal επανεφερε deployment πριν τον ελεγχο.",

    # S17
    "s17_title": "Αναπτυξη με AI Agent (Κεφ. 8)",
    "s17_stat1": "Claude Code sessions",
    "s17_stat2": "συνολικα commits στο repository",
    "s17_stat3": "commits Co-Authored by AI",
    "s17_stat4": "CLAUDE.md αρχεια (knowledge protocol)",
    "s17_loop_title": "DDT ως Agent Feedback Loop",
    "s17_loop": [
        "1. validate  --> parse JSON report",
        "2. diagnose  --> εντοπισμος αποτυχημενης αρχης + check",
        "3. remediate --> δημιουργια fix (π.χ. maxLength)",
        "4. re-validate --> επιβεβαιωση fix (ντετερμινιστικη)",
        "",
        '"Η αποτελεσματικοτητα του agent κλιμακωνεται',
        'με την ποιοτητα της υποδομης γνωσης, οχι',
        'με τη δυνατοτητα του μοντελου μονο." -- Κεφ. 8',
    ],

    # S18
    "s18_title": "Ερευνητικα Ερωτηματα -- Απαντησεις (Κεφ. 10)",
    "s18_rqs": [
        ("RQ1: Μπορει η DDT να παρεχει\nαυτοματοποιημενη QA συγκρισιμη\nμε manual proσεγγισεις;",
         "NAI",
         "100% ανιχνευση, 0 false positives.\nΠολυεπιπεδη διακριση: καθε layer\nπιανει ξεχωριστη κατηγορια ελαττωματων."),
        ("RQ2: Σε τι βαθμο μπορουν\nοι κανονες validation να\nπαραχθουν αυτοματα απο specs;",
         "ΠΛΗΡΩΣ",
         "7/9 checkers παραγουν ολο το\nvalidation απο το spec. Zero manual\ntest authorship επιβεβαιωμενο σε κλιμακα."),
        ("RQ3: Ποσο αποτελεσματικα\nενσωματωνεται σε\nGitOps pipelines;",
         "ΑΠΟΤΕΛ/ΚΑ",
         "1 CR = ~44 resources σε <15s.\nSurgical coexistence σε production\ncluster. 5 APIs ταυτοχρονα."),
        ("RQ4: Μπορουν τα quality gates\nνα εκφραστουν ως οντολογικα\nspecifications;",
         "NAI",
         "XSDLC = discoverable, inspectable,\ncomposable Kubernetes-native\nοντολογικο specification."),
    ],

    # S19
    "s19_title": "Τεσσερις Συνεισφορες (Κεφ. 10)",
    "s19_contribs": [
        ("1. Μεθοδολογια DDT", [
            "3 αξιωματα, 9 αρχες, 4 validation modes",
            "Zero manual test authorship",
            "Σταθερο αξιωματικο πλαισιο, εξελιξιμες αρχες",
            "Severity model: 2 critical, 7 warning",
        ]),
        ("2. DriveBy Framework", [
            "Go CLI: 11 packages, 7,170 LOC",
            "APISpec abstraction (Adapter pattern)",
            "7 principle checkers (Strategy pattern)",
            "JSON + Markdown reports, 4-code exit protocol",
        ]),
        ("3. XSDLC Pipeline", [
            "Crossplane CRD: ~25 YAML --> ~44 resources",
            "Dual-provider: provider-kubernetes + provider-upjet-github",
            "BYOCI: developer owns CI, XSDLC owns delivery",
            "Level-triggered reconciliation, self-healing",
        ]),
        ("4. Agent-Driven Development", [
            "CLAUDE.md knowledge protocol (16 αρχεια)",
            "57 sessions, CI-enforced documentation",
            "Αμφιδρομη ιχνηλασιμοτητα code <--> thesis",
            "DDT output ως agent feedback infrastructure",
        ]),
    ],

    # S20
    "s20_title": "Περιορισμοι & Μελλοντικη Εργασια (Κεφ. 10)",
    "s20_lim_title": "Περιορισμοι",
    "s20_lims": [
        "P006, P007: η λογικη υπαρχει αλλα χωρις PrincipleChecker wrapper",
        "P008: keyword-based matching στο info.description (ευθραυστο)",
        "Μονο OpenAPI 3.x + Swagger 2.0 (χωρις GraphQL, gRPC, AsyncAPI)",
        "Στατικα load test thresholds (χωρις αυτοπροσαρμογη)",
        "Αξιολογηση: ενα IDP (Novelcore), τα public APIs δεν εκπροσωπουν enterprise",
        "Δυαδικο pass/fail scoring (χωρις ποσοστιαια κοκκωση)",
    ],
    "s20_fut_title": "Μελλοντικη Εργασια",
    "s20_future": [
        "Adaptive load thresholds: controller παρακολουθει Prometheus",
        "Agent-driven remediation: AI αναλυει DDT reports, διορθωνει specs",
        "P006/P007 ως PrincipleCheckers στο ενιαιο pipeline",
        "Multi-format: GraphQL, gRPC, AsyncAPI adapters",
        "Enterprise multi-cluster delivery patterns",
        "XSDLC ως child resource μεγαλυτερου operator",
    ],

    # S21
    "s21_quote1": "Σταματας να λες στο συστημα τι να κανει.\nΑρχιζεις να του λες πως πρεπει να ΕΙΝΑΙ ο κοσμος.",
    "s21_quote2": (
        "Απο imperative scripts σε ontological specifications.\n"
        "Το συστημα συγκλινει προς τη δηλωμενη κατασταση συνεχως,\n"
        "χωρις ανθρωπινη εναρξη, χωρις ανθρωπινη ερμηνεια.\n"
        "-- Κεφαλαιο 10"
    ),
    "s21_quote3": (
        "Η διασφαλιση ποιοτητας γινεται δυνατοτητα πλατφορμας\n"
        "αντι για ευθυνη καθε ομαδας ξεχωριστα."
    ),
    "s21_thanks": "Ευχαριστω",
    "s21_contact": "Πετρος Ευαγγελος Τριανταφυλλης  --  Πανεπιστημιο Αιγαιου  --  github.com/meter-peter/driveby",
}


if __name__ == "__main__":
    out_dir = os.path.dirname(os.path.abspath(__file__))

    for lang, suffix in [("gr", "GR"), ("en", "EN")]:
        prs = build(lang)
        path = os.path.join(out_dir, f"DDT-Presentation-{suffix}.pptx")
        prs.save(path)
        print(f"Created: {path}")
