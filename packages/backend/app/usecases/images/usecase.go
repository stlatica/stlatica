package images

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/images"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images/ports"
)

// ImageUseCase is the interface for handling image.
type ImageUseCase interface {
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageBinary []byte) (types.ImageID, error)
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

func (u *imageUseCase) UploadImage(ctx context.Context, imageBinary []byte) (types.ImageID, error) {
	uploader := u.domainFactory.NewImageUploader()
	portImpl := &imagePortImpl{
		imageAdapter: u.imageAdapter,
	}
	return uploader.UploadImage(ctx, imageBinary, portImpl)
}

type imagePortImpl struct {
	imageAdapter ports.ImageAdapter
}

func (p *imagePortImpl) UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error {
	return p.imageAdapter.UploadImage(ctx, imageID, imageBinary)
}
