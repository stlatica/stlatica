package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

type PlatGetInPort interface {
	// GetPlat returns Plat. 
	GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error)
}