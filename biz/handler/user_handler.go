package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/service"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

type UserController struct{}

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := &UserController{}
	{
		r.POST("/update_user", userController.UpdateUser)
		r.GET("/find_user", userController.FindUser)
	}
}

func (ins *UserController) UpdateUser(c *gin.Context) {
	req := &bo.UpdateUserRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetUserService().UpdateUser(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERNOTLOGIN, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) FindUser(c *gin.Context) {
	req := &bo.FindUserRequest{}
	resp, err := service.GetUserService().FindUser(c, req)
	if err != nil {
		if utils.IsIncludedByErrors(err, common.USERNOTLOGIN, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, err.(common.Error))
			return
		} else {
			utils.ResponseServerError(c, err.(common.Error))
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
