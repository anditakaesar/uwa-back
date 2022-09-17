package database

import (
	"time"

	"github.com/anditakaesar/uwa-back/domain"
)

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

func (db *Database) CreateUser(user *domain.User) {
	db.create(user)
}
