package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	set404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	// 全局中间件配置
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func set404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// 获取标头中的 Accept
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 类型，返回 404
			c.String(http.StatusNotFound, "页面未找到，返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error":         404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
