package dao

import (
	"errors"
	"fmt"
	"site/app/site-user-svc/pkg/models"
	"site/protocol/shared"
	"strconv"
)

// CreateAccount 插入单条账户记录
func (d *Dao) CreateAccount(account *models.Account) error {
	if account == nil {
		return errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	result := tx.Model(models.Account{}).Create(account)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// FindOneAccountById 通过 id 查询账户信息
func (d *Dao) FindOneAccountById(id int64) (*models.Account, error) {
	var a models.Account
	result := d.DB.First(&a, uint(id))
	if result.Error != nil {
		return nil, result.Error
	}
	if a.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &a, nil
}

// FindOneAccountByUid 通过 uid 查询账户信息
func (d *Dao) FindOneAccountByUid(uid int64) (*models.Account, error) {
	var a models.Account
	result := d.DB.Where(&models.Account{Uid: uid}).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	if a.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &a, nil
}

// FindOneAccountByPhone 通过电话查询账户信息
func (d *Dao) FindOneAccountByPhone(phone string) (*models.Account, error) {
	var a models.Account
	result := d.DB.Where(&models.Account{Phone: phone}).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	if a.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &a, nil
}

// FindOneAccountByEmail 通过邮箱查询账户信息
func (d *Dao) FindOneAccountByEmail(email string) (*models.Account, error) {
	var a models.Account
	result := d.DB.Where(&models.Account{Email: email}).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	if a.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &a, nil
}

// FindOneAccountByUserId 通过关联 userinfo 表的外键 userId 查询账户信息 两表之间是 1:1 关联的
func (d *Dao) FindOneAccountByUserId(userId uint) (*models.Account, error) {
	var a models.Account
	result := d.DB.Where(&models.Account{UserId: userId}).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	if a.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &a, nil
}

// UpdatePasswordByUid 修改密码
func (d *Dao) UpdatePasswordByUid(uid int64, password string) error {
	a, err := d.FindOneAccountByUid(uid)
	if err != nil {
		return err
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}
	a.Password = password
	result := tx.Save(a)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// UpdateEmailById 修改邮箱
func (d *Dao) UpdateEmailById(id int64, email string) error {
	a, err := d.FindOneAccountById(id)
	if err != nil {
		return err
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}
	a.Email = email
	result := d.DB.Save(a)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// UnsubscribeById 注销账号 用 id
// TODO: 需要处理关联数据
func (d *Dao) UnsubscribeById(id int64) error {
	a, err := d.FindOneAccountById(id)
	if err != nil {
		return err
	}
	result := d.DB.Delete(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UnsubscribeByUId 注销账号 用 uid
func (d *Dao) UnsubscribeByUId(uid int64) error {
	a, err := d.FindOneAccountByUid(uid)
	if err != nil {
		return err
	}
	result := d.DB.Delete(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindOneAccountByAccount 通过账号查询账户信息
func (d *Dao) FindOneAccountByAccount(accountType int, account string) (*models.Account, error) {
	var a *models.Account
	var err error
	switch accountType {
	case shared.AccountTypeUid:
		var uid int64
		uid, err = strconv.ParseInt(account, 10, 64)
		if err != nil {
			return nil, err
		}
		a, err = d.FindOneAccountByUid(uid)
		if err != nil {
			return nil, err
		}
	case shared.AccountTypePhone:
		a, err = d.FindOneAccountByPhone(account)
		if err != nil {
			return nil, err
		}
	case shared.AccountTypeEmail:
		a, err = d.FindOneAccountByEmail(account)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.AccountInvalid))
	}
	return a, nil
}

// FindOneUserByAccount 通过账号查询用户信息
func (d *Dao) FindOneUserByAccount(accountType int, account string) (*models.Account, error) {
	a, err := d.FindOneAccountByAccount(accountType, account)
	if err != nil {
		return nil, err
	}
	var info *models.UserInfo
	info, err = d.FindUserInfoById(int64(a.UserId))
	if err != nil {
		return nil, err
	}
	a.UserInfo = *info
	var relation *models.UserRelation
	relation, err = d.FindRelationById(int64(a.RelationId))
	if err != nil {
		return nil, err
	}
	a.UserRelation = *relation
	return a, nil
}

// FindOneUserByUid 通过 Uid 查询用户信息
func (d *Dao) FindOneUserByUid(uid int64) (*models.Account, error) {
	a, err := d.FindOneAccountByUid(uid)
	if err != nil {
		return nil, err
	}
	var info *models.UserInfo
	info, err = d.FindUserInfoById(int64(a.UserId))
	if err != nil {
		return nil, err
	}
	a.UserInfo = *info
	var relation *models.UserRelation
	relation, err = d.FindRelationById(int64(a.RelationId))
	if err != nil {
		return nil, err
	}
	a.UserRelation = *relation
	return a, nil
}

// TODO: 批量查询后期控制单次查询数量上限

// FindUsersLikePhone 通过手机模糊查询用户信息
func (d *Dao) FindUsersLikePhone(paramPhone string) ([]*models.IgnoreAccount, error) {
	accounts := make([]*models.Account, 0)
	result := d.DB.Model(&models.Account{}).Where("phone LIKE ?", fmt.Sprintf("%%%v%%", paramPhone)).Find(&accounts)
	if result.Error != nil {
		return make([]*models.IgnoreAccount, 0), result.Error
	}
	retAccounts := make([]*models.IgnoreAccount, 0, len(accounts))
	var info *models.UserInfo
	var err error
	for _, account := range accounts {
		info, err = d.FindUserInfoById(int64(account.UserId))
		if err != nil {
			continue
		}
		retAccounts = append(retAccounts, &models.IgnoreAccount{
			Phone:    account.Phone,
			Email:    account.Email,
			Username: info.Username,
			Icon:     info.Icon,
		})
	}
	return retAccounts, nil
}

// FindUsersLikeEmail 通过邮箱模糊查询用户信息
func (d *Dao) FindUsersLikeEmail(paramEmail string) ([]*models.IgnoreAccount, error) {
	accounts := make([]*models.Account, 0)
	result := d.DB.Model(&models.Account{}).Where("email LIKE ?", fmt.Sprintf("%%%v%%", paramEmail)).Find(&accounts)
	if result.Error != nil {
		return make([]*models.IgnoreAccount, 0), result.Error
	}
	retAccounts := make([]*models.IgnoreAccount, 0, len(accounts))
	var info *models.UserInfo
	var err error
	for _, account := range accounts {
		info, err = d.FindUserInfoById(int64(account.UserId))
		if err != nil {
			continue
		}
		retAccounts = append(retAccounts, &models.IgnoreAccount{
			Phone:    account.Phone,
			Email:    account.Email,
			Username: info.Username,
			Icon:     info.Icon,
		})
	}
	return retAccounts, nil
}

// FindUsersLikeName 通过用户名模糊查询用户信息
func (d *Dao) FindUsersLikeName(paramName string) ([]*models.IgnoreAccount, error) {
	infos, err := d.FindUserInfosLikeName(paramName)
	if err != nil {
		return nil, err
	}
	var account *models.Account
	retAccounts := make([]*models.IgnoreAccount, 0, len(infos))
	for _, info := range infos {
		account, err = d.FindOneAccountByUserId(info.ID)
		if err != nil {
			continue
		}
		retAccounts = append(retAccounts, &models.IgnoreAccount{
			Phone:    account.Phone,
			Email:    account.Email,
			Username: info.Username,
			Icon:     info.Icon,
		})
	}
	return retAccounts, nil
}
