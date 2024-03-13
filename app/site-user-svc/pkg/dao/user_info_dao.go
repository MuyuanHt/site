package dao

import (
	"fmt"
	"site/app/site-user-svc/pkg/models"
)

// CreateUser 插入用户信息
func (d *Dao) CreateUser(info *models.UserInfo) error {
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
	return u, nil
}

// UpdateUserInfo 修改用户信息
func (d *Dao) UpdateUserInfo(id int64, user *models.UserInfo) error {
	result, err := d.FindUserInfoById(id)
	if err != nil {
		return err
	}
	result.Username = user.Username
	result.Age = user.Age
	result.Gender = user.Gender
	result.Region = user.Region
	result.Icon = user.Icon
	result.Description = user.Description
	result.QRCode = user.QRCode
	result.Birthday = user.Birthday
	result.LastLoginTime = user.LastLoginTime
	d.DB.Model(models.UserInfo{}).Save(result)
	return nil
}

// FindUserInfosLikeName 通过用户名模糊查询对应 id
func (d *Dao) FindUserInfosLikeName(paramName string) ([]*models.IgnoreUserInfo, error) {
	infos := make([]*models.UserInfo, 0)
	d.DB.Model(&models.UserInfo{}).Where("username Like ?", fmt.Sprintf("%%%v%%", paramName)).Find(&infos)
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
