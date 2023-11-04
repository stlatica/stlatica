package actors

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/actors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
)

// ActorUseCase is the interface for getting actor.
type ActorUseCase interface {
	// GetActor returns actor.
	GetActor(ctx context.Context, actorID types.ActorID) (*entities.Actor, error)
	// GetActorByName returns actor by name.
	GetActorByName(ctx context.Context, actorName string) (*entities.Actor, error)
}

// NewActorUseCase returns ActorUseCase.
func NewActorUseCase(domainFactory actors.Factory, actorDAO dao.ActorDAO) ActorUseCase {
	return &actorUseCase{
		actorDAO:      actorDAO,
		domainFactory: domainFactory,
	}
}

type actorUseCase struct {
	actorDAO      dao.ActorDAO
	domainFactory actors.Factory
}

func (u *actorUseCase) GetActor(ctx context.Context, actorID types.ActorID) (*entities.Actor, error) {
	getter := u.domainFactory.NewActorGetter()
	portImpl := &actorPortImpl{
		actorDAO: u.actorDAO,
	}
	return getter.GetActor(ctx, actorID, portImpl)
}

func (u *actorUseCase) GetActorByName(ctx context.Context, actorName string) (*entities.Actor, error) {
	getter := u.domainFactory.NewActorGetter()
	portImpl := &actorPortImpl{
		actorDAO: u.actorDAO,
	}
	return getter.GetActorByName(ctx, actorName, portImpl)
}

type actorPortImpl struct {
	actorDAO dao.ActorDAO
}

func (p *actorPortImpl) GetActor(ctx context.Context, actorID types.ActorID) (*entities.Actor, error) {
	return p.actorDAO.GetActor(ctx, actorID)
}

func (p *actorPortImpl) GetActorByName(ctx context.Context, actorName string) (*entities.Actor, error) {
	return p.actorDAO.GetActorByName(ctx, actorName)
}
