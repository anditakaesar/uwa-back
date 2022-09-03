package log

import "go.uber.org/zap"

func BuildLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	return logger
}
