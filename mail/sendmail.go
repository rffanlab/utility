package mail

import (
	"net/smtp"
	"strings"
)

type Email struct {
	Host        string
	Username    string
	Password    string
	Nickname    string
	To          []string
	Subject     string
	ContentType string
	Body        string
}

func (c *Email) Send() (err error) {
	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)
	msg := []byte("To:" + strings.Join(c.To, ",") + "\r\nForm:" + c.Nickname + "<" + c.Username + ">\r\nSubject:" + c.Subject + "\r\n" + c.ContentType + "\r\n\r\n" + c.Body)
	err = smtp.SendMail(c.Host+":25", auth, c.Username, c.To, msg)
	return

}
