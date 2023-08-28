package main

import (
	v1controllers "github.com/atlatica/stlatica/app/controllers/internalapi/v1"
	"github.com/labstack/echo/v4"
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
