package bo

type HelloRequest struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"-"`
	Avatar   string `json:"avatar" binding:"-"`
	Sign     string `json:"sign" binding:"-"`
}
