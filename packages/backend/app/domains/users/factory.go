package users

import "github.com/stlatica/stlatica/packages/backend/app/logger"

// Factory is a factory of users package.
type Factory interface {
	// NewUserGetter returns a new user getter.
	NewUserGetter() UserGetter
}

// NewFactory returns a new factory of users package.
func NewFactory(appLogger *logger.AppLogger) Factory {
	return &factory{
		appLogger: appLogger,
	}
}

type factory struct {
	appLogger *logger.AppLogger
}

func (f *factory) NewUserGetter() UserGetter {
	return &userGetter{
		appLogger: f.appLogger,
	}
}
