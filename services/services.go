package services

import (
	"time"

	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/utils"
	"gorm.io/gorm"
)

type Context struct {
	Log     log.LogInterface
	Crypter utils.Crypter
	DBI     database.DBInterface
	DB      *gorm.DB
	TimeNow *time.Time
}
