package validators

import (
	"gohub/pkg/captcha"
	"net/url"
)

func ValidateCaptcha(captchaID, captchaAnswer string, errs url.Values) url.Values {

	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码答案错误")
	}

	return errs
}
