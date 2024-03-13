package models

import (
	"gorm.io/gorm"
	"time"
)

// UserInfo 用户基本信息表
type UserInfo struct {
	gorm.Model
	Age           int8      `gorm:"age"`             // 年龄
	Gender        int8      `gorm:"gender"`          // 性别 0-未知 1-男 2-女
	Username      string    `gorm:"username"`        // 名称
	Region        string    `gorm:"region"`          // 地区
	Icon          string    `gorm:"icon"`            // 头像
	Description   string    `gorm:"description"`     // 描述
	QRCode        string    `gorm:"qr_code"`         // 二维码
	Birthday      time.Time `gorm:"birthday"`        // 生日
	LastLoginTime time.Time `gorm:"last_login_time"` // 最后登录时间
}

// IgnoreUserInfo 用于批量查询时只返回关键信息
type IgnoreUserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}
