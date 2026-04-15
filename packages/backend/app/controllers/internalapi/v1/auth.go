package v1

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	authusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/auth"
)

const (
	accessTokenCookieName  = "access_token"
	refreshTokenCookieName = "refresh_token"
)

type authController struct {
	authUseCase authusecase.AuthenticationUseCase
}

// LoginResponse is the response of Login.
type LoginResponse struct {
	UserID string `json:"user_id"`
}

func (c *authController) Login(
	ectx echo.Context,
	mailAddress string,
	preferredUserID string,
	password string,
) (*LoginResponse, error) {
	result, err := c.authUseCase.Login(ectx.Request().Context(), authusecase.LoginParams{
		MailAddress:     mailAddress,
		PreferredUserID: preferredUserID,
		Password:        password,
	})
	if err != nil {
		return nil, err
	}

	ectx.SetCookie(newAuthCookie(
		accessTokenCookieName,
		result.AccessToken,
		result.AccessTokenExpiresAt,
	))
	ectx.SetCookie(newAuthCookie(
		refreshTokenCookieName,
		result.RefreshToken,
		result.RefreshTokenExpiresAt,
	))

	return &LoginResponse{UserID: result.PreferredUserID}, nil
}

func (c *authController) Refresh(ectx echo.Context) error {
	refreshCookie, err := ectx.Cookie(refreshTokenCookieName)
	if err != nil {
		return unauthorizedError(err)
	}

	result, err := c.authUseCase.Refresh(ectx.Request().Context(), refreshCookie.Value)
	if err != nil {
		return err
	}

	ectx.SetCookie(newAuthCookie(
		accessTokenCookieName,
		result.AccessToken,
		result.AccessTokenExpiresAt,
	))
	ectx.SetCookie(newAuthCookie(
		refreshTokenCookieName,
		result.RefreshToken,
		result.RefreshTokenExpiresAt,
	))

	return nil
}

func newAuthCookie(name string, value string, expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  expiresAt,
		MaxAge:   int(time.Until(expiresAt).Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   isSecureCookieEnv(),
	}
}

func isSecureCookieEnv() bool {
	switch os.Getenv("GO_ENV") {
	case "local", "local.docker":
		return false
	default:
		return true
	}
}

func unauthorizedError(err error) error {
	return domainerrors.NewDomainError(err, domainerrors.DomainErrorTypeUnauthorized, "unauthorized")
}
