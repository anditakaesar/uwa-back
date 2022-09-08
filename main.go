package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/log"
	"github.com/anditakaesar/uwa-back/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger := log.BuildLogger()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	appContext := application.Context{
		Log:     logger,
		Crypter: application.BuildCustomCrypter(),
		DB:      db,
	}
	logger.Info("=====Building Routes=====")
	r := router.NewRouter(appContext)
	logger.Info(fmt.Sprintf("App running at port%v", env.Port()))
	http.ListenAndServe(env.Port(), r)
}
