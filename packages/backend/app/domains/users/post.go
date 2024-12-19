package users

import (
	"context"
	"errors"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// UserCreator is the interface for creating user.
type UserCreator interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, userName string, preferredUserID string,
		mailAddress string, iconImageID types.ImageID, outPort ports.UserCreateOutPort) (*entities.User, error)
	FollowUser(ctx context.Context, userID types.UserID, followUserID types.UserID, outPort ports.UserCreateOutPort) error
}

type userCreator struct {
	appLogger *logger.AppLogger
}

func (g *userCreator) CreateUser(ctx context.Context, userName string, preferredUserID string,
	mailAddress string, iconImageID types.ImageID, outPort ports.UserCreateOutPort) (*entities.User, error) {
	// TODO: mail addressのvalidationを実装する https://github.com/stlatica/stlatica/issues/604
	userID := types.NewUserID()
	user := entities.UserBase{
		UserID:            userID,
		PreferredUserName: userName,
		PreferredUserID:   preferredUserID,
		MailAddress:       mailAddress,
		IconImageID:       iconImageID,
	}
	return outPort.CreateUser(ctx, user)
}

func (g *userCreator) FollowUser(ctx context.Context, userID types.UserID,
	followUserID types.UserID, outPort ports.UserCreateOutPort) error {
	followUser := entities.UserRelationBase{
		FollowerUserID: userID,
		FollowUserID:   followUserID,
	}
	err := outPort.FollowUser(ctx, followUser)
	var domainError domainerrors.DomainError
	if errors.As(err, &domainError) &&
		domainError.DomainErrorType() == domainerrors.DomainErrorTypeDuplicateEntry {
		return nil
	}
	return err
}
