package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/services"
	"github.com/anditakaesar/uwa-back/utils"
)

func GetUsers(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		requestParam := services.GetAllUsersRequest{}
		err := requestParam.GetParamFromRequest(r)
		if err != nil {
			return res.SetBadRequest(err)
		}

		result := appCtx.Services.UserService.GetAllUsers(requestParam)

		return res.SetOK(result)
	}
}

func GetUserTokenExpiry(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		userCredential := appCtx.DBI.GetUserCredentialByToken(utils.GetBearerToken(r))

		isExpired := !userCredential.ExpiredAt.After(time.Now())

		result := map[string]string{
			"isExpired": fmt.Sprint(isExpired),
			"expiredAt": userCredential.ExpiredAt.UTC().Format(time.RFC3339),
		}

		return res.SetOK(result)
	}
}
