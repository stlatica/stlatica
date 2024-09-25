package dao

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	domainentities "github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	domainerrors "github.com/stlatica/stlatica/packages/backend/app/domains/errors"
	domainports "github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/repositories/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// PlatDAO is the interface for getting plat.
type PlatDAO interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID) (*domainentities.Plat, error)
	// CreatePlat creates plat.
	CreatePlat(ctx context.Context, userID types.UserID, content string) (*domainentities.Plat, error)
	// GetPlats returns plats.
	GetPlats(ctx context.Context, params domainports.PlatsGetParams) ([]*domainentities.Plat, error)
}

// NewPlatDAO returns PlatDAO.
func NewPlatDAO() PlatDAO {
	ctxExecutor := boil.GetContextDB()
	return &platDAO{
		ctxExecutor: ctxExecutor,
	}
}

type platDAO struct {
	ctxExecutor boil.ContextExecutor
}

func (dao *platDAO) GetPlat(ctx context.Context, platID types.PlatID) (*domainentities.Plat, error) {
	entity, err := entities.FindPlatBase(ctx, dao.ctxExecutor, platID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.NewDomainError(err,
				domainerrors.DomainErrorTypeNotFound, "user not found")
		}
		return nil, err
	}

	return &domainentities.Plat{
		PlatBase: domainentities.PlatBase{
			PlatID:    entity.PlatID,
			UserID:    entity.UserID,
			Content:   entity.Content,
			CreatedAt: entity.CreatedAt,
		},
	}, nil
}

func (dao *platDAO) CreatePlat(ctx context.Context, userID types.UserID, content string) (
	*domainentities.Plat, error) {
	plat := entities.PlatBase{
		PlatID:    types.NewPlatID(),
		UserID:    userID,
		Content:   content,
		CreatedAt: types.NewUnixTimeFromCurrentTime(),
	}
	ctx = boil.SkipTimestamps(ctx)
	err := plat.Insert(ctx, dao.ctxExecutor, boil.Infer())
	return &domainentities.Plat{
		PlatBase: domainentities.PlatBase{
			PlatID:    plat.PlatID,
			UserID:    plat.UserID,
			Content:   plat.Content,
			CreatedAt: plat.CreatedAt,
		},
	}, err
}

func (dao *platDAO) GetPlats(ctx context.Context, params domainports.PlatsGetParams) (
	[]*domainentities.Plat, error) {
	var dateCondition qm.QueryMod
	if params.FromDate != 0 {
		dateCondition = entities.PlatBaseWhere.CreatedAt.GTE(params.FromDate)
	} else {
		dateCondition = entities.PlatBaseWhere.CreatedAt.LTE(params.ToDate)
	}

	userIDs := make([]interface{}, 0, len(params.UserIDs))
	for _, userID := range params.UserIDs {
		userIDs = append(userIDs, userID)
	}
	query := entities.Plats(
		dateCondition,
		qm.WhereIn("user_id IN ?", userIDs...),
		qm.Limit(params.Limit),
		qm.OrderBy("created_at DESC"),
	)

	records, err := query.All(ctx, dao.ctxExecutor)
	if err != nil {
		return nil, err
	}

	plats := make([]*domainentities.Plat, 0, len(records))
	for _, record := range records {
		plats = append(plats, &domainentities.Plat{
			PlatBase: domainentities.PlatBase{
				PlatID:    record.PlatID,
				UserID:    record.UserID,
				Content:   record.Content,
				CreatedAt: record.CreatedAt,
			},
		})
	}
	return plats, nil
}
