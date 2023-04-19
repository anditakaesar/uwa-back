package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/anditakaesar/uwa-back/internal/env"
)

type BearerToken = []byte

func (m Middleware) Guest(h http.Handler) http.HandlerFunc {
	return m.Group(h, false, m.ApiToken, m.AccessToken)
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

func (m Middleware) AccessToken(w http.ResponseWriter, r *http.Request) (*http.Request, *Error) {
	_, err := GetBearerToken(r)
	UserContextKey := env.UserContext()
	if err != nil {
		return nil, &Error{err, http.StatusUnauthorized}
	}

	// user, err := m.UserDomain.GetUserByToken(r.Context(), string(token))
	// if err != nil {
	// 	m.Log.Error("[Middleware][AccessToken] get user by token failed", err)
	// 	return nil, &Error{err, http.StatusUnauthorized}
	// }

	// if user == nil {
	// 	return nil, &Error{err, http.StatusUnauthorized}
	// }

	type UserProfile struct {
		Name      string
		Email     string
		CreatedAt time.Time
	}

	requestedUser := UserProfile{
		Name:      "user.Name",
		Email:     "user.Email",
		CreatedAt: time.Now(),
	}

	requestWithContext := r.WithContext(context.WithValue(r.Context(), UserContextKey, requestedUser))

	return requestWithContext, nil
}

// GetBearerToken ...
func GetBearerToken(r *http.Request) (BearerToken, error) {
	authorizationHeader := r.Header.Get("Authorization")
	splitAuthorizationHeader := strings.Split(authorizationHeader, "Bearer")

	if len(splitAuthorizationHeader) != 2 {
		return nil, errors.New("invalid authorization bearer header")
	}

	token := strings.TrimSpace(splitAuthorizationHeader[1])

	return []byte(token), nil
}
