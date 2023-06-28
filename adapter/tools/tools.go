package tools

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

type ToolsRoute struct {
	Context    router.Context
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

func NewAdapter(d RouteDependecy) {
	route := ToolsRoute(d)
	route.InitEndpoints()
}

func (r ToolsRoute) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:     r.Logger,
		AppContext: r.AppContext,
	})

	r.Context.RegisterEndpoint(r.GetRequesterIpRoute(h))
}

func (r ToolsRoute) GetRequesterIpRoute(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod:    http.MethodGet,
		URLPattern:    "/tools/ip",
		Handler:       h.ReturnIP(),
		Verifications: []constants.VerificationType{},
	}
}
