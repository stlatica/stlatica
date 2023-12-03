package plats

// Factory is a factory of plats package.
type Factory interface {
	// NewPlatGetter returns a new plat getter.
	NewPlatGetter() PlatGetter
}

// NewFactory returns a new factory of plats package.
func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

func (f *factory) NewPlatGetter() PlatGetter {
	return &platGetter{}
}