package service

import (
	"errors"
	"fmt"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"site/app/site-user-svc/pkg/dao"
	"site/app/site-user-svc/pkg/models"
	"site/app/site-user-svc/pkg/utils"
	"site/common/tools"
	"site/conf"
	"site/protocol/shared"
	"time"
)

// TODO: proto 传输时间格式不使用时间戳

type UserService struct {
	D *dao.Dao
	S *snowflake.Snowflake
}

// InitUserService 初始化 UserService
func InitUserService(d *dao.Dao, s *snowflake.Snowflake) UserService {
	return UserService{
		D: d,
		S: s,
	}
}

// checkZeroValue 检查是零值则返回第二个参数 用于防止修改时传入参数为零值影响原数据
func checkZeroValue(input interface{}, replace interface{}) interface{} {
	if tools.IsZeroValue(input) {
		return replace
	}
	return input
}

// FindOneUserByAccount 查询单个用户
func (s *UserService) FindOneUserByAccount(accountType int32, account string) (*models.Account, error) {
	info, err := s.D.FindOneUserByAccount(int(accountType), account)
	if info != nil {
		return info, nil
	}
	return nil, err
}

// FindOneUserByUid 查询单个用户
func (s *UserService) FindOneUserByUid(uid int64) (*models.Account, error) {
	info, err := s.D.FindOneAccountByUid(uid)
	if info != nil {
		return info, nil
	}
	return nil, err
}

// CreateOneUser 创建单个用户
func (s *UserService) CreateOneUser(accountType int32, account string, password string) (*models.Account, error) {
	result, err := s.FindOneUserByAccount(accountType, account)
	if err == nil || result != nil {
		return result, errors.New(shared.CodeMessageIgnoreCode(shared.UserExists))
	}
	// TODO: 生成uid, 默认头像
	icon, err := conf.GetConfigParam("defaultIconUrl")
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.CreateUserError))
	}
	description, err := conf.GetConfigParam("defaultDescription")
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.CreateUserError))
	}
	info := &models.Account{
		Uid:      utils.GetNextSnowflakeId(s.S),
		Password: password,
		UserInfo: models.UserInfo{
			Username:      fmt.Sprintf("user_%v", account),
			Gender:        shared.GenderTypeNull,
			Icon:          icon,
			Description:   description,
			LastLoginTime: time.Now(),
		},
		UserRelation: models.UserRelation{
			SearchLimit: shared.SearchLimitY,
			VisitLimit:  shared.VisitLimitN,
			AddLimit:    shared.AddLimitAgree,
		},
	}
	switch accountType {
	case shared.AccountTypePhone:
		info.Phone = account
	case shared.AccountTypeEmail:
		info.Email = account
	default:
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.AccountInvalid))
	}
	err = s.D.CreateAccount(info)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.CreateUserError))
	}
	return info, nil
}

// UpdatePassword 修改密码
func (s *UserService) UpdatePassword(uid int64, oldPwd string, newPwd string) (bool, error) {
	result, err := s.FindOneUserByUid(uid)
	if err != nil {
		return false, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	if !tools.CheckHashPassword(result.Password, oldPwd) {
		return false, errors.New(shared.CodeMessageIgnoreCode(shared.OldPasswordError))
	}
	if !tools.CheckPwdRegexp(newPwd) {
		return false, errors.New(shared.CodeMessageIgnoreCode(shared.PasswordInvalid))
	}
	if tools.CheckHashPassword(result.Password, newPwd) {
		return false, errors.New(shared.CodeMessageIgnoreCode(shared.OldPwdSameAsNewPwd))
	}
	newPassword, err := tools.HashPassword(newPwd)
	if err != nil {
		return false, err
	}
	err = s.D.UpdatePasswordByUid(uid, newPassword)
	if err != nil {
		return false, errors.New(shared.CodeMessageIgnoreCode(shared.UpdatePasswordError))
	}
	return true, nil
}

// UpdateUserInfo 修改用户信息
func (s *UserService) UpdateUserInfo(uid int64, info *models.UserInfo) (*models.Account, error) {
	if info == nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	result, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	// 处理零值 传入参数为零值时不修改原始数据
	newInfo := &models.UserInfo{
		Age:           checkZeroValue(info.Age, result.UserInfo.Age).(int8),
		Gender:        checkZeroValue(info.Gender, result.UserInfo.Gender).(int8),
		Username:      checkZeroValue(info.Username, result.UserInfo.Username).(string),
		Region:        checkZeroValue(info.Region, result.UserInfo.Region).(string),
		Icon:          checkZeroValue(info.Icon, result.UserInfo.Icon).(string),
		Description:   checkZeroValue(info.Description, result.UserInfo.Description).(string),
		QRCode:        checkZeroValue(info.QRCode, result.UserInfo.QRCode).(string),
		Birthday:      checkZeroValue(info.Birthday, result.UserInfo.Birthday).(time.Time),
		LastLoginTime: checkZeroValue(info.LastLoginTime, result.UserInfo.LastLoginTime).(time.Time),
	}
	err = s.D.UpdateUserInfo(int64(result.UserId), newInfo)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UpdateUserInfoError))
	}
	resAccount, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	return resAccount, nil
}

// mergeIgnoreAccounts 合并多个 IgnoreAccounts
func mergeIgnoreAccounts(accounts1 []*models.IgnoreAccount, accounts2 []*models.IgnoreAccount) []*models.IgnoreAccount {
	accounts := make([]*models.IgnoreAccount, 0, len(accounts1)+len(accounts2))
	tmpMap := make(map[string]*models.IgnoreAccount, 0)
	for _, v1 := range accounts1 {
		tmpMap[v1.Phone] = v1
	}
	for _, v2 := range accounts2 {
		tmpMap[v2.Phone] = v2
	}
	for _, v := range tmpMap {
		accounts = append(accounts, v)
	}
	return accounts
}

// FuzzyQueryUsers 模糊查询用户信息
func (s *UserService) FuzzyQueryUsers(param string) ([]*models.IgnoreAccount, error) {
	var accounts []*models.IgnoreAccount
	usersByPhone, err := s.D.FindUsersLikePhone(param)
	if err != nil {
		return nil, err
	}
	usersByEmail, err := s.D.FindUsersLikeEmail(param)
	if err != nil {
		return usersByPhone, nil
	}
	accounts = mergeIgnoreAccounts(usersByPhone, usersByEmail)
	usersByName, err := s.D.FindUsersLikeName(param)
	if err != nil {
		return accounts, nil
	}
	accounts = mergeIgnoreAccounts(accounts, usersByName)
	return accounts, nil
}
