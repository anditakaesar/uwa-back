package mailer

import (
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/anditakaesar/uwa-back/internal/mailer"
)

type MailerInterface interface {
	SendMail(address string, subject string, body string) error
}

func NewMailerAdapter(logger log.LoggerInterface) MailerInterface {
	m := mailer.NewInternalMailer()
	m.Connect(mailer.MailerDependency{
		From:     env.EmailFrom(),
		SmtpUser: env.EmailSmtpUser(),
		Password: env.EmailPassword(),
		Host:     env.EmailHost(),
		Port:     env.EmailPort(),
		Logger:   logger,
	})

	return m
}
