package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/internal/env"
	"go.uber.org/zap"
)

type BearerToken = []byte

func (m Middleware) Guest(h http.Handler) http.HandlerFunc {
	return m.Group(h, false, m.ApiToken)
}

func (m Middleware) ApiToken(w http.ResponseWriter, r *http.Request) (*http.Request, *Error) {
	token, err := GetBearerToken(r)

	if err != nil {
		return nil, &Error{err, http.StatusUnauthorized}
	}

	if env.APIToken() != string(token) {
		return nil, &Error{errors.New("invalid api token"), http.StatusUnauthorized}
	}

	return r, nil
}

func (m Middleware) JWT(w http.ResponseWriter, r *http.Request) (*http.Request, *Error) {
	token, err := GetBearerToken(r)
	unauthorizedErr := &Error{err, http.StatusUnauthorized}
	if err != nil {
		m.Log.Error("error getting token", err)
		return nil, unauthorizedErr
	}
	claims, err := m.UtilInterface.ValidateAndGetClaims(string(token))
	if err != nil {
		m.Log.Warning("error validating token", zap.Error(err))
		return nil, unauthorizedErr
	}

	requestWithContext := r.WithContext(context.WithValue(r.Context(), env.JWTClaimsKey, claims))
	return requestWithContext, nil
}

// GetBearerToken returns the bearer token from the Authorization header of an HTTP request.
//
// It takes an HTTP request object as a parameter and returns the bearer token and an error if any.
func GetBearerToken(r *http.Request) (BearerToken, error) {
	authorizationHeader := r.Header.Get("Authorization")
	splitAuthorizationHeader := strings.Split(authorizationHeader, "Bearer")

	if len(splitAuthorizationHeader) != 2 {
		return nil, errors.New("invalid authorization bearer header")
	}

	token := strings.TrimSpace(splitAuthorizationHeader[1])

	return []byte(token), nil
}
