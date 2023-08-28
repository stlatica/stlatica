package v1

import (
	"net/http"

	"github.com/atlatica/stlatica/app/controllers/internalapi/v1/openapi"
	"github.com/labstack/echo/v4"
)

type handler struct {
	*sampleUserController
}

func newHandler() openapi.ServerInterface {
	return &handler{
		sampleUserController: &sampleUserController{},
	}
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(server *echo.Echo) {
	handler := newHandler()
	openapi.RegisterHandlers(server, handler)
}

// GetSampleUser is the handler for GET /sample_users/{sample_user_id}, ServerInterface implementation.
func (h *handler) GetSampleUser(ectx echo.Context, sampleUserID string) error {
	response, err := h.sampleUserController.GetSampleUser(sampleUserID)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}
