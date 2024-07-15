package images

import (
	"context"
	"io"

	"github.com/stlatica/stlatica/packages/backend/app/domains/images/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// ImageGetter is the interface for getting image.
type ImageGetter interface {
	// GetImage gets image.
	GetImage(ctx context.Context, imageID types.ImageID, outPort ports.ImageGetOutPort) (io.ReadCloser, error)
}

type imageGetter struct {
	appLogger *logger.AppLogger
}

func (g *imageGetter) GetImage(ctx context.Context,
	imageID types.ImageID, outPort ports.ImageGetOutPort) (io.ReadCloser, error) {
	return outPort.GetImage(ctx, imageID)
}
