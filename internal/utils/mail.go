package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

var (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = 587
	smtpUsername = "your@gmail.com"
	smtpPassword = "owushoxkwzfmfxwu"
)

func SendEmail(toEmail, subject, body string) error {
	from := smtpUsername
	password := smtpPassword

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-version: 1.0;\r\n"+
		"Content-Type: text/plain; charset=\"UTF-8\";\r\n\r\n"+
		"%s", from, toEmail, subject, body)

	conn, err := net.Dial("tcp", net.JoinHostPort(smtpHost, fmt.Sprintf("%d", smtpPort)))
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}

	// Bắt đầu TLS
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		return err
	}

	// Xác thực
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err = c.Auth(auth); err != nil {
		return err
	}

	// Gửi từ ai, đến ai
	if err = c.Mail(from); err != nil {
		return err
	}
	if err = c.Rcpt(toEmail); err != nil {
		return err
	}

	// Gửi nội dung
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
