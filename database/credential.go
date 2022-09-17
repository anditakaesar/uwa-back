package database

import "github.com/anditakaesar/uwa-back/domain"

func (db *Database) GetUserCredentialByToken(userToken string) *domain.UserCredential {
	var userCredential domain.UserCredential
	db.mainDB.First(&userCredential, "user_token = ?", userToken)

	return &userCredential
}

func (db *Database) UpdateUserCredential(userCredential *domain.UserCredential) {
	db.save(userCredential)
}
