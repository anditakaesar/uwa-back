package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/anditakaesar/uwa-back/adapter/httpserver"
	"github.com/anditakaesar/uwa-back/adapter/middleware"
	"github.com/anditakaesar/uwa-back/internal/client"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/anditakaesar/uwa-back/internal/way"

	routerSvc "github.com/anditakaesar/uwa-back/application/services/router"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting application %s", env.AppName())
	}
}

func run() error {
	addr := flag.String("addr", env.AppPort(), "http service address")
	flag.Parse()

	internalLogger := log.BuildNewLogger()
	internalRouter := way.NewRouter()
	httpServerAdapter := httpserver.NewAdapter(&httpserver.Adapter{
		Log:    internalLogger,
		Router: internalRouter,
	})

	newClient := client.New()
	middlewareAdapter := middleware.NewAdapter(middleware.Middleware{
		Client: newClient,
		Log:    internalLogger,
	})

	routerService := routerSvc.NewService(httpServerAdapter.Server, httpserver.RouterHelper{}, middlewareAdapter, "/api")
	routerService.InitOptionsRoute()

	//-------------domain here

	//-------------service or use cases

	//-------------adapters

	//--------------------------

	internalRouter.ReArrange()

	s := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      internalRouter,
		Addr:         *addr,
	}

	internalLogger.Info("+++++++++++++++++++++++++++++++++++++++++++++++++++")
	internalLogger.Info(fmt.Sprint(env.AppName(), "backend service started on port", env.AppPort()))
	internalLogger.Info("+++++++++++++++++++++++++++++++++++++++++++++++++++")

	return s.ListenAndServe()
}
