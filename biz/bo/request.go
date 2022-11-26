package bo

type HelloRequest struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
