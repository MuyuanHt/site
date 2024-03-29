package shared

// This file use shell http_code.sh auto created
const (
	CodeNotFound            = 10101 // 状态码异常
	ParseParamError         = 10102 // 参数解析失败
	ServerError             = 10103 // 服务器内部出错
	ParamError              = 10104 // 参数错误
	UserNotFound            = 10201 // 用户不存在
	UserExists              = 10202 // 用户已存在
	PasswordError           = 10203 // 密码错误
	UserOrPasswordError     = 10204 // 用户名或密码有误
	CreateUidError          = 10205 // 生成用户序列号失败
	OldPasswordError        = 10206 // 原始密码有误
	InputsInconsistent      = 10207 // 两次输入不一致
	UpdatePasswordError     = 10208 // 修改密码失败
	UpdateUserInfoError     = 10209 // 修改用户信息失败
	CreateUserError         = 10210 // 创建用户失败
	OldPwdSameAsNewPwd      = 10211 // 新密码与原始密码相同
	FuzzyQueryError         = 10212 // 未查询到相关用户
	QueryParamNilError      = 10213 // 查询参数为空
	UpdateUserRelationError = 10214 // 修改用户关系信息失败
	RecordNotFound          = 10215 // 记录不存在
	AccountInvalid          = 10401 // 账号无效
	AccountTypeInvalid      = 10402 // 账号类型无效
	UnAuthorizedError       = 10403 // 账号未授权
	PasswordInvalid         = 10404 // 密码不合法，请输入6~16位且至少包含数字与字母
	GenerateTokenError      = 10411 // 生成令牌失败
	TokenParsingError       = 10412 // 令牌解析失败
	TokenExpiration         = 10413 // 令牌已过期
	TokenInvalid            = 10414 // 令牌无效
	ParseKeyError           = 10415 // 解析密钥失败
	TimeInvalid             = 10451 // 时间无效
	ConfNotFound            = 10452 // 无配置信息
	HashPwdError            = 10453 // 生成密码失败
	AddError                = 10601 // 添加用户或加入群组失败
	DeleteError             = 10602 // 删除用户或退出群组失败
	TopError                = 10603 // 置顶会话失败
	DelTopError             = 10604 // 取消置顶会话失败
	BlackError              = 10605 // 拉黑失败
	DelBlackError           = 10606 // 取消拉黑失败
	CreateGroupError        = 10607 // 创建群组失败
	DeleteGroupError        = 10608 // 解散群组失败
	ExistsFriendOrGroup     = 10609 // 已经添加好友或加入群组
	AddYourselfError        = 10610 // 无法添加自己为好友
	TagExists               = 10630 // 标签已存在或名称重复
	TagRelationExist        = 10631 // 已与标签关联
	TagRelationNotExist     = 10632 // 未与标签关联
)

var resultCodeText = map[int]string{
	CodeNotFound:            "状态码异常",
	ParseParamError:         "参数解析失败",
	ServerError:             "服务器内部出错",
	ParamError:              "参数错误",
	UserNotFound:            "用户不存在",
	UserExists:              "用户已存在",
	PasswordError:           "密码错误",
	UserOrPasswordError:     "用户名或密码有误",
	CreateUidError:          "生成用户序列号失败",
	OldPasswordError:        "原始密码有误",
	InputsInconsistent:      "两次输入不一致",
	UpdatePasswordError:     "修改密码失败",
	UpdateUserInfoError:     "修改用户信息失败",
	CreateUserError:         "创建用户失败",
	OldPwdSameAsNewPwd:      "新密码与原始密码相同",
	FuzzyQueryError:         "未查询到相关用户",
	QueryParamNilError:      "查询参数为空",
	UpdateUserRelationError: "修改用户关系信息失败",
	RecordNotFound:          "记录不存在",
	AccountInvalid:          "账号无效",
	AccountTypeInvalid:      "账号类型无效",
	UnAuthorizedError:       "账号未授权",
	PasswordInvalid:         "密码不合法，请输入6~16位且至少包含数字与字母",
	GenerateTokenError:      "生成令牌失败",
	TokenParsingError:       "令牌解析失败",
	TokenExpiration:         "令牌已过期",
	TokenInvalid:            "令牌无效",
	ParseKeyError:           "解析密钥失败",
	TimeInvalid:             "时间无效",
	ConfNotFound:            "无配置信息",
	HashPwdError:            "生成密码失败",
	AddError:                "添加用户或加入群组失败",
	DeleteError:             "删除用户或退出群组失败",
	TopError:                "置顶会话失败",
	DelTopError:             "取消置顶会话失败",
	BlackError:              "拉黑失败",
	DelBlackError:           "取消拉黑失败",
	CreateGroupError:        "创建群组失败",
	DeleteGroupError:        "解散群组失败",
	ExistsFriendOrGroup:     "已经添加好友或加入群组",
	AddYourselfError:        "无法添加自己为好友",
	TagExists:               "标签已存在或名称重复",
	TagRelationExist:        "已与标签关联",
	TagRelationNotExist:     "未与标签关联",
}

// CodeMessage 获取 code 对应的 message
func CodeMessage(code int) (string, bool) {
	message, ok := resultCodeText[code]
	return message, ok
}

// CodeMessageIgnoreCode 获取 code 对应的 message 未查询到状态码时返回指定的状态码异常错误
func CodeMessageIgnoreCode(code int) string {
	message, ok := resultCodeText[code]
	if !ok {
		return resultCodeText[CodeNotFound]
	}
	return message
}
