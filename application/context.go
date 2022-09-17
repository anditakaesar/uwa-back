package application

import (
	"time"

	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/services"
	"gorm.io/gorm"
)

type Context struct {
	Log      log.LogInterface
	DBI      database.DBInterface
	DB       *gorm.DB
	TimeNow  *time.Time
	Services Services
}

type Services struct {
	AuthService    services.AuthServiceInterface
	DBToolsService services.DBToolsServiceInterface
	UserService    services.UserServiceInterface
}
