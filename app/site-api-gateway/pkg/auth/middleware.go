package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
	"site/protocol/shared"
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
