package server

import (
	"context"
	"net/http"
	"site/app/site-user-svc/pkg/models"
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
	userInfo, err := s.Svc.FindOneUserByAccount(req.AccountType, req.Account)
	if userInfo == nil || err != nil {
		return &user.FindOneUserResp{
			Msg: &user.RetMsg{
				Status: http.StatusNotFound,
				Error:  shared.CodeMessageIgnoreCode(shared.UserNotFound),
			},
		}, nil
	}
	birthday, err := tools.TimeToInt64(userInfo.UserInfo.Birthday)
	if err != nil {
		birthday = 0
	}
	return &user.FindOneUserResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
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
	userInfo, err := s.Svc.FindOneUserByAccount(req.AccountType, req.Account)
	if userInfo == nil || err != nil {
		return &user.FindAccountResp{
			Msg: &user.RetMsg{
				Status: http.StatusNotFound,
				Error:  shared.CodeMessageIgnoreCode(shared.UserNotFound),
			},
		}, nil
	}
	return &user.FindAccountResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
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
			Msg: &user.RetMsg{
				Status: http.StatusConflict,
				Error:  err.Error(),
			},
		}, nil
	}
	birthday, err := tools.TimeToInt64(userInfo.UserInfo.Birthday)
	if err != nil {
		birthday = 0
	}
	return &user.CreateUserResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
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

// UpdatePassword 修改密码
func (s *UserServer) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (*user.UpdatePasswordResp, error) {
	ok, err := s.Svc.UpdatePassword(req.Uid, req.OldPassword, req.NewPassword)
	if err != nil {
		return &user.UpdatePasswordResp{
			Msg: &user.RetMsg{
				Status: http.StatusConflict,
				Error:  err.Error(),
			},
		}, nil
	}
	return &user.UpdatePasswordResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
		Success: ok,
	}, nil
}

// UpdateUserInfo 修改用户信息
func (s *UserServer) UpdateUserInfo(ctx context.Context, req *user.ChangeUserInfoReq) (*user.ChangeUserInfoResp, error) {
	birthdayTime, err := tools.Int64ToTime(req.Data.Birthday)
	if err != nil {
		return &user.ChangeUserInfoResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	info := &models.UserInfo{
		Age:         int8(req.Data.Age),
		Gender:      int8(req.Data.Gender),
		Username:    req.Data.Username,
		Region:      req.Data.Region,
		Icon:        req.Data.Icon,
		Description: req.Data.Description,
		Birthday:    tools.FormatBirthdayTime(birthdayTime),
	}
	userInfo, err := s.Svc.UpdateUserInfo(req.Data.Uid, info)
	if err != nil {
		return &user.ChangeUserInfoResp{
			Msg: &user.RetMsg{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			},
		}, nil
	}
	birthdayInt, err := tools.TimeToInt64(userInfo.UserInfo.Birthday)
	if err != nil {
		birthdayInt = 0
	}
	return &user.ChangeUserInfoResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
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
			Birthday:    birthdayInt,
		},
	}, nil
}

// FuzzyQueryUsers 模糊查询用户信息
func (s *UserServer) FuzzyQueryUsers(ctx context.Context, req *user.FuzzyQueryUsersReq) (*user.FuzzyQueryUsersResp, error) {
	accounts, err := s.Svc.FuzzyQueryUsers(req.Param)
	if err != nil {
		return &user.FuzzyQueryUsersResp{
			Msg: &user.RetMsg{
				Status: http.StatusNotFound,
				Error:  shared.CodeMessageIgnoreCode(shared.FuzzyQueryError),
			},
		}, nil
	}
	data := make([]*user.IgnoreUserData, 0, len(accounts))
	for _, ac := range accounts {
		data = append(data, &user.IgnoreUserData{
			Phone:    ac.Phone,
			Email:    ac.Email,
			Username: ac.Username,
			Icon:     ac.Icon,
		})
	}
	return &user.FuzzyQueryUsersResp{
		Msg: &user.RetMsg{
			Status: http.StatusOK,
		},
		Data: data,
	}, nil
}
