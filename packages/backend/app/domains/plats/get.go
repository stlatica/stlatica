package plats

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/entities"
	"github.com/stlatica/stlatica/packages/backend/app/domains/plats/ports"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// PlatGetter is the interface for getting plat.
type PlatGetter interface {
	// GetPlat returns plat.
	GetPlat(ctx context.Context, platID types.PlatID, inPort ports.PlatGetInPort) (*entities.Plat, error)
	// GetPlats returns plats.
	GetPlats(ctx context.Context, getParams ports.PlatsGetParams, inPort ports.PlatGetInPort) ([]*entities.Plat, error)
}

type platGetter struct {
	appLogger *logger.AppLogger
}

func (g *platGetter) GetPlat(ctx context.Context,
	platID types.PlatID, inPort ports.PlatGetInPort) (*entities.Plat, error) {
	// TODO: 閲覧可否の判定ロジックを実装する https://github.com/stlatica/stlatica/issues/209
	return inPort.GetPlat(ctx, platID)
}

func (g *platGetter) GetPlats(ctx context.Context,
	getParams ports.PlatsGetParams, inPort ports.PlatGetInPort) ([]*entities.Plat, error) {
	if getParams.ToDate == 0 && getParams.FromDate == 0 {
		return nil, errors.New("must specify either to_date or from_date")
	}
	if getParams.ToDate != 0 && getParams.FromDate != 0 {
		return nil, errors.New("cannot specify both to_date and from_date")
	}
	var limit uint64
	if getParams.Limit == 0 {
		limit = 100
	}
	params := ports.PlatsGetParams{
		UserIDs:  getParams.UserIDs,
		FromDate: getParams.FromDate,
		ToDate:   getParams.ToDate,
		Limit:    limit,
	}
	return inPort.GetPlats(ctx, params)
}
