package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/anditakaesar/uwa-back/domain"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/utils"
	"github.com/thoas/go-funk"
)

type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthServiceInterface interface {
	AuthUser(authParam AuthParam) (string, error)
}

type AuthService struct {
	Ctx *Context
}

func NewAuthService(ctx *Context) AuthServiceInterface {
	return &AuthService{
		Ctx: ctx,
	}
}

func (as *AuthService) AuthUser(authParam AuthParam) (string, error) {
	userCredential := &domain.UserCredential{}
	now := time.Now()

	user := as.Ctx.DBI.GetUserByUsername(authParam.Username)
	crypter := utils.GetDefaultCrypter()

	if funk.IsEmpty(user) {
		as.Ctx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s", authParam.Username))
		return "", errors.New("[Services][Auth] unauthorized")
	}

	authPassHash := user.Password
	ok := crypter.CompareHash(authPassHash, authParam.Password)
	if !ok {
		as.Ctx.Log.Warn(fmt.Sprintf("[Services][Auth] auth attempt with user: %s, pass: %s", authParam.Username, authParam.Password))
		return "", errors.New("[Services][Auth] unauthorized")
	}

	userCredential.User = *user
	userToken, err := utils.GenerateSecureToken(env.UserTokenLength())
	if err != nil {
		as.Ctx.Log.Warn(fmt.Sprintf("[Services][Auth] generate secure token failed: %d, err:%v", env.UserTokenLength(), err))
		return "", err
	}

	userCredential.UserToken = userToken
	expiredAt := now.Add(24 * time.Hour)
	userCredential.ExpiredAt = &expiredAt
	userCredential = as.Ctx.DBI.GetOrCreateUserCredential(userCredential, &now)

	return userCredential.UserToken, nil
}
