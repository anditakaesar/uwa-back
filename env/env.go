package env

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DefaultPort            = ":5000"
	DefaultAppName         = "AppName Unset"
	DefaultEnv             = "development"
	DefaultSqliteDBName    = "test.db"
	DefaultUserTokenLength = "128"
)

const (
	EnvProduction = "production"
)

func Port() string {
	port := os.Getenv("Port")
	if valid, _ := regexp.MatchString(`\:\d{4}`, port); valid {
		return port
	}

	return DefaultPort
}

func AppName() string {
	appName := os.Getenv("AppName")
	appName = strings.Trim(appName, " ")
	if appName != "" {
		return appName
	}

	return DefaultAppName
}

func AppEnv() string {
	env := os.Getenv("AppEnv")
	if env != "" {
		return env
	}
	return DefaultEnv
}

func AppToken() string {
	return os.Getenv("AppToken")
}

func SqliteDBName() string {
	dbname := os.Getenv("SqliteDBName")
	if dbname == "" {
		return DefaultSqliteDBName
	}

	return dbname
}

func UserTokenLength() int {
	lengthStr := os.Getenv("UserTokenLength")
	if valid, _ := regexp.MatchString(`\d{4}`, lengthStr); valid {
		userTokenLength, _ := strconv.Atoi(lengthStr)
		return userTokenLength
	}

	userTokenLength, _ := strconv.Atoi(DefaultUserTokenLength)
	return userTokenLength
}
