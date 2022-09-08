package services

import (
	"fmt"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/thoas/go-funk"
)

const (
	errorAutoMigrateModel = "[services][AutoMigrate] error on models %s, err: %v"
)

type User domain.User

func AutoMigrate(appCtx application.Context) error {
	// user table
	err := appCtx.DB.AutoMigrate(&User{})
	if err != nil {
		appCtx.Log.Fatal(fmt.Sprintf(errorAutoMigrateModel, "User", err))
		return err
	}

	return nil
}

func SeedUser(appCtx application.Context) error {
	var user User
	users := []User{
		{
			Username: "anditakaesar",
			Password: "42a9798b99d4afcec9995e47a1d246b98ebc96be7a732323eee39d924006ee1d",
		},
		{
			Username: "usertwo",
			Password: "4a34219b9b4f66a1932428cccae29846b4c5fce07ce7c390b9c5b27e0fea378d",
		},
	}

	for _, u := range users {
		appCtx.DB.First(&user, "username = ?", u.Username)
		if funk.IsEmpty(user) && user.Username != u.Username {
			appCtx.DB.Create(&u)
		}
	}

	return nil
}
