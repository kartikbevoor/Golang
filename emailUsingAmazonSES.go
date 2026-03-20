package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"gopkg.in/gomail.v2"
)

func sesEmail(to string) error {
	// Load AWS config
	// config.LoadDefaultConfig: Loads AWS configuration using the default credential and config chain
	// Environment variables, AWS config files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// Explicitly sets the AWS region: AWS services are region-specific
		config.WithRegion("ap-south-1"), // Mumbai region
		// the above function return values:
		// cfg: aws.Config, contains: Credentials, Region, Retry settings, HTTP client config
	)
	if err != nil {
		return err
	}

	// Create SES client
	client := ses.NewFromConfig(cfg) // Initializes a client for interacting with Amazon SES.

	// Email content
	input := &ses.SendEmailInput{ // This struct contains everything SES needs to send the email.
		Source: awsString("your-verified-email@gmail.com"),
		// types.Destination → a struct (data structure) defined in the AWS SES SDK, & - creates pointer
		Destination: &types.Destination{
			ToAddresses: []string{to}, // can add multiple to addresses here
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: awsString("Welcome to Our Platform 🎉"),
			},
			Body: &types.Body{
				Html: &types.Content{
					Data: awsString(`
						<h1>Welcome!</h1>
						<p>Your account has been successfully created.</p>
					`),
				},
			},
		},
	}

	// Send email
	_, err = client.SendEmail(context.TODO(), input)
	return err
}

func awsString(s string) *string {
	return &s
}

func toReadFileFromDisk() {
	fileContent, err := os.ReadFile("certificate.pdf")
	if err != nil {
		panic(err)
	}

	SendEmailWithAttachment("receiver@gmail.com", fileContent)
}

func SendEmailWithAttachment(to string, fileContent []byte) error {
	boundary := "my-boundary" // This separates different parts of the email: html content and attachment

	// Email Headers
	header := ""
	header += "From: your-verified-email@gmail.com\r\n"
	header += "To: " + to + "\r\n"
	header += "Subject: Certificate\n"
	header += "MIME-Version: 1.0\r\n" // Required for emails with attachments
	header += "Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\n"

	// Email Body
	body := "--" + boundary + "\n"
	body += "Content-Type: text/html; charset=\"UTF-8\"\n\n"
	body += "<h1>Your Certificate</h1>\n\n"

	// Attachment
	body += "--" + boundary + "\n"
	body += "Content-Type: application/pdf\n"
	body += "Content-Transfer-Encoding: base64\n"
	body += "Content-Disposition: attachment; filename=\"certificate.pdf\"\n\n"

	// Encoding the File
	encoded := base64.StdEncoding.EncodeToString(fileContent)
	body += encoded + "\n"
	body += "--" + boundary + "--"

	rawMessage := header + body

	// Load AWS Config
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	// Create SES Client
	client := ses.NewFromConfig(cfg)

	// Send Email
	_, err := client.SendRawEmail(context.TODO(), &ses.SendRawEmailInput{
		RawMessage: &types.RawMessage{
			Data: []byte(rawMessage),
		},
	})

	return err
}

// for url
// <a href="https://yourapp.com/verify?token=abc123">
//     Verify Account
// </a>

func sendEmailWithAttachment2() {

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create SES client
	svc := ses.NewFromConfig(cfg)

	from := "your-verified-email@example.com"
	to := "receiver@example.com"
	subject := "Test Email with Attachment & URL"

	// HTML body
	body := `
	<html>
		<body>
			<h2>Welcome!</h2>
			<p>Click below:</p>
			<a href="https://yourapp.com/verify">Verify Account</a>
		</body>
	</html>
	`

	// Attachment
	fileName := "test.txt"
	fileContent := []byte("This is a test attachment")

	encodedFile := base64.StdEncoding.EncodeToString(fileContent)

	// Build MIME message
	var msg bytes.Buffer
	boundary := "NextPartBoundary"

	msg.WriteString(fmt.Sprintf("From: %s\r\n", from))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", to))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n")

	// Body part
	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	msg.WriteString("Content-Transfer-Encoding: 7bit\r\n\r\n")
	msg.WriteString(body + "\r\n\r\n")

	// Attachment part
	msg.WriteString("--" + boundary + "\r\n")
	msg.WriteString("Content-Type: application/octet-stream\r\n")
	msg.WriteString("Content-Transfer-Encoding: base64\r\n")
	msg.WriteString("Content-Disposition: attachment; filename=\"" + fileName + "\"\r\n\r\n")
	msg.WriteString(encodedFile + "\r\n\r\n")

	// End boundary
	msg.WriteString("--" + boundary + "--")

	// Send email
	_, err = svc.SendRawEmail(context.TODO(), &ses.SendRawEmailInput{
		RawMessage: &types.RawMessage{
			Data: msg.Bytes(),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}

// Using gomail
func sendMailUsingGomail() {

	// SMTP Configuration (SES)
	smtpHost := "email-smtp.us-east-1.amazonaws.com"
	smtpPort := 587
	smtpUser := "YOUR_SMTP_USERNAME"
	smtpPass := "YOUR_SMTP_PASSWORD"

	from := "your-verified-email@example.com"
	to := "receiver@example.com"

	// Create new message
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome! Verify your account")

	// HTML Body with URL
	htmlBody := `
		<h2>Welcome to Our Platform 🚀</h2>
		<p>Please verify your account by clicking below:</p>
		<a href="https://yourapp.com/verify?token=abc123">Verify Account</a>
	`

	m.SetBody("text/html", htmlBody)

	// ----------------------------
	// OPTION 1: Attach existing file
	// ----------------------------
	m.Attach("sample.pdf")

	// ----------------------------
	// OPTION 2: Attach dynamic file (VERY IMPORTANT for your use case)
	// ----------------------------
	// m.Attach("report.txt", gomail.SetCopyFunc(func(w io.Writer) error {
	// 	content := "This is dynamically generated content (like certificate/report)"
	// 	_, err := w.Write([]byte(content))
	// 	return err
	// }))

	// Create dialer
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
