package services

import (
	"errors"
	"fmt"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/thoas/go-funk"
)

type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthUser(appCtx application.Context, authParam AuthParam) error {
	var user domain.User
	appCtx.DB.First(&user, "username = ?", authParam.Username)

	if funk.IsEmpty(user) {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s", authParam.Username))
		return errors.New("[Services][Auth] unauthorized")
	}

	authPassHash := user.Password
	ok := appCtx.Crypter.CompareHash(authPassHash, authParam.Password)
	if !ok {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s, pass: %s", authParam.Username, authParam.Password))
		return errors.New("[Services][Auth] unauthorized")
	}

	return nil
}
