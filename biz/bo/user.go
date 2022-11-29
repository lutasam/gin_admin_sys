package bo

import "github.com/lutasam/gin_admin_sys/biz/vo"

type UpdateUserRequest struct {
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	Sign     string `json:"required" binding:"-"`
}

type UpdateUserResponse struct{}

type FindUserRequest struct{}

type FindUserResponse struct {
	User *vo.UserVO `json:"user"`
}
