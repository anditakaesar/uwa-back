package env

import (
	"os"
)

type UserContextKey string
type ContextKey string

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

func HostURL() string {
	return os.Getenv("HostURL")
}

func LogglyBaseUrl() string {
	return os.Getenv("LogglyBaseUrl")
}

func LogglyToken() string {
	return os.Getenv("LogglyToken")
}

func LogglyTag() string {
	if os.Getenv("LogglyTag") != "" {
		return os.Getenv("LogglyTag")
	}

	return os.Getenv("Env")
}
