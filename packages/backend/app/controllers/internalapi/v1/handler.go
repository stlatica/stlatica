package v1

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users"
)

type handler struct {
	*userController
	*platController
	*imageController
}

func newHandler(initContent ControllerInitContents) openapi.ServerInterface {
	return &handler{
		userController: &userController{
			userUseCase: initContent.UserUseCase,
		},
		platController: &platController{
			platUseCase: initContent.PlatUseCase,
		},
		imageController: &imageController{
			imageUseCase: initContent.ImageUseCase,
		},
	}
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(initContent ControllerInitContents, server *echo.Echo, appLogger *logger.AppLogger) {
	handler := newHandler(initContent)
	openapi.RegisterHandlers(server, handler)

	server.HTTPErrorHandler = NewErrorHandler(server, appLogger)
}

// GetUser is the handler for GET /users/{user_id}, ServerInterface implementation.
func (h *handler) GetUser(ectx echo.Context, userID string) error {
	response, err := h.userController.GetUser(ectx, userID)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) GetUsers(_ echo.Context, _ openapi.GetUsersParams) error {
	panic("implement me")
}

func (h *handler) CreateUser(_ echo.Context) error {
	panic("implement me")
}

func (h *handler) DeleteUser(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) PostPlat(ectx echo.Context) error {
	var plat openapi.PlatPost
	err := ectx.Bind(&plat)
	if err != nil {
		return err
	}
	response, err := h.platController.PostPlat(ectx, plat.UserId, plat.Content)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusCreated, response)
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

func (h *handler) GetTimelineByQuery(ectx echo.Context, params openapi.GetTimelineByQueryParams) error {
	response, err := h.platController.GetPlatsByQuery(ectx,
		string(params.Type), params.ToDate, params.FromDate, params.Limit)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) GetImage(_ echo.Context, _ string) error {
	panic("implement me")
}

func (h *handler) UploadImage(ectx echo.Context) error {
	imageBase64, err := io.ReadAll(ectx.Request().Body)
	if err != nil {
		return err
	}
	response, err := h.imageController.UploadImage(ectx, string(imageBase64))
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusCreated, response)
}

func (h *handler) PostFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) DeleteFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) Login(_ echo.Context) error {
	panic("implement me")
}

// ControllerInitContents is the struct to hold the dependencies for the controller.
type ControllerInitContents struct {
	UserUseCase  users.UserUseCase
	PlatUseCase  plats.PlatUseCase
	ImageUseCase images.ImageUseCase
}
