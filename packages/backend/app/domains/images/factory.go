package images

import "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"

// Factory is a factory of images package.
type Factory interface {
	// NewImageGetter returns a new image getter.
	NewImageGetter() ImageGetter
	// NewImageUploader returns a new image uploader.
	NewImageUploader() ImageUploader
	// NewImageDeleter returns a new image deleter.
	NewImageDeleter() ImageDeleter
}

// NewFactory returns a new factory of images package.
func NewFactory(appLogger *logger.AppLogger) Factory {
	return &factory{
		appLogger: appLogger,
	}
}

type factory struct {
	appLogger *logger.AppLogger
}

func (f *factory) NewImageGetter() ImageGetter {
	return &imageGetter{
		appLogger: f.appLogger,
	}
}

func (f *factory) NewImageUploader() ImageUploader {
	return &imageUploader{
		appLogger: f.appLogger,
	}
}

func (f *factory) NewImageDeleter() ImageDeleter {
	return &imageDeleter{
		appLogger: f.appLogger,
	}
}
