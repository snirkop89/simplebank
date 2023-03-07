package mail

import (
	"testing"

	"github.com/snirkop89/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	smtpConfig := SMTPConfig{
		Host:          config.SmtpHost,
		Port:          config.SmtpPort,
		Username:      config.SmtpUsername,
		Password:      config.SmtpPassword,
		SenderName:    config.SmtpSenderName,
		SenderAddress: config.SmtpSenderAddress,
	}
	sender := NewSender(smtpConfig)
	subjet := "A test email"
	content := `
	<h1>Hello world!</h1>
	<p>This is a test message from your bank SimpleBank</p>
	`
	to := []string{"john@example.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subjet, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
