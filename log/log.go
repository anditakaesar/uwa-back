package log

import (
	"github.com/anditakaesar/uwa-back/env"
	"go.uber.org/zap"
)

func BuildLogger() *zap.Logger {
	zapOpt := zap.NewDevelopmentConfig()
	zapOpt.DisableStacktrace = true
	zapOpt.OutputPaths = []string{env.LogFilePath()}
	if env.AppEnv() == env.EnvProduction {
		zapProdOpt := zap.NewProductionConfig()
		zapProdOpt.OutputPaths = []string{env.LogFilePath()}
		zapLogger, _ := zapProdOpt.Build()
		defer zapLogger.Sync()
		return zapLogger
	}

	zapLogger, _ := zapOpt.Build()
	defer zapLogger.Sync()
	return zapLogger
}
