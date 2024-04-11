package plats

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/logger"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
)

// PlatUseCase is the interface for getting plat.
type PlatUseCase interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error)
	CreatePlat(ctx context.Context, actorID types.ActorID, content string) (*entities.Plat, error)
}

// NewPlatUseCase returns PlatUseCase.
func NewPlatUseCase(appLogger *logger.AppLogger, domainFactory plats.Factory, platDAO dao.PlatDAO) PlatUseCase {
	return &platUseCase{
		appLogger:     appLogger,
		platDAO:       platDAO,
		domainFactory: domainFactory,
	}
}

type platUseCase struct {
	appLogger     *logger.AppLogger
	platDAO       dao.PlatDAO
	domainFactory plats.Factory
}

func (u *platUseCase) GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error) {
	getter := u.domainFactory.NewPlatGetter()
	portImpl := &platPortImpl{
		platDAO: u.platDAO,
	}
	return getter.GetPlat(ctx, platID, portImpl)
}

func (u *platUseCase) CreatePlat(ctx context.Context, actorID types.ActorID, content string) (*entities.Plat, error) {
	creator := u.domainFactory.NewPlatCreator()
	portImpl := &platPortImpl{
		platDAO: u.platDAO,
	}
	return creator.CreatePlat(ctx, actorID, content, portImpl)
}

type platPortImpl struct {
	platDAO dao.PlatDAO
}

func (p *platPortImpl) GetPlat(ctx context.Context, platID types.PlatID) (*entities.Plat, error) {
	return p.platDAO.GetPlat(ctx, platID)
}

func (p *platPortImpl) CreatePlat(ctx context.Context, actorID types.ActorID, content string) (*entities.Plat, error) {
	return p.platDAO.CreatePlat(ctx, actorID, content)
}
