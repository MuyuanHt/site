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

type UpdateFriendInfoReq struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
	IsTop    bool   `json:"is_top"`
	IsBlack  bool   `json:"is_black"`
	Label    string `json:"label"`
}

type UpdateFriendInfoResp struct {
	Uid      string `json:"uid"`
	FriendId string `json:"friend_id"`
	IsTop    bool   `json:"is_top"`
	IsBlack  bool   `json:"is_black"`
	Label    string `json:"label"`
}
