package v1

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/logger"
)

// NewErrorHandler returns a new HTTPErrorHandler.
func NewErrorHandler(e *echo.Echo, appLogger *logger.AppLogger) echo.HTTPErrorHandler {
	return func(err error, ectx echo.Context) {
		// TODO: trace IDとuser IDをcontextから取得してログに出力する https://github.com/stlatica/stlatica/issues/403
		appLogger.Error(ectx.Request().Context(), err.Error(), "", time.Now(), "")
		e.DefaultHTTPErrorHandler(err, ectx)
	}
}
