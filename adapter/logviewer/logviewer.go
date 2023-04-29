package logviewer

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

type HealthCheckRoute struct {
	Context    router.Context
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

func NewAdapter(d RouteDependecy) {
	route := HealthCheckRoute(d)
	route.InitEndpoints()
}

func (r HealthCheckRoute) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:     r.Logger,
		AppContext: r.AppContext,
	})

	r.Context.RegisterEndpoint(r.GetAvailableLogs(h))
	r.Context.RegisterEndpoint(r.GetLog(h))
}

func (r HealthCheckRoute) GetAvailableLogs(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/logviewer",
		Handler:    h.GetAvailableLogs(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r HealthCheckRoute) GetLog(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/logviewer/:logfile",
		Handler:    h.GetLog(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}
