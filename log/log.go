package log

import (
	"github.com/anditakaesar/uwa-back/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Interface interface {
	Info(string, ...zapcore.Field)
	Error(str string, f ...zapcore.Field)
	Warning(str string, f ...zapcore.Field)
	Flush()
}

type Logger struct {
	zap *zap.Logger
}

func New() Logger {
	logger := Logger{
		zap: buildZapLogger(),
	}

	defer logger.zap.Sync()

	return logger
}

func buildZapLogger() *zap.Logger {
	zapLogger, _ := zap.NewDevelopment()
	if env.AppEnv() == env.EnvProduction {
		zapLogger, _ = zap.NewProduction()
	}
	return zapLogger
}

func (l *Logger) Info(str string, f ...zapcore.Field) {
	l.zap.Info(str, f...)
}

func (l *Logger) Error(str string, f ...zapcore.Field) {
	l.zap.Error(str, f...)
}

func (l *Logger) Warning(str string, f ...zapcore.Field) {
	l.zap.Warn(str, f...)
}

func (l *Logger) Flush() {
	l.zap.Sync()
}
