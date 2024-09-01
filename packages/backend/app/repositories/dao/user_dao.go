package dao

import (
	"context"
	"reflect"

	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserDAO is the interface for getting user.
type UserDAO interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*domainentities.User, error)
	// CreateUser creates user.
	CreateUser(ctx context.Context, userName string, mailAddress string) (*domainentities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserName string) (*domainentities.User, error)
	// GetFollows returns follows of user.
	GetFollows(ctx context.Context, params domainports.FollowsGetParams) ([]*domainentities.User, error)
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

func (dao *userDAO) CreateUser(ctx context.Context, userName string, mailAddress string) (
	*domainentities.User, error) {
	newUserID := types.NewUserID()
	user := entities.UserBase{
		UserID:            newUserID,
		PreferredUserID:   newUserID.String(),
		PreferredUserName: userName,
		RegisteredAt:      types.NewUnixTimeFromCurrentTime(),
		IsPublic:          true,
		MailAddress:       mailAddress,
	}
	ctx = boil.SkipTimestamps(ctx)
	err := user.Insert(ctx, dao.ctxExecutor, boil.Infer())
	return &domainentities.User{
		UserBase: domainentities.UserBase{
			UserID:            user.UserID,
			PreferredUserID:   user.PreferredUserID,
			PreferredUserName: user.PreferredUserName,
			RegisteredAt:      user.RegisteredAt,
			IsPublic:          user.IsPublic,
			MailAddress:       user.MailAddress,
			CreatedAt:         user.RegisteredAt,
			UpdatedAt:         user.RegisteredAt,
		},
	}, err
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

func (dao *userDAO) GetFollows(ctx context.Context,
	params domainports.FollowsGetParams) ([]*domainentities.User, error) {
	var userRelations entities.UserRelationBaseSlice
	var err error
	if !reflect.ValueOf(params.UserPaginationID).IsZero() {
		userRelations, err = entities.UserRelations(
			entities.UserRelationBaseWhere.FollowUserID.EQ(params.UserID),
			entities.UserRelationBaseWhere.FollowerUserID.GT(params.UserPaginationID),
			qm.Limit(int(params.Limit)),
			qm.OrderBy("follower_user_id DESC"),
		).All(ctx, dao.ctxExecutor)
	} else {
		userRelations, err = entities.UserRelations(
			entities.UserRelationBaseWhere.FollowUserID.EQ(params.UserID),
			qm.Limit(int(params.Limit)),
			qm.OrderBy("follower_user_id DESC"),
		).All(ctx, dao.ctxExecutor)
	}
	if err != nil {
		return nil, err
	}

	followUserIDs := make([]interface{}, 0, len(userRelations))
	for _, relation := range userRelations {
		followUserIDs = append(followUserIDs, relation.FollowerUserID)
	}
	userQuery := entities.Users(
		qm.WhereIn("user_id IN ?", followUserIDs...),
	)
	users, err := userQuery.All(ctx, dao.ctxExecutor)
	if err != nil {
		return nil, err
	}

	follows := make([]*domainentities.User, 0, len(users))
	for _, user := range users {
		follows = append(follows, convertUserEntityToDomainEntity(user))
	}

	return follows, nil
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
