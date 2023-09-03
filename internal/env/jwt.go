package env

import (
	"os"
	"strconv"
	"time"
)

var JWTClaimsKey ContextKey = ContextKey(JWTClaimsToken())

func JWTPrivateFile() string {
	return os.Getenv("JWTPrivateFile")
}

func JWTPublicFile() string {
	return os.Getenv("JWTPublicFile")
}

func JWTExpiresIn() time.Duration {
	expirationInHours := 720 * time.Hour
	if os.Getenv("JWTExpiresIn") != "" {
		parsedInt, err := strconv.Atoi(os.Getenv("JWTExpiresIn"))
		if err != nil {
			return expirationInHours
		}
		return time.Duration(parsedInt) * time.Hour
	}

	return expirationInHours
}

func JWTClaimsToken() string {
	return os.Getenv("JWTClaimsToken")
}
