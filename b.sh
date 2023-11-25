#!/bin/bash

# 遍历子目录下的所有 .proto 文件
find . -type f -name "*.proto" | while read file; do
  # 获取当前文件的相对路径
  relative_path=$(dirname "${file}")

  # 获取文件名（不包含扩展名）
  file_name=$(basename "${file}" .proto)

  # 构建 protoc 命令
  protoc_command="protoc -I=/root/src -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${file}"

  # 运行 protoc 命令
  eval "${protoc_command}"

  echo "已处理文件：${file}"
done

echo "脚本执行完成"
