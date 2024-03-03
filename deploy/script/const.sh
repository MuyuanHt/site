#!/usr/bin/env bash
# shellcheck disable=SC2129

# 删除旧的 const_type.go 存在则删除
cd ../../protocol/shared || exit
if [ -f "const_type.go" ]; then
  rm "const_type.go"
  echo "deleted file successfully"
fi

# 清除文件内容 没有文件则创建 用于创建新的 const_type.go
# 使用 > 表示覆盖 使用 >> 表示追加
:> "const_type.txt"

# 头部
{
  cat <<EOF
  package shared

  // This file use shell const.sh auto created
  const (
EOF
} >> const_type.txt

# 主体部分 逐行读取文件内容并处理
while IFS= read -r line; do
  # 跳过包含 # 的行
  if echo "$line" | grep -q '#'; then
    continue
  fi
  numbers=$(echo "$line" | cut -d'-' -f 1)  # 获取原始行中的数字部分
  defines=$(echo "$line" | cut -d'-' -f 2)  # 获取原始行中的定义部分
  comment=$(echo "$line" | cut -d'-' -f 3)  # 获取原始行中的注释部分
  # 如果任何一个变量为空跳过本次循环
  if [[ -z $numbers || -z $defines || -z $comment ]]; then
    continue
  fi
  echo "$defines = $numbers // $comment"  # 将处理后的内容输出
done < const.txt >> const_type.txt  # 将处理后的内容添加到 const_type.txt

# 尾部
echo ")" >> const_type.txt

echo "write const_type.txt successfully"
mv const_type.txt const_type.go
go fmt const_type.go
echo "created const_type.go successfully"