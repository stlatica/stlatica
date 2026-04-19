package v1

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	authusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/auth"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/images"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users"
)

type handler struct {
	*authController
	*userController
	*platController
	*imageController
}

func newHandler(initContent ControllerInitContents) *handler {
	return &handler{
		authController: &authController{
			authUseCase: initContent.AuthUseCase,
		},
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
	serverHandler := newHandler(initContent)
	wrapper := openapi.ServerInterfaceWrapper{Handler: serverHandler}
	protectedRoutes := server.Group("", requireAuthentication(serverHandler.authController))

	server.POST("/internal/v1/auth/refresh", wrapper.RefreshAuth)
	server.GET("/internal/v1/images/:image_id", wrapper.GetImage)
	server.POST("/internal/v1/login", wrapper.Login)
	server.GET("/internal/v1/plats/:plat_id", wrapper.GetPlat)
	server.GET("/internal/v1/timelines", wrapper.GetTimelineByQuery)
	server.GET("/internal/v1/timelines/:timeline_id", wrapper.GetTimeline)
	server.GET("/internal/v1/users", wrapper.GetUsers)
	server.POST("/internal/v1/users", wrapper.CreateUser)
	server.GET("/internal/v1/users/:user_id", wrapper.GetUser)
	server.GET("/internal/v1/users/:user_id/followers", wrapper.GetFollowers)
	server.GET("/internal/v1/users/:user_id/follows", wrapper.GetFollows)
	server.GET("/internal/v1/users/:user_id/icon", wrapper.GetUserIcon)

	protectedRoutes.POST("/internal/v1/images", wrapper.UploadImage)
	protectedRoutes.DELETE("/internal/v1/images/:image_id", wrapper.DeleteImage)
	protectedRoutes.POST("/internal/v1/plats", wrapper.PostPlat)
	protectedRoutes.DELETE("/internal/v1/plats/:plat_id", wrapper.DeletePlat)
	protectedRoutes.DELETE("/internal/v1/plats/:plat_id/favorites", wrapper.DeleteFavorite)
	protectedRoutes.POST("/internal/v1/plats/:plat_id/favorites", wrapper.PostFavorite)
	protectedRoutes.DELETE("/internal/v1/users/:user_id", wrapper.DeleteUser)
	protectedRoutes.DELETE("/internal/v1/users/:user_id/follow", wrapper.DeleteFollow)
	protectedRoutes.POST("/internal/v1/users/:user_id/follow", wrapper.PostFollow)

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

func (h *handler) CreateUser(ectx echo.Context) error {
	var user openapi.CreateUserJSONBody
	err := ectx.Bind(&user)
	if err != nil {
		return err
	}
	response, err := h.userController.CreateUser(ectx, user.Name, user.UserId, user.Email, user.IconImageId)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusCreated, response)
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
	userID, err := authenticatedUserID(ectx)
	if err != nil {
		return err
	}
	response, err := h.platController.PostPlat(ectx, userID, plat.Content)
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
		params.UserId, string(params.Type), params.ToDate, params.FromDate, params.Limit)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) GetImage(ectx echo.Context, imageIDStr string) error {
	imageStream, err := h.imageController.GetImage(ectx, imageIDStr)
	if err != nil {
		return err
	}
	response := ectx.Response()
	_, err = io.Copy(response, imageStream)
	if err != nil {
		return err
	}
	return nil
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

func (h *handler) DeleteImage(ectx echo.Context, imageIDStr string) error {
	err := h.imageController.DeleteImage(ectx, imageIDStr)
	if err != nil {
		return err
	}
	return ectx.NoContent(http.StatusNoContent)
}

func (h *handler) PostFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) DeleteFavorite(_ echo.Context, _ openapi.PlatId) error {
	panic("implement me")
}

func (h *handler) GetFollows(ctx echo.Context, userID openapi.UserId, params openapi.GetFollowsParams) error {
	response, err := h.userController.GetFollows(ctx, userID, params.UserPaginationId, params.Limit)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *handler) PostFollow(ectx echo.Context, userID openapi.UserId) error {
	authenticatedActorID, err := authenticatedUserID(ectx)
	if err != nil {
		return err
	}
	err = h.userController.PostFollow(ectx, userID, authenticatedActorID)
	if err != nil {
		return err
	}
	return ectx.NoContent(http.StatusOK)
}

func (h *handler) DeleteFollow(_ echo.Context, _ openapi.UserId) error {
	panic("implement me")
}

func (h *handler) GetFollowers(ctx echo.Context, userID openapi.UserId, params openapi.GetFollowersParams) error {
	response, err := h.userController.GetFollowers(ctx, userID, params.UserPaginationId, params.Limit)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *handler) Login(ectx echo.Context) error {
	var loginParams openapi.LoginJSONBody
	err := ectx.Bind(&loginParams)
	if err != nil {
		return err
	}

	response, err := h.authController.Login(
		ectx,
		stringValue(loginParams.MailAddress),
		stringValue(loginParams.PreferredUserId),
		stringValue(loginParams.Password),
	)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, response)
}

func (h *handler) RefreshAuth(ectx echo.Context) error {
	if err := h.authController.Refresh(ectx); err != nil {
		return err
	}
	return ectx.NoContent(http.StatusNoContent)
}

func (h *handler) GetUserIcon(ectx echo.Context, userID openapi.UserId) error {
	imageStream, err := h.userController.GetUserIcon(ectx, userID)
	if err != nil {
		return err
	}
	response := ectx.Response()
	_, err = io.Copy(response, imageStream)
	if err != nil {
		return err
	}
	return nil
}

// ControllerInitContents is the struct to hold the dependencies for the controller.
type ControllerInitContents struct {
	AuthUseCase  authusecase.AuthenticationUseCase
	UserUseCase  users.UserUseCase
	PlatUseCase  plats.PlatUseCase
	ImageUseCase images.ImageUseCase
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
