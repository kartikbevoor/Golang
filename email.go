package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// basic mail using smtp
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

func SendEmailWithAttachement(to string, subject string, body string) error {

	// load .env file
	err1 := godotenv.Load()
	fmt.Println(err1)

	// get email and password
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// creates new msg: Think of it as a container where you build your email (headers, body, attachments).
	m := gomail.NewMessage() // This initializes a new email object

	// Set email headers: defines metadata of the email
	// m.SetAddressHeader("From", "your_email@gmail.com", "Your App Name") custom sender name
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	// for multiple ppl
	// m.SetHeader("To", "a@gmail.com", "b@gmail.com")

	// Email body
	m.SetBody("text/plain", body)

	// attach file
	m.Attach("order.txt")

	// can also attach multiple files
	// m.Attach("file1.pdf")
	// m.Attach("image.png")

	// smtp configuration: This creates a dialer (connection config to the email server).
	d := gomail.NewDialer(
		"smtp.gmail.com", // SMTP server
		587,              // Port
		from,             // sender
		password,         // app password
	)

	// send mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// to set HTML body
// m.SetBody("text/html", "<h1>Hello</h1><p>Account created!</p>")

func EmailWithDynamicAttachementAndUrl(to string, subject string, body string, url string) error {
	// load .env file
	err1 := godotenv.Load()
	fmt.Println(err1)

	// get email and password
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// create new msg
	m := gomail.NewMessage()

	// headers
	m.SetHeader("From", from)
	m.SetHeader("To", "to")
	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body+url)

	// function to generate content of file
	// fileContent := generateDynamicFile()

	m.Attach("order.txt")

	// m.Attach("report.txt", gomail.SetCopyFunc(func(w *gomail.FileSetting) error {
	// 	_, err := w.Write(fileContent.Bytes())
	// 	return err
	// }))

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		from,
		password,
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func generateDynamicFile() *bytes.Buffer {
	var buf bytes.Buffer
	buf.WriteString("Report generated dynamically!\n")
	buf.WriteString("Date: 2026-03-18\n")
	return &buf
}

// HTML Body with URL
// body := fmt.Sprintf(`
// 	<h2>Welcome!</h2>
// 	<p>Your account has been created successfully.</p>
// 	<p>Click below to access your dashboard:</p>
// 	<a href="%s">Go to Dashboard</a>
// 	<br><br>
// 	<p>Download your file from attachment.</p>
// `, downloadURL)

// to generate pdf
// pdfBytes := generatePDF()
// m.Attach("certificate.pdf", gomail.SetCopyFunc(func(w *gomail.FileSetting) error {
// 	_, err := w.Write(pdfBytes)
// 	return err
// }))

// ____________________________________________________
// ____________________________________________________
// ____________________________________________________
// go get github.com/emersion/go-mail : the problem with below code is unable to install this package

// the problem with below is
// alternate way infact bit easy way: go-mail, go-smtp
// func sendEmailgomail(to string, url string) error {

// 	// Step 1: Create email message
// 	m := mail.NewMessage()

// 	m.SetHeader("From", "your-email@gmail.com")
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", "Welcome 🚀")

// 	// Step 2: HTML Body with URL
// 	body := fmt.Sprintf(`
// 		<h2>Welcome to our platform</h2>
// 		<p>Your account is created successfully.</p>
// 		<p><a href="%s">Click here to access dashboard</a></p>
// 	`, url)

// 	m.SetBody("text/html", body)

// 	// Step 3: Dynamic Attachment
// 	fileBytes := generateDynamicFile()

// 	m.AttachReader("report.txt", bytes.NewReader(fileBytes))

// 	// Step 4: SMTP Config
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	auth := smtp.PlainAuth(
// 		"",
// 		"your-email@gmail.com",
// 		"your-app-password", // use App Password
// 		smtpHost,
// 	)

// 	// Step 5: Send Email
// 	err := mail.SendMail(
// 		smtpHost+":"+smtpPort,
// 		auth,
// 		"your-email@gmail.com",
// 		[]string{to},
// 		m,
// 	)

// 	return err
// }
