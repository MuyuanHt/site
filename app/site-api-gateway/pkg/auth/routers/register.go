package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/auth/api"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
)

// Register 注册路由
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := api.RegisterReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.Register(context.Background(), &pb.RegisterReq{
		Account:  body.Account,
		Password: body.Password,
	})
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	data := &api.RegisterRespBody{
		Account: res.Account,
	}
	middleware.CheckStatusCode(ctx, int(res.Status), res.Error, data)
}
