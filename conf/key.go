package conf

import (
	"os"
)

var (
	PriKey []byte // PriKey 私钥
	PubKey []byte // PubKey 公钥
)

// InitKey 初始化密钥
func InitKey() {
	var err error
	PriKey, err = LoadPrivateKey()
	if err != nil {
		panic(err)
	}
	PubKey, err = LoadPublicKey()
	if err != nil {
		panic(err)
	}
}

// LoadPrivateKey 加载私钥内容
func LoadPrivateKey() ([]byte, error) {
	file := filePath(privateKeyFile, keyConfDir)
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LoadPublicKey 加载公钥内容
func LoadPublicKey() ([]byte, error) {
	file := filePath(publicKeyFile, keyConfDir)
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
