package context

import (
	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/adapter/mailer"
	"github.com/anditakaesar/uwa-back/adapter/models/iplog"
	"github.com/anditakaesar/uwa-back/adapter/redis"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type AppContext struct {
	DB         database.DatabaseInterface
	Logger     log.LoggerInterface
	Redis      redis.RedisInterface
	Mailer     mailer.MailerInterface
	IplogModel iplog.IplogModelInterface
}

type AppContextDependency struct {
	DB         database.DatabaseInterface
	Logger     log.LoggerInterface
	Redis      redis.RedisInterface
	Mailer     mailer.MailerInterface
	IplogModel iplog.IplogModelInterface
}

func NewAppContext(d AppContextDependency) AppContext {
	return AppContext{
		DB:         d.DB,
		Logger:     d.Logger,
		Redis:      d.Redis,
		Mailer:     d.Mailer,
		IplogModel: d.IplogModel,
	}
}
