package sendto

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	smtpConfig := global.Config.SMTP

	contentEmail := Mail {
		From: EmailAddress{
			Address: from,
			Name: "test",
		},
		To: to,
		Subject: "OTP",
		Body: fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", smtpConfig.User, smtpConfig.Pass, smtpConfig.Host)
	err := smtp.SendMail(smtpConfig.Host + ":" + smtpConfig.Port, auth, smtpConfig.User, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Error sending email", zap.Error(err))
		return err
	}
	return nil
}
