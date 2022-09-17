package services

import (
	"time"

	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/log"
	"gorm.io/gorm"
)

type Context struct {
	Log     log.LogInterface
	DBI     database.DBInterface
	DB      *gorm.DB
	TimeNow *time.Time
}
