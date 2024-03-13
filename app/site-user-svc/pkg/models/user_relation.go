package models

import "gorm.io/gorm"

// UserRelation 用户关系表
type UserRelation struct {
	gorm.Model
	// 权限相关
	SearchLimit int8 `gorm:"search_limit"` // 搜索限制
	VisitLimit  int8 `gorm:"visit_limit"`  // 访问限制
	AddLimit    int8 `gorm:"add_limit"`    // 添加限制

	// 好友与群组相关信息
	FriendNum      int32 `gorm:"friend_num"`       // 全部好友数量
	TopFriendNum   int32 `gorm:"top_friend_num"`   // 置顶好友数量
	BlackFriendNum int32 `gorm:"black_friend_num"` // 拉黑好友数量
	GroupNum       int32 `gorm:"group_num"`        // 全部群组数量
	OwnerGroupNum  int32 `gorm:"owner_group_num"`  // 创建群组数量
	AdminGroupNum  int32 `gorm:"admin_group_num"`  // 管理群组数量
}

// UserLimit 用户隐私权限 用于修改访问权限时使用
type UserLimit struct {
	SearchLimit int8 `json:"search_limit"` // 搜索限制
	VisitLimit  int8 `json:"visit_limit"`  // 访问限制
	AddLimit    int8 `json:"add_limit"`    // 添加限制
}

// UserRelationNum 用户关系数量 用于修改用户关系数量使用
type UserRelationNum struct {
	FriendNum      int32 `gorm:"friend_num"`       // 全部好友数量
	TopFriendNum   int32 `gorm:"top_friend_num"`   // 置顶好友数量
	BlackFriendNum int32 `gorm:"black_friend_num"` // 拉黑好友数量
	GroupNum       int32 `gorm:"group_num"`        // 全部群组数量
	OwnerGroupNum  int32 `gorm:"owner_group_num"`  // 创建群组数量
	AdminGroupNum  int32 `gorm:"admin_group_num"`  // 管理群组数量
}
