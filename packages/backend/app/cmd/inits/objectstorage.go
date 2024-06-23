package inits

import (
	"context"
	"fmt"
	"os"

	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/objectstorage"
)

// NewObjectStorageClient creates a new object storage client.
func NewObjectStorageClient(ctx context.Context, appLogger *logger.AppLogger) objectstorage.Client {
	return objectstorage.NewClient(ctx, appLogger,
		fmt.Sprintf("%s:%s", os.Getenv("STLATICA_OBJECT_STORAGE_HOST"), os.Getenv("STLATICA_OBJECT_STORAGE_PORT")),
		os.Getenv("STLATICA_OBJECT_STORAGE_REGION"),
		os.Getenv("STLATICA_OBJECT_STORAGE_ACCESS_KEY"),
		os.Getenv("STLATICA_OBJECT_STORAGE_SECRET_KEY"),
	)
}
