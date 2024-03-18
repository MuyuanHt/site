package server

import (
	"context"
	"net/http"
	"site/app/site-user-svc/pkg/models"
	"site/protocol/shared"
	"site/protocol/user"
)

// UpdateUserFriendNum 添加好友
func (s *UserServer) UpdateUserFriendNum(ctx context.Context, req *user.UpdateUserFriendNumReq) (*user.UpdateUserFriendNumResp, error) {
	if req == nil {
		return &user.UpdateUserFriendNumResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	err := s.Svc.AddFriendByUid(req.UserId, req.FriendId, req.DelFlag)
	if err != nil {
		return &user.UpdateUserFriendNumResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &user.UpdateUserFriendNumResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
	}, nil
}

// UpdateUserLimit 修改用户隐私权限
func (s *UserServer) UpdateUserLimit(ctx context.Context, req *user.UpdateUserLimitReq) (*user.UpdateUserLimitResp, error) {
	if req == nil {
		return &user.UpdateUserLimitResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	limit := &models.UserLimit{
		SearchLimit: int8(req.Data.SearchLimit),
		VisitLimit:  int8(req.Data.VisitLimit),
		AddLimit:    int8(req.Data.AddLimit),
	}
	account, err := s.Svc.UpdateUserLimit(req.Uid, limit)
	if err != nil {
		return &user.UpdateUserLimitResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &user.UpdateUserLimitResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
		Data: &user.UserLimitData{
			SearchLimit: int32(account.UserRelation.SearchLimit),
			VisitLimit:  int32(account.UserRelation.VisitLimit),
			AddLimit:    int32(account.UserRelation.AddLimit),
		},
	}, nil
}

// UpdateUserRelation 修改用户关系数量
func (s *UserServer) UpdateUserRelation(ctx context.Context, req *user.UpdateUserRelationReq) (*user.UpdateUserRelationResp, error) {
	if req == nil {
		return &user.UpdateUserRelationResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  shared.CodeMessageIgnoreCode(shared.ParamError),
			},
		}, nil
	}
	deltaRelation := &models.UserRelationNum{
		FriendNum:      req.Data.FriendNum,
		TopFriendNum:   req.Data.TopFriendNum,
		BlackFriendNum: req.Data.BlackFriendNum,
		GroupNum:       req.Data.GroupNum,
		OwnerGroupNum:  req.Data.OwnerGroupNum,
		AdminGroupNum:  req.Data.AdminGroupNum,
		TopGroupNum:    req.Data.TopGroupNum,
		BlackGroupNum:  req.Data.BlackGroupNum,
	}
	account, err := s.Svc.UpdateUserRelationNum(req.Uid, deltaRelation)
	if err != nil {
		return &user.UpdateUserRelationResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	return &user.UpdateUserRelationResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
		Data: &user.UserRelationData{
			FriendNum:      account.UserRelation.FriendNum,
			TopFriendNum:   account.UserRelation.TopFriendNum,
			BlackFriendNum: account.UserRelation.BlackFriendNum,
			GroupNum:       account.UserRelation.GroupNum,
			OwnerGroupNum:  account.UserRelation.OwnerGroupNum,
			AdminGroupNum:  account.UserRelation.AdminGroupNum,
			TopGroupNum:    account.UserRelation.TopGroupNum,
			BlackGroupNum:  account.UserRelation.BlackGroupNum,
		},
	}, nil
}
