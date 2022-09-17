package services

import (
	"fmt"
	"reflect"

	"github.com/anditakaesar/uwa-back/domain"
	"github.com/thoas/go-funk"
)

const (
	errorAutoMigrateModel = "[services][AutoMigrate] error on models %s, err: %v"
)

type DBToolsServiceInterface interface {
	AutoMigrate() error
	Seed(table string) error
}

type DBToolsService struct {
	Ctx *Context
}

func NewDBToolsService(ctx *Context) DBToolsServiceInterface {
	return &DBToolsService{
		Ctx: ctx,
	}
}

func (dbt *DBToolsService) AutoMigrate() error {
	var user domain.User
	var role domain.Role
	var userRole domain.UserRole
	var userCredential domain.UserCredential
	tableDomains := []interface{}{
		&user, &role, &userRole, &userCredential,
	}

	for _, domain := range tableDomains {
		err := dbt.Ctx.DBI.GetConnectedDB().AutoMigrate(domain)
		if err != nil {
			dbt.Ctx.Log.Fatal(fmt.Sprintf(errorAutoMigrateModel, reflect.TypeOf(domain), err))
			return err
		}
	}

	return nil
}

type SeedFunc func(Ctx *Context) error

func (dbt *DBToolsService) Seed(table string) error {
	availableSeeds := map[string]SeedFunc{
		"user":     SeedUser,
		"role":     SeedRole,
		"userrole": SeedUserRole,
	}

	if fn, ok := availableSeeds[table]; ok {
		return fn(dbt.Ctx)
	}

	dbt.Ctx.Log.Warn(fmt.Sprintf("[services][Seed] seed attempt using table name:%s", table))
	return nil
}

func SeedUser(ctx *Context) error {
	users := []domain.User{
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
		user := ctx.DBI.GetUserByUsername(u.Username)
		if funk.IsEmpty(user) && user.Username != u.Username {
			ctx.DBI.CreateUser(&u)
		}
	}

	return nil
}

func SeedRole(ctx *Context) error {
	roles := []domain.Role{
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
		role := ctx.DBI.GetRoleByName(r.Name)
		if funk.IsEmpty(role) && role.Name != r.Name {
			ctx.DBI.CreateRole(&r)
		}
	}

	return nil
}

func SeedUserRole(ctx *Context) error {
	var userRole1 domain.UserRole
	user1 := ctx.DBI.GetUserByUsername("anditakaesar")
	role1 := ctx.DBI.GetRoleByName("admin")
	userRole1.User = *user1
	userRole1.Role = *role1
	ctx.DBI.CreateUserRole(&userRole1)

	return nil
}
