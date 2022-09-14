package handler

import (
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/domain"
	"github.com/anditakaesar/uwa-back/services"
)

func GetUsers(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		requestParam := services.GetAllUsersRequest{}
		err := requestParam.GetParamFromRequest(r)
		if err != nil {
			return res.SetBadRequest(err)
		}

		result := services.GetAllUsers(appCtx, requestParam)

		return res.SetOK(result)
	}
}

func GetUserTokenExpiry(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		userToken := common.GetBearerToken(r)
		var userCredential domain.UserCredential
		appCtx.DB.First(&userCredential, "user_token = ?", userToken)

		result := map[string]string{
			"expiredAt": userCredential.ExpiredAt.UTC().Format(time.RFC3339),
		}

		return res.SetOK(result)
	}
}
