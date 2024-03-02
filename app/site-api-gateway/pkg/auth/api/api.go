package api

// 注册 =================================================================

type RegisterReqBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRespBody struct {
	Account string `json:"account"`
}

// 登录 =================================================================

type LoginReqBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginRespBody struct {
	Token string `json:"token"`
}
