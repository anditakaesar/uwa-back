package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/database"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/router"
	"github.com/anditakaesar/uwa-back/services"
	"github.com/anditakaesar/uwa-back/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	logger := log.BuildLogger()

	crypter := utils.BuildCustomCrypter()
	db := database.NewConnection()
	err := db.Connect()
	if err != nil {
		panic("failed to connect db")
	}

	serviceCtx := services.Context{
		Log:     logger,
		Crypter: crypter,
		DB:      db.GetConnectedDB(),
		DBI:     db,
		TimeNow: &now,
	}

	appContext := application.Context{
		Log:     logger,
		Crypter: crypter,
		DB:      db.GetConnectedDB(),
		DBI:     db,
		TimeNow: &now,
		Services: application.Services{
			AuthService:    services.NewAuthService(&serviceCtx),
			DBToolsService: services.NewDBToolsService(&serviceCtx),
			UserService:    services.NewUserService(&serviceCtx),
		},
	}
	logger.Info("=====Building Routes=====")
	r := router.NewRouter(appContext)
	logger.Info(fmt.Sprintf("App running at port%v", env.Port()))
	http.ListenAndServe(env.Port(), r)
}
