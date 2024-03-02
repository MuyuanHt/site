package server

import (
	"context"
	"net/http"
	"site/app/site-user-svc/pkg/service"
	"site/common/tools"
	"site/protocol/shared"
	"site/protocol/user"
)

type UserServer struct {
	Svc service.UserService
	user.UnimplementedUserServiceServer
}

// FindOneUser 查询一个用户
func (s *UserServer) FindOneUser(ctx context.Context, req *user.FindOneUserReq) (*user.FindOneUserResp, error) {
	userInfo, err := s.Svc.FindOneUser(req.AccountType, req.Account)
	if userInfo == nil || err != nil {
		return &user.FindOneUserResp{
			Status: http.StatusNotFound,
			Error:  shared.CodeMessageIgnoreCode(shared.UserNotFound),
		}, nil
	}
	birthday, err := tools.TimeToInt64(userInfo.UserInfo.Birthday)
	if err != nil {
		birthday = 0
	}
	return &user.FindOneUserResp{
		Status: http.StatusOK,
		Data: &user.UserData{
			Uid:         userInfo.Uid,
			Phone:       userInfo.Phone,
			Email:       userInfo.Email,
			Username:    userInfo.UserInfo.Username,
			Age:         int32(userInfo.UserInfo.Age),
			Gender:      int32(userInfo.UserInfo.Gender),
			Region:      userInfo.UserInfo.Region,
			Icon:        userInfo.UserInfo.Icon,
			Description: userInfo.UserInfo.Description,
			Birthday:    birthday,
		},
	}, nil
}

// FindAccount 查询账户信息
func (s *UserServer) FindAccount(ctx context.Context, req *user.FindAccountReq) (*user.FindAccountResp, error) {
	userInfo, err := s.Svc.FindOneUser(req.AccountType, req.Account)
	if userInfo == nil || err != nil {
		return &user.FindAccountResp{
			Status: http.StatusNotFound,
			Error:  shared.CodeMessageIgnoreCode(shared.UserNotFound),
		}, nil
	}
	return &user.FindAccountResp{
		Status: http.StatusOK,
		Data: &user.AccountData{
			Phone:    userInfo.Phone,
			Email:    userInfo.Email,
			Password: userInfo.Password,
			Id:       userInfo.Uid, // TODO: 这里返回 uid 后续看会不会更改
		},
	}, nil
}

// CreateUser 创建用户
func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	userInfo, err := s.Svc.CreateOneUser(req.AccountType, req.Account, req.Password)
	if err != nil {
		return &user.CreateUserResp{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}
	birthday, err := tools.TimeToInt64(userInfo.UserInfo.Birthday)
	if err != nil {
		birthday = 0
	}
	return &user.CreateUserResp{
		Status: http.StatusOK,
		Data: &user.UserData{
			Uid:         userInfo.Uid,
			Phone:       userInfo.Phone,
			Email:       userInfo.Email,
			Username:    userInfo.UserInfo.Username,
			Age:         int32(userInfo.UserInfo.Age),
			Gender:      int32(userInfo.UserInfo.Gender),
			Region:      userInfo.UserInfo.Region,
			Icon:        userInfo.UserInfo.Icon,
			Description: userInfo.UserInfo.Description,
			Birthday:    birthday,
		},
	}, nil
}
