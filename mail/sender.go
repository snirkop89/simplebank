package mail

import (
	"gopkg.in/mail.v2"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

type SMTPConfig struct {
	Host          string
	Port          int
	Username      string
	Password      string
	SenderName    string
	SenderAddress string
}

type Sender struct {
	SMTPConfig
}

var _ EmailSender = &Sender{}

func NewSender(config SMTPConfig) *Sender {
	return &Sender{
		config,
	}
}

func (sender *Sender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	d := mail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	msg := mail.NewMessage()
	msg.SetHeader("From", sender.SenderAddress)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("To", to...)
	msg.SetHeader("Cc", cc...)
	msg.SetHeader("Bcc", bcc...)
	msg.SetBody("text/html", content)

	for _, f := range attachFiles {
		msg.Attach(f)
	}

	return d.DialAndSend(msg)
}
