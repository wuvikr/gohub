package mail

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) SendMail(email Email) bool {
	mailServerConfig := config.GetStringMapString("mail.smtp")
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "邮件配置", e)

	err := e.Send(
		fmt.Sprintf("%s:%s", mailServerConfig["host"], mailServerConfig["port"]),
		smtp.PlainAuth(
			"",
			mailServerConfig["username"],
			mailServerConfig["password"],
			mailServerConfig["host"],
		),
	)

	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}

	logger.DebugJSON("发送邮件", "发送成功", "")
	return true
}
