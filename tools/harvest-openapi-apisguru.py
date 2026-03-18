#!/usr/bin/env python3
"""
harvest-openapi-apisguru.py

Fetch a small sample (max 20) of public OpenAPI 3.x specifications from APIs.guru
and write them to a CSV compatible with DriveBy tools (name,url,host).

Usage:
  tools/harvest-openapi-apisguru.py output.csv [min_version]

min_version:
  - \"3.0\" (default): accept OpenAPI 3.x (openapi=\"3.*\")
  - \"3.1\": accept only OpenAPI 3.1 (openapi=\"3.1*\")

Notes:
- This script fetches https://api.apis.guru/v2/list.json, then for each API:
  - Looks at the preferred version metadata.
  - Fetches the preferred spec JSON and checks openapi version against min_version.
  - Derives the host from the first server URL (OpenAPI 3.x) or from the spec URL.
"""

import csv
import json
import sys
from typing import Any, Dict, List, Optional, Tuple

import urllib.error
import urllib.parse
import urllib.request


APIS_GURU_URL = "https://api.apis.guru/v2/list.json"
MAX_APIS = 20


def fetch_json(url: str, timeout: float = 15.0) -> Optional[Dict[str, Any]]:
  try:
    headers = {
      "Accept": "application/json",
      "User-Agent": "DriveByHarvester/1.0",
    }
    req = urllib.request.Request(url, headers=headers)
    with urllib.request.urlopen(req, timeout=timeout) as resp:
      data = resp.read()
    return json.loads(data.decode("utf-8"))
  except Exception:
    return None


def resolve_preferred_spec(meta: Dict[str, Any]) -> Optional[str]:
  # APIs.guru metadata usually includes 'swaggerUrl' or 'openapiUrl'
  url = meta.get("openapiUrl") or meta.get("swaggerUrl")
  if isinstance(url, str) and url:
    return url
  return None


def fetch_spec_with_min_version(spec_url: str, min_version: str) -> Optional[Dict[str, Any]]:
  spec = fetch_json(spec_url)
  if not spec:
    return None
  openapi_ver = spec.get("openapi")

  if not isinstance(openapi_ver, str):
    return None

  if min_version == "3.1":
    if openapi_ver.startswith("3.1"):
      return spec
    return None

  # Default: 3.0 – accept any 3.x spec
  if openapi_ver.startswith("3."):
    return spec
  return None


def extract_host(spec: Dict[str, Any], spec_url: str) -> str:
  servers = spec.get("servers")
  if isinstance(servers, list) and servers:
    first = servers[0]
    if isinstance(first, dict):
      url = first.get("url")
      if isinstance(url, str) and url:
        parsed = urllib.parse.urlparse(url)
        if parsed.hostname:
          return parsed.hostname
  # Fallback to host from spec URL
  parsed = urllib.parse.urlparse(spec_url)
  if parsed.hostname:
    return parsed.hostname
  return spec_url


def harvest(output_csv: str, min_version: str) -> int:
  root = fetch_json(APIS_GURU_URL)
  if not root:
    print("Failed to fetch APIs.guru list.json", file=sys.stderr)
    return 2

  rows: List[Tuple[str, str, str]] = []

  for api_name, versions in root.items():
    if len(rows) >= MAX_APIS:
      break
    if not isinstance(versions, dict):
      continue

    preferred_meta: Optional[Dict[str, Any]] = None
    for ver, meta in versions.items():
      if not isinstance(meta, dict):
        continue
      if meta.get("preferred"):
        preferred_meta = meta
        break
    if not preferred_meta:
      # fallback to any version
      for meta in versions.values():
        if isinstance(meta, dict):
          preferred_meta = meta
          break
    if not preferred_meta:
      continue

    spec_url = resolve_preferred_spec(preferred_meta)
    if not spec_url:
      continue

    spec = fetch_spec_with_min_version(spec_url, min_version)
    if not spec:
      continue

    host = extract_host(spec, spec_url)
    # Use api_name plus version label from meta if present
    version_label = preferred_meta.get("version") or preferred_meta.get("info", {}).get("version")
    name = f"{api_name}-{version_label}" if version_label else api_name
    rows.append((name, spec_url, host))

  with open(output_csv, "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["name", "url", "host"])
    for name, url, host in rows:
      writer.writerow([name, url, host])

  print(f"Wrote {len(rows)} OpenAPI 3.1 APIs to {output_csv}")
  return 0


def main(argv: List[str]) -> int:
  if len(argv) < 2 or len(argv) > 3:
    print("Usage: harvest-openapi-apisguru.py output.csv [min_version]", file=sys.stderr)
    return 3
  output_csv = argv[1]
  min_version = "3.0"
  if len(argv) == 3:
    min_version = argv[2]
  return harvest(output_csv, min_version)


if __name__ == "__main__":
  raise SystemExit(main(sys.argv))

