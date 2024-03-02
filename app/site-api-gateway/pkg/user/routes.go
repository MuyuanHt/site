package user

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth"
	"site/app/site-api-gateway/pkg/user/routers"
	"site/conf"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine, c *conf.ServiceConf, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	rs := r.Group("/user")
	// 针对 rs 路由组下的路由使用中间件鉴权
	rs.Use(a.AuthRequired)

	find := rs.Group("/find")
	find.POST("/user", svc.FindOneUser)
}

func (svc *ServiceClient) FindOneUser(ctx *gin.Context) {
	routers.FindOneUser(ctx, svc.Client)
}
