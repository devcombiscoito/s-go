#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <files...>"
  exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
  sha256sum "$@" > SHA256SUMS
elif command -v shasum >/dev/null 2>&1; then
  shasum -a 256 "$@" > SHA256SUMS
else
  echo "sha256sum or shasum not found"
  exit 1
fi

echo "Wrote SHA256SUMS"
