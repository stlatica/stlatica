package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
)

type platController struct {
	platUseCase plats.PlatUseCase
}

// GetPlatResponse is the response of GetPlat.
type GetPlatResponse struct {
	PlatID string `json:"plat_id"`
}

func (c *platController) PostPlat(ectx echo.Context, actorID string, content string) (*GetPlatResponse, error) {
	plat, err := c.platUseCase.CreatePlat(ectx.Request().Context(), actorID, content)
	if err != nil {
		return &GetPlatResponse{}, err
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID(),
	}, nil
}

func (c *platController) GetPlat(ectx echo.Context, platID string) (*GetPlatResponse, error) {
	plat, err := c.platUseCase.GetPlat(ectx.Request().Context(), platID)
	if err != nil {
		return &GetPlatResponse{}, err
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID(),
	}, nil
}

