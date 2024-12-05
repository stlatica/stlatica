package ports

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
)

// AuthenticateOutPort is the interface for authenticating user.
type AuthenticateOutPort interface {
	// AuthAndGenerateToken authenticates user and generates token.
	AuthAndGenerateToken(ctx context.Context, userID types.UserID, password string) (string, error)
}
