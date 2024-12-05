package auth

import "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"

// Factory is a factory of auth package.
type Factory interface {
	// NewAuthenticator returns a new authenticator.
	NewAuthenticator() Authenticator
}

// NewFactory returns a new factory of auth package.
func NewFactory(appLogger *logger.AppLogger) Factory {
	return &factory{
		appLogger: appLogger,
	}
}

type factory struct {
	appLogger *logger.AppLogger
}

func (f *factory) NewAuthenticator() Authenticator {
	return &authenticator{
		appLogger: f.appLogger,
	}
}
