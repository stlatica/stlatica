package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
)

// UserCreateInPort is the interface for creating user.
type UserCreateInPort interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user entities.UserBase) (*entities.User, error)
}
