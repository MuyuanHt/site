package models

import "gorm.io/gorm"

// Account 账号信息表
type Account struct {
	gorm.Model
	Uid      int64    `json:"uid"`                // uid 与账号关联
	UserId   uint     `json:"user_id"`            // 关联用户信息表 用户Id
	Phone    string   `json:"phone"`              // 电话
	Password string   `json:"password"`           // 密码
	Email    string   `json:"email"`              // 邮箱
	UserInfo UserInfo `gorm:"foreignKey:user_id"` // 关联用户信息表
}
