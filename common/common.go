package common

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/log"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

type CommonResponseJSON struct {
	err     error
	data    interface{}
	message string
	errCode int
	status  int
}

var commonRender *render.Render = render.New()
var commonLogger *zap.Logger = log.BuildLogger()

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

func (res *CommonResponseJSON) SetWithStatus(httpStatus int, data interface{}) CommonResponseJSON {
	return CommonResponseJSON{
		err:     nil,
		data:    data,
		message: "",
		errCode: 0,
		status:  httpStatus,
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
