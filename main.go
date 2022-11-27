package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	InitRouterAndMiddleware(r)

	err := r.Run(":" + utils.GetConfigString("server.port"))
	if err != nil {
		panic(err)
	}
}
