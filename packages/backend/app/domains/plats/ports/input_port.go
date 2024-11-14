package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// PlatGetInPort is the interface for getting plat.
type PlatGetInPort interface {
	// GetPlat returns Plat.
	GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error)
	// GetPlats returns Plats.
	GetPlats(ctx context.Context, params PlatsGetParams) ([]*entities.Plat, error)
}
