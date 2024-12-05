package verifycode

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"gohub/pkg/mail"
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
func (vc *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机的验证码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 本地开发环境，方便本地调试，生成的验证码都是固定的
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("VerifyCode", "验证码", map[string]string{"key": code})

	// 将生成的验证码保存到 Redis 中
	vc.Store.Set(key, code)
	return code
}

// SendEmail 通过 Email 发送验证码
//
// 1. 生成随机的验证码
// 2. 将生成的验证码保存到 Redis 中
// 3. 发送到用户的 Email
//
// 在本地开发环境中，发送的验证码都是固定的
// 在生产环境中，发送的验证码是随机的
func (vc *VerifyCode) SendEmail(to string) error {

	// 生成随机的验证码
	code := vc.generateVerifyCode(to)

	if !app.IsProduction() && strings.HasSuffix(to, config.GetString("verifycode.debug_email_suffix")) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的验证码是 %s</h1>", code)

	emailData := mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{to},
		Subject: "GoHub 验证码",
		Text:    []byte(content),
		HTML:    []byte(content),
	}

	logger.DebugJSON("VerifyCode", "发送 Email", emailData)
	mail.NewMailer().SendMail(emailData)

	return nil
}
