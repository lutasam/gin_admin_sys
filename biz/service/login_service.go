package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/dal"
	"github.com/lutasam/gin_admin_sys/biz/model"
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
	err = utils.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	jwt, err := utils.GenerateJWTByUserInfo(user)
	if err != nil {
		return nil, err
	}
	return &bo.LoginResponse{
		Username: user.Username,
		Token:    jwt,
	}, nil
}

func (ins *LoginService) DoRegister(c *gin.Context, req *bo.RegisterRequest) (*bo.RegisterResponse, error) {
	if req.Username == "" || req.Password == "" || req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().GetUserByUsername(c, req.Username)
	if err != nil && errors.Is(err, common.DATABASEERROR) {
		return nil, err
	}
	if err == nil { // username is duplicate with other guys
		return nil, common.USEREXISTED
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user = &model.User{
		ID:       utils.GenerateUserID(),
		Username: req.Username,
		Password: encryptPass,
		NickName: req.Nickname,
		Avatar:   req.Avatar,
		Sign:     req.Sign,
	}
	if req.Nickname == "" {
		user.NickName = common.DEFAULTNICKNAME
	}
	if req.Avatar == "" {
		user.Avatar = common.DEFAULTAVATARURL
	}
	if req.Sign == "" {
		user.Sign = common.DEFAULTSIGN
	}
	err = dal.GetUserDal().CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.RegisterResponse{}, nil
}
