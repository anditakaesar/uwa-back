package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	adapterDB "github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/adapter/healthcheck"
	"github.com/anditakaesar/uwa-back/adapter/httpserver"
	"github.com/anditakaesar/uwa-back/adapter/logviewer"
	"github.com/anditakaesar/uwa-back/adapter/mailer"
	"github.com/anditakaesar/uwa-back/adapter/middleware"
	"github.com/anditakaesar/uwa-back/adapter/migration"
	"github.com/anditakaesar/uwa-back/internal/client"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/anditakaesar/uwa-back/internal/postgres"
	internalRedis "github.com/anditakaesar/uwa-back/internal/redis"
	"github.com/anditakaesar/uwa-back/internal/way"

	applicationContext "github.com/anditakaesar/uwa-back/application/context"
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
	internalRouter := way.NewRouter(internalLogger)
	httpServerAdapter := httpserver.NewAdapter(&httpserver.Adapter{
		Log:    internalLogger,
		Router: internalRouter,
	})

	database := postgres.NewDatabase()
	dbErr := database.Connect()
	if dbErr != nil {
		return dbErr
	}

	newClient := client.New()
	middlewareAdapter := middleware.NewAdapter(middleware.Middleware{
		Client:    newClient,
		Log:       internalLogger,
		IpLogRepo: adapterDB.NewIpLogRepository(database),
	})

	routerService := routerSvc.NewService(httpServerAdapter.Server, httpserver.RouterHelper{}, middlewareAdapter, "/api")
	routerService.InitOptionsRoute()

	redis := internalRedis.NewInternalRedis()
	redisErr := redis.Connect()
	if redisErr != nil {
		return redisErr
	}

	internalMailer := mailer.NewMailerAdapter(internalLogger)

	appContext := applicationContext.NewAppContext(applicationContext.AppContextDependency{
		DB:        database,
		Logger:    internalLogger,
		Redis:     redis,
		Mailer:    internalMailer,
		IpLogRepo: adapterDB.NewIpLogRepository(database),
	})

	//-------------domain here

	//-------------service or use cases

	//-------------adapters

	migration.NewAdapter(migration.RouteDependecy{
		Context:    routerService,
		Logger:     internalLogger,
		AppContext: appContext,
	})

	healthcheck.NewAdapter(healthcheck.RouteDependecy{
		Context:    routerService,
		Logger:     internalLogger,
		AppContext: appContext,
	})

	logviewer.NewAdapter(logviewer.RouteDependecy{
		Context:    routerService,
		Logger:     internalLogger,
		AppContext: appContext,
	})

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
