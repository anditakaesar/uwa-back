package router

import (
	"net/http"
	"path"

	"github.com/anditakaesar/uwa-back/adapter/httpserver"
	ma "github.com/anditakaesar/uwa-back/adapter/middleware"
	"github.com/anditakaesar/uwa-back/internal/constants"
	"github.com/anditakaesar/uwa-back/internal/handler"

	"github.com/thoas/go-funk"
)

// Context ...
type Context struct {
	router     httpserver.RouterInterface
	middleware *ma.Middleware
	prefix     string
	Helper     httpserver.HelperInterface
}

type EndpointInfo struct {
	HTTPMethod    string
	URLPattern    string
	Handler       handler.EndpointHandler
	Verifications []constants.VerificationType
}

func NewService(router httpserver.RouterInterface, helper httpserver.HelperInterface, middleware ma.Middleware, prefix string) Context {
	return Context{
		router:     router,
		middleware: &middleware,
		prefix:     prefix,
		Helper:     helper,
	}
}

func (r *Context) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

// RegisterEndpoint ...
func (r *Context) RegisterEndpoint(info EndpointInfo) {
	r.RegisterEndpointWithPrefix(info, r.prefix)
}

// RegisterEndpointWithPrefix ...
func (r *Context) RegisterEndpointWithPrefix(info EndpointInfo, prefix string) {
	m := r.middleware
	urlPattern := getFullURLPattern(info, prefix)

	verificationFns := getVerificationMethod(m, info.Verifications)

	r.router.Handle(info.HTTPMethod, urlPattern,
		m.Cors(
			m.Verify(info.Handler, verificationFns...)))
}

func getVerificationMethod(m *ma.Middleware, verifications []constants.VerificationType) []ma.MiddlewareFunc {
	return funk.Map(verifications, func(t constants.VerificationType) ma.MiddlewareFunc {
		switch t {
		case constants.AccessTokenValue:
			return m.AccessToken
		default:
			return m.ApiToken
		}

	}).([]ma.MiddlewareFunc)
}

func getFullURLPattern(info EndpointInfo, prefix string) string {
	if prefix == "" {
		return info.URLPattern
	}
	return path.Join(prefix, info.URLPattern)
}

type webHandler func(http.ResponseWriter, *http.Request) error

func (fn webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (r *Context) InitOptionsRoute() {
	r.router.Handle(
		http.MethodOptions,
		"/api/",
		r.middleware.Cors(handleCors()),
	)
}

func handleCors() webHandler {
	return func(w http.ResponseWriter, req *http.Request) error {
		return nil
	}
}
