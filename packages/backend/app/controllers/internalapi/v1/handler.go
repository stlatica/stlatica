package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/actors"
	"github.com/tidwall/gjson"
)

type handler struct {
	*userController
	*platController
}

func newHandler(initContent ControllerInitContents) openapi.ServerInterface {
	return &handler{
		userController: &userController{
			actorUseCase: initContent.ActorUseCase,
		},
	}
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(initContent ControllerInitContents, server *echo.Echo) {
	handler := newHandler(initContent)
	openapi.RegisterHandlers(server, handler)
}

// GetUser is the handler for GET /users/{user_id}, ServerInterface implementation.
func (h *handler) GetUser(ectx echo.Context, userID string) error {
	response, err := h.userController.GetUser(ectx, userID)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) PostPlat(ectx echo.Context) error {
	var plat string
	err1 := ectx.Bind(&plat)
	if err1 != nil {
		return err1
	}
	actorIDStr := gjson.Get(plat, "user_id")
	content := gjson.Get(plat, "content")
	response, err2 := h.platController.PostPlat(ectx, actorIDStr.String(), content.String())
	if err2 != nil {
		return err2
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) DeletePlat(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) GetPlat(ectx echo.Context, platIDStr string) error {
	response, err := h.platController.GetPlat(ectx, platIDStr)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) GetTimeline(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) GetImage(_ echo.Context, _ string) error {
	panic("implement me")
}

// ControllerInitContents is the struct to hold the dependencies for the controller.
type ControllerInitContents struct {
	ActorUseCase actors.ActorUseCase
}
