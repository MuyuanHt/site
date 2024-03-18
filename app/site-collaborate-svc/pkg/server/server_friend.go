package server

import (
	"context"
	"net/http"
	"site/app/site-collaborate-svc/pkg/models"
	"site/app/site-collaborate-svc/pkg/service"
	"site/protocol/collaborate"
	"site/protocol/shared"
)

type CollaborateServer struct {
	FriendSvc service.FriendService
	collaborate.UnimplementedCollaborateServiceServer
}

// AddFriend 添加好友
func (s *CollaborateServer) AddFriend(ctx context.Context, req *collaborate.AddFriendReq) (*collaborate.AddFriendResp, error) {
	if req == nil {
		return &collaborate.AddFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	err := s.FriendSvc.AddFriend(req.UserId, req.FriendId, req.UserLabel, req.FriendLabel)
	if err != nil {
		return &collaborate.AddFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &collaborate.AddFriendResp{
		Msg: &collaborate.RetMsg{
			Status: http.StatusOK,
		},
	}, nil
}

// DeleteFriend 删除好友
func (s *CollaborateServer) DeleteFriend(ctx context.Context, req *collaborate.DeleteFriendReq) (*collaborate.DeleteFriendResp, error) {
	if req == nil {
		return &collaborate.DeleteFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	err := s.FriendSvc.DeleteFriend(req.UserId, req.FriendId)
	if err != nil {
		return &collaborate.DeleteFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &collaborate.DeleteFriendResp{
		Msg: &collaborate.RetMsg{
			Status: http.StatusOK,
		},
	}, nil
}

// UpdateFriend 更新好友信息
func (s *CollaborateServer) UpdateFriend(ctx context.Context, req *collaborate.UpdateFriendReq) (*collaborate.UpdateFriendResp, error) {
	if req == nil {
		return &collaborate.UpdateFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	info := &models.UserFriend{
		UserId:   req.Data.UserId,
		FriendId: req.Data.FriendId,
		IsTop:    req.Data.IsTop,
		IsBlack:  req.Data.IsBlack,
		Label:    req.Data.Label,
	}
	retInfo, err := s.FriendSvc.UpdateFriendInfo(info)
	if err != nil {
		return &collaborate.UpdateFriendResp{
			Msg: &collaborate.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &collaborate.UpdateFriendResp{
		Msg: &collaborate.RetMsg{
			Status: http.StatusOK,
		},
		Data: &collaborate.FriendRecordData{
			UserId:   retInfo.UserId,
			FriendId: retInfo.FriendId,
			IsTop:    retInfo.IsTop,
			IsBlack:  retInfo.IsBlack,
			Label:    retInfo.Label,
		},
	}, nil
}
