package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/dao"
	"github.com/stlatica/stlatica/packages/backend/app/usecases/users/ports"
)

// UserUseCase is the interface for getting user.
type UserUseCase interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*entities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error)
	// GetFollows returns follows of user.
	GetFollows(ctx context.Context, params ports.FollowsGetParams) ([]*entities.User, error)
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

func (u *userUseCase) GetFollows(ctx context.Context, params ports.FollowsGetParams) ([]*entities.User, error) {
	getter := u.domainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	user, err := getter.GetUserByPreferredUserID(ctx, params.PreferredUserID, portImpl)
	if err != nil {
		return nil, err
	}
	var paginationUser *entities.User
	if params.PreferredUserPaginationID != "" {
		paginationUser, err = getter.GetUserByPreferredUserID(ctx, params.PreferredUserPaginationID, portImpl)
		if err != nil {
			return nil, err
		}
	} else {
		paginationUser = &entities.User{}
	}

	domainParams := domainports.FollowsGetParams{
		UserID:           user.GetUserID(),
		UserPaginationID: paginationUser.GetUserID(),
		Limit:            params.Limit,
	}
	return getter.GetFollows(ctx, domainParams, portImpl)
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

func (p *userPortImpl) GetFollows(ctx context.Context,
	getParams domainports.FollowsGetParams) ([]*entities.User, error) {
	return p.userDAO.GetFollows(ctx, getParams)
}
