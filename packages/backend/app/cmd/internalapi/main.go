package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/cmd/inits"
	v1controllers "github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1"
	actordomain "github.com/stlatica/stlatica/packages/backend/app/domains/actors"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	actorusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/actors"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	// initialize echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// initialize database
	inits.NewDB()

	// initialize controllers
	actorDAO := dao.NewActorDAO()
	actorFactory := actordomain.NewFactory()
	initContent := &v1controllers.ControllerInitContents{
		ActorUseCase: actorusecase.NewActorUseCase(actorFactory, actorDAO),
	}
	v1controllers.RegisterHandlers(*initContent, e)

	// TODO: port to be read from environment variables https://github.com/stlatica/stlatica/issues/50
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
