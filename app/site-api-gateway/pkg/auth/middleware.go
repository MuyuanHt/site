package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
	"site/protocol/shared"
	"strconv"
	"strings"
)

// MiddlewareConfig 继承 ServiceClient
type MiddlewareConfig struct {
	svc *ServiceClient
}

// InitAuthMiddleware 初始化认证中间件
func InitAuthMiddleware(svc *ServiceClient) MiddlewareConfig {
	return MiddlewareConfig{svc}
}

// AuthRequired 通过 token 认证并授权
func (c *MiddlewareConfig) AuthRequired(ctx *gin.Context) {
	// 获取请求头携带的认证信息
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		// gin 的中间件是栈调用的，后执行的流程先结束，当使用 Abort 时后续流程不会执行，但是本流程将继续执行结束
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			middleware.NewErrResult(http.StatusUnauthorized, shared.UnAuthorizedError),
		)
		return
	}

	// 验证 token
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			middleware.NewErrResult(http.StatusUnauthorized, shared.UnAuthorizedError),
		)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateReq{
		Token: token[1],
	})
	if err != nil || res.Msg.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			middleware.NewErrResult(http.StatusUnauthorized, shared.UnAuthorizedError),
		)
		return
	}

	ctx.Set("uid", res.Uid)
	// 调用后续的处理函数或调用下一个中间件
	ctx.Next()
}

// CheckTokenToUser 检验 token 与写操作对象是否一致，防止篡改其他对象数据
func (c *MiddlewareConfig) CheckTokenToUser(ctx *gin.Context) {
	ctxUid := ctx.GetInt64("uid")

	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			middleware.NewErrResult(http.StatusBadRequest, shared.ParseParamError),
		)
		return
	}

	var body map[string]interface{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			middleware.NewErrResult(http.StatusBadRequest, shared.ParseParamError),
		)
		return
	}

	if paramUid, ok := body["uid"]; ok {
		var uid int64
		uid, err = strconv.ParseInt(paramUid.(string), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				middleware.NewErrResult(http.StatusBadRequest, shared.ParseParamError),
			)
		}
		if uid != ctxUid {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				middleware.NewErrResult(http.StatusUnauthorized, shared.UnAuthorizedError),
			)
			return
		}
	}
	// 将数据写回 ctx 防止后续读取时出现 EOF error
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	ctx.Next()
}
