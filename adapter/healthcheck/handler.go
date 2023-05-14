package healthcheck

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/anditakaesar/uwa-back/internal/log"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
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
		err := h.AppContext.Mailer.SendMail("anditakaesar@gmail.com", "Some Test Subject", "This is a body, Thank You!")

		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, "error")
		}

		return h.Resp.SetOk(nil)
	}
}
