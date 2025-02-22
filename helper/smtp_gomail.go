package helper

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpEmail, smtpPassword)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	fmt.Printf("Email sent successfully")
	return nil
}
