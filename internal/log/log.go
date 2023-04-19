package log

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const NoticeLevel zapcore.Level = zapcore.DebugLevel - 1

type LoggerInterface interface {
	Info(string, ...zapcore.Field)
	Error(string, error, ...zapcore.Field)
	Warning(string, ...zapcore.Field)
	Flush()
}

type LoggerDependency struct {
	Zap *zap.Logger
}

type Logger struct {
	zap *zap.Logger
}

func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Error(msg string, err error, fields ...zapcore.Field) {
	fields = append(fields, zap.String("internalError", err.Error()))
	l.zap.Error(msg, fields...)
}

func (l *Logger) Warning(msg string, fields ...zapcore.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Logger) Flush() {
	l.zap.Sync()
}

func NewLogger(ld *LoggerDependency) LoggerInterface {
	defer ld.Zap.Sync()

	return &Logger{
		zap: ld.Zap,
	}
}

func BuildNewLogger() LoggerInterface {
	executionTime := time.Now()
	todayOutputPath := fmt.Sprintf("%04d-%02d-%02d.log", executionTime.Year(), executionTime.Month(), executionTime.Day())
	var err error
	devConf := zap.NewDevelopmentConfig()
	devConf.OutputPaths = []string{todayOutputPath, "stderr"}
	devConf.Level = zap.NewAtomicLevelAt(NoticeLevel)
	devConf.DisableStacktrace = true
	devConf.EncoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		var col string
		txt := l.CapitalString()
		pae.AppendString(col + txt)
	}

	newZap, err := devConf.Build(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		log.Fatal("cannot init logger in local", err)
	}

	defer newZap.Sync()
	return NewLogger(&LoggerDependency{
		Zap: newZap,
	})
}
