package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

const authenticatedUserIDContextKey = "authenticated_user_id"

func requireAuthentication(controller *authController) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			accessCookie, err := ectx.Cookie(accessTokenCookieName)
			if err != nil {
				return unauthorizedError(err)
			}

			userID, err := controller.authUseCase.AuthenticateAccessToken(
				ectx.Request().Context(),
				accessCookie.Value,
			)
			if err != nil {
				return err
			}

			ectx.Set(authenticatedUserIDContextKey, userID)
			return next(ectx)
		}
	}
}

func authenticatedUserID(ectx echo.Context) (types.UserID, error) {
	value := ectx.Get(authenticatedUserIDContextKey)
	userID, ok := value.(string)
	if !ok || userID == "" {
		return types.UserID{}, unauthorizedError(echo.ErrUnauthorized)
	}

	parsedUserID, err := types.NewUserIDFromString(userID)
	if err != nil {
		return types.UserID{}, unauthorizedError(err)
	}
	return parsedUserID, nil
}
