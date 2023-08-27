package migration

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	json "github.com/anditakaesar/uwa-back/internal/json"

	"github.com/anditakaesar/uwa-back/application/context"
	"github.com/anditakaesar/uwa-back/internal/handler"
	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func (h Handler) DoMigration() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		type MigrationRequest struct {
			Number uint `json:"number"`
		}

		var request MigrationRequest

		err := json.Decode(&request, r.Body)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusBadRequest, err, 1, err.Error())
		}

		defer r.Body.Close()

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

		err = m.Migrate(request.Number)
		if err != nil {
			return h.Resp.SetErrorWithStatus(http.StatusInternalServerError, err, 1, err.Error())
		}
		return h.Resp.SetOk(request)
	}
}

func (h Handler) GetAvailableMigration() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		files, err := os.ReadDir("./migrations")
		if err != nil {
			h.AppContext.Logger.Error("[Handler][GetAvailableMigration] ioutil err", err)
			return h.Resp.SetError(err, 1, "io error")
		}

		results := []string{}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".up.sql") {
				results = append(results, file.Name())
			}
		}
		return h.Resp.SetOk(results)
	}
}
