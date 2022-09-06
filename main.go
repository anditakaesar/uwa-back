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
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger := log.BuildLogger()
	appContext := application.Context{
		Log: logger,
	}
	logger.Info("=====Building Routes=====")
	r := router.NewRouter(appContext)
	logger.Info(fmt.Sprintf("App running at port%v", env.Port()))
	http.ListenAndServe(env.Port(), r)
}
