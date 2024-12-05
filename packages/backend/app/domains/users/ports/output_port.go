package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
)

// UserCreateOutPort is the interface for creating user.
type UserCreateOutPort interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user entities.UserBase) (*entities.User, error)
	FollowUser(ctx context.Context, followUser entities.UserRelationBase) error
}
