package tools

import consts "site/protocol/shared"

// IsGender 判断是否是性别
func IsGender(num int) bool {
	return num >= consts.GenderTypeNull && num <= consts.GenderTypeWoman
}
