#!/bin/bash

# 设置 Go package 的基础路径
go_package_base="github.com/stmtc233/bili-grpc-api-go"

# 遍历子目录下的所有 .proto 文件
find . -type f -name "*.proto" | while read file; do
  # 获取当前文件的相对路径
  relative_path=$(dirname "${file}")

  # 获取规范的绝对路径
  absolute_path=$(realpath --relative-to=. "${relative_path}")

  # 拼接完整的 Go package 路径
  go_package="${go_package_base}/${absolute_path}"

  # 在 .proto 文件的最后添加两个空行
  echo -e "\n\n" >> "${file}"

  # 添加 option go_package 到 .proto 文件的最后一行
  echo "option go_package = \"$go_package\";" >> "${file}"

  echo "已更新文件：${file}"
done

echo "脚本执行完成"
