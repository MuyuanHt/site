package inits

import (
	"site/common/logs"
	"site/conf"
)

// AppInitialize 全局配置初始化
func AppInitialize() {
	conf.InitGlobalAppConfig() // 全局服务配置
	conf.InitParamConfig()     // 全局通用参数配置
	conf.InitKey()             // 初始化密钥

}

// ToolInitialize 全局工具初始化
func ToolInitialize() {
	logs.InitLogger() // 全局日志工具
}
