package handler

import (
	"net/http"

	json "github.com/anditakaesar/uwa-back/internal/json"

	"github.com/anditakaesar/uwa-back/internal/constants"
)

// EndpointHandler ...
type EndpointHandler func(http.ResponseWriter, *http.Request) ResponseInterface

func (fn EndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := fn(w, r)

	switch res.GetContentType() {
	case constants.AvailableMimeType.TextPlain:
		handleTextResponse(w, res)
	default:
		handleJsonResponse(w, r, res)
	}
}

func handleJsonResponse(w http.ResponseWriter, r *http.Request, res ResponseInterface) {
	w.Header().Set("Content-Type", constants.AvailableMimeType.ApplicationJson)
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
	err := json.Encode(data, w)
	if err != nil {
		// log := util.BuildLogger()
		// log.Warning(fmt.Sprintf("Error encode: %s", err.Error()))

		http.Error(w, "Error encode response", http.StatusInternalServerError)
	}
}

func handleTextResponse(w http.ResponseWriter, res ResponseInterface) {
	w.Header().Set("Content-Type", constants.AvailableMimeType.TextPlain)
	w.WriteHeader(res.GetStatus())

	w.Write([]byte(res.GetData().(string)))
}
