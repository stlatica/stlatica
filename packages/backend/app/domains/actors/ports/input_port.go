package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// ActorGetInPort is the interface for getting actor.
type ActorGetInPort interface {
	// GetActor returns actor.
	GetActor(ctx context.Context, actorID types.ActorID) (*entities.Actor, error)
	// GetActorByName returns actor by name.
	GetActorByName(ctx context.Context, actorName string) (*entities.Actor, error)
}
