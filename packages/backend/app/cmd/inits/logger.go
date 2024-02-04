package inits

import (
	"github.com/stlatica/stlatica/packages/backend/app/logger"
	pkglogger "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// LoggerFactory is a factory of logger
type LoggerFactory interface {
	// AppLogger returns a new app logger
	AppLogger() *logger.AppLogger
}

// NewLoggerFactory creates LoggerFactory
func NewLoggerFactory() LoggerFactory {
	zapLogger := pkglogger.NewZapLogger()
	return &loggerFactory{
		logger: zapLogger,
	}
}

type loggerFactory struct {
	logger pkglogger.Logger
}

func (f *loggerFactory) AppLogger() *logger.AppLogger {
	return logger.NewAppLogger(f.logger)
}
