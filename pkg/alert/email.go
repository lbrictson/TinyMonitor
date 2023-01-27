package alert

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type NewEmailAlerterInput struct {
	From string
	Host string
	Port int
	User string
	Pass string
	To   []string
	CC   []string
	BCC  []string
}

func NewEmailAlerter(input NewEmailAlerterInput) *EmailAlerter {
	d := gomail.NewDialer(input.Host, input.Port, input.User, input.Pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &EmailAlerter{
		conn: d,
		from: input.From,
		to:   input.To,
		cc:   input.CC,
		bcc:  input.BCC,
	}
}

type EmailAlerter struct {
	conn *gomail.Dialer
	from string
	to   []string
	cc   []string
	bcc  []string
}

func (a *EmailAlerter) SendDown(monitorName string, message string) error {
	m := gomail.NewMessage()
	m.SetBody("text/html", message)
	m.SetHeader("From", a.from)
	m.SetHeader("To", a.to...)
	m.SetHeader("Cc", a.cc...)
	m.SetHeader("Bcc", a.bcc...)
	m.SetHeader("Subject", "TinyStatus Monitor Down: "+monitorName)
	return a.conn.DialAndSend(m)
}

func (a *EmailAlerter) SendUp(monitorName string, message string) error {
	m := gomail.NewMessage()
	m.SetBody("text/html", message)
	m.SetHeader("From", a.from)
	m.SetHeader("To", a.to...)
	m.SetHeader("Cc", a.cc...)
	m.SetHeader("Bcc", a.bcc...)
	m.SetHeader("Subject", "TinyStatus Monitor Up: "+monitorName)
	return a.conn.DialAndSend(m)
}
