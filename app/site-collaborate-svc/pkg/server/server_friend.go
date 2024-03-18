package server

import (
	"context"
	"net/http"
	"site/app/site-collaborate-svc/pkg/service"
	"site/protocol/collaborate"
)

type CollaborateServer struct {
	FriendSvc service.FriendService
	collaborate.UnimplementedCollaborateServiceServer
}

// AddFriend 添加好友
func (s *CollaborateServer) AddFriend(ctx context.Context, req *collaborate.AddFriendReq) (*collaborate.AddFriendResp, error) {
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
