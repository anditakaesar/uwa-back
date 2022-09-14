package router

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/anditakaesar/uwa-back/handler"
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
		userToken := common.GetBearerToken(r)

		var userCredential domain.UserCredential
		appCtx.DB.First(&userCredential, "user_token = ?", userToken)

		if funk.NotEmpty(userCredential) {
			if userCredential.ExpiredAt.After(*appCtx.TimeNow) {
				userCredential.LastAccessAt = appCtx.TimeNow
				appCtx.DB.Save(&userCredential)
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
