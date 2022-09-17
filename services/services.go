package services

import (
	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/log"
)

type Context struct {
	Log log.LogInterface
	DBI database.DBInterface
}
