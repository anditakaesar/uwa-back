package env

import "os"

func EmailFrom() string {
	return os.Getenv("EmailFrom")
}

func EmailSmtpUser() string {
	return os.Getenv("EmailSmtpUser")
}

func EmailPassword() string {
	return os.Getenv("EmailPassword")
}

func EmailHost() string {
	return os.Getenv("EmailHost")
}

func EmailPort() string {
	return os.Getenv("EmailPort")
}
