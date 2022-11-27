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
	if !utils.IsValidEmail(req.Email) || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().FindUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	if user.IsActive == false {
		return nil, common.USERNOTACTIVE
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
		Email: user.Email,
		Token: jwt,
	}, nil
}

func (ins *LoginService) DoRegister(c *gin.Context, req *bo.RegisterRequest) (*bo.RegisterResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" || req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().FindUserByEmail(c, req.Email)
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
		Email:    req.Email,
		Password: encryptPass,
		NickName: req.Nickname,
		Avatar:   req.Avatar,
		Sign:     req.Sign,
		IsActive: false,
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
	go func() {
		err := sendActiveEmail(c, user.Email)
		if err != nil {
			panic(err)
		}
	}()
	return &bo.RegisterResponse{}, nil
}

func (ins *LoginService) ActiveAccount(c *gin.Context, req *bo.ActiveRequest) (*bo.ActiveResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	err = dal.GetUserDal().ActiveUser(c, req.Email)
	if err != nil {
		return nil, err
	}
	return &bo.ActiveResponse{}, nil
}

func sendActiveEmail(c *gin.Context, email string) error {
	activeCode := utils.GenerateActiveCode()
	err := dal.GetActiveDal().SetActiveCode(c, email, activeCode)
	if err != nil {
		return err
	}
	subject := "[验证激活码]某某后台管理系统"
	body := "验证码：" + activeCode + "\n5分钟之内有效。"
	err = utils.SendMail(email, subject, body)
	if err != nil {
		return err
	}
	return nil
}
