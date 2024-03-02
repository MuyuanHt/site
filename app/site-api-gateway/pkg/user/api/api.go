package api

// 查找单个用户 =================================================================

type FindOneUserReqBody struct {
	AccountType string `json:"account_type"` // 1-手机 2-邮箱 ...
	Account     string `json:"account"`
}

type FindOneUserRespBody struct {
	Uid         int64  `json:"uid"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Age         int32  `json:"age"`
	Gender      int32  `json:"gender"`
	Region      string `json:"region"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Birthday    int64  `json:"birthday"`
}
