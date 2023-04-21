package env

import (
	"os"
	"strconv"
)

func RedisHost() string {
	return os.Getenv("RedisHost")
}

func RedisPort() string {
	return os.Getenv("RedisPort")
}

func RedisPassword() string {
	return os.Getenv("RedisPassword")
}

func RedisDB() int {
	defaultDB := 0
	if os.Getenv("RedisDB") != "" {
		parsedInt, err := strconv.Atoi(os.Getenv("RedisDB"))
		if err != nil {
			return defaultDB
		}

		return parsedInt
	}

	return defaultDB
}
