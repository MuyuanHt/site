package dao

import (
	"site/app/site-user-svc/pkg/models"
)

// CreateUser 插入用户信息
func (d *Dao) CreateUser(user *models.UserInfo) {
	d.DB.Create(user)
}

// FindUserInfoById 通过 id 查询用户信息
func (d *Dao) FindUserInfoById(id int64) (*models.UserInfo, error) {
	var u *models.UserInfo
	result := d.DB.First(&u, uint(id))
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
