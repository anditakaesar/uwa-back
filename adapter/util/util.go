package util

import (
	"context"

	"github.com/anditakaesar/uwa-back/internal/util"
	"github.com/anditakaesar/uwa-back/internal/xjwt"
)

type UtilInterface interface {
	IsPhoneNumberValid(phoneNumber string) bool
	GenerateNewToken(claims xjwt.CustomClaims) (string, error)
	ValidateAndGetClaims(tokenString string) (xjwt.CustomClaims, error)
	GetClaimsFromContext(ctx context.Context) (xjwt.CustomClaims, error)
}

func NewUtilInterface() UtilInterface {
	return &util.InternalUtil{}
}
