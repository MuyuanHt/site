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
	userId, err := strconv.ParseInt(body.Uid, 10, 64)
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
		Uid:      body.Uid,
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
	userId, err := strconv.ParseInt(body.Uid, 10, 64)
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
		Uid:      body.Uid,
		FriendId: body.FriendId,
	}
	middleware.OkWithData(ctx, data)
}

// UpdateFriend 更新好友信息
func UpdateFriend(ctx *gin.Context, c pb.CollaborateServiceClient) {
	body := api.UpdateFriendInfoReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	userId, err := strconv.ParseInt(body.Uid, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	friendId, err := strconv.ParseInt(body.FriendId, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.UpdateFriend(ctx, &pb.UpdateFriendReq{
		Data: &pb.FriendRecordData{
			UserId:   userId,
			FriendId: friendId,
			IsTop:    body.IsTop,
			IsBlack:  body.IsBlack,
			Label:    body.Label,
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
	data := &api.UpdateFriendInfoRespBody{
		Uid:      body.Uid,
		FriendId: body.FriendId,
		IsTop:    res.Data.IsTop,
		IsBlack:  res.Data.IsBlack,
		Label:    res.Data.Label,
	}
	middleware.OkWithData(ctx, data)
}

// FindAllFriends 更新好友信息
func FindAllFriends(ctx *gin.Context, c pb.CollaborateServiceClient) {
	option := ctx.Param("option") // 用于筛选置顶或拉黑等条件
	body := api.FindAllFriendsReqBody{}
	if err := ctx.BindJSON(&body); err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	uid, err := strconv.ParseInt(body.Uid, 10, 64)
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ParseParamError))
		return
	}
	res, err := c.FindAllFriends(ctx, &pb.FindAllFriendsReq{
		Uid: uid,
	})
	if err != nil {
		middleware.Fail(ctx, http.StatusBadRequest, shared.CodeMessageIgnoreCode(shared.ServerError))
		return
	}
	if !tools.CompareInt32Int(res.Msg.Status, http.StatusOK) {
		middleware.CheckStatusCode(ctx, int(res.Msg.Status), res.Msg.Error, nil)
		return
	}
	friends := make([]*api.FriendData, 0, len(res.Data))
	if option != "top" && option != "black" {
		for _, v := range res.Data {
			friends = append(friends, &api.FriendData{
				FriendId: strconv.FormatInt(v.FriendId, 10),
				IsTop:    v.IsTop,
				IsBlack:  v.IsBlack,
				Label:    v.Label,
			})
		}
	} else if option == "top" {
		for _, v := range res.Data {
			if v.IsTop {
				friends = append(friends, &api.FriendData{
					FriendId: strconv.FormatInt(v.FriendId, 10),
					IsTop:    v.IsTop,
					IsBlack:  v.IsBlack,
					Label:    v.Label,
				})
			}
		}
	} else if option == "black" {
		for _, v := range res.Data {
			if v.IsBlack {
				friends = append(friends, &api.FriendData{
					FriendId: strconv.FormatInt(v.FriendId, 10),
					IsTop:    v.IsTop,
					IsBlack:  v.IsBlack,
					Label:    v.Label,
				})
			}
		}
	}
	data := api.FindAllFriendsRespBody{
		FriendNum: len(friends),
		Friends:   friends,
	}
	middleware.OkWithData(ctx, data)
}
