package user

import (
	"github.com/anditakaesar/uwa-back/adapter/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	models.ChangesInfo
	GUID     uuid.UUID
	Username string `gorm:"unique,index"`
	Password string
	Email    string `gorm:"unique"`
}
