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
	CreatePlat(ctx context.Context, userID types.UserID, content string) (*domainentities.Plat, error)
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
		UserID:    entity.UserID,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}, nil
}

func (dao *platDAO) CreatePlat(ctx context.Context, userID types.UserID, content string) (
	*domainentities.Plat, error) {
	plat := entities.Plat{
		PlatID:    types.NewPlatID(),
		UserID:    userID,
		Content:   content,
		CreatedAt: types.NewUnixTimeFromCurrentTime(),
	}
	ctx = boil.SkipTimestamps(ctx)
	err := plat.Insert(ctx, dao.ctxExecutor, boil.Infer())
	return &domainentities.Plat{
		PlatID:    plat.PlatID,
		UserID:    plat.UserID,
		Content:   plat.Content,
		CreatedAt: plat.CreatedAt,
	}, err
}
