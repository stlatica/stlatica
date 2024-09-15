package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// UserCreator is the interface for creating user.
type UserCreator interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, userName string, preferredUserID string,
		mailAddress string, iconImageID types.ImageID, inPort ports.UserCreateInPort) (*entities.User, error)
}

type userCreator struct {
	appLogger *logger.AppLogger
}

func (g *userCreator) CreateUser(ctx context.Context, userName string, preferredUserID string,
	mailAddress string, iconImageID types.ImageID, inPort ports.UserCreateInPort) (*entities.User, error) {
	// TODO: mail addressのvalidationを実装する https://github.com/stlatica/stlatica/issues/604
	userID := types.NewUserID()
	user := entities.UserBase{
		UserID:            userID,
		PreferredUserName: userName,
		PreferredUserID:   preferredUserID,
		MailAddress:       mailAddress,
		IconImageID:       iconImageID,
	}
	return inPort.CreateUser(ctx, user)
}
