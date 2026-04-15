package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

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
	authusecase "github.com/stlatica/stlatica/packages/backend/app/usecases/auth"
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
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:3000",
				"http://127.0.0.1:3000",
				"http://localhost:5173",
				"http://127.0.0.1:5173",
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
				echo.HeaderCookie,
				echo.HeaderAuthorization,
			},
			AllowCredentials: true,
		}))
	}

	// initialize database
	inits.NewDB()

	// initialize controllers
	kvsClient := inits.NewKvsClient(ctx, appLogger)
	userDAO := dao.NewUserDAO()
	userAuthCredentialDAO := dao.NewUserAuthCredentialDAO()
	platDAO := dao.NewPlatDAO()
	objectStorageClient := inits.NewObjectStorageClient(ctx, appLogger)
	imageAdapter := images.NewAdapter(objectStorageClient)
	userFactory := userdomain.NewFactory(appLogger)
	platFactory := platdomain.NewFactory(appLogger)
	imageFactory := imagedomain.NewFactory(appLogger)
	authConfig := authusecase.Config{
		JWTSecret:       authSecret(),
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 30 * 24 * time.Hour,
	}
	initContent := &v1controllers.ControllerInitContents{
		AuthUseCase: authusecase.NewAuthenticationUseCase(
			userDAO,
			userAuthCredentialDAO,
			kvsClient,
			authConfig,
		),
		UserUseCase:  userusecase.NewUserUseCase(appLogger, userFactory, imageFactory, userDAO, imageAdapter),
		PlatUseCase:  platusecase.NewPlatUseCase(appLogger, platFactory, platDAO),
		ImageUseCase: imageusecase.NewImageUseCase(appLogger, imageAdapter, imageFactory),
	}
	v1controllers.RegisterHandlers(*initContent, e, appLogger)

	err = e.Start(fmt.Sprintf(`:%s`, os.Getenv("STLATICA_SERVER_PORT")))
	if err != nil {
		panic(err)
	}
}

func authSecret() string {
	secret := os.Getenv("STLATICA_AUTH_JWT_SECRET")
	if secret != "" {
		return secret
	}
	panic("STLATICA_AUTH_JWT_SECRET is required")
}
