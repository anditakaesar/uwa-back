package httpserver

import (
	"context"
	"net/http"

	"github.com/anditakaesar/uwa-back/internal/way"

	"github.com/anditakaesar/uwa-back/internal/log"
)

type RouterHelper struct{}

func (r RouterHelper) GetParam(ctx context.Context, param string) string {
	return way.Param(ctx, param)
}

type HelperInterface interface {
	GetParam(ctx context.Context, param string) string
}

type RouterInterface interface {
	Handle(method string, pattern string, handler http.Handler)
	HandleFunc(method string, pattern string, fn http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Router struct {
	Server *way.Router
}

type Adapter struct {
	Log    log.LoggerInterface
	Router *way.Router
}

func NewAdapter(a *Adapter) *Router {
	return &Router{
		Server: a.Router,
	}
}
