package images

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/images/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// ImageDeleter is the interface for deleting image.
type ImageDeleter interface {
	// DeleteImage deletes image.
	DeleteImage(ctx context.Context, imageID types.ImageID, outPort ports.ImageDeleteOutPort) error
}

type imageDeleter struct {
	appLogger *logger.AppLogger
}

func (g *imageDeleter) DeleteImage(ctx context.Context,
	imageID types.ImageID, outPort ports.ImageDeleteOutPort) error {
	return outPort.DeleteImage(ctx, imageID)
}
