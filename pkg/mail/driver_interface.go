package mail

type Driver interface {
	SendMail(email Email) bool
}
