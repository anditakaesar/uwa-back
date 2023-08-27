package tool

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/application/services/router"
	"github.com/anditakaesar/uwa-back/internal/constants"
)

type Dependecy struct {
	Context    router.Context
	AppContext context.AppContext
}

type Route struct {
	Context    router.Context
	AppContext context.AppContext
}

func NewDomain(d Dependecy) {
	route := Route(d)
	route.InitEndpoints()
}

func (r Route) InitEndpoints() {
	h := NewHandler(HandlerDependency{
		AppContext: r.AppContext,
	})

	r.Context.RegisterEndpoint(r.GetRequesterIpRoute(h))
}

func (r Route) GetRequesterIpRoute(h Handler) router.EndpointInfo {
	return router.EndpointInfo{
		HTTPMethod:    http.MethodGet,
		URLPattern:    "/tools/ip",
		Handler:       h.ReturnIP(),
		Verifications: []constants.VerificationType{},
	}
}
