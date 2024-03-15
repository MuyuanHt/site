package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/auth/api"
	"site/app/site-api-gateway/pkg/middleware"
	pb "site/protocol/auth"
	"site/protocol/shared"
)

// Register 注册路由
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := api.RegisterReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.Register(context.Background(), &pb.RegisterReq{
		Account:  body.Account,
		Password: body.Password,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	data := &api.RegisterRespBody{
		Account: res.Account,
	}
	middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, data)
}
