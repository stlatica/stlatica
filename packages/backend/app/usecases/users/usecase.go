package users

import (
	"context"
	"io"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/images"
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
	// GetFollowers returns followers of user.
	GetFollowers(ctx context.Context, params ports.FollowersGetParams) ([]*entities.User, error)
	// GetUserIcon returns user icon.
	GetUserIcon(ctx context.Context, preferredUserID string) (io.ReadCloser, error)
}

// NewUserUseCase returns UserUseCase.
func NewUserUseCase(appLogger *logger.AppLogger,
	userDomainFactory users.Factory, imageDomainFactory images.Factory,
	userDAO dao.UserDAO, imageAdapter ports.ImageAdapter) UserUseCase {
	return &userUseCase{
		appLogger:          appLogger,
		userDAO:            userDAO,
		imageAdapter:       imageAdapter,
		userDomainFactory:  userDomainFactory,
		imageDomainFactory: imageDomainFactory,
	}
}

type userUseCase struct {
	appLogger          *logger.AppLogger
	userDAO            dao.UserDAO
	imageAdapter       ports.ImageAdapter
	userDomainFactory  users.Factory
	imageDomainFactory images.Factory
}

func (u *userUseCase) GetUser(ctx context.Context, userID types.UserID) (*entities.User, error) {
	getter := u.userDomainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	return getter.GetUser(ctx, userID, portImpl)
}

func (u *userUseCase) GetUserByPreferredUserID(ctx context.Context, preferredUserID string) (*entities.User, error) {
	getter := u.userDomainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	return getter.GetUserByPreferredUserID(ctx, preferredUserID, portImpl)
}

func (u *userUseCase) GetFollows(ctx context.Context, params ports.FollowsGetParams) ([]*entities.User, error) {
	getter := u.userDomainFactory.NewUserGetter()
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

func (u *userUseCase) GetFollowers(ctx context.Context, params ports.FollowersGetParams) ([]*entities.User, error) {
	getter := u.userDomainFactory.NewUserGetter()
	portImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	user, err := getter.GetUserByPreferredUserID(ctx, params.PreferredUserID, portImpl)
	if err != nil {
		return nil, err
	}
	var paginationUsers *entities.User
	if params.PreferredUserPaginationID != "" {
		paginationUsers, err = getter.GetUserByPreferredUserID(ctx, params.PreferredUserPaginationID, portImpl)
		if err != nil {
			return nil, err
		}
	} else {
		paginationUsers = &entities.User{}
	}
	domainParams := domainports.FollowersGetParams{
		UserID:           user.GetUserID(),
		UserPaginationID: paginationUsers.GetUserID(),
		Limit:            params.Limit,
	}
	return getter.GetFollowers(ctx, domainParams, portImpl)
}

func (u *userUseCase) GetUserIcon(ctx context.Context, preferredUserID string) (io.ReadCloser, error) {
	userGetter := u.userDomainFactory.NewUserGetter()
	userPortImpl := &userPortImpl{
		userDAO: u.userDAO,
	}
	user, err := userGetter.GetUserByPreferredUserID(ctx, preferredUserID, userPortImpl)
	if err != nil {
		return nil, err
	}

	imageGetter := u.imageDomainFactory.NewImageGetter()
	imagePortImpl := &imagePortImpl{
		imageAdapter: u.imageAdapter,
	}
	return imageGetter.GetImage(ctx, user.GetIconImageID(), imagePortImpl)
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

func (p *userPortImpl) GetFollowers(ctx context.Context,
	getParams domainports.FollowersGetParams) ([]*entities.User, error) {
	return p.userDAO.GetFollowers(ctx, getParams)
}

type imagePortImpl struct {
	imageAdapter ports.ImageAdapter
}

func (i *imagePortImpl) GetImage(ctx context.Context, imageID types.ImageID) (io.ReadCloser, error) {
	return i.imageAdapter.GetImage(ctx, imageID)
}
