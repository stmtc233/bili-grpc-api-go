#!/usr/bin/env bash

set -euo pipefail

repo_root="${1:-.}"
module_path="$(awk '/^module / { print $2; exit }' "${repo_root}/go.mod")"

if [[ -z "${module_path}" ]]; then
  echo "failed to read module path from ${repo_root}/go.mod" >&2
  exit 1
fi

updated_count=0

while IFS= read -r -d '' file; do
  relative_file="${file#./}"
  relative_dir="$(dirname "${relative_file}")"

  if [[ "${relative_dir}" == "." ]]; then
    expected_go_package="${module_path}"
  else
    expected_go_package="${module_path}/${relative_dir}"
  fi

  tmp_file="$(mktemp)"

  awk -v expected="${expected_go_package}" '
    BEGIN {
      replaced = 0
    }
    /^[[:space:]]*option go_package = ".*";[[:space:]]*$/ {
      if (!replaced) {
        print "option go_package = \"" expected "\";"
        replaced = 1
      }
      next
    }
    {
      print
    }
    END {
      if (!replaced) {
        print ""
        print "option go_package = \"" expected "\";"
      }
    }
  ' "${repo_root}/${relative_file}" > "${tmp_file}"

  if ! cmp -s "${repo_root}/${relative_file}" "${tmp_file}"; then
    mv "${tmp_file}" "${repo_root}/${relative_file}"
    updated_count=$((updated_count + 1))
  else
    rm -f "${tmp_file}"
  fi
done < <(
  cd "${repo_root}"
  find . -type f -name '*.proto' ! -path './.git/*' ! -path './upstream/*' -print0
)

echo "updated go_package in ${updated_count} proto file(s)"
