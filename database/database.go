package database

import (
	"time"

	"github.com/anditakaesar/uwa-back/domain"
	"github.com/anditakaesar/uwa-back/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DefaultPageSize    = 10
	DefaultCurrentPage = 1
)

type DBInterface interface {
	Connect() error
	GetConnectedDB() *gorm.DB

	GetUserByUsername(username string) *domain.User
	GetOrCreateUserCredential(userCredential *domain.UserCredential, timeNow *time.Time) *domain.UserCredential
	GetUser(paging *domain.Paging) []domain.User
}

type Database struct {
	mainDB *gorm.DB
}

func NewConnection() DBInterface {
	db := &Database{}
	db.Connect()

	return db
}

func (db *Database) Connect() error {
	gormDB, err := gorm.Open(sqlite.Open(env.SqliteDBName()), &gorm.Config{})
	if err != nil {
		return err
	}

	db.mainDB = gormDB

	return nil
}

func (db *Database) GetConnectedDB() *gorm.DB {
	if db.mainDB == nil {
		db.Connect()
	}

	return db.mainDB
}

func (db *Database) GetUserByUsername(username string) *domain.User {
	var user domain.User
	db.mainDB.First(&user, "username = ?", username)
	return &user
}

func (db *Database) GetOrCreateUserCredential(userCredential *domain.UserCredential, timeNow *time.Time) *domain.UserCredential {
	db.mainDB.FirstOrCreate(userCredential, "expired_at >= ?", timeNow)
	return userCredential
}

func (db *Database) GetUser(paging *domain.Paging) []domain.User {
	var result []domain.User
	query := db.mainDB.Model(result)

	if paging != nil {
		enrichQueryWithPaging(query, paging)
	}

	query.Find(&result)

	return result
}

func enrichQueryWithPaging(query *gorm.DB, paging *domain.Paging) {
	o := 0
	p := DefaultPageSize
	var count int64
	query.Count(&count)
	paging.Count = uint64(count)
	if paging.PageSize > 0 {
		p = paging.PageSize
		if paging.CurrentPage > 0 {
			o = (paging.CurrentPage - 1) * p
		} else {
			paging.CurrentPage = DefaultCurrentPage
		}
	} else {
		paging.PageSize = p
		paging.CurrentPage = DefaultCurrentPage
	}

	query.Offset(o).Limit(p)
}
