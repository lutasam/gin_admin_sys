package utils

import (
	"gopkg.in/gomail.v2"
)

func SendMail(userName, host, mailTo, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(GetConfigString("mail.address"), GetConfigString("mail.username")))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, 25, userName, GetConfigString("mail.auth"))
	err := d.DialAndSend(m)
	return err
}
