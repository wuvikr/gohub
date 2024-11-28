package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// 内部使用的 Captcha 实例
var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}

		// 使用全局 Reids 对象，并配置 Key Prefix
		store := RedisStore{
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
			RedisClient: redis.Redis,
		}

		// 配置 base64Captcha driver
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 背景点数，越大字体越模糊
		)

		// 实例化 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha

}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id, b64s, answer string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证图片验证码是否正确
func (c *Captcha) VerifyCaptcha(id, answer string) bool {

	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}

	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
