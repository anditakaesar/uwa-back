package context

import (
	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/adapter/redis"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type AppContext struct {
	DB     database.DatabaseInterface
	Logger log.LoggerInterface
	Redis  redis.RedisInterface
}

type AppContextDependency struct {
	DB     database.DatabaseInterface
	Logger log.LoggerInterface
	Redis  redis.RedisInterface
}

func NewAppContext(d AppContextDependency) AppContext {
	return AppContext{
		DB:     d.DB,
		Logger: d.Logger,
		Redis:  d.Redis,
	}
}
