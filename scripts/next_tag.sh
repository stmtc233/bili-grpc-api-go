#!/usr/bin/env bash

set -euo pipefail

requested_tag="${1:-}"
tag_prefix="${2:-v}"

if [[ -z "${tag_prefix}" || ! "${tag_prefix}" =~ ^[A-Za-z0-9._-]+$ ]]; then
  echo "tag prefix must contain only letters, numbers, dots, underscores, and hyphens" >&2
  exit 1
fi

if [[ -n "${requested_tag}" ]]; then
  requested_version="${requested_tag#"${tag_prefix}"}"
  if [[ "${requested_tag}" == "${requested_version}" || ! "${requested_version}" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "tag must match ${tag_prefix}<major>.<minor>.<patch>" >&2
    exit 1
  fi

  echo "${requested_tag}"
  exit 0
fi

latest_tag=""
while IFS= read -r candidate; do
  candidate_version="${candidate#"${tag_prefix}"}"
  if [[ "${candidate}" != "${candidate_version}" && "${candidate_version}" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    latest_tag="${candidate}"
    break
  fi
done < <(git tag --list "${tag_prefix}*" --sort=-version:refname)

if [[ -z "${latest_tag}" ]]; then
  echo "${tag_prefix}1.0.0"
  exit 0
fi

latest_version="${latest_tag#"${tag_prefix}"}"
if [[ ! "${latest_version}" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
  echo "latest tag is not a supported semver tag: ${latest_tag}" >&2
  exit 1
fi

major="${BASH_REMATCH[1]}"
minor="${BASH_REMATCH[2]}"
patch="${BASH_REMATCH[3]}"

echo "${tag_prefix}${major}.${minor}.$((patch + 1))"
