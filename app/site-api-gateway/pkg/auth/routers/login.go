package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/auth/api"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
)

// Login 登录路由
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := api.LoginReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.Login(ctx, &pb.LoginReq{
		Account:  body.Account,
		Password: body.Password,
	})
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	data := &api.LoginRespBody{
		Token: res.Token,
	}
	middleware.CheckStatusCode(ctx, int(res.Status), res.Error, data)
}
