package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserCredential struct {
	gorm.Model
	UserID       uint
	User         User
	UserToken    string
	ExpiredAt    *time.Time
	LastAccessAt *time.Time
}
