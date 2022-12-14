package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/bo"
	"github.com/lutasam/gin_admin_sys/biz/service"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

type DemoController struct{}

func RegisterDemoRouter(r *gin.RouterGroup) {
	demoController := &DemoController{}
	{
		r.GET("/ping", demoController.Ping)
		r.POST("/hello", demoController.Hello)
	}
}

func (ins *DemoController) Ping(c *gin.Context) {
	pong, err := service.GetDemoService().Ping(c)
	if err != nil {
		utils.Response(c, 400, 400, "server error", nil)
		return
	}
	utils.Response(c, 200, 200, "OK", pong)
}

func (ins *DemoController) Hello(c *gin.Context) {
	req := &bo.HelloRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.Response(c, 400, 400, "server error", nil)
		return
	}
	hello, err := service.GetDemoService().Hello(c, req)
	if err != nil {
		utils.Response(c, 400, 400, "server error", nil)
		return
	}
	utils.Response(c, 200, 200, "OK", hello)
}
