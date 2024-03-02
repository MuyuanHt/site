package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/middleware"
	"site/app/site-api-gateway/pkg/user/api"
	"site/common/tools"
	"site/protocol/shared"
	pb "site/protocol/user"
	"strconv"
)

// FindOneUser 获取单个用户信息的路由函数
func FindOneUser(ctx *gin.Context, c pb.UserServiceClient) {
	body := api.FindOneUserReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	at, err := strconv.ParseInt(body.AccountType, 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.FindOneUser(context.Background(), &pb.FindOneUserReq{
		AccountType: int32(at),
		Account:     body.Account,
	})
	if !tools.CompareInt32Int(res.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Status), res.Error, res.Data)
		return
	}
	if res.Data == nil {
		middleware.FailWithMsg(ctx, shared.UserNotFound)
		return
	}
	data := api.FindOneUserRespBody{
		Uid:         res.Data.Uid,
		Phone:       res.Data.Phone,
		Email:       res.Data.Email,
		Username:    res.Data.Username,
		Age:         res.Data.Age,
		Gender:      res.Data.Gender,
		Region:      res.Data.Region,
		Icon:        res.Data.Icon,
		Description: res.Data.Description,
		Birthday:    res.Data.Birthday,
	}
	middleware.OkWithData(ctx, data)
}
