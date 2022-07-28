package log

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

type Logger interface {
	Info(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Debug(string, ...zap.Field)
	Error(string, ...zap.Field)
}

func init() {
	logger, _ = zap.NewProduction()
}

func SyncLogger() {
	logger.Sync()
}

// info level log
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// fatal level log
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// debug level log
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// error level log
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
