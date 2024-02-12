package logger

import (
	"context"

	"go.uber.org/zap"
)

type zapLogger struct {
	orgLogger *zap.Logger
}

// NewZapLogger creates new instance that implements Logger interface
func NewZapLogger() Logger {
	logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &zapLogger{
		orgLogger: logger,
	}
}

func (l *zapLogger) Info(_ context.Context, msg string) {
	l.orgLogger.Info(msg)
}

func (l *zapLogger) Error(_ context.Context, msg string) {
	l.orgLogger.Error(msg)
}

func (l *zapLogger) Warn(_ context.Context, msg string) {
	l.orgLogger.Warn(msg)
}
