# Submission Package — current state (2026-08-17)

> Supersedes the 2026-05-05 version of this file, which described the previous
> (610-comment) round; that text remains in git history.

## Contents of `thesis/submission-package/`

| File | What it is | Built |
|------|-----------|-------|
| `01-thesis.pdf` | Final thesis, **219 pages**, all 480 final-round comments incorporated + post-round consistency/language audit fixes | 2026-08-17, pdflatex+biber, clean log (0 errors, 0 undefined refs/citations, 1 cosmetic 0.5 pt overfull) |
| `02-presentation.pdf` | Final defense presentation (rev3, 20–30 min format) | 2026-06-19 |

## What accompanies the package

- Cover email: `review/final-submission-email.md` (short cover, Greek).
- Full reply to the 480 comments: `review/kritikos-480-reply-email.md`
  (all 15 headline points, the 19 author decisions, justified deferrals,
  the #427 answer, and the cluster-rerun caveat).

## Pre-send checklist

- [x] Thesis PDF in package is the current build (not a stale round)
- [x] Committee names on both title pages (Kokolakis, Livieris)
- [x] "Diploma Thesis" designation on both English and Greek title pages
- [x] DDT→SDT rename complete (footnoted historical mentions only)
- [x] All cross-chapter numbers reconciled (gaps=7, misses accounting per Ch.7,
      latency 40–70 s median 69 s, ~35-line CR, 44–47/45–50 resources, vegeta not k6,
      strict = P001–P005+P008, commit split 22+4+26)
- [x] Bibliography author corrections (Dygalo, Senart, Becker, Lieret+Yao)
- [ ] Reply email reviewed and sent by the author
- [ ] Optional: re-run cluster PoC gates under corrected specs (est. half a day)

## Provenance

Fix round of 2026-08-17: six-agent audit (consistency, citations, cross-refs,
language, claims-vs-code, review coverage) followed by a full fix pass —
see `thesis/review/` git history and the audit artifact.
