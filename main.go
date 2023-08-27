package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/anditakaesar/uwa-back/internal"
	"github.com/anditakaesar/uwa-back/internal/env"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting application %s", env.AppName())
	}
}

func run() error {
	addr := flag.String("addr", env.AppPort(), "http service address")
	flag.Parse()

	infra := internal.NewInfrastructure()
	err := infra.Init()
	if err != nil {
		infra.Logger.Error(err.Error(), err)
		return err
	}

	s := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      infra.InternalRouter,
		Addr:         *addr,
	}

	infra.Logger.Info("+++++++++++++++++++++++++++++++++++++++++++++++++++")
	infra.Logger.Info(fmt.Sprint(env.AppName(), "backend service started on port", env.AppPort()))
	infra.Logger.Info("+++++++++++++++++++++++++++++++++++++++++++++++++++")

	return s.ListenAndServe()
}
