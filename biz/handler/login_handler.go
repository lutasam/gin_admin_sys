package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/service"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

type LoginController struct{}

func RegisterLoginRouter(r *gin.RouterGroup) {
	loginController := &LoginController{}
	{
		r.POST("/do_login", loginController.DoLogin)
		r.POST("/apply_register", loginController.ApplyRegister)
		r.POST("/active_user", loginController.ActiveUser)
		r.POST("/reset_password", loginController.ResetPassword)
		r.POST("/active_reset_password", loginController.ActiveResetPassword)
	}
}

func (ins *LoginController) DoLogin(c *gin.Context) {
	req := &bo.LoginRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().DoLogin(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERDOESNOTEXIST, common.PASSWORDISERROR, common.USERINPUTERROR, common.UNKNOWNERROR, common.USERNOTACTIVE) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ApplyRegister(c *gin.Context) {
	req := &bo.ApplyRegisterRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().ApplyRegister(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR, common.USEREXISTED) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ActiveUser(c *gin.Context) {
	req := &bo.ActiveUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveUser(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR, common.USEREXISTED, common.USERDOESNOTEXIST, common.ACTIVECODEERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ResetPassword(c *gin.Context) {
	req := &bo.ResetPasswordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().ResetPassword(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR, common.USEREXISTED) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) ActiveResetPassword(c *gin.Context) {
	req := &bo.ActiveResetPasswordRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveResetPassword(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR, common.USEREXISTED, common.USERDOESNOTEXIST, common.ACTIVECODEERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
