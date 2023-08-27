package internal

import (
	"github.com/anditakaesar/uwa-back/adapter/httpserver"
	"github.com/anditakaesar/uwa-back/adapter/mailer"
	"github.com/anditakaesar/uwa-back/adapter/middleware"
	"github.com/anditakaesar/uwa-back/adapter/models/iplog"
	applicationContext "github.com/anditakaesar/uwa-back/application/context"
	routerSvc "github.com/anditakaesar/uwa-back/application/services/router"
	"github.com/anditakaesar/uwa-back/internal/client"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/anditakaesar/uwa-back/internal/postgres"
	"github.com/anditakaesar/uwa-back/internal/redis"
	"github.com/anditakaesar/uwa-back/internal/way"
)

type Infrastructure struct {
	Client         *client.Client
	Logger         log.LoggerInterface
	PostgresDb     *postgres.Database
	InternalRouter *way.Router
	Redis          *redis.InternalRedis
	Mailer         mailer.MailerInterface
	Middleware     middleware.Middleware
	HttpServer     *httpserver.Router
}

func NewInfrastructure() *Infrastructure {
	newClient := client.New()
	logglyCore := log.NewLogglyZapCore(log.NewLogglyLogWriter(
		log.LogglyLogWriterDependency{
			HttpClient:    newClient.HttpClient,
			BaseUrl:       env.LogglyBaseUrl(),
			CustomerToken: env.LogglyToken(),
			Tag:           env.LogglyTag(),
		},
	))
	database := postgres.NewDatabase()
	internalLogger := log.BuildNewLogger(logglyCore)
	internalRedis := redis.NewInternalRedis()
	internalMailer := mailer.NewMailerAdapter(internalLogger)

	return &Infrastructure{
		Client:     newClient,
		Logger:     internalLogger,
		PostgresDb: database,
		Redis:      internalRedis,
		Mailer:     internalMailer,
	}
}

type InfraModels struct {
	iplogModel iplog.IplogModelInterface
}

func InitInfraModels(db *postgres.Database) *InfraModels {
	return &InfraModels{
		iplogModel: iplog.NewIplogModel(db),
	}
}

type InfraServices struct {
	routerSvc routerSvc.Context
}

func InitInfraServices(i *Infrastructure) *InfraServices {
	routerService := routerSvc.NewService(
		i.HttpServer.Server, httpserver.RouterHelper{}, i.Middleware, "/api")
	return &InfraServices{
		routerSvc: routerService,
	}
}

func (i *Infrastructure) Init() error {
	err := i.PostgresDb.Connect()
	if err != nil {
		return err
	}

	err = i.Redis.Connect()
	if err != nil {
		return err
	}

	// prepare the models
	infraModels := InitInfraModels(i.PostgresDb)

	i.InternalRouter = way.NewRouter(
		way.NotFoundHandlerWithIpLogging(i.Logger, infraModels.iplogModel),
	)

	i.Middleware = middleware.NewAdapter(middleware.Middleware{
		Client:     i.Client,
		Log:        i.Logger,
		IplogModel: infraModels.iplogModel,
	})

	i.HttpServer = httpserver.NewAdapter(&httpserver.Adapter{
		Log:    i.Logger,
		Router: i.InternalRouter,
	})

	// prepare the services
	infraServices := InitInfraServices(i)

	// appContext
	appContext := applicationContext.NewAppContext(applicationContext.AppContextDependency{
		DB:         i.PostgresDb,
		Logger:     i.Logger,
		Redis:      i.Redis,
		Mailer:     i.Mailer,
		IplogModel: infraModels.iplogModel,
	})

	// initialize domain
	err = initDomains(infraModels, infraServices, i, &appContext)
	if err != nil {
		return err
	}

	return nil
}
