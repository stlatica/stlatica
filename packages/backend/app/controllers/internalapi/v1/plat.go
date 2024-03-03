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
	actorID, err1 := ulid.Parse(actorIDStr)
	if err1 != nil {
		return &GetPlatResponse{}, err1
	}
	plat, err2 := c.platUseCase.CreatePlat(ectx.Request().Context(), types.ActorID(actorID), content)
	if err2 != nil {
		return &GetPlatResponse{}, err2
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID().String(),
	}, nil
}

func (c *platController) GetPlat(ectx echo.Context, platIDStr string) (*GetPlatResponse, error) {
	platID, err1 := ulid.Parse(platIDStr)
	if err1 != nil {
		return &GetPlatResponse{}, err1
	}
	plat, err2 := c.platUseCase.GetPlat(ectx.Request().Context(), types.PlatID(platID))
	if err2 != nil {
		return &GetPlatResponse{}, err2
	}
	return &GetPlatResponse{
		PlatID: plat.GetPlatID().String(),
	}, nil
}
