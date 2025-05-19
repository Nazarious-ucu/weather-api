package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type EmailService struct {
	User     string
	Host     string
	Port     string
	Password string
	From     string
}

func NewEmailService() *EmailService {
	return &EmailService{
		User:     os.Getenv("SMTP_USER"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Password: os.Getenv("SMTP_PASS"),
		From:     os.Getenv("SMTP_FROM"),
	}
}

func (e *EmailService) SendConfirmationEmail(toEmail, token string) error {
	tmpl, err := template.ParseFiles("internal/templates/confirm_email.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, map[string]string{
		"Email": toEmail,
		"Link":  fmt.Sprintf("http://localhost:8080/confirm/%s", token),
	})
	if err != nil {
		return err
	}

	if e.Host == "" || e.Port == "" || e.User == "" || e.Password == "" {
		log.Fatal("❌ SMTP credentials are not set properly in .env")
	} else {
		log.Println(e.Host, e.Port, e.User, e.Password)
	}

	auth := smtp.PlainAuth("", e.User, e.Password, e.Host)
	msg := []byte("From: " + e.From + "\r\n" +
		"To: " + toEmail + "\r\n" +
		"Subject: Confirm Your Weather Subscription\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body.String())

	addr := fmt.Sprintf("%s:%s", e.Host, e.Port)
	return smtp.SendMail(addr, auth, e.From, []string{toEmail}, msg)
}

func (s *EmailService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.User, s.Password, s.Host)

	msg := "From: " + s.From + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	addr := s.Host + ":" + s.Port
	return smtp.SendMail(addr, auth, s.User, []string{to}, []byte(msg))
}
