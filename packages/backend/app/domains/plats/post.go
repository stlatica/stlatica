package plats

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// PlatCreator is the interface for createing plat.
type PlatCreator interface {
	// CreatePlat returns new plat.
	CreatePlat(ctx context.Context, userID types.UserID, content string,
		inPort ports.PlatCreateInPort) (*entities.Plat, error)
}

type platCreator struct {
	appLogger *logger.AppLogger
}

func (g *platCreator) CreatePlat(ctx context.Context, userID types.UserID,
	content string, inPort ports.PlatCreateInPort) (*entities.Plat, error) {
	return inPort.CreatePlat(ctx, userID, content)
}
