package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/anditakaesar/uwa-back/common"
	"github.com/anditakaesar/uwa-back/services"
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

		err = services.AuthUser(appContext, request)
		if err != nil {
			return res.SetWithStatus(http.StatusUnauthorized, map[string]string{"message": err.Error()})
		}

		return res.SetOK(map[string]string{"message": "Post Auth Api"})
	}
}

func GetGreetName(appContext application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		name := mux.Vars(r)["name"]

		return res.SetOK(map[string]string{"name": name})
	}
}

func GetHashString(appContext application.Context) common.EndpointHandlerJSON {
	return func(w http.ResponseWriter, r *http.Request) (res common.CommonResponseJSON) {
		pass := mux.Vars(r)["pass"]

		return res.SetOK(map[string]string{"pass": pass, "hash": appContext.Crypter.GenerateHash(pass)})
	}
}
