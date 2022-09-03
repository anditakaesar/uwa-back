package main

import (
	"fmt"

	"github.com/anditakaesar/uwa-back/env"
	"github.com/anditakaesar/uwa-back/log"
)

func main() {
	logger := log.BuildLogger()

	logger.Info(fmt.Sprintf("Application %s running on port%s", env.AppName(), env.Port()))
}
