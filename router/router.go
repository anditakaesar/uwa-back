package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/handler"
	"github.com/anditakaesar/uwa-back/log"
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
	HandlerFunc http.HandlerFunc
	Middlewares []Middleware
}

type Middleware func(http.Handler) http.Handler

func NewRouter(appContext application.Context) *mux.Router {
	router := mux.NewRouter()
	routes := registerRoutes(appContext)
	for _, route := range *routes {
		subRouter := router.PathPrefix(route.PathPrefix).Subrouter()

		subRouter.Methods(route.Method).
			Path(route.Pattern).
			Handler(chainMiddlewares(route.HandlerFunc, route.Middlewares...))
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
			HandlerFunc: handler.Index(appContext),
			Middlewares: []Middleware{
				loggingMiddleware,
			},
		},
		{
			PathPrefix:  indexPrefix,
			Method:      http.MethodGet,
			Pattern:     "/greet/{name}",
			HandlerFunc: handler.GetGreetName(appContext),
			Middlewares: []Middleware{
				loggingMiddleware,
			},
		},
	}
}

func InitApiAuthRouter(appContext application.Context) []Route {
	return []Route{
		{
			PathPrefix:  apiPrefix,
			Method:      http.MethodGet,
			Pattern:     "/auth",
			HandlerFunc: handler.GetAuth(appContext),
			Middlewares: []Middleware{
				loggingMiddleware,
				appTokenMiddleware,
			},
		},
		{
			PathPrefix:  apiPrefix,
			Method:      http.MethodPost,
			Pattern:     "/auth",
			HandlerFunc: handler.PostAuth(appContext),
			Middlewares: []Middleware{
				loggingMiddleware,
				appTokenMiddleware,
			},
		},
	}
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New()
		logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL))

		h.ServeHTTP(w, r)
	})
}

func appTokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		reqToken = splitToken[1]
		if reqToken != "" && reqToken == env.AppToken() {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func chainMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
