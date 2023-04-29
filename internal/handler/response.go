package handler

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/internal/log"
)

type HTTPResponse struct {
	status      int
	data        interface{}
	err         error
	errCode     int
	message     string
	noContent   bool
	contentType string
	Logger      log.LoggerInterface
}

type Dep struct {
	Log log.LoggerInterface
}

func NewResponse(dep Dep) *HTTPResponse {
	return &HTTPResponse{
		Logger: dep.Log,
	}
}

// SetOk ...
func (res HTTPResponse) SetOk(data interface{}) HTTPResponse {
	res.status = http.StatusOK
	res.data = data
	res.message = "success"

	return res
}

// SetOkWithStatus ...
func (res HTTPResponse) SetOkWithStatus(status int, data interface{}) HTTPResponse {
	res.status = status
	res.data = data
	res.message = "success"
	res.noContent = status == http.StatusNoContent

	return res
}

// SetError ...
func (res HTTPResponse) SetError(err error, errCode int, message string) HTTPResponse {
	res.status = http.StatusInternalServerError
	res.err = err
	res.errCode = errCode
	res.message = message
	res.data = nil

	return res
}

// SetErrorWithStatus ...
func (res HTTPResponse) SetErrorWithStatus(status int, err error, errCode int, message string) HTTPResponse {
	res.status = status
	res.err = err
	res.errCode = errCode
	res.message = message

	return res
}

// HasError ...
func (res HTTPResponse) HasError() bool {
	return res.err != nil
}

// GetData ...
func (res HTTPResponse) GetData() interface{} {
	return res.data
}

// GetError ...
func (res HTTPResponse) GetError() error {
	return res.err
}

// GetStatus ...
func (res HTTPResponse) GetStatus() int {
	if res.status != 0 {
		return res.status
	}
	return http.StatusInternalServerError
}

// GetErrCode ...
func (res HTTPResponse) GetErrCode() int {
	if res.errCode != 0 {
		return res.errCode
	}

	return 500
}

// GetErrorMessage get error message from message or error object
func (res HTTPResponse) GetErrorMessage() string {
	if res.message != "" {
		return res.message
	}

	return res.err.Error()
}

// GetErrorMessageVerbose get full string with error code, message and error object
func (res HTTPResponse) GetErrorMessageVerbose() string {
	return fmt.Sprintf("Error Code: %d, Message: %s. Detail: %s", res.errCode, res.message, res.err.Error())
}

// HasNoContent ...
func (res HTTPResponse) HasNoContent() bool {
	return res.noContent
}

func (res HTTPResponse) SetErrorWithData(errParam ErrorResponseParam, data interface{}) HTTPResponse {
	res.status = errParam.Status
	res.err = errParam.Err
	res.errCode = errParam.ErrCode
	res.message = errParam.Message
	res.data = data

	return res
}

func (res HTTPResponse) GetContentType() string {
	return res.contentType
}

func (res HTTPResponse) SetOkWithText(contentType string, data string) HTTPResponse {
	res.contentType = contentType
	res.data = data
	res.status = http.StatusOK

	return res
}

type ErrorResponseParam struct {
	Status  int
	Err     error
	ErrCode int
	Message string
	Data    map[string]interface{}
}

type ResponseInterface interface {
	SetOk(data interface{}) HTTPResponse
	SetOkWithStatus(status int, data interface{}) HTTPResponse
	SetOkWithText(contentType string, data string) HTTPResponse
	SetError(err error, errCode int, message string) HTTPResponse
	SetErrorWithStatus(status int, err error, errCode int, message string) HTTPResponse
	SetErrorWithData(errParam ErrorResponseParam, data interface{}) HTTPResponse
	HasError() bool
	GetData() interface{}
	GetError() error
	GetStatus() int
	GetErrCode() int
	GetErrorMessage() string
	GetErrorMessageVerbose() string
	HasNoContent() bool
	GetContentType() string
}
