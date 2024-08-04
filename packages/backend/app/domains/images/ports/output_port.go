package ports

import (
	"context"
	"io"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// ImageGetOutPort is the interface for getting image.
type ImageGetOutPort interface {
	// GetImage gets image.
	GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error)
}

// ImageUploadOutPort is the interface for uploading image.
type ImageUploadOutPort interface {
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error
}

// ImageDeleteOutPort is the interface for deleting image.
type ImageDeleteOutPort interface {
	// DeleteImage deletes image.
	DeleteImage(ctx context.Context, imageID types.ImageID) error
}
