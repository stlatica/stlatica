package actors

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/actors/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/logger"
)

// ActorGetter is the interface for getting actor.
type ActorGetter interface {
	// GetActor returns actor.
	GetActor(ctx context.Context, actorID types.ActorID, inPort ports.ActorGetInPort) (*entities.Actor, error)
	// GetActorByName returns actor by name.
	GetActorByName(ctx context.Context, actorName string, inPort ports.ActorGetInPort) (*entities.Actor, error)
}

type actorGetter struct {
	appLogger *logger.AppLogger
}

func (g *actorGetter) GetActor(ctx context.Context,
	actorID types.ActorID, inPort ports.ActorGetInPort) (*entities.Actor, error) {
	return inPort.GetActor(ctx, actorID)
}

func (g *actorGetter) GetActorByName(ctx context.Context,
	actorName string, inPort ports.ActorGetInPort) (*entities.Actor, error) {
	return inPort.GetActorByName(ctx, actorName)
}
