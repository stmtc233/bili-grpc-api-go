#!/usr/bin/env bash

set -euo pipefail

repo_root="${1:-.}"

(
  cd "${repo_root}"

  proto_files="$(find . -type f -name '*.proto' ! -path './.git/*' ! -path './upstream/*' | LC_ALL=C sort)"

  if [[ -z "${proto_files}" ]]; then
    echo "no proto files found"
    exit 0
  fi

  while IFS= read -r file; do
    [[ -z "${file}" ]] && continue

    protoc \
      -I=. \
      --go_out=. \
      --go_opt=paths=source_relative \
      --go-grpc_out=. \
      --go-grpc_opt=paths=source_relative \
      "${file}"
  done <<< "${proto_files}"
)

echo "generated go stubs from proto files"
