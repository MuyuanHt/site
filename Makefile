
.PHONY: p a b c
p: proto const code key

# 构建协议
proto:
	cd deploy/script && sh protocol.sh

# 构建常量
const:
	cd deploy/script && sh const.sh

# 构建自定义状态码
code:
	cd deploy/script && sh http_code.sh

# 构建密钥
key:
	cd deploy/script && sh rsa_key.sh