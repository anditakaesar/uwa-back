package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/handler"
	"github.com/anditakaesar/uwa-back/utils"
	"github.com/thoas/go-funk"
)

func InitUserRouter(appCtx application.Context) []Route {
	return []Route{
		{
			PathPrefix: apiPrefix,
			Method:     http.MethodGet,
			Pattern:    "/user",
			Handler:    handler.GetUsers(appCtx),
			Middlewares: []Middleware{
				userTokenMiddleware,
			},
		},
	}
}

func userTokenMiddleware(h http.Handler, appCtx application.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userToken := utils.GetBearerToken(r)
		userCredential := appCtx.DBI.GetUserCredentialByToken(userToken)
		timeNow := time.Now()

		if funk.NotEmpty(userCredential) {
			if userCredential.ExpiredAt.After(timeNow) {
				userCredential.LastAccessAt = &timeNow
				appCtx.DBI.UpdateUserCredential(userCredential)
				h.ServeHTTP(w, r)
			} else {
				appCtx.Log.Warn(fmt.Sprintf("Login attempt! with expired token:%v", userToken))
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		} else {
			appCtx.Log.Warn(fmt.Sprintf("Login attempt! non exist token:%v", userToken))
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
