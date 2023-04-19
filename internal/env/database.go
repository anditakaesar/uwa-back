package env

import "os"

func DBUser() string {
	return os.Getenv("DBUser")
}

func DBPassword() string {
	return os.Getenv("DBPassword")
}

func DBAddress() string {
	return os.Getenv("DBAddress")
}

func DBName() string {
	return os.Getenv("DBName")
}

func DBPort() string {
	return os.Getenv("DBPort")
}
