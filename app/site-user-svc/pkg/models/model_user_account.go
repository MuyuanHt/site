package models

import "gorm.io/gorm"

// Account 账号信息表
type Account struct {
	gorm.Model
	Uid          int64        `gorm:"uid"`                                 // uid 与账号关联
	Phone        string       `gorm:"phone"`                               // 电话
	Password     string       `gorm:"password"`                            // 密码
	Email        string       `gorm:"email"`                               // 邮箱
	UserId       uint         `gorm:"column:user_id"`                      // 关联用户信息表 用户Id
	RelationId   uint         `gorm:"column:relation_id"`                  // 关联用户关系表 关系Id
	UserInfo     UserInfo     `gorm:"foreignKey:UserId;references:ID"`     // 关联用户信息表
	UserRelation UserRelation `gorm:"foreignKey:RelationId;references:ID"` // 关联用户关系表
}

// IgnoreAccount 用于批量查询时只返回关键信息
type IgnoreAccount struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}
