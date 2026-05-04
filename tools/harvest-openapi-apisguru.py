#!/usr/bin/env python3
"""
harvest-openapi-apisguru.py

Fetch a small sample (max 20) of public OpenAPI / Swagger specifications
from APIs.guru and write them to a CSV compatible with DriveBy tools
(name,url,host).

Usage:
  tools/harvest-openapi-apisguru.py output.csv [version_selector]

version_selector:
  - "3.0" (default): accept OpenAPI 3.x (openapi="3.*")
  - "3.1": accept only OpenAPI 3.1 (openapi="3.1*")
  - "2.0": accept only Swagger 2.0 (swagger="2.0")

Notes:
- This script fetches https://api.apis.guru/v2/list.json, then for each API:
  - Looks at the preferred version metadata.
  - Fetches the preferred spec JSON and checks the version against the
    selector (the "openapi" field for 3.x, the "swagger" field for 2.0).
  - Derives the host from the first server URL (OpenAPI 3.x) or from the
    "host" field (Swagger 2.0); falls back to the host of the spec URL.
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


def fetch_spec_with_min_version(spec_url: str, version_selector: str) -> Optional[Dict[str, Any]]:
  spec = fetch_json(spec_url)
  if not spec:
    return None

  # Swagger 2.0 selector: accept only specs with swagger:"2.0".
  if version_selector == "2.0":
    swagger_ver = spec.get("swagger")
    if isinstance(swagger_ver, str) and swagger_ver == "2.0":
      return spec
    return None

  # OpenAPI 3.x selectors operate on the "openapi" field.
  openapi_ver = spec.get("openapi")
  if not isinstance(openapi_ver, str):
    return None

  if version_selector == "3.1":
    if openapi_ver.startswith("3.1"):
      return spec
    return None

  # Default: 3.0 – accept any 3.x spec
  if openapi_ver.startswith("3."):
    return spec
  return None


def extract_host(spec: Dict[str, Any], spec_url: str) -> str:
  # OpenAPI 3.x: first server URL.
  servers = spec.get("servers")
  if isinstance(servers, list) and servers:
    first = servers[0]
    if isinstance(first, dict):
      url = first.get("url")
      if isinstance(url, str) and url:
        parsed = urllib.parse.urlparse(url)
        if parsed.hostname:
          return parsed.hostname
  # Swagger 2.0: top-level "host" field.
  swagger_host = spec.get("host")
  if isinstance(swagger_host, str) and swagger_host:
    # The "host" field may include a port; urlparse needs a scheme to extract it.
    parsed = urllib.parse.urlparse(f"http://{swagger_host}")
    if parsed.hostname:
      return parsed.hostname
  # Fallback to host from spec URL
  parsed = urllib.parse.urlparse(spec_url)
  if parsed.hostname:
    return parsed.hostname
  return spec_url


def harvest(output_csv: str, version_selector: str) -> int:
  root = fetch_json(APIS_GURU_URL)
  if not root:
    print("Failed to fetch APIs.guru list.json", file=sys.stderr)
    return 2

  rows: List[Tuple[str, str, str]] = []

  # For OpenAPI 3.x selectors we follow the original behaviour and look at
  # the API's *preferred* version only (APIs.guru's curated default).
  # For Swagger 2.0 we scan *every* version of each API entry, because
  # APIs.guru aggressively prefers 3.x when an upgraded spec exists, so
  # the preferred-only path returns near-zero 2.0 hits even though
  # thousands of APIs still ship 2.0 specs as legacy/historical versions.
  scan_all_versions = version_selector == "2.0"

  # Cap how many APIs are accepted from a single provider (the part of
  # the api_name before the first ':' or '/'). This avoids datasets
  # dominated by a single provider's many sub-APIs (e.g. amadeus.com
  # exposes ~30+ Swagger 2.0 sub-APIs).
  PER_PROVIDER_CAP = 2
  per_provider_count: Dict[str, int] = {}

  for api_name, entry in root.items():
    if len(rows) >= MAX_APIS:
      break
    if not isinstance(entry, dict):
      continue

    # The APIs.guru list.json schema is:
    #   {api_name: {added, preferred: <version_label>, versions: {<v>: <meta>}}}
    # Each <meta> contains swaggerUrl/openapiUrl, info, externalDocs, ...
    versions_dict = entry.get("versions") if isinstance(entry.get("versions"), dict) else entry
    preferred_label = entry.get("preferred") if isinstance(entry.get("preferred"), str) else None

    candidate_metas: List[Tuple[str, Dict[str, Any]]] = []
    if scan_all_versions:
      for ver, meta in versions_dict.items():
        if isinstance(meta, dict):
          candidate_metas.append((ver, meta))
    else:
      preferred_meta: Optional[Dict[str, Any]] = None
      preferred_ver: Optional[str] = None
      if preferred_label and isinstance(versions_dict.get(preferred_label), dict):
        preferred_meta = versions_dict[preferred_label]
        preferred_ver = preferred_label
      if not preferred_meta:
        for ver, meta in versions_dict.items():
          if isinstance(meta, dict):
            preferred_meta = meta
            preferred_ver = ver
            break
      if preferred_meta:
        candidate_metas.append((preferred_ver or "?", preferred_meta))

    provider = api_name.split(":")[0].split("/")[0]
    if per_provider_count.get(provider, 0) >= PER_PROVIDER_CAP:
      continue

    for ver, meta in candidate_metas:
      if len(rows) >= MAX_APIS:
        break
      spec_url = resolve_preferred_spec(meta)
      if not spec_url:
        continue
      spec = fetch_spec_with_min_version(spec_url, version_selector)
      if not spec:
        continue

      host = extract_host(spec, spec_url)
      version_label = meta.get("version") or meta.get("info", {}).get("version") or ver
      name = f"{api_name}-{version_label}" if version_label else api_name
      rows.append((name, spec_url, host))
      per_provider_count[provider] = per_provider_count.get(provider, 0) + 1
      # In multi-version scans, take only the first matching version per API
      # entry to maximise breadth across distinct APIs.
      if scan_all_versions:
        break

  with open(output_csv, "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["name", "url", "host"])
    for name, url, host in rows:
      writer.writerow([name, url, host])

  label = {"2.0": "Swagger 2.0", "3.1": "OpenAPI 3.1", "3.0": "OpenAPI 3.x"}.get(
      version_selector, version_selector)
  print(f"Wrote {len(rows)} {label} APIs to {output_csv}")
  return 0


def main(argv: List[str]) -> int:
  if len(argv) < 2 or len(argv) > 3:
    print("Usage: harvest-openapi-apisguru.py output.csv [version_selector]", file=sys.stderr)
    print("  version_selector: 3.0 (default) | 3.1 | 2.0", file=sys.stderr)
    return 3
  output_csv = argv[1]
  version_selector = "3.0"
  if len(argv) == 3:
    version_selector = argv[2]
  return harvest(output_csv, version_selector)


if __name__ == "__main__":
  raise SystemExit(main(sys.argv))

