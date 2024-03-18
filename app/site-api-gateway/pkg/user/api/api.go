package api

// 查找单个用户 =================================================================

type FindOneUserReqBody struct {
	AccountType string `json:"account_type"` // 1-手机 2-邮箱 ...
	Account     string `json:"account"`
}

type FindOneUserRespBody struct {
	Uid         string `json:"uid"`
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

// 修改用户密码 =================================================================

type UpdatePasswordReqBody struct {
	Uid         string `json:"uid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdatePasswordRespBody struct {
	Uid     string `json:"uid"`
	Success bool   `json:"success"`
}

// 修改个人信息 =================================================================

type UpdateUserInfoReqBody struct {
	Uid         string `json:"uid"`
	Username    string `json:"username"`
	Age         int32  `json:"age"`
	Gender      int32  `json:"gender"`
	Region      string `json:"region"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Birthday    int64  `json:"birthday"`
}

type UpdateUserInfoRespBody struct {
	Uid         string `json:"uid"`
	Username    string `json:"username"`
	Age         int32  `json:"age"`
	Gender      int32  `json:"gender"`
	Region      string `json:"region"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Birthday    int64  `json:"birthday"`
}

// 批量查询用户信息

type FuzzyQueryUsersData struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}

type FuzzyQueryUsersReqBody struct {
	Param string `json:"param"`
}

type FuzzyQueryUsersRespBody struct {
	UserNum int                    `json:"user_num"`
	Users   []*FuzzyQueryUsersData `json:"users"`
}

// 修改用户隐私权限

type UpdateUserLimitReqBody struct {
	Uid         string `json:"uid"`
	SearchLimit int32  `json:"search_limit"`
	VisitLimit  int32  `json:"visit_limit"`
	AddLimit    int32  `json:"add_limit"`
}

type UpdateUserLimitRespBody struct {
	SearchLimit int32 `json:"search_limit"`
	VisitLimit  int32 `json:"visit_limit"`
	AddLimit    int32 `json:"add_limit"`
}
