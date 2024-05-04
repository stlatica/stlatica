package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// ImageAdapter is the interface for getting image.
type ImageAdapter interface {
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error
}
