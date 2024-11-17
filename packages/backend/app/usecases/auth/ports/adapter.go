package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// AuthAdapter is the interface for authenticating user.
type AuthAdapter interface {
	// Authenticate authenticates user
	Authenticate(ctx context.Context, userID types.UserID, password string) (token string, err error)
}
