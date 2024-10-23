package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/auth"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// NewAuthMiddleware creates a new Auth middleware
func NewAuthMiddleware[T auth.SignType](logger logger.AppLogger, client kvs.Client,
	alg jwa.SignatureAlgorithm, key *T, lifetime int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			header := ctx.Request().Header.Get("Authorization")
			token := auth.NewAccessToken[T]([]byte(header))

			requestPath := ctx.Request().URL.Path

			var isVerified bool
			var err error
			if requestPath == "/internal/v1/timelines/" {
				isVerified, err = token.QuickVerify(ctx.Request().Context(), client, alg, key, lifetime)
			} else {
				isVerified, err = token.Verify(ctx.Request().Context(), client, alg, key, lifetime)
			}

			if err != nil {
				logger.Error(ctx.Request().Context(), "Error verifying", "", time.Now(), "")
				return errors.NewDomainError(err,
					errors.DomainErrorTypeInternalServerError, "Internal Server Error")
			}
			if !isVerified {
				return errors.NewDomainError(err,
					errors.DomainErrorTypeUnauthorized, "Unauthorized")
			}
			return nil
		}
	}
}
