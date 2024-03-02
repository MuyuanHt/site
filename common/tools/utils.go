package tools

import (
	"os"
	"path"
	"regexp"
	consts "site/protocol/shared"
	"strings"
)

// GetCurrentPath 当前执行程序的文件路径
func GetCurrentPath() string {
	// runtime.Caller(skip) 获取调用者路径 参数 skip 表示要提升的堆栈帧数 0-当前函数 1-上一层函数
	// os.Getwd() 获取程序运行时路径
	cur, _ := os.Getwd()
	// 处理 windows 的反斜杠路径分割符 将 \ 全部替换为 /
	filePath := strings.ReplaceAll(cur, "\\", "/")
	return filePath
}

// GetCurrentDir 获取路径的最后一级目录
func GetCurrentDir(filePath string) string {
	return path.Base(filePath)
}

// GetAccountType 获取账号所属类型
func GetAccountType(account string) int32 {
	// 正则调用规则
	phoneRegexp := regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`) // 手机号
	emailRegexp := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)                                // 邮箱
	// 返回 MatchString 是否匹配
	if phoneRegexp.MatchString(account) {
		return consts.AccountTypePhone
	}
	if emailRegexp.MatchString(account) {
		return consts.AccountTypeEmail
	}
	return -1
}
