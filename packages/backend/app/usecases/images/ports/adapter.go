package ports

import (
	"context"
	"io"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// ImageAdapter is the interface for getting image.
type ImageAdapter interface {
	// GetImage gets image.
	GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error)
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error
	// DeleteImage deletes image.
	DeleteImage(ctx context.Context, imageID types.ImageID) error
}
