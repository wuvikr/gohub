package middlewares

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// 解析成功，保存用户信息到上下文
		user := user.Get(claims.UserID)
		if user.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，可能已经删除")
			return
		}

		c.Set("current_user_id", user.GetStringID())
		c.Set("current_user_name", user.Name)
		c.Set("current_user", user)
		c.Next()
	}
}
