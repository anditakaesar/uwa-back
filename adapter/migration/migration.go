package migration

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/application/services/router"
	"github.com/anditakaesar/uwa-back/internal/constants"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type RouteDependecy struct {
	Context    router.Context
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

type MigrationRoute struct {
	Context    router.Context
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

func NewAdapter(d RouteDependecy) {
	route := MigrationRoute(d)
	route.InitEndpoints()
}

func (r MigrationRoute) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:     r.Logger,
		AppContext: r.AppContext,
	})

	r.Context.RegisterEndpoint(r.PostUpMigration(h))
}

func (r MigrationRoute) PostUpMigration(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/migration/up",
		Handler:    h.UpMigration(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}
