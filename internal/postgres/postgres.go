package postgres

import (
	"fmt"

	"github.com/anditakaesar/uwa-back/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	db := &Database{}

	return db
}

func GenerateDSN(host string, user string, pwd string, dbname string, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, pwd, dbname, port)
}

func (db *Database) Connect() error {
	dsn := GenerateDSN(env.DBAddress(), env.DBUser(), env.DBPassword(), env.DBName(), env.DBPort())
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.DB = gormDB
	return nil
}

func (db *Database) Get() *gorm.DB {
	return db.DB
}
