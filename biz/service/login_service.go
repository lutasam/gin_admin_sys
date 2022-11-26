package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/dal"
	"github.com/lutasam/gin_admin_sys/biz/utils"
	"sync"
)

type LoginService struct{}

var (
	loginService     *LoginService
	loginServiceOnce sync.Once
)

func GetLoginService() *LoginService {
	loginServiceOnce.Do(func() {
		loginService = &LoginService{}
	})
	return loginService
}

func (ins *LoginService) DoLogin(c *gin.Context, req *bo.LoginRequest) (*bo.LoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().GetUserByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}
	//err = utils.ValidatePassword(user.Password, req.Password)
	//if err != nil {
	//	return nil, err
	//}
	jwt, err := utils.GenerateJWTByUserInfo(user)
	if err != nil {
		return nil, err
	}
	return &bo.LoginResponse{
		Username: user.Username,
		Token:    jwt,
	}, nil
}
