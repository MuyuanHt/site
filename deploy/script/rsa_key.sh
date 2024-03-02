#!/usr/bin/env bash

cd ../../conf/key || exit

if [ -f "private_key.pem" ]&&[ -f "public_key.pem" ]; then
  rm "private_key.pem" "public_key.pem"
  echo "deleted files successfully"
fi

# 生成 rsa 私钥
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048

# 根据 rsa 私钥生成公钥
openssl rsa -pubout -in private_key.pem -out public_key.pem