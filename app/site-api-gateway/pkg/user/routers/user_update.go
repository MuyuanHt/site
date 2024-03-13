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

// UpdatePassword 修改用户密码的路由函数
func UpdatePassword(ctx *gin.Context, c pb.UserServiceClient) {
	body := api.UpdatePasswordReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	uid, err := strconv.ParseInt(body.Uid, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.UpdatePassword(context.Background(), &pb.UpdatePasswordReq{
		Uid:         uid,
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	data := api.UpdatePasswordRespBody{
		Uid:     body.Uid,
		Success: res.Success,
	}
	middleware.OkWithData(ctx, data)
}

// UpdateUserInfo 修改用户信息的路由函数
func UpdateUserInfo(ctx *gin.Context, c pb.UserServiceClient) {
	body := api.UpdateUserInfoReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	uid, err := strconv.ParseInt(body.Uid, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	if !tools.IsGender(int(body.Gender)) {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.UpdateUserInfo(ctx, &pb.ChangeUserInfoReq{
		Data: &pb.UserData{
			Uid:         uid,
			Age:         body.Age,
			Gender:      body.Gender,
			Username:    body.Username,
			Region:      body.Region,
			Icon:        body.Icon,
			Description: body.Description,
			Birthday:    body.Birthday,
		},
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	data := api.UpdateUserInfoRespBody{
		Uid:         body.Uid,
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

// UpdateUserLimit 修改用户隐私权限的路由函数
func UpdateUserLimit(ctx *gin.Context, c pb.UserServiceClient) {
	body := api.UpdateUserLimitReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	uid, err := strconv.ParseInt(body.Uid, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	if !tools.IsSearchLimit(int(body.SearchLimit)) || !tools.IsVisitLimit(int(body.VisitLimit)) || !tools.IsAddLimit(int(body.AddLimit)) {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParamError))
		return
	}
	res, err := c.UpdateUserLimit(ctx, &pb.UpdateUserLimitReq{
		Uid: uid,
		Data: &pb.UserLimitData{
			SearchLimit: body.SearchLimit,
			VisitLimit:  body.VisitLimit,
			AddLimit:    body.AddLimit,
		},
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	data := api.UpdateUserLimitRespBody{
		SearchLimit: res.Data.SearchLimit,
		VisitLimit:  res.Data.VisitLimit,
		AddLimit:    res.Data.AddLimit,
	}
	middleware.OkWithData(ctx, data)
}
