package collaborate

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth"
	"site/app/site-api-gateway/pkg/collaborate/routers"
	"site/app/site-api-gateway/pkg/middleware"
	"site/conf"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine, c *conf.ServiceConf, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	logger := middleware.InitLogMiddleware()
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	rs := r.Group("/team")
	// 针对 rs 路由组下的路由使用中间件鉴权
	rs.Use(a.AuthRequired)
	// 记录请求日志中间件
	rs.Use(logger.RecordRequestLog)
	// 用户关系路由组
	friend := rs.Group("/friend")
	friend.POST("/add", svc.AddFriend)
	friend.POST("/delete", svc.DeleteFriend)
}

func (svc *ServiceClient) AddFriend(ctx *gin.Context) {
	routers.AddFriend(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteFriend(ctx *gin.Context) {
	routers.DeleteFriend(ctx, svc.Client)
}
