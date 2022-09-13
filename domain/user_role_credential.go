package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique,index"`
	Password string
}

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
	User   User
	Role   Role
}

const (
	RoleAdmin  = "admin"
	RoleEditor = "editor"
)

type UserCredential struct {
	gorm.Model
	UserID       uint
	User         User
	UserToken    string
	ExpiredAt    *time.Time
	LastAccessAt *time.Time
}
