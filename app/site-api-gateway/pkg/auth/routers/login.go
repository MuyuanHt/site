package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/auth/api"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
	"site/protocol/shared"
)

// Login 登录路由
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := api.LoginReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.Login(ctx, &pb.LoginReq{
		Account:  body.Account,
		Password: body.Password,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	data := &api.LoginRespBody{
		Token: res.Token,
	}
	middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, data)
}
