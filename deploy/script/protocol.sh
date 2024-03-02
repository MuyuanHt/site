#!/usr/bin/env bash

# 删除旧的 protoc 编译文件 *.pb.go
echo "delete old *.pb.go files start"
cd ../../protocol || exit
current_path=$(pwd)

# 使用 find 命令获取文件列表并删除 *.pb.go 文件
while IFS= read -r -d '' file; do
  echo "deleted: ${file}"
  rm "${file}"
done < <(find "${current_path}" -type f -name "*.pb.go" -print0)

echo "delete old *.pb.go files successfully"
echo

# 重新编译 proto 文件
echo "building proto files start"
cd ..
protoc --go_out=paths=source_relative:. \
       --go-grpc_out=paths=source_relative:. \
       protocol/auth/*.proto \
       protocol/calendar/*.proto \
       protocol/chat/*.proto \
       protocol/collaborate/*.proto \
       protocol/document/*.proto \
       protocol/user/*.proto
echo "building proto files successfully"