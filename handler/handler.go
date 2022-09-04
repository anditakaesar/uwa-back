package handler

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/application"
)

func Index(appContext application.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appContext.Log.Info("Index Hit")
		fmt.Fprintf(w, "Index with logging")
	}
}
