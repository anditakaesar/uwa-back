package tool

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/handler"
)

type HandlerDependency struct {
	AppContext context.AppContext
}

type Handler struct {
	Resp       handler.ResponseInterface
	AppContext context.AppContext
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		Resp:       handler.NewResponse(handler.Dep{Log: d.AppContext.Logger}),
		AppContext: d.AppContext,
	}
}

func (h Handler) ReturnIP() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		currentIp := r.Header.Get(env.IPHeaderKey())
		return h.Resp.SetOk(currentIp)
	}
}
