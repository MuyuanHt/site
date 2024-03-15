package conf

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"site/protocol/shared"
)

// ParamConfig 全局参数配置信息
var ParamConfig map[string]string

// InitParamConfig 初始化全局通用参数配置
func InitParamConfig() {
	err := LoadParamConf()
	if err != nil {
		panic(err)
	}
}

// LoadParamConf 加载参数相关配置
func LoadParamConf() error {
	ParamConfig = make(map[string]string)
	file := filePath(paramConfFile, paramDir)
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &ParamConfig)
	if err != nil {
		return err
	}
	return nil
}

// GetConfigParam 获取配置信息
func GetConfigParam(name string) (string, error) {
	params, ok := ParamConfig[name]
	if !ok {
		return "", errors.New(shared.CodeMessageIgnoreCode(shared.ConfNotFound))
	}
	return params, nil
}
