#!/usr/bin/env bash

# 删除旧的 http_code_type.go 存在则删除
cd ../../protocol/shared || exit
if [ -f "http_code_type.go" ]; then
  rm "http_code_type.go"
  echo "deleted file successfully"
fi

:> "http_code_type.txt"

# 头部
# shellcheck disable=SC2129
{
  echo "package shared"
  echo "// This file use shell \"http_code.sh\" auto created"
  echo "const ("
} >> http_code_type.txt

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
  if [[ -z $numbers || -z $defines ||-z $comment ]]; then
    continue
  fi
  echo "$defines = $numbers // $comment"  # 将处理后的内容输出
done < http_code.txt >> http_code_type.txt  # 将处理后的内容添加到 http_code_type.txt

# 中间
{
  echo ")"
  echo
  echo "var resultCodeText = map[int]string{"
} >> http_code_type.txt

while IFS= read -r line; do
  # 跳过包含 # 的行
  if echo "$line" | grep -q '#'; then
    continue
  fi
  numbers=$(echo "$line" | cut -d'-' -f 1)  # 获取原始行中的数字部分
  defines=$(echo "$line" | cut -d'-' -f 2)  # 获取原始行中的定义部分
  comment=$(echo "$line" | cut -d'-' -f 3)  # 获取原始行中的注释部分
  # 如果任何一个变量为空跳过本次循环
  if [[ -z $numbers || -z $defines ||-z $comment ]]; then
    continue
  fi
  echo "$defines:\"$comment\","  # 将处理后的内容输出
done < http_code.txt >> http_code_type.txt  # 将处理后的内容添加到 http_code_type.txt

# 末尾
{
  echo "}"
  echo
  echo "// CodeMessage 获取 code 对应的 message"
  echo "func CodeMessage(code int) (string, bool) {"
  echo "message, ok := resultCodeText[code]"
  echo "return message, ok }"

  echo "// CodeMessageIgnoreCode 获取 code 对应的 message 未查询到状态码时返回指定的状态码异常错误"
  echo "func CodeMessageIgnoreCode(code int) string {"
  echo "message, ok := resultCodeText[code]"
  echo "if !ok { return resultCodeText[CodeNotFound] }" # 状态码异常 10101
  echo "return message}"
} >> http_code_type.txt

echo "write http_code_type.txt successfully"
mv http_code_type.txt http_code_type.go
go fmt http_code_type.go
echo "create http_code_type.go successfully"