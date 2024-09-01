package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// UserGetInPort is the interface for getting user.
type UserGetInPort interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*entities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error)
	// GetFollows returns follows of user.
	GetFollows(ctx context.Context, params FollowsGetParams) ([]*entities.User, error)
	// GetFollowers returns followers of user
	GetFollowers(ctx context.Context, params FollowersGetParams) ([]*entities.User, error)
}

// UserCreateInPort is the interface for creating user.
type UserCreateInPort interface {
	// CreateUSer creates a new user.
	CreateUser(ctx context.Context, userName string, mailAddress string) (*entities.User, error)
}
