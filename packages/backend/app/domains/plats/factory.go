package plats

import "github.com/stlatica/stlatica/packages/backend/app/logger"

// Factory is a factory of plats package.
type Factory interface {
	// NewPlatGetter returns a new plat getter.
	NewPlatGetter() PlatGetter
	// NewPlatCreator returns a new plat creator.
	NewPlatCreator() PlatCreator
}

// NewFactory returns a new factory of plats package.
func NewFactory(appLogger *logger.AppLogger) Factory {
	return &factory{
		appLogger: appLogger,
	}
}

type factory struct {
	appLogger *logger.AppLogger
}

func (f *factory) NewPlatGetter() PlatGetter {
	return &platGetter{
		appLogger: f.appLogger,
	}
}

func (f *factory) NewPlatCreator() PlatCreator {
	return &platCreator{
		appLogger: f.appLogger,
	}
}
