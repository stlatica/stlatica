package actors

import "github.com/stlatica/stlatica/packages/backend/app/logger"

// Factory is a factory of actors package.
type Factory interface {
	// NewActorGetter returns a new actor getter.
	NewActorGetter() ActorGetter
}

// NewFactory returns a new factory of actors package.
func NewFactory(appLogger *logger.AppLogger) Factory {
	return &factory{
		appLogger: appLogger,
	}
}

type factory struct {
	appLogger *logger.AppLogger
}

func (f *factory) NewActorGetter() ActorGetter {
	return &actorGetter{
		appLogger: f.appLogger,
	}
}
