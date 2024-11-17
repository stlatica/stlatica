package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// AuthenticateInPort is the interface for authenticating user.
type AuthenticateInPort interface {
	// GetUserIDByPreferredUserID returns user ID by preferred user ID.
	GetUserIDByPreferredUserID(ctx context.Context, preferredUserID string) (types.UserID, error)
}
