package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site/protocol/shared"
)

// Result 服务器统一响应的 json 格式
type Result struct {
	Code int         `json:"code"`    // 状态码
	Msg  string      `json:"message"` // 响应消息
	Data interface{} `json:"data"`    // 响应数据
}

// NewResult 创建返回对象
func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Data: data}
}

// NewErrResult 创建异常时返回对象 code 为返回状态码 ownerCode 用于查找自定义错误
func NewErrResult(code int, ownerCode int) *Result {
	return &Result{Code: code, Msg: shared.CodeMessageIgnoreCode(ownerCode)}
}

// NewResultOk 创建成功返回信息
func NewResultOk(c *gin.Context, r *Result) {
	c.JSON(http.StatusOK, r)
}

// NewResultError 创建错误返回信息
func NewResultError(c *gin.Context, r *Result) {
	c.JSON(r.Code, r)
}

// Ok 成功 无返回
func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

// OkWithData 成功 并返回数据
func OkWithData(c *gin.Context, data interface{}) {
	NewResultOk(c, NewResult(http.StatusOK, "Success", data))
}

// OkWithMsg 成功 只返回信息
func OkWithMsg(c *gin.Context, msg string) {
	NewResultOk(c, NewResult(http.StatusOK, msg, nil))
}

// Fail 失败 返回异常信息
func Fail(c *gin.Context, code int, msg string) {
	NewResultError(c, NewResult(code, msg, nil))
}

// FailWithMsg 失败 返回自定义异常信息
func FailWithMsg(c *gin.Context, code int) {
	NewResultError(c, NewResult(code, shared.CodeMessageIgnoreCode(code), nil))
}

// CheckStatusCode 校验得到协议消息的状态码并返回对应消息
func CheckStatusCode(c *gin.Context, code int, err string, data interface{}) {
	if code != http.StatusOK {
		Fail(c, code, err)
		return
	}
	OkWithData(c, data)
}
