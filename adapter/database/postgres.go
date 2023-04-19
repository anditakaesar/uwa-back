package database

import "gorm.io/gorm"

type DatabaseInterface interface {
	Connect() error
	Get() *gorm.DB
}
