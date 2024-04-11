package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
)

type platController struct {
	platUseCase plats.PlatUseCase
}

// GetPlatResponse is the response of GetPlat.
type GetPlatResponse struct {
	PlatID string `json:"plat_id"`
}

func (c *platController) PostPlat(ectx echo.Context, actorIDStr string, content string) (*GetPlatResponse, error) {
	actorID, err := ulid.Parse(actorIDStr)
	if err != nil {
		return nil, err
	}
	plat, err := c.platUseCase.CreatePlat(ectx.Request().Context(), types.ActorID(actorID), content)
	if err != nil {
		return nil, err
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID().String(),
	}, nil
}

func (c *platController) GetPlat(ectx echo.Context, platIDStr string) (*GetPlatResponse, error) {
	platID, err := ulid.Parse(platIDStr)
	if err != nil {
		return nil, err
	}
	plat, err := c.platUseCase.GetPlat(ectx.Request().Context(), types.PlatID(platID))
	if err != nil {
		return nil, err
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID().String(),
	}, nil
}
