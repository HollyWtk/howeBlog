package user

type LoginReq struct {
	Username string `json:"username" form:"account"`
	Password string `json:"password" form:"password"`
}

type LoginRsp struct {
	User  User      `json:"member"`
	Token TokenList `json:"tokenList"`
}

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Role       string `json:"role"`
}

type TokenList struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	TokenType      string `json:"tokenType"`
	AccessTokenExp int64  `json:"accessTokenExp"`
}
