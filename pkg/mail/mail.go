package mail

import (
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // 文本(可选)
	HTML    []byte // HTML(可选)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMail *Mailer

func NewMailer() *Mailer {
	once.Do(func() {
		internalMail = &Mailer{Driver: &SMTP{}}
	})

	return internalMail
}

func (m *Mailer) SendMail(email Email) bool {
	return m.Driver.SendMail(email)
}
