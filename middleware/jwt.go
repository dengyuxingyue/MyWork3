package middleware

import (
	"time"
	"work/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //token无权限
			} else if time.Now().Unix() > claim.ExpiresAt.Unix() {
				code = 401 //token过期
			}
		}

		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "token解析错误",
			})
			//阻止后续的处理器
			c.Abort()
			return
		}
		c.Next()

	}
}
