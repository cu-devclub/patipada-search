package tests

import (
	"auth-service/config"
	"testing"

	"auth-service/users/entities"
	emailService "auth-service/users/repositories"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config.InitializeViper("../")
	cfg := config.GetConfig()

	sender := emailService.NewUserJordanWrightEmailing(cfg.Email.SenderName, cfg.Email.SenderEmail, cfg.Email.SenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from Dhammava service</p>
	`
	to := []string{cfg.Email.ReceiverTestEmail}

	in := &entities.Email{
		Subject:     subject,
		Content:     content,
		To:          to,
		AttachFiles: nil,
		CC:          nil,
		BCC:         nil,
	}

	err := sender.SendEmail(in)
	require.NoError(t, err)
}
