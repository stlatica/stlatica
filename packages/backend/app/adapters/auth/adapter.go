package auth

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/auth"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/kvs"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/auth/ports"
)

// authAdapter implements the AuthAdapter interface.
type authAdapter[T auth.SignType] struct {
	client kvs.Client
	config *Config[T]
}

// NewAdapter returns an AuthAdapter.
func NewAdapter[T auth.SignType](client kvs.Client, config *Config[T]) ports.AuthAdapter {
	return &authAdapter[T]{
		client: client,
		config: config,
	}
}

// Authenticate authenticates a user using their userID and password.
func (a *authAdapter[T]) Authenticate(ctx context.Context, userID types.UserID, password string) (string, error) {
	userCredential := auth.NewPWAuth(userID.String(), password, a.config.SignatureAlgorithm, a.config.SignaturePrivateKey)
	accessToken, err := userCredential.Auth(ctx, a.client)
	if err != nil {
		return "", err
	}
	return string(accessToken.Export()), nil
}
