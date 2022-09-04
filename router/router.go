package router

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(appContext application.Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := registerRoutes(appContext)
	for _, route := range *routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func registerRoutes(appContext application.Context) *[]Route {
	r := []Route{}
	r = append(r, InitIndexRouter(appContext)...)
	return &r
}
