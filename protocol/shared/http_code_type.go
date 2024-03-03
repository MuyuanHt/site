package shared

// This file use shell http_code.sh auto created
const (
	CodeNotFound        = 10101 // 状态码异常
	ParseParamError     = 10102 // 参数解析失败
	UserNotFound        = 10201 // 用户不存在
	UserExists          = 10202 // 用户已存在
	PasswordError       = 10203 // 密码错误
	UserOrPasswordError = 10204 // 用户名或密码有误
	CreateUidError      = 10205 // 生成用户序列号失败
	AccountInvalid      = 10401 // 账号无效
	AccountTypeInvalid  = 10402 // 账号类型无效
	UnAuthorizedError   = 10403 // 账号未授权
	GenerateTokenError  = 10411 // 生成令牌失败
	TokenParsingError   = 10412 // 令牌解析失败
	TokenExpiration     = 10413 // 令牌已过期
	TokenInvalid        = 10414 // 令牌无效
	ParseKeyError       = 10415 // 解析密钥失败
	TimeInvalid         = 10451 // 时间无效
	ConfNotFound        = 10452 // 无配置信息
)

var resultCodeText = map[int]string{
	CodeNotFound:        "状态码异常",
	ParseParamError:     "参数解析失败",
	UserNotFound:        "用户不存在",
	UserExists:          "用户已存在",
	PasswordError:       "密码错误",
	UserOrPasswordError: "用户名或密码有误",
	CreateUidError:      "生成用户序列号失败",
	AccountInvalid:      "账号无效",
	AccountTypeInvalid:  "账号类型无效",
	UnAuthorizedError:   "账号未授权",
	GenerateTokenError:  "生成令牌失败",
	TokenParsingError:   "令牌解析失败",
	TokenExpiration:     "令牌已过期",
	TokenInvalid:        "令牌无效",
	ParseKeyError:       "解析密钥失败",
	TimeInvalid:         "时间无效",
	ConfNotFound:        "无配置信息",
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
