package dao

import (
	"context"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// PlatDAO is the interface for getting plat.
type PlatDAO interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID) (*domainentities.Plat, error)
	CreatePlat(ctx context.Context, ActorID types.ActorID, content string) (*domainentities.Plat, error)
}

// NewPlatDAO returns PlatDAO.
func NewPlatDAO() PlatDAO {
	ctxExecutor := boil.GetContextDB()
	return &platDAO{
		ctxExecutor: ctxExecutor,
	}
}

type platDAO struct {
	ctxExecutor boil.ContextExecutor
}

func (dao *platDAO) GetPlat(ctx context.Context, platID types.PlatID) (*domainentities.Plat, error) {
	entity, err := entities.FindPlat(ctx, dao.ctxExecutor, platID)
	if err != nil {
		return nil, err
	}

	return &domainentities.Plat{
		PlatID:    entity.PlatID,
		ActorID:   entity.ActorID,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}, nil
}

func (dao *platDAO) CreatePlat(ctx context.Context, actorID types.ActorID, content string) (
	*domainentities.Plat, error) {
	plat := entities.Plat{
		PlatID:  types.NewPlatID(),
		ActorID: actorID,
		Content: content,
	}
	err := plat.Insert(ctx, dao.ctxExecutor, boil.Infer())
	return &domainentities.Plat{
		PlatID:    plat.PlatID,
		ActorID:   plat.ActorID,
		Content:   plat.Content,
		CreatedAt: plat.CreatedAt,
	}, err
}
