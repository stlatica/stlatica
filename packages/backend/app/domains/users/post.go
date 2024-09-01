package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// UserCreator is the interface for creating user.
type UserCreator interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, userName string, mailAddress string, inPort ports.UserCreateInPort) (
		*entities.User, error)
}

type userCreator struct {
	appLogger *logger.AppLogger
}

func (g *userCreator) CreateUser(ctx context.Context, userName string,
	mailAddress string, inPort ports.UserCreateInPort) (*entities.User, error) {
	return inPort.CreateUser(ctx, userName, mailAddress)
}
