package server

import (
	"context"
	"net/http"
	"site/app/site-auth-svc/pkg/service"
	"site/protocol/auth"
	"site/protocol/shared"
)

// TODO: 增加 jwt 并使用中间件处理

type AuthServer struct {
	Svc service.AuthService
	auth.UnimplementedAuthServiceServer
}

// Register 用户注册
func (s *AuthServer) Register(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error) {
	err := s.Svc.Register(req.Account, req.Password)
	if err != nil {
		return &auth.RegisterResp{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}
	return &auth.RegisterResp{
		Status:  http.StatusOK,
		Account: req.Account,
	}, nil
}

// Login 用户登录
func (s *AuthServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	token, err := s.Svc.Login(req.Account, req.Password)
	if err != nil {
		return &auth.LoginResp{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	return &auth.LoginResp{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

// Validate 校验 token 有效性
func (s *AuthServer) Validate(ctx context.Context, req *auth.ValidateReq) (*auth.ValidateResp, error) {
	uid, err := s.Svc.Validate(req.Token)
	if err != nil {
		return &auth.ValidateResp{
			Status: http.StatusUnauthorized,
			Error:  shared.CodeMessageIgnoreCode(shared.TokenInvalid),
		}, nil
	}

	return &auth.ValidateResp{
		Status: http.StatusOK,
		Uid:    uid,
	}, nil
}
