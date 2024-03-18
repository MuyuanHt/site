package user

import (
	"github.com/gin-gonic/gin"
	"site/app/site-api-gateway/pkg/auth"
	"site/app/site-api-gateway/pkg/middleware"
	"site/app/site-api-gateway/pkg/user/routers"
	"site/conf"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine, c *conf.ServiceConf, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	logger := middleware.InitLogMiddleware()
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	rs := r.Group("/user")
	// 针对 rs 路由组下的路由使用中间件鉴权
	rs.Use(a.AuthRequired)
	// 记录请求日志中间件
	rs.Use(logger.RecordRequestLog)
	{
		// 查询用户信息路由组
		find := rs.Group("/find")
		{
			find.POST("/user", svc.FindOneUser)
			find.POST("/users", svc.FuzzyQueryUsers)
		}

		// 更新用户信息路由组
		update := rs.Group("/update")
		// TODO: 为方便开发期间测试，暂时屏蔽身份验证
		// update.Use(a.CheckTokenToUser)
		{
			update.POST("/password", svc.UpdatePassword)
			update.POST("/userInfo", svc.UpdateUserInfo)
			update.POST("/userLimit", svc.UpdateUserLimit)
		}
	}
}

func (svc *ServiceClient) FindOneUser(ctx *gin.Context) {
	routers.FindOneUser(ctx, svc.Client)
}

func (svc *ServiceClient) UpdatePassword(ctx *gin.Context) {
	routers.UpdatePassword(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUserInfo(ctx *gin.Context) {
	routers.UpdateUserInfo(ctx, svc.Client)
}

func (svc *ServiceClient) FuzzyQueryUsers(ctx *gin.Context) {
	routers.FuzzyQueryUsers(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUserLimit(ctx *gin.Context) {
	routers.UpdateUserLimit(ctx, svc.Client)
}
