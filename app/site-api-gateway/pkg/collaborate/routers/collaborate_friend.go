package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site/app/site-api-gateway/pkg/collaborate/api"
	"site/app/site-api-gateway/pkg/middleware"
	"site/common/tools"
	pb "site/protocol/collaborate"
	"site/protocol/shared"
	"strconv"
)

// AddFriend 添加好友
func AddFriend(ctx *gin.Context, c pb.CollaborateServiceClient) {
	body := api.AddFriendReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	userId, err := strconv.ParseInt(body.UserId, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	friendId, err := strconv.ParseInt(body.FriendId, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.AddFriend(ctx, &pb.AddFriendReq{
		UserId:      userId,
		FriendId:    friendId,
		UserLabel:   body.UserLabel,
		FriendLabel: body.FriendLabel,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	data := &api.AddFriendRespBody{
		UserId:   body.UserId,
		FriendId: body.FriendId,
	}
	middleware.OkWithData(ctx, data)
}

// DeleteFriend 添加好友
func DeleteFriend(ctx *gin.Context, c pb.CollaborateServiceClient) {
	body := api.DelFriendReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	userId, err := strconv.ParseInt(body.UserId, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	friendId, err := strconv.ParseInt(body.FriendId, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.DeleteFriend(ctx, &pb.DeleteFriendReq{
		UserId:   userId,
		FriendId: friendId,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	data := &api.DelFriendRespBody{
		UserId:   body.UserId,
		FriendId: body.FriendId,
	}
	middleware.OkWithData(ctx, data)
}
