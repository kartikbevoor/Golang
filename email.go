package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmail(to string, subject string, body string) error {

	err1 := godotenv.Load()
	fmt.Println(err1)

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD") // Gmail App Password: must enable 2FA and generate an App Password

	// SMTP Server Configuration
	smtpHost := "smtp.gmail.com" // SMTP server address: the server responsible for sending mails
	smtpPort := "587"            // network port: used to communicate with the SMTP server.

	// contents of email
	message := []byte(fmt.Sprintf( // Converted to []byte because SendMail expects raw bytes
		"Subject: %s\r\n\r\n%s",
		subject,
		body,
	))

	// creates an authentication object: uses plain authentication
	// Logins in using: email(from), password and restricts authentication to smtpHost (smtp.gmail.com)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Connects to an SMTP server → authenticates → sends your email → returns an error (if any)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort, // combines smtpHost and server: Connect to this server on this port
		auth,                  // Contains your login credentials: Used to authenticate with the SMTP server
		from,
		[]string{to}, // slice of recipients email address, ex: to := []string{"a@example.com", "b@example.com"}
		message,      // email content
	)

	if err != nil {
		return err
	}

	return nil
}

// message := []byte("MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
// "<h1>Welcome!</h1><p>Your account is created.</p>")
