package dao

import (
	"errors"
	"site/app/site-user-svc/pkg/models"
	"site/protocol/shared"
	"strconv"
)

// CreateAccount 插入账户记录
func (d *Dao) CreateAccount(account *models.Account) {
	d.DB.Create(account)
}

// FindOneAccountById 通过 id 查询账户信息
func (d *Dao) FindOneAccountById(id int64) (*models.Account, error) {
	var a models.Account
	result := d.DB.First(&a, uint(id))
	if result.Error != nil {
		return nil, result.Error
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
	return &a, nil
}

// FindOneAccountByPhone 通过电话查询账户信息
func (d *Dao) FindOneAccountByPhone(phone string) (*models.Account, error) {
	var a models.Account
	result := d.DB.Where(&models.Account{Phone: phone}).First(&a)
	if result.Error != nil {
		return nil, result.Error
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
	return &a, nil
}

// UpdatePasswordByUid 修改密码
func (d *Dao) UpdatePasswordByUid(uid int64, password string) error {
	a, err := d.FindOneAccountByUid(uid)
	if err != nil {
		return err
	}
	a.Password = password
	d.DB.Save(a)
	return nil
}

// UpdateEmailById 修改邮箱
func (d *Dao) UpdateEmailById(id int64, email string) error {
	a, err := d.FindOneAccountById(id)
	if err != nil {
		return err
	}
	a.Email = email
	d.DB.Save(a)
	return nil
}

// UnsubscribeById 注销账号 用 id
func (d *Dao) UnsubscribeById(id int64) error {
	a, err := d.FindOneAccountById(id)
	if err != nil {
		return err
	}
	d.DB.Delete(a)
	return nil
}

// UnsubscribeByUId 注销账号 用 uid
func (d *Dao) UnsubscribeByUId(uid int64) error {
	a, err := d.FindOneAccountByUid(uid)
	if err != nil {
		return err
	}
	d.DB.Delete(a)
	return nil
}

// FindOneAccountByAccount 通过账号查询账户信息
func (d *Dao) FindOneAccountByAccount(accountType int, account string) (*models.Account, error) {
	var a *models.Account
	var err error
	switch accountType {
	case shared.AccountTypeUid:
		var uid int
		uid, err = strconv.Atoi(account)
		if err != nil {
			return nil, err
		}
		a, err = d.FindOneAccountByUid(int64(uid))
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
	return a, nil
}

// FindUsersLikePhone 通过手机模糊查询用户信息
func (d *Dao) FindUsersLikePhone(paramPhone string) ([]*models.Account, error) {
	ret := make([]*models.Account, 0)

	return ret, nil
}
