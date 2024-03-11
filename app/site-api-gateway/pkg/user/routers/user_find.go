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
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	at, err := strconv.ParseInt(body.AccountType, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.FindOneUser(context.Background(), &pb.FindOneUserReq{
		AccountType: int32(at),
		Account:     body.Account,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	if res.Data == nil {
		middleware.FailWithMsg(ctx, shared.UserNotFound)
		return
	}
	data := api.FindOneUserRespBody{
		Uid:         strconv.FormatInt(res.Data.Uid, 10),
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

// FuzzyQueryUsers 批量查询用户信息的路由函数
func FuzzyQueryUsers(ctx *gin.Context, c pb.UserServiceClient) {
	body := api.FuzzyQueryUsersReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.FuzzyQueryUsers(ctx, &pb.FuzzyQueryUsersReq{
		Param: body.Param,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	if res.Data == nil {
		middleware.FailWithMsg(ctx, shared.UserNotFound)
		return
	}
	users := make([]*api.FuzzyQueryUsersData, 0, len(res.Data))
	for _, v := range res.Data {
		users = append(users, &api.FuzzyQueryUsersData{
			Phone:    v.Phone,
			Email:    v.Email,
			Username: v.Username,
			Icon:     v.Icon,
		})
	}
	data := api.FuzzyQueryUsersRespBody{
		Data: users,
	}
	middleware.OkWithData(ctx, data)
}
