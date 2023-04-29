package auth

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type HandlerDependency struct {
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

type Handler struct {
	Resp       handler.ResponseInterface
	Log        log.LoggerInterface
	AppContext context.AppContext
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		Resp:       handler.NewResponse(handler.Dep{Log: d.Logger}),
		Log:        d.Logger,
		AppContext: d.AppContext,
	}
}

func (h Handler) RegisterUser() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		return h.Resp.SetOk(nil)
	}
}
