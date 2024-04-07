package plats

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/logger"
)

// PlatGetter is the interface for getting plat.
type PlatGetter interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID, inPort ports.PlatGetInPort) (*entities.Plat, error)
}

type platGetter struct {
	appLogger *logger.AppLogger
}

func (g *platGetter) GetPlat(ctx context.Context,
	platID types.PlatID, inPort ports.PlatGetInPort) (*entities.Plat, error) {
	// TODO: 閲覧可否の判定ロジックを実装する https://github.com/stlatica/stlatica/issues/209
	return inPort.GetPlat(ctx, platID)
}
