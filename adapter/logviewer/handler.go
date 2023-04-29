package logviewer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/constants"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/anditakaesar/uwa-back/internal/way"
	"go.uber.org/zap"
)

type HandlerDependency struct {
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

type Handler struct {
	Resp       handler.ResponseInterface
	Log        log.LoggerInterface
	AppContext context.AppContext
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		Resp:       handler.NewResponse(handler.Dep{Log: d.Logger}),
		Log:        d.Logger,
		AppContext: d.AppContext,
	}
}

func (h Handler) GetAvailableLogs() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			h.Log.Error("[Handler][GetAvailableLogs] ioutil err", err)
			return h.Resp.SetError(err, 1, "io error")
		}

		results := []string{}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".log") {
				newStr := fmt.Sprintf("%s/api/logviewer/%s", env.HostURL(), file.Name())
				results = append(results, newStr)
			}
		}
		return h.Resp.SetOk(results)
	}
}

func (h Handler) GetLog() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		fileName := way.Param(r.Context(), "logfile")

		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			h.Log.Warning("[Handler][GetLog] file not exist", zap.String("fileName", fileName))
			return h.Resp.SetErrorWithStatus(http.StatusNotFound, err, 1, "404 not found")
		}

		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			h.Log.Error("[Handler][GetLog] ioutil err", err, zap.String("fileName", fileName))
			return h.Resp.SetErrorWithStatus(http.StatusNotFound, err, 1, "404 not found")
		}

		return h.Resp.SetOkWithText(constants.AvailableMimeType.TextPlain, string(data))
	}
}
