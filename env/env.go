package env

import (
	"os"
	"regexp"
)

func Port() string {
	port := os.Getenv("Port")
	if valid, _ := regexp.MatchString(`\:\d{4}`, port); valid {
		return port
	}

	return ":5000"
}

func AppName() string {
	return os.Getenv("AppName")
}
