package main

import (
	"github.com/labstack/echo/v4"
	mockv1controllers "github.com/stlatica/stlatica/packages/backend/mock/controllers/internalapi/v1"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	mockv1controllers.RegisterHandlers(e)

	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
