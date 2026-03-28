#!/usr/bin/env bash

set -euo pipefail

upstream_dir="${1:-upstream/grpc_api}"
repo_root="${2:-.}"
manifest_file="${3:-${repo_root}/.proto-sync-manifest}"

if [[ ! -d "${upstream_dir}" ]]; then
  echo "upstream proto directory not found: ${upstream_dir}" >&2
  exit 1
fi

tmp_manifest="$(mktemp)"
old_manifest="$(mktemp)"

if [[ -f "${manifest_file}" ]]; then
  grep -v '^[[:space:]]*$' "${manifest_file}" | sort -u > "${old_manifest}"
fi

copied_count=0

while IFS= read -r -d '' source_file; do
  relative_path="${source_file#${upstream_dir}/}"
  destination_file="${repo_root}/${relative_path}"

  mkdir -p "$(dirname "${destination_file}")"
  cp "${source_file}" "${destination_file}"
  printf '%s\n' "${relative_path}" >> "${tmp_manifest}"
  copied_count=$((copied_count + 1))
done < <(find "${upstream_dir}" -type f -name '*.proto' -print0)

sort -u "${tmp_manifest}" -o "${tmp_manifest}"

if [[ -s "${old_manifest}" ]]; then
  while IFS= read -r previous_path; do
    if ! grep -Fxq "${previous_path}" "${tmp_manifest}"; then
      destination_file="${repo_root}/${previous_path}"

      rm -f "${destination_file}"

      cleanup_dir="$(dirname "${destination_file}")"
      while [[ "${cleanup_dir}" != "${repo_root}" && "${cleanup_dir}" != "." ]]; do
        rmdir "${cleanup_dir}" 2>/dev/null || break
        cleanup_dir="$(dirname "${cleanup_dir}")"
      done
    fi
  done < "${old_manifest}"
fi

mv "${tmp_manifest}" "${manifest_file}"
rm -f "${old_manifest}"

echo "synced ${copied_count} proto file(s)"
