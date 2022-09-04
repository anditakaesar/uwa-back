package router

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/handler"
)

func InitIndexRouter(appContext application.Context) []Route {
	return []Route{
		{
			Name:        "IndexRoute",
			Method:      http.MethodGet,
			Pattern:     "/",
			HandlerFunc: handler.Index(appContext),
		},
	}
}
