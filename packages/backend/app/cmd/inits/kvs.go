package inits

import (
	"context"
	"fmt"
	"os"

	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// NewKvsClient creates a new key-value store client.
func NewKvsClient(ctx context.Context, appLogger *logger.AppLogger) kvs.Client {
	return kvs.NewClient(ctx, appLogger,
		fmt.Sprintf("%s:%s", os.Getenv("STLATICA_KVS_HOST"), os.Getenv("STLATICA_KVS_PORT")),
		os.Getenv("STLATICA_KVS_PASSWORD"),
		0)
}
