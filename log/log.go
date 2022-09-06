package log

import (
	"github.com/anditakaesar/uwa-back/env"
	"go.uber.org/zap"
)

func BuildLogger() *zap.Logger {
	zapOpt := zap.NewDevelopmentConfig()
	zapOpt.DisableStacktrace = true
	zapLogger, _ := zapOpt.Build()
	if env.AppEnv() == env.EnvProduction {
		zapLogger, _ = zap.NewProduction()
	}

	defer zapLogger.Sync()

	return zapLogger
}
