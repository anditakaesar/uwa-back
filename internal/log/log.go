package log

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const NoticeLevel zapcore.Level = zapcore.DebugLevel - 1

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	ResetColor   = "\033[0m"
)

type LoggerInterface interface {
	Info(string, ...zapcore.Field)
	Error(string, error, ...zapcore.Field)
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
		switch l {
		case NoticeLevel:
			col += NoticeColor
			txt = "NOTICE"
		case zapcore.DebugLevel:
			col += DebugColor
		case zapcore.InfoLevel:
			col += InfoColor
		case zapcore.WarnLevel:
			col += WarningColor
		case zapcore.ErrorLevel:
			col += ErrorColor
		case zapcore.DPanicLevel:
			col += ErrorColor
		case zapcore.PanicLevel:
			col += ErrorColor
		case zapcore.FatalLevel:
			col += ErrorColor
		}

		pae.AppendString(col + txt + ResetColor)
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
