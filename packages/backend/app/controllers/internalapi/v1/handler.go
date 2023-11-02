package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
)

type handler struct {
	*userController
}

func newHandler() openapi.ServerInterface {
	return &handler{
		userController: &userController{},
	}
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(server *echo.Echo) {
	handler := newHandler()
	openapi.RegisterHandlers(server, handler)
}

// GetUser is the handler for GET /users/{user_id}, ServerInterface implementation.
func (h *handler) GetUser(ectx echo.Context, userID string) error {
	response, err := h.userController.GetUser(userID)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}
