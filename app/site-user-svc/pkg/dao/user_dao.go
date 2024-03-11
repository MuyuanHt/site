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
	result.Birthday = user.Birthday
	result.LastLoginTime = user.LastLoginTime
	d.DB.Save(result)
	return nil
}
