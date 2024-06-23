package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// UserGetter is the interface for getting user.
type UserGetter interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID, inPort ports.UserGetInPort) (*entities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context,
		preferredUserID string, inPort ports.UserGetInPort) (*entities.User, error)
}

type userGetter struct {
	appLogger *logger.AppLogger
}

func (g *userGetter) GetUser(ctx context.Context,
	userID types.UserID, inPort ports.UserGetInPort) (*entities.User, error) {
	return inPort.GetUser(ctx, userID)
}

func (g *userGetter) GetUserByPreferredUserID(ctx context.Context,
	preferredUserID string, inPort ports.UserGetInPort) (*entities.User, error) {
	return inPort.GetUserByPreferredUserID(ctx, preferredUserID)
}
