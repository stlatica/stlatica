package auth

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/auth/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// Authenticator is an authenticator.
type Authenticator interface {
	// Authenticate authenticates user.
	Authenticate(ctx context.Context, preferredUserID string, password string,
		inPort ports.AuthenticateInPort, outPort ports.AuthenticateOutPort) (string, error)
}

type authenticator struct {
	appLogger *logger.AppLogger
}

func (a *authenticator) Authenticate(ctx context.Context, preferredUserID string, password string,
	inPort ports.AuthenticateInPort, outPort ports.AuthenticateOutPort) (string, error) {
	userID, err := inPort.GetUserIDByPreferredUserID(ctx, preferredUserID)
	if err != nil {
		return "", err
	}

	return outPort.AuthAndGenerateToken(ctx, userID, password)
}
