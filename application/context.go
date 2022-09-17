package application

import (
	"time"

	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/services"
	"github.com/anditakaesar/uwa-back/utils"
	"gorm.io/gorm"
)

type Context struct {
	Log      log.LogInterface
	Crypter  utils.Crypter
	DB       *gorm.DB
	TimeNow  *time.Time
	Services Services
}

type Services struct {
	AuthService    services.AuthServiceInterface
	DBToolsService services.DBToolsServiceInterface
	UserService    services.UserServiceInterface
}
