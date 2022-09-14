package services

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/thoas/go-funk"
)

type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthUser(appCtx application.Context, authParam AuthParam) (string, error) {
	var user domain.User
	var userCredential domain.UserCredential
	appCtx.DB.First(&user, "username = ?", authParam.Username)

	if funk.IsEmpty(user) {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s", authParam.Username))
		return "", errors.New("[Services][Auth] unauthorized")
	}

	authPassHash := user.Password
	ok := appCtx.Crypter.CompareHash(authPassHash, authParam.Password)
	if !ok {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s, pass: %s", authParam.Username, authParam.Password))
		return "", errors.New("[Services][Auth] unauthorized")
	}

	userCredential.User = user
	userToken, err := GenerateSecureToken(env.UserTokenLength())
	if err != nil {
		appCtx.Log.Warn(fmt.Sprintf("[Services][Auth] generate secure token failed: %d, err:%v", env.UserTokenLength(), err))
		return "", err
	}

	userCredential.UserToken = userToken
	expiredAt := appCtx.TimeNow.Add(24 * time.Hour)
	userCredential.ExpiredAt = &expiredAt
	appCtx.DB.FirstOrCreate(&userCredential, "expired_at >= ?", appCtx.TimeNow)

	return userCredential.UserToken, nil
}

func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
