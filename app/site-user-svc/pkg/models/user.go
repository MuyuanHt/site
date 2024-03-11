package models

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	Age           int8      `json:"age"`             // 年龄
	Gender        int8      `json:"Gender"`          // 性别 0-未知 1-男 2-女
	Username      string    `json:"username"`        // 用户名
	Region        string    `json:"Region"`          // 地区
	Icon          string    `json:"Icon"`            // 头像
	Description   string    `json:"Description"`     // 描述
	Birthday      time.Time `json:"birthday"`        // 生日
	LastLoginTime time.Time `json:"last_login_time"` // 最后登录时间
}
