package validators

import (
	"gohub/pkg/captcha"
	"gohub/pkg/verifycode"
	"net/url"
)

// ValidateCaptcha 验证图片验证码
func ValidateCaptcha(captchaID, captchaAnswer string, errs url.Values) url.Values {

	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码答案错误")
	}

	return errs
}

// ValidatePasswordConfirm 验证两次输入的密码是否相同
func ValidatePasswordConfirm(password, passwordConfirm string, errs url.Values) url.Values {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不相同")
	}

	return errs
}

// ValidateVerifyCode 验证短信验证码
func ValidateVerifyCode(key, answer string, errs url.Values) url.Values {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}

	return errs
}
