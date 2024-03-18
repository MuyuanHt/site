package models

import (
	"gorm.io/gorm"
)

// UserFriend 好友表
type UserFriend struct {
	gorm.Model
	IsTop    bool   `gorm:"is_top"`    // 是否置顶
	IsBlack  bool   `gorm:"is_black"`  // 是否拉黑
	UserId   int64  `gorm:"user_id"`   // 用户ID
	FriendId int64  `gorm:"friend_id"` // 好友ID
	Label    string `gorm:"label"`     // 好友备注
}
