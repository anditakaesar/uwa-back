package internal

import (
	applicationContext "github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/domain/healthcheck"
	"github.com/anditakaesar/uwa-back/domain/migration"
	"github.com/anditakaesar/uwa-back/domain/tool"
	"github.com/anditakaesar/uwa-back/domain/user"
)

func initDomains(m *InfraModels, s *InfraServices, i *Infrastructure, a *applicationContext.AppContext) error {
	tool.NewDomain(tool.Dependecy{
		AppContext: *a,
		Context:    s.routerSvc,
	})

	healthcheck.NewDomain(healthcheck.Dependecy{
		Context:    s.routerSvc,
		AppContext: *a,
	})

	migration.NewDomain(migration.RouteDependecy{
		Context:    s.routerSvc,
		AppContext: *a,
		DB:         i.PostgresDb,
	})

	user.NewDomain(user.Dependecy{
		Context:     s.routerSvc,
		AppContext:  *a,
		UserService: s.userSvc,
		RoleService: s.roleSvc,
	})

	return nil
}
