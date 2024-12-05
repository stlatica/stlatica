package auth

import (
	"context"

	domainauth "github.com/stlatica/stlatica/packages/backend/app/domains/auth"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/auth/ports"
)

// AuthenticationUseCase is the interface for authenticating user.
type AuthenticationUseCase interface {
	// Authenticate authenticates user.
	Authenticate(ctx context.Context, preferredUserID string, password string) (string, error)
}

// NewAuthenticationUseCase returns AuthenticationUseCase.
func NewAuthenticationUseCase(appLogger *logger.AppLogger,
	domainFactory domainauth.Factory, authAdapter ports.AuthAdapter, userDAO dao.UserDAO) AuthenticationUseCase {
	return &authenticationUseCase{
		appLogger:     appLogger,
		userDAO:       userDAO,
		domainFactory: domainFactory,
		authAdapter:   authAdapter,
	}
}

type authenticationUseCase struct {
	appLogger     *logger.AppLogger
	userDAO       dao.UserDAO
	domainFactory domainauth.Factory
	authAdapter   ports.AuthAdapter
}

func (u *authenticationUseCase) Authenticate(ctx context.Context,
	preferredUserID string, password string) (string, error) {
	authenticator := u.domainFactory.NewAuthenticator()
	portImpl := &authPortImpl{
		userDAO:     u.userDAO,
		authAdapter: u.authAdapter,
	}
	return authenticator.Authenticate(ctx, preferredUserID, password, portImpl, portImpl)
}

type authPortImpl struct {
	userDAO     dao.UserDAO
	authAdapter ports.AuthAdapter
}

func (p *authPortImpl) GetUserIDByPreferredUserID(ctx context.Context, preferredUserID string) (types.UserID, error) {
	user, err := p.userDAO.GetUserByPreferredUserID(ctx, preferredUserID)
	if err != nil {
		return types.UserID{}, err
	}
	return user.GetUserID(), nil
}

func (p *authPortImpl) AuthAndGenerateToken(ctx context.Context, userID types.UserID, password string) (string, error) {
	return p.authAdapter.Authenticate(ctx, userID, password)
}
