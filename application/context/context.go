package context

import (
	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type AppContext struct {
	DB     database.DatabaseInterface
	Logger log.LoggerInterface
}

func NewAppContext(dbi database.DatabaseInterface, logger log.LoggerInterface) *AppContext {
	return &AppContext{
		DB:     dbi,
		Logger: logger,
	}
}
