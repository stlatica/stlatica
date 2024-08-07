package images

import (
	"context"
	"io"

	"github.com/stlatica/stlatica/packages/backend/app/domains/images"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images/ports"
)

// ImageUseCase is the interface for handling image.
type ImageUseCase interface {
	// GetImage gets image.
	GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error)
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageBinary []byte) (types.ImageID, error)
	// DeleteImage deletes image.
	DeleteImage(ctx context.Context, imageID types.ImageID) error
}

// NewImageUseCase returns ImageUseCase.
func NewImageUseCase(appLogger *logger.AppLogger,
	imageAdapter ports.ImageAdapter, domainFactory images.Factory) ImageUseCase {
	return &imageUseCase{
		appLogger:     appLogger,
		imageAdapter:  imageAdapter,
		domainFactory: domainFactory,
	}
}

type imageUseCase struct {
	appLogger     *logger.AppLogger
	imageAdapter  ports.ImageAdapter
	domainFactory images.Factory
}

func (u *imageUseCase) GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error) {
	getter := u.domainFactory.NewImageGetter()
	portImpl := &imagePortImpl{
		imageAdapter: u.imageAdapter,
	}
	return getter.GetImage(ctx, imageID, portImpl)
}

func (u *imageUseCase) UploadImage(ctx context.Context, imageBinary []byte) (types.ImageID, error) {
	uploader := u.domainFactory.NewImageUploader()
	portImpl := &imagePortImpl{
		imageAdapter: u.imageAdapter,
	}
	return uploader.UploadImage(ctx, imageBinary, portImpl)
}

func (u *imageUseCase) DeleteImage(ctx context.Context, imageID types.ImageID) error {
	deleter := u.domainFactory.NewImageDeleter()
	portImpl := &imagePortImpl{
		imageAdapter: u.imageAdapter,
	}
	return deleter.DeleteImage(ctx, imageID, portImpl)
}

type imagePortImpl struct {
	imageAdapter ports.ImageAdapter
}

func (p *imagePortImpl) GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error) {
	return p.imageAdapter.GetImage(ctx, imageID)
}

func (p *imagePortImpl) UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error {
	return p.imageAdapter.UploadImage(ctx, imageID, imageBinary)
}

func (p *imagePortImpl) DeleteImage(ctx context.Context, imageID types.ImageID) error {
	return p.imageAdapter.DeleteImage(ctx, imageID)
}
