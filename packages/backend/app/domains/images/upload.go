package images

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/images/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// ImageUploader is the interface for uploading image.
type ImageUploader interface {
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageBinary []byte, outPort ports.ImageUploadOutPort) (types.ImageID, error)
}

type imageUploader struct {
	appLogger *logger.AppLogger
}

func (g *imageUploader) UploadImage(ctx context.Context,
	imageBinary []byte, outPort ports.ImageUploadOutPort) (types.ImageID, error) {
	imageID := types.NewImageID()
	return imageID, outPort.UploadImage(ctx, imageID, imageBinary)
}
