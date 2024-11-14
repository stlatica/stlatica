package dao

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// UserAuthCredentialDAO is the interface for getting user_auth_credential.
type UserAuthCredentialDAO interface {
	// GetUserAuthCredential returns user_auth_credential.
	GetUserAuthCredential(ctx context.Context, userID types.UserID) (*domainentities.UserAuthCredential, error)
	// CreateUserAuthCredential creates user_auth_credential.
	CreateUserAuthCredential(ctx context.Context,
		userID types.UserID, encryptedPassword string) (*domainentities.UserAuthCredential, error)
}

// NewUserAuthCredentialDAO returns UserAuthCredentialDAO.
func NewUserAuthCredentialDAO() UserAuthCredentialDAO {
	ctxExecutor := boil.GetContextDB()
	return &userAuthCredentialDAO{
		ctxExecutor: ctxExecutor,
	}
}

type userAuthCredentialDAO struct {
	ctxExecutor boil.ContextExecutor
}

func (dao *userAuthCredentialDAO) GetUserAuthCredential(ctx context.Context, userID types.UserID) (
	*domainentities.UserAuthCredential, error) {
	entity, err := entities.FindUserAuthCredentialBase(ctx, dao.ctxExecutor, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.NewDomainError(err,
				domainerrors.DomainErrorTypeNotFound, "user_auth_credential not found")
		}
		return nil, err
	}

	return &domainentities.UserAuthCredential{
		UserAuthCredentialBase: domainentities.UserAuthCredentialBase{
			UserID:            entity.UserID,
			EncryptedPassword: entity.EncryptedPassword,
			CreatedAt:         entity.CreatedAt,
			UpdatedAt:         entity.UpdatedAt,
		},
	}, nil
}

func (dao *userAuthCredentialDAO) CreateUserAuthCredential(ctx context.Context,
	userID types.UserID, encryptedPassword string) (*domainentities.UserAuthCredential, error) {
	entity := entities.UserAuthCredentialBase{
		UserID:            userID,
		EncryptedPassword: encryptedPassword,
		CreatedAt:         types.NewUnixTimeFromCurrentTime(),
		UpdatedAt:         types.NewUnixTimeFromCurrentTime(),
	}
	ctx = boil.SkipTimestamps(ctx)
	err := entity.Insert(ctx, dao.ctxExecutor, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &domainentities.UserAuthCredential{
		UserAuthCredentialBase: domainentities.UserAuthCredentialBase{
			UserID:            entity.UserID,
			EncryptedPassword: entity.EncryptedPassword,
			CreatedAt:         entity.CreatedAt,
			UpdatedAt:         entity.UpdatedAt,
		},
	}, nil
}
