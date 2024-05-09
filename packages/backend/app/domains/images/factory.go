package images

import "github.com/stlatica/stlatica/packages/backend/app/pkg/logger"

// Factory is a factory of images package.
type Factory interface {
	// NewImageUploader returns a new image uploader.
	NewImageUploader() ImageUploader
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

func (f *factory) NewImageUploader() ImageUploader {
	return &imageUploader{
		appLogger: f.appLogger,
	}
}
