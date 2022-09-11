package services

import (
	"fmt"
	"reflect"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/thoas/go-funk"
)

const (
	errorAutoMigrateModel = "[services][AutoMigrate] error on models %s, err: %v"
)

type User domain.User
type Role domain.Role

func AutoMigrate(appCtx application.Context) error {
	tableDomains := []interface{}{
		&User{}, &Role{},
	}

	for _, domain := range tableDomains {
		err := appCtx.DB.AutoMigrate(domain)
		if err != nil {
			appCtx.Log.Fatal(fmt.Sprintf(errorAutoMigrateModel, reflect.TypeOf(domain), err))
			return err
		}
	}

	return nil
}

type SeedFunc func(appCtx application.Context) error

func Seed(appCtx application.Context, table string) error {
	availableSeeds := map[string]SeedFunc{
		"user": SeedUser,
		"role": SeedRole,
	}

	if fn, ok := availableSeeds[table]; ok {
		return fn(appCtx)
	}

	appCtx.Log.Warn(fmt.Sprintf("[services][Seed] seed attempt using table name:%s", table))
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

func SeedRole(appCtx application.Context) error {
	var role Role
	roles := []Role{
		{
			Name:        "admin",
			Description: "main admin that can access all",
		},
		{
			Name:        "editor",
			Description: "main editor that can only access articles",
		},
	}

	for _, r := range roles {
		appCtx.DB.First(&role, "name = ?", r.Name)
		if funk.IsEmpty(role) && role.Name != r.Name {
			appCtx.DB.Create(&r)
		}
	}

	return nil
}
