package v1

import (
	"time"

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
	PlatID        string    `json:"plat_id"`
	UserID        string    `json:"user_id"`
	Content       string    `json:"content"`
	ImageURLs     []string  `json:"image_urls"`
	ReplyCount    uint64    `json:"reply_count"`
	RePlatCount   uint64    `json:"replat_count"`
	FavoriteCount uint64    `json:"favorite_count"`
	CreatedAt     time.Time `json:"created_at"`
}

func (c *platController) PostPlat(ectx echo.Context, userIDStr string, content string) (*GetPlatResponse, error) {
	userID, err := ulid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}
	plat, err := c.platUseCase.CreatePlat(ectx.Request().Context(), types.UserID(userID), content)
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
		PlatID:        plat.GetPlatID().String(),
		UserID:        plat.GetUserID().String(),
		Content:       plat.GetContent(),
		ImageURLs:     []string{}, // TODO: Return actual value https://github.com/stlatica/stlatica/issues/441
		ReplyCount:    0,          // TODO: Return actual value https://github.com/stlatica/stlatica/issues/441
		RePlatCount:   0,          // TODO: Return actual value https://github.com/stlatica/stlatica/issues/441
		FavoriteCount: 0,          // TODO: Return actual value https://github.com/stlatica/stlatica/issues/441
		CreatedAt:     plat.GetCreatedAt().ConvertToTime(),
	}, nil
}
