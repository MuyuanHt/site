package service

import (
	"errors"
	"site/app/site-auth-svc/pkg/client"
	"site/app/site-auth-svc/pkg/utils"
	"site/common/tools"
	"site/protocol/shared"
	"strconv"
)

type AuthService struct {
	UserSvc client.UserServiceClient
	Jwt     *utils.JwtWrapper
}

// Register 注册账户
func (s *AuthService) Register(account string, password string) error {
	accountType := tools.GetAccountType(account)
	if accountType == 0 {
		return errors.New(shared.CodeMessageIgnoreCode(shared.AccountTypeInvalid))
	}
	if !tools.CheckPwdRegexp(password) {
		return errors.New(shared.CodeMessageIgnoreCode(shared.PasswordInvalid))
	}
	encryptPwd, err := tools.HashPassword(password)
	if err != nil {
		return err
	}
	resp, err := s.UserSvc.CreateUser(accountType, account, encryptPwd)
	if err != nil {
		return err
	}
	if resp.Msg.Error != "" || resp.Msg.Status != 200 {
		return errors.New(resp.Msg.Error)
	}
	return nil
}

// Login 登录账户 返回的 string 类型为 token
func (s *AuthService) Login(account string, password string) (string, error) {
	accountType := tools.GetAccountType(account)
	if accountType < 0 {
		return "", errors.New(shared.CodeMessageIgnoreCode(shared.AccountTypeInvalid))
	}

	resp, err := s.UserSvc.FindAccount(accountType, account)
	if err != nil {
		return "", err
	}
	if resp.Msg.Error != "" || resp.Msg.Status != 200 {
		return "", errors.New(resp.Msg.Error)
	}

	hashPwd := resp.Data.Password
	if !tools.CheckHashPassword(hashPwd, password) {
		return "", errors.New(shared.CodeMessageIgnoreCode(shared.UserOrPasswordError))
	}
	token, err := s.Jwt.GenerateTokenUsingRS256(resp.Data.Id, account)
	if err != nil {
		return "", errors.New(shared.CodeMessageIgnoreCode(shared.GenerateTokenError))
	}
	loginTimeResp, err := s.UserSvc.UpdateLastLoginTime(resp.Data.Id)
	if err != nil {
		return "", err
	}
	if loginTimeResp.Msg.Error != "" || loginTimeResp.Msg.Status != 200 {
		return "", errors.New(resp.Msg.Error)
	}
	return token, nil
}

// Validate 校验 token 有效性
func (s *AuthService) Validate(token string) (int64, error) {
	claims, err := s.Jwt.ParseTokenRs256(token)
	if err != nil {
		return -1, err
	}
	// 去数据库再校验一下有无此对象，防止由于对象删除等原因出现错误
	uidStr := strconv.Itoa(int(claims.UserId))
	resp, err := s.UserSvc.FindAccount(shared.AccountTypeUid, uidStr)
	if err != nil {
		return -1, err
	}
	if resp.Msg.Error != "" || resp.Msg.Status != 200 {
		return -1, errors.New(resp.Msg.Error)
	}

	return claims.UserId, nil
}
