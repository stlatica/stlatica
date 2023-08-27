package main

import (
	"github.com/labstack/echo/v4"
	v1controllers "github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1"
)

func main() {
	// initialize echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	v1controllers.RegisterHandlers(e)

	// TODO: port to be read from environment variables https://github.com/stlatica/stlatica/issues/50
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
