package logger

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	return logger
}
