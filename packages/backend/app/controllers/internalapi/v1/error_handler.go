package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1/openapi"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

type errorResponse struct {
	Code    openapi.ErrorResponseCode `json:"code"`
	Message string                    `json:"message"`
}

// NewErrorHandler returns a new HTTPErrorHandler.
func NewErrorHandler(_ *echo.Echo, appLogger *logger.AppLogger) echo.HTTPErrorHandler {
	return func(err error, ectx echo.Context) {
		var responseStatusCode int
		var responseErrorCode openapi.ErrorResponseCode
		var responseMessage string

		var domainError domainerrors.DomainError
		if errors.As(err, &domainError) {
			switch domainError.DomainErrorType() {
			case domainerrors.DomainErrorTypeNotFound:
				responseStatusCode = http.StatusNotFound
				responseErrorCode = openapi.NOTFOUND
			case domainerrors.DomainErrorTypeInvalidData:
				responseStatusCode = http.StatusBadRequest
				responseErrorCode = openapi.BADREQUEST
			case domainerrors.DomainErrorTypeUnauthorized:
				responseStatusCode = http.StatusUnauthorized
				responseErrorCode = openapi.UNAUTHORIZED
			case domainerrors.DomainErrorTypeInternalServerError:
				responseStatusCode = http.StatusInternalServerError
				responseErrorCode = openapi.INTERNALSERVERERROR
			default:
				responseStatusCode = http.StatusInternalServerError
				responseErrorCode = openapi.INTERNALSERVERERROR
			}
			responseMessage = domainError.Error()
		} else {
			responseStatusCode = http.StatusInternalServerError
			responseErrorCode = openapi.INTERNALSERVERERROR
			responseMessage = err.Error()
		}
		// TODO: trace IDとuser IDをcontextから取得してログに出力する https://github.com/stlatica/stlatica/issues/403
		appLogger.Error(ectx.Request().Context(), err.Error(), "", time.Now(), "")

		err = ectx.JSON(responseStatusCode,
			errorResponse{
				Code:    responseErrorCode,
				Message: responseMessage,
			})
		if err != nil {
			// TODO: trace IDとuser IDをcontextから取得してログに出力する https://github.com/stlatica/stlatica/issues/403
			appLogger.Error(ectx.Request().Context(),
				fmt.Sprintf("failed to send error response: %s", err.Error()), "", time.Now(), "")
		}
	}
}
