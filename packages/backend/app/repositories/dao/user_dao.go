package dao

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/friendsofgo/errors"
	"github.com/go-sql-driver/mysql"
	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/users/ports"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	// DuplicateEntry in database
	DuplicateEntry = 1062
)

// UserDAO is the interface for getting user.
type UserDAO interface {
	// GetUser returns user.
	GetUser(ctx context.Context, userID types.UserID) (*domainentities.User, error)
	// CreateUser creates user.
	CreateUser(ctx context.Context, user domainentities.UserBase) (*domainentities.User, error)
	// GetUserByPreferredUserID returns user by preferred user ID.
	GetUserByPreferredUserID(ctx context.Context, preferredUserName string) (*domainentities.User, error)
	// GetUserByMailAddress returns user by mail address.
	GetUserByMailAddress(ctx context.Context, mailAddress string) (*domainentities.User, error)
	// GetFollows returns follows of user.
	GetFollows(ctx context.Context, params domainports.FollowsGetParams) ([]*domainentities.User, error)
	// GetFollowers returns followers of user.
	GetFollowers(ctx context.Context, params domainports.FollowersGetParams) ([]*domainentities.User, error)
	// FollowUser creates a state where one user is following another user.
	FollowUser(ctx context.Context, params domainentities.UserRelationBase) error
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.NewDomainError(err,
				domainerrors.DomainErrorTypeNotFound, "user not found")
		}
		return nil, err
	}

	return convertUserEntityToDomainEntity(entity), nil
}

func (dao *userDAO) CreateUser(ctx context.Context, user domainentities.UserBase) (
	*domainentities.User, error) {
	/*user := entities.UserBase{
		PreferredUserID:   preferredUserID,
		PreferredUserName: userName,
		RegisteredAt:      types.NewUnixTimeFromCurrentTime(),
		IsPublic:          true,
		MailAddress:       mailAddress,
	}*/
	ctx = boil.SkipTimestamps(ctx)
	// err := entities.UserBase
	repositoriesUser := entities.UserBase{
		UserID:            user.UserID,
		PreferredUserID:   user.PreferredUserID,
		PreferredUserName: user.PreferredUserName,
		RegisteredAt:      user.RegisteredAt,
		IsPublic:          user.IsPublic,
		MailAddress:       user.MailAddress,
		IconImageID:       user.IconImageID,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
	err := repositoriesUser.Insert(ctx, dao.ctxExecutor, boil.Infer())
	return &domainentities.User{
		UserBase: domainentities.UserBase{
			UserID:            user.UserID,
			PreferredUserID:   user.PreferredUserID,
			PreferredUserName: user.PreferredUserName,
			RegisteredAt:      user.RegisteredAt,
			IsPublic:          user.IsPublic,
			MailAddress:       user.MailAddress,
			IconImageID:       user.IconImageID,
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.NewDomainError(err,
				domainerrors.DomainErrorTypeNotFound, "user not found")
		}
		return nil, err
	}

	return convertUserEntityToDomainEntity(entity), nil
}

func (dao *userDAO) GetUserByMailAddress(ctx context.Context,
	mailAddress string) (*domainentities.User, error) {
	query := entities.UserBaseWhere.MailAddress.EQ(mailAddress)
	entity, err := entities.Users(query).One(ctx, dao.ctxExecutor)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.NewDomainError(err,
				domainerrors.DomainErrorTypeNotFound, "user not found")
		}
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
			qm.Limit(params.Limit),
			qm.OrderBy("follower_user_id DESC"),
		).All(ctx, dao.ctxExecutor)
	} else {
		userRelations, err = entities.UserRelations(
			entities.UserRelationBaseWhere.FollowUserID.EQ(params.UserID),
			qm.Limit(params.Limit),
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

func (dao *userDAO) GetFollowers(ctx context.Context,
	params domainports.FollowersGetParams) ([]*domainentities.User, error) {
	var userRelations entities.UserRelationBaseSlice
	var err error
	if reflect.ValueOf(params.UserPaginationID).IsZero() {
		userRelations, err = entities.UserRelations(
			entities.UserRelationBaseWhere.FollowerUserID.EQ(params.UserID),
			qm.Limit(params.Limit),
			qm.OrderBy("follow_user_id DESC"),
		).All(ctx, dao.ctxExecutor)
	} else {
		userRelations, err = entities.UserRelations(
			entities.UserRelationBaseWhere.FollowerUserID.EQ(params.UserID),
			entities.UserRelationBaseWhere.FollowUserID.GT(params.UserPaginationID),
			qm.Limit(params.Limit),
			qm.OrderBy("follow_user_id DESC"),
		).All(ctx, dao.ctxExecutor)
	}
	if err != nil {
		return nil, err
	}

	followerUserIDs := make([]interface{}, 0, len(userRelations))
	for _, relation := range userRelations {
		followerUserIDs = append(followerUserIDs, relation.FollowUserID)
	}
	userQuery := entities.Users(
		qm.WhereIn("user_id IN ?", followerUserIDs...),
	)
	users, err := userQuery.All(ctx, dao.ctxExecutor)
	if err != nil {
		return nil, err
	}

	followers := make([]*domainentities.User, 0, len(users))
	for _, user := range users {
		followers = append(followers, convertUserEntityToDomainEntity(user))
	}
	return followers, nil
}

func (dao *userDAO) FollowUser(ctx context.Context, params domainentities.UserRelationBase) error {
	ctx = boil.SkipTimestamps(ctx)
	repositoriesUserRelationBase := entities.UserRelationBase{
		FollowUserID:   params.FollowUserID,
		FollowerUserID: params.FollowerUserID,
		CreatedAt:      params.CreatedAt,
		UpdatedAt:      params.UpdatedAt,
	}
	err := repositoriesUserRelationBase.Insert(ctx, dao.ctxExecutor, boil.Infer())
	if err != nil {
		e := errors.Cause(err)
		if e != nil {
			var mySQLError *mysql.MySQLError
			if errors.As(e, &mySQLError) && mySQLError.Number == DuplicateEntry {
				return domainerrors.NewDomainError(err,
					domainerrors.DomainErrorTypeDuplicateEntry, "already following")
			}
		}
	}
	return err
}

func convertUserEntityToDomainEntity(entity *entities.UserBase) *domainentities.User {
	return &domainentities.User{
		UserBase: domainentities.UserBase{
			UserID:            entity.UserID,
			PreferredUserID:   entity.PreferredUserID,
			PreferredUserName: entity.PreferredUserName,
			IconImageID:       entity.IconImageID,
			RegisteredAt:      entity.RegisteredAt,
			IsPublic:          entity.IsPublic,
			MailAddress:       entity.MailAddress,
			CreatedAt:         entity.CreatedAt,
			UpdatedAt:         entity.UpdatedAt,
		},
	}
}
