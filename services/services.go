package services

import (
	"time"

	"github.com/anditakaesar/uwa-back/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Context struct {
	Log     *zap.Logger
	Crypter utils.Crypter
	DB      *gorm.DB
	TimeNow *time.Time
}
