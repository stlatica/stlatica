//nolint:revive,lll // TODO: msg以外のparameterについても出力できるように修正する https://github.com/stlatica/stlatica/issues/295, https://github.com/stlatica/stlatica/issues/296
package logger

import (
	"context"
	"time"

	pkglogger "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// AppLogger is a logger for application.
// It is used to log application specific logs.
// ref: https://github.com/stlatica/stlatica/blob/main/docs/designs/log_design.md
type AppLogger struct {
	logger pkglogger.Logger
}

// NewAppLogger creates AppLogger
func NewAppLogger(logger pkglogger.Logger) *AppLogger {
	return &AppLogger{
		logger: logger,
	}
}

// Info prints info log
func (l *AppLogger) Info(ctx context.Context, msg string, traceID string, executedTime time.Time, userID string) {
	l.logger.Info(ctx, msg)
}

// Error prints error log
func (l *AppLogger) Error(ctx context.Context, msg string, traceID string, executedTime time.Time, userID string) {
	l.logger.Error(ctx, msg)
}

// Warn prints warn log
func (l *AppLogger) Warn(ctx context.Context, msg string, traceID string, executedTime time.Time, userID string) {
	l.logger.Warn(ctx, msg)
}
