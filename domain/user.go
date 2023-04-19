package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GUID     uuid.UUID
	Username string `gorm:"unique,index"`
	Password string
	Email    string `gorm:"unique"`
}
