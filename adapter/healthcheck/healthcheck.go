package healthcheck

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

	r.Context.RegisterEndpoint(r.GetHealtCheckStatus(h))
	r.Context.RegisterEndpoint(r.PostTestMail(h))
	r.Context.RegisterRootEndpoint(r.ServeStatic(h))
}

func (r HealthCheckRoute) GetHealtCheckStatus(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodGet,
		URLPattern: "/healthcheck/status",
		Handler:    h.CheckHealth(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r HealthCheckRoute) PostTestMail(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod: http.MethodPost,
		URLPattern: "/healthcheck/sendmail",
		Handler:    h.SendTestMail(),
		Verifications: []constants.VerificationType{
			constants.VerificationTypeConstants.APIToken,
		},
	}
}

func (r HealthCheckRoute) ServeStatic(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod:    http.MethodGet,
		URLPattern:    "/",
		Handler:       nil,
		Verifications: []constants.VerificationType{},
	}
}
