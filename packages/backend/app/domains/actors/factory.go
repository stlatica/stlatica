package actors

// Factory is a factory of actors package.
type Factory interface {
	// NewActorGetter returns a new actor getter.
	NewActorGetter() ActorGetter
}

// NewFactory returns a new factory of actors package.
func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

func (f *factory) NewActorGetter() ActorGetter {
	return &actorGetter{}
}
