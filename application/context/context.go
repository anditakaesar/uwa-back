package context

import (
	"github.com/anditakaesar/uwa-back/adapter/mailer"
	"github.com/anditakaesar/uwa-back/adapter/models/iplog"
	"github.com/anditakaesar/uwa-back/adapter/redis"
	"github.com/anditakaesar/uwa-back/adapter/util"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type AppContext struct {
	Logger        log.LoggerInterface
	Redis         redis.RedisInterface
	Mailer        mailer.MailerInterface
	IplogModel    iplog.IplogModelInterface
	UtilInterface util.UtilInterface
}

type AppContextDependency struct {
	Logger        log.LoggerInterface
	Redis         redis.RedisInterface
	Mailer        mailer.MailerInterface
	IplogModel    iplog.IplogModelInterface
	UtilInterface util.UtilInterface
}

func NewAppContext(d AppContextDependency) AppContext {
	return AppContext(d)
}
