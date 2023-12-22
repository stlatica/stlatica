package plats

import  (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/plats"
	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
)

// PlatsUseCase is the interface for getting plat.
type PlatUseCase interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Cntent, platID types.PlatID) (*entities.Plat, error)
}

// NewPlatUseCase returns PlatUseCase.
func NewPlatUseCase(domainFactory plats.Factory, platDAO dao.PlatDAO) PlatUseCase {
	return &PlatUseCase{
		platDAO:       platDAO,
		domainFactory: domainFactory,
	}
}

type PlatUseCase struct {
	platDAO       dao.PlatDAO
	domainFactory plats.Factory
}