package env

import (
	"os"
)

type UserContextKey string

func AppName() string {
	return os.Getenv("AppName")
}

func AppPort() string {
	return os.Getenv("AppPort")
}

func APIToken() string {
	return os.Getenv("APIToken")
}

func Env() string {
	return os.Getenv("Env")
}

func CorsOrigin() string {
	if os.Getenv("CorsOrigin") != "" {
		return os.Getenv("CorsOrigin")
	}

	return "*"
}

func UserContext() UserContextKey {
	if os.Getenv("UserContext") != "" {
		return UserContextKey(os.Getenv("UserContext"))
	}

	return "defaultUserContext"
}

func IPHeaderKey() string {
	if os.Getenv("IPHeaderKey") != "" {
		return os.Getenv("IPHeaderKey")
	}

	return "X-Header"
}
