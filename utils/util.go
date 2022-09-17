package utils

import (
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) string {
	userToken := r.Header.Get("Authorization")
	splitToken := strings.Split(userToken, "Bearer ")

	if len(splitToken) > 1 {
		return splitToken[1]
	}

	return ""
}
