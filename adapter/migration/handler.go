package migration

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anditakaesar/uwa-back/adapter/database"
	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/anditakaesar/uwa-back/internal/log"
	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type HandlerDependency struct {
	DB         database.DatabaseInterface
	Logger     log.LoggerInterface
	AppContext context.AppContext
}

type Handler struct {
	DB         database.DatabaseInterface
	Resp       handler.ResponseInterface
	Log        log.LoggerInterface
	AppContext context.AppContext
}

func NewHandler(d HandlerDependency) Handler {
	return Handler{
		DB:         d.DB,
		Resp:       handler.NewResponse(handler.Dep{Log: d.Logger}),
		Log:        d.Logger,
		AppContext: d.AppContext,
	}
}

func (h Handler) UpMigration() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		sqlDB, _ := h.AppContext.DB.Get().DB()
		mydir, _ := os.Getwd()
		driver, err := migratePg.WithInstance(sqlDB, &migratePg.Config{})
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, err.Error())
		}

		m, err := migrate.NewWithDatabaseInstance(fmt.Sprint("file://", mydir, "/migrations"), "uwa_back", driver)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, err.Error())
		}

		err = m.Up()
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, err.Error())
		}
		return h.Resp.SetOk(nil)
	}
}
