package common

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/log"
	"github.com/unrolled/render"
)

type CommonResponseJSON struct {
	err     error
	data    interface{}
	message string
	errCode int
	status  int
}

var commonRender *render.Render = render.New()
var commonLogger log.LogInterface = log.BuildLogger()

func (res *CommonResponseJSON) GetStatus() int {
	return res.status
}

func (res *CommonResponseJSON) HasError() bool {
	return res.err != nil
}

func (res *CommonResponseJSON) GetErrCode() int {
	return res.errCode
}

func (res *CommonResponseJSON) GetErrMessage() string {
	return res.err.Error()
}

func (res *CommonResponseJSON) GetCompleteErrMessage() string {
	return fmt.Sprintf("Code: %d, Message: %s, Detail: %s", res.errCode, res.message, res.err.Error())
}

func (res *CommonResponseJSON) SetOK(data interface{}) CommonResponseJSON {
	return CommonResponseJSON{
		err:     nil,
		data:    data,
		message: "",
		errCode: 0,
		status:  http.StatusOK,
	}
}

func (res *CommonResponseJSON) SetOKWithSuccessMessage() CommonResponseJSON {
	return CommonResponseJSON{
		err:     nil,
		data:    map[string]string{"message": "success"},
		message: "",
		errCode: 0,
		status:  http.StatusOK,
	}
}

func (res *CommonResponseJSON) SetWithStatus(httpStatus int, data interface{}) CommonResponseJSON {
	return CommonResponseJSON{
		err:     nil,
		data:    data,
		message: "",
		errCode: 0,
		status:  httpStatus,
	}
}

func (res *CommonResponseJSON) SetInternalError(err error) CommonResponseJSON {
	return CommonResponseJSON{
		err:     err,
		data:    nil,
		message: err.Error(),
		errCode: 0,
		status:  http.StatusInternalServerError,
	}
}

func (res *CommonResponseJSON) SetBadRequest(err error) CommonResponseJSON {
	return CommonResponseJSON{
		err:     err,
		data:    nil,
		message: err.Error(),
		errCode: 0,
		status:  http.StatusBadRequest,
	}
}

type EndpointHandlerJSON func(http.ResponseWriter, *http.Request) CommonResponseJSON

func (fn EndpointHandlerJSON) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	commonLogger.Info(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
	res := fn(w, r)
	if res.HasError() {
		commonLogger.Error(res.GetCompleteErrMessage())
		data := struct {
			Code    int
			Message string
		}{
			Code:    res.GetErrCode(),
			Message: res.GetErrMessage(),
		}

		commonRender.JSON(w, res.status, data)
	} else {
		commonRender.JSON(w, res.status, res.data)
	}
}

func GetBearerToken(r *http.Request) string {
	userToken := r.Header.Get("Authorization")
	splitToken := strings.Split(userToken, "Bearer ")

	if len(splitToken) > 1 {
		return splitToken[1]
	}

	return ""
}
