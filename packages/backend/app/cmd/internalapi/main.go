package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stlatica/stlatica/packages/backend/app/adapters/images"
	"github.com/stlatica/stlatica/packages/backend/app/cmd/inits"
	v1controllers "github.com/stlatica/stlatica/packages/backend/app/controllers/internalapi/v1"
	imagedomain "github.com/stlatica/stlatica/packages/backend/app/domains/images"
	platdomain "github.com/stlatica/stlatica/packages/backend/app/domains/plats"
	userdomain "github.com/stlatica/stlatica/packages/backend/app/domains/users"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	imageusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/images"
	platusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/plats"
	userusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/users"

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

	// initialize context
	ctx := context.Background()

	// initialize echo server
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// CORS
	if os.Getenv("GO_ENV") == "local" || os.Getenv("GO_ENV") == "local.docker" {
		e.Use(middleware.CORS())
	}

	// initialize database
	inits.NewDB()

	// initialize controllers
	userDAO := dao.NewUserDAO()
	platDAO := dao.NewPlatDAO()
	objectStorageClient := inits.NewObjectStorageClient(ctx, appLogger)
	imageAdapter := images.NewAdapter(objectStorageClient)
	userFactory := userdomain.NewFactory(appLogger)
	platFactory := platdomain.NewFactory(appLogger)
	imageFactory := imagedomain.NewFactory(appLogger)
	initContent := &v1controllers.ControllerInitContents{
		UserUseCase:  userusecase.NewUserUseCase(appLogger, userFactory, userDAO),
		PlatUseCase:  platusecase.NewPlatUseCase(appLogger, platFactory, platDAO),
		ImageUseCase: imageusecase.NewImageUseCase(appLogger, imageAdapter, imageFactory),
	}
	v1controllers.RegisterHandlers(*initContent, e, appLogger)

	err = e.Start(fmt.Sprintf(`:%s`, os.Getenv("STLATICA_SERVER_PORT")))
	if err != nil {
		panic(err)
	}
}
