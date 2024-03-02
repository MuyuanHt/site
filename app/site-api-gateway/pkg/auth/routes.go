package auth

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth/routers"
	"site/conf"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine, c *conf.ServiceConf) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	rs := r.Group("/auth")
	rs.POST("/register", svc.Register)
	rs.POST("/login", svc.Login)
	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routers.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routers.Login(ctx, svc.Client)
}
