package migration

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/application/services/router"
	"github.com/anditakaesar/uwa-back/internal/constants"
)

type RouteDependecy struct {
	Context    router.Context
	AppContext context.AppContext
}

type MigrationRoute struct {
	Context    router.Context
	AppContext context.AppContext
}

func NewDomain(d RouteDependecy) {
	route := MigrationRoute(d)
	route.InitEndpoints()
}

func (r MigrationRoute) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		AppContext: r.AppContext,
	})

	r.Context.RegisterEndpoint(r.PostUpMigration(h))
	r.Context.RegisterEndpoint(r.PostDoMigration(h))
	r.Context.RegisterEndpoint(r.GetListMigration(h))
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

func (r MigrationRoute) PostDoMigration(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/migration/do",
		Handler:    h.DoMigration(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r MigrationRoute) GetListMigration(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/migration",
		Handler:    h.GetAvailableMigration(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}
