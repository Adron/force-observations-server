package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// GetLogger returns the configured logger instance
func GetLogger() *zap.Logger {
	return logger
}
