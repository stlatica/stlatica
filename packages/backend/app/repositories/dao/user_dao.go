package dao

import (
	"context"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// UserDAO is the interface for getting user.
type UserDAO interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*domainentities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserName string) (*domainentities.User, error)
}

// NewUserDAO returns UserDAO.
func NewUserDAO() UserDAO {
	ctxExecutor := boil.GetContextDB()
	return &userDAO{
		ctxExecutor: ctxExecutor,
	}
}

type userDAO struct {
	ctxExecutor boil.ContextExecutor
}

func (dao *userDAO) GetUser(ctx context.Context, userID types.UserID) (*domainentities.User, error) {
	entity, err := entities.FindUserBase(ctx, dao.ctxExecutor, userID)
	if err != nil {
		return nil, err
	}

	return convertUserEntityToDomainEntity(entity), nil
}

func (dao *userDAO) GetUserByPreferredUserID(ctx context.Context,
	preferredUserID string) (*domainentities.User, error) {
	query := entities.UserBaseWhere.PreferredUserID.EQ(preferredUserID)
	entity, err := entities.Users(query).One(ctx, dao.ctxExecutor)
	if err != nil {
		return nil, err
	}

	return convertUserEntityToDomainEntity(entity), nil
}

func convertUserEntityToDomainEntity(entity *entities.UserBase) *domainentities.User {
	return &domainentities.User{
		UserBase: domainentities.UserBase{
			UserID:            entity.UserID,
			PreferredUserID:   entity.PreferredUserID,
			PreferredUserName: entity.PreferredUserName,
			RegisteredAt:      entity.RegisteredAt,
			IsPublic:          entity.IsPublic,
			MailAddress:       entity.MailAddress,
			CreatedAt:         entity.CreatedAt,
			UpdatedAt:         entity.UpdatedAt,
		},
	}
}
