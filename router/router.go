package router

import (
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/handler"
	"github.com/gorilla/mux"
)

const (
	indexPrefix = ""
	apiPrefix   = "/api"
)

type Route struct {
	PathPrefix  string
	Method      string
	Pattern     string
	Handler     common.EndpointHandlerJSON
	Middlewares []Middleware
}

type Middleware func(http.Handler, application.Context) http.Handler

func NewRouter(appContext application.Context) *mux.Router {
	router := mux.NewRouter()
	routes := registerRoutes(appContext)
	for _, route := range *routes {
		subRouter := router.PathPrefix(route.PathPrefix).Subrouter()

		subRouter.Methods(route.Method).
			Path(route.Pattern).
			Handler(chainMiddlewares(route.Handler, appContext, route.Middlewares...))
	}

	return router
}

func registerRoutes(appContext application.Context) *[]Route {
	r := []Route{}
	r = append(r, InitIndexRouter(appContext)...)
	r = append(r, InitApiAuthRouter(appContext)...)
	return &r
}

func InitIndexRouter(appContext application.Context) []Route {
	return []Route{
		{
			PathPrefix:  indexPrefix,
			Method:      http.MethodGet,
			Pattern:     "/",
			Handler:     handler.Index(appContext),
			Middlewares: []Middleware{},
		},
		{
			PathPrefix:  indexPrefix,
			Method:      http.MethodGet,
			Pattern:     "/greet/{name}",
			Handler:     handler.GetGreetName(appContext),
			Middlewares: []Middleware{},
		},
	}
}

func InitApiAuthRouter(appContext application.Context) []Route {
	return []Route{
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodGet,
			Pattern:    "/auth",
			Handler:    handler.GetAuth(appContext),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodPost,
			Pattern:    "/auth",
			Handler:    handler.PostAuth(appContext),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodGet,
			Pattern:    "/hash/{pass}",
			Handler:    handler.GetHashString(appContext),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
	}
}

func appTokenMiddleware(h http.Handler, appCtx application.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		reqToken = splitToken[1]
		if reqToken != "" && reqToken == env.AppToken() {
			h.ServeHTTP(w, r)
		} else {
			appCtx.Log.Warn("Login attempt!")
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func chainMiddlewares(h http.Handler, appContext application.Context, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h, appContext)
	}
	return h
}
