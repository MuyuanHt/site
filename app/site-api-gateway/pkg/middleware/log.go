package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"site/common/logs"
	"site/protocol/shared"
	"time"
)

type LogMiddleware struct{}

// loggerBody 记录请求相关信息
type loggerBody struct {
	HandleCost int64  `json:"handle_cost"` // 处理耗时
	ReqUserId  int64  `json:"req_user_id"` // 发起人
	ReqPath    string `json:"req_path"`    // 请求路由
	ReqMethod  string `json:"req_method"`  // 请求方法
	ReqTime    string `json:"req_time"`    // 请求日期
	ReqBody    string `json:"req_body"`    // 请求体
	RespBody   string `json:"resp_body"`   // 响应体
}

// 自定义一个结构体，实现 gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

// InitLogMiddleware 初始化日志中间件
func InitLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

// RecordRequestLog 记录请求日志
func (m *LogMiddleware) RecordRequestLog(ctx *gin.Context) {
	body := &loggerBody{}
	defer func() {
		_ = logs.Logger.Sync() // 刷新日志
	}()
	start := time.Now()
	body.ReqTime = start.Format(time.DateTime) // 请求日期
	body.ReqPath = ctx.Request.URL.String()    // 请求路径
	body.ReqMethod = ctx.Request.Method        // 请求方式
	reqBody, err := ctx.GetRawData()           // 请求体
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			NewErrResult(http.StatusBadRequest, shared.ParseParamError),
		)
		return
	}
	// 将数据写回 ctx 防止后续读取时出现 EOF error
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer
	ctx.Next()

	body.ReqBody = string(reqBody)
	body.HandleCost = time.Since(start).Milliseconds() // 毫秒
	body.RespBody = writer.b.String()
	body.ReqUserId = ctx.GetInt64("uid")
	logStr, err := json.Marshal(body)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			NewErrResult(http.StatusBadRequest, shared.ParseParamError),
		)
		return
	}
	logs.Logger.Info(string(logStr))
}
