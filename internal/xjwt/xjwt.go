package xjwt

import (
	"errors"
	"os"
	"time"

	"github.com/anditakaesar/uwa-back/internal/env"
	goJwt "github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	goJwt.RegisteredClaims
}

func GenerateNewToken(claims CustomClaims, duration time.Duration) (string, error) {
	claims.ExpiresAt = goJwt.NewNumericDate(time.Now().Add(duration))
	newToken := goJwt.NewWithClaims(goJwt.SigningMethodRS256, claims)
	privateKey, err := os.ReadFile(env.JWTPrivateFile())
	if err != nil {
		return "", err
	}
	key, err := goJwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}
	signedString, err := newToken.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ValidateAndGetClaims(tokenString string) (CustomClaims, error) {
	token, err := goJwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *goJwt.Token) (interface{}, error) {
		publicKey, err := os.ReadFile(env.JWTPublicFile())
		if err != nil {
			return nil, err
		}
		key, err := goJwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	if err != nil {
		return CustomClaims{}, err
	}
	if token.Valid {
		return *token.Claims.(*CustomClaims), nil
	}
	return CustomClaims{}, errors.New("invalid token")
}
