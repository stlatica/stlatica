// Package logger abstracts the log output destinations.
// When used from other packages, it is possible to use it regardless of which log output destination is used.
package logger

import (
	"context"
)

// Logger is an interface to pkg logger.
// This interface can be commonly used by other packages regardless of the log output destination.
type Logger interface {
	// Info prints info log
	Info(ctx context.Context, msg string)
	// Error prints error log
	Error(ctx context.Context, msg string)
	// Warn prints warn log
	Warn(ctx context.Context, msg string)
}
