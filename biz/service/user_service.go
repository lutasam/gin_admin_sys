package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/dal"
	"github.com/lutasam/gin_admin_sys/biz/model"
	"github.com/lutasam/gin_admin_sys/biz/utils"
	"github.com/lutasam/gin_admin_sys/biz/vo"
	"sync"
)

type UserService struct{}

var (
	userService     *UserService
	userServiceOnce sync.Once
)

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (ins *UserService) UpdateUser(c *gin.Context, req *bo.UpdateUserRequest) (*bo.UpdateUserResponse, error) {
	if !utils.IsValidURL(req.Avatar) || req.Password == "" || req.Nickname == "" {
		return nil, common.USERINPUTERROR
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, userInfo.Email)
	if err != nil {
		return nil, err
	}
	user.Password = req.Password
	user.Avatar = req.Avatar
	user.NickName = req.Nickname
	user.Sign = req.Sign
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateUserResponse{}, nil
}

func (ins *UserService) FindUser(c *gin.Context, req *bo.FindUserRequest) (*bo.FindUserResponse, error) {
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, userInfo.Email)
	if err != nil {
		return nil, err
	}
	return &bo.FindUserResponse{
		User: convertToUserVO(user),
	}, nil
}

func convertToUserVO(user *model.User) *vo.UserVO {
	return &vo.UserVO{
		ID:        user.ID,
		Email:     user.Email,
		NickName:  user.NickName,
		Avatar:    user.Avatar,
		Sign:      user.Sign,
		CreatedAt: user.CreatedAt,
	}
}
