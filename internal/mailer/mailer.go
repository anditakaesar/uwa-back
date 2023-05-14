package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/anditakaesar/uwa-back/internal/log"
	"go.uber.org/zap"
)

type Mailer struct {
	From     string
	Host     string
	FullHost string
	Auth     smtp.Auth
	Logger   log.LoggerInterface
}

type MailerDependency struct {
	From     string
	SmtpUser string
	Password string
	Host     string
	Port     string
	Logger   log.LoggerInterface
}

func NewInternalMailer() *Mailer {
	return &Mailer{}
}

func (m *Mailer) Connect(d MailerDependency) {
	m.From = d.From
	m.Host = d.Host
	m.FullHost = fmt.Sprintf("%s:%s", d.Host, d.Port)
	m.Auth = smtp.PlainAuth("", d.SmtpUser, d.Password, d.Host)
	m.Logger = d.Logger
}

func (m *Mailer) SendMail(address string, subject string, body string) error {
	fullMessage := "From: " + m.From + "\n" + "To: " + address + "\n" + "Subject: " + subject + "\n\n" + body
	mailDetail := map[string]string{
		"from":    m.From,
		"address": address,
		"subject": subject,
		"body":    body,
	}

	err := verifySendMail(address, subject, body)
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] verifySendMail", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.Host,
	}

	conn, err := tls.Dial("tcp", m.FullHost, tlsconfig)
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] tls.Dial", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	c, err := smtp.NewClient(conn, m.Host)
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] smtp.NewClient", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	// Auth
	if err = c.Auth(m.Auth); err != nil {
		m.Logger.Error("[Mailer][SendMail] c.Auth", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	// To && From
	if err = c.Mail(m.From); err != nil {
		m.Logger.Error("[Mailer][SendMail] c.Mail", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	if err = c.Rcpt(address); err != nil {
		m.Logger.Error("[Mailer][SendMail] c.Rcpt", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] c.Data", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	_, err = w.Write([]byte(fullMessage))
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] w.Write", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	err = w.Close()
	if err != nil {
		m.Logger.Error("[Mailer][SendMail] w.Close", err, zap.Any("mailDetail", mailDetail))
		return err
	}

	return c.Quit()
}

func verifySendMail(address string, subject string, body string) error {
	fields := []string{}

	if address == "" {
		fields = append(fields, "address")
	}

	if subject == "" {
		fields = append(fields, "subject")
	}

	if body == "" {
		fields = append(fields, "body")
	}

	if len(fields) > 0 {
		return fmt.Errorf("VerifySendMail failed these fields are required: %s", fields)
	}

	return nil
}
