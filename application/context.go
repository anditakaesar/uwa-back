package application

import (
	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/services"
)

type Context struct {
	Log log.LogInterface
	DBI database.DBInterface
	Services
}

type Services struct {
	AuthService    services.AuthServiceInterface
	DBToolsService services.DBToolsServiceInterface
	UserService    services.UserServiceInterface
}
