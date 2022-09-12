package router

import (
	"fmt"
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

func NewRouter(appCtx application.Context) *mux.Router {
	router := mux.NewRouter()
	routes := registerRoutes(appCtx)
	for _, route := range *routes {
		subRouter := router.PathPrefix(route.PathPrefix).Subrouter()

		subRouter.Methods(route.Method).
			Path(route.Pattern).
			Handler(chainMiddlewares(route.Handler, appCtx, route.Middlewares...))
	}

	return router
}

func registerRoutes(appCtx application.Context) *[]Route {
	r := []Route{}
	r = append(r, InitIndexRouter(appCtx)...)
	r = append(r, InitApiAuthRouter(appCtx)...)
	r = append(r, InitUserRouter(appCtx)...)
	return &r
}

func InitIndexRouter(appCtx application.Context) []Route {
	return []Route{
		{
			PathPrefix:  indexPrefix,
			Method:      http.MethodGet,
			Pattern:     "/",
			Handler:     handler.Index(appCtx),
			Middlewares: []Middleware{},
		},
		{
			PathPrefix:  indexPrefix,
			Method:      http.MethodGet,
			Pattern:     "/greet/{name}",
			Handler:     handler.GetGreetName(appCtx),
			Middlewares: []Middleware{},
		},
	}
}

func InitApiAuthRouter(appCtx application.Context) []Route {
	return []Route{
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodGet,
			Pattern:    "/auth",
			Handler:    handler.GetAuth(appCtx),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodPost,
			Pattern:    "/auth",
			Handler:    handler.PostAuth(appCtx),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodGet,
			Pattern:    "/hash/{pass}",
			Handler:    handler.GetHashString(appCtx),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodPost,
			Pattern:    "/tools/migrate/all",
			Handler:    handler.MigrateAll(appCtx),
			Middlewares: []Middleware{
				appTokenMiddleware,
			},
		},
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodPost,
			Pattern:    "/tools/seed/{table}",
			Handler:    handler.SeedOne(appCtx),
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
			appCtx.Log.Warn(fmt.Sprintf("Login attempt! with token:%v", reqToken))
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func chainMiddlewares(h http.Handler, appCtx application.Context, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h, appCtx)
	}
	return h
}
