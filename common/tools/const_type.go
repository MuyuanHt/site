package tools

import consts "site/protocol/shared"

// IsGender 判断是否是性别
func IsGender(num int) bool {
	return num >= consts.GenderTypeNull && num <= consts.GenderTypeWoman
}

// IsSearchLimit 判断是否是搜索限制
func IsSearchLimit(num int) bool {
	return num >= consts.SearchLimitY && num <= consts.SearchLimitN
}

// IsVisitLimit 判断是否是查看限制
func IsVisitLimit(num int) bool {
	return num >= consts.VisitLimitY && num <= consts.VisitLimitFriendY
}

// IsAddLimit 判断是否是添加限制
func IsAddLimit(num int) bool {
	return num >= consts.AddLimitY && num <= consts.AddLimitAgree
}
