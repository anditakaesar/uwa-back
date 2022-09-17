package log

import (
	"github.com/anditakaesar/uwa-back/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogInterface interface {
	Info(message string, fields ...zapcore.Field)
	Error(message string, fields ...zapcore.Field)
	Warn(message string, fields ...zapcore.Field)
	Fatal(message string, fields ...zapcore.Field)
}

func BuildLogger() LogInterface {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableStacktrace = true

	if env.AppEnv() == env.EnvProduction {
		zapProdConfig := zap.NewProductionConfig()
		zapProdConfig.OutputPaths = []string{env.LogFilePath()}
		zapLogger, _ := zapProdConfig.Build()
		defer zapLogger.Sync()
		return zapLogger
	}

	zapLogger, _ := zapConfig.Build()
	defer zapLogger.Sync()
	return zapLogger
}
