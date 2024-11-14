package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	// r := gin.Default()
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 处理 404 请求
	r.NoRoute(func(c *gin.Context) {
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

	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}
