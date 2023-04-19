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

func DBUser() string {
	return os.Getenv("DBUser")
}

func DBPassword() string {
	return os.Getenv("DBPassword")
}

func DBAddress() string {
	return os.Getenv("DBAddress")
}

func DBDatabase() string {
	return os.Getenv("DBDatabase")
}

func DBUrl() string {
	return os.Getenv("DBUrl")
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
