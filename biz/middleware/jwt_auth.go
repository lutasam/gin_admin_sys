package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/gin_admin_sys/biz/common"
	"github.com/lutasam/gin_admin_sys/biz/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.ResponseClientError(c, common.USERNOTLOGIN)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			token = strings.Split(token, " ")[1]
		}
		jwtStruct, err := utils.ParseJWTToken(token)
		if err != nil {
			utils.ResponseClientError(c, common.EXCEEDTIMELIMIT)
			c.Abort()
			return
		}
		c.Set("jwtStruct", jwtStruct)
		c.Next()
	}
}
