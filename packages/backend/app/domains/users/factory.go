package users

import "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"

// Factory is a factory of users package.
type Factory interface {
	// NewUserGetter returns a new user getter.
	NewUserGetter() UserGetter
	// NewUserCreator returns a new user creator.
	NewUserCreator() UserCreator
	// NewUserUpdater returns a new user updater.
	NewUserUpdater() UserUpdater
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

func (f *factory) NewUserCreator() UserCreator {
	return &userCreator{
		appLogger: f.appLogger,
	}
}

func (f *factory) NewUserUpdater() UserUpdater {
	return &userUpdater{
		appLogger: f.appLogger,
	}
}
