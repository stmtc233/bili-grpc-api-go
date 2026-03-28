#!/usr/bin/env bash

set -euo pipefail

requested_tag="${1:-}"

if [[ -n "${requested_tag}" ]]; then
  if [[ ! "${requested_tag}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "tag must match v<major>.<minor>.<patch>" >&2
    exit 1
  fi

  echo "${requested_tag}"
  exit 0
fi

latest_tag="$(git tag --list 'v*' --sort=-version:refname | head -n 1)"

if [[ -z "${latest_tag}" ]]; then
  echo "v1.0.0"
  exit 0
fi

if [[ ! "${latest_tag}" =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
  echo "latest tag is not a supported semver tag: ${latest_tag}" >&2
  exit 1
fi

major="${BASH_REMATCH[1]}"
minor="${BASH_REMATCH[2]}"
patch="${BASH_REMATCH[3]}"

echo "v${major}.${minor}.$((patch + 1))"
