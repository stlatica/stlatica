package dao

import (
	"context"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// ActorDAO is the interface for getting actor.
type ActorDAO interface {
	// GetActor returns actor.
	GetActor(ctx context.Context, actorID types.ActorID) (*domainentities.Actor, error)
	// GetActorByName returns actor by name.
	GetActorByName(ctx context.Context, actorName string) (*domainentities.Actor, error)
}

// NewActorDAO returns ActorDAO.
func NewActorDAO() ActorDAO {
	ctxExecutor := boil.GetContextDB()
	return &actorDAO{
		ctxExecutor: ctxExecutor,
	}
}

type actorDAO struct {
	ctxExecutor boil.ContextExecutor
}

func (dao *actorDAO) GetActor(ctx context.Context, actorID types.ActorID) (*domainentities.Actor, error) {
	entity, err := entities.FindActor(ctx, dao.ctxExecutor, actorID)
	if err != nil {
		return nil, err
	}

	return &domainentities.Actor{
		ActorID:   entity.ActorID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
	}, nil
}

func (dao *actorDAO) GetActorByName(ctx context.Context, actorName string) (*domainentities.Actor, error) {
	query := entities.ActorWhere.Name.EQ(actorName)
	entity, err := entities.Actors(query).One(ctx, dao.ctxExecutor)
	if err != nil {
		return nil, err
	}

	return &domainentities.Actor{
		ActorID:   entity.ActorID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
	}, nil
}
