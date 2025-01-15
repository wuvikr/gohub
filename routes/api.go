package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// v1 路由组
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// group 中的限流中间件，测试时可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机号是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			// 使用手机号短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 使用账号密码进行登录， 支持手机号，邮箱和用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新 Token
			authGroup.POST("/login/refresh-token", middlewares.GuestJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)

		}

		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			// 创建分类
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)

			// 分类列表
			cgcGroup.GET("", cgc.Index)

			// 删除分类
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			// 创建话题
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)

			// 更新话题
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
		}

	}
}
