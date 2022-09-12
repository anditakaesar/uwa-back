package domain

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
	User   User
	Role   Role
}
