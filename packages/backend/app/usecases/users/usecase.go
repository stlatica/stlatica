package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users"
	"github.com/stlatica/stlatica/packages/backend/app/logger"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
)

// UserUseCase is the interface for getting user.
type UserUseCase interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*entities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error)
}

// NewUserUseCase returns UserUseCase.
func NewUserUseCase(appLogger *logger.AppLogger, domainFactory users.Factory, userDAO dao.UserDAO) UserUseCase {
	return &userUseCase{
		appLogger:     appLogger,
		userDAO:       userDAO,
		domainFactory: domainFactory,
	}
}

type userUseCase struct {
	appLogger     *logger.AppLogger
	userDAO       dao.UserDAO
	domainFactory users.Factory
}

func (u *userUseCase) GetUser(ctx context.Context, userID types.UserID) (*entities.User, error) {
	getter := u.domainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	return getter.GetUser(ctx, userID, portImpl)
}

func (u *userUseCase) GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error) {
	getter := u.domainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	return getter.GetUserByPreferredUserID(ctx, preferredUserID, portImpl)
}

type userPortImpl struct {
	userDAO dao.UserDAO
}

func (p *userPortImpl) GetUser(ctx context.Context, userID types.UserID) (*entities.User, error) {
	return p.userDAO.GetUser(ctx, userID)
}

func (p *userPortImpl) GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error) {
	return p.userDAO.GetUserByPreferredUserID(ctx, preferredUserID)
}
