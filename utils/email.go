package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SendEmail(recievers []string, resetURL string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error fetching smtp credentials, error: %v", err)
	}

	sender := os.Getenv("EmailAddr")
	password := os.Getenv("SMTPpwd")

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error fetching smtp credentials, error: %v", err)
	}

	var body bytes.Buffer
	templatePath := filepath.Join(dir, "utils", "emailTemplate.html")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing html email template, err : %v", err)
	}
	t.Execute(&body, struct{ ResetURL string }{ResetURL: resetURL})

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	auth := smtp.PlainAuth("", sender, password, host)

	message := []byte(
		"To: recipient@example.com\r\n" +
			"Subject: Password Reset\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			body.String())

	err = smtp.SendMail(address, auth, sender, recievers, message)
	if err != nil {
		return fmt.Errorf("error sending Email:", err)
	}
	return nil
}
