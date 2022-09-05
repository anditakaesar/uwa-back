package handler

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
	"github.com/gorilla/mux"
)

func Index(appContext application.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appContext.Log.Info("Index Hit")

		type XData struct {
			Data string
		}

		response := struct {
			Data string
			Arr  []XData
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

		appContext.Render.JSON(w, http.StatusOK, response)
	}
}

func GetAuth(appContext application.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appContext.Log.Info("Get Auth")
		fmt.Fprintf(w, "Get Auth api")
	}
}

func PostAuth(appContext application.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appContext.Log.Info("Post Auth")
		fmt.Fprintf(w, "Post Auth api")
	}
}

func GetGreetName(appContext application.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]

		fmt.Fprintf(w, "Hello %s", name)
	}
}
