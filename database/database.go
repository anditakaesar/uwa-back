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
	CreateUser(user *domain.User)

	GetRoleByName(roleName string) *domain.Role
	CreateRole(role *domain.Role)

	CreateUserRole(userRole *domain.UserRole)

	GetUserCredentialByToken(userToken string) *domain.UserCredential
	UpdateUserCredential(userCredential *domain.UserCredential)
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

func (db *Database) create(object interface{}) {
	db.mainDB.Create(object)
}

func (db *Database) save(object interface{}) {
	db.mainDB.Save(object)
}
