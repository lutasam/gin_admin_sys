package service

import (
	"errors"
	"fmt"
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
	user, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
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
		Email: user.Email,
		Token: jwt,
	}, nil
}

func (ins *LoginService) ApplyRegister(c *gin.Context, req *bo.ApplyRegisterRequest) (*bo.ApplyRegisterResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil && errors.Is(err, common.DATABASEERROR) {
		return nil, err
	}
	if err == nil { // username is duplicate with other guys
		return nil, common.USEREXISTED
	}
	go func() {
		err := sendActiveUserEmail(c, req.Email)
		if err != nil {
			panic(err)
		}
	}()
	return &bo.ApplyRegisterResponse{}, nil
}

func (ins *LoginService) ActiveUser(c *gin.Context, req *bo.ActiveUserRequest) (*bo.ActiveUserResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" || req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:       utils.GenerateUserID(),
		Email:    req.Email,
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
	return &bo.ActiveUserResponse{}, nil
}

func (ins *LoginService) ResetPassword(c *gin.Context, req *bo.ResetPasswordRequest) (*bo.ResetPasswordResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	go func() {
		err := sendActiveUserEmail(c, req.Email)
		if err != nil {
			panic(err)
		}
	}()
	return &bo.ResetPasswordResponse{}, nil
}

func (ins *LoginService) ActiveResetPassword(c *gin.Context, req *bo.ActiveResetPasswordRequest) (*bo.ActiveResetPasswordResponse, error) {
	if !utils.IsValidEmail(req.Email) || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	code, err := dal.GetActiveDal().GetActiveCodeIfExist(c, req.Email)
	if err != nil {
		return nil, err
	}
	if code != req.ActiveCode {
		return nil, common.ACTIVECODEERROR
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptPass
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.ActiveResetPasswordResponse{}, nil
}

func sendActiveUserEmail(c *gin.Context, email string) error {
	activeCode := utils.GenerateActiveCode()
	err := dal.GetActiveDal().SetActiveCode(c, email, activeCode)
	if err != nil {
		return err
	}
	subject := "[验证激活码]某某后台管理系统"
	body := `
验证码：%s。5分钟之内有效。<br>
如果不是您本人操作，请忽视该邮件。
`
	err = utils.SendMail(email, subject, fmt.Sprintf(body, activeCode))
	if err != nil {
		return err
	}
	return nil
}
