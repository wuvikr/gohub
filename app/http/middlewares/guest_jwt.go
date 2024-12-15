package middlewares

import (
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {
			// 验证是否已经登录
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "用户已登录, 请先退出登录再进行操作")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
