package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stlatica/stlatica/packages/backend/app/cmd/inits"
	v1controllers "github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1"
	userdomain "github.com/stlatica/stlatica/packages/backend/app/domains/users"
	platdomain "github.com/stlatica/stlatica/packages/backend/app/domains/plats"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	userusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/users"
	platusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	// load environment variables
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	// initialize logger
	loggerFactory := inits.NewLoggerFactory()
	appLogger := loggerFactory.AppLogger()

	// initialize echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// initialize database
	inits.NewDB()

	// initialize controllers
	userDAO := dao.NewUserDAO()
	platDAO := dao.NewPlatDAO()
	userFactory := userdomain.NewFactory(appLogger)
	platFactory := platdomain.NewFactory(appLogger)
	initContent := &v1controllers.ControllerInitContents{
		UserUseCase: userusecase.NewUserUseCase(appLogger, userFactory, userDAO),
		PlatUseCase:  platusecase.NewPlatUseCase(appLogger, platFactory, platDAO),
	}
	v1controllers.RegisterHandlers(*initContent, e)

	err = e.Start(fmt.Sprintf(`:%s`, os.Getenv("STLATICA_SERVER_PORT")))
	if err != nil {
		panic(err)
	}
}
