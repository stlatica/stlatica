package v1

import (
	"encoding/base64"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images"
)

type imageController struct {
	imageUseCase images.ImageUseCase
}

// PostImageResponse is the response of UploadImage.
type PostImageResponse struct {
	ImageID string `json:"image_id"`
}

// GetImageResponse is the response of GetImage.
type GetImageResponse struct {
	ImageStream io.ReadCloser
}

func (c *imageController) GetImage(ectx echo.Context, imageIDStr string) (io.ReadCloser, error) {
	imageID, err := ulid.Parse(imageIDStr)
	if err != nil {
		return nil, err
	}
	return c.imageUseCase.GetImage(ectx.Request().Context(), types.ImageID(imageID))
}

func (c *imageController) UploadImage(ectx echo.Context, imageBase64 string) (*PostImageResponse, error) {
	imageBinary, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}
	imageID, err := c.imageUseCase.UploadImage(ectx.Request().Context(), imageBinary)
	if err != nil {
		return nil, err
	}

	return &PostImageResponse{
		ImageID: imageID.String(),
	}, nil
}

func (c *imageController) DeleteImage(ectx echo.Context, imageIDStr string) error {
	imageID, err := ulid.Parse(imageIDStr)
	if err != nil {
		return err
	}
	return c.imageUseCase.DeleteImage(ectx.Request().Context(), types.ImageID(imageID))
}
