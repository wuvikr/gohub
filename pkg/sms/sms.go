package sms

import (
	"sync"
)

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

type SMS struct {
	Driver Driver
}

var once sync.Once
var internalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})

	return internalSMS
}

func (s *SMS) Send(phone string, message Message) bool {
	return s.Driver.Send(phone, message)
}
