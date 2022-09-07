package services

import (
	"errors"
	"fmt"

	"github.com/anditakaesar/uwa-back/application"
)

type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthUser(appCtx application.Context, authParam AuthParam) error {
	users := map[string]string{
		"anditakaesar": "42a9798b99d4afcec9995e47a1d246b98ebc96be7a732323eee39d924006ee1d",
		"usertwo":      "4a34219b9b4f66a1932428cccae29846b4c5fce07ce7c390b9c5b27e0fea378d",
	}

	authPassHash := users[authParam.Username]
	if authPassHash == "" {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s", authParam.Username))
		return errors.New("[Services][Auth] unauthorized")
	}

	ok := appCtx.Crypter.CompareHash(authPassHash, authParam.Password)
	if !ok {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s, pass: %s", authParam.Username, authParam.Password))
		return errors.New("[Services][Auth] unauthorized")
	}

	return nil
}
