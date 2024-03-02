package service

import (
	"errors"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"site/app/site-user-svc/pkg/dao"
	"site/app/site-user-svc/pkg/models"
	"site/app/site-user-svc/pkg/utils"
	"site/protocol/shared"
	"time"
)

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

// FindOneUser 查询单个用户
func (s *UserService) FindOneUser(accountType int32, account string) (*models.Account, error) {
	info, err := s.D.FindOneUserByAccount(int(accountType), account)
	if info != nil {
		return info, nil
	}
	return nil, err
}

// CreateOneUser 创建单个用户
func (s *UserService) CreateOneUser(accountType int32, account string, password string) (*models.Account, error) {
	result, err := s.FindOneUser(accountType, account)
	if err == nil || result != nil {
		return result, errors.New(shared.CodeMessageIgnoreCode(shared.UserExists))
	}
	// TODO: 生成uid, 默认头像
	info := &models.Account{
		Uid:      utils.GetNextSnowflakeId(s.S),
		Password: password,
		UserInfo: models.UserInfo{
			Username:      account, // 默认用户名为账号
			Gender:        shared.GenderTypeNull,
			Icon:          "images/xxx.png",
			LastLoginTime: time.Now(),
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
	s.D.CreateAccount(info)
	return info, nil
}
