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
		r.POST("/do_register", loginController.DoRegister)
		r.POST("/active_user", loginController.ActiveUser)
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

func (ins *LoginController) DoRegister(c *gin.Context) {
	req := &bo.RegisterRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().DoRegister(c, req)
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
	req := &bo.ActiveRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
	}
	resp, err := service.GetLoginService().ActiveAccount(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERINPUTERROR, common.USERDOESNOTEXIST, common.ACTIVECODEERROR) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
