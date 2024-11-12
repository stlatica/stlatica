package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// PlatCreateInPort is the interface for posting plat.
type PlatCreateInPort interface {
	// CreatePlat returns Plat.
	CreatePlat(ctx context.Context, UserID types.UserID, content string) (*entities.Plat, error)
}
