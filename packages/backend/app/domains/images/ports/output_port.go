package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// ImageUploadOutPort is the interface for uploading image.
type ImageUploadOutPort interface {
	// UploadImage uploads image.
	UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error
}
