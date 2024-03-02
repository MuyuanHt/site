package conf

import (
	"os"
)

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
