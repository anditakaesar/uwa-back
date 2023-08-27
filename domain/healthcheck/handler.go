package healthcheck

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/handler"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
)

type HandlerDependency struct {
	AppContext context.AppContext
}

type Handler struct {
	Resp       handler.ResponseInterface
	AppContext context.AppContext
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		Resp:       handler.NewResponse(handler.Dep{Log: d.AppContext.Logger}),
		AppContext: d.AppContext,
	}
}

func (h Handler) CheckHealth() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		newKey := fmt.Sprintf("test:%s", uuid.NewString())

		redis := h.AppContext.Redis
		redis.SetRedisValue(newKey, time.Now().Format(time.RFC3339), 2*time.Minute)
		return h.Resp.SetOk(newKey)
	}
}

func (h Handler) SendTestMail() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		err := h.AppContext.Mailer.SendMail(env.EmailTestTo(), "Some Test Subject", "This is a body, Thank You!")

		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, "error")
		}

		return h.Resp.SetOk(nil)
	}
}

func (h Handler) GetIpLog() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		ipLog, _ := h.AppContext.IplogModel.GetIplogByAddress("127.0.0.1")
		return h.Resp.SetOk(ipLog)
	}
}
