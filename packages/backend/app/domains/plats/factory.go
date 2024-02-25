package plats

// Factory is a factory of plats package.
type Factory interface {
	// NewPlatGetter returns a new plat getter.
	NewPlatGetter() PlatGetter
	// NewPlatCreator returns a new plat creator.
	NewPlatCreator() PlatCreator
}

// NewFactory returns a new factory of plats package.
func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

func (f *factory) NewPlatGetter() PlatGetter {
	return &platGetter{}
}

func (f *factory) NewPlatCreator() PlatCreator {
	return &platCreator{}
}
