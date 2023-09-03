package internal

import (
	"github.com/anditakaesar/uwa-back/adapter/httpserver"
	"github.com/anditakaesar/uwa-back/adapter/mailer"
	"github.com/anditakaesar/uwa-back/adapter/middleware"
	"github.com/anditakaesar/uwa-back/adapter/models/iplog"
	roleModel "github.com/anditakaesar/uwa-back/adapter/models/role"
	userModel "github.com/anditakaesar/uwa-back/adapter/models/user"
	"github.com/anditakaesar/uwa-back/adapter/util"
	applicationContext "github.com/anditakaesar/uwa-back/application/context"
	roleSvc "github.com/anditakaesar/uwa-back/application/services/role"
	routerSvc "github.com/anditakaesar/uwa-back/application/services/router"
	userSvc "github.com/anditakaesar/uwa-back/application/services/user"
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
	userModel  userModel.UserModelInterface
	roleModel  roleModel.RoleModelInterface
}

func InitInfraModels(db *postgres.Database) *InfraModels {
	return &InfraModels{
		iplogModel: iplog.NewIplogModel(db),
		userModel:  userModel.NewUserModel(db),
		roleModel:  roleModel.NewRoleModel(db),
	}
}

type InfraServices struct {
	routerSvc routerSvc.Context
	userSvc   userSvc.UserSeviceInterface
	roleSvc   roleSvc.RoleSeviceInterface
}

func InitInfraServices(i *Infrastructure, m *InfraModels) *InfraServices {
	routerService := routerSvc.NewService(
		i.HttpServer.Server, httpserver.RouterHelper{}, i.Middleware, "/api")
	return &InfraServices{
		routerSvc: routerService,
		userSvc:   userSvc.NewUserService(m.userModel),
		roleSvc:   roleSvc.NewRoleService(m.roleModel),
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

	// prepare util interface
	utilInt := util.NewUtilInterface()

	i.InternalRouter = way.NewRouter(
		way.NotFoundHandlerWithIpLogging(i.Logger, infraModels.iplogModel),
	)

	i.Middleware = middleware.NewAdapter(middleware.Middleware{
		Client:        i.Client,
		Log:           i.Logger,
		UtilInterface: utilInt,
		IplogModel:    infraModels.iplogModel,
	})

	i.HttpServer = httpserver.NewAdapter(&httpserver.Adapter{
		Log:    i.Logger,
		Router: i.InternalRouter,
	})

	// prepare the services
	infraServices := InitInfraServices(i, infraModels)

	// appContext
	appContext := applicationContext.NewAppContext(applicationContext.AppContextDependency{
		Logger:        i.Logger,
		Redis:         i.Redis,
		Mailer:        i.Mailer,
		IplogModel:    infraModels.iplogModel,
		UtilInterface: utilInt,
	})

	// initialize domain
	err = initDomains(infraModels, infraServices, i, &appContext)
	if err != nil {
		return err
	}

	return nil
}
