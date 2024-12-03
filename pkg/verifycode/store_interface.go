package verifycode

type Store interface {
	// 保存验证码
	Set(id, value string) bool

	// 获取验证码
	Get(id string, clear bool) string

	// 校验验证码
	Verify(id, answer string, clear bool) bool
}
