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
}
