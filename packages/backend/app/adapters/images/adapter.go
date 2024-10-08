package images

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/objectstorage"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images/ports"
)

const (
	// BucketName is the name of the bucket where images are stored.
	BucketName = "stlatica"
	// JpegMineType is the mine type for JPEG images.
	JpegMineType = "image/jpeg"
	// PngMineType is the mine type for PNG images.
	PngMineType = "image/png"
	// GifMineType is the mine type for GIF images.
	GifMineType = "image/gif"
)

type imageAdapter struct {
	client objectstorage.Client
}

// NewAdapter returns ImageAdapter.
func NewAdapter(client objectstorage.Client) ports.ImageAdapter {
	return &imageAdapter{
		client: client,
	}
}

func (a *imageAdapter) GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error) {
	return a.client.GetObject(ctx, BucketName, imageID.String())
}

func (a *imageAdapter) UploadImage(ctx context.Context, imageID types.ImageID, imageBinary []byte) error {
	mine, err := validateMineType(imageBinary)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(imageBinary)
	return a.client.PutObject(ctx, BucketName, imageID.String(), reader, mine)
}

func (a *imageAdapter) DeleteImage(ctx context.Context, imageID types.ImageID) error {
	return a.client.DeleteObject(ctx, BucketName, imageID.String())
}

func validateMineType(data []byte) (string, error) {
	mine := http.DetectContentType(data)

	switch {
	case mine == JpegMineType:
		return mine, nil
	case mine == PngMineType:
		return mine, nil
	case mine == GifMineType:
		return mine, nil
	default:
		return "", domainerrors.NewDomainError(fmt.Errorf("invalid file type, mine type: %s", mine),
			domainerrors.DomainErrorTypeInvalidData, "invalid file type")
	}
}
