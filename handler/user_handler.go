package handler

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
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
