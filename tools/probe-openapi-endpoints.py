#!/usr/bin/env python3
"""
probe-openapi-endpoints.py

Given a CSV of OpenAPI specs (name,url,host), probe each API to find
whether at least one simple GET endpoint behaves as documented.

This is intended to help you build a small curated list of public APIs
where BOTH:
- The OpenAPI spec is reachable.
- At least one documented endpoint returns a documented status code
  (typically 2xx/3xx) and a response Content-Type that matches the spec.

Usage:
  tools/probe-openapi-endpoints.py input.csv output.csv

Limitations:
- Only JSON OpenAPI/Swagger documents are supported (no YAML parsing).
- Only unauthenticated GET endpoints are probed.
"""

import csv
import json
import sys
import urllib.parse
from typing import Any, Dict, List, Optional, Tuple

import urllib.request
import urllib.error


def fetch_json(url: str, timeout: float = 10.0) -> Optional[Dict[str, Any]]:
  try:
    req = urllib.request.Request(url, headers={"Accept": "application/json"})
    with urllib.request.urlopen(req, timeout=timeout) as resp:
      body = resp.read()
    return json.loads(body.decode("utf-8"))
  except Exception:
    return None


def build_base_url(host: str) -> str:
  if host.startswith("http://") or host.startswith("https://"):
    return host.rstrip("/")
  return f"https://{host.rstrip('/')}"


def iter_openapi3_paths(spec: Dict[str, Any]) -> List[Tuple[str, Dict[str, Any]]]:
  paths = spec.get("paths") or {}
  items: List[Tuple[str, Dict[str, Any]]] = []
  for path, operations in paths.items():
    if not isinstance(operations, dict):
      continue
    items.append((path, operations))
  return items


def pick_simple_endpoints(spec: Dict[str, Any]) -> List[Tuple[str, str, Dict[str, Any]]]:
  """
  Return a small list of (method, path, operation_obj) for simple GET endpoints:
  - method GET
  - path without path parameters ('{')
  - no required query parameters
  """
  candidates: List[Tuple[str, str, Dict[str, Any]]] = []
  for path, operations in iter_openapi3_paths(spec):
    if "{" in path:
      continue
    op = operations.get("get") or operations.get("GET")
    if not isinstance(op, dict):
      continue
    params = op.get("parameters") or []
    if any(
      isinstance(p, dict)
      and p.get("in") == "query"
      and p.get("required") is True
      for p in params
    ):
      continue
    candidates.append(("GET", path, op))
  return candidates[:5]


def documented_success_statuses(operation: Dict[str, Any]) -> List[int]:
  responses = operation.get("responses") or {}
  codes: List[int] = []
  for code_str in responses.keys():
    if code_str == "default":
      continue
    try:
      code_int = int(code_str)
    except ValueError:
      continue
    if 200 <= code_int < 400:
      codes.append(code_int)
  return codes


def content_types_for_status(operation: Dict[str, Any], status: int) -> List[str]:
  responses = operation.get("responses") or {}
  resp_obj = responses.get(str(status)) or responses.get("default") or {}
  content = resp_obj.get("content") or {}
  return [ct.lower() for ct in content.keys()]


def http_get(url: str, timeout: float = 5.0) -> Tuple[int, str]:
  req = urllib.request.Request(url, headers={"Accept": "application/json"})
  try:
    with urllib.request.urlopen(req, timeout=timeout) as resp:
      status = resp.getcode()
      ct = resp.headers.get("Content-Type", "") or ""
      return status, ct.lower()
  except urllib.error.HTTPError as e:
    ct = e.headers.get("Content-Type", "") if e.headers else ""
    return e.code, (ct or "").lower()
  except Exception:
    return 0, ""


def probe_api(name: str, spec_url: str, host: str) -> bool:
  spec = fetch_json(spec_url)
  if not spec:
    return False

  candidates = pick_simple_endpoints(spec)
  if not candidates:
    return False

  base = build_base_url(host)

  for method, path, op in candidates:
    success_codes = documented_success_statuses(op)
    if not success_codes:
      continue

    url = urllib.parse.urljoin(base + "/", path.lstrip("/"))
    status, ct = http_get(url)
    if status in success_codes:
      expected_cts = content_types_for_status(op, status)
      if not expected_cts:
        return True
      if any(ct.startswith(ect) for ect in expected_cts):
        return True

  return False


def main(argv: List[str]) -> int:
  if len(argv) != 3:
    print("Usage: probe-openapi-endpoints.py input.csv output.csv", file=sys.stderr)
    return 3

  input_csv, output_csv = argv[1], argv[2]

  kept_rows: List[Tuple[str, str, str]] = []

  with open(input_csv, newline="") as f:
    reader = csv.DictReader(f)
    for row in reader:
      name = (row.get("name") or "").strip()
      url = (row.get("url") or "").strip()
      host = (row.get("host") or "").strip()
      if not name or not url or not host:
        continue
      ok = probe_api(name, url, host)
      if ok:
        kept_rows.append((name, url, host))

  with open(output_csv, "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["name", "url", "host"])
    for row in kept_rows:
      writer.writerow(list(row))

  print(f"Probed {len(kept_rows)} working APIs out of input CSV.")
  return 0


if __name__ == "__main__":
  raise SystemExit(main(sys.argv))

