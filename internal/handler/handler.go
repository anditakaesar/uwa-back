package handler

import (
	jsonEncoding "encoding/json"
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"go.uber.org/zap"
)

// DefaultDecoder ...
// var DefaultDecoder = schema.NewDecoder()

// EndpointHandler ...
type EndpointHandler func(http.ResponseWriter, *http.Request) ResponseInterface

func (fn EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := log.BuildNewLogger()
	requestIp := r.Header.Get(env.IPHeaderKey())
	logger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.String()), zap.Any("ip", requestIp))

	res := fn(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.GetStatus())

	if res.HasError() {
		handleErrorResponse(w, r, res)
	} else {
		handleOKResponse(w, res)
	}
}

func handleOKResponse(w http.ResponseWriter, res ResponseInterface) {
	if res.HasNoContent() {
		return
	}

	resp := make(map[string]interface{})

	resp["message"] = "success"
	data := res.GetData()

	if data == nil {
		emptyData := make(map[string]string)
		resp["data"] = emptyData
	} else {
		resp["data"] = data
	}

	encodeResponse(w, resp)
}

func handleErrorResponse(w http.ResponseWriter, r *http.Request, res ResponseInterface) {
	data := res.GetData()

	resp := map[string]interface{}{
		"data":    make(map[string]string),
		"code":    res.GetErrCode(),
		"message": res.GetErrorMessage(),
	}

	if data != nil {
		resp["data"] = data
	}

	encodeResponse(w, resp)
}

func encodeResponse(w http.ResponseWriter, data interface{}) {
	err := jsonEncoding.NewEncoder(w).Encode(data)
	if err != nil {
		// log := util.BuildLogger()
		// log.Warning(fmt.Sprintf("Error encode: %s", err.Error()))

		http.Error(w, "Error encode response", http.StatusInternalServerError)
	}
}
