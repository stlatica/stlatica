package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

type tmpErrorMessage struct {
	Message string `json:"message"`
}

// NewErrorHandler returns a new HTTPErrorHandler.
func NewErrorHandler(_ *echo.Echo, _ *logger.AppLogger) echo.HTTPErrorHandler {
	return func(err error, ectx echo.Context) {
		// TODO: 実際のエラーハンドリング処理を実装する https://github.com/stlatica/stlatica/issues/590
		_ = ectx.JSON(http.StatusInternalServerError,
			tmpErrorMessage{
				Message: err.Error(),
			})
		// TODO: trace IDとuser IDをcontextから取得してログに出力する https://github.com/stlatica/stlatica/issues/403
		// appLogger.Error(ectx.Request().Context(), err.Error(), "", time.Now(), "")
		// e.DefaultHTTPErrorHandler(err, ectx)
	}
}
