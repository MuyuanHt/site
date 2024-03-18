package api

// 添加好友 =================================================================

type AddFriendReqBody struct {
	UserId      string `json:"user_id"`
	FriendId    string `json:"friend_id"`
	UserLabel   string `json:"user_label"`
	FriendLabel string `json:"friend_label"`
}

type AddFriendRespBody struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

// 删除好友 =================================================================

type DelFriendReqBody struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

type DelFriendRespBody struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}
