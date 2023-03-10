package mail

import (
	"net/smtp"
)

type Mailer struct {
}

var mailSender Mailer

func InitMailSender() Mailer {
	return mailSender
}

var form = "jia@qq.com"
var auth = smtp.PlainAuth("", "", "", "localhost")

func (m Mailer) SendMail(to []string, msg string) error {
	err := smtp.SendMail("localhost:1025", auth, form, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
