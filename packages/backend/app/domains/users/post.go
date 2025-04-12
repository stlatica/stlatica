package users

import (
	"context"

	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// UserCreator is the interface for creating user.
type UserCreator interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, userName string, preferredUserID string,
		mailAddress string, iconImageID types.ImageID, outPort ports.UserCreateOutPort) (*entities.User, error)
}

type userCreator struct {
	appLogger *logger.AppLogger
}

type userValidator struct {
	MailAddress string `validate:"required,email"`
}

func validateUser(user userValidator) error {
	validate := validator.New()
	return validate.Struct(user)
}

func (g *userCreator) CreateUser(ctx context.Context, userName string, preferredUserID string,
	mailAddress string, iconImageID types.ImageID, outPort ports.UserCreateOutPort) (*entities.User, error) {
	userValidate := userValidator{
		MailAddress: mailAddress,
	}
	err := validateUser(userValidate)
	if err != nil {
		return nil, domainerrors.NewDomainError(err,
			domainerrors.DomainErrorTypeInvalidData, "Invalid data format.")
	}
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
