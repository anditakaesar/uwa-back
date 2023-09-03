package util

import (
	"context"
	"errors"
	"regexp"

	"github.com/anditakaesar/uwa-back/internal/constants"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/xjwt"
)

type InternalUtil struct{}

func (u *InternalUtil) IsPhoneNumberValid(phoneNumber string) bool {
	return regexp.MustCompile(constants.PhoneNumberRegex).MatchString(phoneNumber)
}

func (u *InternalUtil) GenerateNewToken(claims xjwt.CustomClaims) (string, error) {
	return xjwt.GenerateNewToken(claims, env.JWTExpiresIn())
}

func (u *InternalUtil) ValidateAndGetClaims(tokenString string) (xjwt.CustomClaims, error) {
	return xjwt.ValidateAndGetClaims(tokenString)
}

func (u *InternalUtil) GetClaimsFromContext(ctx context.Context) (xjwt.CustomClaims, error) {
	rawClaims := ctx.Value(env.JWTClaimsKey)
	claims, ok := rawClaims.(xjwt.CustomClaims)
	if !ok {
		return xjwt.CustomClaims{}, errors.New("invalid claims")
	}

	return claims, nil
}
