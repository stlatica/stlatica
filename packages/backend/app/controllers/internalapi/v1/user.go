package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/actors"
)

type userController struct {
	actorUseCase actors.ActorUseCase
}

// GetUserResponse is the response of GetUser.
type GetUserResponse struct {
	UserID string `json:"user_id"`
}

// GetUser converts request data and calls usecase to get actor.
func (c *userController) GetUser(ectx echo.Context, userID string) (*GetUserResponse, error) {
	actor, err := c.actorUseCase.GetActorByName(ectx.Request().Context(), userID)
	if err != nil {
		return &GetUserResponse{}, err
	}

	return &GetUserResponse{
		UserID: actor.GetName(),
	}, nil
}
