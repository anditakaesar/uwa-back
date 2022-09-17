package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/services"
	"github.com/anditakaesar/uwa-back/utils"
	"github.com/gorilla/mux"
)

func Index(appContext application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		appContext.Log.Info("Index Hit")

		type XData struct {
			Data string
		}

		response := struct {
			Data string  `json:"parentData"`
			Arr  []XData `json:"childs"`
		}{
			Data: "some data",
			Arr: []XData{
				{
					Data: "one",
				},
				{
					Data: "two",
				},
			},
		}

		return res.SetOK(response)
	}
}

func GetAuth(appContext application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		return res.SetOK(map[string]string{"message": "Get Auth Api"})
	}
}

func PostAuth(appContext application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		request := services.AuthParam{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			return res.SetWithStatus(http.StatusInternalServerError, map[string]string{"message": "error decoder"})
		}

		authService := appContext.Services.AuthService

		userToken, err := authService.AuthUser(request)
		if err != nil {
			return res.SetWithStatus(http.StatusUnauthorized, map[string]string{"message": err.Error()})
		}

		return res.SetOK(map[string]string{"message": "Post Auth Api", "token": userToken})
	}
}

func GetGreetName(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		name := mux.Vars(r)["name"]

		return res.SetOK(map[string]string{"name": name})
	}
}

func GetHashString(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		pass := mux.Vars(r)["pass"]
		crypter := utils.GetDefaultCrypter()

		return res.SetOK(map[string]string{"pass": pass, "hash": crypter.GenerateHash(pass)})
	}
}

func MigrateAll(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		err := appCtx.Services.DBToolsService.AutoMigrate()
		if err != nil {
			return res.SetInternalError(err)
		}

		return res.SetOKWithSuccessMessage()
	}
}

func SeedOne(appCtx application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		table := mux.Vars(r)["table"]
		if table == "" {
			return res.SetWithStatus(http.StatusBadRequest, map[string]string{"message": "table name required"})
		}
		err := appCtx.Services.DBToolsService.Seed(table)
		if err != nil {
			return res.SetInternalError(err)
		}

		return res.SetOKWithSuccessMessage()
	}
}
