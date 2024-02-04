package plats

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// PlatPoster is the interface for posting plat.
type PlatPoster interface {
	// PostPlat returns new plat.
	PostPlat(ctx context.Context, actorID types.ActorID, content string,
		inPort ports.PlatPostInPort) (*entities.Plat, error)
}

type platPoster struct{}

func (g *platPoster) PostPlat(ctx context.Context, actorID types.ActorID,
	content string, inPort ports.PlatPostInPort) (*entities.Plat, error) {
	return inPort.PostPlat(ctx, actorID, content)
}
