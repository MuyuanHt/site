package dao

import (
	"errors"
	"fmt"
	"site/app/site-user-svc/pkg/models"
	"site/protocol/shared"
)

// CreateUser 插入用户信息
func (d *Dao) CreateUser(info *models.UserInfo) error {
	if info == nil {
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
	result := tx.Model(models.UserInfo{}).Create(info)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// FindUserInfoById 通过 id 查询用户信息
func (d *Dao) FindUserInfoById(id int64) (*models.UserInfo, error) {
	var u *models.UserInfo
	result := d.DB.Model(models.UserInfo{}).First(&u, uint(id))
	if result.Error != nil {
		return nil, result.Error
	}
	if u.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return u, nil
}

// UpdateUserInfo 修改用户信息
func (d *Dao) UpdateUserInfo(id int64, user *models.UserInfo) error {
	if user == nil {
		return errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	info, err := d.FindUserInfoById(id)
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
	info.Username = user.Username
	info.Age = user.Age
	info.Gender = user.Gender
	info.Region = user.Region
	info.Icon = user.Icon
	info.Description = user.Description
	info.QRCode = user.QRCode
	info.Birthday = user.Birthday
	info.LastLoginTime = user.LastLoginTime
	result := tx.Save(info)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// FindUserInfosLikeName 通过用户名模糊查询对应 id
func (d *Dao) FindUserInfosLikeName(paramName string) ([]*models.IgnoreUserInfo, error) {
	infos := make([]*models.UserInfo, 0)
	result := d.DB.Model(&models.UserInfo{}).Where("username LIKE ?", fmt.Sprintf("%%%v%%", paramName)).Find(&infos)
	if result.Error != nil {
		return make([]*models.IgnoreUserInfo, 0), result.Error
	}
	retInfos := make([]*models.IgnoreUserInfo, 0, len(infos))
	for _, v := range infos {
		retInfos = append(retInfos, &models.IgnoreUserInfo{
			ID:       v.ID,
			Username: v.Username,
			Icon:     v.Icon,
		})
	}
	return retInfos, nil
}
