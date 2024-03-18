package api

// 添加好友 =================================================================

type AddFriendReqBody struct {
	Uid         string `json:"uid"`
	FriendId    string `json:"friend_id"`
	UserLabel   string `json:"user_label"`
	FriendLabel string `json:"friend_label"`
}

type AddFriendRespBody struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
}

// 删除好友 =================================================================

type DelFriendReqBody struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
}

type DelFriendRespBody struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
}

// 更新好友信息 =================================================================

type UpdateFriendInfoReqBody struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
	IsTop    bool   `json:"is_top"`
	IsBlack  bool   `json:"is_black"`
	Label    string `json:"label"`
}

type UpdateFriendInfoRespBody struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
	IsTop    bool   `json:"is_top"`
	IsBlack  bool   `json:"is_black"`
	Label    string `json:"label"`
}

// 查询全部好友

type FriendData struct {
	FriendId string `json:"friend_id"`
	IsTop    bool   `json:"is_top"`
	IsBlack  bool   `json:"is_black"`
	Label    string `json:"label"`
}

type FindAllFriendsReqBody struct {
	Uid string `json:"uid"`
}

type FindAllFriendsRespBody struct {
	FriendNum int           `json:"friend_num"`
	Friends   []*FriendData `json:"friends"`
}
