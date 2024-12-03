package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"gohub/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信
func (vc *VerifyCode) SendSMS(phone string) bool {
	code := vc.generateVerifyCode(phone)

	// 本地开发环境，方便本地调试
	if app.IsLocal() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	// 发送短信
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})

}

// CheckAnswer 检查用户提交的验证码是否正确
func (vc *VerifyCode) CheckAnswer(key, answer string) bool {
	logger.DebugJSON("VerifyCode", "用户输入的验证码", map[string]string{"key": key, "answer": answer})

	if !app.IsProduction() && (strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix")) ||
		strings.HasPrefix(key, config.GetString("verifycode.debug_code_prefix"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码
func (vc *VerifyCode) generateVerifyCode(phone string) string {
	// 生成随机的验证码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 本地开发环境，方便本地调试，生成的验证码都是固定的
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("VerifyCode", "验证码", map[string]string{"key": code})

	// 将生成的验证码保存到 Redis 中
	vc.Store.Set(phone, code)
	return code
}
