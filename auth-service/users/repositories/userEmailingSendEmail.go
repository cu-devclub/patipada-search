package repositories

import (
	"auth-service/config"
	"auth-service/users/entities"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func (u *userJordanWrightEmailing) SendEmail(in *entities.Email) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", u.name, u.fromEmailAddress)
	e.Subject = in.Subject
	e.HTML = []byte(in.Content)
	e.To = in.To
	e.Cc = in.CC
	e.Bcc = in.BCC

	for _, f := range in.AttachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	cfg := config.GetConfig()
	smtpAuthAddress := cfg.Email.Host
	smtpServerAddress := fmt.Sprintf("%s:%d", smtpAuthAddress, cfg.Email.Port)
	smtpAuth := smtp.PlainAuth("", u.fromEmailAddress, u.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
