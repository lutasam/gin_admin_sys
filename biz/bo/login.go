package bo

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ApplyRegisterRequest struct {
	Email string `json:"email" binding:"required"`
}

type ApplyRegisterResponse struct{}

type ActiveUserRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Nickname   string `json:"nickname" binding:"-"`
	Avatar     string `json:"avatar" binding:"-"`
	Sign       string `json:"sign" binding:"-"`
}

type ActiveUserResponse struct{}

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordResponse struct{}

type ActiveResetPasswordRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type ActiveResetPasswordResponse struct{}
